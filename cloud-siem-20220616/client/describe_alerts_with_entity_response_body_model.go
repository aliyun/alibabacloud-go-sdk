// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertsWithEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAlertsWithEntityResponseBody
	GetCode() *int32
	SetData(v *DescribeAlertsWithEntityResponseBodyData) *DescribeAlertsWithEntityResponseBody
	GetData() *DescribeAlertsWithEntityResponseBodyData
	SetMessage(v string) *DescribeAlertsWithEntityResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAlertsWithEntityResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAlertsWithEntityResponseBody
	GetSuccess() *bool
}

type DescribeAlertsWithEntityResponseBody struct {
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
	Data *DescribeAlertsWithEntityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DescribeAlertsWithEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsWithEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEntityResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAlertsWithEntityResponseBody) GetData() *DescribeAlertsWithEntityResponseBodyData {
	return s.Data
}

func (s *DescribeAlertsWithEntityResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertsWithEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertsWithEntityResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertsWithEntityResponseBody) SetCode(v int32) *DescribeAlertsWithEntityResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBody) SetData(v *DescribeAlertsWithEntityResponseBodyData) *DescribeAlertsWithEntityResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertsWithEntityResponseBody) SetMessage(v string) *DescribeAlertsWithEntityResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBody) SetRequestId(v string) *DescribeAlertsWithEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBody) SetSuccess(v bool) *DescribeAlertsWithEntityResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertsWithEntityResponseBodyData struct {
	// The pagination information.
	PageInfo *DescribeAlertsWithEntityResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*DescribeAlertsWithEntityResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeAlertsWithEntityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsWithEntityResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEntityResponseBodyData) GetPageInfo() *DescribeAlertsWithEntityResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *DescribeAlertsWithEntityResponseBodyData) GetResponseData() []*DescribeAlertsWithEntityResponseBodyDataResponseData {
	return s.ResponseData
}

func (s *DescribeAlertsWithEntityResponseBodyData) SetPageInfo(v *DescribeAlertsWithEntityResponseBodyDataPageInfo) *DescribeAlertsWithEntityResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyData) SetResponseData(v []*DescribeAlertsWithEntityResponseBodyDataResponseData) *DescribeAlertsWithEntityResponseBodyData {
	s.ResponseData = v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertsWithEntityResponseBodyDataPageInfo struct {
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

func (s DescribeAlertsWithEntityResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsWithEntityResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEntityResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAlertsWithEntityResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAlertsWithEntityResponseBodyDataPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAlertsWithEntityResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeAlertsWithEntityResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeAlertsWithEntityResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeAlertsWithEntityResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertsWithEntityResponseBodyDataResponseData struct {
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
	// The alert description in English.
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
	AlertInfoList []*DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList `json:"AlertInfoList,omitempty" xml:"AlertInfoList,omitempty" type:"Repeated"`
	// The risk level. Valid values:
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
	// The name of the alert.
	//
	// example:
	//
	// Try SNMP weak password
	AlertNameEn *string `json:"AlertNameEn,omitempty" xml:"AlertNameEn,omitempty"`
	// The source of the alert.
	//
	// example:
	//
	// sas
	AlertSrcProd *string `json:"AlertSrcProd,omitempty" xml:"AlertSrcProd,omitempty"`
	// The sub-module of the alert source.
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
	// The alert title in English.
	//
	// example:
	//
	// Scan-Try SNMP weak password
	AlertTitleEn *string `json:"AlertTitleEn,omitempty" xml:"AlertTitleEn,omitempty"`
	// The type of the alert.
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
	// The alert type in English.
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
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EntityList *string `json:"EntityList,omitempty" xml:"EntityList,omitempty"`
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
	// Specifies whether an attack is defended. Valid values:
	//
	// 	- 0: detected
	//
	// 	- 1: blocked
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
	// The time when the alert was triggered.
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
	VendorId    *string `json:"VendorId,omitempty" xml:"VendorId,omitempty"`
}

func (s DescribeAlertsWithEntityResponseBodyDataResponseData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsWithEntityResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertDesc() *string {
	return s.AlertDesc
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertDescCode() *string {
	return s.AlertDescCode
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertDescEn() *string {
	return s.AlertDescEn
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertDetail() *string {
	return s.AlertDetail
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertInfoList() []*DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList {
	return s.AlertInfoList
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertLevel() *string {
	return s.AlertLevel
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertName() *string {
	return s.AlertName
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertNameCode() *string {
	return s.AlertNameCode
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertNameEn() *string {
	return s.AlertNameEn
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertSrcProd() *string {
	return s.AlertSrcProd
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertSrcProdModule() *string {
	return s.AlertSrcProdModule
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertTitle() *string {
	return s.AlertTitle
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertTitleEn() *string {
	return s.AlertTitleEn
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertType() *string {
	return s.AlertType
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertTypeCode() *string {
	return s.AlertTypeCode
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertTypeEn() *string {
	return s.AlertTypeEn
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAssetList() *string {
	return s.AssetList
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetAttCk() *string {
	return s.AttCk
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetCloudCode() *string {
	return s.CloudCode
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetDetectionRuleId() *string {
	return s.DetectionRuleId
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetEntityList() *string {
	return s.EntityList
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetId() *int64 {
	return s.Id
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetIsDefend() *string {
	return s.IsDefend
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetLogTime() *string {
	return s.LogTime
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetLogUuid() *string {
	return s.LogUuid
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetMainUserId() *int64 {
	return s.MainUserId
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetOccurTime() *string {
	return s.OccurTime
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetProductId() *string {
	return s.ProductId
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetSubUserName() *string {
	return s.SubUserName
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) GetVendorId() *string {
	return s.VendorId
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertDesc(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertDesc = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertDescCode(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertDescCode = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertDescEn(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertDescEn = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertDetail(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertDetail = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertInfoList(v []*DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertInfoList = v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertLevel(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertLevel = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertName(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertNameCode(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertNameCode = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertNameEn(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertNameEn = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertSrcProd(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertSrcProd = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertSrcProdModule(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertSrcProdModule = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertTitle(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertTitle = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertTitleEn(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertTitleEn = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertType(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertTypeCode(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertTypeCode = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertTypeEn(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertTypeEn = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAlertUuid(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAssetList(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AssetList = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetAttCk(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.AttCk = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetCloudCode(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.CloudCode = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetDetectionRuleId(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.DetectionRuleId = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetEndTime(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetEntityList(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.EntityList = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetGmtCreate(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetGmtModified(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetId(v int64) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetIncidentUuid(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetIsDefend(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.IsDefend = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetLogTime(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.LogTime = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetLogUuid(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.LogUuid = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetMainUserId(v int64) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.MainUserId = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetOccurTime(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.OccurTime = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetProductId(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.ProductId = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetStartTime(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetSubUserId(v int64) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetSubUserName(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.SubUserName = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) SetVendorId(v string) *DescribeAlertsWithEntityResponseBodyDataResponseData {
	s.VendorId = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseData) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList struct {
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

func (s DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) GetKey() *string {
	return s.Key
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) GetKeyName() *string {
	return s.KeyName
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) GetValues() *string {
	return s.Values
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) SetKey(v string) *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList {
	s.Key = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) SetKeyName(v string) *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList {
	s.KeyName = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) SetValues(v string) *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList {
	s.Values = &v
	return s
}

func (s *DescribeAlertsWithEntityResponseBodyDataResponseDataAlertInfoList) Validate() error {
	return dara.Validate(s)
}
