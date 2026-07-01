// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmarttagJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SubmitSmarttagJobRequest
	GetContent() *string
	SetContentAddr(v string) *SubmitSmarttagJobRequest
	GetContentAddr() *string
	SetContentType(v string) *SubmitSmarttagJobRequest
	GetContentType() *string
	SetInput(v *SubmitSmarttagJobRequestInput) *SubmitSmarttagJobRequest
	GetInput() *SubmitSmarttagJobRequestInput
	SetNotifyUrl(v string) *SubmitSmarttagJobRequest
	GetNotifyUrl() *string
	SetParams(v string) *SubmitSmarttagJobRequest
	GetParams() *string
	SetScheduleConfig(v *SubmitSmarttagJobRequestScheduleConfig) *SubmitSmarttagJobRequest
	GetScheduleConfig() *SubmitSmarttagJobRequestScheduleConfig
	SetTemplateConfig(v string) *SubmitSmarttagJobRequest
	GetTemplateConfig() *string
	SetTemplateId(v string) *SubmitSmarttagJobRequest
	GetTemplateId() *string
	SetTitle(v string) *SubmitSmarttagJobRequest
	GetTitle() *string
	SetUserData(v string) *SubmitSmarttagJobRequest
	GetUserData() *string
}

type SubmitSmarttagJobRequest struct {
	// The description of the video content can contain Chinese characters, English letters, digits, and hyphens (-). It cannot start with a special character and must not exceed 1 KB.
	//
	// example:
	//
	// example content ****
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Deprecated.
	ContentAddr *string `json:"ContentAddr,omitempty" xml:"ContentAddr,omitempty"`
	// Deprecated.
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The input file for the job.
	Input *SubmitSmarttagJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The callback URL. Only HTTP and HTTPS URLs are supported.
	//
	// example:
	//
	// https://example.com/endpoint/aliyun/ai?id=76401125000***
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// Additional request parameters, specified as a JSON string. For example: `{"needAsrData":true, "needOcrData":false}`.
	//
	// - `needAsrData`: Specifies whether to include the raw Automatic Speech Recognition (ASR) results in the analysis output. The default is `false`.
	//
	// - `needOcrData`: Specifies whether to include the raw Optical Character Recognition (OCR) results in the analysis output. The default is `false`.
	//
	// - `needMetaData`: Specifies whether to include metadata in the analysis output. The default is `false`.
	//
	// - `nlpParams`: A JSON object that specifies the input parameters for the Natural Language Processing (NLP) operator. If left empty, the operator is not used. For details, see the `nlpParams` table below.
	//
	// example:
	//
	// {"needAsrData":true, "needOcrData":false, "nlpParams":{"sourceLanguage":"cn"}}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The scheduling configurations.
	ScheduleConfig *SubmitSmarttagJobRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// Dynamic parameters for the job, which temporarily override or supplement the base template specified by `TemplateId`. The service merges the dynamic and template parameters to generate the final configuration for the current job and validates it before execution.
	//
	// - Merge rules:
	//
	// 1. Values in the request override corresponding values in the template.
	//
	// 2. Fields in the request that do not exist in the template are added to the configuration.
	//
	// - Currently supported dynamic fields:
	//
	// 1. `FaceCategoryIds`: A list of face library IDs for recognition, separated by commas (,). You can include both system and custom library IDs.
	//
	// - Note: These dynamic parameters affect only the current job and do not modify the template itself.
	//
	// example:
	//
	// {"FaceCategoryIds":"custom_face_lib1"}
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The ID of the template that specifies the analysis algorithms to use.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The video title can contain Chinese characters, English letters, digits, and hyphens (-). It cannot start with a special character and must not exceed 256 bytes.
	//
	// example:
	//
	// example-title-****
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// Custom data to include in the callback. If you use Message Service (MNS) for callbacks, this data is included in the message. The maximum length is 1 KB.
	//
	// example:
	//
	// {“a”:"test"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitSmarttagJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmarttagJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmarttagJobRequest) GetContent() *string {
	return s.Content
}

func (s *SubmitSmarttagJobRequest) GetContentAddr() *string {
	return s.ContentAddr
}

func (s *SubmitSmarttagJobRequest) GetContentType() *string {
	return s.ContentType
}

func (s *SubmitSmarttagJobRequest) GetInput() *SubmitSmarttagJobRequestInput {
	return s.Input
}

func (s *SubmitSmarttagJobRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *SubmitSmarttagJobRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitSmarttagJobRequest) GetScheduleConfig() *SubmitSmarttagJobRequestScheduleConfig {
	return s.ScheduleConfig
}

func (s *SubmitSmarttagJobRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *SubmitSmarttagJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitSmarttagJobRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitSmarttagJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSmarttagJobRequest) SetContent(v string) *SubmitSmarttagJobRequest {
	s.Content = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetContentAddr(v string) *SubmitSmarttagJobRequest {
	s.ContentAddr = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetContentType(v string) *SubmitSmarttagJobRequest {
	s.ContentType = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetInput(v *SubmitSmarttagJobRequestInput) *SubmitSmarttagJobRequest {
	s.Input = v
	return s
}

func (s *SubmitSmarttagJobRequest) SetNotifyUrl(v string) *SubmitSmarttagJobRequest {
	s.NotifyUrl = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetParams(v string) *SubmitSmarttagJobRequest {
	s.Params = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetScheduleConfig(v *SubmitSmarttagJobRequestScheduleConfig) *SubmitSmarttagJobRequest {
	s.ScheduleConfig = v
	return s
}

func (s *SubmitSmarttagJobRequest) SetTemplateConfig(v string) *SubmitSmarttagJobRequest {
	s.TemplateConfig = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetTemplateId(v string) *SubmitSmarttagJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetTitle(v string) *SubmitSmarttagJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitSmarttagJobRequest) SetUserData(v string) *SubmitSmarttagJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSmarttagJobRequest) Validate() error {
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

type SubmitSmarttagJobRequestInput struct {
	// - If you set the `Type` parameter to `OSS`, specify the OSS URL of the media file. Example: `OSS://test-bucket/video/202208/test.mp4`.
	//
	// - If you set the `Type` parameter to `Media`, specify the media ID. Example: `c5c62d8f0361337cab312dce8e77dc6d`.
	//
	// - If you set the `Type` parameter to `URL`, specify the HTTP or HTTPS URL of the media file. Example: `https://zc-test.oss-cn-shanghai.aliyuncs.com/test/unknowFace.mp4`.
	//
	// example:
	//
	// c5c62d8f0361337cab312dce8e77dc6d
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the input media file. Valid values:
	//
	// - OSS
	//
	// - Media
	//
	// - URL
	//
	// example:
	//
	// Media
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SubmitSmarttagJobRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmarttagJobRequestInput) GoString() string {
	return s.String()
}

func (s *SubmitSmarttagJobRequestInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitSmarttagJobRequestInput) GetType() *string {
	return s.Type
}

func (s *SubmitSmarttagJobRequestInput) SetMedia(v string) *SubmitSmarttagJobRequestInput {
	s.Media = &v
	return s
}

func (s *SubmitSmarttagJobRequestInput) SetType(v string) *SubmitSmarttagJobRequestInput {
	s.Type = &v
	return s
}

func (s *SubmitSmarttagJobRequestInput) Validate() error {
	return dara.Validate(s)
}

type SubmitSmarttagJobRequestScheduleConfig struct {
	// The ID of the pipeline. Pipelines separate business workloads and bind message notifications.
	//
	// If you do not specify this parameter, the default pipeline is used. The default pipeline has a concurrency of 2. To increase the concurrency, submit a ticket.
	//
	// example:
	//
	// acdbfe4323bcfdae
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The priority of the job. This feature is not yet implemented. You can leave this parameter empty or specify any value.
	//
	// example:
	//
	// 4
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s SubmitSmarttagJobRequestScheduleConfig) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmarttagJobRequestScheduleConfig) GoString() string {
	return s.String()
}

func (s *SubmitSmarttagJobRequestScheduleConfig) GetPipelineId() *string {
	return s.PipelineId
}

func (s *SubmitSmarttagJobRequestScheduleConfig) GetPriority() *string {
	return s.Priority
}

func (s *SubmitSmarttagJobRequestScheduleConfig) SetPipelineId(v string) *SubmitSmarttagJobRequestScheduleConfig {
	s.PipelineId = &v
	return s
}

func (s *SubmitSmarttagJobRequestScheduleConfig) SetPriority(v string) *SubmitSmarttagJobRequestScheduleConfig {
	s.Priority = &v
	return s
}

func (s *SubmitSmarttagJobRequestScheduleConfig) Validate() error {
	return dara.Validate(s)
}
