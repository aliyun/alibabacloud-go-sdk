// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesForRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleName(v string) *ListPoliciesForRoleRequest
	GetRoleName() *string
}

type ListPoliciesForRoleRequest struct {
	// The name of the RAM role.
	//
	// example:
	//
	// AdminRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s ListPoliciesForRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesForRoleRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesForRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *ListPoliciesForRoleRequest) SetRoleName(v string) *ListPoliciesForRoleRequest {
	s.RoleName = &v
	return s
}

func (s *ListPoliciesForRoleRequest) Validate() error {
	return dara.Validate(s)
}
