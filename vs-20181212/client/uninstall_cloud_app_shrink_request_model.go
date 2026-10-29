// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallCloudAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UninstallCloudAppShrinkRequest
	GetAppId() *string
	SetPageNumber(v int32) *UninstallCloudAppShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *UninstallCloudAppShrinkRequest
	GetPageSize() *int32
	SetPatchId(v string) *UninstallCloudAppShrinkRequest
	GetPatchId() *string
	SetProjectId(v string) *UninstallCloudAppShrinkRequest
	GetProjectId() *string
	SetRenderingInstanceId(v string) *UninstallCloudAppShrinkRequest
	GetRenderingInstanceId() *string
	SetRenderingInstanceIdsShrink(v string) *UninstallCloudAppShrinkRequest
	GetRenderingInstanceIdsShrink() *string
}

type UninstallCloudAppShrinkRequest struct {
	// The ID of the cloud application.
	//
	// This parameter is required.
	//
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The page number for instance associations in a project. By default, a paged query sorts results in descending order of instance association time. This parameter is valid only when ProjectId is not empty. It limits the maximum number of instances on which the action is executed in the project to control the impact scope of the action. Default value: 1.
	//
	// 1. Valid values of PageNumber:
	//
	//   a. [Recommended] Method 1: Calculate the upper limit based on the total number of instances associated with the project. You can call the ListRenderingProjectInstances operation to obtain the total number.
	//
	//   b. Method 2: Determine whether PageNumber reaches the upper limit for the project based on the response. [Advantage: No need to calculate the valid value range]. PageNumber reaches the upper limit if the response meets any of the following conditions:
	//
	//     i. The response returns HTTP status code 403 and error code 200301.
	//
	//     ii. The sum of SuccessInstanceCount and FailedInstanceCount in the response is less than PageSize.
	//
	// 2. Scenario example:
	//
	//   a. Uninstall instances in a project in full: If the number of instances in the project is greater than PageSize (default value: 100), you can call the Uninstall operation multiple times and increment PageNumber by 1 each time to uninstall all instances. You can call the ListCloudAppInstallations operation to query the uninstallation progress of the instances in the project.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of selected instances in a project. This parameter is valid only when ProjectId is not empty. It limits the maximum number of instances on which the action is executed in the project to control the impact scope of the action. Default value: 100. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the patch package to uninstall. This parameter is supported only in Windows scenarios.
	//
	// 1. By default, the StablePatchId is uninstalled.
	//
	// 2. If you set this parameter to origin, the original version is uninstalled.
	//
	// 3. If you set this parameter to all, all installed versions are uninstalled.
	//
	// example:
	//
	// patch-03fa76e8e13a49b6a966b063d9d309b4
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// The ID of the project.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The cloud application service instance ID.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// The list of cloud application service instance IDs.
	RenderingInstanceIdsShrink *string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty"`
}

func (s UninstallCloudAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallCloudAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *UninstallCloudAppShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UninstallCloudAppShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *UninstallCloudAppShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *UninstallCloudAppShrinkRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *UninstallCloudAppShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UninstallCloudAppShrinkRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *UninstallCloudAppShrinkRequest) GetRenderingInstanceIdsShrink() *string {
	return s.RenderingInstanceIdsShrink
}

func (s *UninstallCloudAppShrinkRequest) SetAppId(v string) *UninstallCloudAppShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UninstallCloudAppShrinkRequest) SetPageNumber(v int32) *UninstallCloudAppShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *UninstallCloudAppShrinkRequest) SetPageSize(v int32) *UninstallCloudAppShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *UninstallCloudAppShrinkRequest) SetPatchId(v string) *UninstallCloudAppShrinkRequest {
	s.PatchId = &v
	return s
}

func (s *UninstallCloudAppShrinkRequest) SetProjectId(v string) *UninstallCloudAppShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *UninstallCloudAppShrinkRequest) SetRenderingInstanceId(v string) *UninstallCloudAppShrinkRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *UninstallCloudAppShrinkRequest) SetRenderingInstanceIdsShrink(v string) *UninstallCloudAppShrinkRequest {
	s.RenderingInstanceIdsShrink = &v
	return s
}

func (s *UninstallCloudAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
