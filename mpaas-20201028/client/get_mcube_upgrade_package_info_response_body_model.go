// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcubeUpgradePackageInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGetPackageResult(v *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult) *GetMcubeUpgradePackageInfoResponseBody
	GetGetPackageResult() *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult
	SetRequestId(v string) *GetMcubeUpgradePackageInfoResponseBody
	GetRequestId() *string
	SetResultCode(v string) *GetMcubeUpgradePackageInfoResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *GetMcubeUpgradePackageInfoResponseBody
	GetResultMessage() *string
}

type GetMcubeUpgradePackageInfoResponseBody struct {
	GetPackageResult *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult `json:"GetPackageResult,omitempty" xml:"GetPackageResult,omitempty" type:"Struct"`
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode       *string                                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage    *string                                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetMcubeUpgradePackageInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeUpgradePackageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcubeUpgradePackageInfoResponseBody) GetGetPackageResult() *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult {
	return s.GetPackageResult
}

func (s *GetMcubeUpgradePackageInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcubeUpgradePackageInfoResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetMcubeUpgradePackageInfoResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *GetMcubeUpgradePackageInfoResponseBody) SetGetPackageResult(v *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult) *GetMcubeUpgradePackageInfoResponseBody {
	s.GetPackageResult = v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBody) SetRequestId(v string) *GetMcubeUpgradePackageInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBody) SetResultCode(v string) *GetMcubeUpgradePackageInfoResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBody) SetResultMessage(v string) *GetMcubeUpgradePackageInfoResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMcubeUpgradePackageInfoResponseBodyGetPackageResult struct {
	ErrorCode   *string                                                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	PackageInfo *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfo `json:"PackageInfo,omitempty" xml:"PackageInfo,omitempty" type:"Struct"`
	RequestId   *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg   *string                                                            `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success     *bool                                                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMcubeUpgradePackageInfoResponseBodyGetPackageResult) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeUpgradePackageInfoResponseBodyGetPackageResult) GoString() string {
	return s.String()
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult) GetPackageInfo() *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfo {
	return s.PackageInfo
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult) SetErrorCode(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult {
	s.ErrorCode = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult) SetPackageInfo(v *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfo) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult {
	s.PackageInfo = v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult) SetRequestId(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult {
	s.RequestId = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult) SetResultMsg(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult {
	s.ResultMsg = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult) SetSuccess(v bool) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult {
	s.Success = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResult) Validate() error {
	return dara.Validate(s)
}

type GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfo struct {
	MobileTestFlightConfigDO *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO `json:"MobileTestFlightConfigDO,omitempty" xml:"MobileTestFlightConfigDO,omitempty" type:"Struct"`
	UpgradeBaseInfoDO        *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO        `json:"UpgradeBaseInfoDO,omitempty" xml:"UpgradeBaseInfoDO,omitempty" type:"Struct"`
}

func (s GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfo) GoString() string {
	return s.String()
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfo) GetMobileTestFlightConfigDO() *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO {
	return s.MobileTestFlightConfigDO
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfo) GetUpgradeBaseInfoDO() *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	return s.UpgradeBaseInfoDO
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfo) SetMobileTestFlightConfigDO(v *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfo {
	s.MobileTestFlightConfigDO = v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfo) SetUpgradeBaseInfoDO(v *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfo {
	s.UpgradeBaseInfoDO = v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfo) Validate() error {
	return dara.Validate(s)
}

type GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO struct {
	GmtCreate     *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified   *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id            *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InstallAmount *int32  `json:"InstallAmount,omitempty" xml:"InstallAmount,omitempty"`
	InvalidTime   *string `json:"InvalidTime,omitempty" xml:"InvalidTime,omitempty"`
	UpgradeId     *int64  `json:"UpgradeId,omitempty" xml:"UpgradeId,omitempty"`
}

func (s GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) GoString() string {
	return s.String()
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) GetId() *int64 {
	return s.Id
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) GetInstallAmount() *int32 {
	return s.InstallAmount
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) GetInvalidTime() *string {
	return s.InvalidTime
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) GetUpgradeId() *int64 {
	return s.UpgradeId
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) SetGmtCreate(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO {
	s.GmtCreate = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) SetGmtModified(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO {
	s.GmtModified = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) SetId(v int64) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO {
	s.Id = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) SetInstallAmount(v int32) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO {
	s.InstallAmount = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) SetInvalidTime(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO {
	s.InvalidTime = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) SetUpgradeId(v int64) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO {
	s.UpgradeId = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoMobileTestFlightConfigDO) Validate() error {
	return dara.Validate(s)
}

type GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO struct {
	AllowCreateTask  *bool   `json:"AllowCreateTask,omitempty" xml:"AllowCreateTask,omitempty"`
	AppCode          *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppstoreUrl      *string `json:"AppstoreUrl,omitempty" xml:"AppstoreUrl,omitempty"`
	BackLog          *string `json:"BackLog,omitempty" xml:"BackLog,omitempty"`
	ChangeLog        *string `json:"ChangeLog,omitempty" xml:"ChangeLog,omitempty"`
	ClientFileSize   *int32  `json:"ClientFileSize,omitempty" xml:"ClientFileSize,omitempty"`
	ClientName       *string `json:"ClientName,omitempty" xml:"ClientName,omitempty"`
	CpId             *string `json:"CpId,omitempty" xml:"CpId,omitempty"`
	Creator          *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DownloadUrl      *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	GlobalVariables  *string `json:"GlobalVariables,omitempty" xml:"GlobalVariables,omitempty"`
	GmtCreate        *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtCreateStr     *string `json:"GmtCreateStr,omitempty" xml:"GmtCreateStr,omitempty"`
	GmtModified      *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GmtModifiedStr   *string `json:"GmtModifiedStr,omitempty" xml:"GmtModifiedStr,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InnerVersion     *string `json:"InnerVersion,omitempty" xml:"InnerVersion,omitempty"`
	IosSymbol        *string `json:"IosSymbol,omitempty" xml:"IosSymbol,omitempty"`
	IsEnterprise     *int32  `json:"IsEnterprise,omitempty" xml:"IsEnterprise,omitempty"`
	IsRc             *int32  `json:"IsRc,omitempty" xml:"IsRc,omitempty"`
	IsRelease        *int32  `json:"IsRelease,omitempty" xml:"IsRelease,omitempty"`
	MaxVersion       *string `json:"MaxVersion,omitempty" xml:"MaxVersion,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Modifier         *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	NeedCheck        *int32  `json:"NeedCheck,omitempty" xml:"NeedCheck,omitempty"`
	OssPath          *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	PackageType      *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	Platform         *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ProductId        *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductName      *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	ProductVersion   *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	PublishPeriod    *int32  `json:"PublishPeriod,omitempty" xml:"PublishPeriod,omitempty"`
	QrcodeUrl        *string `json:"QrcodeUrl,omitempty" xml:"QrcodeUrl,omitempty"`
	ReleaseType      *string `json:"ReleaseType,omitempty" xml:"ReleaseType,omitempty"`
	ReleaseWindow    *string `json:"ReleaseWindow,omitempty" xml:"ReleaseWindow,omitempty"`
	ScmDownloadUrl   *string `json:"ScmDownloadUrl,omitempty" xml:"ScmDownloadUrl,omitempty"`
	ServerVersion    *int32  `json:"ServerVersion,omitempty" xml:"ServerVersion,omitempty"`
	VerificationCode *string `json:"VerificationCode,omitempty" xml:"VerificationCode,omitempty"`
	VerifyResult     *int32  `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
	VersionCode      *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GoString() string {
	return s.String()
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetAllowCreateTask() *bool {
	return s.AllowCreateTask
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetAppCode() *string {
	return s.AppCode
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetAppstoreUrl() *string {
	return s.AppstoreUrl
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetBackLog() *string {
	return s.BackLog
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetChangeLog() *string {
	return s.ChangeLog
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetClientFileSize() *int32 {
	return s.ClientFileSize
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetClientName() *string {
	return s.ClientName
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetCpId() *string {
	return s.CpId
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetCreator() *string {
	return s.Creator
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetGlobalVariables() *string {
	return s.GlobalVariables
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetGmtCreateStr() *string {
	return s.GmtCreateStr
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetGmtModifiedStr() *string {
	return s.GmtModifiedStr
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetId() *int64 {
	return s.Id
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetInnerVersion() *string {
	return s.InnerVersion
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetIosSymbol() *string {
	return s.IosSymbol
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetIsEnterprise() *int32 {
	return s.IsEnterprise
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetIsRc() *int32 {
	return s.IsRc
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetIsRelease() *int32 {
	return s.IsRelease
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetMaxVersion() *string {
	return s.MaxVersion
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetMd5() *string {
	return s.Md5
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetModifier() *string {
	return s.Modifier
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetNeedCheck() *int32 {
	return s.NeedCheck
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetOssPath() *string {
	return s.OssPath
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetPackageType() *string {
	return s.PackageType
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetPlatform() *string {
	return s.Platform
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetProductId() *string {
	return s.ProductId
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetProductName() *string {
	return s.ProductName
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetPublishPeriod() *int32 {
	return s.PublishPeriod
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetQrcodeUrl() *string {
	return s.QrcodeUrl
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetReleaseType() *string {
	return s.ReleaseType
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetReleaseWindow() *string {
	return s.ReleaseWindow
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetScmDownloadUrl() *string {
	return s.ScmDownloadUrl
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetServerVersion() *int32 {
	return s.ServerVersion
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetVerificationCode() *string {
	return s.VerificationCode
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetVerifyResult() *int32 {
	return s.VerifyResult
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) GetVersionCode() *string {
	return s.VersionCode
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetAllowCreateTask(v bool) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.AllowCreateTask = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetAppCode(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.AppCode = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetAppstoreUrl(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.AppstoreUrl = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetBackLog(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.BackLog = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetChangeLog(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.ChangeLog = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetClientFileSize(v int32) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.ClientFileSize = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetClientName(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.ClientName = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetCpId(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.CpId = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetCreator(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.Creator = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetDownloadUrl(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.DownloadUrl = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetGlobalVariables(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.GlobalVariables = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetGmtCreate(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.GmtCreate = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetGmtCreateStr(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.GmtCreateStr = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetGmtModified(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.GmtModified = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetGmtModifiedStr(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.GmtModifiedStr = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetId(v int64) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.Id = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetInnerVersion(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.InnerVersion = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetIosSymbol(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.IosSymbol = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetIsEnterprise(v int32) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.IsEnterprise = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetIsRc(v int32) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.IsRc = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetIsRelease(v int32) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.IsRelease = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetMaxVersion(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.MaxVersion = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetMd5(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.Md5 = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetModifier(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.Modifier = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetNeedCheck(v int32) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.NeedCheck = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetOssPath(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.OssPath = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetPackageType(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.PackageType = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetPlatform(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.Platform = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetProductId(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.ProductId = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetProductName(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.ProductName = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetProductVersion(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.ProductVersion = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetPublishPeriod(v int32) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.PublishPeriod = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetQrcodeUrl(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.QrcodeUrl = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetReleaseType(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.ReleaseType = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetReleaseWindow(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.ReleaseWindow = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetScmDownloadUrl(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.ScmDownloadUrl = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetServerVersion(v int32) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.ServerVersion = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetVerificationCode(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.VerificationCode = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetVerifyResult(v int32) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.VerifyResult = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) SetVersionCode(v string) *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO {
	s.VersionCode = &v
	return s
}

func (s *GetMcubeUpgradePackageInfoResponseBodyGetPackageResultPackageInfoUpgradeBaseInfoDO) Validate() error {
	return dara.Validate(s)
}
