// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAppInstallationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstallationInfos(v []*ListCloudAppInstallationsResponseBodyInstallationInfos) *ListCloudAppInstallationsResponseBody
	GetInstallationInfos() []*ListCloudAppInstallationsResponseBodyInstallationInfos
	SetPageNumber(v int64) *ListCloudAppInstallationsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCloudAppInstallationsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListCloudAppInstallationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCloudAppInstallationsResponseBody
	GetTotalCount() *int64
}

type ListCloudAppInstallationsResponseBody struct {
	InstallationInfos []*ListCloudAppInstallationsResponseBodyInstallationInfos `json:"InstallationInfos,omitempty" xml:"InstallationInfos,omitempty" type:"Repeated"`
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
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudAppInstallationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAppInstallationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudAppInstallationsResponseBody) GetInstallationInfos() []*ListCloudAppInstallationsResponseBodyInstallationInfos {
	return s.InstallationInfos
}

func (s *ListCloudAppInstallationsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCloudAppInstallationsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCloudAppInstallationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudAppInstallationsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCloudAppInstallationsResponseBody) SetInstallationInfos(v []*ListCloudAppInstallationsResponseBodyInstallationInfos) *ListCloudAppInstallationsResponseBody {
	s.InstallationInfos = v
	return s
}

func (s *ListCloudAppInstallationsResponseBody) SetPageNumber(v int64) *ListCloudAppInstallationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCloudAppInstallationsResponseBody) SetPageSize(v int64) *ListCloudAppInstallationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCloudAppInstallationsResponseBody) SetRequestId(v string) *ListCloudAppInstallationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudAppInstallationsResponseBody) SetTotalCount(v int64) *ListCloudAppInstallationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCloudAppInstallationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCloudAppInstallationsResponseBodyInstallationInfos struct {
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
	// 1.5.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// 2024-05-28T14:48:34+08:00
	InstallationTime *string `json:"InstallationTime,omitempty" xml:"InstallationTime,omitempty"`
	// example:
	//
	// patch-7bdf679812484df08a956b73e0b3bdf6
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// example:
	//
	// installed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// install success
	StatusDescription *string `json:"StatusDescription,omitempty" xml:"StatusDescription,omitempty"`
	// example:
	//
	// 2024-05-28T14:50:04+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListCloudAppInstallationsResponseBodyInstallationInfos) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAppInstallationsResponseBodyInstallationInfos) GoString() string {
	return s.String()
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) GetAppId() *string {
	return s.AppId
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) GetAppName() *string {
	return s.AppName
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) GetInstallationTime() *string {
	return s.InstallationTime
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) GetPatchId() *string {
	return s.PatchId
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) GetStatus() *string {
	return s.Status
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) GetStatusDescription() *string {
	return s.StatusDescription
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) SetAppId(v string) *ListCloudAppInstallationsResponseBodyInstallationInfos {
	s.AppId = &v
	return s
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) SetAppName(v string) *ListCloudAppInstallationsResponseBodyInstallationInfos {
	s.AppName = &v
	return s
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) SetAppVersion(v string) *ListCloudAppInstallationsResponseBodyInstallationInfos {
	s.AppVersion = &v
	return s
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) SetInstallationTime(v string) *ListCloudAppInstallationsResponseBodyInstallationInfos {
	s.InstallationTime = &v
	return s
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) SetPatchId(v string) *ListCloudAppInstallationsResponseBodyInstallationInfos {
	s.PatchId = &v
	return s
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) SetRenderingInstanceId(v string) *ListCloudAppInstallationsResponseBodyInstallationInfos {
	s.RenderingInstanceId = &v
	return s
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) SetStatus(v string) *ListCloudAppInstallationsResponseBodyInstallationInfos {
	s.Status = &v
	return s
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) SetStatusDescription(v string) *ListCloudAppInstallationsResponseBodyInstallationInfos {
	s.StatusDescription = &v
	return s
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) SetUpdateTime(v string) *ListCloudAppInstallationsResponseBodyInstallationInfos {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudAppInstallationsResponseBodyInstallationInfos) Validate() error {
	return dara.Validate(s)
}
