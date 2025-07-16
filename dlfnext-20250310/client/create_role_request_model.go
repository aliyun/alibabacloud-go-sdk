// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateRoleRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateRoleRequest
	GetDisplayName() *string
	SetRoleName(v string) *CreateRoleRequest
	GetRoleName() *string
}

type CreateRoleRequest struct {
	// example:
	//
	// role_description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// role_display_name
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// role_name
	RoleName *string `json:"roleName,omitempty" xml:"roleName,omitempty"`
}

func (s CreateRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRoleRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateRoleRequest) SetDescription(v string) *CreateRoleRequest {
	s.Description = &v
	return s
}

func (s *CreateRoleRequest) SetDisplayName(v string) *CreateRoleRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateRoleRequest) SetRoleName(v string) *CreateRoleRequest {
	s.RoleName = &v
	return s
}

func (s *CreateRoleRequest) Validate() error {
	return dara.Validate(s)
}
