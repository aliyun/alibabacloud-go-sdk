// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGlobalQuestionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnswers(v string) *ModifyGlobalQuestionRequest
	GetAnswers() *string
	SetGlobalQuestionId(v string) *ModifyGlobalQuestionRequest
	GetGlobalQuestionId() *string
	SetGlobalQuestionName(v string) *ModifyGlobalQuestionRequest
	GetGlobalQuestionName() *string
	SetGlobalQuestionType(v string) *ModifyGlobalQuestionRequest
	GetGlobalQuestionType() *string
	SetInstanceId(v string) *ModifyGlobalQuestionRequest
	GetInstanceId() *string
	SetQuestions(v string) *ModifyGlobalQuestionRequest
	GetQuestions() *string
	SetScriptId(v string) *ModifyGlobalQuestionRequest
	GetScriptId() *string
}

type ModifyGlobalQuestionRequest struct {
	// List of response scripts
	//
	// This parameter is required.
	//
	// example:
	//
	// ["你好,我是你的专属客服顾问."]
	Answers *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// Global question ID. This parameter is the unique value of the source global question.
	//
	// This parameter is required.
	//
	// example:
	//
	// ad80de88-1661-445a-92ec-bf88dc45d581
	GlobalQuestionId *string `json:"GlobalQuestionId,omitempty" xml:"GlobalQuestionId,omitempty"`
	// Global question name
	//
	// This parameter is required.
	//
	// example:
	//
	// 未识别全局问题
	GlobalQuestionName *string `json:"GlobalQuestionName,omitempty" xml:"GlobalQuestionName,omitempty"`
	// Global question type
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
	// e5035654-1745-484a-8c5b-165f7c7bcd79
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// List of questions
	//
	// This parameter is required.
	//
	// example:
	//
	// ["你是谁","你叫什么"]
	Questions *string `json:"Questions,omitempty" xml:"Questions,omitempty"`
	// Scenario ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 0fe7f71c-8771-42ef-9bb1-19aa16ae7120
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ModifyGlobalQuestionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyGlobalQuestionRequest) GoString() string {
	return s.String()
}

func (s *ModifyGlobalQuestionRequest) GetAnswers() *string {
	return s.Answers
}

func (s *ModifyGlobalQuestionRequest) GetGlobalQuestionId() *string {
	return s.GlobalQuestionId
}

func (s *ModifyGlobalQuestionRequest) GetGlobalQuestionName() *string {
	return s.GlobalQuestionName
}

func (s *ModifyGlobalQuestionRequest) GetGlobalQuestionType() *string {
	return s.GlobalQuestionType
}

func (s *ModifyGlobalQuestionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyGlobalQuestionRequest) GetQuestions() *string {
	return s.Questions
}

func (s *ModifyGlobalQuestionRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ModifyGlobalQuestionRequest) SetAnswers(v string) *ModifyGlobalQuestionRequest {
	s.Answers = &v
	return s
}

func (s *ModifyGlobalQuestionRequest) SetGlobalQuestionId(v string) *ModifyGlobalQuestionRequest {
	s.GlobalQuestionId = &v
	return s
}

func (s *ModifyGlobalQuestionRequest) SetGlobalQuestionName(v string) *ModifyGlobalQuestionRequest {
	s.GlobalQuestionName = &v
	return s
}

func (s *ModifyGlobalQuestionRequest) SetGlobalQuestionType(v string) *ModifyGlobalQuestionRequest {
	s.GlobalQuestionType = &v
	return s
}

func (s *ModifyGlobalQuestionRequest) SetInstanceId(v string) *ModifyGlobalQuestionRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyGlobalQuestionRequest) SetQuestions(v string) *ModifyGlobalQuestionRequest {
	s.Questions = &v
	return s
}

func (s *ModifyGlobalQuestionRequest) SetScriptId(v string) *ModifyGlobalQuestionRequest {
	s.ScriptId = &v
	return s
}

func (s *ModifyGlobalQuestionRequest) Validate() error {
	return dara.Validate(s)
}
