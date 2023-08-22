package commands

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) DefaultBehavior(inputMessage *tgbotapi.Message) {
	fmt.Printf("[%s] %s\n", inputMessage.From.UserName, inputMessage.Text)

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	_, _ = c.bot.Send(msg)
}
