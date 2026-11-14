// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWarmPoolScheduledPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *WarmPoolScheduledPolicy
	GetEndTime() *string
	SetName(v string) *WarmPoolScheduledPolicy
	GetName() *string
	SetScheduleExpression(v string) *WarmPoolScheduledPolicy
	GetScheduleExpression() *string
	SetStartTime(v string) *WarmPoolScheduledPolicy
	GetStartTime() *string
	SetTarget(v int64) *WarmPoolScheduledPolicy
	GetTarget() *int64
	SetTimeZone(v string) *WarmPoolScheduledPolicy
	GetTimeZone() *string
}

type WarmPoolScheduledPolicy struct {
	// The end time when the rule expires, in RFC 3339 format.
	//
	// example:
	//
	// 2026-12-31T23:59:59Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The name of the scheduled rule. It must be unique within the same warm pool.
	//
	// example:
	//
	// daily-scale
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The cron expression for the scheduled rule, in the cron (second minute hour day month week) format.
	//
	// example:
	//
	// cron(0 0 0 	- 	- *)
	ScheduleExpression *string `json:"scheduleExpression,omitempty" xml:"scheduleExpression,omitempty"`
	// The start time when the rule takes effect, in RFC 3339 format.
	//
	// example:
	//
	// 2026-12-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The target number of ready sandboxes after the rule is triggered. The value can be 0.
	//
	// example:
	//
	// 0
	Target *int64 `json:"target,omitempty" xml:"target,omitempty"`
	// The IANA time zone used to execute the scheduled rule.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s WarmPoolScheduledPolicy) String() string {
	return dara.Prettify(s)
}

func (s WarmPoolScheduledPolicy) GoString() string {
	return s.String()
}

func (s *WarmPoolScheduledPolicy) GetEndTime() *string {
	return s.EndTime
}

func (s *WarmPoolScheduledPolicy) GetName() *string {
	return s.Name
}

func (s *WarmPoolScheduledPolicy) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *WarmPoolScheduledPolicy) GetStartTime() *string {
	return s.StartTime
}

func (s *WarmPoolScheduledPolicy) GetTarget() *int64 {
	return s.Target
}

func (s *WarmPoolScheduledPolicy) GetTimeZone() *string {
	return s.TimeZone
}

func (s *WarmPoolScheduledPolicy) SetEndTime(v string) *WarmPoolScheduledPolicy {
	s.EndTime = &v
	return s
}

func (s *WarmPoolScheduledPolicy) SetName(v string) *WarmPoolScheduledPolicy {
	s.Name = &v
	return s
}

func (s *WarmPoolScheduledPolicy) SetScheduleExpression(v string) *WarmPoolScheduledPolicy {
	s.ScheduleExpression = &v
	return s
}

func (s *WarmPoolScheduledPolicy) SetStartTime(v string) *WarmPoolScheduledPolicy {
	s.StartTime = &v
	return s
}

func (s *WarmPoolScheduledPolicy) SetTarget(v int64) *WarmPoolScheduledPolicy {
	s.Target = &v
	return s
}

func (s *WarmPoolScheduledPolicy) SetTimeZone(v string) *WarmPoolScheduledPolicy {
	s.TimeZone = &v
	return s
}

func (s *WarmPoolScheduledPolicy) Validate() error {
	return dara.Validate(s)
}
