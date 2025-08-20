// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribePriceResponseBody
	GetCode() *string
	SetCoupons(v []*DescribePriceResponseBodyCoupons) *DescribePriceResponseBody
	GetCoupons() []*DescribePriceResponseBodyCoupons
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
	SetMessage(v string) *DescribePriceResponseBody
	GetMessage() *string
	SetOriginalPrice(v float32) *DescribePriceResponseBody
	GetOriginalPrice() *float32
	SetProductCode(v string) *DescribePriceResponseBody
	GetProductCode() *string
	SetPromotionRules(v []*DescribePriceResponseBodyPromotionRules) *DescribePriceResponseBody
	GetPromotionRules() []*DescribePriceResponseBodyPromotionRules
	SetRequestId(v string) *DescribePriceResponseBody
	GetRequestId() *string
	SetTradePrice(v float32) *DescribePriceResponseBody
	GetTradePrice() *float32
}

type DescribePriceResponseBody struct {
	// example:
	//
	// 200
	Code    *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Coupons []*DescribePriceResponseBodyCoupons `json:"Coupons,omitempty" xml:"Coupons,omitempty" type:"Repeated"`
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	Cuxiao   *bool   `json:"Cuxiao,omitempty" xml:"Cuxiao,omitempty"`
	// example:
	//
	// once
	Cycle *string `json:"Cycle,omitempty" xml:"Cycle,omitempty"`
	// example:
	//
	// 0.0
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// example:
	//
	// 3
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// NO_AVAILABLE_PAYMENT_METHOD
	ExpressionCode *string `json:"ExpressionCode,omitempty" xml:"ExpressionCode,omitempty"`
	// example:
	//
	// ExpressionMessage
	ExpressionMessage *string `json:"ExpressionMessage,omitempty" xml:"ExpressionMessage,omitempty"`
	// example:
	//
	// 139.41
	InfoTitle *string `json:"InfoTitle,omitempty" xml:"InfoTitle,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2099.0
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// csas
	ProductCode    *string                                    `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	PromotionRules []*DescribePriceResponseBodyPromotionRules `json:"PromotionRules,omitempty" xml:"PromotionRules,omitempty" type:"Repeated"`
	// example:
	//
	// 054AF571-C86F-533F-8A7B-F62206FA4367
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1320.0
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribePriceResponseBody) GetCoupons() []*DescribePriceResponseBodyCoupons {
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

func (s *DescribePriceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribePriceResponseBody) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribePriceResponseBody) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribePriceResponseBody) GetPromotionRules() []*DescribePriceResponseBodyPromotionRules {
	return s.PromotionRules
}

func (s *DescribePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePriceResponseBody) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribePriceResponseBody) SetCode(v string) *DescribePriceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePriceResponseBody) SetCoupons(v []*DescribePriceResponseBodyCoupons) *DescribePriceResponseBody {
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

func (s *DescribePriceResponseBody) SetMessage(v string) *DescribePriceResponseBody {
	s.Message = &v
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

func (s *DescribePriceResponseBody) SetPromotionRules(v []*DescribePriceResponseBodyPromotionRules) *DescribePriceResponseBody {
	s.PromotionRules = v
	return s
}

func (s *DescribePriceResponseBody) SetRequestId(v string) *DescribePriceResponseBody {
	s.RequestId = &v
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
	// example:
	//
	// CanPromFee
	CanPromFee *int64  `json:"CanPromFee,omitempty" xml:"CanPromFee,omitempty"`
	CouponDesc *string `json:"CouponDesc,omitempty" xml:"CouponDesc,omitempty"`
	// example:
	//
	// CouponName
	CouponName *string `json:"CouponName,omitempty" xml:"CouponName,omitempty"`
	// example:
	//
	// CouponOptionCode
	CouponOptionCode *string `json:"CouponOptionCode,omitempty" xml:"CouponOptionCode,omitempty"`
	// example:
	//
	// CouponOptionNo
	CouponOptionNo *string `json:"CouponOptionNo,omitempty" xml:"CouponOptionNo,omitempty"`
	// example:
	//
	// IsSelected
	IsSelected *bool `json:"IsSelected,omitempty" xml:"IsSelected,omitempty"`
	// example:
	//
	// PromotionDesc
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
}

func (s DescribePriceResponseBodyCoupons) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyCoupons) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyCoupons) GetCanPromFee() *int64 {
	return s.CanPromFee
}

func (s *DescribePriceResponseBodyCoupons) GetCouponDesc() *string {
	return s.CouponDesc
}

func (s *DescribePriceResponseBodyCoupons) GetCouponName() *string {
	return s.CouponName
}

func (s *DescribePriceResponseBodyCoupons) GetCouponOptionCode() *string {
	return s.CouponOptionCode
}

func (s *DescribePriceResponseBodyCoupons) GetCouponOptionNo() *string {
	return s.CouponOptionNo
}

func (s *DescribePriceResponseBodyCoupons) GetIsSelected() *bool {
	return s.IsSelected
}

func (s *DescribePriceResponseBodyCoupons) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *DescribePriceResponseBodyCoupons) SetCanPromFee(v int64) *DescribePriceResponseBodyCoupons {
	s.CanPromFee = &v
	return s
}

func (s *DescribePriceResponseBodyCoupons) SetCouponDesc(v string) *DescribePriceResponseBodyCoupons {
	s.CouponDesc = &v
	return s
}

func (s *DescribePriceResponseBodyCoupons) SetCouponName(v string) *DescribePriceResponseBodyCoupons {
	s.CouponName = &v
	return s
}

func (s *DescribePriceResponseBodyCoupons) SetCouponOptionCode(v string) *DescribePriceResponseBodyCoupons {
	s.CouponOptionCode = &v
	return s
}

func (s *DescribePriceResponseBodyCoupons) SetCouponOptionNo(v string) *DescribePriceResponseBodyCoupons {
	s.CouponOptionNo = &v
	return s
}

func (s *DescribePriceResponseBodyCoupons) SetIsSelected(v bool) *DescribePriceResponseBodyCoupons {
	s.IsSelected = &v
	return s
}

func (s *DescribePriceResponseBodyCoupons) SetPromotionDesc(v string) *DescribePriceResponseBodyCoupons {
	s.PromotionDesc = &v
	return s
}

func (s *DescribePriceResponseBodyCoupons) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPromotionRules struct {
	// example:
	//
	// alb-xdnghlhvm9vvo1sk54_Accelerator
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1827975
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// 12
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribePriceResponseBodyPromotionRules) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPromotionRules) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPromotionRules) GetName() *string {
	return s.Name
}

func (s *DescribePriceResponseBodyPromotionRules) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribePriceResponseBodyPromotionRules) GetTitle() *string {
	return s.Title
}

func (s *DescribePriceResponseBodyPromotionRules) SetName(v string) *DescribePriceResponseBodyPromotionRules {
	s.Name = &v
	return s
}

func (s *DescribePriceResponseBodyPromotionRules) SetRuleId(v string) *DescribePriceResponseBodyPromotionRules {
	s.RuleId = &v
	return s
}

func (s *DescribePriceResponseBodyPromotionRules) SetTitle(v string) *DescribePriceResponseBodyPromotionRules {
	s.Title = &v
	return s
}

func (s *DescribePriceResponseBodyPromotionRules) Validate() error {
	return dara.Validate(s)
}
