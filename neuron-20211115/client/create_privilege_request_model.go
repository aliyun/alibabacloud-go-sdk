// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreatePrivilegeCmd) *CreatePrivilegeRequest
	GetBody() *CreatePrivilegeCmd
}

type CreatePrivilegeRequest struct {
	Body *CreatePrivilegeCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivilegeRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivilegeRequest) GetBody() *CreatePrivilegeCmd {
	return s.Body
}

func (s *CreatePrivilegeRequest) SetBody(v *CreatePrivilegeCmd) *CreatePrivilegeRequest {
	s.Body = v
	return s
}

func (s *CreatePrivilegeRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
