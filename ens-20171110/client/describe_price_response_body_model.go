// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *DescribePriceResponseBodyPriceInfo) *DescribePriceResponseBody
	GetPriceInfo() *DescribePriceResponseBodyPriceInfo
	SetRequestId(v string) *DescribePriceResponseBody
	GetRequestId() *string
}

type DescribePriceResponseBody struct {
	// The price information.
	PriceInfo *DescribePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 694EB8FF-446C-5B1E-8146-9629E492551F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) GetPriceInfo() *DescribePriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *DescribePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePriceResponseBody) SetPriceInfo(v *DescribePriceResponseBodyPriceInfo) *DescribePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribePriceResponseBody) SetRequestId(v string) *DescribePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePriceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPriceInfo struct {
	// The price.
	Price *DescribePriceResponseBodyPriceInfoPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
}

func (s DescribePriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfo) GetPrice() *DescribePriceResponseBodyPriceInfoPrice {
	return s.Price
}

func (s *DescribePriceResponseBodyPriceInfo) SetPrice(v *DescribePriceResponseBodyPriceInfoPrice) *DescribePriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPriceInfoPrice struct {
	// The type of currency.
	//
	// 	- USD: US dollar
	//
	// 	- JPY: Japanese Yen
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discount.
	//
	// example:
	//
	// 1
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 16125.00
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The final price.
	//
	// example:
	//
	// 126.75
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribePriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) Validate() error {
	return dara.Validate(s)
}
