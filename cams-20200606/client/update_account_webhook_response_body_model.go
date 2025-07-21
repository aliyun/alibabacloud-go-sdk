// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAccountWebhookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateAccountWebhookResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateAccountWebhookResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateAccountWebhookResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAccountWebhookResponseBody
	GetRequestId() *string
}

type UpdateAccountWebhookResponseBody struct {
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
	// The error message returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 744c4b5c79c9432497a075bdfca3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAccountWebhookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAccountWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAccountWebhookResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateAccountWebhookResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAccountWebhookResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAccountWebhookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAccountWebhookResponseBody) SetAccessDeniedDetail(v string) *UpdateAccountWebhookResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateAccountWebhookResponseBody) SetCode(v string) *UpdateAccountWebhookResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAccountWebhookResponseBody) SetMessage(v string) *UpdateAccountWebhookResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAccountWebhookResponseBody) SetRequestId(v string) *UpdateAccountWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAccountWebhookResponseBody) Validate() error {
	return dara.Validate(s)
}
