// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceModifyPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetResourceModifyPriceResponseBodyData) *GetResourceModifyPriceResponseBody
	GetData() *GetResourceModifyPriceResponseBodyData
	SetRequestId(v string) *GetResourceModifyPriceResponseBody
	GetRequestId() *string
}

type GetResourceModifyPriceResponseBody struct {
	// The quote for the resource scale-out of the delivery group, which includes the pricing details and pricing calculation rules.
	Data *GetResourceModifyPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The unique ID of the request, which is used to locate and troubleshoot request issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetResourceModifyPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceModifyPriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceModifyPriceResponseBody) GetData() *GetResourceModifyPriceResponseBodyData {
	return s.Data
}

func (s *GetResourceModifyPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceModifyPriceResponseBody) SetData(v *GetResourceModifyPriceResponseBodyData) *GetResourceModifyPriceResponseBody {
	s.Data = v
	return s
}

func (s *GetResourceModifyPriceResponseBody) SetRequestId(v string) *GetResourceModifyPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceModifyPriceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceModifyPriceResponseBodyData struct {
	// The pricing details of the scale-out. For the currency of the amounts, see Currency. For the coupon selection status, see Promotions.
	Price *GetResourceModifyPriceResponseBodyDataPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// The list of pricing calculation rules for the quote, which is used to describe the applicable pricing or discount rules.
	Rules []*GetResourceModifyPriceResponseBodyDataRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s GetResourceModifyPriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetResourceModifyPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetResourceModifyPriceResponseBodyData) GetPrice() *GetResourceModifyPriceResponseBodyDataPrice {
	return s.Price
}

func (s *GetResourceModifyPriceResponseBodyData) GetRules() []*GetResourceModifyPriceResponseBodyDataRules {
	return s.Rules
}

func (s *GetResourceModifyPriceResponseBodyData) SetPrice(v *GetResourceModifyPriceResponseBodyDataPrice) *GetResourceModifyPriceResponseBodyData {
	s.Price = v
	return s
}

func (s *GetResourceModifyPriceResponseBodyData) SetRules(v []*GetResourceModifyPriceResponseBodyDataRules) *GetResourceModifyPriceResponseBodyData {
	s.Rules = v
	return s
}

func (s *GetResourceModifyPriceResponseBodyData) Validate() error {
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

type GetResourceModifyPriceResponseBodyDataPrice struct {
	// The currency code of the price amount. CNY indicates Chinese Yuan. Refer to this field to understand the amounts of OriginalPrice, DiscountPrice, and TradePrice.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discount amount that can be deducted from the scale-out price, not the discount percentage.
	//
	// The discounted price TradePrice is obtained by subtracting this field from the original price OriginalPrice. For the currency of the amount, see Currency.
	//
	// example:
	//
	// 1.00
	DiscountPrice *float64 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price of the scale-out, which is the amount before discounts. For the currency of the amount, see Currency.
	//
	// example:
	//
	// 11.00
	OriginalPrice *float64 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The list of optional coupons returned in the quote. Each item includes the coupon ID, name, description, and selection status.
	Promotions []*GetResourceModifyPriceResponseBodyDataPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The discounted price of the scale-out.
	//
	// Calculation relationship: TradePrice = OriginalPrice - DiscountPrice. For the currency of the amount, see Currency. This operation returns only a quote and does not execute payment.
	//
	// example:
	//
	// 10.00
	TradePrice *float64 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s GetResourceModifyPriceResponseBodyDataPrice) String() string {
	return dara.Prettify(s)
}

func (s GetResourceModifyPriceResponseBodyDataPrice) GoString() string {
	return s.String()
}

func (s *GetResourceModifyPriceResponseBodyDataPrice) GetCurrency() *string {
	return s.Currency
}

func (s *GetResourceModifyPriceResponseBodyDataPrice) GetDiscountPrice() *float64 {
	return s.DiscountPrice
}

func (s *GetResourceModifyPriceResponseBodyDataPrice) GetOriginalPrice() *float64 {
	return s.OriginalPrice
}

func (s *GetResourceModifyPriceResponseBodyDataPrice) GetPromotions() []*GetResourceModifyPriceResponseBodyDataPricePromotions {
	return s.Promotions
}

func (s *GetResourceModifyPriceResponseBodyDataPrice) GetTradePrice() *float64 {
	return s.TradePrice
}

func (s *GetResourceModifyPriceResponseBodyDataPrice) SetCurrency(v string) *GetResourceModifyPriceResponseBodyDataPrice {
	s.Currency = &v
	return s
}

func (s *GetResourceModifyPriceResponseBodyDataPrice) SetDiscountPrice(v float64) *GetResourceModifyPriceResponseBodyDataPrice {
	s.DiscountPrice = &v
	return s
}

func (s *GetResourceModifyPriceResponseBodyDataPrice) SetOriginalPrice(v float64) *GetResourceModifyPriceResponseBodyDataPrice {
	s.OriginalPrice = &v
	return s
}

func (s *GetResourceModifyPriceResponseBodyDataPrice) SetPromotions(v []*GetResourceModifyPriceResponseBodyDataPricePromotions) *GetResourceModifyPriceResponseBodyDataPrice {
	s.Promotions = v
	return s
}

func (s *GetResourceModifyPriceResponseBodyDataPrice) SetTradePrice(v float64) *GetResourceModifyPriceResponseBodyDataPrice {
	s.TradePrice = &v
	return s
}

func (s *GetResourceModifyPriceResponseBodyDataPrice) Validate() error {
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

type GetResourceModifyPriceResponseBodyDataPricePromotions struct {
	// The option code of the coupon, which identifies the coupon returned in the quote together with PromotionId.
	//
	// example:
	//
	// coupon****
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The description of the coupon, which is used to display the information about the discount.
	//
	// example:
	//
	// **Discount during the promotion
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The ID of the coupon, which is used to identify the coupon returned in the quote.
	//
	// example:
	//
	// 17440009****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The name of the coupon, which is used to display the optional discount in the current quote.
	//
	// example:
	//
	// **Discount
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// Specifies whether the coupon is selected in the quote. Valid values:
	//
	// - true: Selected. The coupon is applied in the quote.
	//
	// - false: Not selected. The coupon is not applied in the quote.
	//
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s GetResourceModifyPriceResponseBodyDataPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s GetResourceModifyPriceResponseBodyDataPricePromotions) GoString() string {
	return s.String()
}

func (s *GetResourceModifyPriceResponseBodyDataPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *GetResourceModifyPriceResponseBodyDataPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *GetResourceModifyPriceResponseBodyDataPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *GetResourceModifyPriceResponseBodyDataPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *GetResourceModifyPriceResponseBodyDataPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *GetResourceModifyPriceResponseBodyDataPricePromotions) SetOptionCode(v string) *GetResourceModifyPriceResponseBodyDataPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *GetResourceModifyPriceResponseBodyDataPricePromotions) SetPromotionDesc(v string) *GetResourceModifyPriceResponseBodyDataPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *GetResourceModifyPriceResponseBodyDataPricePromotions) SetPromotionId(v string) *GetResourceModifyPriceResponseBodyDataPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *GetResourceModifyPriceResponseBodyDataPricePromotions) SetPromotionName(v string) *GetResourceModifyPriceResponseBodyDataPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *GetResourceModifyPriceResponseBodyDataPricePromotions) SetSelected(v bool) *GetResourceModifyPriceResponseBodyDataPricePromotions {
	s.Selected = &v
	return s
}

func (s *GetResourceModifyPriceResponseBodyDataPricePromotions) Validate() error {
	return dara.Validate(s)
}

type GetResourceModifyPriceResponseBodyDataRules struct {
	// The description of the pricing calculation rule, such as the user discount applicable to the current quote.
	//
	// example:
	//
	// User discount
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the pricing calculation rule, which is used to identify the rule applied in the quote.
	//
	// example:
	//
	// 102002100393****
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetResourceModifyPriceResponseBodyDataRules) String() string {
	return dara.Prettify(s)
}

func (s GetResourceModifyPriceResponseBodyDataRules) GoString() string {
	return s.String()
}

func (s *GetResourceModifyPriceResponseBodyDataRules) GetDescription() *string {
	return s.Description
}

func (s *GetResourceModifyPriceResponseBodyDataRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetResourceModifyPriceResponseBodyDataRules) SetDescription(v string) *GetResourceModifyPriceResponseBodyDataRules {
	s.Description = &v
	return s
}

func (s *GetResourceModifyPriceResponseBodyDataRules) SetRuleId(v int64) *GetResourceModifyPriceResponseBodyDataRules {
	s.RuleId = &v
	return s
}

func (s *GetResourceModifyPriceResponseBodyDataRules) Validate() error {
	return dara.Validate(s)
}
