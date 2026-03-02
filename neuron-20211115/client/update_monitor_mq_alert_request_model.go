// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMonitorMqAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorMqAlertUpdateCmd) *UpdateMonitorMqAlertRequest
	GetBody() *MonitorMqAlertUpdateCmd
}

type UpdateMonitorMqAlertRequest struct {
	// This parameter is required.
	Body *MonitorMqAlertUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMonitorMqAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMonitorMqAlertRequest) GoString() string {
	return s.String()
}

func (s *UpdateMonitorMqAlertRequest) GetBody() *MonitorMqAlertUpdateCmd {
	return s.Body
}

func (s *UpdateMonitorMqAlertRequest) SetBody(v *MonitorMqAlertUpdateCmd) *UpdateMonitorMqAlertRequest {
	s.Body = v
	return s
}

func (s *UpdateMonitorMqAlertRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
