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
	SetAlertLevel(v string) *UpdateDetectionRuleRequest
	GetAlertLevel() *string
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
	// example:
	//
	// Discovery
	AlertAttCk *string `json:"AlertAttCk,omitempty" xml:"AlertAttCk,omitempty"`
	// example:
	//
	// 1
	AlertLevel *string `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// example:
	//
	// ALERT_ACTIVITY
	AlertSchemaId *string `json:"AlertSchemaId,omitempty" xml:"AlertSchemaId,omitempty"`
	// example:
	//
	// TA0042
	AlertTacticId *string `json:"AlertTacticId,omitempty" xml:"AlertTacticId,omitempty"`
	// example:
	//
	// 10
	AlertThresholdCount *int32 `json:"AlertThresholdCount,omitempty" xml:"AlertThresholdCount,omitempty"`
	// example:
	//
	// alert_type,ip
	AlertThresholdGroup *string `json:"AlertThresholdGroup,omitempty" xml:"AlertThresholdGroup,omitempty"`
	// example:
	//
	// 5m
	AlertThresholdPeriod *string `json:"AlertThresholdPeriod,omitempty" xml:"AlertThresholdPeriod,omitempty"`
	// example:
	//
	// siem_rule_type_alert_storm
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
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
	// example:
	//
	// sql
	DetectionExpressionType *string `json:"DetectionExpressionType,omitempty" xml:"DetectionExpressionType,omitempty"`
	// example:
	//
	// Check the enumeration behavior of local system groups. An attacker may attempt to find the Local Systems group and its permission settings.
	DetectionRuleDescription *string `json:"DetectionRuleDescription,omitempty" xml:"DetectionRuleDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// jndi-attack-success_http_dns
	DetectionRuleId *string `json:"DetectionRuleId,omitempty" xml:"DetectionRuleId,omitempty"`
	// example:
	//
	// CTDR Port Scan Behavior
	DetectionRuleName *string `json:"DetectionRuleName,omitempty" xml:"DetectionRuleName,omitempty"`
	// example:
	//
	// enabled
	DetectionRuleStatus *string `json:"DetectionRuleStatus,omitempty" xml:"DetectionRuleStatus,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// custom
	DetectionRuleType *string `json:"DetectionRuleType,omitempty" xml:"DetectionRuleType,omitempty"`
	// example:
	//
	// [{\\"NormalizationSchemaId\\":\\"host\\",\\"NormalizationFieldMappings\\":[{\\"NormalizationFieldName\\":\\"uuid\\",\\"MappingFieldName\\":\\"host\\",\\"NormalizationFieldType\\":\\"varchar\\"}]}]
	EntityMappings *string `json:"EntityMappings,omitempty" xml:"EntityMappings,omitempty"`
	// example:
	//
	// 60m
	IncidentAggregationExpression *string `json:"IncidentAggregationExpression,omitempty" xml:"IncidentAggregationExpression,omitempty"`
	// example:
	//
	// window
	IncidentAggregationType *string `json:"IncidentAggregationType,omitempty" xml:"IncidentAggregationType,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// NETWORK_CATEGORY
	LogCategoryId *string `json:"LogCategoryId,omitempty" xml:"LogCategoryId,omitempty"`
	// example:
	//
	// API_RISK_ACTIVITY
	LogSchemaId *string `json:"LogSchemaId,omitempty" xml:"LogSchemaId,omitempty"`
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
	// example:
	//
	// 31568394-7a86-487c-b8ec-b3f42b59****
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 1733269771123
	ScheduleBeginTime *int64 `json:"ScheduleBeginTime,omitempty" xml:"ScheduleBeginTime,omitempty"`
	// example:
	//
	// 1h
	ScheduleExpression *string `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	// example:
	//
	// 1
	ScheduleMaxRetries *int32 `json:"ScheduleMaxRetries,omitempty" xml:"ScheduleMaxRetries,omitempty"`
	// example:
	//
	// 60
	ScheduleMaxTimeout *int32 `json:"ScheduleMaxTimeout,omitempty" xml:"ScheduleMaxTimeout,omitempty"`
	// example:
	//
	// fixed_rate
	ScheduleType *string `json:"ScheduleType,omitempty" xml:"ScheduleType,omitempty"`
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

func (s *UpdateDetectionRuleRequest) GetAlertLevel() *string {
	return s.AlertLevel
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

func (s *UpdateDetectionRuleRequest) SetAlertLevel(v string) *UpdateDetectionRuleRequest {
	s.AlertLevel = &v
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
