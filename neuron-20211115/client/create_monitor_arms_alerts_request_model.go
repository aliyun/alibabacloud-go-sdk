// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorArmsAlertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorArmsAlerts) *CreateMonitorArmsAlertsRequest
	GetBody() *MonitorArmsAlerts
}

type CreateMonitorArmsAlertsRequest struct {
	// This parameter is required.
	Body *MonitorArmsAlerts `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorArmsAlertsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorArmsAlertsRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorArmsAlertsRequest) GetBody() *MonitorArmsAlerts {
	return s.Body
}

func (s *CreateMonitorArmsAlertsRequest) SetBody(v *MonitorArmsAlerts) *CreateMonitorArmsAlertsRequest {
	s.Body = v
	return s
}

func (s *CreateMonitorArmsAlertsRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
