// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlanMaintenanceWindowsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *DescribePlanMaintenanceWindowsShrinkRequest
	GetEnable() *bool
	SetMaxResults(v int32) *DescribePlanMaintenanceWindowsShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePlanMaintenanceWindowsShrinkRequest
	GetNextToken() *string
	SetPlanWindowId(v string) *DescribePlanMaintenanceWindowsShrinkRequest
	GetPlanWindowId() *string
	SetPlanWindowName(v string) *DescribePlanMaintenanceWindowsShrinkRequest
	GetPlanWindowName() *string
	SetRegionId(v string) *DescribePlanMaintenanceWindowsShrinkRequest
	GetRegionId() *string
	SetTargetResourceGroupId(v string) *DescribePlanMaintenanceWindowsShrinkRequest
	GetTargetResourceGroupId() *string
	SetTargetResourceTagsShrink(v string) *DescribePlanMaintenanceWindowsShrinkRequest
	GetTargetResourceTagsShrink() *string
}

type DescribePlanMaintenanceWindowsShrinkRequest struct {
	// Indicates whether the maintenance window is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the last query as the value of NextToken.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the maintenance window.
	//
	// example:
	//
	// pw-bp1jarob1aup7yvlrdt6
	PlanWindowId *string `json:"PlanWindowId,omitempty" xml:"PlanWindowId,omitempty"`
	// The name of the maintenance window.
	//
	// example:
	//
	// WIndowName
	PlanWindowName *string `json:"PlanWindowName,omitempty" xml:"PlanWindowName,omitempty"`
	// The ID of the region where the ECS instance is located. You can call the DescribeRegions operation to query the latest list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the maintenance window applies.
	//
	// example:
	//
	// rg-d85g5yocioezmdrll
	TargetResourceGroupId *string `json:"TargetResourceGroupId,omitempty" xml:"TargetResourceGroupId,omitempty"`
	// The tags of the resources to which the maintenance window applies.
	TargetResourceTagsShrink *string `json:"TargetResourceTags,omitempty" xml:"TargetResourceTags,omitempty"`
}

func (s DescribePlanMaintenanceWindowsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlanMaintenanceWindowsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetEnable() *bool {
	return s.Enable
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetPlanWindowId() *string {
	return s.PlanWindowId
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetPlanWindowName() *string {
	return s.PlanWindowName
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetTargetResourceGroupId() *string {
	return s.TargetResourceGroupId
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) GetTargetResourceTagsShrink() *string {
	return s.TargetResourceTagsShrink
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetEnable(v bool) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.Enable = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetMaxResults(v int32) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetNextToken(v string) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetPlanWindowId(v string) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.PlanWindowId = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetPlanWindowName(v string) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.PlanWindowName = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetRegionId(v string) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetTargetResourceGroupId(v string) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.TargetResourceGroupId = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) SetTargetResourceTagsShrink(v string) *DescribePlanMaintenanceWindowsShrinkRequest {
	s.TargetResourceTagsShrink = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
