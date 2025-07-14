// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceUsersRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleId(v int64) *UpdateWorkspaceUsersRoleRequest
	GetRoleId() *int64
	SetUserIds(v string) *UpdateWorkspaceUsersRoleRequest
	GetUserIds() *string
	SetWorkspaceId(v string) *UpdateWorkspaceUsersRoleRequest
	GetWorkspaceId() *string
}

type UpdateWorkspaceUsersRoleRequest struct {
	// Preset space role ID, existing roles will be overwritten. Value range:
	//
	// - 25: Space Administrator
	//
	// - 26: Space Developer
	//
	// - 27: Space Analyst
	//
	// - 30: Space Viewer
	//
	// This parameter is required.
	//
	// example:
	//
	// 25
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// User ID. This is the UserID for Quick BI, not the UID for Alibaba Cloud.
	//
	// - Supports batch parameters, separate user IDs with a comma (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 136516262323****,124498444445****
	UserIds *string `json:"UserIds,omitempty" xml:"UserIds,omitempty"`
	// Workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateWorkspaceUsersRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceUsersRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUsersRoleRequest) GetRoleId() *int64 {
	return s.RoleId
}

func (s *UpdateWorkspaceUsersRoleRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *UpdateWorkspaceUsersRoleRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateWorkspaceUsersRoleRequest) SetRoleId(v int64) *UpdateWorkspaceUsersRoleRequest {
	s.RoleId = &v
	return s
}

func (s *UpdateWorkspaceUsersRoleRequest) SetUserIds(v string) *UpdateWorkspaceUsersRoleRequest {
	s.UserIds = &v
	return s
}

func (s *UpdateWorkspaceUsersRoleRequest) SetWorkspaceId(v string) *UpdateWorkspaceUsersRoleRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateWorkspaceUsersRoleRequest) Validate() error {
	return dara.Validate(s)
}
