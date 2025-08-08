// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCoupons(v *DescribePriceResponseBodyCoupons) *DescribePriceResponseBody
	GetCoupons() *DescribePriceResponseBodyCoupons
	SetCurrency(v string) *DescribePriceResponseBody
	GetCurrency() *string
	SetCuxiao(v bool) *DescribePriceResponseBody
	GetCuxiao() *bool
	SetCycle(v string) *DescribePriceResponseBody
	GetCycle() *string
	SetDiscountPrice(v float32) *DescribePriceResponseBody
	GetDiscountPrice() *float32
	SetDuration(v int32) *DescribePriceResponseBody
	GetDuration() *int32
	SetExpressionCode(v string) *DescribePriceResponseBody
	GetExpressionCode() *string
	SetExpressionMessage(v string) *DescribePriceResponseBody
	GetExpressionMessage() *string
	SetInfoTitle(v string) *DescribePriceResponseBody
	GetInfoTitle() *string
	SetOriginalPrice(v float32) *DescribePriceResponseBody
	GetOriginalPrice() *float32
	SetProductCode(v string) *DescribePriceResponseBody
	GetProductCode() *string
	SetPromotionRules(v *DescribePriceResponseBodyPromotionRules) *DescribePriceResponseBody
	GetPromotionRules() *DescribePriceResponseBodyPromotionRules
	SetTradePrice(v float32) *DescribePriceResponseBody
	GetTradePrice() *float32
}

type DescribePriceResponseBody struct {
	Coupons *DescribePriceResponseBodyCoupons `json:"Coupons,omitempty" xml:"Coupons,omitempty" type:"Struct"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// example:
	//
	// true
	Cuxiao *bool `json:"Cuxiao,omitempty" xml:"Cuxiao,omitempty"`
	// example:
	//
	// MONTH
	Cycle *string `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// example:
	//
	// 178.20
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// ORDER.NO_REAL_NAME_AUTHENTICATION
	ExpressionCode    *string `json:"ExpressionCode,omitempty" xml:"ExpressionCode,omitempty"`
	ExpressionMessage *string `json:"ExpressionMessage,omitempty" xml:"ExpressionMessage,omitempty"`
	InfoTitle         *string `json:"InfoTitle,omitempty" xml:"InfoTitle,omitempty"`
	// example:
	//
	// 198.00
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// cmgj01****
	ProductCode    *string                                  `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	PromotionRules *DescribePriceResponseBodyPromotionRules `json:"PromotionRules,omitempty" xml:"PromotionRules,omitempty" type:"Struct"`
	// example:
	//
	// 19.80
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) GetCoupons() *DescribePriceResponseBodyCoupons {
	return s.Coupons
}

func (s *DescribePriceResponseBody) GetCurrency() *string {
	return s.Currency
}

func (s *DescribePriceResponseBody) GetCuxiao() *bool {
	return s.Cuxiao
}

func (s *DescribePriceResponseBody) GetCycle() *string {
	return s.Cycle
}

func (s *DescribePriceResponseBody) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribePriceResponseBody) GetDuration() *int32 {
	return s.Duration
}

func (s *DescribePriceResponseBody) GetExpressionCode() *string {
	return s.ExpressionCode
}

func (s *DescribePriceResponseBody) GetExpressionMessage() *string {
	return s.ExpressionMessage
}

func (s *DescribePriceResponseBody) GetInfoTitle() *string {
	return s.InfoTitle
}

func (s *DescribePriceResponseBody) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribePriceResponseBody) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribePriceResponseBody) GetPromotionRules() *DescribePriceResponseBodyPromotionRules {
	return s.PromotionRules
}

func (s *DescribePriceResponseBody) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribePriceResponseBody) SetCoupons(v *DescribePriceResponseBodyCoupons) *DescribePriceResponseBody {
	s.Coupons = v
	return s
}

func (s *DescribePriceResponseBody) SetCurrency(v string) *DescribePriceResponseBody {
	s.Currency = &v
	return s
}

func (s *DescribePriceResponseBody) SetCuxiao(v bool) *DescribePriceResponseBody {
	s.Cuxiao = &v
	return s
}

func (s *DescribePriceResponseBody) SetCycle(v string) *DescribePriceResponseBody {
	s.Cycle = &v
	return s
}

func (s *DescribePriceResponseBody) SetDiscountPrice(v float32) *DescribePriceResponseBody {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBody) SetDuration(v int32) *DescribePriceResponseBody {
	s.Duration = &v
	return s
}

func (s *DescribePriceResponseBody) SetExpressionCode(v string) *DescribePriceResponseBody {
	s.ExpressionCode = &v
	return s
}

func (s *DescribePriceResponseBody) SetExpressionMessage(v string) *DescribePriceResponseBody {
	s.ExpressionMessage = &v
	return s
}

func (s *DescribePriceResponseBody) SetInfoTitle(v string) *DescribePriceResponseBody {
	s.InfoTitle = &v
	return s
}

func (s *DescribePriceResponseBody) SetOriginalPrice(v float32) *DescribePriceResponseBody {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBody) SetProductCode(v string) *DescribePriceResponseBody {
	s.ProductCode = &v
	return s
}

func (s *DescribePriceResponseBody) SetPromotionRules(v *DescribePriceResponseBodyPromotionRules) *DescribePriceResponseBody {
	s.PromotionRules = v
	return s
}

func (s *DescribePriceResponseBody) SetTradePrice(v float32) *DescribePriceResponseBody {
	s.TradePrice = &v
	return s
}

func (s *DescribePriceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyCoupons struct {
	Coupon []*DescribePriceResponseBodyCouponsCoupon `json:"Coupon,omitempty" xml:"Coupon,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyCoupons) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyCoupons) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyCoupons) GetCoupon() []*DescribePriceResponseBodyCouponsCoupon {
	return s.Coupon
}

func (s *DescribePriceResponseBodyCoupons) SetCoupon(v []*DescribePriceResponseBodyCouponsCoupon) *DescribePriceResponseBodyCoupons {
	s.Coupon = v
	return s
}

func (s *DescribePriceResponseBodyCoupons) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyCouponsCoupon struct {
	// example:
	//
	// 100.00
	CanPromFee *float32 `json:"CanPromFee,omitempty" xml:"CanPromFee,omitempty"`
	CouponDesc *string  `json:"CouponDesc,omitempty" xml:"CouponDesc,omitempty"`
	CouponName *string  `json:"CouponName,omitempty" xml:"CouponName,omitempty"`
	// example:
	//
	// ActiveCoupon
	CouponOptionCode *string `json:"CouponOptionCode,omitempty" xml:"CouponOptionCode,omitempty"`
	// example:
	//
	// 3874923111
	CouponOptionNo *string `json:"CouponOptionNo,omitempty" xml:"CouponOptionNo,omitempty"`
	// example:
	//
	// false
	IsSelected *bool `json:"IsSelected,omitempty" xml:"IsSelected,omitempty"`
}

func (s DescribePriceResponseBodyCouponsCoupon) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyCouponsCoupon) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyCouponsCoupon) GetCanPromFee() *float32 {
	return s.CanPromFee
}

func (s *DescribePriceResponseBodyCouponsCoupon) GetCouponDesc() *string {
	return s.CouponDesc
}

func (s *DescribePriceResponseBodyCouponsCoupon) GetCouponName() *string {
	return s.CouponName
}

func (s *DescribePriceResponseBodyCouponsCoupon) GetCouponOptionCode() *string {
	return s.CouponOptionCode
}

func (s *DescribePriceResponseBodyCouponsCoupon) GetCouponOptionNo() *string {
	return s.CouponOptionNo
}

func (s *DescribePriceResponseBodyCouponsCoupon) GetIsSelected() *bool {
	return s.IsSelected
}

func (s *DescribePriceResponseBodyCouponsCoupon) SetCanPromFee(v float32) *DescribePriceResponseBodyCouponsCoupon {
	s.CanPromFee = &v
	return s
}

func (s *DescribePriceResponseBodyCouponsCoupon) SetCouponDesc(v string) *DescribePriceResponseBodyCouponsCoupon {
	s.CouponDesc = &v
	return s
}

func (s *DescribePriceResponseBodyCouponsCoupon) SetCouponName(v string) *DescribePriceResponseBodyCouponsCoupon {
	s.CouponName = &v
	return s
}

func (s *DescribePriceResponseBodyCouponsCoupon) SetCouponOptionCode(v string) *DescribePriceResponseBodyCouponsCoupon {
	s.CouponOptionCode = &v
	return s
}

func (s *DescribePriceResponseBodyCouponsCoupon) SetCouponOptionNo(v string) *DescribePriceResponseBodyCouponsCoupon {
	s.CouponOptionNo = &v
	return s
}

func (s *DescribePriceResponseBodyCouponsCoupon) SetIsSelected(v bool) *DescribePriceResponseBodyCouponsCoupon {
	s.IsSelected = &v
	return s
}

func (s *DescribePriceResponseBodyCouponsCoupon) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPromotionRules struct {
	PromotionRule []*DescribePriceResponseBodyPromotionRulesPromotionRule `json:"PromotionRule,omitempty" xml:"PromotionRule,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyPromotionRules) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPromotionRules) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPromotionRules) GetPromotionRule() []*DescribePriceResponseBodyPromotionRulesPromotionRule {
	return s.PromotionRule
}

func (s *DescribePriceResponseBodyPromotionRules) SetPromotionRule(v []*DescribePriceResponseBodyPromotionRulesPromotionRule) *DescribePriceResponseBodyPromotionRules {
	s.PromotionRule = v
	return s
}

func (s *DescribePriceResponseBodyPromotionRules) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPromotionRulesPromotionRule struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 102112
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	Title  *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribePriceResponseBodyPromotionRulesPromotionRule) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPromotionRulesPromotionRule) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPromotionRulesPromotionRule) GetName() *string {
	return s.Name
}

func (s *DescribePriceResponseBodyPromotionRulesPromotionRule) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribePriceResponseBodyPromotionRulesPromotionRule) GetTitle() *string {
	return s.Title
}

func (s *DescribePriceResponseBodyPromotionRulesPromotionRule) SetName(v string) *DescribePriceResponseBodyPromotionRulesPromotionRule {
	s.Name = &v
	return s
}

func (s *DescribePriceResponseBodyPromotionRulesPromotionRule) SetRuleId(v string) *DescribePriceResponseBodyPromotionRulesPromotionRule {
	s.RuleId = &v
	return s
}

func (s *DescribePriceResponseBodyPromotionRulesPromotionRule) SetTitle(v string) *DescribePriceResponseBodyPromotionRulesPromotionRule {
	s.Title = &v
	return s
}

func (s *DescribePriceResponseBodyPromotionRulesPromotionRule) Validate() error {
	return dara.Validate(s)
}
