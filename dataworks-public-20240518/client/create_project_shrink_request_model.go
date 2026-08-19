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
	// The ID of the Alibaba Cloud resource group to which the workspace belongs. You can log on to the [Resource Management console](https://resourcemanager.console.aliyun.com/resource-groups) and go to the resource group list page to obtain the ID.
	//
	// This parameter is used to manage the DataWorks workspace that you create within the specified Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfmzbn7****
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// The list of tags.
	AliyunResourceTagsShrink *string `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty"`
	// The description of the workspace.
	//
	// example:
	//
	// Financial analysis group project data development
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the development environment. Valid values:
	//
	// - true: The development environment is enabled for the workspace, which supports isolation between the development and production environments.
	//
	// - false: Only the production environment is used.
	//
	// example:
	//
	// false
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// Specifies whether to disable the development role. Valid values:
	//
	// - false: The development role is enabled. This is the default value.
	//
	// - true: The development role is disabled.
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
	// The workspace name.
	//
	// Constraints:
	//
	// - The workspace name must be unique within the region.
	//
	// - The name must start with a letter and can contain only letters, digits, and underscores (_).
	//
	// - The name must be 3 to 28 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// sora_finance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to enable PAI task scheduling. Valid values:
	//
	// - true: You can create Machine Learning Platform for AI (PAI) nodes in the DataWorks workspace and run them on a periodic schedule based on the node configurations.
	//
	// - false: PAI task scheduling is disabled.
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
