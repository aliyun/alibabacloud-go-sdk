// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlanMaintenanceWindowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *DescribePlanMaintenanceWindowsRequest
	GetEnable() *bool
	SetMaxResults(v int32) *DescribePlanMaintenanceWindowsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePlanMaintenanceWindowsRequest
	GetNextToken() *string
	SetPlanWindowId(v string) *DescribePlanMaintenanceWindowsRequest
	GetPlanWindowId() *string
	SetPlanWindowName(v string) *DescribePlanMaintenanceWindowsRequest
	GetPlanWindowName() *string
	SetRegionId(v string) *DescribePlanMaintenanceWindowsRequest
	GetRegionId() *string
	SetTargetResourceGroupId(v string) *DescribePlanMaintenanceWindowsRequest
	GetTargetResourceGroupId() *string
	SetTargetResourceTags(v *DescribePlanMaintenanceWindowsRequestTargetResourceTags) *DescribePlanMaintenanceWindowsRequest
	GetTargetResourceTags() *DescribePlanMaintenanceWindowsRequestTargetResourceTags
}

type DescribePlanMaintenanceWindowsRequest struct {
	// Specifies whether the window is enabled or disabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The number of entries per page for a paged query. Maximum value: 100. Default value: If the value is not specified or is less than 10, the default value is 10. If the value is greater than 100, the default value is 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the O&M window.
	//
	// example:
	//
	// pw-bp1jarob1aup7yvlrdt6
	PlanWindowId *string `json:"PlanWindowId,omitempty" xml:"PlanWindowId,omitempty"`
	// The name of the O&M window.
	//
	// example:
	//
	// WIndowName
	PlanWindowName *string `json:"PlanWindowName,omitempty" xml:"PlanWindowName,omitempty"`
	// The region ID of the instance. You can call DescribeRegions to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the O&M window applies.
	//
	// example:
	//
	// rg-d85g5yocioezmdrll
	TargetResourceGroupId *string `json:"TargetResourceGroupId,omitempty" xml:"TargetResourceGroupId,omitempty"`
	// The tags to which the O&M window applies.
	TargetResourceTags *DescribePlanMaintenanceWindowsRequestTargetResourceTags `json:"TargetResourceTags,omitempty" xml:"TargetResourceTags,omitempty" type:"Struct"`
}

func (s DescribePlanMaintenanceWindowsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePlanMaintenanceWindowsRequest) GoString() string {
	return s.String()
}

func (s *DescribePlanMaintenanceWindowsRequest) GetEnable() *bool {
	return s.Enable
}

func (s *DescribePlanMaintenanceWindowsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePlanMaintenanceWindowsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePlanMaintenanceWindowsRequest) GetPlanWindowId() *string {
	return s.PlanWindowId
}

func (s *DescribePlanMaintenanceWindowsRequest) GetPlanWindowName() *string {
	return s.PlanWindowName
}

func (s *DescribePlanMaintenanceWindowsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePlanMaintenanceWindowsRequest) GetTargetResourceGroupId() *string {
	return s.TargetResourceGroupId
}

func (s *DescribePlanMaintenanceWindowsRequest) GetTargetResourceTags() *DescribePlanMaintenanceWindowsRequestTargetResourceTags {
	return s.TargetResourceTags
}

func (s *DescribePlanMaintenanceWindowsRequest) SetEnable(v bool) *DescribePlanMaintenanceWindowsRequest {
	s.Enable = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsRequest) SetMaxResults(v int32) *DescribePlanMaintenanceWindowsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsRequest) SetNextToken(v string) *DescribePlanMaintenanceWindowsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsRequest) SetPlanWindowId(v string) *DescribePlanMaintenanceWindowsRequest {
	s.PlanWindowId = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsRequest) SetPlanWindowName(v string) *DescribePlanMaintenanceWindowsRequest {
	s.PlanWindowName = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsRequest) SetRegionId(v string) *DescribePlanMaintenanceWindowsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsRequest) SetTargetResourceGroupId(v string) *DescribePlanMaintenanceWindowsRequest {
	s.TargetResourceGroupId = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsRequest) SetTargetResourceTags(v *DescribePlanMaintenanceWindowsRequestTargetResourceTags) *DescribePlanMaintenanceWindowsRequest {
	s.TargetResourceTags = v
	return s
}

func (s *DescribePlanMaintenanceWindowsRequest) Validate() error {
	if s.TargetResourceTags != nil {
		if err := s.TargetResourceTags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePlanMaintenanceWindowsRequestTargetResourceTags struct {
	// The key of the tag to which the window applies.
	//
	// example:
	//
	// tagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag to which the window applies.
	//
	// example:
	//
	// tagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePlanMaintenanceWindowsRequestTargetResourceTags) String() string {
	return dara.Prettify(s)
}

func (s DescribePlanMaintenanceWindowsRequestTargetResourceTags) GoString() string {
	return s.String()
}

func (s *DescribePlanMaintenanceWindowsRequestTargetResourceTags) GetKey() *string {
	return s.Key
}

func (s *DescribePlanMaintenanceWindowsRequestTargetResourceTags) GetValue() *string {
	return s.Value
}

func (s *DescribePlanMaintenanceWindowsRequestTargetResourceTags) SetKey(v string) *DescribePlanMaintenanceWindowsRequestTargetResourceTags {
	s.Key = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsRequestTargetResourceTags) SetValue(v string) *DescribePlanMaintenanceWindowsRequestTargetResourceTags {
	s.Value = &v
	return s
}

func (s *DescribePlanMaintenanceWindowsRequestTargetResourceTags) Validate() error {
	return dara.Validate(s)
}
