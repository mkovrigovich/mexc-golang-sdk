package mexchttpmarket

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mkovrigovich/mexc-golang-sdk/consts"
)

// BatchOrders https://www.mexc.com/api-docs/spot-v3/spot-account-trade#batch-orders
func (s *Service) BatchOrders(ctx context.Context, requests []BatchOrdersRequest) ([]BatchOrdersResponse, error) {
	allParams := make(map[string]string)
	allParams["batchOrders"] = "["

	for idx, req := range requests {
		params := make(map[string]string)
		params["symbol"] = req.Symbol
		params["side"] = string(req.Side)
		params["type"] = string(req.Type)
		params["timestamp"] = s.getTimestamp()

		if req.Quantity != nil {
			params["quantity"] = *req.Quantity
		}
		if req.QuoteOrderQty != nil {
			params["quoteOrderQty"] = *req.QuoteOrderQty
		}
		if req.Price != nil {
			params["price"] = *req.Price
		}
		if req.NewClientOrderId != nil {
			params["newClientOrderId"] = *req.NewClientOrderId
		}
		if req.RecvWindow != nil {
			params["recvWindow"] = fmt.Sprintf("%d", *req.RecvWindow)
		}

		bytes, err := json.Marshal(params)
		if err != nil {
			return nil, err
		}
		allParams["batchOrders"] = allParams["batchOrders"] + string(bytes)
		if (idx + 1) < len(requests) {
			allParams["batchOrders"] = allParams["batchOrders"] + ","
		}
	}

	allParams["batchOrders"] = allParams["batchOrders"] + "]"
	fmt.Println(allParams)

	res, err := s.client.SendRequest(ctx, http.MethodPost, consts.EndpointBatchOrders, allParams)
	if err != nil {
		return nil, err
	}

	var orderResponse []BatchOrdersResponse
	err = json.Unmarshal(res, &orderResponse)
	if err != nil {
		return nil, err
	}

	return orderResponse, nil
}

type BatchOrdersRequest struct {
	Symbol           string  `json:"symbol"`
	Side             Side    `json:"side"`
	Type             Type    `json:"type"`
	Quantity         *string `json:"quantity,omitempty"`
	QuoteOrderQty    *string `json:"quoteOrderQty,omitempty"`
	Price            *string `json:"price,omitempty"`
	NewClientOrderId *string `json:"newClientOrderId,omitempty"`
	RecvWindow       *int64  `json:"recvWindow,omitempty"`
}

type BatchOrdersResponse struct {
	Symbol      string `json:"symbol"`
	OrderId     string `json:"orderId"`
	OrderListId int    `json:"orderListId"`

	Msg  string `json:"msg"`
	Code string `json:"code"`
	//Price        decimal.Decimal `json:"price"`
	//OrigQty      decimal.Decimal `json:"origQty"`
	//Type         Type            `json:"type"`
	//Side         Side            `json:"side"`
	//TransactTime int64           `json:"transactTime"`
}
