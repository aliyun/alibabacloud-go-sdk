// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleId(v int32) *DeleteRoleRequest
	GetRoleId() *int32
}

type DeleteRoleRequest struct {
	// The ID of the RAM role.
	//
	// This parameter is required.
	//
	// example:
	//
	// 99999999
	RoleId *int32 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s DeleteRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoleRequest) GetRoleId() *int32 {
	return s.RoleId
}

func (s *DeleteRoleRequest) SetRoleId(v int32) *DeleteRoleRequest {
	s.RoleId = &v
	return s
}

func (s *DeleteRoleRequest) Validate() error {
	return dara.Validate(s)
}
