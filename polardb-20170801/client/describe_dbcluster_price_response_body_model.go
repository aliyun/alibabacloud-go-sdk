// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderType(v string) *DescribeDBClusterPriceResponseBody
	GetOrderType() *string
	SetPriceInfo(v *DescribeDBClusterPriceResponseBodyPriceInfo) *DescribeDBClusterPriceResponseBody
	GetPriceInfo() *DescribeDBClusterPriceResponseBodyPriceInfo
	SetRequestId(v string) *DescribeDBClusterPriceResponseBody
	GetRequestId() *string
	SetRules(v *DescribeDBClusterPriceResponseBodyRules) *DescribeDBClusterPriceResponseBody
	GetRules() *DescribeDBClusterPriceResponseBodyRules
	SetShowDiscount(v bool) *DescribeDBClusterPriceResponseBody
	GetShowDiscount() *bool
}

type DescribeDBClusterPriceResponseBody struct {
	// The order type. Valid values:
	//
	// 	- BUY: new purchase.
	//
	// 	- UPGRADE: specification change.
	//
	// 	- RENEW: renewal.
	//
	// 	- CONVERT: billing method conversion.
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The price details.
	PriceInfo *DescribeDBClusterPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7E2FE3BB-C677-5FF9-9FC5-XXX
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rules     *DescribeDBClusterPriceResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	// Indicates whether discounts are allowed.
	//
	// example:
	//
	// False
	ShowDiscount *bool `json:"ShowDiscount,omitempty" xml:"ShowDiscount,omitempty"`
}

func (s DescribeDBClusterPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPriceResponseBody) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeDBClusterPriceResponseBody) GetPriceInfo() *DescribeDBClusterPriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *DescribeDBClusterPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterPriceResponseBody) GetRules() *DescribeDBClusterPriceResponseBodyRules {
	return s.Rules
}

func (s *DescribeDBClusterPriceResponseBody) GetShowDiscount() *bool {
	return s.ShowDiscount
}

func (s *DescribeDBClusterPriceResponseBody) SetOrderType(v string) *DescribeDBClusterPriceResponseBody {
	s.OrderType = &v
	return s
}

func (s *DescribeDBClusterPriceResponseBody) SetPriceInfo(v *DescribeDBClusterPriceResponseBodyPriceInfo) *DescribeDBClusterPriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribeDBClusterPriceResponseBody) SetRequestId(v string) *DescribeDBClusterPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterPriceResponseBody) SetRules(v *DescribeDBClusterPriceResponseBodyRules) *DescribeDBClusterPriceResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeDBClusterPriceResponseBody) SetShowDiscount(v bool) *DescribeDBClusterPriceResponseBody {
	s.ShowDiscount = &v
	return s
}

func (s *DescribeDBClusterPriceResponseBody) Validate() error {
	if s.PriceInfo != nil {
		if err := s.PriceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		if err := s.Rules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterPriceResponseBodyPriceInfo struct {
	Coupons *DescribeDBClusterPriceResponseBodyPriceInfoCoupons `json:"Coupons,omitempty" xml:"Coupons,omitempty" type:"Struct"`
	// The currency unit.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discount amount.
	//
	// example:
	//
	// 1978.2
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 1099.0
	OriginalPrice *float32                                            `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	RuleIds       *DescribeDBClusterPriceResponseBodyPriceInfoRuleIds `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Struct"`
	// The final price, which is the original price minus the discount.
	//
	// example:
	//
	// 165.0
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeDBClusterPriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfo) GetCoupons() *DescribeDBClusterPriceResponseBodyPriceInfoCoupons {
	return s.Coupons
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfo) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfo) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfo) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfo) GetRuleIds() *DescribeDBClusterPriceResponseBodyPriceInfoRuleIds {
	return s.RuleIds
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfo) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfo) SetCoupons(v *DescribeDBClusterPriceResponseBodyPriceInfoCoupons) *DescribeDBClusterPriceResponseBodyPriceInfo {
	s.Coupons = v
	return s
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfo) SetCurrency(v string) *DescribeDBClusterPriceResponseBodyPriceInfo {
	s.Currency = &v
	return s
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfo) SetDiscountPrice(v float32) *DescribeDBClusterPriceResponseBodyPriceInfo {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfo) SetOriginalPrice(v float32) *DescribeDBClusterPriceResponseBodyPriceInfo {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfo) SetRuleIds(v *DescribeDBClusterPriceResponseBodyPriceInfoRuleIds) *DescribeDBClusterPriceResponseBodyPriceInfo {
	s.RuleIds = v
	return s
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfo) SetTradePrice(v float32) *DescribeDBClusterPriceResponseBodyPriceInfo {
	s.TradePrice = &v
	return s
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfo) Validate() error {
	if s.Coupons != nil {
		if err := s.Coupons.Validate(); err != nil {
			return err
		}
	}
	if s.RuleIds != nil {
		if err := s.RuleIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBClusterPriceResponseBodyPriceInfoCoupons struct {
	Coupon []*DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon `json:"Coupon,omitempty" xml:"Coupon,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPriceResponseBodyPriceInfoCoupons) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPriceResponseBodyPriceInfoCoupons) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfoCoupons) GetCoupon() []*DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon {
	return s.Coupon
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfoCoupons) SetCoupon(v []*DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon) *DescribeDBClusterPriceResponseBodyPriceInfoCoupons {
	s.Coupon = v
	return s
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfoCoupons) Validate() error {
	if s.Coupon != nil {
		for _, item := range s.Coupon {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon struct {
	CouponNo   *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	IsSelected *string `json:"IsSelected,omitempty" xml:"IsSelected,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon) GetCouponNo() *string {
	return s.CouponNo
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon) GetIsSelected() *string {
	return s.IsSelected
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon) GetName() *string {
	return s.Name
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon) SetCouponNo(v string) *DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon {
	s.CouponNo = &v
	return s
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon) SetIsSelected(v string) *DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon {
	s.IsSelected = &v
	return s
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon) SetName(v string) *DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfoCouponsCoupon) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterPriceResponseBodyPriceInfoRuleIds struct {
	RuleId []*string `json:"RuleId,omitempty" xml:"RuleId,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPriceResponseBodyPriceInfoRuleIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPriceResponseBodyPriceInfoRuleIds) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfoRuleIds) GetRuleId() []*string {
	return s.RuleId
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfoRuleIds) SetRuleId(v []*string) *DescribeDBClusterPriceResponseBodyPriceInfoRuleIds {
	s.RuleId = v
	return s
}

func (s *DescribeDBClusterPriceResponseBodyPriceInfoRuleIds) Validate() error {
	return dara.Validate(s)
}

type DescribeDBClusterPriceResponseBodyRules struct {
	Rule []*DescribeDBClusterPriceResponseBodyRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeDBClusterPriceResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPriceResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPriceResponseBodyRules) GetRule() []*DescribeDBClusterPriceResponseBodyRulesRule {
	return s.Rule
}

func (s *DescribeDBClusterPriceResponseBodyRules) SetRule(v []*DescribeDBClusterPriceResponseBodyRulesRule) *DescribeDBClusterPriceResponseBodyRules {
	s.Rule = v
	return s
}

func (s *DescribeDBClusterPriceResponseBodyRules) Validate() error {
	if s.Rule != nil {
		for _, item := range s.Rule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterPriceResponseBodyRulesRule struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RuleId *int64  `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeDBClusterPriceResponseBodyRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterPriceResponseBodyRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPriceResponseBodyRulesRule) GetName() *string {
	return s.Name
}

func (s *DescribeDBClusterPriceResponseBodyRulesRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeDBClusterPriceResponseBodyRulesRule) SetName(v string) *DescribeDBClusterPriceResponseBodyRulesRule {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPriceResponseBodyRulesRule) SetRuleId(v int64) *DescribeDBClusterPriceResponseBodyRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeDBClusterPriceResponseBodyRulesRule) Validate() error {
	return dara.Validate(s)
}
