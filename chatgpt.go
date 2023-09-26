package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ChatGPT struct {
	client *http.Client
	token  string
}

func NewChatGPT(token string) *ChatGPT {
	return &ChatGPT{
		token:  token,
		client: http.DefaultClient,
	}
}

func (c *ChatGPT) Connect() {
	fmt.Println("chatgpt connected")
}

func (c *ChatGPT) CreateChatCompletion(name string) (message string, err error) {
	url := "https://api.openai.com/v1/chat/completions"

	request := RequestChatCompletion{
		Model:    "gpt-3.5-turbo",
		Messages: session[name],
	}
	reqBody, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	// log.Println(string(reqBody))

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+c.token)
	req.Header.Set("Content-Type", "application/json")

	response, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	rawRes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// log.Println(string(rawRes))

	var resStruct ResponseChatCompletion
	err = json.Unmarshal(rawRes, &resStruct)
	if err != nil {
		return "", err
	}

	message = resStruct.Choices[0].Message.Content
	AddChat(name, Assistant, message)
	return message, nil
}
