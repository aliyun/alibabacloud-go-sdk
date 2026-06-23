// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDataQualityRulesResponseBodyPagingInfo) *ListDataQualityRulesResponseBody
	GetPagingInfo() *ListDataQualityRulesResponseBodyPagingInfo
	SetRequestId(v string) *ListDataQualityRulesResponseBody
	GetRequestId() *string
}

type ListDataQualityRulesResponseBody struct {
	// Paginated query result of the rule list.
	PagingInfo *ListDataQualityRulesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// Request ID.
	//
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataQualityRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBody) GetPagingInfo() *ListDataQualityRulesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDataQualityRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataQualityRulesResponseBody) SetPagingInfo(v *ListDataQualityRulesResponseBodyPagingInfo) *ListDataQualityRulesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDataQualityRulesResponseBody) SetRequestId(v string) *ListDataQualityRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataQualityRulesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataQualityRulesResponseBodyPagingInfo struct {
	// Specific rule list.
	DataQualityRules []*ListDataQualityRulesResponseBodyPagingInfoDataQualityRules `json:"DataQualityRules,omitempty" xml:"DataQualityRules,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of entries.
	//
	// example:
	//
	// 294
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfo) GetDataQualityRules() []*ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	return s.DataQualityRules
}

func (s *ListDataQualityRulesResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityRulesResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityRulesResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataQualityRulesResponseBodyPagingInfo) SetDataQualityRules(v []*ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) *ListDataQualityRulesResponseBodyPagingInfo {
	s.DataQualityRules = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfo) SetPageNumber(v int32) *ListDataQualityRulesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfo) SetPageSize(v int32) *ListDataQualityRulesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfo) SetTotalCount(v int32) *ListDataQualityRulesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfo) Validate() error {
	if s.DataQualityRules != nil {
		for _, item := range s.DataQualityRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRules struct {
	// Sample validation settings.
	CheckingConfig *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// Rule description. Maximum length: 500 characters.
	//
	// example:
	//
	// this is a odps _sql task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether the data quality rule is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// List of issue handlers for data quality rule validation.
	ErrorHandlers []*ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// Rule ID.
	//
	// example:
	//
	// 22130
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Rule name.
	//
	// example:
	//
	// The table cannot be empty.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// DataWorks workspace ID.
	//
	// example:
	//
	// 100001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Settings required for sample collection.
	SamplingConfig *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// Severity level of the rule for the business (corresponding to strong/weak rules on the page). Valid enumerated values:
	//
	// - Normal
	//
	// - High
	//
	// example:
	//
	// High
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// Object monitored by the rule.
	Target *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// Unique identifier of the rule template referenced by the rule.
	//
	// example:
	//
	// SYSTEM:user_defined_sql
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) GetCheckingConfig() *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig {
	return s.CheckingConfig
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) GetDescription() *string {
	return s.Description
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) GetErrorHandlers() []*ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers {
	return s.ErrorHandlers
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) GetId() *int64 {
	return s.Id
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) GetName() *string {
	return s.Name
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) GetSamplingConfig() *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig {
	return s.SamplingConfig
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) GetSeverity() *string {
	return s.Severity
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) GetTarget() *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget {
	return s.Target
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetCheckingConfig(v *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.CheckingConfig = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetDescription(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.Description = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetEnabled(v bool) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.Enabled = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetErrorHandlers(v []*ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.ErrorHandlers = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetId(v int64) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.Id = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetName(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.Name = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetProjectId(v int64) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetSamplingConfig(v *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.SamplingConfig = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetSeverity(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.Severity = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetTarget(v *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.Target = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) SetTemplateCode(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules {
	s.TemplateCode = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRules) Validate() error {
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

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig struct {
	// Some types of thresholds require querying some reference samples, then aggregating the values of the reference samples to obtain the threshold for comparison. An expression is used here to represent the query method for the reference samples.
	//
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// Threshold settings.
	Thresholds *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
	// Threshold calculation method.
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

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) GetReferencedSamplesFilter() *string {
	return s.ReferencedSamplesFilter
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) GetThresholds() *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds {
	return s.Thresholds
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) GetType() *string {
	return s.Type
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) SetReferencedSamplesFilter(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) SetThresholds(v *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig {
	s.Thresholds = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) SetType(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig {
	s.Type = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig) Validate() error {
	if s.Thresholds != nil {
		if err := s.Thresholds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds struct {
	// Threshold settings for critical warnings.
	Critical *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	// Expected threshold settings.
	Expected *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	// Threshold settings for normal warnings.
	Warned *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned `json:"Warned,omitempty" xml:"Warned,omitempty" type:"Struct"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) GetCritical() *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical {
	return s.Critical
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) GetExpected() *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected {
	return s.Expected
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) GetWarned() *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned {
	return s.Warned
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) SetCritical(v *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds {
	s.Critical = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) SetExpected(v *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds {
	s.Expected = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) SetWarned(v *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds {
	s.Warned = v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds) Validate() error {
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

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical struct {
	// Threshold expression.
	//
	// Fluctuation rate type rules must use the expression method to indicate the fluctuation threshold. Examples:
	//
	// - Fluctuation increase greater than 0.01: $checkValue > 0.01
	//
	// - Fluctuation decrease greater than 0.01: $checkValue < -0.01
	//
	// - Absolute value of the fluctuation rate: abs($checkValue) > 0.01
	//
	// Fixed value type rules can also configure thresholds using expressions. If both are configured, the expression takes precedence over Operator and Value.
	//
	// example:
	//
	// $checkValue > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// Comparison operator.
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
	// Threshold value.
	//
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical) GetExpression() *string {
	return s.Expression
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical) GetOperator() *string {
	return s.Operator
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical) GetValue() *string {
	return s.Value
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical) SetExpression(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical {
	s.Expression = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical) SetOperator(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical {
	s.Operator = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical) SetValue(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical {
	s.Value = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected struct {
	// Threshold expression.
	//
	// Fluctuation rate type rules must use the expression method to indicate the fluctuation threshold. Examples:
	//
	// - Fluctuation increase greater than 0.01: $checkValue > 0.01
	//
	// - Fluctuation decrease greater than 0.01: $checkValue < -0.01
	//
	// - Absolute value of the fluctuation rate: abs($checkValue) > 0.01
	//
	// Fixed value type rules can also configure thresholds using expressions. If both are configured, the expression takes precedence over Operator and Value.
	//
	// example:
	//
	// $checkValue > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// Comparison operator.
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
	// Threshold value.
	//
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected) GetExpression() *string {
	return s.Expression
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected) GetOperator() *string {
	return s.Operator
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected) GetValue() *string {
	return s.Value
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected) SetExpression(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected {
	s.Expression = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected) SetOperator(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected {
	s.Operator = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected) SetValue(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected {
	s.Value = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned struct {
	// Threshold expression.
	//
	// Fluctuation rate type rules must use the expression method to indicate the fluctuation threshold. Examples:
	//
	// - Fluctuation increase greater than 0.01: $checkValue > 0.01
	//
	// - Fluctuation decrease greater than 0.01: $checkValue < -0.01
	//
	// - Absolute value of the fluctuation rate: abs($checkValue) > 0.01
	//
	// Fixed value type rules can also configure thresholds using expressions. If both are configured, the expression takes precedence over Operator and Value.
	//
	// example:
	//
	// $checkValue > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// Comparison operator.
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
	// Threshold value.
	//
	// example:
	//
	// 100.0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned) GetExpression() *string {
	return s.Expression
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned) GetOperator() *string {
	return s.Operator
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned) GetValue() *string {
	return s.Value
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned) SetExpression(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned {
	s.Expression = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned) SetOperator(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned {
	s.Operator = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned) SetValue(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned {
	s.Value = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsWarned) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers struct {
	// For custom SQL rules, the user needs to specify SQL to filter problem data.
	//
	// example:
	//
	// SELECT 	- FROM tb_api_log WHERE id IS NULL
	ErrorDataFilter *string `json:"ErrorDataFilter,omitempty" xml:"ErrorDataFilter,omitempty"`
	// Handler type:
	//
	// - SaveErrorData
	//
	// example:
	//
	// SaveErrorData
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers) GetErrorDataFilter() *string {
	return s.ErrorDataFilter
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers) GetType() *string {
	return s.Type
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers) SetErrorDataFilter(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers {
	s.ErrorDataFilter = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers) SetType(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers {
	s.Type = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig struct {
	// Sampling metric name.
	//
	// - Count: number of table rows.
	//
	// - Min: minimum value of the field.
	//
	// - Max: maximum value of the field.
	//
	// - Avg: average value of the field.
	//
	// - DistinctCount: number of distinct values of the field.
	//
	// - DistinctPercent: ratio of the number of distinct values of the field to the number of data rows.
	//
	// - DuplicatedCount: number of duplicate values of the field.
	//
	// - DuplicatedPercent: ratio of the number of duplicate values of the field to the number of data rows.
	//
	// - TableSize: table size.
	//
	// - NullValueCount: number of rows where the field value is null.
	//
	// - NullValuePercent: percentage of rows where the field value is null.
	//
	// - GroupCount: each value and its corresponding number of data rows after aggregation by field value.
	//
	// - CountNotIn: number of rows whose enumerated values do not match.
	//
	// - CountDistinctNotIn: number of distinct values whose enumerated values do not match.
	//
	// - UserDefinedSql: sample collection via custom SQL.
	//
	// example:
	//
	// Max
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// Parameters required during sample collection.
	//
	// example:
	//
	// { "Columns": [ "id", "name" ] , "SQL": "select count(1) from table;"}
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// Condition for the secondary filtering of data that is not of concern during sampling. Maximum length: 16,777,215 characters.
	//
	// example:
	//
	// id IS NULL
	SamplingFilter *string `json:"SamplingFilter,omitempty" xml:"SamplingFilter,omitempty"`
	// Runtime parameter setting statements inserted and executed before the sampling statement is actually executed. Maximum length: 1,000 characters. Currently only MaxCompute is supported.
	//
	// example:
	//
	// SET odps.sql.udf.timeout=600s;
	//
	// SET odps.sql.python.version=cp27;
	SettingConfig *string `json:"SettingConfig,omitempty" xml:"SettingConfig,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) GetMetric() *string {
	return s.Metric
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) GetMetricParameters() *string {
	return s.MetricParameters
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) GetSamplingFilter() *string {
	return s.SamplingFilter
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) GetSettingConfig() *string {
	return s.SettingConfig
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) SetMetric(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig {
	s.Metric = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) SetMetricParameters(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) SetSamplingFilter(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig {
	s.SamplingFilter = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) SetSettingConfig(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig {
	s.SettingConfig = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget struct {
	// For table-type datasets, the database type to which the table belongs.
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
	// Unique ID of the table to which the rule applies in Data Map.
	//
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// Type of the monitored object.
	//
	// - Table
	//
	// example:
	//
	// Table
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) GoString() string {
	return s.String()
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) GetTableGuid() *string {
	return s.TableGuid
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) GetType() *string {
	return s.Type
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) SetDatabaseType(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget {
	s.DatabaseType = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) SetTableGuid(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget {
	s.TableGuid = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) SetType(v string) *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget {
	s.Type = &v
	return s
}

func (s *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget) Validate() error {
	return dara.Validate(s)
}
