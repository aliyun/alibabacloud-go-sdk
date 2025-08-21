// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAppInstallationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListCloudAppInstallationsRequest
	GetAppId() *string
	SetAppName(v string) *ListCloudAppInstallationsRequest
	GetAppName() *string
	SetAppVersion(v string) *ListCloudAppInstallationsRequest
	GetAppVersion() *string
	SetEndTime(v string) *ListCloudAppInstallationsRequest
	GetEndTime() *string
	SetPageNumber(v int64) *ListCloudAppInstallationsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCloudAppInstallationsRequest
	GetPageSize() *int64
	SetPatchId(v string) *ListCloudAppInstallationsRequest
	GetPatchId() *string
	SetProjectId(v string) *ListCloudAppInstallationsRequest
	GetProjectId() *string
	SetRenderingInstanceId(v string) *ListCloudAppInstallationsRequest
	GetRenderingInstanceId() *string
	SetStartTime(v string) *ListCloudAppInstallationsRequest
	GetStartTime() *string
}

type ListCloudAppInstallationsRequest struct {
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// com.aaa.bbb
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// 1.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// patch-7bdf679812484df08a956b73e0b3bdf6
	PatchId   *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListCloudAppInstallationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAppInstallationsRequest) GoString() string {
	return s.String()
}

func (s *ListCloudAppInstallationsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListCloudAppInstallationsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListCloudAppInstallationsRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListCloudAppInstallationsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListCloudAppInstallationsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCloudAppInstallationsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCloudAppInstallationsRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *ListCloudAppInstallationsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListCloudAppInstallationsRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ListCloudAppInstallationsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListCloudAppInstallationsRequest) SetAppId(v string) *ListCloudAppInstallationsRequest {
	s.AppId = &v
	return s
}

func (s *ListCloudAppInstallationsRequest) SetAppName(v string) *ListCloudAppInstallationsRequest {
	s.AppName = &v
	return s
}

func (s *ListCloudAppInstallationsRequest) SetAppVersion(v string) *ListCloudAppInstallationsRequest {
	s.AppVersion = &v
	return s
}

func (s *ListCloudAppInstallationsRequest) SetEndTime(v string) *ListCloudAppInstallationsRequest {
	s.EndTime = &v
	return s
}

func (s *ListCloudAppInstallationsRequest) SetPageNumber(v int64) *ListCloudAppInstallationsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCloudAppInstallationsRequest) SetPageSize(v int64) *ListCloudAppInstallationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudAppInstallationsRequest) SetPatchId(v string) *ListCloudAppInstallationsRequest {
	s.PatchId = &v
	return s
}

func (s *ListCloudAppInstallationsRequest) SetProjectId(v string) *ListCloudAppInstallationsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListCloudAppInstallationsRequest) SetRenderingInstanceId(v string) *ListCloudAppInstallationsRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *ListCloudAppInstallationsRequest) SetStartTime(v string) *ListCloudAppInstallationsRequest {
	s.StartTime = &v
	return s
}

func (s *ListCloudAppInstallationsRequest) Validate() error {
	return dara.Validate(s)
}
