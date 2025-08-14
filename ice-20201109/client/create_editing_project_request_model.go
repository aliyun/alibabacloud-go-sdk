// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEditingProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessConfig(v string) *CreateEditingProjectRequest
	GetBusinessConfig() *string
	SetClipsParam(v string) *CreateEditingProjectRequest
	GetClipsParam() *string
	SetCoverURL(v string) *CreateEditingProjectRequest
	GetCoverURL() *string
	SetDescription(v string) *CreateEditingProjectRequest
	GetDescription() *string
	SetMaterialMaps(v string) *CreateEditingProjectRequest
	GetMaterialMaps() *string
	SetProjectType(v string) *CreateEditingProjectRequest
	GetProjectType() *string
	SetTemplateId(v string) *CreateEditingProjectRequest
	GetTemplateId() *string
	SetTemplateType(v string) *CreateEditingProjectRequest
	GetTemplateType() *string
	SetTimeline(v string) *CreateEditingProjectRequest
	GetTimeline() *string
	SetTitle(v string) *CreateEditingProjectRequest
	GetTitle() *string
}

type CreateEditingProjectRequest struct {
	// The business configuration of the project. This parameter can be ignored for general editing projects.
	//
	// For a live stream editing project, observe the following rules: OutputMediaConfig.StorageLocation is required. OutputMediaConfig.Path is optional. If you do not specify this option, the live streaming clips are stored in the root directory by default.
	//
	// Valid values of OutputMediaTarget include vod-media and oss-object. If you do not specify OutputMediaTarget, the default value oss-object is used.
	//
	// If you set OutputMediaTarget to vod-media, the setting of OutputMediaConfig.Path does not take effect.
	//
	// example:
	//
	// { "OutputMediaConfig" : { "StorageLocation": "test-bucket.oss-cn-shanghai.aliyuncs.com", "Path": "test-path" }, "OutputMediaTarget": "oss-object", "ReservationTime": "2021-06-21T08:05:00Z" }
	BusinessConfig *string `json:"BusinessConfig,omitempty" xml:"BusinessConfig,omitempty"`
	// The material parameter corresponding to the template, in the JSON format. If TemplateId is specified, ClipsParam must also be specified. For more information<props="china">, see [Create and use a regular template](https://help.aliyun.com/document_detail/328557.html) and [Create and use an advanced template](https://help.aliyun.com/document_detail/291418.html).
	ClipsParam *string `json:"ClipsParam,omitempty" xml:"ClipsParam,omitempty"`
	// The thumbnail URL of the online editing project.
	//
	// example:
	//
	// https://example.com/example.png
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The description of the online editing project.
	//
	// example:
	//
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The material associated with the project. Separate multiple material IDs with commas (,). Each type supports up to 10 material IDs.
	//
	// example:
	//
	// {"video":"*****2e057304fcd9b145c5cafc*****", "image":"****8021a8d493da643c8acd98*****,*****cb6307a4edea614d8b3f3c*****", "liveStream": "[{\\"appName\\":\\"testrecord\\",\\"domainName\\":\\"test.alivecdn.com\\",\\"liveUrl\\":\\"rtmp://test.alivecdn.com/testrecord/teststream\\",\\"streamName\\":\\"teststream\\"}]", "editingProject": "*****9b145c5cafc2e057304fcd*****"}
	MaterialMaps *string `json:"MaterialMaps,omitempty" xml:"MaterialMaps,omitempty"`
	// The type of the editing project. Valid values: EditingProject and LiveEditingProject. A value of EditingProject indicates a regular editing project, and a value of LiveEditingProject indicates a live stream editing project.
	//
	// example:
	//
	// LiveEditingProject
	ProjectType *string `json:"ProjectType,omitempty" xml:"ProjectType,omitempty"`
	// The template ID. This parameter is used to quickly build a timeline with ease. Note: Only one of Timeline and TemplateId can be specified. If TemplateId is specified, ClipsParam must also be specified.
	//
	// example:
	//
	// ****96e8864746a0b6f3****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template type. This parameter is required if you create a template-based online editing project. Default value: Timeline. Valid values:
	//
	// 	- Timeline: a regular template.
	//
	// 	- VETemplate: an advanced template.
	//
	// example:
	//
	// Timeline
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// example:
	//
	// {"VideoTracks":[{"VideoTrackClips":[{"MediaId":"****4d7cf14dc7b83b0e801c****"},{"MediaId":"****4d7cf14dc7b83b0e801c****"}]}]}
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// The title of the online editing project.
	//
	// This parameter is required.
	//
	// example:
	//
	// example
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s CreateEditingProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateEditingProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateEditingProjectRequest) GetBusinessConfig() *string {
	return s.BusinessConfig
}

func (s *CreateEditingProjectRequest) GetClipsParam() *string {
	return s.ClipsParam
}

func (s *CreateEditingProjectRequest) GetCoverURL() *string {
	return s.CoverURL
}

func (s *CreateEditingProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateEditingProjectRequest) GetMaterialMaps() *string {
	return s.MaterialMaps
}

func (s *CreateEditingProjectRequest) GetProjectType() *string {
	return s.ProjectType
}

func (s *CreateEditingProjectRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateEditingProjectRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *CreateEditingProjectRequest) GetTimeline() *string {
	return s.Timeline
}

func (s *CreateEditingProjectRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateEditingProjectRequest) SetBusinessConfig(v string) *CreateEditingProjectRequest {
	s.BusinessConfig = &v
	return s
}

func (s *CreateEditingProjectRequest) SetClipsParam(v string) *CreateEditingProjectRequest {
	s.ClipsParam = &v
	return s
}

func (s *CreateEditingProjectRequest) SetCoverURL(v string) *CreateEditingProjectRequest {
	s.CoverURL = &v
	return s
}

func (s *CreateEditingProjectRequest) SetDescription(v string) *CreateEditingProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateEditingProjectRequest) SetMaterialMaps(v string) *CreateEditingProjectRequest {
	s.MaterialMaps = &v
	return s
}

func (s *CreateEditingProjectRequest) SetProjectType(v string) *CreateEditingProjectRequest {
	s.ProjectType = &v
	return s
}

func (s *CreateEditingProjectRequest) SetTemplateId(v string) *CreateEditingProjectRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateEditingProjectRequest) SetTemplateType(v string) *CreateEditingProjectRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateEditingProjectRequest) SetTimeline(v string) *CreateEditingProjectRequest {
	s.Timeline = &v
	return s
}

func (s *CreateEditingProjectRequest) SetTitle(v string) *CreateEditingProjectRequest {
	s.Title = &v
	return s
}

func (s *CreateEditingProjectRequest) Validate() error {
	return dara.Validate(s)
}
