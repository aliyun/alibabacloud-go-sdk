// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBudgetUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBudgetDimensionRefId(v string) *DescribeBudgetUsageRequest
	GetBudgetDimensionRefId() *string
	SetBudgetDimensionType(v string) *DescribeBudgetUsageRequest
	GetBudgetDimensionType() *string
	SetGwClusterId(v string) *DescribeBudgetUsageRequest
	GetGwClusterId() *string
	SetRegionId(v string) *DescribeBudgetUsageRequest
	GetRegionId() *string
}

type DescribeBudgetUsageRequest struct {
	// The budget dimension reference ID.
	//
	// example:
	//
	// ref-001
	BudgetDimensionRefId *string `json:"BudgetDimensionRefId,omitempty" xml:"BudgetDimensionRefId,omitempty"`
	// The budget dimension type.
	//
	// example:
	//
	// cluster
	BudgetDimensionType *string `json:"BudgetDimensionType,omitempty" xml:"BudgetDimensionType,omitempty"`
	// The ID of the gateway instance.
	//
	// example:
	//
	// pc-123
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The region ID of the gateway instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBudgetUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBudgetUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeBudgetUsageRequest) GetBudgetDimensionRefId() *string {
	return s.BudgetDimensionRefId
}

func (s *DescribeBudgetUsageRequest) GetBudgetDimensionType() *string {
	return s.BudgetDimensionType
}

func (s *DescribeBudgetUsageRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeBudgetUsageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBudgetUsageRequest) SetBudgetDimensionRefId(v string) *DescribeBudgetUsageRequest {
	s.BudgetDimensionRefId = &v
	return s
}

func (s *DescribeBudgetUsageRequest) SetBudgetDimensionType(v string) *DescribeBudgetUsageRequest {
	s.BudgetDimensionType = &v
	return s
}

func (s *DescribeBudgetUsageRequest) SetGwClusterId(v string) *DescribeBudgetUsageRequest {
	s.GwClusterId = &v
	return s
}

func (s *DescribeBudgetUsageRequest) SetRegionId(v string) *DescribeBudgetUsageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBudgetUsageRequest) Validate() error {
	return dara.Validate(s)
}
