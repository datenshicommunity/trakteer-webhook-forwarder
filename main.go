package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "bytes"
    "os"
)

// Payload structure to match the incoming JSON
type Payload struct {
    CreatedAt        string  `json:"created_at"`
    TransactionID    string  `json:"transaction_id"`
    Type             string  `json:"type"`
    SupporterName    string  `json:"supporter_name"`
    SupporterAvatar  string  `json:"supporter_avatar"`
    SupporterMessage string  `json:"supporter_message"`
    Media            *string `json:"media"`
    Unit             string  `json:"unit"`
    UnitIcon         string  `json:"unit_icon"`
    Quantity         int     `json:"quantity"`
    Price            int     `json:"price"`
    NetAmount        int     `json:"net_amount"`
}

// DiscordWebhook structure for the output JSON to Discord
type DiscordWebhook struct {
    Content    string    `json:"content"`
    Embeds     []Embed   `json:"embeds"`
    Attachments []string `json:"attachments"`
}

type Embed struct {
    Description string  `json:"description"`
    Color       int     `json:"color"`
    Author      Author  `json:"author"`
    Footer      Footer  `json:"footer"`
    Thumbnail   Thumbnail `json:"thumbnail"`
}

type Author struct {
    Name    string `json:"name"`
    IconURL string `json:"icon_url"`
}

type Footer struct {
    Text string `json:"text"`
}

type Thumbnail struct {
    URL string `json:"url"`
}

var discordWebhookURL string
var validToken string

func init() {
    discordWebhookURL = os.Getenv("DISCORD_WEBHOOK_URL")
    validToken = os.Getenv("VALID_TOKEN")
}

// Function to handle incoming webhook requests
func webhookHandler(w http.ResponseWriter, r *http.Request) {
    // Extract token from Authorization header
    token := r.Header.Get("Authorization")
    if token == "" {
        http.Error(w, "Missing Authorization token", http.StatusUnauthorized)
        return
    }

    // Validate the token
    if token != validToken {
        http.Error(w, "Invalid Authorization token", http.StatusUnauthorized)
        return
    }

    // Read the body of the request
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Unable to read request body", http.StatusBadRequest)
        return
    }

    // Parse the incoming payload
    var payload Payload
    if err := json.Unmarshal(body, &payload); err != nil {
        http.Error(w, "Invalid JSON format", http.StatusBadRequest)
        return
    }

    // Create the Discord webhook message
    discordMessage := DiscordWebhook{
        Embeds: []Embed{
            {
                Description: fmt.Sprintf("**%s** mentraktir **%d %s** senilai **Rp %d** dengan pesan _**\"%s\"**_",
                    payload.SupporterName, payload.Quantity, payload.Unit, payload.Price, payload.SupporterMessage),
                Color: 261932, // Custom color for the embed
                Author: Author{
                    Name:    payload.SupporterName,
                    IconURL: payload.SupporterAvatar,
                },
                Footer: Footer{
                    Text: "Dukung Datenshi Community di https://trakteer.id/datenshi",
                },
                Thumbnail: Thumbnail{
                    URL: payload.UnitIcon,
                },
            },
        },
        Attachments: []string{},
    }

    // Convert the message to JSON
    discordPayload, err := json.Marshal(discordMessage)
    if err != nil {
        http.Error(w, "Error creating Discord message", http.StatusInternalServerError)
        return
    }

    // Send the message to the Discord webhook
    resp, err := http.Post(discordWebhookURL, "application/json", bytes.NewBuffer(discordPayload))
    if err != nil || resp.StatusCode != http.StatusOK {
        http.Error(w, "Failed to send message to Discord", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Payload sent to Discord successfully"))
}

func main() {
    http.HandleFunc("/webhook", webhookHandler)
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}