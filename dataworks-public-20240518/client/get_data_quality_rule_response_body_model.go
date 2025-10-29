// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataQualityRule(v *GetDataQualityRuleResponseBodyDataQualityRule) *GetDataQualityRuleResponseBody
	GetDataQualityRule() *GetDataQualityRuleResponseBodyDataQualityRule
	SetRequestId(v string) *GetDataQualityRuleResponseBody
	GetRequestId() *string
}

type GetDataQualityRuleResponseBody struct {
	// The information about the rule.
	DataQualityRule *GetDataQualityRuleResponseBodyDataQualityRule `json:"DataQualityRule,omitempty" xml:"DataQualityRule,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataQualityRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleResponseBody) GetDataQualityRule() *GetDataQualityRuleResponseBodyDataQualityRule {
	return s.DataQualityRule
}

func (s *GetDataQualityRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataQualityRuleResponseBody) SetDataQualityRule(v *GetDataQualityRuleResponseBodyDataQualityRule) *GetDataQualityRuleResponseBody {
	s.DataQualityRule = v
	return s
}

func (s *GetDataQualityRuleResponseBody) SetRequestId(v string) *GetDataQualityRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataQualityRuleResponseBody) Validate() error {
	if s.DataQualityRule != nil {
		if err := s.DataQualityRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityRuleResponseBodyDataQualityRule struct {
	// The check settings for sample data.
	CheckingConfig *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
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
	ErrorHandlers []*GetDataQualityRuleResponseBodyDataQualityRuleErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// example:
	//
	// 16033
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The rule name.
	//
	// example:
	//
	// The table cannot be empty.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 1948
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sampling settings.
	SamplingConfig *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// Rule for the business level (corresponding to the strong and weak rules on the page), optional enumeration value:
	//
	// - Normal
	//
	// - High
	//
	// example:
	//
	// High
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The monitored object of the rule.
	Target *GetDataQualityRuleResponseBodyDataQualityRuleTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The ID of the template used by the rule.
	//
	// example:
	//
	// system::user_defined
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s GetDataQualityRuleResponseBodyDataQualityRule) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleResponseBodyDataQualityRule) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) GetCheckingConfig() *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig {
	return s.CheckingConfig
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) GetDescription() *string {
	return s.Description
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) GetErrorHandlers() []*GetDataQualityRuleResponseBodyDataQualityRuleErrorHandlers {
	return s.ErrorHandlers
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) GetName() *string {
	return s.Name
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) GetSamplingConfig() *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig {
	return s.SamplingConfig
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) GetSeverity() *string {
	return s.Severity
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) GetTarget() *GetDataQualityRuleResponseBodyDataQualityRuleTarget {
	return s.Target
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) SetCheckingConfig(v *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig) *GetDataQualityRuleResponseBodyDataQualityRule {
	s.CheckingConfig = v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) SetDescription(v string) *GetDataQualityRuleResponseBodyDataQualityRule {
	s.Description = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) SetEnabled(v bool) *GetDataQualityRuleResponseBodyDataQualityRule {
	s.Enabled = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) SetErrorHandlers(v []*GetDataQualityRuleResponseBodyDataQualityRuleErrorHandlers) *GetDataQualityRuleResponseBodyDataQualityRule {
	s.ErrorHandlers = v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) SetId(v int64) *GetDataQualityRuleResponseBodyDataQualityRule {
	s.Id = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) SetName(v string) *GetDataQualityRuleResponseBodyDataQualityRule {
	s.Name = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) SetProjectId(v int64) *GetDataQualityRuleResponseBodyDataQualityRule {
	s.ProjectId = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) SetSamplingConfig(v *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig) *GetDataQualityRuleResponseBodyDataQualityRule {
	s.SamplingConfig = v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) SetSeverity(v string) *GetDataQualityRuleResponseBodyDataQualityRule {
	s.Severity = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) SetTarget(v *GetDataQualityRuleResponseBodyDataQualityRuleTarget) *GetDataQualityRuleResponseBodyDataQualityRule {
	s.Target = v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) SetTemplateCode(v string) *GetDataQualityRuleResponseBodyDataQualityRule {
	s.TemplateCode = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRule) Validate() error {
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

type GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig struct {
	// The method that is used to query the referenced samples. To obtain some types of thresholds, you need to query reference values. In this example, an expression is used to indicate the query method of referenced samples.
	//
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// The threshold settings.
	Thresholds *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
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

func (s GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig) GetReferencedSamplesFilter() *string {
	return s.ReferencedSamplesFilter
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig) GetThresholds() *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds {
	return s.Thresholds
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig) GetType() *string {
	return s.Type
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig) SetReferencedSamplesFilter(v string) *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig) SetThresholds(v *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds) *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig {
	s.Thresholds = v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig) SetType(v string) *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig {
	s.Type = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfig) Validate() error {
	if s.Thresholds != nil {
		if err := s.Thresholds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds struct {
	// The threshold settings for critical alerts.
	Critical *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	// The expected threshold setting.
	Expected *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	// The threshold settings for normal alerts.
	Warned *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned `json:"Warned,omitempty" xml:"Warned,omitempty" type:"Struct"`
}

func (s GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds) GetCritical() *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical {
	return s.Critical
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds) GetExpected() *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected {
	return s.Expected
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds) GetWarned() *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned {
	return s.Warned
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds) SetCritical(v *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical) *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds {
	s.Critical = v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds) SetExpected(v *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected) *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds {
	s.Expected = v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds) SetWarned(v *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned) *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds {
	s.Warned = v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholds) Validate() error {
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

type GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical struct {
	// The threshold expression.
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

func (s GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical) GetExpression() *string {
	return s.Expression
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical) GetOperator() *string {
	return s.Operator
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical) GetValue() *string {
	return s.Value
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical) SetExpression(v string) *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical {
	s.Expression = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical) SetOperator(v string) *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical {
	s.Operator = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical) SetValue(v string) *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical {
	s.Value = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsCritical) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected struct {
	// The threshold expression.
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

func (s GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected) GetExpression() *string {
	return s.Expression
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected) GetOperator() *string {
	return s.Operator
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected) GetValue() *string {
	return s.Value
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected) SetExpression(v string) *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected {
	s.Expression = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected) SetOperator(v string) *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected {
	s.Operator = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected) SetValue(v string) *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected {
	s.Value = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsExpected) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned struct {
	// The threshold expression.
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

func (s GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned) GetExpression() *string {
	return s.Expression
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned) GetOperator() *string {
	return s.Operator
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned) GetValue() *string {
	return s.Value
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned) SetExpression(v string) *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned {
	s.Expression = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned) SetOperator(v string) *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned {
	s.Operator = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned) SetValue(v string) *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned {
	s.Value = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleCheckingConfigThresholdsWarned) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityRuleResponseBodyDataQualityRuleErrorHandlers struct {
	// The SQL statement that is used to filter failed tasks. If you define the rule by using custom SQL statements, you must specify an SQL statement to filter failed tasks.
	//
	// example:
	//
	// SELECT 	- FROM tb_api_log WHERE id IS NULL
	ErrorDataFilter *string `json:"ErrorDataFilter,omitempty" xml:"ErrorDataFilter,omitempty"`
	// Processor type:
	//
	// - SaveErrorData
	//
	// example:
	//
	// SaveErrorData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityRuleResponseBodyDataQualityRuleErrorHandlers) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleResponseBodyDataQualityRuleErrorHandlers) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleErrorHandlers) GetErrorDataFilter() *string {
	return s.ErrorDataFilter
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleErrorHandlers) GetType() *string {
	return s.Type
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleErrorHandlers) SetErrorDataFilter(v string) *GetDataQualityRuleResponseBodyDataQualityRuleErrorHandlers {
	s.ErrorDataFilter = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleErrorHandlers) SetType(v string) *GetDataQualityRuleResponseBodyDataQualityRuleErrorHandlers {
	s.Type = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleErrorHandlers) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig struct {
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
	// Max
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

func (s GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig) GetMetric() *string {
	return s.Metric
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig) GetMetricParameters() *string {
	return s.MetricParameters
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig) GetSamplingFilter() *string {
	return s.SamplingFilter
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig) GetSettingConfig() *string {
	return s.SettingConfig
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig) SetMetric(v string) *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig {
	s.Metric = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig) SetMetricParameters(v string) *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig) SetSamplingFilter(v string) *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig {
	s.SamplingFilter = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig) SetSettingConfig(v string) *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig {
	s.SettingConfig = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleSamplingConfig) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityRuleResponseBodyDataQualityRuleTarget struct {
	// The dataset of the table type. The database type to which the table belongs.
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
	// The configuration of the partitioned table.
	//
	// example:
	//
	// ds=$[yyyymmdd-1]
	PartitionSpec *string `json:"PartitionSpec,omitempty" xml:"PartitionSpec,omitempty"`
	// The ID of the table that is limited by the rule in Data Map.
	//
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// Monitoring object type
	//
	// - Table
	//
	// example:
	//
	// Table
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDataQualityRuleResponseBodyDataQualityRuleTarget) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityRuleResponseBodyDataQualityRuleTarget) GoString() string {
	return s.String()
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleTarget) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleTarget) GetPartitionSpec() *string {
	return s.PartitionSpec
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleTarget) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleTarget) GetType() *string {
	return s.Type
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleTarget) SetDatabaseType(v string) *GetDataQualityRuleResponseBodyDataQualityRuleTarget {
	s.DatabaseType = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleTarget) SetPartitionSpec(v string) *GetDataQualityRuleResponseBodyDataQualityRuleTarget {
	s.PartitionSpec = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleTarget) SetTableGuid(v string) *GetDataQualityRuleResponseBodyDataQualityRuleTarget {
	s.TableGuid = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleTarget) SetType(v string) *GetDataQualityRuleResponseBodyDataQualityRuleTarget {
	s.Type = &v
	return s
}

func (s *GetDataQualityRuleResponseBodyDataQualityRuleTarget) Validate() error {
	return dara.Validate(s)
}
