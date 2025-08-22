// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListTasksShrinkRequest
	GetGroupId() *string
	SetKeyword(v string) *ListTasksShrinkRequest
	GetKeyword() *string
	SetModuleId(v string) *ListTasksShrinkRequest
	GetModuleId() *string
	SetPageNumber(v int32) *ListTasksShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTasksShrinkRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListTasksShrinkRequest
	GetProjectId() *string
	SetStatus(v string) *ListTasksShrinkRequest
	GetStatus() *string
	SetTagShrink(v string) *ListTasksShrinkRequest
	GetTagShrink() *string
	SetTaskId(v string) *ListTasksShrinkRequest
	GetTaskId() *string
}

type ListTasksShrinkRequest struct {
	// example:
	//
	// g-59d8d22e78792ffe3d3eb6154d727
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// key
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// mod-1525e992f1b62139d1c437d64ae
	ModuleId *string `json:"moduleId,omitempty" xml:"moduleId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// p-433aead7560572f8d95b25775c
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// example:
	//
	// task-433aead756057fffeaba4828f5195
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListTasksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTasksShrinkRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListTasksShrinkRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTasksShrinkRequest) GetModuleId() *string {
	return s.ModuleId
}

func (s *ListTasksShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTasksShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTasksShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListTasksShrinkRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTasksShrinkRequest) GetTagShrink() *string {
	return s.TagShrink
}

func (s *ListTasksShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTasksShrinkRequest) SetGroupId(v string) *ListTasksShrinkRequest {
	s.GroupId = &v
	return s
}

func (s *ListTasksShrinkRequest) SetKeyword(v string) *ListTasksShrinkRequest {
	s.Keyword = &v
	return s
}

func (s *ListTasksShrinkRequest) SetModuleId(v string) *ListTasksShrinkRequest {
	s.ModuleId = &v
	return s
}

func (s *ListTasksShrinkRequest) SetPageNumber(v int32) *ListTasksShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksShrinkRequest) SetPageSize(v int32) *ListTasksShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksShrinkRequest) SetProjectId(v string) *ListTasksShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTasksShrinkRequest) SetStatus(v string) *ListTasksShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListTasksShrinkRequest) SetTagShrink(v string) *ListTasksShrinkRequest {
	s.TagShrink = &v
	return s
}

func (s *ListTasksShrinkRequest) SetTaskId(v string) *ListTasksShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *ListTasksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
