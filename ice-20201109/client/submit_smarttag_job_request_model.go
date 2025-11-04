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
	SetTemplateId(v string) *SubmitSmarttagJobRequest
	GetTemplateId() *string
	SetTitle(v string) *SubmitSmarttagJobRequest
	GetTitle() *string
	SetUserData(v string) *SubmitSmarttagJobRequest
	GetUserData() *string
}

type SubmitSmarttagJobRequest struct {
	// The video description. The description can contain letters, digits, and hyphens (-) and cannot start with a special character. The description can be up to 1 KB in length.
	//
	// example:
	//
	// example content ****
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is discontinued.
	//
	// example:
	//
	// http://123.com/testVideo.mp4
	ContentAddr *string `json:"ContentAddr,omitempty" xml:"ContentAddr,omitempty"`
	// This parameter is discontinued.
	//
	// example:
	//
	// application/zip
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The job input.
	Input *SubmitSmarttagJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The URL for receiving callbacks. Set the value to an HTTP URL or an HTTPS URL.
	//
	// example:
	//
	// https://example.com/endpoint/aliyun/ai?id=76401125000***
	NotifyUrl *string `json:"NotifyUrl,omitempty" xml:"NotifyUrl,omitempty"`
	// The additional request parameters. The value is a JSON string. Example: {"needAsrData":true, "needOcrData":false}. The following parameters are supported:
	//
	// 	- needAsrData: specifies whether to query the automatic speech recognition (ASR) data. The value is of the BOOLEAN type. Default value: false. Valid values: true and false.
	//
	// 	- needOcrData: specifies whether to query the optical character recognition (OCR) data. The value is of the BOOLEAN type. Default value: false. Valid values: true and false.
	//
	// 	- needMetaData: specifies whether to query the metadata. The value is of the BOOLEAN type. Default value: false. Valid values: true and false.
	//
	// 	- nlpParams: the input parameters of the natural language processing (NLP) operator. The value is a JSON object. This parameter is empty by default, which indicates that the NLP operator is not used. For more information, see the "nlpParams" section of this topic.
	//
	// example:
	//
	// {"needAsrData":true, "needOcrData":false}
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The scheduling configurations.
	ScheduleConfig *SubmitSmarttagJobRequestScheduleConfig `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty" type:"Struct"`
	// The ID of the template that specifies the analysis algorithms. For more information about template operations, see [Configure templates](https://help.aliyun.com/document_detail/445702.html).
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The video title. The title can contain letters, digits, and hyphens (-) and cannot start with a special character. The title can be up to 256 bytes in length.
	//
	// example:
	//
	// example-title-****
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The data to be passed through Simple Message Queue (SMQ, formerly MNS) during callbacks. The data can be up to 1 KB in length. For more information about how to specify an SMQ queue for receiving callbacks, see UpdatePipeline.
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
	// If Type is set to OSS, specify an OSS path. Example: OSS://test-bucket/video/202208/test.mp4.
	//
	// If Type is set to Media, specify a media asset ID. Example: c5c62d8f0361337cab312dce8e77dc6d.
	//
	// If Type is set to URL, specify an HTTP URL. Example: https://zc-test.oss-cn-shanghai.aliyuncs.com/test/unknowFace.mp4.
	//
	// example:
	//
	// c5c62d8f0361337cab312dce8e77dc6d
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The media type. Valid values:
	//
	// 	- OSS
	//
	// 	- Media
	//
	// 	- URL
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
	// The ID of the ApsaraVideo Media Processing (MPS) queue to which you want to submit the smart tagging job. The MPS queue is bound to an SMQ queue. This parameter specifies the default MPS queue. By default, an MPS queue can process a maximum of two concurrent smart tagging jobs. To increase the limit, submit a ticket.
	//
	// example:
	//
	// acdbfe4323bcfdae
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// The job priority. This parameter is not implemented. You can leave this parameter empty or enter a random value.
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
