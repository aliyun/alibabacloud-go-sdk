// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceGroupPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrice(v *DescribeInstanceGroupPriceResponseBodyPrice) *DescribeInstanceGroupPriceResponseBody
	GetPrice() *DescribeInstanceGroupPriceResponseBodyPrice
	SetRequestId(v string) *DescribeInstanceGroupPriceResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeInstanceGroupPriceResponseBodyRules) *DescribeInstanceGroupPriceResponseBody
	GetRules() []*DescribeInstanceGroupPriceResponseBodyRules
}

type DescribeInstanceGroupPriceResponseBody struct {
	Price *DescribeInstanceGroupPriceResponseBodyPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// example:
	//
	// 90E4F2B7-BBDA-534D-B872-CCC798996F69
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*DescribeInstanceGroupPriceResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeInstanceGroupPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGroupPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGroupPriceResponseBody) GetPrice() *DescribeInstanceGroupPriceResponseBodyPrice {
	return s.Price
}

func (s *DescribeInstanceGroupPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceGroupPriceResponseBody) GetRules() []*DescribeInstanceGroupPriceResponseBodyRules {
	return s.Rules
}

func (s *DescribeInstanceGroupPriceResponseBody) SetPrice(v *DescribeInstanceGroupPriceResponseBodyPrice) *DescribeInstanceGroupPriceResponseBody {
	s.Price = v
	return s
}

func (s *DescribeInstanceGroupPriceResponseBody) SetRequestId(v string) *DescribeInstanceGroupPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceGroupPriceResponseBody) SetRules(v []*DescribeInstanceGroupPriceResponseBodyRules) *DescribeInstanceGroupPriceResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeInstanceGroupPriceResponseBody) Validate() error {
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

type DescribeInstanceGroupPriceResponseBodyPrice struct {
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 12618.0
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// 18134.0
	OriginalPrice *float32                                                 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	Promotions    []*DescribeInstanceGroupPriceResponseBodyPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// example:
	//
	// 70.0
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeInstanceGroupPriceResponseBodyPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGroupPriceResponseBodyPrice) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGroupPriceResponseBodyPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeInstanceGroupPriceResponseBodyPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeInstanceGroupPriceResponseBodyPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeInstanceGroupPriceResponseBodyPrice) GetPromotions() []*DescribeInstanceGroupPriceResponseBodyPricePromotions {
	return s.Promotions
}

func (s *DescribeInstanceGroupPriceResponseBodyPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeInstanceGroupPriceResponseBodyPrice) SetCurrency(v string) *DescribeInstanceGroupPriceResponseBodyPrice {
	s.Currency = &v
	return s
}

func (s *DescribeInstanceGroupPriceResponseBodyPrice) SetDiscountPrice(v float32) *DescribeInstanceGroupPriceResponseBodyPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeInstanceGroupPriceResponseBodyPrice) SetOriginalPrice(v float32) *DescribeInstanceGroupPriceResponseBodyPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeInstanceGroupPriceResponseBodyPrice) SetPromotions(v []*DescribeInstanceGroupPriceResponseBodyPricePromotions) *DescribeInstanceGroupPriceResponseBodyPrice {
	s.Promotions = v
	return s
}

func (s *DescribeInstanceGroupPriceResponseBodyPrice) SetTradePrice(v float32) *DescribeInstanceGroupPriceResponseBodyPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeInstanceGroupPriceResponseBodyPrice) Validate() error {
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

type DescribeInstanceGroupPriceResponseBodyPricePromotions struct {
	// example:
	//
	// code
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// example:
	//
	// 日常优惠
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// example:
	//
	// 86744342****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// example:
	//
	// 日常优惠
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// example:
	//
	// true
	Selected *bool `json:"selected,omitempty" xml:"selected,omitempty"`
}

func (s DescribeInstanceGroupPriceResponseBodyPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGroupPriceResponseBodyPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGroupPriceResponseBodyPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribeInstanceGroupPriceResponseBodyPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *DescribeInstanceGroupPriceResponseBodyPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeInstanceGroupPriceResponseBodyPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribeInstanceGroupPriceResponseBodyPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *DescribeInstanceGroupPriceResponseBodyPricePromotions) SetOptionCode(v string) *DescribeInstanceGroupPriceResponseBodyPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeInstanceGroupPriceResponseBodyPricePromotions) SetPromotionDesc(v string) *DescribeInstanceGroupPriceResponseBodyPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeInstanceGroupPriceResponseBodyPricePromotions) SetPromotionId(v string) *DescribeInstanceGroupPriceResponseBodyPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeInstanceGroupPriceResponseBodyPricePromotions) SetPromotionName(v string) *DescribeInstanceGroupPriceResponseBodyPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeInstanceGroupPriceResponseBodyPricePromotions) SetSelected(v bool) *DescribeInstanceGroupPriceResponseBodyPricePromotions {
	s.Selected = &v
	return s
}

func (s *DescribeInstanceGroupPriceResponseBodyPricePromotions) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceGroupPriceResponseBodyRules struct {
	// example:
	//
	// 日常优惠
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// yhpcxb4qhgp2k4****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeInstanceGroupPriceResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGroupPriceResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGroupPriceResponseBodyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceGroupPriceResponseBodyRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeInstanceGroupPriceResponseBodyRules) SetDescription(v string) *DescribeInstanceGroupPriceResponseBodyRules {
	s.Description = &v
	return s
}

func (s *DescribeInstanceGroupPriceResponseBodyRules) SetRuleId(v string) *DescribeInstanceGroupPriceResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeInstanceGroupPriceResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
