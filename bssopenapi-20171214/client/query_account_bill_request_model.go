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
	// The ID of the account that owns and actually uses the resource.
	//
	// example:
	//
	// 122
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The billing date. Format: YYYY-MM.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-07
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only when you set Granularity to DAILY. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2021-06-01
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The granularity for querying the bill. Valid values:
	//
	// - MONTHLY: Queries bills on a monthly basis. Corresponds to the monthly bills in **Expenses & Costs&#xA;**> **Bills&#xA;**> **Bill Details&#xA;**.
	//
	// - DAILY: Queries bills on a daily basis. Corresponds to the daily bills in **Expenses & Costs&#xA;**> **Bills&#xA;**> **Bill Details&#xA;**.
	//
	// > If you set `Granularity` to `DAILY`, you must also specify the `BillingDate` parameter.
	//
	// example:
	//
	// MONTHLY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// Specifies whether to aggregate the results by product code.
	//
	// - true: Results are aggregated by product code.
	//
	// - false: Results are not aggregated by product code.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IsGroupByProduct *bool  `json:"IsGroupByProduct,omitempty" xml:"IsGroupByProduct,omitempty"`
	OwnerID          *int64 `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	// The page number. The default value is 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. The default value is 20. The maximum value is 300.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product code.
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
