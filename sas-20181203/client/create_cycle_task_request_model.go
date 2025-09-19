// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCycleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v int32) *CreateCycleTaskRequest
	GetEnable() *int32
	SetFirstDateStr(v int64) *CreateCycleTaskRequest
	GetFirstDateStr() *int64
	SetIntervalPeriod(v int32) *CreateCycleTaskRequest
	GetIntervalPeriod() *int32
	SetParam(v string) *CreateCycleTaskRequest
	GetParam() *string
	SetPeriodUnit(v string) *CreateCycleTaskRequest
	GetPeriodUnit() *string
	SetSource(v string) *CreateCycleTaskRequest
	GetSource() *string
	SetTargetEndTime(v int32) *CreateCycleTaskRequest
	GetTargetEndTime() *int32
	SetTargetStartTime(v int32) *CreateCycleTaskRequest
	GetTargetStartTime() *int32
	SetTaskName(v string) *CreateCycleTaskRequest
	GetTaskName() *string
	SetTaskType(v string) *CreateCycleTaskRequest
	GetTaskType() *string
}

type CreateCycleTaskRequest struct {
	// Specifies whether to enable the task. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The first time when the task is performed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1650556800000
	FirstDateStr *int64 `json:"FirstDateStr,omitempty" xml:"FirstDateStr,omitempty"`
	// The interval of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	IntervalPeriod *int32 `json:"IntervalPeriod,omitempty" xml:"IntervalPeriod,omitempty"`
	// The additional information.
	//
	// example:
	//
	// {
	//
	//       "targetInfo": [
	//
	//             {
	//
	//                   "type": "groupId",
	//
	//                   "name": "TI HOST",
	//
	//                   "target": 10597***
	//
	//             },
	//
	//             {
	//
	//                   "type": "groupId",
	//
	//                   "name": "expense HOST",
	//
	//                   "target": 10597***
	//
	//             }
	//
	//       ]
	//
	// }
	Param *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The unit of the scan interval. Valid values:
	//
	// 	- **day**: days
	//
	// 	- **hour**: hours
	//
	// This parameter is required.
	//
	// example:
	//
	// day
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The additional source for the task.
	//
	// example:
	//
	// console_batch
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The time when the task ends. Unit: hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	TargetEndTime *int32 `json:"TargetEndTime,omitempty" xml:"TargetEndTime,omitempty"`
	// The time when the task is started. Unit: hours.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	TargetStartTime *int32 `json:"TargetStartTime,omitempty" xml:"TargetStartTime,omitempty"`
	// The name of the task. Valid values:
	//
	// 	- **VIRUS_VUL_SCHEDULE_SCAN**: virus scan task
	//
	// 	- **IMAGE_SCAN**: image scan task
	//
	// 	- **EMG_VUL_SCHEDULE_SCAN**: urgent vulnerability scan task
	//
	// This parameter is required.
	//
	// example:
	//
	// EMG_VUL_SCHEDULE_SCAN
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **VIRUS_VUL_SCHEDULE_SCAN**: virus scan task
	//
	// 	- **IMAGE_SCAN**: image scan task
	//
	// 	- **EMG_VUL_SCHEDULE_SCAN**: urgent vulnerability scan task
	//
	// This parameter is required.
	//
	// example:
	//
	// VIRUS_VUL_SCHEDULE_SCAN
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s CreateCycleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCycleTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateCycleTaskRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *CreateCycleTaskRequest) GetFirstDateStr() *int64 {
	return s.FirstDateStr
}

func (s *CreateCycleTaskRequest) GetIntervalPeriod() *int32 {
	return s.IntervalPeriod
}

func (s *CreateCycleTaskRequest) GetParam() *string {
	return s.Param
}

func (s *CreateCycleTaskRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateCycleTaskRequest) GetSource() *string {
	return s.Source
}

func (s *CreateCycleTaskRequest) GetTargetEndTime() *int32 {
	return s.TargetEndTime
}

func (s *CreateCycleTaskRequest) GetTargetStartTime() *int32 {
	return s.TargetStartTime
}

func (s *CreateCycleTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateCycleTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *CreateCycleTaskRequest) SetEnable(v int32) *CreateCycleTaskRequest {
	s.Enable = &v
	return s
}

func (s *CreateCycleTaskRequest) SetFirstDateStr(v int64) *CreateCycleTaskRequest {
	s.FirstDateStr = &v
	return s
}

func (s *CreateCycleTaskRequest) SetIntervalPeriod(v int32) *CreateCycleTaskRequest {
	s.IntervalPeriod = &v
	return s
}

func (s *CreateCycleTaskRequest) SetParam(v string) *CreateCycleTaskRequest {
	s.Param = &v
	return s
}

func (s *CreateCycleTaskRequest) SetPeriodUnit(v string) *CreateCycleTaskRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateCycleTaskRequest) SetSource(v string) *CreateCycleTaskRequest {
	s.Source = &v
	return s
}

func (s *CreateCycleTaskRequest) SetTargetEndTime(v int32) *CreateCycleTaskRequest {
	s.TargetEndTime = &v
	return s
}

func (s *CreateCycleTaskRequest) SetTargetStartTime(v int32) *CreateCycleTaskRequest {
	s.TargetStartTime = &v
	return s
}

func (s *CreateCycleTaskRequest) SetTaskName(v string) *CreateCycleTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateCycleTaskRequest) SetTaskType(v string) *CreateCycleTaskRequest {
	s.TaskType = &v
	return s
}

func (s *CreateCycleTaskRequest) Validate() error {
	return dara.Validate(s)
}
