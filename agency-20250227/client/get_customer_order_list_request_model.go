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
	// Customer Account
	//
	// example:
	//
	// test_123
	CustomerAccount *string `json:"CustomerAccount,omitempty" xml:"CustomerAccount,omitempty"`
	// Customer UID
	//
	// example:
	//
	// 123456
	CustomerUid *int64 `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// The UNIX timestamp indicating the start time of order creation. The time range must not exceed six months.
	//
	// The time range for order creation and the time range for order payment cannot both be empty.
	//
	// example:
	//
	// 1727789348000
	OrderCreateAfter *int64 `json:"OrderCreateAfter,omitempty" xml:"OrderCreateAfter,omitempty"`
	// The UNIX timestamp indicating the end time of order creation. The time range must not exceed six months.
	//
	// The time range for order creation and the time range for order payment cannot both be empty.
	//
	// example:
	//
	// 1741008566000
	OrderCreateBefore *int64 `json:"OrderCreateBefore,omitempty" xml:"OrderCreateBefore,omitempty"`
	// Order ID
	//
	// example:
	//
	// 13595216
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// Order payment start UNIX timestamp. The time range must not exceed six months.
	//
	// The order creation time range and the order payment time range cannot both be empty.
	//
	// example:
	//
	// 1727789348000
	OrderPayAfter *int64 `json:"OrderPayAfter,omitempty" xml:"OrderPayAfter,omitempty"`
	// Order payment end UNIX timestamp. The time range must not exceed six months.
	//
	// The order creation time range and the order payment time range cannot both be empty.
	//
	// example:
	//
	// 1741008566000
	OrderPayBefore *int64 `json:"OrderPayBefore,omitempty" xml:"OrderPayBefore,omitempty"`
	// Order status:
	//
	// - 1 Unpaid
	//
	// - 2 Discarded
	//
	// - 3 Paid
	//
	// example:
	//
	// 3
	OrderStatus *int32 `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// Order type List
	OrderTypeList []*string `json:"OrderTypeList,omitempty" xml:"OrderTypeList,omitempty" type:"Repeated"`
	// Page number
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Page size
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Minimum paid amount
	//
	// example:
	//
	// 1
	PayAmountAfter *float64 `json:"PayAmountAfter,omitempty" xml:"PayAmountAfter,omitempty"`
	// Actual payment amount up to this point
	//
	// example:
	//
	// 1000
	PayAmountBefore *float64 `json:"PayAmountBefore,omitempty" xml:"PayAmountBefore,omitempty"`
	// Payment Type:
	//
	// 1: Non-agent payment;
	//
	// 2: Agent payment
	//
	// example:
	//
	// 1
	PayType *int32 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// Product code
	//
	// example:
	//
	// ecs
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// Product Name
	//
	// example:
	//
	// 弹性计算
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// Opportunity ID
	//
	// example:
	//
	// 202502002231
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Customer follow-up staff
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
