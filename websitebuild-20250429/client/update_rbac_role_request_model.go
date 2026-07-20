// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRbacRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *UpdateRbacRoleRequest
	GetBizId() *string
	SetRoleData(v string) *UpdateRbacRoleRequest
	GetRoleData() *string
	SetRoleId(v string) *UpdateRbacRoleRequest
	GetRoleId() *string
}

type UpdateRbacRoleRequest struct {
	// The application instance ID.
	//
	// example:
	//
	// WD20250814102215000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The role data.
	//
	// example:
	//
	// {"name":"customer","label":"xx1
	//
	// ","is_default":true,"is_system":false}
	RoleData *string `json:"RoleData,omitempty" xml:"RoleData,omitempty"`
	// The role ID.
	//
	// example:
	//
	// agent@ly-online
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s UpdateRbacRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRbacRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRbacRoleRequest) GetBizId() *string {
	return s.BizId
}

func (s *UpdateRbacRoleRequest) GetRoleData() *string {
	return s.RoleData
}

func (s *UpdateRbacRoleRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *UpdateRbacRoleRequest) SetBizId(v string) *UpdateRbacRoleRequest {
	s.BizId = &v
	return s
}

func (s *UpdateRbacRoleRequest) SetRoleData(v string) *UpdateRbacRoleRequest {
	s.RoleData = &v
	return s
}

func (s *UpdateRbacRoleRequest) SetRoleId(v string) *UpdateRbacRoleRequest {
	s.RoleId = &v
	return s
}

func (s *UpdateRbacRoleRequest) Validate() error {
	return dara.Validate(s)
}
