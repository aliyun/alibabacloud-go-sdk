// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoScalingActivitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAutoScalingActivitiesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListAutoScalingActivitiesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAutoScalingActivitiesResponseBody
	GetRequestId() *string
	SetScalingActivities(v []*ListAutoScalingActivitiesResponseBodyScalingActivities) *ListAutoScalingActivitiesResponseBody
	GetScalingActivities() []*ListAutoScalingActivitiesResponseBodyScalingActivities
	SetTotalCount(v int32) *ListAutoScalingActivitiesResponseBody
	GetTotalCount() *int32
}

type ListAutoScalingActivitiesResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The scaling activities.
	ScalingActivities []*ListAutoScalingActivitiesResponseBodyScalingActivities `json:"ScalingActivities,omitempty" xml:"ScalingActivities,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAutoScalingActivitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingActivitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAutoScalingActivitiesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAutoScalingActivitiesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAutoScalingActivitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAutoScalingActivitiesResponseBody) GetScalingActivities() []*ListAutoScalingActivitiesResponseBodyScalingActivities {
	return s.ScalingActivities
}

func (s *ListAutoScalingActivitiesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAutoScalingActivitiesResponseBody) SetMaxResults(v int32) *ListAutoScalingActivitiesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBody) SetNextToken(v string) *ListAutoScalingActivitiesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBody) SetRequestId(v string) *ListAutoScalingActivitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBody) SetScalingActivities(v []*ListAutoScalingActivitiesResponseBodyScalingActivities) *ListAutoScalingActivitiesResponseBody {
	s.ScalingActivities = v
	return s
}

func (s *ListAutoScalingActivitiesResponseBody) SetTotalCount(v int32) *ListAutoScalingActivitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAutoScalingActivitiesResponseBodyScalingActivities struct {
	// The ID of the scaling activity.
	//
	// example:
	//
	// asa-36373b084d6b4b13aa50f4129a9e****
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// The status of the scaling activity. Valid values:
	//
	// 	- REJECTED
	//
	// 	- SUCCESSFUL
	//
	// 	- FAILED
	//
	// 	- IN_PROGRESS
	//
	// example:
	//
	// IN_PROGRESS
	ActivityState *string `json:"ActivityState,omitempty" xml:"ActivityState,omitempty"`
	// The type of the scaling activity. Valid values:
	//
	// 	- SCALE_OUT
	//
	// 	- SCALE_IN
	//
	// example:
	//
	// SCALE_OUT
	ActivityType *string `json:"ActivityType,omitempty" xml:"ActivityType,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The description of the scaling activity.
	//
	// example:
	//
	// clusterId not exist
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The end time of the scaling. Unit: milliseconds.
	//
	// example:
	//
	// 1639715634819
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The number of added or removed instances.
	//
	// example:
	//
	// 10
	ExpectNum         *int32            `json:"ExpectNum,omitempty" xml:"ExpectNum,omitempty"`
	InstanceTypeToNum map[string]*int32 `json:"InstanceTypeToNum,omitempty" xml:"InstanceTypeToNum,omitempty"`
	// The ID of the node group.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The name of the node group.
	//
	// example:
	//
	// task-01
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// The operation ID.
	//
	// example:
	//
	// op-13c37a77c505****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	PolicyType  *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The name of the scaling rule.
	//
	// example:
	//
	// scaling-out-memory
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The start time of the scaling. Unit: milliseconds.
	//
	// example:
	//
	// 1639714634819
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListAutoScalingActivitiesResponseBodyScalingActivities) String() string {
	return dara.Prettify(s)
}

func (s ListAutoScalingActivitiesResponseBodyScalingActivities) GoString() string {
	return s.String()
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) GetActivityId() *string {
	return s.ActivityId
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) GetActivityState() *string {
	return s.ActivityState
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) GetActivityType() *string {
	return s.ActivityType
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) GetDescription() *string {
	return s.Description
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) GetExpectNum() *int32 {
	return s.ExpectNum
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) GetInstanceTypeToNum() map[string]*int32 {
	return s.InstanceTypeToNum
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) GetOperationId() *string {
	return s.OperationId
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) GetRuleName() *string {
	return s.RuleName
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) SetActivityId(v string) *ListAutoScalingActivitiesResponseBodyScalingActivities {
	s.ActivityId = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) SetActivityState(v string) *ListAutoScalingActivitiesResponseBodyScalingActivities {
	s.ActivityState = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) SetActivityType(v string) *ListAutoScalingActivitiesResponseBodyScalingActivities {
	s.ActivityType = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) SetClusterId(v string) *ListAutoScalingActivitiesResponseBodyScalingActivities {
	s.ClusterId = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) SetDescription(v string) *ListAutoScalingActivitiesResponseBodyScalingActivities {
	s.Description = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) SetEndTime(v int64) *ListAutoScalingActivitiesResponseBodyScalingActivities {
	s.EndTime = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) SetExpectNum(v int32) *ListAutoScalingActivitiesResponseBodyScalingActivities {
	s.ExpectNum = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) SetInstanceTypeToNum(v map[string]*int32) *ListAutoScalingActivitiesResponseBodyScalingActivities {
	s.InstanceTypeToNum = v
	return s
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) SetNodeGroupId(v string) *ListAutoScalingActivitiesResponseBodyScalingActivities {
	s.NodeGroupId = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) SetNodeGroupName(v string) *ListAutoScalingActivitiesResponseBodyScalingActivities {
	s.NodeGroupName = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) SetOperationId(v string) *ListAutoScalingActivitiesResponseBodyScalingActivities {
	s.OperationId = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) SetPolicyType(v string) *ListAutoScalingActivitiesResponseBodyScalingActivities {
	s.PolicyType = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) SetRuleName(v string) *ListAutoScalingActivitiesResponseBodyScalingActivities {
	s.RuleName = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) SetStartTime(v int64) *ListAutoScalingActivitiesResponseBodyScalingActivities {
	s.StartTime = &v
	return s
}

func (s *ListAutoScalingActivitiesResponseBodyScalingActivities) Validate() error {
	return dara.Validate(s)
}
