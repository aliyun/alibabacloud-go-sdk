// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkspaceRolesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleIds(v []*string) *DeleteWorkspaceRolesRequest
	GetRoleIds() []*string
}

type DeleteWorkspaceRolesRequest struct {
	// The IDs of the roles to delete.
	RoleIds []*string `json:"RoleIds,omitempty" xml:"RoleIds,omitempty" type:"Repeated"`
}

func (s DeleteWorkspaceRolesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkspaceRolesRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkspaceRolesRequest) GetRoleIds() []*string {
	return s.RoleIds
}

func (s *DeleteWorkspaceRolesRequest) SetRoleIds(v []*string) *DeleteWorkspaceRolesRequest {
	s.RoleIds = v
	return s
}

func (s *DeleteWorkspaceRolesRequest) Validate() error {
	return dara.Validate(s)
}
