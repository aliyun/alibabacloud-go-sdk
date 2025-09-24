// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAmortizedCostByConsumePeriodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmortizationPeriodFilter(v []*string) *DescribeInstanceAmortizedCostByConsumePeriodRequest
	GetAmortizationPeriodFilter() []*string
	SetBillOwnerIdList(v []*string) *DescribeInstanceAmortizedCostByConsumePeriodRequest
	GetBillOwnerIdList() []*string
	SetBillUserIdList(v []*string) *DescribeInstanceAmortizedCostByConsumePeriodRequest
	GetBillUserIdList() []*string
	SetBillingCycle(v string) *DescribeInstanceAmortizedCostByConsumePeriodRequest
	GetBillingCycle() *string
	SetCostUnitCode(v string) *DescribeInstanceAmortizedCostByConsumePeriodRequest
	GetCostUnitCode() *string
	SetInstanceIdList(v []*string) *DescribeInstanceAmortizedCostByConsumePeriodRequest
	GetInstanceIdList() []*string
	SetMaxResults(v int32) *DescribeInstanceAmortizedCostByConsumePeriodRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInstanceAmortizedCostByConsumePeriodRequest
	GetNextToken() *string
	SetProductCode(v string) *DescribeInstanceAmortizedCostByConsumePeriodRequest
	GetProductCode() *string
	SetProductDetail(v string) *DescribeInstanceAmortizedCostByConsumePeriodRequest
	GetProductDetail() *string
	SetSubscriptionType(v string) *DescribeInstanceAmortizedCostByConsumePeriodRequest
	GetSubscriptionType() *string
}

type DescribeInstanceAmortizedCostByConsumePeriodRequest struct {
	AmortizationPeriodFilter []*string `json:"AmortizationPeriodFilter,omitempty" xml:"AmortizationPeriodFilter,omitempty" type:"Repeated"`
	BillOwnerIdList          []*string `json:"BillOwnerIdList,omitempty" xml:"BillOwnerIdList,omitempty" type:"Repeated"`
	BillUserIdList           []*string `json:"BillUserIdList,omitempty" xml:"BillUserIdList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2022-10
	BillingCycle *string `json:"BillingCycle,omitempty" xml:"BillingCycle,omitempty"`
	// example:
	//
	// 123#
	CostUnitCode   *string   `json:"CostUnitCode,omitempty" xml:"CostUnitCode,omitempty"`
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// CAESEgoQCg4KCmdtdF9jcmVhdGUEARgBIkgKCQBwhGmPcAEAAAo7AzYAAAAxTDgwMDcxMjg3ZDJhNmM3ZDguTDgwMDAwMDAwMDAwMzE1MTIuTDgwMDcyZDMyZTJkYzg3N2U
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// rds
	ProductDetail *string `json:"ProductDetail,omitempty" xml:"ProductDetail,omitempty"`
	// example:
	//
	// Subscription
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s DescribeInstanceAmortizedCostByConsumePeriodRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAmortizedCostByConsumePeriodRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) GetAmortizationPeriodFilter() []*string {
	return s.AmortizationPeriodFilter
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) GetBillOwnerIdList() []*string {
	return s.BillOwnerIdList
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) GetBillUserIdList() []*string {
	return s.BillUserIdList
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) GetCostUnitCode() *string {
	return s.CostUnitCode
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) SetAmortizationPeriodFilter(v []*string) *DescribeInstanceAmortizedCostByConsumePeriodRequest {
	s.AmortizationPeriodFilter = v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) SetBillOwnerIdList(v []*string) *DescribeInstanceAmortizedCostByConsumePeriodRequest {
	s.BillOwnerIdList = v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) SetBillUserIdList(v []*string) *DescribeInstanceAmortizedCostByConsumePeriodRequest {
	s.BillUserIdList = v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) SetBillingCycle(v string) *DescribeInstanceAmortizedCostByConsumePeriodRequest {
	s.BillingCycle = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) SetCostUnitCode(v string) *DescribeInstanceAmortizedCostByConsumePeriodRequest {
	s.CostUnitCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) SetInstanceIdList(v []*string) *DescribeInstanceAmortizedCostByConsumePeriodRequest {
	s.InstanceIdList = v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) SetMaxResults(v int32) *DescribeInstanceAmortizedCostByConsumePeriodRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) SetNextToken(v string) *DescribeInstanceAmortizedCostByConsumePeriodRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) SetProductCode(v string) *DescribeInstanceAmortizedCostByConsumePeriodRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) SetProductDetail(v string) *DescribeInstanceAmortizedCostByConsumePeriodRequest {
	s.ProductDetail = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) SetSubscriptionType(v string) *DescribeInstanceAmortizedCostByConsumePeriodRequest {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByConsumePeriodRequest) Validate() error {
	return dara.Validate(s)
}
