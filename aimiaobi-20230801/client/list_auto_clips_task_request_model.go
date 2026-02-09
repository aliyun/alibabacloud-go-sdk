// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAutoClipsTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTimeEnd(v string) *ListAutoClipsTaskRequest
	GetCreateTimeEnd() *string
	SetCreateTimeStart(v string) *ListAutoClipsTaskRequest
	GetCreateTimeStart() *string
	SetCurrent(v int32) *ListAutoClipsTaskRequest
	GetCurrent() *int32
	SetMaxResults(v int32) *ListAutoClipsTaskRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAutoClipsTaskRequest
	GetNextToken() *string
	SetSize(v int32) *ListAutoClipsTaskRequest
	GetSize() *int32
	SetSkip(v int32) *ListAutoClipsTaskRequest
	GetSkip() *int32
	SetTaskName(v string) *ListAutoClipsTaskRequest
	GetTaskName() *string
	SetTaskStatus(v int32) *ListAutoClipsTaskRequest
	GetTaskStatus() *int32
	SetTaskType(v string) *ListAutoClipsTaskRequest
	GetTaskType() *string
	SetWorkspaceId(v string) *ListAutoClipsTaskRequest
	GetWorkspaceId() *string
}

type ListAutoClipsTaskRequest struct {
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
	// null
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// null
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// null
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// example:
	//
	// task001
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// 0
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// type001
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListAutoClipsTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAutoClipsTaskRequest) GoString() string {
	return s.String()
}

func (s *ListAutoClipsTaskRequest) GetCreateTimeEnd() *string {
	return s.CreateTimeEnd
}

func (s *ListAutoClipsTaskRequest) GetCreateTimeStart() *string {
	return s.CreateTimeStart
}

func (s *ListAutoClipsTaskRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListAutoClipsTaskRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAutoClipsTaskRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAutoClipsTaskRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListAutoClipsTaskRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListAutoClipsTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *ListAutoClipsTaskRequest) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ListAutoClipsTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListAutoClipsTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAutoClipsTaskRequest) SetCreateTimeEnd(v string) *ListAutoClipsTaskRequest {
	s.CreateTimeEnd = &v
	return s
}

func (s *ListAutoClipsTaskRequest) SetCreateTimeStart(v string) *ListAutoClipsTaskRequest {
	s.CreateTimeStart = &v
	return s
}

func (s *ListAutoClipsTaskRequest) SetCurrent(v int32) *ListAutoClipsTaskRequest {
	s.Current = &v
	return s
}

func (s *ListAutoClipsTaskRequest) SetMaxResults(v int32) *ListAutoClipsTaskRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAutoClipsTaskRequest) SetNextToken(v string) *ListAutoClipsTaskRequest {
	s.NextToken = &v
	return s
}

func (s *ListAutoClipsTaskRequest) SetSize(v int32) *ListAutoClipsTaskRequest {
	s.Size = &v
	return s
}

func (s *ListAutoClipsTaskRequest) SetSkip(v int32) *ListAutoClipsTaskRequest {
	s.Skip = &v
	return s
}

func (s *ListAutoClipsTaskRequest) SetTaskName(v string) *ListAutoClipsTaskRequest {
	s.TaskName = &v
	return s
}

func (s *ListAutoClipsTaskRequest) SetTaskStatus(v int32) *ListAutoClipsTaskRequest {
	s.TaskStatus = &v
	return s
}

func (s *ListAutoClipsTaskRequest) SetTaskType(v string) *ListAutoClipsTaskRequest {
	s.TaskType = &v
	return s
}

func (s *ListAutoClipsTaskRequest) SetWorkspaceId(v string) *ListAutoClipsTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListAutoClipsTaskRequest) Validate() error {
	return dara.Validate(s)
}
