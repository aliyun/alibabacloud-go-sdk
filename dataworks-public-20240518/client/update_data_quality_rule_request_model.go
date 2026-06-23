// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckingConfig(v *UpdateDataQualityRuleRequestCheckingConfig) *UpdateDataQualityRuleRequest
	GetCheckingConfig() *UpdateDataQualityRuleRequestCheckingConfig
	SetDescription(v string) *UpdateDataQualityRuleRequest
	GetDescription() *string
	SetEnabled(v bool) *UpdateDataQualityRuleRequest
	GetEnabled() *bool
	SetErrorHandlers(v []*UpdateDataQualityRuleRequestErrorHandlers) *UpdateDataQualityRuleRequest
	GetErrorHandlers() []*UpdateDataQualityRuleRequestErrorHandlers
	SetId(v int64) *UpdateDataQualityRuleRequest
	GetId() *int64
	SetName(v string) *UpdateDataQualityRuleRequest
	GetName() *string
	SetProjectId(v int64) *UpdateDataQualityRuleRequest
	GetProjectId() *int64
	SetSamplingConfig(v *UpdateDataQualityRuleRequestSamplingConfig) *UpdateDataQualityRuleRequest
	GetSamplingConfig() *UpdateDataQualityRuleRequestSamplingConfig
	SetSeverity(v string) *UpdateDataQualityRuleRequest
	GetSeverity() *string
	SetTemplateCode(v string) *UpdateDataQualityRuleRequest
	GetTemplateCode() *string
}

type UpdateDataQualityRuleRequest struct {
	// The sample verification settings.
	CheckingConfig *UpdateDataQualityRuleRequestCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// The rule description. The maximum length is 500 characters.
	//
	// example:
	//
	// this is a odps _sql task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether the rule is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The list of issue handlers for data quality rule verification.
	ErrorHandlers []*UpdateDataQualityRuleRequestErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The rule name. The name can be a combination of digits, English letters, Chinese characters, and half-width or full-width punctuation. The maximum length is 255 characters.
	//
	// example:
	//
	// The table cannot be empty.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the DataWorks workspace. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace Settings page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The settings required for sample collection.
	SamplingConfig *UpdateDataQualityRuleRequestSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// The severity level of the rule for the business (corresponding to strong/weak rules on the page). Valid values:
	//
	// - Normal
	//
	// - High
	//
	// example:
	//
	// High
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The unique identifier of the rule template referenced by the rule.
	//
	// example:
	//
	// SYSTEM:table:table_count:fixed
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s UpdateDataQualityRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleRequest) GetCheckingConfig() *UpdateDataQualityRuleRequestCheckingConfig {
	return s.CheckingConfig
}

func (s *UpdateDataQualityRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDataQualityRuleRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateDataQualityRuleRequest) GetErrorHandlers() []*UpdateDataQualityRuleRequestErrorHandlers {
	return s.ErrorHandlers
}

func (s *UpdateDataQualityRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateDataQualityRuleRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDataQualityRuleRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateDataQualityRuleRequest) GetSamplingConfig() *UpdateDataQualityRuleRequestSamplingConfig {
	return s.SamplingConfig
}

func (s *UpdateDataQualityRuleRequest) GetSeverity() *string {
	return s.Severity
}

func (s *UpdateDataQualityRuleRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *UpdateDataQualityRuleRequest) SetCheckingConfig(v *UpdateDataQualityRuleRequestCheckingConfig) *UpdateDataQualityRuleRequest {
	s.CheckingConfig = v
	return s
}

func (s *UpdateDataQualityRuleRequest) SetDescription(v string) *UpdateDataQualityRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateDataQualityRuleRequest) SetEnabled(v bool) *UpdateDataQualityRuleRequest {
	s.Enabled = &v
	return s
}

func (s *UpdateDataQualityRuleRequest) SetErrorHandlers(v []*UpdateDataQualityRuleRequestErrorHandlers) *UpdateDataQualityRuleRequest {
	s.ErrorHandlers = v
	return s
}

func (s *UpdateDataQualityRuleRequest) SetId(v int64) *UpdateDataQualityRuleRequest {
	s.Id = &v
	return s
}

func (s *UpdateDataQualityRuleRequest) SetName(v string) *UpdateDataQualityRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateDataQualityRuleRequest) SetProjectId(v int64) *UpdateDataQualityRuleRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateDataQualityRuleRequest) SetSamplingConfig(v *UpdateDataQualityRuleRequestSamplingConfig) *UpdateDataQualityRuleRequest {
	s.SamplingConfig = v
	return s
}

func (s *UpdateDataQualityRuleRequest) SetSeverity(v string) *UpdateDataQualityRuleRequest {
	s.Severity = &v
	return s
}

func (s *UpdateDataQualityRuleRequest) SetTemplateCode(v string) *UpdateDataQualityRuleRequest {
	s.TemplateCode = &v
	return s
}

func (s *UpdateDataQualityRuleRequest) Validate() error {
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
	return nil
}

type UpdateDataQualityRuleRequestCheckingConfig struct {
	// Some types of thresholds require querying reference samples and then aggregating the values of those reference samples to derive the threshold used for comparison. An expression is used here to indicate how the reference samples are queried.
	//
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// The threshold settings.
	Thresholds *UpdateDataQualityRuleRequestCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
	// The threshold calculation method. This parameter is not required when a template is used.
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

func (s UpdateDataQualityRuleRequestCheckingConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleRequestCheckingConfig) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleRequestCheckingConfig) GetReferencedSamplesFilter() *string {
	return s.ReferencedSamplesFilter
}

func (s *UpdateDataQualityRuleRequestCheckingConfig) GetThresholds() *UpdateDataQualityRuleRequestCheckingConfigThresholds {
	return s.Thresholds
}

func (s *UpdateDataQualityRuleRequestCheckingConfig) GetType() *string {
	return s.Type
}

func (s *UpdateDataQualityRuleRequestCheckingConfig) SetReferencedSamplesFilter(v string) *UpdateDataQualityRuleRequestCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *UpdateDataQualityRuleRequestCheckingConfig) SetThresholds(v *UpdateDataQualityRuleRequestCheckingConfigThresholds) *UpdateDataQualityRuleRequestCheckingConfig {
	s.Thresholds = v
	return s
}

func (s *UpdateDataQualityRuleRequestCheckingConfig) SetType(v string) *UpdateDataQualityRuleRequestCheckingConfig {
	s.Type = &v
	return s
}

func (s *UpdateDataQualityRuleRequestCheckingConfig) Validate() error {
	if s.Thresholds != nil {
		if err := s.Thresholds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDataQualityRuleRequestCheckingConfigThresholds struct {
	// The threshold settings for critical warnings.
	Critical *UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	// The expected threshold settings.
	Expected *UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	// The threshold settings for normal warnings.
	Warned *UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned `json:"Warned,omitempty" xml:"Warned,omitempty" type:"Struct"`
}

func (s UpdateDataQualityRuleRequestCheckingConfigThresholds) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleRequestCheckingConfigThresholds) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholds) GetCritical() *UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical {
	return s.Critical
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholds) GetExpected() *UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected {
	return s.Expected
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholds) GetWarned() *UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned {
	return s.Warned
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholds) SetCritical(v *UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical) *UpdateDataQualityRuleRequestCheckingConfigThresholds {
	s.Critical = v
	return s
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholds) SetExpected(v *UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected) *UpdateDataQualityRuleRequestCheckingConfigThresholds {
	s.Expected = v
	return s
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholds) SetWarned(v *UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned) *UpdateDataQualityRuleRequestCheckingConfigThresholds {
	s.Warned = v
	return s
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholds) Validate() error {
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

type UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical struct {
	// The threshold expression.
	//
	// Fluctuation-type rules must use an expression to represent the fluctuation threshold. Examples:
	//
	// - Upward fluctuation greater than 0.01: $checkValue > 0.01
	//
	// - Downward fluctuation greater than 0.01: $checkValue < -0.01
	//
	// - Absolute fluctuation rate: abs($checkValue) > 0.01
	//
	// Fixed-value rules can also use an expression to configure the threshold. If both are configured, the expression takes precedence over Operator and Value.
	//
	// example:
	//
	// $checkValue > 0.05
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator.
	//
	// - \\>
	//
	// - \\>=
	//
	// - <
	//
	// - <=
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

func (s UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical) GetOperator() *string {
	return s.Operator
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical) GetValue() *string {
	return s.Value
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical) SetExpression(v string) *UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical {
	s.Expression = &v
	return s
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical) SetOperator(v string) *UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical {
	s.Operator = &v
	return s
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical) SetValue(v string) *UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical {
	s.Value = &v
	return s
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected struct {
	// The threshold expression.
	//
	// Fluctuation-type rules must use an expression to represent the fluctuation threshold. Examples:
	//
	// - Upward fluctuation greater than 0.01: $checkValue > 0.01
	//
	// - Downward fluctuation greater than 0.01: $checkValue < -0.01
	//
	// - Absolute fluctuation rate: abs($checkValue) > 0.01
	//
	// Fixed-value rules can also use an expression to configure the threshold. If both are configured, the expression takes precedence over Operator and Value.
	//
	// example:
	//
	// $checkValue <= 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator.
	//
	// - \\>
	//
	// - \\>=
	//
	// - <
	//
	// - <=
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

func (s UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected) GetOperator() *string {
	return s.Operator
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected) GetValue() *string {
	return s.Value
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected) SetExpression(v string) *UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected {
	s.Expression = &v
	return s
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected) SetOperator(v string) *UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected {
	s.Operator = &v
	return s
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected) SetValue(v string) *UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected {
	s.Value = &v
	return s
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned struct {
	// The threshold expression.
	//
	// Fluctuation-type rules must use an expression to represent the fluctuation threshold. Examples:
	//
	// - Upward fluctuation greater than 0.01: $checkValue > 0.01
	//
	// - Downward fluctuation greater than 0.01: $checkValue < -0.01
	//
	// - Absolute fluctuation rate: abs($checkValue) > 0.01
	//
	// Fixed-value rules can also use an expression to configure the threshold. If both are configured, the expression takes precedence over Operator and Value.
	//
	// example:
	//
	// $checkValue > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator.
	//
	// - \\>
	//
	// - \\>=
	//
	// - <
	//
	// - <=
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

func (s UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned) GetExpression() *string {
	return s.Expression
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned) GetOperator() *string {
	return s.Operator
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned) GetValue() *string {
	return s.Value
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned) SetExpression(v string) *UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned {
	s.Expression = &v
	return s
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned) SetOperator(v string) *UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned {
	s.Operator = &v
	return s
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned) SetValue(v string) *UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned {
	s.Value = &v
	return s
}

func (s *UpdateDataQualityRuleRequestCheckingConfigThresholdsWarned) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityRuleRequestErrorHandlers struct {
	// For a custom SQL rule, you must specify the SQL used to filter problematic data.
	//
	// example:
	//
	// SELECT 	- FROM tb_api_log WHERE id IS NULL
	ErrorDataFilter *string `json:"ErrorDataFilter,omitempty" xml:"ErrorDataFilter,omitempty"`
	// The handler type.
	//
	// - SaveErrorData
	//
	// example:
	//
	// SaveErrorData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateDataQualityRuleRequestErrorHandlers) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleRequestErrorHandlers) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleRequestErrorHandlers) GetErrorDataFilter() *string {
	return s.ErrorDataFilter
}

func (s *UpdateDataQualityRuleRequestErrorHandlers) GetType() *string {
	return s.Type
}

func (s *UpdateDataQualityRuleRequestErrorHandlers) SetErrorDataFilter(v string) *UpdateDataQualityRuleRequestErrorHandlers {
	s.ErrorDataFilter = &v
	return s
}

func (s *UpdateDataQualityRuleRequestErrorHandlers) SetType(v string) *UpdateDataQualityRuleRequestErrorHandlers {
	s.Type = &v
	return s
}

func (s *UpdateDataQualityRuleRequestErrorHandlers) Validate() error {
	return dara.Validate(s)
}

type UpdateDataQualityRuleRequestSamplingConfig struct {
	// The name of the metric to sample. This parameter is not required when a template is used.
	//
	// - Count: the number of rows in the table.
	//
	// - Min: the minimum value of the field.
	//
	// - Max: the maximum value of the field.
	//
	// - Avg: the average value of the field.
	//
	// - DistinctCount: the number of distinct values in the field.
	//
	// - DistinctPercent: the ratio of the number of distinct values in the field to the total number of rows.
	//
	// - DuplicatedCount: the number of duplicate values in the field.
	//
	// - DuplicatedPercent: the ratio of the number of duplicate values in the field to the total number of rows.
	//
	// - TableSize: the size of the table.
	//
	// - NullValueCount: the number of rows in which the field is null.
	//
	// - NullValuePercent: the percentage of rows in which the field is null.
	//
	// - GroupCount: the number of data rows for each value after aggregation by field value.
	//
	// - CountNotIn: the number of rows that do not match the enumerated values.
	//
	// - CountDistinctNotIn: the number of distinct values that do not match the enumerated values.
	//
	// - UserDefinedSql: sample collection by using custom SQL.
	//
	// example:
	//
	// Min
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The parameters required for sample collection.
	//
	// example:
	//
	// { "Columns": [ "id", "name" ] , "SQL": "select count(1) from table;"}
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// The condition used to apply secondary filtering on data that is not of interest during sampling. The maximum length is 16,777,215 characters.
	//
	// example:
	//
	// id IS NULL
	SamplingFilter *string `json:"SamplingFilter,omitempty" xml:"SamplingFilter,omitempty"`
	// The runtime parameter setting statements to be inserted and executed before the actual sampling statements. The maximum length is 1,000 characters. Only MaxCompute is supported.
	//
	// example:
	//
	// SET odps.sql.udf.timeout=600s;
	//
	// SET odps.sql.python.version=cp27;
	SettingConfig *string `json:"SettingConfig,omitempty" xml:"SettingConfig,omitempty"`
}

func (s UpdateDataQualityRuleRequestSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityRuleRequestSamplingConfig) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityRuleRequestSamplingConfig) GetMetric() *string {
	return s.Metric
}

func (s *UpdateDataQualityRuleRequestSamplingConfig) GetMetricParameters() *string {
	return s.MetricParameters
}

func (s *UpdateDataQualityRuleRequestSamplingConfig) GetSamplingFilter() *string {
	return s.SamplingFilter
}

func (s *UpdateDataQualityRuleRequestSamplingConfig) GetSettingConfig() *string {
	return s.SettingConfig
}

func (s *UpdateDataQualityRuleRequestSamplingConfig) SetMetric(v string) *UpdateDataQualityRuleRequestSamplingConfig {
	s.Metric = &v
	return s
}

func (s *UpdateDataQualityRuleRequestSamplingConfig) SetMetricParameters(v string) *UpdateDataQualityRuleRequestSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *UpdateDataQualityRuleRequestSamplingConfig) SetSamplingFilter(v string) *UpdateDataQualityRuleRequestSamplingConfig {
	s.SamplingFilter = &v
	return s
}

func (s *UpdateDataQualityRuleRequestSamplingConfig) SetSettingConfig(v string) *UpdateDataQualityRuleRequestSamplingConfig {
	s.SettingConfig = &v
	return s
}

func (s *UpdateDataQualityRuleRequestSamplingConfig) Validate() error {
	return dara.Validate(s)
}
