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
	// example:
	//
	// 0
	AccountQuantity *int64 `json:"AccountQuantity,omitempty" xml:"AccountQuantity,omitempty"`
	// example:
	//
	// 190311111111****
	AliUid     *int64      `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Components interface{} `json:"Components,omitempty" xml:"Components,omitempty"`
	// example:
	//
	// 0.0
	CouponPrice *float32 `json:"CouponPrice,omitempty" xml:"CouponPrice,omitempty"`
	// example:
	//
	// 1531191564000
	CreatedOn *int64 `json:"CreatedOn,omitempty" xml:"CreatedOn,omitempty"`
	// List
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// 202211111111111
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// NORMAL
	OrderStatus *string `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// example:
	//
	// NEW
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// example:
	//
	// 10.0
	OriginalPrice *float32 `json:"OriginalPrice,omitempty" xml:"OriginalPrice,omitempty"`
	// example:
	//
	// 1531191675000
	PaidOn *int64 `json:"PaidOn,omitempty" xml:"PaidOn,omitempty"`
	// example:
	//
	// PAID
	PayStatus *string `json:"PayStatus,omitempty" xml:"PayStatus,omitempty"`
	// example:
	//
	// 0.0
	PaymentPrice *float32 `json:"PaymentPrice,omitempty" xml:"PaymentPrice,omitempty"`
	// example:
	//
	// MONTH
	PeriodType *string `json:"PeriodType,omitempty" xml:"PeriodType,omitempty"`
	// example:
	//
	// cmgj02****
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// cmgj02****-prepay
	ProductSkuCode *string `json:"ProductSkuCode,omitempty" xml:"ProductSkuCode,omitempty"`
	// example:
	//
	// 1
	Quantity *int32 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	// example:
	//
	// 6EF60BEC-****-****-****-270359FB54A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
