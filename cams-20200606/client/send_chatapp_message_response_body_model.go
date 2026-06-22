// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatappMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SendChatappMessageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SendChatappMessageResponseBody
	GetCode() *string
	SetMessage(v string) *SendChatappMessageResponseBody
	GetMessage() *string
	SetMessageId(v string) *SendChatappMessageResponseBody
	GetMessageId() *string
	SetRequestId(v string) *SendChatappMessageResponseBody
	GetRequestId() *string
}

type SendChatappMessageResponseBody struct {
	// Details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code of the request.
	//
	// - A value of OK indicates a successful request.
	//
	// - For other status codes, see [error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the message.
	//
	// example:
	//
	// 61851ccb2f1365b16aee****
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendChatappMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendChatappMessageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SendChatappMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendChatappMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendChatappMessageResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *SendChatappMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendChatappMessageResponseBody) SetAccessDeniedDetail(v string) *SendChatappMessageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SendChatappMessageResponseBody) SetCode(v string) *SendChatappMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendChatappMessageResponseBody) SetMessage(v string) *SendChatappMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendChatappMessageResponseBody) SetMessageId(v string) *SendChatappMessageResponseBody {
	s.MessageId = &v
	return s
}

func (s *SendChatappMessageResponseBody) SetRequestId(v string) *SendChatappMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendChatappMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
