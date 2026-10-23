// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudPhoneNodeRenewPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPrice(v *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) *DescribeCloudPhoneNodeRenewPriceResponseBody
	GetPrice() *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice
	SetRequestId(v string) *DescribeCloudPhoneNodeRenewPriceResponseBody
	GetRequestId() *string
	SetRules(v []*DescribeCloudPhoneNodeRenewPriceResponseBodyRules) *DescribeCloudPhoneNodeRenewPriceResponseBody
	GetRules() []*DescribeCloudPhoneNodeRenewPriceResponseBodyRules
}

type DescribeCloudPhoneNodeRenewPriceResponseBody struct {
	Price *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// example:
	//
	// 1A923337-44D9-5CAD-9A53-95084BD4****
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     []*DescribeCloudPhoneNodeRenewPriceResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeCloudPhoneNodeRenewPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodeRenewPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBody) GetPrice() *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice {
	return s.Price
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBody) GetRules() []*DescribeCloudPhoneNodeRenewPriceResponseBodyRules {
	return s.Rules
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBody) SetPrice(v *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) *DescribeCloudPhoneNodeRenewPriceResponseBody {
	s.Price = v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBody) SetRequestId(v string) *DescribeCloudPhoneNodeRenewPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBody) SetRules(v []*DescribeCloudPhoneNodeRenewPriceResponseBodyRules) *DescribeCloudPhoneNodeRenewPriceResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBody) Validate() error {
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

type DescribeCloudPhoneNodeRenewPriceResponseBodyPrice struct {
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
	// 1860
	OriginalPrice *float32                                                        `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	PriceDetail   []*DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail `json:"PriceDetail,omitempty" xml:"PriceDetail,omitempty" type:"Repeated"`
	Promotions    []*DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions  `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// example:
	//
	// 769.3
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) GetPriceDetail() []*DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail {
	return s.PriceDetail
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) GetPromotions() []*DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions {
	return s.Promotions
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) SetCurrency(v string) *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice {
	s.Currency = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) SetDiscountPrice(v float32) *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) SetOriginalPrice(v float32) *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) SetPriceDetail(v []*DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice {
	s.PriceDetail = v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) SetPromotions(v []*DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions) *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice {
	s.Promotions = v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) SetTradePrice(v float32) *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPrice) Validate() error {
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

type DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail struct {
	// example:
	//
	// 274.75
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// cpn-9rz9ou1l0wq30****
	InstanceId  *string                                                                    `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModulePrice []*DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice `json:"ModulePrice,omitempty" xml:"ModulePrice,omitempty" type:"Repeated"`
	// example:
	//
	// 2087.0
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// CloudPhoneMatrix
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 1099.0
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) GetModulePrice() []*DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice {
	return s.ModulePrice
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) SetDiscountPrice(v float32) *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) SetInstanceId(v string) *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) SetModulePrice(v []*DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice) *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail {
	s.ModulePrice = v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) SetOriginalPrice(v float32) *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) SetResourceType(v string) *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail {
	s.ResourceType = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) SetTradePrice(v float32) *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail {
	s.TradePrice = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetail) Validate() error {
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

type DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice struct {
	// example:
	//
	// 878.15
	DiscountPrice *string `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// PhoneData
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// example:
	//
	// 1090
	OriginalPrice *string `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// 73.87
	TradePrice *string `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice) GetDiscountPrice() *string {
	return s.DiscountPrice
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice) GetModuleName() *string {
	return s.ModuleName
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice) GetOriginalPrice() *string {
	return s.OriginalPrice
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice) GetTradePrice() *string {
	return s.TradePrice
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice) SetDiscountPrice(v string) *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice) SetModuleName(v string) *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice {
	s.ModuleName = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice) SetOriginalPrice(v string) *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice) SetTradePrice(v string) *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePriceDetailModulePrice) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions struct {
	// example:
	//
	// youhui_quan
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
	Selected *bool `json:"selected,omitempty" xml:"selected,omitempty"`
}

func (s DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions) SetOptionCode(v string) *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions) SetPromotionDesc(v string) *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions) SetPromotionId(v string) *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions) SetPromotionName(v string) *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions) SetSelected(v bool) *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions {
	s.Selected = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyPricePromotions) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudPhoneNodeRenewPriceResponseBodyRules struct {
	// example:
	//
	// 满40减20
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2023110611****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeCloudPhoneNodeRenewPriceResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudPhoneNodeRenewPriceResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyRules) SetDescription(v string) *DescribeCloudPhoneNodeRenewPriceResponseBodyRules {
	s.Description = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyRules) SetRuleId(v string) *DescribeCloudPhoneNodeRenewPriceResponseBodyRules {
	s.RuleId = &v
	return s
}

func (s *DescribeCloudPhoneNodeRenewPriceResponseBodyRules) Validate() error {
	return dara.Validate(s)
}
