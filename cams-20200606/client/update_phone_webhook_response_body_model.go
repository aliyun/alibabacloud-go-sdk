// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePhoneWebhookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdatePhoneWebhookResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdatePhoneWebhookResponseBody
	GetCode() *string
	SetMessage(v string) *UpdatePhoneWebhookResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePhoneWebhookResponseBody
	GetRequestId() *string
}

type UpdatePhoneWebhookResponseBody struct {
	// Access denied for detailed information.
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
	// Prompt message, there is a value when an exception is returned.
	//
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePhoneWebhookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePhoneWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePhoneWebhookResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdatePhoneWebhookResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdatePhoneWebhookResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePhoneWebhookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePhoneWebhookResponseBody) SetAccessDeniedDetail(v string) *UpdatePhoneWebhookResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdatePhoneWebhookResponseBody) SetCode(v string) *UpdatePhoneWebhookResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePhoneWebhookResponseBody) SetMessage(v string) *UpdatePhoneWebhookResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePhoneWebhookResponseBody) SetRequestId(v string) *UpdatePhoneWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePhoneWebhookResponseBody) Validate() error {
	return dara.Validate(s)
}
