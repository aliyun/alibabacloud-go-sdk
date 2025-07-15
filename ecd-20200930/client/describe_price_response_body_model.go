// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *DescribePriceResponseBodyPriceInfo) *DescribePriceResponseBody
	GetPriceInfo() *DescribePriceResponseBodyPriceInfo
	SetRequestId(v string) *DescribePriceResponseBody
	GetRequestId() *string
}

type DescribePriceResponseBody struct {
	// The price details.
	PriceInfo *DescribePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B1175630-3C44-4389-A3C1-15639FFC8EBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) GetPriceInfo() *DescribePriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *DescribePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePriceResponseBody) SetPriceInfo(v *DescribePriceResponseBodyPriceInfo) *DescribePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribePriceResponseBody) SetRequestId(v string) *DescribePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePriceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPriceInfo struct {
	// Indicates whether a free enterprise drive is available.
	//
	// example:
	//
	// true
	FreeCdsQuota *bool `json:"FreeCdsQuota,omitempty" xml:"FreeCdsQuota,omitempty"`
	// The free capacity provided by the enterprise drive. Unit: GiB.
	//
	// example:
	//
	// 100
	FreeCdsSize *int64 `json:"FreeCdsSize,omitempty" xml:"FreeCdsSize,omitempty"`
	// The price.
	Price *DescribePriceResponseBodyPriceInfoPrice `json:"Price,omitempty" xml:"Price,omitempty" type:"Struct"`
	// The details of the promotion rules.
	Rules []*DescribePriceResponseBodyPriceInfoRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfo) GetFreeCdsQuota() *bool {
	return s.FreeCdsQuota
}

func (s *DescribePriceResponseBodyPriceInfo) GetFreeCdsSize() *int64 {
	return s.FreeCdsSize
}

func (s *DescribePriceResponseBodyPriceInfo) GetPrice() *DescribePriceResponseBodyPriceInfoPrice {
	return s.Price
}

func (s *DescribePriceResponseBodyPriceInfo) GetRules() []*DescribePriceResponseBodyPriceInfoRules {
	return s.Rules
}

func (s *DescribePriceResponseBodyPriceInfo) SetFreeCdsQuota(v bool) *DescribePriceResponseBodyPriceInfo {
	s.FreeCdsQuota = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetFreeCdsSize(v int64) *DescribePriceResponseBodyPriceInfo {
	s.FreeCdsSize = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetPrice(v *DescribePriceResponseBodyPriceInfoPrice) *DescribePriceResponseBodyPriceInfo {
	s.Price = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetRules(v []*DescribePriceResponseBodyPriceInfoRules) *DescribePriceResponseBodyPriceInfo {
	s.Rules = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPriceInfoPrice struct {
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
	// 0
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The orders.
	OrderLines map[string]*string `json:"OrderLines,omitempty" xml:"OrderLines,omitempty"`
	// The original price.
	//
	// example:
	//
	// 2.796
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The promotions.
	Promotions []*DescribePriceResponseBodyPriceInfoPricePromotions `json:"Promotions,omitempty" xml:"Promotions,omitempty" type:"Repeated"`
	// The price under an effective savings plan.
	//
	// example:
	//
	// 50.00
	SpPrice *int64 `json:"SpPrice,omitempty" xml:"SpPrice,omitempty"`
	// The actual price. The original price minus the discounted amount equals the actual price.
	//
	// example:
	//
	// 2.796
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPrice) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetCurrency() *string {
	return s.Currency
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetOrderLines() map[string]*string {
	return s.OrderLines
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetPromotions() []*DescribePriceResponseBodyPriceInfoPricePromotions {
	return s.Promotions
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetSpPrice() *int64 {
	return s.SpPrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetCurrency(v string) *DescribePriceResponseBodyPriceInfoPrice {
	s.Currency = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetDiscountPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetOrderLines(v map[string]*string) *DescribePriceResponseBodyPriceInfoPrice {
	s.OrderLines = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetOriginalPrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetPromotions(v []*DescribePriceResponseBodyPriceInfoPricePromotions) *DescribePriceResponseBodyPriceInfoPrice {
	s.Promotions = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetSpPrice(v int64) *DescribePriceResponseBodyPriceInfoPrice {
	s.SpPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) SetTradePrice(v float32) *DescribePriceResponseBodyPriceInfoPrice {
	s.TradePrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPrice) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPriceInfoPricePromotions struct {
	// The description of the promotion rule.
	//
	// example:
	//
	// test
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The description of the promotion.
	//
	// example:
	//
	// Get started with new services with a discount.
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// 123456
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The promotion name.
	//
	// example:
	//
	// Special Offer
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// Indicates whether an item is selected.
	//
	// example:
	//
	// false
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoPricePromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoPricePromotions) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) GetPromotionId() *string {
	return s.PromotionId
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) GetSelected() *bool {
	return s.Selected
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetOptionCode(v string) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.OptionCode = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetPromotionDesc(v string) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.PromotionDesc = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetPromotionId(v string) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.PromotionId = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetPromotionName(v string) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.PromotionName = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) SetSelected(v bool) *DescribePriceResponseBodyPriceInfoPricePromotions {
	s.Selected = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoPricePromotions) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPriceInfoRules struct {
	// The description of the rule.
	//
	// example:
	//
	// Receive a 15% discount on a one-year subscription.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 587
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoRules) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoRules) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoRules) GetDescription() *string {
	return s.Description
}

func (s *DescribePriceResponseBodyPriceInfoRules) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribePriceResponseBodyPriceInfoRules) SetDescription(v string) *DescribePriceResponseBodyPriceInfoRules {
	s.Description = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoRules) SetRuleId(v int64) *DescribePriceResponseBodyPriceInfoRules {
	s.RuleId = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoRules) Validate() error {
	return dara.Validate(s)
}
