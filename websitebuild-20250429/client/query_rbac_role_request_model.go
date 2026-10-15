// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRbacRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QueryRbacRoleRequest
	GetBizId() *string
	SetRoleId(v string) *QueryRbacRoleRequest
	GetRoleId() *string
}

type QueryRbacRoleRequest struct {
	// The business ID of the application instance.
	//
	// example:
	//
	// WD20250821161210000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The storage management role ID, including preset roles and custom roles.
	//
	// - 25: storage management administrator (preset role)
	//
	// - 26: developer (preset role)
	//
	// - 27: analyst (preset role)
	//
	// - 30: viewer (preset role)
	//
	// - Custom role: the role ID of the custom role
	//
	// example:
	//
	// 75b2f16f-35a5-4e90-949f-295ea14a4dc8
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s QueryRbacRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRbacRoleRequest) GoString() string {
	return s.String()
}

func (s *QueryRbacRoleRequest) GetBizId() *string {
	return s.BizId
}

func (s *QueryRbacRoleRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *QueryRbacRoleRequest) SetBizId(v string) *QueryRbacRoleRequest {
	s.BizId = &v
	return s
}

func (s *QueryRbacRoleRequest) SetRoleId(v string) *QueryRbacRoleRequest {
	s.RoleId = &v
	return s
}

func (s *QueryRbacRoleRequest) Validate() error {
	return dara.Validate(s)
}
