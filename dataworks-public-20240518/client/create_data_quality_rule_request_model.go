// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckingConfig(v *CreateDataQualityRuleRequestCheckingConfig) *CreateDataQualityRuleRequest
	GetCheckingConfig() *CreateDataQualityRuleRequestCheckingConfig
	SetDescription(v string) *CreateDataQualityRuleRequest
	GetDescription() *string
	SetEnabled(v bool) *CreateDataQualityRuleRequest
	GetEnabled() *bool
	SetErrorHandlers(v []*CreateDataQualityRuleRequestErrorHandlers) *CreateDataQualityRuleRequest
	GetErrorHandlers() []*CreateDataQualityRuleRequestErrorHandlers
	SetName(v string) *CreateDataQualityRuleRequest
	GetName() *string
	SetProjectId(v int64) *CreateDataQualityRuleRequest
	GetProjectId() *int64
	SetSamplingConfig(v *CreateDataQualityRuleRequestSamplingConfig) *CreateDataQualityRuleRequest
	GetSamplingConfig() *CreateDataQualityRuleRequestSamplingConfig
	SetSeverity(v string) *CreateDataQualityRuleRequest
	GetSeverity() *string
	SetTarget(v *CreateDataQualityRuleRequestTarget) *CreateDataQualityRuleRequest
	GetTarget() *CreateDataQualityRuleRequestTarget
	SetTemplateCode(v string) *CreateDataQualityRuleRequest
	GetTemplateCode() *string
}

type CreateDataQualityRuleRequest struct {
	// The sample check settings.
	CheckingConfig *CreateDataQualityRuleRequestCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// The description of the rule. The maximum length is 500 characters.
	//
	// example:
	//
	// this is a odps _sql task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the data quality rule.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of issue handlers for the data quality rule check.
	ErrorHandlers []*CreateDataQualityRuleRequestErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// The name of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// The table cannot be empty.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10726
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The settings required for sample collection.
	SamplingConfig *CreateDataQualityRuleRequestSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// The severity of the rule for the business (corresponding to the strong/weak rule on the page). Valid values:
	//
	// - Normal
	//
	// - High
	//
	// example:
	//
	// Normal
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The object monitored by the rule.
	Target *CreateDataQualityRuleRequestTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The unique identifier of the rule template that the rule references.
	//
	// example:
	//
	// SYSTEM:table:table_count:fixed
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s CreateDataQualityRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleRequest) GetCheckingConfig() *CreateDataQualityRuleRequestCheckingConfig {
	return s.CheckingConfig
}

func (s *CreateDataQualityRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataQualityRuleRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateDataQualityRuleRequest) GetErrorHandlers() []*CreateDataQualityRuleRequestErrorHandlers {
	return s.ErrorHandlers
}

func (s *CreateDataQualityRuleRequest) GetName() *string {
	return s.Name
}

func (s *CreateDataQualityRuleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDataQualityRuleRequest) GetSamplingConfig() *CreateDataQualityRuleRequestSamplingConfig {
	return s.SamplingConfig
}

func (s *CreateDataQualityRuleRequest) GetSeverity() *string {
	return s.Severity
}

func (s *CreateDataQualityRuleRequest) GetTarget() *CreateDataQualityRuleRequestTarget {
	return s.Target
}

func (s *CreateDataQualityRuleRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *CreateDataQualityRuleRequest) SetCheckingConfig(v *CreateDataQualityRuleRequestCheckingConfig) *CreateDataQualityRuleRequest {
	s.CheckingConfig = v
	return s
}

func (s *CreateDataQualityRuleRequest) SetDescription(v string) *CreateDataQualityRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateDataQualityRuleRequest) SetEnabled(v bool) *CreateDataQualityRuleRequest {
	s.Enabled = &v
	return s
}

func (s *CreateDataQualityRuleRequest) SetErrorHandlers(v []*CreateDataQualityRuleRequestErrorHandlers) *CreateDataQualityRuleRequest {
	s.ErrorHandlers = v
	return s
}

func (s *CreateDataQualityRuleRequest) SetName(v string) *CreateDataQualityRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateDataQualityRuleRequest) SetProjectId(v int64) *CreateDataQualityRuleRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDataQualityRuleRequest) SetSamplingConfig(v *CreateDataQualityRuleRequestSamplingConfig) *CreateDataQualityRuleRequest {
	s.SamplingConfig = v
	return s
}

func (s *CreateDataQualityRuleRequest) SetSeverity(v string) *CreateDataQualityRuleRequest {
	s.Severity = &v
	return s
}

func (s *CreateDataQualityRuleRequest) SetTarget(v *CreateDataQualityRuleRequestTarget) *CreateDataQualityRuleRequest {
	s.Target = v
	return s
}

func (s *CreateDataQualityRuleRequest) SetTemplateCode(v string) *CreateDataQualityRuleRequest {
	s.TemplateCode = &v
	return s
}

func (s *CreateDataQualityRuleRequest) Validate() error {
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

type CreateDataQualityRuleRequestCheckingConfig struct {
	// Some types of thresholds require querying for reference samples and then summarizing the values of the reference samples to derive the threshold to be compared. An expression is used here to represent the query method for the reference samples.
	//
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// The threshold settings.
	Thresholds *CreateDataQualityRuleRequestCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
	// The threshold calculation method. You do not need to specify this parameter when a template is used.
	//
	// - Fixed
	//
	// - Fluctation
	//
	// - FluctationDiscreate
	//
	// - Auto
	//
	// - Average
	//
	// - Variance
	//
	// example:
	//
	// Fixed
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataQualityRuleRequestCheckingConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleRequestCheckingConfig) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleRequestCheckingConfig) GetReferencedSamplesFilter() *string {
	return s.ReferencedSamplesFilter
}

func (s *CreateDataQualityRuleRequestCheckingConfig) GetThresholds() *CreateDataQualityRuleRequestCheckingConfigThresholds {
	return s.Thresholds
}

func (s *CreateDataQualityRuleRequestCheckingConfig) GetType() *string {
	return s.Type
}

func (s *CreateDataQualityRuleRequestCheckingConfig) SetReferencedSamplesFilter(v string) *CreateDataQualityRuleRequestCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *CreateDataQualityRuleRequestCheckingConfig) SetThresholds(v *CreateDataQualityRuleRequestCheckingConfigThresholds) *CreateDataQualityRuleRequestCheckingConfig {
	s.Thresholds = v
	return s
}

func (s *CreateDataQualityRuleRequestCheckingConfig) SetType(v string) *CreateDataQualityRuleRequestCheckingConfig {
	s.Type = &v
	return s
}

func (s *CreateDataQualityRuleRequestCheckingConfig) Validate() error {
	if s.Thresholds != nil {
		if err := s.Thresholds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataQualityRuleRequestCheckingConfigThresholds struct {
	// The threshold settings for a critical warning.
	Critical *CreateDataQualityRuleRequestCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	// The expected threshold settings.
	Expected *CreateDataQualityRuleRequestCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	// The threshold settings for a normal warning.
	Warned *CreateDataQualityRuleRequestCheckingConfigThresholdsWarned `json:"Warned,omitempty" xml:"Warned,omitempty" type:"Struct"`
}

func (s CreateDataQualityRuleRequestCheckingConfigThresholds) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleRequestCheckingConfigThresholds) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholds) GetCritical() *CreateDataQualityRuleRequestCheckingConfigThresholdsCritical {
	return s.Critical
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholds) GetExpected() *CreateDataQualityRuleRequestCheckingConfigThresholdsExpected {
	return s.Expected
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholds) GetWarned() *CreateDataQualityRuleRequestCheckingConfigThresholdsWarned {
	return s.Warned
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholds) SetCritical(v *CreateDataQualityRuleRequestCheckingConfigThresholdsCritical) *CreateDataQualityRuleRequestCheckingConfigThresholds {
	s.Critical = v
	return s
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholds) SetExpected(v *CreateDataQualityRuleRequestCheckingConfigThresholdsExpected) *CreateDataQualityRuleRequestCheckingConfigThresholds {
	s.Expected = v
	return s
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholds) SetWarned(v *CreateDataQualityRuleRequestCheckingConfigThresholdsWarned) *CreateDataQualityRuleRequestCheckingConfigThresholds {
	s.Warned = v
	return s
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholds) Validate() error {
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

type CreateDataQualityRuleRequestCheckingConfigThresholdsCritical struct {
	// The threshold expression.
	//
	// For a fluctuation-type rule, you must use an expression to represent the fluctuation threshold. Examples:
	//
	// - Fluctuation rises above 0.01: $checkValue > 0.01
	//
	// - Fluctuation drops below -0.01: $checkValue < -0.01
	//
	// - Absolute value of the fluctuation: abs($checkValue) > 0.01
	//
	// For a fixed-value type rule, you can also use an expression to configure the threshold. If both are configured at the same time, the expression takes precedence over Operator and Value.
	//
	// example:
	//
	// $checkValue > 0.05
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator:
	//
	// - \\>
	//
	// - \\>=
	//
	// - \\<
	//
	// - \\<=
	//
	// - !=
	//
	// - =
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

func (s CreateDataQualityRuleRequestCheckingConfigThresholdsCritical) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleRequestCheckingConfigThresholdsCritical) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsCritical) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsCritical) GetOperator() *string {
	return s.Operator
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsCritical) GetValue() *string {
	return s.Value
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsCritical) SetExpression(v string) *CreateDataQualityRuleRequestCheckingConfigThresholdsCritical {
	s.Expression = &v
	return s
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsCritical) SetOperator(v string) *CreateDataQualityRuleRequestCheckingConfigThresholdsCritical {
	s.Operator = &v
	return s
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsCritical) SetValue(v string) *CreateDataQualityRuleRequestCheckingConfigThresholdsCritical {
	s.Value = &v
	return s
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsCritical) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityRuleRequestCheckingConfigThresholdsExpected struct {
	// The threshold expression.
	//
	// For a fluctuation-type rule, you must use an expression to represent the fluctuation threshold. Examples:
	//
	// - Fluctuation rises above 0.01: $checkValue > 0.01
	//
	// - Fluctuation drops below -0.01: $checkValue < -0.01
	//
	// - Absolute value of the fluctuation: abs($checkValue) > 0.01
	//
	// For a fixed-value type rule, you can also use an expression to configure the threshold. If both are configured at the same time, the expression takes precedence over Operator and Value.
	//
	// example:
	//
	// $checkValue <= 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator:
	//
	// - \\>
	//
	// - \\>=
	//
	// - \\<
	//
	// - \\<=
	//
	// - !=
	//
	// - =
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

func (s CreateDataQualityRuleRequestCheckingConfigThresholdsExpected) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleRequestCheckingConfigThresholdsExpected) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsExpected) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsExpected) GetOperator() *string {
	return s.Operator
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsExpected) GetValue() *string {
	return s.Value
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsExpected) SetExpression(v string) *CreateDataQualityRuleRequestCheckingConfigThresholdsExpected {
	s.Expression = &v
	return s
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsExpected) SetOperator(v string) *CreateDataQualityRuleRequestCheckingConfigThresholdsExpected {
	s.Operator = &v
	return s
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsExpected) SetValue(v string) *CreateDataQualityRuleRequestCheckingConfigThresholdsExpected {
	s.Value = &v
	return s
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsExpected) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityRuleRequestCheckingConfigThresholdsWarned struct {
	// The threshold expression.
	//
	// For a fluctuation-type rule, you must use an expression to represent the fluctuation threshold. Examples:
	//
	// - Fluctuation rises above 0.01: $checkValue > 0.01
	//
	// - Fluctuation drops below -0.01: $checkValue < -0.01
	//
	// - Absolute value of the fluctuation: abs($checkValue) > 0.01
	//
	// For a fixed-value type rule, you can also use an expression to configure the threshold. If both are configured at the same time, the expression takes precedence over Operator and Value.
	//
	// example:
	//
	// $checkValue > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator:
	//
	// - \\>
	//
	// - \\>=
	//
	// - \\<
	//
	// - \\<=
	//
	// - !=
	//
	// - =
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

func (s CreateDataQualityRuleRequestCheckingConfigThresholdsWarned) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleRequestCheckingConfigThresholdsWarned) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsWarned) GetExpression() *string {
	return s.Expression
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsWarned) GetOperator() *string {
	return s.Operator
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsWarned) GetValue() *string {
	return s.Value
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsWarned) SetExpression(v string) *CreateDataQualityRuleRequestCheckingConfigThresholdsWarned {
	s.Expression = &v
	return s
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsWarned) SetOperator(v string) *CreateDataQualityRuleRequestCheckingConfigThresholdsWarned {
	s.Operator = &v
	return s
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsWarned) SetValue(v string) *CreateDataQualityRuleRequestCheckingConfigThresholdsWarned {
	s.Value = &v
	return s
}

func (s *CreateDataQualityRuleRequestCheckingConfigThresholdsWarned) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityRuleRequestErrorHandlers struct {
	// For a custom SQL rule, you must specify an SQL statement to filter problematic data.
	//
	// example:
	//
	// SELECT 	- FROM tb_api_log WHERE id IS NULL
	ErrorDataFilter *string `json:"ErrorDataFilter,omitempty" xml:"ErrorDataFilter,omitempty"`
	// The handler type:
	//
	// - SaveErrorData
	//
	// example:
	//
	// SaveErrorData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataQualityRuleRequestErrorHandlers) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleRequestErrorHandlers) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleRequestErrorHandlers) GetErrorDataFilter() *string {
	return s.ErrorDataFilter
}

func (s *CreateDataQualityRuleRequestErrorHandlers) GetType() *string {
	return s.Type
}

func (s *CreateDataQualityRuleRequestErrorHandlers) SetErrorDataFilter(v string) *CreateDataQualityRuleRequestErrorHandlers {
	s.ErrorDataFilter = &v
	return s
}

func (s *CreateDataQualityRuleRequestErrorHandlers) SetType(v string) *CreateDataQualityRuleRequestErrorHandlers {
	s.Type = &v
	return s
}

func (s *CreateDataQualityRuleRequestErrorHandlers) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityRuleRequestSamplingConfig struct {
	// The name of the metric to be sampled. You do not need to specify this parameter when a template is used.
	//
	// - Count: the number of rows in the table.
	//
	// - Min: the minimum value of the field.
	//
	// - Max: the maximum value of the field.
	//
	// - Avg: the average value of the field.
	//
	// - DistinctCount: the number of distinct values of the field.
	//
	// - DistinctPercent: the ratio of the number of distinct values of the field to the number of data rows.
	//
	// - DuplicatedCount: the number of duplicate values of the field.
	//
	// - DuplicatedPercent: the ratio of the number of duplicate values of the field to the number of data rows.
	//
	// - TableSize: the size of the table.
	//
	// - NullValueCount: the number of rows in which the field is null.
	//
	// - NullValuePercent: the ratio of rows in which the field is null.
	//
	// - GroupCount: the values aggregated by field value and the corresponding number of data rows for each value.
	//
	// - CountNotIn: the number of rows whose enum values do not match.
	//
	// - CountDistinctNotIn: the number of distinct values that do not match the enum values.
	//
	// - UserDefinedSql: collects samples by using a custom SQL statement.
	//
	// example:
	//
	// Count
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The parameters required during sample collection.
	//
	// example:
	//
	// { "Columns": [ "id", "name" ] , "SQL": "select count(1) from table;"}
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// The condition used to perform a secondary filter on the irrelevant data during sampling. The maximum length is 16777215 characters.
	//
	// example:
	//
	// id IS NULL
	SamplingFilter *string `json:"SamplingFilter,omitempty" xml:"SamplingFilter,omitempty"`
	// The runtime parameter setting statements to be inserted and executed before the sampling statement is executed. The maximum length is 1000 characters. Only MaxCompute is supported.
	//
	// example:
	//
	// SET odps.sql.udf.timeout=600s;
	//
	// SET odps.sql.python.version=cp27;
	SettingConfig *string `json:"SettingConfig,omitempty" xml:"SettingConfig,omitempty"`
}

func (s CreateDataQualityRuleRequestSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleRequestSamplingConfig) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleRequestSamplingConfig) GetMetric() *string {
	return s.Metric
}

func (s *CreateDataQualityRuleRequestSamplingConfig) GetMetricParameters() *string {
	return s.MetricParameters
}

func (s *CreateDataQualityRuleRequestSamplingConfig) GetSamplingFilter() *string {
	return s.SamplingFilter
}

func (s *CreateDataQualityRuleRequestSamplingConfig) GetSettingConfig() *string {
	return s.SettingConfig
}

func (s *CreateDataQualityRuleRequestSamplingConfig) SetMetric(v string) *CreateDataQualityRuleRequestSamplingConfig {
	s.Metric = &v
	return s
}

func (s *CreateDataQualityRuleRequestSamplingConfig) SetMetricParameters(v string) *CreateDataQualityRuleRequestSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *CreateDataQualityRuleRequestSamplingConfig) SetSamplingFilter(v string) *CreateDataQualityRuleRequestSamplingConfig {
	s.SamplingFilter = &v
	return s
}

func (s *CreateDataQualityRuleRequestSamplingConfig) SetSettingConfig(v string) *CreateDataQualityRuleRequestSamplingConfig {
	s.SettingConfig = &v
	return s
}

func (s *CreateDataQualityRuleRequestSamplingConfig) Validate() error {
	return dara.Validate(s)
}

type CreateDataQualityRuleRequestTarget struct {
	// For a table-type dataset, the type of database to which the table belongs.
	//
	// - maxcompute
	//
	// - emr
	//
	// - cdh
	//
	// - hologres
	//
	// - analyticdb_for_postgresql
	//
	// - analyticdb_for_mysql
	//
	// - starrocks
	//
	// example:
	//
	// maxcompute
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The partition settings of the partitioned table.
	//
	// example:
	//
	// ds=$[yyyymmdd-1]
	PartitionSpec *string `json:"PartitionSpec,omitempty" xml:"PartitionSpec,omitempty"`
	// The unique ID of the table in Data Map on which the rule operates.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The type of the monitored object.
	//
	// - Table
	//
	// example:
	//
	// Table
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDataQualityRuleRequestTarget) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityRuleRequestTarget) GoString() string {
	return s.String()
}

func (s *CreateDataQualityRuleRequestTarget) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *CreateDataQualityRuleRequestTarget) GetPartitionSpec() *string {
	return s.PartitionSpec
}

func (s *CreateDataQualityRuleRequestTarget) GetTableGuid() *string {
	return s.TableGuid
}

func (s *CreateDataQualityRuleRequestTarget) GetType() *string {
	return s.Type
}

func (s *CreateDataQualityRuleRequestTarget) SetDatabaseType(v string) *CreateDataQualityRuleRequestTarget {
	s.DatabaseType = &v
	return s
}

func (s *CreateDataQualityRuleRequestTarget) SetPartitionSpec(v string) *CreateDataQualityRuleRequestTarget {
	s.PartitionSpec = &v
	return s
}

func (s *CreateDataQualityRuleRequestTarget) SetTableGuid(v string) *CreateDataQualityRuleRequestTarget {
	s.TableGuid = &v
	return s
}

func (s *CreateDataQualityRuleRequestTarget) SetType(v string) *CreateDataQualityRuleRequestTarget {
	s.Type = &v
	return s
}

func (s *CreateDataQualityRuleRequestTarget) Validate() error {
	return dara.Validate(s)
}
