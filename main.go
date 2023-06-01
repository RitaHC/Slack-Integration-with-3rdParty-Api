package main

import (
	// Slack Dependency
    "github.com/slack-go/slack"
	"fmt"
)

// Create Slack client with OAuth Token
func createSlackClient() *slack.Client {
    token := "xoxb-5377582754640-5356335573060-4iavJvIiL2F6pECHUP9SHgAY"
    return slack.New(token)
}

// Create a Slack client that can interact with Slack API

func sendMessageToChannel(client *slack.Client, channelID, message string) error {
    _, _, err := client.PostMessage(channelID, slack.MsgOptionText(message, false))
    return err
}


func main() {
    client := createSlackClient()
    channelID := "C05AATL3E9K"
    message := "Hello from Golang!"
    err := sendMessageToChannel(client, channelID, message)
    if err != nil {
        fmt.Printf("Failed to send message: %v\n", err)
    }
}
