// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertWebhooksShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWebhookIdsShrink(v string) *DeleteAlertWebhooksShrinkRequest
	GetWebhookIdsShrink() *string
}

type DeleteAlertWebhooksShrinkRequest struct {
	// This parameter is required.
	WebhookIdsShrink *string `json:"webhookIds,omitempty" xml:"webhookIds,omitempty"`
}

func (s DeleteAlertWebhooksShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertWebhooksShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertWebhooksShrinkRequest) GetWebhookIdsShrink() *string {
	return s.WebhookIdsShrink
}

func (s *DeleteAlertWebhooksShrinkRequest) SetWebhookIdsShrink(v string) *DeleteAlertWebhooksShrinkRequest {
	s.WebhookIdsShrink = &v
	return s
}

func (s *DeleteAlertWebhooksShrinkRequest) Validate() error {
	return dara.Validate(s)
}
