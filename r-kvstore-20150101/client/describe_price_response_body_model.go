// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrder(v *DescribePriceResponseBodyOrder) *DescribePriceResponseBody
	GetOrder() *DescribePriceResponseBodyOrder
	SetOrderParams(v string) *DescribePriceResponseBody
	GetOrderParams() *string
	SetRequestId(v string) *DescribePriceResponseBody
	GetRequestId() *string
	SetRules(v *DescribePriceResponseBodyRules) *DescribePriceResponseBody
	GetRules() *DescribePriceResponseBodyRules
	SetSubOrders(v *DescribePriceResponseBodySubOrders) *DescribePriceResponseBody
	GetSubOrders() *DescribePriceResponseBodySubOrders
}

type DescribePriceResponseBody struct {
	// The information about the order.
	Order *DescribePriceResponseBodyOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	// The parameters of the order. When OrderParamOut is set to `true`, this parameter is returned.
	//
	// example:
	//
	// String
	OrderParams *string `json:"OrderParams,omitempty" xml:"OrderParams,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3A40BE4E-1890-4972-889C-FEFA37663635
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about promotion rules.
	Rules *DescribePriceResponseBodyRules `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Struct"`
	// The orders that match the coupons.
	SubOrders *DescribePriceResponseBodySubOrders `json:"SubOrders,omitempty" xml:"SubOrders,omitempty" type:"Struct"`
}

func (s DescribePriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBody) GetOrder() *DescribePriceResponseBodyOrder {
	return s.Order
}

func (s *DescribePriceResponseBody) GetOrderParams() *string {
	return s.OrderParams
}

func (s *DescribePriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePriceResponseBody) GetRules() *DescribePriceResponseBodyRules {
	return s.Rules
}

func (s *DescribePriceResponseBody) GetSubOrders() *DescribePriceResponseBodySubOrders {
	return s.SubOrders
}

func (s *DescribePriceResponseBody) SetOrder(v *DescribePriceResponseBodyOrder) *DescribePriceResponseBody {
	s.Order = v
	return s
}

func (s *DescribePriceResponseBody) SetOrderParams(v string) *DescribePriceResponseBody {
	s.OrderParams = &v
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

func (s *DescribePriceResponseBody) SetSubOrders(v *DescribePriceResponseBodySubOrders) *DescribePriceResponseBody {
	s.SubOrders = v
	return s
}

func (s *DescribePriceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyOrder struct {
	// The order code.
	//
	// example:
	//
	// ""
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details about coupons.
	Coupons *DescribePriceResponseBodyOrderCoupons `json:"Coupons,omitempty" xml:"Coupons,omitempty" type:"Struct"`
	// The currency used for payment. A value of CNY is used when the order was generated on the China site (aliyun.com), and a value of USD is used when the order was generated on the international site (alibabacloud.com).
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The activity information.
	DepreciateInfo *DescribePriceResponseBodyOrderDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// The discount amount of the order.
	//
	// example:
	//
	// 0.21
	DiscountAmount *string `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The service fees of the order.
	//
	// example:
	//
	// 0.1
	HandlingFeeAmount *string `json:"HandlingFeeAmount,omitempty" xml:"HandlingFeeAmount,omitempty"`
	// Indicates whether eligibility for the contracted discount is met.
	//
	// example:
	//
	// false
	IsContractActivity *bool `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	// The information about the order.
	//
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The original price of the order.
	//
	// example:
	//
	// 0.21
	OriginalAmount *string `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// The rule IDs.
	RuleIds *DescribePriceResponseBodyOrderRuleIds `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Struct"`
	// Indicates whether the discount information is displayed.
	ShowDiscountInfo *bool `json:"ShowDiscountInfo,omitempty" xml:"ShowDiscountInfo,omitempty"`
	// The discount.
	//
	// example:
	//
	// ****
	StandDiscountPrice *int64 `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	// The discounted price.
	//
	// example:
	//
	// 0
	StandPrice *int64 `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// The transaction price of the order.
	//
	// example:
	//
	// 10
	TradeAmount *string `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s DescribePriceResponseBodyOrder) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyOrder) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyOrder) GetCode() *string {
	return s.Code
}

func (s *DescribePriceResponseBodyOrder) GetCoupons() *DescribePriceResponseBodyOrderCoupons {
	return s.Coupons
}

func (s *DescribePriceResponseBodyOrder) GetCurrency() *string {
	return s.Currency
}

func (s *DescribePriceResponseBodyOrder) GetDepreciateInfo() *DescribePriceResponseBodyOrderDepreciateInfo {
	return s.DepreciateInfo
}

func (s *DescribePriceResponseBodyOrder) GetDiscountAmount() *string {
	return s.DiscountAmount
}

func (s *DescribePriceResponseBodyOrder) GetHandlingFeeAmount() *string {
	return s.HandlingFeeAmount
}

func (s *DescribePriceResponseBodyOrder) GetIsContractActivity() *bool {
	return s.IsContractActivity
}

func (s *DescribePriceResponseBodyOrder) GetMessage() *string {
	return s.Message
}

func (s *DescribePriceResponseBodyOrder) GetOriginalAmount() *string {
	return s.OriginalAmount
}

func (s *DescribePriceResponseBodyOrder) GetRuleIds() *DescribePriceResponseBodyOrderRuleIds {
	return s.RuleIds
}

func (s *DescribePriceResponseBodyOrder) GetShowDiscountInfo() *bool {
	return s.ShowDiscountInfo
}

func (s *DescribePriceResponseBodyOrder) GetStandDiscountPrice() *int64 {
	return s.StandDiscountPrice
}

func (s *DescribePriceResponseBodyOrder) GetStandPrice() *int64 {
	return s.StandPrice
}

func (s *DescribePriceResponseBodyOrder) GetTradeAmount() *string {
	return s.TradeAmount
}

func (s *DescribePriceResponseBodyOrder) SetCode(v string) *DescribePriceResponseBodyOrder {
	s.Code = &v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetCoupons(v *DescribePriceResponseBodyOrderCoupons) *DescribePriceResponseBodyOrder {
	s.Coupons = v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetCurrency(v string) *DescribePriceResponseBodyOrder {
	s.Currency = &v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetDepreciateInfo(v *DescribePriceResponseBodyOrderDepreciateInfo) *DescribePriceResponseBodyOrder {
	s.DepreciateInfo = v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetDiscountAmount(v string) *DescribePriceResponseBodyOrder {
	s.DiscountAmount = &v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetHandlingFeeAmount(v string) *DescribePriceResponseBodyOrder {
	s.HandlingFeeAmount = &v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetIsContractActivity(v bool) *DescribePriceResponseBodyOrder {
	s.IsContractActivity = &v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetMessage(v string) *DescribePriceResponseBodyOrder {
	s.Message = &v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetOriginalAmount(v string) *DescribePriceResponseBodyOrder {
	s.OriginalAmount = &v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetRuleIds(v *DescribePriceResponseBodyOrderRuleIds) *DescribePriceResponseBodyOrder {
	s.RuleIds = v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetShowDiscountInfo(v bool) *DescribePriceResponseBodyOrder {
	s.ShowDiscountInfo = &v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetStandDiscountPrice(v int64) *DescribePriceResponseBodyOrder {
	s.StandDiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetStandPrice(v int64) *DescribePriceResponseBodyOrder {
	s.StandPrice = &v
	return s
}

func (s *DescribePriceResponseBodyOrder) SetTradeAmount(v string) *DescribePriceResponseBodyOrder {
	s.TradeAmount = &v
	return s
}

func (s *DescribePriceResponseBodyOrder) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyOrderCoupons struct {
	Coupon []*DescribePriceResponseBodyOrderCouponsCoupon `json:"Coupon,omitempty" xml:"Coupon,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyOrderCoupons) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyOrderCoupons) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyOrderCoupons) GetCoupon() []*DescribePriceResponseBodyOrderCouponsCoupon {
	return s.Coupon
}

func (s *DescribePriceResponseBodyOrderCoupons) SetCoupon(v []*DescribePriceResponseBodyOrderCouponsCoupon) *DescribePriceResponseBodyOrderCoupons {
	s.Coupon = v
	return s
}

func (s *DescribePriceResponseBodyOrderCoupons) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyOrderCouponsCoupon struct {
	// The coupon ID.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The coupon description.
	//
	// example:
	//
	// coupondemo
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

func (s DescribePriceResponseBodyOrderCouponsCoupon) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyOrderCouponsCoupon) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyOrderCouponsCoupon) GetCouponNo() *string {
	return s.CouponNo
}

func (s *DescribePriceResponseBodyOrderCouponsCoupon) GetDescription() *string {
	return s.Description
}

func (s *DescribePriceResponseBodyOrderCouponsCoupon) GetIsSelected() *string {
	return s.IsSelected
}

func (s *DescribePriceResponseBodyOrderCouponsCoupon) GetName() *string {
	return s.Name
}

func (s *DescribePriceResponseBodyOrderCouponsCoupon) SetCouponNo(v string) *DescribePriceResponseBodyOrderCouponsCoupon {
	s.CouponNo = &v
	return s
}

func (s *DescribePriceResponseBodyOrderCouponsCoupon) SetDescription(v string) *DescribePriceResponseBodyOrderCouponsCoupon {
	s.Description = &v
	return s
}

func (s *DescribePriceResponseBodyOrderCouponsCoupon) SetIsSelected(v string) *DescribePriceResponseBodyOrderCouponsCoupon {
	s.IsSelected = &v
	return s
}

func (s *DescribePriceResponseBodyOrderCouponsCoupon) SetName(v string) *DescribePriceResponseBodyOrderCouponsCoupon {
	s.Name = &v
	return s
}

func (s *DescribePriceResponseBodyOrderCouponsCoupon) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyOrderDepreciateInfo struct {
	// The price reduction rate.
	//
	// example:
	//
	// 30%
	CheapRate *int64 `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// The new total price displayed on the official website.
	//
	// example:
	//
	// 9*
	CheapStandAmount *int64 `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// The activity information.
	ContractActivity *DescribePriceResponseBodyOrderDepreciateInfoContractActivity `json:"ContractActivity,omitempty" xml:"ContractActivity,omitempty" type:"Struct"`
	// The promotional offer (displayed in the total order amount).
	//
	// example:
	//
	// **
	Differential *int64 `json:"Differential,omitempty" xml:"Differential,omitempty"`
	// The name of the promotional offer.
	//
	// example:
	//
	// XXXXX
	DifferentialName *string `json:"DifferentialName,omitempty" xml:"DifferentialName,omitempty"`
	// Indicates whether eligibility for the contracted discount is met.
	//
	// example:
	//
	// false
	IsContractActivity *bool `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	// Indicates whether the price reduction rate is displayed.
	//
	// example:
	//
	// true
	IsShow *bool `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	// The list price.
	//
	// example:
	//
	// 1*
	ListPrice *int64 `json:"ListPrice,omitempty" xml:"ListPrice,omitempty"`
	// The monthly price.
	//
	// example:
	//
	// **
	MonthPrice *int64 `json:"MonthPrice,omitempty" xml:"MonthPrice,omitempty"`
	// The original total price displayed on the official website.
	//
	// example:
	//
	// 12*
	OriginalStandAmount *int64 `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
}

func (s DescribePriceResponseBodyOrderDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyOrderDepreciateInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) GetCheapRate() *int64 {
	return s.CheapRate
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) GetCheapStandAmount() *int64 {
	return s.CheapStandAmount
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) GetContractActivity() *DescribePriceResponseBodyOrderDepreciateInfoContractActivity {
	return s.ContractActivity
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) GetDifferential() *int64 {
	return s.Differential
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) GetDifferentialName() *string {
	return s.DifferentialName
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) GetIsContractActivity() *bool {
	return s.IsContractActivity
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) GetListPrice() *int64 {
	return s.ListPrice
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) GetMonthPrice() *int64 {
	return s.MonthPrice
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) GetOriginalStandAmount() *int64 {
	return s.OriginalStandAmount
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) SetCheapRate(v int64) *DescribePriceResponseBodyOrderDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) SetCheapStandAmount(v int64) *DescribePriceResponseBodyOrderDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) SetContractActivity(v *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) *DescribePriceResponseBodyOrderDepreciateInfo {
	s.ContractActivity = v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) SetDifferential(v int64) *DescribePriceResponseBodyOrderDepreciateInfo {
	s.Differential = &v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) SetDifferentialName(v string) *DescribePriceResponseBodyOrderDepreciateInfo {
	s.DifferentialName = &v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) SetIsContractActivity(v bool) *DescribePriceResponseBodyOrderDepreciateInfo {
	s.IsContractActivity = &v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) SetIsShow(v bool) *DescribePriceResponseBodyOrderDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) SetListPrice(v int64) *DescribePriceResponseBodyOrderDepreciateInfo {
	s.ListPrice = &v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) SetMonthPrice(v int64) *DescribePriceResponseBodyOrderDepreciateInfo {
	s.MonthPrice = &v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) SetOriginalStandAmount(v int64) *DescribePriceResponseBodyOrderDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyOrderDepreciateInfoContractActivity struct {
	// The activity ID.
	//
	// example:
	//
	// ****
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// The activity name.
	//
	// example:
	//
	// contract promotion_order_xxx discount
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// The preferential price.
	FinalFee *float64 `json:"FinalFee,omitempty" xml:"FinalFee,omitempty"`
	// The total discount amount.
	FinalPromFee *float64 `json:"FinalPromFee,omitempty" xml:"FinalPromFee,omitempty"`
	// The promotion code.
	//
	// example:
	//
	// ****
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The promotion IDs.
	OptionIds *DescribePriceResponseBodyOrderDepreciateInfoContractActivityOptionIds `json:"OptionIds,omitempty" xml:"OptionIds,omitempty" type:"Struct"`
	// The original price.
	//
	// example:
	//
	// ****
	ProdFee *float64 `json:"ProdFee,omitempty" xml:"ProdFee,omitempty"`
}

func (s DescribePriceResponseBodyOrderDepreciateInfoContractActivity) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyOrderDepreciateInfoContractActivity) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) GetActivityName() *string {
	return s.ActivityName
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) GetFinalFee() *float64 {
	return s.FinalFee
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) GetFinalPromFee() *float64 {
	return s.FinalPromFee
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) GetOptionIds() *DescribePriceResponseBodyOrderDepreciateInfoContractActivityOptionIds {
	return s.OptionIds
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) GetProdFee() *float64 {
	return s.ProdFee
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) SetActivityId(v int64) *DescribePriceResponseBodyOrderDepreciateInfoContractActivity {
	s.ActivityId = &v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) SetActivityName(v string) *DescribePriceResponseBodyOrderDepreciateInfoContractActivity {
	s.ActivityName = &v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) SetFinalFee(v float64) *DescribePriceResponseBodyOrderDepreciateInfoContractActivity {
	s.FinalFee = &v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) SetFinalPromFee(v float64) *DescribePriceResponseBodyOrderDepreciateInfoContractActivity {
	s.FinalPromFee = &v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) SetOptionCode(v string) *DescribePriceResponseBodyOrderDepreciateInfoContractActivity {
	s.OptionCode = &v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) SetOptionIds(v *DescribePriceResponseBodyOrderDepreciateInfoContractActivityOptionIds) *DescribePriceResponseBodyOrderDepreciateInfoContractActivity {
	s.OptionIds = v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) SetProdFee(v float64) *DescribePriceResponseBodyOrderDepreciateInfoContractActivity {
	s.ProdFee = &v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivity) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyOrderDepreciateInfoContractActivityOptionIds struct {
	OptionId []*int64 `json:"OptionId,omitempty" xml:"OptionId,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyOrderDepreciateInfoContractActivityOptionIds) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyOrderDepreciateInfoContractActivityOptionIds) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivityOptionIds) GetOptionId() []*int64 {
	return s.OptionId
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivityOptionIds) SetOptionId(v []*int64) *DescribePriceResponseBodyOrderDepreciateInfoContractActivityOptionIds {
	s.OptionId = v
	return s
}

func (s *DescribePriceResponseBodyOrderDepreciateInfoContractActivityOptionIds) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodyOrderRuleIds struct {
	RuleId []*string `json:"RuleId,omitempty" xml:"RuleId,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodyOrderRuleIds) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyOrderRuleIds) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyOrderRuleIds) GetRuleId() []*string {
	return s.RuleId
}

func (s *DescribePriceResponseBodyOrderRuleIds) SetRuleId(v []*string) *DescribePriceResponseBodyOrderRuleIds {
	s.RuleId = v
	return s
}

func (s *DescribePriceResponseBodyOrderRuleIds) Validate() error {
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
	return dara.Validate(s)
}

type DescribePriceResponseBodyRulesRule struct {
	// The name of the rule.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 1111111111
	RuleDescId *int64 `json:"RuleDescId,omitempty" xml:"RuleDescId,omitempty"`
	// The title of the rule.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribePriceResponseBodyRulesRule) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodyRulesRule) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodyRulesRule) GetName() *string {
	return s.Name
}

func (s *DescribePriceResponseBodyRulesRule) GetRuleDescId() *int64 {
	return s.RuleDescId
}

func (s *DescribePriceResponseBodyRulesRule) GetTitle() *string {
	return s.Title
}

func (s *DescribePriceResponseBodyRulesRule) SetName(v string) *DescribePriceResponseBodyRulesRule {
	s.Name = &v
	return s
}

func (s *DescribePriceResponseBodyRulesRule) SetRuleDescId(v int64) *DescribePriceResponseBodyRulesRule {
	s.RuleDescId = &v
	return s
}

func (s *DescribePriceResponseBodyRulesRule) SetTitle(v string) *DescribePriceResponseBodyRulesRule {
	s.Title = &v
	return s
}

func (s *DescribePriceResponseBodyRulesRule) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodySubOrders struct {
	SubOrder []*DescribePriceResponseBodySubOrdersSubOrder `json:"SubOrder,omitempty" xml:"SubOrder,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodySubOrders) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodySubOrders) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrders) GetSubOrder() []*DescribePriceResponseBodySubOrdersSubOrder {
	return s.SubOrder
}

func (s *DescribePriceResponseBodySubOrders) SetSubOrder(v []*DescribePriceResponseBodySubOrdersSubOrder) *DescribePriceResponseBodySubOrders {
	s.SubOrder = v
	return s
}

func (s *DescribePriceResponseBodySubOrders) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodySubOrdersSubOrder struct {
	// Indicates whether eligibility for the contracted discount is met.
	//
	// example:
	//
	// ****
	ContractActivity *bool `json:"ContractActivity,omitempty" xml:"ContractActivity,omitempty"`
	// The price reduction information.
	DepreciateInfo *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// The discount amount of the order.
	//
	// example:
	//
	// 0.21
	DiscountAmount *string `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// r-bp1xxxxxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether eligibility for the contracted discount is met.
	//
	// example:
	//
	// ****
	IsContractActivity *bool `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	// The configuration details for each instance of an order line item.
	ModuleInstance *DescribePriceResponseBodySubOrdersSubOrderModuleInstance `json:"ModuleInstance,omitempty" xml:"ModuleInstance,omitempty" type:"Struct"`
	// The optional promotions.
	OptionalPromotions *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotions `json:"OptionalPromotions,omitempty" xml:"OptionalPromotions,omitempty" type:"Struct"`
	// The original price of the order.
	//
	// example:
	//
	// 0.21
	OriginalAmount *string `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
	// Details about promotions.
	PromDetailList *DescribePriceResponseBodySubOrdersSubOrderPromDetailList `json:"PromDetailList,omitempty" xml:"PromDetailList,omitempty" type:"Struct"`
	// The hit rule IDs.
	RuleIds *DescribePriceResponseBodySubOrdersSubOrderRuleIds `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Struct"`
	// The discounted price.
	//
	// example:
	//
	// ****
	StandDiscountPrice *int64 `json:"StandDiscountPrice,omitempty" xml:"StandDiscountPrice,omitempty"`
	// The discounted price.
	//
	// example:
	//
	// ****
	StandPrice *int64 `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// The transaction price of the order.
	//
	// example:
	//
	// 10
	TradeAmount *string `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
}

func (s DescribePriceResponseBodySubOrdersSubOrder) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrder) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) GetContractActivity() *bool {
	return s.ContractActivity
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) GetDepreciateInfo() *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo {
	return s.DepreciateInfo
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) GetDiscountAmount() *string {
	return s.DiscountAmount
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) GetIsContractActivity() *bool {
	return s.IsContractActivity
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) GetModuleInstance() *DescribePriceResponseBodySubOrdersSubOrderModuleInstance {
	return s.ModuleInstance
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) GetOptionalPromotions() *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotions {
	return s.OptionalPromotions
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) GetOriginalAmount() *string {
	return s.OriginalAmount
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) GetPromDetailList() *DescribePriceResponseBodySubOrdersSubOrderPromDetailList {
	return s.PromDetailList
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) GetRuleIds() *DescribePriceResponseBodySubOrdersSubOrderRuleIds {
	return s.RuleIds
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) GetStandDiscountPrice() *int64 {
	return s.StandDiscountPrice
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) GetStandPrice() *int64 {
	return s.StandPrice
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) GetTradeAmount() *string {
	return s.TradeAmount
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetContractActivity(v bool) *DescribePriceResponseBodySubOrdersSubOrder {
	s.ContractActivity = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetDepreciateInfo(v *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) *DescribePriceResponseBodySubOrdersSubOrder {
	s.DepreciateInfo = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetDiscountAmount(v string) *DescribePriceResponseBodySubOrdersSubOrder {
	s.DiscountAmount = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetInstanceId(v string) *DescribePriceResponseBodySubOrdersSubOrder {
	s.InstanceId = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetIsContractActivity(v bool) *DescribePriceResponseBodySubOrdersSubOrder {
	s.IsContractActivity = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetModuleInstance(v *DescribePriceResponseBodySubOrdersSubOrderModuleInstance) *DescribePriceResponseBodySubOrdersSubOrder {
	s.ModuleInstance = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetOptionalPromotions(v *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotions) *DescribePriceResponseBodySubOrdersSubOrder {
	s.OptionalPromotions = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetOriginalAmount(v string) *DescribePriceResponseBodySubOrdersSubOrder {
	s.OriginalAmount = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetPromDetailList(v *DescribePriceResponseBodySubOrdersSubOrderPromDetailList) *DescribePriceResponseBodySubOrdersSubOrder {
	s.PromDetailList = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetRuleIds(v *DescribePriceResponseBodySubOrdersSubOrderRuleIds) *DescribePriceResponseBodySubOrdersSubOrder {
	s.RuleIds = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetStandDiscountPrice(v int64) *DescribePriceResponseBodySubOrdersSubOrder {
	s.StandDiscountPrice = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetStandPrice(v int64) *DescribePriceResponseBodySubOrdersSubOrder {
	s.StandPrice = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) SetTradeAmount(v string) *DescribePriceResponseBodySubOrdersSubOrder {
	s.TradeAmount = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrder) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo struct {
	// The price reduction rate.
	//
	// example:
	//
	// ****
	CheapRate *int64 `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// The new total price displayed on the official website.
	//
	// example:
	//
	// ****
	CheapStandAmount *int64 `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// The contract promotion.
	ContractActivity *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity `json:"ContractActivity,omitempty" xml:"ContractActivity,omitempty" type:"Struct"`
	// The promotional offer (displayed in the total order amount).
	//
	// example:
	//
	// ****
	Differential *int64 `json:"Differential,omitempty" xml:"Differential,omitempty"`
	// The name of the promotional offer.
	//
	// example:
	//
	// ****
	DifferentialName *string `json:"DifferentialName,omitempty" xml:"DifferentialName,omitempty"`
	// Indicates whether eligibility for the contracted discount is met.
	//
	// example:
	//
	// false
	IsContractActivity *bool `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	// The list price.
	//
	// example:
	//
	// ****
	ListPrice *int64 `json:"ListPrice,omitempty" xml:"ListPrice,omitempty"`
	// The monthly price.
	//
	// example:
	//
	// ****
	MonthPrice *int64 `json:"MonthPrice,omitempty" xml:"MonthPrice,omitempty"`
	// The original total price displayed on the official website.
	//
	// example:
	//
	// ****
	OriginalStandAmount *int64 `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
	// The start time of the activity.
	//
	// example:
	//
	// 2024-11-18T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) GetCheapRate() *int64 {
	return s.CheapRate
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) GetCheapStandAmount() *int64 {
	return s.CheapStandAmount
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) GetContractActivity() *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity {
	return s.ContractActivity
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) GetDifferential() *int64 {
	return s.Differential
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) GetDifferentialName() *string {
	return s.DifferentialName
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) GetIsContractActivity() *bool {
	return s.IsContractActivity
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) GetListPrice() *int64 {
	return s.ListPrice
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) GetMonthPrice() *int64 {
	return s.MonthPrice
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) GetOriginalStandAmount() *int64 {
	return s.OriginalStandAmount
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) SetCheapRate(v int64) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) SetCheapStandAmount(v int64) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) SetContractActivity(v *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo {
	s.ContractActivity = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) SetDifferential(v int64) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo {
	s.Differential = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) SetDifferentialName(v string) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo {
	s.DifferentialName = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) SetIsContractActivity(v bool) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo {
	s.IsContractActivity = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) SetListPrice(v int64) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo {
	s.ListPrice = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) SetMonthPrice(v int64) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo {
	s.MonthPrice = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) SetOriginalStandAmount(v int64) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) SetStartTime(v string) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo {
	s.StartTime = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity struct {
	// The activity ID.
	//
	// example:
	//
	// 1412025702634847
	ActivityId *int64 `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// The activity name.
	//
	// example:
	//
	// ****
	ActivityName *string `json:"ActivityName,omitempty" xml:"ActivityName,omitempty"`
	// The preferential price.
	//
	// example:
	//
	// ****
	FinalFee *float64 `json:"FinalFee,omitempty" xml:"FinalFee,omitempty"`
	// The total discount amount.
	//
	// example:
	//
	// ****
	FinalPromFee *float64 `json:"FinalPromFee,omitempty" xml:"FinalPromFee,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// ****
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The promotion IDs.
	OptionIds *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivityOptionIds `json:"OptionIds,omitempty" xml:"OptionIds,omitempty" type:"Struct"`
	// The original price.
	//
	// example:
	//
	// ****
	ProdFee *float64 `json:"ProdFee,omitempty" xml:"ProdFee,omitempty"`
}

func (s DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) GetActivityId() *int64 {
	return s.ActivityId
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) GetActivityName() *string {
	return s.ActivityName
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) GetFinalFee() *float64 {
	return s.FinalFee
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) GetFinalPromFee() *float64 {
	return s.FinalPromFee
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) GetOptionIds() *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivityOptionIds {
	return s.OptionIds
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) GetProdFee() *float64 {
	return s.ProdFee
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) SetActivityId(v int64) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity {
	s.ActivityId = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) SetActivityName(v string) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity {
	s.ActivityName = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) SetFinalFee(v float64) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity {
	s.FinalFee = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) SetFinalPromFee(v float64) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity {
	s.FinalPromFee = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) SetOptionCode(v string) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity {
	s.OptionCode = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) SetOptionIds(v *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivityOptionIds) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity {
	s.OptionIds = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) SetProdFee(v float64) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity {
	s.ProdFee = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivity) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivityOptionIds struct {
	OptionId []*int64 `json:"OptionId,omitempty" xml:"OptionId,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivityOptionIds) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivityOptionIds) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivityOptionIds) GetOptionId() []*int64 {
	return s.OptionId
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivityOptionIds) SetOptionId(v []*int64) *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivityOptionIds {
	s.OptionId = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderDepreciateInfoContractActivityOptionIds) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodySubOrdersSubOrderModuleInstance struct {
	ModuleInstance []*DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance `json:"ModuleInstance,omitempty" xml:"ModuleInstance,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodySubOrdersSubOrderModuleInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrderModuleInstance) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstance) GetModuleInstance() []*DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance {
	return s.ModuleInstance
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstance) SetModuleInstance(v []*DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) *DescribePriceResponseBodySubOrdersSubOrderModuleInstance {
	s.ModuleInstance = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstance) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance struct {
	// Indicates whether eligibility for the contracted discount is met.
	//
	// example:
	//
	// ****
	ContractActivity *bool `json:"ContractActivity,omitempty" xml:"ContractActivity,omitempty"`
	// The price reduction information.
	DepreciateInfo *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo `json:"DepreciateInfo,omitempty" xml:"DepreciateInfo,omitempty" type:"Struct"`
	// The discount.
	//
	// example:
	//
	// ****
	DiscountFee *float64 `json:"DiscountFee,omitempty" xml:"DiscountFee,omitempty"`
	// The module attributes.
	ModuleAttrs *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrs `json:"ModuleAttrs,omitempty" xml:"ModuleAttrs,omitempty" type:"Struct"`
	// The module code.
	//
	// example:
	//
	// ****
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// The module ID.
	//
	// example:
	//
	// ****
	ModuleId *string `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// The module name.
	//
	// example:
	//
	// ****
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// Indicates whether the order is paid.
	//
	// example:
	//
	// true
	NeedOrderPay *bool `json:"NeedOrderPay,omitempty" xml:"NeedOrderPay,omitempty"`
	// The actual amount paid.
	//
	// example:
	//
	// 0.0
	PayFee *float64 `json:"PayFee,omitempty" xml:"PayFee,omitempty"`
	// Indicates whether the item is billed.
	//
	// example:
	//
	// ****
	PricingModule *bool `json:"PricingModule,omitempty" xml:"PricingModule,omitempty"`
	// The discounted price.
	//
	// example:
	//
	// ****
	StandPrice *float64 `json:"StandPrice,omitempty" xml:"StandPrice,omitempty"`
	// The original price of the instance.
	//
	// example:
	//
	// 0.0
	TotalProductFee *float64 `json:"TotalProductFee,omitempty" xml:"TotalProductFee,omitempty"`
}

func (s DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) GetContractActivity() *bool {
	return s.ContractActivity
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) GetDepreciateInfo() *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo {
	return s.DepreciateInfo
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) GetDiscountFee() *float64 {
	return s.DiscountFee
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) GetModuleAttrs() *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrs {
	return s.ModuleAttrs
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) GetModuleId() *string {
	return s.ModuleId
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) GetModuleName() *string {
	return s.ModuleName
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) GetNeedOrderPay() *bool {
	return s.NeedOrderPay
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) GetPayFee() *float64 {
	return s.PayFee
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) GetPricingModule() *bool {
	return s.PricingModule
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) GetStandPrice() *float64 {
	return s.StandPrice
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) GetTotalProductFee() *float64 {
	return s.TotalProductFee
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) SetContractActivity(v bool) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance {
	s.ContractActivity = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) SetDepreciateInfo(v *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance {
	s.DepreciateInfo = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) SetDiscountFee(v float64) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance {
	s.DiscountFee = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) SetModuleAttrs(v *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrs) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance {
	s.ModuleAttrs = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) SetModuleCode(v string) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance {
	s.ModuleCode = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) SetModuleId(v string) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance {
	s.ModuleId = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) SetModuleName(v string) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance {
	s.ModuleName = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) SetNeedOrderPay(v bool) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance {
	s.NeedOrderPay = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) SetPayFee(v float64) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance {
	s.PayFee = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) SetPricingModule(v bool) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance {
	s.PricingModule = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) SetStandPrice(v float64) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance {
	s.StandPrice = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) SetTotalProductFee(v float64) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance {
	s.TotalProductFee = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstance) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo struct {
	// The price reduction rate.
	//
	// example:
	//
	// ******
	CheapRate *float64 `json:"CheapRate,omitempty" xml:"CheapRate,omitempty"`
	// The new total price displayed on the official website.
	//
	// example:
	//
	// ****
	CheapStandAmount *float64 `json:"CheapStandAmount,omitempty" xml:"CheapStandAmount,omitempty"`
	// The promotional offer (displayed in the total order amount).
	//
	// example:
	//
	// ****
	Differential *float64 `json:"Differential,omitempty" xml:"Differential,omitempty"`
	// The name of the promotional offer.
	//
	// example:
	//
	// ****
	DifferentialName *string `json:"DifferentialName,omitempty" xml:"DifferentialName,omitempty"`
	// Indicates whether eligibility for the contracted discount is met.
	//
	// example:
	//
	// ****
	IsContractActivity *bool `json:"IsContractActivity,omitempty" xml:"IsContractActivity,omitempty"`
	// Indicates whether the price reduction rate is displayed.
	//
	// example:
	//
	// true
	IsShow *bool `json:"IsShow,omitempty" xml:"IsShow,omitempty"`
	// The list price.
	//
	// example:
	//
	// ****
	ListPrice *float64 `json:"ListPrice,omitempty" xml:"ListPrice,omitempty"`
	// The monthly price.
	//
	// example:
	//
	// ****
	MonthPrice *float64 `json:"MonthPrice,omitempty" xml:"MonthPrice,omitempty"`
	// The original total price displayed on the official website.
	//
	// example:
	//
	// ****
	OriginalStandAmount *float64 `json:"OriginalStandAmount,omitempty" xml:"OriginalStandAmount,omitempty"`
	// The start time of the activity.
	//
	// example:
	//
	// 2024-09-23T14:00:01Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) GetCheapRate() *float64 {
	return s.CheapRate
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) GetCheapStandAmount() *float64 {
	return s.CheapStandAmount
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) GetDifferential() *float64 {
	return s.Differential
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) GetDifferentialName() *string {
	return s.DifferentialName
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) GetIsContractActivity() *bool {
	return s.IsContractActivity
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) GetIsShow() *bool {
	return s.IsShow
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) GetListPrice() *float64 {
	return s.ListPrice
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) GetMonthPrice() *float64 {
	return s.MonthPrice
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) GetOriginalStandAmount() *float64 {
	return s.OriginalStandAmount
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) SetCheapRate(v float64) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo {
	s.CheapRate = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) SetCheapStandAmount(v float64) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo {
	s.CheapStandAmount = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) SetDifferential(v float64) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo {
	s.Differential = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) SetDifferentialName(v string) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo {
	s.DifferentialName = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) SetIsContractActivity(v bool) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo {
	s.IsContractActivity = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) SetIsShow(v bool) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo {
	s.IsShow = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) SetListPrice(v float64) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo {
	s.ListPrice = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) SetMonthPrice(v float64) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo {
	s.MonthPrice = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) SetOriginalStandAmount(v float64) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo {
	s.OriginalStandAmount = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) SetStartTime(v string) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo {
	s.StartTime = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceDepreciateInfo) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrs struct {
	ModuleAttr []*DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr `json:"moduleAttr,omitempty" xml:"moduleAttr,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrs) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrs) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrs) GetModuleAttr() []*DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr {
	return s.ModuleAttr
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrs) SetModuleAttr(v []*DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrs {
	s.ModuleAttr = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrs) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr struct {
	// The attribute code.
	//
	// example:
	//
	// ****
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The attribute name.
	//
	// example:
	//
	// ****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The attribute type.
	//
	// example:
	//
	// ****
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The attribute value.
	//
	// example:
	//
	// ****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr) GetCode() *string {
	return s.Code
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr) GetName() *string {
	return s.Name
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr) GetType() *int64 {
	return s.Type
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr) GetValue() *string {
	return s.Value
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr) SetCode(v string) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr {
	s.Code = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr) SetName(v string) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr {
	s.Name = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr) SetType(v int64) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr {
	s.Type = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr) SetValue(v string) *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr {
	s.Value = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderModuleInstanceModuleInstanceModuleAttrsModuleAttr) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodySubOrdersSubOrderOptionalPromotions struct {
	OptionalPromotion []*DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion `json:"OptionalPromotion,omitempty" xml:"OptionalPromotion,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodySubOrdersSubOrderOptionalPromotions) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrderOptionalPromotions) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotions) GetOptionalPromotion() []*DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion {
	return s.OptionalPromotion
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotions) SetOptionalPromotion(v []*DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotions {
	s.OptionalPromotion = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotions) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion struct {
	// The additional activity information.
	//
	// example:
	//
	// ****
	ActivityExtInfo map[string]interface{} `json:"ActivityExtInfo,omitempty" xml:"ActivityExtInfo,omitempty"`
	// The amount that can be deducted by using the coupon.
	//
	// example:
	//
	// ****
	CanPromFee *string `json:"CanPromFee,omitempty" xml:"CanPromFee,omitempty"`
	// Indicates whether coupons are used. Valid values:
	//
	// 	- default or null (default): Coupons are used.
	//
	// 	- youhuiquan_promotion_option_id_for_blank: Coupons are not used.
	//
	// example:
	//
	// default
	CouponNo *string `json:"CouponNo,omitempty" xml:"CouponNo,omitempty"`
	// The description of the coupon.
	//
	// example:
	//
	// ****
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The coupon name.
	//
	// example:
	//
	// ****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The promotion code.
	//
	// example:
	//
	// ****
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The promotion name.
	//
	// example:
	//
	// ****
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// youhuiquan_promotion_option_id_for_blank
	PromotionOptionNo *string `json:"PromotionOptionNo,omitempty" xml:"PromotionOptionNo,omitempty"`
	// Indicates whether the promotion option is selected.
	//
	// example:
	//
	// false
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
	// Indicates whether the discount is displayed.
	//
	// example:
	//
	// False
	Show *bool `json:"Show,omitempty" xml:"Show,omitempty"`
}

func (s DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) GetActivityExtInfo() map[string]interface{} {
	return s.ActivityExtInfo
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) GetCanPromFee() *string {
	return s.CanPromFee
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) GetCouponNo() *string {
	return s.CouponNo
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) GetDescription() *string {
	return s.Description
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) GetName() *string {
	return s.Name
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) GetPromotionOptionNo() *string {
	return s.PromotionOptionNo
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) GetSelected() *bool {
	return s.Selected
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) GetShow() *bool {
	return s.Show
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) SetActivityExtInfo(v map[string]interface{}) *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion {
	s.ActivityExtInfo = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) SetCanPromFee(v string) *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion {
	s.CanPromFee = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) SetCouponNo(v string) *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion {
	s.CouponNo = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) SetDescription(v string) *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion {
	s.Description = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) SetName(v string) *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion {
	s.Name = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) SetOptionCode(v string) *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion {
	s.OptionCode = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) SetPromotionName(v string) *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion {
	s.PromotionName = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) SetPromotionOptionNo(v string) *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion {
	s.PromotionOptionNo = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) SetSelected(v bool) *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion {
	s.Selected = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) SetShow(v bool) *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion {
	s.Show = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderOptionalPromotionsOptionalPromotion) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodySubOrdersSubOrderPromDetailList struct {
	PromDetail []*DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail `json:"PromDetail,omitempty" xml:"PromDetail,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodySubOrdersSubOrderPromDetailList) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrderPromDetailList) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailList) GetPromDetail() []*DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail {
	return s.PromDetail
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailList) SetPromDetail(v []*DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) *DescribePriceResponseBodySubOrdersSubOrderPromDetailList {
	s.PromDetail = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailList) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail struct {
	// The additional activity information.
	//
	// example:
	//
	// ****
	ActivityExtInfo map[string]interface{} `json:"ActivityExtInfo,omitempty" xml:"ActivityExtInfo,omitempty"`
	// The derived promotion type.
	//
	// example:
	//
	// ****
	DerivedPromType *string `json:"DerivedPromType,omitempty" xml:"DerivedPromType,omitempty"`
	// The discount amount.
	//
	// example:
	//
	// ****
	FinalPromFee *float64 `json:"FinalPromFee,omitempty" xml:"FinalPromFee,omitempty"`
	// The coupon code.
	//
	// example:
	//
	// ****
	OptionCode *string `json:"OptionCode,omitempty" xml:"OptionCode,omitempty"`
	// The promotion type.
	//
	// example:
	//
	// ****
	PromType *string `json:"PromType,omitempty" xml:"PromType,omitempty"`
	// The promotion code.
	//
	// example:
	//
	// ****
	PromotionCode *string `json:"PromotionCode,omitempty" xml:"PromotionCode,omitempty"`
	// The promotion ID.
	//
	// example:
	//
	// ****
	PromotionId *int64 `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The promotion name.
	//
	// example:
	//
	// ****
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
}

func (s DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) GetActivityExtInfo() map[string]interface{} {
	return s.ActivityExtInfo
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) GetDerivedPromType() *string {
	return s.DerivedPromType
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) GetFinalPromFee() *float64 {
	return s.FinalPromFee
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) GetOptionCode() *string {
	return s.OptionCode
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) GetPromType() *string {
	return s.PromType
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) GetPromotionCode() *string {
	return s.PromotionCode
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) GetPromotionId() *int64 {
	return s.PromotionId
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) GetPromotionName() *string {
	return s.PromotionName
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) SetActivityExtInfo(v map[string]interface{}) *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail {
	s.ActivityExtInfo = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) SetDerivedPromType(v string) *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail {
	s.DerivedPromType = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) SetFinalPromFee(v float64) *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail {
	s.FinalPromFee = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) SetOptionCode(v string) *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail {
	s.OptionCode = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) SetPromType(v string) *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail {
	s.PromType = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) SetPromotionCode(v string) *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail {
	s.PromotionCode = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) SetPromotionId(v int64) *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail {
	s.PromotionId = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) SetPromotionName(v string) *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail {
	s.PromotionName = &v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderPromDetailListPromDetail) Validate() error {
	return dara.Validate(s)
}

type DescribePriceResponseBodySubOrdersSubOrderRuleIds struct {
	RuleId []*string `json:"RuleId,omitempty" xml:"RuleId,omitempty" type:"Repeated"`
}

func (s DescribePriceResponseBodySubOrdersSubOrderRuleIds) String() string {
	return dara.Prettify(s)
}

func (s DescribePriceResponseBodySubOrdersSubOrderRuleIds) GoString() string {
	return s.String()
}

func (s *DescribePriceResponseBodySubOrdersSubOrderRuleIds) GetRuleId() []*string {
	return s.RuleId
}

func (s *DescribePriceResponseBodySubOrdersSubOrderRuleIds) SetRuleId(v []*string) *DescribePriceResponseBodySubOrdersSubOrderRuleIds {
	s.RuleId = v
	return s
}

func (s *DescribePriceResponseBodySubOrdersSubOrderRuleIds) Validate() error {
	return dara.Validate(s)
}
