// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateQuestionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTemplateQuestionsResponseBody
	GetCode() *int32
	SetDetails(v string) *GetTemplateQuestionsResponseBody
	GetDetails() *string
	SetErrorCode(v string) *GetTemplateQuestionsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *GetTemplateQuestionsResponseBody
	GetMessage() *string
	SetQuestionConfigs(v []*QuestionPlugin) *GetTemplateQuestionsResponseBody
	GetQuestionConfigs() []*QuestionPlugin
	SetRequestId(v string) *GetTemplateQuestionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTemplateQuestionsResponseBody
	GetSuccess() *bool
}

type GetTemplateQuestionsResponseBody struct {
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
	// Return message of the request
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// List of question configurations
	QuestionConfigs []*QuestionPlugin `json:"QuestionConfigs,omitempty" xml:"QuestionConfigs,omitempty" type:"Repeated"`
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

func (s GetTemplateQuestionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateQuestionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateQuestionsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTemplateQuestionsResponseBody) GetDetails() *string {
	return s.Details
}

func (s *GetTemplateQuestionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetTemplateQuestionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTemplateQuestionsResponseBody) GetQuestionConfigs() []*QuestionPlugin {
	return s.QuestionConfigs
}

func (s *GetTemplateQuestionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateQuestionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTemplateQuestionsResponseBody) SetCode(v int32) *GetTemplateQuestionsResponseBody {
	s.Code = &v
	return s
}

func (s *GetTemplateQuestionsResponseBody) SetDetails(v string) *GetTemplateQuestionsResponseBody {
	s.Details = &v
	return s
}

func (s *GetTemplateQuestionsResponseBody) SetErrorCode(v string) *GetTemplateQuestionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetTemplateQuestionsResponseBody) SetMessage(v string) *GetTemplateQuestionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetTemplateQuestionsResponseBody) SetQuestionConfigs(v []*QuestionPlugin) *GetTemplateQuestionsResponseBody {
	s.QuestionConfigs = v
	return s
}

func (s *GetTemplateQuestionsResponseBody) SetRequestId(v string) *GetTemplateQuestionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateQuestionsResponseBody) SetSuccess(v bool) *GetTemplateQuestionsResponseBody {
	s.Success = &v
	return s
}

func (s *GetTemplateQuestionsResponseBody) Validate() error {
	if s.QuestionConfigs != nil {
		for _, item := range s.QuestionConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
