// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskTemplateQuestionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTaskTemplateQuestionsResponseBody
	GetCode() *int32
	SetDetails(v string) *GetTaskTemplateQuestionsResponseBody
	GetDetails() *string
	SetErrorCode(v string) *GetTaskTemplateQuestionsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetTaskTemplateQuestionsResponseBody
	GetMessage() *string
	SetQuestions(v []*QuestionPlugin) *GetTaskTemplateQuestionsResponseBody
	GetQuestions() []*QuestionPlugin
	SetRequestId(v string) *GetTaskTemplateQuestionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskTemplateQuestionsResponseBody
	GetSuccess() *bool
}

type GetTaskTemplateQuestionsResponseBody struct {
	// Total amount of data under the current request conditions. This parameter is optional and does not need to be returned by default.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details
	//
	// example:
	//
	// null
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// error code
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Response message of the request
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Question list
	Questions []*QuestionPlugin `json:"Questions,omitempty" xml:"Questions,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1F29E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// is succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTaskTemplateQuestionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskTemplateQuestionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskTemplateQuestionsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTaskTemplateQuestionsResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetTaskTemplateQuestionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTaskTemplateQuestionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskTemplateQuestionsResponseBody) GetQuestions() []*QuestionPlugin {
	return s.Questions
}

func (s *GetTaskTemplateQuestionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskTemplateQuestionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskTemplateQuestionsResponseBody) SetCode(v int32) *GetTaskTemplateQuestionsResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskTemplateQuestionsResponseBody) SetDetails(v string) *GetTaskTemplateQuestionsResponseBody {
	s.Details = &v
	return s
}

func (s *GetTaskTemplateQuestionsResponseBody) SetErrorCode(v string) *GetTaskTemplateQuestionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTaskTemplateQuestionsResponseBody) SetMessage(v string) *GetTaskTemplateQuestionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskTemplateQuestionsResponseBody) SetQuestions(v []*QuestionPlugin) *GetTaskTemplateQuestionsResponseBody {
	s.Questions = v
	return s
}

func (s *GetTaskTemplateQuestionsResponseBody) SetRequestId(v string) *GetTaskTemplateQuestionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskTemplateQuestionsResponseBody) SetSuccess(v bool) *GetTaskTemplateQuestionsResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskTemplateQuestionsResponseBody) Validate() error {
	if s.Questions != nil {
		for _, item := range s.Questions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
