// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBudgetPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetDimensionRefId(v string) *DescribeBudgetPoliciesRequest
	GetBudgetDimensionRefId() *string
	SetBudgetDimensionType(v string) *DescribeBudgetPoliciesRequest
	GetBudgetDimensionType() *string
	SetBudgetPolicyId(v string) *DescribeBudgetPoliciesRequest
	GetBudgetPolicyId() *string
	SetGwClusterId(v string) *DescribeBudgetPoliciesRequest
	GetGwClusterId() *string
	SetPageNumber(v int32) *DescribeBudgetPoliciesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBudgetPoliciesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeBudgetPoliciesRequest
	GetRegionId() *string
	SetStatus(v string) *DescribeBudgetPoliciesRequest
	GetStatus() *string
}

type DescribeBudgetPoliciesRequest struct {
	// example:
	//
	// cg-p3gk2oh55c**
	BudgetDimensionRefId *string `json:"BudgetDimensionRefId,omitempty" xml:"BudgetDimensionRefId,omitempty"`
	// example:
	//
	// ConsumerGroup
	BudgetDimensionType *string `json:"BudgetDimensionType,omitempty" xml:"BudgetDimensionType,omitempty"`
	// example:
	//
	// 023aacc1effc4b56bb154bfbec6baxxx
	BudgetPolicyId *string `json:"BudgetPolicyId,omitempty" xml:"BudgetPolicyId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeBudgetPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeBudgetPoliciesRequest) GetBudgetDimensionRefId() *string {
	return s.BudgetDimensionRefId
}

func (s *DescribeBudgetPoliciesRequest) GetBudgetDimensionType() *string {
	return s.BudgetDimensionType
}

func (s *DescribeBudgetPoliciesRequest) GetBudgetPolicyId() *string {
	return s.BudgetPolicyId
}

func (s *DescribeBudgetPoliciesRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeBudgetPoliciesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBudgetPoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBudgetPoliciesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBudgetPoliciesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeBudgetPoliciesRequest) SetBudgetDimensionRefId(v string) *DescribeBudgetPoliciesRequest {
	s.BudgetDimensionRefId = &v
	return s
}

func (s *DescribeBudgetPoliciesRequest) SetBudgetDimensionType(v string) *DescribeBudgetPoliciesRequest {
	s.BudgetDimensionType = &v
	return s
}

func (s *DescribeBudgetPoliciesRequest) SetBudgetPolicyId(v string) *DescribeBudgetPoliciesRequest {
	s.BudgetPolicyId = &v
	return s
}

func (s *DescribeBudgetPoliciesRequest) SetGwClusterId(v string) *DescribeBudgetPoliciesRequest {
	s.GwClusterId = &v
	return s
}

func (s *DescribeBudgetPoliciesRequest) SetPageNumber(v int32) *DescribeBudgetPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBudgetPoliciesRequest) SetPageSize(v int32) *DescribeBudgetPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBudgetPoliciesRequest) SetRegionId(v string) *DescribeBudgetPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBudgetPoliciesRequest) SetStatus(v string) *DescribeBudgetPoliciesRequest {
	s.Status = &v
	return s
}

func (s *DescribeBudgetPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
