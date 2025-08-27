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
	EndTime            *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	Name               *string `json:"name,omitempty" xml:"name,omitempty"`
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	StartTime          *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	Target             *int64  `json:"target,omitempty" xml:"target,omitempty"`
	TimeZone           *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
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
