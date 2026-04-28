// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentity(v *Identity) *ListIdentityRoleRequest
	GetIdentity() *Identity
}

type ListIdentityRoleRequest struct {
	// This parameter is required.
	Identity *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
}

func (s ListIdentityRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityRoleRequest) GoString() string {
	return s.String()
}

func (s *ListIdentityRoleRequest) GetIdentity() *Identity {
	return s.Identity
}

func (s *ListIdentityRoleRequest) SetIdentity(v *Identity) *ListIdentityRoleRequest {
	s.Identity = v
	return s
}

func (s *ListIdentityRoleRequest) Validate() error {
	if s.Identity != nil {
		if err := s.Identity.Validate(); err != nil {
			return err
		}
	}
	return nil
}
