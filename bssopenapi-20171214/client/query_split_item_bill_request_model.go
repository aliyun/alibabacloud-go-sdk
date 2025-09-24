// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySplitItemBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *QuerySplitItemBillRequest
	GetBillOwnerId() *int64
	SetBillingCycle(v string) *QuerySplitItemBillRequest
	GetBillingCycle() *string
	SetOwnerId(v int64) *QuerySplitItemBillRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *QuerySplitItemBillRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QuerySplitItemBillRequest
	GetPageSize() *int32
	SetProductCode(v string) *QuerySplitItemBillRequest
	GetProductCode() *string
	SetProductType(v string) *QuerySplitItemBillRequest
	GetProductType() *string
	SetSubscriptionType(v string) *QuerySplitItemBillRequest
	GetSubscriptionType() *string
}

type QuerySplitItemBillRequest struct {
	// The ID of the member. If you specify a value for this parameter, you can query the split bills of the specified member. If you leave this parameter empty, the split bills of the current account are queried by default.
	//
	// example:
	//
	// 123
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The billing cycle, in the YYYY-MM format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-03
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The type of the service.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The billing method. Valid values: Subscription: subscription PayAsYouGo: pay-as-you-go This parameter must be used together with the ProductCode parameter.
	//
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s QuerySplitItemBillRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySplitItemBillRequest) GoString() string {
	return s.String()
}

func (s *QuerySplitItemBillRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *QuerySplitItemBillRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QuerySplitItemBillRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySplitItemBillRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QuerySplitItemBillRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySplitItemBillRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *QuerySplitItemBillRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QuerySplitItemBillRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QuerySplitItemBillRequest) SetBillOwnerId(v int64) *QuerySplitItemBillRequest {
	s.BillOwnerId = &v
	return s
}

func (s *QuerySplitItemBillRequest) SetBillingCycle(v string) *QuerySplitItemBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *QuerySplitItemBillRequest) SetOwnerId(v int64) *QuerySplitItemBillRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySplitItemBillRequest) SetPageNum(v int32) *QuerySplitItemBillRequest {
	s.PageNum = &v
	return s
}

func (s *QuerySplitItemBillRequest) SetPageSize(v int32) *QuerySplitItemBillRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySplitItemBillRequest) SetProductCode(v string) *QuerySplitItemBillRequest {
	s.ProductCode = &v
	return s
}

func (s *QuerySplitItemBillRequest) SetProductType(v string) *QuerySplitItemBillRequest {
	s.ProductType = &v
	return s
}

func (s *QuerySplitItemBillRequest) SetSubscriptionType(v string) *QuerySplitItemBillRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QuerySplitItemBillRequest) Validate() error {
	return dara.Validate(s)
}
