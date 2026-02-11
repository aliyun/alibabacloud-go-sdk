// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostCustomizeRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *PostCustomizeRuleResponseBody
	GetCode() *int32
	SetData(v *PostCustomizeRuleResponseBodyData) *PostCustomizeRuleResponseBody
	GetData() *PostCustomizeRuleResponseBodyData
	SetMessage(v string) *PostCustomizeRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *PostCustomizeRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PostCustomizeRuleResponseBody
	GetSuccess() *bool
}

type PostCustomizeRuleResponseBody struct {
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
	Data *PostCustomizeRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s PostCustomizeRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PostCustomizeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *PostCustomizeRuleResponseBody) GetData() *PostCustomizeRuleResponseBodyData {
	return s.Data
}

func (s *PostCustomizeRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PostCustomizeRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PostCustomizeRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PostCustomizeRuleResponseBody) SetCode(v int32) *PostCustomizeRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PostCustomizeRuleResponseBody) SetData(v *PostCustomizeRuleResponseBodyData) *PostCustomizeRuleResponseBody {
	s.Data = v
	return s
}

func (s *PostCustomizeRuleResponseBody) SetMessage(v string) *PostCustomizeRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PostCustomizeRuleResponseBody) SetRequestId(v string) *PostCustomizeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PostCustomizeRuleResponseBody) SetSuccess(v bool) *PostCustomizeRuleResponseBody {
	s.Success = &v
	return s
}

func (s *PostCustomizeRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PostCustomizeRuleResponseBodyData struct {
	// The risk type.
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
	// The ID of the Alibaba Cloud account that is used to purchase the threat analysis feature.
	//
	// example:
	//
	// 127608589417****
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// 告警附加字段attck
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	// 自动化响应规则条件字段数据类型。
	//
	// example:
	//
	// varchar
	DataType *int32 `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The extended information about event generation. If eventTransferType is set to allToSingle, the value of this parameter indicates the length and unit of the alert aggregation window. The HTML escape characters are reversed.
	//
	// example:
	//
	// {&quot;time&quot;:&quot;1&quot;,&quot;unit&quot;:&quot;MINUTE&quot;}
	EventTransferExt *string `json:"EventTransferExt,omitempty" xml:"EventTransferExt,omitempty"`
	// Indicates whether the system generates an event for the alert. Valid values:
	//
	// 	- 0: no
	//
	// 	- 1: yes
	//
	// example:
	//
	// 1
	EventTransferSwitch *int32 `json:"EventTransferSwitch,omitempty" xml:"EventTransferSwitch,omitempty"`
	// The event generation method. Valid values:
	//
	// 	- default: The default method is used.
	//
	// 	- singleToSingle: The system generates an event for each alert.
	//
	// 	- allToSingle: The system generates an event for alerts within a period of time.
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
	// ${security_event_config.event_name.webshellName_clientav}
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
	// The log aggregation field of the rule. The value is a JSON string. The HTML escape characters are reversed.
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
	// The threshold configuration of the rule. The value is in the JSON format. The HTML escape characters are reversed.
	//
	// example:
	//
	// {&quot;aggregateFunction&quot;:&quot;count&quot;,&quot;aggregateFunctionName&quot;:&quot;count&quot;,&quot;field&quot;:&quot;activity_name&quot;,&quot;operator&quot;:&quot;&lt;=&quot;,&quot;value&quot;:1}
	RuleThreshold *string `json:"RuleThreshold,omitempty" xml:"RuleThreshold,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- predefine
	//
	// 	- customize
	//
	// example:
	//
	// customize
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The rule status. Valid values:
	//
	// 	- 0: The rule is in the initial state.
	//
	// 	- 10: The simulation data is tested.
	//
	// 	- 15: The business data is being tested.
	//
	// 	- 20: The business data test ends.
	//
	// 	- 100: The rule takes effect.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
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
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s PostCustomizeRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PostCustomizeRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleResponseBodyData) GetAlertType() *string {
	return s.AlertType
}

func (s *PostCustomizeRuleResponseBodyData) GetAlertTypeMds() *string {
	return s.AlertTypeMds
}

func (s *PostCustomizeRuleResponseBodyData) GetAliuid() *int64 {
	return s.Aliuid
}

func (s *PostCustomizeRuleResponseBodyData) GetAttCk() *string {
	return s.AttCk
}

func (s *PostCustomizeRuleResponseBodyData) GetDataType() *int32 {
	return s.DataType
}

func (s *PostCustomizeRuleResponseBodyData) GetEventTransferExt() *string {
	return s.EventTransferExt
}

func (s *PostCustomizeRuleResponseBodyData) GetEventTransferSwitch() *int32 {
	return s.EventTransferSwitch
}

func (s *PostCustomizeRuleResponseBodyData) GetEventTransferType() *string {
	return s.EventTransferType
}

func (s *PostCustomizeRuleResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *PostCustomizeRuleResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *PostCustomizeRuleResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *PostCustomizeRuleResponseBodyData) GetLogSource() *string {
	return s.LogSource
}

func (s *PostCustomizeRuleResponseBodyData) GetLogSourceMds() *string {
	return s.LogSourceMds
}

func (s *PostCustomizeRuleResponseBodyData) GetLogType() *string {
	return s.LogType
}

func (s *PostCustomizeRuleResponseBodyData) GetLogTypeMds() *string {
	return s.LogTypeMds
}

func (s *PostCustomizeRuleResponseBodyData) GetQueryCycle() *string {
	return s.QueryCycle
}

func (s *PostCustomizeRuleResponseBodyData) GetRuleCondition() *string {
	return s.RuleCondition
}

func (s *PostCustomizeRuleResponseBodyData) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *PostCustomizeRuleResponseBodyData) GetRuleGroup() *string {
	return s.RuleGroup
}

func (s *PostCustomizeRuleResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *PostCustomizeRuleResponseBodyData) GetRuleThreshold() *string {
	return s.RuleThreshold
}

func (s *PostCustomizeRuleResponseBodyData) GetRuleType() *string {
	return s.RuleType
}

func (s *PostCustomizeRuleResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *PostCustomizeRuleResponseBodyData) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *PostCustomizeRuleResponseBodyData) SetAlertType(v string) *PostCustomizeRuleResponseBodyData {
	s.AlertType = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetAlertTypeMds(v string) *PostCustomizeRuleResponseBodyData {
	s.AlertTypeMds = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetAliuid(v int64) *PostCustomizeRuleResponseBodyData {
	s.Aliuid = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetAttCk(v string) *PostCustomizeRuleResponseBodyData {
	s.AttCk = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetDataType(v int32) *PostCustomizeRuleResponseBodyData {
	s.DataType = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetEventTransferExt(v string) *PostCustomizeRuleResponseBodyData {
	s.EventTransferExt = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetEventTransferSwitch(v int32) *PostCustomizeRuleResponseBodyData {
	s.EventTransferSwitch = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetEventTransferType(v string) *PostCustomizeRuleResponseBodyData {
	s.EventTransferType = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetGmtCreate(v string) *PostCustomizeRuleResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetGmtModified(v string) *PostCustomizeRuleResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetId(v int64) *PostCustomizeRuleResponseBodyData {
	s.Id = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetLogSource(v string) *PostCustomizeRuleResponseBodyData {
	s.LogSource = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetLogSourceMds(v string) *PostCustomizeRuleResponseBodyData {
	s.LogSourceMds = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetLogType(v string) *PostCustomizeRuleResponseBodyData {
	s.LogType = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetLogTypeMds(v string) *PostCustomizeRuleResponseBodyData {
	s.LogTypeMds = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetQueryCycle(v string) *PostCustomizeRuleResponseBodyData {
	s.QueryCycle = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleCondition(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleCondition = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleDesc(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleDesc = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleGroup(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleGroup = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleName(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleThreshold(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleThreshold = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetRuleType(v string) *PostCustomizeRuleResponseBodyData {
	s.RuleType = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetStatus(v int32) *PostCustomizeRuleResponseBodyData {
	s.Status = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) SetThreatLevel(v string) *PostCustomizeRuleResponseBodyData {
	s.ThreatLevel = &v
	return s
}

func (s *PostCustomizeRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
