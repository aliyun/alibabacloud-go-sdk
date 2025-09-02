// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateProjectShrinkRequest
	GetClientToken() *string
	SetDisableDevelopment(v bool) *CreateProjectShrinkRequest
	GetDisableDevelopment() *bool
	SetIsAllowDownload(v int32) *CreateProjectShrinkRequest
	GetIsAllowDownload() *int32
	SetProjectDescription(v string) *CreateProjectShrinkRequest
	GetProjectDescription() *string
	SetProjectIdentifier(v string) *CreateProjectShrinkRequest
	GetProjectIdentifier() *string
	SetProjectMode(v int32) *CreateProjectShrinkRequest
	GetProjectMode() *int32
	SetProjectName(v string) *CreateProjectShrinkRequest
	GetProjectName() *string
	SetResourceManagerResourceGroupId(v string) *CreateProjectShrinkRequest
	GetResourceManagerResourceGroupId() *string
	SetTagsShrink(v string) *CreateProjectShrinkRequest
	GetTagsShrink() *string
}

type CreateProjectShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. This parameter can be left empty.
	//
	// example:
	//
	// ABFUOEUOTRTRJKE
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to disable the Develop role. Valid values:
	//
	// 	- **false*	- (default)
	//
	// 	- **true**
	//
	// example:
	//
	// false
	DisableDevelopment *bool `json:"DisableDevelopment,omitempty" xml:"DisableDevelopment,omitempty"`
	// Specifies whether to allow you to download the query result from DataStudio. Valid values:
	//
	// 	- **1*	- (default): allows you to download the query result from DataStudio.
	//
	// 	- **0**: does not allow you to download the query result from DataStudio.
	//
	// example:
	//
	// 1
	IsAllowDownload *int32 `json:"IsAllowDownload,omitempty" xml:"IsAllowDownload,omitempty"`
	// The description of the workspace.
	//
	// example:
	//
	// test_describe
	ProjectDescription *string `json:"ProjectDescription,omitempty" xml:"ProjectDescription,omitempty"`
	// The name of the workspace. The name can contain only letters, digits, and underscores (_) and must start with a letter or digit.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// The mode of the workspace. For more information about the differences between the modes of workspaces, see [Differences between workspaces in basic mode and workspaces in standard mode](https://help.aliyun.com/document_detail/85772.html).
	//
	// Valid values:
	//
	// 	- **2*	- (default): basic mode
	//
	// 	- **3**: standard mode
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
	// The tags added to the workspace.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s CreateProjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateProjectShrinkRequest) GetDisableDevelopment() *bool {
	return s.DisableDevelopment
}

func (s *CreateProjectShrinkRequest) GetIsAllowDownload() *int32 {
	return s.IsAllowDownload
}

func (s *CreateProjectShrinkRequest) GetProjectDescription() *string {
	return s.ProjectDescription
}

func (s *CreateProjectShrinkRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *CreateProjectShrinkRequest) GetProjectMode() *int32 {
	return s.ProjectMode
}

func (s *CreateProjectShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateProjectShrinkRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateProjectShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateProjectShrinkRequest) SetClientToken(v string) *CreateProjectShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetDisableDevelopment(v bool) *CreateProjectShrinkRequest {
	s.DisableDevelopment = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetIsAllowDownload(v int32) *CreateProjectShrinkRequest {
	s.IsAllowDownload = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetProjectDescription(v string) *CreateProjectShrinkRequest {
	s.ProjectDescription = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetProjectIdentifier(v string) *CreateProjectShrinkRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetProjectMode(v int32) *CreateProjectShrinkRequest {
	s.ProjectMode = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetProjectName(v string) *CreateProjectShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetResourceManagerResourceGroupId(v string) *CreateProjectShrinkRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetTagsShrink(v string) *CreateProjectShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateProjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
