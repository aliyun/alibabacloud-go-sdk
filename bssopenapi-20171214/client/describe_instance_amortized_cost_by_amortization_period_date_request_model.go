// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAmortizedCostByAmortizationPeriodDateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmortizationDateEnd(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest
	GetAmortizationDateEnd() *string
	SetAmortizationDateStart(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest
	GetAmortizationDateStart() *string
	SetBillOwnerIdList(v []*string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest
	GetBillOwnerIdList() []*string
	SetBillUserIdList(v []*string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest
	GetBillUserIdList() []*string
	SetBillingCycle(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest
	GetBillingCycle() *string
	SetCostUnitCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest
	GetCostUnitCode() *string
	SetInstanceIdList(v []*string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest
	GetInstanceIdList() []*string
	SetMaxResults(v int32) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest
	GetNextToken() *string
	SetProductCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest
	GetProductCode() *string
	SetProductDetail(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest
	GetProductDetail() *string
	SetSubscriptionType(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest
	GetSubscriptionType() *string
}

type DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-10
	AmortizationDateEnd *string `json:"AmortizationDateEnd,omitempty" xml:"AmortizationDateEnd,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2023-05-10
	AmortizationDateStart *string   `json:"AmortizationDateStart,omitempty" xml:"AmortizationDateStart,omitempty"`
	BillOwnerIdList       []*string `json:"BillOwnerIdList,omitempty" xml:"BillOwnerIdList,omitempty" type:"Repeated"`
	BillUserIdList        []*string `json:"BillUserIdList,omitempty" xml:"BillUserIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 2023-05
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

func (s DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) GetAmortizationDateEnd() *string {
	return s.AmortizationDateEnd
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) GetAmortizationDateStart() *string {
	return s.AmortizationDateStart
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) GetBillOwnerIdList() []*string {
	return s.BillOwnerIdList
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) GetBillUserIdList() []*string {
	return s.BillUserIdList
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) GetBillingCycle() *string {
	return s.BillingCycle
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) GetCostUnitCode() *string {
	return s.CostUnitCode
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) GetProductDetail() *string {
	return s.ProductDetail
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) GetSubscriptionType() *string {
	return s.SubscriptionType
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) SetAmortizationDateEnd(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest {
	s.AmortizationDateEnd = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) SetAmortizationDateStart(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest {
	s.AmortizationDateStart = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) SetBillOwnerIdList(v []*string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest {
	s.BillOwnerIdList = v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) SetBillUserIdList(v []*string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest {
	s.BillUserIdList = v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) SetBillingCycle(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest {
	s.BillingCycle = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) SetCostUnitCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest {
	s.CostUnitCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) SetInstanceIdList(v []*string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest {
	s.InstanceIdList = v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) SetMaxResults(v int32) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) SetNextToken(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) SetProductCode(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) SetProductDetail(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest {
	s.ProductDetail = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) SetSubscriptionType(v string) *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest {
	s.SubscriptionType = &v
	return s
}

func (s *DescribeInstanceAmortizedCostByAmortizationPeriodDateRequest) Validate() error {
	return dara.Validate(s)
}
