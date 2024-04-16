

```markdown
# Bulk USDT Transfer

This Go program allows you to send USDT (Tether) tokens to multiple Ethereum addresses in bulk.

## Prerequisites

Before running this program, ensure you have the following:

- Ethereum private key: You need the private key of an Ethereum account that holds USDT tokens.
- Recipient addresses: Prepare a list of Ethereum addresses to which you want to send USDT tokens. Store these addresses in a text file named `recipients.txt`.
- Ethereum client URL: Update the `clientURL` variable in the code with the URL of your Ethereum client (e.g., Infura, Alchemy, or a local node).

## Installation

1. Clone this repository to your local machine.

2. Install the necessary dependencies by running:

   ```bash
   go mod tidy
   ```

3. Ensure you have a properly formatted ABI file for the USDT contract. Save the ABI JSON in a file named `usdt_abi.json` in the project directory.

## Usage

1. Open `recipients.txt` and add the Ethereum addresses of the recipients, with each address on a new line.

2. Update the `privateKey` variable in `main.go` with your Ethereum private key.

3. Adjust the `amount` variable in `main.go` to specify the amount of USDT tokens to send to each recipient.

4. Run the program using:

   ```bash
   go run main.go
   ```

5. The program will read the recipient addresses from `recipients.txt`, send the specified amount of USDT tokens to each address, and print a success message for each transaction.

## Disclaimer

- Use this program responsibly. Ensure that you have permission to send tokens from the specified Ethereum account.
- Be cautious when handling private keys and sensitive information. Keep your private keys secure and never share them with anyone.

## Author

Ashikur Rahman

For further interaction, my social media handles are as follows:
- Twitter: [Alex_Ashu_07](https://twitter.com/Alex_Ashu_07)
- LinkedIn: [ashik2mk](https://www.linkedin.com/in/ashik2mk/)
- GitHub: [Ashikurrahaman287](https://github.com/Ashikurrahaman287)
  
For prompt communication:
- WhatsApp: [+1 (409) 965-5574](https://wa.me/14099655574)
- Skype: [Join Skype Meeting](https://join.skype.com/invite/sBFuPnQCPN8O)

To explore my portfolio, please visit: [www.ashikurrahaman.com](https://www.ashikurrahaman.com)
```

Feel free to adjust the formatting or content as needed!
