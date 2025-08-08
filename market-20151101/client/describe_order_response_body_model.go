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
	// example:
	//
	// 0
	AccountQuantity *int64 `json:"AccountQuantity,omitempty" xml:"AccountQuantity,omitempty"`
	// example:
	//
	// 190311111111****
	AliUid     *int64                 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	Components map[string]interface{} `json:"Components,omitempty" xml:"Components,omitempty"`
	// example:
	//
	// 0.0
	CouponPrice *float32 `json:"CouponPrice,omitempty" xml:"CouponPrice,omitempty"`
	// example:
	//
	// 1531191564000
	CreatedOn   *int64                                `json:"CreatedOn,omitempty" xml:"CreatedOn,omitempty"`
	InstanceIds *DescribeOrderResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
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
	// 6EF60BEC-0242-43AF-BB20-270359FB54A7
	RequestId           *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SupplierCompanyName *string                                      `json:"SupplierCompanyName,omitempty" xml:"SupplierCompanyName,omitempty"`
	SupplierTelephones  *DescribeOrderResponseBodySupplierTelephones `json:"SupplierTelephones,omitempty" xml:"SupplierTelephones,omitempty" type:"Struct"`
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
	return dara.Validate(s)
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
