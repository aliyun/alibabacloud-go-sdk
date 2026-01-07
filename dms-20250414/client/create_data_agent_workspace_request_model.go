// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAgentWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *CreateDataAgentWorkspaceRequest
	GetDMSUnit() *string
	SetWorkspaceDesc(v string) *CreateDataAgentWorkspaceRequest
	GetWorkspaceDesc() *string
	SetWorkspaceName(v string) *CreateDataAgentWorkspaceRequest
	GetWorkspaceName() *string
}

type CreateDataAgentWorkspaceRequest struct {
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// example:
	//
	// Space for test
	WorkspaceDesc *string `json:"WorkspaceDesc,omitempty" xml:"WorkspaceDesc,omitempty"`
	// example:
	//
	// workspaceTest
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s CreateDataAgentWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAgentWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateDataAgentWorkspaceRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *CreateDataAgentWorkspaceRequest) GetWorkspaceDesc() *string {
	return s.WorkspaceDesc
}

func (s *CreateDataAgentWorkspaceRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *CreateDataAgentWorkspaceRequest) SetDMSUnit(v string) *CreateDataAgentWorkspaceRequest {
	s.DMSUnit = &v
	return s
}

func (s *CreateDataAgentWorkspaceRequest) SetWorkspaceDesc(v string) *CreateDataAgentWorkspaceRequest {
	s.WorkspaceDesc = &v
	return s
}

func (s *CreateDataAgentWorkspaceRequest) SetWorkspaceName(v string) *CreateDataAgentWorkspaceRequest {
	s.WorkspaceName = &v
	return s
}

func (s *CreateDataAgentWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
