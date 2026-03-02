// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMonitorSlsAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorSlsAlertUpdateCmd) *UpdateMonitorSlsAlertRequest
	GetBody() *MonitorSlsAlertUpdateCmd
}

type UpdateMonitorSlsAlertRequest struct {
	// This parameter is required.
	Body *MonitorSlsAlertUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMonitorSlsAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMonitorSlsAlertRequest) GoString() string {
	return s.String()
}

func (s *UpdateMonitorSlsAlertRequest) GetBody() *MonitorSlsAlertUpdateCmd {
	return s.Body
}

func (s *UpdateMonitorSlsAlertRequest) SetBody(v *MonitorSlsAlertUpdateCmd) *UpdateMonitorSlsAlertRequest {
	s.Body = v
	return s
}

func (s *UpdateMonitorSlsAlertRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
