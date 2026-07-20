// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRbacRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *DeleteRbacRoleRequest
	GetBizId() *string
	SetRoleId(v string) *DeleteRbacRoleRequest
	GetRoleId() *string
}

type DeleteRbacRoleRequest struct {
	// The business ID of the application instance.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The role ID.
	//
	// example:
	//
	// 75b2f16f-35a5-4e90-949f-295ea14a4dc8
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s DeleteRbacRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRbacRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRbacRoleRequest) GetBizId() *string {
	return s.BizId
}

func (s *DeleteRbacRoleRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *DeleteRbacRoleRequest) SetBizId(v string) *DeleteRbacRoleRequest {
	s.BizId = &v
	return s
}

func (s *DeleteRbacRoleRequest) SetRoleId(v string) *DeleteRbacRoleRequest {
	s.RoleId = &v
	return s
}

func (s *DeleteRbacRoleRequest) Validate() error {
	return dara.Validate(s)
}
