package mexcwsmarket

import (
	"github.com/mkovrigovich/mexc-golang-sdk/websocket"
)

type Service struct {
	client *mexcws.MEXCWebSocket
}

func New(client *mexcws.MEXCWebSocket) *Service {
	return &Service{
		client: client,
	}
}
