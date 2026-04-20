// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDmsUserId(v string) *UpdateWorkspaceUserRequest
	GetDmsUserId() *string
	SetRoleIds(v string) *UpdateWorkspaceUserRequest
	GetRoleIds() *string
	SetWorkspaceId(v string) *UpdateWorkspaceUserRequest
	GetWorkspaceId() *string
}

type UpdateWorkspaceUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123***
	DmsUserId *string `json:"DmsUserId,omitempty" xml:"DmsUserId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 11,12
	RoleIds *string `json:"RoleIds,omitempty" xml:"RoleIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateWorkspaceUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceUserRequest) GetDmsUserId() *string {
	return s.DmsUserId
}

func (s *UpdateWorkspaceUserRequest) GetRoleIds() *string {
	return s.RoleIds
}

func (s *UpdateWorkspaceUserRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateWorkspaceUserRequest) SetDmsUserId(v string) *UpdateWorkspaceUserRequest {
	s.DmsUserId = &v
	return s
}

func (s *UpdateWorkspaceUserRequest) SetRoleIds(v string) *UpdateWorkspaceUserRequest {
	s.RoleIds = &v
	return s
}

func (s *UpdateWorkspaceUserRequest) SetWorkspaceId(v string) *UpdateWorkspaceUserRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateWorkspaceUserRequest) Validate() error {
	return dara.Validate(s)
}
