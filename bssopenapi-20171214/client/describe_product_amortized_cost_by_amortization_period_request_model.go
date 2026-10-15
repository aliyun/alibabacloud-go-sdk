// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductAmortizedCostByAmortizationPeriodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerIdList(v []*string) *DescribeProductAmortizedCostByAmortizationPeriodRequest
	GetBillOwnerIdList() []*string
	SetBillUserIdList(v []*string) *DescribeProductAmortizedCostByAmortizationPeriodRequest
	GetBillUserIdList() []*string
	SetBillingCycle(v string) *DescribeProductAmortizedCostByAmortizationPeriodRequest
	GetBillingCycle() *string
	SetConsumePeriodFilter(v []*string) *DescribeProductAmortizedCostByAmortizationPeriodRequest
	GetConsumePeriodFilter() []*string
	SetCostUnitCode(v string) *DescribeProductAmortizedCostByAmortizationPeriodRequest
	GetCostUnitCode() *string
	SetMaxResults(v int32) *DescribeProductAmortizedCostByAmortizationPeriodRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeProductAmortizedCostByAmortizationPeriodRequest
	GetNextToken() *string
	SetProductCode(v string) *DescribeProductAmortizedCostByAmortizationPeriodRequest
	GetProductCode() *string
	SetProductDetail(v string) *DescribeProductAmortizedCostByAmortizationPeriodRequest
	GetProductDetail() *string
	SetSubscriptionType(v string) *DescribeProductAmortizedCostByAmortizationPeriodRequest
	GetSubscriptionType() *string
}

type DescribeProductAmortizedCostByAmortizationPeriodRequest struct {
	// The IDs of resource owners. The resource owner is the account that uses the resources.
	//
	// You can specify a maximum of 10 IDs.
	BillOwnerIdList []*string `json:"BillOwnerIdList,omitempty" xml:"BillOwnerIdList,omitempty" type:"Repeated"`
	// The IDs of billable accounts. This parameter is used to filter bills by the RAM user who pays for the resources. If you specify this parameter, you can query the bills of the specified RAM user. If you leave this parameter empty, the bills of the current account and all its RAM users are queried.
	//
	// You can specify a maximum of 10 IDs.
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
	// The filter for billing cycle months.
	//
	// You can specify a maximum of 10 elements.
	ConsumePeriodFilter []*string `json:"ConsumePeriodFilter,omitempty" xml:"ConsumePeriodFilter,omitempty" type:"Repeated"`
	// The code of the cost center.
	//
	// example:
	//
	// 123#
	CostUnitCode *string `json:"CostUnitCode,omitempty" xml:"CostUnitCode,omitempty"`
	// The maximum number of entries to return. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position to start the query. If you leave this parameter empty, the query starts from the beginning. Set this parameter to the value of NextToken that is returned in the last response to query the next page of results.
	//
	// example:
	//
	// CAESEgoQCg4KCmdtdF9jcmVhdGUEARgBIkgKCQBwhGmPcAEAAAo7AzYAAAAxTDgwMDcxMjg3ZDJhNmM3ZDguTDgwMDAwMDAwMDAwMzE1MTIuTDgwMDcyZDMyZTJkYzg3N2U
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The product code.
	//
	// For more information about how to obtain this parameter, see QueryProductList and DescribeResourcePackageProduct.
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

func (s DescribeProductAmortizedCostByAmortizationPeriodRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductAmortizedCostByAmortizationPeriodRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) GetBillOwnerIdList() []*string {
	return s.BillOwnerIdList
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) GetBillUserIdList() []*string {
	return s.BillUserIdList
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) GetConsumePeriodFilter() []*string {
	return s.ConsumePeriodFilter
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) GetCostUnitCode() *string {
	return s.CostUnitCode
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) SetBillOwnerIdList(v []*string) *DescribeProductAmortizedCostByAmortizationPeriodRequest {
	s.BillOwnerIdList = v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) SetBillUserIdList(v []*string) *DescribeProductAmortizedCostByAmortizationPeriodRequest {
	s.BillUserIdList = v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) SetBillingCycle(v string) *DescribeProductAmortizedCostByAmortizationPeriodRequest {
	s.BillingCycle = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) SetConsumePeriodFilter(v []*string) *DescribeProductAmortizedCostByAmortizationPeriodRequest {
	s.ConsumePeriodFilter = v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) SetCostUnitCode(v string) *DescribeProductAmortizedCostByAmortizationPeriodRequest {
	s.CostUnitCode = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) SetMaxResults(v int32) *DescribeProductAmortizedCostByAmortizationPeriodRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) SetNextToken(v string) *DescribeProductAmortizedCostByAmortizationPeriodRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) SetProductCode(v string) *DescribeProductAmortizedCostByAmortizationPeriodRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) SetProductDetail(v string) *DescribeProductAmortizedCostByAmortizationPeriodRequest {
	s.ProductDetail = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) SetSubscriptionType(v string) *DescribeProductAmortizedCostByAmortizationPeriodRequest {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeProductAmortizedCostByAmortizationPeriodRequest) Validate() error {
	return dara.Validate(s)
}
