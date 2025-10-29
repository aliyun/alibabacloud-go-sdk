// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataQualityResult interface {
	dara.Model
	String() string
	GoString() string
	SetDetails(v []*DataQualityResultDetails) *DataQualityResult
	GetDetails() []*DataQualityResultDetails
	SetId(v int64) *DataQualityResult
	GetId() *int64
	SetRule(v *DataQualityResultRule) *DataQualityResult
	GetRule() *DataQualityResultRule
	SetSample(v string) *DataQualityResult
	GetSample() *string
	SetStatus(v string) *DataQualityResult
	GetStatus() *string
	SetTaskInstanceId(v int64) *DataQualityResult
	GetTaskInstanceId() *int64
}

type DataQualityResult struct {
	Details []*DataQualityResultDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// example:
	//
	// 10001
	Id   *int64                 `json:"Id,omitempty" xml:"Id,omitempty"`
	Rule *DataQualityResultRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
	// example:
	//
	// [   {     "gender": "male",     "_count": 100   }, {     "gender": "female",     "_count": 100   } ]
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// example:
	//
	// Passed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 20001
	TaskInstanceId *int64 `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
}

func (s DataQualityResult) String() string {
	return dara.Prettify(s)
}

func (s DataQualityResult) GoString() string {
	return s.String()
}

func (s *DataQualityResult) GetDetails() []*DataQualityResultDetails {
	return s.Details
}

func (s *DataQualityResult) GetId() *int64 {
	return s.Id
}

func (s *DataQualityResult) GetRule() *DataQualityResultRule {
	return s.Rule
}

func (s *DataQualityResult) GetSample() *string {
	return s.Sample
}

func (s *DataQualityResult) GetStatus() *string {
	return s.Status
}

func (s *DataQualityResult) GetTaskInstanceId() *int64 {
	return s.TaskInstanceId
}

func (s *DataQualityResult) SetDetails(v []*DataQualityResultDetails) *DataQualityResult {
	s.Details = v
	return s
}

func (s *DataQualityResult) SetId(v int64) *DataQualityResult {
	s.Id = &v
	return s
}

func (s *DataQualityResult) SetRule(v *DataQualityResultRule) *DataQualityResult {
	s.Rule = v
	return s
}

func (s *DataQualityResult) SetSample(v string) *DataQualityResult {
	s.Sample = &v
	return s
}

func (s *DataQualityResult) SetStatus(v string) *DataQualityResult {
	s.Status = &v
	return s
}

func (s *DataQualityResult) SetTaskInstanceId(v int64) *DataQualityResult {
	s.TaskInstanceId = &v
	return s
}

func (s *DataQualityResult) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Rule != nil {
		if err := s.Rule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataQualityResultDetails struct {
	// example:
	//
	// 100.0
	CheckedValue *string `json:"CheckedValue,omitempty" xml:"CheckedValue,omitempty"`
	// example:
	//
	// 0.0
	ReferencedValue *string `json:"ReferencedValue,omitempty" xml:"ReferencedValue,omitempty"`
	// example:
	//
	// Passed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DataQualityResultDetails) String() string {
	return dara.Prettify(s)
}

func (s DataQualityResultDetails) GoString() string {
	return s.String()
}

func (s *DataQualityResultDetails) GetCheckedValue() *string {
	return s.CheckedValue
}

func (s *DataQualityResultDetails) GetReferencedValue() *string {
	return s.ReferencedValue
}

func (s *DataQualityResultDetails) GetStatus() *string {
	return s.Status
}

func (s *DataQualityResultDetails) SetCheckedValue(v string) *DataQualityResultDetails {
	s.CheckedValue = &v
	return s
}

func (s *DataQualityResultDetails) SetReferencedValue(v string) *DataQualityResultDetails {
	s.ReferencedValue = &v
	return s
}

func (s *DataQualityResultDetails) SetStatus(v string) *DataQualityResultDetails {
	s.Status = &v
	return s
}

func (s *DataQualityResultDetails) Validate() error {
	return dara.Validate(s)
}

type DataQualityResultRule struct {
	CheckingConfig *DataQualityResultRuleCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// example:
	//
	// this is a odps _sql task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	Enabled       *bool                                 `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	ErrorHandlers []*DataQualityResultRuleErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// example:
	//
	// 100001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 表不能为空
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 100
	ProjectId      *int64                               `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	SamplingConfig *DataQualityResultRuleSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// example:
	//
	// High
	Severity *string                      `json:"Severity,omitempty" xml:"Severity,omitempty"`
	Target   *DataQualityResultRuleTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// example:
	//
	// SYSTEM:user_defined_sql
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// example:
	//
	// 1
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DataQualityResultRule) String() string {
	return dara.Prettify(s)
}

func (s DataQualityResultRule) GoString() string {
	return s.String()
}

func (s *DataQualityResultRule) GetCheckingConfig() *DataQualityResultRuleCheckingConfig {
	return s.CheckingConfig
}

func (s *DataQualityResultRule) GetDescription() *string {
	return s.Description
}

func (s *DataQualityResultRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *DataQualityResultRule) GetErrorHandlers() []*DataQualityResultRuleErrorHandlers {
	return s.ErrorHandlers
}

func (s *DataQualityResultRule) GetId() *int64 {
	return s.Id
}

func (s *DataQualityResultRule) GetName() *string {
	return s.Name
}

func (s *DataQualityResultRule) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *DataQualityResultRule) GetSamplingConfig() *DataQualityResultRuleSamplingConfig {
	return s.SamplingConfig
}

func (s *DataQualityResultRule) GetSeverity() *string {
	return s.Severity
}

func (s *DataQualityResultRule) GetTarget() *DataQualityResultRuleTarget {
	return s.Target
}

func (s *DataQualityResultRule) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DataQualityResultRule) GetTenantId() *int64 {
	return s.TenantId
}

func (s *DataQualityResultRule) SetCheckingConfig(v *DataQualityResultRuleCheckingConfig) *DataQualityResultRule {
	s.CheckingConfig = v
	return s
}

func (s *DataQualityResultRule) SetDescription(v string) *DataQualityResultRule {
	s.Description = &v
	return s
}

func (s *DataQualityResultRule) SetEnabled(v bool) *DataQualityResultRule {
	s.Enabled = &v
	return s
}

func (s *DataQualityResultRule) SetErrorHandlers(v []*DataQualityResultRuleErrorHandlers) *DataQualityResultRule {
	s.ErrorHandlers = v
	return s
}

func (s *DataQualityResultRule) SetId(v int64) *DataQualityResultRule {
	s.Id = &v
	return s
}

func (s *DataQualityResultRule) SetName(v string) *DataQualityResultRule {
	s.Name = &v
	return s
}

func (s *DataQualityResultRule) SetProjectId(v int64) *DataQualityResultRule {
	s.ProjectId = &v
	return s
}

func (s *DataQualityResultRule) SetSamplingConfig(v *DataQualityResultRuleSamplingConfig) *DataQualityResultRule {
	s.SamplingConfig = v
	return s
}

func (s *DataQualityResultRule) SetSeverity(v string) *DataQualityResultRule {
	s.Severity = &v
	return s
}

func (s *DataQualityResultRule) SetTarget(v *DataQualityResultRuleTarget) *DataQualityResultRule {
	s.Target = v
	return s
}

func (s *DataQualityResultRule) SetTemplateCode(v string) *DataQualityResultRule {
	s.TemplateCode = &v
	return s
}

func (s *DataQualityResultRule) SetTenantId(v int64) *DataQualityResultRule {
	s.TenantId = &v
	return s
}

func (s *DataQualityResultRule) Validate() error {
	if s.CheckingConfig != nil {
		if err := s.CheckingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ErrorHandlers != nil {
		for _, item := range s.ErrorHandlers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SamplingConfig != nil {
		if err := s.SamplingConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataQualityResultRuleCheckingConfig struct {
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string                                        `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	Thresholds              *DataQualityResultRuleCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
	// example:
	//
	// Fixed
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityResultRuleCheckingConfig) String() string {
	return dara.Prettify(s)
}

func (s DataQualityResultRuleCheckingConfig) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleCheckingConfig) GetReferencedSamplesFilter() *string {
	return s.ReferencedSamplesFilter
}

func (s *DataQualityResultRuleCheckingConfig) GetThresholds() *DataQualityResultRuleCheckingConfigThresholds {
	return s.Thresholds
}

func (s *DataQualityResultRuleCheckingConfig) GetType() *string {
	return s.Type
}

func (s *DataQualityResultRuleCheckingConfig) SetReferencedSamplesFilter(v string) *DataQualityResultRuleCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *DataQualityResultRuleCheckingConfig) SetThresholds(v *DataQualityResultRuleCheckingConfigThresholds) *DataQualityResultRuleCheckingConfig {
	s.Thresholds = v
	return s
}

func (s *DataQualityResultRuleCheckingConfig) SetType(v string) *DataQualityResultRuleCheckingConfig {
	s.Type = &v
	return s
}

func (s *DataQualityResultRuleCheckingConfig) Validate() error {
	if s.Thresholds != nil {
		if err := s.Thresholds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataQualityResultRuleCheckingConfigThresholds struct {
	Critical *DataQualityResultRuleCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	Expected *DataQualityResultRuleCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	Warned   *DataQualityResultRuleCheckingConfigThresholdsWarned   `json:"Warned,omitempty" xml:"Warned,omitempty" type:"Struct"`
}

func (s DataQualityResultRuleCheckingConfigThresholds) String() string {
	return dara.Prettify(s)
}

func (s DataQualityResultRuleCheckingConfigThresholds) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleCheckingConfigThresholds) GetCritical() *DataQualityResultRuleCheckingConfigThresholdsCritical {
	return s.Critical
}

func (s *DataQualityResultRuleCheckingConfigThresholds) GetExpected() *DataQualityResultRuleCheckingConfigThresholdsExpected {
	return s.Expected
}

func (s *DataQualityResultRuleCheckingConfigThresholds) GetWarned() *DataQualityResultRuleCheckingConfigThresholdsWarned {
	return s.Warned
}

func (s *DataQualityResultRuleCheckingConfigThresholds) SetCritical(v *DataQualityResultRuleCheckingConfigThresholdsCritical) *DataQualityResultRuleCheckingConfigThresholds {
	s.Critical = v
	return s
}

func (s *DataQualityResultRuleCheckingConfigThresholds) SetExpected(v *DataQualityResultRuleCheckingConfigThresholdsExpected) *DataQualityResultRuleCheckingConfigThresholds {
	s.Expected = v
	return s
}

func (s *DataQualityResultRuleCheckingConfigThresholds) SetWarned(v *DataQualityResultRuleCheckingConfigThresholdsWarned) *DataQualityResultRuleCheckingConfigThresholds {
	s.Warned = v
	return s
}

func (s *DataQualityResultRuleCheckingConfigThresholds) Validate() error {
	if s.Critical != nil {
		if err := s.Critical.Validate(); err != nil {
			return err
		}
	}
	if s.Expected != nil {
		if err := s.Expected.Validate(); err != nil {
			return err
		}
	}
	if s.Warned != nil {
		if err := s.Warned.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataQualityResultRuleCheckingConfigThresholdsCritical struct {
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataQualityResultRuleCheckingConfigThresholdsCritical) String() string {
	return dara.Prettify(s)
}

func (s DataQualityResultRuleCheckingConfigThresholdsCritical) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleCheckingConfigThresholdsCritical) GetOperator() *string {
	return s.Operator
}

func (s *DataQualityResultRuleCheckingConfigThresholdsCritical) GetValue() *string {
	return s.Value
}

func (s *DataQualityResultRuleCheckingConfigThresholdsCritical) SetOperator(v string) *DataQualityResultRuleCheckingConfigThresholdsCritical {
	s.Operator = &v
	return s
}

func (s *DataQualityResultRuleCheckingConfigThresholdsCritical) SetValue(v string) *DataQualityResultRuleCheckingConfigThresholdsCritical {
	s.Value = &v
	return s
}

func (s *DataQualityResultRuleCheckingConfigThresholdsCritical) Validate() error {
	return dara.Validate(s)
}

type DataQualityResultRuleCheckingConfigThresholdsExpected struct {
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataQualityResultRuleCheckingConfigThresholdsExpected) String() string {
	return dara.Prettify(s)
}

func (s DataQualityResultRuleCheckingConfigThresholdsExpected) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleCheckingConfigThresholdsExpected) GetOperator() *string {
	return s.Operator
}

func (s *DataQualityResultRuleCheckingConfigThresholdsExpected) GetValue() *string {
	return s.Value
}

func (s *DataQualityResultRuleCheckingConfigThresholdsExpected) SetOperator(v string) *DataQualityResultRuleCheckingConfigThresholdsExpected {
	s.Operator = &v
	return s
}

func (s *DataQualityResultRuleCheckingConfigThresholdsExpected) SetValue(v string) *DataQualityResultRuleCheckingConfigThresholdsExpected {
	s.Value = &v
	return s
}

func (s *DataQualityResultRuleCheckingConfigThresholdsExpected) Validate() error {
	return dara.Validate(s)
}

type DataQualityResultRuleCheckingConfigThresholdsWarned struct {
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DataQualityResultRuleCheckingConfigThresholdsWarned) String() string {
	return dara.Prettify(s)
}

func (s DataQualityResultRuleCheckingConfigThresholdsWarned) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleCheckingConfigThresholdsWarned) GetOperator() *string {
	return s.Operator
}

func (s *DataQualityResultRuleCheckingConfigThresholdsWarned) GetValue() *string {
	return s.Value
}

func (s *DataQualityResultRuleCheckingConfigThresholdsWarned) SetOperator(v string) *DataQualityResultRuleCheckingConfigThresholdsWarned {
	s.Operator = &v
	return s
}

func (s *DataQualityResultRuleCheckingConfigThresholdsWarned) SetValue(v string) *DataQualityResultRuleCheckingConfigThresholdsWarned {
	s.Value = &v
	return s
}

func (s *DataQualityResultRuleCheckingConfigThresholdsWarned) Validate() error {
	return dara.Validate(s)
}

type DataQualityResultRuleErrorHandlers struct {
	// example:
	//
	// SELECT 	- FROM tb_api_log WHERE id IS NULL
	ErrorDataFilter *string `json:"ErrorDataFilter,omitempty" xml:"ErrorDataFilter,omitempty"`
	// example:
	//
	// SaveErrorData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DataQualityResultRuleErrorHandlers) String() string {
	return dara.Prettify(s)
}

func (s DataQualityResultRuleErrorHandlers) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleErrorHandlers) GetErrorDataFilter() *string {
	return s.ErrorDataFilter
}

func (s *DataQualityResultRuleErrorHandlers) GetType() *string {
	return s.Type
}

func (s *DataQualityResultRuleErrorHandlers) SetErrorDataFilter(v string) *DataQualityResultRuleErrorHandlers {
	s.ErrorDataFilter = &v
	return s
}

func (s *DataQualityResultRuleErrorHandlers) SetType(v string) *DataQualityResultRuleErrorHandlers {
	s.Type = &v
	return s
}

func (s *DataQualityResultRuleErrorHandlers) Validate() error {
	return dara.Validate(s)
}

type DataQualityResultRuleSamplingConfig struct {
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

func (s DataQualityResultRuleSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s DataQualityResultRuleSamplingConfig) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleSamplingConfig) GetMetric() *string {
	return s.Metric
}

func (s *DataQualityResultRuleSamplingConfig) GetMetricParameters() *string {
	return s.MetricParameters
}

func (s *DataQualityResultRuleSamplingConfig) GetSamplingFilter() *string {
	return s.SamplingFilter
}

func (s *DataQualityResultRuleSamplingConfig) GetSettingConfig() *string {
	return s.SettingConfig
}

func (s *DataQualityResultRuleSamplingConfig) SetMetric(v string) *DataQualityResultRuleSamplingConfig {
	s.Metric = &v
	return s
}

func (s *DataQualityResultRuleSamplingConfig) SetMetricParameters(v string) *DataQualityResultRuleSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *DataQualityResultRuleSamplingConfig) SetSamplingFilter(v string) *DataQualityResultRuleSamplingConfig {
	s.SamplingFilter = &v
	return s
}

func (s *DataQualityResultRuleSamplingConfig) SetSettingConfig(v string) *DataQualityResultRuleSamplingConfig {
	s.SettingConfig = &v
	return s
}

func (s *DataQualityResultRuleSamplingConfig) Validate() error {
	return dara.Validate(s)
}

type DataQualityResultRuleTarget struct {
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

func (s DataQualityResultRuleTarget) String() string {
	return dara.Prettify(s)
}

func (s DataQualityResultRuleTarget) GoString() string {
	return s.String()
}

func (s *DataQualityResultRuleTarget) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *DataQualityResultRuleTarget) GetPartitionSpec() *string {
	return s.PartitionSpec
}

func (s *DataQualityResultRuleTarget) GetTableGuid() *string {
	return s.TableGuid
}

func (s *DataQualityResultRuleTarget) GetType() *string {
	return s.Type
}

func (s *DataQualityResultRuleTarget) SetDatabaseType(v string) *DataQualityResultRuleTarget {
	s.DatabaseType = &v
	return s
}

func (s *DataQualityResultRuleTarget) SetPartitionSpec(v string) *DataQualityResultRuleTarget {
	s.PartitionSpec = &v
	return s
}

func (s *DataQualityResultRuleTarget) SetTableGuid(v string) *DataQualityResultRuleTarget {
	s.TableGuid = &v
	return s
}

func (s *DataQualityResultRuleTarget) SetType(v string) *DataQualityResultRuleTarget {
	s.Type = &v
	return s
}

func (s *DataQualityResultRuleTarget) Validate() error {
	return dara.Validate(s)
}
