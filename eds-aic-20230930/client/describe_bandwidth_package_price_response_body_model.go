// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBandwidthPackagePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrice(v *DescribeBandwidthPackagePriceResponseBodyPrice) *DescribeBandwidthPackagePriceResponseBody
	GetPrice() *DescribeBandwidthPackagePriceResponseBodyPrice
	SetRequestId(v string) *DescribeBandwidthPackagePriceResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeBandwidthPackagePriceResponseBodyRules) *DescribeBandwidthPackagePriceResponseBody
	GetRules() []*DescribeBandwidthPackagePriceResponseBodyRules
}

type DescribeBandwidthPackagePriceResponseBody struct {
	Price *DescribeBandwidthPackagePriceResponseBodyPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// example:
	//
	// 6C8439B9-7DBF-57F4-92AE-55A9B9D3****
	RequestId *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*DescribeBandwidthPackagePriceResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeBandwidthPackagePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackagePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackagePriceResponseBody) GetPrice() *DescribeBandwidthPackagePriceResponseBodyPrice {
	return s.Price
}

func (s *DescribeBandwidthPackagePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBandwidthPackagePriceResponseBody) GetRules() []*DescribeBandwidthPackagePriceResponseBodyRules {
	return s.Rules
}

func (s *DescribeBandwidthPackagePriceResponseBody) SetPrice(v *DescribeBandwidthPackagePriceResponseBodyPrice) *DescribeBandwidthPackagePriceResponseBody {
	s.Price = v
	return s
}

func (s *DescribeBandwidthPackagePriceResponseBody) SetRequestId(v string) *DescribeBandwidthPackagePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBandwidthPackagePriceResponseBody) SetRules(v []*DescribeBandwidthPackagePriceResponseBodyRules) *DescribeBandwidthPackagePriceResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeBandwidthPackagePriceResponseBody) Validate() error {
	if s.Price != nil {
		if err := s.Price.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBandwidthPackagePriceResponseBodyPrice struct {
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 274.75
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// 8800
	OriginalPrice *float32                                                    `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	Promotions    []*DescribeBandwidthPackagePriceResponseBodyPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// example:
	//
	// 213.52
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeBandwidthPackagePriceResponseBodyPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackagePriceResponseBodyPrice) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackagePriceResponseBodyPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeBandwidthPackagePriceResponseBodyPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeBandwidthPackagePriceResponseBodyPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeBandwidthPackagePriceResponseBodyPrice) GetPromotions() []*DescribeBandwidthPackagePriceResponseBodyPricePromotions {
	return s.Promotions
}

func (s *DescribeBandwidthPackagePriceResponseBodyPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeBandwidthPackagePriceResponseBodyPrice) SetCurrency(v string) *DescribeBandwidthPackagePriceResponseBodyPrice {
	s.Currency = &v
	return s
}

func (s *DescribeBandwidthPackagePriceResponseBodyPrice) SetDiscountPrice(v float32) *DescribeBandwidthPackagePriceResponseBodyPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeBandwidthPackagePriceResponseBodyPrice) SetOriginalPrice(v float32) *DescribeBandwidthPackagePriceResponseBodyPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeBandwidthPackagePriceResponseBodyPrice) SetPromotions(v []*DescribeBandwidthPackagePriceResponseBodyPricePromotions) *DescribeBandwidthPackagePriceResponseBodyPrice {
	s.Promotions = v
	return s
}

func (s *DescribeBandwidthPackagePriceResponseBodyPrice) SetTradePrice(v float32) *DescribeBandwidthPackagePriceResponseBodyPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeBandwidthPackagePriceResponseBodyPrice) Validate() error {
	if s.Promotions != nil {
		for _, item := range s.Promotions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBandwidthPackagePriceResponseBodyPricePromotions struct {
	// example:
	//
	// youhui_quan
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// example:
	//
	// youhui
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// example:
	//
	// 91381816****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// example:
	//
	// youhui
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// example:
	//
	// ture
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s DescribeBandwidthPackagePriceResponseBodyPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackagePriceResponseBodyPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackagePriceResponseBodyPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribeBandwidthPackagePriceResponseBodyPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *DescribeBandwidthPackagePriceResponseBodyPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeBandwidthPackagePriceResponseBodyPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribeBandwidthPackagePriceResponseBodyPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *DescribeBandwidthPackagePriceResponseBodyPricePromotions) SetOptionCode(v string) *DescribeBandwidthPackagePriceResponseBodyPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeBandwidthPackagePriceResponseBodyPricePromotions) SetPromotionDesc(v string) *DescribeBandwidthPackagePriceResponseBodyPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeBandwidthPackagePriceResponseBodyPricePromotions) SetPromotionId(v string) *DescribeBandwidthPackagePriceResponseBodyPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeBandwidthPackagePriceResponseBodyPricePromotions) SetPromotionName(v string) *DescribeBandwidthPackagePriceResponseBodyPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeBandwidthPackagePriceResponseBodyPricePromotions) SetSelected(v bool) *DescribeBandwidthPackagePriceResponseBodyPricePromotions {
	s.Selected = &v
	return s
}

func (s *DescribeBandwidthPackagePriceResponseBodyPricePromotions) Validate() error {
	return dara.Validate(s)
}

type DescribeBandwidthPackagePriceResponseBodyRules struct {
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// a7gesyly7o34qy****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeBandwidthPackagePriceResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeBandwidthPackagePriceResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeBandwidthPackagePriceResponseBodyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeBandwidthPackagePriceResponseBodyRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeBandwidthPackagePriceResponseBodyRules) SetDescription(v string) *DescribeBandwidthPackagePriceResponseBodyRules {
	s.Description = &v
	return s
}

func (s *DescribeBandwidthPackagePriceResponseBodyRules) SetRuleId(v string) *DescribeBandwidthPackagePriceResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeBandwidthPackagePriceResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
