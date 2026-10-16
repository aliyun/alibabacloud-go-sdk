// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateRoleRequest
	GetDescription() *string
	SetRoleName(v string) *CreateRoleRequest
	GetRoleName() *string
	SetUserPoolName(v string) *CreateRoleRequest
	GetUserPoolName() *string
}

type CreateRoleRequest struct {
	// The description of the role. The description must be 0 to 128 characters in length.
	//
	// example:
	//
	// Data analyst role
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the role. The name must be unique within the user pool and must be 3 to 64 characters in length.
	//
	// example:
	//
	// Analyst
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// The name of the user pool.
	//
	// example:
	//
	// my-agent-userpool
	UserPoolName *string `json:"UserPoolName,omitempty" xml:"UserPoolName,omitempty"`
}

func (s CreateRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *CreateRoleRequest) GetUserPoolName() *string {
	return s.UserPoolName
}

func (s *CreateRoleRequest) SetDescription(v string) *CreateRoleRequest {
	s.Description = &v
	return s
}

func (s *CreateRoleRequest) SetRoleName(v string) *CreateRoleRequest {
	s.RoleName = &v
	return s
}

func (s *CreateRoleRequest) SetUserPoolName(v string) *CreateRoleRequest {
	s.UserPoolName = &v
	return s
}

func (s *CreateRoleRequest) Validate() error {
	return dara.Validate(s)
}
