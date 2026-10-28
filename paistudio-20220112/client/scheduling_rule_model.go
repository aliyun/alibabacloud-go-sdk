// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSchedulingRule interface {
	dara.Model
	String() string
	GoString() string
	SetCronTab(v string) *SchedulingRule
	GetCronTab() *string
	SetEndAt(v string) *SchedulingRule
	GetEndAt() *string
	SetExecuteOnce(v bool) *SchedulingRule
	GetExecuteOnce() *bool
	SetStartAt(v string) *SchedulingRule
	GetStartAt() *string
}

type SchedulingRule struct {
	// The cron tab. Only the minute, hour, day, month, and week fields are currently supported. Second-level configurations are discarded.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 0 0 	- 	- 7
	CronTab *string `json:"CronTab,omitempty" xml:"CronTab,omitempty"`
	// The stop time of the task.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2025-06-22T00:00:00Z
	EndAt *string `json:"EndAt,omitempty" xml:"EndAt,omitempty"`
	// Specifies whether the task is executed only once.
	//
	// if can be null:
	// true
	ExecuteOnce *bool `json:"ExecuteOnce,omitempty" xml:"ExecuteOnce,omitempty"`
	// The start time of the effective period for the node.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2023-06-22T00:00:00Z
	StartAt *string `json:"StartAt,omitempty" xml:"StartAt,omitempty"`
}

func (s SchedulingRule) String() string {
	return dara.Prettify(s)
}

func (s SchedulingRule) GoString() string {
	return s.String()
}

func (s *SchedulingRule) GetCronTab() *string {
	return s.CronTab
}

func (s *SchedulingRule) GetEndAt() *string {
	return s.EndAt
}

func (s *SchedulingRule) GetExecuteOnce() *bool {
	return s.ExecuteOnce
}

func (s *SchedulingRule) GetStartAt() *string {
	return s.StartAt
}

func (s *SchedulingRule) SetCronTab(v string) *SchedulingRule {
	s.CronTab = &v
	return s
}

func (s *SchedulingRule) SetEndAt(v string) *SchedulingRule {
	s.EndAt = &v
	return s
}

func (s *SchedulingRule) SetExecuteOnce(v bool) *SchedulingRule {
	s.ExecuteOnce = &v
	return s
}

func (s *SchedulingRule) SetStartAt(v string) *SchedulingRule {
	s.StartAt = &v
	return s
}

func (s *SchedulingRule) Validate() error {
	return dara.Validate(s)
}
