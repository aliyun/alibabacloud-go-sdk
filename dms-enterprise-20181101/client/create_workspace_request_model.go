// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateWorkspaceRequest
	GetClientToken() *string
	SetDescription(v string) *CreateWorkspaceRequest
	GetDescription() *string
	SetVpcId(v string) *CreateWorkspaceRequest
	GetVpcId() *string
	SetWorkspaceName(v string) *CreateWorkspaceRequest
	GetWorkspaceName() *string
	SetWorkspaceRegion(v string) *CreateWorkspaceRequest
	GetWorkspaceRegion() *string
}

type CreateWorkspaceRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// token-xxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The VPC ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-xxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The name of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// workspace_xxx
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	WorkspaceRegion *string `json:"WorkspaceRegion,omitempty" xml:"WorkspaceRegion,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateWorkspaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkspaceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateWorkspaceRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *CreateWorkspaceRequest) GetWorkspaceRegion() *string {
	return s.WorkspaceRegion
}

func (s *CreateWorkspaceRequest) SetClientToken(v string) *CreateWorkspaceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDescription(v string) *CreateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceRequest) SetVpcId(v string) *CreateWorkspaceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateWorkspaceRequest) SetWorkspaceName(v string) *CreateWorkspaceRequest {
	s.WorkspaceName = &v
	return s
}

func (s *CreateWorkspaceRequest) SetWorkspaceRegion(v string) *CreateWorkspaceRequest {
	s.WorkspaceRegion = &v
	return s
}

func (s *CreateWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
