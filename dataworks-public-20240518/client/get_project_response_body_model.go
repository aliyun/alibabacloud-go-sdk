// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProject(v *GetProjectResponseBodyProject) *GetProjectResponseBody
	GetProject() *GetProjectResponseBodyProject
	SetRequestId(v string) *GetProjectResponseBody
	GetRequestId() *string
}

type GetProjectResponseBody struct {
	// The information about the workspace.
	Project *GetProjectResponseBodyProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// 22C97E95-F023-56B5-8852-B1A77A17XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) GetProject() *GetProjectResponseBodyProject {
	return s.Project
}

func (s *GetProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectResponseBody) SetProject(v *GetProjectResponseBodyProject) *GetProjectResponseBody {
	s.Project = v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectResponseBody) Validate() error {
	if s.Project != nil {
		if err := s.Project.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProjectResponseBodyProject struct {
	// The ID of the Alibaba Cloud resource group to which the workspace belongs.
	//
	// example:
	//
	// rg-acfmzbn7pti3zfa
	AliyunResourceGroupId *string `json:"AliyunResourceGroupId,omitempty" xml:"AliyunResourceGroupId,omitempty"`
	// The tags.
	AliyunResourceTags []*GetProjectResponseBodyProjectAliyunResourceTags `json:"AliyunResourceTags,omitempty" xml:"AliyunResourceTags,omitempty" type:"Repeated"`
	// The description of the workspace.
	//
	// example:
	//
	// Financial analysis group project data development
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the development environment is enabled. Valid values:
	//
	// 	- true: The development environment is enabled. In this case, the development environment is isolated from the production environment in the workspace.
	//
	// 	- false: The development environment is disabled. In this case, only the production environment is used in the workspace.
	//
	// example:
	//
	// true
	DevEnvironmentEnabled *bool `json:"DevEnvironmentEnabled,omitempty" xml:"DevEnvironmentEnabled,omitempty"`
	// Indicates whether the Develop role is disabled. Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	DevRoleDisabled *bool `json:"DevRoleDisabled,omitempty" xml:"DevRoleDisabled,omitempty"`
	// The display name of the workspace.
	//
	// example:
	//
	// Sora financial analysis
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 28477242
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// sora_finance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the Alibaba Cloud account to which the workspace belongs.
	//
	// example:
	//
	// 207947397706614299
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// Indicates whether scheduling of PAI tasks is enabled. Valid values:
	//
	// 	- true: Scheduling of PAI tasks is enabled. In this case, you can create a PAI node in a DataWorks workspace and configure scheduling properties for the node to implement periodic scheduling of PAI tasks.
	//
	// 	- false: Scheduling of PAI tasks is disabled.
	//
	// example:
	//
	// true
	PaiTaskEnabled *bool `json:"PaiTaskEnabled,omitempty" xml:"PaiTaskEnabled,omitempty"`
	// The status of the workspace. Valid values:
	//
	// 	- Available
	//
	// 	- Initializing
	//
	// 	- InitFailed
	//
	// 	- Forbidden
	//
	// 	- Deleting
	//
	// 	- DeleteFailed
	//
	// 	- Frozen
	//
	// 	- Updating
	//
	// 	- UpdateFailed
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetProjectResponseBodyProject) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyProject) GetAliyunResourceGroupId() *string {
	return s.AliyunResourceGroupId
}

func (s *GetProjectResponseBodyProject) GetAliyunResourceTags() []*GetProjectResponseBodyProjectAliyunResourceTags {
	return s.AliyunResourceTags
}

func (s *GetProjectResponseBodyProject) GetDescription() *string {
	return s.Description
}

func (s *GetProjectResponseBodyProject) GetDevEnvironmentEnabled() *bool {
	return s.DevEnvironmentEnabled
}

func (s *GetProjectResponseBodyProject) GetDevRoleDisabled() *bool {
	return s.DevRoleDisabled
}

func (s *GetProjectResponseBodyProject) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetProjectResponseBodyProject) GetId() *int64 {
	return s.Id
}

func (s *GetProjectResponseBodyProject) GetName() *string {
	return s.Name
}

func (s *GetProjectResponseBodyProject) GetOwner() *string {
	return s.Owner
}

func (s *GetProjectResponseBodyProject) GetPaiTaskEnabled() *bool {
	return s.PaiTaskEnabled
}

func (s *GetProjectResponseBodyProject) GetStatus() *string {
	return s.Status
}

func (s *GetProjectResponseBodyProject) SetAliyunResourceGroupId(v string) *GetProjectResponseBodyProject {
	s.AliyunResourceGroupId = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetAliyunResourceTags(v []*GetProjectResponseBodyProjectAliyunResourceTags) *GetProjectResponseBodyProject {
	s.AliyunResourceTags = v
	return s
}

func (s *GetProjectResponseBodyProject) SetDescription(v string) *GetProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetDevEnvironmentEnabled(v bool) *GetProjectResponseBodyProject {
	s.DevEnvironmentEnabled = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetDevRoleDisabled(v bool) *GetProjectResponseBodyProject {
	s.DevRoleDisabled = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetDisplayName(v string) *GetProjectResponseBodyProject {
	s.DisplayName = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetId(v int64) *GetProjectResponseBodyProject {
	s.Id = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetName(v string) *GetProjectResponseBodyProject {
	s.Name = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetOwner(v string) *GetProjectResponseBodyProject {
	s.Owner = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetPaiTaskEnabled(v bool) *GetProjectResponseBodyProject {
	s.PaiTaskEnabled = &v
	return s
}

func (s *GetProjectResponseBodyProject) SetStatus(v string) *GetProjectResponseBodyProject {
	s.Status = &v
	return s
}

func (s *GetProjectResponseBodyProject) Validate() error {
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

type GetProjectResponseBodyProjectAliyunResourceTags struct {
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

func (s GetProjectResponseBodyProjectAliyunResourceTags) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBodyProjectAliyunResourceTags) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBodyProjectAliyunResourceTags) GetKey() *string {
	return s.Key
}

func (s *GetProjectResponseBodyProjectAliyunResourceTags) GetValue() *string {
	return s.Value
}

func (s *GetProjectResponseBodyProjectAliyunResourceTags) SetKey(v string) *GetProjectResponseBodyProjectAliyunResourceTags {
	s.Key = &v
	return s
}

func (s *GetProjectResponseBodyProjectAliyunResourceTags) SetValue(v string) *GetProjectResponseBodyProjectAliyunResourceTags {
	s.Value = &v
	return s
}

func (s *GetProjectResponseBodyProjectAliyunResourceTags) Validate() error {
	return dara.Validate(s)
}
