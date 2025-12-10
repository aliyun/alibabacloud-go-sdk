// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceLinkedRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleName(v string) *DeleteServiceLinkedRoleRequest
	GetRoleName() *string
}

type DeleteServiceLinkedRoleRequest struct {
	// The name of the role.
	//
	// This parameter is required.
	//
	// example:
	//
	// AliyunServiceRoleForPolarDB
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s DeleteServiceLinkedRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceLinkedRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *DeleteServiceLinkedRoleRequest) SetRoleName(v string) *DeleteServiceLinkedRoleRequest {
	s.RoleName = &v
	return s
}

func (s *DeleteServiceLinkedRoleRequest) Validate() error {
	return dara.Validate(s)
}
