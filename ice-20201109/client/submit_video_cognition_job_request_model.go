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
	Input *SubmitVideoCognitionJobRequestInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type SubmitVideoCognitionJobRequestInput struct {
	// example:
	//
	// c5c62d8f03613************c6d
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
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
