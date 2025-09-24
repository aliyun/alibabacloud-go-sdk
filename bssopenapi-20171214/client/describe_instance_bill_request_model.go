// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceBillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerId(v int64) *DescribeInstanceBillRequest
	GetBillOwnerId() *int64
	SetBillingCycle(v string) *DescribeInstanceBillRequest
	GetBillingCycle() *string
	SetBillingDate(v string) *DescribeInstanceBillRequest
	GetBillingDate() *string
	SetGranularity(v string) *DescribeInstanceBillRequest
	GetGranularity() *string
	SetInstanceID(v string) *DescribeInstanceBillRequest
	GetInstanceID() *string
	SetIsBillingItem(v bool) *DescribeInstanceBillRequest
	GetIsBillingItem() *bool
	SetIsHideZeroCharge(v bool) *DescribeInstanceBillRequest
	GetIsHideZeroCharge() *bool
	SetMaxResults(v int32) *DescribeInstanceBillRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInstanceBillRequest
	GetNextToken() *string
	SetOwnerId(v int64) *DescribeInstanceBillRequest
	GetOwnerId() *int64
	SetPipCode(v string) *DescribeInstanceBillRequest
	GetPipCode() *string
	SetProductCode(v string) *DescribeInstanceBillRequest
	GetProductCode() *string
	SetProductType(v string) *DescribeInstanceBillRequest
	GetProductType() *string
	SetSubscriptionType(v string) *DescribeInstanceBillRequest
	GetSubscriptionType() *string
}

type DescribeInstanceBillRequest struct {
	// The ID of the member. If you specify this parameter, the bills of the member are queried. If you do not specify this parameter, the bills of the current account are queried by default.
	//
	// example:
	//
	// 122
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The billing cycle. Specify the parameter in the YYYY-MM format.
	//
	// Only the latest 18 month billing cycle is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-03
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only when the Granularity parameter is set to DAILY. Format: YYYY-MM-DD.
	//
	// example:
	//
	// 2020-03-02
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The granularity at which bills are queried. Valid values:
	//
	// 	- MONTHLY: queries bills on a monthly basis. The data that you query is the same as the data searched by instances on the Billing Details tab of the Bill Details page in the User Center console.
	//
	// 	- DAILY: queries bills on a daily basis. The data that you query is the same as the data searched by days on the Billing Details tab of the Bill Details page in the User Center console.
	//
	// The BillingDate parameter is required if you set the Granularity parameter to DAILY.
	//
	// example:
	//
	// MONTHLY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// abc
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// Specifies whether to query data by billable items. Valid values:
	//
	// 	- false: The data that you query is the same as the data searched by instances on the Billing Details tab of the Bill Details page in the User Center console.
	//
	// 	- true: The data that you query is the same as the data searched by billable items on the Billing Details tab of the Bill Details page in the User Center console.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	IsBillingItem *bool `json:"IsBillingItem,omitempty" xml:"IsBillingItem,omitempty"`
	// Specifies whether to filter bills if both the pretax gross amount and pretax amount are 0. Valid values:
	//
	// 	- false: does not filter bills.
	//
	// 	- true: filters bills.
	//
	// example:
	//
	// false
	IsHideZeroCharge *bool `json:"IsHideZeroCharge,omitempty" xml:"IsHideZeroCharge,omitempty"`
	// The maximum number of entries to return. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to indicate the position where the results for the current call start. The parameter must be left empty or set to the value of the NextToken parameter that is returned from the last call. Otherwise, an error is returned. If the parameter is left empty, data is queried from the first item.
	//
	// example:
	//
	// CAESEgoQCg4KCm
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The code of the service. The code is the same as that in Cost Center.
	//
	// example:
	//
	// rds
	PipCode *string `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
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
	// 	- Subscription: the subscription billing method.
	//
	// 	- PayAsYouGo: the pay-as-you-go billing method.
	//
	// example:
	//
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s DescribeInstanceBillRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceBillRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceBillRequest) GetBillOwnerId() *int64 {
	return s.BillOwnerId
}

func (s *DescribeInstanceBillRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeInstanceBillRequest) GetBillingDate() *string {
	return s.BillingDate
}

func (s *DescribeInstanceBillRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *DescribeInstanceBillRequest) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeInstanceBillRequest) GetIsBillingItem() *bool {
	return s.IsBillingItem
}

func (s *DescribeInstanceBillRequest) GetIsHideZeroCharge() *bool {
	return s.IsHideZeroCharge
}

func (s *DescribeInstanceBillRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceBillRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceBillRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceBillRequest) GetPipCode() *string {
	return s.PipCode
}

func (s *DescribeInstanceBillRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstanceBillRequest) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeInstanceBillRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeInstanceBillRequest) SetBillOwnerId(v int64) *DescribeInstanceBillRequest {
	s.BillOwnerId = &v
	return s
}

func (s *DescribeInstanceBillRequest) SetBillingCycle(v string) *DescribeInstanceBillRequest {
	s.BillingCycle = &v
	return s
}

func (s *DescribeInstanceBillRequest) SetBillingDate(v string) *DescribeInstanceBillRequest {
	s.BillingDate = &v
	return s
}

func (s *DescribeInstanceBillRequest) SetGranularity(v string) *DescribeInstanceBillRequest {
	s.Granularity = &v
	return s
}

func (s *DescribeInstanceBillRequest) SetInstanceID(v string) *DescribeInstanceBillRequest {
	s.InstanceID = &v
	return s
}

func (s *DescribeInstanceBillRequest) SetIsBillingItem(v bool) *DescribeInstanceBillRequest {
	s.IsBillingItem = &v
	return s
}

func (s *DescribeInstanceBillRequest) SetIsHideZeroCharge(v bool) *DescribeInstanceBillRequest {
	s.IsHideZeroCharge = &v
	return s
}

func (s *DescribeInstanceBillRequest) SetMaxResults(v int32) *DescribeInstanceBillRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceBillRequest) SetNextToken(v string) *DescribeInstanceBillRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceBillRequest) SetOwnerId(v int64) *DescribeInstanceBillRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceBillRequest) SetPipCode(v string) *DescribeInstanceBillRequest {
	s.PipCode = &v
	return s
}

func (s *DescribeInstanceBillRequest) SetProductCode(v string) *DescribeInstanceBillRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceBillRequest) SetProductType(v string) *DescribeInstanceBillRequest {
	s.ProductType = &v
	return s
}

func (s *DescribeInstanceBillRequest) SetSubscriptionType(v string) *DescribeInstanceBillRequest {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeInstanceBillRequest) Validate() error {
	return dara.Validate(s)
}
