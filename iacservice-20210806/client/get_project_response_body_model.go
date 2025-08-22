// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProject(v *GetProjectResponseBodyProject) *GetProjectResponseBody
	GetProject() *GetProjectResponseBodyProject
	SetRequestId(v string) *GetProjectResponseBody
	GetRequestId() *string
}

type GetProjectResponseBody struct {
	Project *GetProjectResponseBodyProject `json:"project,omitempty" xml:"project,omitempty" type:"Struct"`
	// example:
	//
	// 7FA0FF4A-ABD4-54F6-BEAC-B4273EBA10A2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) GetProject() *GetProjectResponseBodyProject {
	return s.Project
}

func (s *GetProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectResponseBody) SetProject(v *GetProjectResponseBodyProject) *GetProjectResponseBody {
	s.Project = v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetProjectResponseBodyProject struct {
	// example:
	//
	// 2022-09-06T06:11:27Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// abc
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// abc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// p-433aead7560572f8d95b25775c
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// example:
	//
	// 2
	TaskCnt *int64 `json:"taskCnt,omitempty" xml:"taskCnt,omitempty"`
}

func (s GetProjectResponseBodyProject) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyProject) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProjectResponseBodyProject) GetDescription() *string {
	return s.Description
}

func (s *GetProjectResponseBodyProject) GetName() *string {
	return s.Name
}

func (s *GetProjectResponseBodyProject) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetProjectResponseBodyProject) GetTaskCnt() *int64 {
	return s.TaskCnt
}

func (s *GetProjectResponseBodyProject) SetCreateTime(v string) *GetProjectResponseBodyProject {
	s.CreateTime = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetDescription(v string) *GetProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetName(v string) *GetProjectResponseBodyProject {
	s.Name = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetProjectId(v string) *GetProjectResponseBodyProject {
	s.ProjectId = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetTaskCnt(v int64) *GetProjectResponseBodyProject {
	s.TaskCnt = &v
	return s
}

func (s *GetProjectResponseBodyProject) Validate() error {
	return dara.Validate(s)
}
