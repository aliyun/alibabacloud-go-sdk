// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunResourceGroupId(v string) *CreateProjectShrinkRequest
	GetAliyunResourceGroupId() *string
	SetAliyunResourceTagsShrink(v string) *CreateProjectShrinkRequest
	GetAliyunResourceTagsShrink() *string
	SetDescription(v string) *CreateProjectShrinkRequest
	GetDescription() *string
	SetDevEnvironmentEnabled(v bool) *CreateProjectShrinkRequest
	GetDevEnvironmentEnabled() *bool
	SetDevRoleDisabled(v bool) *CreateProjectShrinkRequest
	GetDevRoleDisabled() *bool
	SetDisplayName(v string) *CreateProjectShrinkRequest
	GetDisplayName() *string
	SetName(v string) *CreateProjectShrinkRequest
	GetName() *string
	SetPaiTaskEnabled(v bool) *CreateProjectShrinkRequest
	GetPaiTaskEnabled() *bool
}

type CreateProjectShrinkRequest struct {
	// The ID of the Alibaba Cloud resource group to which the workspace belongs. You can log on to the [Resource Management console](https://resourcemanager.console.aliyun.com/resource-groups) and go to the Resource Group page to query the ID.
	//
	// You must configure this parameter to specify an Alibaba Cloud resource group for the workspace that you want to create.
	//
	// example:
	//
	// rg-acfmzbn7pti3zff
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// The tags.
	AliyunResourceTagsShrink *string `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty"`
	// The description of the workspace.
	//
	// example:
	//
	// Financial analysis group project data development
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the development environment. Valid values:
	//
	// 	- true : enables the development environment. In this case, the development environment is isolated from the production environment in the workspace.
	//
	// 	- false: disables the development environment. In this case, only the production environment is used in the workspace.
	//
	// example:
	//
	// false
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// Specifies whether to disable the Develop role. Valid values:
	//
	// 	- false (default)
	//
	// 	- true
	//
	// example:
	//
	// true
	DevRoleDisabled *bool `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	// The display name of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sora financial analysis
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The name of the workspace.
	//
	// Limits:
	//
	// 	- The workspace name must be unqiue in a region.
	//
	// 	- The workspace name can contain letters, digits, and underscores (_), and must start with a letter.
	//
	// 	- The workspace name must be 3 to 28 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// sora_finance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to enable scheduling of Platform for AI (PAI) tasks. Valid values:
	//
	// 	- true: enables scheduling of PAI tasks. In this case, you can create a PAI node in a DataWorks workspace and configure scheduling properties for the node to implement periodic scheduling of PAI tasks.
	//
	// 	- false: disables scheduling of PAI tasks.
	//
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
}

func (s CreateProjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectShrinkRequest) GetAliyunResourceGroupId() *string {
	return s.AliyunResourceGroupId
}

func (s *CreateProjectShrinkRequest) GetAliyunResourceTagsShrink() *string {
	return s.AliyunResourceTagsShrink
}

func (s *CreateProjectShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProjectShrinkRequest) GetDevEnvironmentEnabled() *bool {
	return s.DevEnvironmentEnabled
}

func (s *CreateProjectShrinkRequest) GetDevRoleDisabled() *bool {
	return s.DevRoleDisabled
}

func (s *CreateProjectShrinkRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateProjectShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateProjectShrinkRequest) GetPaiTaskEnabled() *bool {
	return s.PaiTaskEnabled
}

func (s *CreateProjectShrinkRequest) SetAliyunResourceGroupId(v string) *CreateProjectShrinkRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetAliyunResourceTagsShrink(v string) *CreateProjectShrinkRequest {
	s.AliyunResourceTagsShrink = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDescription(v string) *CreateProjectShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDevEnvironmentEnabled(v bool) *CreateProjectShrinkRequest {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDevRoleDisabled(v bool) *CreateProjectShrinkRequest {
	s.DevRoleDisabled = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDisplayName(v string) *CreateProjectShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetName(v string) *CreateProjectShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetPaiTaskEnabled(v bool) *CreateProjectShrinkRequest {
	s.PaiTaskEnabled = &v
	return s
}

func (s *CreateProjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
