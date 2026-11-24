// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOrderForIsvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountQuantity(v int64) *DescribeOrderForIsvResponseBody
	GetAccountQuantity() *int64
	SetAliUid(v int64) *DescribeOrderForIsvResponseBody
	GetAliUid() *int64
	SetComponents(v interface{}) *DescribeOrderForIsvResponseBody
	GetComponents() interface{}
	SetCouponPrice(v float32) *DescribeOrderForIsvResponseBody
	GetCouponPrice() *float32
	SetCreatedOn(v int64) *DescribeOrderForIsvResponseBody
	GetCreatedOn() *int64
	SetInstanceIds(v []*string) *DescribeOrderForIsvResponseBody
	GetInstanceIds() []*string
	SetOrderId(v int64) *DescribeOrderForIsvResponseBody
	GetOrderId() *int64
	SetOrderStatus(v string) *DescribeOrderForIsvResponseBody
	GetOrderStatus() *string
	SetOrderType(v string) *DescribeOrderForIsvResponseBody
	GetOrderType() *string
	SetOriginalPrice(v float32) *DescribeOrderForIsvResponseBody
	GetOriginalPrice() *float32
	SetPaidOn(v int64) *DescribeOrderForIsvResponseBody
	GetPaidOn() *int64
	SetPayStatus(v string) *DescribeOrderForIsvResponseBody
	GetPayStatus() *string
	SetPaymentPrice(v float32) *DescribeOrderForIsvResponseBody
	GetPaymentPrice() *float32
	SetPeriodType(v string) *DescribeOrderForIsvResponseBody
	GetPeriodType() *string
	SetProductCode(v string) *DescribeOrderForIsvResponseBody
	GetProductCode() *string
	SetProductName(v string) *DescribeOrderForIsvResponseBody
	GetProductName() *string
	SetProductSkuCode(v string) *DescribeOrderForIsvResponseBody
	GetProductSkuCode() *string
	SetQuantity(v int32) *DescribeOrderForIsvResponseBody
	GetQuantity() *int32
	SetRequestId(v string) *DescribeOrderForIsvResponseBody
	GetRequestId() *string
	SetTotalPrice(v float32) *DescribeOrderForIsvResponseBody
	GetTotalPrice() *float32
}

type DescribeOrderForIsvResponseBody struct {
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
	// The order module information. The metric description for the parameters in the example is as follows:
	//
	// - package_quantity: the number of usage times (a module specific to API products)
	//
	// - package_version: the version
	//
	// - ord_time: the duration
	//
	// - order_num: the quantity
	//
	// - code: the module code
	//
	// - globarKey: the module key
	//
	// - name: the module name
	//
	// - value: the module value.
	//
	// example:
	//
	// {"package_version":{"code":"package_version","globarKey":"cmgj0******8_package_version_package_version_yuncode4176*****1","name":"0.01","value":"yuncode4176*****1"},"ord_time":{"code":"ord_time","globarKey":"cmgj0******8_ord_time_ord_time_1:Month","name":"1 month","value":"1:Month"},"order_num":{"code":"order_num","globarKey":"cmgj0******8_order_num_order_num_1","name":"1","value":"1"}}
	Components interface{} `json:"Components,omitempty" xml:"Components,omitempty"`
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
	CreatedOn *int64 `json:"CreatedOn,omitempty" xml:"CreatedOn,omitempty"`
	// List
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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
	// - UPGRADE: Upgrade order.
	//
	// - DOWNGRADE: Downgrade order.
	//
	// - PURCHASE: Conversion order.
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
	// The SKU code of the subscribed product.
	//
	// example:
	//
	// cmgj02****-prepay
	ProductSkuCode *string `json:"ProductSkuCode,omitempty" xml:"ProductSkuCode,omitempty"`
	// The quantity of the order.
	//
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// The request ID. This ID is used for troubleshooting when an error occurs.
	//
	// example:
	//
	// 6EF60BEC-****-****-****-270359FB54A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total receivable amount of the order, in CNY.
	//
	// example:
	//
	// 0.0
	TotalPrice *float32 `json:"TotalPrice,omitempty" xml:"TotalPrice,omitempty"`
}

func (s DescribeOrderForIsvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOrderForIsvResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOrderForIsvResponseBody) GetAccountQuantity() *int64 {
	return s.AccountQuantity
}

func (s *DescribeOrderForIsvResponseBody) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeOrderForIsvResponseBody) GetComponents() interface{} {
	return s.Components
}

func (s *DescribeOrderForIsvResponseBody) GetCouponPrice() *float32 {
	return s.CouponPrice
}

func (s *DescribeOrderForIsvResponseBody) GetCreatedOn() *int64 {
	return s.CreatedOn
}

func (s *DescribeOrderForIsvResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeOrderForIsvResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DescribeOrderForIsvResponseBody) GetOrderStatus() *string {
	return s.OrderStatus
}

func (s *DescribeOrderForIsvResponseBody) GetOrderType() *string {
	return s.OrderType
}

func (s *DescribeOrderForIsvResponseBody) GetOriginalPrice() *float32 {
	return s.OriginalPrice
}

func (s *DescribeOrderForIsvResponseBody) GetPaidOn() *int64 {
	return s.PaidOn
}

func (s *DescribeOrderForIsvResponseBody) GetPayStatus() *string {
	return s.PayStatus
}

func (s *DescribeOrderForIsvResponseBody) GetPaymentPrice() *float32 {
	return s.PaymentPrice
}

func (s *DescribeOrderForIsvResponseBody) GetPeriodType() *string {
	return s.PeriodType
}

func (s *DescribeOrderForIsvResponseBody) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeOrderForIsvResponseBody) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeOrderForIsvResponseBody) GetProductSkuCode() *string {
	return s.ProductSkuCode
}

func (s *DescribeOrderForIsvResponseBody) GetQuantity() *int32 {
	return s.Quantity
}

func (s *DescribeOrderForIsvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOrderForIsvResponseBody) GetTotalPrice() *float32 {
	return s.TotalPrice
}

func (s *DescribeOrderForIsvResponseBody) SetAccountQuantity(v int64) *DescribeOrderForIsvResponseBody {
	s.AccountQuantity = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetAliUid(v int64) *DescribeOrderForIsvResponseBody {
	s.AliUid = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetComponents(v interface{}) *DescribeOrderForIsvResponseBody {
	s.Components = v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetCouponPrice(v float32) *DescribeOrderForIsvResponseBody {
	s.CouponPrice = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetCreatedOn(v int64) *DescribeOrderForIsvResponseBody {
	s.CreatedOn = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetInstanceIds(v []*string) *DescribeOrderForIsvResponseBody {
	s.InstanceIds = v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetOrderId(v int64) *DescribeOrderForIsvResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetOrderStatus(v string) *DescribeOrderForIsvResponseBody {
	s.OrderStatus = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetOrderType(v string) *DescribeOrderForIsvResponseBody {
	s.OrderType = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetOriginalPrice(v float32) *DescribeOrderForIsvResponseBody {
	s.OriginalPrice = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetPaidOn(v int64) *DescribeOrderForIsvResponseBody {
	s.PaidOn = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetPayStatus(v string) *DescribeOrderForIsvResponseBody {
	s.PayStatus = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetPaymentPrice(v float32) *DescribeOrderForIsvResponseBody {
	s.PaymentPrice = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetPeriodType(v string) *DescribeOrderForIsvResponseBody {
	s.PeriodType = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetProductCode(v string) *DescribeOrderForIsvResponseBody {
	s.ProductCode = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetProductName(v string) *DescribeOrderForIsvResponseBody {
	s.ProductName = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetProductSkuCode(v string) *DescribeOrderForIsvResponseBody {
	s.ProductSkuCode = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetQuantity(v int32) *DescribeOrderForIsvResponseBody {
	s.Quantity = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetRequestId(v string) *DescribeOrderForIsvResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) SetTotalPrice(v float32) *DescribeOrderForIsvResponseBody {
	s.TotalPrice = &v
	return s
}

func (s *DescribeOrderForIsvResponseBody) Validate() error {
	return dara.Validate(s)
}
