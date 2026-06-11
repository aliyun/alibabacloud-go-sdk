// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRoleId(v int64) *AddUserToWorkspaceRequest
	GetRoleId() *int64
	SetUserId(v string) *AddUserToWorkspaceRequest
	GetUserId() *string
	SetWorkspaceId(v string) *AddUserToWorkspaceRequest
	GetWorkspaceId() *string
}

type AddUserToWorkspaceRequest struct {
	// The ID of the role in the workspace. Predefined roles:
	//
	// - 25: workspace administrator
	//
	// - 26: workspace developer
	//
	// - 27: workspace analyst
	//
	// - 30: workspace viewer
	//
	// > Custom role IDs are supported in version 6.0.3 and later.
	//
	// This parameter is required.
	//
	// example:
	//
	// 25
	RoleId *int64 `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// The ID of the Quick BI user to be added.
	//
	// This parameter is required.
	//
	// example:
	//
	// de4bc5f9429141cc8091cdd1c15b****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The ID of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 95296e95-ca89-4c7d-8af9-dedf0ad0****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddUserToWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserToWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *AddUserToWorkspaceRequest) GetRoleId() *int64 {
	return s.RoleId
}

func (s *AddUserToWorkspaceRequest) GetUserId() *string {
	return s.UserId
}

func (s *AddUserToWorkspaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddUserToWorkspaceRequest) SetRoleId(v int64) *AddUserToWorkspaceRequest {
	s.RoleId = &v
	return s
}

func (s *AddUserToWorkspaceRequest) SetUserId(v string) *AddUserToWorkspaceRequest {
	s.UserId = &v
	return s
}

func (s *AddUserToWorkspaceRequest) SetWorkspaceId(v string) *AddUserToWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddUserToWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
