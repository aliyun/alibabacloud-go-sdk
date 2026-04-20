// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddWorkspaceUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDmsUserIds(v string) *AddWorkspaceUserRequest
	GetDmsUserIds() *string
	SetRoleId(v string) *AddWorkspaceUserRequest
	GetRoleId() *string
	SetRoleSource(v string) *AddWorkspaceUserRequest
	GetRoleSource() *string
	SetWorkspaceId(v string) *AddWorkspaceUserRequest
	GetWorkspaceId() *string
}

type AddWorkspaceUserRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123***
	DmsUserIds *string `json:"DmsUserIds,omitempty" xml:"DmsUserIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 31****
	RoleId *string `json:"RoleId,omitempty" xml:"RoleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// INNER
	RoleSource *string `json:"RoleSource,omitempty" xml:"RoleSource,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddWorkspaceUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AddWorkspaceUserRequest) GoString() string {
	return s.String()
}

func (s *AddWorkspaceUserRequest) GetDmsUserIds() *string {
	return s.DmsUserIds
}

func (s *AddWorkspaceUserRequest) GetRoleId() *string {
	return s.RoleId
}

func (s *AddWorkspaceUserRequest) GetRoleSource() *string {
	return s.RoleSource
}

func (s *AddWorkspaceUserRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddWorkspaceUserRequest) SetDmsUserIds(v string) *AddWorkspaceUserRequest {
	s.DmsUserIds = &v
	return s
}

func (s *AddWorkspaceUserRequest) SetRoleId(v string) *AddWorkspaceUserRequest {
	s.RoleId = &v
	return s
}

func (s *AddWorkspaceUserRequest) SetRoleSource(v string) *AddWorkspaceUserRequest {
	s.RoleSource = &v
	return s
}

func (s *AddWorkspaceUserRequest) SetWorkspaceId(v string) *AddWorkspaceUserRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddWorkspaceUserRequest) Validate() error {
	return dara.Validate(s)
}
