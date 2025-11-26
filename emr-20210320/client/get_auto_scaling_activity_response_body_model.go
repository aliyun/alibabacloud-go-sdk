// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoScalingActivityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAutoScalingActivityResponseBody
	GetRequestId() *string
	SetScalingActivity(v *GetAutoScalingActivityResponseBodyScalingActivity) *GetAutoScalingActivityResponseBody
	GetScalingActivity() *GetAutoScalingActivityResponseBodyScalingActivity
}

type GetAutoScalingActivityResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the scaling activity.
	ScalingActivity *GetAutoScalingActivityResponseBodyScalingActivity `json:"ScalingActivity,omitempty" xml:"ScalingActivity,omitempty" type:"Struct"`
}

func (s GetAutoScalingActivityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingActivityResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoScalingActivityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutoScalingActivityResponseBody) GetScalingActivity() *GetAutoScalingActivityResponseBodyScalingActivity {
	return s.ScalingActivity
}

func (s *GetAutoScalingActivityResponseBody) SetRequestId(v string) *GetAutoScalingActivityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoScalingActivityResponseBody) SetScalingActivity(v *GetAutoScalingActivityResponseBodyScalingActivity) *GetAutoScalingActivityResponseBody {
	s.ScalingActivity = v
	return s
}

func (s *GetAutoScalingActivityResponseBody) Validate() error {
	if s.ScalingActivity != nil {
		if err := s.ScalingActivity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutoScalingActivityResponseBodyScalingActivity struct {
	// The ID of the scaling activity.
	//
	// example:
	//
	// asa-36373b084d6b4b13aa50f4129a9e****
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// The instances that correspond to the scaling activity.
	ActivityResults []*ScalingActivityResult `json:"ActivityResults,omitempty" xml:"ActivityResults,omitempty" type:"Repeated"`
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
	// The type of the scaling activity. Valid value:
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
	// The time when scaling ended.
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
	ExpectNum *int32 `json:"ExpectNum,omitempty" xml:"ExpectNum,omitempty"`
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
	// The policy type.
	//
	// example:
	//
	// AUTO
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The description of the scaling rule.
	RuleDetail *ScalingRule `json:"RuleDetail,omitempty" xml:"RuleDetail,omitempty"`
	// The name of the scaling rule.
	//
	// example:
	//
	// scaling-out-memory
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The time when scaling started.
	//
	// example:
	//
	// 1639714634819
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetAutoScalingActivityResponseBodyScalingActivity) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingActivityResponseBodyScalingActivity) GoString() string {
	return s.String()
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) GetActivityId() *string {
	return s.ActivityId
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) GetActivityResults() []*ScalingActivityResult {
	return s.ActivityResults
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) GetActivityState() *string {
	return s.ActivityState
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) GetActivityType() *string {
	return s.ActivityType
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) GetDescription() *string {
	return s.Description
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) GetExpectNum() *int32 {
	return s.ExpectNum
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) GetOperationId() *string {
	return s.OperationId
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) GetRuleDetail() *ScalingRule {
	return s.RuleDetail
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) GetRuleName() *string {
	return s.RuleName
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) SetActivityId(v string) *GetAutoScalingActivityResponseBodyScalingActivity {
	s.ActivityId = &v
	return s
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) SetActivityResults(v []*ScalingActivityResult) *GetAutoScalingActivityResponseBodyScalingActivity {
	s.ActivityResults = v
	return s
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) SetActivityState(v string) *GetAutoScalingActivityResponseBodyScalingActivity {
	s.ActivityState = &v
	return s
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) SetActivityType(v string) *GetAutoScalingActivityResponseBodyScalingActivity {
	s.ActivityType = &v
	return s
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) SetClusterId(v string) *GetAutoScalingActivityResponseBodyScalingActivity {
	s.ClusterId = &v
	return s
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) SetDescription(v string) *GetAutoScalingActivityResponseBodyScalingActivity {
	s.Description = &v
	return s
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) SetEndTime(v int64) *GetAutoScalingActivityResponseBodyScalingActivity {
	s.EndTime = &v
	return s
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) SetExpectNum(v int32) *GetAutoScalingActivityResponseBodyScalingActivity {
	s.ExpectNum = &v
	return s
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) SetNodeGroupId(v string) *GetAutoScalingActivityResponseBodyScalingActivity {
	s.NodeGroupId = &v
	return s
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) SetNodeGroupName(v string) *GetAutoScalingActivityResponseBodyScalingActivity {
	s.NodeGroupName = &v
	return s
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) SetOperationId(v string) *GetAutoScalingActivityResponseBodyScalingActivity {
	s.OperationId = &v
	return s
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) SetPolicyType(v string) *GetAutoScalingActivityResponseBodyScalingActivity {
	s.PolicyType = &v
	return s
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) SetRuleDetail(v *ScalingRule) *GetAutoScalingActivityResponseBodyScalingActivity {
	s.RuleDetail = v
	return s
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) SetRuleName(v string) *GetAutoScalingActivityResponseBodyScalingActivity {
	s.RuleName = &v
	return s
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) SetStartTime(v int64) *GetAutoScalingActivityResponseBodyScalingActivity {
	s.StartTime = &v
	return s
}

func (s *GetAutoScalingActivityResponseBodyScalingActivity) Validate() error {
	if s.ActivityResults != nil {
		for _, item := range s.ActivityResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RuleDetail != nil {
		if err := s.RuleDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}
