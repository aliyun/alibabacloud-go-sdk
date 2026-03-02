// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorGroupMemberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *MonitorGroupMemberCreateCmd) *CreateMonitorGroupMemberRequest
	GetBody() *MonitorGroupMemberCreateCmd
}

type CreateMonitorGroupMemberRequest struct {
	// This parameter is required.
	Body *MonitorGroupMemberCreateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorGroupMemberRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *CreateMonitorGroupMemberRequest) GetBody() *MonitorGroupMemberCreateCmd {
	return s.Body
}

func (s *CreateMonitorGroupMemberRequest) SetBody(v *MonitorGroupMemberCreateCmd) *CreateMonitorGroupMemberRequest {
	s.Body = v
	return s
}

func (s *CreateMonitorGroupMemberRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
