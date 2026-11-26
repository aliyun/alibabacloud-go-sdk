// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAppInstallationsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListCloudAppInstallationsShrinkRequest
	GetAppId() *string
	SetAppName(v string) *ListCloudAppInstallationsShrinkRequest
	GetAppName() *string
	SetAppVersion(v string) *ListCloudAppInstallationsShrinkRequest
	GetAppVersion() *string
	SetEndTime(v string) *ListCloudAppInstallationsShrinkRequest
	GetEndTime() *string
	SetPageNumber(v int64) *ListCloudAppInstallationsShrinkRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCloudAppInstallationsShrinkRequest
	GetPageSize() *int64
	SetPatchId(v string) *ListCloudAppInstallationsShrinkRequest
	GetPatchId() *string
	SetProjectId(v string) *ListCloudAppInstallationsShrinkRequest
	GetProjectId() *string
	SetRenderingInstanceId(v string) *ListCloudAppInstallationsShrinkRequest
	GetRenderingInstanceId() *string
	SetRenderingInstanceIdsShrink(v string) *ListCloudAppInstallationsShrinkRequest
	GetRenderingInstanceIdsShrink() *string
	SetStartTime(v string) *ListCloudAppInstallationsShrinkRequest
	GetStartTime() *string
}

type ListCloudAppInstallationsShrinkRequest struct {
	// The ID of the cloud application, which corresponds to a unique application package.
	//
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// com.aaa.bbb
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The version of the application.
	//
	// example:
	//
	// 1.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The start time of the time range filter. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2025-06-02T02:18:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number of the list to query. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The Settings for the number of entries per page during paging in a paged query. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the patch package. This parameter is only supported in Windows scenarios. Valid values:
	//
	// - origin: Only the original version is filtered.
	//
	// example:
	//
	// patch-7bdf679812484df08a956b73e0b3bdf6
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// The ID of the project.
	//
	// example:
	//
	// project-b93ea81de76f48609eed3cd420f0399f
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The instance ID of the cloud application service. This parameter is used to query the cloud applications installed on the specified instance.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// The list of cloud application service instance IDs.
	//
	// example:
	//
	// [\"render-eb01436f39ac459a96ac566c418ca926\"]
	RenderingInstanceIdsShrink *string `json:"RenderingInstanceIds,omitempty" xml:"RenderingInstanceIds,omitempty"`
	// The start time of the time range filter. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2025-05-26T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListCloudAppInstallationsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAppInstallationsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListCloudAppInstallationsShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListCloudAppInstallationsShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListCloudAppInstallationsShrinkRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListCloudAppInstallationsShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListCloudAppInstallationsShrinkRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCloudAppInstallationsShrinkRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCloudAppInstallationsShrinkRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *ListCloudAppInstallationsShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListCloudAppInstallationsShrinkRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ListCloudAppInstallationsShrinkRequest) GetRenderingInstanceIdsShrink() *string {
	return s.RenderingInstanceIdsShrink
}

func (s *ListCloudAppInstallationsShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListCloudAppInstallationsShrinkRequest) SetAppId(v string) *ListCloudAppInstallationsShrinkRequest {
	s.AppId = &v
	return s
}

func (s *ListCloudAppInstallationsShrinkRequest) SetAppName(v string) *ListCloudAppInstallationsShrinkRequest {
	s.AppName = &v
	return s
}

func (s *ListCloudAppInstallationsShrinkRequest) SetAppVersion(v string) *ListCloudAppInstallationsShrinkRequest {
	s.AppVersion = &v
	return s
}

func (s *ListCloudAppInstallationsShrinkRequest) SetEndTime(v string) *ListCloudAppInstallationsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListCloudAppInstallationsShrinkRequest) SetPageNumber(v int64) *ListCloudAppInstallationsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloudAppInstallationsShrinkRequest) SetPageSize(v int64) *ListCloudAppInstallationsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudAppInstallationsShrinkRequest) SetPatchId(v string) *ListCloudAppInstallationsShrinkRequest {
	s.PatchId = &v
	return s
}

func (s *ListCloudAppInstallationsShrinkRequest) SetProjectId(v string) *ListCloudAppInstallationsShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListCloudAppInstallationsShrinkRequest) SetRenderingInstanceId(v string) *ListCloudAppInstallationsShrinkRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *ListCloudAppInstallationsShrinkRequest) SetRenderingInstanceIdsShrink(v string) *ListCloudAppInstallationsShrinkRequest {
	s.RenderingInstanceIdsShrink = &v
	return s
}

func (s *ListCloudAppInstallationsShrinkRequest) SetStartTime(v string) *ListCloudAppInstallationsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListCloudAppInstallationsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
