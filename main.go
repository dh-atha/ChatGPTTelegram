package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err)
	}

	chatgpt := NewChatGPT(os.Getenv("OPENAPI_KEY"))

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEBOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	log.Println("Waiting messages....")

	for update := range updates {
		if update.Message != nil {
			name := update.Message.From.FirstName
			_, ok := session[name]
			if !ok && update.Message.Text != "/start" {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "send /start to start session")
				bot.Send(msg)

				continue
			}

			if !ok && update.Message.Text == "/start" {
				reply := NewSession(name)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
				bot.Send(msg)

				continue
			}

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			AddChat(name, User, update.Message.Text)
			message, err := chatgpt.CreateChatCompletion(name)
			if err != nil {
				log.Println(err)
				continue
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)

			bot.Send(msg)
		}
	}
}
