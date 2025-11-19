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
	// The list of online editing projects.
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
	// The thumbnail URL of the online editing project.
	//
	// example:
	//
	// cover_url
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the online editing project was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the online editing project.
	//
	// example:
	//
	// test project 001
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The duration of the online editing project, which must be consistent with the duration of the timeline.
	//
	// > The Timeline parameter is not included in response parameters.
	//
	// example:
	//
	// 22.65
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The last time when the online editing project was modified. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-01-11T13:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the online editing project.
	//
	// example:
	//
	// 25cfc178d2de4*****e77aebed6afcd
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region where the online editing project was created.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the online editing project. Separate multiple states with commas (,). By default, all online editing projects were queried. Valid values:
	//
	// 	- **Normal**: indicates that the online editing project is in draft.
	//
	// 	- **Producing**: indicates that the video is being produced.
	//
	// 	- **Produced**: indicates that the video was produced.
	//
	// 	- **ProduceFailed**: indicates that the video failed to be produced.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The path of the Object Storage Service (OSS) bucket where the produced video is stored.
	//
	// > To view the path of the OSS bucket, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com/?spm=a2c4g.11186623.2.15.6948257eaZ4m54#/vod/settings/censored), and choose **Configuration Management*	- > **Media Management*	- > **Storage**. On the Storage page, you can view the path of the OSS bucket.
	//
	// example:
	//
	// location_s
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The title of the online editing project.
	//
	// example:
	//
	// video_150873681****
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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
