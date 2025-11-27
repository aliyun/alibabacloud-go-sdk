// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenewalPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPriceInfo(v *DescribeRenewalPriceResponseBodyPriceInfo) *DescribeRenewalPriceResponseBody
	GetPriceInfo() *DescribeRenewalPriceResponseBodyPriceInfo
	SetRequestId(v string) *DescribeRenewalPriceResponseBody
	GetRequestId() *string
	SetRules(v *DescribeRenewalPriceResponseBodyRules) *DescribeRenewalPriceResponseBody
	GetRules() *DescribeRenewalPriceResponseBodyRules
}

type DescribeRenewalPriceResponseBody struct {
	// Details of price information.
	PriceInfo *DescribeRenewalPriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// DC9F4EF6-D038-4405-B497-1F48E722C9F2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the details of the promotion rule.
	Rules *DescribeRenewalPriceResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
}

func (s DescribeRenewalPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBody) GetPriceInfo() *DescribeRenewalPriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *DescribeRenewalPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRenewalPriceResponseBody) GetRules() *DescribeRenewalPriceResponseBodyRules {
	return s.Rules
}

func (s *DescribeRenewalPriceResponseBody) SetPriceInfo(v *DescribeRenewalPriceResponseBodyPriceInfo) *DescribeRenewalPriceResponseBody {
	s.PriceInfo = v
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

func (s *DescribeRenewalPriceResponseBody) Validate() error {
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

type DescribeRenewalPriceResponseBodyPriceInfo struct {
	// The information about the promotion.
	ActivityInfo *DescribeRenewalPriceResponseBodyPriceInfoActivityInfo `json:"ActivityInfo,omitempty" xml:"ActivityInfo,omitempty" type:"Struct"`
	// An array that consists of information about the coupon.
	Coupons *DescribeRenewalPriceResponseBodyPriceInfoCoupons `json:"Coupons,omitempty" xml:"Coupons,omitempty" type:"Struct"`
	// The currency unit.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The discount.
	//
	// example:
	//
	// 27
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The original price.
	//
	// example:
	//
	// 138
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// An array that consists of the ID of the promotion rule.
	RuleIds *DescribeRenewalPriceResponseBodyPriceInfoRuleIds `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Struct"`
	// The transaction price, which is equal to the original price minus the discount.
	//
	// example:
	//
	// 111
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) GetActivityInfo() *DescribeRenewalPriceResponseBodyPriceInfoActivityInfo {
	return s.ActivityInfo
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) GetCoupons() *DescribeRenewalPriceResponseBodyPriceInfoCoupons {
	return s.Coupons
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) GetCurrency() *string {
	return s.Currency
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) GetRuleIds() *DescribeRenewalPriceResponseBodyPriceInfoRuleIds {
	return s.RuleIds
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) SetActivityInfo(v *DescribeRenewalPriceResponseBodyPriceInfoActivityInfo) *DescribeRenewalPriceResponseBodyPriceInfo {
	s.ActivityInfo = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) SetCoupons(v *DescribeRenewalPriceResponseBodyPriceInfoCoupons) *DescribeRenewalPriceResponseBodyPriceInfo {
	s.Coupons = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) SetCurrency(v string) *DescribeRenewalPriceResponseBodyPriceInfo {
	s.Currency = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) SetDiscountPrice(v float32) *DescribeRenewalPriceResponseBodyPriceInfo {
	s.DiscountPrice = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) SetOriginalPrice(v float32) *DescribeRenewalPriceResponseBodyPriceInfo {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) SetRuleIds(v *DescribeRenewalPriceResponseBodyPriceInfoRuleIds) *DescribeRenewalPriceResponseBodyPriceInfo {
	s.RuleIds = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) SetTradePrice(v float32) *DescribeRenewalPriceResponseBodyPriceInfo {
	s.TradePrice = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfo) Validate() error {
	if s.ActivityInfo != nil {
		if err := s.ActivityInfo.Validate(); err != nil {
			return err
		}
	}
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

type DescribeRenewalPriceResponseBodyPriceInfoActivityInfo struct {
	// The returned message.
	//
	// example:
	//
	// Error description
	CheckErrMsg *string `json:"CheckErrMsg,omitempty" xml:"CheckErrMsg,omitempty"`
	// The error code that is returned.
	//
	// example:
	//
	// 123456
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// Success
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoActivityInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoActivityInfo) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoActivityInfo) GetCheckErrMsg() *string {
	return s.CheckErrMsg
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoActivityInfo) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoActivityInfo) GetSuccess() *string {
	return s.Success
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoActivityInfo) SetCheckErrMsg(v string) *DescribeRenewalPriceResponseBodyPriceInfoActivityInfo {
	s.CheckErrMsg = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoActivityInfo) SetErrorCode(v string) *DescribeRenewalPriceResponseBodyPriceInfoActivityInfo {
	s.ErrorCode = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoActivityInfo) SetSuccess(v string) *DescribeRenewalPriceResponseBodyPriceInfoActivityInfo {
	s.Success = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoActivityInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeRenewalPriceResponseBodyPriceInfoCoupons struct {
	Coupon []*DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon `json:"Coupon,omitempty" xml:"Coupon,omitempty" type:"Repeated"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoCoupons) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoCoupons) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoCoupons) GetCoupon() []*DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon {
	return s.Coupon
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoCoupons) SetCoupon(v []*DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon) *DescribeRenewalPriceResponseBodyPriceInfoCoupons {
	s.Coupon = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoCoupons) Validate() error {
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

type DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon struct {
	// The coupon ID.
	//
	// example:
	//
	// 123456
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The description of the coupon.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the coupon is selected.
	//
	// example:
	//
	// true
	IsSelected *string `json:"IsSelected,omitempty" xml:"IsSelected,omitempty"`
	// The name of the coupon.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon) GetCouponNo() *string {
	return s.CouponNo
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon) GetDescription() *string {
	return s.Description
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon) GetIsSelected() *string {
	return s.IsSelected
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon) GetName() *string {
	return s.Name
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon) SetCouponNo(v string) *DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon {
	s.CouponNo = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon) SetDescription(v string) *DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon {
	s.Description = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon) SetIsSelected(v string) *DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon {
	s.IsSelected = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon) SetName(v string) *DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon {
	s.Name = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoCouponsCoupon) Validate() error {
	return dara.Validate(s)
}

type DescribeRenewalPriceResponseBodyPriceInfoRuleIds struct {
	RuleId []*string `json:"RuleId,omitempty" xml:"RuleId,omitempty" type:"Repeated"`
}

func (s DescribeRenewalPriceResponseBodyPriceInfoRuleIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyPriceInfoRuleIds) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRuleIds) GetRuleId() []*string {
	return s.RuleId
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRuleIds) SetRuleId(v []*string) *DescribeRenewalPriceResponseBodyPriceInfoRuleIds {
	s.RuleId = v
	return s
}

func (s *DescribeRenewalPriceResponseBodyPriceInfoRuleIds) Validate() error {
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
	// The description of the activity.
	//
	// example:
	//
	// Content
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the promotion rule.
	//
	// example:
	//
	// 1001199213
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeRenewalPriceResponseBodyRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenewalPriceResponseBodyRulesRule) GoString() string {
	return s.String()
}

func (s *DescribeRenewalPriceResponseBodyRulesRule) GetDescription() *string {
	return s.Description
}

func (s *DescribeRenewalPriceResponseBodyRulesRule) GetName() *string {
	return s.Name
}

func (s *DescribeRenewalPriceResponseBodyRulesRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribeRenewalPriceResponseBodyRulesRule) SetDescription(v string) *DescribeRenewalPriceResponseBodyRulesRule {
	s.Description = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyRulesRule) SetName(v string) *DescribeRenewalPriceResponseBodyRulesRule {
	s.Name = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyRulesRule) SetRuleId(v int64) *DescribeRenewalPriceResponseBodyRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribeRenewalPriceResponseBodyRulesRule) Validate() error {
	return dara.Validate(s)
}
