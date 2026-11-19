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
	// The filter condition for the resource ownership account. The ID of the resource ownership account. The resource ownership account is the account that actually uses the resources. Maximum number of elements: 10.
	BillOwnerIdList []*string `json:"BillOwnerIdList,omitempty" xml:"BillOwnerIdList,omitempty" type:"Repeated"`
	// The filter condition for the bill purchaser ID. The ID of the RAM user. Specify this parameter to view the bills of RAM users in Finance Cloud. If you do not specify this parameter, the bills of the current caller account and all its RAM users are viewed by default. Maximum number of elements: 10.
	BillUserIdList []*string `json:"BillUserIdList,omitempty" xml:"BillUserIdList,omitempty" type:"Repeated"`
	// The amortization month. Format: YYYY-MM.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-10
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// The filter condition for the billing month. Maximum number of elements: 10.
	ConsumePeriodFilter []*string `json:"ConsumePeriodFilter,omitempty" xml:"ConsumePeriodFilter,omitempty" type:"Repeated"`
	// The code of the financial unit.
	//
	// example:
	//
	// 123#
	CostUnitCode *string `json:"CostUnitCode,omitempty" xml:"CostUnitCode,omitempty"`
	// The maximum number of data records to return in the current call. Default value: 20. Maximum value: 300.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that specifies the position from which the current call starts reading. The value must be empty or set to the NextToken value returned in the previous response. Otherwise, an error is reported. An empty value indicates that the reading starts from the beginning.
	//
	// example:
	//
	// CAESEgoQCg4KCmdtdF9jcmVhdGUEARgBIkgKCQBwhGmPcAEAAAo7AzYAAAAxTDgwMDcxMjg3ZDJhNmM3ZDguTDgwMDAwMDAwMDAwMzE1MTIuTDgwMDcyZDMyZTJkYzg3N2U
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The product code. Reference value sources: Call QueryProductList to query product information, or call DescribeResourcePackageProduct to query the information about a specific resource plan product.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product detail.
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
