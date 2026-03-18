// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterChatCompletionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ModelRouterChatCompletionsResponseBody
	GetData() interface{}
	SetErrCode(v string) *ModelRouterChatCompletionsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterChatCompletionsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterChatCompletionsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterChatCompletionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterChatCompletionsResponseBody
	GetSuccess() *bool
}

type ModelRouterChatCompletionsResponseBody struct {
	// example:
	//
	// []
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterChatCompletionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterChatCompletionsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterChatCompletionsResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ModelRouterChatCompletionsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterChatCompletionsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterChatCompletionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterChatCompletionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterChatCompletionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterChatCompletionsResponseBody) SetData(v interface{}) *ModelRouterChatCompletionsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterChatCompletionsResponseBody) SetErrCode(v string) *ModelRouterChatCompletionsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterChatCompletionsResponseBody) SetErrMessage(v string) *ModelRouterChatCompletionsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterChatCompletionsResponseBody) SetHttpStatusCode(v int32) *ModelRouterChatCompletionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterChatCompletionsResponseBody) SetRequestId(v string) *ModelRouterChatCompletionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterChatCompletionsResponseBody) SetSuccess(v bool) *ModelRouterChatCompletionsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterChatCompletionsResponseBody) Validate() error {
	return dara.Validate(s)
}
