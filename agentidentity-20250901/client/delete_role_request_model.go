// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleName(v string) *DeleteRoleRequest
	GetRoleName() *string
	SetUserPoolName(v string) *DeleteRoleRequest
	GetUserPoolName() *string
}

type DeleteRoleRequest struct {
	// example:
	//
	// Analyst
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s DeleteRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *DeleteRoleRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *DeleteRoleRequest) SetRoleName(v string) *DeleteRoleRequest {
	s.RoleName = &v
	return s
}

func (s *DeleteRoleRequest) SetUserPoolName(v string) *DeleteRoleRequest {
	s.UserPoolName = &v
	return s
}

func (s *DeleteRoleRequest) Validate() error {
	return dara.Validate(s)
}
