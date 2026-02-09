// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcubeHotpatchTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryHotpatchTaskDetailResult(v *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult) *QueryMcubeHotpatchTaskDetailResponseBody
	GetQueryHotpatchTaskDetailResult() *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult
	SetRequestId(v string) *QueryMcubeHotpatchTaskDetailResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryMcubeHotpatchTaskDetailResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *QueryMcubeHotpatchTaskDetailResponseBody
	GetResultMessage() *string
}

type QueryMcubeHotpatchTaskDetailResponseBody struct {
	QueryHotpatchTaskDetailResult *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult `json:"QueryHotpatchTaskDetailResult,omitempty" xml:"QueryHotpatchTaskDetailResult,omitempty" type:"Struct"`
	// example:
	//
	// 11E66B29-9E5E-5C10-B64E-B5A0E0F26355
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// OK
	ResultCode *string `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// example:
	//
	// success
	ResultMessage *string `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMcubeHotpatchTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeHotpatchTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMcubeHotpatchTaskDetailResponseBody) GetQueryHotpatchTaskDetailResult() *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult {
	return s.QueryHotpatchTaskDetailResult
}

func (s *QueryMcubeHotpatchTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMcubeHotpatchTaskDetailResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryMcubeHotpatchTaskDetailResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryMcubeHotpatchTaskDetailResponseBody) SetQueryHotpatchTaskDetailResult(v *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult) *QueryMcubeHotpatchTaskDetailResponseBody {
	s.QueryHotpatchTaskDetailResult = v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBody) SetRequestId(v string) *QueryMcubeHotpatchTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBody) SetResultCode(v string) *QueryMcubeHotpatchTaskDetailResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBody) SetResultMessage(v string) *QueryMcubeHotpatchTaskDetailResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBody) Validate() error {
	if s.QueryHotpatchTaskDetailResult != nil {
		if err := s.QueryHotpatchTaskDetailResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult struct {
	// example:
	//
	// Success
	ErrorCode          *string                                                                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HotpatchTaskDetail *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail `json:"HotpatchTaskDetail,omitempty" xml:"HotpatchTaskDetail,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 61B9F63C-4E6F-5D8E-A694-899450987B48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	ResultMsg *string `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult) GoString() string {
	return s.String()
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult) GetHotpatchTaskDetail() *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	return s.HotpatchTaskDetail
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult) SetErrorCode(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult {
	s.ErrorCode = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult) SetHotpatchTaskDetail(v *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult {
	s.HotpatchTaskDetail = v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult) SetRequestId(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult {
	s.RequestId = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult) SetResultMsg(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult {
	s.ResultMsg = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult) SetSuccess(v bool) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult {
	s.Success = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResult) Validate() error {
	if s.HotpatchTaskDetail != nil {
		if err := s.HotpatchTaskDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail struct {
	// example:
	//
	// ONEXPRE22BA951112038-default
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// example:
	//
	// ALIPUB9A63274111812
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 0
	BaseInfoId *int64    `json:"BaseInfoId,omitempty" xml:"BaseInfoId,omitempty"`
	Bundles    []*string `json:"Bundles,omitempty" xml:"Bundles,omitempty" type:"Repeated"`
	// example:
	//
	// ***
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// https://xxxxx
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// example:
	//
	// 117
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// 1766111313000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2026-01-09 10:14:46
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 2026-01-09 10:14:46
	GmtModifiedStr *string `json:"GmtModifiedStr,omitempty" xml:"GmtModifiedStr,omitempty"`
	GreyConfigInfo *string `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	// example:
	//
	// 2024-01-01 12:00:00
	GreyEndtimeData *string `json:"GreyEndtimeData,omitempty" xml:"GreyEndtimeData,omitempty"`
	// example:
	//
	// 100
	GreyNum *int64 `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	// example:
	//
	// 14332
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 15cf3772223630be907c7aaefe8d51c6
	Md5  *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Memo *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	// example:
	//
	// ***
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// example:
	//
	// 1664552
	PackageId *int64 `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// example:
	//
	// iOS
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// ONEXPRE22BA951112038_ANDROID-default
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// example:
	//
	// 1.0.0
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	// example:
	//
	// 1
	PublishMode *int64 `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	// example:
	//
	// 2
	PublishPeriod *int64 `json:"PublishPeriod,omitempty" xml:"PublishPeriod,omitempty"`
	// example:
	//
	// 3
	PublishType *int64 `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	// example:
	//
	// 0
	QuickRollback *int64 `json:"QuickRollback,omitempty" xml:"QuickRollback,omitempty"`
	// example:
	//
	// 81c90a2cafdc6dfc54201e70845b5708
	ReleaseVersion *string                                                                                                `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	RuleJsonList   []*QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList `json:"RuleJsonList,omitempty" xml:"RuleJsonList,omitempty" type:"Repeated"`
	// example:
	//
	// mpaas.jar
	SourceName *string `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	// example:
	//
	// 1
	TaskStatus *int64 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 1770442895017
	TaskVersion *int64                                                                                              `json:"TaskVersion,omitempty" xml:"TaskVersion,omitempty"`
	Whitelist   []*QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
	// example:
	//
	// 825827
	WhitelistIds *string `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GoString() string {
	return s.String()
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetAppCode() *string {
	return s.AppCode
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetAppId() *string {
	return s.AppId
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetBaseInfoId() *int64 {
	return s.BaseInfoId
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetBundles() []*string {
	return s.Bundles
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetCreator() *string {
	return s.Creator
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetFileSize() *string {
	return s.FileSize
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetGmtModifiedStr() *string {
	return s.GmtModifiedStr
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetGreyEndtimeData() *string {
	return s.GreyEndtimeData
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetGreyNum() *int64 {
	return s.GreyNum
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetId() *int64 {
	return s.Id
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetMd5() *string {
	return s.Md5
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetMemo() *string {
	return s.Memo
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetModifier() *string {
	return s.Modifier
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetPackageId() *int64 {
	return s.PackageId
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetPlatform() *string {
	return s.Platform
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetProductId() *string {
	return s.ProductId
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetPublishMode() *int64 {
	return s.PublishMode
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetPublishPeriod() *int64 {
	return s.PublishPeriod
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetPublishType() *int64 {
	return s.PublishType
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetQuickRollback() *int64 {
	return s.QuickRollback
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetRuleJsonList() []*QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList {
	return s.RuleJsonList
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetSourceName() *string {
	return s.SourceName
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetTaskStatus() *int64 {
	return s.TaskStatus
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetTaskVersion() *int64 {
	return s.TaskVersion
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetWhitelist() []*QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist {
	return s.Whitelist
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetAppCode(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.AppCode = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetAppId(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.AppId = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetBaseInfoId(v int64) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.BaseInfoId = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetBundles(v []*string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.Bundles = v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetCreator(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.Creator = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetDownloadUrl(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.DownloadUrl = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetFileSize(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.FileSize = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetGmtCreate(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.GmtCreate = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetGmtModified(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.GmtModified = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetGmtModifiedStr(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.GmtModifiedStr = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetGreyConfigInfo(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.GreyConfigInfo = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetGreyEndtimeData(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.GreyEndtimeData = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetGreyNum(v int64) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.GreyNum = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetId(v int64) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.Id = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetMd5(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.Md5 = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetMemo(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.Memo = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetModifier(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.Modifier = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetPackageId(v int64) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.PackageId = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetPlatform(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.Platform = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetProductId(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.ProductId = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetProductVersion(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.ProductVersion = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetPublishMode(v int64) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.PublishMode = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetPublishPeriod(v int64) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.PublishPeriod = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetPublishType(v int64) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.PublishType = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetQuickRollback(v int64) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.QuickRollback = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetReleaseVersion(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.ReleaseVersion = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetRuleJsonList(v []*QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.RuleJsonList = v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetSourceName(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.SourceName = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetTaskStatus(v int64) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.TaskStatus = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetTaskVersion(v int64) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.TaskVersion = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetWhitelist(v []*QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.Whitelist = v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetWhitelistIds(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.WhitelistIds = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) SetWorkspaceId(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail) Validate() error {
	if s.RuleJsonList != nil {
		for _, item := range s.RuleJsonList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Whitelist != nil {
		for _, item := range s.Whitelist {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList struct {
	// example:
	//
	// and
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// city
	RuleElement *string `json:"RuleElement,omitempty" xml:"RuleElement,omitempty"`
	// example:
	//
	// 0
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// example:
	//
	// smtp.qiye.aliyun.com
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList) GoString() string {
	return s.String()
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList) GetOperation() *string {
	return s.Operation
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList) GetRuleElement() *string {
	return s.RuleElement
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList) GetRuleType() *string {
	return s.RuleType
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList) GetValue() *string {
	return s.Value
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList) SetOperation(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList {
	s.Operation = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList) SetRuleElement(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList {
	s.RuleElement = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList) SetRuleType(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList {
	s.RuleType = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList) SetValue(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList {
	s.Value = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList) Validate() error {
	return dara.Validate(s)
}

type QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist struct {
	// example:
	//
	// ONEXPRE40DB571151920-default
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// example:
	//
	// business
	Business *string `json:"Business,omitempty" xml:"Business,omitempty"`
	// example:
	//
	// 1760754049000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 9952804
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// userId
	IdType *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	// example:
	//
	// iOS
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 100
	WhiteListCount *int64 `json:"WhiteListCount,omitempty" xml:"WhiteListCount,omitempty"`
	// example:
	//
	// whitelistName
	WhiteListName *string `json:"WhiteListName,omitempty" xml:"WhiteListName,omitempty"`
}

func (s QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) GoString() string {
	return s.String()
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) GetAppCode() *string {
	return s.AppCode
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) GetBusiness() *string {
	return s.Business
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) GetId() *int64 {
	return s.Id
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) GetIdType() *string {
	return s.IdType
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) GetPlatform() *string {
	return s.Platform
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) GetStatus() *int64 {
	return s.Status
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) GetWhiteListCount() *int64 {
	return s.WhiteListCount
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) GetWhiteListName() *string {
	return s.WhiteListName
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) SetAppCode(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist {
	s.AppCode = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) SetBusiness(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist {
	s.Business = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) SetGmtModified(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist {
	s.GmtModified = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) SetId(v int64) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist {
	s.Id = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) SetIdType(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist {
	s.IdType = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) SetPlatform(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist {
	s.Platform = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) SetStatus(v int64) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist {
	s.Status = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) SetWhiteListCount(v int64) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist {
	s.WhiteListCount = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) SetWhiteListName(v string) *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist {
	s.WhiteListName = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist) Validate() error {
	return dara.Validate(s)
}
