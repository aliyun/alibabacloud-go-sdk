// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorMqAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorMqAlert) *CreateMonitorMqAlertRequest
	GetBody() *MonitorMqAlert
}

type CreateMonitorMqAlertRequest struct {
	// This parameter is required.
	Body *MonitorMqAlert `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorMqAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorMqAlertRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorMqAlertRequest) GetBody() *MonitorMqAlert {
	return s.Body
}

func (s *CreateMonitorMqAlertRequest) SetBody(v *MonitorMqAlert) *CreateMonitorMqAlertRequest {
	s.Body = v
	return s
}

func (s *CreateMonitorMqAlertRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
