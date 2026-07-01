// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitMediaCensorJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBarrages(v string) *SubmitMediaCensorJobRequest
	GetBarrages() *string
	SetCoverImages(v string) *SubmitMediaCensorJobRequest
	GetCoverImages() *string
	SetDescription(v string) *SubmitMediaCensorJobRequest
	GetDescription() *string
	SetInput(v *SubmitMediaCensorJobRequestInput) *SubmitMediaCensorJobRequest
	GetInput() *SubmitMediaCensorJobRequestInput
	SetNotifyUrl(v string) *SubmitMediaCensorJobRequest
	GetNotifyUrl() *string
	SetOutput(v string) *SubmitMediaCensorJobRequest
	GetOutput() *string
	SetScheduleConfig(v *SubmitMediaCensorJobRequestScheduleConfig) *SubmitMediaCensorJobRequest
	GetScheduleConfig() *SubmitMediaCensorJobRequestScheduleConfig
	SetTemplateId(v string) *SubmitMediaCensorJobRequest
	GetTemplateId() *string
	SetTitle(v string) *SubmitMediaCensorJobRequest
	GetTitle() *string
	SetUserData(v string) *SubmitMediaCensorJobRequest
	GetUserData() *string
}

type SubmitMediaCensorJobRequest struct {
	// The video barrages (on-screen comments).
	//
	// > If specified, it overrides the barrages specified in the Media object.
	//
	// example:
	//
	// hello world
	Barrages *string `json:"Barrages,omitempty" xml:"Barrages,omitempty"`
	// The Object Storage Service (OSS) files for the cover images, specified as a JSON array. You can specify up to five cover images.
	//
	// > If specified, this parameter overrides the cover image information in the **Media*	- object.
	//
	// example:
	//
	// [{"Bucket":"example-bucket-****","Location":"oss-cn-shanghai","Object":"example-****.jpeg","RoleArn":"acs:ram::1997018457688683:role/AliyunICEDefaultRole"}]
	CoverImages *string `json:"CoverImages,omitempty" xml:"CoverImages,omitempty"`
	// The video description. The maximum length is 128 bytes.
	//
	// > If specified, this parameter overrides the description specified in the Media object.
	//
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The input file to censor.
	Input *SubmitMediaCensorJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The callback path. Both Message Service (MNS) and HTTP callbacks are supported.
	//
	// example:
	//
	// mns://125340688170****.oss-cn-shanghai.aliyuncs.com/queues/example-pipeline
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The output location for screenshots. The censor job generates screenshots and a result JSON file in the OSS location specified by this parameter.
	//
	// - Example format: `oss://bucket/snapshot-{Count}.jpg`, where `bucket` is the name of an OSS bucket in the same region as the project, and `{Count}` is a placeholder for the screenshot sequence number.
	//
	// - The detailed censor results are saved to a file named `{jobId}.output` in the same OSS folder as the value of `Output`. For information about the fields in the output file, see [Media censor result file fields](https://help.aliyun.com/document_detail/609211.html).
	//
	// example:
	//
	// oss://sashimi-cn-shanghai/censor/snapshot-{Count}.jpg
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The scheduling configuration.
	ScheduleConfig *SubmitMediaCensorJobRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The template ID. If this parameter is left empty, the service uses the default template for the censor job.
	//
	// example:
	//
	// S00000001-100060
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The video title. The maximum length is 64 bytes.
	//
	// > If specified, this parameter overrides the title specified in the Media object.
	//
	// example:
	//
	// Hello World
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The user-defined data. The maximum length is 128 bytes.
	//
	// example:
	//
	// UserDatatestid-001-****
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitMediaCensorJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaCensorJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitMediaCensorJobRequest) GetBarrages() *string {
	return s.Barrages
}

func (s *SubmitMediaCensorJobRequest) GetCoverImages() *string {
	return s.CoverImages
}

func (s *SubmitMediaCensorJobRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitMediaCensorJobRequest) GetInput() *SubmitMediaCensorJobRequestInput {
	return s.Input
}

func (s *SubmitMediaCensorJobRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *SubmitMediaCensorJobRequest) GetOutput() *string {
	return s.Output
}

func (s *SubmitMediaCensorJobRequest) GetScheduleConfig() *SubmitMediaCensorJobRequestScheduleConfig {
	return s.ScheduleConfig
}

func (s *SubmitMediaCensorJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitMediaCensorJobRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitMediaCensorJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitMediaCensorJobRequest) SetBarrages(v string) *SubmitMediaCensorJobRequest {
	s.Barrages = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetCoverImages(v string) *SubmitMediaCensorJobRequest {
	s.CoverImages = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetDescription(v string) *SubmitMediaCensorJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetInput(v *SubmitMediaCensorJobRequestInput) *SubmitMediaCensorJobRequest {
	s.Input = v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetNotifyUrl(v string) *SubmitMediaCensorJobRequest {
	s.NotifyUrl = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetOutput(v string) *SubmitMediaCensorJobRequest {
	s.Output = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetScheduleConfig(v *SubmitMediaCensorJobRequestScheduleConfig) *SubmitMediaCensorJobRequest {
	s.ScheduleConfig = v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetTemplateId(v string) *SubmitMediaCensorJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetTitle(v string) *SubmitMediaCensorJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) SetUserData(v string) *SubmitMediaCensorJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitMediaCensorJobRequest) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduleConfig != nil {
		if err := s.ScheduleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitMediaCensorJobRequestInput struct {
	// The identifier for the input file. You can specify either an OSS URL or a media ID.
	//
	// Valid OSS URL formats:
	//
	// 1\\. `oss://bucket/object`
	//
	// 2\\. `http(s)://bucket.oss-[regionId].aliyuncs.com/object`
	//
	// The `bucket` must be in the same region as the project, and `object` is the path to the file.
	//
	// example:
	//
	// 1b1b9cd148034739af413150fded****
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the input file. Valid values:
	//
	// `OSS`: an OSS URL
	//
	// `Media`: a media ID
	//
	// example:
	//
	// Media
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitMediaCensorJobRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaCensorJobRequestInput) GoString() string {
	return s.String()
}

func (s *SubmitMediaCensorJobRequestInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitMediaCensorJobRequestInput) GetType() *string {
	return s.Type
}

func (s *SubmitMediaCensorJobRequestInput) SetMedia(v string) *SubmitMediaCensorJobRequestInput {
	s.Media = &v
	return s
}

func (s *SubmitMediaCensorJobRequestInput) SetType(v string) *SubmitMediaCensorJobRequestInput {
	s.Type = &v
	return s
}

func (s *SubmitMediaCensorJobRequestInput) Validate() error {
	return dara.Validate(s)
}

type SubmitMediaCensorJobRequestScheduleConfig struct {
	// The pipeline ID.
	//
	// example:
	//
	// 5246b8d12a62433ab77845074039****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The task priority. A higher value indicates a higher priority. Valid values range from 1 to 10.
	//
	// example:
	//
	// 6
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s SubmitMediaCensorJobRequestScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitMediaCensorJobRequestScheduleConfig) GoString() string {
	return s.String()
}

func (s *SubmitMediaCensorJobRequestScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitMediaCensorJobRequestScheduleConfig) GetPriority() *int32 {
	return s.Priority
}

func (s *SubmitMediaCensorJobRequestScheduleConfig) SetPipelineId(v string) *SubmitMediaCensorJobRequestScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *SubmitMediaCensorJobRequestScheduleConfig) SetPriority(v int32) *SubmitMediaCensorJobRequestScheduleConfig {
	s.Priority = &v
	return s
}

func (s *SubmitMediaCensorJobRequestScheduleConfig) Validate() error {
	return dara.Validate(s)
}
