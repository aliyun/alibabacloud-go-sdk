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
	// The description of the workspace. The description cannot exceed 80 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 这是一个工作空间描述示例。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// We recommend that you name the workspace based on its business attribute to facilitate identification of its purpose. If you do not configure this parameter, the workspace name is used by default.
	//
	// example:
	//
	// demo工作空间
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The environments included in the workspace:
	//
	// - The simple mode contains only the production environment (prod).
	//
	// - The standard mode contains both the development environment (dev) and the production environment (prod).
	//
	// This parameter is required.
	EnvTypes []*string `json:"EnvTypes,omitempty" xml:"EnvTypes,omitempty" type:"Repeated"`
	// The resource group ID. For information about how to obtain the resource group ID, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-acfmwp7rkyq****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the workspace. The format is as follows:
	//
	// - The length is 3 to 23 characters and can contain letters, underscores (_), or digits.
	//
	// - It must start with a letter (uppercase or lowercase).
	//
	// - It must be unique within the current region.
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
