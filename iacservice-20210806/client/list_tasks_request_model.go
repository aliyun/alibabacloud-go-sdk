// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ListTasksRequest
	GetGroupId() *string
	SetKeyword(v string) *ListTasksRequest
	GetKeyword() *string
	SetModuleId(v string) *ListTasksRequest
	GetModuleId() *string
	SetPageNumber(v int32) *ListTasksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTasksRequest
	GetPageSize() *int32
	SetProjectId(v string) *ListTasksRequest
	GetProjectId() *string
	SetStatus(v string) *ListTasksRequest
	GetStatus() *string
	SetTag(v []*ListTasksRequestTag) *ListTasksRequest
	GetTag() []*ListTasksRequestTag
	SetTaskId(v string) *ListTasksRequest
	GetTaskId() *string
}

type ListTasksRequest struct {
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
	ProjectId *string                `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Status    *string                `json:"status,omitempty" xml:"status,omitempty"`
	Tag       []*ListTasksRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
	// example:
	//
	// task-433aead756057fffeaba4828f5195
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s ListTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *ListTasksRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListTasksRequest) GetModuleId() *string {
	return s.ModuleId
}

func (s *ListTasksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTasksRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListTasksRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTasksRequest) GetTag() []*ListTasksRequestTag {
	return s.Tag
}

func (s *ListTasksRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListTasksRequest) SetGroupId(v string) *ListTasksRequest {
	s.GroupId = &v
	return s
}

func (s *ListTasksRequest) SetKeyword(v string) *ListTasksRequest {
	s.Keyword = &v
	return s
}

func (s *ListTasksRequest) SetModuleId(v string) *ListTasksRequest {
	s.ModuleId = &v
	return s
}

func (s *ListTasksRequest) SetPageNumber(v int32) *ListTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTasksRequest) SetPageSize(v int32) *ListTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListTasksRequest) SetProjectId(v string) *ListTasksRequest {
	s.ProjectId = &v
	return s
}

func (s *ListTasksRequest) SetStatus(v string) *ListTasksRequest {
	s.Status = &v
	return s
}

func (s *ListTasksRequest) SetTag(v []*ListTasksRequestTag) *ListTasksRequest {
	s.Tag = v
	return s
}

func (s *ListTasksRequest) SetTaskId(v string) *ListTasksRequest {
	s.TaskId = &v
	return s
}

func (s *ListTasksRequest) Validate() error {
	return dara.Validate(s)
}

type ListTasksRequestTag struct {
	TagKey   *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListTasksRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTasksRequestTag) GoString() string {
	return s.String()
}

func (s *ListTasksRequestTag) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTasksRequestTag) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTasksRequestTag) SetTagKey(v string) *ListTasksRequestTag {
	s.TagKey = &v
	return s
}

func (s *ListTasksRequestTag) SetTagValue(v string) *ListTasksRequestTag {
	s.TagValue = &v
	return s
}

func (s *ListTasksRequestTag) Validate() error {
	return dara.Validate(s)
}
