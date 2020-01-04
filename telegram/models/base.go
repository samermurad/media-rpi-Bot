package models

import "time"

type ServerResponse struct {
	Ok     bool        `json:"ok"`
	Result interface{} `json:"result"`
}

type TextEntity struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

type From struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	IsBot     bool   `json:"is_bot"`
	LangCode  string `json:"language_code"`
}

type Chat struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type Message struct {
	MessageId int64         `json:"message_id"`
	From      From          `json:"from"`
	Chat      Chat          `json:"chat"`
	Entities  []TextEntity  `json:"entities"`
	Text      string        `json:"text"`
	Date      time.Duration `json:"date"`
}

type Update struct {
	UpdateId int64   `json:"update_id"`
	Message  Message `json:"message"`
}
type BotMessageParseMode string

const (
	HTML       BotMessageParseMode = "HTML"
	Markdown   BotMessageParseMode = "Markdown"
	MarkdownV2 BotMessageParseMode = "MarkdownV2"
)

type BotMessage struct {
	ChatId    int64               `json:"chat_id"`
	Text      string              `json:"text"`
	ParseMode BotMessageParseMode `json:"parse_mode"`
	MessageId int64               `json:"message_id,omitempty"`
}
