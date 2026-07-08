// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *InstallCloudAppShrinkRequest
	GetAppId() *string
	SetPageNumber(v int32) *InstallCloudAppShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *InstallCloudAppShrinkRequest
	GetPageSize() *int32
	SetPatchId(v string) *InstallCloudAppShrinkRequest
	GetPatchId() *string
	SetProjectId(v string) *InstallCloudAppShrinkRequest
	GetProjectId() *string
	SetRenderingInstanceId(v string) *InstallCloudAppShrinkRequest
	GetRenderingInstanceId() *string
	SetRenderingInstanceIdsShrink(v string) *InstallCloudAppShrinkRequest
	GetRenderingInstanceIdsShrink() *string
}

type InstallCloudAppShrinkRequest struct {
	// Cloud application ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Page number for paged queries of instance associations under the project. Paged queries default to reverse order by instance association time. This applies only when ProjectId is not empty. It limits the maximum number of instances for actions within the project, controlling the impact scope. Default is 1.
	//
	// 1. PageNumber value range:
	//
	//    a. Method one (recommended): Calculate the upper limit using the total number of instances associated with the project. The ListRenderingProjectInstances interface provides this count.
	//
	//    b. Method two: Determine if PageNumber reaches the project\\"s upper limit by checking the interface return value. This avoids calculating the range. PageNumber reaches the upper limit if the interface returns any of these conditions:
	//
	//    ⅰ. A 403 status code and error code 200301.
	//
	//    ⅱ. The sum of \\`SuccessInstanceCount\\` and \\`FailedInstanceCount\\` is less than \\`PageSize\\`.
	//
	// 2. Scenario examples:
	//
	//    a. Full installation for project instances: If the number of project instances exceeds \\`PageSize\\` (default 100), invoke Install multiple times. Increment PageNumber by 1 for each call to complete the full installation. Get project instance installation progress using the ListCloudAppInstallations interface.
	//
	//    b. New instance installation for a project: Start with \\`PageNumber=1\\`. Paged queries default to reverse order by instance association time. The \\`PageNumber=1\\` page shows the latest new instances.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Maximum number of instances selected for the project. This applies only when ProjectId is not empty. It limits the maximum number of instances for actions within the project, controlling the impact scope. Default is 100. The value range is 1-100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Patch package ID to install. This is only for Windows scenarios.
	//
	// 1. Install \\`StablePatchId\\` by default.
	//
	// 2. Enter \\`origin\\` to install the original version.
	//
	// example:
	//
	// patch-7bdf679812484df08a956b73e0b3bdf6
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// Project ID
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Cloud application service instance ID
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// List of cloud application service instance IDs
	RenderingInstanceIdsShrink *string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty"`
}

func (s InstallCloudAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *InstallCloudAppShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *InstallCloudAppShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *InstallCloudAppShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *InstallCloudAppShrinkRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *InstallCloudAppShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *InstallCloudAppShrinkRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *InstallCloudAppShrinkRequest) GetRenderingInstanceIdsShrink() *string {
	return s.RenderingInstanceIdsShrink
}

func (s *InstallCloudAppShrinkRequest) SetAppId(v string) *InstallCloudAppShrinkRequest {
	s.AppId = &v
	return s
}

func (s *InstallCloudAppShrinkRequest) SetPageNumber(v int32) *InstallCloudAppShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *InstallCloudAppShrinkRequest) SetPageSize(v int32) *InstallCloudAppShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *InstallCloudAppShrinkRequest) SetPatchId(v string) *InstallCloudAppShrinkRequest {
	s.PatchId = &v
	return s
}

func (s *InstallCloudAppShrinkRequest) SetProjectId(v string) *InstallCloudAppShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *InstallCloudAppShrinkRequest) SetRenderingInstanceId(v string) *InstallCloudAppShrinkRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *InstallCloudAppShrinkRequest) SetRenderingInstanceIdsShrink(v string) *InstallCloudAppShrinkRequest {
	s.RenderingInstanceIdsShrink = &v
	return s
}

func (s *InstallCloudAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
