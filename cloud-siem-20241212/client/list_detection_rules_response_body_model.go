// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDetectionRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetectionRules(v []*ListDetectionRulesResponseBodyDetectionRules) *ListDetectionRulesResponseBody
	GetDetectionRules() []*ListDetectionRulesResponseBodyDetectionRules
	SetMaxResults(v int32) *ListDetectionRulesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDetectionRulesResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListDetectionRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDetectionRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDetectionRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDetectionRulesResponseBody
	GetTotalCount() *int32
}

type ListDetectionRulesResponseBody struct {
	// The list of detection rules.
	DetectionRules []*ListDetectionRulesResponseBodyDetectionRules `json:"DetectionRules,omitempty" xml:"DetectionRules,omitempty" type:"Repeated"`
	// The maximum number of data entries to read.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to start the next query.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The pagination parameter. This specifies the current page number.
	//
	// example:
	//
	// 2
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The pagination parameter. This specifies the number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 508DCFFD-4508-54BF-A8A0-E97A0FA6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDetectionRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDetectionRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDetectionRulesResponseBody) GetDetectionRules() []*ListDetectionRulesResponseBodyDetectionRules {
	return s.DetectionRules
}

func (s *ListDetectionRulesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDetectionRulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDetectionRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDetectionRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDetectionRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDetectionRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDetectionRulesResponseBody) SetDetectionRules(v []*ListDetectionRulesResponseBodyDetectionRules) *ListDetectionRulesResponseBody {
	s.DetectionRules = v
	return s
}

func (s *ListDetectionRulesResponseBody) SetMaxResults(v int32) *ListDetectionRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDetectionRulesResponseBody) SetNextToken(v string) *ListDetectionRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDetectionRulesResponseBody) SetPageNumber(v int32) *ListDetectionRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDetectionRulesResponseBody) SetPageSize(v int32) *ListDetectionRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDetectionRulesResponseBody) SetRequestId(v string) *ListDetectionRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDetectionRulesResponseBody) SetTotalCount(v int32) *ListDetectionRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDetectionRulesResponseBody) Validate() error {
	if s.DetectionRules != nil {
		for _, item := range s.DetectionRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDetectionRulesResponseBodyDetectionRules struct {
	// The ATT\\&CK technique of the alert.
	//
	// example:
	//
	// Discovery
	AlertAttCk        *string `json:"AlertAttCk,omitempty" xml:"AlertAttCk,omitempty"`
	AlertAttCkMapping *string `json:"AlertAttCkMapping,omitempty" xml:"AlertAttCkMapping,omitempty"`
	// The alert description. You can use $$ to reference the output fields of the query.
	//
	// example:
	//
	// Alert from: $product_code$, detected network attack from $src_ip$, affected assets include: $dst_ip$
	AlertDescription *string `json:"AlertDescription,omitempty" xml:"AlertDescription,omitempty"`
	// The threat level of the alert. Valid values:
	//
	// - 5: critical.
	//
	// - 4: important.
	//
	// - 3: medium.
	//
	// - 2: low.
	//
	// - 1: informational.
	//
	// example:
	//
	// 1
	AlertLevel        *string `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	AlertLevelMapping *string `json:"AlertLevelMapping,omitempty" xml:"AlertLevelMapping,omitempty"`
	// The alert name. You can use $$ to reference the output fields of the query.
	//
	// example:
	//
	// Detected high-frequency multi-type network attacks from $src_ip$
	AlertName *string `json:"AlertName,omitempty" xml:"AlertName,omitempty"`
	// The ID of the alert template for the detection rule. Valid values:
	//
	// - ALERT_ACTIVITY: other alerts.
	//
	// - EDR_ALERT_ACTIVITY: Endpoint Detection and Response (EDR) alerts.
	//
	// - FIREWALL_ALERT_ACTIVITY: firewall alerts.
	//
	// - WAF_ALERT_ACTIVITY: Web Application Firewall (WAF) alerts.
	//
	// example:
	//
	// ALERT_ACTIVITY
	AlertSchemaId *string `json:"AlertSchemaId,omitempty" xml:"AlertSchemaId,omitempty"`
	// The alert tactic phase.
	//
	// example:
	//
	// TA0042
	AlertTacticId *string `json:"AlertTacticId,omitempty" xml:"AlertTacticId,omitempty"`
	// The number of alert thresholds.
	//
	// example:
	//
	// 10
	AlertThresholdCount *int32 `json:"AlertThresholdCount,omitempty" xml:"AlertThresholdCount,omitempty"`
	// The list of alert threshold fields, separated by commas.
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
	// The alert type.
	//
	// example:
	//
	// WebShell
	AlertType        *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	AlertTypeMapping *string `json:"AlertTypeMapping,omitempty" xml:"AlertTypeMapping,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2023-03-21 13:47:01
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
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
	// The type of the detection rule expression.
	//
	// example:
	//
	// playbook
	DetectionExpressionType *string `json:"DetectionExpressionType,omitempty" xml:"DetectionExpressionType,omitempty"`
	// The description of the detection rule.
	//
	// example:
	//
	// Check the enumeration behavior of local system groups. An attacker may attempt to find the Local Systems group and its permission settings.
	DetectionRuleDescription *string `json:"DetectionRuleDescription,omitempty" xml:"DetectionRuleDescription,omitempty"`
	// The ID of the detection rule.
	//
	// example:
	//
	// jndi-attack-success_http_netstat
	DetectionRuleId *string `json:"DetectionRuleId,omitempty" xml:"DetectionRuleId,omitempty"`
	// The name of the detection rule.
	//
	// example:
	//
	// Detect Discovery Behavior for Local Systems Groups
	DetectionRuleName *string `json:"DetectionRuleName,omitempty" xml:"DetectionRuleName,omitempty"`
	// The status of the detection rule. Valid values:
	//
	// - enabled: The rule is enabled.
	//
	// - disabled: The rule is disabled.
	//
	// - testing: The rule is being tested.
	//
	// example:
	//
	// enabled
	DetectionRuleStatus *string `json:"DetectionRuleStatus,omitempty" xml:"DetectionRuleStatus,omitempty"`
	// The type of the detection rule. Valid values:
	//
	// - preset: a built-in detection rule.
	//
	// - custom: a custom detection rule.
	//
	// - custom_template: a rule template.
	//
	// example:
	//
	// custom
	DetectionRuleType *string `json:"DetectionRuleType,omitempty" xml:"DetectionRuleType,omitempty"`
	// The entity mapping configuration.
	EntityMappings []*ListDetectionRulesResponseBodyDetectionRulesEntityMappings `json:"EntityMappings,omitempty" xml:"EntityMappings,omitempty" type:"Repeated"`
	// The configuration of the event aggregation period.
	//
	// example:
	//
	// 5m
	IncidentAggregationExpression *string `json:"IncidentAggregationExpression,omitempty" xml:"IncidentAggregationExpression,omitempty"`
	// The event aggregation type. Valid values:
	//
	// - none: No events are generated.
	//
	// - graph_compute: graph computing (supported by predefined rules).
	//
	// - expert: expert rule.
	//
	// - passthrough: alert passthrough (one-to-one).
	//
	// - window: aggregation of similar alerts (by window).
	//
	// example:
	//
	// passthrough
	IncidentAggregationType *string `json:"IncidentAggregationType,omitempty" xml:"IncidentAggregationType,omitempty"`
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
	// The configuration of the playbook.
	Playbook *ListDetectionRulesResponseBodyDetectionRulesPlaybook `json:"Playbook,omitempty" xml:"Playbook,omitempty" type:"Struct"`
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
	// dde983ed-eb56-45ea-ac2e-3b12b2a9****
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
	// The start time of the scheduling. This is a 13-digit timestamp.
	//
	// example:
	//
	// 1733269771123
	ScheduleBeginTime *int64 `json:"ScheduleBeginTime,omitempty" xml:"ScheduleBeginTime,omitempty"`
	// The cron expression for scheduling. This parameter is required when ScheduleType is set to cron.
	//
	// example:
	//
	// 5m
	ScheduleExpression *string `json:"ScheduleExpression,omitempty" xml:"ScheduleExpression,omitempty"`
	// The maximum number of retries after a timeout. The value can be from 1 to 100.
	//
	// example:
	//
	// 1
	ScheduleMaxRetries *int32 `json:"ScheduleMaxRetries,omitempty" xml:"ScheduleMaxRetries,omitempty"`
	// The maximum timeout period, in seconds. The value can be from 60 to 1800.
	//
	// example:
	//
	// 60
	ScheduleMaxTimeout *int32 `json:"ScheduleMaxTimeout,omitempty" xml:"ScheduleMaxTimeout,omitempty"`
	// The scheduling type. Valid values:
	//
	// - fixed_rate: fixed interval
	//
	// - cron: cron expression
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
	// The update time.
	//
	// example:
	//
	// 2023-04-16 10:51:00
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDetectionRulesResponseBodyDetectionRules) String() string {
	return dara.Prettify(s)
}

func (s ListDetectionRulesResponseBodyDetectionRules) GoString() string {
	return s.String()
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetAlertAttCk() *string {
	return s.AlertAttCk
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetAlertAttCkMapping() *string {
	return s.AlertAttCkMapping
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetAlertDescription() *string {
	return s.AlertDescription
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetAlertLevel() *string {
	return s.AlertLevel
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetAlertLevelMapping() *string {
	return s.AlertLevelMapping
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetAlertName() *string {
	return s.AlertName
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetAlertSchemaId() *string {
	return s.AlertSchemaId
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetAlertTacticId() *string {
	return s.AlertTacticId
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetAlertThresholdCount() *int32 {
	return s.AlertThresholdCount
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetAlertThresholdGroup() *string {
	return s.AlertThresholdGroup
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetAlertThresholdPeriod() *string {
	return s.AlertThresholdPeriod
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetAlertType() *string {
	return s.AlertType
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetAlertTypeMapping() *string {
	return s.AlertTypeMapping
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetDetectionExpressionContent() *string {
	return s.DetectionExpressionContent
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetDetectionExpressionType() *string {
	return s.DetectionExpressionType
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetDetectionRuleDescription() *string {
	return s.DetectionRuleDescription
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetDetectionRuleId() *string {
	return s.DetectionRuleId
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetDetectionRuleName() *string {
	return s.DetectionRuleName
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetDetectionRuleStatus() *string {
	return s.DetectionRuleStatus
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetDetectionRuleType() *string {
	return s.DetectionRuleType
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetEntityMappings() []*ListDetectionRulesResponseBodyDetectionRulesEntityMappings {
	return s.EntityMappings
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetIncidentAggregationExpression() *string {
	return s.IncidentAggregationExpression
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetIncidentAggregationType() *string {
	return s.IncidentAggregationType
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetLogCategoryId() *string {
	return s.LogCategoryId
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetLogSchemaId() *string {
	return s.LogSchemaId
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetPlaybook() *ListDetectionRulesResponseBodyDetectionRulesPlaybook {
	return s.Playbook
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetPlaybookParameters() *string {
	return s.PlaybookParameters
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetScheduleBeginTime() *int64 {
	return s.ScheduleBeginTime
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetScheduleExpression() *string {
	return s.ScheduleExpression
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetScheduleMaxRetries() *int32 {
	return s.ScheduleMaxRetries
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetScheduleMaxTimeout() *int32 {
	return s.ScheduleMaxTimeout
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetScheduleType() *string {
	return s.ScheduleType
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetScheduleWindow() *string {
	return s.ScheduleWindow
}

func (s *ListDetectionRulesResponseBodyDetectionRules) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetAlertAttCk(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.AlertAttCk = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetAlertAttCkMapping(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.AlertAttCkMapping = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetAlertDescription(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.AlertDescription = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetAlertLevel(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.AlertLevel = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetAlertLevelMapping(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.AlertLevelMapping = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetAlertName(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.AlertName = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetAlertSchemaId(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.AlertSchemaId = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetAlertTacticId(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.AlertTacticId = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetAlertThresholdCount(v int32) *ListDetectionRulesResponseBodyDetectionRules {
	s.AlertThresholdCount = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetAlertThresholdGroup(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.AlertThresholdGroup = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetAlertThresholdPeriod(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.AlertThresholdPeriod = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetAlertType(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.AlertType = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetAlertTypeMapping(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.AlertTypeMapping = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetCreateTime(v int64) *ListDetectionRulesResponseBodyDetectionRules {
	s.CreateTime = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetDetectionExpressionContent(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.DetectionExpressionContent = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetDetectionExpressionType(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.DetectionExpressionType = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetDetectionRuleDescription(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.DetectionRuleDescription = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetDetectionRuleId(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.DetectionRuleId = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetDetectionRuleName(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.DetectionRuleName = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetDetectionRuleStatus(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.DetectionRuleStatus = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetDetectionRuleType(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.DetectionRuleType = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetEntityMappings(v []*ListDetectionRulesResponseBodyDetectionRulesEntityMappings) *ListDetectionRulesResponseBodyDetectionRules {
	s.EntityMappings = v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetIncidentAggregationExpression(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.IncidentAggregationExpression = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetIncidentAggregationType(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.IncidentAggregationType = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetLogCategoryId(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.LogCategoryId = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetLogSchemaId(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.LogSchemaId = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetPlaybook(v *ListDetectionRulesResponseBodyDetectionRulesPlaybook) *ListDetectionRulesResponseBodyDetectionRules {
	s.Playbook = v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetPlaybookParameters(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.PlaybookParameters = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetPlaybookUuid(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.PlaybookUuid = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetScheduleBeginTime(v int64) *ListDetectionRulesResponseBodyDetectionRules {
	s.ScheduleBeginTime = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetScheduleExpression(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.ScheduleExpression = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetScheduleMaxRetries(v int32) *ListDetectionRulesResponseBodyDetectionRules {
	s.ScheduleMaxRetries = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetScheduleMaxTimeout(v int32) *ListDetectionRulesResponseBodyDetectionRules {
	s.ScheduleMaxTimeout = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetScheduleType(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.ScheduleType = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetScheduleWindow(v string) *ListDetectionRulesResponseBodyDetectionRules {
	s.ScheduleWindow = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) SetUpdateTime(v int64) *ListDetectionRulesResponseBodyDetectionRules {
	s.UpdateTime = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRules) Validate() error {
	if s.EntityMappings != nil {
		for _, item := range s.EntityMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Playbook != nil {
		if err := s.Playbook.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDetectionRulesResponseBodyDetectionRulesEntityMappings struct {
	// The entity mapping configuration.
	NormalizationFieldMappings []*ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings `json:"NormalizationFieldMappings,omitempty" xml:"NormalizationFieldMappings,omitempty" type:"Repeated"`
	// The entity type ID.
	//
	// example:
	//
	// ip
	NormalizationSchemaId *string `json:"NormalizationSchemaId,omitempty" xml:"NormalizationSchemaId,omitempty"`
}

func (s ListDetectionRulesResponseBodyDetectionRulesEntityMappings) String() string {
	return dara.Prettify(s)
}

func (s ListDetectionRulesResponseBodyDetectionRulesEntityMappings) GoString() string {
	return s.String()
}

func (s *ListDetectionRulesResponseBodyDetectionRulesEntityMappings) GetNormalizationFieldMappings() []*ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings {
	return s.NormalizationFieldMappings
}

func (s *ListDetectionRulesResponseBodyDetectionRulesEntityMappings) GetNormalizationSchemaId() *string {
	return s.NormalizationSchemaId
}

func (s *ListDetectionRulesResponseBodyDetectionRulesEntityMappings) SetNormalizationFieldMappings(v []*ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings) *ListDetectionRulesResponseBodyDetectionRulesEntityMappings {
	s.NormalizationFieldMappings = v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRulesEntityMappings) SetNormalizationSchemaId(v string) *ListDetectionRulesResponseBodyDetectionRulesEntityMappings {
	s.NormalizationSchemaId = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRulesEntityMappings) Validate() error {
	if s.NormalizationFieldMappings != nil {
		for _, item := range s.NormalizationFieldMappings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings struct {
	// The mapped field.
	//
	// example:
	//
	// ip
	MappingFieldName *string `json:"MappingFieldName,omitempty" xml:"MappingFieldName,omitempty"`
	// The standard entity field.
	//
	// example:
	//
	// src_ip
	NormalizationFieldName *string `json:"NormalizationFieldName,omitempty" xml:"NormalizationFieldName,omitempty"`
	// The data type of the standard entity field.
	//
	// example:
	//
	// vachar
	NormalizationFieldType *string `json:"NormalizationFieldType,omitempty" xml:"NormalizationFieldType,omitempty"`
}

func (s ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings) String() string {
	return dara.Prettify(s)
}

func (s ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings) GoString() string {
	return s.String()
}

func (s *ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings) GetMappingFieldName() *string {
	return s.MappingFieldName
}

func (s *ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings) GetNormalizationFieldName() *string {
	return s.NormalizationFieldName
}

func (s *ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings) GetNormalizationFieldType() *string {
	return s.NormalizationFieldType
}

func (s *ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings) SetMappingFieldName(v string) *ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings {
	s.MappingFieldName = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings) SetNormalizationFieldName(v string) *ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings {
	s.NormalizationFieldName = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings) SetNormalizationFieldType(v string) *ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings {
	s.NormalizationFieldType = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRulesEntityMappingsNormalizationFieldMappings) Validate() error {
	return dara.Validate(s)
}

type ListDetectionRulesResponseBodyDetectionRulesPlaybook struct {
	// The configuration of the input parameters for invoking the playbook.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "name": "expireDay",
	//
	//         "dataType": "Integer",
	//
	//         "required": true,
	//
	//         "isArray": false,
	//
	//         "example": "7",
	//
	//         "description": "desc",
	//
	//         "typeName": "Integer",
	//
	//         "dataClass": "normal",
	//
	//         "stanchDefaultValue": "7"
	//
	//     }
	//
	// ]
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The flow configuration of the playbook.
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "id": "EndEvent_1fqpq4h",
	//
	//         "zIndex": 1,
	//
	//         "data": {
	//
	//             "nodeType": "endEvent",
	//
	//             "appType": "basic",
	//
	//             "valueData": {
	//
	//             },
	//
	//             "icon": "icon-radio-off-full"
	//
	//         },
	//
	//         "position": {
	//
	//             "x": 1369,
	//
	//             "y": 174
	//
	//         }
	//
	//     }
	//
	// ]
	Flow *string `json:"Flow,omitempty" xml:"Flow,omitempty"`
}

func (s ListDetectionRulesResponseBodyDetectionRulesPlaybook) String() string {
	return dara.Prettify(s)
}

func (s ListDetectionRulesResponseBodyDetectionRulesPlaybook) GoString() string {
	return s.String()
}

func (s *ListDetectionRulesResponseBodyDetectionRulesPlaybook) GetConfig() *string {
	return s.Config
}

func (s *ListDetectionRulesResponseBodyDetectionRulesPlaybook) GetFlow() *string {
	return s.Flow
}

func (s *ListDetectionRulesResponseBodyDetectionRulesPlaybook) SetConfig(v string) *ListDetectionRulesResponseBodyDetectionRulesPlaybook {
	s.Config = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRulesPlaybook) SetFlow(v string) *ListDetectionRulesResponseBodyDetectionRulesPlaybook {
	s.Flow = &v
	return s
}

func (s *ListDetectionRulesResponseBodyDetectionRulesPlaybook) Validate() error {
	return dara.Validate(s)
}
