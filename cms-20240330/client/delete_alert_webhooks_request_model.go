// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertWebhooksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetWebhookIds(v []*string) *DeleteAlertWebhooksRequest
	GetWebhookIds() []*string
}

type DeleteAlertWebhooksRequest struct {
	// This parameter is required.
	WebhookIds []*string `json:"webhookIds,omitempty" xml:"webhookIds,omitempty" type:"Repeated"`
}

func (s DeleteAlertWebhooksRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertWebhooksRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertWebhooksRequest) GetWebhookIds() []*string {
	return s.WebhookIds
}

func (s *DeleteAlertWebhooksRequest) SetWebhookIds(v []*string) *DeleteAlertWebhooksRequest {
	s.WebhookIds = v
	return s
}

func (s *DeleteAlertWebhooksRequest) Validate() error {
	return dara.Validate(s)
}
