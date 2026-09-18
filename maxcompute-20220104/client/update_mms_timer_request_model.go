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
	SetTableBlackList(v []*string) *UpdateMmsTimerRequest
	GetTableBlackList() []*string
	SetTableWhiteList(v []*string) *UpdateMmsTimerRequest
	GetTableWhiteList() []*string
	SetValue(v string) *UpdateMmsTimerRequest
	GetValue() *string
}

type UpdateMmsTimerRequest struct {
	// The scheduling type of the scheduled task.
	//
	// example:
	//
	// Daily
	ScheduleType *string `json:"scheduleType,omitempty" xml:"scheduleType,omitempty"`
	// Indicates whether the scheduled task is stopped.
	Stopped *bool `json:"stopped,omitempty" xml:"stopped,omitempty"`
	// The tables to exclude when type is set to Database.
	TableBlackList []*string `json:"tableBlackList,omitempty" xml:"tableBlackList,omitempty" type:"Repeated"`
	// The tables to migrate when type is set to Database.
	TableWhiteList []*string `json:"tableWhiteList,omitempty" xml:"tableWhiteList,omitempty" type:"Repeated"`
	// The scheduling time of the scheduled task. If scheduleType is set to Daily, the value is in the HH:MM format. If scheduleType is set to Hourly, the value is in the MM format.
	//
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

func (s *UpdateMmsTimerRequest) GetTableBlackList() []*string {
	return s.TableBlackList
}

func (s *UpdateMmsTimerRequest) GetTableWhiteList() []*string {
	return s.TableWhiteList
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

func (s *UpdateMmsTimerRequest) SetTableBlackList(v []*string) *UpdateMmsTimerRequest {
	s.TableBlackList = v
	return s
}

func (s *UpdateMmsTimerRequest) SetTableWhiteList(v []*string) *UpdateMmsTimerRequest {
	s.TableWhiteList = v
	return s
}

func (s *UpdateMmsTimerRequest) SetValue(v string) *UpdateMmsTimerRequest {
	s.Value = &v
	return s
}

func (s *UpdateMmsTimerRequest) Validate() error {
	return dara.Validate(s)
}
