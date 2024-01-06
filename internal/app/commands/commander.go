package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/telegram-bot-api/internal/service/product"
)

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewRouter(bot *tgbotapi.BotAPI, productService *product.Service) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}
