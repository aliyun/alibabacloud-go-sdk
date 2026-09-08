// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebhookContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateWebhookContactResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateWebhookContactResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateWebhookContactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWebhookContactResponseBody
	GetSuccess() *bool
}

type UpdateWebhookContactResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// /
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateWebhookContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebhookContactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWebhookContactResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateWebhookContactResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWebhookContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWebhookContactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWebhookContactResponseBody) SetCode(v string) *UpdateWebhookContactResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWebhookContactResponseBody) SetMessage(v string) *UpdateWebhookContactResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWebhookContactResponseBody) SetRequestId(v string) *UpdateWebhookContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWebhookContactResponseBody) SetSuccess(v bool) *UpdateWebhookContactResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWebhookContactResponseBody) Validate() error {
	return dara.Validate(s)
}
