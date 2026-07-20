// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRbacRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *CreateRbacRoleRequest
	GetBizId() *string
	SetRoleData(v string) *CreateRbacRoleRequest
	GetRoleData() *string
}

type CreateRbacRoleRequest struct {
	// The business ID of the customer.
	//
	// example:
	//
	// WS20250731233102000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The role data.
	//
	// example:
	//
	// {"name":"customer","label":"xx
	//
	// ","is_default":true,"is_system":false}
	RoleData *string `json:"RoleData,omitempty" xml:"RoleData,omitempty"`
}

func (s CreateRbacRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRbacRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRbacRoleRequest) GetBizId() *string {
	return s.BizId
}

func (s *CreateRbacRoleRequest) GetRoleData() *string {
	return s.RoleData
}

func (s *CreateRbacRoleRequest) SetBizId(v string) *CreateRbacRoleRequest {
	s.BizId = &v
	return s
}

func (s *CreateRbacRoleRequest) SetRoleData(v string) *CreateRbacRoleRequest {
	s.RoleData = &v
	return s
}

func (s *CreateRbacRoleRequest) Validate() error {
	return dara.Validate(s)
}
