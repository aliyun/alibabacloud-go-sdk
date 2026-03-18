// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterDeleteConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModelRouterDeleteConversationResponseBody
	GetData() *bool
	SetErrCode(v string) *ModelRouterDeleteConversationResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterDeleteConversationResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterDeleteConversationResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterDeleteConversationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterDeleteConversationResponseBody
	GetSuccess() *bool
}

type ModelRouterDeleteConversationResponseBody struct {
	// example:
	//
	// []
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
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

func (s ModelRouterDeleteConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterDeleteConversationResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterDeleteConversationResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModelRouterDeleteConversationResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterDeleteConversationResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterDeleteConversationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterDeleteConversationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterDeleteConversationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterDeleteConversationResponseBody) SetData(v bool) *ModelRouterDeleteConversationResponseBody {
	s.Data = &v
	return s
}

func (s *ModelRouterDeleteConversationResponseBody) SetErrCode(v string) *ModelRouterDeleteConversationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterDeleteConversationResponseBody) SetErrMessage(v string) *ModelRouterDeleteConversationResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterDeleteConversationResponseBody) SetHttpStatusCode(v int32) *ModelRouterDeleteConversationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterDeleteConversationResponseBody) SetRequestId(v string) *ModelRouterDeleteConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterDeleteConversationResponseBody) SetSuccess(v bool) *ModelRouterDeleteConversationResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterDeleteConversationResponseBody) Validate() error {
	return dara.Validate(s)
}
