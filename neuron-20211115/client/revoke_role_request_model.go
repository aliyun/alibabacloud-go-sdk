// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *RoleRevokeCmd) *RevokeRoleRequest
	GetBody() *RoleRevokeCmd
}

type RevokeRoleRequest struct {
	// This parameter is required.
	Body *RoleRevokeCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeRoleRequest) GoString() string {
	return s.String()
}

func (s *RevokeRoleRequest) GetBody() *RoleRevokeCmd {
	return s.Body
}

func (s *RevokeRoleRequest) SetBody(v *RoleRevokeCmd) *RevokeRoleRequest {
	s.Body = v
	return s
}

func (s *RevokeRoleRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
