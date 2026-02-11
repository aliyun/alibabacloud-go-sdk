// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomizeRuleTestResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListCustomizeRuleTestResultResponseBody
	GetCode() *int32
	SetData(v *ListCustomizeRuleTestResultResponseBodyData) *ListCustomizeRuleTestResultResponseBody
	GetData() *ListCustomizeRuleTestResultResponseBodyData
	SetMessage(v string) *ListCustomizeRuleTestResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCustomizeRuleTestResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCustomizeRuleTestResultResponseBody
	GetSuccess() *bool
}

type ListCustomizeRuleTestResultResponseBody struct {
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
	Data *ListCustomizeRuleTestResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListCustomizeRuleTestResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizeRuleTestResultResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListCustomizeRuleTestResultResponseBody) GetData() *ListCustomizeRuleTestResultResponseBodyData {
	return s.Data
}

func (s *ListCustomizeRuleTestResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCustomizeRuleTestResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomizeRuleTestResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCustomizeRuleTestResultResponseBody) SetCode(v int32) *ListCustomizeRuleTestResultResponseBody {
	s.Code = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBody) SetData(v *ListCustomizeRuleTestResultResponseBodyData) *ListCustomizeRuleTestResultResponseBody {
	s.Data = v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBody) SetMessage(v string) *ListCustomizeRuleTestResultResponseBody {
	s.Message = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBody) SetRequestId(v string) *ListCustomizeRuleTestResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBody) SetSuccess(v bool) *ListCustomizeRuleTestResultResponseBody {
	s.Success = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomizeRuleTestResultResponseBodyData struct {
	// The pagination information.
	PageInfo *ListCustomizeRuleTestResultResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*ListCustomizeRuleTestResultResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListCustomizeRuleTestResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizeRuleTestResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultResponseBodyData) GetPageInfo() *ListCustomizeRuleTestResultResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *ListCustomizeRuleTestResultResponseBodyData) GetResponseData() []*ListCustomizeRuleTestResultResponseBodyDataResponseData {
	return s.ResponseData
}

func (s *ListCustomizeRuleTestResultResponseBodyData) SetPageInfo(v *ListCustomizeRuleTestResultResponseBodyDataPageInfo) *ListCustomizeRuleTestResultResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyData) SetResponseData(v []*ListCustomizeRuleTestResultResponseBodyDataResponseData) *ListCustomizeRuleTestResultResponseBodyData {
	s.ResponseData = v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyData) Validate() error {
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

type ListCustomizeRuleTestResultResponseBodyDataPageInfo struct {
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
	TotalCount    *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VerifiedCount *int64 `json:"VerifiedCount,omitempty" xml:"VerifiedCount,omitempty"`
}

func (s ListCustomizeRuleTestResultResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizeRuleTestResultResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) GetVerifiedCount() *int64 {
	return s.VerifiedCount
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListCustomizeRuleTestResultResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) SetPageSize(v int32) *ListCustomizeRuleTestResultResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) SetTotalCount(v int64) *ListCustomizeRuleTestResultResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) SetVerifiedCount(v int64) *ListCustomizeRuleTestResultResponseBodyDataPageInfo {
	s.VerifiedCount = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListCustomizeRuleTestResultResponseBodyDataResponseData struct {
	// The description of the alert.
	//
	// example:
	//
	// The account you logged in this time is not in the legal account category defined by you. Please confirm the legality of the login behavior.
	AlertDesc *string `json:"AlertDesc,omitempty" xml:"AlertDesc,omitempty"`
	// The alert details in the JSON format.
	//
	// example:
	//
	// {"main_user_id": "165295629792****";"log_uuid_count": "99";"attack_ip": "218.92.XX.XX"}
	AlertDetail *string `json:"AlertDetail,omitempty" xml:"AlertDetail,omitempty"`
	// The source of the alert.
	//
	// example:
	//
	// sas
	AlertSrcProd *string `json:"AlertSrcProd,omitempty" xml:"AlertSrcProd,omitempty"`
	// The sub-module of the source.
	//
	// example:
	//
	// waf
	AlertSrcProdModule *string `json:"AlertSrcProdModule,omitempty" xml:"AlertSrcProdModule,omitempty"`
	// The tag of the ATT\\&CK attack.
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	// The name of the alert, which corresponds to the name of the custom rule.
	//
	// example:
	//
	// waf_scan
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The threat type, which indicates the alert type.
	//
	// example:
	//
	// WEBSHELL
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The threat level. Valid values:
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
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The log source of the rule.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The time when the alert was recorded.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	LogTime *string `json:"LogTime,omitempty" xml:"LogTime,omitempty"`
	// The log type of the rule.
	//
	// example:
	//
	// ALERT_ACTIVITY
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The ID of the Alibaba Cloud account that is associated with the alert in SIEM.
	//
	// example:
	//
	// 127608589417****
	MainUserId *string `json:"MainUserId,omitempty" xml:"MainUserId,omitempty"`
	// The status of the alert data. Valid values:
	//
	// 	- test: business test data.
	//
	// 	- online: online data.
	//
	// example:
	//
	// test
	OnlineStatus *string `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// The ID of the Alibaba Cloud account within which the alert is generated.
	//
	// example:
	//
	// 176555323***
	SubUserId *string `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
	// The UUID of the alert.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	Uuid       *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	VerifyType *string `json:"VerifyType,omitempty" xml:"VerifyType,omitempty"`
}

func (s ListCustomizeRuleTestResultResponseBodyDataResponseData) String() string {
	return dara.Prettify(s)
}

func (s ListCustomizeRuleTestResultResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetAlertDesc() *string {
	return s.AlertDesc
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetAlertDetail() *string {
	return s.AlertDetail
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetAlertSrcProd() *string {
	return s.AlertSrcProd
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetAlertSrcProdModule() *string {
	return s.AlertSrcProdModule
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetAttCk() *string {
	return s.AttCk
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetEventName() *string {
	return s.EventName
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetEventType() *string {
	return s.EventType
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetLevel() *string {
	return s.Level
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetLogSource() *string {
	return s.LogSource
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetLogTime() *string {
	return s.LogTime
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetLogType() *string {
	return s.LogType
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetMainUserId() *string {
	return s.MainUserId
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetOnlineStatus() *string {
	return s.OnlineStatus
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetSubUserId() *string {
	return s.SubUserId
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetUuid() *string {
	return s.Uuid
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) GetVerifyType() *string {
	return s.VerifyType
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetAlertDesc(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.AlertDesc = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetAlertDetail(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.AlertDetail = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetAlertSrcProd(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.AlertSrcProd = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetAlertSrcProdModule(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.AlertSrcProdModule = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetAttCk(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.AttCk = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetEventName(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.EventName = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetEventType(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.EventType = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetLevel(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.Level = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetLogSource(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.LogSource = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetLogTime(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.LogTime = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetLogType(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.LogType = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetMainUserId(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.MainUserId = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetOnlineStatus(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.OnlineStatus = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetSubUserId(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetUuid(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.Uuid = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) SetVerifyType(v string) *ListCustomizeRuleTestResultResponseBodyDataResponseData {
	s.VerifyType = &v
	return s
}

func (s *ListCustomizeRuleTestResultResponseBodyDataResponseData) Validate() error {
	return dara.Validate(s)
}
