// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateRoleRequest
	GetDescription() *string
	SetRoleName(v string) *UpdateRoleRequest
	GetRoleName() *string
	SetUserPoolName(v string) *UpdateRoleRequest
	GetUserPoolName() *string
}

type UpdateRoleRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Analyst
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s UpdateRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *UpdateRoleRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *UpdateRoleRequest) SetDescription(v string) *UpdateRoleRequest {
	s.Description = &v
	return s
}

func (s *UpdateRoleRequest) SetRoleName(v string) *UpdateRoleRequest {
	s.RoleName = &v
	return s
}

func (s *UpdateRoleRequest) SetUserPoolName(v string) *UpdateRoleRequest {
	s.UserPoolName = &v
	return s
}

func (s *UpdateRoleRequest) Validate() error {
	return dara.Validate(s)
}
