package main

import "fmt"

var session = make(map[string][]Chat)

type role string

const (
	User      role = "user"
	Assistant role = "assistant"
)

type Chat struct {
	Role    role   `json:"role"`
	Content string `json:"content"`
}

func NewSession(name string) string {
	session[name] = []Chat{}
	return fmt.Sprintf("Hello %s! How can I help you?", name)
}

func AddChat(name string, role role, content string) {
	session[name] = append(session[name], Chat{
		Role:    role,
		Content: content,
	})
}
