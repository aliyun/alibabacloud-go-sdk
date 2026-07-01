// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSmarttagJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SubmitSmarttagJobShrinkRequest
	GetContent() *string
	SetContentAddr(v string) *SubmitSmarttagJobShrinkRequest
	GetContentAddr() *string
	SetContentType(v string) *SubmitSmarttagJobShrinkRequest
	GetContentType() *string
	SetInputShrink(v string) *SubmitSmarttagJobShrinkRequest
	GetInputShrink() *string
	SetNotifyUrl(v string) *SubmitSmarttagJobShrinkRequest
	GetNotifyUrl() *string
	SetParams(v string) *SubmitSmarttagJobShrinkRequest
	GetParams() *string
	SetScheduleConfigShrink(v string) *SubmitSmarttagJobShrinkRequest
	GetScheduleConfigShrink() *string
	SetTemplateConfig(v string) *SubmitSmarttagJobShrinkRequest
	GetTemplateConfig() *string
	SetTemplateId(v string) *SubmitSmarttagJobShrinkRequest
	GetTemplateId() *string
	SetTitle(v string) *SubmitSmarttagJobShrinkRequest
	GetTitle() *string
	SetUserData(v string) *SubmitSmarttagJobShrinkRequest
	GetUserData() *string
}

type SubmitSmarttagJobShrinkRequest struct {
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
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
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
	ScheduleConfigShrink *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
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

func (s SubmitSmarttagJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSmarttagJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitSmarttagJobShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *SubmitSmarttagJobShrinkRequest) GetContentAddr() *string {
	return s.ContentAddr
}

func (s *SubmitSmarttagJobShrinkRequest) GetContentType() *string {
	return s.ContentType
}

func (s *SubmitSmarttagJobShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *SubmitSmarttagJobShrinkRequest) GetNotifyUrl() *string {
	return s.NotifyUrl
}

func (s *SubmitSmarttagJobShrinkRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitSmarttagJobShrinkRequest) GetScheduleConfigShrink() *string {
	return s.ScheduleConfigShrink
}

func (s *SubmitSmarttagJobShrinkRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *SubmitSmarttagJobShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitSmarttagJobShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitSmarttagJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSmarttagJobShrinkRequest) SetContent(v string) *SubmitSmarttagJobShrinkRequest {
	s.Content = &v
	return s
}

func (s *SubmitSmarttagJobShrinkRequest) SetContentAddr(v string) *SubmitSmarttagJobShrinkRequest {
	s.ContentAddr = &v
	return s
}

func (s *SubmitSmarttagJobShrinkRequest) SetContentType(v string) *SubmitSmarttagJobShrinkRequest {
	s.ContentType = &v
	return s
}

func (s *SubmitSmarttagJobShrinkRequest) SetInputShrink(v string) *SubmitSmarttagJobShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *SubmitSmarttagJobShrinkRequest) SetNotifyUrl(v string) *SubmitSmarttagJobShrinkRequest {
	s.NotifyUrl = &v
	return s
}

func (s *SubmitSmarttagJobShrinkRequest) SetParams(v string) *SubmitSmarttagJobShrinkRequest {
	s.Params = &v
	return s
}

func (s *SubmitSmarttagJobShrinkRequest) SetScheduleConfigShrink(v string) *SubmitSmarttagJobShrinkRequest {
	s.ScheduleConfigShrink = &v
	return s
}

func (s *SubmitSmarttagJobShrinkRequest) SetTemplateConfig(v string) *SubmitSmarttagJobShrinkRequest {
	s.TemplateConfig = &v
	return s
}

func (s *SubmitSmarttagJobShrinkRequest) SetTemplateId(v string) *SubmitSmarttagJobShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitSmarttagJobShrinkRequest) SetTitle(v string) *SubmitSmarttagJobShrinkRequest {
	s.Title = &v
	return s
}

func (s *SubmitSmarttagJobShrinkRequest) SetUserData(v string) *SubmitSmarttagJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSmarttagJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
