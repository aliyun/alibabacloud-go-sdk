// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMdsUpgradeTaskDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryMdsUpgradeTaskDetailResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryMdsUpgradeTaskDetailResponseBody
	GetResultCode() *string
	SetResultContent(v *QueryMdsUpgradeTaskDetailResponseBodyResultContent) *QueryMdsUpgradeTaskDetailResponseBody
	GetResultContent() *QueryMdsUpgradeTaskDetailResponseBodyResultContent
	SetResultMessage(v string) *QueryMdsUpgradeTaskDetailResponseBody
	GetResultMessage() *string
}

type QueryMdsUpgradeTaskDetailResponseBody struct {
	RequestId     *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                             `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *QueryMdsUpgradeTaskDetailResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                             `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMdsUpgradeTaskDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMdsUpgradeTaskDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMdsUpgradeTaskDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMdsUpgradeTaskDetailResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryMdsUpgradeTaskDetailResponseBody) GetResultContent() *QueryMdsUpgradeTaskDetailResponseBodyResultContent {
	return s.ResultContent
}

func (s *QueryMdsUpgradeTaskDetailResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryMdsUpgradeTaskDetailResponseBody) SetRequestId(v string) *QueryMdsUpgradeTaskDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBody) SetResultCode(v string) *QueryMdsUpgradeTaskDetailResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBody) SetResultContent(v *QueryMdsUpgradeTaskDetailResponseBodyResultContent) *QueryMdsUpgradeTaskDetailResponseBody {
	s.ResultContent = v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBody) SetResultMessage(v string) *QueryMdsUpgradeTaskDetailResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMdsUpgradeTaskDetailResponseBodyResultContent struct {
	Data      *QueryMdsUpgradeTaskDetailResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMdsUpgradeTaskDetailResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s QueryMdsUpgradeTaskDetailResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContent) GetData() *QueryMdsUpgradeTaskDetailResponseBodyResultContentData {
	return s.Data
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContent) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContent) SetData(v *QueryMdsUpgradeTaskDetailResponseBodyResultContentData) *QueryMdsUpgradeTaskDetailResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContent) SetRequestId(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContent {
	s.RequestId = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}

type QueryMdsUpgradeTaskDetailResponseBodyResultContentData struct {
	Content   *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	ErrorCode *string                                                        `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	RequestId *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultMsg *string                                                        `json:"ResultMsg,omitempty" xml:"ResultMsg,omitempty"`
	Success   *bool                                                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMdsUpgradeTaskDetailResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s QueryMdsUpgradeTaskDetailResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentData) GetContent() *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	return s.Content
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentData) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentData) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentData) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentData) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentData) SetContent(v *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) *QueryMdsUpgradeTaskDetailResponseBodyResultContentData {
	s.Content = v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentData) SetErrorCode(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentData {
	s.ErrorCode = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentData) SetRequestId(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentData {
	s.RequestId = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentData) SetResultMsg(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentData {
	s.ResultMsg = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentData) SetSuccess(v bool) *QueryMdsUpgradeTaskDetailResponseBodyResultContentData {
	s.Success = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}

type QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent struct {
	AppCode             *string                                                                      `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	AppId               *string                                                                      `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Appstoreurl         *string                                                                      `json:"Appstoreurl,omitempty" xml:"Appstoreurl,omitempty"`
	ChannelContains     *string                                                                      `json:"ChannelContains,omitempty" xml:"ChannelContains,omitempty"`
	ChannelExcludes     *string                                                                      `json:"ChannelExcludes,omitempty" xml:"ChannelExcludes,omitempty"`
	CityContains        *string                                                                      `json:"CityContains,omitempty" xml:"CityContains,omitempty"`
	CityExcludes        *string                                                                      `json:"CityExcludes,omitempty" xml:"CityExcludes,omitempty"`
	Creator             *string                                                                      `json:"Creator,omitempty" xml:"Creator,omitempty"`
	DeviceGreyNum       *int64                                                                       `json:"DeviceGreyNum,omitempty" xml:"DeviceGreyNum,omitempty"`
	DevicePercent       *int64                                                                       `json:"DevicePercent,omitempty" xml:"DevicePercent,omitempty"`
	DownloadUrl         *string                                                                      `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ExecutionOrder      *int64                                                                       `json:"ExecutionOrder,omitempty" xml:"ExecutionOrder,omitempty"`
	GmtCreateStr        *string                                                                      `json:"GmtCreateStr,omitempty" xml:"GmtCreateStr,omitempty"`
	GreyConfigInfo      *string                                                                      `json:"GreyConfigInfo,omitempty" xml:"GreyConfigInfo,omitempty"`
	GreyEndtimeData     *string                                                                      `json:"GreyEndtimeData,omitempty" xml:"GreyEndtimeData,omitempty"`
	GreyNotice          *int64                                                                       `json:"GreyNotice,omitempty" xml:"GreyNotice,omitempty"`
	GreyNum             *int64                                                                       `json:"GreyNum,omitempty" xml:"GreyNum,omitempty"`
	GreyUv              *int64                                                                       `json:"GreyUv,omitempty" xml:"GreyUv,omitempty"`
	Id                  *int64                                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	InnerVersion        *string                                                                      `json:"InnerVersion,omitempty" xml:"InnerVersion,omitempty"`
	IsEnterprise        *int64                                                                       `json:"IsEnterprise,omitempty" xml:"IsEnterprise,omitempty"`
	IsOfficial          *int64                                                                       `json:"IsOfficial,omitempty" xml:"IsOfficial,omitempty"`
	IsPush              *int64                                                                       `json:"IsPush,omitempty" xml:"IsPush,omitempty"`
	IsRc                *int64                                                                       `json:"IsRc,omitempty" xml:"IsRc,omitempty"`
	IsRelease           *int64                                                                       `json:"IsRelease,omitempty" xml:"IsRelease,omitempty"`
	Memo                *string                                                                      `json:"Memo,omitempty" xml:"Memo,omitempty"`
	MobileModelContains *string                                                                      `json:"MobileModelContains,omitempty" xml:"MobileModelContains,omitempty"`
	MobileModelExcludes *string                                                                      `json:"MobileModelExcludes,omitempty" xml:"MobileModelExcludes,omitempty"`
	Modifier            *string                                                                      `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	NetType             *string                                                                      `json:"NetType,omitempty" xml:"NetType,omitempty"`
	OsVersion           *string                                                                      `json:"OsVersion,omitempty" xml:"OsVersion,omitempty"`
	PackageInfoId       *int64                                                                       `json:"PackageInfoId,omitempty" xml:"PackageInfoId,omitempty"`
	PackageType         *string                                                                      `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	Platform            *string                                                                      `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ProductId           *string                                                                      `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	ProductVersion      *string                                                                      `json:"ProductVersion,omitempty" xml:"ProductVersion,omitempty"`
	PublishMode         *int64                                                                       `json:"PublishMode,omitempty" xml:"PublishMode,omitempty"`
	PublishType         *int64                                                                       `json:"PublishType,omitempty" xml:"PublishType,omitempty"`
	PushContent         *string                                                                      `json:"PushContent,omitempty" xml:"PushContent,omitempty"`
	QrcodeUrl           *string                                                                      `json:"QrcodeUrl,omitempty" xml:"QrcodeUrl,omitempty"`
	ReleaseType         *string                                                                      `json:"ReleaseType,omitempty" xml:"ReleaseType,omitempty"`
	RuleJsonList        []*QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList `json:"RuleJsonList,omitempty" xml:"RuleJsonList,omitempty" type:"Repeated"`
	SilentType          *int64                                                                       `json:"SilentType,omitempty" xml:"SilentType,omitempty"`
	SyncMode            *string                                                                      `json:"SyncMode,omitempty" xml:"SyncMode,omitempty"`
	SyncResult          *string                                                                      `json:"SyncResult,omitempty" xml:"SyncResult,omitempty"`
	TaskStatus          *int64                                                                       `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	UpgradeContent      *string                                                                      `json:"UpgradeContent,omitempty" xml:"UpgradeContent,omitempty"`
	UpgradeType         *int64                                                                       `json:"UpgradeType,omitempty" xml:"UpgradeType,omitempty"`
	UpgradeValidTime    *int64                                                                       `json:"UpgradeValidTime,omitempty" xml:"UpgradeValidTime,omitempty"`
	Whitelist           []*QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist    `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Repeated"`
	WhitelistIds        *string                                                                      `json:"WhitelistIds,omitempty" xml:"WhitelistIds,omitempty"`
}

func (s QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) String() string {
	return dara.Prettify(s)
}

func (s QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GoString() string {
	return s.String()
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetAppCode() *string {
	return s.AppCode
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetAppId() *string {
	return s.AppId
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetAppstoreurl() *string {
	return s.Appstoreurl
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetChannelContains() *string {
	return s.ChannelContains
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetChannelExcludes() *string {
	return s.ChannelExcludes
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetCityContains() *string {
	return s.CityContains
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetCityExcludes() *string {
	return s.CityExcludes
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetCreator() *string {
	return s.Creator
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetDeviceGreyNum() *int64 {
	return s.DeviceGreyNum
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetDevicePercent() *int64 {
	return s.DevicePercent
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetExecutionOrder() *int64 {
	return s.ExecutionOrder
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetGmtCreateStr() *string {
	return s.GmtCreateStr
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetGreyConfigInfo() *string {
	return s.GreyConfigInfo
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetGreyEndtimeData() *string {
	return s.GreyEndtimeData
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetGreyNotice() *int64 {
	return s.GreyNotice
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetGreyNum() *int64 {
	return s.GreyNum
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetGreyUv() *int64 {
	return s.GreyUv
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetId() *int64 {
	return s.Id
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetInnerVersion() *string {
	return s.InnerVersion
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetIsEnterprise() *int64 {
	return s.IsEnterprise
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetIsOfficial() *int64 {
	return s.IsOfficial
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetIsPush() *int64 {
	return s.IsPush
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetIsRc() *int64 {
	return s.IsRc
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetIsRelease() *int64 {
	return s.IsRelease
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetMemo() *string {
	return s.Memo
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetMobileModelContains() *string {
	return s.MobileModelContains
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetMobileModelExcludes() *string {
	return s.MobileModelExcludes
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetModifier() *string {
	return s.Modifier
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetNetType() *string {
	return s.NetType
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetOsVersion() *string {
	return s.OsVersion
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetPackageInfoId() *int64 {
	return s.PackageInfoId
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetPackageType() *string {
	return s.PackageType
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetPlatform() *string {
	return s.Platform
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetProductId() *string {
	return s.ProductId
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetProductVersion() *string {
	return s.ProductVersion
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetPublishMode() *int64 {
	return s.PublishMode
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetPublishType() *int64 {
	return s.PublishType
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetPushContent() *string {
	return s.PushContent
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetQrcodeUrl() *string {
	return s.QrcodeUrl
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetReleaseType() *string {
	return s.ReleaseType
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetRuleJsonList() []*QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList {
	return s.RuleJsonList
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetSilentType() *int64 {
	return s.SilentType
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetSyncMode() *string {
	return s.SyncMode
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetSyncResult() *string {
	return s.SyncResult
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetTaskStatus() *int64 {
	return s.TaskStatus
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetUpgradeContent() *string {
	return s.UpgradeContent
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetUpgradeType() *int64 {
	return s.UpgradeType
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetUpgradeValidTime() *int64 {
	return s.UpgradeValidTime
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetWhitelist() []*QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist {
	return s.Whitelist
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) GetWhitelistIds() *string {
	return s.WhitelistIds
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetAppCode(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.AppCode = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetAppId(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.AppId = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetAppstoreurl(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.Appstoreurl = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetChannelContains(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.ChannelContains = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetChannelExcludes(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.ChannelExcludes = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetCityContains(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.CityContains = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetCityExcludes(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.CityExcludes = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetCreator(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.Creator = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetDeviceGreyNum(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.DeviceGreyNum = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetDevicePercent(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.DevicePercent = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetDownloadUrl(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.DownloadUrl = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetExecutionOrder(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.ExecutionOrder = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetGmtCreateStr(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.GmtCreateStr = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetGreyConfigInfo(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.GreyConfigInfo = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetGreyEndtimeData(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.GreyEndtimeData = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetGreyNotice(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.GreyNotice = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetGreyNum(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.GreyNum = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetGreyUv(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.GreyUv = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetId(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.Id = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetInnerVersion(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.InnerVersion = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetIsEnterprise(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.IsEnterprise = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetIsOfficial(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.IsOfficial = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetIsPush(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.IsPush = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetIsRc(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.IsRc = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetIsRelease(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.IsRelease = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetMemo(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.Memo = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetMobileModelContains(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.MobileModelContains = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetMobileModelExcludes(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.MobileModelExcludes = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetModifier(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.Modifier = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetNetType(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.NetType = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetOsVersion(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.OsVersion = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetPackageInfoId(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.PackageInfoId = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetPackageType(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.PackageType = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetPlatform(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.Platform = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetProductId(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.ProductId = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetProductVersion(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.ProductVersion = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetPublishMode(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.PublishMode = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetPublishType(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.PublishType = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetPushContent(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.PushContent = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetQrcodeUrl(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.QrcodeUrl = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetReleaseType(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.ReleaseType = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetRuleJsonList(v []*QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.RuleJsonList = v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetSilentType(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.SilentType = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetSyncMode(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.SyncMode = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetSyncResult(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.SyncResult = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetTaskStatus(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.TaskStatus = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetUpgradeContent(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.UpgradeContent = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetUpgradeType(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.UpgradeType = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetUpgradeValidTime(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.UpgradeValidTime = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetWhitelist(v []*QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.Whitelist = v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) SetWhitelistIds(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent {
	s.WhitelistIds = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContent) Validate() error {
	return dara.Validate(s)
}

type QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList struct {
	Operation   *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	RuleElement *string `json:"RuleElement,omitempty" xml:"RuleElement,omitempty"`
	RuleType    *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList) String() string {
	return dara.Prettify(s)
}

func (s QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList) GoString() string {
	return s.String()
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList) GetOperation() *string {
	return s.Operation
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList) GetRuleElement() *string {
	return s.RuleElement
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList) GetRuleType() *string {
	return s.RuleType
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList) GetValue() *string {
	return s.Value
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList) SetOperation(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList {
	s.Operation = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList) SetRuleElement(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList {
	s.RuleElement = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList) SetRuleType(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList {
	s.RuleType = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList) SetValue(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList {
	s.Value = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentRuleJsonList) Validate() error {
	return dara.Validate(s)
}

type QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist struct {
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

func (s QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) String() string {
	return dara.Prettify(s)
}

func (s QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) GoString() string {
	return s.String()
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) GetAppCode() *string {
	return s.AppCode
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) GetBusiness() *string {
	return s.Business
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) GetId() *int64 {
	return s.Id
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) GetIdType() *string {
	return s.IdType
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) GetPlatform() *string {
	return s.Platform
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) GetStatus() *int64 {
	return s.Status
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) GetWhiteListCount() *int64 {
	return s.WhiteListCount
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) GetWhiteListName() *string {
	return s.WhiteListName
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) SetAppCode(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist {
	s.AppCode = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) SetBusiness(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist {
	s.Business = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) SetGmtModified(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist {
	s.GmtModified = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) SetId(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist {
	s.Id = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) SetIdType(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist {
	s.IdType = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) SetPlatform(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist {
	s.Platform = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) SetStatus(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist {
	s.Status = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) SetWhiteListCount(v int64) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist {
	s.WhiteListCount = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) SetWhiteListName(v string) *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist {
	s.WhiteListName = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailResponseBodyResultContentDataContentWhitelist) Validate() error {
	return dara.Validate(s)
}
