// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcubeNebulaTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGetMcubeNebulaTaskDetailResult(v *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult) *GetMcubeNebulaTaskDetailResponseBody
	GetGetMcubeNebulaTaskDetailResult() *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult
	SetRequestId(v string) *GetMcubeNebulaTaskDetailResponseBody
	GetRequestId() *string
	SetResultCode(v string) *GetMcubeNebulaTaskDetailResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *GetMcubeNebulaTaskDetailResponseBody
	GetResultMessage() *string
}

type GetMcubeNebulaTaskDetailResponseBody struct {
	GetMcubeNebulaTaskDetailResult *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult `json:"GetMcubeNebulaTaskDetailResult,omitempty" xml:"GetMcubeNebulaTaskDetailResult,omitempty" type:"Struct"`
	RequestId                      *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                     *string                                                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage                  *string                                                             `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetMcubeNebulaTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeNebulaTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcubeNebulaTaskDetailResponseBody) GetGetMcubeNebulaTaskDetailResult() *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult {
	return s.GetMcubeNebulaTaskDetailResult
}

func (s *GetMcubeNebulaTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcubeNebulaTaskDetailResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetMcubeNebulaTaskDetailResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *GetMcubeNebulaTaskDetailResponseBody) SetGetMcubeNebulaTaskDetailResult(v *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult) *GetMcubeNebulaTaskDetailResponseBody {
	s.GetMcubeNebulaTaskDetailResult = v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBody) SetRequestId(v string) *GetMcubeNebulaTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBody) SetResultCode(v string) *GetMcubeNebulaTaskDetailResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBody) SetResultMessage(v string) *GetMcubeNebulaTaskDetailResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult struct {
	ErrorCode        *string                                                                             `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	NebulaTaskDetail *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail `json:"NebulaTaskDetail,omitempty" xml:"NebulaTaskDetail,omitempty" type:"Struct"`
	RequestId        *string                                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg        *string                                                                             `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success          *bool                                                                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult) GoString() string {
	return s.String()
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult) GetNebulaTaskDetail() *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	return s.NebulaTaskDetail
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult) SetErrorCode(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult {
	s.ErrorCode = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult) SetNebulaTaskDetail(v *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult {
	s.NebulaTaskDetail = v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult) SetRequestId(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult {
	s.RequestId = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult) SetResultMsg(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult {
	s.ResultMsg = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult) SetSuccess(v bool) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult {
	s.Success = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResult) Validate() error {
	return dara.Validate(s)
}

type GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail struct {
	AppCode          *string                                                                                           `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppId            *string                                                                                           `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Atomic           *int32                                                                                            `json:"Atomic,omitempty" xml:"Atomic,omitempty"`
	BaseInfoId       *int64                                                                                            `json:"BaseInfoId,omitempty" xml:"BaseInfoId,omitempty"`
	BizType          *string                                                                                           `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Creator          *string                                                                                           `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Cronexpress      *int32                                                                                            `json:"Cronexpress,omitempty" xml:"Cronexpress,omitempty"`
	DownloadUrl      *string                                                                                           `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ExtraData        *string                                                                                           `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	FileSize         *string                                                                                           `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	FullRepair       *int32                                                                                            `json:"FullRepair,omitempty" xml:"FullRepair,omitempty"`
	GmtCreate        *string                                                                                           `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *string                                                                                           `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GmtModifiedStr   *string                                                                                           `json:"GmtModifiedStr,omitempty" xml:"GmtModifiedStr,omitempty"`
	GreyConfigInfo   *string                                                                                           `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	GreyEndtime      *string                                                                                           `json:"GreyEndtime,omitempty" xml:"GreyEndtime,omitempty"`
	GreyEndtimeData  *string                                                                                           `json:"GreyEndtimeData,omitempty" xml:"GreyEndtimeData,omitempty"`
	GreyEndtimeStr   *string                                                                                           `json:"GreyEndtimeStr,omitempty" xml:"GreyEndtimeStr,omitempty"`
	GreyNum          *int32                                                                                            `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	GreyUrl          *string                                                                                           `json:"GreyUrl,omitempty" xml:"GreyUrl,omitempty"`
	Id               *int64                                                                                            `json:"Id,omitempty" xml:"Id,omitempty"`
	IssueDesc        *string                                                                                           `json:"IssueDesc,omitempty" xml:"IssueDesc,omitempty"`
	Memo             *string                                                                                           `json:"Memo,omitempty" xml:"Memo,omitempty"`
	Modifier         *string                                                                                           `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	OssPath          *string                                                                                           `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	PackageId        *int64                                                                                            `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	Percent          *int32                                                                                            `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Platform         *string                                                                                           `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ProductId        *string                                                                                           `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductVersion   *string                                                                                           `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	PublishMode      *int32                                                                                            `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	PublishPeriod    *int32                                                                                            `json:"PublishPeriod,omitempty" xml:"PublishPeriod,omitempty"`
	PublishType      *int32                                                                                            `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	QuickRollback    *int32                                                                                            `json:"QuickRollback,omitempty" xml:"QuickRollback,omitempty"`
	ReleaseVersion   *string                                                                                           `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	RuleJsonList     []*GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList `json:"RuleJsonList,omitempty" xml:"RuleJsonList,omitempty" type:"Repeated"`
	SourceId         *string                                                                                           `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceName       *string                                                                                           `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	SourceType       *string                                                                                           `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Status           *int32                                                                                            `json:"Status,omitempty" xml:"Status,omitempty"`
	SyncResult       *string                                                                                           `json:"SyncResult,omitempty" xml:"SyncResult,omitempty"`
	SyncType         *int32                                                                                            `json:"SyncType,omitempty" xml:"SyncType,omitempty"`
	TaskName         *string                                                                                           `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskStatus       *int32                                                                                            `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskType         *int32                                                                                            `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskVersion      *int64                                                                                            `json:"TaskVersion,omitempty" xml:"TaskVersion,omitempty"`
	UpgradeNoticeNum *int64                                                                                            `json:"UpgradeNoticeNum,omitempty" xml:"UpgradeNoticeNum,omitempty"`
	UpgradeProgress  *string                                                                                           `json:"UpgradeProgress,omitempty" xml:"UpgradeProgress,omitempty"`
	WhitelistIds     *string                                                                                           `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
	WorkspaceId      *string                                                                                           `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GoString() string {
	return s.String()
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetAppCode() *string {
	return s.AppCode
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetAppId() *string {
	return s.AppId
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetAtomic() *int32 {
	return s.Atomic
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetBaseInfoId() *int64 {
	return s.BaseInfoId
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetBizType() *string {
	return s.BizType
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetCreator() *string {
	return s.Creator
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetCronexpress() *int32 {
	return s.Cronexpress
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetExtraData() *string {
	return s.ExtraData
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetFileSize() *string {
	return s.FileSize
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetFullRepair() *int32 {
	return s.FullRepair
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetGmtModifiedStr() *string {
	return s.GmtModifiedStr
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetGreyEndtime() *string {
	return s.GreyEndtime
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetGreyEndtimeData() *string {
	return s.GreyEndtimeData
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetGreyEndtimeStr() *string {
	return s.GreyEndtimeStr
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetGreyNum() *int32 {
	return s.GreyNum
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetGreyUrl() *string {
	return s.GreyUrl
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetId() *int64 {
	return s.Id
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetIssueDesc() *string {
	return s.IssueDesc
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetMemo() *string {
	return s.Memo
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetModifier() *string {
	return s.Modifier
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetOssPath() *string {
	return s.OssPath
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetPackageId() *int64 {
	return s.PackageId
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetPercent() *int32 {
	return s.Percent
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetPlatform() *string {
	return s.Platform
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetProductId() *string {
	return s.ProductId
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetPublishMode() *int32 {
	return s.PublishMode
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetPublishPeriod() *int32 {
	return s.PublishPeriod
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetPublishType() *int32 {
	return s.PublishType
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetQuickRollback() *int32 {
	return s.QuickRollback
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetRuleJsonList() []*GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList {
	return s.RuleJsonList
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetSourceId() *string {
	return s.SourceId
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetSourceName() *string {
	return s.SourceName
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetSourceType() *string {
	return s.SourceType
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetStatus() *int32 {
	return s.Status
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetSyncResult() *string {
	return s.SyncResult
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetSyncType() *int32 {
	return s.SyncType
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetTaskName() *string {
	return s.TaskName
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetTaskType() *int32 {
	return s.TaskType
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetTaskVersion() *int64 {
	return s.TaskVersion
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetUpgradeNoticeNum() *int64 {
	return s.UpgradeNoticeNum
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetUpgradeProgress() *string {
	return s.UpgradeProgress
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetAppCode(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.AppCode = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetAppId(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.AppId = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetAtomic(v int32) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.Atomic = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetBaseInfoId(v int64) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.BaseInfoId = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetBizType(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.BizType = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetCreator(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.Creator = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetCronexpress(v int32) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.Cronexpress = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetDownloadUrl(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.DownloadUrl = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetExtraData(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.ExtraData = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetFileSize(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.FileSize = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetFullRepair(v int32) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.FullRepair = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetGmtCreate(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.GmtCreate = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetGmtModified(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.GmtModified = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetGmtModifiedStr(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.GmtModifiedStr = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetGreyConfigInfo(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.GreyConfigInfo = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetGreyEndtime(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.GreyEndtime = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetGreyEndtimeData(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.GreyEndtimeData = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetGreyEndtimeStr(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.GreyEndtimeStr = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetGreyNum(v int32) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.GreyNum = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetGreyUrl(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.GreyUrl = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetId(v int64) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.Id = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetIssueDesc(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.IssueDesc = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetMemo(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.Memo = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetModifier(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.Modifier = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetOssPath(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.OssPath = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetPackageId(v int64) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.PackageId = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetPercent(v int32) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.Percent = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetPlatform(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.Platform = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetProductId(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.ProductId = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetProductVersion(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.ProductVersion = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetPublishMode(v int32) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.PublishMode = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetPublishPeriod(v int32) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.PublishPeriod = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetPublishType(v int32) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.PublishType = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetQuickRollback(v int32) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.QuickRollback = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetReleaseVersion(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.ReleaseVersion = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetRuleJsonList(v []*GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.RuleJsonList = v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetSourceId(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.SourceId = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetSourceName(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.SourceName = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetSourceType(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.SourceType = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetStatus(v int32) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.Status = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetSyncResult(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.SyncResult = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetSyncType(v int32) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.SyncType = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetTaskName(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.TaskName = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetTaskStatus(v int32) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.TaskStatus = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetTaskType(v int32) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.TaskType = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetTaskVersion(v int64) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.TaskVersion = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetUpgradeNoticeNum(v int64) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.UpgradeNoticeNum = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetUpgradeProgress(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.UpgradeProgress = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetWhitelistIds(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.WhitelistIds = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) SetWorkspaceId(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail {
	s.WorkspaceId = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetail) Validate() error {
	return dara.Validate(s)
}

type GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList struct {
	Operation   *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	RuleElement *string `json:"RuleElement,omitempty" xml:"RuleElement,omitempty"`
	RuleType    *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList) GoString() string {
	return s.String()
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList) GetOperation() *string {
	return s.Operation
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList) GetRuleElement() *string {
	return s.RuleElement
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList) GetRuleType() *string {
	return s.RuleType
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList) GetValue() *string {
	return s.Value
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList) SetOperation(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList {
	s.Operation = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList) SetRuleElement(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList {
	s.RuleElement = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList) SetRuleType(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList {
	s.RuleType = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList) SetValue(v string) *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList {
	s.Value = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailResponseBodyGetMcubeNebulaTaskDetailResultNebulaTaskDetailRuleJsonList) Validate() error {
	return dara.Validate(s)
}
