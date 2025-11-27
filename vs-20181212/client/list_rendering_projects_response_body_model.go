// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjects(v []*ListRenderingProjectsResponseBodyProjects) *ListRenderingProjectsResponseBody
	GetProjects() []*ListRenderingProjectsResponseBodyProjects
	SetRequestId(v string) *ListRenderingProjectsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRenderingProjectsResponseBody
	GetTotalCount() *int64
}

type ListRenderingProjectsResponseBody struct {
	Projects []*ListRenderingProjectsResponseBodyProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRenderingProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRenderingProjectsResponseBody) GetProjects() []*ListRenderingProjectsResponseBodyProjects {
	return s.Projects
}

func (s *ListRenderingProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRenderingProjectsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRenderingProjectsResponseBody) SetProjects(v []*ListRenderingProjectsResponseBodyProjects) *ListRenderingProjectsResponseBody {
	s.Projects = v
	return s
}

func (s *ListRenderingProjectsResponseBody) SetRequestId(v string) *ListRenderingProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRenderingProjectsResponseBody) SetTotalCount(v int64) *ListRenderingProjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRenderingProjectsResponseBody) Validate() error {
	if s.Projects != nil {
		for _, item := range s.Projects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRenderingProjectsResponseBodyProjects struct {
	// example:
	//
	// 2024-09-09T18:44:49+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// prod-project
	ProjectName    *string                                                  `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SessionAttribs *ListRenderingProjectsResponseBodyProjectsSessionAttribs `json:"SessionAttribs,omitempty" xml:"SessionAttribs,omitempty" type:"Struct"`
	// example:
	//
	// 2024-10-09T18:44:49+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListRenderingProjectsResponseBodyProjects) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingProjectsResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListRenderingProjectsResponseBodyProjects) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListRenderingProjectsResponseBodyProjects) GetDescription() *string {
	return s.Description
}

func (s *ListRenderingProjectsResponseBodyProjects) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListRenderingProjectsResponseBodyProjects) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListRenderingProjectsResponseBodyProjects) GetSessionAttribs() *ListRenderingProjectsResponseBodyProjectsSessionAttribs {
	return s.SessionAttribs
}

func (s *ListRenderingProjectsResponseBodyProjects) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListRenderingProjectsResponseBodyProjects) SetCreationTime(v string) *ListRenderingProjectsResponseBodyProjects {
	s.CreationTime = &v
	return s
}

func (s *ListRenderingProjectsResponseBodyProjects) SetDescription(v string) *ListRenderingProjectsResponseBodyProjects {
	s.Description = &v
	return s
}

func (s *ListRenderingProjectsResponseBodyProjects) SetProjectId(v string) *ListRenderingProjectsResponseBodyProjects {
	s.ProjectId = &v
	return s
}

func (s *ListRenderingProjectsResponseBodyProjects) SetProjectName(v string) *ListRenderingProjectsResponseBodyProjects {
	s.ProjectName = &v
	return s
}

func (s *ListRenderingProjectsResponseBodyProjects) SetSessionAttribs(v *ListRenderingProjectsResponseBodyProjectsSessionAttribs) *ListRenderingProjectsResponseBodyProjects {
	s.SessionAttribs = v
	return s
}

func (s *ListRenderingProjectsResponseBodyProjects) SetUpdateTime(v string) *ListRenderingProjectsResponseBodyProjects {
	s.UpdateTime = &v
	return s
}

func (s *ListRenderingProjectsResponseBodyProjects) Validate() error {
	if s.SessionAttribs != nil {
		if err := s.SessionAttribs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRenderingProjectsResponseBodyProjectsSessionAttribs struct {
	// example:
	//
	// Sync
	StartMode *string `json:"StartMode,omitempty" xml:"StartMode,omitempty"`
}

func (s ListRenderingProjectsResponseBodyProjectsSessionAttribs) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingProjectsResponseBodyProjectsSessionAttribs) GoString() string {
	return s.String()
}

func (s *ListRenderingProjectsResponseBodyProjectsSessionAttribs) GetStartMode() *string {
	return s.StartMode
}

func (s *ListRenderingProjectsResponseBodyProjectsSessionAttribs) SetStartMode(v string) *ListRenderingProjectsResponseBodyProjectsSessionAttribs {
	s.StartMode = &v
	return s
}

func (s *ListRenderingProjectsResponseBodyProjectsSessionAttribs) Validate() error {
	return dara.Validate(s)
}
