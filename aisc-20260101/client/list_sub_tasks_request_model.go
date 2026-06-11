// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListSubTasksRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListSubTasksRequest
	GetPageSize() *int32
	SetRootTaskId(v string) *ListSubTasksRequest
	GetRootTaskId() *string
	SetTaskType(v string) *ListSubTasksRequest
	GetTaskType() *string
}

type ListSubTasksRequest struct {
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 110146ac282314fe4f7cd11afc7540eb
	RootTaskId *string `json:"RootTaskId,omitempty" xml:"RootTaskId,omitempty"`
	// example:
	//
	// SKILL_CHECK
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListSubTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubTasksRequest) GoString() string {
	return s.String()
}

func (s *ListSubTasksRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListSubTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSubTasksRequest) GetRootTaskId() *string {
	return s.RootTaskId
}

func (s *ListSubTasksRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *ListSubTasksRequest) SetCurrentPage(v int32) *ListSubTasksRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListSubTasksRequest) SetPageSize(v int32) *ListSubTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListSubTasksRequest) SetRootTaskId(v string) *ListSubTasksRequest {
	s.RootTaskId = &v
	return s
}

func (s *ListSubTasksRequest) SetTaskType(v string) *ListSubTasksRequest {
	s.TaskType = &v
	return s
}

func (s *ListSubTasksRequest) Validate() error {
	return dara.Validate(s)
}
