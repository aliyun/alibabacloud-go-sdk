// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAppsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudApps(v []*ListCloudAppsResponseBodyCloudApps) *ListCloudAppsResponseBody
	GetCloudApps() []*ListCloudAppsResponseBodyCloudApps
	SetPageNumber(v int64) *ListCloudAppsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListCloudAppsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListCloudAppsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCloudAppsResponseBody
	GetTotalCount() *int64
}

type ListCloudAppsResponseBody struct {
	// The list of cloud applications.
	CloudApps []*ListCloudAppsResponseBodyCloudApps `json:"CloudApps,omitempty" xml:"CloudApps,omitempty" type:"Repeated"`
	// The page number of the returned list.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The Settings for the number of entries returned on each page during paging for a paged query.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of cloud applications that match the query conditions.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudAppsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAppsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudAppsResponseBody) GetCloudApps() []*ListCloudAppsResponseBodyCloudApps {
	return s.CloudApps
}

func (s *ListCloudAppsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListCloudAppsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListCloudAppsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudAppsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCloudAppsResponseBody) SetCloudApps(v []*ListCloudAppsResponseBodyCloudApps) *ListCloudAppsResponseBody {
	s.CloudApps = v
	return s
}

func (s *ListCloudAppsResponseBody) SetPageNumber(v int64) *ListCloudAppsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCloudAppsResponseBody) SetPageSize(v int64) *ListCloudAppsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCloudAppsResponseBody) SetRequestId(v string) *ListCloudAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudAppsResponseBody) SetTotalCount(v int64) *ListCloudAppsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCloudAppsResponseBody) Validate() error {
	if s.CloudApps != nil {
		for _, item := range s.CloudApps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloudAppsResponseBodyCloudApps struct {
	// The ID of the application.
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
	// 1.5.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The format of the installation package.
	//
	// example:
	//
	// apk
	PkgFormat *string `json:"PkgFormat,omitempty" xml:"PkgFormat,omitempty"`
	// The type of the installation package, such as android or win.
	//
	// example:
	//
	// android
	PkgType *string `json:"PkgType,omitempty" xml:"PkgType,omitempty"`
	// The relative path of the post-command in the application package. This parameter is supported only for Windows applications.
	//
	// example:
	//
	// install.ps1
	PostCommandPath *string `json:"PostCommandPath,omitempty" xml:"PostCommandPath,omitempty"`
	// The timeout period for executing the post-command, in seconds. This parameter is supported only for Windows applications.
	//
	// example:
	//
	// 10
	PostCommandTimeoutSec *int32 `json:"PostCommandTimeoutSec,omitempty" xml:"PostCommandTimeoutSec,omitempty"`
	// The ID of the stable patch package.
	//
	// example:
	//
	// patch-7bdf679812484df08a956b73e0b3bdf6
	StablePatchId *string `json:"StablePatchId,omitempty" xml:"StablePatchId,omitempty"`
	// The upload status of the application. Valid values:
	//
	// - Created
	//
	// - Doing
	//
	// - Success: This is a final state.
	//
	// - Failed: This is a final state.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The description of the status.
	//
	// example:
	//
	// upload success
	StatusDescription *string `json:"StatusDescription,omitempty" xml:"StatusDescription,omitempty"`
	// The time when the status was last updated. The time is displayed in the ISO 8601 standard with a time zone offset, in the YYYY-MM-DDTHH:mm:ss+HH:mm format.
	//
	// example:
	//
	// 2024-05-28T14:48:34+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The time when the application was uploaded. The time is displayed in the ISO 8601 standard with a time zone offset, in the YYYY-MM-DDTHH:mm:ss+HH:mm format.
	//
	// example:
	//
	// 2024-05-28T14:28:14+08:00
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s ListCloudAppsResponseBodyCloudApps) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAppsResponseBodyCloudApps) GoString() string {
	return s.String()
}

func (s *ListCloudAppsResponseBodyCloudApps) GetAppId() *string {
	return s.AppId
}

func (s *ListCloudAppsResponseBodyCloudApps) GetAppName() *string {
	return s.AppName
}

func (s *ListCloudAppsResponseBodyCloudApps) GetAppVersion() *string {
	return s.AppVersion
}

func (s *ListCloudAppsResponseBodyCloudApps) GetDescription() *string {
	return s.Description
}

func (s *ListCloudAppsResponseBodyCloudApps) GetPkgFormat() *string {
	return s.PkgFormat
}

func (s *ListCloudAppsResponseBodyCloudApps) GetPkgType() *string {
	return s.PkgType
}

func (s *ListCloudAppsResponseBodyCloudApps) GetPostCommandPath() *string {
	return s.PostCommandPath
}

func (s *ListCloudAppsResponseBodyCloudApps) GetPostCommandTimeoutSec() *int32 {
	return s.PostCommandTimeoutSec
}

func (s *ListCloudAppsResponseBodyCloudApps) GetStablePatchId() *string {
	return s.StablePatchId
}

func (s *ListCloudAppsResponseBodyCloudApps) GetStatus() *string {
	return s.Status
}

func (s *ListCloudAppsResponseBodyCloudApps) GetStatusDescription() *string {
	return s.StatusDescription
}

func (s *ListCloudAppsResponseBodyCloudApps) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCloudAppsResponseBodyCloudApps) GetUploadTime() *string {
	return s.UploadTime
}

func (s *ListCloudAppsResponseBodyCloudApps) SetAppId(v string) *ListCloudAppsResponseBodyCloudApps {
	s.AppId = &v
	return s
}

func (s *ListCloudAppsResponseBodyCloudApps) SetAppName(v string) *ListCloudAppsResponseBodyCloudApps {
	s.AppName = &v
	return s
}

func (s *ListCloudAppsResponseBodyCloudApps) SetAppVersion(v string) *ListCloudAppsResponseBodyCloudApps {
	s.AppVersion = &v
	return s
}

func (s *ListCloudAppsResponseBodyCloudApps) SetDescription(v string) *ListCloudAppsResponseBodyCloudApps {
	s.Description = &v
	return s
}

func (s *ListCloudAppsResponseBodyCloudApps) SetPkgFormat(v string) *ListCloudAppsResponseBodyCloudApps {
	s.PkgFormat = &v
	return s
}

func (s *ListCloudAppsResponseBodyCloudApps) SetPkgType(v string) *ListCloudAppsResponseBodyCloudApps {
	s.PkgType = &v
	return s
}

func (s *ListCloudAppsResponseBodyCloudApps) SetPostCommandPath(v string) *ListCloudAppsResponseBodyCloudApps {
	s.PostCommandPath = &v
	return s
}

func (s *ListCloudAppsResponseBodyCloudApps) SetPostCommandTimeoutSec(v int32) *ListCloudAppsResponseBodyCloudApps {
	s.PostCommandTimeoutSec = &v
	return s
}

func (s *ListCloudAppsResponseBodyCloudApps) SetStablePatchId(v string) *ListCloudAppsResponseBodyCloudApps {
	s.StablePatchId = &v
	return s
}

func (s *ListCloudAppsResponseBodyCloudApps) SetStatus(v string) *ListCloudAppsResponseBodyCloudApps {
	s.Status = &v
	return s
}

func (s *ListCloudAppsResponseBodyCloudApps) SetStatusDescription(v string) *ListCloudAppsResponseBodyCloudApps {
	s.StatusDescription = &v
	return s
}

func (s *ListCloudAppsResponseBodyCloudApps) SetUpdateTime(v string) *ListCloudAppsResponseBodyCloudApps {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudAppsResponseBodyCloudApps) SetUploadTime(v string) *ListCloudAppsResponseBodyCloudApps {
	s.UploadTime = &v
	return s
}

func (s *ListCloudAppsResponseBodyCloudApps) Validate() error {
	return dara.Validate(s)
}
