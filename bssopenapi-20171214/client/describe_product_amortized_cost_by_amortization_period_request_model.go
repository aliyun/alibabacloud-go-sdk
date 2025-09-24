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
	// The instance ID that is used to filter bills. You can specify multiple instance IDs to query bills of multiple instances. If you leave this parameter empty, the bills of all instances are queried by default. You can specify a maximum of 10 instance IDs.
	BillOwnerIdList []*string `json:"BillOwnerIdList,omitempty" xml:"BillOwnerIdList,omitempty" type:"Repeated"`
	// The ID of the member that needs to settle the bill. The member ID is used to filter bills. If you specify a value for this parameter, you can query the bills of the specified member. If you leave this parameter empty, the bills of the current account and all members of the current account are queried by default. You can specify a maximum of 10 IDs.
	BillUserIdList []*string `json:"BillUserIdList,omitempty" xml:"BillUserIdList,omitempty" type:"Repeated"`
	// The allocation month. Format: YYYY-MM.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-10
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The billing cycle that is used to filter bills. You can specify a maximum of 10 billing cycles.
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
	// The position from which the query starts. The parameter must be left empty or set to the value of the NextToken parameter returned from the last call. Otherwise, an error is returned. If this parameter is left empty, data is queried from the beginning.
	//
	// example:
	//
	// CAESEgoQCg4KCmdtdF9jcmVhdGUEARgBIkgKCQBwhGmPcAEAAAo7AzYAAAAxTDgwMDcxMjg3ZDJhNmM3ZDguTDgwMDAwMDAwMDAwMzE1MTIuTDgwMDcyZDMyZTJkYzg3N2U
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The code of the service. You can obtain the value of this parameter by calling the QueryProductList operation or the DescribeResourcePackageProduct operation.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The specific service resource.
	//
	// example:
	//
	// rds
	ProductDetail *string `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	// The billing method. Valid values:
	//
	// 	- Subscription: the subscription billing method
	//
	// 	- PayAsYouGo: the pay-as-you-go billing method
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
