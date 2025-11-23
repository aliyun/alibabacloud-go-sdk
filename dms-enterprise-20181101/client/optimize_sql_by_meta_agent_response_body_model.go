// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOptimizeSqlByMetaAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *OptimizeSqlByMetaAgentResponseBodyData) *OptimizeSqlByMetaAgentResponseBody
	GetData() *OptimizeSqlByMetaAgentResponseBodyData
	SetErrorCode(v string) *OptimizeSqlByMetaAgentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *OptimizeSqlByMetaAgentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *OptimizeSqlByMetaAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OptimizeSqlByMetaAgentResponseBody
	GetSuccess() *bool
}

type OptimizeSqlByMetaAgentResponseBody struct {
	// The data returned.
	Data *OptimizeSqlByMetaAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed.
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

func (s OptimizeSqlByMetaAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OptimizeSqlByMetaAgentResponseBody) GoString() string {
	return s.String()
}

func (s *OptimizeSqlByMetaAgentResponseBody) GetData() *OptimizeSqlByMetaAgentResponseBodyData {
	return s.Data
}

func (s *OptimizeSqlByMetaAgentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *OptimizeSqlByMetaAgentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *OptimizeSqlByMetaAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OptimizeSqlByMetaAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OptimizeSqlByMetaAgentResponseBody) SetData(v *OptimizeSqlByMetaAgentResponseBodyData) *OptimizeSqlByMetaAgentResponseBody {
	s.Data = v
	return s
}

func (s *OptimizeSqlByMetaAgentResponseBody) SetErrorCode(v string) *OptimizeSqlByMetaAgentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *OptimizeSqlByMetaAgentResponseBody) SetErrorMessage(v string) *OptimizeSqlByMetaAgentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *OptimizeSqlByMetaAgentResponseBody) SetRequestId(v string) *OptimizeSqlByMetaAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *OptimizeSqlByMetaAgentResponseBody) SetSuccess(v bool) *OptimizeSqlByMetaAgentResponseBody {
	s.Success = &v
	return s
}

func (s *OptimizeSqlByMetaAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OptimizeSqlByMetaAgentResponseBodyData struct {
	// The answer.
	//
	// example:
	//
	// SQL优化结果...
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The session ID.
	//
	// example:
	//
	// f63a6eed-0e3c-4564-8533-b1295db8d6ff
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s OptimizeSqlByMetaAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s OptimizeSqlByMetaAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *OptimizeSqlByMetaAgentResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *OptimizeSqlByMetaAgentResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *OptimizeSqlByMetaAgentResponseBodyData) SetContent(v string) *OptimizeSqlByMetaAgentResponseBodyData {
	s.Content = &v
	return s
}

func (s *OptimizeSqlByMetaAgentResponseBodyData) SetSessionId(v string) *OptimizeSqlByMetaAgentResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *OptimizeSqlByMetaAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
