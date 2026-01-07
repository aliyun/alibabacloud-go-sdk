// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAgentWorkspaceMemberRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *UpdateDataAgentWorkspaceMemberRoleRequest
	GetDMSUnit() *string
	SetMemberId(v string) *UpdateDataAgentWorkspaceMemberRoleRequest
	GetMemberId() *string
	SetRoleName(v string) *UpdateDataAgentWorkspaceMemberRoleRequest
	GetRoleName() *string
	SetWorkspaceId(v string) *UpdateDataAgentWorkspaceMemberRoleRequest
	GetWorkspaceId() *string
}

type UpdateDataAgentWorkspaceMemberRoleRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// example:
	//
	// 20282*****7591
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// admin
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// example:
	//
	// 12****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateDataAgentWorkspaceMemberRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentWorkspaceMemberRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentWorkspaceMemberRoleRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *UpdateDataAgentWorkspaceMemberRoleRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *UpdateDataAgentWorkspaceMemberRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *UpdateDataAgentWorkspaceMemberRoleRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateDataAgentWorkspaceMemberRoleRequest) SetDMSUnit(v string) *UpdateDataAgentWorkspaceMemberRoleRequest {
	s.DMSUnit = &v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleRequest) SetMemberId(v string) *UpdateDataAgentWorkspaceMemberRoleRequest {
	s.MemberId = &v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleRequest) SetRoleName(v string) *UpdateDataAgentWorkspaceMemberRoleRequest {
	s.RoleName = &v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleRequest) SetWorkspaceId(v string) *UpdateDataAgentWorkspaceMemberRoleRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDataAgentWorkspaceMemberRoleRequest) Validate() error {
	return dara.Validate(s)
}
