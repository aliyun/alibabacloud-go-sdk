// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudSiemCustomizeRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListCloudSiemCustomizeRulesResponseBody
	GetCode() *int32
	SetData(v *ListCloudSiemCustomizeRulesResponseBodyData) *ListCloudSiemCustomizeRulesResponseBody
	GetData() *ListCloudSiemCustomizeRulesResponseBodyData
	SetMessage(v string) *ListCloudSiemCustomizeRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCloudSiemCustomizeRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCloudSiemCustomizeRulesResponseBody
	GetSuccess() *bool
}

type ListCloudSiemCustomizeRulesResponseBody struct {
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
	Data *ListCloudSiemCustomizeRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCloudSiemCustomizeRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListCloudSiemCustomizeRulesResponseBody) GetData() *ListCloudSiemCustomizeRulesResponseBodyData {
	return s.Data
}

func (s *ListCloudSiemCustomizeRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCloudSiemCustomizeRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudSiemCustomizeRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCloudSiemCustomizeRulesResponseBody) SetCode(v int32) *ListCloudSiemCustomizeRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBody) SetData(v *ListCloudSiemCustomizeRulesResponseBodyData) *ListCloudSiemCustomizeRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBody) SetMessage(v string) *ListCloudSiemCustomizeRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBody) SetRequestId(v string) *ListCloudSiemCustomizeRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBody) SetSuccess(v bool) *ListCloudSiemCustomizeRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudSiemCustomizeRulesResponseBodyData struct {
	// The pagination information.
	PageInfo *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*ListCloudSiemCustomizeRulesResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s ListCloudSiemCustomizeRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesResponseBodyData) GetPageInfo() *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *ListCloudSiemCustomizeRulesResponseBodyData) GetResponseData() []*ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	return s.ResponseData
}

func (s *ListCloudSiemCustomizeRulesResponseBodyData) SetPageInfo(v *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) *ListCloudSiemCustomizeRulesResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyData) SetResponseData(v []*ListCloudSiemCustomizeRulesResponseBodyDataResponseData) *ListCloudSiemCustomizeRulesResponseBodyData {
	s.ResponseData = v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyData) Validate() error {
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

type ListCloudSiemCustomizeRulesResponseBodyDataPageInfo struct {
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

func (s ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) SetCurrentPage(v int32) *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) SetPageSize(v int32) *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) SetTotalCount(v int64) *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListCloudSiemCustomizeRulesResponseBodyDataResponseData struct {
	// The type of the risk.
	//
	// example:
	//
	// WEBSHELL
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The internal code of the risk type.
	//
	// example:
	//
	// ${siem_rule_type_process_abnormal_command}
	AlertTypeMds *string `json:"AlertTypeMds,omitempty" xml:"AlertTypeMds,omitempty"`
	// The ID of the Alibaba Cloud account in SIEM.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The alert additional field for ATT\\&CK.
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	// The type of the view. Valid values:
	//
	// 0: view of the current Alibaba Cloud account. 1: view of all accounts for the enterprise.
	//
	// example:
	//
	// 1
	DataType *int32 `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The extended information about event generation. If the value of **eventTransferType*	- is **allToSingle**, the value of this parameter indicates the length and unit of the alert aggregation window. The HTML escape characters are reversed.
	//
	// example:
	//
	// {&quot;time&quot;:&quot;1&quot;,&quot;unit&quot;:&quot;MINUTE&quot;}
	EventTransferExt *string `json:"EventTransferExt,omitempty" xml:"EventTransferExt,omitempty"`
	// Indicates whether the system generates an event for the alert. Valid values:
	//
	// 	- **0**: no.
	//
	// 	- **1**: yes.
	//
	// example:
	//
	// 1
	EventTransferSwitch *int32 `json:"EventTransferSwitch,omitempty" xml:"EventTransferSwitch,omitempty"`
	// The method that is used to generate an event. Valid values:
	//
	// 	- **default**: built-in method.
	//
	// 	- **singleToSingle**: The system generates an event for each alert.
	//
	// 	- **allToSingle**: The system generates an event for alerts within a period of time.
	//
	// example:
	//
	// allToSingle
	EventTransferType *string `json:"EventTransferType,omitempty" xml:"EventTransferType,omitempty"`
	// The time when the custom rule was created.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the custom rule was last updated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the custom rule.
	//
	// example:
	//
	// 123456789
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The log source of the rule.
	//
	// example:
	//
	// cloud_siem_aegis_sas_alert
	LogSource *string `json:"LogSource,omitempty" xml:"LogSource,omitempty"`
	// The internal code of the log source.
	//
	// example:
	//
	// ${sas.cloudsiem.prod.cloud_siem_aegis_sas_alert}
	LogSourceMds *string `json:"LogSourceMds,omitempty" xml:"LogSourceMds,omitempty"`
	// The log type of the rule.
	//
	// example:
	//
	// ALERT_ACTIVITY
	LogType *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	// The internal code of the log type.
	//
	// example:
	//
	// ${sas.cloudsiem.prod.alert_activity}
	LogTypeMds *string `json:"LogTypeMds,omitempty" xml:"LogTypeMds,omitempty"`
	// The window length of the rule. The HTML escape characters are reversed.
	//
	// example:
	//
	// {&quot;time&quot;:&quot;1&quot;,&quot;unit&quot;:&quot;HOUR&quot;}
	QueryCycle *string `json:"QueryCycle,omitempty" xml:"QueryCycle,omitempty"`
	// The query condition of the rule. The value is in the JSON format. The HTML escape characters are reversed.
	//
	// example:
	//
	// [[{&quot;not&quot;:false,&quot;left&quot;:&quot;alert_name&quot;,&quot;operator&quot;:&quot;=&quot;,&quot;right&quot;:&quot;WEBSHELL&quot;}]]
	RuleCondition *string `json:"RuleCondition,omitempty" xml:"RuleCondition,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// this rule is for waf scan
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// The log aggregation field. The value is in the JSON format. The HTML escape characters are reversed.
	//
	// example:
	//
	// [&quot;asset_id&quot;]
	RuleGroup *string `json:"RuleGroup,omitempty" xml:"RuleGroup,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// waf_scan
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The threshold configurations of the rule in the JSON format. The HTML escape characters are reversed.
	//
	// example:
	//
	// {&quot;aggregateFunction&quot;:&quot;count&quot;,&quot;aggregateFunctionName&quot;:&quot;count&quot;,&quot;field&quot;:&quot;activity_name&quot;,&quot;operator&quot;:&quot;&lt;=&quot;,&quot;value&quot;:1}
	RuleThreshold *string `json:"RuleThreshold,omitempty" xml:"RuleThreshold,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **predefine**
	//
	// 	- **customize**
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **0**: The rule is in the initial state.
	//
	// 	- **10**: The simulation data is tested.
	//
	// 	- **15**: The business data is being tested.
	//
	// 	- **20**: The business data test is complete.
	//
	// 	- **100**: The rule is in effect.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The risk level. Valid values:
	//
	// 	- **serious**: high-risk.
	//
	// 	- **suspicious**: medium-risk.
	//
	// 	- **remind**: low-risk.
	//
	// example:
	//
	// remind
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s ListCloudSiemCustomizeRulesResponseBodyDataResponseData) String() string {
	return dara.Prettify(s)
}

func (s ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetAlertType() *string {
	return s.AlertType
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetAlertTypeMds() *string {
	return s.AlertTypeMds
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetAliuid() *int64 {
	return s.Aliuid
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetAttCk() *string {
	return s.AttCk
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetDataType() *int32 {
	return s.DataType
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetEventTransferExt() *string {
	return s.EventTransferExt
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetEventTransferSwitch() *int32 {
	return s.EventTransferSwitch
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetEventTransferType() *string {
	return s.EventTransferType
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetId() *int64 {
	return s.Id
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetLogSource() *string {
	return s.LogSource
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetLogSourceMds() *string {
	return s.LogSourceMds
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetLogType() *string {
	return s.LogType
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetLogTypeMds() *string {
	return s.LogTypeMds
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetQueryCycle() *string {
	return s.QueryCycle
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetRuleCondition() *string {
	return s.RuleCondition
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetRuleGroup() *string {
	return s.RuleGroup
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetRuleName() *string {
	return s.RuleName
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetRuleThreshold() *string {
	return s.RuleThreshold
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetRuleType() *string {
	return s.RuleType
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetStatus() *int32 {
	return s.Status
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetAlertType(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.AlertType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetAlertTypeMds(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.AlertTypeMds = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetAliuid(v int64) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetAttCk(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.AttCk = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetDataType(v int32) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.DataType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetEventTransferExt(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.EventTransferExt = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetEventTransferSwitch(v int32) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.EventTransferSwitch = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetEventTransferType(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.EventTransferType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetGmtCreate(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetGmtModified(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetId(v int64) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetLogSource(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.LogSource = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetLogSourceMds(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.LogSourceMds = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetLogType(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.LogType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetLogTypeMds(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.LogTypeMds = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetQueryCycle(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.QueryCycle = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleCondition(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleCondition = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleDesc(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleDesc = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleGroup(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleGroup = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleName(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleName = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleThreshold(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleThreshold = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetRuleType(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.RuleType = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetStatus(v int32) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.Status = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) SetThreatLevel(v string) *ListCloudSiemCustomizeRulesResponseBodyDataResponseData {
	s.ThreatLevel = &v
	return s
}

func (s *ListCloudSiemCustomizeRulesResponseBodyDataResponseData) Validate() error {
	return dara.Validate(s)
}
