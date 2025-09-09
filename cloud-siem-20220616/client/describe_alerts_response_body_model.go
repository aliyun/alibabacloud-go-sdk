// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAlertsResponseBody
	GetCode() *int32
	SetData(v *DescribeAlertsResponseBodyData) *DescribeAlertsResponseBody
	GetData() *DescribeAlertsResponseBodyData
	SetMessage(v string) *DescribeAlertsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAlertsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAlertsResponseBody
	GetSuccess() *bool
}

type DescribeAlertsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeAlertsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAlertsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAlertsResponseBody) GetData() *DescribeAlertsResponseBodyData {
	return s.Data
}

func (s *DescribeAlertsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertsResponseBody) SetCode(v int32) *DescribeAlertsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertsResponseBody) SetData(v *DescribeAlertsResponseBodyData) *DescribeAlertsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertsResponseBody) SetMessage(v string) *DescribeAlertsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertsResponseBody) SetRequestId(v string) *DescribeAlertsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertsResponseBody) SetSuccess(v bool) *DescribeAlertsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertsResponseBodyData struct {
	// The pagination information.
	PageInfo *DescribeAlertsResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*DescribeAlertsResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeAlertsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertsResponseBodyData) GetPageInfo() *DescribeAlertsResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *DescribeAlertsResponseBodyData) GetResponseData() []*DescribeAlertsResponseBodyDataResponseData {
	return s.ResponseData
}

func (s *DescribeAlertsResponseBodyData) SetPageInfo(v *DescribeAlertsResponseBodyDataPageInfo) *DescribeAlertsResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeAlertsResponseBodyData) SetResponseData(v []*DescribeAlertsResponseBodyDataResponseData) *DescribeAlertsResponseBodyData {
	s.ResponseData = v
	return s
}

func (s *DescribeAlertsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertsResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAlertsResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAlertsResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAlertsResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAlertsResponseBodyDataPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAlertsResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeAlertsResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeAlertsResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeAlertsResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertsResponseBodyDataResponseData struct {
	// The description of the alert.
	//
	// example:
	//
	// The detection model found a suspicious Webshell file on your server, which may be a backdoor file implanted to maintain permissions after the attacker successfully invaded the website.
	AlertDesc *string `json:"AlertDesc,omitempty" xml:"AlertDesc,omitempty"`
	// The internal code of the alert description.
	//
	// example:
	//
	// security_event_config.event_name.webshell
	AlertDescCode *string `json:"AlertDescCode,omitempty" xml:"AlertDescCode,omitempty"`
	// The description of the alert in English.
	//
	// example:
	//
	// The detection model found a suspicious Webshell file on your server, which may be a backdoor file implanted to maintain permissions after the attacker successfully invaded the website.
	AlertDescEn *string `json:"AlertDescEn,omitempty" xml:"AlertDescEn,omitempty"`
	// The details of the alert.
	//
	// example:
	//
	// {"main_user_id": "165295629792****";"log_uuid_count": "99";"attack_ip": "21.92.*.*"}
	AlertDetail *string `json:"AlertDetail,omitempty" xml:"AlertDetail,omitempty"`
	// The displayed details of the alert.
	//
	// example:
	//
	// aliyun
	AlertInfoList []*DescribeAlertsResponseBodyDataResponseDataAlertInfoList `json:"AlertInfoList,omitempty" xml:"AlertInfoList,omitempty" type:"Repeated"`
	// The threat level. Valid values:
	//
	// 	- serious: high
	//
	// 	- suspicious: medium
	//
	// 	- remind: low
	//
	// example:
	//
	// remind
	AlertLevel *string `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// The name of the alert.
	//
	// example:
	//
	// Try SNMP weak password
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The internal code of the alert name.
	//
	// example:
	//
	// security_event_config.event_name.webshell
	AlertNameCode *string `json:"AlertNameCode,omitempty" xml:"AlertNameCode,omitempty"`
	// The name of the alert in English.
	//
	// example:
	//
	// Try SNMP weak password
	AlertNameEn *string `json:"AlertNameEn,omitempty" xml:"AlertNameEn,omitempty"`
	// The service for which the alert associated with the event is generated.
	//
	// example:
	//
	// sas
	AlertSrcProd *string `json:"AlertSrcProd,omitempty" xml:"AlertSrcProd,omitempty"`
	// The sub-module of ther alert source.
	//
	// example:
	//
	// waf
	AlertSrcProdModule *string `json:"AlertSrcProdModule,omitempty" xml:"AlertSrcProdModule,omitempty"`
	// The title of the alert.
	//
	// example:
	//
	// Scan-Try SNMP weak password
	AlertTitle *string `json:"AlertTitle,omitempty" xml:"AlertTitle,omitempty"`
	// The title of the alert in English.
	//
	// example:
	//
	// Scan-Try SNMP weak password
	AlertTitleEn *string `json:"AlertTitleEn,omitempty" xml:"AlertTitleEn,omitempty"`
	// The alert type.
	//
	// example:
	//
	// Scan
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The internal code of the alert type.
	//
	// example:
	//
	// security_event_config.event_name.webshellName
	AlertTypeCode *string `json:"AlertTypeCode,omitempty" xml:"AlertTypeCode,omitempty"`
	// The type of the alert in English.
	//
	// example:
	//
	// Scan
	AlertTypeEn *string `json:"AlertTypeEn,omitempty" xml:"AlertTypeEn,omitempty"`
	// The UUID of the alert.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The details of the asset.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "is_main_asset": "1",
	//
	//             "asset_name": "47.245.*",
	//
	//             "port": "22",
	//
	//             "ip": "47.245.*",
	//
	//             "asset_type": "ip",
	//
	//             "location": "ap-southeast-1",
	//
	//             "asset_id": "47.245.*",
	//
	//             "net_connect_dir": "in"
	//
	//       }
	//
	// ]
	AssetList *string `json:"AssetList,omitempty" xml:"AssetList,omitempty"`
	// The tag of the ATT\\&CK attack.
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	// The cloud code. Valid values:
	//
	// 	- aliyun: Alibaba Cloud
	//
	// 	- qcloud: Tencent Cloud
	//
	// 	- hcloud: Huawei Cloud
	//
	// example:
	//
	// aliyun
	CloudCode       *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	DetectionRuleId *string `json:"DetectionRuleId,omitempty" xml:"DetectionRuleId,omitempty"`
	// The time when the alert was closed.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EntityList    *string `json:"EntityList,omitempty" xml:"EntityList,omitempty"`
	ExtendContent *string `json:"ExtendContent,omitempty" xml:"ExtendContent,omitempty"`
	// The time when the alert was received.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the alert was last updated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The unique ID of the alert.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// Indicates whether an attack is defended. Valid values:
	//
	// 	- 0: detected.
	//
	// 	- 1: blocked.
	//
	// example:
	//
	// 1
	IsDefend *string `json:"IsDefend,omitempty" xml:"IsDefend,omitempty"`
	// The time when the alert was recorded.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	LogTime *string `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	// The UUID of the alert log.
	//
	// example:
	//
	// cfw_d12e285a-a042-4d7e-be89-f8a795ef****
	LogUuid *string `json:"LogUuid,omitempty" xml:"LogUuid,omitempty"`
	// The ID of the Alibaba Cloud account that is associated with the alert in SIEM.
	//
	// example:
	//
	// 127608589417****
	MainUserId *int64 `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The time when the alert is triggered.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	OccurTime *string `json:"OccurTime,omitempty" xml:"OccurTime,omitempty"`
	ProductId *string `json:"ProductId,omitempty" xml:"ProductId,omitempty"`
	// The time at which the alert was first generated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The ID of the Alibaba Cloud account within which the alert is generated.
	//
	// example:
	//
	// 176555323***
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// example:
	//
	// 176555323***
	SubUserName *string `json:"SubUserName,omitempty" xml:"SubUserName,omitempty"`
	// example:
	//
	// aliyun
	VendorId *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s DescribeAlertsResponseBodyDataResponseData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertDesc() *string {
	return s.AlertDesc
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertDescCode() *string {
	return s.AlertDescCode
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertDescEn() *string {
	return s.AlertDescEn
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertDetail() *string {
	return s.AlertDetail
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertInfoList() []*DescribeAlertsResponseBodyDataResponseDataAlertInfoList {
	return s.AlertInfoList
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertLevel() *string {
	return s.AlertLevel
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertName() *string {
	return s.AlertName
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertNameCode() *string {
	return s.AlertNameCode
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertNameEn() *string {
	return s.AlertNameEn
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertSrcProd() *string {
	return s.AlertSrcProd
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertSrcProdModule() *string {
	return s.AlertSrcProdModule
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertTitle() *string {
	return s.AlertTitle
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertTitleEn() *string {
	return s.AlertTitleEn
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertType() *string {
	return s.AlertType
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertTypeCode() *string {
	return s.AlertTypeCode
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertTypeEn() *string {
	return s.AlertTypeEn
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAssetList() *string {
	return s.AssetList
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetAttCk() *string {
	return s.AttCk
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetCloudCode() *string {
	return s.CloudCode
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetDetectionRuleId() *string {
	return s.DetectionRuleId
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetEntityList() *string {
	return s.EntityList
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetExtendContent() *string {
	return s.ExtendContent
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetId() *int64 {
	return s.Id
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetIsDefend() *string {
	return s.IsDefend
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetLogTime() *string {
	return s.LogTime
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetLogUuid() *string {
	return s.LogUuid
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetMainUserId() *int64 {
	return s.MainUserId
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetOccurTime() *string {
	return s.OccurTime
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetProductId() *string {
	return s.ProductId
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetSubUserName() *string {
	return s.SubUserName
}

func (s *DescribeAlertsResponseBodyDataResponseData) GetVendorId() *string {
	return s.VendorId
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertDesc(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertDesc = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertDescCode(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertDescCode = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertDescEn(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertDescEn = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertDetail(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertDetail = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertInfoList(v []*DescribeAlertsResponseBodyDataResponseDataAlertInfoList) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertInfoList = v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertLevel(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertLevel = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertName(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertNameCode(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertNameCode = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertNameEn(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertNameEn = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertSrcProd(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertSrcProd = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertSrcProdModule(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertSrcProdModule = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertTitle(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertTitle = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertTitleEn(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertTitleEn = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertType(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertTypeCode(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertTypeCode = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertTypeEn(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertTypeEn = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAlertUuid(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAssetList(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AssetList = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetAttCk(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.AttCk = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetCloudCode(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.CloudCode = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetDetectionRuleId(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.DetectionRuleId = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetEndTime(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetEntityList(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.EntityList = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetExtendContent(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.ExtendContent = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetGmtCreate(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetGmtModified(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetId(v int64) *DescribeAlertsResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetIncidentUuid(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetIsDefend(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.IsDefend = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetLogTime(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.LogTime = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetLogUuid(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.LogUuid = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetMainUserId(v int64) *DescribeAlertsResponseBodyDataResponseData {
	s.MainUserId = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetOccurTime(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.OccurTime = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetProductId(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.ProductId = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetStartTime(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetSubUserId(v int64) *DescribeAlertsResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetSubUserName(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.SubUserName = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) SetVendorId(v string) *DescribeAlertsResponseBodyDataResponseData {
	s.VendorId = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseData) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertsResponseBodyDataResponseDataAlertInfoList struct {
	// The attribute key.
	//
	// example:
	//
	// suspicious.wbd.wb.trojanpath
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the key.
	//
	// example:
	//
	// Trojan Path
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The value of the key.
	//
	// example:
	//
	// /root/test33.php
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeAlertsResponseBodyDataResponseDataAlertInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsResponseBodyDataResponseDataAlertInfoList) GoString() string {
	return s.String()
}

func (s *DescribeAlertsResponseBodyDataResponseDataAlertInfoList) GetKey() *string {
	return s.Key
}

func (s *DescribeAlertsResponseBodyDataResponseDataAlertInfoList) GetKeyName() *string {
	return s.KeyName
}

func (s *DescribeAlertsResponseBodyDataResponseDataAlertInfoList) GetValues() *string {
	return s.Values
}

func (s *DescribeAlertsResponseBodyDataResponseDataAlertInfoList) SetKey(v string) *DescribeAlertsResponseBodyDataResponseDataAlertInfoList {
	s.Key = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseDataAlertInfoList) SetKeyName(v string) *DescribeAlertsResponseBodyDataResponseDataAlertInfoList {
	s.KeyName = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseDataAlertInfoList) SetValues(v string) *DescribeAlertsResponseBodyDataResponseDataAlertInfoList {
	s.Values = &v
	return s
}

func (s *DescribeAlertsResponseBodyDataResponseDataAlertInfoList) Validate() error {
	return dara.Validate(s)
}
