// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceIdentityRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleName(v string) *CreateServiceIdentityRoleRequest
	GetRoleName() *string
}

type CreateServiceIdentityRoleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// AliyunServiceRoleForFeatureStore
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s CreateServiceIdentityRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceIdentityRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceIdentityRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateServiceIdentityRoleRequest) SetRoleName(v string) *CreateServiceIdentityRoleRequest {
	s.RoleName = &v
	return s
}

func (s *CreateServiceIdentityRoleRequest) Validate() error {
	return dara.Validate(s)
}
