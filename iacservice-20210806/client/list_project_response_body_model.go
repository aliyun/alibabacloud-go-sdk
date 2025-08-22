// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListProjectResponseBody
	GetCount() *int64
	SetPageNumber(v int64) *ListProjectResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListProjectResponseBody
	GetPageSize() *int64
	SetProjects(v []*ListProjectResponseBodyProjects) *ListProjectResponseBody
	GetProjects() []*ListProjectResponseBodyProjects
	SetRequestId(v string) *ListProjectResponseBody
	GetRequestId() *string
}

type ListProjectResponseBody struct {
	// example:
	//
	// 3
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64                             `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Projects []*ListProjectResponseBodyProjects `json:"projects,omitempty" xml:"projects,omitempty" type:"Repeated"`
	// example:
	//
	// 136B3926-DD90-5DB2-96EC-8BAD6407D1C9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListProjectResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListProjectResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListProjectResponseBody) GetProjects() []*ListProjectResponseBodyProjects {
	return s.Projects
}

func (s *ListProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectResponseBody) SetCount(v int64) *ListProjectResponseBody {
	s.Count = &v
	return s
}

func (s *ListProjectResponseBody) SetPageNumber(v int64) *ListProjectResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListProjectResponseBody) SetPageSize(v int64) *ListProjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProjectResponseBody) SetProjects(v []*ListProjectResponseBodyProjects) *ListProjectResponseBody {
	s.Projects = v
	return s
}

func (s *ListProjectResponseBody) SetRequestId(v string) *ListProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListProjectResponseBodyProjects struct {
	// example:
	//
	// 2022-05-10T10:08:34Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// abc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1234
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// p-148e7853433574fffe9fec72ed9b72
	ProjectId *string                                `json:"projectId,omitempty" xml:"projectId,omitempty"`
	Tags      []*ListProjectResponseBodyProjectsTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TaskCnt *int64 `json:"taskCnt,omitempty" xml:"taskCnt,omitempty"`
}

func (s ListProjectResponseBodyProjects) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBodyProjects) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListProjectResponseBodyProjects) GetDescription() *string {
	return s.Description
}

func (s *ListProjectResponseBodyProjects) GetName() *string {
	return s.Name
}

func (s *ListProjectResponseBodyProjects) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListProjectResponseBodyProjects) GetTags() []*ListProjectResponseBodyProjectsTags {
	return s.Tags
}

func (s *ListProjectResponseBodyProjects) GetTaskCnt() *int64 {
	return s.TaskCnt
}

func (s *ListProjectResponseBodyProjects) SetCreateTime(v string) *ListProjectResponseBodyProjects {
	s.CreateTime = &v
	return s
}

func (s *ListProjectResponseBodyProjects) SetDescription(v string) *ListProjectResponseBodyProjects {
	s.Description = &v
	return s
}

func (s *ListProjectResponseBodyProjects) SetName(v string) *ListProjectResponseBodyProjects {
	s.Name = &v
	return s
}

func (s *ListProjectResponseBodyProjects) SetProjectId(v string) *ListProjectResponseBodyProjects {
	s.ProjectId = &v
	return s
}

func (s *ListProjectResponseBodyProjects) SetTags(v []*ListProjectResponseBodyProjectsTags) *ListProjectResponseBodyProjects {
	s.Tags = v
	return s
}

func (s *ListProjectResponseBodyProjects) SetTaskCnt(v int64) *ListProjectResponseBodyProjects {
	s.TaskCnt = &v
	return s
}

func (s *ListProjectResponseBodyProjects) Validate() error {
	return dara.Validate(s)
}

type ListProjectResponseBodyProjectsTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListProjectResponseBodyProjectsTags) String() string {
	return dara.Prettify(s)
}

func (s ListProjectResponseBodyProjectsTags) GoString() string {
	return s.String()
}

func (s *ListProjectResponseBodyProjectsTags) GetKey() *string {
	return s.Key
}

func (s *ListProjectResponseBodyProjectsTags) GetValue() *string {
	return s.Value
}

func (s *ListProjectResponseBodyProjectsTags) SetKey(v string) *ListProjectResponseBodyProjectsTags {
	s.Key = &v
	return s
}

func (s *ListProjectResponseBodyProjectsTags) SetValue(v string) *ListProjectResponseBodyProjectsTags {
	s.Value = &v
	return s
}

func (s *ListProjectResponseBodyProjectsTags) Validate() error {
	return dara.Validate(s)
}
