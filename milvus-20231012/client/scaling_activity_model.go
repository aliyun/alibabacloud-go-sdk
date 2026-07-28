// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingActivity interface {
	dara.Model
	String() string
	GoString() string
	SetComponentTypes(v string) *ScalingActivity
	GetComponentTypes() *string
	SetDescription(v string) *ScalingActivity
	GetDescription() *string
	SetEndTime(v string) *ScalingActivity
	GetEndTime() *string
	SetInstanceId(v string) *ScalingActivity
	GetInstanceId() *string
	SetPolicyType(v string) *ScalingActivity
	GetPolicyType() *string
	SetScalingActivityId(v string) *ScalingActivity
	GetScalingActivityId() *string
	SetScalingActivityState(v string) *ScalingActivity
	GetScalingActivityState() *string
	SetScalingPolicyId(v string) *ScalingActivity
	GetScalingPolicyId() *string
	SetScalingRuleDetail(v string) *ScalingActivity
	GetScalingRuleDetail() *string
	SetScalingRuleId(v string) *ScalingActivity
	GetScalingRuleId() *string
	SetScalingRuleName(v string) *ScalingActivity
	GetScalingRuleName() *string
	SetStartTime(v string) *ScalingActivity
	GetStartTime() *string
	SetTimeZone(v string) *ScalingActivity
	GetTimeZone() *string
}

type ScalingActivity struct {
	// The types of components involved in the scaling activity.
	ComponentTypes *string `json:"componentTypes,omitempty" xml:"componentTypes,omitempty"`
	// The description of the scaling activity.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The end time of the scaling activity.
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The type of the scaling policy.
	PolicyType *string `json:"policyType,omitempty" xml:"policyType,omitempty"`
	// The ID of the scaling activity.
	ScalingActivityId *string `json:"scalingActivityId,omitempty" xml:"scalingActivityId,omitempty"`
	// The state of the scaling activity.
	ScalingActivityState *string `json:"scalingActivityState,omitempty" xml:"scalingActivityState,omitempty"`
	// The ID of the scaling policy.
	ScalingPolicyId *string `json:"scalingPolicyId,omitempty" xml:"scalingPolicyId,omitempty"`
	// The details of the scaling rule.
	ScalingRuleDetail *string `json:"scalingRuleDetail,omitempty" xml:"scalingRuleDetail,omitempty"`
	// The ID of the scaling rule.
	ScalingRuleId *string `json:"scalingRuleId,omitempty" xml:"scalingRuleId,omitempty"`
	// The name of the scaling rule that triggered the activity.
	ScalingRuleName *string `json:"scalingRuleName,omitempty" xml:"scalingRuleName,omitempty"`
	// The start time of the scaling activity.
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The time zone of the scaling activity.
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s ScalingActivity) String() string {
	return dara.Prettify(s)
}

func (s ScalingActivity) GoString() string {
	return s.String()
}

func (s *ScalingActivity) GetComponentTypes() *string {
	return s.ComponentTypes
}

func (s *ScalingActivity) GetDescription() *string {
	return s.Description
}

func (s *ScalingActivity) GetEndTime() *string {
	return s.EndTime
}

func (s *ScalingActivity) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ScalingActivity) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ScalingActivity) GetScalingActivityId() *string {
	return s.ScalingActivityId
}

func (s *ScalingActivity) GetScalingActivityState() *string {
	return s.ScalingActivityState
}

func (s *ScalingActivity) GetScalingPolicyId() *string {
	return s.ScalingPolicyId
}

func (s *ScalingActivity) GetScalingRuleDetail() *string {
	return s.ScalingRuleDetail
}

func (s *ScalingActivity) GetScalingRuleId() *string {
	return s.ScalingRuleId
}

func (s *ScalingActivity) GetScalingRuleName() *string {
	return s.ScalingRuleName
}

func (s *ScalingActivity) GetStartTime() *string {
	return s.StartTime
}

func (s *ScalingActivity) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ScalingActivity) SetComponentTypes(v string) *ScalingActivity {
	s.ComponentTypes = &v
	return s
}

func (s *ScalingActivity) SetDescription(v string) *ScalingActivity {
	s.Description = &v
	return s
}

func (s *ScalingActivity) SetEndTime(v string) *ScalingActivity {
	s.EndTime = &v
	return s
}

func (s *ScalingActivity) SetInstanceId(v string) *ScalingActivity {
	s.InstanceId = &v
	return s
}

func (s *ScalingActivity) SetPolicyType(v string) *ScalingActivity {
	s.PolicyType = &v
	return s
}

func (s *ScalingActivity) SetScalingActivityId(v string) *ScalingActivity {
	s.ScalingActivityId = &v
	return s
}

func (s *ScalingActivity) SetScalingActivityState(v string) *ScalingActivity {
	s.ScalingActivityState = &v
	return s
}

func (s *ScalingActivity) SetScalingPolicyId(v string) *ScalingActivity {
	s.ScalingPolicyId = &v
	return s
}

func (s *ScalingActivity) SetScalingRuleDetail(v string) *ScalingActivity {
	s.ScalingRuleDetail = &v
	return s
}

func (s *ScalingActivity) SetScalingRuleId(v string) *ScalingActivity {
	s.ScalingRuleId = &v
	return s
}

func (s *ScalingActivity) SetScalingRuleName(v string) *ScalingActivity {
	s.ScalingRuleName = &v
	return s
}

func (s *ScalingActivity) SetStartTime(v string) *ScalingActivity {
	s.StartTime = &v
	return s
}

func (s *ScalingActivity) SetTimeZone(v string) *ScalingActivity {
	s.TimeZone = &v
	return s
}

func (s *ScalingActivity) Validate() error {
	return dara.Validate(s)
}
