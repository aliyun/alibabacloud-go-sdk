// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceUserRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleId(v int64) *UpdateWorkspaceUserRoleRequest
	GetRoleId() *int64
	SetRoleIds(v string) *UpdateWorkspaceUserRoleRequest
	GetRoleIds() *string
	SetUserId(v string) *UpdateWorkspaceUserRoleRequest
	GetUserId() *string
	SetWorkspaceId(v string) *UpdateWorkspaceUserRoleRequest
	GetWorkspaceId() *string
}

type UpdateWorkspaceUserRoleRequest struct {
	// Deprecated
	//
	// The ID of the preset workspace role. This operation overwrites existing roles. Valid values:
	//
	// - 25: workspace administrator
	//
	// - 26: workspace developer
	//
	// - 27: workspace analyst
	//
	// - 30: workspace viewer
	//
	// example:
	//
	// 25
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The IDs of multiple workspace roles, separated by commas. If this parameter is specified, it takes precedence. <notice>roleId and roleIds cannot be both empty.</notice>
	//
	// example:
	//
	// 25,26
	RoleIds *string `json:"RoleIds,omitempty" xml:"RoleIds,omitempty"`
	// The Quick BI user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f5698bedeb384b1986afccd9e434****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateWorkspaceUserRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceUserRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUserRoleRequest) GetRoleId() *int64 {
	return s.RoleId
}

func (s *UpdateWorkspaceUserRoleRequest) GetRoleIds() *string {
	return s.RoleIds
}

func (s *UpdateWorkspaceUserRoleRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateWorkspaceUserRoleRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateWorkspaceUserRoleRequest) SetRoleId(v int64) *UpdateWorkspaceUserRoleRequest {
	s.RoleId = &v
	return s
}

func (s *UpdateWorkspaceUserRoleRequest) SetRoleIds(v string) *UpdateWorkspaceUserRoleRequest {
	s.RoleIds = &v
	return s
}

func (s *UpdateWorkspaceUserRoleRequest) SetUserId(v string) *UpdateWorkspaceUserRoleRequest {
	s.UserId = &v
	return s
}

func (s *UpdateWorkspaceUserRoleRequest) SetWorkspaceId(v string) *UpdateWorkspaceUserRoleRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateWorkspaceUserRoleRequest) Validate() error {
	return dara.Validate(s)
}
