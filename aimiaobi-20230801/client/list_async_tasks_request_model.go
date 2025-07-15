// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListAsyncTasksRequest
	GetAgentKey() *string
	SetCreateTimeEnd(v string) *ListAsyncTasksRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *ListAsyncTasksRequest
	GetCreateTimeStart() *string
	SetCurrent(v int32) *ListAsyncTasksRequest
	GetCurrent() *int32
	SetSize(v int32) *ListAsyncTasksRequest
	GetSize() *int32
	SetTaskCode(v string) *ListAsyncTasksRequest
	GetTaskCode() *string
	SetTaskName(v string) *ListAsyncTasksRequest
	GetTaskName() *string
	SetTaskStatus(v int32) *ListAsyncTasksRequest
	GetTaskStatus() *int32
	SetTaskStatusList(v []*int32) *ListAsyncTasksRequest
	GetTaskStatusList() []*int32
	SetTaskType(v string) *ListAsyncTasksRequest
	GetTaskType() *string
	SetTaskTypeList(v []*string) *ListAsyncTasksRequest
	GetTaskTypeList() []*string
}

type ListAsyncTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cd327c3d5d5e44159cc716e23bfa530e_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 2023-03-18 02:00:00
	CreateTimeEnd *string `json:"CreateTimeEnd,omitempty" xml:"CreateTimeEnd,omitempty"`
	// example:
	//
	// 2023-02-19 07:28:11
	CreateTimeStart *string `json:"CreateTimeStart,omitempty" xml:"CreateTimeStart,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// MaterialDocumentUpload
	TaskCode *string `json:"TaskCode,omitempty" xml:"TaskCode,omitempty"`
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 1
	TaskStatus     *int32    `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusList []*int32  `json:"TaskStatusList,omitempty" xml:"TaskStatusList,omitempty" type:"Repeated"`
	TaskType       *string   `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeList   []*string `json:"TaskTypeList,omitempty" xml:"TaskTypeList,omitempty" type:"Repeated"`
}

func (s ListAsyncTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTasksRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListAsyncTasksRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *ListAsyncTasksRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *ListAsyncTasksRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListAsyncTasksRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListAsyncTasksRequest) GetTaskCode() *string {
	return s.TaskCode
}

func (s *ListAsyncTasksRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListAsyncTasksRequest) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ListAsyncTasksRequest) GetTaskStatusList() []*int32 {
	return s.TaskStatusList
}

func (s *ListAsyncTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListAsyncTasksRequest) GetTaskTypeList() []*string {
	return s.TaskTypeList
}

func (s *ListAsyncTasksRequest) SetAgentKey(v string) *ListAsyncTasksRequest {
	s.AgentKey = &v
	return s
}

func (s *ListAsyncTasksRequest) SetCreateTimeEnd(v string) *ListAsyncTasksRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListAsyncTasksRequest) SetCreateTimeStart(v string) *ListAsyncTasksRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListAsyncTasksRequest) SetCurrent(v int32) *ListAsyncTasksRequest {
	s.Current = &v
	return s
}

func (s *ListAsyncTasksRequest) SetSize(v int32) *ListAsyncTasksRequest {
	s.Size = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskCode(v string) *ListAsyncTasksRequest {
	s.TaskCode = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskName(v string) *ListAsyncTasksRequest {
	s.TaskName = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskStatus(v int32) *ListAsyncTasksRequest {
	s.TaskStatus = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskStatusList(v []*int32) *ListAsyncTasksRequest {
	s.TaskStatusList = v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskType(v string) *ListAsyncTasksRequest {
	s.TaskType = &v
	return s
}

func (s *ListAsyncTasksRequest) SetTaskTypeList(v []*string) *ListAsyncTasksRequest {
	s.TaskTypeList = v
	return s
}

func (s *ListAsyncTasksRequest) Validate() error {
	return dara.Validate(s)
}
