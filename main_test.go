package main_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	main "github.com/kawax/lambda-alarm-chatwork"
)

func TestHandler(t *testing.T) {

	inputJson := readJsonFromFile(t, "./testdata/sns-event.json")

	var event events.SNSEvent

	if err := json.Unmarshal(inputJson, &event); err != nil {
		t.Errorf("could not unmarshal event. details: %v", err)
	}

	response, err := main.Handler(context.Background(), event)

	// ChatWorkへの投稿以降はひとまず無視する。apikeyがなくてエラーになることを確認。
	// interface でモックをやろうとしてたけどこれでそこまでやる必要はなさそうなのでやめた。
	if err == nil {
		t.Error(err)
	}

	if response != "error" {
		t.Error(response)
	}
}

func readJsonFromFile(t *testing.T, inputFile string) []byte {
	inputJson, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("could not open test file. details: %v", err)
	}

	return inputJson
}
