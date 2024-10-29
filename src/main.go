package main

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Replace with your bot token
	botToken := readToken()
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN error")
	}

	// Create a new bot instance
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Create an update config
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Start receiving updates
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-message updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Respond to the user
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You said: "+update.Message.Text)
		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}
