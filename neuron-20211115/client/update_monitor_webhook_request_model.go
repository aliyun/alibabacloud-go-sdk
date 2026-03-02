// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMonitorWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorWebhookUpdateCmd) *UpdateMonitorWebhookRequest
	GetBody() *MonitorWebhookUpdateCmd
}

type UpdateMonitorWebhookRequest struct {
	// This parameter is required.
	Body *MonitorWebhookUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMonitorWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMonitorWebhookRequest) GoString() string {
	return s.String()
}

func (s *UpdateMonitorWebhookRequest) GetBody() *MonitorWebhookUpdateCmd {
	return s.Body
}

func (s *UpdateMonitorWebhookRequest) SetBody(v *MonitorWebhookUpdateCmd) *UpdateMonitorWebhookRequest {
	s.Body = v
	return s
}

func (s *UpdateMonitorWebhookRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
