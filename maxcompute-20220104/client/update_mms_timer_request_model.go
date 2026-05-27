// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmsTimerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScheduleType(v string) *UpdateMmsTimerRequest
	GetScheduleType() *string
	SetStopped(v bool) *UpdateMmsTimerRequest
	GetStopped() *bool
	SetValue(v string) *UpdateMmsTimerRequest
	GetValue() *string
}

type UpdateMmsTimerRequest struct {
	// example:
	//
	// Daily
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	Stopped      *bool   `json:"stopped,omitempty" xml:"stopped,omitempty"`
	// example:
	//
	// 2025-09-20
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateMmsTimerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsTimerRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmsTimerRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *UpdateMmsTimerRequest) GetStopped() *bool {
	return s.Stopped
}

func (s *UpdateMmsTimerRequest) GetValue() *string {
	return s.Value
}

func (s *UpdateMmsTimerRequest) SetScheduleType(v string) *UpdateMmsTimerRequest {
	s.ScheduleType = &v
	return s
}

func (s *UpdateMmsTimerRequest) SetStopped(v bool) *UpdateMmsTimerRequest {
	s.Stopped = &v
	return s
}

func (s *UpdateMmsTimerRequest) SetValue(v string) *UpdateMmsTimerRequest {
	s.Value = &v
	return s
}

func (s *UpdateMmsTimerRequest) Validate() error {
	return dara.Validate(s)
}
