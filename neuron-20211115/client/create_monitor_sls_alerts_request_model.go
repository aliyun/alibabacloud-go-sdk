// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorSlsAlertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorSlsAlerts) *CreateMonitorSlsAlertsRequest
	GetBody() *MonitorSlsAlerts
}

type CreateMonitorSlsAlertsRequest struct {
	// This parameter is required.
	Body *MonitorSlsAlerts `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorSlsAlertsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorSlsAlertsRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorSlsAlertsRequest) GetBody() *MonitorSlsAlerts {
	return s.Body
}

func (s *CreateMonitorSlsAlertsRequest) SetBody(v *MonitorSlsAlerts) *CreateMonitorSlsAlertsRequest {
	s.Body = v
	return s
}

func (s *CreateMonitorSlsAlertsRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
