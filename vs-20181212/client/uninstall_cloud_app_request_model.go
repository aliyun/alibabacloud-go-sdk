// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallCloudAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UninstallCloudAppRequest
	GetAppId() *string
	SetPageNumber(v int32) *UninstallCloudAppRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *UninstallCloudAppRequest
	GetPageSize() *int32
	SetPatchId(v string) *UninstallCloudAppRequest
	GetPatchId() *string
	SetProjectId(v string) *UninstallCloudAppRequest
	GetProjectId() *string
	SetRenderingInstanceId(v string) *UninstallCloudAppRequest
	GetRenderingInstanceId() *string
	SetRenderingInstanceIds(v []*string) *UninstallCloudAppRequest
	GetRenderingInstanceIds() []*string
}

type UninstallCloudAppRequest struct {
	// Cloud application ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Page number for paged queries of instance associations in the project. Results are sorted by association time in descending order. This parameter applies only when ProjectId is not empty. It limits the maximum number of instances affected by this operation to control impact scope. Default value: 1.
	//
	// 1. Valid PageNumber range:
	//
	//    a. Recommended method: Calculate the upper limit based on the total number of instances associated with the project. You can get this count using the ListRenderingProjectInstances API.
	//
	//    b. Alternative method: Check the API response to determine whether PageNumber has reached the upper limit. This avoids manual calculation. PageNumber has reached the upper limit if either of the following occurs:
	//
	//    ⅰ. The API returns HTTP status 403 and error code 200301.
	//
	//    ⅱ. The sum of SuccessInstanceCount and FailedInstanceCount in the response is less than PageSize.
	//
	// 2. Example scenario:
	//
	//    a. Full uninstall across all project instances: If the project has more instances than PageSize (default 100), call UninstallCloudApp multiple times, incrementing PageNumber by 1 each time. Track uninstall progress using the ListCloudAppInstallations API.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Maximum number of instances selected in the project. This parameter applies only when ProjectId is not empty. It limits the maximum number of instances affected by this operation to control impact scope. Default value: 100. Valid values: 1 to 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// ID of the patch package to uninstall. Supported only on Windows.
	//
	// 1. Default: uninstall the StablePatchId.
	//
	// 2. Set to origin to uninstall the original version.
	//
	// 3. Set to all to uninstall all installed versions.
	//
	// example:
	//
	// patch-03fa76e8e13a49b6a966b063d9d309b4
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// Project ID
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Cloud application instance ID
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// List of cloud application instance IDs
	RenderingInstanceIds []*string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty" type:"Repeated"`
}

func (s UninstallCloudAppRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallCloudAppRequest) GoString() string {
	return s.String()
}

func (s *UninstallCloudAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *UninstallCloudAppRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *UninstallCloudAppRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *UninstallCloudAppRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *UninstallCloudAppRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *UninstallCloudAppRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *UninstallCloudAppRequest) GetRenderingInstanceIds() []*string {
	return s.RenderingInstanceIds
}

func (s *UninstallCloudAppRequest) SetAppId(v string) *UninstallCloudAppRequest {
	s.AppId = &v
	return s
}

func (s *UninstallCloudAppRequest) SetPageNumber(v int32) *UninstallCloudAppRequest {
	s.PageNumber = &v
	return s
}

func (s *UninstallCloudAppRequest) SetPageSize(v int32) *UninstallCloudAppRequest {
	s.PageSize = &v
	return s
}

func (s *UninstallCloudAppRequest) SetPatchId(v string) *UninstallCloudAppRequest {
	s.PatchId = &v
	return s
}

func (s *UninstallCloudAppRequest) SetProjectId(v string) *UninstallCloudAppRequest {
	s.ProjectId = &v
	return s
}

func (s *UninstallCloudAppRequest) SetRenderingInstanceId(v string) *UninstallCloudAppRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *UninstallCloudAppRequest) SetRenderingInstanceIds(v []*string) *UninstallCloudAppRequest {
	s.RenderingInstanceIds = v
	return s
}

func (s *UninstallCloudAppRequest) Validate() error {
	return dara.Validate(s)
}
