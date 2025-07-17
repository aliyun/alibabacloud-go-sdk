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
	// The result returned. Valid values:
	//
	// 	- `true`: The modification is successful.
	//
	// 	- `false`: The modification fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 16AF921B-8187-489F-9913-43C808B4****
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
