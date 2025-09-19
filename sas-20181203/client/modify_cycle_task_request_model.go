// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCycleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v string) *ModifyCycleTaskRequest
	GetConfigId() *string
	SetEnable(v int32) *ModifyCycleTaskRequest
	GetEnable() *int32
	SetFirstDateStr(v int64) *ModifyCycleTaskRequest
	GetFirstDateStr() *int64
	SetIntervalPeriod(v int32) *ModifyCycleTaskRequest
	GetIntervalPeriod() *int32
	SetParam(v string) *ModifyCycleTaskRequest
	GetParam() *string
	SetPeriodUnit(v string) *ModifyCycleTaskRequest
	GetPeriodUnit() *string
	SetTargetEndTime(v int32) *ModifyCycleTaskRequest
	GetTargetEndTime() *int32
	SetTargetStartTime(v int32) *ModifyCycleTaskRequest
	GetTargetStartTime() *int32
	SetTaskName(v string) *ModifyCycleTaskRequest
	GetTaskName() *string
	SetTaskType(v string) *ModifyCycleTaskRequest
	GetTaskType() *string
}

type ModifyCycleTaskRequest struct {
	// The ID of the task configuration.
	//
	// >  You can call the [DescribeCycleTaskList](~~DescribeCycleTaskList~~) operation to query the IDs of task configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00cfa8161da093089e6804ba6a33****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Specifies whether to enable the task. Valid values:
	//
	// 	- **1**: enables the task.
	//
	// 	- **0**: disables the task.
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The time when the task first started.
	//
	// example:
	//
	// 1664380800000
	FirstDateStr *int64 `json:"FirstDateStr,omitempty" xml:"FirstDateStr,omitempty"`
	// The interval at which the task is run.
	//
	// example:
	//
	// 14
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
	// 	- **day**
	//
	// 	- **hour**
	//
	// example:
	//
	// day
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The end time of the task. The time must be a time frame.
	//
	// example:
	//
	// 6
	TargetEndTime *int32 `json:"TargetEndTime,omitempty" xml:"TargetEndTime,omitempty"`
	// The start time of the task. The start time must be a time frame.
	//
	// example:
	//
	// 0
	TargetStartTime *int32 `json:"TargetStartTime,omitempty" xml:"TargetStartTime,omitempty"`
	// The name of the task. Valid values:
	//
	// 	- **VIRUS_VUL_SCHEDULE_SCAN**: virus detection task
	//
	// 	- **IMAGE_SCAN**: image scan task
	//
	// 	- **EMG_VUL_SCHEDULE_SCAN**: urgent vulnerability scan task
	//
	// Valid values:
	//
	// 	- VIRUS_VUL_SCHEDULE_SCAN
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     virus detection task
	//
	//     <!-- -->
	//
	// 	- IMAGE_SCAN
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     image scan task
	//
	//     <!-- -->
	//
	// 	- EMG_VUL_SCHEDULE_SCAN
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     urgent vulnerability scan task
	//
	//     <!-- -->
	//
	// example:
	//
	// EMG_VUL_SCHEDULE_SCAN
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// The type of the task. Valid values:
	//
	// 	- **VIRUS_VUL_SCHEDULE_SCAN**: virus detection task
	//
	// 	- **IMAGE_SCAN**: image scan task
	//
	// 	- **EMG_VUL_SCHEDULE_SCAN**: urgent vulnerability scan task
	//
	// Valid values:
	//
	// 	- VIRUS_VUL_SCHEDULE_SCAN
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     virus detection task
	//
	//     <!-- -->
	//
	// 	- IMAGE_SCAN
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     image scan task
	//
	//     <!-- -->
	//
	// 	- EMG_VUL_SCHEDULE_SCAN
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     urgent vulnerability scan task
	//
	//     <!-- -->
	//
	// example:
	//
	// VIRUS_VUL_SCHEDULE_SCAN
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ModifyCycleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCycleTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyCycleTaskRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *ModifyCycleTaskRequest) GetEnable() *int32 {
	return s.Enable
}

func (s *ModifyCycleTaskRequest) GetFirstDateStr() *int64 {
	return s.FirstDateStr
}

func (s *ModifyCycleTaskRequest) GetIntervalPeriod() *int32 {
	return s.IntervalPeriod
}

func (s *ModifyCycleTaskRequest) GetParam() *string {
	return s.Param
}

func (s *ModifyCycleTaskRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyCycleTaskRequest) GetTargetEndTime() *int32 {
	return s.TargetEndTime
}

func (s *ModifyCycleTaskRequest) GetTargetStartTime() *int32 {
	return s.TargetStartTime
}

func (s *ModifyCycleTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ModifyCycleTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ModifyCycleTaskRequest) SetConfigId(v string) *ModifyCycleTaskRequest {
	s.ConfigId = &v
	return s
}

func (s *ModifyCycleTaskRequest) SetEnable(v int32) *ModifyCycleTaskRequest {
	s.Enable = &v
	return s
}

func (s *ModifyCycleTaskRequest) SetFirstDateStr(v int64) *ModifyCycleTaskRequest {
	s.FirstDateStr = &v
	return s
}

func (s *ModifyCycleTaskRequest) SetIntervalPeriod(v int32) *ModifyCycleTaskRequest {
	s.IntervalPeriod = &v
	return s
}

func (s *ModifyCycleTaskRequest) SetParam(v string) *ModifyCycleTaskRequest {
	s.Param = &v
	return s
}

func (s *ModifyCycleTaskRequest) SetPeriodUnit(v string) *ModifyCycleTaskRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyCycleTaskRequest) SetTargetEndTime(v int32) *ModifyCycleTaskRequest {
	s.TargetEndTime = &v
	return s
}

func (s *ModifyCycleTaskRequest) SetTargetStartTime(v int32) *ModifyCycleTaskRequest {
	s.TargetStartTime = &v
	return s
}

func (s *ModifyCycleTaskRequest) SetTaskName(v string) *ModifyCycleTaskRequest {
	s.TaskName = &v
	return s
}

func (s *ModifyCycleTaskRequest) SetTaskType(v string) *ModifyCycleTaskRequest {
	s.TaskType = &v
	return s
}

func (s *ModifyCycleTaskRequest) Validate() error {
	return dara.Validate(s)
}
