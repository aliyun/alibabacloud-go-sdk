// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *QueryBillRequest
	GetBillOwnerId() *int64
	SetBillingCycle(v string) *QueryBillRequest
	GetBillingCycle() *string
	SetIsDisplayLocalCurrency(v bool) *QueryBillRequest
	GetIsDisplayLocalCurrency() *bool
	SetIsHideZeroCharge(v bool) *QueryBillRequest
	GetIsHideZeroCharge() *bool
	SetOwnerId(v int64) *QueryBillRequest
	GetOwnerId() *int64
	SetPageNum(v int32) *QueryBillRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryBillRequest
	GetPageSize() *int32
	SetProductCode(v string) *QueryBillRequest
	GetProductCode() *string
	SetProductType(v string) *QueryBillRequest
	GetProductType() *string
	SetSubscriptionType(v string) *QueryBillRequest
	GetSubscriptionType() *string
	SetType(v string) *QueryBillRequest
	GetType() *string
}

type QueryBillRequest struct {
	// The ID of the member.
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
	// 2018-07
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// Specifies whether to display local currency information in bills. The parameter will be discontinued.
	//
	// example:
	//
	// false
	IsDisplayLocalCurrency *bool `json:"IsDisplayLocalCurrency,omitempty" xml:"IsDisplayLocalCurrency,omitempty"`
	// Specifies whether to filter out a bill whose pretax gross amount is 0. By default, a bill whose pretax gross amount is 0 is not filtered out. Valid values:
	//
	// 	- true: filters out a bill whose pretax gross amount is 0.
	//
	// 	- false: does not filter out a bill whose pretax gross amount is 0.
	//
	// example:
	//
	// true
	IsHideZeroCharge *bool  `json:"IsHideZeroCharge,omitempty" xml:"IsHideZeroCharge,omitempty"`
	OwnerId          *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The billing method. Valid values:
	//
	// 	- Subscription
	//
	// 	- PayAsYouGo
	//
	// This parameter must be used together with the ProductCode parameter.
	//
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
	// The type of the bill. Valid values:
	//
	// 	- SubscriptionOrder
	//
	// 	- PayAsYouGoBill
	//
	// 	- Refund
	//
	// 	- Adjustment
	//
	// example:
	//
	// SubscriptionOrder
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryBillRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryBillRequest) GoString() string {
	return s.String()
}

func (s *QueryBillRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *QueryBillRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *QueryBillRequest) GetIsDisplayLocalCurrency() *bool {
	return s.IsDisplayLocalCurrency
}

func (s *QueryBillRequest) GetIsHideZeroCharge() *bool {
	return s.IsHideZeroCharge
}

func (s *QueryBillRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryBillRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryBillRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryBillRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *QueryBillRequest) GetProductType() *string {
	return s.ProductType
}

func (s *QueryBillRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *QueryBillRequest) GetType() *string {
	return s.Type
}

func (s *QueryBillRequest) SetBillOwnerId(v int64) *QueryBillRequest {
	s.BillOwnerId = &v
	return s
}

func (s *QueryBillRequest) SetBillingCycle(v string) *QueryBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *QueryBillRequest) SetIsDisplayLocalCurrency(v bool) *QueryBillRequest {
	s.IsDisplayLocalCurrency = &v
	return s
}

func (s *QueryBillRequest) SetIsHideZeroCharge(v bool) *QueryBillRequest {
	s.IsHideZeroCharge = &v
	return s
}

func (s *QueryBillRequest) SetOwnerId(v int64) *QueryBillRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryBillRequest) SetPageNum(v int32) *QueryBillRequest {
	s.PageNum = &v
	return s
}

func (s *QueryBillRequest) SetPageSize(v int32) *QueryBillRequest {
	s.PageSize = &v
	return s
}

func (s *QueryBillRequest) SetProductCode(v string) *QueryBillRequest {
	s.ProductCode = &v
	return s
}

func (s *QueryBillRequest) SetProductType(v string) *QueryBillRequest {
	s.ProductType = &v
	return s
}

func (s *QueryBillRequest) SetSubscriptionType(v string) *QueryBillRequest {
	s.SubscriptionType = &v
	return s
}

func (s *QueryBillRequest) SetType(v string) *QueryBillRequest {
	s.Type = &v
	return s
}

func (s *QueryBillRequest) Validate() error {
	return dara.Validate(s)
}
