# Trakteer Webhook Forwarder

This project is a webhook forwarder for the Trakteer platform to Discord. It addresses the limitation of Trakteer, which does not support multiple Discord webhooks.

## Features

- Forwards webhooks from Trakteer to Discord channels.
- Easy to configure and deploy.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/datenshicommunity/trakteer-webhook-forwarder.git
    ```
2. Install dependencies:
    ```sh
    cd trakteer-webhook-forwarder
    go build
    ```

## Configuration

1. Create a `.env` file in the root directory and add your configuration:
    ```env
    DISCORD_WEBHOOK_URL=https://discord.com/api/webhooks/your_webhook_url_1
    VALID_TOKEN=TOKEN_FROM_TRAKTEER
    ```

## Usage

1. Start the server:
    ```sh
    ./main
    ```

2. Add this backend URL into your Trakteer Platform in Integrations -> Webhook

3. The server will listen for incoming webhooks from Trakteer and forward them to the configured Discord webhooks.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License.
