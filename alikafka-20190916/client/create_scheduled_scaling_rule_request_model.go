// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDurationMinutes(v int32) *CreateScheduledScalingRuleRequest
	GetDurationMinutes() *int32
	SetEnable(v bool) *CreateScheduledScalingRuleRequest
	GetEnable() *bool
	SetFirstScheduledTime(v int64) *CreateScheduledScalingRuleRequest
	GetFirstScheduledTime() *int64
	SetInstanceId(v string) *CreateScheduledScalingRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *CreateScheduledScalingRuleRequest
	GetRegionId() *string
	SetRepeatType(v string) *CreateScheduledScalingRuleRequest
	GetRepeatType() *string
	SetReservedPubFlow(v int32) *CreateScheduledScalingRuleRequest
	GetReservedPubFlow() *int32
	SetReservedSubFlow(v int32) *CreateScheduledScalingRuleRequest
	GetReservedSubFlow() *int32
	SetRuleName(v string) *CreateScheduledScalingRuleRequest
	GetRuleName() *string
	SetScheduleType(v string) *CreateScheduledScalingRuleRequest
	GetScheduleType() *string
	SetTimeZone(v string) *CreateScheduledScalingRuleRequest
	GetTimeZone() *string
	SetWeeklyTypes(v []*string) *CreateScheduledScalingRuleRequest
	GetWeeklyTypes() []*string
}

type CreateScheduledScalingRuleRequest struct {
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
	WeeklyTypes []*string `json:"WeeklyTypes,omitempty" xml:"WeeklyTypes,omitempty" type:"Repeated"`
}

func (s CreateScheduledScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledScalingRuleRequest) GetDurationMinutes() *int32 {
	return s.DurationMinutes
}

func (s *CreateScheduledScalingRuleRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateScheduledScalingRuleRequest) GetFirstScheduledTime() *int64 {
	return s.FirstScheduledTime
}

func (s *CreateScheduledScalingRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateScheduledScalingRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateScheduledScalingRuleRequest) GetRepeatType() *string {
	return s.RepeatType
}

func (s *CreateScheduledScalingRuleRequest) GetReservedPubFlow() *int32 {
	return s.ReservedPubFlow
}

func (s *CreateScheduledScalingRuleRequest) GetReservedSubFlow() *int32 {
	return s.ReservedSubFlow
}

func (s *CreateScheduledScalingRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateScheduledScalingRuleRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *CreateScheduledScalingRuleRequest) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateScheduledScalingRuleRequest) GetWeeklyTypes() []*string {
	return s.WeeklyTypes
}

func (s *CreateScheduledScalingRuleRequest) SetDurationMinutes(v int32) *CreateScheduledScalingRuleRequest {
	s.DurationMinutes = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetEnable(v bool) *CreateScheduledScalingRuleRequest {
	s.Enable = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetFirstScheduledTime(v int64) *CreateScheduledScalingRuleRequest {
	s.FirstScheduledTime = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetInstanceId(v string) *CreateScheduledScalingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetRegionId(v string) *CreateScheduledScalingRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetRepeatType(v string) *CreateScheduledScalingRuleRequest {
	s.RepeatType = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetReservedPubFlow(v int32) *CreateScheduledScalingRuleRequest {
	s.ReservedPubFlow = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetReservedSubFlow(v int32) *CreateScheduledScalingRuleRequest {
	s.ReservedSubFlow = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetRuleName(v string) *CreateScheduledScalingRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetScheduleType(v string) *CreateScheduledScalingRuleRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetTimeZone(v string) *CreateScheduledScalingRuleRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateScheduledScalingRuleRequest) SetWeeklyTypes(v []*string) *CreateScheduledScalingRuleRequest {
	s.WeeklyTypes = v
	return s
}

func (s *CreateScheduledScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
