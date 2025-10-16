// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenewalPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *DescribeRenewalPriceResponseBodyPriceInfo) *DescribeRenewalPriceResponseBody
	GetPriceInfo() *DescribeRenewalPriceResponseBodyPriceInfo
	SetRequestId(v string) *DescribeRenewalPriceResponseBody
	GetRequestId() *string
}

type DescribeRenewalPriceResponseBody struct {
	// The price details.
	PriceInfo *DescribeRenewalPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 72E47B1E-6B11-5A11-A27C-7A80F866****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRenewalPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBody) GetPriceInfo() *DescribeRenewalPriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *DescribeRenewalPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRenewalPriceResponseBody) SetPriceInfo(v *DescribeRenewalPriceResponseBodyPriceInfo) *DescribeRenewalPriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribeRenewalPriceResponseBody) SetRequestId(v string) *DescribeRenewalPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRenewalPriceResponseBody) Validate() error {
	if s.PriceInfo != nil {
		if err := s.PriceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRenewalPriceResponseBodyPriceInfo struct {
	// The price.
	Price *DescribeRenewalPriceResponseBodyPriceInfoPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// The promotion rules.
	Rules []*DescribeRenewalPriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) GetPrice() *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	return s.Price
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) GetRules() []*DescribeRenewalPriceResponseBodyPriceInfoRules {
	return s.Rules
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) SetPrice(v *DescribeRenewalPriceResponseBodyPriceInfoPrice) *DescribeRenewalPriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) SetRules(v []*DescribeRenewalPriceResponseBodyPriceInfoRules) *DescribeRenewalPriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) Validate() error {
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

type DescribeRenewalPriceResponseBodyPriceInfoPrice struct {
	// The unit of currency (USD).
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discounted amount.
	//
	// example:
	//
	// 15.8
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The orders.
	OrderLines map[string]*string `json:"OrderLines,omitempty" xml:"OrderLines,omitempty"`
	// The original price.
	//
	// example:
	//
	// 79.0
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The promotions.
	Promotions []*DescribeRenewalPriceResponseBodyPriceInfoPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The actual price. The actual price is the original price minus the discount.
	//
	// example:
	//
	// 63.2
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) GetOrderLines() map[string]*string {
	return s.OrderLines
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) GetPromotions() []*DescribeRenewalPriceResponseBodyPriceInfoPricePromotions {
	return s.Promotions
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetOrderLines(v map[string]*string) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.OrderLines = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetPromotions(v []*DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.Promotions = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribeRenewalPriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPrice) Validate() error {
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

type DescribeRenewalPriceResponseBodyPriceInfoPricePromotions struct {
	// The description of the promotion rule.
	//
	// example:
	//
	// test
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The promotion description.
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// promo_option
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The promotion name.
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// Indicates whether an item is selected.
	//
	// example:
	//
	// false
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) SetOptionCode(v string) *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) SetPromotionDesc(v string) *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) SetPromotionId(v string) *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) SetPromotionName(v string) *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) SetSelected(v bool) *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions {
	s.Selected = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoPricePromotions) Validate() error {
	return dara.Validate(s)
}

type DescribeRenewalPriceResponseBodyPriceInfoRules struct {
	// The rule description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 29644
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRules) SetDescription(v string) *DescribeRenewalPriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRules) SetRuleId(v int64) *DescribeRenewalPriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRules) Validate() error {
	return dara.Validate(s)
}
