// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *RoleInfoUpdateCmd) *UpdateRoleRequest
	GetBody() *RoleInfoUpdateCmd
}

type UpdateRoleRequest struct {
	// This parameter is required.
	Body *RoleInfoUpdateCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleRequest) GetBody() *RoleInfoUpdateCmd {
	return s.Body
}

func (s *UpdateRoleRequest) SetBody(v *RoleInfoUpdateCmd) *UpdateRoleRequest {
	s.Body = v
	return s
}

func (s *UpdateRoleRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
