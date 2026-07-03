// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDetectionRulesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertAttCk(v string) *ListDetectionRulesShrinkRequest
	GetAlertAttCk() *string
	SetAlertLevel(v string) *ListDetectionRulesShrinkRequest
	GetAlertLevel() *string
	SetAlertTacticId(v string) *ListDetectionRulesShrinkRequest
	GetAlertTacticId() *string
	SetAlertType(v string) *ListDetectionRulesShrinkRequest
	GetAlertType() *string
	SetDetectionExpressionType(v string) *ListDetectionRulesShrinkRequest
	GetDetectionExpressionType() *string
	SetDetectionRuleId(v string) *ListDetectionRulesShrinkRequest
	GetDetectionRuleId() *string
	SetDetectionRuleIdsShrink(v string) *ListDetectionRulesShrinkRequest
	GetDetectionRuleIdsShrink() *string
	SetDetectionRuleName(v string) *ListDetectionRulesShrinkRequest
	GetDetectionRuleName() *string
	SetDetectionRuleStatus(v string) *ListDetectionRulesShrinkRequest
	GetDetectionRuleStatus() *string
	SetDetectionRuleType(v string) *ListDetectionRulesShrinkRequest
	GetDetectionRuleType() *string
	SetIncidentAggregationType(v string) *ListDetectionRulesShrinkRequest
	GetIncidentAggregationType() *string
	SetLang(v string) *ListDetectionRulesShrinkRequest
	GetLang() *string
	SetLogCategoryId(v string) *ListDetectionRulesShrinkRequest
	GetLogCategoryId() *string
	SetLogSchemaId(v string) *ListDetectionRulesShrinkRequest
	GetLogSchemaId() *string
	SetMaxResults(v int32) *ListDetectionRulesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDetectionRulesShrinkRequest
	GetNextToken() *string
	SetOrderDirection(v string) *ListDetectionRulesShrinkRequest
	GetOrderDirection() *string
	SetOrderFieldName(v string) *ListDetectionRulesShrinkRequest
	GetOrderFieldName() *string
	SetPageNumber(v int32) *ListDetectionRulesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDetectionRulesShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListDetectionRulesShrinkRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDetectionRulesShrinkRequest
	GetRoleFor() *int64
}

type ListDetectionRulesShrinkRequest struct {
	// The ATT\\&CK technique of the alert.
	//
	// example:
	//
	// Discovery
	AlertAttCk *string `json:"AlertAttCk,omitempty" xml:"AlertAttCk,omitempty"`
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
	AlertLevel *string `json:"AlertLevel,omitempty" xml:"AlertLevel,omitempty"`
	// The alert tactic phase.
	//
	// example:
	//
	// TA0042
	AlertTacticId *string `json:"AlertTacticId,omitempty" xml:"AlertTacticId,omitempty"`
	// The alert type.
	//
	// example:
	//
	// siem_rule_type_alert_storm
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// The type of the detection rule expression.
	//
	// example:
	//
	// playbook
	DetectionExpressionType *string `json:"DetectionExpressionType,omitempty" xml:"DetectionExpressionType,omitempty"`
	// The ID of the detection rule.
	//
	// example:
	//
	// dr-ppa85gfw69tgwkys****
	DetectionRuleId *string `json:"DetectionRuleId,omitempty" xml:"DetectionRuleId,omitempty"`
	// The list of detection rule IDs.
	DetectionRuleIdsShrink *string `json:"DetectionRuleIds,omitempty" xml:"DetectionRuleIds,omitempty"`
	// The name of the detection rule.
	//
	// example:
	//
	// Detect Discovery Behavior for Local Systems Groups
	DetectionRuleName *string `json:"DetectionRuleName,omitempty" xml:"DetectionRuleName,omitempty"`
	// The status of the detection rule.
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
	// preset
	DetectionRuleType *string `json:"DetectionRuleType,omitempty" xml:"DetectionRuleType,omitempty"`
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
	// graph_compute
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
	// The sort order. Valid values:
	//
	// - **asc**: ascending order. This is the default value.
	//
	// - **desc**: descending order.
	//
	// example:
	//
	// asc
	OrderDirection *string `json:"OrderDirection,omitempty" xml:"OrderDirection,omitempty"`
	// The field to sort by. Valid values:
	//
	// - GmtCreate: the creation time.
	//
	// - GmtModified: the update time.
	//
	// example:
	//
	// GmtModified
	OrderFieldName *string `json:"OrderFieldName,omitempty" xml:"OrderFieldName,omitempty"`
	// The pagination parameter. This specifies the current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The pagination parameter. This specifies the number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the Data Management center of threat analysis is located. Select a region based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets are outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID that the administrator uses to switch to the perspective of another member.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListDetectionRulesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDetectionRulesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListDetectionRulesShrinkRequest) GetAlertAttCk() *string {
	return s.AlertAttCk
}

func (s *ListDetectionRulesShrinkRequest) GetAlertLevel() *string {
	return s.AlertLevel
}

func (s *ListDetectionRulesShrinkRequest) GetAlertTacticId() *string {
	return s.AlertTacticId
}

func (s *ListDetectionRulesShrinkRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *ListDetectionRulesShrinkRequest) GetDetectionExpressionType() *string {
	return s.DetectionExpressionType
}

func (s *ListDetectionRulesShrinkRequest) GetDetectionRuleId() *string {
	return s.DetectionRuleId
}

func (s *ListDetectionRulesShrinkRequest) GetDetectionRuleIdsShrink() *string {
	return s.DetectionRuleIdsShrink
}

func (s *ListDetectionRulesShrinkRequest) GetDetectionRuleName() *string {
	return s.DetectionRuleName
}

func (s *ListDetectionRulesShrinkRequest) GetDetectionRuleStatus() *string {
	return s.DetectionRuleStatus
}

func (s *ListDetectionRulesShrinkRequest) GetDetectionRuleType() *string {
	return s.DetectionRuleType
}

func (s *ListDetectionRulesShrinkRequest) GetIncidentAggregationType() *string {
	return s.IncidentAggregationType
}

func (s *ListDetectionRulesShrinkRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDetectionRulesShrinkRequest) GetLogCategoryId() *string {
	return s.LogCategoryId
}

func (s *ListDetectionRulesShrinkRequest) GetLogSchemaId() *string {
	return s.LogSchemaId
}

func (s *ListDetectionRulesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDetectionRulesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDetectionRulesShrinkRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListDetectionRulesShrinkRequest) GetOrderFieldName() *string {
	return s.OrderFieldName
}

func (s *ListDetectionRulesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDetectionRulesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDetectionRulesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDetectionRulesShrinkRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDetectionRulesShrinkRequest) SetAlertAttCk(v string) *ListDetectionRulesShrinkRequest {
	s.AlertAttCk = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetAlertLevel(v string) *ListDetectionRulesShrinkRequest {
	s.AlertLevel = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetAlertTacticId(v string) *ListDetectionRulesShrinkRequest {
	s.AlertTacticId = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetAlertType(v string) *ListDetectionRulesShrinkRequest {
	s.AlertType = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetDetectionExpressionType(v string) *ListDetectionRulesShrinkRequest {
	s.DetectionExpressionType = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetDetectionRuleId(v string) *ListDetectionRulesShrinkRequest {
	s.DetectionRuleId = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetDetectionRuleIdsShrink(v string) *ListDetectionRulesShrinkRequest {
	s.DetectionRuleIdsShrink = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetDetectionRuleName(v string) *ListDetectionRulesShrinkRequest {
	s.DetectionRuleName = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetDetectionRuleStatus(v string) *ListDetectionRulesShrinkRequest {
	s.DetectionRuleStatus = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetDetectionRuleType(v string) *ListDetectionRulesShrinkRequest {
	s.DetectionRuleType = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetIncidentAggregationType(v string) *ListDetectionRulesShrinkRequest {
	s.IncidentAggregationType = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetLang(v string) *ListDetectionRulesShrinkRequest {
	s.Lang = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetLogCategoryId(v string) *ListDetectionRulesShrinkRequest {
	s.LogCategoryId = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetLogSchemaId(v string) *ListDetectionRulesShrinkRequest {
	s.LogSchemaId = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetMaxResults(v int32) *ListDetectionRulesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetNextToken(v string) *ListDetectionRulesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetOrderDirection(v string) *ListDetectionRulesShrinkRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetOrderFieldName(v string) *ListDetectionRulesShrinkRequest {
	s.OrderFieldName = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetPageNumber(v int32) *ListDetectionRulesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetPageSize(v int32) *ListDetectionRulesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetRegionId(v string) *ListDetectionRulesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) SetRoleFor(v int64) *ListDetectionRulesShrinkRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDetectionRulesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
