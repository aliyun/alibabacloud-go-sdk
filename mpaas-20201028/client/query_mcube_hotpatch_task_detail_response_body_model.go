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
	RequestId                     *string                                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode                    *string                                                                `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultMessage                 *string                                                                `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
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
	ErrorCode          *string                                                                                  `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HotpatchTaskDetail *QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetail `json:"HotpatchTaskDetail,omitempty" xml:"HotpatchTaskDetail,omitempty" type:"Struct"`
	RequestId          *string                                                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg          *string                                                                                  `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success            *bool                                                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
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
	AppCode         *string                                                                                                `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppId           *string                                                                                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BaseInfoId      *int64                                                                                                 `json:"BaseInfoId,omitempty" xml:"BaseInfoId,omitempty"`
	Bundles         []*string                                                                                              `json:"Bundles,omitempty" xml:"Bundles,omitempty" type:"Repeated"`
	Creator         *string                                                                                                `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DownloadUrl     *string                                                                                                `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	FileSize        *string                                                                                                `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	GmtCreate       *string                                                                                                `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string                                                                                                `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GmtModifiedStr  *string                                                                                                `json:"GmtModifiedStr,omitempty" xml:"GmtModifiedStr,omitempty"`
	GreyConfigInfo  *string                                                                                                `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	GreyEndtimeData *string                                                                                                `json:"GreyEndtimeData,omitempty" xml:"GreyEndtimeData,omitempty"`
	GreyNum         *int64                                                                                                 `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	Id              *int64                                                                                                 `json:"Id,omitempty" xml:"Id,omitempty"`
	Md5             *string                                                                                                `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Memo            *string                                                                                                `json:"Memo,omitempty" xml:"Memo,omitempty"`
	Modifier        *string                                                                                                `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	PackageId       *int64                                                                                                 `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	Platform        *string                                                                                                `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ProductId       *string                                                                                                `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductVersion  *string                                                                                                `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	PublishMode     *int64                                                                                                 `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	PublishPeriod   *int64                                                                                                 `json:"PublishPeriod,omitempty" xml:"PublishPeriod,omitempty"`
	PublishType     *int64                                                                                                 `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	QuickRollback   *int64                                                                                                 `json:"QuickRollback,omitempty" xml:"QuickRollback,omitempty"`
	ReleaseVersion  *string                                                                                                `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	RuleJsonList    []*QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailRuleJsonList `json:"RuleJsonList,omitempty" xml:"RuleJsonList,omitempty" type:"Repeated"`
	SourceName      *string                                                                                                `json:"SourceName,omitempty" xml:"SourceName,omitempty"`
	TaskStatus      *int64                                                                                                 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskVersion     *int64                                                                                                 `json:"TaskVersion,omitempty" xml:"TaskVersion,omitempty"`
	Whitelist       []*QueryMcubeHotpatchTaskDetailResponseBodyQueryHotpatchTaskDetailResultHotpatchTaskDetailWhitelist    `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
	WhitelistIds    *string                                                                                                `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
	WorkspaceId     *string                                                                                                `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
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
	Operation   *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	RuleElement *string `json:"RuleElement,omitempty" xml:"RuleElement,omitempty"`
	RuleType    *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
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
	AppCode        *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	Business       *string `json:"Business,omitempty" xml:"Business,omitempty"`
	GmtModified    *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id             *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	IdType         *string `json:"IdType,omitempty" xml:"IdType,omitempty"`
	Platform       *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Status         *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	WhiteListCount *int64  `json:"WhiteListCount,omitempty" xml:"WhiteListCount,omitempty"`
	WhiteListName  *string `json:"WhiteListName,omitempty" xml:"WhiteListName,omitempty"`
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
