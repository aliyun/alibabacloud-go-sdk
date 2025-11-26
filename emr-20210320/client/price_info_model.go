// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPriceInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCurrency(v string) *PriceInfo
	GetCurrency() *string
	SetDiscountPrice(v string) *PriceInfo
	GetDiscountPrice() *string
	SetOriginalPrice(v string) *PriceInfo
	GetOriginalPrice() *string
	SetPayType(v string) *PriceInfo
	GetPayType() *string
	SetPromotionResults(v []*PromotionInfo) *PriceInfo
	GetPromotionResults() []*PromotionInfo
	SetResourceType(v string) *PriceInfo
	GetResourceType() *string
	SetSpotInstanceTypeOriginalPrice(v string) *PriceInfo
	GetSpotInstanceTypeOriginalPrice() *string
	SetSpotInstanceTypePrice(v string) *PriceInfo
	GetSpotInstanceTypePrice() *string
	SetSpotOriginalPrice(v string) *PriceInfo
	GetSpotOriginalPrice() *string
	SetSpotPrice(v string) *PriceInfo
	GetSpotPrice() *string
	SetTaxPrice(v string) *PriceInfo
	GetTaxPrice() *string
	SetTradePrice(v string) *PriceInfo
	GetTradePrice() *string
}

type PriceInfo struct {
	Currency                      *string          `json:"Currency,omitempty" xml:"Currency,omitempty"`
	DiscountPrice                 *string          `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	OriginalPrice                 *string          `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	PayType                       *string          `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PromotionResults              []*PromotionInfo `json:"PromotionResults,omitempty" xml:"PromotionResults,omitempty" type:"Repeated"`
	ResourceType                  *string          `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SpotInstanceTypeOriginalPrice *string          `json:"SpotInstanceTypeOriginalPrice,omitempty" xml:"SpotInstanceTypeOriginalPrice,omitempty"`
	SpotInstanceTypePrice         *string          `json:"SpotInstanceTypePrice,omitempty" xml:"SpotInstanceTypePrice,omitempty"`
	SpotOriginalPrice             *string          `json:"SpotOriginalPrice,omitempty" xml:"SpotOriginalPrice,omitempty"`
	SpotPrice                     *string          `json:"SpotPrice,omitempty" xml:"SpotPrice,omitempty"`
	TaxPrice                      *string          `json:"TaxPrice,omitempty" xml:"TaxPrice,omitempty"`
	TradePrice                    *string          `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s PriceInfo) String() string {
	return dara.Prettify(s)
}

func (s PriceInfo) GoString() string {
	return s.String()
}

func (s *PriceInfo) GetCurrency() *string {
	return s.Currency
}

func (s *PriceInfo) GetDiscountPrice() *string {
	return s.DiscountPrice
}

func (s *PriceInfo) GetOriginalPrice() *string {
	return s.OriginalPrice
}

func (s *PriceInfo) GetPayType() *string {
	return s.PayType
}

func (s *PriceInfo) GetPromotionResults() []*PromotionInfo {
	return s.PromotionResults
}

func (s *PriceInfo) GetResourceType() *string {
	return s.ResourceType
}

func (s *PriceInfo) GetSpotInstanceTypeOriginalPrice() *string {
	return s.SpotInstanceTypeOriginalPrice
}

func (s *PriceInfo) GetSpotInstanceTypePrice() *string {
	return s.SpotInstanceTypePrice
}

func (s *PriceInfo) GetSpotOriginalPrice() *string {
	return s.SpotOriginalPrice
}

func (s *PriceInfo) GetSpotPrice() *string {
	return s.SpotPrice
}

func (s *PriceInfo) GetTaxPrice() *string {
	return s.TaxPrice
}

func (s *PriceInfo) GetTradePrice() *string {
	return s.TradePrice
}

func (s *PriceInfo) SetCurrency(v string) *PriceInfo {
	s.Currency = &v
	return s
}

func (s *PriceInfo) SetDiscountPrice(v string) *PriceInfo {
	s.DiscountPrice = &v
	return s
}

func (s *PriceInfo) SetOriginalPrice(v string) *PriceInfo {
	s.OriginalPrice = &v
	return s
}

func (s *PriceInfo) SetPayType(v string) *PriceInfo {
	s.PayType = &v
	return s
}

func (s *PriceInfo) SetPromotionResults(v []*PromotionInfo) *PriceInfo {
	s.PromotionResults = v
	return s
}

func (s *PriceInfo) SetResourceType(v string) *PriceInfo {
	s.ResourceType = &v
	return s
}

func (s *PriceInfo) SetSpotInstanceTypeOriginalPrice(v string) *PriceInfo {
	s.SpotInstanceTypeOriginalPrice = &v
	return s
}

func (s *PriceInfo) SetSpotInstanceTypePrice(v string) *PriceInfo {
	s.SpotInstanceTypePrice = &v
	return s
}

func (s *PriceInfo) SetSpotOriginalPrice(v string) *PriceInfo {
	s.SpotOriginalPrice = &v
	return s
}

func (s *PriceInfo) SetSpotPrice(v string) *PriceInfo {
	s.SpotPrice = &v
	return s
}

func (s *PriceInfo) SetTaxPrice(v string) *PriceInfo {
	s.TaxPrice = &v
	return s
}

func (s *PriceInfo) SetTradePrice(v string) *PriceInfo {
	s.TradePrice = &v
	return s
}

func (s *PriceInfo) Validate() error {
	if s.PromotionResults != nil {
		for _, item := range s.PromotionResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
