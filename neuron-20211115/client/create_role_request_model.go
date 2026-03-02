// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Role) *CreateRoleRequest
	GetBody() *Role
}

type CreateRoleRequest struct {
	// This parameter is required.
	Body *Role `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleRequest) GetBody() *Role {
	return s.Body
}

func (s *CreateRoleRequest) SetBody(v *Role) *CreateRoleRequest {
	s.Body = v
	return s
}

func (s *CreateRoleRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
