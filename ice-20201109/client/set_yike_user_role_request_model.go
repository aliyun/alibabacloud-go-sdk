// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetYikeUserRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleName(v string) *SetYikeUserRoleRequest
	GetRoleName() *string
	SetYikeUserId(v string) *SetYikeUserRoleRequest
	GetYikeUserId() *string
}

type SetYikeUserRoleRequest struct {
	// The user role. Valid values:
	//
	// - SuperAdmin: Super Admin.
	//
	// - Admin: Admin.
	//
	// - RegularUser: Regular User.
	//
	// example:
	//
	// Admin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The sub-account user ID.
	//
	// example:
	//
	// id
	YikeUserId *string `json:"YikeUserId,omitempty" xml:"YikeUserId,omitempty"`
}

func (s SetYikeUserRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s SetYikeUserRoleRequest) GoString() string {
	return s.String()
}

func (s *SetYikeUserRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *SetYikeUserRoleRequest) GetYikeUserId() *string {
	return s.YikeUserId
}

func (s *SetYikeUserRoleRequest) SetRoleName(v string) *SetYikeUserRoleRequest {
	s.RoleName = &v
	return s
}

func (s *SetYikeUserRoleRequest) SetYikeUserId(v string) *SetYikeUserRoleRequest {
	s.YikeUserId = &v
	return s
}

func (s *SetYikeUserRoleRequest) Validate() error {
	return dara.Validate(s)
}
