// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorWebhook) *CreateMonitorWebhookRequest
	GetBody() *MonitorWebhook
}

type CreateMonitorWebhookRequest struct {
	// This parameter is required.
	Body *MonitorWebhook `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorWebhookRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorWebhookRequest) GetBody() *MonitorWebhook {
	return s.Body
}

func (s *CreateMonitorWebhookRequest) SetBody(v *MonitorWebhook) *CreateMonitorWebhookRequest {
	s.Body = v
	return s
}

func (s *CreateMonitorWebhookRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
