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
	// The request ID.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
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
	// The business configuration of the project. This parameter can be ignored for general editing projects.
	//
	// example:
	//
	// { "OutputMediaConfig" : { "StorageLocation": "test-bucket.oss-cn-shanghai.aliyuncs.com", "Path": "test-path" }, "OutputMediaTarget": "oss-object", "ReservationTime": "2021-06-21T08:05:00Z" }
	BusinessConfig *string `json:"BusinessConfig,omitempty" xml:"BusinessConfig,omitempty"`
	// The business status of the project. This parameter can be ignored for general editing projects. Valid values:
	//
	// Reserving
	//
	// ReservationCanceled
	//
	// BroadCasting
	//
	// LoadingFailed
	//
	// LiveFinished
	//
	// example:
	//
	// Reserving
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The material parameter corresponding to the template, in the JSON format. If TemplateId is specified, ClipsParam must also be specified. For more information<props="china">, see [Create and use a regular template](https://help.aliyun.com/document_detail/328557.html) and [Create and use an advanced template](https://help.aliyun.com/document_detail/291418.html).
	ClipsParam *string `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	// The thumbnail URL of the online editing project.
	//
	// example:
	//
	// oss://example-bucket/example.jpg
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
	// OpenAPI
	CreateSource *string `json:"CreateSource,omitempty" xml:"CreateSource,omitempty"`
	// The time when the online editing project was created.
	//
	// example:
	//
	// 2020-12-20T12:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the online editing project.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The total duration of the online editing project.
	//
	// example:
	//
	// 24.120000
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
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
	// OpenAPI
	ModifiedSource *string `json:"ModifiedSource,omitempty" xml:"ModifiedSource,omitempty"`
	// The time when the online editing project was last modified.
	//
	// example:
	//
	// 2020-12-20T13:00:00Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The ID of the online editing project.
	//
	// example:
	//
	// ****fb2101bf24b2754cb318787dc****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the editing project. Default value: EditingProject. Valid values:
	//
	// \\- EditingProject: a regular editing project.
	//
	// \\- LiveEditingProject: a live stream editing project.
	//
	// example:
	//
	// EditingProject
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
	// The status of the online editing project. Valid values:
	//
	// \\- Draft
	//
	// \\- Editing
	//
	// \\- Producing
	//
	// \\- Produced
	//
	// \\- ProduceFailed
	//
	// \\- Deleted
	//
	// example:
	//
	// Editing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	// The timeline of the online editing project.
	//
	// example:
	//
	// {"VideoTracks":[{"VideoTrackClips":[{"MediaId":"****9b4d7cf14dc7b83b0e801cbe****"},{"MediaId":"****9b4d7cf14dc7b83b0e801cbe****"},{"MediaId":"****1656bca4474999c961a6d2a2****"}]}]}
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// The error message returned if the project conversion failed. The error message displays the detailed information about the failure, and is returned only if the value of TimelineConvertStatus is ConvertFailed.
	//
	// example:
	//
	// The StorageLocation must be in the same division(apiRegion) as ICE service access point.
	TimelineConvertErrorMessage *string `json:"TimelineConvertErrorMessage,omitempty" xml:"TimelineConvertErrorMessage,omitempty"`
	// The project conversion status. Conversion of an API-style timeline into a frontend-style timeline is an asynchronous process and takes effect only if RequestSource:WebSDK is specified.
	//
	// \\- Unconverted
	//
	// \\- Converting
	//
	// \\- Converted
	//
	// \\- ConvertFailed
	//
	// example:
	//
	// Converted
	TimelineConvertStatus *string `json:"TimelineConvertStatus,omitempty" xml:"TimelineConvertStatus,omitempty"`
	// The title of the online editing project.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetEditingProjectResponseBodyProject) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *GetEditingProjectResponseBodyProject) GetBusinessConfig() *string {
	return s.BusinessConfig
}

func (s *GetEditingProjectResponseBodyProject) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *GetEditingProjectResponseBodyProject) GetClipsParam() *string {
	return s.ClipsParam
}

func (s *GetEditingProjectResponseBodyProject) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetEditingProjectResponseBodyProject) GetCreateSource() *string {
	return s.CreateSource
}

func (s *GetEditingProjectResponseBodyProject) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetEditingProjectResponseBodyProject) GetDescription() *string {
	return s.Description
}

func (s *GetEditingProjectResponseBodyProject) GetDuration() *int64 {
	return s.Duration
}

func (s *GetEditingProjectResponseBodyProject) GetModifiedSource() *string {
	return s.ModifiedSource
}

func (s *GetEditingProjectResponseBodyProject) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetEditingProjectResponseBodyProject) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetEditingProjectResponseBodyProject) GetProjectType() *string {
	return s.ProjectType
}

func (s *GetEditingProjectResponseBodyProject) GetStatus() *string {
	return s.Status
}

func (s *GetEditingProjectResponseBodyProject) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetEditingProjectResponseBodyProject) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetEditingProjectResponseBodyProject) GetTimeline() *string {
	return s.Timeline
}

func (s *GetEditingProjectResponseBodyProject) GetTimelineConvertErrorMessage() *string {
	return s.TimelineConvertErrorMessage
}

func (s *GetEditingProjectResponseBodyProject) GetTimelineConvertStatus() *string {
	return s.TimelineConvertStatus
}

func (s *GetEditingProjectResponseBodyProject) GetTitle() *string {
	return s.Title
}

func (s *GetEditingProjectResponseBodyProject) SetBusinessConfig(v string) *GetEditingProjectResponseBodyProject {
	s.BusinessConfig = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetBusinessStatus(v string) *GetEditingProjectResponseBodyProject {
	s.BusinessStatus = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetClipsParam(v string) *GetEditingProjectResponseBodyProject {
	s.ClipsParam = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetCoverURL(v string) *GetEditingProjectResponseBodyProject {
	s.CoverURL = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetCreateSource(v string) *GetEditingProjectResponseBodyProject {
	s.CreateSource = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetCreateTime(v string) *GetEditingProjectResponseBodyProject {
	s.CreateTime = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetDescription(v string) *GetEditingProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetDuration(v int64) *GetEditingProjectResponseBodyProject {
	s.Duration = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetModifiedSource(v string) *GetEditingProjectResponseBodyProject {
	s.ModifiedSource = &v
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

func (s *GetEditingProjectResponseBodyProject) SetProjectType(v string) *GetEditingProjectResponseBodyProject {
	s.ProjectType = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetStatus(v string) *GetEditingProjectResponseBodyProject {
	s.Status = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetTemplateId(v string) *GetEditingProjectResponseBodyProject {
	s.TemplateId = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetTemplateType(v string) *GetEditingProjectResponseBodyProject {
	s.TemplateType = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetTimeline(v string) *GetEditingProjectResponseBodyProject {
	s.Timeline = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetTimelineConvertErrorMessage(v string) *GetEditingProjectResponseBodyProject {
	s.TimelineConvertErrorMessage = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetTimelineConvertStatus(v string) *GetEditingProjectResponseBodyProject {
	s.TimelineConvertStatus = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) SetTitle(v string) *GetEditingProjectResponseBodyProject {
	s.Title = &v
	return s
}

func (s *GetEditingProjectResponseBodyProject) Validate() error {
	return dara.Validate(s)
}
