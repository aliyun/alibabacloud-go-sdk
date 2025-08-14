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
	SetTemplateId(v string) *SubmitSmarttagJobShrinkRequest
	GetTemplateId() *string
	SetTitle(v string) *SubmitSmarttagJobShrinkRequest
	GetTitle() *string
	SetUserData(v string) *SubmitSmarttagJobShrinkRequest
	GetUserData() *string
}

type SubmitSmarttagJobShrinkRequest struct {
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
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
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
	ScheduleConfigShrink *string `json:"ScheduleConfig,omitempty" xml:"ScheduleConfig,omitempty"`
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
