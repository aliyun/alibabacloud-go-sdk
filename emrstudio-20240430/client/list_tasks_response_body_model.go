// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListTasksResponseBodyData) *ListTasksResponseBody
	GetData() []*ListTasksResponseBodyData
	SetNextToken(v string) *ListTasksResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTasksResponseBody
	GetRequestId() *string
	SetTotalSize(v int32) *ListTasksResponseBody
	GetTotalSize() *int32
}

type ListTasksResponseBody struct {
	Data []*ListTasksResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 123abc***
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) GetData() []*ListTasksResponseBodyData {
	return s.Data
}

func (s *ListTasksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTasksResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListTasksResponseBody) SetData(v []*ListTasksResponseBodyData) *ListTasksResponseBody {
	s.Data = v
	return s
}

func (s *ListTasksResponseBody) SetNextToken(v string) *ListTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetTotalSize(v int32) *ListTasksResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListTasksResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTasksResponseBodyData struct {
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
	// r-3q9jo749ne5****
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
	// 0
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
}

func (s ListTasksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTasksResponseBodyData) GetDelayTime() *int32 {
	return s.DelayTime
}

func (s *ListTasksResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListTasksResponseBodyData) GetFailRetryInterval() *int32 {
	return s.FailRetryInterval
}

func (s *ListTasksResponseBodyData) GetFailRetryTimes() *int32 {
	return s.FailRetryTimes
}

func (s *ListTasksResponseBodyData) GetFlag() *string {
	return s.Flag
}

func (s *ListTasksResponseBodyData) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListTasksResponseBodyData) GetResourceIds() *string {
	return s.ResourceIds
}

func (s *ListTasksResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTasksResponseBodyData) GetTaskName() *string {
	return s.TaskName
}

func (s *ListTasksResponseBodyData) GetTaskParams() *string {
	return s.TaskParams
}

func (s *ListTasksResponseBodyData) GetTaskPriority() *string {
	return s.TaskPriority
}

func (s *ListTasksResponseBodyData) GetTaskType() *string {
	return s.TaskType
}

func (s *ListTasksResponseBodyData) GetTimeout() *int32 {
	return s.Timeout
}

func (s *ListTasksResponseBodyData) GetTimeoutFlag() *string {
	return s.TimeoutFlag
}

func (s *ListTasksResponseBodyData) GetTimeoutNotifyStrategy() *string {
	return s.TimeoutNotifyStrategy
}

func (s *ListTasksResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListTasksResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *ListTasksResponseBodyData) SetCreateTime(v string) *ListTasksResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListTasksResponseBodyData) SetDelayTime(v int32) *ListTasksResponseBodyData {
	s.DelayTime = &v
	return s
}

func (s *ListTasksResponseBodyData) SetDescription(v string) *ListTasksResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListTasksResponseBodyData) SetFailRetryInterval(v int32) *ListTasksResponseBodyData {
	s.FailRetryInterval = &v
	return s
}

func (s *ListTasksResponseBodyData) SetFailRetryTimes(v int32) *ListTasksResponseBodyData {
	s.FailRetryTimes = &v
	return s
}

func (s *ListTasksResponseBodyData) SetFlag(v string) *ListTasksResponseBodyData {
	s.Flag = &v
	return s
}

func (s *ListTasksResponseBodyData) SetProjectId(v string) *ListTasksResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *ListTasksResponseBodyData) SetResourceIds(v string) *ListTasksResponseBodyData {
	s.ResourceIds = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTaskId(v string) *ListTasksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTaskName(v string) *ListTasksResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTaskParams(v string) *ListTasksResponseBodyData {
	s.TaskParams = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTaskPriority(v string) *ListTasksResponseBodyData {
	s.TaskPriority = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTaskType(v string) *ListTasksResponseBodyData {
	s.TaskType = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTimeout(v int32) *ListTasksResponseBodyData {
	s.Timeout = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTimeoutFlag(v string) *ListTasksResponseBodyData {
	s.TimeoutFlag = &v
	return s
}

func (s *ListTasksResponseBodyData) SetTimeoutNotifyStrategy(v string) *ListTasksResponseBodyData {
	s.TimeoutNotifyStrategy = &v
	return s
}

func (s *ListTasksResponseBodyData) SetUpdateTime(v string) *ListTasksResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListTasksResponseBodyData) SetVersion(v string) *ListTasksResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListTasksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
