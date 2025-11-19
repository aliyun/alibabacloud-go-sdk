// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEditingProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProject(v *GetEditingProjectResponseBodyProject) *GetEditingProjectResponseBody
	GetProject() *GetEditingProjectResponseBodyProject
	SetRequestId(v string) *GetEditingProjectResponseBody
	GetRequestId() *string
}

type GetEditingProjectResponseBody struct {
	// The information about the online editing project.
	Project *GetEditingProjectResponseBodyProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 63E8B7C7-4812-46*****AD-0FA56029AC86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEditingProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetEditingProjectResponseBody) GetProject() *GetEditingProjectResponseBodyProject {
	return s.Project
}

func (s *GetEditingProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEditingProjectResponseBody) SetProject(v *GetEditingProjectResponseBodyProject) *GetEditingProjectResponseBody {
	s.Project = v
	return s
}

func (s *GetEditingProjectResponseBody) SetRequestId(v string) *GetEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEditingProjectResponseBody) Validate() error {
	if s.Project != nil {
		if err := s.Project.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEditingProjectResponseBodyProject struct {
	// The thumbnail URL of the online editing project.
	//
	// example:
	//
	// https://****.com/6AB4D0E1E1C74468883516C2349****.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the online editing project was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-10-23T13:33:40Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the online editing project.
	//
	// example:
	//
	// testdescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The last time when the online editing project was modified. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-10-23T14:27:26Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the online editing project.
	//
	// example:
	//
	// fb2101bf24b27*****54cb318787dc
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
	// The path of the Object Storage Service (OSS) bucket where the online editing project is stored.
	//
	// > To view the path of the OSS bucket, log on to the [ApsaraVideo VOD console](https://vod.console.aliyun.com/?spm=a2c4g.11186623.2.15.6948257eaZ4m54#/vod/settings/censored), and choose **Configuration Management*	- > **Media Management*	- > **Storage**. On the Storage page, you can view the path of the OSS bucket.
	//
	// example:
	//
	// location_s
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The timeline of the online editing project.
	//
	// example:
	//
	// {\\"TimelineIn\\":0,\\"TimelineOut\\":9.42}
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// The title of the online editing project.
	//
	// example:
	//
	// video_1508736815000
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetEditingProjectResponseBodyProject) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *GetEditingProjectResponseBodyProject) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetEditingProjectResponseBodyProject) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetEditingProjectResponseBodyProject) GetDescription() *string {
	return s.Description
}

func (s *GetEditingProjectResponseBodyProject) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetEditingProjectResponseBodyProject) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetEditingProjectResponseBodyProject) GetRegionId() *string {
	return s.RegionId
}

func (s *GetEditingProjectResponseBodyProject) GetStatus() *string {
	return s.Status
}

func (s *GetEditingProjectResponseBodyProject) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *GetEditingProjectResponseBodyProject) GetTimeline() *string {
	return s.Timeline
}

func (s *GetEditingProjectResponseBodyProject) GetTitle() *string {
	return s.Title
}

func (s *GetEditingProjectResponseBodyProject) SetCoverURL(v string) *GetEditingProjectResponseBodyProject {
	s.CoverURL = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetCreationTime(v string) *GetEditingProjectResponseBodyProject {
	s.CreationTime = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetDescription(v string) *GetEditingProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetModifiedTime(v string) *GetEditingProjectResponseBodyProject {
	s.ModifiedTime = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetProjectId(v string) *GetEditingProjectResponseBodyProject {
	s.ProjectId = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetRegionId(v string) *GetEditingProjectResponseBodyProject {
	s.RegionId = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetStatus(v string) *GetEditingProjectResponseBodyProject {
	s.Status = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetStorageLocation(v string) *GetEditingProjectResponseBodyProject {
	s.StorageLocation = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetTimeline(v string) *GetEditingProjectResponseBodyProject {
	s.Timeline = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetTitle(v string) *GetEditingProjectResponseBodyProject {
	s.Title = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) Validate() error {
	return dara.Validate(s)
}
