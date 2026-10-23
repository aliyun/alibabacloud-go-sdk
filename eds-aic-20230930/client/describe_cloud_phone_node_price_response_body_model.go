// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudPhoneNodePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrice(v *DescribeCloudPhoneNodePriceResponseBodyPrice) *DescribeCloudPhoneNodePriceResponseBody
	GetPrice() *DescribeCloudPhoneNodePriceResponseBodyPrice
	SetRequestId(v string) *DescribeCloudPhoneNodePriceResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeCloudPhoneNodePriceResponseBodyRules) *DescribeCloudPhoneNodePriceResponseBody
	GetRules() []*DescribeCloudPhoneNodePriceResponseBodyRules
}

type DescribeCloudPhoneNodePriceResponseBody struct {
	Price *DescribeCloudPhoneNodePriceResponseBodyPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*DescribeCloudPhoneNodePriceResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeCloudPhoneNodePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodePriceResponseBody) GetPrice() *DescribeCloudPhoneNodePriceResponseBodyPrice {
	return s.Price
}

func (s *DescribeCloudPhoneNodePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudPhoneNodePriceResponseBody) GetRules() []*DescribeCloudPhoneNodePriceResponseBodyRules {
	return s.Rules
}

func (s *DescribeCloudPhoneNodePriceResponseBody) SetPrice(v *DescribeCloudPhoneNodePriceResponseBodyPrice) *DescribeCloudPhoneNodePriceResponseBody {
	s.Price = v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBody) SetRequestId(v string) *DescribeCloudPhoneNodePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBody) SetRules(v []*DescribeCloudPhoneNodePriceResponseBodyRules) *DescribeCloudPhoneNodePriceResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBody) Validate() error {
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

type DescribeCloudPhoneNodePriceResponseBodyPrice struct {
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 878.15
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// 2710
	OriginalPrice *float32                                                   `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	PriceDetail   []*DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail `json:"PriceDetail,omitempty" xml:"PriceDetail,omitempty" type:"Repeated"`
	Promotions    []*DescribeCloudPhoneNodePriceResponseBodyPricePromotions  `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// example:
	//
	// 99.0
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeCloudPhoneNodePriceResponseBodyPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodePriceResponseBodyPrice) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPrice) GetPriceDetail() []*DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail {
	return s.PriceDetail
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPrice) GetPromotions() []*DescribeCloudPhoneNodePriceResponseBodyPricePromotions {
	return s.Promotions
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPrice) SetCurrency(v string) *DescribeCloudPhoneNodePriceResponseBodyPrice {
	s.Currency = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPrice) SetDiscountPrice(v float32) *DescribeCloudPhoneNodePriceResponseBodyPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPrice) SetOriginalPrice(v float32) *DescribeCloudPhoneNodePriceResponseBodyPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPrice) SetPriceDetail(v []*DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail) *DescribeCloudPhoneNodePriceResponseBodyPrice {
	s.PriceDetail = v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPrice) SetPromotions(v []*DescribeCloudPhoneNodePriceResponseBodyPricePromotions) *DescribeCloudPhoneNodePriceResponseBodyPrice {
	s.Promotions = v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPrice) SetTradePrice(v float32) *DescribeCloudPhoneNodePriceResponseBodyPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPrice) Validate() error {
	if s.PriceDetail != nil {
		for _, item := range s.PriceDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail struct {
	// example:
	//
	// 1978.2
	DiscountPrice *float32                                                              `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	ModulePrice   []*DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice `json:"ModulePrice,omitempty" xml:"ModulePrice,omitempty" type:"Repeated"`
	// example:
	//
	// 4199
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// CloudPhoneMatrix
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 892.0
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail) GetModulePrice() []*DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice {
	return s.ModulePrice
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail) SetDiscountPrice(v float32) *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail) SetModulePrice(v []*DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice) *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail {
	s.ModulePrice = v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail) SetOriginalPrice(v float32) *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail) SetResourceType(v string) *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail {
	s.ResourceType = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail) SetTradePrice(v float32) *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail {
	s.TradePrice = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetail) Validate() error {
	if s.ModulePrice != nil {
		for _, item := range s.ModulePrice {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice struct {
	// example:
	//
	// 1978.2
	DiscountPrice *string `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// PhoneData
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// example:
	//
	// 8990
	OriginalPrice *string `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// 1666.0
	TradePrice *string `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice) GetDiscountPrice() *string {
	return s.DiscountPrice
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice) GetModuleName() *string {
	return s.ModuleName
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice) GetOriginalPrice() *string {
	return s.OriginalPrice
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice) GetTradePrice() *string {
	return s.TradePrice
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice) SetDiscountPrice(v string) *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice) SetModuleName(v string) *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice {
	s.ModuleName = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice) SetOriginalPrice(v string) *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice) SetTradePrice(v string) *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePriceDetailModulePrice) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudPhoneNodePriceResponseBodyPricePromotions struct {
	// example:
	//
	// daily
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// example:
	//
	// 日常优惠
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// example:
	//
	// 82842468****
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

func (s DescribeCloudPhoneNodePriceResponseBodyPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodePriceResponseBodyPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePromotions) SetOptionCode(v string) *DescribeCloudPhoneNodePriceResponseBodyPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePromotions) SetPromotionDesc(v string) *DescribeCloudPhoneNodePriceResponseBodyPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePromotions) SetPromotionId(v string) *DescribeCloudPhoneNodePriceResponseBodyPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePromotions) SetPromotionName(v string) *DescribeCloudPhoneNodePriceResponseBodyPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePromotions) SetSelected(v bool) *DescribeCloudPhoneNodePriceResponseBodyPricePromotions {
	s.Selected = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyPricePromotions) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudPhoneNodePriceResponseBodyRules struct {
	// example:
	//
	// 日常优惠
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// t1uf7xt141ut5i****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeCloudPhoneNodePriceResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodePriceResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodePriceResponseBodyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeCloudPhoneNodePriceResponseBodyRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeCloudPhoneNodePriceResponseBodyRules) SetDescription(v string) *DescribeCloudPhoneNodePriceResponseBodyRules {
	s.Description = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyRules) SetRuleId(v string) *DescribeCloudPhoneNodePriceResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeCloudPhoneNodePriceResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
