// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduledAction interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ScheduledAction
	GetEndTime() *string
	SetName(v string) *ScheduledAction
	GetName() *string
	SetScheduleExpression(v string) *ScheduledAction
	GetScheduleExpression() *string
	SetStartTime(v string) *ScheduledAction
	GetStartTime() *string
	SetTarget(v int64) *ScheduledAction
	GetTarget() *int64
	SetTimeZone(v string) *ScheduledAction
	GetTimeZone() *string
}

type ScheduledAction struct {
	// example:
	//
	// 2024-03-10T10:10:10
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cron(0 0 22 	- 	- *)
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	// example:
	//
	// 2023-03-10T10:10:10
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 50
	Target *int64 `json:"target,omitempty" xml:"target,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s ScheduledAction) String() string {
	return dara.Prettify(s)
}

func (s ScheduledAction) GoString() string {
	return s.String()
}

func (s *ScheduledAction) GetEndTime() *string {
	return s.EndTime
}

func (s *ScheduledAction) GetName() *string {
	return s.Name
}

func (s *ScheduledAction) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *ScheduledAction) GetStartTime() *string {
	return s.StartTime
}

func (s *ScheduledAction) GetTarget() *int64 {
	return s.Target
}

func (s *ScheduledAction) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ScheduledAction) SetEndTime(v string) *ScheduledAction {
	s.EndTime = &v
	return s
}

func (s *ScheduledAction) SetName(v string) *ScheduledAction {
	s.Name = &v
	return s
}

func (s *ScheduledAction) SetScheduleExpression(v string) *ScheduledAction {
	s.ScheduleExpression = &v
	return s
}

func (s *ScheduledAction) SetStartTime(v string) *ScheduledAction {
	s.StartTime = &v
	return s
}

func (s *ScheduledAction) SetTarget(v int64) *ScheduledAction {
	s.Target = &v
	return s
}

func (s *ScheduledAction) SetTimeZone(v string) *ScheduledAction {
	s.TimeZone = &v
	return s
}

func (s *ScheduledAction) Validate() error {
	return dara.Validate(s)
}
