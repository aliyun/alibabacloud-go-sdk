// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrganizationRoleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleId(v int64) *QueryOrganizationRoleConfigRequest
	GetRoleId() *int64
}

type QueryOrganizationRoleConfigRequest struct {
	// Organization role ID, including predefined roles and custom roles:
	//
	// - Organization Administrator (predefined role): 111111111
	//
	// - Permission Administrator (predefined role): 111111112
	//
	// - Regular User (predefined role): 111111113
	//
	// - Custom Role: The corresponding role ID of the custom role
	//
	// This parameter is required.
	//
	// example:
	//
	// 111111111
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s QueryOrganizationRoleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOrganizationRoleConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryOrganizationRoleConfigRequest) GetRoleId() *int64 {
	return s.RoleId
}

func (s *QueryOrganizationRoleConfigRequest) SetRoleId(v int64) *QueryOrganizationRoleConfigRequest {
	s.RoleId = &v
	return s
}

func (s *QueryOrganizationRoleConfigRequest) Validate() error {
	return dara.Validate(s)
}
