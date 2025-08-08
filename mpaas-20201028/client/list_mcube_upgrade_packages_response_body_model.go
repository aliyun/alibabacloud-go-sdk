// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeUpgradePackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListPackagesResult(v *ListMcubeUpgradePackagesResponseBodyListPackagesResult) *ListMcubeUpgradePackagesResponseBody
	GetListPackagesResult() *ListMcubeUpgradePackagesResponseBodyListPackagesResult
	SetRequestId(v string) *ListMcubeUpgradePackagesResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMcubeUpgradePackagesResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListMcubeUpgradePackagesResponseBody
	GetResultMessage() *string
}

type ListMcubeUpgradePackagesResponseBody struct {
	ListPackagesResult *ListMcubeUpgradePackagesResponseBodyListPackagesResult `json:"ListPackagesResult,omitempty" xml:"ListPackagesResult,omitempty" type:"Struct"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string                                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string                                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMcubeUpgradePackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeUpgradePackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcubeUpgradePackagesResponseBody) GetListPackagesResult() *ListMcubeUpgradePackagesResponseBodyListPackagesResult {
	return s.ListPackagesResult
}

func (s *ListMcubeUpgradePackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeUpgradePackagesResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMcubeUpgradePackagesResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMcubeUpgradePackagesResponseBody) SetListPackagesResult(v *ListMcubeUpgradePackagesResponseBodyListPackagesResult) *ListMcubeUpgradePackagesResponseBody {
	s.ListPackagesResult = v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBody) SetRequestId(v string) *ListMcubeUpgradePackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBody) SetResultCode(v string) *ListMcubeUpgradePackagesResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBody) SetResultMessage(v string) *ListMcubeUpgradePackagesResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMcubeUpgradePackagesResponseBodyListPackagesResult struct {
	CurrentPage *int32                                                            `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	ErrorCode   *string                                                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HasMore     *bool                                                             `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	Packages    []*ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Repeated"`
	PageSize    *int32                                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg   *string                                                           `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success     *bool                                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int64                                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMcubeUpgradePackagesResponseBodyListPackagesResult) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeUpgradePackagesResponseBodyListPackagesResult) GoString() string {
	return s.String()
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) GetPackages() []*ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	return s.Packages
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) SetCurrentPage(v int32) *ListMcubeUpgradePackagesResponseBodyListPackagesResult {
	s.CurrentPage = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) SetErrorCode(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResult {
	s.ErrorCode = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) SetHasMore(v bool) *ListMcubeUpgradePackagesResponseBodyListPackagesResult {
	s.HasMore = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) SetPackages(v []*ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) *ListMcubeUpgradePackagesResponseBodyListPackagesResult {
	s.Packages = v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) SetPageSize(v int32) *ListMcubeUpgradePackagesResponseBodyListPackagesResult {
	s.PageSize = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) SetRequestId(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResult {
	s.RequestId = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) SetResultMsg(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResult {
	s.ResultMsg = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) SetSuccess(v bool) *ListMcubeUpgradePackagesResponseBodyListPackagesResult {
	s.Success = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) SetTotalCount(v int64) *ListMcubeUpgradePackagesResponseBodyListPackagesResult {
	s.TotalCount = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) Validate() error {
	return dara.Validate(s)
}

type ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages struct {
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

func (s ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GoString() string {
	return s.String()
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetAllowCreateTask() *bool {
	return s.AllowCreateTask
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetAppCode() *string {
	return s.AppCode
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetAppstoreUrl() *string {
	return s.AppstoreUrl
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetBackLog() *string {
	return s.BackLog
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetChangeLog() *string {
	return s.ChangeLog
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetClientFileSize() *int32 {
	return s.ClientFileSize
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetClientName() *string {
	return s.ClientName
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetCpId() *string {
	return s.CpId
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetCreator() *string {
	return s.Creator
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetGlobalVariables() *string {
	return s.GlobalVariables
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetGmtCreateStr() *string {
	return s.GmtCreateStr
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetGmtModifiedStr() *string {
	return s.GmtModifiedStr
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetId() *int64 {
	return s.Id
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetInnerVersion() *string {
	return s.InnerVersion
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetIosSymbol() *string {
	return s.IosSymbol
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetIsEnterprise() *int32 {
	return s.IsEnterprise
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetIsRc() *int32 {
	return s.IsRc
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetIsRelease() *int32 {
	return s.IsRelease
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetMaxVersion() *string {
	return s.MaxVersion
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetMd5() *string {
	return s.Md5
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetModifier() *string {
	return s.Modifier
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetNeedCheck() *int32 {
	return s.NeedCheck
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetOssPath() *string {
	return s.OssPath
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetPackageType() *string {
	return s.PackageType
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetPlatform() *string {
	return s.Platform
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetProductId() *string {
	return s.ProductId
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetProductName() *string {
	return s.ProductName
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetPublishPeriod() *int32 {
	return s.PublishPeriod
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetQrcodeUrl() *string {
	return s.QrcodeUrl
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetReleaseType() *string {
	return s.ReleaseType
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetReleaseWindow() *string {
	return s.ReleaseWindow
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetScmDownloadUrl() *string {
	return s.ScmDownloadUrl
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetServerVersion() *int32 {
	return s.ServerVersion
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetVerificationCode() *string {
	return s.VerificationCode
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetVerifyResult() *int32 {
	return s.VerifyResult
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetVersionCode() *string {
	return s.VersionCode
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetAllowCreateTask(v bool) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.AllowCreateTask = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetAppCode(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.AppCode = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetAppstoreUrl(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.AppstoreUrl = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetBackLog(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.BackLog = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetChangeLog(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.ChangeLog = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetClientFileSize(v int32) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.ClientFileSize = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetClientName(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.ClientName = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetCpId(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.CpId = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetCreator(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.Creator = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetDownloadUrl(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.DownloadUrl = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetGlobalVariables(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.GlobalVariables = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetGmtCreate(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.GmtCreate = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetGmtCreateStr(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.GmtCreateStr = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetGmtModified(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.GmtModified = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetGmtModifiedStr(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.GmtModifiedStr = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetId(v int64) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.Id = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetInnerVersion(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.InnerVersion = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetIosSymbol(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.IosSymbol = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetIsEnterprise(v int32) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.IsEnterprise = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetIsRc(v int32) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.IsRc = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetIsRelease(v int32) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.IsRelease = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetMaxVersion(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.MaxVersion = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetMd5(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.Md5 = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetModifier(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.Modifier = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetNeedCheck(v int32) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.NeedCheck = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetOssPath(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.OssPath = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetPackageType(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.PackageType = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetPlatform(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.Platform = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetProductId(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.ProductId = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetProductName(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.ProductName = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetProductVersion(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.ProductVersion = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetPublishPeriod(v int32) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.PublishPeriod = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetQrcodeUrl(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.QrcodeUrl = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetReleaseType(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.ReleaseType = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetReleaseWindow(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.ReleaseWindow = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetScmDownloadUrl(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.ScmDownloadUrl = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetServerVersion(v int32) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.ServerVersion = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetVerificationCode(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.VerificationCode = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetVerifyResult(v int32) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.VerifyResult = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetVersionCode(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.VersionCode = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) Validate() error {
	return dara.Validate(s)
}
