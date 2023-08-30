package main

import (
	"log"
	"os"

	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/yaml.v3"
)

func main() {

	//read file with rune mappings
	mappingFile := "mapping.map"
	mapping, err := os.ReadFile(mappingFile)
	if err != nil {
		log.Panic("Error:", err)
	}

	//read file with rune mappings
	tokenFile := "token.tkn"
	token, err := os.ReadFile(tokenFile)
	if err != nil {
		log.Panic("Error:", err)
	}

	var data map[int]int

	err = yaml.Unmarshal(mapping, &data)
	if err != nil {
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(string(token))
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("Chat.ID:%d [%s:%d] %s\n", update.Message.Chat.ID, update.Message.From.UserName, update.Message.From.ID, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, remapString(update.Message.Text, data))

		bot.Send(msg)
	}
}

// remapString сhanges characters according to the data specified in the mapping.map file
func remapString(s string, m map[int]int) string {
	var sb strings.Builder
	for _, v := range s {
		if val, ok := m[int(v)]; ok {
			sb.WriteString(string(rune(val)))
		} else {
			sb.WriteString(string(v))
		}
	}
	return sb.String()
}
