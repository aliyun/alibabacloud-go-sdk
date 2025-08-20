// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryChatInstanceSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryChatInstanceSessionsResponseBody
	GetCode() *string
	SetData(v []*ChatSessionInfo) *QueryChatInstanceSessionsResponseBody
	GetData() []*ChatSessionInfo
	SetHttpStatusCode(v int32) *QueryChatInstanceSessionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryChatInstanceSessionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryChatInstanceSessionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryChatInstanceSessionsResponseBody
	GetSuccess() *bool
}

type QueryChatInstanceSessionsResponseBody struct {
	// example:
	//
	// 200
	Code *string            `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ChatSessionInfo `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 7239F9E5-A4DB-55BA-B701-0CE47DBDB0A8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s QueryChatInstanceSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryChatInstanceSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryChatInstanceSessionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryChatInstanceSessionsResponseBody) GetData() []*ChatSessionInfo {
	return s.Data
}

func (s *QueryChatInstanceSessionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryChatInstanceSessionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryChatInstanceSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryChatInstanceSessionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryChatInstanceSessionsResponseBody) SetCode(v string) *QueryChatInstanceSessionsResponseBody {
	s.Code = &v
	return s
}

func (s *QueryChatInstanceSessionsResponseBody) SetData(v []*ChatSessionInfo) *QueryChatInstanceSessionsResponseBody {
	s.Data = v
	return s
}

func (s *QueryChatInstanceSessionsResponseBody) SetHttpStatusCode(v int32) *QueryChatInstanceSessionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryChatInstanceSessionsResponseBody) SetMessage(v string) *QueryChatInstanceSessionsResponseBody {
	s.Message = &v
	return s
}

func (s *QueryChatInstanceSessionsResponseBody) SetRequestId(v string) *QueryChatInstanceSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryChatInstanceSessionsResponseBody) SetSuccess(v bool) *QueryChatInstanceSessionsResponseBody {
	s.Success = &v
	return s
}

func (s *QueryChatInstanceSessionsResponseBody) Validate() error {
	return dara.Validate(s)
}
