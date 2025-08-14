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
	SetTemplateId(v string) *SubmitVideoCognitionJobShrinkRequest
	GetTemplateId() *string
	SetTitle(v string) *SubmitVideoCognitionJobShrinkRequest
	GetTitle() *string
	SetUserData(v string) *SubmitVideoCognitionJobShrinkRequest
	GetUserData() *string
}

type SubmitVideoCognitionJobShrinkRequest struct {
	InputShrink *string `json:"Input,omitempty" xml:"Input,omitempty"`
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
	// example:
	//
	// 39f8e0bc00***************
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// example-title-****
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
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
