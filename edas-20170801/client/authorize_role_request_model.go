// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleIds(v string) *AuthorizeRoleRequest
	GetRoleIds() *string
	SetTargetUserId(v string) *AuthorizeRoleRequest
	GetTargetUserId() *string
}

type AuthorizeRoleRequest struct {
	// The ID of the role to be assigned. If you want to assign multiple roles to the specified RAM user, separate the IDs of the roles with semicolons (;). If you leave this parameter empty, the roles assigned to the specified RAM user are revoked.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1;2
	RoleIds *string `json:"RoleIds,omitempty" xml:"RoleIds,omitempty"`
	// The ID of the RAM user to which you want to assign the roles.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@133******
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s AuthorizeRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeRoleRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeRoleRequest) GetRoleIds() *string {
	return s.RoleIds
}

func (s *AuthorizeRoleRequest) GetTargetUserId() *string {
	return s.TargetUserId
}

func (s *AuthorizeRoleRequest) SetRoleIds(v string) *AuthorizeRoleRequest {
	s.RoleIds = &v
	return s
}

func (s *AuthorizeRoleRequest) SetTargetUserId(v string) *AuthorizeRoleRequest {
	s.TargetUserId = &v
	return s
}

func (s *AuthorizeRoleRequest) Validate() error {
	return dara.Validate(s)
}
