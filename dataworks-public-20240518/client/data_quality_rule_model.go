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
	// The check settings for sample data.
	CheckingConfig *DataQualityRuleCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// The description of the rule. The description can be up to 500 characters in length.
	//
	// example:
	//
	// this is a odps _sql task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the rule is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The operations that you can perform after the rule-based check fails.
	ErrorHandlers []*DataQualityRuleErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The rule name. The name can be up to 255 characters in length and can contain digits, letters, and punctuation marks.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The settings for sampling.
	SamplingConfig *DataQualityRuleSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// The strength of the rule. Valid values:
	//
	// 	- Normal
	//
	// 	- High
	//
	// example:
	//
	// High
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The monitored object of the rule.
	Target *DataQualityRuleTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The template used by the rule.
	//
	// example:
	//
	// system::user_defined
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The DataWorks tenant ID.
	//
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

type DataQualityRuleCheckingConfig struct {
	// The method that is used to query the referenced samples. To obtain some types of thresholds, you need to query reference values. In this example, an expression is used to indicate the query method of referenced samples.
	//
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// The threshold settings.
	Thresholds *DataQualityRuleCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
	// The threshold calculation method. Valid values:
	//
	// 	- Fixed
	//
	// 	- Fluctation
	//
	// 	- FluctationDiscreate
	//
	// 	- Auto
	//
	// 	- Average
	//
	// 	- Variance
	//
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
	if s.Thresholds != nil {
		if err := s.Thresholds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DataQualityRuleCheckingConfigThresholds struct {
	// The threshold settings for critical alerts.
	Critical *DataQualityRuleCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	// The expected threshold setting.
	Expected *DataQualityRuleCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	// The threshold settings for normal alerts.
	Warned *DataQualityRuleCheckingConfigThresholdsWarned `json:"Warned,omitempty" xml:"Warned,omitempty" type:"Struct"`
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

type DataQualityRuleCheckingConfigThresholdsCritical struct {
	// example:
	//
	// 波动率类型规则使用，通过表达式来表示波动阈值。如：波动上升大于0.01： $checkValue > 0.01  波动下降大于0.01：$checkValue < -0.01   波动率绝对值：abs($checkValue) > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator. Valid values:
	//
	// 	- />
	//
	// 	- />=
	//
	// 	- <
	//
	// 	- <=
	//
	// 	- !=
	//
	// 	- \\=
	//
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The threshold value.
	//
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
	// The comparison operator. Valid values:
	//
	// 	- />
	//
	// 	- />=
	//
	// 	- <
	//
	// 	- <=
	//
	// 	- !=
	//
	// 	- \\=
	//
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The threshold value.
	//
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
	// The comparison operator. Valid values:
	//
	// 	- />
	//
	// 	- />=
	//
	// 	- <
	//
	// 	- <=
	//
	// 	- !=
	//
	// 	- \\=
	//
	// example:
	//
	// >
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The threshold value.
	//
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
	// The SQL statement that is used to filter failed tasks. If the rule is defined by custom SQL statements, you must specify an SQL statement to filter failed tasks.
	//
	// example:
	//
	// SELECT 	- FROM tb_api_log WHERE id IS NULL
	ErrorDataFilter *string `json:"ErrorDataFilter,omitempty" xml:"ErrorDataFilter,omitempty"`
	// The type of the operation. Valid values:
	//
	// 	- SaveErrorData
	//
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
	// The metrics used for sampling. Valid values:
	//
	// 	- Count: the number of rows in the table.
	//
	// 	- Min: the minimum value of the field.
	//
	// 	- Max: the maximum value of the field.
	//
	// 	- Avg: the average value of the field.
	//
	// 	- DistinctCount: the number of unique values of the field after deduplication.
	//
	// 	- DistinctPercent: the percentage of the number of unique values of the field after deduplication to the number of rows in the table.
	//
	// 	- DuplicatedCount: the number of duplicated values in the field.
	//
	// 	- DuplicatedPercent: the percentage of the number of duplicated values of the field to the number of rows in the table.
	//
	// 	- TableSize: the table size.
	//
	// 	- NullValueCount: the number of rows in which the field is set to null.
	//
	// 	- NullValuePercent: the percentage of the number of rows in which the field is set to null to the number of rows in the table.
	//
	// 	- GroupCount: the field value and the number of rows for each field value.
	//
	// 	- CountNotIn: the number of rows in which the field values are different from the referenced values that you specified in the rule.
	//
	// 	- CountDistinctNotIn: the number of unique values that are different from the referenced values that you specified in the rule after deduplication.
	//
	// 	- UserDefinedSql: indicates that the data is sampled by executing custom SQL statements.
	//
	// example:
	//
	// Min
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The parameters required for sampling.
	//
	// example:
	//
	// { "Columns": [ "id", "name" ] }
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// The statements that are used to filter unnecessary data during sampling. The statements can be up to 16,777,215 characters in length.
	//
	// example:
	//
	// id IS NULL
	SamplingFilter *string `json:"SamplingFilter,omitempty" xml:"SamplingFilter,omitempty"`
	// The statements that are used to configure the parameters required for sampling before you execute the sampling statements. The statements can be up to 1,000 characters in length. Only the MaxCompute database is supported.
	//
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
	// The type of the database to which the table belongs. Valid values:
	//
	// 	- maxcompute
	//
	// 	- emr
	//
	// 	- cdh
	//
	// 	- hologres
	//
	// 	- analyticdb_for_postgresql
	//
	// 	- analyticdb_for_mysql
	//
	// 	- starrocks
	//
	// example:
	//
	// maxcompute
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The configuration of the partitioned table.
	//
	// example:
	//
	// ds=$[yyyymmdd-1]
	PartitionSpec *string `json:"PartitionSpec,omitempty" xml:"PartitionSpec,omitempty"`
	// The ID of the table in Data Map.
	//
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The type of the monitored object. Valid values:
	//
	// 	- Table
	//
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
