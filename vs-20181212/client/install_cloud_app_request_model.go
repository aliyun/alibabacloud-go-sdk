// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *InstallCloudAppRequest
	GetAppId() *string
	SetPageNumber(v int32) *InstallCloudAppRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *InstallCloudAppRequest
	GetPageSize() *int32
	SetPatchId(v string) *InstallCloudAppRequest
	GetPatchId() *string
	SetProjectId(v string) *InstallCloudAppRequest
	GetProjectId() *string
	SetRenderingInstanceId(v string) *InstallCloudAppRequest
	GetRenderingInstanceId() *string
	SetRenderingInstanceIds(v []*string) *InstallCloudAppRequest
	GetRenderingInstanceIds() []*string
}

type InstallCloudAppRequest struct {
	// The cloud application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The page number for the paged query of instance associations under the project. By default, the paged query is sorted by instance association time in descending order. This parameter takes effect only when ProjectId is not empty. It limits the maximum number of instances on which the operation is executed under the project to control the impact scope. Default value: 1.
	//
	// 1. About the value range of PageNumber:
	//
	//   a. [Recommended] Method 1: Calculate the upper limit based on the total number of instances associated with the project (you can call the ListRenderingProjectInstances operation to obtain the total number).
	//
	//   b. Method 2: Determine whether PageNumber has reached the upper limit of the project based on the response. [Advantage: You do not need to calculate the value range.] If the response meets either of the following conditions, PageNumber has reached the upper limit:
	//
	//     i. The response returns HTTP status code 403 and error code 200301.
	//
	//     ii. The response returns SuccessInstanceCount + FailedInstanceCount < PageSize.
	//
	// 2. Example scenarios:
	//
	//   a. Full installation for instances under a project: If the number of instances in the project is greater than PageSize (default value: 100), you can call the Install operation multiple times by incrementing PageNumber by 1 each time to perform a full installation. You can call the ListCloudAppInstallations operation to query the installation progress of the instances.
	//
	//   b. Installation for newly added instances under a project: Start with PageNumber=1. By default, the paged query is sorted by instance association time in descending order. The page with PageNumber=1 contains the most recently added instances.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of instances that can be selected under the project. This parameter takes effect only when ProjectId is not empty. It limits the maximum number of instances on which the operation is performed under the project to control the impact scope. Default value: 100. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The patch package ID to install. This parameter is supported only in Windows scenarios.
	//
	// 1. By default, StablePatchId is installed.
	//
	// 2. If you set the value to origin, the original version is installed.
	//
	// example:
	//
	// patch-7bdf679812484df08a956b73e0b3bdf6
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// The project ID.
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
	RenderingInstanceIds []*string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty" type:"Repeated"`
}

func (s InstallCloudAppRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudAppRequest) GoString() string {
	return s.String()
}

func (s *InstallCloudAppRequest) GetAppId() *string {
	return s.AppId
}

func (s *InstallCloudAppRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *InstallCloudAppRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *InstallCloudAppRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *InstallCloudAppRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *InstallCloudAppRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *InstallCloudAppRequest) GetRenderingInstanceIds() []*string {
	return s.RenderingInstanceIds
}

func (s *InstallCloudAppRequest) SetAppId(v string) *InstallCloudAppRequest {
	s.AppId = &v
	return s
}

func (s *InstallCloudAppRequest) SetPageNumber(v int32) *InstallCloudAppRequest {
	s.PageNumber = &v
	return s
}

func (s *InstallCloudAppRequest) SetPageSize(v int32) *InstallCloudAppRequest {
	s.PageSize = &v
	return s
}

func (s *InstallCloudAppRequest) SetPatchId(v string) *InstallCloudAppRequest {
	s.PatchId = &v
	return s
}

func (s *InstallCloudAppRequest) SetProjectId(v string) *InstallCloudAppRequest {
	s.ProjectId = &v
	return s
}

func (s *InstallCloudAppRequest) SetRenderingInstanceId(v string) *InstallCloudAppRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *InstallCloudAppRequest) SetRenderingInstanceIds(v []*string) *InstallCloudAppRequest {
	s.RenderingInstanceIds = v
	return s
}

func (s *InstallCloudAppRequest) Validate() error {
	return dara.Validate(s)
}
