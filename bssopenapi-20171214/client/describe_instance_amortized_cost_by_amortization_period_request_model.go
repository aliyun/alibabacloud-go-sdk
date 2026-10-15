// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAmortizedCostByAmortizationPeriodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerIdList(v []*string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest
	GetBillOwnerIdList() []*string
	SetBillUserIdList(v []*string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest
	GetBillUserIdList() []*string
	SetBillingCycle(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest
	GetBillingCycle() *string
	SetConsumePeriodFilter(v []*string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest
	GetConsumePeriodFilter() []*string
	SetCostUnitCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest
	GetCostUnitCode() *string
	SetInstanceIdList(v []*string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest
	GetInstanceIdList() []*string
	SetMaxResults(v int32) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest
	GetNextToken() *string
	SetProductCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest
	GetProductCode() *string
	SetProductDetail(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest
	GetProductDetail() *string
	SetSubscriptionType(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest
	GetSubscriptionType() *string
}

type DescribeInstanceAmortizedCostByAmortizationPeriodRequest struct {
	// The filter condition for the account to which the resource belongs.
	//
	// The resource ownership account is the account that actually uses the resource.
	//
	// Maximum number of elements: 10.
	BillOwnerIdList []*string `json:"BillOwnerIdList,omitempty" xml:"BillOwnerIdList,omitempty" type:"Repeated"`
	// The filter condition for the ID of the bill owner.
	//
	// This parameter specifies the ID of a RAM user. You can set this parameter to query the bills of a RAM user. If you do not set this parameter, the bills of the current account and all RAM users of the current account are queried.
	//
	// Maximum number of elements: 10.
	BillUserIdList []*string `json:"BillUserIdList,omitempty" xml:"BillUserIdList,omitempty" type:"Repeated"`
	// The amortization month.
	//
	// Format: YYYY-MM.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-10
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The filter condition for the billing cycle.
	//
	// Maximum number of elements: 10.
	ConsumePeriodFilter []*string `json:"ConsumePeriodFilter,omitempty" xml:"ConsumePeriodFilter,omitempty" type:"Repeated"`
	// The code of the cost center.
	//
	// example:
	//
	// 123#
	CostUnitCode *string `json:"CostUnitCode,omitempty" xml:"CostUnitCode,omitempty"`
	// The filter condition for the instance ID.
	//
	// You can query multiple instances. If you do not set this parameter, all instances are queried.
	//
	// Maximum number of elements: 10.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	// The maximum number of data records to return. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The position from which to start reading the current call. The value of this parameter must be empty or the value of NextToken that is returned in the last call. Otherwise, an error is reported. If this parameter is empty, the query starts from the beginning.
	//
	// example:
	//
	// CAESEgoQCg4KCmdtdF9jcmVhdGUEARgBIkgKCQBwhGmPcAEAAAo7AzYAAAAxTDgwMDcxMjg3ZDJhNmM3ZDguTDgwMDAwMDAwMDAwMzE1MTIuTDgwMDcyZDMyZTJkYzg3N2U
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The product code. To obtain the value of this parameter, call QueryProductList or DescribeResourcePackageProduct.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The details of the product.
	//
	// example:
	//
	// rds
	ProductDetail *string `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	// The subscription type. Valid values:
	//
	// - Subscription: subscription.
	//
	// - PayAsYouGo: pay-as-you-go.
	//
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s DescribeInstanceAmortizedCostByAmortizationPeriodRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAmortizedCostByAmortizationPeriodRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) GetBillOwnerIdList() []*string {
	return s.BillOwnerIdList
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) GetBillUserIdList() []*string {
	return s.BillUserIdList
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) GetConsumePeriodFilter() []*string {
	return s.ConsumePeriodFilter
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) GetCostUnitCode() *string {
	return s.CostUnitCode
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) SetBillOwnerIdList(v []*string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest {
	s.BillOwnerIdList = v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) SetBillUserIdList(v []*string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest {
	s.BillUserIdList = v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) SetBillingCycle(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest {
	s.BillingCycle = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) SetConsumePeriodFilter(v []*string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest {
	s.ConsumePeriodFilter = v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) SetCostUnitCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest {
	s.CostUnitCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) SetInstanceIdList(v []*string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest {
	s.InstanceIdList = v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) SetMaxResults(v int32) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) SetNextToken(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) SetProductCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) SetProductDetail(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest {
	s.ProductDetail = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) SetSubscriptionType(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodRequest {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodRequest) Validate() error {
	return dara.Validate(s)
}
