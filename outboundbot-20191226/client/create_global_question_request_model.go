// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGlobalQuestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnswers(v string) *CreateGlobalQuestionRequest
	GetAnswers() *string
	SetGlobalQuestionName(v string) *CreateGlobalQuestionRequest
	GetGlobalQuestionName() *string
	SetGlobalQuestionType(v string) *CreateGlobalQuestionRequest
	GetGlobalQuestionType() *string
	SetInstanceId(v string) *CreateGlobalQuestionRequest
	GetInstanceId() *string
	SetQuestions(v string) *CreateGlobalQuestionRequest
	GetQuestions() *string
	SetScriptId(v string) *CreateGlobalQuestionRequest
	GetScriptId() *string
}

type CreateGlobalQuestionRequest struct {
	// This parameter is required.
	Answers *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// This parameter is required.
	GlobalQuestionName *string `json:"GlobalQuestionName,omitempty" xml:"GlobalQuestionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// COMMON
	GlobalQuestionType *string `json:"GlobalQuestionType,omitempty" xml:"GlobalQuestionType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 174952ab-9825-4cc9-a5e2-de82d7fa4cdd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Questions *string `json:"Questions,omitempty" xml:"Questions,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 36fea72b-d6fa-4974-ace7-19ffe3f622fb
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s CreateGlobalQuestionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGlobalQuestionRequest) GoString() string {
	return s.String()
}

func (s *CreateGlobalQuestionRequest) GetAnswers() *string {
	return s.Answers
}

func (s *CreateGlobalQuestionRequest) GetGlobalQuestionName() *string {
	return s.GlobalQuestionName
}

func (s *CreateGlobalQuestionRequest) GetGlobalQuestionType() *string {
	return s.GlobalQuestionType
}

func (s *CreateGlobalQuestionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateGlobalQuestionRequest) GetQuestions() *string {
	return s.Questions
}

func (s *CreateGlobalQuestionRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *CreateGlobalQuestionRequest) SetAnswers(v string) *CreateGlobalQuestionRequest {
	s.Answers = &v
	return s
}

func (s *CreateGlobalQuestionRequest) SetGlobalQuestionName(v string) *CreateGlobalQuestionRequest {
	s.GlobalQuestionName = &v
	return s
}

func (s *CreateGlobalQuestionRequest) SetGlobalQuestionType(v string) *CreateGlobalQuestionRequest {
	s.GlobalQuestionType = &v
	return s
}

func (s *CreateGlobalQuestionRequest) SetInstanceId(v string) *CreateGlobalQuestionRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateGlobalQuestionRequest) SetQuestions(v string) *CreateGlobalQuestionRequest {
	s.Questions = &v
	return s
}

func (s *CreateGlobalQuestionRequest) SetScriptId(v string) *CreateGlobalQuestionRequest {
	s.ScriptId = &v
	return s
}

func (s *CreateGlobalQuestionRequest) Validate() error {
	return dara.Validate(s)
}
