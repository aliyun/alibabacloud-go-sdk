// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeNebulaTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListMcubeNebulaTaskResult(v *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult) *ListMcubeNebulaTasksResponseBody
	GetListMcubeNebulaTaskResult() *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult
	SetRequestId(v string) *ListMcubeNebulaTasksResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMcubeNebulaTasksResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListMcubeNebulaTasksResponseBody
	GetResultMessage() *string
}

type ListMcubeNebulaTasksResponseBody struct {
	ListMcubeNebulaTaskResult *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult `json:"ListMcubeNebulaTaskResult,omitempty" xml:"ListMcubeNebulaTaskResult,omitempty" type:"Struct"`
	RequestId                 *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                *string                                                    `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage             *string                                                    `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMcubeNebulaTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeNebulaTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcubeNebulaTasksResponseBody) GetListMcubeNebulaTaskResult() *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult {
	return s.ListMcubeNebulaTaskResult
}

func (s *ListMcubeNebulaTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeNebulaTasksResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMcubeNebulaTasksResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMcubeNebulaTasksResponseBody) SetListMcubeNebulaTaskResult(v *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult) *ListMcubeNebulaTasksResponseBody {
	s.ListMcubeNebulaTaskResult = v
	return s
}

func (s *ListMcubeNebulaTasksResponseBody) SetRequestId(v string) *ListMcubeNebulaTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBody) SetResultCode(v string) *ListMcubeNebulaTasksResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBody) SetResultMessage(v string) *ListMcubeNebulaTasksResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult struct {
	ErrorCode      *string                                                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	NebulaTaskInfo []*ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo `json:"NebulaTaskInfo,omitempty" xml:"NebulaTaskInfo,omitempty" type:"Repeated"`
	RequestId      *string                                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg      *string                                                                    `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success        *bool                                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult) GoString() string {
	return s.String()
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult) GetNebulaTaskInfo() []*ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	return s.NebulaTaskInfo
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult) SetErrorCode(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult {
	s.ErrorCode = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult) SetNebulaTaskInfo(v []*ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult {
	s.NebulaTaskInfo = v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult) SetRequestId(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult {
	s.RequestId = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult) SetResultMsg(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult {
	s.ResultMsg = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult) SetSuccess(v bool) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult {
	s.Success = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResult) Validate() error {
	return dara.Validate(s)
}

type ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo struct {
	AppCode          *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	BizType          *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	Creator          *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	GmtCreate        *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GmtModifiedStr   *string `json:"GmtModifiedStr,omitempty" xml:"GmtModifiedStr,omitempty"`
	GreyConfigInfo   *string `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	GreyEndtime      *string `json:"GreyEndtime,omitempty" xml:"GreyEndtime,omitempty"`
	GreyEndtimeData  *string `json:"GreyEndtimeData,omitempty" xml:"GreyEndtimeData,omitempty"`
	GreyEndtimeStr   *string `json:"GreyEndtimeStr,omitempty" xml:"GreyEndtimeStr,omitempty"`
	GreyNum          *int32  `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	GreyUrl          *string `json:"GreyUrl,omitempty" xml:"GreyUrl,omitempty"`
	Id               *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Memo             *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	Modifier         *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	PackageId        *int64  `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	Percent          *int32  `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Platform         *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ProductId        *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductVersion   *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	PublishMode      *int32  `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	PublishType      *int32  `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	ReleaseVersion   *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	Status           *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	SyncResult       *string `json:"SyncResult,omitempty" xml:"SyncResult,omitempty"`
	TaskName         *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskStatus       *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskType         *int32  `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	TaskVersion      *int64  `json:"TaskVersion,omitempty" xml:"TaskVersion,omitempty"`
	UpgradeNoticeNum *int64  `json:"UpgradeNoticeNum,omitempty" xml:"UpgradeNoticeNum,omitempty"`
	UpgradeProgress  *string `json:"UpgradeProgress,omitempty" xml:"UpgradeProgress,omitempty"`
	WhitelistIds     *string `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
}

func (s ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GoString() string {
	return s.String()
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetAppCode() *string {
	return s.AppCode
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetBizType() *string {
	return s.BizType
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetCreator() *string {
	return s.Creator
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetGmtModifiedStr() *string {
	return s.GmtModifiedStr
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetGreyEndtime() *string {
	return s.GreyEndtime
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetGreyEndtimeData() *string {
	return s.GreyEndtimeData
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetGreyEndtimeStr() *string {
	return s.GreyEndtimeStr
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetGreyNum() *int32 {
	return s.GreyNum
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetGreyUrl() *string {
	return s.GreyUrl
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetId() *int64 {
	return s.Id
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetMemo() *string {
	return s.Memo
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetModifier() *string {
	return s.Modifier
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetPackageId() *int64 {
	return s.PackageId
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetPercent() *int32 {
	return s.Percent
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetPlatform() *string {
	return s.Platform
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetProductId() *string {
	return s.ProductId
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetPublishMode() *int32 {
	return s.PublishMode
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetPublishType() *int32 {
	return s.PublishType
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetStatus() *int32 {
	return s.Status
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetSyncResult() *string {
	return s.SyncResult
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetTaskName() *string {
	return s.TaskName
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetTaskType() *int32 {
	return s.TaskType
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetTaskVersion() *int64 {
	return s.TaskVersion
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetUpgradeNoticeNum() *int64 {
	return s.UpgradeNoticeNum
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetUpgradeProgress() *string {
	return s.UpgradeProgress
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetAppCode(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.AppCode = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetBizType(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.BizType = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetCreator(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.Creator = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetGmtCreate(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.GmtCreate = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetGmtModified(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.GmtModified = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetGmtModifiedStr(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.GmtModifiedStr = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetGreyConfigInfo(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.GreyConfigInfo = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetGreyEndtime(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.GreyEndtime = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetGreyEndtimeData(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.GreyEndtimeData = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetGreyEndtimeStr(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.GreyEndtimeStr = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetGreyNum(v int32) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.GreyNum = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetGreyUrl(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.GreyUrl = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetId(v int64) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.Id = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetMemo(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.Memo = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetModifier(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.Modifier = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetPackageId(v int64) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.PackageId = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetPercent(v int32) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.Percent = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetPlatform(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.Platform = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetProductId(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.ProductId = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetProductVersion(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.ProductVersion = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetPublishMode(v int32) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.PublishMode = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetPublishType(v int32) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.PublishType = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetReleaseVersion(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.ReleaseVersion = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetStatus(v int32) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.Status = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetSyncResult(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.SyncResult = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetTaskName(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.TaskName = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetTaskStatus(v int32) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.TaskStatus = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetTaskType(v int32) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.TaskType = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetTaskVersion(v int64) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.TaskVersion = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetUpgradeNoticeNum(v int64) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.UpgradeNoticeNum = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetUpgradeProgress(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.UpgradeProgress = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) SetWhitelistIds(v string) *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo {
	s.WhitelistIds = &v
	return s
}

func (s *ListMcubeNebulaTasksResponseBodyListMcubeNebulaTaskResultNebulaTaskInfo) Validate() error {
	return dara.Validate(s)
}
