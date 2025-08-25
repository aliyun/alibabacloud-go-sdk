// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcubeNebulaResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGetNebulaResourceResult(v *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult) *GetMcubeNebulaResourceResponseBody
	GetGetNebulaResourceResult() *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult
	SetRequestId(v string) *GetMcubeNebulaResourceResponseBody
	GetRequestId() *string
	SetResultCode(v string) *GetMcubeNebulaResourceResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *GetMcubeNebulaResourceResponseBody
	GetResultMessage() *string
}

type GetMcubeNebulaResourceResponseBody struct {
	GetNebulaResourceResult *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult `json:"GetNebulaResourceResult,omitempty" xml:"GetNebulaResourceResult,omitempty" type:"Struct"`
	RequestId               *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode              *string                                                    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage           *string                                                    `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetMcubeNebulaResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeNebulaResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcubeNebulaResourceResponseBody) GetGetNebulaResourceResult() *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult {
	return s.GetNebulaResourceResult
}

func (s *GetMcubeNebulaResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcubeNebulaResourceResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetMcubeNebulaResourceResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *GetMcubeNebulaResourceResponseBody) SetGetNebulaResourceResult(v *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult) *GetMcubeNebulaResourceResponseBody {
	s.GetNebulaResourceResult = v
	return s
}

func (s *GetMcubeNebulaResourceResponseBody) SetRequestId(v string) *GetMcubeNebulaResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBody) SetResultCode(v string) *GetMcubeNebulaResourceResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBody) SetResultMessage(v string) *GetMcubeNebulaResourceResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult struct {
	ErrorCode          *string                                                                      `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	NebulaResourceInfo *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo `json:"NebulaResourceInfo,omitempty" xml:"NebulaResourceInfo,omitempty" type:"Struct"`
	RequestId          *string                                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg          *string                                                                      `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success            *bool                                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult) GoString() string {
	return s.String()
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult) GetNebulaResourceInfo() *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	return s.NebulaResourceInfo
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult) SetErrorCode(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult {
	s.ErrorCode = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult) SetNebulaResourceInfo(v *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult {
	s.NebulaResourceInfo = v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult) SetRequestId(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult {
	s.RequestId = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult) SetResultMsg(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult {
	s.ResultMsg = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult) SetSuccess(v bool) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult {
	s.Success = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResult) Validate() error {
	return dara.Validate(s)
}

type GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo struct {
	AppCode          *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AutoInstall      *int32  `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	ClientVersionMax *string `json:"ClientVersionMax,omitempty" xml:"ClientVersionMax,omitempty"`
	ClientVersionMin *string `json:"ClientVersionMin,omitempty" xml:"ClientVersionMin,omitempty"`
	Creator          *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
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
	MainUrl          *string `json:"MainUrl,omitempty" xml:"MainUrl,omitempty"`
	Memo             *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	MetaId           *int64  `json:"MetaId,omitempty" xml:"MetaId,omitempty"`
	Modifier         *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	PackageType      *int32  `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	Platform         *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	PublishPeriod    *int32  `json:"PublishPeriod,omitempty" xml:"PublishPeriod,omitempty"`
	ResourceType     *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Vhost            *string `json:"Vhost,omitempty" xml:"Vhost,omitempty"`
}

func (s GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GoString() string {
	return s.String()
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetAppCode() *string {
	return s.AppCode
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetAutoInstall() *int32 {
	return s.AutoInstall
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetClientVersionMax() *string {
	return s.ClientVersionMax
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetClientVersionMin() *string {
	return s.ClientVersionMin
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetExtraData() *string {
	return s.ExtraData
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetFallbackBaseUrl() *string {
	return s.FallbackBaseUrl
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetFileSize() *string {
	return s.FileSize
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetH5Id() *string {
	return s.H5Id
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetH5Name() *string {
	return s.H5Name
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetH5Version() *string {
	return s.H5Version
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetId() *int64 {
	return s.Id
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetInstallType() *int32 {
	return s.InstallType
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetMainUrl() *string {
	return s.MainUrl
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetMemo() *string {
	return s.Memo
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetMetaId() *int64 {
	return s.MetaId
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetModifier() *string {
	return s.Modifier
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetPackageType() *int32 {
	return s.PackageType
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetPlatform() *string {
	return s.Platform
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetPublishPeriod() *int32 {
	return s.PublishPeriod
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetStatus() *int32 {
	return s.Status
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) GetVhost() *string {
	return s.Vhost
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetAppCode(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.AppCode = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetAutoInstall(v int32) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.AutoInstall = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetClientVersionMax(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.ClientVersionMax = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetClientVersionMin(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.ClientVersionMin = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetCreator(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.Creator = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetDownloadUrl(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.DownloadUrl = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetExtendInfo(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.ExtendInfo = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetExtraData(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.ExtraData = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetFallbackBaseUrl(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.FallbackBaseUrl = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetFileSize(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.FileSize = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetGmtCreate(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetGmtModified(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.GmtModified = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetH5Id(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.H5Id = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetH5Name(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.H5Name = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetH5Version(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.H5Version = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetId(v int64) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.Id = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetInstallType(v int32) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.InstallType = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetMainUrl(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.MainUrl = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetMemo(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.Memo = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetMetaId(v int64) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.MetaId = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetModifier(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.Modifier = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetPackageType(v int32) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.PackageType = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetPlatform(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.Platform = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetPublishPeriod(v int32) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.PublishPeriod = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetResourceType(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.ResourceType = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetStatus(v int32) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.Status = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) SetVhost(v string) *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo {
	s.Vhost = &v
	return s
}

func (s *GetMcubeNebulaResourceResponseBodyGetNebulaResourceResultNebulaResourceInfo) Validate() error {
	return dara.Validate(s)
}
