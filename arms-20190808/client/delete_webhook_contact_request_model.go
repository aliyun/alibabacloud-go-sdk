// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebhookContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWebhookId(v int64) *DeleteWebhookContactRequest
	GetWebhookId() *int64
}

type DeleteWebhookContactRequest struct {
	// The ID of the webhook alert contact.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	WebhookId *int64 `json:"WebhookId,omitempty" xml:"WebhookId,omitempty"`
}

func (s DeleteWebhookContactRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebhookContactRequest) GoString() string {
	return s.String()
}

func (s *DeleteWebhookContactRequest) GetWebhookId() *int64 {
	return s.WebhookId
}

func (s *DeleteWebhookContactRequest) SetWebhookId(v int64) *DeleteWebhookContactRequest {
	s.WebhookId = &v
	return s
}

func (s *DeleteWebhookContactRequest) Validate() error {
	return dara.Validate(s)
}
