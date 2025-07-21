// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendChatappMassMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SendChatappMassMessageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SendChatappMassMessageResponseBody
	GetCode() *string
	SetGroupMessageId(v string) *SendChatappMassMessageResponseBody
	GetGroupMessageId() *string
	SetMessage(v string) *SendChatappMassMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendChatappMassMessageResponseBody
	GetRequestId() *string
}

type SendChatappMassMessageResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the message group.
	//
	// example:
	//
	// 890000010002****
	GroupMessageId *string `json:"GroupMessageId,omitempty" xml:"GroupMessageId,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// User not authorized to operate on the specified resource.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendChatappMassMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendChatappMassMessageResponseBody) GoString() string {
	return s.String()
}

func (s *SendChatappMassMessageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SendChatappMassMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendChatappMassMessageResponseBody) GetGroupMessageId() *string {
	return s.GroupMessageId
}

func (s *SendChatappMassMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendChatappMassMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendChatappMassMessageResponseBody) SetAccessDeniedDetail(v string) *SendChatappMassMessageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SendChatappMassMessageResponseBody) SetCode(v string) *SendChatappMassMessageResponseBody {
	s.Code = &v
	return s
}

func (s *SendChatappMassMessageResponseBody) SetGroupMessageId(v string) *SendChatappMassMessageResponseBody {
	s.GroupMessageId = &v
	return s
}

func (s *SendChatappMassMessageResponseBody) SetMessage(v string) *SendChatappMassMessageResponseBody {
	s.Message = &v
	return s
}

func (s *SendChatappMassMessageResponseBody) SetRequestId(v string) *SendChatappMassMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendChatappMassMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
