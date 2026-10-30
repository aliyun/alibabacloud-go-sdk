// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMobileAgentPackagePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrice(v *DescribeMobileAgentPackagePriceResponseBodyPrice) *DescribeMobileAgentPackagePriceResponseBody
	GetPrice() *DescribeMobileAgentPackagePriceResponseBodyPrice
	SetRequestId(v string) *DescribeMobileAgentPackagePriceResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeMobileAgentPackagePriceResponseBodyRules) *DescribeMobileAgentPackagePriceResponseBody
	GetRules() []*DescribeMobileAgentPackagePriceResponseBodyRules
}

type DescribeMobileAgentPackagePriceResponseBody struct {
	// The price.
	Price *DescribeMobileAgentPackagePriceResponseBodyPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The promotion rules.
	Rules []*DescribeMobileAgentPackagePriceResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeMobileAgentPackagePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceResponseBody) GetPrice() *DescribeMobileAgentPackagePriceResponseBodyPrice {
	return s.Price
}

func (s *DescribeMobileAgentPackagePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMobileAgentPackagePriceResponseBody) GetRules() []*DescribeMobileAgentPackagePriceResponseBodyRules {
	return s.Rules
}

func (s *DescribeMobileAgentPackagePriceResponseBody) SetPrice(v *DescribeMobileAgentPackagePriceResponseBodyPrice) *DescribeMobileAgentPackagePriceResponseBody {
	s.Price = v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBody) SetRequestId(v string) *DescribeMobileAgentPackagePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBody) SetRules(v []*DescribeMobileAgentPackagePriceResponseBodyRules) *DescribeMobileAgentPackagePriceResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBody) Validate() error {
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

type DescribeMobileAgentPackagePriceResponseBodyPrice struct {
	// The currency.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discounted price.
	//
	// example:
	//
	// 200
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 1000
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The price details.
	PriceDetail []*DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail `json:"PriceDetail,omitempty" xml:"PriceDetail,omitempty" type:"Repeated"`
	// The promotion information.
	Promotions []*DescribeMobileAgentPackagePriceResponseBodyPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The transaction price.
	//
	// example:
	//
	// 800
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeMobileAgentPackagePriceResponseBodyPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceResponseBodyPrice) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPrice) GetPriceDetail() []*DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail {
	return s.PriceDetail
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPrice) GetPromotions() []*DescribeMobileAgentPackagePriceResponseBodyPricePromotions {
	return s.Promotions
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPrice) SetCurrency(v string) *DescribeMobileAgentPackagePriceResponseBodyPrice {
	s.Currency = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPrice) SetDiscountPrice(v float32) *DescribeMobileAgentPackagePriceResponseBodyPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPrice) SetOriginalPrice(v float32) *DescribeMobileAgentPackagePriceResponseBodyPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPrice) SetPriceDetail(v []*DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail) *DescribeMobileAgentPackagePriceResponseBodyPrice {
	s.PriceDetail = v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPrice) SetPromotions(v []*DescribeMobileAgentPackagePriceResponseBodyPricePromotions) *DescribeMobileAgentPackagePriceResponseBodyPrice {
	s.Promotions = v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPrice) SetTradePrice(v float32) *DescribeMobileAgentPackagePriceResponseBodyPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPrice) Validate() error {
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

type DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail struct {
	// The discounted price.
	//
	// example:
	//
	// 800
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The prices of the modules.
	ModulePrice []*DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice `json:"ModulePrice,omitempty" xml:"ModulePrice,omitempty" type:"Repeated"`
	// The original price.
	//
	// example:
	//
	// 1000
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The resource type.
	//
	// example:
	//
	// CloudPhoneMobileAgent
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The transaction price.
	//
	// example:
	//
	// 200
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail) GetModulePrice() []*DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice {
	return s.ModulePrice
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail) SetDiscountPrice(v float32) *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail) SetModulePrice(v []*DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice) *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail {
	s.ModulePrice = v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail) SetOriginalPrice(v float32) *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail) SetResourceType(v string) *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail {
	s.ResourceType = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail) SetTradePrice(v float32) *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail {
	s.TradePrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetail) Validate() error {
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

type DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice struct {
	// The discounted price.
	//
	// example:
	//
	// 800
	DiscountPrice *string `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The name of the module.
	//
	// example:
	//
	// Advanced Edition
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The original price.
	//
	// example:
	//
	// 1000
	OriginalPrice *string `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The transaction price.
	//
	// example:
	//
	// 200
	TradePrice *string `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice) GetDiscountPrice() *string {
	return s.DiscountPrice
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice) GetModuleName() *string {
	return s.ModuleName
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice) GetOriginalPrice() *string {
	return s.OriginalPrice
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice) GetTradePrice() *string {
	return s.TradePrice
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice) SetDiscountPrice(v string) *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice) SetModuleName(v string) *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice {
	s.ModuleName = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice) SetOriginalPrice(v string) *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice) SetTradePrice(v string) *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePriceDetailModulePrice) Validate() error {
	return dara.Validate(s)
}

type DescribeMobileAgentPackagePriceResponseBodyPricePromotions struct {
	// The coupon code.
	//
	// example:
	//
	// daily
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The coupon description.
	//
	// example:
	//
	// Regular discount
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The ID of the promotion.
	//
	// example:
	//
	// 90712091****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The name of the promotion.
	//
	// example:
	//
	// Regular discount
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// Indicates whether the promotion is selected.
	//
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s DescribeMobileAgentPackagePriceResponseBodyPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceResponseBodyPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePromotions) SetOptionCode(v string) *DescribeMobileAgentPackagePriceResponseBodyPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePromotions) SetPromotionDesc(v string) *DescribeMobileAgentPackagePriceResponseBodyPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePromotions) SetPromotionId(v string) *DescribeMobileAgentPackagePriceResponseBodyPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePromotions) SetPromotionName(v string) *DescribeMobileAgentPackagePriceResponseBodyPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePromotions) SetSelected(v bool) *DescribeMobileAgentPackagePriceResponseBodyPricePromotions {
	s.Selected = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyPricePromotions) Validate() error {
	return dara.Validate(s)
}

type DescribeMobileAgentPackagePriceResponseBodyRules struct {
	// The rule description.
	//
	// example:
	//
	// Regular discount
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// f07tkae80jbiy****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeMobileAgentPackagePriceResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceResponseBodyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeMobileAgentPackagePriceResponseBodyRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeMobileAgentPackagePriceResponseBodyRules) SetDescription(v string) *DescribeMobileAgentPackagePriceResponseBodyRules {
	s.Description = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyRules) SetRuleId(v string) *DescribeMobileAgentPackagePriceResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
