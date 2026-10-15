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
	// The ID of the resource ownership account. The resource ownership account is the account that actually uses the resource.
	//
	// example:
	//
	// 122
	BillOwnerId *int64 `json:"BillOwnerId,omitempty" xml:"BillOwnerId,omitempty"`
	// The billing cycle in the YYYY-MM format. Only the last 18 billing cycles are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-03
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing date. This parameter is required only when Granularity is set to DAILY. Format: YYYY-MM-DD.
	//
	// > The value of BillingDate must be in the same month as the value of BillingCycle. For example, if BillingCycle is set to 2020-03, BillingDate must be set to 2020-03-DD. If the months are inconsistent, the query results are returned based on the time specified by BillingCycle.
	//
	// example:
	//
	// 2020-03-02
	BillingDate *string `json:"BillingDate,omitempty" xml:"BillingDate,omitempty"`
	// The granularity at which bills are queried. Valid values:
	//
	// - MONTHLY: monthly. The data is consistent with the monthly bills in Expenses and Costs > Bills > Bill Details > Monthly Bills.
	//
	// - DAILY: daily. The data is consistent with the daily bills in Expenses and Costs > Bills > Bill Details > Daily Bills.
	//
	// > If you select DAILY, you must specify BillingDate.
	//
	// example:
	//
	// MONTHLY
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// abc
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// Specifies whether to retrieve data by billable item. Valid values:
	//
	// - false (default): The data is consistent with the instance bills in Expenses and Costs > Bills > Bill Details > Instance Bills.
	//
	// - true: The data is consistent with the billable item bills in Expenses and Costs > Bills > Bill Details > Billable Item Bills.
	//
	// example:
	//
	// false
	IsBillingItem *bool `json:"IsBillingItem,omitempty" xml:"IsBillingItem,omitempty"`
	// Specifies whether to filter out records where both the original price (PretaxGrossAmount) and the payable amount (PretaxAmount) are 0. Valid values:
	//
	// - false
	//
	// - true
	//
	// example:
	//
	// false
	IsHideZeroCharge *bool `json:"IsHideZeroCharge,omitempty" xml:"IsHideZeroCharge,omitempty"`
	// The maximum number of data records to read in the current call. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The position from which the current call starts to read data. The value must be empty or set to the NextToken value returned in the previous response. Otherwise, an error occurs. An empty value indicates that the reading starts from the beginning.
	//
	// example:
	//
	// CAESEgoQCg4KCm
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The product code, which is consistent with the product code in the bills of the User Center.
	//
	// example:
	//
	// rds
	PipCode *string `json:"PipCode,omitempty" xml:"PipCode,omitempty"`
	// The product code.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product type.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The subscription type. Valid values:
	//
	// - Subscription: subscription.
	//
	// - PayAsYouGo: pay-as-you-go.
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
