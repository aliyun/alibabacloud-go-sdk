// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeMiniPackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListMiniPackageResult(v *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) *ListMcubeMiniPackagesResponseBody
	GetListMiniPackageResult() *ListMcubeMiniPackagesResponseBodyListMiniPackageResult
	SetRequestId(v string) *ListMcubeMiniPackagesResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMcubeMiniPackagesResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListMcubeMiniPackagesResponseBody
	GetResultMessage() *string
}

type ListMcubeMiniPackagesResponseBody struct {
	ListMiniPackageResult *ListMcubeMiniPackagesResponseBodyListMiniPackageResult `json:"ListMiniPackageResult,omitempty" xml:"ListMiniPackageResult,omitempty" type:"Struct"`
	RequestId             *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode            *string                                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage         *string                                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMcubeMiniPackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeMiniPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcubeMiniPackagesResponseBody) GetListMiniPackageResult() *ListMcubeMiniPackagesResponseBodyListMiniPackageResult {
	return s.ListMiniPackageResult
}

func (s *ListMcubeMiniPackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeMiniPackagesResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMcubeMiniPackagesResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMcubeMiniPackagesResponseBody) SetListMiniPackageResult(v *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) *ListMcubeMiniPackagesResponseBody {
	s.ListMiniPackageResult = v
	return s
}

func (s *ListMcubeMiniPackagesResponseBody) SetRequestId(v string) *ListMcubeMiniPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBody) SetResultCode(v string) *ListMcubeMiniPackagesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBody) SetResultMessage(v string) *ListMcubeMiniPackagesResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMcubeMiniPackagesResponseBodyListMiniPackageResult struct {
	CurrentPage     *int32                                                                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	HasMore         *bool                                                                    `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	MiniPackageList []*ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList `json:"MiniPackageList,omitempty" xml:"MiniPackageList,omitempty" type:"Repeated"`
	PageSize        *int32                                                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResultMsg       *string                                                                  `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success         *bool                                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount      *int64                                                                   `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMcubeMiniPackagesResponseBodyListMiniPackageResult) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeMiniPackagesResponseBodyListMiniPackageResult) GoString() string {
	return s.String()
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) GetMiniPackageList() []*ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	return s.MiniPackageList
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) SetCurrentPage(v int32) *ListMcubeMiniPackagesResponseBodyListMiniPackageResult {
	s.CurrentPage = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) SetHasMore(v bool) *ListMcubeMiniPackagesResponseBodyListMiniPackageResult {
	s.HasMore = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) SetMiniPackageList(v []*ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) *ListMcubeMiniPackagesResponseBodyListMiniPackageResult {
	s.MiniPackageList = v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) SetPageSize(v int32) *ListMcubeMiniPackagesResponseBodyListMiniPackageResult {
	s.PageSize = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) SetResultMsg(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResult {
	s.ResultMsg = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) SetSuccess(v bool) *ListMcubeMiniPackagesResponseBodyListMiniPackageResult {
	s.Success = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) SetTotalCount(v int64) *ListMcubeMiniPackagesResponseBodyListMiniPackageResult {
	s.TotalCount = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResult) Validate() error {
	return dara.Validate(s)
}

type ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList struct {
	AppCode          *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AutoInstall      *int64  `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	ClientVersionMax *string `json:"ClientVersionMax,omitempty" xml:"ClientVersionMax,omitempty"`
	ClientVersionMin *string `json:"ClientVersionMin,omitempty" xml:"ClientVersionMin,omitempty"`
	DownloadUrl      *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ExtendInfo       *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	ExtraData        *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	FallbackBaseUrl  *string `json:"FallbackBaseUrl,omitempty" xml:"FallbackBaseUrl,omitempty"`
	GmtCreate        *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	H5Id             *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	H5Name           *string `json:"H5Name,omitempty" xml:"H5Name,omitempty"`
	H5Version        *string `json:"H5Version,omitempty" xml:"H5Version,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstallType      *int64  `json:"InstallType,omitempty" xml:"InstallType,omitempty"`
	MainUrl          *string `json:"MainUrl,omitempty" xml:"MainUrl,omitempty"`
	Memo             *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	PackageType      *int64  `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	Platform         *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	PublishPeriod    *int64  `json:"PublishPeriod,omitempty" xml:"PublishPeriod,omitempty"`
	ResourceType     *int64  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status           *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GoString() string {
	return s.String()
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetAppCode() *string {
	return s.AppCode
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetAutoInstall() *int64 {
	return s.AutoInstall
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetClientVersionMax() *string {
	return s.ClientVersionMax
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetClientVersionMin() *string {
	return s.ClientVersionMin
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetExtraData() *string {
	return s.ExtraData
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetFallbackBaseUrl() *string {
	return s.FallbackBaseUrl
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetH5Id() *string {
	return s.H5Id
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetH5Name() *string {
	return s.H5Name
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetH5Version() *string {
	return s.H5Version
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetId() *int64 {
	return s.Id
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetInstallType() *int64 {
	return s.InstallType
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetMainUrl() *string {
	return s.MainUrl
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetMemo() *string {
	return s.Memo
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetPackageType() *int64 {
	return s.PackageType
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetPlatform() *string {
	return s.Platform
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetPublishPeriod() *int64 {
	return s.PublishPeriod
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetResourceType() *int64 {
	return s.ResourceType
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) GetStatus() *int64 {
	return s.Status
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetAppCode(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.AppCode = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetAutoInstall(v int64) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.AutoInstall = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetClientVersionMax(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.ClientVersionMax = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetClientVersionMin(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.ClientVersionMin = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetDownloadUrl(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.DownloadUrl = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetExtendInfo(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.ExtendInfo = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetExtraData(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.ExtraData = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetFallbackBaseUrl(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.FallbackBaseUrl = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetGmtCreate(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.GmtCreate = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetGmtModified(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.GmtModified = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetH5Id(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.H5Id = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetH5Name(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.H5Name = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetH5Version(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.H5Version = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetId(v int64) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.Id = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetInstallType(v int64) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.InstallType = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetMainUrl(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.MainUrl = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetMemo(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.Memo = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetPackageType(v int64) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.PackageType = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetPlatform(v string) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.Platform = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetPublishPeriod(v int64) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.PublishPeriod = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetResourceType(v int64) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.ResourceType = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) SetStatus(v int64) *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList {
	s.Status = &v
	return s
}

func (s *ListMcubeMiniPackagesResponseBodyListMiniPackageResultMiniPackageList) Validate() error {
	return dara.Validate(s)
}
