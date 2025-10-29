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
	// The pagination information.
	PagingInfo *ListDataQualityRulesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
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
	// The rules.
	DataQualityRules []*ListDataQualityRulesResponseBodyPagingInfoDataQualityRules `json:"DataQualityRules,omitempty" xml:"DataQualityRules,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
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
	// The check settings for sample data.
	CheckingConfig *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
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
	ErrorHandlers []*ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// The rule ID.
	//
	// example:
	//
	// 22130
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
	// 100001
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The settings for sampling.
	SamplingConfig *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
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
	Target *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The ID of the template used by the rule.
	//
	// example:
	//
	// system::user_defined
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
	// The method that is used to query the referenced samples. To obtain some types of thresholds, you need to query reference values. In this example, an expression is used to indicate the query method of referenced samples.
	//
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// The threshold settings.
	Thresholds *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
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
	// The threshold settings for critical alerts.
	Critical *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	// The expected threshold setting.
	Expected *ListDataQualityRulesResponseBodyPagingInfoDataQualityRulesCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	// The threshold settings for normal alerts.
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
	// The threshold expression.
	//
	// If the template specified by the TemplateCode parameter is about fluctuation, you must use an expression to represent the threshold for fluctuation. Examples:
	//
	// 	- $checkValue > 0.01
	//
	// 	- $checkValue < -0.01
	//
	// 	- abs($checkValue) > 0.01
	//
	// If the template specified by the TemplateCode parameter is about fixed value, you can also use an expression to represent the threshold. If you configure the Expression, Operator, and Value parameters for the threshold at the same time, the Expression parameter takes precedence over the Operator and Value parameters.
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
	// The threshold expression.
	//
	// If the template specified by the TemplateCode parameter is about fluctuation, you must use an expression to represent the threshold for fluctuation. Examples:
	//
	// 	- $checkValue > 0.01
	//
	// 	- $checkValue < -0.01
	//
	// 	- abs($checkValue) > 0.01
	//
	// If the template specified by the TemplateCode parameter is about fixed value, you can also use an expression to represent the threshold. If you configure the Expression, Operator, and Value parameters for the threshold at the same time, the Expression parameter takes precedence over the Operator and Value parameters.
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
	// The threshold expression.
	//
	// If the template specified by the TemplateCode parameter is about fluctuation, you must use an expression to represent the threshold for fluctuation. Examples:
	//
	// 	- $checkValue > 0.01
	//
	// 	- $checkValue < -0.01
	//
	// 	- abs($checkValue) > 0.01
	//
	// If the template specified by the TemplateCode parameter is about fixed value, you can also use an expression to represent the threshold. If you configure the Expression, Operator, and Value parameters for the threshold at the same time, the Expression parameter takes precedence over the Operator and Value parameters.
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
	// The ID of the table that is limited by the rule in Data Map.
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
