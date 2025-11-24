package mexcwsmarket

import (
	"context"
	"fmt"
	"github.com/kattana-io/mexc-golang-sdk/websocket/dto"
)

const (
	TradesMinInterval    = "10ms"
	TradesMaxInterval    = "100ms"
	TradesRequestPattern = "spot@public.aggre.deals.v3.api.pb@%s@%s"
)

// TradesSubscribe subscribes to aggregated trades updates for specified symbols with interval
// https://mexcdevelop.github.io/apidocs/spot_v3_en/#aggregate-deal-record
func (s *Service) TradesSubscribe(ctx context.Context, symbols []string, interval string, callback func(api *dto.PublicAggreDealsV3Api, symbol string)) error {
	lstnr := func(message *dto.PushDataV3ApiWrapper) {
		var symbol string
		if message.Symbol != nil {
			symbol = *message.Symbol
		}

		switch msg := message.Body.(type) {
		case *dto.PushDataV3ApiWrapper_PublicAggreDeals:
			callback(msg.PublicAggreDeals, symbol)
		default:
			fmt.Println("Trades callback unknown type:", message.Body)
		}
	}

	for _, symbol := range symbols {
		channel := fmt.Sprintf(TradesRequestPattern, interval, symbol)
		if err := s.client.Subscribe(ctx, channel, nil, lstnr); err != nil {
			return err
		}
	}

	return nil
}

// TradesUnsubscribe unsubscribes from trades updates for specified symbols
func (s *Service) TradesUnsubscribe(symbols []string, interval string) error {
	for _, symbol := range symbols {
		channel := fmt.Sprintf(TradesRequestPattern, interval, symbol)
		if err := s.client.Unsubscribe(channel); err != nil {
			return err
		}
	}

	return nil
}
