// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeUpgradeTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetListTaskResult(v *ListMcubeUpgradeTasksResponseBodyListTaskResult) *ListMcubeUpgradeTasksResponseBody
	GetListTaskResult() *ListMcubeUpgradeTasksResponseBodyListTaskResult
	SetRequestId(v string) *ListMcubeUpgradeTasksResponseBody
	GetRequestId() *string
	SetResultCode(v string) *ListMcubeUpgradeTasksResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *ListMcubeUpgradeTasksResponseBody
	GetResultMessage() *string
}

type ListMcubeUpgradeTasksResponseBody struct {
	ListTaskResult *ListMcubeUpgradeTasksResponseBodyListTaskResult `json:"ListTaskResult,omitempty" xml:"ListTaskResult,omitempty" type:"Struct"`
	RequestId      *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode     *string                                          `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage  *string                                          `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s ListMcubeUpgradeTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeUpgradeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListMcubeUpgradeTasksResponseBody) GetListTaskResult() *ListMcubeUpgradeTasksResponseBodyListTaskResult {
	return s.ListTaskResult
}

func (s *ListMcubeUpgradeTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeUpgradeTasksResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *ListMcubeUpgradeTasksResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *ListMcubeUpgradeTasksResponseBody) SetListTaskResult(v *ListMcubeUpgradeTasksResponseBodyListTaskResult) *ListMcubeUpgradeTasksResponseBody {
	s.ListTaskResult = v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBody) SetRequestId(v string) *ListMcubeUpgradeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBody) SetResultCode(v string) *ListMcubeUpgradeTasksResponseBody {
	s.ResultCode = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBody) SetResultMessage(v string) *ListMcubeUpgradeTasksResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMcubeUpgradeTasksResponseBodyListTaskResult struct {
	ErrorCode *string                                                    `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string                                                    `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool                                                      `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskInfo  []*ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Repeated"`
}

func (s ListMcubeUpgradeTasksResponseBodyListTaskResult) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeUpgradeTasksResponseBodyListTaskResult) GoString() string {
	return s.String()
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResult) GetSuccess() *bool {
	return s.Success
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResult) GetTaskInfo() []*ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	return s.TaskInfo
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResult) SetErrorCode(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResult {
	s.ErrorCode = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResult) SetRequestId(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResult {
	s.RequestId = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResult) SetResultMsg(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResult {
	s.ResultMsg = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResult) SetSuccess(v bool) *ListMcubeUpgradeTasksResponseBodyListTaskResult {
	s.Success = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResult) SetTaskInfo(v []*ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) *ListMcubeUpgradeTasksResponseBodyListTaskResult {
	s.TaskInfo = v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResult) Validate() error {
	return dara.Validate(s)
}

type ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo struct {
	AppCode        *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	Creator        *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	GmtCreate      *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified    *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GreyConfigInfo *string `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	GreyEndtime    *string `json:"GreyEndtime,omitempty" xml:"GreyEndtime,omitempty"`
	GreyNum        *int32  `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	HistoryForce   *int32  `json:"HistoryForce,omitempty" xml:"HistoryForce,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IsEnterprise   *int32  `json:"IsEnterprise,omitempty" xml:"IsEnterprise,omitempty"`
	Memo           *string `json:"Memo,omitempty" xml:"Memo,omitempty"`
	Modifier       *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	PackageInfoId  *int64  `json:"PackageInfoId,omitempty" xml:"PackageInfoId,omitempty"`
	Platform       *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ProductId      *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductVersion *string `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	PublishMode    *int32  `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	PublishType    *int32  `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	PushContent    *string `json:"PushContent,omitempty" xml:"PushContent,omitempty"`
	TaskStatus     *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	UpgradeContent *string `json:"UpgradeContent,omitempty" xml:"UpgradeContent,omitempty"`
	UpgradeType    *int32  `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
	WhitelistIds   *string `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
}

func (s ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GoString() string {
	return s.String()
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetAppCode() *string {
	return s.AppCode
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetCreator() *string {
	return s.Creator
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetGreyEndtime() *string {
	return s.GreyEndtime
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetGreyNum() *int32 {
	return s.GreyNum
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetHistoryForce() *int32 {
	return s.HistoryForce
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetId() *int64 {
	return s.Id
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetIsEnterprise() *int32 {
	return s.IsEnterprise
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetMemo() *string {
	return s.Memo
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetModifier() *string {
	return s.Modifier
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetPackageInfoId() *int64 {
	return s.PackageInfoId
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetPlatform() *string {
	return s.Platform
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetProductId() *string {
	return s.ProductId
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetPublishMode() *int32 {
	return s.PublishMode
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetPublishType() *int32 {
	return s.PublishType
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetPushContent() *string {
	return s.PushContent
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetUpgradeContent() *string {
	return s.UpgradeContent
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetUpgradeType() *int32 {
	return s.UpgradeType
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetAppCode(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.AppCode = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetCreator(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.Creator = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetGmtCreate(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.GmtCreate = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetGmtModified(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.GmtModified = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetGreyConfigInfo(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.GreyConfigInfo = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetGreyEndtime(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.GreyEndtime = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetGreyNum(v int32) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.GreyNum = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetHistoryForce(v int32) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.HistoryForce = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetId(v int64) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.Id = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetIsEnterprise(v int32) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.IsEnterprise = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetMemo(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.Memo = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetModifier(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.Modifier = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetPackageInfoId(v int64) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.PackageInfoId = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetPlatform(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.Platform = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetProductId(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.ProductId = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetProductVersion(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.ProductVersion = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetPublishMode(v int32) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.PublishMode = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetPublishType(v int32) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.PublishType = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetPushContent(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.PushContent = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetTaskStatus(v int32) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.TaskStatus = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetUpgradeContent(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.UpgradeContent = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetUpgradeType(v int32) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.UpgradeType = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) SetWhitelistIds(v string) *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo {
	s.WhitelistIds = &v
	return s
}

func (s *ListMcubeUpgradeTasksResponseBodyListTaskResultTaskInfo) Validate() error {
	return dara.Validate(s)
}
