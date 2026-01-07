// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAgentSpaceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *UpdateDataAgentSpaceInfoRequest
	GetDMSUnit() *string
	SetWorkspaceDesc(v string) *UpdateDataAgentSpaceInfoRequest
	GetWorkspaceDesc() *string
	SetWorkspaceId(v string) *UpdateDataAgentSpaceInfoRequest
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *UpdateDataAgentSpaceInfoRequest
	GetWorkspaceName() *string
}

type UpdateDataAgentSpaceInfoRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// example:
	//
	// space for test new
	WorkspaceDesc *string `json:"WorkspaceDesc,omitempty" xml:"WorkspaceDesc,omitempty"`
	// example:
	//
	// 20923*****7291
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// yunqitest_v2
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s UpdateDataAgentSpaceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAgentSpaceInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataAgentSpaceInfoRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *UpdateDataAgentSpaceInfoRequest) GetWorkspaceDesc() *string {
	return s.WorkspaceDesc
}

func (s *UpdateDataAgentSpaceInfoRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateDataAgentSpaceInfoRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *UpdateDataAgentSpaceInfoRequest) SetDMSUnit(v string) *UpdateDataAgentSpaceInfoRequest {
	s.DMSUnit = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoRequest) SetWorkspaceDesc(v string) *UpdateDataAgentSpaceInfoRequest {
	s.WorkspaceDesc = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoRequest) SetWorkspaceId(v string) *UpdateDataAgentSpaceInfoRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoRequest) SetWorkspaceName(v string) *UpdateDataAgentSpaceInfoRequest {
	s.WorkspaceName = &v
	return s
}

func (s *UpdateDataAgentSpaceInfoRequest) Validate() error {
	return dara.Validate(s)
}
