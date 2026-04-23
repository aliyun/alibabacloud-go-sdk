// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchEditingProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjectList(v *SearchEditingProjectResponseBodyProjectList) *SearchEditingProjectResponseBody
	GetProjectList() *SearchEditingProjectResponseBodyProjectList
	SetRequestId(v string) *SearchEditingProjectResponseBody
	GetRequestId() *string
	SetTotal(v int32) *SearchEditingProjectResponseBody
	GetTotal() *int32
}

type SearchEditingProjectResponseBody struct {
	ProjectList *SearchEditingProjectResponseBodyProjectList `json:"ProjectList,omitempty" xml:"ProjectList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 9262E3DA-07FA-48*****62-FCBB6BC61D08
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of online editing projects returned.
	//
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SearchEditingProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponseBody) GetProjectList() *SearchEditingProjectResponseBodyProjectList {
	return s.ProjectList
}

func (s *SearchEditingProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchEditingProjectResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *SearchEditingProjectResponseBody) SetProjectList(v *SearchEditingProjectResponseBodyProjectList) *SearchEditingProjectResponseBody {
	s.ProjectList = v
	return s
}

func (s *SearchEditingProjectResponseBody) SetRequestId(v string) *SearchEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchEditingProjectResponseBody) SetTotal(v int32) *SearchEditingProjectResponseBody {
	s.Total = &v
	return s
}

func (s *SearchEditingProjectResponseBody) Validate() error {
	if s.ProjectList != nil {
		if err := s.ProjectList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchEditingProjectResponseBodyProjectList struct {
	Project []*SearchEditingProjectResponseBodyProjectListProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Repeated"`
}

func (s SearchEditingProjectResponseBodyProjectList) String() string {
	return dara.Prettify(s)
}

func (s SearchEditingProjectResponseBodyProjectList) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponseBodyProjectList) GetProject() []*SearchEditingProjectResponseBodyProjectListProject {
	return s.Project
}

func (s *SearchEditingProjectResponseBodyProjectList) SetProject(v []*SearchEditingProjectResponseBodyProjectListProject) *SearchEditingProjectResponseBodyProjectList {
	s.Project = v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectList) Validate() error {
	if s.Project != nil {
		for _, item := range s.Project {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchEditingProjectResponseBodyProjectListProject struct {
	CoverURL        *string  `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	CreationTime    *string  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description     *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration        *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ModifiedTime    *string  `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	ProjectId       *string  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId        *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status          *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageLocation *string  `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	Title           *string  `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s SearchEditingProjectResponseBodyProjectListProject) String() string {
	return dara.Prettify(s)
}

func (s SearchEditingProjectResponseBodyProjectListProject) GoString() string {
	return s.String()
}

func (s *SearchEditingProjectResponseBodyProjectListProject) GetCoverURL() *string {
	return s.CoverURL
}

func (s *SearchEditingProjectResponseBodyProjectListProject) GetCreationTime() *string {
	return s.CreationTime
}

func (s *SearchEditingProjectResponseBodyProjectListProject) GetDescription() *string {
	return s.Description
}

func (s *SearchEditingProjectResponseBodyProjectListProject) GetDuration() *float32 {
	return s.Duration
}

func (s *SearchEditingProjectResponseBodyProjectListProject) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *SearchEditingProjectResponseBodyProjectListProject) GetProjectId() *string {
	return s.ProjectId
}

func (s *SearchEditingProjectResponseBodyProjectListProject) GetRegionId() *string {
	return s.RegionId
}

func (s *SearchEditingProjectResponseBodyProjectListProject) GetStatus() *string {
	return s.Status
}

func (s *SearchEditingProjectResponseBodyProjectListProject) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *SearchEditingProjectResponseBodyProjectListProject) GetTitle() *string {
	return s.Title
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetCoverURL(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.CoverURL = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetCreationTime(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.CreationTime = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetDescription(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.Description = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetDuration(v float32) *SearchEditingProjectResponseBodyProjectListProject {
	s.Duration = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetModifiedTime(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.ModifiedTime = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetProjectId(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.ProjectId = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetRegionId(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.RegionId = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetStatus(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.Status = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetStorageLocation(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.StorageLocation = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) SetTitle(v string) *SearchEditingProjectResponseBodyProjectListProject {
	s.Title = &v
	return s
}

func (s *SearchEditingProjectResponseBodyProjectListProject) Validate() error {
	return dara.Validate(s)
}
