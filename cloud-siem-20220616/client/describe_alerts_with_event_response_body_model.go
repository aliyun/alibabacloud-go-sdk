// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertsWithEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAlertsWithEventResponseBody
	GetCode() *int32
	SetData(v *DescribeAlertsWithEventResponseBodyData) *DescribeAlertsWithEventResponseBody
	GetData() *DescribeAlertsWithEventResponseBodyData
	SetMessage(v string) *DescribeAlertsWithEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAlertsWithEventResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAlertsWithEventResponseBody
	GetSuccess() *bool
}

type DescribeAlertsWithEventResponseBody struct {
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
	Data *DescribeAlertsWithEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s DescribeAlertsWithEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsWithEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEventResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAlertsWithEventResponseBody) GetData() *DescribeAlertsWithEventResponseBodyData {
	return s.Data
}

func (s *DescribeAlertsWithEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAlertsWithEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAlertsWithEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAlertsWithEventResponseBody) SetCode(v int32) *DescribeAlertsWithEventResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBody) SetData(v *DescribeAlertsWithEventResponseBodyData) *DescribeAlertsWithEventResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAlertsWithEventResponseBody) SetMessage(v string) *DescribeAlertsWithEventResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBody) SetRequestId(v string) *DescribeAlertsWithEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBody) SetSuccess(v bool) *DescribeAlertsWithEventResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAlertsWithEventResponseBodyData struct {
	// The pagination information.
	PageInfo *DescribeAlertsWithEventResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*DescribeAlertsWithEventResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeAlertsWithEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsWithEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEventResponseBodyData) GetPageInfo() *DescribeAlertsWithEventResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *DescribeAlertsWithEventResponseBodyData) GetResponseData() []*DescribeAlertsWithEventResponseBodyDataResponseData {
	return s.ResponseData
}

func (s *DescribeAlertsWithEventResponseBodyData) SetPageInfo(v *DescribeAlertsWithEventResponseBodyDataPageInfo) *DescribeAlertsWithEventResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyData) SetResponseData(v []*DescribeAlertsWithEventResponseBodyDataResponseData) *DescribeAlertsWithEventResponseBodyData {
	s.ResponseData = v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyData) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ResponseData != nil {
		for _, item := range s.ResponseData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlertsWithEventResponseBodyDataPageInfo struct {
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

func (s DescribeAlertsWithEventResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsWithEventResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEventResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeAlertsWithEventResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAlertsWithEventResponseBodyDataPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAlertsWithEventResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeAlertsWithEventResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeAlertsWithEventResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeAlertsWithEventResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeAlertsWithEventResponseBodyDataResponseData struct {
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
	AlertInfoList []*DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList `json:"AlertInfoList,omitempty" xml:"AlertInfoList,omitempty" type:"Repeated"`
	// The risk level. Valid values:
	//
	// 	- serious: high.
	//
	// 	- suspicious: medium.
	//
	// 	- remind: low.
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
	// The alert name in English.
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
	// The tag of the ATT\\&CK technique.
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	// The code of the cloud service provider. Valid values:
	//
	// 	- aliyun: Alibaba Cloud.
	//
	// 	- qcloud: Tencent Cloud.
	//
	// 	- hcloud: Huawei Cloud.
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
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The details of the entity.
	//
	// example:
	//
	// [{&quot;entity_user_id&quot;:&quot;198921674491****&quot;,&quot;entity_account_id&quot;:&quot;N/A&quot;,&quot;entity_uuid&quot;:&quot;6245f979d5dd9ef8dd19bdc72228****&quot;,&quot;entity_type&quot;:&quot;host&quot;,&quot;entity_name&quot;:&quot;zhh-test-20240409&quot;,&quot;is_comprised&quot;:&quot;1&quot;,&quot;os_type&quot;:&quot;linux&quot;,&quot;entity_id&quot;:&quot;a88f44dd-b8d4-4ded-831c-77a4835****&quot;,&quot;host_uuid&quot;:&quot;a88f44dd-b8d4-4ded-831c-77a4835****&quot;,&quot;host_name&quot;:&quot;zhh-test-2024****&quot;}]
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
	// Indicates whether an attack is defended against. Valid values:
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

func (s DescribeAlertsWithEventResponseBodyDataResponseData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsWithEventResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertDesc() *string {
	return s.AlertDesc
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertDescCode() *string {
	return s.AlertDescCode
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertDescEn() *string {
	return s.AlertDescEn
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertDetail() *string {
	return s.AlertDetail
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertInfoList() []*DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList {
	return s.AlertInfoList
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertLevel() *string {
	return s.AlertLevel
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertName() *string {
	return s.AlertName
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertNameCode() *string {
	return s.AlertNameCode
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertNameEn() *string {
	return s.AlertNameEn
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertSrcProd() *string {
	return s.AlertSrcProd
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertSrcProdModule() *string {
	return s.AlertSrcProdModule
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertTitle() *string {
	return s.AlertTitle
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertTitleEn() *string {
	return s.AlertTitleEn
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertType() *string {
	return s.AlertType
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertTypeCode() *string {
	return s.AlertTypeCode
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertTypeEn() *string {
	return s.AlertTypeEn
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAssetList() *string {
	return s.AssetList
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetAttCk() *string {
	return s.AttCk
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetCloudCode() *string {
	return s.CloudCode
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetDetectionRuleId() *string {
	return s.DetectionRuleId
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetEntityList() *string {
	return s.EntityList
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetExtendContent() *string {
	return s.ExtendContent
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetId() *int64 {
	return s.Id
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetIsDefend() *string {
	return s.IsDefend
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetLogTime() *string {
	return s.LogTime
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetLogUuid() *string {
	return s.LogUuid
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetMainUserId() *int64 {
	return s.MainUserId
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetOccurTime() *string {
	return s.OccurTime
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetProductId() *string {
	return s.ProductId
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetSubUserName() *string {
	return s.SubUserName
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) GetVendorId() *string {
	return s.VendorId
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertDesc(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertDesc = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertDescCode(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertDescCode = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertDescEn(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertDescEn = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertDetail(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertDetail = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertInfoList(v []*DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertInfoList = v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertLevel(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertLevel = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertName(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertName = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertNameCode(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertNameCode = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertNameEn(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertNameEn = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertSrcProd(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertSrcProd = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertSrcProdModule(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertSrcProdModule = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertTitle(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertTitle = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertTitleEn(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertTitleEn = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertType(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertType = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertTypeCode(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertTypeCode = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertTypeEn(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertTypeEn = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAlertUuid(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAssetList(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AssetList = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetAttCk(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.AttCk = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetCloudCode(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.CloudCode = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetDetectionRuleId(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.DetectionRuleId = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetEndTime(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.EndTime = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetEntityList(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.EntityList = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetExtendContent(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.ExtendContent = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetGmtCreate(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetGmtModified(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetId(v int64) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetIncidentUuid(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetIsDefend(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.IsDefend = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetLogTime(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.LogTime = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetLogUuid(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.LogUuid = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetMainUserId(v int64) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.MainUserId = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetOccurTime(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.OccurTime = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetProductId(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.ProductId = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetStartTime(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.StartTime = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetSubUserId(v int64) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetSubUserName(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.SubUserName = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) SetVendorId(v string) *DescribeAlertsWithEventResponseBodyDataResponseData {
	s.VendorId = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseData) Validate() error {
	if s.AlertInfoList != nil {
		for _, item := range s.AlertInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList struct {
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

func (s DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) GoString() string {
	return s.String()
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) GetKey() *string {
	return s.Key
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) GetKeyName() *string {
	return s.KeyName
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) GetValues() *string {
	return s.Values
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) SetKey(v string) *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList {
	s.Key = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) SetKeyName(v string) *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList {
	s.KeyName = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) SetValues(v string) *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList {
	s.Values = &v
	return s
}

func (s *DescribeAlertsWithEventResponseBodyDataResponseDataAlertInfoList) Validate() error {
	return dara.Validate(s)
}
