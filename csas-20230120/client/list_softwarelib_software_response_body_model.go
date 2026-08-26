// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSoftwarelibSoftwareResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*ListSoftwarelibSoftwareResponseBodyDataList) *ListSoftwarelibSoftwareResponseBody
	GetDataList() []*ListSoftwarelibSoftwareResponseBodyDataList
	SetMaxResults(v int32) *ListSoftwarelibSoftwareResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSoftwarelibSoftwareResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSoftwarelibSoftwareResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListSoftwarelibSoftwareResponseBody
	GetTotalCount() *int32
}

type ListSoftwarelibSoftwareResponseBody struct {
	// The software list.
	DataList []*ListSoftwarelibSoftwareResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The maximum number of entries per page. This parameter is not returned by this operation.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. This parameter is not returned by this operation.
	//
	// example:
	//
	// FFM+3L1WZbKngBeLWcDmQrzLuGDDwAw7JA5q2AjvTSJm9WyhQ0MwJoOWpky9ZhgcWfIgtGpZ+4NQX97+EIwsqUNQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BE4FB974-11BC-5453-9BE1-1606A73EACA6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of software entries that match the query conditions.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSoftwarelibSoftwareResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwarelibSoftwareResponseBody) GoString() string {
	return s.String()
}

func (s *ListSoftwarelibSoftwareResponseBody) GetDataList() []*ListSoftwarelibSoftwareResponseBodyDataList {
	return s.DataList
}

func (s *ListSoftwarelibSoftwareResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSoftwarelibSoftwareResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSoftwarelibSoftwareResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSoftwarelibSoftwareResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSoftwarelibSoftwareResponseBody) SetDataList(v []*ListSoftwarelibSoftwareResponseBodyDataList) *ListSoftwarelibSoftwareResponseBody {
	s.DataList = v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBody) SetMaxResults(v int32) *ListSoftwarelibSoftwareResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBody) SetNextToken(v string) *ListSoftwarelibSoftwareResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBody) SetRequestId(v string) *ListSoftwarelibSoftwareResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBody) SetTotalCount(v int32) *ListSoftwarelibSoftwareResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSoftwarelibSoftwareResponseBodyDataList struct {
	// The associated built-in software ID.
	//
	// example:
	//
	// softwarelib-software-1da844a39729****
	BuiltinSoftwareId *string `json:"BuiltinSoftwareId,omitempty" xml:"BuiltinSoftwareId,omitempty"`
	// The software BundleId used for verification.
	//
	// example:
	//
	// test software
	CheckBundleId *string `json:"CheckBundleId,omitempty" xml:"CheckBundleId,omitempty"`
	// The software name used for verification.
	//
	// example:
	//
	// test software
	CheckSoftwareName *string `json:"CheckSoftwareName,omitempty" xml:"CheckSoftwareName,omitempty"`
	// The software classification ID.
	//
	// example:
	//
	// softwarelib-classify-61b7ccc63cae****
	ClassifyId *string `json:"ClassifyId,omitempty" xml:"ClassifyId,omitempty"`
	// The time when the software was created, in seconds-level UNIX timestamp.
	//
	// example:
	//
	// 1781748302
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The software description.
	//
	// example:
	//
	// This is a demo software.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The list of associated terminal device IDs.
	DevTags []*string `json:"DevTags,omitempty" xml:"DevTags,omitempty" type:"Repeated"`
	// The list of associated device group IDs.
	DeviceGroupIds []*string `json:"DeviceGroupIds,omitempty" xml:"DeviceGroupIds,omitempty" type:"Repeated"`
	// The number of times the software has been manually downloaded from the client.
	//
	// example:
	//
	// 1
	DownloadTimes *int64 `json:"DownloadTimes,omitempty" xml:"DownloadTimes,omitempty"`
	// Indicates whether a new version is available for the software.
	//
	// example:
	//
	// false
	HasNewVersion *bool `json:"HasNewVersion,omitempty" xml:"HasNewVersion,omitempty"`
	// The URL of the software logo.
	//
	// example:
	//
	// https://img.alicdn.com/imgextra/i4/O1CN01mXKAQX1P3a5fbS0Dp_!!6000000001785-2-tps-40-40.png
	LogoUrl *string `json:"LogoUrl,omitempty" xml:"LogoUrl,omitempty"`
	// The latest software version number for Mac (Apple).
	//
	// example:
	//
	// 1.0.0
	MacAppleVersion *string `json:"MacAppleVersion,omitempty" xml:"MacAppleVersion,omitempty"`
	// The latest software version number for Mac (Intel).
	//
	// example:
	//
	// 1.0.0
	MacIntelVersion *string `json:"MacIntelVersion,omitempty" xml:"MacIntelVersion,omitempty"`
	// The policy matching target type. Valid values:
	//
	// - **UserGroupAll**: all users.
	//
	// - **UserGroupNormal**: specified user groups.
	//
	// - **DevTagNormal**: specified devices.
	//
	// - **DeviceGroupNormal**: specified device groups.
	//
	// - **DevTagAll**: all devices.
	//
	// - **None**: not configured.
	//
	// example:
	//
	// UserGroupAll
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The official download URL of the software.
	//
	// example:
	//
	// http://xxx.com/****
	OfficialDownloadUrl *string `json:"OfficialDownloadUrl,omitempty" xml:"OfficialDownloadUrl,omitempty"`
	// The execution account (only supported on Windows).
	//
	// example:
	//
	// admin
	RunAsAccount *string `json:"RunAsAccount,omitempty" xml:"RunAsAccount,omitempty"`
	// The software ID.
	//
	// example:
	//
	// softwarelib-software-1da844a39729****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	// The software name.
	//
	// example:
	//
	// Thunder
	SoftwareName *string `json:"SoftwareName,omitempty" xml:"SoftwareName,omitempty"`
	// Indicates whether the built-in library source has been deleted.
	//
	// example:
	//
	// false
	SourceRemoved *bool `json:"SourceRemoved,omitempty" xml:"SourceRemoved,omitempty"`
	// The software source. Valid values:
	//
	// - **custom**: custom software.
	//
	// - **builtin**: built-in software library.
	//
	// example:
	//
	// custom
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The list of associated user group IDs.
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
	// The software version list. This field is not returned by this operation. Call [ListSoftwarelibVersion](~~ListSoftwarelibVersion~~) to query software versions.
	Versions []*ListSoftwarelibSoftwareResponseBodyDataListVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
	// The latest software version number for Windows.
	//
	// example:
	//
	// 1.0.0
	WindowsVersion *string `json:"WindowsVersion,omitempty" xml:"WindowsVersion,omitempty"`
}

func (s ListSoftwarelibSoftwareResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwarelibSoftwareResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetBuiltinSoftwareId() *string {
	return s.BuiltinSoftwareId
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetCheckBundleId() *string {
	return s.CheckBundleId
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetCheckSoftwareName() *string {
	return s.CheckSoftwareName
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetClassifyId() *string {
	return s.ClassifyId
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetDevTags() []*string {
	return s.DevTags
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetDeviceGroupIds() []*string {
	return s.DeviceGroupIds
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetDownloadTimes() *int64 {
	return s.DownloadTimes
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetHasNewVersion() *bool {
	return s.HasNewVersion
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetLogoUrl() *string {
	return s.LogoUrl
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetMacAppleVersion() *string {
	return s.MacAppleVersion
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetMacIntelVersion() *string {
	return s.MacIntelVersion
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetOfficialDownloadUrl() *string {
	return s.OfficialDownloadUrl
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetRunAsAccount() *string {
	return s.RunAsAccount
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetSoftwareName() *string {
	return s.SoftwareName
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetSourceRemoved() *bool {
	return s.SourceRemoved
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetSourceType() *string {
	return s.SourceType
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetVersions() []*ListSoftwarelibSoftwareResponseBodyDataListVersions {
	return s.Versions
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) GetWindowsVersion() *string {
	return s.WindowsVersion
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetBuiltinSoftwareId(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.BuiltinSoftwareId = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetCheckBundleId(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.CheckBundleId = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetCheckSoftwareName(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.CheckSoftwareName = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetClassifyId(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.ClassifyId = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetCreateTime(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetDescription(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetDevTags(v []*string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.DevTags = v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetDeviceGroupIds(v []*string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.DeviceGroupIds = v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetDownloadTimes(v int64) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.DownloadTimes = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetHasNewVersion(v bool) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.HasNewVersion = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetLogoUrl(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.LogoUrl = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetMacAppleVersion(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.MacAppleVersion = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetMacIntelVersion(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.MacIntelVersion = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetMatchMode(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.MatchMode = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetOfficialDownloadUrl(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.OfficialDownloadUrl = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetRunAsAccount(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.RunAsAccount = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetSoftwareId(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.SoftwareId = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetSoftwareName(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.SoftwareName = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetSourceRemoved(v bool) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.SourceRemoved = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetSourceType(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.SourceType = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetUserGroupIds(v []*string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.UserGroupIds = v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetVersions(v []*ListSoftwarelibSoftwareResponseBodyDataListVersions) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.Versions = v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) SetWindowsVersion(v string) *ListSoftwarelibSoftwareResponseBodyDataList {
	s.WindowsVersion = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataList) Validate() error {
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSoftwarelibSoftwareResponseBodyDataListVersions struct {
	// The time when the software version was created.
	//
	// example:
	//
	// 2026-08-05 18:03:58
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of times the software has been downloaded from the client.
	//
	// example:
	//
	// 1
	DownloadTimes *int64 `json:"DownloadTimes,omitempty" xml:"DownloadTimes,omitempty"`
	// The MD5 value of the software package.
	//
	// example:
	//
	// 0b5824cdd509d3ed560e2d20d29a1bcb
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The time when the software version was last modified.
	//
	// example:
	//
	// 2026-08-05 18:03:58
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The operating system to which the software package applies. Valid values:
	//
	// - **Windows**: Windows.
	//
	// - **Mac(Apple)**: macOS with Apple silicon.
	//
	// - **Mac(Intel)**: macOS with Intel processors.
	//
	// example:
	//
	// Windows
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The software publisher type. Valid values:
	//
	// - **local**: locally uploaded.
	//
	// - **thirdparty**: third-party link.
	//
	// example:
	//
	// local
	PublisherType *string `json:"PublisherType,omitempty" xml:"PublisherType,omitempty"`
	// The ID of the software to which this version belongs.
	//
	// example:
	//
	// softwarelib-software-1da844a39729****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	// The name of the software package.
	//
	// example:
	//
	// test softwarename
	SoftwarePkgName *string `json:"SoftwarePkgName,omitempty" xml:"SoftwarePkgName,omitempty"`
	// The size of the software package.
	//
	// example:
	//
	// 100
	SoftwarePkgSize *int64 `json:"SoftwarePkgSize,omitempty" xml:"SoftwarePkgSize,omitempty"`
	// The download URL of the software package.
	//
	// example:
	//
	// https://****.com/****
	SoftwareUrl *string `json:"SoftwareUrl,omitempty" xml:"SoftwareUrl,omitempty"`
	// The version publish status. Valid values:
	//
	// - **published**: published.
	//
	// - **unpublished**: not published.
	//
	// example:
	//
	// published
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The software version number.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The software version ID.
	//
	// example:
	//
	// softwarelib-version-21ae186e2ac9****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListSoftwarelibSoftwareResponseBodyDataListVersions) String() string {
	return dara.Prettify(s)
}

func (s ListSoftwarelibSoftwareResponseBodyDataListVersions) GoString() string {
	return s.String()
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) GetDownloadTimes() *int64 {
	return s.DownloadTimes
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) GetMd5() *string {
	return s.Md5
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) GetOs() *string {
	return s.Os
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) GetPublisherType() *string {
	return s.PublisherType
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) GetSoftwarePkgName() *string {
	return s.SoftwarePkgName
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) GetSoftwarePkgSize() *int64 {
	return s.SoftwarePkgSize
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) GetSoftwareUrl() *string {
	return s.SoftwareUrl
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) GetStatus() *string {
	return s.Status
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) GetVersion() *string {
	return s.Version
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) GetVersionId() *string {
	return s.VersionId
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) SetCreateTime(v string) *ListSoftwarelibSoftwareResponseBodyDataListVersions {
	s.CreateTime = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) SetDownloadTimes(v int64) *ListSoftwarelibSoftwareResponseBodyDataListVersions {
	s.DownloadTimes = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) SetMd5(v string) *ListSoftwarelibSoftwareResponseBodyDataListVersions {
	s.Md5 = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) SetModifyTime(v string) *ListSoftwarelibSoftwareResponseBodyDataListVersions {
	s.ModifyTime = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) SetOs(v string) *ListSoftwarelibSoftwareResponseBodyDataListVersions {
	s.Os = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) SetPublisherType(v string) *ListSoftwarelibSoftwareResponseBodyDataListVersions {
	s.PublisherType = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) SetSoftwareId(v string) *ListSoftwarelibSoftwareResponseBodyDataListVersions {
	s.SoftwareId = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) SetSoftwarePkgName(v string) *ListSoftwarelibSoftwareResponseBodyDataListVersions {
	s.SoftwarePkgName = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) SetSoftwarePkgSize(v int64) *ListSoftwarelibSoftwareResponseBodyDataListVersions {
	s.SoftwarePkgSize = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) SetSoftwareUrl(v string) *ListSoftwarelibSoftwareResponseBodyDataListVersions {
	s.SoftwareUrl = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) SetStatus(v string) *ListSoftwarelibSoftwareResponseBodyDataListVersions {
	s.Status = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) SetVersion(v string) *ListSoftwarelibSoftwareResponseBodyDataListVersions {
	s.Version = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) SetVersionId(v string) *ListSoftwarelibSoftwareResponseBodyDataListVersions {
	s.VersionId = &v
	return s
}

func (s *ListSoftwarelibSoftwareResponseBodyDataListVersions) Validate() error {
	return dara.Validate(s)
}
