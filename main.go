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

var (
	apikey = os.Getenv("CHATWORK_API_KEY")
	room   = os.Getenv("CHATWORK_ROOM_ID")
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

	if len(snsEvent.Records) == 0 {
		return "error", errors.New("SNSRecord is empty")
	}

	snsRecord := snsEvent.Records[0].SNS

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

	if len(apikey) == 0 {
		return "error", errors.New("CHATWORK_API_KEY is empty")
	}

	if len(room) == 0 {
		return "error", errors.New("CHATWORK_ROOM_ID is empty")
	}

	cw := chatwork.NewClient(apikey)

	response, err := cw.PostRoomMessage(room, postMessage)

	return string(response), err
}

func main() {
	lambda.Start(Handler)
}
