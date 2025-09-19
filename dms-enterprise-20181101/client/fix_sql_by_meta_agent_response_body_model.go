// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFixSqlByMetaAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *FixSqlByMetaAgentResponseBodyData) *FixSqlByMetaAgentResponseBody
	GetData() *FixSqlByMetaAgentResponseBodyData
	SetErrorCode(v string) *FixSqlByMetaAgentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *FixSqlByMetaAgentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *FixSqlByMetaAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FixSqlByMetaAgentResponseBody
	GetSuccess() *bool
}

type FixSqlByMetaAgentResponseBody struct {
	Data *FixSqlByMetaAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FixSqlByMetaAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FixSqlByMetaAgentResponseBody) GoString() string {
	return s.String()
}

func (s *FixSqlByMetaAgentResponseBody) GetData() *FixSqlByMetaAgentResponseBodyData {
	return s.Data
}

func (s *FixSqlByMetaAgentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *FixSqlByMetaAgentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *FixSqlByMetaAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FixSqlByMetaAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FixSqlByMetaAgentResponseBody) SetData(v *FixSqlByMetaAgentResponseBodyData) *FixSqlByMetaAgentResponseBody {
	s.Data = v
	return s
}

func (s *FixSqlByMetaAgentResponseBody) SetErrorCode(v string) *FixSqlByMetaAgentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *FixSqlByMetaAgentResponseBody) SetErrorMessage(v string) *FixSqlByMetaAgentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *FixSqlByMetaAgentResponseBody) SetRequestId(v string) *FixSqlByMetaAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *FixSqlByMetaAgentResponseBody) SetSuccess(v bool) *FixSqlByMetaAgentResponseBody {
	s.Success = &v
	return s
}

func (s *FixSqlByMetaAgentResponseBody) Validate() error {
	return dara.Validate(s)
}

type FixSqlByMetaAgentResponseBodyData struct {
	// example:
	//
	// SQL修复结果...
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// f63a6eed-0e3c-4564-8533-b1295db8d6ff
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s FixSqlByMetaAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FixSqlByMetaAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *FixSqlByMetaAgentResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *FixSqlByMetaAgentResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *FixSqlByMetaAgentResponseBodyData) SetContent(v string) *FixSqlByMetaAgentResponseBodyData {
	s.Content = &v
	return s
}

func (s *FixSqlByMetaAgentResponseBodyData) SetSessionId(v string) *FixSqlByMetaAgentResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *FixSqlByMetaAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
