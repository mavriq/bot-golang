package botgolang

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type BotOption interface {
	Type() string
	Value() interface{}
}

type BotApiURL string

func (o BotApiURL) Type() string {
	return "api_url"
}

func (o BotApiURL) Value() interface{} {
	return string(o)
}

type BotDebug bool

func (o BotDebug) Type() string {
	return "debug"
}

func (o BotDebug) Value() interface{} {
	return bool(o)
}

type BotHTTPClient http.Client

func (o BotHTTPClient) Type() string {
	return "http_client"
}

func (o BotHTTPClient) Value() interface{} {
	return http.Client(o)
}

type botLogger struct {
	*logrus.Logger
}

func BotLogger(l *logrus.Logger) botLogger {
	return botLogger{l}
}

func (o botLogger) Type() string {
	return "logger"
}

func (o botLogger) Value() interface{} {
	return (*logrus.Logger)(o.Logger)
}
