// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateWorkspaceRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateWorkspaceRequest
	GetDisplayName() *string
	SetEnvTypes(v []*string) *CreateWorkspaceRequest
	GetEnvTypes() []*string
	SetResourceGroupId(v string) *CreateWorkspaceRequest
	GetResourceGroupId() *string
	SetWorkspaceName(v string) *CreateWorkspaceRequest
	GetWorkspaceName() *string
}

type CreateWorkspaceRequest struct {
	// The description of the workspace. The description can be up to 80 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the workspace. You can set it based on the purpose of the workspace. If left empty, the name of the workspace is used.
	//
	// example:
	//
	// display name
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The environment of the workspace.
	//
	// 	- Workspaces in basic mode can run only in the production environment (prod).
	//
	// 	- Workspaces in standard mode can run in both the development and production environments (dev and prod).
	//
	// This parameter is required.
	EnvTypes        []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the workspace. Format:
	//
	// 	- The name must be 3 to 23 characters in length, and can contain letters, underscores (_), and digits.
	//
	// 	- The name must start with a letter.
	//
	// 	- It must be unique in the current region.
	//
	// This parameter is required.
	//
	// example:
	//
	// workspace_example
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateWorkspaceRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateWorkspaceRequest) GetEnvTypes() []*string {
	return s.EnvTypes
}

func (s *CreateWorkspaceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateWorkspaceRequest) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *CreateWorkspaceRequest) SetDescription(v string) *CreateWorkspaceRequest {
	s.Description = &v
	return s
}

func (s *CreateWorkspaceRequest) SetDisplayName(v string) *CreateWorkspaceRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateWorkspaceRequest) SetEnvTypes(v []*string) *CreateWorkspaceRequest {
	s.EnvTypes = v
	return s
}

func (s *CreateWorkspaceRequest) SetResourceGroupId(v string) *CreateWorkspaceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateWorkspaceRequest) SetWorkspaceName(v string) *CreateWorkspaceRequest {
	s.WorkspaceName = &v
	return s
}

func (s *CreateWorkspaceRequest) Validate() error {
	return dara.Validate(s)
}
