// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDetectionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertAttCk(v string) *CreateDetectionRuleRequest
	GetAlertAttCk() *string
	SetAlertAttCkMapping(v string) *CreateDetectionRuleRequest
	GetAlertAttCkMapping() *string
	SetAlertDescription(v string) *CreateDetectionRuleRequest
	GetAlertDescription() *string
	SetAlertLevel(v string) *CreateDetectionRuleRequest
	GetAlertLevel() *string
	SetAlertLevelMapping(v string) *CreateDetectionRuleRequest
	GetAlertLevelMapping() *string
	SetAlertName(v string) *CreateDetectionRuleRequest
	GetAlertName() *string
	SetAlertSchemaId(v string) *CreateDetectionRuleRequest
	GetAlertSchemaId() *string
	SetAlertTacticId(v string) *CreateDetectionRuleRequest
	GetAlertTacticId() *string
	SetAlertThresholdCount(v int32) *CreateDetectionRuleRequest
	GetAlertThresholdCount() *int32
	SetAlertThresholdGroup(v string) *CreateDetectionRuleRequest
	GetAlertThresholdGroup() *string
	SetAlertThresholdPeriod(v string) *CreateDetectionRuleRequest
	GetAlertThresholdPeriod() *string
	SetAlertType(v string) *CreateDetectionRuleRequest
	GetAlertType() *string
	SetAlertTypeMapping(v string) *CreateDetectionRuleRequest
	GetAlertTypeMapping() *string
	SetDetectionExpressionContent(v string) *CreateDetectionRuleRequest
	GetDetectionExpressionContent() *string
	SetDetectionExpressionType(v string) *CreateDetectionRuleRequest
	GetDetectionExpressionType() *string
	SetDetectionRuleDescription(v string) *CreateDetectionRuleRequest
	GetDetectionRuleDescription() *string
	SetDetectionRuleName(v string) *CreateDetectionRuleRequest
	GetDetectionRuleName() *string
	SetDetectionRuleStatus(v string) *CreateDetectionRuleRequest
	GetDetectionRuleStatus() *string
	SetDetectionRuleTemplateId(v string) *CreateDetectionRuleRequest
	GetDetectionRuleTemplateId() *string
	SetDetectionRuleTemplateVersion(v string) *CreateDetectionRuleRequest
	GetDetectionRuleTemplateVersion() *string
	SetDetectionRuleType(v string) *CreateDetectionRuleRequest
	GetDetectionRuleType() *string
	SetDetectionRules(v string) *CreateDetectionRuleRequest
	GetDetectionRules() *string
	SetEntityMappings(v string) *CreateDetectionRuleRequest
	GetEntityMappings() *string
	SetIncidentAggregationExpression(v string) *CreateDetectionRuleRequest
	GetIncidentAggregationExpression() *string
	SetIncidentAggregationType(v string) *CreateDetectionRuleRequest
	GetIncidentAggregationType() *string
	SetLang(v string) *CreateDetectionRuleRequest
	GetLang() *string
	SetLogCategoryId(v string) *CreateDetectionRuleRequest
	GetLogCategoryId() *string
	SetLogSchemaId(v string) *CreateDetectionRuleRequest
	GetLogSchemaId() *string
	SetPlaybookParameters(v string) *CreateDetectionRuleRequest
	GetPlaybookParameters() *string
	SetPlaybookUuid(v string) *CreateDetectionRuleRequest
	GetPlaybookUuid() *string
	SetRegionId(v string) *CreateDetectionRuleRequest
	GetRegionId() *string
	SetRoleFor(v int64) *CreateDetectionRuleRequest
	GetRoleFor() *int64
	SetScheduleBeginTime(v int64) *CreateDetectionRuleRequest
	GetScheduleBeginTime() *int64
	SetScheduleExpression(v string) *CreateDetectionRuleRequest
	GetScheduleExpression() *string
	SetScheduleMaxRetries(v int32) *CreateDetectionRuleRequest
	GetScheduleMaxRetries() *int32
	SetScheduleMaxTimeout(v int32) *CreateDetectionRuleRequest
	GetScheduleMaxTimeout() *int32
	SetScheduleType(v string) *CreateDetectionRuleRequest
	GetScheduleType() *string
	SetScheduleWindow(v string) *CreateDetectionRuleRequest
	GetScheduleWindow() *string
}

type CreateDetectionRuleRequest struct {
	// The alert ATT&CK technique.
	//
	// example:
	//
	// Discovery
	AlertAttCk        *string `json:"AlertAttCk,omitempty" xml:"AlertAttCk,omitempty"`
	AlertAttCkMapping *string `json:"AlertAttCkMapping,omitempty" xml:"AlertAttCkMapping,omitempty"`
	// The alert description. You can use $$ to reference query output fields.
	//
	// example:
	//
	// Alert from: $product_code$, detected network attack from $src_ip$, affected assets include: $dst_ip$
	AlertDescription *string `json:"AlertDescription,omitempty" xml:"AlertDescription,omitempty"`
	// The threat level of the alert. Valid values:
	//
	// - 5: Critical.
	//
	// - 4: High.
	//
	// - 3: Medium.
	//
	// - 2: Low.
	//
	// - 1: Informational.
	//
	// example:
	//
	// 1
	AlertLevel        *string `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	AlertLevelMapping *string `json:"AlertLevelMapping,omitempty" xml:"AlertLevelMapping,omitempty"`
	// The alert name. You can use $$ to reference query output fields.
	//
	// example:
	//
	// Detected high-frequency multi-type network attacks from $src_ip$
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The ID of the detection rule alert template. Valid values:
	//
	// - ALERT_ACTIVITY: other alerts.
	//
	// - EDR_ALERT_ACTIVITY: endpoint detection and response alerts.
	//
	// - FIREWALL_ALERT_ACTIVITY: firewall alerts.
	//
	// - WAF_ALERT_ACTIVITY: web application firewall alerts.
	//
	// example:
	//
	// ALERT_ACTIVITY
	AlertSchemaId *string `json:"AlertSchemaId,omitempty" xml:"AlertSchemaId,omitempty"`
	// The alert tactic stage.
	//
	// example:
	//
	// TA0042
	AlertTacticId *string `json:"AlertTacticId,omitempty" xml:"AlertTacticId,omitempty"`
	// The alert threshold count.
	//
	// example:
	//
	// 10
	AlertThresholdCount *int32 `json:"AlertThresholdCount,omitempty" xml:"AlertThresholdCount,omitempty"`
	// The list of alert threshold fields, separated by commas (,).
	//
	// example:
	//
	// alert_type,ip
	AlertThresholdGroup *string `json:"AlertThresholdGroup,omitempty" xml:"AlertThresholdGroup,omitempty"`
	// The length of the alert threshold period.
	//
	// example:
	//
	// 5m
	AlertThresholdPeriod *string `json:"AlertThresholdPeriod,omitempty" xml:"AlertThresholdPeriod,omitempty"`
	// The Alarm Metric of the alerting rule.
	//
	// example:
	//
	// WebShell
	AlertType        *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AlertTypeMapping *string `json:"AlertTypeMapping,omitempty" xml:"AlertTypeMapping,omitempty"`
	// The content of the detection rule expression.
	//
	// example:
	//
	// *|set session mode=scan;SELECT 	- FROM log
	//
	// WHERE schema = \\"PROCESS_START_ACTIVITY\\"
	//
	// AND (
	//
	//     proc_path LIKE \\"%/groups\\"
	//
	//     OR (
	//
	//         (
	//
	//             proc_path LIKE \\"%/cat\\"
	//
	//             OR proc_path LIKE \\"%/head\\"
	//
	//             OR proc_path LIKE \\"%/tail\\"
	//
	//             OR proc_path LIKE \\"%/more\\"
	//
	//         )
	//
	//         AND cmdline LIKE \\"%/etc/group%\\"
	//
	//     )
	//
	// )
	DetectionExpressionContent *string `json:"DetectionExpressionContent,omitempty" xml:"DetectionExpressionContent,omitempty"`
	// The type of the detection rule expression. Valid values:
	//
	// - sql: SQL.
	//
	// - playbook: playbook.
	//
	// example:
	//
	// sql
	DetectionExpressionType *string `json:"DetectionExpressionType,omitempty" xml:"DetectionExpressionType,omitempty"`
	// The description of the detection rule.
	//
	// example:
	//
	// dr-123
	DetectionRuleDescription *string `json:"DetectionRuleDescription,omitempty" xml:"DetectionRuleDescription,omitempty"`
	// The name of the detection rule.
	//
	// example:
	//
	// dr-ha1i09ob3zmqrs85****
	DetectionRuleName *string `json:"DetectionRuleName,omitempty" xml:"DetectionRuleName,omitempty"`
	// The status of the detection rule.
	//
	// example:
	//
	// 0
	DetectionRuleStatus *string `json:"DetectionRuleStatus,omitempty" xml:"DetectionRuleStatus,omitempty"`
	// The ID of the detection rule template.
	//
	// example:
	//
	// cfw-out-ip_aegis-netstat
	DetectionRuleTemplateId *string `json:"DetectionRuleTemplateId,omitempty" xml:"DetectionRuleTemplateId,omitempty"`
	// The version of the detection rule template.
	//
	// example:
	//
	// v1.0.0
	DetectionRuleTemplateVersion *string `json:"DetectionRuleTemplateVersion,omitempty" xml:"DetectionRuleTemplateVersion,omitempty"`
	// The type of the detection rule. Valid values:
	//
	// - preset: preset detection rule.
	//
	// - custom: custom detection rule.
	//
	// - custom_template: rule template.
	//
	// example:
	//
	// custom
	DetectionRuleType *string `json:"DetectionRuleType,omitempty" xml:"DetectionRuleType,omitempty"`
	DetectionRules    *string `json:"DetectionRules,omitempty" xml:"DetectionRules,omitempty"`
	// The entity mapping configuration.
	//
	// example:
	//
	// [{\\"NormalizationSchemaId\\":\\"host\\",\\"NormalizationFieldMappings\\":[{\\"NormalizationFieldName\\":\\"uuid\\",\\"MappingFieldName\\":\\"host\\",\\"NormalizationFieldType\\":\\"varchar\\"}]}]
	EntityMappings *string `json:"EntityMappings,omitempty" xml:"EntityMappings,omitempty"`
	// The event aggregation period configuration.
	//
	// example:
	//
	// 5m
	IncidentAggregationExpression *string `json:"IncidentAggregationExpression,omitempty" xml:"IncidentAggregationExpression,omitempty"`
	// The event aggregation type. Valid values:
	//
	// - none: No event is generated.
	//
	// - graph_compute: Graph computing. This value is supported by predefined rules.
	//
	// - expert: Expert rule.
	//
	// - passthrough: Alerting pass-through (one-to-one).
	//
	// - window: Same-type aggregation (window).
	//
	// example:
	//
	// window
	IncidentAggregationType *string `json:"IncidentAggregationType,omitempty" xml:"IncidentAggregationType,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the log normalization category.
	//
	// example:
	//
	// NETWORK_CATEGORY
	LogCategoryId *string `json:"LogCategoryId,omitempty" xml:"LogCategoryId,omitempty"`
	// The ID of the log normalization schema.
	//
	// example:
	//
	// API_RISK_ACTIVITY
	LogSchemaId *string `json:"LogSchemaId,omitempty" xml:"LogSchemaId,omitempty"`
	// The custom parameters of the playbook.
	//
	// example:
	//
	// {
	//
	//     "ip": {
	//
	//         "ip": "124.23.*.*"
	//
	//     }
	//
	// }
	PlaybookParameters *string `json:"PlaybookParameters,omitempty" xml:"PlaybookParameters,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// system_aliyun_clb_process_book
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Your assets reside in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets reside outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member to which the administrator switches the view.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// The scheduling start time. The value is a 13-digit UNIX timestamp.
	//
	// example:
	//
	// 1733269771123
	ScheduleBeginTime *int64 `json:"ScheduleBeginTime,omitempty" xml:"ScheduleBeginTime,omitempty"`
	// The scheduling cron expression. This parameter is required when ScheduleType is set to cron.
	//
	// example:
	//
	// 0/5 	- 	- 	- *
	ScheduleExpression *string `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	// The maximum number of retries upon timeout. Valid values: 1 to 100.
	//
	// example:
	//
	// 1
	ScheduleMaxRetries *int32 `json:"ScheduleMaxRetries,omitempty" xml:"ScheduleMaxRetries,omitempty"`
	// The maximum timeout period, in seconds. Valid values: 60 to 1800.
	//
	// example:
	//
	// 60
	ScheduleMaxTimeout *int32 `json:"ScheduleMaxTimeout,omitempty" xml:"ScheduleMaxTimeout,omitempty"`
	// The scheduling type. Valid values:
	//
	// - fixed_rate: fixed interval.
	//
	// - cron: cron expression.
	//
	// example:
	//
	// fixed_rate
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
	// The length of the scheduling window.
	//
	// example:
	//
	// 5m
	ScheduleWindow *string `json:"ScheduleWindow,omitempty" xml:"ScheduleWindow,omitempty"`
}

func (s CreateDetectionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDetectionRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDetectionRuleRequest) GetAlertAttCk() *string {
	return s.AlertAttCk
}

func (s *CreateDetectionRuleRequest) GetAlertAttCkMapping() *string {
	return s.AlertAttCkMapping
}

func (s *CreateDetectionRuleRequest) GetAlertDescription() *string {
	return s.AlertDescription
}

func (s *CreateDetectionRuleRequest) GetAlertLevel() *string {
	return s.AlertLevel
}

func (s *CreateDetectionRuleRequest) GetAlertLevelMapping() *string {
	return s.AlertLevelMapping
}

func (s *CreateDetectionRuleRequest) GetAlertName() *string {
	return s.AlertName
}

func (s *CreateDetectionRuleRequest) GetAlertSchemaId() *string {
	return s.AlertSchemaId
}

func (s *CreateDetectionRuleRequest) GetAlertTacticId() *string {
	return s.AlertTacticId
}

func (s *CreateDetectionRuleRequest) GetAlertThresholdCount() *int32 {
	return s.AlertThresholdCount
}

func (s *CreateDetectionRuleRequest) GetAlertThresholdGroup() *string {
	return s.AlertThresholdGroup
}

func (s *CreateDetectionRuleRequest) GetAlertThresholdPeriod() *string {
	return s.AlertThresholdPeriod
}

func (s *CreateDetectionRuleRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *CreateDetectionRuleRequest) GetAlertTypeMapping() *string {
	return s.AlertTypeMapping
}

func (s *CreateDetectionRuleRequest) GetDetectionExpressionContent() *string {
	return s.DetectionExpressionContent
}

func (s *CreateDetectionRuleRequest) GetDetectionExpressionType() *string {
	return s.DetectionExpressionType
}

func (s *CreateDetectionRuleRequest) GetDetectionRuleDescription() *string {
	return s.DetectionRuleDescription
}

func (s *CreateDetectionRuleRequest) GetDetectionRuleName() *string {
	return s.DetectionRuleName
}

func (s *CreateDetectionRuleRequest) GetDetectionRuleStatus() *string {
	return s.DetectionRuleStatus
}

func (s *CreateDetectionRuleRequest) GetDetectionRuleTemplateId() *string {
	return s.DetectionRuleTemplateId
}

func (s *CreateDetectionRuleRequest) GetDetectionRuleTemplateVersion() *string {
	return s.DetectionRuleTemplateVersion
}

func (s *CreateDetectionRuleRequest) GetDetectionRuleType() *string {
	return s.DetectionRuleType
}

func (s *CreateDetectionRuleRequest) GetDetectionRules() *string {
	return s.DetectionRules
}

func (s *CreateDetectionRuleRequest) GetEntityMappings() *string {
	return s.EntityMappings
}

func (s *CreateDetectionRuleRequest) GetIncidentAggregationExpression() *string {
	return s.IncidentAggregationExpression
}

func (s *CreateDetectionRuleRequest) GetIncidentAggregationType() *string {
	return s.IncidentAggregationType
}

func (s *CreateDetectionRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateDetectionRuleRequest) GetLogCategoryId() *string {
	return s.LogCategoryId
}

func (s *CreateDetectionRuleRequest) GetLogSchemaId() *string {
	return s.LogSchemaId
}

func (s *CreateDetectionRuleRequest) GetPlaybookParameters() *string {
	return s.PlaybookParameters
}

func (s *CreateDetectionRuleRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *CreateDetectionRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDetectionRuleRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CreateDetectionRuleRequest) GetScheduleBeginTime() *int64 {
	return s.ScheduleBeginTime
}

func (s *CreateDetectionRuleRequest) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *CreateDetectionRuleRequest) GetScheduleMaxRetries() *int32 {
	return s.ScheduleMaxRetries
}

func (s *CreateDetectionRuleRequest) GetScheduleMaxTimeout() *int32 {
	return s.ScheduleMaxTimeout
}

func (s *CreateDetectionRuleRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *CreateDetectionRuleRequest) GetScheduleWindow() *string {
	return s.ScheduleWindow
}

func (s *CreateDetectionRuleRequest) SetAlertAttCk(v string) *CreateDetectionRuleRequest {
	s.AlertAttCk = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetAlertAttCkMapping(v string) *CreateDetectionRuleRequest {
	s.AlertAttCkMapping = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetAlertDescription(v string) *CreateDetectionRuleRequest {
	s.AlertDescription = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetAlertLevel(v string) *CreateDetectionRuleRequest {
	s.AlertLevel = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetAlertLevelMapping(v string) *CreateDetectionRuleRequest {
	s.AlertLevelMapping = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetAlertName(v string) *CreateDetectionRuleRequest {
	s.AlertName = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetAlertSchemaId(v string) *CreateDetectionRuleRequest {
	s.AlertSchemaId = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetAlertTacticId(v string) *CreateDetectionRuleRequest {
	s.AlertTacticId = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetAlertThresholdCount(v int32) *CreateDetectionRuleRequest {
	s.AlertThresholdCount = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetAlertThresholdGroup(v string) *CreateDetectionRuleRequest {
	s.AlertThresholdGroup = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetAlertThresholdPeriod(v string) *CreateDetectionRuleRequest {
	s.AlertThresholdPeriod = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetAlertType(v string) *CreateDetectionRuleRequest {
	s.AlertType = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetAlertTypeMapping(v string) *CreateDetectionRuleRequest {
	s.AlertTypeMapping = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetDetectionExpressionContent(v string) *CreateDetectionRuleRequest {
	s.DetectionExpressionContent = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetDetectionExpressionType(v string) *CreateDetectionRuleRequest {
	s.DetectionExpressionType = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetDetectionRuleDescription(v string) *CreateDetectionRuleRequest {
	s.DetectionRuleDescription = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetDetectionRuleName(v string) *CreateDetectionRuleRequest {
	s.DetectionRuleName = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetDetectionRuleStatus(v string) *CreateDetectionRuleRequest {
	s.DetectionRuleStatus = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetDetectionRuleTemplateId(v string) *CreateDetectionRuleRequest {
	s.DetectionRuleTemplateId = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetDetectionRuleTemplateVersion(v string) *CreateDetectionRuleRequest {
	s.DetectionRuleTemplateVersion = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetDetectionRuleType(v string) *CreateDetectionRuleRequest {
	s.DetectionRuleType = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetDetectionRules(v string) *CreateDetectionRuleRequest {
	s.DetectionRules = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetEntityMappings(v string) *CreateDetectionRuleRequest {
	s.EntityMappings = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetIncidentAggregationExpression(v string) *CreateDetectionRuleRequest {
	s.IncidentAggregationExpression = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetIncidentAggregationType(v string) *CreateDetectionRuleRequest {
	s.IncidentAggregationType = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetLang(v string) *CreateDetectionRuleRequest {
	s.Lang = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetLogCategoryId(v string) *CreateDetectionRuleRequest {
	s.LogCategoryId = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetLogSchemaId(v string) *CreateDetectionRuleRequest {
	s.LogSchemaId = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetPlaybookParameters(v string) *CreateDetectionRuleRequest {
	s.PlaybookParameters = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetPlaybookUuid(v string) *CreateDetectionRuleRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetRegionId(v string) *CreateDetectionRuleRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetRoleFor(v int64) *CreateDetectionRuleRequest {
	s.RoleFor = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetScheduleBeginTime(v int64) *CreateDetectionRuleRequest {
	s.ScheduleBeginTime = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetScheduleExpression(v string) *CreateDetectionRuleRequest {
	s.ScheduleExpression = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetScheduleMaxRetries(v int32) *CreateDetectionRuleRequest {
	s.ScheduleMaxRetries = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetScheduleMaxTimeout(v int32) *CreateDetectionRuleRequest {
	s.ScheduleMaxTimeout = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetScheduleType(v string) *CreateDetectionRuleRequest {
	s.ScheduleType = &v
	return s
}

func (s *CreateDetectionRuleRequest) SetScheduleWindow(v string) *CreateDetectionRuleRequest {
	s.ScheduleWindow = &v
	return s
}

func (s *CreateDetectionRuleRequest) Validate() error {
	return dara.Validate(s)
}
