// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWarmPoolScheduledPolicyInput interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *WarmPoolScheduledPolicyInput
	GetEndTime() *string
	SetName(v string) *WarmPoolScheduledPolicyInput
	GetName() *string
	SetScheduleExpression(v string) *WarmPoolScheduledPolicyInput
	GetScheduleExpression() *string
	SetStartTime(v string) *WarmPoolScheduledPolicyInput
	GetStartTime() *string
	SetTarget(v int64) *WarmPoolScheduledPolicyInput
	GetTarget() *int64
	SetTimeZone(v string) *WarmPoolScheduledPolicyInput
	GetTimeZone() *string
}

type WarmPoolScheduledPolicyInput struct {
	// The end time when the rule expires, in RFC 3339 format.
	//
	// example:
	//
	// 2026-12-31T23:59:59Z
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The name of the scheduled rule. It must be unique within the same warm pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// daily-scale
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The cron expression for the scheduled rule, in the cron (second minute hour day month week) format.
	//
	// This parameter is required.
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
	// This parameter is required.
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

func (s WarmPoolScheduledPolicyInput) String() string {
	return dara.Prettify(s)
}

func (s WarmPoolScheduledPolicyInput) GoString() string {
	return s.String()
}

func (s *WarmPoolScheduledPolicyInput) GetEndTime() *string {
	return s.EndTime
}

func (s *WarmPoolScheduledPolicyInput) GetName() *string {
	return s.Name
}

func (s *WarmPoolScheduledPolicyInput) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *WarmPoolScheduledPolicyInput) GetStartTime() *string {
	return s.StartTime
}

func (s *WarmPoolScheduledPolicyInput) GetTarget() *int64 {
	return s.Target
}

func (s *WarmPoolScheduledPolicyInput) GetTimeZone() *string {
	return s.TimeZone
}

func (s *WarmPoolScheduledPolicyInput) SetEndTime(v string) *WarmPoolScheduledPolicyInput {
	s.EndTime = &v
	return s
}

func (s *WarmPoolScheduledPolicyInput) SetName(v string) *WarmPoolScheduledPolicyInput {
	s.Name = &v
	return s
}

func (s *WarmPoolScheduledPolicyInput) SetScheduleExpression(v string) *WarmPoolScheduledPolicyInput {
	s.ScheduleExpression = &v
	return s
}

func (s *WarmPoolScheduledPolicyInput) SetStartTime(v string) *WarmPoolScheduledPolicyInput {
	s.StartTime = &v
	return s
}

func (s *WarmPoolScheduledPolicyInput) SetTarget(v int64) *WarmPoolScheduledPolicyInput {
	s.Target = &v
	return s
}

func (s *WarmPoolScheduledPolicyInput) SetTimeZone(v string) *WarmPoolScheduledPolicyInput {
	s.TimeZone = &v
	return s
}

func (s *WarmPoolScheduledPolicyInput) Validate() error {
	return dara.Validate(s)
}
