// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebhookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *UpdateWebhookResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *UpdateWebhookResponseBody
	GetRequestId() *string
}

type UpdateWebhookResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateWebhookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebhookResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWebhookResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *UpdateWebhookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWebhookResponseBody) SetIsSuccess(v bool) *UpdateWebhookResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *UpdateWebhookResponseBody) SetRequestId(v string) *UpdateWebhookResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWebhookResponseBody) Validate() error {
	return dara.Validate(s)
}
