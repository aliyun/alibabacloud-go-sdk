// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeUpgradePackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListMcubeUpgradePackagesResponseBody
	GetCurrentPage() *int32
	SetHasMore(v bool) *ListMcubeUpgradePackagesResponseBody
	GetHasMore() *bool
	SetListPackagesResult(v *ListMcubeUpgradePackagesResponseBodyListPackagesResult) *ListMcubeUpgradePackagesResponseBody
	GetListPackagesResult() *ListMcubeUpgradePackagesResponseBodyListPackagesResult
	SetPageSize(v int32) *ListMcubeUpgradePackagesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListMcubeUpgradePackagesResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMcubeUpgradePackagesResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListMcubeUpgradePackagesResponseBody
	GetResultMessage() *string
	SetTotalCount(v int64) *ListMcubeUpgradePackagesResponseBody
	GetTotalCount() *int64
}

type ListMcubeUpgradePackagesResponseBody struct {
	CurrentPage        *int32                                                  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	HasMore            *bool                                                   `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	ListPackagesResult *ListMcubeUpgradePackagesResponseBodyListPackagesResult `json:"ListPackagesResult,omitempty" xml:"ListPackagesResult,omitempty" type:"Struct"`
	PageSize           *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode         *string                                                 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage      *string                                                 `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
	TotalCount         *int64                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMcubeUpgradePackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeUpgradePackagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcubeUpgradePackagesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListMcubeUpgradePackagesResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMcubeUpgradePackagesResponseBody) GetListPackagesResult() *ListMcubeUpgradePackagesResponseBodyListPackagesResult {
	return s.ListPackagesResult
}

func (s *ListMcubeUpgradePackagesResponseBody) GetPageSize() *int32 {
	return s.PageSize
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

func (s *ListMcubeUpgradePackagesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMcubeUpgradePackagesResponseBody) SetCurrentPage(v int32) *ListMcubeUpgradePackagesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBody) SetHasMore(v bool) *ListMcubeUpgradePackagesResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBody) SetListPackagesResult(v *ListMcubeUpgradePackagesResponseBodyListPackagesResult) *ListMcubeUpgradePackagesResponseBody {
	s.ListPackagesResult = v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBody) SetPageSize(v int32) *ListMcubeUpgradePackagesResponseBody {
	s.PageSize = &v
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

func (s *ListMcubeUpgradePackagesResponseBody) SetTotalCount(v int64) *ListMcubeUpgradePackagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMcubeUpgradePackagesResponseBodyListPackagesResult struct {
	ErrorCode *string                                                           `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	Packages  []*ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages `json:"Packages,omitempty" xml:"Packages,omitempty" type:"Repeated"`
	RequestId *string                                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string                                                           `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool                                                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMcubeUpgradePackagesResponseBodyListPackagesResult) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeUpgradePackagesResponseBodyListPackagesResult) GoString() string {
	return s.String()
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) GetPackages() []*ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	return s.Packages
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

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) SetErrorCode(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResult {
	s.ErrorCode = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) SetPackages(v []*ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) *ListMcubeUpgradePackagesResponseBodyListPackagesResult {
	s.Packages = v
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

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResult) Validate() error {
	return dara.Validate(s)
}

type ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages struct {
	AllowCreateTask  *bool   `json:"AllowCreateTask,omitempty" xml:"AllowCreateTask,omitempty"`
	AppCode          *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppstoreUrl      *string `json:"AppstoreUrl,omitempty" xml:"AppstoreUrl,omitempty"`
	ChangeLog        *string `json:"ChangeLog,omitempty" xml:"ChangeLog,omitempty"`
	Creator          *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DownloadUrl      *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	GmtCreate        *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IsEnterprise     *int32  `json:"IsEnterprise,omitempty" xml:"IsEnterprise,omitempty"`
	Md5              *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Modifier         *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	NeedCheck        *int32  `json:"NeedCheck,omitempty" xml:"NeedCheck,omitempty"`
	PackageType      *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	Platform         *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ProductId        *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductVersion   *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	PublishPeriod    *int32  `json:"PublishPeriod,omitempty" xml:"PublishPeriod,omitempty"`
	VerificationCode *string `json:"VerificationCode,omitempty" xml:"VerificationCode,omitempty"`
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

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetChangeLog() *string {
	return s.ChangeLog
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetCreator() *string {
	return s.Creator
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetId() *int64 {
	return s.Id
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetIsEnterprise() *int32 {
	return s.IsEnterprise
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

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetPackageType() *string {
	return s.PackageType
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetPlatform() *string {
	return s.Platform
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetProductId() *string {
	return s.ProductId
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetPublishPeriod() *int32 {
	return s.PublishPeriod
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) GetVerificationCode() *string {
	return s.VerificationCode
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

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetChangeLog(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.ChangeLog = &v
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

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetGmtCreate(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.GmtCreate = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetGmtModified(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.GmtModified = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetId(v int64) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.Id = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetIsEnterprise(v int32) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.IsEnterprise = &v
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

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetProductVersion(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.ProductVersion = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetPublishPeriod(v int32) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.PublishPeriod = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) SetVerificationCode(v string) *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages {
	s.VerificationCode = &v
	return s
}

func (s *ListMcubeUpgradePackagesResponseBodyListPackagesResultPackages) Validate() error {
	return dara.Validate(s)
}
