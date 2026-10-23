// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceGroupPriceForRenewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrice(v *DescribeInstanceGroupPriceForRenewResponseBodyPrice) *DescribeInstanceGroupPriceForRenewResponseBody
	GetPrice() *DescribeInstanceGroupPriceForRenewResponseBodyPrice
	SetRequestId(v string) *DescribeInstanceGroupPriceForRenewResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeInstanceGroupPriceForRenewResponseBodyRules) *DescribeInstanceGroupPriceForRenewResponseBody
	GetRules() []*DescribeInstanceGroupPriceForRenewResponseBodyRules
}

type DescribeInstanceGroupPriceForRenewResponseBody struct {
	// The price.
	Price *DescribeInstanceGroupPriceForRenewResponseBodyPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F8451630-77B9-5259-A9BC-950E9BC586B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The promotion rules.
	Rules []*DescribeInstanceGroupPriceForRenewResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeInstanceGroupPriceForRenewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGroupPriceForRenewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGroupPriceForRenewResponseBody) GetPrice() *DescribeInstanceGroupPriceForRenewResponseBodyPrice {
	return s.Price
}

func (s *DescribeInstanceGroupPriceForRenewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceGroupPriceForRenewResponseBody) GetRules() []*DescribeInstanceGroupPriceForRenewResponseBodyRules {
	return s.Rules
}

func (s *DescribeInstanceGroupPriceForRenewResponseBody) SetPrice(v *DescribeInstanceGroupPriceForRenewResponseBodyPrice) *DescribeInstanceGroupPriceForRenewResponseBody {
	s.Price = v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponseBody) SetRequestId(v string) *DescribeInstanceGroupPriceForRenewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponseBody) SetRules(v []*DescribeInstanceGroupPriceForRenewResponseBodyRules) *DescribeInstanceGroupPriceForRenewResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponseBody) Validate() error {
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

type DescribeInstanceGroupPriceForRenewResponseBodyPrice struct {
	// The currency unit.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discounted price.
	//
	// example:
	//
	// 1098.99
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 5007.0
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The promotion information.
	Promotions []*DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The transaction amount.
	//
	// example:
	//
	// 34.02
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeInstanceGroupPriceForRenewResponseBodyPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGroupPriceForRenewResponseBodyPrice) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPrice) GetPromotions() []*DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions {
	return s.Promotions
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPrice) SetCurrency(v string) *DescribeInstanceGroupPriceForRenewResponseBodyPrice {
	s.Currency = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPrice) SetDiscountPrice(v float32) *DescribeInstanceGroupPriceForRenewResponseBodyPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPrice) SetOriginalPrice(v float32) *DescribeInstanceGroupPriceForRenewResponseBodyPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPrice) SetPromotions(v []*DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions) *DescribeInstanceGroupPriceForRenewResponseBodyPrice {
	s.Promotions = v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPrice) SetTradePrice(v float32) *DescribeInstanceGroupPriceForRenewResponseBodyPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPrice) Validate() error {
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

type DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions struct {
	// The option code.
	//
	// example:
	//
	// youhui_quan
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The promotion description.
	//
	// example:
	//
	// 20 off on orders over 40
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 87996380****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The promotion name.
	//
	// example:
	//
	// 20 off on orders over 40
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// Indicates whether the promotion is selected.
	//
	// example:
	//
	// true
	Selected *bool `json:"selected,omitempty" xml:"selected,omitempty"`
}

func (s DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions) SetOptionCode(v string) *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions) SetPromotionDesc(v string) *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions) SetPromotionId(v string) *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions) SetPromotionName(v string) *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions) SetSelected(v bool) *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions {
	s.Selected = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyPricePromotions) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceGroupPriceForRenewResponseBodyRules struct {
	// The activity description.
	//
	// example:
	//
	// 20 off on orders over 40
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// t1uf7xt141ut5i****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeInstanceGroupPriceForRenewResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceGroupPriceForRenewResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyRules) SetDescription(v string) *DescribeInstanceGroupPriceForRenewResponseBodyRules {
	s.Description = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyRules) SetRuleId(v string) *DescribeInstanceGroupPriceForRenewResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeInstanceGroupPriceForRenewResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
