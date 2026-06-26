// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduledPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ScheduledPolicy
	GetEndTime() *string
	SetName(v string) *ScheduledPolicy
	GetName() *string
	SetScheduleExpression(v string) *ScheduledPolicy
	GetScheduleExpression() *string
	SetStartTime(v string) *ScheduledPolicy
	GetStartTime() *string
	SetTarget(v int64) *ScheduledPolicy
	GetTarget() *int64
	SetTimeZone(v string) *ScheduledPolicy
	GetTimeZone() *string
}

type ScheduledPolicy struct {
	// The end time.
	//
	// example:
	//
	// 1633449590000
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The policy name.
	//
	// example:
	//
	// student_app_shop_analyzer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The schedule configuration.
	//
	// example:
	//
	// 。
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1764432000000
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The current number of target resources. If a metric-based auto scaling policy or a scheduled policy is in effect, this parameter specifies the number of resources calculated by the policy. Otherwise, this parameter specifies the default number of provisioned instances.
	//
	// > How is this different from defaultTarget?<br>
	//
	// > Assume that you set the number of provisioned instances to 1 and then add a scheduled auto scaling policy to set the number to 5 for a specific time period.<br>
	//
	// >
	//
	// > - When the scheduled policy is active, target is 5 and defaultTarget is 1.
	//
	// >
	//
	// > - When the scheduled policy is inactive, both target and defaultTarget are 1.
	//
	// example:
	//
	// 5
	Target *int64 `json:"target,omitempty" xml:"target,omitempty"`
	// The time zone. If this parameter is left empty, the times for startTime, endTime, and scheduleExpression must be in UTC format.
	//
	// example:
	//
	// 。
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s ScheduledPolicy) String() string {
	return dara.Prettify(s)
}

func (s ScheduledPolicy) GoString() string {
	return s.String()
}

func (s *ScheduledPolicy) GetEndTime() *string {
	return s.EndTime
}

func (s *ScheduledPolicy) GetName() *string {
	return s.Name
}

func (s *ScheduledPolicy) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *ScheduledPolicy) GetStartTime() *string {
	return s.StartTime
}

func (s *ScheduledPolicy) GetTarget() *int64 {
	return s.Target
}

func (s *ScheduledPolicy) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ScheduledPolicy) SetEndTime(v string) *ScheduledPolicy {
	s.EndTime = &v
	return s
}

func (s *ScheduledPolicy) SetName(v string) *ScheduledPolicy {
	s.Name = &v
	return s
}

func (s *ScheduledPolicy) SetScheduleExpression(v string) *ScheduledPolicy {
	s.ScheduleExpression = &v
	return s
}

func (s *ScheduledPolicy) SetStartTime(v string) *ScheduledPolicy {
	s.StartTime = &v
	return s
}

func (s *ScheduledPolicy) SetTarget(v int64) *ScheduledPolicy {
	s.Target = &v
	return s
}

func (s *ScheduledPolicy) SetTimeZone(v string) *ScheduledPolicy {
	s.TimeZone = &v
	return s
}

func (s *ScheduledPolicy) Validate() error {
	return dara.Validate(s)
}
