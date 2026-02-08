// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalQuestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeGlobalQuestionResponseBody
	GetCode() *string
	SetGlobalQuestion(v *DescribeGlobalQuestionResponseBodyGlobalQuestion) *DescribeGlobalQuestionResponseBody
	GetGlobalQuestion() *DescribeGlobalQuestionResponseBodyGlobalQuestion
	SetHttpStatusCode(v int32) *DescribeGlobalQuestionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DescribeGlobalQuestionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeGlobalQuestionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeGlobalQuestionResponseBody
	GetSuccess() *bool
}

type DescribeGlobalQuestionResponseBody struct {
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Global question information
	GlobalQuestion *DescribeGlobalQuestionResponseBodyGlobalQuestion `json:"GlobalQuestion,omitempty" xml:"GlobalQuestion,omitempty" type:"Struct"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeGlobalQuestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalQuestionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeGlobalQuestionResponseBody) GetGlobalQuestion() *DescribeGlobalQuestionResponseBodyGlobalQuestion {
	return s.GlobalQuestion
}

func (s *DescribeGlobalQuestionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeGlobalQuestionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeGlobalQuestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGlobalQuestionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeGlobalQuestionResponseBody) SetCode(v string) *DescribeGlobalQuestionResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeGlobalQuestionResponseBody) SetGlobalQuestion(v *DescribeGlobalQuestionResponseBodyGlobalQuestion) *DescribeGlobalQuestionResponseBody {
	s.GlobalQuestion = v
	return s
}

func (s *DescribeGlobalQuestionResponseBody) SetHttpStatusCode(v int32) *DescribeGlobalQuestionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeGlobalQuestionResponseBody) SetMessage(v string) *DescribeGlobalQuestionResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeGlobalQuestionResponseBody) SetRequestId(v string) *DescribeGlobalQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalQuestionResponseBody) SetSuccess(v bool) *DescribeGlobalQuestionResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeGlobalQuestionResponseBody) Validate() error {
	if s.GlobalQuestion != nil {
		if err := s.GlobalQuestion.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeGlobalQuestionResponseBodyGlobalQuestion struct {
	// Answer responses
	//
	// example:
	//
	// ["你好,我是你的专属客服顾问."]
	Answers *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// Global script ID
	//
	// example:
	//
	// f160ec2e-94f2-4c03-87be-ece5b52d5dd9
	GlobalQuestionId *string `json:"GlobalQuestionId,omitempty" xml:"GlobalQuestionId,omitempty"`
	// Global question name
	//
	// example:
	//
	// 你是谁-全局问题
	GlobalQuestionName *string `json:"GlobalQuestionName,omitempty" xml:"GlobalQuestionName,omitempty"`
	// Global script type
	//
	// example:
	//
	// COMMON
	GlobalQuestionType *string `json:"GlobalQuestionType,omitempty" xml:"GlobalQuestionType,omitempty"`
	// Global questions
	//
	// example:
	//
	// ["你是谁","你叫什么"]
	Questions *string `json:"Questions,omitempty" xml:"Questions,omitempty"`
	// Scenario ID
	//
	// example:
	//
	// 290e06a5-6de2-4cc8-8a9c-72b7c152256c
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s DescribeGlobalQuestionResponseBodyGlobalQuestion) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalQuestionResponseBodyGlobalQuestion) GoString() string {
	return s.String()
}

func (s *DescribeGlobalQuestionResponseBodyGlobalQuestion) GetAnswers() *string {
	return s.Answers
}

func (s *DescribeGlobalQuestionResponseBodyGlobalQuestion) GetGlobalQuestionId() *string {
	return s.GlobalQuestionId
}

func (s *DescribeGlobalQuestionResponseBodyGlobalQuestion) GetGlobalQuestionName() *string {
	return s.GlobalQuestionName
}

func (s *DescribeGlobalQuestionResponseBodyGlobalQuestion) GetGlobalQuestionType() *string {
	return s.GlobalQuestionType
}

func (s *DescribeGlobalQuestionResponseBodyGlobalQuestion) GetQuestions() *string {
	return s.Questions
}

func (s *DescribeGlobalQuestionResponseBodyGlobalQuestion) GetScriptId() *string {
	return s.ScriptId
}

func (s *DescribeGlobalQuestionResponseBodyGlobalQuestion) SetAnswers(v string) *DescribeGlobalQuestionResponseBodyGlobalQuestion {
	s.Answers = &v
	return s
}

func (s *DescribeGlobalQuestionResponseBodyGlobalQuestion) SetGlobalQuestionId(v string) *DescribeGlobalQuestionResponseBodyGlobalQuestion {
	s.GlobalQuestionId = &v
	return s
}

func (s *DescribeGlobalQuestionResponseBodyGlobalQuestion) SetGlobalQuestionName(v string) *DescribeGlobalQuestionResponseBodyGlobalQuestion {
	s.GlobalQuestionName = &v
	return s
}

func (s *DescribeGlobalQuestionResponseBodyGlobalQuestion) SetGlobalQuestionType(v string) *DescribeGlobalQuestionResponseBodyGlobalQuestion {
	s.GlobalQuestionType = &v
	return s
}

func (s *DescribeGlobalQuestionResponseBodyGlobalQuestion) SetQuestions(v string) *DescribeGlobalQuestionResponseBodyGlobalQuestion {
	s.Questions = &v
	return s
}

func (s *DescribeGlobalQuestionResponseBodyGlobalQuestion) SetScriptId(v string) *DescribeGlobalQuestionResponseBodyGlobalQuestion {
	s.ScriptId = &v
	return s
}

func (s *DescribeGlobalQuestionResponseBodyGlobalQuestion) Validate() error {
	return dara.Validate(s)
}
