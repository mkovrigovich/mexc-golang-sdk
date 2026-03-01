package mexcwstypes

import "github.com/mkovrigovich/mexc-golang-sdk/websocket/dto"

type OnReceive func(message *dto.PushDataV3ApiWrapper)
type OnError func(connClosed bool, err error)

type WsReq struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
}
