// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGlobalQuestionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGlobalQuestionsResponseBody
	GetCode() *string
	SetGlobalQuestions(v *ListGlobalQuestionsResponseBodyGlobalQuestions) *ListGlobalQuestionsResponseBody
	GetGlobalQuestions() *ListGlobalQuestionsResponseBodyGlobalQuestions
	SetHttpStatusCode(v int32) *ListGlobalQuestionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListGlobalQuestionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGlobalQuestionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListGlobalQuestionsResponseBody
	GetSuccess() *bool
}

type ListGlobalQuestionsResponseBody struct {
	// example:
	//
	// OK
	Code            *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	GlobalQuestions *ListGlobalQuestionsResponseBodyGlobalQuestions `json:"GlobalQuestions,omitempty" xml:"GlobalQuestions,omitempty" type:"Struct"`
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

func (s ListGlobalQuestionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGlobalQuestionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGlobalQuestionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGlobalQuestionsResponseBody) GetGlobalQuestions() *ListGlobalQuestionsResponseBodyGlobalQuestions {
	return s.GlobalQuestions
}

func (s *ListGlobalQuestionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListGlobalQuestionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGlobalQuestionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGlobalQuestionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListGlobalQuestionsResponseBody) SetCode(v string) *ListGlobalQuestionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListGlobalQuestionsResponseBody) SetGlobalQuestions(v *ListGlobalQuestionsResponseBodyGlobalQuestions) *ListGlobalQuestionsResponseBody {
	s.GlobalQuestions = v
	return s
}

func (s *ListGlobalQuestionsResponseBody) SetHttpStatusCode(v int32) *ListGlobalQuestionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListGlobalQuestionsResponseBody) SetMessage(v string) *ListGlobalQuestionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListGlobalQuestionsResponseBody) SetRequestId(v string) *ListGlobalQuestionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGlobalQuestionsResponseBody) SetSuccess(v bool) *ListGlobalQuestionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListGlobalQuestionsResponseBody) Validate() error {
	if s.GlobalQuestions != nil {
		if err := s.GlobalQuestions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGlobalQuestionsResponseBodyGlobalQuestions struct {
	List []*ListGlobalQuestionsResponseBodyGlobalQuestionsList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 12
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGlobalQuestionsResponseBodyGlobalQuestions) String() string {
	return dara.Prettify(s)
}

func (s ListGlobalQuestionsResponseBodyGlobalQuestions) GoString() string {
	return s.String()
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestions) GetList() []*ListGlobalQuestionsResponseBodyGlobalQuestionsList {
	return s.List
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestions) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestions) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestions) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestions) SetList(v []*ListGlobalQuestionsResponseBodyGlobalQuestionsList) *ListGlobalQuestionsResponseBodyGlobalQuestions {
	s.List = v
	return s
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestions) SetPageNumber(v int32) *ListGlobalQuestionsResponseBodyGlobalQuestions {
	s.PageNumber = &v
	return s
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestions) SetPageSize(v int32) *ListGlobalQuestionsResponseBodyGlobalQuestions {
	s.PageSize = &v
	return s
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestions) SetTotalCount(v int32) *ListGlobalQuestionsResponseBodyGlobalQuestions {
	s.TotalCount = &v
	return s
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestions) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGlobalQuestionsResponseBodyGlobalQuestionsList struct {
	Answers *string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// example:
	//
	// 53c27755-d41e-46a6-bb3c-4f66f257d50c
	GlobalQuestionId   *string `json:"GlobalQuestionId,omitempty" xml:"GlobalQuestionId,omitempty"`
	GlobalQuestionName *string `json:"GlobalQuestionName,omitempty" xml:"GlobalQuestionName,omitempty"`
	// example:
	//
	// SYSTEM
	GlobalQuestionType *string `json:"GlobalQuestionType,omitempty" xml:"GlobalQuestionType,omitempty"`
	Questions          *string `json:"Questions,omitempty" xml:"Questions,omitempty"`
	// example:
	//
	// aa279896-64a6-4182-864c-4f2b04ec8d17
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ListGlobalQuestionsResponseBodyGlobalQuestionsList) String() string {
	return dara.Prettify(s)
}

func (s ListGlobalQuestionsResponseBodyGlobalQuestionsList) GoString() string {
	return s.String()
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestionsList) GetAnswers() *string {
	return s.Answers
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestionsList) GetGlobalQuestionId() *string {
	return s.GlobalQuestionId
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestionsList) GetGlobalQuestionName() *string {
	return s.GlobalQuestionName
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestionsList) GetGlobalQuestionType() *string {
	return s.GlobalQuestionType
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestionsList) GetQuestions() *string {
	return s.Questions
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestionsList) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestionsList) SetAnswers(v string) *ListGlobalQuestionsResponseBodyGlobalQuestionsList {
	s.Answers = &v
	return s
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestionsList) SetGlobalQuestionId(v string) *ListGlobalQuestionsResponseBodyGlobalQuestionsList {
	s.GlobalQuestionId = &v
	return s
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestionsList) SetGlobalQuestionName(v string) *ListGlobalQuestionsResponseBodyGlobalQuestionsList {
	s.GlobalQuestionName = &v
	return s
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestionsList) SetGlobalQuestionType(v string) *ListGlobalQuestionsResponseBodyGlobalQuestionsList {
	s.GlobalQuestionType = &v
	return s
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestionsList) SetQuestions(v string) *ListGlobalQuestionsResponseBodyGlobalQuestionsList {
	s.Questions = &v
	return s
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestionsList) SetScriptId(v string) *ListGlobalQuestionsResponseBodyGlobalQuestionsList {
	s.ScriptId = &v
	return s
}

func (s *ListGlobalQuestionsResponseBodyGlobalQuestionsList) Validate() error {
	return dara.Validate(s)
}
