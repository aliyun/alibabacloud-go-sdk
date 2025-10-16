// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceRenewPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetResourceRenewPriceResponseBodyData) *GetResourceRenewPriceResponseBody
	GetData() *GetResourceRenewPriceResponseBodyData
	SetRequestId(v string) *GetResourceRenewPriceResponseBody
	GetRequestId() *string
}

type GetResourceRenewPriceResponseBody struct {
	// The price object.
	Data *GetResourceRenewPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceRenewPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceRenewPriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponseBody) GetData() *GetResourceRenewPriceResponseBodyData {
	return s.Data
}

func (s *GetResourceRenewPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceRenewPriceResponseBody) SetData(v *GetResourceRenewPriceResponseBodyData) *GetResourceRenewPriceResponseBody {
	s.Data = v
	return s
}

func (s *GetResourceRenewPriceResponseBody) SetRequestId(v string) *GetResourceRenewPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceRenewPriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceRenewPriceResponseBodyData struct {
	// The price details.
	Price *GetResourceRenewPriceResponseBodyDataPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// The price calculation rules.
	Rules []*GetResourceRenewPriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s GetResourceRenewPriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetResourceRenewPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponseBodyData) GetPrice() *GetResourceRenewPriceResponseBodyDataPrice {
	return s.Price
}

func (s *GetResourceRenewPriceResponseBodyData) GetRules() []*GetResourceRenewPriceResponseBodyDataRules {
	return s.Rules
}

func (s *GetResourceRenewPriceResponseBodyData) SetPrice(v *GetResourceRenewPriceResponseBodyDataPrice) *GetResourceRenewPriceResponseBodyData {
	s.Price = v
	return s
}

func (s *GetResourceRenewPriceResponseBodyData) SetRules(v []*GetResourceRenewPriceResponseBodyDataRules) *GetResourceRenewPriceResponseBodyData {
	s.Rules = v
	return s
}

func (s *GetResourceRenewPriceResponseBodyData) Validate() error {
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

type GetResourceRenewPriceResponseBodyDataPrice struct {
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
	// 1
	DiscountPrice *string `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 11
	OriginalPrice *string `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The coupon description.
	Promotions []*GetResourceRenewPriceResponseBodyDataPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The actual price. The actual price is calculated based on the following formula: Actual price = Original price - Discount.
	//
	// example:
	//
	// 10
	TradePrice *string `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s GetResourceRenewPriceResponseBodyDataPrice) String() string {
	return dara.Prettify(s)
}

func (s GetResourceRenewPriceResponseBodyDataPrice) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) GetCurrency() *string {
	return s.Currency
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) GetDiscountPrice() *string {
	return s.DiscountPrice
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) GetOriginalPrice() *string {
	return s.OriginalPrice
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) GetPromotions() []*GetResourceRenewPriceResponseBodyDataPricePromotions {
	return s.Promotions
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) GetTradePrice() *string {
	return s.TradePrice
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) SetCurrency(v string) *GetResourceRenewPriceResponseBodyDataPrice {
	s.Currency = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) SetDiscountPrice(v string) *GetResourceRenewPriceResponseBodyDataPrice {
	s.DiscountPrice = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) SetOriginalPrice(v string) *GetResourceRenewPriceResponseBodyDataPrice {
	s.OriginalPrice = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) SetPromotions(v []*GetResourceRenewPriceResponseBodyDataPricePromotions) *GetResourceRenewPriceResponseBodyDataPrice {
	s.Promotions = v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) SetTradePrice(v string) *GetResourceRenewPriceResponseBodyDataPrice {
	s.TradePrice = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPrice) Validate() error {
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

type GetResourceRenewPriceResponseBodyDataPricePromotions struct {
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
	// 139965*****
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

func (s GetResourceRenewPriceResponseBodyDataPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s GetResourceRenewPriceResponseBodyDataPricePromotions) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) SetOptionCode(v string) *GetResourceRenewPriceResponseBodyDataPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) SetPromotionDesc(v string) *GetResourceRenewPriceResponseBodyDataPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) SetPromotionId(v string) *GetResourceRenewPriceResponseBodyDataPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) SetPromotionName(v string) *GetResourceRenewPriceResponseBodyDataPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) SetSelected(v bool) *GetResourceRenewPriceResponseBodyDataPricePromotions {
	s.Selected = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataPricePromotions) Validate() error {
	return dara.Validate(s)
}

type GetResourceRenewPriceResponseBodyDataRules struct {
	// The description of the price calculation rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the price calculation rule.
	//
	// example:
	//
	// 20002****
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetResourceRenewPriceResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s GetResourceRenewPriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponseBodyDataRules) GetDescription() *string {
	return s.Description
}

func (s *GetResourceRenewPriceResponseBodyDataRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetResourceRenewPriceResponseBodyDataRules) SetDescription(v string) *GetResourceRenewPriceResponseBodyDataRules {
	s.Description = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataRules) SetRuleId(v int64) *GetResourceRenewPriceResponseBodyDataRules {
	s.RuleId = &v
	return s
}

func (s *GetResourceRenewPriceResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}
