// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappPhoneNumberDeregisterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ChatappPhoneNumberDeregisterResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ChatappPhoneNumberDeregisterResponseBody
	GetCode() *string
	SetMessage(v string) *ChatappPhoneNumberDeregisterResponseBody
	GetMessage() *string
	SetRequestId(v string) *ChatappPhoneNumberDeregisterResponseBody
	GetRequestId() *string
}

type ChatappPhoneNumberDeregisterResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other values indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChatappPhoneNumberDeregisterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatappPhoneNumberDeregisterResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappPhoneNumberDeregisterResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ChatappPhoneNumberDeregisterResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChatappPhoneNumberDeregisterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChatappPhoneNumberDeregisterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatappPhoneNumberDeregisterResponseBody) SetAccessDeniedDetail(v string) *ChatappPhoneNumberDeregisterResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappPhoneNumberDeregisterResponseBody) SetCode(v string) *ChatappPhoneNumberDeregisterResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappPhoneNumberDeregisterResponseBody) SetMessage(v string) *ChatappPhoneNumberDeregisterResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappPhoneNumberDeregisterResponseBody) SetRequestId(v string) *ChatappPhoneNumberDeregisterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatappPhoneNumberDeregisterResponseBody) Validate() error {
	return dara.Validate(s)
}
