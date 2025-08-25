// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcubeUpgradeTaskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGetTaskResult(v *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult) *GetMcubeUpgradeTaskInfoResponseBody
	GetGetTaskResult() *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult
	SetRequestId(v string) *GetMcubeUpgradeTaskInfoResponseBody
	GetRequestId() *string
	SetResultCode(v string) *GetMcubeUpgradeTaskInfoResponseBody
	GetResultCode() *string
	SetResultMessage(v string) *GetMcubeUpgradeTaskInfoResponseBody
	GetResultMessage() *string
}

type GetMcubeUpgradeTaskInfoResponseBody struct {
	GetTaskResult *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult `json:"GetTaskResult,omitempty" xml:"GetTaskResult,omitempty" type:"Struct"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                           `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage *string                                           `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s GetMcubeUpgradeTaskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeUpgradeTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcubeUpgradeTaskInfoResponseBody) GetGetTaskResult() *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult {
	return s.GetTaskResult
}

func (s *GetMcubeUpgradeTaskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcubeUpgradeTaskInfoResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *GetMcubeUpgradeTaskInfoResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *GetMcubeUpgradeTaskInfoResponseBody) SetGetTaskResult(v *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult) *GetMcubeUpgradeTaskInfoResponseBody {
	s.GetTaskResult = v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBody) SetRequestId(v string) *GetMcubeUpgradeTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBody) SetResultCode(v string) *GetMcubeUpgradeTaskInfoResponseBody {
	s.ResultCode = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBody) SetResultMessage(v string) *GetMcubeUpgradeTaskInfoResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult struct {
	ErrorCode *string                                                   `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string                                                   `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskInfo  *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo `json:"TaskInfo,omitempty" xml:"TaskInfo,omitempty" type:"Struct"`
}

func (s GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult) GoString() string {
	return s.String()
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult) GetTaskInfo() *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	return s.TaskInfo
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult) SetErrorCode(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult {
	s.ErrorCode = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult) SetRequestId(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult {
	s.RequestId = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult) SetResultMsg(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult {
	s.ResultMsg = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult) SetSuccess(v bool) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult {
	s.Success = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult) SetTaskInfo(v *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult {
	s.TaskInfo = v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResult) Validate() error {
	return dara.Validate(s)
}

type GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo struct {
	AppCode          *string                                                                 `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppId            *string                                                                 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppstoreUrl      *string                                                                 `json:"AppstoreUrl,omitempty" xml:"AppstoreUrl,omitempty"`
	Creator          *string                                                                 `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DownloadUrl      *string                                                                 `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	GreyConfigInfo   *string                                                                 `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	GreyEndtimeData  *string                                                                 `json:"GreyEndtimeData,omitempty" xml:"GreyEndtimeData,omitempty"`
	GreyNum          *int32                                                                  `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	HistoryForce     *int32                                                                  `json:"HistoryForce,omitempty" xml:"HistoryForce,omitempty"`
	Id               *int64                                                                  `json:"Id,omitempty" xml:"Id,omitempty"`
	IsEnterprise     *int32                                                                  `json:"IsEnterprise,omitempty" xml:"IsEnterprise,omitempty"`
	Memo             *string                                                                 `json:"Memo,omitempty" xml:"Memo,omitempty"`
	Modifier         *string                                                                 `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	PackageInfoId    *int64                                                                  `json:"PackageInfoId,omitempty" xml:"PackageInfoId,omitempty"`
	PackageType      *string                                                                 `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	Platform         *string                                                                 `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ProductId        *string                                                                 `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	PublishMode      *int32                                                                  `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	PublishType      *int32                                                                  `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	RuleJsonList     []*GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList `json:"RuleJsonList,omitempty" xml:"RuleJsonList,omitempty" type:"Repeated"`
	SilentType       *int32                                                                  `json:"SilentType,omitempty" xml:"SilentType,omitempty"`
	TaskStatus       *int32                                                                  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	UpgradeContent   *string                                                                 `json:"UpgradeContent,omitempty" xml:"UpgradeContent,omitempty"`
	UpgradeType      *int32                                                                  `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
	UpgradeValidTime *int32                                                                  `json:"UpgradeValidTime,omitempty" xml:"UpgradeValidTime,omitempty"`
	Whitelist        []*GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist    `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
	WhitelistIds     *string                                                                 `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
	WorkspaceId      *string                                                                 `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GoString() string {
	return s.String()
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetAppCode() *string {
	return s.AppCode
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetAppId() *string {
	return s.AppId
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetAppstoreUrl() *string {
	return s.AppstoreUrl
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetCreator() *string {
	return s.Creator
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetGreyEndtimeData() *string {
	return s.GreyEndtimeData
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetGreyNum() *int32 {
	return s.GreyNum
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetHistoryForce() *int32 {
	return s.HistoryForce
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetId() *int64 {
	return s.Id
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetIsEnterprise() *int32 {
	return s.IsEnterprise
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetMemo() *string {
	return s.Memo
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetModifier() *string {
	return s.Modifier
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetPackageInfoId() *int64 {
	return s.PackageInfoId
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetPackageType() *string {
	return s.PackageType
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetPlatform() *string {
	return s.Platform
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetProductId() *string {
	return s.ProductId
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetPublishMode() *int32 {
	return s.PublishMode
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetPublishType() *int32 {
	return s.PublishType
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetRuleJsonList() []*GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList {
	return s.RuleJsonList
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetSilentType() *int32 {
	return s.SilentType
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetUpgradeContent() *string {
	return s.UpgradeContent
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetUpgradeType() *int32 {
	return s.UpgradeType
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetUpgradeValidTime() *int32 {
	return s.UpgradeValidTime
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetWhitelist() []*GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist {
	return s.Whitelist
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetAppCode(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.AppCode = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetAppId(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.AppId = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetAppstoreUrl(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.AppstoreUrl = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetCreator(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.Creator = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetDownloadUrl(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.DownloadUrl = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetGreyConfigInfo(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.GreyConfigInfo = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetGreyEndtimeData(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.GreyEndtimeData = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetGreyNum(v int32) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.GreyNum = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetHistoryForce(v int32) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.HistoryForce = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetId(v int64) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.Id = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetIsEnterprise(v int32) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.IsEnterprise = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetMemo(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.Memo = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetModifier(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.Modifier = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetPackageInfoId(v int64) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.PackageInfoId = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetPackageType(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.PackageType = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetPlatform(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.Platform = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetProductId(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.ProductId = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetPublishMode(v int32) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.PublishMode = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetPublishType(v int32) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.PublishType = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetRuleJsonList(v []*GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.RuleJsonList = v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetSilentType(v int32) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.SilentType = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetTaskStatus(v int32) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.TaskStatus = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetUpgradeContent(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.UpgradeContent = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetUpgradeType(v int32) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.UpgradeType = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetUpgradeValidTime(v int32) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.UpgradeValidTime = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetWhitelist(v []*GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.Whitelist = v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetWhitelistIds(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.WhitelistIds = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) SetWorkspaceId(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo {
	s.WorkspaceId = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfo) Validate() error {
	return dara.Validate(s)
}

type GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList struct {
	Operation   *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	RuleElement *string `json:"RuleElement,omitempty" xml:"RuleElement,omitempty"`
	RuleType    *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList) GoString() string {
	return s.String()
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList) GetOperation() *string {
	return s.Operation
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList) GetRuleElement() *string {
	return s.RuleElement
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList) GetRuleType() *string {
	return s.RuleType
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList) GetValue() *string {
	return s.Value
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList) SetOperation(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList {
	s.Operation = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList) SetRuleElement(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList {
	s.RuleElement = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList) SetRuleType(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList {
	s.RuleType = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList) SetValue(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList {
	s.Value = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoRuleJsonList) Validate() error {
	return dara.Validate(s)
}

type GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist struct {
	AppCode        *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	Platform       *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UserType       *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
	WhiteListCount *int64  `json:"WhiteListCount,omitempty" xml:"WhiteListCount,omitempty"`
	WhiteListName  *string `json:"WhiteListName,omitempty" xml:"WhiteListName,omitempty"`
	WhitelistType  *string `json:"WhitelistType,omitempty" xml:"WhitelistType,omitempty"`
}

func (s GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) GoString() string {
	return s.String()
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) GetAppCode() *string {
	return s.AppCode
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) GetId() *int64 {
	return s.Id
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) GetIdType() *string {
	return s.IdType
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) GetPlatform() *string {
	return s.Platform
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) GetStatus() *int32 {
	return s.Status
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) GetUserType() *string {
	return s.UserType
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) GetWhiteListCount() *int64 {
	return s.WhiteListCount
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) GetWhiteListName() *string {
	return s.WhiteListName
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) GetWhitelistType() *string {
	return s.WhitelistType
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) SetAppCode(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist {
	s.AppCode = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) SetId(v int64) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist {
	s.Id = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) SetIdType(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist {
	s.IdType = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) SetPlatform(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist {
	s.Platform = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) SetStatus(v int32) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist {
	s.Status = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) SetUserType(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist {
	s.UserType = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) SetWhiteListCount(v int64) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist {
	s.WhiteListCount = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) SetWhiteListName(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist {
	s.WhiteListName = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) SetWhitelistType(v string) *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist {
	s.WhitelistType = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoResponseBodyGetTaskResultTaskInfoWhitelist) Validate() error {
	return dara.Validate(s)
}
