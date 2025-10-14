// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDetectionRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlertAttCk(v string) *ListDetectionRulesRequest
	GetAlertAttCk() *string
	SetAlertLevel(v string) *ListDetectionRulesRequest
	GetAlertLevel() *string
	SetAlertTacticId(v string) *ListDetectionRulesRequest
	GetAlertTacticId() *string
	SetAlertType(v string) *ListDetectionRulesRequest
	GetAlertType() *string
	SetDetectionExpressionType(v string) *ListDetectionRulesRequest
	GetDetectionExpressionType() *string
	SetDetectionRuleId(v string) *ListDetectionRulesRequest
	GetDetectionRuleId() *string
	SetDetectionRuleIds(v []*string) *ListDetectionRulesRequest
	GetDetectionRuleIds() []*string
	SetDetectionRuleName(v string) *ListDetectionRulesRequest
	GetDetectionRuleName() *string
	SetDetectionRuleStatus(v string) *ListDetectionRulesRequest
	GetDetectionRuleStatus() *string
	SetDetectionRuleType(v string) *ListDetectionRulesRequest
	GetDetectionRuleType() *string
	SetIncidentAggregationType(v string) *ListDetectionRulesRequest
	GetIncidentAggregationType() *string
	SetLang(v string) *ListDetectionRulesRequest
	GetLang() *string
	SetLogCategoryId(v string) *ListDetectionRulesRequest
	GetLogCategoryId() *string
	SetLogSchemaId(v string) *ListDetectionRulesRequest
	GetLogSchemaId() *string
	SetMaxResults(v int32) *ListDetectionRulesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDetectionRulesRequest
	GetNextToken() *string
	SetOrderDirection(v string) *ListDetectionRulesRequest
	GetOrderDirection() *string
	SetOrderFieldName(v string) *ListDetectionRulesRequest
	GetOrderFieldName() *string
	SetPageNumber(v int32) *ListDetectionRulesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDetectionRulesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListDetectionRulesRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ListDetectionRulesRequest
	GetRoleFor() *int64
}

type ListDetectionRulesRequest struct {
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
	// TA0042
	AlertTacticId *string `json:"AlertTacticId,omitempty" xml:"AlertTacticId,omitempty"`
	// example:
	//
	// siem_rule_type_alert_storm
	AlertType *string `json:"AlertType,omitempty" xml:"AlertType,omitempty"`
	// example:
	//
	// playbook
	DetectionExpressionType *string `json:"DetectionExpressionType,omitempty" xml:"DetectionExpressionType,omitempty"`
	// example:
	//
	// dr-ppa85gfw69tgwkys****
	DetectionRuleId  *string   `json:"DetectionRuleId,omitempty" xml:"DetectionRuleId,omitempty"`
	DetectionRuleIds []*string `json:"DetectionRuleIds,omitempty" xml:"DetectionRuleIds,omitempty" type:"Repeated"`
	// example:
	//
	// Detect Discovery Behavior for Local Systems Groups
	DetectionRuleName *string `json:"DetectionRuleName,omitempty" xml:"DetectionRuleName,omitempty"`
	// example:
	//
	// enabled
	DetectionRuleStatus *string `json:"DetectionRuleStatus,omitempty" xml:"DetectionRuleStatus,omitempty"`
	// example:
	//
	// preset
	DetectionRuleType *string `json:"DetectionRuleType,omitempty" xml:"DetectionRuleType,omitempty"`
	// example:
	//
	// graph_compute
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
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// asc
	OrderDirection *string `json:"OrderDirection,omitempty" xml:"OrderDirection,omitempty"`
	// example:
	//
	// GmtModified
	OrderFieldName *string `json:"OrderFieldName,omitempty" xml:"OrderFieldName,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ListDetectionRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDetectionRulesRequest) GoString() string {
	return s.String()
}

func (s *ListDetectionRulesRequest) GetAlertAttCk() *string {
	return s.AlertAttCk
}

func (s *ListDetectionRulesRequest) GetAlertLevel() *string {
	return s.AlertLevel
}

func (s *ListDetectionRulesRequest) GetAlertTacticId() *string {
	return s.AlertTacticId
}

func (s *ListDetectionRulesRequest) GetAlertType() *string {
	return s.AlertType
}

func (s *ListDetectionRulesRequest) GetDetectionExpressionType() *string {
	return s.DetectionExpressionType
}

func (s *ListDetectionRulesRequest) GetDetectionRuleId() *string {
	return s.DetectionRuleId
}

func (s *ListDetectionRulesRequest) GetDetectionRuleIds() []*string {
	return s.DetectionRuleIds
}

func (s *ListDetectionRulesRequest) GetDetectionRuleName() *string {
	return s.DetectionRuleName
}

func (s *ListDetectionRulesRequest) GetDetectionRuleStatus() *string {
	return s.DetectionRuleStatus
}

func (s *ListDetectionRulesRequest) GetDetectionRuleType() *string {
	return s.DetectionRuleType
}

func (s *ListDetectionRulesRequest) GetIncidentAggregationType() *string {
	return s.IncidentAggregationType
}

func (s *ListDetectionRulesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListDetectionRulesRequest) GetLogCategoryId() *string {
	return s.LogCategoryId
}

func (s *ListDetectionRulesRequest) GetLogSchemaId() *string {
	return s.LogSchemaId
}

func (s *ListDetectionRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDetectionRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDetectionRulesRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListDetectionRulesRequest) GetOrderFieldName() *string {
	return s.OrderFieldName
}

func (s *ListDetectionRulesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDetectionRulesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDetectionRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDetectionRulesRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ListDetectionRulesRequest) SetAlertAttCk(v string) *ListDetectionRulesRequest {
	s.AlertAttCk = &v
	return s
}

func (s *ListDetectionRulesRequest) SetAlertLevel(v string) *ListDetectionRulesRequest {
	s.AlertLevel = &v
	return s
}

func (s *ListDetectionRulesRequest) SetAlertTacticId(v string) *ListDetectionRulesRequest {
	s.AlertTacticId = &v
	return s
}

func (s *ListDetectionRulesRequest) SetAlertType(v string) *ListDetectionRulesRequest {
	s.AlertType = &v
	return s
}

func (s *ListDetectionRulesRequest) SetDetectionExpressionType(v string) *ListDetectionRulesRequest {
	s.DetectionExpressionType = &v
	return s
}

func (s *ListDetectionRulesRequest) SetDetectionRuleId(v string) *ListDetectionRulesRequest {
	s.DetectionRuleId = &v
	return s
}

func (s *ListDetectionRulesRequest) SetDetectionRuleIds(v []*string) *ListDetectionRulesRequest {
	s.DetectionRuleIds = v
	return s
}

func (s *ListDetectionRulesRequest) SetDetectionRuleName(v string) *ListDetectionRulesRequest {
	s.DetectionRuleName = &v
	return s
}

func (s *ListDetectionRulesRequest) SetDetectionRuleStatus(v string) *ListDetectionRulesRequest {
	s.DetectionRuleStatus = &v
	return s
}

func (s *ListDetectionRulesRequest) SetDetectionRuleType(v string) *ListDetectionRulesRequest {
	s.DetectionRuleType = &v
	return s
}

func (s *ListDetectionRulesRequest) SetIncidentAggregationType(v string) *ListDetectionRulesRequest {
	s.IncidentAggregationType = &v
	return s
}

func (s *ListDetectionRulesRequest) SetLang(v string) *ListDetectionRulesRequest {
	s.Lang = &v
	return s
}

func (s *ListDetectionRulesRequest) SetLogCategoryId(v string) *ListDetectionRulesRequest {
	s.LogCategoryId = &v
	return s
}

func (s *ListDetectionRulesRequest) SetLogSchemaId(v string) *ListDetectionRulesRequest {
	s.LogSchemaId = &v
	return s
}

func (s *ListDetectionRulesRequest) SetMaxResults(v int32) *ListDetectionRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDetectionRulesRequest) SetNextToken(v string) *ListDetectionRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDetectionRulesRequest) SetOrderDirection(v string) *ListDetectionRulesRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListDetectionRulesRequest) SetOrderFieldName(v string) *ListDetectionRulesRequest {
	s.OrderFieldName = &v
	return s
}

func (s *ListDetectionRulesRequest) SetPageNumber(v int32) *ListDetectionRulesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDetectionRulesRequest) SetPageSize(v int32) *ListDetectionRulesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDetectionRulesRequest) SetRegionId(v string) *ListDetectionRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDetectionRulesRequest) SetRoleFor(v int64) *ListDetectionRulesRequest {
	s.RoleFor = &v
	return s
}

func (s *ListDetectionRulesRequest) Validate() error {
	return dara.Validate(s)
}
