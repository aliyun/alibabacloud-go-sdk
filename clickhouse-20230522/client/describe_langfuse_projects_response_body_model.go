// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeLangfuseProjectsResponseBodyData) *DescribeLangfuseProjectsResponseBody
	GetData() *DescribeLangfuseProjectsResponseBodyData
	SetRequestId(v string) *DescribeLangfuseProjectsResponseBody
	GetRequestId() *string
}

type DescribeLangfuseProjectsResponseBody struct {
	// The returned result.
	Data *DescribeLangfuseProjectsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD8660E9CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLangfuseProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseProjectsResponseBody) GetData() *DescribeLangfuseProjectsResponseBodyData {
	return s.Data
}

func (s *DescribeLangfuseProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLangfuseProjectsResponseBody) SetData(v *DescribeLangfuseProjectsResponseBodyData) *DescribeLangfuseProjectsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLangfuseProjectsResponseBody) SetRequestId(v string) *DescribeLangfuseProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLangfuseProjectsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLangfuseProjectsResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 30
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of Langfuse projects.
	Projects []*DescribeLangfuseProjectsResponseBodyDataProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeLangfuseProjectsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseProjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseProjectsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeLangfuseProjectsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeLangfuseProjectsResponseBodyData) GetProjects() []*DescribeLangfuseProjectsResponseBodyDataProjects {
	return s.Projects
}

func (s *DescribeLangfuseProjectsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeLangfuseProjectsResponseBodyData) SetPageNumber(v int64) *DescribeLangfuseProjectsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeLangfuseProjectsResponseBodyData) SetPageSize(v int64) *DescribeLangfuseProjectsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeLangfuseProjectsResponseBodyData) SetProjects(v []*DescribeLangfuseProjectsResponseBodyDataProjects) *DescribeLangfuseProjectsResponseBodyData {
	s.Projects = v
	return s
}

func (s *DescribeLangfuseProjectsResponseBodyData) SetTotalCount(v int64) *DescribeLangfuseProjectsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeLangfuseProjectsResponseBodyData) Validate() error {
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

type DescribeLangfuseProjectsResponseBodyDataProjects struct {
	// The time when the Langfuse project was created.
	//
	// example:
	//
	// 2026-05-27T08:23:45Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The Langfuse project name.
	//
	// example:
	//
	// project_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The organization ID to which the Langfuse project belongs.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	OrganizationId *string `json:"OrganizationId,omitempty" xml:"OrganizationId,omitempty"`
	// The Langfuse project ID.
	//
	// example:
	//
	// cmrbhzx930005jw2q****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The time when the Langfuse project was last updated.
	//
	// example:
	//
	// 2026-06-09T10:27:35
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s DescribeLangfuseProjectsResponseBodyDataProjects) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseProjectsResponseBodyDataProjects) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseProjectsResponseBodyDataProjects) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *DescribeLangfuseProjectsResponseBodyDataProjects) GetName() *string {
	return s.Name
}

func (s *DescribeLangfuseProjectsResponseBodyDataProjects) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DescribeLangfuseProjectsResponseBodyDataProjects) GetProjectId() *string {
	return s.ProjectId
}

func (s *DescribeLangfuseProjectsResponseBodyDataProjects) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *DescribeLangfuseProjectsResponseBodyDataProjects) SetCreatedAt(v string) *DescribeLangfuseProjectsResponseBodyDataProjects {
	s.CreatedAt = &v
	return s
}

func (s *DescribeLangfuseProjectsResponseBodyDataProjects) SetName(v string) *DescribeLangfuseProjectsResponseBodyDataProjects {
	s.Name = &v
	return s
}

func (s *DescribeLangfuseProjectsResponseBodyDataProjects) SetOrganizationId(v string) *DescribeLangfuseProjectsResponseBodyDataProjects {
	s.OrganizationId = &v
	return s
}

func (s *DescribeLangfuseProjectsResponseBodyDataProjects) SetProjectId(v string) *DescribeLangfuseProjectsResponseBodyDataProjects {
	s.ProjectId = &v
	return s
}

func (s *DescribeLangfuseProjectsResponseBodyDataProjects) SetUpdatedAt(v string) *DescribeLangfuseProjectsResponseBodyDataProjects {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeLangfuseProjectsResponseBodyDataProjects) Validate() error {
	return dara.Validate(s)
}
