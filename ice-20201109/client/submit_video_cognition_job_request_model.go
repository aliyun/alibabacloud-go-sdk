// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoCognitionJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v *SubmitVideoCognitionJobRequestInput) *SubmitVideoCognitionJobRequest
	GetInput() *SubmitVideoCognitionJobRequestInput
	SetParams(v string) *SubmitVideoCognitionJobRequest
	GetParams() *string
	SetTemplateId(v string) *SubmitVideoCognitionJobRequest
	GetTemplateId() *string
	SetTitle(v string) *SubmitVideoCognitionJobRequest
	GetTitle() *string
	SetUserData(v string) *SubmitVideoCognitionJobRequest
	GetUserData() *string
}

type SubmitVideoCognitionJobRequest struct {
	// The media input object.
	Input *SubmitVideoCognitionJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// Additional request parameters, provided as a JSON string. This is used to pass specific settings for various AI analysis modules, such as Natural Language Processing (NLP), shot segmentation, tagging, and action recognition.
	//
	// example:
	//
	// {
	//
	// 	"nlpParams": {
	//
	// 		"sourceLanguage": "cn",
	//
	// 		"diarizationEnabled": true,
	//
	// 		"speakerCount": 2,
	//
	// 		"summarizationEnabled": true,
	//
	// 		"summarizationTypes": "Paragraph,Conversational,QuestionsAnswering,MindMap",
	//
	// 		"translationEnabled": true,
	//
	// 		"targetLanguages": "en",
	//
	// 		"autoChaptersEnabled": true,
	//
	// 		"meetingAssistanceEnabled": true
	//
	// 	}
	//
	// }
	Params *string `json:"Params,omitempty" xml:"Params,omitempty"`
	// The ID of the template that specifies the analysis algorithms to be used. For details, see [CreateCustomTemplate](https://help.aliyun.com/zh/ims/developer-reference/api-ice-2020-11-09-createcustomtemplate?spm=a2c4g.11186623.help-menu-193643.d_5_0_3_3_0_0.17b66afamjKySv) and [smart tagging template](https://help.aliyun.com/zh/ims/user-guide/smart-tagging-template?spm=a2c4g.11186623.0.i15).
	//
	// example:
	//
	// 39f8e0bc00***************
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The video title. It supports letters, digits, and hyphens (-), and cannot start with a special character. Max length: 256 bytes.
	//
	// example:
	//
	// example-title-****
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The user-defined data that is passed through and returned as-is in the response. Max length: 1,024 bytes.
	//
	// example:
	//
	// {"test":1}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitVideoCognitionJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoCognitionJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoCognitionJobRequest) GetInput() *SubmitVideoCognitionJobRequestInput {
	return s.Input
}

func (s *SubmitVideoCognitionJobRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitVideoCognitionJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitVideoCognitionJobRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitVideoCognitionJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitVideoCognitionJobRequest) SetInput(v *SubmitVideoCognitionJobRequestInput) *SubmitVideoCognitionJobRequest {
	s.Input = v
	return s
}

func (s *SubmitVideoCognitionJobRequest) SetParams(v string) *SubmitVideoCognitionJobRequest {
	s.Params = &v
	return s
}

func (s *SubmitVideoCognitionJobRequest) SetTemplateId(v string) *SubmitVideoCognitionJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitVideoCognitionJobRequest) SetTitle(v string) *SubmitVideoCognitionJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitVideoCognitionJobRequest) SetUserData(v string) *SubmitVideoCognitionJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitVideoCognitionJobRequest) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitVideoCognitionJobRequestInput struct {
	// If Type is set to OSS, specify an OSS path. Example: OSS://test-bucket/video/202208/test.mp4.
	//
	// If Type is set to Media, specify a media asset ID. Example: c5c62d8f0361337cab312dce8e77dc6d.
	//
	// If Type is set to URL, specify an HTTP URL. Example: https://zc-test.oss-cn-shanghai.aliyuncs.com/test/unknowFace.mp4.
	//
	// example:
	//
	// c5c62d8f03613************c6d
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of media input. Valid values:
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

func (s SubmitVideoCognitionJobRequestInput) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoCognitionJobRequestInput) GoString() string {
	return s.String()
}

func (s *SubmitVideoCognitionJobRequestInput) GetMedia() *string {
	return s.Media
}

func (s *SubmitVideoCognitionJobRequestInput) GetType() *string {
	return s.Type
}

func (s *SubmitVideoCognitionJobRequestInput) SetMedia(v string) *SubmitVideoCognitionJobRequestInput {
	s.Media = &v
	return s
}

func (s *SubmitVideoCognitionJobRequestInput) SetType(v string) *SubmitVideoCognitionJobRequestInput {
	s.Type = &v
	return s
}

func (s *SubmitVideoCognitionJobRequestInput) Validate() error {
	return dara.Validate(s)
}
