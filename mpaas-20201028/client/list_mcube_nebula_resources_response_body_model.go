// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeNebulaResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListMcubeNebulaResourceResult(v *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) *ListMcubeNebulaResourcesResponseBody
	GetListMcubeNebulaResourceResult() *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult
	SetRequestId(v string) *ListMcubeNebulaResourcesResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMcubeNebulaResourcesResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListMcubeNebulaResourcesResponseBody
	GetResultMessage() *string
}

type ListMcubeNebulaResourcesResponseBody struct {
	ListMcubeNebulaResourceResult *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult `json:"ListMcubeNebulaResourceResult,omitempty" xml:"ListMcubeNebulaResourceResult,omitempty" type:"Struct"`
	RequestId                     *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                    *string                                                            `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage                 *string                                                            `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMcubeNebulaResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeNebulaResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcubeNebulaResourcesResponseBody) GetListMcubeNebulaResourceResult() *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult {
	return s.ListMcubeNebulaResourceResult
}

func (s *ListMcubeNebulaResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeNebulaResourcesResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMcubeNebulaResourcesResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMcubeNebulaResourcesResponseBody) SetListMcubeNebulaResourceResult(v *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) *ListMcubeNebulaResourcesResponseBody {
	s.ListMcubeNebulaResourceResult = v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBody) SetRequestId(v string) *ListMcubeNebulaResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBody) SetResultCode(v string) *ListMcubeNebulaResourcesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBody) SetResultMessage(v string) *ListMcubeNebulaResourcesResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult struct {
	CurrentPage        *int32                                                                                 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ErrorCode          *string                                                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HasMore            *bool                                                                                  `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	NebulaResourceInfo []*ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo `json:"NebulaResourceInfo,omitempty" xml:"NebulaResourceInfo,omitempty" type:"Repeated"`
	PageSize           *int32                                                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg          *string                                                                                `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success            *bool                                                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount         *int64                                                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) GoString() string {
	return s.String()
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) GetNebulaResourceInfo() []*ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	return s.NebulaResourceInfo
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) SetCurrentPage(v int32) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult {
	s.CurrentPage = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) SetErrorCode(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult {
	s.ErrorCode = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) SetHasMore(v bool) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult {
	s.HasMore = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) SetNebulaResourceInfo(v []*ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult {
	s.NebulaResourceInfo = v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) SetPageSize(v int32) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult {
	s.PageSize = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) SetRequestId(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult {
	s.RequestId = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) SetResultMsg(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult {
	s.ResultMsg = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) SetSuccess(v bool) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult {
	s.Success = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) SetTotalCount(v int64) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult {
	s.TotalCount = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResult) Validate() error {
	return dara.Validate(s)
}

type ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo struct {
	AppCode          *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AutoInstall      *int32  `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	ClientVersionMax *string `json:"ClientVersionMax,omitempty" xml:"ClientVersionMax,omitempty"`
	ClientVersionMin *string `json:"ClientVersionMin,omitempty" xml:"ClientVersionMin,omitempty"`
	Creator          *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DebugUrl         *string `json:"DebugUrl,omitempty" xml:"DebugUrl,omitempty"`
	DownloadUrl      *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ExtendInfo       *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	ExtraData        *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	FallbackBaseUrl  *string `json:"FallbackBaseUrl,omitempty" xml:"FallbackBaseUrl,omitempty"`
	FileSize         *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	GmtCreate        *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	H5Id             *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	H5Name           *string `json:"H5Name,omitempty" xml:"H5Name,omitempty"`
	H5Version        *string `json:"H5Version,omitempty" xml:"H5Version,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstallType      *int32  `json:"InstallType,omitempty" xml:"InstallType,omitempty"`
	LazyLoad         *int32  `json:"LazyLoad,omitempty" xml:"LazyLoad,omitempty"`
	MainUrl          *string `json:"MainUrl,omitempty" xml:"MainUrl,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Memo             *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	MetaId           *int64  `json:"MetaId,omitempty" xml:"MetaId,omitempty"`
	Modifier         *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	PackageType      *int32  `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	Platform         *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	PublishPeriod    *int32  `json:"PublishPeriod,omitempty" xml:"PublishPeriod,omitempty"`
	ReleaseVersion   *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	ResourceType     *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Vhost            *string `json:"Vhost,omitempty" xml:"Vhost,omitempty"`
}

func (s ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GoString() string {
	return s.String()
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetAppCode() *string {
	return s.AppCode
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetAutoInstall() *int32 {
	return s.AutoInstall
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetClientVersionMax() *string {
	return s.ClientVersionMax
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetClientVersionMin() *string {
	return s.ClientVersionMin
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetCreator() *string {
	return s.Creator
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetDebugUrl() *string {
	return s.DebugUrl
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetExtraData() *string {
	return s.ExtraData
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetFallbackBaseUrl() *string {
	return s.FallbackBaseUrl
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetH5Id() *string {
	return s.H5Id
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetH5Name() *string {
	return s.H5Name
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetH5Version() *string {
	return s.H5Version
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetId() *int64 {
	return s.Id
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetInstallType() *int32 {
	return s.InstallType
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetLazyLoad() *int32 {
	return s.LazyLoad
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetMainUrl() *string {
	return s.MainUrl
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetMd5() *string {
	return s.Md5
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetMemo() *string {
	return s.Memo
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetMetaId() *int64 {
	return s.MetaId
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetModifier() *string {
	return s.Modifier
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetPackageType() *int32 {
	return s.PackageType
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetPlatform() *string {
	return s.Platform
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetPublishPeriod() *int32 {
	return s.PublishPeriod
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetStatus() *int32 {
	return s.Status
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) GetVhost() *string {
	return s.Vhost
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetAppCode(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.AppCode = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetAutoInstall(v int32) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.AutoInstall = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetClientVersionMax(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.ClientVersionMax = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetClientVersionMin(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.ClientVersionMin = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetCreator(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.Creator = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetDebugUrl(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.DebugUrl = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetDownloadUrl(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.DownloadUrl = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetExtendInfo(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.ExtendInfo = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetExtraData(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.ExtraData = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetFallbackBaseUrl(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.FallbackBaseUrl = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetFileSize(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.FileSize = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetGmtCreate(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.GmtCreate = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetGmtModified(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.GmtModified = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetH5Id(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.H5Id = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetH5Name(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.H5Name = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetH5Version(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.H5Version = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetId(v int64) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.Id = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetInstallType(v int32) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.InstallType = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetLazyLoad(v int32) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.LazyLoad = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetMainUrl(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.MainUrl = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetMd5(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.Md5 = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetMemo(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.Memo = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetMetaId(v int64) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.MetaId = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetModifier(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.Modifier = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetPackageType(v int32) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.PackageType = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetPlatform(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.Platform = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetPublishPeriod(v int32) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.PublishPeriod = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetReleaseVersion(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.ReleaseVersion = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetResourceType(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.ResourceType = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetStatus(v int32) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.Status = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) SetVhost(v string) *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo {
	s.Vhost = &v
	return s
}

func (s *ListMcubeNebulaResourcesResponseBodyListMcubeNebulaResourceResultNebulaResourceInfo) Validate() error {
	return dara.Validate(s)
}
