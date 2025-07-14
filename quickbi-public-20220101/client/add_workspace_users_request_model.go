// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleId(v int64) *AddWorkspaceUsersRequest
	GetRoleId() *int64
	SetUserIds(v string) *AddWorkspaceUsersRequest
	GetUserIds() *string
	SetWorkspaceId(v string) *AddWorkspaceUsersRequest
	GetWorkspaceId() *string
}

type AddWorkspaceUsersRequest struct {
	// Preset space role ID. Value range:
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
	// User ID. This is the UserID for Quick BI, not the Alibaba Cloud UID.
	//
	// - Supports batch parameters, with user IDs separated by commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
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

func (s AddWorkspaceUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceUsersRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceUsersRequest) GetRoleId() *int64 {
	return s.RoleId
}

func (s *AddWorkspaceUsersRequest) GetUserIds() *string {
	return s.UserIds
}

func (s *AddWorkspaceUsersRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddWorkspaceUsersRequest) SetRoleId(v int64) *AddWorkspaceUsersRequest {
	s.RoleId = &v
	return s
}

func (s *AddWorkspaceUsersRequest) SetUserIds(v string) *AddWorkspaceUsersRequest {
	s.UserIds = &v
	return s
}

func (s *AddWorkspaceUsersRequest) SetWorkspaceId(v string) *AddWorkspaceUsersRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddWorkspaceUsersRequest) Validate() error {
	return dara.Validate(s)
}
