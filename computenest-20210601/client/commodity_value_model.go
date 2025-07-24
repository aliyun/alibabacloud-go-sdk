// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommodityValue interface {
	dara.Model
	String() string
	GoString() string
	SetResult(v *CommodityValueResult) *CommodityValue
	GetResult() *CommodityValueResult
}

type CommodityValue struct {
	// Result模型。
	Result *CommodityValueResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s CommodityValue) String() string {
	return dara.Prettify(s)
}

func (s CommodityValue) GoString() string {
	return s.String()
}

func (s *CommodityValue) GetResult() *CommodityValueResult {
	return s.Result
}

func (s *CommodityValue) SetResult(v *CommodityValueResult) *CommodityValue {
	s.Result = v
	return s
}

func (s *CommodityValue) Validate() error {
	return dara.Validate(s)
}

type CommodityValueResult struct {
	// 订单信息。
	Order *CommodityValueResultOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	// 询价类型，可选值：
	//
	// 1. Buy：新购询价。
	//
	// 2. ModificationBuy：变配询价。
	//
	// example:
	//
	// Buy
	InquiryType *string `json:"InquiryType,omitempty" xml:"InquiryType,omitempty"`
	// 订单子项。
	SubOrders *CommodityValueResultSubOrders `json:"SubOrders,omitempty" xml:"SubOrders,omitempty" type:"Struct"`
	// 优惠券。
	Coupons []*CommodityValueResultCoupons `json:"Coupons,omitempty" xml:"Coupons,omitempty" type:"Repeated"`
}

func (s CommodityValueResult) String() string {
	return dara.Prettify(s)
}

func (s CommodityValueResult) GoString() string {
	return s.String()
}

func (s *CommodityValueResult) GetOrder() *CommodityValueResultOrder {
	return s.Order
}

func (s *CommodityValueResult) GetInquiryType() *string {
	return s.InquiryType
}

func (s *CommodityValueResult) GetSubOrders() *CommodityValueResultSubOrders {
	return s.SubOrders
}

func (s *CommodityValueResult) GetCoupons() []*CommodityValueResultCoupons {
	return s.Coupons
}

func (s *CommodityValueResult) SetOrder(v *CommodityValueResultOrder) *CommodityValueResult {
	s.Order = v
	return s
}

func (s *CommodityValueResult) SetInquiryType(v string) *CommodityValueResult {
	s.InquiryType = &v
	return s
}

func (s *CommodityValueResult) SetSubOrders(v *CommodityValueResultSubOrders) *CommodityValueResult {
	s.SubOrders = v
	return s
}

func (s *CommodityValueResult) SetCoupons(v []*CommodityValueResultCoupons) *CommodityValueResult {
	s.Coupons = v
	return s
}

func (s *CommodityValueResult) Validate() error {
	return dara.Validate(s)
}

type CommodityValueResultOrder struct {
	// 货币代码。
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// 优惠后。
	//
	// example:
	//
	// 9.99
	TradeAmount *string `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
	// 抵扣金额。
	//
	// example:
	//
	// 1.99
	DiscountAmount *string `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// 优惠前。
	//
	// example:
	//
	// 11.98
	OriginalAmount *string `json:"OriginalAmount,omitempty" xml:"OriginalAmount,omitempty"`
}

func (s CommodityValueResultOrder) String() string {
	return dara.Prettify(s)
}

func (s CommodityValueResultOrder) GoString() string {
	return s.String()
}

func (s *CommodityValueResultOrder) GetCurrency() *string {
	return s.Currency
}

func (s *CommodityValueResultOrder) GetTradeAmount() *string {
	return s.TradeAmount
}

func (s *CommodityValueResultOrder) GetDiscountAmount() *string {
	return s.DiscountAmount
}

func (s *CommodityValueResultOrder) GetOriginalAmount() *string {
	return s.OriginalAmount
}

func (s *CommodityValueResultOrder) SetCurrency(v string) *CommodityValueResultOrder {
	s.Currency = &v
	return s
}

func (s *CommodityValueResultOrder) SetTradeAmount(v string) *CommodityValueResultOrder {
	s.TradeAmount = &v
	return s
}

func (s *CommodityValueResultOrder) SetDiscountAmount(v string) *CommodityValueResultOrder {
	s.DiscountAmount = &v
	return s
}

func (s *CommodityValueResultOrder) SetOriginalAmount(v string) *CommodityValueResultOrder {
	s.OriginalAmount = &v
	return s
}

func (s *CommodityValueResultOrder) Validate() error {
	return dara.Validate(s)
}

type CommodityValueResultSubOrders struct {
	// 订单子项。
	SubOrder []*CommodityValueResultSubOrdersSubOrder `json:"SubOrder,omitempty" xml:"SubOrder,omitempty" type:"Repeated"`
}

func (s CommodityValueResultSubOrders) String() string {
	return dara.Prettify(s)
}

func (s CommodityValueResultSubOrders) GoString() string {
	return s.String()
}

func (s *CommodityValueResultSubOrders) GetSubOrder() []*CommodityValueResultSubOrdersSubOrder {
	return s.SubOrder
}

func (s *CommodityValueResultSubOrders) SetSubOrder(v []*CommodityValueResultSubOrdersSubOrder) *CommodityValueResultSubOrders {
	s.SubOrder = v
	return s
}

func (s *CommodityValueResultSubOrders) Validate() error {
	return dara.Validate(s)
}

type CommodityValueResultSubOrdersSubOrder struct {
	// 模块（实例）信息。
	ModuleInstance []*CommodityValueResultSubOrdersSubOrderModuleInstance `json:"ModuleInstance,omitempty" xml:"ModuleInstance,omitempty" type:"Repeated"`
}

func (s CommodityValueResultSubOrdersSubOrder) String() string {
	return dara.Prettify(s)
}

func (s CommodityValueResultSubOrdersSubOrder) GoString() string {
	return s.String()
}

func (s *CommodityValueResultSubOrdersSubOrder) GetModuleInstance() []*CommodityValueResultSubOrdersSubOrderModuleInstance {
	return s.ModuleInstance
}

func (s *CommodityValueResultSubOrdersSubOrder) SetModuleInstance(v []*CommodityValueResultSubOrdersSubOrderModuleInstance) *CommodityValueResultSubOrdersSubOrder {
	s.ModuleInstance = v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrder) Validate() error {
	return dara.Validate(s)
}

type CommodityValueResultSubOrdersSubOrderModuleInstance struct {
	// 模块ID。
	//
	// example:
	//
	// 1234
	ModuleId *int64 `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// 模块名称。
	//
	// example:
	//
	// Rds
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// 模块代码。
	//
	// example:
	//
	// rds_dbtype
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// 产品原价（元）。
	//
	// example:
	//
	// 10.00
	TotalProductFee *float64 `json:"TotalProductFee,omitempty" xml:"TotalProductFee,omitempty"`
	// 折扣费用（元）。
	//
	// example:
	//
	// 1.99
	DiscountFee *float64 `json:"DiscountFee,omitempty" xml:"DiscountFee,omitempty"`
	// 实付金额（元）。
	//
	// example:
	//
	// 8.01
	PayFee *float64 `json:"PayFee,omitempty" xml:"PayFee,omitempty"`
	// 价格单位。
	//
	// example:
	//
	// 元/GB/小时
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// 是否计价项。
	//
	// example:
	//
	// true
	IsPricingModule *bool `json:"IsPricingModule,omitempty" xml:"IsPricingModule,omitempty"`
	// 在订单中是否需要支付。
	//
	// example:
	//
	// true
	NeedOrderPay *bool `json:"NeedOrderPay,omitempty" xml:"NeedOrderPay,omitempty"`
	// 定价类型。
	//
	// example:
	//
	// hourPrice
	PriceType *string `json:"PriceType,omitempty" xml:"PriceType,omitempty"`
	// 模块属性。
	ModuleAttrs []*CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs `json:"ModuleAttrs,omitempty" xml:"ModuleAttrs,omitempty" type:"Repeated"`
}

func (s CommodityValueResultSubOrdersSubOrderModuleInstance) String() string {
	return dara.Prettify(s)
}

func (s CommodityValueResultSubOrdersSubOrderModuleInstance) GoString() string {
	return s.String()
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) GetModuleId() *int64 {
	return s.ModuleId
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) GetModuleName() *string {
	return s.ModuleName
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) GetTotalProductFee() *float64 {
	return s.TotalProductFee
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) GetDiscountFee() *float64 {
	return s.DiscountFee
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) GetPayFee() *float64 {
	return s.PayFee
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) GetPriceUnit() *string {
	return s.PriceUnit
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) GetIsPricingModule() *bool {
	return s.IsPricingModule
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) GetNeedOrderPay() *bool {
	return s.NeedOrderPay
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) GetPriceType() *string {
	return s.PriceType
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) GetModuleAttrs() []*CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs {
	return s.ModuleAttrs
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetModuleId(v int64) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.ModuleId = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetModuleName(v string) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.ModuleName = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetModuleCode(v string) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.ModuleCode = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetTotalProductFee(v float64) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.TotalProductFee = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetDiscountFee(v float64) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.DiscountFee = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetPayFee(v float64) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.PayFee = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetPriceUnit(v string) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.PriceUnit = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetIsPricingModule(v bool) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.IsPricingModule = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetNeedOrderPay(v bool) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.NeedOrderPay = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetPriceType(v string) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.PriceType = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetModuleAttrs(v []*CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.ModuleAttrs = v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) Validate() error {
	return dara.Validate(s)
}

type CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs struct {
	// 属性类型，可选值：
	//
	// 1. 1：商品属性
	//
	// 2. 2：规格属性
	//
	// 3. 3：模块属性
	//
	// 4. 4：外部参数（备用）
	//
	// example:
	//
	// 3
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// Name
	//
	// example:
	//
	// 20GB
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Module attr code
	//
	// example:
	//
	// rds_storage
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Value
	//
	// example:
	//
	// 20
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// Unit
	//
	// example:
	//
	// GB
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) String() string {
	return dara.Prettify(s)
}

func (s CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) GoString() string {
	return s.String()
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) GetType() *int64 {
	return s.Type
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) GetName() *string {
	return s.Name
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) GetCode() *string {
	return s.Code
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) GetValue() *string {
	return s.Value
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) GetUnit() *string {
	return s.Unit
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) SetType(v int64) *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs {
	s.Type = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) SetName(v string) *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs {
	s.Name = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) SetCode(v string) *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs {
	s.Code = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) SetValue(v string) *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs {
	s.Value = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) SetUnit(v string) *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs {
	s.Unit = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs) Validate() error {
	return dara.Validate(s)
}

type CommodityValueResultCoupons struct {
	// 可支付金额。
	//
	// example:
	//
	// 9.99
	CanPromFee *float64 `json:"CanPromFee,omitempty" xml:"CanPromFee,omitempty"`
	// Coupon Description
	//
	// example:
	//
	// 10元优惠券（有效期至2024年9月8日）
	CouponDesc *string `json:"CouponDesc,omitempty" xml:"CouponDesc,omitempty"`
	// Coupon Name
	//
	// example:
	//
	// 10元优惠券
	CouponName *string `json:"CouponName,omitempty" xml:"CouponName,omitempty"`
	// Coupon OptionNo
	//
	// example:
	//
	// 50008800000xxxx
	CouponOptionNo *string `json:"CouponOptionNo,omitempty" xml:"CouponOptionNo,omitempty"`
	// 是否选中。
	//
	// example:
	//
	// true
	Selected *bool `json:"Selected,omitempty" xml:"Selected,omitempty"`
}

func (s CommodityValueResultCoupons) String() string {
	return dara.Prettify(s)
}

func (s CommodityValueResultCoupons) GoString() string {
	return s.String()
}

func (s *CommodityValueResultCoupons) GetCanPromFee() *float64 {
	return s.CanPromFee
}

func (s *CommodityValueResultCoupons) GetCouponDesc() *string {
	return s.CouponDesc
}

func (s *CommodityValueResultCoupons) GetCouponName() *string {
	return s.CouponName
}

func (s *CommodityValueResultCoupons) GetCouponOptionNo() *string {
	return s.CouponOptionNo
}

func (s *CommodityValueResultCoupons) GetSelected() *bool {
	return s.Selected
}

func (s *CommodityValueResultCoupons) SetCanPromFee(v float64) *CommodityValueResultCoupons {
	s.CanPromFee = &v
	return s
}

func (s *CommodityValueResultCoupons) SetCouponDesc(v string) *CommodityValueResultCoupons {
	s.CouponDesc = &v
	return s
}

func (s *CommodityValueResultCoupons) SetCouponName(v string) *CommodityValueResultCoupons {
	s.CouponName = &v
	return s
}

func (s *CommodityValueResultCoupons) SetCouponOptionNo(v string) *CommodityValueResultCoupons {
	s.CouponOptionNo = &v
	return s
}

func (s *CommodityValueResultCoupons) SetSelected(v bool) *CommodityValueResultCoupons {
	s.Selected = &v
	return s
}

func (s *CommodityValueResultCoupons) Validate() error {
	return dara.Validate(s)
}
