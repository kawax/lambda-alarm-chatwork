package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	chatwork "github.com/griffin-stewie/go-chatwork"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Message struct
type Message struct {
	NewStateValue    string `json:"NewStateValue"`
	NewStateReason   string `json:"NewStateReason"`
	AlarmName        string `json:"AlarmName"`
	AlarmDescription string `json:"AlarmDescription"`
}

// Handler ...
func Handler(ctx context.Context, snsEvent events.SNSEvent) (string, error) {
	snsRecord := snsEvent.Records[0].SNS

	if snsRecord.Message == "" {
		return "error", errors.New("SNSRecord is empty")
	}

	message := new(Message)
	jsonBytes := ([]byte)(snsRecord.Message)

	if err := json.Unmarshal(jsonBytes, message); err != nil {
		log.Fatal(err)
	}

	postMessage := fmt.Sprintf(
		"[info][title]%s %s(%s)[/title]%s[/info]",
		message.AlarmName,
		message.AlarmDescription,
		message.NewStateValue,
		message.NewStateReason,
	)

	chatwork := chatwork.NewClient(os.Getenv("CHATWORK_API_KEY"))

	response, err := chatwork.PostRoomMessage(os.Getenv("CHATWORK_ROOM_ID"), postMessage)

	return string(response), err
}

func main() {
	lambda.Start(Handler)
}
