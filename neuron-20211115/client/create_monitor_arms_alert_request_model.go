// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorArmsAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorArmsAlert) *CreateMonitorArmsAlertRequest
	GetBody() *MonitorArmsAlert
}

type CreateMonitorArmsAlertRequest struct {
	// This parameter is required.
	Body *MonitorArmsAlert `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorArmsAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorArmsAlertRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorArmsAlertRequest) GetBody() *MonitorArmsAlert {
	return s.Body
}

func (s *CreateMonitorArmsAlertRequest) SetBody(v *MonitorArmsAlert) *CreateMonitorArmsAlertRequest {
	s.Body = v
	return s
}

func (s *CreateMonitorArmsAlertRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
