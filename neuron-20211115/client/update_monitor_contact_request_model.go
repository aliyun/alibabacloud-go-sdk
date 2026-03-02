// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMonitorContactRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorContactUpdateCmd) *UpdateMonitorContactRequest
	GetBody() *MonitorContactUpdateCmd
}

type UpdateMonitorContactRequest struct {
	// This parameter is required.
	Body *MonitorContactUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMonitorContactRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMonitorContactRequest) GoString() string {
	return s.String()
}

func (s *UpdateMonitorContactRequest) GetBody() *MonitorContactUpdateCmd {
	return s.Body
}

func (s *UpdateMonitorContactRequest) SetBody(v *MonitorContactUpdateCmd) *UpdateMonitorContactRequest {
	s.Body = v
	return s
}

func (s *UpdateMonitorContactRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
