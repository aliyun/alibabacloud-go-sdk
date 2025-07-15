// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledScalingRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDurationMinutes(v int32) *CreateScheduledScalingRuleShrinkRequest
	GetDurationMinutes() *int32
	SetEnable(v bool) *CreateScheduledScalingRuleShrinkRequest
	GetEnable() *bool
	SetFirstScheduledTime(v int64) *CreateScheduledScalingRuleShrinkRequest
	GetFirstScheduledTime() *int64
	SetInstanceId(v string) *CreateScheduledScalingRuleShrinkRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateScheduledScalingRuleShrinkRequest
	GetRegionId() *string
	SetRepeatType(v string) *CreateScheduledScalingRuleShrinkRequest
	GetRepeatType() *string
	SetReservedPubFlow(v int32) *CreateScheduledScalingRuleShrinkRequest
	GetReservedPubFlow() *int32
	SetReservedSubFlow(v int32) *CreateScheduledScalingRuleShrinkRequest
	GetReservedSubFlow() *int32
	SetRuleName(v string) *CreateScheduledScalingRuleShrinkRequest
	GetRuleName() *string
	SetScheduleType(v string) *CreateScheduledScalingRuleShrinkRequest
	GetScheduleType() *string
	SetTimeZone(v string) *CreateScheduledScalingRuleShrinkRequest
	GetTimeZone() *string
	SetWeeklyTypesShrink(v string) *CreateScheduledScalingRuleShrinkRequest
	GetWeeklyTypesShrink() *string
}

type CreateScheduledScalingRuleShrinkRequest struct {
	// The duration of each scheduled scaling task. Unit: minutes.
	//
	// >  The value of this parameter must be greater than or equal to 15.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	DurationMinutes *int32 `json:"DurationMinutes,omitempty" xml:"DurationMinutes,omitempty"`
	// Specifies whether to enable the scheduled scaling rule. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The time when the scheduled scaling task is executed.
	//
	// If you set ScheduleType to at, make sure that the value of this parameter is at least 30 minutes later than the current point in time.
	//
	// 	Notice: To prevent the broker from repeatedly executing instance upgrade and downgrade tasks, make sure that the interval between two consecutive scheduled scaling tasks is at least 60 minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1714467540000
	FirstScheduledTime *int64 `json:"FirstScheduledTime,omitempty" xml:"FirstScheduledTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_serverless-cn-vxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The frequency to execute the scheduled scaling task. This parameter is required only if you set ScheduleType to repeat. Valid values:
	//
	// 	- Daily: The scheduled scaling task is executed every day.
	//
	// 	- Weekly: The scheduled scaling task is executed every week.
	//
	// example:
	//
	// Weekly
	RepeatType *string `json:"RepeatType,omitempty" xml:"RepeatType,omitempty"`
	// The reserved production capacity for scheduled scaling. Unit: MB/s.
	//
	// >  You must specify a higher value than the instance specification for at least one of ReservedPubFlow and ReservedSubFlow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120
	ReservedPubFlow *int32 `json:"ReservedPubFlow,omitempty" xml:"ReservedPubFlow,omitempty"`
	// The reserved consumption capacity for scheduled scaling. Unit: MB/s.
	//
	// >  You must specify a higher value than the instance specification for at least one of ReservedPubFlow and ReservedSubFlow.
	//
	// This parameter is required.
	//
	// example:
	//
	// 120
	ReservedSubFlow *int32 `json:"ReservedSubFlow,omitempty" xml:"ReservedSubFlow,omitempty"`
	// The name of the scheduled scaling rule.
	//
	// >  The name of the scheduled scaling rule cannot be the same as the names of other rules for the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The type of the scheduled scaling task. Valid values:
	//
	// 	- at: The scheduled scaling task is executed only once.
	//
	// 	- repeat: The scheduled scaling task is repeatedly executed.
	//
	// This parameter is required.
	//
	// example:
	//
	// at
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The time zone in Coordinated Universal Time (UTC).
	//
	// This parameter is required.
	//
	// example:
	//
	// GMT+8
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	// The day on which the scheduled scaling task is executed every week. You can specify multiple days.
	WeeklyTypesShrink *string `json:"WeeklyTypes,omitempty" xml:"WeeklyTypes,omitempty"`
}

func (s CreateScheduledScalingRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledScalingRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledScalingRuleShrinkRequest) GetDurationMinutes() *int32 {
	return s.DurationMinutes
}

func (s *CreateScheduledScalingRuleShrinkRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateScheduledScalingRuleShrinkRequest) GetFirstScheduledTime() *int64 {
	return s.FirstScheduledTime
}

func (s *CreateScheduledScalingRuleShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateScheduledScalingRuleShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateScheduledScalingRuleShrinkRequest) GetRepeatType() *string {
	return s.RepeatType
}

func (s *CreateScheduledScalingRuleShrinkRequest) GetReservedPubFlow() *int32 {
	return s.ReservedPubFlow
}

func (s *CreateScheduledScalingRuleShrinkRequest) GetReservedSubFlow() *int32 {
	return s.ReservedSubFlow
}

func (s *CreateScheduledScalingRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateScheduledScalingRuleShrinkRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *CreateScheduledScalingRuleShrinkRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateScheduledScalingRuleShrinkRequest) GetWeeklyTypesShrink() *string {
	return s.WeeklyTypesShrink
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetDurationMinutes(v int32) *CreateScheduledScalingRuleShrinkRequest {
	s.DurationMinutes = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetEnable(v bool) *CreateScheduledScalingRuleShrinkRequest {
	s.Enable = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetFirstScheduledTime(v int64) *CreateScheduledScalingRuleShrinkRequest {
	s.FirstScheduledTime = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetInstanceId(v string) *CreateScheduledScalingRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetRegionId(v string) *CreateScheduledScalingRuleShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetRepeatType(v string) *CreateScheduledScalingRuleShrinkRequest {
	s.RepeatType = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetReservedPubFlow(v int32) *CreateScheduledScalingRuleShrinkRequest {
	s.ReservedPubFlow = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetReservedSubFlow(v int32) *CreateScheduledScalingRuleShrinkRequest {
	s.ReservedSubFlow = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetRuleName(v string) *CreateScheduledScalingRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetScheduleType(v string) *CreateScheduledScalingRuleShrinkRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetTimeZone(v string) *CreateScheduledScalingRuleShrinkRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) SetWeeklyTypesShrink(v string) *CreateScheduledScalingRuleShrinkRequest {
	s.WeeklyTypesShrink = &v
	return s
}

func (s *CreateScheduledScalingRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
