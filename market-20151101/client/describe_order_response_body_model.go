// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountQuantity(v int64) *DescribeOrderResponseBody
	GetAccountQuantity() *int64
	SetAliUid(v int64) *DescribeOrderResponseBody
	GetAliUid() *int64
	SetComponents(v map[string]interface{}) *DescribeOrderResponseBody
	GetComponents() map[string]interface{}
	SetCouponPrice(v float32) *DescribeOrderResponseBody
	GetCouponPrice() *float32
	SetCreatedOn(v int64) *DescribeOrderResponseBody
	GetCreatedOn() *int64
	SetInstanceIds(v *DescribeOrderResponseBodyInstanceIds) *DescribeOrderResponseBody
	GetInstanceIds() *DescribeOrderResponseBodyInstanceIds
	SetOrderId(v int64) *DescribeOrderResponseBody
	GetOrderId() *int64
	SetOrderStatus(v string) *DescribeOrderResponseBody
	GetOrderStatus() *string
	SetOrderType(v string) *DescribeOrderResponseBody
	GetOrderType() *string
	SetOriginalPrice(v float32) *DescribeOrderResponseBody
	GetOriginalPrice() *float32
	SetPaidOn(v int64) *DescribeOrderResponseBody
	GetPaidOn() *int64
	SetPayStatus(v string) *DescribeOrderResponseBody
	GetPayStatus() *string
	SetPaymentPrice(v float32) *DescribeOrderResponseBody
	GetPaymentPrice() *float32
	SetPeriodType(v string) *DescribeOrderResponseBody
	GetPeriodType() *string
	SetProductCode(v string) *DescribeOrderResponseBody
	GetProductCode() *string
	SetProductName(v string) *DescribeOrderResponseBody
	GetProductName() *string
	SetProductSkuCode(v string) *DescribeOrderResponseBody
	GetProductSkuCode() *string
	SetQuantity(v int32) *DescribeOrderResponseBody
	GetQuantity() *int32
	SetRequestId(v string) *DescribeOrderResponseBody
	GetRequestId() *string
	SetSupplierCompanyName(v string) *DescribeOrderResponseBody
	GetSupplierCompanyName() *string
	SetSupplierTelephones(v *DescribeOrderResponseBodySupplierTelephones) *DescribeOrderResponseBody
	GetSupplierTelephones() *DescribeOrderResponseBodySupplierTelephones
	SetTotalPrice(v float32) *DescribeOrderResponseBody
	GetTotalPrice() *float32
}

type DescribeOrderResponseBody struct {
	// The number of accounts.
	//
	// example:
	//
	// 0
	AccountQuantity *int64 `json:"AccountQuantity,omitempty" xml:"AccountQuantity,omitempty"`
	// The unique identifier of the Alibaba Cloud user (aliyunPK).
	//
	// example:
	//
	// 190311111111****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The order component information. The metric description for the parameters in the example is as follows:
	//
	// - package_quantity: the number of usage times (a component specific to API products)
	//
	// - package_version: the version
	//
	// - ord_time: the duration
	//
	// - order_num: the quantity
	//
	// - code: the component code
	//
	// - globarKey: the component key
	//
	// - name: the component name
	//
	// - value: the component value.
	//
	// example:
	//
	// { "package_quantity": { "code": "package_quantity", "globarKey": "cmapi014302_package_quantity_package_quantity_10000", "name": "10,000 times", "value": "10000" }, "package_version": { "code": "package_version", "globarKey": "cmapi014302_package_version_package_version_yuncode830200000", "name": "0 yuan/10,000 times", "value": "yuncode830200000" }, "ord_time": { "code": "ord_time", "globarKey": "cmapi014302_ord_time_ord_time_6:Month", "name": "6 months", "value": "6:Month" }, "order_num": { "code": "order_num", "globarKey": "cmapi014302_order_num_order_num_1", "name": "1", "value": "1" } }
	Components map[string]interface{} `json:"Components,omitempty" xml:"Components,omitempty"`
	// The amount of vouchers used, in CNY.
	//
	// example:
	//
	// 0.0
	CouponPrice *float32 `json:"CouponPrice,omitempty" xml:"CouponPrice,omitempty"`
	// The time when the order was placed.
	//
	// example:
	//
	// 1531191564000
	CreatedOn   *int64                                `json:"CreatedOn,omitempty" xml:"CreatedOn,omitempty"`
	InstanceIds *DescribeOrderResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	// The order ID.
	//
	// example:
	//
	// 202211111111111
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The order status. Valid values:
	//
	// - NORMAL: Normal.
	//
	// - REFUND: Refunded.
	//
	// - DELETE: Voided.
	//
	// example:
	//
	// NORMAL
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// The order type. Valid values:
	//
	// - NEW: New purchase order.
	//
	// - RENEW: Renewal order.
	//
	// - TRIAL: Trial order.
	//
	// example:
	//
	// NEW
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The original price of the order, in CNY.
	//
	// example:
	//
	// 10.0
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// The payment time.
	//
	// example:
	//
	// 1531191675000
	PaidOn *int64 `json:"PaidOn,omitempty" xml:"PaidOn,omitempty"`
	// The payment status. Valid values:
	//
	// - PAID: Paid.
	//
	// - UNPAID: Unpaid.
	//
	// example:
	//
	// PAID
	PayStatus *string `json:"PayStatus,omitempty" xml:"PayStatus,omitempty"`
	// The actual payment amount, in CNY.
	//
	// example:
	//
	// 0.0
	PaymentPrice *float32 `json:"PaymentPrice,omitempty" xml:"PaymentPrice,omitempty"`
	// The time unit for prepaid orders. Valid values:
	//
	// - MONTH: one month
	//
	// - SEASON: one quarter
	//
	// - HALFYEAR: half a year
	//
	// - YEAR: one year
	//
	// - TWOYEARS: two years
	//
	// - ONCE: one-time.
	//
	// example:
	//
	// MONTH
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// The commodity code.
	//
	// example:
	//
	// cmgj02****
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The name of the subscribed product.
	//
	// example:
	//
	// SaaS service product
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The specification code of the subscribed product.
	//
	// example:
	//
	// cmgj02****-prepay
	ProductSkuCode *string `json:"ProductSkuCode,omitempty" xml:"ProductSkuCode,omitempty"`
	// The quantity of the subscribed product.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6EF60BEC-0242-43AF-BB20-270359FB54A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The company name of the seller.
	//
	// example:
	//
	// *	- Technology Co., Ltd.
	SupplierCompanyName *string                                      `json:"SupplierCompanyName,omitempty" xml:"SupplierCompanyName,omitempty"`
	SupplierTelephones  *DescribeOrderResponseBodySupplierTelephones `json:"SupplierTelephones,omitempty" xml:"SupplierTelephones,omitempty" type:"Struct"`
	// The total payable amount of the order, in CNY.
	//
	// example:
	//
	// 0.0
	TotalPrice *float32 `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
}

func (s DescribeOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrderResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOrderResponseBody) GetAccountQuantity() *int64 {
	return s.AccountQuantity
}

func (s *DescribeOrderResponseBody) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeOrderResponseBody) GetComponents() map[string]interface{} {
	return s.Components
}

func (s *DescribeOrderResponseBody) GetCouponPrice() *float32 {
	return s.CouponPrice
}

func (s *DescribeOrderResponseBody) GetCreatedOn() *int64 {
	return s.CreatedOn
}

func (s *DescribeOrderResponseBody) GetInstanceIds() *DescribeOrderResponseBodyInstanceIds {
	return s.InstanceIds
}

func (s *DescribeOrderResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DescribeOrderResponseBody) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *DescribeOrderResponseBody) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeOrderResponseBody) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeOrderResponseBody) GetPaidOn() *int64 {
	return s.PaidOn
}

func (s *DescribeOrderResponseBody) GetPayStatus() *string {
	return s.PayStatus
}

func (s *DescribeOrderResponseBody) GetPaymentPrice() *float32 {
	return s.PaymentPrice
}

func (s *DescribeOrderResponseBody) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeOrderResponseBody) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeOrderResponseBody) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeOrderResponseBody) GetProductSkuCode() *string {
	return s.ProductSkuCode
}

func (s *DescribeOrderResponseBody) GetQuantity() *int32 {
	return s.Quantity
}

func (s *DescribeOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOrderResponseBody) GetSupplierCompanyName() *string {
	return s.SupplierCompanyName
}

func (s *DescribeOrderResponseBody) GetSupplierTelephones() *DescribeOrderResponseBodySupplierTelephones {
	return s.SupplierTelephones
}

func (s *DescribeOrderResponseBody) GetTotalPrice() *float32 {
	return s.TotalPrice
}

func (s *DescribeOrderResponseBody) SetAccountQuantity(v int64) *DescribeOrderResponseBody {
	s.AccountQuantity = &v
	return s
}

func (s *DescribeOrderResponseBody) SetAliUid(v int64) *DescribeOrderResponseBody {
	s.AliUid = &v
	return s
}

func (s *DescribeOrderResponseBody) SetComponents(v map[string]interface{}) *DescribeOrderResponseBody {
	s.Components = v
	return s
}

func (s *DescribeOrderResponseBody) SetCouponPrice(v float32) *DescribeOrderResponseBody {
	s.CouponPrice = &v
	return s
}

func (s *DescribeOrderResponseBody) SetCreatedOn(v int64) *DescribeOrderResponseBody {
	s.CreatedOn = &v
	return s
}

func (s *DescribeOrderResponseBody) SetInstanceIds(v *DescribeOrderResponseBodyInstanceIds) *DescribeOrderResponseBody {
	s.InstanceIds = v
	return s
}

func (s *DescribeOrderResponseBody) SetOrderId(v int64) *DescribeOrderResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribeOrderResponseBody) SetOrderStatus(v string) *DescribeOrderResponseBody {
	s.OrderStatus = &v
	return s
}

func (s *DescribeOrderResponseBody) SetOrderType(v string) *DescribeOrderResponseBody {
	s.OrderType = &v
	return s
}

func (s *DescribeOrderResponseBody) SetOriginalPrice(v float32) *DescribeOrderResponseBody {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeOrderResponseBody) SetPaidOn(v int64) *DescribeOrderResponseBody {
	s.PaidOn = &v
	return s
}

func (s *DescribeOrderResponseBody) SetPayStatus(v string) *DescribeOrderResponseBody {
	s.PayStatus = &v
	return s
}

func (s *DescribeOrderResponseBody) SetPaymentPrice(v float32) *DescribeOrderResponseBody {
	s.PaymentPrice = &v
	return s
}

func (s *DescribeOrderResponseBody) SetPeriodType(v string) *DescribeOrderResponseBody {
	s.PeriodType = &v
	return s
}

func (s *DescribeOrderResponseBody) SetProductCode(v string) *DescribeOrderResponseBody {
	s.ProductCode = &v
	return s
}

func (s *DescribeOrderResponseBody) SetProductName(v string) *DescribeOrderResponseBody {
	s.ProductName = &v
	return s
}

func (s *DescribeOrderResponseBody) SetProductSkuCode(v string) *DescribeOrderResponseBody {
	s.ProductSkuCode = &v
	return s
}

func (s *DescribeOrderResponseBody) SetQuantity(v int32) *DescribeOrderResponseBody {
	s.Quantity = &v
	return s
}

func (s *DescribeOrderResponseBody) SetRequestId(v string) *DescribeOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOrderResponseBody) SetSupplierCompanyName(v string) *DescribeOrderResponseBody {
	s.SupplierCompanyName = &v
	return s
}

func (s *DescribeOrderResponseBody) SetSupplierTelephones(v *DescribeOrderResponseBodySupplierTelephones) *DescribeOrderResponseBody {
	s.SupplierTelephones = v
	return s
}

func (s *DescribeOrderResponseBody) SetTotalPrice(v float32) *DescribeOrderResponseBody {
	s.TotalPrice = &v
	return s
}

func (s *DescribeOrderResponseBody) Validate() error {
	if s.InstanceIds != nil {
		if err := s.InstanceIds.Validate(); err != nil {
			return err
		}
	}
	if s.SupplierTelephones != nil {
		if err := s.SupplierTelephones.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOrderResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s DescribeOrderResponseBodyInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrderResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *DescribeOrderResponseBodyInstanceIds) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DescribeOrderResponseBodyInstanceIds) SetInstanceId(v []*string) *DescribeOrderResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

func (s *DescribeOrderResponseBodyInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DescribeOrderResponseBodySupplierTelephones struct {
	Telephone []*string `json:"Telephone,omitempty" xml:"Telephone,omitempty" type:"Repeated"`
}

func (s DescribeOrderResponseBodySupplierTelephones) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrderResponseBodySupplierTelephones) GoString() string {
	return s.String()
}

func (s *DescribeOrderResponseBodySupplierTelephones) GetTelephone() []*string {
	return s.Telephone
}

func (s *DescribeOrderResponseBodySupplierTelephones) SetTelephone(v []*string) *DescribeOrderResponseBodySupplierTelephones {
	s.Telephone = v
	return s
}

func (s *DescribeOrderResponseBodySupplierTelephones) Validate() error {
	return dara.Validate(s)
}
