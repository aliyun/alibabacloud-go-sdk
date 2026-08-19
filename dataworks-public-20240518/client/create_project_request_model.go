// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunResourceGroupId(v string) *CreateProjectRequest
	GetAliyunResourceGroupId() *string
	SetAliyunResourceTags(v []*CreateProjectRequestAliyunResourceTags) *CreateProjectRequest
	GetAliyunResourceTags() []*CreateProjectRequestAliyunResourceTags
	SetDescription(v string) *CreateProjectRequest
	GetDescription() *string
	SetDevEnvironmentEnabled(v bool) *CreateProjectRequest
	GetDevEnvironmentEnabled() *bool
	SetDevRoleDisabled(v bool) *CreateProjectRequest
	GetDevRoleDisabled() *bool
	SetDisplayName(v string) *CreateProjectRequest
	GetDisplayName() *string
	SetName(v string) *CreateProjectRequest
	GetName() *string
	SetPaiTaskEnabled(v bool) *CreateProjectRequest
	GetPaiTaskEnabled() *bool
}

type CreateProjectRequest struct {
	// The ID of the Alibaba Cloud resource group to which the workspace belongs. You can log on to the [Resource Management console](https://resourcemanager.console.aliyun.com/resource-groups) and go to the resource group list page to obtain the ID.
	//
	// This parameter is used to manage the DataWorks workspace that you create within the specified Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfmzbn7****
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// The list of tags.
	AliyunResourceTags []*CreateProjectRequestAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
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

func (s CreateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) GetAliyunResourceGroupId() *string {
	return s.AliyunResourceGroupId
}

func (s *CreateProjectRequest) GetAliyunResourceTags() []*CreateProjectRequestAliyunResourceTags {
	return s.AliyunResourceTags
}

func (s *CreateProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProjectRequest) GetDevEnvironmentEnabled() *bool {
	return s.DevEnvironmentEnabled
}

func (s *CreateProjectRequest) GetDevRoleDisabled() *bool {
	return s.DevRoleDisabled
}

func (s *CreateProjectRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateProjectRequest) GetName() *string {
	return s.Name
}

func (s *CreateProjectRequest) GetPaiTaskEnabled() *bool {
	return s.PaiTaskEnabled
}

func (s *CreateProjectRequest) SetAliyunResourceGroupId(v string) *CreateProjectRequest {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *CreateProjectRequest) SetAliyunResourceTags(v []*CreateProjectRequestAliyunResourceTags) *CreateProjectRequest {
	s.AliyunResourceTags = v
	return s
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetDevEnvironmentEnabled(v bool) *CreateProjectRequest {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *CreateProjectRequest) SetDevRoleDisabled(v bool) *CreateProjectRequest {
	s.DevRoleDisabled = &v
	return s
}

func (s *CreateProjectRequest) SetDisplayName(v string) *CreateProjectRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRequest) SetPaiTaskEnabled(v bool) *CreateProjectRequest {
	s.PaiTaskEnabled = &v
	return s
}

func (s *CreateProjectRequest) Validate() error {
	if s.AliyunResourceTags != nil {
		for _, item := range s.AliyunResourceTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateProjectRequestAliyunResourceTags struct {
	// The tag key.
	//
	// example:
	//
	// batch
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// blue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateProjectRequestAliyunResourceTags) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequestAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *CreateProjectRequestAliyunResourceTags) GetKey() *string {
	return s.Key
}

func (s *CreateProjectRequestAliyunResourceTags) GetValue() *string {
	return s.Value
}

func (s *CreateProjectRequestAliyunResourceTags) SetKey(v string) *CreateProjectRequestAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *CreateProjectRequestAliyunResourceTags) SetValue(v string) *CreateProjectRequestAliyunResourceTags {
	s.Value = &v
	return s
}

func (s *CreateProjectRequestAliyunResourceTags) Validate() error {
	return dara.Validate(s)
}
