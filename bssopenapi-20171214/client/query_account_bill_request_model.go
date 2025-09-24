// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *QueryAccountBillRequest
	GetBillOwnerId() *int64
	SetBillingCycle(v string) *QueryAccountBillRequest
	GetBillingCycle() *string
	SetBillingDate(v string) *QueryAccountBillRequest
	GetBillingDate() *string
	SetGranularity(v string) *QueryAccountBillRequest
	GetGranularity() *string
	SetIsGroupByProduct(v bool) *QueryAccountBillRequest
	GetIsGroupByProduct() *bool
	SetOwnerID(v int64) *QueryAccountBillRequest
	GetOwnerID() *int64
	SetPageNum(v int32) *QueryAccountBillRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryAccountBillRequest
	GetPageSize() *int32
	SetProductCode(v string) *QueryAccountBillRequest
	GetProductCode() *string
}

type QueryAccountBillRequest struct {
	// The ID of the member. If you specify a value for this parameter, you can query the bills of the specified member. If you leave this parameter empty, the bills of the current account are queried by default.
	//
	// example:
	//
	// 122
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The billing cycle. Format: YYYY-MM.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-07
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only if the Granularity parameter is set to DAILY. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2021-06-01
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The granularity at which bills are queried. Valid values:
	//
	// 	- MONTHLY: queries bills by month. The data queried is consistent with the data that is displayed for the specified billing cycle on the Billing Details tab of the Bill Details page in User Center.
	//
	// 	- DAILY: queries bills by day. The data queried is consistent with the data that is displayed for the specified day on the Billing Details tab of the Bill Details page in User Center.
	//
	// You must set the BillingDate parameter before you can set the Granularity parameter to DAILY.
	//
	// example:
	//
	// Monthly
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// Specifies whether to summarize bills based on service codes. Valid values:
	//
	// 	- true: summarizes bills based on service codes.
	//
	// 	- false: does not summarize bills based on service codes.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IsGroupByProduct *bool  `json:"IsGroupByProduct,omitempty" xml:"IsGroupByProduct,omitempty"`
	OwnerID          *int64 `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries to return on each page. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The code of the service.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
}

func (s QueryAccountBillRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountBillRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountBillRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *QueryAccountBillRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QueryAccountBillRequest) GetBillingDate() *string {
	return s.BillingDate
}

func (s *QueryAccountBillRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *QueryAccountBillRequest) GetIsGroupByProduct() *bool {
	return s.IsGroupByProduct
}

func (s *QueryAccountBillRequest) GetOwnerID() *int64 {
	return s.OwnerID
}

func (s *QueryAccountBillRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryAccountBillRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryAccountBillRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryAccountBillRequest) SetBillOwnerId(v int64) *QueryAccountBillRequest {
	s.BillOwnerId = &v
	return s
}

func (s *QueryAccountBillRequest) SetBillingCycle(v string) *QueryAccountBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *QueryAccountBillRequest) SetBillingDate(v string) *QueryAccountBillRequest {
	s.BillingDate = &v
	return s
}

func (s *QueryAccountBillRequest) SetGranularity(v string) *QueryAccountBillRequest {
	s.Granularity = &v
	return s
}

func (s *QueryAccountBillRequest) SetIsGroupByProduct(v bool) *QueryAccountBillRequest {
	s.IsGroupByProduct = &v
	return s
}

func (s *QueryAccountBillRequest) SetOwnerID(v int64) *QueryAccountBillRequest {
	s.OwnerID = &v
	return s
}

func (s *QueryAccountBillRequest) SetPageNum(v int32) *QueryAccountBillRequest {
	s.PageNum = &v
	return s
}

func (s *QueryAccountBillRequest) SetPageSize(v int32) *QueryAccountBillRequest {
	s.PageSize = &v
	return s
}

func (s *QueryAccountBillRequest) SetProductCode(v string) *QueryAccountBillRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryAccountBillRequest) Validate() error {
	return dara.Validate(s)
}
