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
	// The check settings for sample data.
	CheckingConfig *UpdateDataQualityRuleRequestCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// The description of the rule. The description can be up to 500 characters in length.
	//
	// example:
	//
	// this is a odps _sql task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the rule.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The operations that you can perform after the rule-based check fails.
	ErrorHandlers []*UpdateDataQualityRuleRequestErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the rule. The name can be up to 255 characters in length and can contain digits, letters, and punctuation marks.
	//
	// example:
	//
	// The table cannot be empty.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sampling settings.
	SamplingConfig *UpdateDataQualityRuleRequestSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
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
	// The ID of the template used by the rule.
	//
	// example:
	//
	// system::user_defined
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
	return dara.Validate(s)
}

type UpdateDataQualityRuleRequestCheckingConfig struct {
	// The method that is used to query the referenced samples. To obtain some types of thresholds, you need to query reference values. In this example, an expression is used to specify the query method of referenced samples.
	//
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// The threshold settings.
	Thresholds *UpdateDataQualityRuleRequestCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type UpdateDataQualityRuleRequestCheckingConfigThresholds struct {
	// The threshold settings for critical alerts.
	Critical *UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	// The expected threshold setting.
	Expected *UpdateDataQualityRuleRequestCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	// The threshold settings for normal alerts.
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
	return dara.Validate(s)
}

type UpdateDataQualityRuleRequestCheckingConfigThresholdsCritical struct {
	// The threshold expression.
	//
	// The volatility type rule must use an expression to represent the volatility threshold. For example:
	//
	// - Fluctuation rise greater than 0.01: $checkValue > 0.01
	//
	// - Fluctuation drop greater than 0.01:$checkValue < -0.01
	//
	// - Absolute volatility: abs($checkValue) > 0.01
	//
	// You can also use expressions to configure thresholds for fixed-Value rules. If you configure them at the same time, the expression priority is higher than Operator and Value.
	//
	// example:
	//
	// $checkValue > 0.05
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator. Valid values:
	//
	// 	- \\>
	//
	// 	- \\>=
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
	// The volatility type rule must use an expression to represent the volatility threshold. For example:
	//
	// - Fluctuation rise greater than 0.01: $checkValue > 0.01
	//
	// - Fluctuation drop greater than 0.01:$checkValue < -0.01
	//
	// - Absolute volatility: abs($checkValue) > 0.01
	//
	// You can also use expressions to configure thresholds for fixed-Value rules. If you configure them at the same time, the expression priority is higher than Operator and Value.
	//
	// example:
	//
	// $checkValue <= 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator. Valid values:
	//
	// 	- \\>
	//
	// 	- \\>=
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
	// The volatility type rule must use an expression to represent the volatility threshold. For example:
	//
	// - Fluctuation rise greater than 0.01: $checkValue > 0.01
	//
	// - Fluctuation drop greater than 0.01:$checkValue < -0.01
	//
	// - Absolute volatility: abs($checkValue) > 0.01
	//
	// You can also use expressions to configure thresholds for fixed-Value rules. If you configure them at the same time, the expression priority is higher than Operator and Value.
	//
	// example:
	//
	// $checkValue > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The comparison operator. Valid values:
	//
	// 	- \\>
	//
	// 	- \\>=
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
	// The metrics used for sampling. You can leave this parameter empty if you use a rule template. Valid values:
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
	// 	- DistinctPercent: the proportion of the number of unique values of the field after deduplication to the number of rows in the table.
	//
	// 	- DuplicatedCount: the number of duplicated values of the field.
	//
	// 	- DuplicatedPercent: the proportion of the number of duplicated values of the field to the number of rows in the table.
	//
	// 	- TableSize: the table size.
	//
	// 	- NullValueCount: the number of rows in which the field value is null.
	//
	// 	- NullValuePercent: the proportion of the number of rows in which the field value is null to the number of rows in the table.
	//
	// 	- GroupCount: the field value and the number of rows for each field value.
	//
	// 	- CountNotIn: the number of rows in which the field values are different from the referenced values that you specified in the rule.
	//
	// 	- CountDistinctNotIn: the number of unique values that are different from the referenced values that you specified in the rule after deduplication.
	//
	// 	- UserDefinedSql: indicates that data is sampled by executing custom SQL statements.
	//
	// example:
	//
	// Min
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The parameters required for sampling.
	//
	// example:
	//
	// { "Columns": [ "id", "name" ] , "SQL": "select count(1) from table;"}
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
