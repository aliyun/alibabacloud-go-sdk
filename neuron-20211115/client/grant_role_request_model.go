// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *RoleGrantCmd) *GrantRoleRequest
	GetBody() *RoleGrantCmd
}

type GrantRoleRequest struct {
	// This parameter is required.
	Body *RoleGrantCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantRoleRequest) GoString() string {
	return s.String()
}

func (s *GrantRoleRequest) GetBody() *RoleGrantCmd {
	return s.Body
}

func (s *GrantRoleRequest) SetBody(v *RoleGrantCmd) *GrantRoleRequest {
	s.Body = v
	return s
}

func (s *GrantRoleRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
