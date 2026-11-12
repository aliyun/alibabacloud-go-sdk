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
	// The reference ID of the budget dimension. If BudgetDimensionType is set to Global, set this parameter to Global. If BudgetDimensionType is set to Consumer, set this parameter to the consumer ID returned by DescribeBudgetPolicies, such as consumer-xxx. If BudgetDimensionType is set to ConsumerGroup, set this parameter to the consumer group ID, such as cg-xxx.
	//
	// example:
	//
	// Global
	BudgetDimensionRefId *string `json:"BudgetDimensionRefId,omitempty" xml:"BudgetDimensionRefId,omitempty"`
	// The budget dimension type. Valid values: Global (global budget), Consumer (consumer budget), and ConsumerGroup (consumer group budget).
	//
	// example:
	//
	// Global
	BudgetDimensionType *string `json:"BudgetDimensionType,omitempty" xml:"BudgetDimensionType,omitempty"`
	// The ID of the gateway instance to query. This parameter is required. If this parameter is not specified, the API returns InvalidPolarGatewayInstanceId.NotFound (404).
	//
	// example:
	//
	// pc-123
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The ID of the region where the gateway instance resides.
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
