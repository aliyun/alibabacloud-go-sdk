// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleName(v string) *GetRoleRequest
	GetRoleName() *string
	SetUserPoolName(v string) *GetRoleRequest
	GetUserPoolName() *string
}

type GetRoleRequest struct {
	// example:
	//
	// Analyst
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s GetRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRoleRequest) GoString() string {
	return s.String()
}

func (s *GetRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *GetRoleRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *GetRoleRequest) SetRoleName(v string) *GetRoleRequest {
	s.RoleName = &v
	return s
}

func (s *GetRoleRequest) SetUserPoolName(v string) *GetRoleRequest {
	s.UserPoolName = &v
	return s
}

func (s *GetRoleRequest) Validate() error {
	return dara.Validate(s)
}
