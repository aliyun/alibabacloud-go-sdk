// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenewalPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrder(v *DescribeRenewalPriceResponseBodyOrder) *DescribeRenewalPriceResponseBody
	GetOrder() *DescribeRenewalPriceResponseBodyOrder
	SetRequestId(v string) *DescribeRenewalPriceResponseBody
	GetRequestId() *string
	SetRules(v *DescribeRenewalPriceResponseBodyRules) *DescribeRenewalPriceResponseBody
	GetRules() *DescribeRenewalPriceResponseBodyRules
	SetSubOrders(v *DescribeRenewalPriceResponseBodySubOrders) *DescribeRenewalPriceResponseBody
	GetSubOrders() *DescribeRenewalPriceResponseBodySubOrders
}

type DescribeRenewalPriceResponseBody struct {
	// The list of orders.
	Order *DescribeRenewalPriceResponseBodyOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// EFD65226-08CC-4C4D-B6A4-CB3C382F67B0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about the promotion rules.
	Rules *DescribeRenewalPriceResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	// The rules matching the coupons.
	SubOrders *DescribeRenewalPriceResponseBodySubOrders `json:"SubOrders,omitempty" xml:"SubOrders,omitempty" type:"Struct"`
}

func (s DescribeRenewalPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBody) GetOrder() *DescribeRenewalPriceResponseBodyOrder {
	return s.Order
}

func (s *DescribeRenewalPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRenewalPriceResponseBody) GetRules() *DescribeRenewalPriceResponseBodyRules {
	return s.Rules
}

func (s *DescribeRenewalPriceResponseBody) GetSubOrders() *DescribeRenewalPriceResponseBodySubOrders {
	return s.SubOrders
}

func (s *DescribeRenewalPriceResponseBody) SetOrder(v *DescribeRenewalPriceResponseBodyOrder) *DescribeRenewalPriceResponseBody {
	s.Order = v
	return s
}

func (s *DescribeRenewalPriceResponseBody) SetRequestId(v string) *DescribeRenewalPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRenewalPriceResponseBody) SetRules(v *DescribeRenewalPriceResponseBodyRules) *DescribeRenewalPriceResponseBody {
	s.Rules = v
	return s
}

func (s *DescribeRenewalPriceResponseBody) SetSubOrders(v *DescribeRenewalPriceResponseBodySubOrders) *DescribeRenewalPriceResponseBody {
	s.SubOrders = v
	return s
}

func (s *DescribeRenewalPriceResponseBody) Validate() error {
	if s.Order != nil {
		if err := s.Order.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		if err := s.Rules.Validate(); err != nil {
			return err
		}
	}
	if s.SubOrders != nil {
		if err := s.SubOrders.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRenewalPriceResponseBodyOrder struct {
	// Details about the coupons.
	Coupons *DescribeRenewalPriceResponseBodyOrderCoupons `json:"Coupons,omitempty" xml:"Coupons,omitempty" type:"Struct"`
	// The type of the currency. Valid values:
	//
	// 	- USD: United States dollar
	//
	// 	- JPY: Japanese Yen
	//
	// example:
	//
	// USD
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discount amount of the order.
	//
	// example:
	//
	// 1144.8
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The original price of the order.
	//
	// example:
	//
	// 1144.8
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The IDs of the matched rules.
	RuleIds *DescribeRenewalPriceResponseBodyOrderRuleIds `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Struct"`
	// The actual price of the order.
	//
	// example:
	//
	// 0
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyOrder) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyOrder) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyOrder) GetCoupons() *DescribeRenewalPriceResponseBodyOrderCoupons {
	return s.Coupons
}

func (s *DescribeRenewalPriceResponseBodyOrder) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeRenewalPriceResponseBodyOrder) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *DescribeRenewalPriceResponseBodyOrder) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *DescribeRenewalPriceResponseBodyOrder) GetRuleIds() *DescribeRenewalPriceResponseBodyOrderRuleIds {
	return s.RuleIds
}

func (s *DescribeRenewalPriceResponseBodyOrder) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *DescribeRenewalPriceResponseBodyOrder) SetCoupons(v *DescribeRenewalPriceResponseBodyOrderCoupons) *DescribeRenewalPriceResponseBodyOrder {
	s.Coupons = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyOrder) SetCurrency(v string) *DescribeRenewalPriceResponseBodyOrder {
	s.Currency = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyOrder) SetDiscountAmount(v float32) *DescribeRenewalPriceResponseBodyOrder {
	s.DiscountAmount = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyOrder) SetOriginalAmount(v float32) *DescribeRenewalPriceResponseBodyOrder {
	s.OriginalAmount = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyOrder) SetRuleIds(v *DescribeRenewalPriceResponseBodyOrderRuleIds) *DescribeRenewalPriceResponseBodyOrder {
	s.RuleIds = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyOrder) SetTradeAmount(v float32) *DescribeRenewalPriceResponseBodyOrder {
	s.TradeAmount = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyOrder) Validate() error {
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

type DescribeRenewalPriceResponseBodyOrderCoupons struct {
	Coupon []*DescribeRenewalPriceResponseBodyOrderCouponsCoupon `json:"Coupon,omitempty" xml:"Coupon,omitempty" type:"Repeated"`
}

func (s DescribeRenewalPriceResponseBodyOrderCoupons) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyOrderCoupons) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyOrderCoupons) GetCoupon() []*DescribeRenewalPriceResponseBodyOrderCouponsCoupon {
	return s.Coupon
}

func (s *DescribeRenewalPriceResponseBodyOrderCoupons) SetCoupon(v []*DescribeRenewalPriceResponseBodyOrderCouponsCoupon) *DescribeRenewalPriceResponseBodyOrderCoupons {
	s.Coupon = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyOrderCoupons) Validate() error {
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

type DescribeRenewalPriceResponseBodyOrderCouponsCoupon struct {
	// The coupon number.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The description of the coupon.
	//
	// example:
	//
	// coupondemo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the coupon was selected.
	//
	// example:
	//
	// true
	IsSelected *string `json:"IsSelected,omitempty" xml:"IsSelected,omitempty"`
	// The name of the coupon.
	//
	// example:
	//
	// youhuiquan111
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyOrderCouponsCoupon) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyOrderCouponsCoupon) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyOrderCouponsCoupon) GetCouponNo() *string {
	return s.CouponNo
}

func (s *DescribeRenewalPriceResponseBodyOrderCouponsCoupon) GetDescription() *string {
	return s.Description
}

func (s *DescribeRenewalPriceResponseBodyOrderCouponsCoupon) GetIsSelected() *string {
	return s.IsSelected
}

func (s *DescribeRenewalPriceResponseBodyOrderCouponsCoupon) GetName() *string {
	return s.Name
}

func (s *DescribeRenewalPriceResponseBodyOrderCouponsCoupon) SetCouponNo(v string) *DescribeRenewalPriceResponseBodyOrderCouponsCoupon {
	s.CouponNo = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyOrderCouponsCoupon) SetDescription(v string) *DescribeRenewalPriceResponseBodyOrderCouponsCoupon {
	s.Description = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyOrderCouponsCoupon) SetIsSelected(v string) *DescribeRenewalPriceResponseBodyOrderCouponsCoupon {
	s.IsSelected = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyOrderCouponsCoupon) SetName(v string) *DescribeRenewalPriceResponseBodyOrderCouponsCoupon {
	s.Name = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyOrderCouponsCoupon) Validate() error {
	return dara.Validate(s)
}

type DescribeRenewalPriceResponseBodyOrderRuleIds struct {
	RuleId []*string `json:"RuleId,omitempty" xml:"RuleId,omitempty" type:"Repeated"`
}

func (s DescribeRenewalPriceResponseBodyOrderRuleIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyOrderRuleIds) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyOrderRuleIds) GetRuleId() []*string {
	return s.RuleId
}

func (s *DescribeRenewalPriceResponseBodyOrderRuleIds) SetRuleId(v []*string) *DescribeRenewalPriceResponseBodyOrderRuleIds {
	s.RuleId = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyOrderRuleIds) Validate() error {
	return dara.Validate(s)
}

type DescribeRenewalPriceResponseBodyRules struct {
	Rule []*DescribeRenewalPriceResponseBodyRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribeRenewalPriceResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyRules) GetRule() []*DescribeRenewalPriceResponseBodyRulesRule {
	return s.Rule
}

func (s *DescribeRenewalPriceResponseBodyRules) SetRule(v []*DescribeRenewalPriceResponseBodyRulesRule) *DescribeRenewalPriceResponseBodyRules {
	s.Rule = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyRules) Validate() error {
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

type DescribeRenewalPriceResponseBodyRulesRule struct {
	// The name of the rule.
	//
	// example:
	//
	// demoname
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 11111111
	RuleDescId *int64 `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
	// The title of the rule.
	//
	// example:
	//
	// demo
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyRulesRule) GetName() *string {
	return s.Name
}

func (s *DescribeRenewalPriceResponseBodyRulesRule) GetRuleDescId() *int64 {
	return s.RuleDescId
}

func (s *DescribeRenewalPriceResponseBodyRulesRule) GetTitle() *string {
	return s.Title
}

func (s *DescribeRenewalPriceResponseBodyRulesRule) SetName(v string) *DescribeRenewalPriceResponseBodyRulesRule {
	s.Name = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyRulesRule) SetRuleDescId(v int64) *DescribeRenewalPriceResponseBodyRulesRule {
	s.RuleDescId = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyRulesRule) SetTitle(v string) *DescribeRenewalPriceResponseBodyRulesRule {
	s.Title = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyRulesRule) Validate() error {
	return dara.Validate(s)
}

type DescribeRenewalPriceResponseBodySubOrders struct {
	SubOrder []*DescribeRenewalPriceResponseBodySubOrdersSubOrder `json:"SubOrder,omitempty" xml:"SubOrder,omitempty" type:"Repeated"`
}

func (s DescribeRenewalPriceResponseBodySubOrders) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodySubOrders) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodySubOrders) GetSubOrder() []*DescribeRenewalPriceResponseBodySubOrdersSubOrder {
	return s.SubOrder
}

func (s *DescribeRenewalPriceResponseBodySubOrders) SetSubOrder(v []*DescribeRenewalPriceResponseBodySubOrdersSubOrder) *DescribeRenewalPriceResponseBodySubOrders {
	s.SubOrder = v
	return s
}

func (s *DescribeRenewalPriceResponseBodySubOrders) Validate() error {
	if s.SubOrder != nil {
		for _, item := range s.SubOrder {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRenewalPriceResponseBodySubOrdersSubOrder struct {
	// The discount amount of the order.
	//
	// example:
	//
	// 1144.8
	DiscountAmount *float32 `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// dds-bp12c5b040dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The original price of the order.
	//
	// example:
	//
	// 1144.8
	OriginalAmount *float32 `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The IDs of the matched rules.
	RuleIds *DescribeRenewalPriceResponseBodySubOrdersSubOrderRuleIds `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Struct"`
	// The actual price of the order.
	//
	// example:
	//
	// 0
	TradeAmount *float32 `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s DescribeRenewalPriceResponseBodySubOrdersSubOrder) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodySubOrdersSubOrder) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodySubOrdersSubOrder) GetDiscountAmount() *float32 {
	return s.DiscountAmount
}

func (s *DescribeRenewalPriceResponseBodySubOrdersSubOrder) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeRenewalPriceResponseBodySubOrdersSubOrder) GetOriginalAmount() *float32 {
	return s.OriginalAmount
}

func (s *DescribeRenewalPriceResponseBodySubOrdersSubOrder) GetRuleIds() *DescribeRenewalPriceResponseBodySubOrdersSubOrderRuleIds {
	return s.RuleIds
}

func (s *DescribeRenewalPriceResponseBodySubOrdersSubOrder) GetTradeAmount() *float32 {
	return s.TradeAmount
}

func (s *DescribeRenewalPriceResponseBodySubOrdersSubOrder) SetDiscountAmount(v float32) *DescribeRenewalPriceResponseBodySubOrdersSubOrder {
	s.DiscountAmount = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodySubOrdersSubOrder) SetInstanceId(v string) *DescribeRenewalPriceResponseBodySubOrdersSubOrder {
	s.InstanceId = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodySubOrdersSubOrder) SetOriginalAmount(v float32) *DescribeRenewalPriceResponseBodySubOrdersSubOrder {
	s.OriginalAmount = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodySubOrdersSubOrder) SetRuleIds(v *DescribeRenewalPriceResponseBodySubOrdersSubOrderRuleIds) *DescribeRenewalPriceResponseBodySubOrdersSubOrder {
	s.RuleIds = v
	return s
}

func (s *DescribeRenewalPriceResponseBodySubOrdersSubOrder) SetTradeAmount(v float32) *DescribeRenewalPriceResponseBodySubOrdersSubOrder {
	s.TradeAmount = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodySubOrdersSubOrder) Validate() error {
	if s.RuleIds != nil {
		if err := s.RuleIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRenewalPriceResponseBodySubOrdersSubOrderRuleIds struct {
	RuleId []*string `json:"RuleId,omitempty" xml:"RuleId,omitempty" type:"Repeated"`
}

func (s DescribeRenewalPriceResponseBodySubOrdersSubOrderRuleIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodySubOrdersSubOrderRuleIds) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodySubOrdersSubOrderRuleIds) GetRuleId() []*string {
	return s.RuleId
}

func (s *DescribeRenewalPriceResponseBodySubOrdersSubOrderRuleIds) SetRuleId(v []*string) *DescribeRenewalPriceResponseBodySubOrdersSubOrderRuleIds {
	s.RuleId = v
	return s
}

func (s *DescribeRenewalPriceResponseBodySubOrdersSubOrderRuleIds) Validate() error {
	return dara.Validate(s)
}
