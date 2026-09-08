// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebhookContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteWebhookContactResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteWebhookContactResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteWebhookContactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteWebhookContactResponseBody
	GetSuccess() *bool
}

type DeleteWebhookContactResponseBody struct {
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
	// /
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// /
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteWebhookContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebhookContactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebhookContactResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteWebhookContactResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteWebhookContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWebhookContactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteWebhookContactResponseBody) SetCode(v string) *DeleteWebhookContactResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWebhookContactResponseBody) SetMessage(v string) *DeleteWebhookContactResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWebhookContactResponseBody) SetRequestId(v string) *DeleteWebhookContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWebhookContactResponseBody) SetSuccess(v bool) *DeleteWebhookContactResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteWebhookContactResponseBody) Validate() error {
	return dara.Validate(s)
}
