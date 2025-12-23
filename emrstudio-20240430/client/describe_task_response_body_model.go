// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeTaskResponseBody
	GetCreateTime() *string
	SetDelayTime(v int32) *DescribeTaskResponseBody
	GetDelayTime() *int32
	SetDescription(v string) *DescribeTaskResponseBody
	GetDescription() *string
	SetFailRetryInterval(v int32) *DescribeTaskResponseBody
	GetFailRetryInterval() *int32
	SetFailRetryTimes(v int32) *DescribeTaskResponseBody
	GetFailRetryTimes() *int32
	SetFlag(v string) *DescribeTaskResponseBody
	GetFlag() *string
	SetProjectId(v string) *DescribeTaskResponseBody
	GetProjectId() *string
	SetResourceIds(v string) *DescribeTaskResponseBody
	GetResourceIds() *string
	SetTaskId(v string) *DescribeTaskResponseBody
	GetTaskId() *string
	SetTaskName(v string) *DescribeTaskResponseBody
	GetTaskName() *string
	SetTaskParams(v string) *DescribeTaskResponseBody
	GetTaskParams() *string
	SetTaskPriority(v string) *DescribeTaskResponseBody
	GetTaskPriority() *string
	SetTaskType(v string) *DescribeTaskResponseBody
	GetTaskType() *string
	SetTimeout(v int32) *DescribeTaskResponseBody
	GetTimeout() *int32
	SetTimeoutFlag(v string) *DescribeTaskResponseBody
	GetTimeoutFlag() *string
	SetTimeoutNotifyStrategy(v string) *DescribeTaskResponseBody
	GetTimeoutNotifyStrategy() *string
	SetUpdateTime(v string) *DescribeTaskResponseBody
	GetUpdateTime() *string
	SetVersion(v string) *DescribeTaskResponseBody
	GetVersion() *string
	SetRequestId(v string) *DescribeTaskResponseBody
	GetRequestId() *string
}

type DescribeTaskResponseBody struct {
	// example:
	//
	// 2024-03-27 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 0
	DelayTime *int32 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 1
	FailRetryInterval *int32 `json:"FailRetryInterval,omitempty" xml:"FailRetryInterval,omitempty"`
	// example:
	//
	// 0
	FailRetryTimes *int32 `json:"FailRetryTimes,omitempty" xml:"FailRetryTimes,omitempty"`
	// example:
	//
	// YES
	Flag *string `json:"Flag,omitempty" xml:"Flag,omitempty"`
	// example:
	//
	// p-3q9jo749ne5****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// r-oy98v7n43el****
	ResourceIds *string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty"`
	// example:
	//
	// t-3q9jo749ne5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// test
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// {
	//
	//     "yarnUser": "",
	//
	//     "conditionResult": "null",
	//
	//     "rawScript": "sleep 300",
	//
	//     "submitOnYarnFlag": false,
	//
	//     "emrClusterId": "",
	//
	//     "yarnPriority": "",
	//
	//     "dependence": "null",
	//
	//     "yarnMemory": "",
	//
	//     "localParams": [],
	//
	//     "switchResult": "null",
	//
	//     "resourceList": [],
	//
	//     "yarnQueue": "",
	//
	//     "yarnVCores": "",
	//
	//     "associateManualTaskFlag": false
	//
	// }
	TaskParams *string `json:"TaskParams,omitempty" xml:"TaskParams,omitempty"`
	// example:
	//
	// MEDIUM
	TaskPriority *string `json:"TaskPriority,omitempty" xml:"TaskPriority,omitempty"`
	// example:
	//
	// SHELL
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// 10
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// example:
	//
	// CLOSE
	TimeoutFlag *string `json:"TimeoutFlag,omitempty" xml:"TimeoutFlag,omitempty"`
	// example:
	//
	// WARN
	TimeoutNotifyStrategy *string `json:"TimeoutNotifyStrategy,omitempty" xml:"TimeoutNotifyStrategy,omitempty"`
	// example:
	//
	// 2024-03-27 00:00:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DescribeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeTaskResponseBody) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *DescribeTaskResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeTaskResponseBody) GetFailRetryInterval() *int32 {
	return s.FailRetryInterval
}

func (s *DescribeTaskResponseBody) GetFailRetryTimes() *int32 {
	return s.FailRetryTimes
}

func (s *DescribeTaskResponseBody) GetFlag() *string {
	return s.Flag
}

func (s *DescribeTaskResponseBody) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeTaskResponseBody) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *DescribeTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTaskResponseBody) GetTaskName() *string {
	return s.TaskName
}

func (s *DescribeTaskResponseBody) GetTaskParams() *string {
	return s.TaskParams
}

func (s *DescribeTaskResponseBody) GetTaskPriority() *string {
	return s.TaskPriority
}

func (s *DescribeTaskResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeTaskResponseBody) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DescribeTaskResponseBody) GetTimeoutFlag() *string {
	return s.TimeoutFlag
}

func (s *DescribeTaskResponseBody) GetTimeoutNotifyStrategy() *string {
	return s.TimeoutNotifyStrategy
}

func (s *DescribeTaskResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeTaskResponseBody) GetVersion() *string {
	return s.Version
}

func (s *DescribeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTaskResponseBody) SetCreateTime(v string) *DescribeTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeTaskResponseBody) SetDelayTime(v int32) *DescribeTaskResponseBody {
	s.DelayTime = &v
	return s
}

func (s *DescribeTaskResponseBody) SetDescription(v string) *DescribeTaskResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeTaskResponseBody) SetFailRetryInterval(v int32) *DescribeTaskResponseBody {
	s.FailRetryInterval = &v
	return s
}

func (s *DescribeTaskResponseBody) SetFailRetryTimes(v int32) *DescribeTaskResponseBody {
	s.FailRetryTimes = &v
	return s
}

func (s *DescribeTaskResponseBody) SetFlag(v string) *DescribeTaskResponseBody {
	s.Flag = &v
	return s
}

func (s *DescribeTaskResponseBody) SetProjectId(v string) *DescribeTaskResponseBody {
	s.ProjectId = &v
	return s
}

func (s *DescribeTaskResponseBody) SetResourceIds(v string) *DescribeTaskResponseBody {
	s.ResourceIds = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskId(v string) *DescribeTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskName(v string) *DescribeTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskParams(v string) *DescribeTaskResponseBody {
	s.TaskParams = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskPriority(v string) *DescribeTaskResponseBody {
	s.TaskPriority = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskType(v string) *DescribeTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTimeout(v int32) *DescribeTaskResponseBody {
	s.Timeout = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTimeoutFlag(v string) *DescribeTaskResponseBody {
	s.TimeoutFlag = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTimeoutNotifyStrategy(v string) *DescribeTaskResponseBody {
	s.TimeoutNotifyStrategy = &v
	return s
}

func (s *DescribeTaskResponseBody) SetUpdateTime(v string) *DescribeTaskResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeTaskResponseBody) SetVersion(v string) *DescribeTaskResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeTaskResponseBody) SetRequestId(v string) *DescribeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
