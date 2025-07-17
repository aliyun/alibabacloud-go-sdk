// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataQualityRule interface {
	dara.Model
	String() string
	GoString() string
	SetCheckingConfig(v *DataQualityRuleCheckingConfig) *DataQualityRule
	GetCheckingConfig() *DataQualityRuleCheckingConfig
	SetDescription(v string) *DataQualityRule
	GetDescription() *string
	SetEnabled(v bool) *DataQualityRule
	GetEnabled() *bool
	SetErrorHandlers(v []*DataQualityRuleErrorHandlers) *DataQualityRule
	GetErrorHandlers() []*DataQualityRuleErrorHandlers
	SetId(v int64) *DataQualityRule
	GetId() *int64
	SetName(v string) *DataQualityRule
	GetName() *string
	SetProjectId(v int64) *DataQualityRule
	GetProjectId() *int64
	SetSamplingConfig(v *DataQualityRuleSamplingConfig) *DataQualityRule
	GetSamplingConfig() *DataQualityRuleSamplingConfig
	SetSeverity(v string) *DataQualityRule
	GetSeverity() *string
	SetTarget(v *DataQualityRuleTarget) *DataQualityRule
	GetTarget() *DataQualityRuleTarget
	SetTemplateCode(v string) *DataQualityRule
	GetTemplateCode() *string
	SetTenantId(v int64) *DataQualityRule
	GetTenantId() *int64
}

type DataQualityRule struct {
	CheckingConfig *DataQualityRuleCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// example:
	//
	// this is a odps _sql task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Enabled       *bool                           `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ErrorHandlers []*DataQualityRuleErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 表不能为空
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 100
	ProjectId      *int64                         `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SamplingConfig *DataQualityRuleSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// example:
	//
	// High
	Severity *string                `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Target   *DataQualityRuleTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// example:
	//
	// SYSTEM:user_defined_sql
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// 1
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DataQualityRule) String() string {
	return dara.Prettify(s)
}

func (s DataQualityRule) GoString() string {
	return s.String()
}

func (s *DataQualityRule) GetCheckingConfig() *DataQualityRuleCheckingConfig {
	return s.CheckingConfig
}

func (s *DataQualityRule) GetDescription() *string {
	return s.Description
}

func (s *DataQualityRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *DataQualityRule) GetErrorHandlers() []*DataQualityRuleErrorHandlers {
	return s.ErrorHandlers
}

func (s *DataQualityRule) GetId() *int64 {
	return s.Id
}

func (s *DataQualityRule) GetName() *string {
	return s.Name
}

func (s *DataQualityRule) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DataQualityRule) GetSamplingConfig() *DataQualityRuleSamplingConfig {
	return s.SamplingConfig
}

func (s *DataQualityRule) GetSeverity() *string {
	return s.Severity
}

func (s *DataQualityRule) GetTarget() *DataQualityRuleTarget {
	return s.Target
}

func (s *DataQualityRule) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DataQualityRule) GetTenantId() *int64 {
	return s.TenantId
}

func (s *DataQualityRule) SetCheckingConfig(v *DataQualityRuleCheckingConfig) *DataQualityRule {
	s.CheckingConfig = v
	return s
}

func (s *DataQualityRule) SetDescription(v string) *DataQualityRule {
	s.Description = &v
	return s
}

func (s *DataQualityRule) SetEnabled(v bool) *DataQualityRule {
	s.Enabled = &v
	return s
}

func (s *DataQualityRule) SetErrorHandlers(v []*DataQualityRuleErrorHandlers) *DataQualityRule {
	s.ErrorHandlers = v
	return s
}

func (s *DataQualityRule) SetId(v int64) *DataQualityRule {
	s.Id = &v
	return s
}

func (s *DataQualityRule) SetName(v string) *DataQualityRule {
	s.Name = &v
	return s
}

func (s *DataQualityRule) SetProjectId(v int64) *DataQualityRule {
	s.ProjectId = &v
	return s
}

func (s *DataQualityRule) SetSamplingConfig(v *DataQualityRuleSamplingConfig) *DataQualityRule {
	s.SamplingConfig = v
	return s
}

func (s *DataQualityRule) SetSeverity(v string) *DataQualityRule {
	s.Severity = &v
	return s
}

func (s *DataQualityRule) SetTarget(v *DataQualityRuleTarget) *DataQualityRule {
	s.Target = v
	return s
}

func (s *DataQualityRule) SetTemplateCode(v string) *DataQualityRule {
	s.TemplateCode = &v
	return s
}

func (s *DataQualityRule) SetTenantId(v int64) *DataQualityRule {
	s.TenantId = &v
	return s
}

func (s *DataQualityRule) Validate() error {
	return dara.Validate(s)
}

type DataQualityRuleCheckingConfig struct {
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string                                  `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	Thresholds              *DataQualityRuleCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
	// example:
	//
	// Fixed
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityRuleCheckingConfig) String() string {
	return dara.Prettify(s)
}

func (s DataQualityRuleCheckingConfig) GoString() string {
	return s.String()
}

func (s *DataQualityRuleCheckingConfig) GetReferencedSamplesFilter() *string {
	return s.ReferencedSamplesFilter
}

func (s *DataQualityRuleCheckingConfig) GetThresholds() *DataQualityRuleCheckingConfigThresholds {
	return s.Thresholds
}

func (s *DataQualityRuleCheckingConfig) GetType() *string {
	return s.Type
}

func (s *DataQualityRuleCheckingConfig) SetReferencedSamplesFilter(v string) *DataQualityRuleCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *DataQualityRuleCheckingConfig) SetThresholds(v *DataQualityRuleCheckingConfigThresholds) *DataQualityRuleCheckingConfig {
	s.Thresholds = v
	return s
}

func (s *DataQualityRuleCheckingConfig) SetType(v string) *DataQualityRuleCheckingConfig {
	s.Type = &v
	return s
}

func (s *DataQualityRuleCheckingConfig) Validate() error {
	return dara.Validate(s)
}

type DataQualityRuleCheckingConfigThresholds struct {
	Critical *DataQualityRuleCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Expected *DataQualityRuleCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	Warned   *DataQualityRuleCheckingConfigThresholdsWarned   `json:"Warned,omitempty" xml:"Warned,omitempty" type:"Struct"`
}

func (s DataQualityRuleCheckingConfigThresholds) String() string {
	return dara.Prettify(s)
}

func (s DataQualityRuleCheckingConfigThresholds) GoString() string {
	return s.String()
}

func (s *DataQualityRuleCheckingConfigThresholds) GetCritical() *DataQualityRuleCheckingConfigThresholdsCritical {
	return s.Critical
}

func (s *DataQualityRuleCheckingConfigThresholds) GetExpected() *DataQualityRuleCheckingConfigThresholdsExpected {
	return s.Expected
}

func (s *DataQualityRuleCheckingConfigThresholds) GetWarned() *DataQualityRuleCheckingConfigThresholdsWarned {
	return s.Warned
}

func (s *DataQualityRuleCheckingConfigThresholds) SetCritical(v *DataQualityRuleCheckingConfigThresholdsCritical) *DataQualityRuleCheckingConfigThresholds {
	s.Critical = v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholds) SetExpected(v *DataQualityRuleCheckingConfigThresholdsExpected) *DataQualityRuleCheckingConfigThresholds {
	s.Expected = v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholds) SetWarned(v *DataQualityRuleCheckingConfigThresholdsWarned) *DataQualityRuleCheckingConfigThresholds {
	s.Warned = v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholds) Validate() error {
	return dara.Validate(s)
}

type DataQualityRuleCheckingConfigThresholdsCritical struct {
	// example:
	//
	// 波动率类型规则使用，通过表达式来表示波动阈值。如：波动上升大于0.01： $checkValue > 0.01  波动下降大于0.01：$checkValue < -0.01   波动率绝对值：abs($checkValue) > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataQualityRuleCheckingConfigThresholdsCritical) String() string {
	return dara.Prettify(s)
}

func (s DataQualityRuleCheckingConfigThresholdsCritical) GoString() string {
	return s.String()
}

func (s *DataQualityRuleCheckingConfigThresholdsCritical) GetExpression() *string {
	return s.Expression
}

func (s *DataQualityRuleCheckingConfigThresholdsCritical) GetOperator() *string {
	return s.Operator
}

func (s *DataQualityRuleCheckingConfigThresholdsCritical) GetValue() *string {
	return s.Value
}

func (s *DataQualityRuleCheckingConfigThresholdsCritical) SetExpression(v string) *DataQualityRuleCheckingConfigThresholdsCritical {
	s.Expression = &v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholdsCritical) SetOperator(v string) *DataQualityRuleCheckingConfigThresholdsCritical {
	s.Operator = &v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholdsCritical) SetValue(v string) *DataQualityRuleCheckingConfigThresholdsCritical {
	s.Value = &v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholdsCritical) Validate() error {
	return dara.Validate(s)
}

type DataQualityRuleCheckingConfigThresholdsExpected struct {
	// example:
	//
	// 波动率类型规则使用，通过表达式来表示波动阈值。如：波动上升大于0.01： $checkValue > 0.01  波动下降大于0.01：$checkValue < -0.01   波动率绝对值：abs($checkValue) > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataQualityRuleCheckingConfigThresholdsExpected) String() string {
	return dara.Prettify(s)
}

func (s DataQualityRuleCheckingConfigThresholdsExpected) GoString() string {
	return s.String()
}

func (s *DataQualityRuleCheckingConfigThresholdsExpected) GetExpression() *string {
	return s.Expression
}

func (s *DataQualityRuleCheckingConfigThresholdsExpected) GetOperator() *string {
	return s.Operator
}

func (s *DataQualityRuleCheckingConfigThresholdsExpected) GetValue() *string {
	return s.Value
}

func (s *DataQualityRuleCheckingConfigThresholdsExpected) SetExpression(v string) *DataQualityRuleCheckingConfigThresholdsExpected {
	s.Expression = &v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholdsExpected) SetOperator(v string) *DataQualityRuleCheckingConfigThresholdsExpected {
	s.Operator = &v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholdsExpected) SetValue(v string) *DataQualityRuleCheckingConfigThresholdsExpected {
	s.Value = &v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholdsExpected) Validate() error {
	return dara.Validate(s)
}

type DataQualityRuleCheckingConfigThresholdsWarned struct {
	// example:
	//
	// 波动率类型规则使用，通过表达式来表示波动阈值。如：波动上升大于0.01： $checkValue > 0.01  波动下降大于0.01：$checkValue < -0.01   波动率绝对值：abs($checkValue) > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataQualityRuleCheckingConfigThresholdsWarned) String() string {
	return dara.Prettify(s)
}

func (s DataQualityRuleCheckingConfigThresholdsWarned) GoString() string {
	return s.String()
}

func (s *DataQualityRuleCheckingConfigThresholdsWarned) GetExpression() *string {
	return s.Expression
}

func (s *DataQualityRuleCheckingConfigThresholdsWarned) GetOperator() *string {
	return s.Operator
}

func (s *DataQualityRuleCheckingConfigThresholdsWarned) GetValue() *string {
	return s.Value
}

func (s *DataQualityRuleCheckingConfigThresholdsWarned) SetExpression(v string) *DataQualityRuleCheckingConfigThresholdsWarned {
	s.Expression = &v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholdsWarned) SetOperator(v string) *DataQualityRuleCheckingConfigThresholdsWarned {
	s.Operator = &v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholdsWarned) SetValue(v string) *DataQualityRuleCheckingConfigThresholdsWarned {
	s.Value = &v
	return s
}

func (s *DataQualityRuleCheckingConfigThresholdsWarned) Validate() error {
	return dara.Validate(s)
}

type DataQualityRuleErrorHandlers struct {
	// example:
	//
	// SELECT 	- FROM tb_api_log WHERE id IS NULL
	ErrorDataFilter *string `json:"ErrorDataFilter,omitempty" xml:"ErrorDataFilter,omitempty"`
	// example:
	//
	// SaveErrorData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityRuleErrorHandlers) String() string {
	return dara.Prettify(s)
}

func (s DataQualityRuleErrorHandlers) GoString() string {
	return s.String()
}

func (s *DataQualityRuleErrorHandlers) GetErrorDataFilter() *string {
	return s.ErrorDataFilter
}

func (s *DataQualityRuleErrorHandlers) GetType() *string {
	return s.Type
}

func (s *DataQualityRuleErrorHandlers) SetErrorDataFilter(v string) *DataQualityRuleErrorHandlers {
	s.ErrorDataFilter = &v
	return s
}

func (s *DataQualityRuleErrorHandlers) SetType(v string) *DataQualityRuleErrorHandlers {
	s.Type = &v
	return s
}

func (s *DataQualityRuleErrorHandlers) Validate() error {
	return dara.Validate(s)
}

type DataQualityRuleSamplingConfig struct {
	// example:
	//
	// Min
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// example:
	//
	// { "Columns": [ "id", "name" ] }
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// example:
	//
	// id IS NULL
	SamplingFilter *string `json:"SamplingFilter,omitempty" xml:"SamplingFilter,omitempty"`
	// example:
	//
	// SET odps.sql.udf.timeout=600s;
	SettingConfig *string `json:"SettingConfig,omitempty" xml:"SettingConfig,omitempty"`
}

func (s DataQualityRuleSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s DataQualityRuleSamplingConfig) GoString() string {
	return s.String()
}

func (s *DataQualityRuleSamplingConfig) GetMetric() *string {
	return s.Metric
}

func (s *DataQualityRuleSamplingConfig) GetMetricParameters() *string {
	return s.MetricParameters
}

func (s *DataQualityRuleSamplingConfig) GetSamplingFilter() *string {
	return s.SamplingFilter
}

func (s *DataQualityRuleSamplingConfig) GetSettingConfig() *string {
	return s.SettingConfig
}

func (s *DataQualityRuleSamplingConfig) SetMetric(v string) *DataQualityRuleSamplingConfig {
	s.Metric = &v
	return s
}

func (s *DataQualityRuleSamplingConfig) SetMetricParameters(v string) *DataQualityRuleSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *DataQualityRuleSamplingConfig) SetSamplingFilter(v string) *DataQualityRuleSamplingConfig {
	s.SamplingFilter = &v
	return s
}

func (s *DataQualityRuleSamplingConfig) SetSettingConfig(v string) *DataQualityRuleSamplingConfig {
	s.SettingConfig = &v
	return s
}

func (s *DataQualityRuleSamplingConfig) Validate() error {
	return dara.Validate(s)
}

type DataQualityRuleTarget struct {
	// example:
	//
	// maxcompute
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// example:
	//
	// ds=$[yyyymmdd-1]
	PartitionSpec *string `json:"PartitionSpec,omitempty" xml:"PartitionSpec,omitempty"`
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// example:
	//
	// Table
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityRuleTarget) String() string {
	return dara.Prettify(s)
}

func (s DataQualityRuleTarget) GoString() string {
	return s.String()
}

func (s *DataQualityRuleTarget) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *DataQualityRuleTarget) GetPartitionSpec() *string {
	return s.PartitionSpec
}

func (s *DataQualityRuleTarget) GetTableGuid() *string {
	return s.TableGuid
}

func (s *DataQualityRuleTarget) GetType() *string {
	return s.Type
}

func (s *DataQualityRuleTarget) SetDatabaseType(v string) *DataQualityRuleTarget {
	s.DatabaseType = &v
	return s
}

func (s *DataQualityRuleTarget) SetPartitionSpec(v string) *DataQualityRuleTarget {
	s.PartitionSpec = &v
	return s
}

func (s *DataQualityRuleTarget) SetTableGuid(v string) *DataQualityRuleTarget {
	s.TableGuid = &v
	return s
}

func (s *DataQualityRuleTarget) SetType(v string) *DataQualityRuleTarget {
	s.Type = &v
	return s
}

func (s *DataQualityRuleTarget) Validate() error {
	return dara.Validate(s)
}
