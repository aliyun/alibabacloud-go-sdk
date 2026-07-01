// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoCognitionJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputShrink(v string) *SubmitVideoCognitionJobShrinkRequest
	GetInputShrink() *string
	SetParams(v string) *SubmitVideoCognitionJobShrinkRequest
	GetParams() *string
	SetTemplateConfig(v string) *SubmitVideoCognitionJobShrinkRequest
	GetTemplateConfig() *string
	SetTemplateId(v string) *SubmitVideoCognitionJobShrinkRequest
	GetTemplateId() *string
	SetTitle(v string) *SubmitVideoCognitionJobShrinkRequest
	GetTitle() *string
	SetUserData(v string) *SubmitVideoCognitionJobShrinkRequest
	GetUserData() *string
}

type SubmitVideoCognitionJobShrinkRequest struct {
	// The input media object.
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// A JSON string containing additional parameters for operators such as natural language processing, shot detection, custom tagging, and action recognition.
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
	// The template configuration.
	TemplateConfig *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The ID of the template that specifies the analysis algorithms to use. For more information about managing templates, see [Create Custom Template](https://help.aliyun.com/zh/ims/developer-reference/api-ice-2020-11-09-createcustomtemplate?spm=a2c4g.11186623.help-menu-193643.d_5_0_3_3_0_0.17b66afamjKySv) and [AI-powered tagging template](https://help.aliyun.com/zh/ims/user-guide/smart-tagging-template?spm=a2c4g.11186623.0.i15).
	//
	// example:
	//
	// 39f8e0bc00***************
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The title of the video. The title can contain Chinese characters, English letters, digits, and hyphens (-). The title cannot start with a special character and must not exceed 256 bytes in length.
	//
	// example:
	//
	// example-title-****
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The user-defined data. The service returns this data unmodified in the callback notification. This parameter cannot exceed 1,024 bytes.
	//
	// example:
	//
	// {"test":1}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitVideoCognitionJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoCognitionJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoCognitionJobShrinkRequest) GetInputShrink() *string {
	return s.InputShrink
}

func (s *SubmitVideoCognitionJobShrinkRequest) GetParams() *string {
	return s.Params
}

func (s *SubmitVideoCognitionJobShrinkRequest) GetTemplateConfig() *string {
	return s.TemplateConfig
}

func (s *SubmitVideoCognitionJobShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitVideoCognitionJobShrinkRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitVideoCognitionJobShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitVideoCognitionJobShrinkRequest) SetInputShrink(v string) *SubmitVideoCognitionJobShrinkRequest {
	s.InputShrink = &v
	return s
}

func (s *SubmitVideoCognitionJobShrinkRequest) SetParams(v string) *SubmitVideoCognitionJobShrinkRequest {
	s.Params = &v
	return s
}

func (s *SubmitVideoCognitionJobShrinkRequest) SetTemplateConfig(v string) *SubmitVideoCognitionJobShrinkRequest {
	s.TemplateConfig = &v
	return s
}

func (s *SubmitVideoCognitionJobShrinkRequest) SetTemplateId(v string) *SubmitVideoCognitionJobShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitVideoCognitionJobShrinkRequest) SetTitle(v string) *SubmitVideoCognitionJobShrinkRequest {
	s.Title = &v
	return s
}

func (s *SubmitVideoCognitionJobShrinkRequest) SetUserData(v string) *SubmitVideoCognitionJobShrinkRequest {
	s.UserData = &v
	return s
}

func (s *SubmitVideoCognitionJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
