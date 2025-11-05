// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerOrdersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerAccount(v string) *GetCustomerOrdersRequest
	GetCustomerAccount() *string
	SetCustomerManager(v string) *GetCustomerOrdersRequest
	GetCustomerManager() *string
	SetCustomerUid(v int64) *GetCustomerOrdersRequest
	GetCustomerUid() *int64
	SetEndDate(v string) *GetCustomerOrdersRequest
	GetEndDate() *string
	SetOrderId(v int64) *GetCustomerOrdersRequest
	GetOrderId() *int64
	SetOrderSource(v int32) *GetCustomerOrdersRequest
	GetOrderSource() *int32
	SetOrderStatus(v int32) *GetCustomerOrdersRequest
	GetOrderStatus() *int32
	SetOrderType(v string) *GetCustomerOrdersRequest
	GetOrderType() *string
	SetPageNo(v int32) *GetCustomerOrdersRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetCustomerOrdersRequest
	GetPageSize() *int32
	SetProductType(v string) *GetCustomerOrdersRequest
	GetProductType() *string
	SetStartDate(v string) *GetCustomerOrdersRequest
	GetStartDate() *string
	SetTimeType(v int32) *GetCustomerOrdersRequest
	GetTimeType() *int32
}

type GetCustomerOrdersRequest struct {
	// example:
	//
	// test_123
	CustomerAccount *string `json:"CustomerAccount,omitempty" xml:"CustomerAccount,omitempty"`
	// example:
	//
	// myBd
	CustomerManager *string `json:"CustomerManager,omitempty" xml:"CustomerManager,omitempty"`
	// example:
	//
	// 123456
	CustomerUid *int64 `json:"CustomerUid,omitempty" xml:"CustomerUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-08-23 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 209335720330622
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 0
	OrderSource *int32 `json:"OrderSource,omitempty" xml:"OrderSource,omitempty"`
	// example:
	//
	// 3
	OrderStatus *int32 `json:"OrderStatus,omitempty" xml:"OrderStatus,omitempty"`
	// example:
	//
	// RENEW
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// vm_intl
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024-08-13 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	TimeType *int32 `json:"TimeType,omitempty" xml:"TimeType,omitempty"`
}

func (s GetCustomerOrdersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerOrdersRequest) GoString() string {
	return s.String()
}

func (s *GetCustomerOrdersRequest) GetCustomerAccount() *string {
	return s.CustomerAccount
}

func (s *GetCustomerOrdersRequest) GetCustomerManager() *string {
	return s.CustomerManager
}

func (s *GetCustomerOrdersRequest) GetCustomerUid() *int64 {
	return s.CustomerUid
}

func (s *GetCustomerOrdersRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetCustomerOrdersRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetCustomerOrdersRequest) GetOrderSource() *int32 {
	return s.OrderSource
}

func (s *GetCustomerOrdersRequest) GetOrderStatus() *int32 {
	return s.OrderStatus
}

func (s *GetCustomerOrdersRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *GetCustomerOrdersRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetCustomerOrdersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetCustomerOrdersRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GetCustomerOrdersRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetCustomerOrdersRequest) GetTimeType() *int32 {
	return s.TimeType
}

func (s *GetCustomerOrdersRequest) SetCustomerAccount(v string) *GetCustomerOrdersRequest {
	s.CustomerAccount = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetCustomerManager(v string) *GetCustomerOrdersRequest {
	s.CustomerManager = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetCustomerUid(v int64) *GetCustomerOrdersRequest {
	s.CustomerUid = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetEndDate(v string) *GetCustomerOrdersRequest {
	s.EndDate = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetOrderId(v int64) *GetCustomerOrdersRequest {
	s.OrderId = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetOrderSource(v int32) *GetCustomerOrdersRequest {
	s.OrderSource = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetOrderStatus(v int32) *GetCustomerOrdersRequest {
	s.OrderStatus = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetOrderType(v string) *GetCustomerOrdersRequest {
	s.OrderType = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetPageNo(v int32) *GetCustomerOrdersRequest {
	s.PageNo = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetPageSize(v int32) *GetCustomerOrdersRequest {
	s.PageSize = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetProductType(v string) *GetCustomerOrdersRequest {
	s.ProductType = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetStartDate(v string) *GetCustomerOrdersRequest {
	s.StartDate = &v
	return s
}

func (s *GetCustomerOrdersRequest) SetTimeType(v int32) *GetCustomerOrdersRequest {
	s.TimeType = &v
	return s
}

func (s *GetCustomerOrdersRequest) Validate() error {
	return dara.Validate(s)
}
