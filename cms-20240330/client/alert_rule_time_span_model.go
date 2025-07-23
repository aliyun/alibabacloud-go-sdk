// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAlertRuleTimeSpan interface {
	dara.Model
	String() string
	GoString() string
	SetDayOfWeek(v []*int32) *AlertRuleTimeSpan
	GetDayOfWeek() []*int32
	SetEndTime(v string) *AlertRuleTimeSpan
	GetEndTime() *string
	SetGmtOffset(v string) *AlertRuleTimeSpan
	GetGmtOffset() *string
	SetStartTime(v string) *AlertRuleTimeSpan
	GetStartTime() *string
}

type AlertRuleTimeSpan struct {
	DayOfWeek []*int32 `json:"dayOfWeek,omitempty" xml:"dayOfWeek,omitempty" type:"Repeated"`
	EndTime   *string  `json:"endTime,omitempty" xml:"endTime,omitempty"`
	GmtOffset *string  `json:"gmtOffset,omitempty" xml:"gmtOffset,omitempty"`
	StartTime *string  `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s AlertRuleTimeSpan) String() string {
	return dara.Prettify(s)
}

func (s AlertRuleTimeSpan) GoString() string {
	return s.String()
}

func (s *AlertRuleTimeSpan) GetDayOfWeek() []*int32 {
	return s.DayOfWeek
}

func (s *AlertRuleTimeSpan) GetEndTime() *string {
	return s.EndTime
}

func (s *AlertRuleTimeSpan) GetGmtOffset() *string {
	return s.GmtOffset
}

func (s *AlertRuleTimeSpan) GetStartTime() *string {
	return s.StartTime
}

func (s *AlertRuleTimeSpan) SetDayOfWeek(v []*int32) *AlertRuleTimeSpan {
	s.DayOfWeek = v
	return s
}

func (s *AlertRuleTimeSpan) SetEndTime(v string) *AlertRuleTimeSpan {
	s.EndTime = &v
	return s
}

func (s *AlertRuleTimeSpan) SetGmtOffset(v string) *AlertRuleTimeSpan {
	s.GmtOffset = &v
	return s
}

func (s *AlertRuleTimeSpan) SetStartTime(v string) *AlertRuleTimeSpan {
	s.StartTime = &v
	return s
}

func (s *AlertRuleTimeSpan) Validate() error {
	return dara.Validate(s)
}
