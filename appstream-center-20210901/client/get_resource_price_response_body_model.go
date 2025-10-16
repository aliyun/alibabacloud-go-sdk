// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetResourcePriceResponseBody
	GetCode() *string
	SetMessage(v string) *GetResourcePriceResponseBody
	GetMessage() *string
	SetPriceList(v []*GetResourcePriceResponseBodyPriceList) *GetResourcePriceResponseBody
	GetPriceList() []*GetResourcePriceResponseBodyPriceList
	SetPriceModel(v *GetResourcePriceResponseBodyPriceModel) *GetResourcePriceResponseBody
	GetPriceModel() *GetResourcePriceResponseBodyPriceModel
	SetRequestId(v string) *GetResourcePriceResponseBody
	GetRequestId() *string
}

type GetResourcePriceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter.ProductType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The parameter ProductType is invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The price objects.
	//
	// This parameter is returned only if a value is specified for AppInstanceType.
	PriceList []*GetResourcePriceResponseBodyPriceList `json:"PriceList,omitempty" xml:"PriceList,omitempty" type:"Repeated"`
	// The price object.
	//
	// This parameter is returned only if a value is specified for NodeInstanceType.
	PriceModel *GetResourcePriceResponseBodyPriceModel `json:"PriceModel,omitempty" xml:"PriceModel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourcePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetResourcePriceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetResourcePriceResponseBody) GetPriceList() []*GetResourcePriceResponseBodyPriceList {
	return s.PriceList
}

func (s *GetResourcePriceResponseBody) GetPriceModel() *GetResourcePriceResponseBodyPriceModel {
	return s.PriceModel
}

func (s *GetResourcePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourcePriceResponseBody) SetCode(v string) *GetResourcePriceResponseBody {
	s.Code = &v
	return s
}

func (s *GetResourcePriceResponseBody) SetMessage(v string) *GetResourcePriceResponseBody {
	s.Message = &v
	return s
}

func (s *GetResourcePriceResponseBody) SetPriceList(v []*GetResourcePriceResponseBodyPriceList) *GetResourcePriceResponseBody {
	s.PriceList = v
	return s
}

func (s *GetResourcePriceResponseBody) SetPriceModel(v *GetResourcePriceResponseBodyPriceModel) *GetResourcePriceResponseBody {
	s.PriceModel = v
	return s
}

func (s *GetResourcePriceResponseBody) SetRequestId(v string) *GetResourcePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourcePriceResponseBody) Validate() error {
	if s.PriceList != nil {
		for _, item := range s.PriceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PriceModel != nil {
		if err := s.PriceModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourcePriceResponseBodyPriceList struct {
	// The price details.
	Price *GetResourcePriceResponseBodyPriceListPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// The price type.
	//
	// Valid values:
	//
	// 	- Connected: in use
	//
	// 	- Standby: pending for use.
	//
	// example:
	//
	// Standby
	PriceType *string `json:"PriceType,omitempty" xml:"PriceType,omitempty"`
	// The price calculation rules.
	Rules []*GetResourcePriceResponseBodyPriceListRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s GetResourcePriceResponseBodyPriceList) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceList) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceList) GetPrice() *GetResourcePriceResponseBodyPriceListPrice {
	return s.Price
}

func (s *GetResourcePriceResponseBodyPriceList) GetPriceType() *string {
	return s.PriceType
}

func (s *GetResourcePriceResponseBodyPriceList) GetRules() []*GetResourcePriceResponseBodyPriceListRules {
	return s.Rules
}

func (s *GetResourcePriceResponseBodyPriceList) SetPrice(v *GetResourcePriceResponseBodyPriceListPrice) *GetResourcePriceResponseBodyPriceList {
	s.Price = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceList) SetPriceType(v string) *GetResourcePriceResponseBodyPriceList {
	s.PriceType = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceList) SetRules(v []*GetResourcePriceResponseBodyPriceListRules) *GetResourcePriceResponseBodyPriceList {
	s.Rules = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceList) Validate() error {
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

type GetResourcePriceResponseBodyPriceListPrice struct {
	// The currency type.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discount. The actual price is calculated based on the following formula: Actual price = Original price - Discount.
	//
	// example:
	//
	// 999.0
	DiscountPrice *string `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 6700
	OriginalPrice *string `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The coupon metadata.
	Promotions []*GetResourcePriceResponseBodyPriceListPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The actual price. The actual price is calculated based on the following formula: Actual price = Original price - Discount.
	//
	// example:
	//
	// 5278.0
	TradePrice *string `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceListPrice) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceListPrice) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceListPrice) GetCurrency() *string {
	return s.Currency
}

func (s *GetResourcePriceResponseBodyPriceListPrice) GetDiscountPrice() *string {
	return s.DiscountPrice
}

func (s *GetResourcePriceResponseBodyPriceListPrice) GetOriginalPrice() *string {
	return s.OriginalPrice
}

func (s *GetResourcePriceResponseBodyPriceListPrice) GetPromotions() []*GetResourcePriceResponseBodyPriceListPricePromotions {
	return s.Promotions
}

func (s *GetResourcePriceResponseBodyPriceListPrice) GetTradePrice() *string {
	return s.TradePrice
}

func (s *GetResourcePriceResponseBodyPriceListPrice) SetCurrency(v string) *GetResourcePriceResponseBodyPriceListPrice {
	s.Currency = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPrice) SetDiscountPrice(v string) *GetResourcePriceResponseBodyPriceListPrice {
	s.DiscountPrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPrice) SetOriginalPrice(v string) *GetResourcePriceResponseBodyPriceListPrice {
	s.OriginalPrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPrice) SetPromotions(v []*GetResourcePriceResponseBodyPriceListPricePromotions) *GetResourcePriceResponseBodyPriceListPrice {
	s.Promotions = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPrice) SetTradePrice(v string) *GetResourcePriceResponseBodyPriceListPrice {
	s.TradePrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPrice) Validate() error {
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

type GetResourcePriceResponseBodyPriceListPricePromotions struct {
	// The coupon code.
	//
	// example:
	//
	// coupon****
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The coupon description.
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// 1847709****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The coupon name.
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// Indicates whether the coupon was used.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceListPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceListPricePromotions) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) SetOptionCode(v string) *GetResourcePriceResponseBodyPriceListPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) SetPromotionDesc(v string) *GetResourcePriceResponseBodyPriceListPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) SetPromotionId(v string) *GetResourcePriceResponseBodyPriceListPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) SetPromotionName(v string) *GetResourcePriceResponseBodyPriceListPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) SetSelected(v bool) *GetResourcePriceResponseBodyPriceListPricePromotions {
	s.Selected = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListPricePromotions) Validate() error {
	return dara.Validate(s)
}

type GetResourcePriceResponseBodyPriceListRules struct {
	// The description of the price calculation rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the price calculation rule.
	//
	// example:
	//
	// 260904273633****
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceListRules) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceListRules) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceListRules) GetDescription() *string {
	return s.Description
}

func (s *GetResourcePriceResponseBodyPriceListRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetResourcePriceResponseBodyPriceListRules) SetDescription(v string) *GetResourcePriceResponseBodyPriceListRules {
	s.Description = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListRules) SetRuleId(v int64) *GetResourcePriceResponseBodyPriceListRules {
	s.RuleId = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceListRules) Validate() error {
	return dara.Validate(s)
}

type GetResourcePriceResponseBodyPriceModel struct {
	// The price details.
	Price *GetResourcePriceResponseBodyPriceModelPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// The price calculation rules.
	Rules []*GetResourcePriceResponseBodyPriceModelRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s GetResourcePriceResponseBodyPriceModel) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModel) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModel) GetPrice() *GetResourcePriceResponseBodyPriceModelPrice {
	return s.Price
}

func (s *GetResourcePriceResponseBodyPriceModel) GetRules() []*GetResourcePriceResponseBodyPriceModelRules {
	return s.Rules
}

func (s *GetResourcePriceResponseBodyPriceModel) SetPrice(v *GetResourcePriceResponseBodyPriceModelPrice) *GetResourcePriceResponseBodyPriceModel {
	s.Price = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModel) SetRules(v []*GetResourcePriceResponseBodyPriceModelRules) *GetResourcePriceResponseBodyPriceModel {
	s.Rules = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModel) Validate() error {
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

type GetResourcePriceResponseBodyPriceModelPrice struct {
	// The currency type.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discount. The actual price is calculated based on the following formula: Actual price = Original price - Discount.
	//
	// example:
	//
	// 1.00
	DiscountPrice *string `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 11.00
	OriginalPrice *string `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The coupon metadata.
	Promotions []*GetResourcePriceResponseBodyPriceModelPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The actual price. The actual price is calculated based on the following formula: Actual price = Original price - Discount.
	//
	// example:
	//
	// 10.00
	TradePrice *string `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceModelPrice) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModelPrice) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) GetCurrency() *string {
	return s.Currency
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) GetDiscountPrice() *string {
	return s.DiscountPrice
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) GetOriginalPrice() *string {
	return s.OriginalPrice
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) GetPromotions() []*GetResourcePriceResponseBodyPriceModelPricePromotions {
	return s.Promotions
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) GetTradePrice() *string {
	return s.TradePrice
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetCurrency(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.Currency = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetDiscountPrice(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.DiscountPrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetOriginalPrice(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.OriginalPrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetPromotions(v []*GetResourcePriceResponseBodyPriceModelPricePromotions) *GetResourcePriceResponseBodyPriceModelPrice {
	s.Promotions = v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) SetTradePrice(v string) *GetResourcePriceResponseBodyPriceModelPrice {
	s.TradePrice = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPrice) Validate() error {
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

type GetResourcePriceResponseBodyPriceModelPricePromotions struct {
	// The coupon code.
	//
	// example:
	//
	// coupon****
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The coupon description.
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The coupon ID.
	//
	// example:
	//
	// 17440009****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The coupon name.
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// Indicates whether the coupon was used.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceModelPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModelPricePromotions) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetOptionCode(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetPromotionDesc(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetPromotionId(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetPromotionName(v string) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) SetSelected(v bool) *GetResourcePriceResponseBodyPriceModelPricePromotions {
	s.Selected = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelPricePromotions) Validate() error {
	return dara.Validate(s)
}

type GetResourcePriceResponseBodyPriceModelRules struct {
	// The description of the price calculation rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the price calculation rule.
	//
	// example:
	//
	// 102002100393****
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetResourcePriceResponseBodyPriceModelRules) String() string {
	return dara.Prettify(s)
}

func (s GetResourcePriceResponseBodyPriceModelRules) GoString() string {
	return s.String()
}

func (s *GetResourcePriceResponseBodyPriceModelRules) GetDescription() *string {
	return s.Description
}

func (s *GetResourcePriceResponseBodyPriceModelRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetResourcePriceResponseBodyPriceModelRules) SetDescription(v string) *GetResourcePriceResponseBodyPriceModelRules {
	s.Description = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelRules) SetRuleId(v int64) *GetResourcePriceResponseBodyPriceModelRules {
	s.RuleId = &v
	return s
}

func (s *GetResourcePriceResponseBodyPriceModelRules) Validate() error {
	return dara.Validate(s)
}
