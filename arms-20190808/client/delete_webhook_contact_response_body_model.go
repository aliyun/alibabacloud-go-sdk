// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebhookContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *DeleteWebhookContactResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteWebhookContactResponseBody
	GetRequestId() *string
}

type DeleteWebhookContactResponseBody struct {
	// Indicates whether the webhook alert contact was deleted.
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C21AB7CF-B7AF-410F-BD61-82D1567F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebhookContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebhookContactResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebhookContactResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteWebhookContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWebhookContactResponseBody) SetIsSuccess(v bool) *DeleteWebhookContactResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteWebhookContactResponseBody) SetRequestId(v string) *DeleteWebhookContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWebhookContactResponseBody) Validate() error {
	return dara.Validate(s)
}
