// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMonitorContactGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorContactGroupUpdateCmd) *UpdateMonitorContactGroupRequest
	GetBody() *MonitorContactGroupUpdateCmd
}

type UpdateMonitorContactGroupRequest struct {
	// This parameter is required.
	Body *MonitorContactGroupUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMonitorContactGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMonitorContactGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateMonitorContactGroupRequest) GetBody() *MonitorContactGroupUpdateCmd {
	return s.Body
}

func (s *UpdateMonitorContactGroupRequest) SetBody(v *MonitorContactGroupUpdateCmd) *UpdateMonitorContactGroupRequest {
	s.Body = v
	return s
}

func (s *UpdateMonitorContactGroupRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
