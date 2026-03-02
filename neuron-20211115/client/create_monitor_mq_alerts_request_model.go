// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorMqAlertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorMqAlerts) *CreateMonitorMqAlertsRequest
	GetBody() *MonitorMqAlerts
}

type CreateMonitorMqAlertsRequest struct {
	// This parameter is required.
	Body *MonitorMqAlerts `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorMqAlertsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorMqAlertsRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorMqAlertsRequest) GetBody() *MonitorMqAlerts {
	return s.Body
}

func (s *CreateMonitorMqAlertsRequest) SetBody(v *MonitorMqAlerts) *CreateMonitorMqAlertsRequest {
	s.Body = v
	return s
}

func (s *CreateMonitorMqAlertsRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
