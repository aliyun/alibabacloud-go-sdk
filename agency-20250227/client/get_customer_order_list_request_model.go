// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerOrderListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerAccount(v string) *GetCustomerOrderListRequest
	GetCustomerAccount() *string
	SetCustomerUid(v int64) *GetCustomerOrderListRequest
	GetCustomerUid() *int64
	SetOrderCreateAfter(v int64) *GetCustomerOrderListRequest
	GetOrderCreateAfter() *int64
	SetOrderCreateBefore(v int64) *GetCustomerOrderListRequest
	GetOrderCreateBefore() *int64
	SetOrderId(v int64) *GetCustomerOrderListRequest
	GetOrderId() *int64
	SetOrderPayAfter(v int64) *GetCustomerOrderListRequest
	GetOrderPayAfter() *int64
	SetOrderPayBefore(v int64) *GetCustomerOrderListRequest
	GetOrderPayBefore() *int64
	SetOrderStatus(v int32) *GetCustomerOrderListRequest
	GetOrderStatus() *int32
	SetOrderTypeList(v []*string) *GetCustomerOrderListRequest
	GetOrderTypeList() []*string
	SetPageNo(v int32) *GetCustomerOrderListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetCustomerOrderListRequest
	GetPageSize() *int32
	SetPayAmountAfter(v float64) *GetCustomerOrderListRequest
	GetPayAmountAfter() *float64
	SetPayAmountBefore(v float64) *GetCustomerOrderListRequest
	GetPayAmountBefore() *float64
	SetPayType(v int32) *GetCustomerOrderListRequest
	GetPayType() *int32
	SetProductCode(v string) *GetCustomerOrderListRequest
	GetProductCode() *string
	SetProductName(v string) *GetCustomerOrderListRequest
	GetProductName() *string
	SetProjectId(v int64) *GetCustomerOrderListRequest
	GetProjectId() *int64
	SetRamAccountForCustomerManager(v string) *GetCustomerOrderListRequest
	GetRamAccountForCustomerManager() *string
}

type GetCustomerOrderListRequest struct {
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
	OrderTypeList []*string `json:"OrderTypeList,omitempty" xml:"OrderTypeList,omitempty" type:"Repeated"`
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

func (s GetCustomerOrderListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerOrderListRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerOrderListRequest) GetCustomerAccount() *string {
	return s.CustomerAccount
}

func (s *GetCustomerOrderListRequest) GetCustomerUid() *int64 {
	return s.CustomerUid
}

func (s *GetCustomerOrderListRequest) GetOrderCreateAfter() *int64 {
	return s.OrderCreateAfter
}

func (s *GetCustomerOrderListRequest) GetOrderCreateBefore() *int64 {
	return s.OrderCreateBefore
}

func (s *GetCustomerOrderListRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetCustomerOrderListRequest) GetOrderPayAfter() *int64 {
	return s.OrderPayAfter
}

func (s *GetCustomerOrderListRequest) GetOrderPayBefore() *int64 {
	return s.OrderPayBefore
}

func (s *GetCustomerOrderListRequest) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *GetCustomerOrderListRequest) GetOrderTypeList() []*string {
	return s.OrderTypeList
}

func (s *GetCustomerOrderListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetCustomerOrderListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetCustomerOrderListRequest) GetPayAmountAfter() *float64 {
	return s.PayAmountAfter
}

func (s *GetCustomerOrderListRequest) GetPayAmountBefore() *float64 {
	return s.PayAmountBefore
}

func (s *GetCustomerOrderListRequest) GetPayType() *int32 {
	return s.PayType
}

func (s *GetCustomerOrderListRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetCustomerOrderListRequest) GetProductName() *string {
	return s.ProductName
}

func (s *GetCustomerOrderListRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetCustomerOrderListRequest) GetRamAccountForCustomerManager() *string {
	return s.RamAccountForCustomerManager
}

func (s *GetCustomerOrderListRequest) SetCustomerAccount(v string) *GetCustomerOrderListRequest {
	s.CustomerAccount = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetCustomerUid(v int64) *GetCustomerOrderListRequest {
	s.CustomerUid = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetOrderCreateAfter(v int64) *GetCustomerOrderListRequest {
	s.OrderCreateAfter = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetOrderCreateBefore(v int64) *GetCustomerOrderListRequest {
	s.OrderCreateBefore = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetOrderId(v int64) *GetCustomerOrderListRequest {
	s.OrderId = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetOrderPayAfter(v int64) *GetCustomerOrderListRequest {
	s.OrderPayAfter = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetOrderPayBefore(v int64) *GetCustomerOrderListRequest {
	s.OrderPayBefore = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetOrderStatus(v int32) *GetCustomerOrderListRequest {
	s.OrderStatus = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetOrderTypeList(v []*string) *GetCustomerOrderListRequest {
	s.OrderTypeList = v
	return s
}

func (s *GetCustomerOrderListRequest) SetPageNo(v int32) *GetCustomerOrderListRequest {
	s.PageNo = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetPageSize(v int32) *GetCustomerOrderListRequest {
	s.PageSize = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetPayAmountAfter(v float64) *GetCustomerOrderListRequest {
	s.PayAmountAfter = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetPayAmountBefore(v float64) *GetCustomerOrderListRequest {
	s.PayAmountBefore = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetPayType(v int32) *GetCustomerOrderListRequest {
	s.PayType = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetProductCode(v string) *GetCustomerOrderListRequest {
	s.ProductCode = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetProductName(v string) *GetCustomerOrderListRequest {
	s.ProductName = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetProjectId(v int64) *GetCustomerOrderListRequest {
	s.ProjectId = &v
	return s
}

func (s *GetCustomerOrderListRequest) SetRamAccountForCustomerManager(v string) *GetCustomerOrderListRequest {
	s.RamAccountForCustomerManager = &v
	return s
}

func (s *GetCustomerOrderListRequest) Validate() error {
	return dara.Validate(s)
}
