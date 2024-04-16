package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	// Update these values with your contract address and client URL
	contractAddress = common.HexToAddress("0x55d398326f99059fF775485246999027B3197955")
	clientURL       = "http://localhost:8545" // Update with your client URL
)

// Function to sign transactions using the sender's private key
func signTransaction(privateKey string, transactionData []byte) ([]byte, error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	// Sign the transaction
	signedTx, err := crypto.Sign(transactionData, key)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

// Function to send USDT to multiple addresses
func sendUSDTToMultipleAddresses(privateKey string, recipients []string, amount float64) error {
	// Load USDT contract ABI
	abiJSON, err := ioutil.ReadFile("usdt_abi.json")
	if err != nil {
		return err
	}

	// Parse USDT contract ABI
	contractABI, err := abi.JSON(strings.NewReader(string(abiJSON)))
	if err != nil {
		return err
	}

	// Connect to the Ethereum client
	client, err := ethclient.Dial(clientURL)
	if err != nil {
		return err
	}

	// Get the function selector for transfer function
	transferFnSelector := contractABI.Methods["transfer"].ID

	// Get sender's address from private key
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return err
	}
	publicKey := crypto.PubkeyToAddress(key.PublicKey)

	// Iterate over recipients
	for _, recipient := range recipients {
		// Convert recipient address to common.Address
		recipientAddr := common.HexToAddress(recipient)

		// Prepare the transfer data
		transferData, err := contractABI.Pack("transfer", recipientAddr, 10)
		if err != nil {
			return err
		}

		// Construct the transaction data
		nonce, err := client.PendingNonceAt(context.Background(), publicKey)
		if err != nil {
			return err
		}
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			return err
		}
		gasLimit := uint64(200000)
		txData := ethereum.CallMsg{
			From:     publicKey,
			To:       &contractAddress,
			Value:    nil,
			Gas:      gasLimit,
			GasPrice: gasPrice,
			Data:     transferData,
		}

		// Sign the transaction
		signedTx, err := signTransaction(privateKey, txData)
		if err != nil {
			return err
		}

		// Send the signed transaction to the network
		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			return err
		}

		fmt.Printf("Successfully sent %f USDT to %s\n", amount, recipient)
	}

	return nil
}

func main() {
	// Load private key from file or environment variable
	privateKey := "0xfbac2a544fdc94a38f5050b51cb7fa536d5d8708c1d0e975f2443ce461759c2f"

	// Load recipient addresses from a text file
	file, err := os.Open("recipients.txt")
	if err != nil {
		log.Fatal("Error opening recipients file:", err)
	}
	defer file.Close()

	var recipients []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		recipients = append(recipients, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error scanning recipients file:", err)
	}

	// Amount of USDT to send to each recipient (in decimal)
	amount := 10.0 // Adjust the amount as needed

	// Send USDT to multiple addresses
	err = sendUSDTToMultipleAddresses(privateKey, recipients, amount)
	if err != nil {
		log.Fatal("Error sending USDT:", err)
	}
}
