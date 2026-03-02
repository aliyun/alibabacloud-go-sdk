// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivilegePopRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreatePrivilegePopCmd) *CreatePrivilegePopRequest
	GetBody() *CreatePrivilegePopCmd
}

type CreatePrivilegePopRequest struct {
	Body *CreatePrivilegePopCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrivilegePopRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivilegePopRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivilegePopRequest) GetBody() *CreatePrivilegePopCmd {
	return s.Body
}

func (s *CreatePrivilegePopRequest) SetBody(v *CreatePrivilegePopCmd) *CreatePrivilegePopRequest {
	s.Body = v
	return s
}

func (s *CreatePrivilegePopRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
