// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnswerSqlSyntaxByMetaAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AnswerSqlSyntaxByMetaAgentResponseBodyData) *AnswerSqlSyntaxByMetaAgentResponseBody
	GetData() *AnswerSqlSyntaxByMetaAgentResponseBodyData
	SetErrorCode(v string) *AnswerSqlSyntaxByMetaAgentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AnswerSqlSyntaxByMetaAgentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AnswerSqlSyntaxByMetaAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AnswerSqlSyntaxByMetaAgentResponseBody
	GetSuccess() *bool
}

type AnswerSqlSyntaxByMetaAgentResponseBody struct {
	// The data returned.
	Data *AnswerSqlSyntaxByMetaAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0C1CB646-1DE4-4AD0-B4A4-7D47DD52E931
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AnswerSqlSyntaxByMetaAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AnswerSqlSyntaxByMetaAgentResponseBody) GoString() string {
	return s.String()
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBody) GetData() *AnswerSqlSyntaxByMetaAgentResponseBodyData {
	return s.Data
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBody) SetData(v *AnswerSqlSyntaxByMetaAgentResponseBodyData) *AnswerSqlSyntaxByMetaAgentResponseBody {
	s.Data = v
	return s
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBody) SetErrorCode(v string) *AnswerSqlSyntaxByMetaAgentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBody) SetErrorMessage(v string) *AnswerSqlSyntaxByMetaAgentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBody) SetRequestId(v string) *AnswerSqlSyntaxByMetaAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBody) SetSuccess(v bool) *AnswerSqlSyntaxByMetaAgentResponseBody {
	s.Success = &v
	return s
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AnswerSqlSyntaxByMetaAgentResponseBodyData struct {
	// The answer to the question.
	//
	// example:
	//
	// mysql数据库类型下...
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The session ID.
	//
	// example:
	//
	// f63a6eed-0e3c-4564-8533-b1295db8d6ff
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s AnswerSqlSyntaxByMetaAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AnswerSqlSyntaxByMetaAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBodyData) SetContent(v string) *AnswerSqlSyntaxByMetaAgentResponseBodyData {
	s.Content = &v
	return s
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBodyData) SetSessionId(v string) *AnswerSqlSyntaxByMetaAgentResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *AnswerSqlSyntaxByMetaAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
