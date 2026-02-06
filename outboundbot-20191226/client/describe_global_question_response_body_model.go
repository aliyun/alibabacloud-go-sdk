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
	// example:
	//
	// OK
	Code           *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	GlobalQuestion *DescribeGlobalQuestionResponseBodyGlobalQuestion `json:"GlobalQuestion,omitempty" xml:"GlobalQuestion,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	Answers *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// example:
	//
	// f160ec2e-94f2-4c03-87be-ece5b52d5dd9
	GlobalQuestionId   *string `json:"GlobalQuestionId,omitempty" xml:"GlobalQuestionId,omitempty"`
	GlobalQuestionName *string `json:"GlobalQuestionName,omitempty" xml:"GlobalQuestionName,omitempty"`
	// example:
	//
	// COMMON
	GlobalQuestionType *string `json:"GlobalQuestionType,omitempty" xml:"GlobalQuestionType,omitempty"`
	Questions          *string `json:"Questions,omitempty" xml:"Questions,omitempty"`
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
