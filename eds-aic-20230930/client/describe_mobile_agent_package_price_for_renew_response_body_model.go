// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMobileAgentPackagePriceForRenewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrice(v *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) *DescribeMobileAgentPackagePriceForRenewResponseBody
	GetPrice() *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice
	SetRequestId(v string) *DescribeMobileAgentPackagePriceForRenewResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeMobileAgentPackagePriceForRenewResponseBodyRules) *DescribeMobileAgentPackagePriceForRenewResponseBody
	GetRules() []*DescribeMobileAgentPackagePriceForRenewResponseBodyRules
}

type DescribeMobileAgentPackagePriceForRenewResponseBody struct {
	Price *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// example:
	//
	// 425F351C-3F8E-5218-A520-B6311D0D****
	RequestId *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*DescribeMobileAgentPackagePriceForRenewResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeMobileAgentPackagePriceForRenewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceForRenewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBody) GetPrice() *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice {
	return s.Price
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBody) GetRules() []*DescribeMobileAgentPackagePriceForRenewResponseBodyRules {
	return s.Rules
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBody) SetPrice(v *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) *DescribeMobileAgentPackagePriceForRenewResponseBody {
	s.Price = v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBody) SetRequestId(v string) *DescribeMobileAgentPackagePriceForRenewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBody) SetRules(v []*DescribeMobileAgentPackagePriceForRenewResponseBodyRules) *DescribeMobileAgentPackagePriceForRenewResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBody) Validate() error {
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

type DescribeMobileAgentPackagePriceForRenewResponseBodyPrice struct {
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// 200
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// 1000
	OriginalPrice *float32                                                               `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	PriceDetail   []*DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail `json:"PriceDetail,omitempty" xml:"PriceDetail,omitempty" type:"Repeated"`
	Promotions    []*DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions  `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// example:
	//
	// 800
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) GetPriceDetail() []*DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail {
	return s.PriceDetail
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) GetPromotions() []*DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions {
	return s.Promotions
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) SetCurrency(v string) *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice {
	s.Currency = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) SetDiscountPrice(v float32) *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) SetOriginalPrice(v float32) *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) SetPriceDetail(v []*DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice {
	s.PriceDetail = v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) SetPromotions(v []*DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions) *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice {
	s.Promotions = v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) SetTradePrice(v float32) *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPrice) Validate() error {
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

type DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail struct {
	// example:
	//
	// 200
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// cmag-shdbcudbaib****
	InstanceId  *string                                                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModulePrice []*DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice `json:"ModulePrice,omitempty" xml:"ModulePrice,omitempty" type:"Repeated"`
	// example:
	//
	// 1000
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// CloudPhoneMobileAgent
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 800
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) GetModulePrice() []*DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice {
	return s.ModulePrice
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) SetDiscountPrice(v float32) *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) SetInstanceId(v string) *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail {
	s.InstanceId = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) SetModulePrice(v []*DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice) *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail {
	s.ModulePrice = v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) SetOriginalPrice(v float32) *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) SetResourceType(v string) *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail {
	s.ResourceType = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) SetTradePrice(v float32) *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail {
	s.TradePrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetail) Validate() error {
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

type DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice struct {
	// example:
	//
	// 629.7
	DiscountPrice *string `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// 高级版
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// example:
	//
	// 1680
	OriginalPrice *string `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// 94.85
	TradePrice *string `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice) GetDiscountPrice() *string {
	return s.DiscountPrice
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice) GetModuleName() *string {
	return s.ModuleName
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice) GetOriginalPrice() *string {
	return s.OriginalPrice
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice) GetTradePrice() *string {
	return s.TradePrice
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice) SetDiscountPrice(v string) *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice) SetModuleName(v string) *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice {
	s.ModuleName = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice) SetOriginalPrice(v string) *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice) SetTradePrice(v string) *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePriceDetailModulePrice) Validate() error {
	return dara.Validate(s)
}

type DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions struct {
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
	// 91381816****
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

func (s DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions) SetOptionCode(v string) *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions) SetPromotionDesc(v string) *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions) SetPromotionId(v string) *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions) SetPromotionName(v string) *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions) SetSelected(v bool) *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions {
	s.Selected = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyPricePromotions) Validate() error {
	return dara.Validate(s)
}

type DescribeMobileAgentPackagePriceForRenewResponseBodyRules struct {
	// example:
	//
	// 日常优惠
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 20230629****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeMobileAgentPackagePriceForRenewResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackagePriceForRenewResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyRules) SetDescription(v string) *DescribeMobileAgentPackagePriceForRenewResponseBodyRules {
	s.Description = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyRules) SetRuleId(v string) *DescribeMobileAgentPackagePriceForRenewResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeMobileAgentPackagePriceForRenewResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
