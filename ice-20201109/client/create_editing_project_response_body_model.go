// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEditingProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProject(v *CreateEditingProjectResponseBodyProject) *CreateEditingProjectResponseBody
	GetProject() *CreateEditingProjectResponseBodyProject
	SetRequestId(v string) *CreateEditingProjectResponseBody
	GetRequestId() *string
}

type CreateEditingProjectResponseBody struct {
	// The information about the online editing project.
	Project *CreateEditingProjectResponseBodyProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEditingProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEditingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEditingProjectResponseBody) GetProject() *CreateEditingProjectResponseBodyProject {
	return s.Project
}

func (s *CreateEditingProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEditingProjectResponseBody) SetProject(v *CreateEditingProjectResponseBodyProject) *CreateEditingProjectResponseBody {
	s.Project = v
	return s
}

func (s *CreateEditingProjectResponseBody) SetRequestId(v string) *CreateEditingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEditingProjectResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateEditingProjectResponseBodyProject struct {
	// The business configuration of the project. This parameter can be ignored for general editing projects.
	//
	// example:
	//
	// { "OutputMediaConfig" :    { "StorageLocation": "test-bucket.oss-cn-shanghai.aliyuncs.com", "Path": "test-path"   }, "OutputMediaTarget": "oss-object", "ReservationTime": "2021-06-21T08:05:00Z" }
	BusinessConfig *string `json:"BusinessConfig,omitempty" xml:"BusinessConfig,omitempty"`
	// The business status of the project. This parameter can be ignored for general editing projects. Valid values:
	//
	// 	- Reserving
	//
	// 	- ReservationCanceled
	//
	// 	- BroadCasting
	//
	// 	- LoadingFailed
	//
	// 	- LiveFinished
	//
	// example:
	//
	// Reserving
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The template material parameters.
	ClipsParam *string `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	// The thumbnail URL of the online editing project.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.png?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The method for creating the online editing project. Valid values:
	//
	// \\- OpenAPI
	//
	// \\- AliyunConsole
	//
	// \\- WebSDK
	//
	// \\- LiveEditingOpenAPI
	//
	// \\- LiveEditingConsole
	//
	// example:
	//
	// WebSDK
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// The time when the online editing project was created.
	//
	// example:
	//
	// 2021-01-08T16:52:07Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the online editing project.
	//
	// example:
	//
	// example_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The duration of the online editing project.
	//
	// example:
	//
	// 3.4200000
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The method for editing the online editing project. Valid values:
	//
	// \\- OpenAPI
	//
	// \\- AliyunConsole
	//
	// \\- WebSDK
	//
	// \\- LiveEditingOpenAPI
	//
	// \\- LiveEditingConsole
	//
	// example:
	//
	// WebSDK
	ModifiedSource *string `json:"ModifiedSource,omitempty" xml:"ModifiedSource,omitempty"`
	// The time when the online editing project was last edited.
	//
	// example:
	//
	// 2021-01-08T16:52:07Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the online editing project.
	//
	// example:
	//
	// ****01bf24bf41c78b2754cb3187****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the editing project. Default value: EditingProject. Valid values:
	//
	// \\- EditingProject: a regular editing project.
	//
	// \\- LiveEditingProject: a live stream editing project.
	//
	// example:
	//
	// LiveEditingProject
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
	// The status of the online editing project.
	//
	// Valid values:
	//
	// \\- 1: Draft
	//
	// \\- 2: Editing
	//
	// \\- 3: Producing
	//
	// \\- 4: Produced
	//
	// \\- 5: ProduceFailed
	//
	// \\- 7: Deleted
	//
	// example:
	//
	// 2
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status of the online editing project. For more information, see the status list.
	//
	// example:
	//
	// Editing
	StatusName *string `json:"StatusName,omitempty" xml:"StatusName,omitempty"`
	// The template ID.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template type of the online editing project. Valid values:
	//
	// \\- Timeline
	//
	// \\- VETemplate
	//
	// example:
	//
	// Timeline
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The timeline of the online editing project, in the JSON format.<props="china">For more information about objects in a timeline, see [Timeline configurations](https://help.aliyun.com/document_detail/198823.htm?spm=a2c4g.11186623.2.9.90dc653dF67srN#topic-2024662).  If you leave this parameter empty, an empty timeline is created and the duration of the online editing project is zero.
	//
	// example:
	//
	// {"VideoTracks":[{"VideoTrackClips":[{"MediaId":"****4d7cf14dc7b83b0e801c****"},{"MediaId":"****4d7cf14dc7b83b0e801c****"}]}]}
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// The title of the online editing project.
	//
	// example:
	//
	// example_title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateEditingProjectResponseBodyProject) String() string {
	return dara.Prettify(s)
}

func (s CreateEditingProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *CreateEditingProjectResponseBodyProject) GetBusinessConfig() *string {
	return s.BusinessConfig
}

func (s *CreateEditingProjectResponseBodyProject) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *CreateEditingProjectResponseBodyProject) GetClipsParam() *string {
	return s.ClipsParam
}

func (s *CreateEditingProjectResponseBodyProject) GetCoverURL() *string {
	return s.CoverURL
}

func (s *CreateEditingProjectResponseBodyProject) GetCreateSource() *string {
	return s.CreateSource
}

func (s *CreateEditingProjectResponseBodyProject) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateEditingProjectResponseBodyProject) GetDescription() *string {
	return s.Description
}

func (s *CreateEditingProjectResponseBodyProject) GetDuration() *float32 {
	return s.Duration
}

func (s *CreateEditingProjectResponseBodyProject) GetModifiedSource() *string {
	return s.ModifiedSource
}

func (s *CreateEditingProjectResponseBodyProject) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *CreateEditingProjectResponseBodyProject) GetProjectId() *string {
	return s.ProjectId
}

func (s *CreateEditingProjectResponseBodyProject) GetProjectType() *string {
	return s.ProjectType
}

func (s *CreateEditingProjectResponseBodyProject) GetStatus() *int64 {
	return s.Status
}

func (s *CreateEditingProjectResponseBodyProject) GetStatusName() *string {
	return s.StatusName
}

func (s *CreateEditingProjectResponseBodyProject) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateEditingProjectResponseBodyProject) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateEditingProjectResponseBodyProject) GetTimeline() *string {
	return s.Timeline
}

func (s *CreateEditingProjectResponseBodyProject) GetTitle() *string {
	return s.Title
}

func (s *CreateEditingProjectResponseBodyProject) SetBusinessConfig(v string) *CreateEditingProjectResponseBodyProject {
	s.BusinessConfig = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetBusinessStatus(v string) *CreateEditingProjectResponseBodyProject {
	s.BusinessStatus = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetClipsParam(v string) *CreateEditingProjectResponseBodyProject {
	s.ClipsParam = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetCoverURL(v string) *CreateEditingProjectResponseBodyProject {
	s.CoverURL = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetCreateSource(v string) *CreateEditingProjectResponseBodyProject {
	s.CreateSource = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetCreateTime(v string) *CreateEditingProjectResponseBodyProject {
	s.CreateTime = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetDescription(v string) *CreateEditingProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetDuration(v float32) *CreateEditingProjectResponseBodyProject {
	s.Duration = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetModifiedSource(v string) *CreateEditingProjectResponseBodyProject {
	s.ModifiedSource = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetModifiedTime(v string) *CreateEditingProjectResponseBodyProject {
	s.ModifiedTime = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetProjectId(v string) *CreateEditingProjectResponseBodyProject {
	s.ProjectId = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetProjectType(v string) *CreateEditingProjectResponseBodyProject {
	s.ProjectType = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetStatus(v int64) *CreateEditingProjectResponseBodyProject {
	s.Status = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetStatusName(v string) *CreateEditingProjectResponseBodyProject {
	s.StatusName = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetTemplateId(v string) *CreateEditingProjectResponseBodyProject {
	s.TemplateId = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetTemplateType(v string) *CreateEditingProjectResponseBodyProject {
	s.TemplateType = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetTimeline(v string) *CreateEditingProjectResponseBodyProject {
	s.Timeline = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) SetTitle(v string) *CreateEditingProjectResponseBodyProject {
	s.Title = &v
	return s
}

func (s *CreateEditingProjectResponseBodyProject) Validate() error {
	return dara.Validate(s)
}
