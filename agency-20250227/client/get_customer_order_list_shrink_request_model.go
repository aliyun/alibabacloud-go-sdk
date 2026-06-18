// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerOrderListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerAccount(v string) *GetCustomerOrderListShrinkRequest
	GetCustomerAccount() *string
	SetCustomerUid(v int64) *GetCustomerOrderListShrinkRequest
	GetCustomerUid() *int64
	SetOrderCreateAfter(v int64) *GetCustomerOrderListShrinkRequest
	GetOrderCreateAfter() *int64
	SetOrderCreateBefore(v int64) *GetCustomerOrderListShrinkRequest
	GetOrderCreateBefore() *int64
	SetOrderId(v int64) *GetCustomerOrderListShrinkRequest
	GetOrderId() *int64
	SetOrderPayAfter(v int64) *GetCustomerOrderListShrinkRequest
	GetOrderPayAfter() *int64
	SetOrderPayBefore(v int64) *GetCustomerOrderListShrinkRequest
	GetOrderPayBefore() *int64
	SetOrderStatus(v int32) *GetCustomerOrderListShrinkRequest
	GetOrderStatus() *int32
	SetOrderTypeListShrink(v string) *GetCustomerOrderListShrinkRequest
	GetOrderTypeListShrink() *string
	SetPageNo(v int32) *GetCustomerOrderListShrinkRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetCustomerOrderListShrinkRequest
	GetPageSize() *int32
	SetPayAmountAfter(v float64) *GetCustomerOrderListShrinkRequest
	GetPayAmountAfter() *float64
	SetPayAmountBefore(v float64) *GetCustomerOrderListShrinkRequest
	GetPayAmountBefore() *float64
	SetPayType(v int32) *GetCustomerOrderListShrinkRequest
	GetPayType() *int32
	SetProductCode(v string) *GetCustomerOrderListShrinkRequest
	GetProductCode() *string
	SetProductName(v string) *GetCustomerOrderListShrinkRequest
	GetProductName() *string
	SetProjectId(v int64) *GetCustomerOrderListShrinkRequest
	GetProjectId() *int64
	SetRamAccountForCustomerManager(v string) *GetCustomerOrderListShrinkRequest
	GetRamAccountForCustomerManager() *string
}

type GetCustomerOrderListShrinkRequest struct {
	// The customer account.
	//
	// example:
	//
	// test_123
	CustomerAccount *string `json:"CustomerAccount,omitempty" xml:"CustomerAccount,omitempty"`
	// The customer UID.
	//
	// example:
	//
	// 123456
	CustomerUid *int64 `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// The start timestamp for order creation. The time range cannot exceed 6 months. The order creation time range and the order payment time range cannot both be empty.
	//
	// example:
	//
	// 1727789348000
	OrderCreateAfter *int64 `json:"OrderCreateAfter,omitempty" xml:"OrderCreateAfter,omitempty"`
	// The end timestamp for order creation. The time range cannot exceed 6 months. The order creation time range and the order payment time range cannot both be empty.
	//
	// example:
	//
	// 1741008566000
	OrderCreateBefore *int64 `json:"OrderCreateBefore,omitempty" xml:"OrderCreateBefore,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 13595216
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The start timestamp for order payment. The time range cannot exceed 6 months. The order creation time range and the order payment time range cannot both be empty.
	//
	// example:
	//
	// 1727789348000
	OrderPayAfter *int64 `json:"OrderPayAfter,omitempty" xml:"OrderPayAfter,omitempty"`
	// The end timestamp for order payment. The time range cannot exceed 6 months. The order creation time range and the order payment time range cannot both be empty.
	//
	// example:
	//
	// 1741008566000
	OrderPayBefore *int64 `json:"OrderPayBefore,omitempty" xml:"OrderPayBefore,omitempty"`
	// The order status. Valid values:
	//
	// - 1: unpaid
	//
	// - 2: canceled
	//
	// - 3: paid.
	//
	// example:
	//
	// 3
	OrderStatus *int32 `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// The list of order types.
	OrderTypeListShrink *string `json:"OrderTypeList,omitempty" xml:"OrderTypeList,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The minimum actual payment amount.
	//
	// example:
	//
	// 1
	PayAmountAfter *float64 `json:"PayAmountAfter,omitempty" xml:"PayAmountAfter,omitempty"`
	// The maximum actual payment amount.
	//
	// example:
	//
	// 1000
	PayAmountBefore *float64 `json:"PayAmountBefore,omitempty" xml:"PayAmountBefore,omitempty"`
	// The payment type. Valid values:
	//
	// - 1: non-delegated payment
	//
	// - 2: delegated payment.
	//
	// example:
	//
	// 1
	PayType *int32 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The product code.
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product name.
	//
	// example:
	//
	// 弹性计算
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The opportunity ID.
	//
	// example:
	//
	// 202502002231
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The employee who follows up with the customer.
	//
	// example:
	//
	// 1234532
	RamAccountForCustomerManager *string `json:"RamAccountForCustomerManager,omitempty" xml:"RamAccountForCustomerManager,omitempty"`
}

func (s GetCustomerOrderListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerOrderListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerOrderListShrinkRequest) GetCustomerAccount() *string {
	return s.CustomerAccount
}

func (s *GetCustomerOrderListShrinkRequest) GetCustomerUid() *int64 {
	return s.CustomerUid
}

func (s *GetCustomerOrderListShrinkRequest) GetOrderCreateAfter() *int64 {
	return s.OrderCreateAfter
}

func (s *GetCustomerOrderListShrinkRequest) GetOrderCreateBefore() *int64 {
	return s.OrderCreateBefore
}

func (s *GetCustomerOrderListShrinkRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetCustomerOrderListShrinkRequest) GetOrderPayAfter() *int64 {
	return s.OrderPayAfter
}

func (s *GetCustomerOrderListShrinkRequest) GetOrderPayBefore() *int64 {
	return s.OrderPayBefore
}

func (s *GetCustomerOrderListShrinkRequest) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *GetCustomerOrderListShrinkRequest) GetOrderTypeListShrink() *string {
	return s.OrderTypeListShrink
}

func (s *GetCustomerOrderListShrinkRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetCustomerOrderListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetCustomerOrderListShrinkRequest) GetPayAmountAfter() *float64 {
	return s.PayAmountAfter
}

func (s *GetCustomerOrderListShrinkRequest) GetPayAmountBefore() *float64 {
	return s.PayAmountBefore
}

func (s *GetCustomerOrderListShrinkRequest) GetPayType() *int32 {
	return s.PayType
}

func (s *GetCustomerOrderListShrinkRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetCustomerOrderListShrinkRequest) GetProductName() *string {
	return s.ProductName
}

func (s *GetCustomerOrderListShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetCustomerOrderListShrinkRequest) GetRamAccountForCustomerManager() *string {
	return s.RamAccountForCustomerManager
}

func (s *GetCustomerOrderListShrinkRequest) SetCustomerAccount(v string) *GetCustomerOrderListShrinkRequest {
	s.CustomerAccount = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetCustomerUid(v int64) *GetCustomerOrderListShrinkRequest {
	s.CustomerUid = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetOrderCreateAfter(v int64) *GetCustomerOrderListShrinkRequest {
	s.OrderCreateAfter = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetOrderCreateBefore(v int64) *GetCustomerOrderListShrinkRequest {
	s.OrderCreateBefore = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetOrderId(v int64) *GetCustomerOrderListShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetOrderPayAfter(v int64) *GetCustomerOrderListShrinkRequest {
	s.OrderPayAfter = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetOrderPayBefore(v int64) *GetCustomerOrderListShrinkRequest {
	s.OrderPayBefore = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetOrderStatus(v int32) *GetCustomerOrderListShrinkRequest {
	s.OrderStatus = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetOrderTypeListShrink(v string) *GetCustomerOrderListShrinkRequest {
	s.OrderTypeListShrink = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetPageNo(v int32) *GetCustomerOrderListShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetPageSize(v int32) *GetCustomerOrderListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetPayAmountAfter(v float64) *GetCustomerOrderListShrinkRequest {
	s.PayAmountAfter = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetPayAmountBefore(v float64) *GetCustomerOrderListShrinkRequest {
	s.PayAmountBefore = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetPayType(v int32) *GetCustomerOrderListShrinkRequest {
	s.PayType = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetProductCode(v string) *GetCustomerOrderListShrinkRequest {
	s.ProductCode = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetProductName(v string) *GetCustomerOrderListShrinkRequest {
	s.ProductName = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetProjectId(v int64) *GetCustomerOrderListShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) SetRamAccountForCustomerManager(v string) *GetCustomerOrderListShrinkRequest {
	s.RamAccountForCustomerManager = &v
	return s
}

func (s *GetCustomerOrderListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
