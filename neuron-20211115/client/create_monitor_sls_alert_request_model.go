// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorSlsAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorSlsAlert) *CreateMonitorSlsAlertRequest
	GetBody() *MonitorSlsAlert
}

type CreateMonitorSlsAlertRequest struct {
	// This parameter is required.
	Body *MonitorSlsAlert `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorSlsAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorSlsAlertRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorSlsAlertRequest) GetBody() *MonitorSlsAlert {
	return s.Body
}

func (s *CreateMonitorSlsAlertRequest) SetBody(v *MonitorSlsAlert) *CreateMonitorSlsAlertRequest {
	s.Body = v
	return s
}

func (s *CreateMonitorSlsAlertRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
