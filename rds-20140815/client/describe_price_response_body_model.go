// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderParams(v string) *DescribePriceResponseBody
	GetOrderParams() *string
	SetPriceInfo(v *DescribePriceResponseBodyPriceInfo) *DescribePriceResponseBody
	GetPriceInfo() *DescribePriceResponseBodyPriceInfo
	SetRequestId(v string) *DescribePriceResponseBody
	GetRequestId() *string
	SetRules(v *DescribePriceResponseBodyRules) *DescribePriceResponseBody
	GetRules() *DescribePriceResponseBodyRules
	SetServerlessPrice(v *DescribePriceResponseBodyServerlessPrice) *DescribePriceResponseBody
	GetServerlessPrice() *DescribePriceResponseBodyServerlessPrice
	SetShowDiscount(v bool) *DescribePriceResponseBody
	GetShowDiscount() *bool
	SetTradeMaxRCUAmount(v float32) *DescribePriceResponseBody
	GetTradeMaxRCUAmount() *float32
	SetTradeMinRCUAmount(v float32) *DescribePriceResponseBody
	GetTradeMinRCUAmount() *float32
}

type DescribePriceResponseBody struct {
	// The order parameters.
	//
	// >  If the **OrderParamOut*	- parameter is set to **true**, the value of the OrderParams parameter is returned.
	//
	// example:
	//
	// {\\"autoPay\\":false}"
	OrderParams *string `json:"OrderParams,omitempty" xml:"OrderParams,omitempty"`
	// The price information.
	PriceInfo *DescribePriceResponseBodyPriceInfo `json:"PriceInfo,omitempty" xml:"PriceInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CA0ADDDC-0BEB-4381-A3ED-73B4C79B8CC6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the promotion rule.
	Rules *DescribePriceResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	// The pricing information about a serverless RDS instance.
	ServerlessPrice *DescribePriceResponseBodyServerlessPrice `json:"ServerlessPrice,omitempty" xml:"ServerlessPrice,omitempty" type:"Struct"`
	// Indicates whether discounts can be used.
	//
	// example:
	//
	// True
	ShowDiscount *bool `json:"ShowDiscount,omitempty" xml:"ShowDiscount,omitempty"`
	// The estimated hourly fee that is calculated based on the maximum number of RCUs.
	//
	// example:
	//
	// 2**
	TradeMaxRCUAmount *float32 `json:"TradeMaxRCUAmount,omitempty" xml:"TradeMaxRCUAmount,omitempty"`
	// The estimated hourly fee that is calculated based on the minimum number of RCUs.
	//
	// example:
	//
	// 1**
	TradeMinRCUAmount *float32 `json:"TradeMinRCUAmount,omitempty" xml:"TradeMinRCUAmount,omitempty"`
}

func (s DescribePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) GetOrderParams() *string {
	return s.OrderParams
}

func (s *DescribePriceResponseBody) GetPriceInfo() *DescribePriceResponseBodyPriceInfo {
	return s.PriceInfo
}

func (s *DescribePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePriceResponseBody) GetRules() *DescribePriceResponseBodyRules {
	return s.Rules
}

func (s *DescribePriceResponseBody) GetServerlessPrice() *DescribePriceResponseBodyServerlessPrice {
	return s.ServerlessPrice
}

func (s *DescribePriceResponseBody) GetShowDiscount() *bool {
	return s.ShowDiscount
}

func (s *DescribePriceResponseBody) GetTradeMaxRCUAmount() *float32 {
	return s.TradeMaxRCUAmount
}

func (s *DescribePriceResponseBody) GetTradeMinRCUAmount() *float32 {
	return s.TradeMinRCUAmount
}

func (s *DescribePriceResponseBody) SetOrderParams(v string) *DescribePriceResponseBody {
	s.OrderParams = &v
	return s
}

func (s *DescribePriceResponseBody) SetPriceInfo(v *DescribePriceResponseBodyPriceInfo) *DescribePriceResponseBody {
	s.PriceInfo = v
	return s
}

func (s *DescribePriceResponseBody) SetRequestId(v string) *DescribePriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePriceResponseBody) SetRules(v *DescribePriceResponseBodyRules) *DescribePriceResponseBody {
	s.Rules = v
	return s
}

func (s *DescribePriceResponseBody) SetServerlessPrice(v *DescribePriceResponseBodyServerlessPrice) *DescribePriceResponseBody {
	s.ServerlessPrice = v
	return s
}

func (s *DescribePriceResponseBody) SetShowDiscount(v bool) *DescribePriceResponseBody {
	s.ShowDiscount = &v
	return s
}

func (s *DescribePriceResponseBody) SetTradeMaxRCUAmount(v float32) *DescribePriceResponseBody {
	s.TradeMaxRCUAmount = &v
	return s
}

func (s *DescribePriceResponseBody) SetTradeMinRCUAmount(v float32) *DescribePriceResponseBody {
	s.TradeMinRCUAmount = &v
	return s
}

func (s *DescribePriceResponseBody) Validate() error {
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
	if s.ServerlessPrice != nil {
		if err := s.ServerlessPrice.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePriceResponseBodyPriceInfo struct {
	// The information about the promotion.
	ActivityInfo *DescribePriceResponseBodyPriceInfoActivityInfo `json:"ActivityInfo,omitempty" xml:"ActivityInfo,omitempty" type:"Struct"`
	// The information about the coupon.
	Coupons *DescribePriceResponseBodyPriceInfoCoupons `json:"Coupons,omitempty" xml:"Coupons,omitempty" type:"Struct"`
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
	// 0
	DiscountPrice *float32 `json:"DiscountPrice,omitempty" xml:"DiscountPrice,omitempty"`
	// The order information.
	//
	// example:
	//
	// Order Information
	OrderLines interface{} `json:"OrderLines,omitempty" xml:"OrderLines,omitempty"`
	// The original price.
	//
	// example:
	//
	// 2504
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// An array that consists of the ID of the promotion rule.
	RuleIds *DescribePriceResponseBodyPriceInfoRuleIds `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Struct"`
	// The estimated hourly cost that is calculated based on the maximum number of RCUs you specify.
	//
	// example:
	//
	// 1**
	TradeMaxRCUAmount *float32 `json:"TradeMaxRCUAmount,omitempty" xml:"TradeMaxRCUAmount,omitempty"`
	// The estimated hourly cost that is calculated based on the minimum number of RCUs you specify.
	//
	// example:
	//
	// 2**
	TradeMinRCUAmount *float32 `json:"TradeMinRCUAmount,omitempty" xml:"TradeMinRCUAmount,omitempty"`
	// The transaction price, which is equal to the original price minus the discount.
	//
	// example:
	//
	// 2504
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfo) GetActivityInfo() *DescribePriceResponseBodyPriceInfoActivityInfo {
	return s.ActivityInfo
}

func (s *DescribePriceResponseBodyPriceInfo) GetCoupons() *DescribePriceResponseBodyPriceInfoCoupons {
	return s.Coupons
}

func (s *DescribePriceResponseBodyPriceInfo) GetCurrency() *string {
	return s.Currency
}

func (s *DescribePriceResponseBodyPriceInfo) GetDiscountPrice() *float32 {
	return s.DiscountPrice
}

func (s *DescribePriceResponseBodyPriceInfo) GetOrderLines() interface{} {
	return s.OrderLines
}

func (s *DescribePriceResponseBodyPriceInfo) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribePriceResponseBodyPriceInfo) GetRuleIds() *DescribePriceResponseBodyPriceInfoRuleIds {
	return s.RuleIds
}

func (s *DescribePriceResponseBodyPriceInfo) GetTradeMaxRCUAmount() *float32 {
	return s.TradeMaxRCUAmount
}

func (s *DescribePriceResponseBodyPriceInfo) GetTradeMinRCUAmount() *float32 {
	return s.TradeMinRCUAmount
}

func (s *DescribePriceResponseBodyPriceInfo) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *DescribePriceResponseBodyPriceInfo) SetActivityInfo(v *DescribePriceResponseBodyPriceInfoActivityInfo) *DescribePriceResponseBodyPriceInfo {
	s.ActivityInfo = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetCoupons(v *DescribePriceResponseBodyPriceInfoCoupons) *DescribePriceResponseBodyPriceInfo {
	s.Coupons = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetCurrency(v string) *DescribePriceResponseBodyPriceInfo {
	s.Currency = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetDiscountPrice(v float32) *DescribePriceResponseBodyPriceInfo {
	s.DiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetOrderLines(v interface{}) *DescribePriceResponseBodyPriceInfo {
	s.OrderLines = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetOriginalPrice(v float32) *DescribePriceResponseBodyPriceInfo {
	s.OriginalPrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetRuleIds(v *DescribePriceResponseBodyPriceInfoRuleIds) *DescribePriceResponseBodyPriceInfo {
	s.RuleIds = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetTradeMaxRCUAmount(v float32) *DescribePriceResponseBodyPriceInfo {
	s.TradeMaxRCUAmount = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetTradeMinRCUAmount(v float32) *DescribePriceResponseBodyPriceInfo {
	s.TradeMinRCUAmount = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) SetTradePrice(v float32) *DescribePriceResponseBodyPriceInfo {
	s.TradePrice = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfo) Validate() error {
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

type DescribePriceResponseBodyPriceInfoActivityInfo struct {
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

func (s DescribePriceResponseBodyPriceInfoActivityInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoActivityInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoActivityInfo) GetCheckErrMsg() *string {
	return s.CheckErrMsg
}

func (s *DescribePriceResponseBodyPriceInfoActivityInfo) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribePriceResponseBodyPriceInfoActivityInfo) GetSuccess() *string {
	return s.Success
}

func (s *DescribePriceResponseBodyPriceInfoActivityInfo) SetCheckErrMsg(v string) *DescribePriceResponseBodyPriceInfoActivityInfo {
	s.CheckErrMsg = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoActivityInfo) SetErrorCode(v string) *DescribePriceResponseBodyPriceInfoActivityInfo {
	s.ErrorCode = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoActivityInfo) SetSuccess(v string) *DescribePriceResponseBodyPriceInfoActivityInfo {
	s.Success = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoActivityInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPriceInfoCoupons struct {
	Coupon []*DescribePriceResponseBodyPriceInfoCouponsCoupon `json:"Coupon,omitempty" xml:"Coupon,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyPriceInfoCoupons) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoCoupons) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoCoupons) GetCoupon() []*DescribePriceResponseBodyPriceInfoCouponsCoupon {
	return s.Coupon
}

func (s *DescribePriceResponseBodyPriceInfoCoupons) SetCoupon(v []*DescribePriceResponseBodyPriceInfoCouponsCoupon) *DescribePriceResponseBodyPriceInfoCoupons {
	s.Coupon = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoCoupons) Validate() error {
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

type DescribePriceResponseBodyPriceInfoCouponsCoupon struct {
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
	// The coupon name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribePriceResponseBodyPriceInfoCouponsCoupon) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoCouponsCoupon) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoCouponsCoupon) GetCouponNo() *string {
	return s.CouponNo
}

func (s *DescribePriceResponseBodyPriceInfoCouponsCoupon) GetDescription() *string {
	return s.Description
}

func (s *DescribePriceResponseBodyPriceInfoCouponsCoupon) GetIsSelected() *string {
	return s.IsSelected
}

func (s *DescribePriceResponseBodyPriceInfoCouponsCoupon) GetName() *string {
	return s.Name
}

func (s *DescribePriceResponseBodyPriceInfoCouponsCoupon) SetCouponNo(v string) *DescribePriceResponseBodyPriceInfoCouponsCoupon {
	s.CouponNo = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoCouponsCoupon) SetDescription(v string) *DescribePriceResponseBodyPriceInfoCouponsCoupon {
	s.Description = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoCouponsCoupon) SetIsSelected(v string) *DescribePriceResponseBodyPriceInfoCouponsCoupon {
	s.IsSelected = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoCouponsCoupon) SetName(v string) *DescribePriceResponseBodyPriceInfoCouponsCoupon {
	s.Name = &v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoCouponsCoupon) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyPriceInfoRuleIds struct {
	RuleId []*string `json:"RuleId,omitempty" xml:"RuleId,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyPriceInfoRuleIds) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyPriceInfoRuleIds) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyPriceInfoRuleIds) GetRuleId() []*string {
	return s.RuleId
}

func (s *DescribePriceResponseBodyPriceInfoRuleIds) SetRuleId(v []*string) *DescribePriceResponseBodyPriceInfoRuleIds {
	s.RuleId = v
	return s
}

func (s *DescribePriceResponseBodyPriceInfoRuleIds) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyRules struct {
	Rule []*DescribePriceResponseBodyRulesRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyRules) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyRules) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyRules) GetRule() []*DescribePriceResponseBodyRulesRule {
	return s.Rule
}

func (s *DescribePriceResponseBodyRules) SetRule(v []*DescribePriceResponseBodyRulesRule) *DescribePriceResponseBodyRules {
	s.Rule = v
	return s
}

func (s *DescribePriceResponseBodyRules) Validate() error {
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

type DescribePriceResponseBodyRulesRule struct {
	// The description of the promotion rule.
	//
	// example:
	//
	// Activity Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the promotion rule.
	//
	// example:
	//
	// Rule1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the promotion rule.
	//
	// example:
	//
	// 1020021003939076
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribePriceResponseBodyRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyRulesRule) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyRulesRule) GetDescription() *string {
	return s.Description
}

func (s *DescribePriceResponseBodyRulesRule) GetName() *string {
	return s.Name
}

func (s *DescribePriceResponseBodyRulesRule) GetRuleId() *int64 {
	return s.RuleId
}

func (s *DescribePriceResponseBodyRulesRule) SetDescription(v string) *DescribePriceResponseBodyRulesRule {
	s.Description = &v
	return s
}

func (s *DescribePriceResponseBodyRulesRule) SetName(v string) *DescribePriceResponseBodyRulesRule {
	s.Name = &v
	return s
}

func (s *DescribePriceResponseBodyRulesRule) SetRuleId(v int64) *DescribePriceResponseBodyRulesRule {
	s.RuleId = &v
	return s
}

func (s *DescribePriceResponseBodyRulesRule) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyServerlessPrice struct {
	// The discount amount of the maximum number of RCUs.
	//
	// example:
	//
	// 1**.*
	RCUDiscountMaxAmount *float32 `json:"RCUDiscountMaxAmount,omitempty" xml:"RCUDiscountMaxAmount,omitempty"`
	// The discount amount of the minimum number of RCUs.
	//
	// example:
	//
	// 1*.*
	RCUDiscountMinAmount *float32 `json:"RCUDiscountMinAmount,omitempty" xml:"RCUDiscountMinAmount,omitempty"`
	// The price of the maximum number of RCUs.
	//
	// example:
	//
	// 2**.*
	RCUOriginalMaxAmount *float32 `json:"RCUOriginalMaxAmount,omitempty" xml:"RCUOriginalMaxAmount,omitempty"`
	// The price of the minimum number of RCUs.
	//
	// example:
	//
	// 3*.*
	RCUOriginalMinAmount *float32 `json:"RCUOriginalMinAmount,omitempty" xml:"RCUOriginalMinAmount,omitempty"`
	// The original price of the disk capacity.
	//
	// example:
	//
	// 1*
	StorageOriginalAmount *float32 `json:"StorageOriginalAmount,omitempty" xml:"StorageOriginalAmount,omitempty"`
	// The maximum total price before the discount.
	//
	// example:
	//
	// 2**.*
	TotalOriginalMaxAmount *float32 `json:"TotalOriginalMaxAmount,omitempty" xml:"TotalOriginalMaxAmount,omitempty"`
	// The minimum total price before the discount.
	//
	// example:
	//
	// 2*.*
	TotalOriginalMinAmount *float32 `json:"TotalOriginalMinAmount,omitempty" xml:"TotalOriginalMinAmount,omitempty"`
	// The transaction price of the maximum number of RCUs.
	//
	// example:
	//
	// 1**.*
	TradeMaxRCUAmount *float32 `json:"TradeMaxRCUAmount,omitempty" xml:"TradeMaxRCUAmount,omitempty"`
	// The transaction price of the minimum number of RCUs.
	//
	// example:
	//
	// 2*.*
	TradeMinRCUAmount *float32 `json:"TradeMinRCUAmount,omitempty" xml:"TradeMinRCUAmount,omitempty"`
	// The discounted price of the disk capacity.
	//
	// example:
	//
	// 2.*
	StorageDiscountAmount *float32 `json:"storageDiscountAmount,omitempty" xml:"storageDiscountAmount,omitempty"`
}

func (s DescribePriceResponseBodyServerlessPrice) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyServerlessPrice) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyServerlessPrice) GetRCUDiscountMaxAmount() *float32 {
	return s.RCUDiscountMaxAmount
}

func (s *DescribePriceResponseBodyServerlessPrice) GetRCUDiscountMinAmount() *float32 {
	return s.RCUDiscountMinAmount
}

func (s *DescribePriceResponseBodyServerlessPrice) GetRCUOriginalMaxAmount() *float32 {
	return s.RCUOriginalMaxAmount
}

func (s *DescribePriceResponseBodyServerlessPrice) GetRCUOriginalMinAmount() *float32 {
	return s.RCUOriginalMinAmount
}

func (s *DescribePriceResponseBodyServerlessPrice) GetStorageOriginalAmount() *float32 {
	return s.StorageOriginalAmount
}

func (s *DescribePriceResponseBodyServerlessPrice) GetTotalOriginalMaxAmount() *float32 {
	return s.TotalOriginalMaxAmount
}

func (s *DescribePriceResponseBodyServerlessPrice) GetTotalOriginalMinAmount() *float32 {
	return s.TotalOriginalMinAmount
}

func (s *DescribePriceResponseBodyServerlessPrice) GetTradeMaxRCUAmount() *float32 {
	return s.TradeMaxRCUAmount
}

func (s *DescribePriceResponseBodyServerlessPrice) GetTradeMinRCUAmount() *float32 {
	return s.TradeMinRCUAmount
}

func (s *DescribePriceResponseBodyServerlessPrice) GetStorageDiscountAmount() *float32 {
	return s.StorageDiscountAmount
}

func (s *DescribePriceResponseBodyServerlessPrice) SetRCUDiscountMaxAmount(v float32) *DescribePriceResponseBodyServerlessPrice {
	s.RCUDiscountMaxAmount = &v
	return s
}

func (s *DescribePriceResponseBodyServerlessPrice) SetRCUDiscountMinAmount(v float32) *DescribePriceResponseBodyServerlessPrice {
	s.RCUDiscountMinAmount = &v
	return s
}

func (s *DescribePriceResponseBodyServerlessPrice) SetRCUOriginalMaxAmount(v float32) *DescribePriceResponseBodyServerlessPrice {
	s.RCUOriginalMaxAmount = &v
	return s
}

func (s *DescribePriceResponseBodyServerlessPrice) SetRCUOriginalMinAmount(v float32) *DescribePriceResponseBodyServerlessPrice {
	s.RCUOriginalMinAmount = &v
	return s
}

func (s *DescribePriceResponseBodyServerlessPrice) SetStorageOriginalAmount(v float32) *DescribePriceResponseBodyServerlessPrice {
	s.StorageOriginalAmount = &v
	return s
}

func (s *DescribePriceResponseBodyServerlessPrice) SetTotalOriginalMaxAmount(v float32) *DescribePriceResponseBodyServerlessPrice {
	s.TotalOriginalMaxAmount = &v
	return s
}

func (s *DescribePriceResponseBodyServerlessPrice) SetTotalOriginalMinAmount(v float32) *DescribePriceResponseBodyServerlessPrice {
	s.TotalOriginalMinAmount = &v
	return s
}

func (s *DescribePriceResponseBodyServerlessPrice) SetTradeMaxRCUAmount(v float32) *DescribePriceResponseBodyServerlessPrice {
	s.TradeMaxRCUAmount = &v
	return s
}

func (s *DescribePriceResponseBodyServerlessPrice) SetTradeMinRCUAmount(v float32) *DescribePriceResponseBodyServerlessPrice {
	s.TradeMinRCUAmount = &v
	return s
}

func (s *DescribePriceResponseBodyServerlessPrice) SetStorageDiscountAmount(v float32) *DescribePriceResponseBodyServerlessPrice {
	s.StorageDiscountAmount = &v
	return s
}

func (s *DescribePriceResponseBodyServerlessPrice) Validate() error {
	return dara.Validate(s)
}
