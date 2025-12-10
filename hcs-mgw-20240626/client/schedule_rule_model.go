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
	MaxScheduleCount      *int64  `json:"MaxScheduleCount,omitempty" xml:"MaxScheduleCount,omitempty"`
	StartCronExpression   *string `json:"StartCronExpression,omitempty" xml:"StartCronExpression,omitempty"`
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
