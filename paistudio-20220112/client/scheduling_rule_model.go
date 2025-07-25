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
	// if can be null:
	// true
	CronTab *string `json:"CronTab,omitempty" xml:"CronTab,omitempty"`
	// if can be null:
	// true
	EndAt *string `json:"EndAt,omitempty" xml:"EndAt,omitempty"`
	// if can be null:
	// true
	ExecuteOnce *bool `json:"ExecuteOnce,omitempty" xml:"ExecuteOnce,omitempty"`
	// if can be null:
	// true
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
