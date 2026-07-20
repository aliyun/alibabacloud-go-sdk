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
	BizId  *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
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
