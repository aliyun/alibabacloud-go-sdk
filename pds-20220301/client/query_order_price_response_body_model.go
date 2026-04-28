// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrderPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiscountPrice(v float64) *QueryOrderPriceResponseBody
	GetDiscountPrice() *float64
	SetOriginalPrice(v float64) *QueryOrderPriceResponseBody
	GetOriginalPrice() *float64
	SetTradePrice(v float64) *QueryOrderPriceResponseBody
	GetTradePrice() *float64
}

type QueryOrderPriceResponseBody struct {
	DiscountPrice *float64 `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
	OriginalPrice *float64 `json:"original_price,omitempty" xml:"original_price,omitempty"`
	TradePrice    *float64 `json:"trade_price,omitempty" xml:"trade_price,omitempty"`
}

func (s QueryOrderPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryOrderPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderPriceResponseBody) GetDiscountPrice() *float64 {
	return s.DiscountPrice
}

func (s *QueryOrderPriceResponseBody) GetOriginalPrice() *float64 {
	return s.OriginalPrice
}

func (s *QueryOrderPriceResponseBody) GetTradePrice() *float64 {
	return s.TradePrice
}

func (s *QueryOrderPriceResponseBody) SetDiscountPrice(v float64) *QueryOrderPriceResponseBody {
	s.DiscountPrice = &v
	return s
}

func (s *QueryOrderPriceResponseBody) SetOriginalPrice(v float64) *QueryOrderPriceResponseBody {
	s.OriginalPrice = &v
	return s
}

func (s *QueryOrderPriceResponseBody) SetTradePrice(v float64) *QueryOrderPriceResponseBody {
	s.TradePrice = &v
	return s
}

func (s *QueryOrderPriceResponseBody) Validate() error {
	return dara.Validate(s)
}
