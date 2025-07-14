// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWorkspaceRoleConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleId(v int64) *QueryWorkspaceRoleConfigRequest
	GetRoleId() *int64
}

type QueryWorkspaceRoleConfigRequest struct {
	// Workspace role ID, including predefined roles and custom roles:
	//
	// - 25: Workspace Administrator (predefined role)
	//
	// - 26: Developer (predefined role)
	//
	// - 27: Analyst (predefined role)
	//
	// - 30: Viewer (predefined role)
	//
	// - Custom role: The corresponding role ID for the custom role
	//
	// This parameter is required.
	//
	// example:
	//
	// 25
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
}

func (s QueryWorkspaceRoleConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryWorkspaceRoleConfigRequest) GoString() string {
	return s.String()
}

func (s *QueryWorkspaceRoleConfigRequest) GetRoleId() *int64 {
	return s.RoleId
}

func (s *QueryWorkspaceRoleConfigRequest) SetRoleId(v int64) *QueryWorkspaceRoleConfigRequest {
	s.RoleId = &v
	return s
}

func (s *QueryWorkspaceRoleConfigRequest) Validate() error {
	return dara.Validate(s)
}
