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
	SetAlertLevel(v string) *CreateDetectionRuleRequest
	GetAlertLevel() *string
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
	SetDetectionRuleType(v string) *CreateDetectionRuleRequest
	GetDetectionRuleType() *string
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
	// example:
	//
	// Discovery
	AlertAttCk *string `json:"AlertAttCk,omitempty" xml:"AlertAttCk,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	AlertLevel *string `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// This parameter is required.
	//
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
	// This parameter is required.
	//
	// example:
	//
	// WebShell
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
	// dr-123
	DetectionRuleDescription *string `json:"DetectionRuleDescription,omitempty" xml:"DetectionRuleDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dr-ha1i09ob3zmqrs85****
	DetectionRuleName *string `json:"DetectionRuleName,omitempty" xml:"DetectionRuleName,omitempty"`
	// example:
	//
	// 0
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
	// 5m
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
	// This parameter is required.
	//
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
	// system_aliyun_clb_process_book
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
	// example:
	//
	// 1733269771123
	ScheduleBeginTime *int64 `json:"ScheduleBeginTime,omitempty" xml:"ScheduleBeginTime,omitempty"`
	// example:
	//
	// 0/5 	- 	- 	- *
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

func (s CreateDetectionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDetectionRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDetectionRuleRequest) GetAlertAttCk() *string {
	return s.AlertAttCk
}

func (s *CreateDetectionRuleRequest) GetAlertLevel() *string {
	return s.AlertLevel
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

func (s *CreateDetectionRuleRequest) GetDetectionRuleType() *string {
	return s.DetectionRuleType
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

func (s *CreateDetectionRuleRequest) SetAlertLevel(v string) *CreateDetectionRuleRequest {
	s.AlertLevel = &v
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

func (s *CreateDetectionRuleRequest) SetDetectionRuleType(v string) *CreateDetectionRuleRequest {
	s.DetectionRuleType = &v
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
