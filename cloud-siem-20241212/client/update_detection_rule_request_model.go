// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDetectionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertAttCk(v string) *UpdateDetectionRuleRequest
	GetAlertAttCk() *string
	SetAlertAttCkMapping(v string) *UpdateDetectionRuleRequest
	GetAlertAttCkMapping() *string
	SetAlertDescription(v string) *UpdateDetectionRuleRequest
	GetAlertDescription() *string
	SetAlertLevel(v string) *UpdateDetectionRuleRequest
	GetAlertLevel() *string
	SetAlertLevelMapping(v string) *UpdateDetectionRuleRequest
	GetAlertLevelMapping() *string
	SetAlertName(v string) *UpdateDetectionRuleRequest
	GetAlertName() *string
	SetAlertSchemaId(v string) *UpdateDetectionRuleRequest
	GetAlertSchemaId() *string
	SetAlertTacticId(v string) *UpdateDetectionRuleRequest
	GetAlertTacticId() *string
	SetAlertThresholdCount(v int32) *UpdateDetectionRuleRequest
	GetAlertThresholdCount() *int32
	SetAlertThresholdGroup(v string) *UpdateDetectionRuleRequest
	GetAlertThresholdGroup() *string
	SetAlertThresholdPeriod(v string) *UpdateDetectionRuleRequest
	GetAlertThresholdPeriod() *string
	SetAlertType(v string) *UpdateDetectionRuleRequest
	GetAlertType() *string
	SetAlertTypeMapping(v string) *UpdateDetectionRuleRequest
	GetAlertTypeMapping() *string
	SetDetectionExpressionContent(v string) *UpdateDetectionRuleRequest
	GetDetectionExpressionContent() *string
	SetDetectionExpressionType(v string) *UpdateDetectionRuleRequest
	GetDetectionExpressionType() *string
	SetDetectionRuleDescription(v string) *UpdateDetectionRuleRequest
	GetDetectionRuleDescription() *string
	SetDetectionRuleId(v string) *UpdateDetectionRuleRequest
	GetDetectionRuleId() *string
	SetDetectionRuleName(v string) *UpdateDetectionRuleRequest
	GetDetectionRuleName() *string
	SetDetectionRuleStatus(v string) *UpdateDetectionRuleRequest
	GetDetectionRuleStatus() *string
	SetDetectionRuleType(v string) *UpdateDetectionRuleRequest
	GetDetectionRuleType() *string
	SetEntityMappings(v string) *UpdateDetectionRuleRequest
	GetEntityMappings() *string
	SetIncidentAggregationExpression(v string) *UpdateDetectionRuleRequest
	GetIncidentAggregationExpression() *string
	SetIncidentAggregationType(v string) *UpdateDetectionRuleRequest
	GetIncidentAggregationType() *string
	SetLang(v string) *UpdateDetectionRuleRequest
	GetLang() *string
	SetLogCategoryId(v string) *UpdateDetectionRuleRequest
	GetLogCategoryId() *string
	SetLogSchemaId(v string) *UpdateDetectionRuleRequest
	GetLogSchemaId() *string
	SetPlaybookParameters(v string) *UpdateDetectionRuleRequest
	GetPlaybookParameters() *string
	SetPlaybookUuid(v string) *UpdateDetectionRuleRequest
	GetPlaybookUuid() *string
	SetRegionId(v string) *UpdateDetectionRuleRequest
	GetRegionId() *string
	SetScheduleBeginTime(v int64) *UpdateDetectionRuleRequest
	GetScheduleBeginTime() *int64
	SetScheduleExpression(v string) *UpdateDetectionRuleRequest
	GetScheduleExpression() *string
	SetScheduleMaxRetries(v int32) *UpdateDetectionRuleRequest
	GetScheduleMaxRetries() *int32
	SetScheduleMaxTimeout(v int32) *UpdateDetectionRuleRequest
	GetScheduleMaxTimeout() *int32
	SetScheduleType(v string) *UpdateDetectionRuleRequest
	GetScheduleType() *string
	SetScheduleWindow(v string) *UpdateDetectionRuleRequest
	GetScheduleWindow() *string
}

type UpdateDetectionRuleRequest struct {
	// The ATT\\&CK tactic of the alert.
	//
	// example:
	//
	// Discovery
	AlertAttCk        *string `json:"AlertAttCk,omitempty" xml:"AlertAttCk,omitempty"`
	AlertAttCkMapping *string `json:"AlertAttCkMapping,omitempty" xml:"AlertAttCkMapping,omitempty"`
	// The description of the alert. You can use $$ to reference fields from the query output.
	//
	// example:
	//
	// Alert from: $product_code$, detected network attack from $src_ip$, affected assets include: $dst_ip$
	AlertDescription *string `json:"AlertDescription,omitempty" xml:"AlertDescription,omitempty"`
	// The threat level of the alert. Valid values:
	//
	// - 5: Critical.
	//
	// - 4: Important.
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
	// The name of the alert. You can use $$ to reference fields from the query output.
	//
	// example:
	//
	// Detected high-frequency multi-type network attacks from $src_ip$
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The ID of the alert template for the detection rule.
	//
	// example:
	//
	// ALERT_ACTIVITY
	AlertSchemaId *string `json:"AlertSchemaId,omitempty" xml:"AlertSchemaId,omitempty"`
	// The ATT\\&CK tactic ID of the alert.
	//
	// example:
	//
	// TA0042
	AlertTacticId *string `json:"AlertTacticId,omitempty" xml:"AlertTacticId,omitempty"`
	// The number of alerts for the alert threshold.
	//
	// example:
	//
	// 10
	AlertThresholdCount *int32 `json:"AlertThresholdCount,omitempty" xml:"AlertThresholdCount,omitempty"`
	// The fields for the alert threshold. Separate multiple fields with commas.
	//
	// example:
	//
	// alert_type,ip
	AlertThresholdGroup *string `json:"AlertThresholdGroup,omitempty" xml:"AlertThresholdGroup,omitempty"`
	// The period for the alert threshold.
	//
	// example:
	//
	// 5m
	AlertThresholdPeriod *string `json:"AlertThresholdPeriod,omitempty" xml:"AlertThresholdPeriod,omitempty"`
	// The type of the alert.
	//
	// example:
	//
	// siem_rule_type_alert_storm
	AlertType        *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AlertTypeMapping *string `json:"AlertTypeMapping,omitempty" xml:"AlertTypeMapping,omitempty"`
	// The content of the detection expression.
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
	// The type of the detection expression.
	//
	// example:
	//
	// sql
	DetectionExpressionType *string `json:"DetectionExpressionType,omitempty" xml:"DetectionExpressionType,omitempty"`
	// The description of the detection rule.
	//
	// example:
	//
	// Check the enumeration behavior of local system groups. An attacker may attempt to find the Local Systems group and its permission settings.
	DetectionRuleDescription *string `json:"DetectionRuleDescription,omitempty" xml:"DetectionRuleDescription,omitempty"`
	// The ID of the detection rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// jndi-attack-success_http_dns
	DetectionRuleId *string `json:"DetectionRuleId,omitempty" xml:"DetectionRuleId,omitempty"`
	// The name of the detection rule.
	//
	// example:
	//
	// CTDR Port Scan Behavior
	DetectionRuleName *string `json:"DetectionRuleName,omitempty" xml:"DetectionRuleName,omitempty"`
	// The status of the detection rule.
	//
	// example:
	//
	// enabled
	DetectionRuleStatus *string `json:"DetectionRuleStatus,omitempty" xml:"DetectionRuleStatus,omitempty"`
	// The type of the detection rule. Valid values:
	//
	// - preset: a preset detection rule.
	//
	// - custom: a custom detection rule.
	//
	// - custom_template: a rule template.
	//
	// This parameter is required.
	//
	// example:
	//
	// custom
	DetectionRuleType *string `json:"DetectionRuleType,omitempty" xml:"DetectionRuleType,omitempty"`
	// The entity mapping configuration.
	//
	// example:
	//
	// [{\\"NormalizationSchemaId\\":\\"host\\",\\"NormalizationFieldMappings\\":[{\\"NormalizationFieldName\\":\\"uuid\\",\\"MappingFieldName\\":\\"host\\",\\"NormalizationFieldType\\":\\"varchar\\"}]}]
	EntityMappings *string `json:"EntityMappings,omitempty" xml:"EntityMappings,omitempty"`
	// The configuration of the event aggregation period.
	//
	// example:
	//
	// 60m
	IncidentAggregationExpression *string `json:"IncidentAggregationExpression,omitempty" xml:"IncidentAggregationExpression,omitempty"`
	// The type of event aggregation. Valid values:
	//
	// - none: Events are not generated.
	//
	// - graph_compute: graph computing (supported by predefined rules).
	//
	// - expert: expert rule.
	//
	// - passthrough: pass-through (one-to-one mapping with alerts).
	//
	// - window: window-based aggregation of similar events.
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
	// The custom parameters for the playbook.
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
	// The unique ID of the playbook.
	//
	// example:
	//
	// 31568394-7a86-487c-b8ec-b3f42b59****
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The region of the Data Management center for Threat Analysis. Select the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: The Chinese mainland.
	//
	// - ap-southeast-1: Regions outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time for scheduling. This is a 13-digit UNIX timestamp.
	//
	// example:
	//
	// 1733269771123
	ScheduleBeginTime *int64 `json:"ScheduleBeginTime,omitempty" xml:"ScheduleBeginTime,omitempty"`
	// The cron expression for scheduling. This parameter is required when you set ScheduleType to cron.
	//
	// example:
	//
	// 1h
	ScheduleExpression *string `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	// The maximum number of retries after a timeout. The value must be between 1 and 100.
	//
	// example:
	//
	// 1
	ScheduleMaxRetries *int32 `json:"ScheduleMaxRetries,omitempty" xml:"ScheduleMaxRetries,omitempty"`
	// The maximum timeout period. Unit: seconds. The value must be between 60 and 1800.
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

func (s UpdateDetectionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDetectionRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateDetectionRuleRequest) GetAlertAttCk() *string {
	return s.AlertAttCk
}

func (s *UpdateDetectionRuleRequest) GetAlertAttCkMapping() *string {
	return s.AlertAttCkMapping
}

func (s *UpdateDetectionRuleRequest) GetAlertDescription() *string {
	return s.AlertDescription
}

func (s *UpdateDetectionRuleRequest) GetAlertLevel() *string {
	return s.AlertLevel
}

func (s *UpdateDetectionRuleRequest) GetAlertLevelMapping() *string {
	return s.AlertLevelMapping
}

func (s *UpdateDetectionRuleRequest) GetAlertName() *string {
	return s.AlertName
}

func (s *UpdateDetectionRuleRequest) GetAlertSchemaId() *string {
	return s.AlertSchemaId
}

func (s *UpdateDetectionRuleRequest) GetAlertTacticId() *string {
	return s.AlertTacticId
}

func (s *UpdateDetectionRuleRequest) GetAlertThresholdCount() *int32 {
	return s.AlertThresholdCount
}

func (s *UpdateDetectionRuleRequest) GetAlertThresholdGroup() *string {
	return s.AlertThresholdGroup
}

func (s *UpdateDetectionRuleRequest) GetAlertThresholdPeriod() *string {
	return s.AlertThresholdPeriod
}

func (s *UpdateDetectionRuleRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *UpdateDetectionRuleRequest) GetAlertTypeMapping() *string {
	return s.AlertTypeMapping
}

func (s *UpdateDetectionRuleRequest) GetDetectionExpressionContent() *string {
	return s.DetectionExpressionContent
}

func (s *UpdateDetectionRuleRequest) GetDetectionExpressionType() *string {
	return s.DetectionExpressionType
}

func (s *UpdateDetectionRuleRequest) GetDetectionRuleDescription() *string {
	return s.DetectionRuleDescription
}

func (s *UpdateDetectionRuleRequest) GetDetectionRuleId() *string {
	return s.DetectionRuleId
}

func (s *UpdateDetectionRuleRequest) GetDetectionRuleName() *string {
	return s.DetectionRuleName
}

func (s *UpdateDetectionRuleRequest) GetDetectionRuleStatus() *string {
	return s.DetectionRuleStatus
}

func (s *UpdateDetectionRuleRequest) GetDetectionRuleType() *string {
	return s.DetectionRuleType
}

func (s *UpdateDetectionRuleRequest) GetEntityMappings() *string {
	return s.EntityMappings
}

func (s *UpdateDetectionRuleRequest) GetIncidentAggregationExpression() *string {
	return s.IncidentAggregationExpression
}

func (s *UpdateDetectionRuleRequest) GetIncidentAggregationType() *string {
	return s.IncidentAggregationType
}

func (s *UpdateDetectionRuleRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDetectionRuleRequest) GetLogCategoryId() *string {
	return s.LogCategoryId
}

func (s *UpdateDetectionRuleRequest) GetLogSchemaId() *string {
	return s.LogSchemaId
}

func (s *UpdateDetectionRuleRequest) GetPlaybookParameters() *string {
	return s.PlaybookParameters
}

func (s *UpdateDetectionRuleRequest) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *UpdateDetectionRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDetectionRuleRequest) GetScheduleBeginTime() *int64 {
	return s.ScheduleBeginTime
}

func (s *UpdateDetectionRuleRequest) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *UpdateDetectionRuleRequest) GetScheduleMaxRetries() *int32 {
	return s.ScheduleMaxRetries
}

func (s *UpdateDetectionRuleRequest) GetScheduleMaxTimeout() *int32 {
	return s.ScheduleMaxTimeout
}

func (s *UpdateDetectionRuleRequest) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *UpdateDetectionRuleRequest) GetScheduleWindow() *string {
	return s.ScheduleWindow
}

func (s *UpdateDetectionRuleRequest) SetAlertAttCk(v string) *UpdateDetectionRuleRequest {
	s.AlertAttCk = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetAlertAttCkMapping(v string) *UpdateDetectionRuleRequest {
	s.AlertAttCkMapping = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetAlertDescription(v string) *UpdateDetectionRuleRequest {
	s.AlertDescription = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetAlertLevel(v string) *UpdateDetectionRuleRequest {
	s.AlertLevel = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetAlertLevelMapping(v string) *UpdateDetectionRuleRequest {
	s.AlertLevelMapping = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetAlertName(v string) *UpdateDetectionRuleRequest {
	s.AlertName = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetAlertSchemaId(v string) *UpdateDetectionRuleRequest {
	s.AlertSchemaId = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetAlertTacticId(v string) *UpdateDetectionRuleRequest {
	s.AlertTacticId = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetAlertThresholdCount(v int32) *UpdateDetectionRuleRequest {
	s.AlertThresholdCount = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetAlertThresholdGroup(v string) *UpdateDetectionRuleRequest {
	s.AlertThresholdGroup = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetAlertThresholdPeriod(v string) *UpdateDetectionRuleRequest {
	s.AlertThresholdPeriod = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetAlertType(v string) *UpdateDetectionRuleRequest {
	s.AlertType = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetAlertTypeMapping(v string) *UpdateDetectionRuleRequest {
	s.AlertTypeMapping = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetDetectionExpressionContent(v string) *UpdateDetectionRuleRequest {
	s.DetectionExpressionContent = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetDetectionExpressionType(v string) *UpdateDetectionRuleRequest {
	s.DetectionExpressionType = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetDetectionRuleDescription(v string) *UpdateDetectionRuleRequest {
	s.DetectionRuleDescription = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetDetectionRuleId(v string) *UpdateDetectionRuleRequest {
	s.DetectionRuleId = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetDetectionRuleName(v string) *UpdateDetectionRuleRequest {
	s.DetectionRuleName = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetDetectionRuleStatus(v string) *UpdateDetectionRuleRequest {
	s.DetectionRuleStatus = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetDetectionRuleType(v string) *UpdateDetectionRuleRequest {
	s.DetectionRuleType = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetEntityMappings(v string) *UpdateDetectionRuleRequest {
	s.EntityMappings = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetIncidentAggregationExpression(v string) *UpdateDetectionRuleRequest {
	s.IncidentAggregationExpression = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetIncidentAggregationType(v string) *UpdateDetectionRuleRequest {
	s.IncidentAggregationType = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetLang(v string) *UpdateDetectionRuleRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetLogCategoryId(v string) *UpdateDetectionRuleRequest {
	s.LogCategoryId = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetLogSchemaId(v string) *UpdateDetectionRuleRequest {
	s.LogSchemaId = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetPlaybookParameters(v string) *UpdateDetectionRuleRequest {
	s.PlaybookParameters = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetPlaybookUuid(v string) *UpdateDetectionRuleRequest {
	s.PlaybookUuid = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetRegionId(v string) *UpdateDetectionRuleRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetScheduleBeginTime(v int64) *UpdateDetectionRuleRequest {
	s.ScheduleBeginTime = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetScheduleExpression(v string) *UpdateDetectionRuleRequest {
	s.ScheduleExpression = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetScheduleMaxRetries(v int32) *UpdateDetectionRuleRequest {
	s.ScheduleMaxRetries = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetScheduleMaxTimeout(v int32) *UpdateDetectionRuleRequest {
	s.ScheduleMaxTimeout = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetScheduleType(v string) *UpdateDetectionRuleRequest {
	s.ScheduleType = &v
	return s
}

func (s *UpdateDetectionRuleRequest) SetScheduleWindow(v string) *UpdateDetectionRuleRequest {
	s.ScheduleWindow = &v
	return s
}

func (s *UpdateDetectionRuleRequest) Validate() error {
	return dara.Validate(s)
}
