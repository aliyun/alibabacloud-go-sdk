// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoScalingActivitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListAutoScalingActivitiesRequest
	GetClusterId() *string
	SetEndTime(v int64) *ListAutoScalingActivitiesRequest
	GetEndTime() *int64
	SetMaxResults(v int32) *ListAutoScalingActivitiesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAutoScalingActivitiesRequest
	GetNextToken() *string
	SetNodeGroupId(v string) *ListAutoScalingActivitiesRequest
	GetNodeGroupId() *string
	SetRegionId(v string) *ListAutoScalingActivitiesRequest
	GetRegionId() *string
	SetScalingActivityStates(v []*string) *ListAutoScalingActivitiesRequest
	GetScalingActivityStates() []*string
	SetScalingActivityType(v string) *ListAutoScalingActivitiesRequest
	GetScalingActivityType() *string
	SetScalingPolicyType(v string) *ListAutoScalingActivitiesRequest
	GetScalingPolicyType() *string
	SetScalingRuleName(v string) *ListAutoScalingActivitiesRequest
	GetScalingRuleName() *string
	SetStartTime(v int64) *ListAutoScalingActivitiesRequest
	GetStartTime() *int64
}

type ListAutoScalingActivitiesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The end of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1639718634819
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the node group.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the scaling activity. Number of elements in the array: 1-20.
	//
	// example:
	//
	// ["REJECTED","SUCCESSFUL"]
	ScalingActivityStates []*string `json:"ScalingActivityStates,omitempty" xml:"ScalingActivityStates,omitempty" type:"Repeated"`
	// The type of the scaling activity. Valid values:
	//
	// 	- SCALE_OUT
	//
	// 	- SCALE_IN
	//
	// example:
	//
	// SCALE_IN
	ScalingActivityType *string `json:"ScalingActivityType,omitempty" xml:"ScalingActivityType,omitempty"`
	ScalingPolicyType   *string `json:"ScalingPolicyType,omitempty" xml:"ScalingPolicyType,omitempty"`
	// The name of the scaling rule.
	//
	// example:
	//
	// scale-out-by-memroy
	ScalingRuleName *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	// The beginning of the time range to query. Unit: milliseconds.
	//
	// example:
	//
	// 1639714634819
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAutoScalingActivitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingActivitiesRequest) GoString() string {
	return s.String()
}

func (s *ListAutoScalingActivitiesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListAutoScalingActivitiesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAutoScalingActivitiesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAutoScalingActivitiesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAutoScalingActivitiesRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListAutoScalingActivitiesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAutoScalingActivitiesRequest) GetScalingActivityStates() []*string {
	return s.ScalingActivityStates
}

func (s *ListAutoScalingActivitiesRequest) GetScalingActivityType() *string {
	return s.ScalingActivityType
}

func (s *ListAutoScalingActivitiesRequest) GetScalingPolicyType() *string {
	return s.ScalingPolicyType
}

func (s *ListAutoScalingActivitiesRequest) GetScalingRuleName() *string {
	return s.ScalingRuleName
}

func (s *ListAutoScalingActivitiesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAutoScalingActivitiesRequest) SetClusterId(v string) *ListAutoScalingActivitiesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAutoScalingActivitiesRequest) SetEndTime(v int64) *ListAutoScalingActivitiesRequest {
	s.EndTime = &v
	return s
}

func (s *ListAutoScalingActivitiesRequest) SetMaxResults(v int32) *ListAutoScalingActivitiesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAutoScalingActivitiesRequest) SetNextToken(v string) *ListAutoScalingActivitiesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAutoScalingActivitiesRequest) SetNodeGroupId(v string) *ListAutoScalingActivitiesRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ListAutoScalingActivitiesRequest) SetRegionId(v string) *ListAutoScalingActivitiesRequest {
	s.RegionId = &v
	return s
}

func (s *ListAutoScalingActivitiesRequest) SetScalingActivityStates(v []*string) *ListAutoScalingActivitiesRequest {
	s.ScalingActivityStates = v
	return s
}

func (s *ListAutoScalingActivitiesRequest) SetScalingActivityType(v string) *ListAutoScalingActivitiesRequest {
	s.ScalingActivityType = &v
	return s
}

func (s *ListAutoScalingActivitiesRequest) SetScalingPolicyType(v string) *ListAutoScalingActivitiesRequest {
	s.ScalingPolicyType = &v
	return s
}

func (s *ListAutoScalingActivitiesRequest) SetScalingRuleName(v string) *ListAutoScalingActivitiesRequest {
	s.ScalingRuleName = &v
	return s
}

func (s *ListAutoScalingActivitiesRequest) SetStartTime(v int64) *ListAutoScalingActivitiesRequest {
	s.StartTime = &v
	return s
}

func (s *ListAutoScalingActivitiesRequest) Validate() error {
	return dara.Validate(s)
}
