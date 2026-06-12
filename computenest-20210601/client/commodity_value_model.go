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
	// The Result model.
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
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CommodityValueResult struct {
	// The order information.
	Order *CommodityValueResultOrder `json:"Order,omitempty" xml:"Order,omitempty" type:"Struct"`
	// The inquiry type. Valid values:
	//
	// 1. Buy: an inquiry for a new purchase.
	//
	// 2. ModificationBuy: an inquiry for an upgrade or downgrade.
	//
	// example:
	//
	// Buy
	InquiryType *string `json:"InquiryType,omitempty" xml:"InquiryType,omitempty"`
	// The sub-orders.
	SubOrders *CommodityValueResultSubOrders `json:"SubOrders,omitempty" xml:"SubOrders,omitempty" type:"Struct"`
	// The coupons.
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
	if s.Order != nil {
		if err := s.Order.Validate(); err != nil {
			return err
		}
	}
	if s.SubOrders != nil {
		if err := s.SubOrders.Validate(); err != nil {
			return err
		}
	}
	if s.Coupons != nil {
		for _, item := range s.Coupons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CommodityValueResultOrder struct {
	// The currency code.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The amount after the discount is applied.
	//
	// example:
	//
	// 9.99
	TradeAmount *string `json:"TradeAmount,omitempty" xml:"TradeAmount,omitempty"`
	// The discount amount.
	//
	// example:
	//
	// 1.99
	DiscountAmount *string `json:"DiscountAmount,omitempty" xml:"DiscountAmount,omitempty"`
	// The original amount.
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
	// The sub-orders.
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

type CommodityValueResultSubOrdersSubOrder struct {
	// The module or instance information.
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
	if s.ModuleInstance != nil {
		for _, item := range s.ModuleInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CommodityValueResultSubOrdersSubOrderModuleInstance struct {
	// The module ID.
	//
	// example:
	//
	// 1234
	ModuleId *int64 `json:"ModuleId,omitempty" xml:"ModuleId,omitempty"`
	// The module name.
	//
	// example:
	//
	// 数据库类型
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The module code.
	//
	// example:
	//
	// rds_dbtype
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// The original price of the product, in CNY.
	//
	// example:
	//
	// 10.00
	TotalProductFee *float64 `json:"TotalProductFee,omitempty" xml:"TotalProductFee,omitempty"`
	// The discount, in CNY.
	//
	// example:
	//
	// 1.99
	DiscountFee *float64 `json:"DiscountFee,omitempty" xml:"DiscountFee,omitempty"`
	// The amount payable, in CNY.
	//
	// example:
	//
	// 8.01
	PayFee *float64 `json:"PayFee,omitempty" xml:"PayFee,omitempty"`
	// The unit of the price.
	//
	// example:
	//
	// 元/GB/小时
	PriceUnit *string `json:"PriceUnit,omitempty" xml:"PriceUnit,omitempty"`
	// Indicates whether the module is a billing item.
	//
	// example:
	//
	// true
	IsPricingModule *bool `json:"IsPricingModule,omitempty" xml:"IsPricingModule,omitempty"`
	// Indicates whether payment is required for the module in the order.
	//
	// example:
	//
	// true
	NeedOrderPay *bool `json:"NeedOrderPay,omitempty" xml:"NeedOrderPay,omitempty"`
	// The pricing type.
	//
	// example:
	//
	// hourPrice
	PriceType *string `json:"PriceType,omitempty" xml:"PriceType,omitempty"`
	// The module attributes.
	ModuleAttrs []*CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs `json:"ModuleAttrs,omitempty" xml:"ModuleAttrs,omitempty" type:"Repeated"`
	// The English name of the module.
	//
	// example:
	//
	// 数据单元
	ModuleNameEn *string `json:"ModuleNameEn,omitempty" xml:"ModuleNameEn,omitempty"`
	// The English name of the price unit.
	//
	// example:
	//
	// 元/月
	PriceUnitEn *string `json:"PriceUnitEn,omitempty" xml:"PriceUnitEn,omitempty"`
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

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) GetModuleNameEn() *string {
	return s.ModuleNameEn
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) GetPriceUnitEn() *string {
	return s.PriceUnitEn
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

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetModuleNameEn(v string) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.ModuleNameEn = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) SetPriceUnitEn(v string) *CommodityValueResultSubOrdersSubOrderModuleInstance {
	s.PriceUnitEn = &v
	return s
}

func (s *CommodityValueResultSubOrdersSubOrderModuleInstance) Validate() error {
	if s.ModuleAttrs != nil {
		for _, item := range s.ModuleAttrs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CommodityValueResultSubOrdersSubOrderModuleInstanceModuleAttrs struct {
	// The attribute type. Valid values:
	//
	// 1. 1: product attribute
	//
	// 2. 2: specification attribute
	//
	// 3. 3: module attribute
	//
	// 4. 4: external parameter (reserved)
	//
	// example:
	//
	// 3
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The attribute name.
	//
	// example:
	//
	// 20GB
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The attribute code.
	//
	// example:
	//
	// rds_storage
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The attribute value.
	//
	// example:
	//
	// 20
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The unit of the value.
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
	// The payable amount.
	//
	// example:
	//
	// 9.99
	CanPromFee *float64 `json:"CanPromFee,omitempty" xml:"CanPromFee,omitempty"`
	// The description of the coupon.
	//
	// example:
	//
	// 10元优惠券（有效期至2024年9月8日）
	CouponDesc *string `json:"CouponDesc,omitempty" xml:"CouponDesc,omitempty"`
	// The name of the coupon.
	//
	// example:
	//
	// 10元优惠券
	CouponName *string `json:"CouponName,omitempty" xml:"CouponName,omitempty"`
	// The coupon number.
	//
	// example:
	//
	// 50008800000xxxx
	CouponOptionNo *string `json:"CouponOptionNo,omitempty" xml:"CouponOptionNo,omitempty"`
	// Indicates whether the coupon is selected.
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
