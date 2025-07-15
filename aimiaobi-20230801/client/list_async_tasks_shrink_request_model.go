// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAsyncTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *ListAsyncTasksShrinkRequest
	GetAgentKey() *string
	SetCreateTimeEnd(v string) *ListAsyncTasksShrinkRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *ListAsyncTasksShrinkRequest
	GetCreateTimeStart() *string
	SetCurrent(v int32) *ListAsyncTasksShrinkRequest
	GetCurrent() *int32
	SetSize(v int32) *ListAsyncTasksShrinkRequest
	GetSize() *int32
	SetTaskCode(v string) *ListAsyncTasksShrinkRequest
	GetTaskCode() *string
	SetTaskName(v string) *ListAsyncTasksShrinkRequest
	GetTaskName() *string
	SetTaskStatus(v int32) *ListAsyncTasksShrinkRequest
	GetTaskStatus() *int32
	SetTaskStatusListShrink(v string) *ListAsyncTasksShrinkRequest
	GetTaskStatusListShrink() *string
	SetTaskType(v string) *ListAsyncTasksShrinkRequest
	GetTaskType() *string
	SetTaskTypeListShrink(v string) *ListAsyncTasksShrinkRequest
	GetTaskTypeListShrink() *string
}

type ListAsyncTasksShrinkRequest struct {
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
	TaskStatus           *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskStatusListShrink *string `json:"TaskStatusList,omitempty" xml:"TaskStatusList,omitempty"`
	TaskType             *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskTypeListShrink   *string `json:"TaskTypeList,omitempty" xml:"TaskTypeList,omitempty"`
}

func (s ListAsyncTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAsyncTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAsyncTasksShrinkRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *ListAsyncTasksShrinkRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *ListAsyncTasksShrinkRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *ListAsyncTasksShrinkRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListAsyncTasksShrinkRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListAsyncTasksShrinkRequest) GetTaskCode() *string {
	return s.TaskCode
}

func (s *ListAsyncTasksShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListAsyncTasksShrinkRequest) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ListAsyncTasksShrinkRequest) GetTaskStatusListShrink() *string {
	return s.TaskStatusListShrink
}

func (s *ListAsyncTasksShrinkRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListAsyncTasksShrinkRequest) GetTaskTypeListShrink() *string {
	return s.TaskTypeListShrink
}

func (s *ListAsyncTasksShrinkRequest) SetAgentKey(v string) *ListAsyncTasksShrinkRequest {
	s.AgentKey = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetCreateTimeEnd(v string) *ListAsyncTasksShrinkRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetCreateTimeStart(v string) *ListAsyncTasksShrinkRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetCurrent(v int32) *ListAsyncTasksShrinkRequest {
	s.Current = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetSize(v int32) *ListAsyncTasksShrinkRequest {
	s.Size = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskCode(v string) *ListAsyncTasksShrinkRequest {
	s.TaskCode = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskName(v string) *ListAsyncTasksShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskStatus(v int32) *ListAsyncTasksShrinkRequest {
	s.TaskStatus = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskStatusListShrink(v string) *ListAsyncTasksShrinkRequest {
	s.TaskStatusListShrink = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskType(v string) *ListAsyncTasksShrinkRequest {
	s.TaskType = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) SetTaskTypeListShrink(v string) *ListAsyncTasksShrinkRequest {
	s.TaskTypeListShrink = &v
	return s
}

func (s *ListAsyncTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
