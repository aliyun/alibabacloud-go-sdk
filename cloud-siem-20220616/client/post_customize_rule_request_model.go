// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostCustomizeRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertType(v string) *PostCustomizeRuleRequest
	GetAlertType() *string
	SetAlertTypeMds(v string) *PostCustomizeRuleRequest
	GetAlertTypeMds() *string
	SetAttCk(v string) *PostCustomizeRuleRequest
	GetAttCk() *string
	SetEventTransferExt(v string) *PostCustomizeRuleRequest
	GetEventTransferExt() *string
	SetEventTransferSwitch(v int32) *PostCustomizeRuleRequest
	GetEventTransferSwitch() *int32
	SetEventTransferType(v string) *PostCustomizeRuleRequest
	GetEventTransferType() *string
	SetId(v int64) *PostCustomizeRuleRequest
	GetId() *int64
	SetLogSource(v string) *PostCustomizeRuleRequest
	GetLogSource() *string
	SetLogSourceMds(v string) *PostCustomizeRuleRequest
	GetLogSourceMds() *string
	SetLogType(v string) *PostCustomizeRuleRequest
	GetLogType() *string
	SetLogTypeMds(v string) *PostCustomizeRuleRequest
	GetLogTypeMds() *string
	SetQueryCycle(v string) *PostCustomizeRuleRequest
	GetQueryCycle() *string
	SetRegionId(v string) *PostCustomizeRuleRequest
	GetRegionId() *string
	SetRoleFor(v int64) *PostCustomizeRuleRequest
	GetRoleFor() *int64
	SetRoleType(v int32) *PostCustomizeRuleRequest
	GetRoleType() *int32
	SetRuleCondition(v string) *PostCustomizeRuleRequest
	GetRuleCondition() *string
	SetRuleDesc(v string) *PostCustomizeRuleRequest
	GetRuleDesc() *string
	SetRuleGroup(v string) *PostCustomizeRuleRequest
	GetRuleGroup() *string
	SetRuleName(v string) *PostCustomizeRuleRequest
	GetRuleName() *string
	SetRuleThreshold(v string) *PostCustomizeRuleRequest
	GetRuleThreshold() *string
	SetThreatLevel(v string) *PostCustomizeRuleRequest
	GetThreatLevel() *string
}

type PostCustomizeRuleRequest struct {
	// The threat type.
	//
	// example:
	//
	// WEBSHELL
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The Medusa code of the threat type.
	//
	// example:
	//
	// ${siem_rule_type_process_abnormal_command}
	AlertTypeMds *string `json:"AlertTypeMds,omitempty" xml:"AlertTypeMds,omitempty"`
	// The ATT\\&CK technique.
	//
	// example:
	//
	// T1595.002 Vulnerability Scanning
	AttCk *string `json:"AttCk,omitempty" xml:"AttCk,omitempty"`
	// The extended information for event generation. This parameter is returned only when EventTransferType is set to allToSingle. The value indicates the length and unit of the alert aggregation window.
	//
	// example:
	//
	// {"time":"1","unit":"MINUTE"}
	EventTransferExt *string `json:"EventTransferExt,omitempty" xml:"EventTransferExt,omitempty"`
	// Specifies whether to convert alerts into events. Valid values:
	//
	// - 0: no
	//
	// - 1: yes
	//
	// example:
	//
	// 1
	EventTransferSwitch *int32 `json:"EventTransferSwitch,omitempty" xml:"EventTransferSwitch,omitempty"`
	// The event generation method. Valid values:
	//
	// - default: the default method
	//
	// - singleToSingle: An event is generated for each alert.
	//
	// - allToSingle: An event is generated for all alerts in a period.
	//
	// example:
	//
	// allToSingle
	EventTransferType *string `json:"EventTransferType,omitempty" xml:"EventTransferType,omitempty"`
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
	// The Medusa code of the log source.
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
	// The Medusa code of the log type.
	//
	// example:
	//
	// ${security_event_config.event_name.webshellName_clientav}
	LogTypeMds *string `json:"LogTypeMds,omitempty" xml:"LogTypeMds,omitempty"`
	// The length of the rule window.
	//
	// example:
	//
	// {"time":"1","unit":"HOUR"}
	QueryCycle *string `json:"QueryCycle,omitempty" xml:"QueryCycle,omitempty"`
	// The region where the data management center of Threat Analysis is located. Select the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: assets in the Chinese mainland and China (Hong Kong)
	//
	// - ap-southeast-1: assets outside China
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that is used to switch the administrator\\"s perspective to a member\\"s perspective.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The view type.
	//
	// - 0: the view of the current Alibaba Cloud account.
	//
	// - 1: the view of all accounts that are managed by the administrator.
	//
	// example:
	//
	// 1
	RoleType *int32 `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// The query condition of the rule. The value is a JSON string.
	//
	// example:
	//
	// [[{"not":false,"left":"alert_name","operator":"=","right":"WEBSHELL"}]]
	RuleCondition *string `json:"RuleCondition,omitempty" xml:"RuleCondition,omitempty"`
	// The description of the rule.
	//
	// example:
	//
	// this rule is for waf scan
	RuleDesc *string `json:"RuleDesc,omitempty" xml:"RuleDesc,omitempty"`
	// The fields that are used to group logs. The value is a JSON array.
	//
	// example:
	//
	// ["asset_id"]
	RuleGroup *string `json:"RuleGroup,omitempty" xml:"RuleGroup,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// waf_scan
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The threshold configuration of the rule. The value is a JSON string.
	//
	// example:
	//
	// {"aggregateFunction":"count","aggregateFunctionName":"count","field":"activity_name","operator":"&lt;=","value":1}
	RuleThreshold *string `json:"RuleThreshold,omitempty" xml:"RuleThreshold,omitempty"`
	// The threat level. Valid values:
	//
	// - serious: high
	//
	// - suspicious: medium
	//
	// - remind: low
	//
	// example:
	//
	// remind
	ThreatLevel *string `json:"ThreatLevel,omitempty" xml:"ThreatLevel,omitempty"`
}

func (s PostCustomizeRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s PostCustomizeRuleRequest) GoString() string {
	return s.String()
}

func (s *PostCustomizeRuleRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *PostCustomizeRuleRequest) GetAlertTypeMds() *string {
	return s.AlertTypeMds
}

func (s *PostCustomizeRuleRequest) GetAttCk() *string {
	return s.AttCk
}

func (s *PostCustomizeRuleRequest) GetEventTransferExt() *string {
	return s.EventTransferExt
}

func (s *PostCustomizeRuleRequest) GetEventTransferSwitch() *int32 {
	return s.EventTransferSwitch
}

func (s *PostCustomizeRuleRequest) GetEventTransferType() *string {
	return s.EventTransferType
}

func (s *PostCustomizeRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *PostCustomizeRuleRequest) GetLogSource() *string {
	return s.LogSource
}

func (s *PostCustomizeRuleRequest) GetLogSourceMds() *string {
	return s.LogSourceMds
}

func (s *PostCustomizeRuleRequest) GetLogType() *string {
	return s.LogType
}

func (s *PostCustomizeRuleRequest) GetLogTypeMds() *string {
	return s.LogTypeMds
}

func (s *PostCustomizeRuleRequest) GetQueryCycle() *string {
	return s.QueryCycle
}

func (s *PostCustomizeRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PostCustomizeRuleRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *PostCustomizeRuleRequest) GetRoleType() *int32 {
	return s.RoleType
}

func (s *PostCustomizeRuleRequest) GetRuleCondition() *string {
	return s.RuleCondition
}

func (s *PostCustomizeRuleRequest) GetRuleDesc() *string {
	return s.RuleDesc
}

func (s *PostCustomizeRuleRequest) GetRuleGroup() *string {
	return s.RuleGroup
}

func (s *PostCustomizeRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *PostCustomizeRuleRequest) GetRuleThreshold() *string {
	return s.RuleThreshold
}

func (s *PostCustomizeRuleRequest) GetThreatLevel() *string {
	return s.ThreatLevel
}

func (s *PostCustomizeRuleRequest) SetAlertType(v string) *PostCustomizeRuleRequest {
	s.AlertType = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetAlertTypeMds(v string) *PostCustomizeRuleRequest {
	s.AlertTypeMds = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetAttCk(v string) *PostCustomizeRuleRequest {
	s.AttCk = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetEventTransferExt(v string) *PostCustomizeRuleRequest {
	s.EventTransferExt = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetEventTransferSwitch(v int32) *PostCustomizeRuleRequest {
	s.EventTransferSwitch = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetEventTransferType(v string) *PostCustomizeRuleRequest {
	s.EventTransferType = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetId(v int64) *PostCustomizeRuleRequest {
	s.Id = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetLogSource(v string) *PostCustomizeRuleRequest {
	s.LogSource = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetLogSourceMds(v string) *PostCustomizeRuleRequest {
	s.LogSourceMds = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetLogType(v string) *PostCustomizeRuleRequest {
	s.LogType = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetLogTypeMds(v string) *PostCustomizeRuleRequest {
	s.LogTypeMds = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetQueryCycle(v string) *PostCustomizeRuleRequest {
	s.QueryCycle = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRegionId(v string) *PostCustomizeRuleRequest {
	s.RegionId = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRoleFor(v int64) *PostCustomizeRuleRequest {
	s.RoleFor = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRoleType(v int32) *PostCustomizeRuleRequest {
	s.RoleType = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRuleCondition(v string) *PostCustomizeRuleRequest {
	s.RuleCondition = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRuleDesc(v string) *PostCustomizeRuleRequest {
	s.RuleDesc = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRuleGroup(v string) *PostCustomizeRuleRequest {
	s.RuleGroup = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRuleName(v string) *PostCustomizeRuleRequest {
	s.RuleName = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetRuleThreshold(v string) *PostCustomizeRuleRequest {
	s.RuleThreshold = &v
	return s
}

func (s *PostCustomizeRuleRequest) SetThreatLevel(v string) *PostCustomizeRuleRequest {
	s.ThreatLevel = &v
	return s
}

func (s *PostCustomizeRuleRequest) Validate() error {
	return dara.Validate(s)
}
