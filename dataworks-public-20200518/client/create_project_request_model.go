// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateProjectRequest
	GetClientToken() *string
	SetDisableDevelopment(v bool) *CreateProjectRequest
	GetDisableDevelopment() *bool
	SetIsAllowDownload(v int32) *CreateProjectRequest
	GetIsAllowDownload() *int32
	SetProjectDescription(v string) *CreateProjectRequest
	GetProjectDescription() *string
	SetProjectIdentifier(v string) *CreateProjectRequest
	GetProjectIdentifier() *string
	SetProjectMode(v int32) *CreateProjectRequest
	GetProjectMode() *int32
	SetProjectName(v string) *CreateProjectRequest
	GetProjectName() *string
	SetResourceManagerResourceGroupId(v string) *CreateProjectRequest
	GetResourceManagerResourceGroupId() *string
	SetTags(v []*CreateProjectRequestTags) *CreateProjectRequest
	GetTags() []*CreateProjectRequestTags
}

type CreateProjectRequest struct {
	// The idempotency parameter. This parameter can be left empty.
	//
	// example:
	//
	// ABFUOEUOTRTRJKE
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to disable the development role. Valid values:
	//
	// - **false*	- (default): enables the development role.
	//
	// - **true**: disables the development role.
	//
	// example:
	//
	// false
	DisableDevelopment *bool `json:"DisableDevelopment,omitempty" xml:"DisableDevelopment,omitempty"`
	// Specifies whether to allow downloading query results from the IDE. Valid values:
	//
	// - **1*	- (default): allows downloading.
	//
	// - **0**: does not allow downloading.
	//
	// example:
	//
	// 1
	IsAllowDownload *int32 `json:"IsAllowDownload,omitempty" xml:"IsAllowDownload,omitempty"`
	// The detailed description of the workspace.
	//
	// example:
	//
	// test_describe
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	// The name of the workspace. The name can contain only letters, digits, and underscores (_), and must start with a letter or digit.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// The mode of the workspace. For more information about the differences between modes, see [Must-read: Differences between simple mode and standard mode](https://help.aliyun.com/document_detail/85772.html).
	//
	// Valid values:
	//
	// - **2*	- (default): simple workspace mode.
	//
	// - **3**: standard workspace mode.
	//
	// example:
	//
	// 3
	ProjectMode *int32 `json:"ProjectMode,omitempty" xml:"ProjectMode,omitempty"`
	// The display name of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzbn7pti3****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The list of tags bound to the workspace.
	Tags []*CreateProjectRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s CreateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateProjectRequest) GetDisableDevelopment() *bool {
	return s.DisableDevelopment
}

func (s *CreateProjectRequest) GetIsAllowDownload() *int32 {
	return s.IsAllowDownload
}

func (s *CreateProjectRequest) GetProjectDescription() *string {
	return s.ProjectDescription
}

func (s *CreateProjectRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *CreateProjectRequest) GetProjectMode() *int32 {
	return s.ProjectMode
}

func (s *CreateProjectRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateProjectRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateProjectRequest) GetTags() []*CreateProjectRequestTags {
	return s.Tags
}

func (s *CreateProjectRequest) SetClientToken(v string) *CreateProjectRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProjectRequest) SetDisableDevelopment(v bool) *CreateProjectRequest {
	s.DisableDevelopment = &v
	return s
}

func (s *CreateProjectRequest) SetIsAllowDownload(v int32) *CreateProjectRequest {
	s.IsAllowDownload = &v
	return s
}

func (s *CreateProjectRequest) SetProjectDescription(v string) *CreateProjectRequest {
	s.ProjectDescription = &v
	return s
}

func (s *CreateProjectRequest) SetProjectIdentifier(v string) *CreateProjectRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *CreateProjectRequest) SetProjectMode(v int32) *CreateProjectRequest {
	s.ProjectMode = &v
	return s
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectRequest) SetResourceManagerResourceGroupId(v string) *CreateProjectRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateProjectRequest) SetTags(v []*CreateProjectRequestTags) *CreateProjectRequest {
	s.Tags = v
	return s
}

func (s *CreateProjectRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateProjectRequestTags struct {
	// The tag key.
	//
	// This parameter is required.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// This parameter is required.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateProjectRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequestTags) GoString() string {
	return s.String()
}

func (s *CreateProjectRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateProjectRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateProjectRequestTags) SetKey(v string) *CreateProjectRequestTags {
	s.Key = &v
	return s
}

func (s *CreateProjectRequestTags) SetValue(v string) *CreateProjectRequestTags {
	s.Value = &v
	return s
}

func (s *CreateProjectRequestTags) Validate() error {
	return dara.Validate(s)
}
