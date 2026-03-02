// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMonitorArmsAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorArmsAlertUpdateCmd) *UpdateMonitorArmsAlertRequest
	GetBody() *MonitorArmsAlertUpdateCmd
}

type UpdateMonitorArmsAlertRequest struct {
	// This parameter is required.
	Body *MonitorArmsAlertUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMonitorArmsAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMonitorArmsAlertRequest) GoString() string {
	return s.String()
}

func (s *UpdateMonitorArmsAlertRequest) GetBody() *MonitorArmsAlertUpdateCmd {
	return s.Body
}

func (s *UpdateMonitorArmsAlertRequest) SetBody(v *MonitorArmsAlertUpdateCmd) *UpdateMonitorArmsAlertRequest {
	s.Body = v
	return s
}

func (s *UpdateMonitorArmsAlertRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
