// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseChatInstanceSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CloseChatInstanceSessionsResponseBody
	GetCode() *string
	SetData(v []*ChatSessionInfo) *CloseChatInstanceSessionsResponseBody
	GetData() []*ChatSessionInfo
	SetHttpStatusCode(v int32) *CloseChatInstanceSessionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CloseChatInstanceSessionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloseChatInstanceSessionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CloseChatInstanceSessionsResponseBody
	GetSuccess() *bool
}

type CloseChatInstanceSessionsResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string            `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ChatSessionInfo `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
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

func (s CloseChatInstanceSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseChatInstanceSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *CloseChatInstanceSessionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloseChatInstanceSessionsResponseBody) GetData() []*ChatSessionInfo {
	return s.Data
}

func (s *CloseChatInstanceSessionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CloseChatInstanceSessionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloseChatInstanceSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseChatInstanceSessionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CloseChatInstanceSessionsResponseBody) SetCode(v string) *CloseChatInstanceSessionsResponseBody {
	s.Code = &v
	return s
}

func (s *CloseChatInstanceSessionsResponseBody) SetData(v []*ChatSessionInfo) *CloseChatInstanceSessionsResponseBody {
	s.Data = v
	return s
}

func (s *CloseChatInstanceSessionsResponseBody) SetHttpStatusCode(v int32) *CloseChatInstanceSessionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CloseChatInstanceSessionsResponseBody) SetMessage(v string) *CloseChatInstanceSessionsResponseBody {
	s.Message = &v
	return s
}

func (s *CloseChatInstanceSessionsResponseBody) SetRequestId(v string) *CloseChatInstanceSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseChatInstanceSessionsResponseBody) SetSuccess(v bool) *CloseChatInstanceSessionsResponseBody {
	s.Success = &v
	return s
}

func (s *CloseChatInstanceSessionsResponseBody) Validate() error {
	return dara.Validate(s)
}
