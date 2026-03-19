// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduleExecuteTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *ModifyScheduleExecuteTimeRequest
	GetEventId() *string
	SetScheduleExecuteTime(v string) *ModifyScheduleExecuteTimeRequest
	GetScheduleExecuteTime() *string
}

type ModifyScheduleExecuteTimeRequest struct {
	// example:
	//
	// eb7efbc90864a0***
	EventId *string `json:"eventId,omitempty" xml:"eventId,omitempty"`
	// example:
	//
	// 2026-01-30T08:45:00Z
	ScheduleExecuteTime *string `json:"scheduleExecuteTime,omitempty" xml:"scheduleExecuteTime,omitempty"`
}

func (s ModifyScheduleExecuteTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduleExecuteTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyScheduleExecuteTimeRequest) GetEventId() *string {
	return s.EventId
}

func (s *ModifyScheduleExecuteTimeRequest) GetScheduleExecuteTime() *string {
	return s.ScheduleExecuteTime
}

func (s *ModifyScheduleExecuteTimeRequest) SetEventId(v string) *ModifyScheduleExecuteTimeRequest {
	s.EventId = &v
	return s
}

func (s *ModifyScheduleExecuteTimeRequest) SetScheduleExecuteTime(v string) *ModifyScheduleExecuteTimeRequest {
	s.ScheduleExecuteTime = &v
	return s
}

func (s *ModifyScheduleExecuteTimeRequest) Validate() error {
	return dara.Validate(s)
}
