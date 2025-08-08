// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcubeMiniPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryMiniPackageResult(v *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult) *QueryMcubeMiniPackageResponseBody
	GetQueryMiniPackageResult() *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult
	SetRequestId(v string) *QueryMcubeMiniPackageResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryMcubeMiniPackageResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *QueryMcubeMiniPackageResponseBody
	GetResultMessage() *string
}

type QueryMcubeMiniPackageResponseBody struct {
	QueryMiniPackageResult *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult `json:"QueryMiniPackageResult,omitempty" xml:"QueryMiniPackageResult,omitempty" type:"Struct"`
	RequestId              *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode             *string                                                  `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage          *string                                                  `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMcubeMiniPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeMiniPackageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMcubeMiniPackageResponseBody) GetQueryMiniPackageResult() *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult {
	return s.QueryMiniPackageResult
}

func (s *QueryMcubeMiniPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMcubeMiniPackageResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryMcubeMiniPackageResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryMcubeMiniPackageResponseBody) SetQueryMiniPackageResult(v *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult) *QueryMcubeMiniPackageResponseBody {
	s.QueryMiniPackageResult = v
	return s
}

func (s *QueryMcubeMiniPackageResponseBody) SetRequestId(v string) *QueryMcubeMiniPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBody) SetResultCode(v string) *QueryMcubeMiniPackageResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBody) SetResultMessage(v string) *QueryMcubeMiniPackageResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult struct {
	MiniPackageInfo *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo `json:"MiniPackageInfo,omitempty" xml:"MiniPackageInfo,omitempty" type:"Struct"`
	ResultMsg       *string                                                                 `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success         *bool                                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult) GoString() string {
	return s.String()
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult) GetMiniPackageInfo() *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	return s.MiniPackageInfo
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult) SetMiniPackageInfo(v *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult {
	s.MiniPackageInfo = v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult) SetResultMsg(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult {
	s.ResultMsg = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult) SetSuccess(v bool) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult {
	s.Success = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResult) Validate() error {
	return dara.Validate(s)
}

type QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo struct {
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

func (s QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GoString() string {
	return s.String()
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetAppCode() *string {
	return s.AppCode
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetAutoInstall() *int64 {
	return s.AutoInstall
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetClientVersionMax() *string {
	return s.ClientVersionMax
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetClientVersionMin() *string {
	return s.ClientVersionMin
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetExtraData() *string {
	return s.ExtraData
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetFallbackBaseUrl() *string {
	return s.FallbackBaseUrl
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetH5Id() *string {
	return s.H5Id
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetH5Name() *string {
	return s.H5Name
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetH5Version() *string {
	return s.H5Version
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetId() *int64 {
	return s.Id
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetInstallType() *int64 {
	return s.InstallType
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetMainUrl() *string {
	return s.MainUrl
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetMemo() *string {
	return s.Memo
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetPackageType() *int64 {
	return s.PackageType
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetPlatform() *string {
	return s.Platform
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetPublishPeriod() *int64 {
	return s.PublishPeriod
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetResourceType() *int64 {
	return s.ResourceType
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) GetStatus() *int64 {
	return s.Status
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetAppCode(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.AppCode = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetAutoInstall(v int64) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.AutoInstall = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetClientVersionMax(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.ClientVersionMax = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetClientVersionMin(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.ClientVersionMin = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetDownloadUrl(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.DownloadUrl = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetExtendInfo(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.ExtendInfo = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetExtraData(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.ExtraData = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetFallbackBaseUrl(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.FallbackBaseUrl = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetGmtCreate(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.GmtCreate = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetGmtModified(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.GmtModified = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetH5Id(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.H5Id = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetH5Name(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.H5Name = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetH5Version(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.H5Version = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetId(v int64) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.Id = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetInstallType(v int64) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.InstallType = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetMainUrl(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.MainUrl = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetMemo(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.Memo = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetPackageType(v int64) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.PackageType = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetPlatform(v string) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.Platform = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetPublishPeriod(v int64) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.PublishPeriod = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetResourceType(v int64) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.ResourceType = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) SetStatus(v int64) *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo {
	s.Status = &v
	return s
}

func (s *QueryMcubeMiniPackageResponseBodyQueryMiniPackageResultMiniPackageInfo) Validate() error {
	return dara.Validate(s)
}
