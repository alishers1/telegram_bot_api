package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/telegram-bot-api/internal/service/product"
)

func main() {
	godotenv.Load()

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	for update := range updates {
		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		case "list":
			listCommand(bot, update.Message, productService)
		default:
			defaultBehavior(bot, update.Message)
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, 
		"/help - help\n"+
			"/list - list products",
	)

	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message, productService *product.Service) {
	outputText := "Here are all products: \n\n"

	products := productService.List()

	for _, p := range products {
		outputText += p.Title
		outputText += "\n"
	}

	


	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputText)

	bot.Send(msg)
}



func defaultBehavior(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "You wrote: "+inputMsg.Text)
	msg.ReplyToMessageID = inputMsg.MessageID

	bot.Send(msg)
}
