// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParameterTimedScheduleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *ModifyParameterTimedScheduleTaskRequest
	GetDBInstanceName() *string
	SetSwitchTime(v string) *ModifyParameterTimedScheduleTaskRequest
	GetSwitchTime() *string
	SetTaskId(v int64) *ModifyParameterTimedScheduleTaskRequest
	GetTaskId() *int64
}

type ModifyParameterTimedScheduleTaskRequest struct {
	// The name of the instance.
	//
	// example:
	//
	// pgm-bp102g323jd4****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The scheduled switchover time to set. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2022-05-06T09:24:00Z
	SwitchTime *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// 440437220
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyParameterTimedScheduleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyParameterTimedScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyParameterTimedScheduleTaskRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *ModifyParameterTimedScheduleTaskRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyParameterTimedScheduleTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ModifyParameterTimedScheduleTaskRequest) SetDBInstanceName(v string) *ModifyParameterTimedScheduleTaskRequest {
	s.DBInstanceName = &v
	return s
}

func (s *ModifyParameterTimedScheduleTaskRequest) SetSwitchTime(v string) *ModifyParameterTimedScheduleTaskRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyParameterTimedScheduleTaskRequest) SetTaskId(v int64) *ModifyParameterTimedScheduleTaskRequest {
	s.TaskId = &v
	return s
}

func (s *ModifyParameterTimedScheduleTaskRequest) Validate() error {
	return dara.Validate(s)
}
