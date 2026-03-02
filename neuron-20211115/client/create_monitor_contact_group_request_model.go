// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorContactGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorContactGroupCreateCmd) *CreateMonitorContactGroupRequest
	GetBody() *MonitorContactGroupCreateCmd
}

type CreateMonitorContactGroupRequest struct {
	// This parameter is required.
	Body *MonitorContactGroupCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorContactGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorContactGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorContactGroupRequest) GetBody() *MonitorContactGroupCreateCmd {
	return s.Body
}

func (s *CreateMonitorContactGroupRequest) SetBody(v *MonitorContactGroupCreateCmd) *CreateMonitorContactGroupRequest {
	s.Body = v
	return s
}

func (s *CreateMonitorContactGroupRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
