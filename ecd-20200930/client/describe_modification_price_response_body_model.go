// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModificationPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *DescribeModificationPriceResponseBodyPriceInfo) *DescribeModificationPriceResponseBody
	GetPriceInfo() *DescribeModificationPriceResponseBodyPriceInfo
	SetRequestId(v string) *DescribeModificationPriceResponseBody
	GetRequestId() *string
}

type DescribeModificationPriceResponseBody struct {
	// The price details.
	PriceInfo *DescribeModificationPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 48174475-5EB2-5F99-A9E9-6F892D645****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeModificationPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModificationPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponseBody) GetPriceInfo() *DescribeModificationPriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *DescribeModificationPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModificationPriceResponseBody) SetPriceInfo(v *DescribeModificationPriceResponseBodyPriceInfo) *DescribeModificationPriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribeModificationPriceResponseBody) SetRequestId(v string) *DescribeModificationPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModificationPriceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeModificationPriceResponseBodyPriceInfo struct {
	// The price information.
	Price *DescribeModificationPriceResponseBodyPriceInfoPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// The promotion rules.
	Rules []*DescribeModificationPriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeModificationPriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeModificationPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponseBodyPriceInfo) GetPrice() *DescribeModificationPriceResponseBodyPriceInfoPrice {
	return s.Price
}

func (s *DescribeModificationPriceResponseBodyPriceInfo) GetRules() []*DescribeModificationPriceResponseBodyPriceInfoRules {
	return s.Rules
}

func (s *DescribeModificationPriceResponseBodyPriceInfo) SetPrice(v *DescribeModificationPriceResponseBodyPriceInfoPrice) *DescribeModificationPriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfo) SetRules(v []*DescribeModificationPriceResponseBodyPriceInfoRules) *DescribeModificationPriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeModificationPriceResponseBodyPriceInfoPrice struct {
	// The unit of currency (USD).
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discounted amount.
	//
	// example:
	//
	// 15.8
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The orders.
	OrderLines map[string]*string `json:"OrderLines,omitempty" xml:"OrderLines,omitempty"`
	// The original price.
	//
	// example:
	//
	// 79.0
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The promotion activities.
	Promotions []*DescribeModificationPriceResponseBodyPriceInfoPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The actual price. The actual price is the original price minus the discount.
	//
	// example:
	//
	// 63.2
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeModificationPriceResponseBodyPriceInfoPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribeModificationPriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) GetOrderLines() map[string]*string {
	return s.OrderLines
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) GetPromotions() []*DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	return s.Promotions
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetOrderLines(v map[string]*string) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.OrderLines = v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetPromotions(v []*DescribeModificationPriceResponseBodyPriceInfoPricePromotions) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.Promotions = v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribeModificationPriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPrice) Validate() error {
	return dara.Validate(s)
}

type DescribeModificationPriceResponseBodyPriceInfoPricePromotions struct {
	// The description of the promotion rule.
	//
	// example:
	//
	// test
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The promotion description.
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The ID of the promotion activity.
	//
	// example:
	//
	// promo_option
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The name of the promotion activity.
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// Indicates whether an item is selected.
	//
	// example:
	//
	// false
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s DescribeModificationPriceResponseBodyPriceInfoPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribeModificationPriceResponseBodyPriceInfoPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) SetOptionCode(v string) *DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) SetPromotionDesc(v string) *DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) SetPromotionId(v string) *DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) SetPromotionName(v string) *DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) SetSelected(v bool) *DescribeModificationPriceResponseBodyPriceInfoPricePromotions {
	s.Selected = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoPricePromotions) Validate() error {
	return dara.Validate(s)
}

type DescribeModificationPriceResponseBodyPriceInfoRules struct {
	// The rule description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 14806
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeModificationPriceResponseBodyPriceInfoRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeModificationPriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribeModificationPriceResponseBodyPriceInfoRules) GetDescription() *string {
	return s.Description
}

func (s *DescribeModificationPriceResponseBodyPriceInfoRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeModificationPriceResponseBodyPriceInfoRules) SetDescription(v string) *DescribeModificationPriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoRules) SetRuleId(v int64) *DescribeModificationPriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

func (s *DescribeModificationPriceResponseBodyPriceInfoRules) Validate() error {
	return dara.Validate(s)
}
