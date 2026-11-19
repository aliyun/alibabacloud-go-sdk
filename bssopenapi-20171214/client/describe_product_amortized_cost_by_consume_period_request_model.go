// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductAmortizedCostByConsumePeriodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmortizationPeriodFilter(v []*string) *DescribeProductAmortizedCostByConsumePeriodRequest
	GetAmortizationPeriodFilter() []*string
	SetBillOwnerIdList(v []*string) *DescribeProductAmortizedCostByConsumePeriodRequest
	GetBillOwnerIdList() []*string
	SetBillUserIdList(v []*string) *DescribeProductAmortizedCostByConsumePeriodRequest
	GetBillUserIdList() []*string
	SetBillingCycle(v string) *DescribeProductAmortizedCostByConsumePeriodRequest
	GetBillingCycle() *string
	SetCostUnitCode(v string) *DescribeProductAmortizedCostByConsumePeriodRequest
	GetCostUnitCode() *string
	SetMaxResults(v int32) *DescribeProductAmortizedCostByConsumePeriodRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeProductAmortizedCostByConsumePeriodRequest
	GetNextToken() *string
	SetProductCode(v string) *DescribeProductAmortizedCostByConsumePeriodRequest
	GetProductCode() *string
	SetProductDetail(v string) *DescribeProductAmortizedCostByConsumePeriodRequest
	GetProductDetail() *string
	SetSubscriptionType(v string) *DescribeProductAmortizedCostByConsumePeriodRequest
	GetSubscriptionType() *string
}

type DescribeProductAmortizedCostByConsumePeriodRequest struct {
	// The filter condition for the amortization month.
	//
	// Maximum number of elements: 10.
	AmortizationPeriodFilter []*string `json:"AmortizationPeriodFilter,omitempty" xml:"AmortizationPeriodFilter,omitempty" type:"Repeated"`
	// The filter condition for the resource ownership account.
	//
	// The ID of the resource ownership account. The resource ownership account is the account that actually uses the resources.
	//
	// Maximum number of elements: 10.
	BillOwnerIdList []*string `json:"BillOwnerIdList,omitempty" xml:"BillOwnerIdList,omitempty" type:"Repeated"`
	// The filter condition for the bill purchaser ID.
	//
	// The ID of the RAM user. Specify this parameter to view the bills of RAM users in Financial Cloud. If you do not specify this parameter, the bills of the current caller account and all its RAM users are viewed by default.
	//
	// Maximum number of elements: 10.
	BillUserIdList []*string `json:"BillUserIdList,omitempty" xml:"BillUserIdList,omitempty" type:"Repeated"`
	// The billing cycle month.
	//
	// Format: YYYY-MM.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-10
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
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
	// The token that specifies the position from which the current call starts reading. The value must be empty or set to the NextToken value returned in the previous response. Otherwise, an error occurs. An empty value indicates that the reading starts from the beginning.
	//
	// example:
	//
	// CAESEgoQCg4KCmdtdF9jcmVhdGUEARgBIkgKCQBwhGmPcAEAAAo7AzYAAAAxTDgwMDcxMjg3ZDJhNmM3ZDguTDgwMDAwMDAwMDAwMzE1MTIuTDgwMDcyZDMyZTJkYzg3N2U
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The product code.
	//
	// Reference sources for valid values:
	//
	// Call QueryProductList to query product information,
	//
	// or call DescribeResourcePackageProduct to query the information about a specific resource plan product.
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
	// - Subscription: subscription
	//
	// - PayAsYouGo: pay-as-you-go
	//
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s DescribeProductAmortizedCostByConsumePeriodRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductAmortizedCostByConsumePeriodRequest) GoString() string {
	return s.String()
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) GetAmortizationPeriodFilter() []*string {
	return s.AmortizationPeriodFilter
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) GetBillOwnerIdList() []*string {
	return s.BillOwnerIdList
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) GetBillUserIdList() []*string {
	return s.BillUserIdList
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) GetCostUnitCode() *string {
	return s.CostUnitCode
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) SetAmortizationPeriodFilter(v []*string) *DescribeProductAmortizedCostByConsumePeriodRequest {
	s.AmortizationPeriodFilter = v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) SetBillOwnerIdList(v []*string) *DescribeProductAmortizedCostByConsumePeriodRequest {
	s.BillOwnerIdList = v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) SetBillUserIdList(v []*string) *DescribeProductAmortizedCostByConsumePeriodRequest {
	s.BillUserIdList = v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) SetBillingCycle(v string) *DescribeProductAmortizedCostByConsumePeriodRequest {
	s.BillingCycle = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) SetCostUnitCode(v string) *DescribeProductAmortizedCostByConsumePeriodRequest {
	s.CostUnitCode = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) SetMaxResults(v int32) *DescribeProductAmortizedCostByConsumePeriodRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) SetNextToken(v string) *DescribeProductAmortizedCostByConsumePeriodRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) SetProductCode(v string) *DescribeProductAmortizedCostByConsumePeriodRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) SetProductDetail(v string) *DescribeProductAmortizedCostByConsumePeriodRequest {
	s.ProductDetail = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) SetSubscriptionType(v string) *DescribeProductAmortizedCostByConsumePeriodRequest {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeProductAmortizedCostByConsumePeriodRequest) Validate() error {
	return dara.Validate(s)
}
