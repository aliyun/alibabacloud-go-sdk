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
	// Answer to the global question
	//
	// This parameter is required.
	//
	// example:
	//
	// ["你好,您可以再说一遍吗","不好意思我刚才没有听清"]
	Answers *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// Name of the global question
	//
	// This parameter is required.
	//
	// example:
	//
	// 未识别全局问题
	GlobalQuestionName *string `json:"GlobalQuestionName,omitempty" xml:"GlobalQuestionName,omitempty"`
	// Type of the global question
	//
	// This parameter is required.
	//
	// example:
	//
	// COMMON
	GlobalQuestionType *string `json:"GlobalQuestionType,omitempty" xml:"GlobalQuestionType,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 361c8a53-0e29-42f3-8aa7-c7752d010399
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Global question
	//
	// This parameter is required.
	//
	// example:
	//
	// ["未识别的用户问题"]
	Questions *string `json:"Questions,omitempty" xml:"Questions,omitempty"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// c5c5d8c0-c0f1-48a7-be2b-dc46006d888a
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
