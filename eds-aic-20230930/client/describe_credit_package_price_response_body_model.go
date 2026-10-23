// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditPackagePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrice(v *DescribeCreditPackagePriceResponseBodyPrice) *DescribeCreditPackagePriceResponseBody
	GetPrice() *DescribeCreditPackagePriceResponseBodyPrice
	SetRequestId(v string) *DescribeCreditPackagePriceResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeCreditPackagePriceResponseBodyRules) *DescribeCreditPackagePriceResponseBody
	GetRules() []*DescribeCreditPackagePriceResponseBodyRules
}

type DescribeCreditPackagePriceResponseBody struct {
	Price *DescribeCreditPackagePriceResponseBodyPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// example:
	//
	// 6C8439B9-7DBF-57F4-92AE-55A9B9D3****
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*DescribeCreditPackagePriceResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeCreditPackagePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackagePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackagePriceResponseBody) GetPrice() *DescribeCreditPackagePriceResponseBodyPrice {
	return s.Price
}

func (s *DescribeCreditPackagePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCreditPackagePriceResponseBody) GetRules() []*DescribeCreditPackagePriceResponseBodyRules {
	return s.Rules
}

func (s *DescribeCreditPackagePriceResponseBody) SetPrice(v *DescribeCreditPackagePriceResponseBodyPrice) *DescribeCreditPackagePriceResponseBody {
	s.Price = v
	return s
}

func (s *DescribeCreditPackagePriceResponseBody) SetRequestId(v string) *DescribeCreditPackagePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBody) SetRules(v []*DescribeCreditPackagePriceResponseBodyRules) *DescribeCreditPackagePriceResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeCreditPackagePriceResponseBody) Validate() error {
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

type DescribeCreditPackagePriceResponseBodyPrice struct {
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 800
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// 1000
	OriginalPrice *float32                                                  `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	PriceDetail   []*DescribeCreditPackagePriceResponseBodyPricePriceDetail `json:"PriceDetail,omitempty" xml:"PriceDetail,omitempty" type:"Repeated"`
	Promotions    []*DescribeCreditPackagePriceResponseBodyPricePromotions  `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeCreditPackagePriceResponseBodyPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackagePriceResponseBodyPrice) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackagePriceResponseBodyPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeCreditPackagePriceResponseBodyPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeCreditPackagePriceResponseBodyPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeCreditPackagePriceResponseBodyPrice) GetPriceDetail() []*DescribeCreditPackagePriceResponseBodyPricePriceDetail {
	return s.PriceDetail
}

func (s *DescribeCreditPackagePriceResponseBodyPrice) GetPromotions() []*DescribeCreditPackagePriceResponseBodyPricePromotions {
	return s.Promotions
}

func (s *DescribeCreditPackagePriceResponseBodyPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeCreditPackagePriceResponseBodyPrice) SetCurrency(v string) *DescribeCreditPackagePriceResponseBodyPrice {
	s.Currency = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPrice) SetDiscountPrice(v float32) *DescribeCreditPackagePriceResponseBodyPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPrice) SetOriginalPrice(v float32) *DescribeCreditPackagePriceResponseBodyPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPrice) SetPriceDetail(v []*DescribeCreditPackagePriceResponseBodyPricePriceDetail) *DescribeCreditPackagePriceResponseBodyPrice {
	s.PriceDetail = v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPrice) SetPromotions(v []*DescribeCreditPackagePriceResponseBodyPricePromotions) *DescribeCreditPackagePriceResponseBodyPrice {
	s.Promotions = v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPrice) SetTradePrice(v float32) *DescribeCreditPackagePriceResponseBodyPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPrice) Validate() error {
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

type DescribeCreditPackagePriceResponseBodyPricePriceDetail struct {
	// example:
	//
	// 800
	DiscountPrice *float32                                                             `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	ModulePrice   []*DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice `json:"ModulePrice,omitempty" xml:"ModulePrice,omitempty" type:"Repeated"`
	// example:
	//
	// 1000
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// CloudPhoneAIPackage
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 200
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeCreditPackagePriceResponseBodyPricePriceDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackagePriceResponseBodyPricePriceDetail) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetail) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetail) GetModulePrice() []*DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice {
	return s.ModulePrice
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetail) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetail) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetail) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetail) SetDiscountPrice(v float32) *DescribeCreditPackagePriceResponseBodyPricePriceDetail {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetail) SetModulePrice(v []*DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice) *DescribeCreditPackagePriceResponseBodyPricePriceDetail {
	s.ModulePrice = v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetail) SetOriginalPrice(v float32) *DescribeCreditPackagePriceResponseBodyPricePriceDetail {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetail) SetResourceType(v string) *DescribeCreditPackagePriceResponseBodyPricePriceDetail {
	s.ResourceType = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetail) SetTradePrice(v float32) *DescribeCreditPackagePriceResponseBodyPricePriceDetail {
	s.TradePrice = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetail) Validate() error {
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

type DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice struct {
	// example:
	//
	// 800
	DiscountPrice *string `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// 1000积分
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// example:
	//
	// 1000
	OriginalPrice *string `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// 200
	TradePrice *string `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice) GetDiscountPrice() *string {
	return s.DiscountPrice
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice) GetModuleName() *string {
	return s.ModuleName
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice) GetOriginalPrice() *string {
	return s.OriginalPrice
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice) GetTradePrice() *string {
	return s.TradePrice
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice) SetDiscountPrice(v string) *DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice) SetModuleName(v string) *DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice {
	s.ModuleName = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice) SetOriginalPrice(v string) *DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice) SetTradePrice(v string) *DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPricePriceDetailModulePrice) Validate() error {
	return dara.Validate(s)
}

type DescribeCreditPackagePriceResponseBodyPricePromotions struct {
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
	// 50003308011****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// example:
	//
	// 日常优惠
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s DescribeCreditPackagePriceResponseBodyPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackagePriceResponseBodyPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackagePriceResponseBodyPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribeCreditPackagePriceResponseBodyPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *DescribeCreditPackagePriceResponseBodyPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeCreditPackagePriceResponseBodyPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribeCreditPackagePriceResponseBodyPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *DescribeCreditPackagePriceResponseBodyPricePromotions) SetOptionCode(v string) *DescribeCreditPackagePriceResponseBodyPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPricePromotions) SetPromotionDesc(v string) *DescribeCreditPackagePriceResponseBodyPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPricePromotions) SetPromotionId(v string) *DescribeCreditPackagePriceResponseBodyPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPricePromotions) SetPromotionName(v string) *DescribeCreditPackagePriceResponseBodyPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPricePromotions) SetSelected(v bool) *DescribeCreditPackagePriceResponseBodyPricePromotions {
	s.Selected = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyPricePromotions) Validate() error {
	return dara.Validate(s)
}

type DescribeCreditPackagePriceResponseBodyRules struct {
	// example:
	//
	// 日常优惠
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 3s2f42fis7pmb9****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeCreditPackagePriceResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditPackagePriceResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeCreditPackagePriceResponseBodyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeCreditPackagePriceResponseBodyRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeCreditPackagePriceResponseBodyRules) SetDescription(v string) *DescribeCreditPackagePriceResponseBodyRules {
	s.Description = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyRules) SetRuleId(v string) *DescribeCreditPackagePriceResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeCreditPackagePriceResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
