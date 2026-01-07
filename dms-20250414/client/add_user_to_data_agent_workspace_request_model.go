// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserToDataAgentWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *AddUserToDataAgentWorkspaceRequest
	GetDMSUnit() *string
	SetMemberId(v string) *AddUserToDataAgentWorkspaceRequest
	GetMemberId() *string
	SetRoleName(v string) *AddUserToDataAgentWorkspaceRequest
	GetRoleName() *string
	SetWorkspaceId(v string) *AddUserToDataAgentWorkspaceRequest
	GetWorkspaceId() *string
}

type AddUserToDataAgentWorkspaceRequest struct {
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
	// 20923*****7291
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AddUserToDataAgentWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserToDataAgentWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *AddUserToDataAgentWorkspaceRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *AddUserToDataAgentWorkspaceRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *AddUserToDataAgentWorkspaceRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *AddUserToDataAgentWorkspaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AddUserToDataAgentWorkspaceRequest) SetDMSUnit(v string) *AddUserToDataAgentWorkspaceRequest {
	s.DMSUnit = &v
	return s
}

func (s *AddUserToDataAgentWorkspaceRequest) SetMemberId(v string) *AddUserToDataAgentWorkspaceRequest {
	s.MemberId = &v
	return s
}

func (s *AddUserToDataAgentWorkspaceRequest) SetRoleName(v string) *AddUserToDataAgentWorkspaceRequest {
	s.RoleName = &v
	return s
}

func (s *AddUserToDataAgentWorkspaceRequest) SetWorkspaceId(v string) *AddUserToDataAgentWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AddUserToDataAgentWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
