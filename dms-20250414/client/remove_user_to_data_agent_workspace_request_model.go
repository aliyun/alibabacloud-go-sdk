// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserToDataAgentWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *RemoveUserToDataAgentWorkspaceRequest
	GetDMSUnit() *string
	SetMemberId(v string) *RemoveUserToDataAgentWorkspaceRequest
	GetMemberId() *string
	SetWorkspaceId(v string) *RemoveUserToDataAgentWorkspaceRequest
	GetWorkspaceId() *string
}

type RemoveUserToDataAgentWorkspaceRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// example:
	//
	// 21482*****7584
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// example:
	//
	// 20923*****7291
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RemoveUserToDataAgentWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserToDataAgentWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *RemoveUserToDataAgentWorkspaceRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *RemoveUserToDataAgentWorkspaceRequest) GetMemberId() *string {
	return s.MemberId
}

func (s *RemoveUserToDataAgentWorkspaceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RemoveUserToDataAgentWorkspaceRequest) SetDMSUnit(v string) *RemoveUserToDataAgentWorkspaceRequest {
	s.DMSUnit = &v
	return s
}

func (s *RemoveUserToDataAgentWorkspaceRequest) SetMemberId(v string) *RemoveUserToDataAgentWorkspaceRequest {
	s.MemberId = &v
	return s
}

func (s *RemoveUserToDataAgentWorkspaceRequest) SetWorkspaceId(v string) *RemoveUserToDataAgentWorkspaceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RemoveUserToDataAgentWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
