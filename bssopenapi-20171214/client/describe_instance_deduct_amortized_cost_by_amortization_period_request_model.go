// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillOwnerIdList(v []*string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest
	GetBillOwnerIdList() []*string
	SetBillUserIdList(v []*string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest
	GetBillUserIdList() []*string
	SetBillingCycle(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest
	GetBillingCycle() *string
	SetCostUnitCode(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest
	GetCostUnitCode() *string
	SetInstanceIdList(v []*string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest
	GetInstanceIdList() []*string
	SetMaxResults(v int32) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest
	GetNextToken() *string
	SetProductCode(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest
	GetProductCode() *string
	SetProductDetail(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest
	GetProductDetail() *string
	SetSubscriptionType(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest
	GetSubscriptionType() *string
}

type DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest struct {
	BillOwnerIdList []*string `json:"BillOwnerIdList,omitempty" xml:"BillOwnerIdList,omitempty" type:"Repeated"`
	BillUserIdList  []*string `json:"BillUserIdList,omitempty" xml:"BillUserIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-04
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
	// PayAsYouGo
	SubscriptionType *string `json:"SubscriptionType,omitempty" xml:"SubscriptionType,omitempty"`
}

func (s DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) GetBillOwnerIdList() []*string {
	return s.BillOwnerIdList
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) GetBillUserIdList() []*string {
	return s.BillUserIdList
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) GetCostUnitCode() *string {
	return s.CostUnitCode
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) SetBillOwnerIdList(v []*string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest {
	s.BillOwnerIdList = v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) SetBillUserIdList(v []*string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest {
	s.BillUserIdList = v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) SetBillingCycle(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest {
	s.BillingCycle = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) SetCostUnitCode(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest {
	s.CostUnitCode = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) SetInstanceIdList(v []*string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest {
	s.InstanceIdList = v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) SetMaxResults(v int32) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) SetNextToken(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) SetProductCode(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) SetProductDetail(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest {
	s.ProductDetail = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) SetSubscriptionType(v string) *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeInstanceDeductAmortizedCostByAmortizationPeriodRequest) Validate() error {
	return dara.Validate(s)
}
