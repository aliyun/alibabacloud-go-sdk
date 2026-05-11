// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduleRule interface {
	dara.Model
	String() string
	GoString() string
	SetMaxScheduleCount(v int64) *ScheduleRule
	GetMaxScheduleCount() *int64
	SetStartCronExpression(v string) *ScheduleRule
	GetStartCronExpression() *string
	SetSuspendCronExpression(v string) *ScheduleRule
	GetSuspendCronExpression() *string
}

type ScheduleRule struct {
	// The maximum number of times that the migration task is automatically scheduled. Each time the migration task is run, the execution ID increases by one until the task is run the specified number of times. The task is automatically scheduled based on the specified start time and pause time. The task is no longer automatically scheduled after the task is run the specified number of times. However, you can still manually start the task.
	//
	// example:
	//
	// 1
	MaxScheduleCount *int64 `json:"MaxScheduleCount,omitempty" xml:"MaxScheduleCount,omitempty"`
	// The time when the migration task started. You can use a CRON expression to specify the time. The interval at which the migration task is run is at least 1 hour.
	//
	// example:
	//
	// 0 0 	- 	- 	- ?
	StartCronExpression *string `json:"StartCronExpression,omitempty" xml:"StartCronExpression,omitempty"`
	// The time when the migration task paused. You can use a CRON expression to specify the time. The interval at which the migration task is run is at least 1 hour.
	//
	// example:
	//
	// 0 0 	- 	- 	- ?
	SuspendCronExpression *string `json:"SuspendCronExpression,omitempty" xml:"SuspendCronExpression,omitempty"`
}

func (s ScheduleRule) String() string {
	return dara.Prettify(s)
}

func (s ScheduleRule) GoString() string {
	return s.String()
}

func (s *ScheduleRule) GetMaxScheduleCount() *int64 {
	return s.MaxScheduleCount
}

func (s *ScheduleRule) GetStartCronExpression() *string {
	return s.StartCronExpression
}

func (s *ScheduleRule) GetSuspendCronExpression() *string {
	return s.SuspendCronExpression
}

func (s *ScheduleRule) SetMaxScheduleCount(v int64) *ScheduleRule {
	s.MaxScheduleCount = &v
	return s
}

func (s *ScheduleRule) SetStartCronExpression(v string) *ScheduleRule {
	s.StartCronExpression = &v
	return s
}

func (s *ScheduleRule) SetSuspendCronExpression(v string) *ScheduleRule {
	s.SuspendCronExpression = &v
	return s
}

func (s *ScheduleRule) Validate() error {
	return dara.Validate(s)
}
