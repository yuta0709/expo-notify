package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	sdk "github.com/oliveroneill/exponent-server-sdk-golang/sdk"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Expo Push Token: ")
	token, _ := reader.ReadString('\n')
	token = token[:len(token)-1]

	fmt.Print("Enter Title: ")
	title, _ := reader.ReadString('\n')
	title = title[:len(title)-1]

	fmt.Print("Enter Body: ")
	body, _ := reader.ReadString('\n')
	body = body[:len(body)-1]

	client := sdk.NewPushClient(nil)
	t, err := sdk.NewExponentPushToken(token)
	if err != nil {
		log.Fatalln(err)
	}

	message := &sdk.PushMessage{
		To:    []sdk.ExponentPushToken{t},
		Body:  body,
		Title: title,
	}

	response, err := client.Publish(message)
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	fmt.Println("Message sent successfully:", response)
}
