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
	// The cloud editing project.
	Project *CreateEditingProjectResponseBodyProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// Id of the request
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
	if s.Project != nil {
		if err := s.Project.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateEditingProjectResponseBodyProject struct {
	// The business configuration of the project. This parameter can be ignored for standard editing projects.
	//
	// example:
	//
	// { "OutputMediaConfig" :    { "StorageLocation": "test-bucket.oss-cn-shanghai.aliyuncs.com", "Path": "test-path"   }, "OutputMediaTarget": "oss-object", "ReservationTime": "2021-06-21T08:05:00Z" }
	BusinessConfig *string `json:"BusinessConfig,omitempty" xml:"BusinessConfig,omitempty"`
	// The business status of the project. This parameter can be ignored for standard editing projects.
	//
	// - Reserving: The live stream is being reserved.
	//
	// - ReservationCanceled: The reservation is canceled.
	//
	// - BroadCasting: The live stream is broadcasting.
	//
	// - LoadingFailed: Loading failed.
	//
	// - LiveFinished: The live stream has ended.
	//
	// example:
	//
	// Reserving
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The template material parameters.
	//
	// example:
	//
	// See the template user guide.
	ClipsParam *string `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	// The cover URL of the cloud editing project.
	//
	// example:
	//
	// http://example-bucket.oss-cn-shanghai.aliyuncs.com/example.png?Expires=<ExpireTime>&OSSAccessKeyId=<OSSAccessKeyId>&Signature=<Signature>&security-token=<SecurityToken>
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The creation source of the cloud editing project.
	//
	// - OpenAPI
	//
	// - AliyunConsole
	//
	// - WebSDK
	//
	// - LiveEditingOpenAPI
	//
	// - LiveEditingConsole
	//
	// example:
	//
	// WebSDK
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// The creation time of the cloud editing project.
	//
	// example:
	//
	// 2021-01-08T16:52:07Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the project.
	//
	// example:
	//
	// example_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The duration of the cloud editing project.
	//
	// example:
	//
	// 3.4200000
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The modification source of the cloud editing project.
	//
	// - OpenAPI
	//
	// - AliyunConsole
	//
	// - WebSDK
	//
	// - LiveEditingOpenAPI
	//
	// - LiveEditingConsole
	//
	// example:
	//
	// WebSDK
	ModifiedSource *string `json:"ModifiedSource,omitempty" xml:"ModifiedSource,omitempty"`
	// The modification time of the cloud editing project.
	//
	// example:
	//
	// 2021-01-08T16:52:07Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the cloud editing project.
	//
	// example:
	//
	// ****01bf24bf41c78b2754cb3187****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the editing project. Default value: EditingProject.
	//
	// - EditingProject: standard editing project.
	//
	// - LiveEditingProject: live editing project.
	//
	// example:
	//
	// LiveEditingProject
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
	// The status of the cloud editing project.
	//
	// Valid values:
	//
	// - 1: Draft.
	//
	// - 2: Editing.
	//
	// - 3: Producing.
	//
	// - 4: Produced.
	//
	// - 5: ProduceFailed.
	//
	// - 7: Deleted.
	//
	// example:
	//
	// 2
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status name of the cloud editing project, corresponding to the status name in the status list.
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
	// The template type of the cloud editing project.
	//
	// - Timeline
	//
	// - VETemplate
	//
	// example:
	//
	// Timeline
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// The timeline of the cloud editing project in JSON format. For more information about the structure, see [TimeLine](~~198823#topic-2024662~~). If this field is empty, an empty timeline is created and the total duration of the cloud editing project is 0.
	//
	// example:
	//
	// {"VideoTracks":[{"VideoTrackClips":[{"MediaId":"****4d7cf14dc7b83b0e801c****"},{"MediaId":"****4d7cf14dc7b83b0e801c****"}]}]}
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// The title of the cloud editing project.
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
