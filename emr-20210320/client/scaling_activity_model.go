// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingActivity interface {
	dara.Model
	String() string
	GoString() string
	SetCause(v string) *ScalingActivity
	GetCause() *string
	SetDescription(v string) *ScalingActivity
	GetDescription() *string
	SetEndTime(v int64) *ScalingActivity
	GetEndTime() *int64
	SetEssScalingRuleId(v string) *ScalingActivity
	GetEssScalingRuleId() *string
	SetExpectNum(v int32) *ScalingActivity
	GetExpectNum() *int32
	SetHostGroupName(v string) *ScalingActivity
	GetHostGroupName() *string
	SetId(v string) *ScalingActivity
	GetId() *string
	SetInstanceIds(v string) *ScalingActivity
	GetInstanceIds() *string
	SetScalingGroupId(v string) *ScalingActivity
	GetScalingGroupId() *string
	SetScalingRuleName(v string) *ScalingActivity
	GetScalingRuleName() *string
	SetStartTime(v int64) *ScalingActivity
	GetStartTime() *int64
	SetStatus(v string) *ScalingActivity
	GetStatus() *string
	SetTotalCapacity(v int32) *ScalingActivity
	GetTotalCapacity() *int32
	SetTransition(v string) *ScalingActivity
	GetTransition() *string
}

type ScalingActivity struct {
	Cause            *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime          *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EssScalingRuleId *string `json:"EssScalingRuleId,omitempty" xml:"EssScalingRuleId,omitempty"`
	ExpectNum        *int32  `json:"ExpectNum,omitempty" xml:"ExpectNum,omitempty"`
	HostGroupName    *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	Id               *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InstanceIds      *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	ScalingGroupId   *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
	ScalingRuleName  *string `json:"ScalingRuleName,omitempty" xml:"ScalingRuleName,omitempty"`
	StartTime        *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalCapacity    *int32  `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	Transition       *string `json:"Transition,omitempty" xml:"Transition,omitempty"`
}

func (s ScalingActivity) String() string {
	return dara.Prettify(s)
}

func (s ScalingActivity) GoString() string {
	return s.String()
}

func (s *ScalingActivity) GetCause() *string {
	return s.Cause
}

func (s *ScalingActivity) GetDescription() *string {
	return s.Description
}

func (s *ScalingActivity) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ScalingActivity) GetEssScalingRuleId() *string {
	return s.EssScalingRuleId
}

func (s *ScalingActivity) GetExpectNum() *int32 {
	return s.ExpectNum
}

func (s *ScalingActivity) GetHostGroupName() *string {
	return s.HostGroupName
}

func (s *ScalingActivity) GetId() *string {
	return s.Id
}

func (s *ScalingActivity) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ScalingActivity) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *ScalingActivity) GetScalingRuleName() *string {
	return s.ScalingRuleName
}

func (s *ScalingActivity) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ScalingActivity) GetStatus() *string {
	return s.Status
}

func (s *ScalingActivity) GetTotalCapacity() *int32 {
	return s.TotalCapacity
}

func (s *ScalingActivity) GetTransition() *string {
	return s.Transition
}

func (s *ScalingActivity) SetCause(v string) *ScalingActivity {
	s.Cause = &v
	return s
}

func (s *ScalingActivity) SetDescription(v string) *ScalingActivity {
	s.Description = &v
	return s
}

func (s *ScalingActivity) SetEndTime(v int64) *ScalingActivity {
	s.EndTime = &v
	return s
}

func (s *ScalingActivity) SetEssScalingRuleId(v string) *ScalingActivity {
	s.EssScalingRuleId = &v
	return s
}

func (s *ScalingActivity) SetExpectNum(v int32) *ScalingActivity {
	s.ExpectNum = &v
	return s
}

func (s *ScalingActivity) SetHostGroupName(v string) *ScalingActivity {
	s.HostGroupName = &v
	return s
}

func (s *ScalingActivity) SetId(v string) *ScalingActivity {
	s.Id = &v
	return s
}

func (s *ScalingActivity) SetInstanceIds(v string) *ScalingActivity {
	s.InstanceIds = &v
	return s
}

func (s *ScalingActivity) SetScalingGroupId(v string) *ScalingActivity {
	s.ScalingGroupId = &v
	return s
}

func (s *ScalingActivity) SetScalingRuleName(v string) *ScalingActivity {
	s.ScalingRuleName = &v
	return s
}

func (s *ScalingActivity) SetStartTime(v int64) *ScalingActivity {
	s.StartTime = &v
	return s
}

func (s *ScalingActivity) SetStatus(v string) *ScalingActivity {
	s.Status = &v
	return s
}

func (s *ScalingActivity) SetTotalCapacity(v int32) *ScalingActivity {
	s.TotalCapacity = &v
	return s
}

func (s *ScalingActivity) SetTransition(v string) *ScalingActivity {
	s.Transition = &v
	return s
}

func (s *ScalingActivity) Validate() error {
	return dara.Validate(s)
}
