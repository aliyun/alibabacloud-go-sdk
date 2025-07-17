// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataQualityResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDataQualityResultsResponseBodyPagingInfo) *ListDataQualityResultsResponseBody
	GetPagingInfo() *ListDataQualityResultsResponseBodyPagingInfo
	SetRequestId(v string) *ListDataQualityResultsResponseBody
	GetRequestId() *string
}

type ListDataQualityResultsResponseBody struct {
	// The pagination information.
	PagingInfo *ListDataQualityResultsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataQualityResultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBody) GetPagingInfo() *ListDataQualityResultsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDataQualityResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataQualityResultsResponseBody) SetPagingInfo(v *ListDataQualityResultsResponseBodyPagingInfo) *ListDataQualityResultsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDataQualityResultsResponseBody) SetRequestId(v string) *ListDataQualityResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataQualityResultsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityResultsResponseBodyPagingInfo struct {
	// The data quality check results.
	DataQualityResults []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResults `json:"DataQualityResults,omitempty" xml:"DataQualityResults,omitempty" type:"Repeated"`
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
	// 219
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfo) GetDataQualityResults() []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResults {
	return s.DataQualityResults
}

func (s *ListDataQualityResultsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDataQualityResultsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataQualityResultsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataQualityResultsResponseBodyPagingInfo) SetDataQualityResults(v []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) *ListDataQualityResultsResponseBodyPagingInfo {
	s.DataQualityResults = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfo) SetPageNumber(v int32) *ListDataQualityResultsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfo) SetPageSize(v int32) *ListDataQualityResultsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfo) SetTotalCount(v int32) *ListDataQualityResultsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResults struct {
	// The time when the data quality check result was generated.
	//
	// example:
	//
	// 1708284916414
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The information about the data quality check.
	Details []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The ID of the check result.
	//
	// example:
	//
	// 16033
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The snapshot of the rule configuration when the check starts.
	Rule *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
	// The sample values used for the check.
	//
	// example:
	//
	// [
	//
	//   {
	//
	//     "gender": "male",
	//
	//     "_count": 100
	//
	//   }, {
	//
	//     "gender": "female",
	//
	//     "_count": 100
	//
	//   }
	//
	// ]
	Sample *string `json:"Sample,omitempty" xml:"Sample,omitempty"`
	// The status of the check result. Valid values:
	//
	// 	- Running
	//
	// 	- Error
	//
	// 	- Passed
	//
	// 	- Warned
	//
	// 	- Critical
	//
	// example:
	//
	// PASSED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the instance generated by the check.
	//
	// example:
	//
	// 200001
	TaskInstanceId *int64 `json:"TaskInstanceId,omitempty" xml:"TaskInstanceId,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) GetDetails() []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails {
	return s.Details
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) GetId() *int64 {
	return s.Id
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) GetRule() *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	return s.Rule
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) GetSample() *string {
	return s.Sample
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) GetStatus() *string {
	return s.Status
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) GetTaskInstanceId() *int64 {
	return s.TaskInstanceId
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) SetCreateTime(v int64) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults {
	s.CreateTime = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) SetDetails(v []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults {
	s.Details = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) SetId(v int64) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults {
	s.Id = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) SetRule(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults {
	s.Rule = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) SetSample(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults {
	s.Sample = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) SetStatus(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults {
	s.Status = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) SetTaskInstanceId(v int64) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults {
	s.TaskInstanceId = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResults) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails struct {
	// The value that is used for comparison with the threshold.
	//
	// example:
	//
	// 100.0
	CheckedValue *string `json:"CheckedValue,omitempty" xml:"CheckedValue,omitempty"`
	// The value that is calculated based on sample data. The value serves as a baseline value during the calculation of the value of the CheckedValue parameter.
	//
	// example:
	//
	// 0.0
	ReferencedValue *string `json:"ReferencedValue,omitempty" xml:"ReferencedValue,omitempty"`
	// The comparison result between the value of CheckedValue and the threshold. Valid values:
	//
	// 	- Error
	//
	// 	- Passed
	//
	// 	- Warned
	//
	// 	- Critical
	//
	// example:
	//
	// PASSED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) GetCheckedValue() *string {
	return s.CheckedValue
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) GetReferencedValue() *string {
	return s.ReferencedValue
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) GetStatus() *string {
	return s.Status
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) SetCheckedValue(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails {
	s.CheckedValue = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) SetReferencedValue(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails {
	s.ReferencedValue = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) SetStatus(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails {
	s.Status = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule struct {
	// The check settings for sample data.
	CheckingConfig *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
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
	ErrorHandlers []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// The rule ID.
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
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The sampling settings.
	SamplingConfig *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// The strength of the rule. Valid values:
	//
	// 	- High
	//
	// 	- Normal
	//
	// example:
	//
	// NORMAL
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The monitored object of the rule.
	Target *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// The code of the template that is referenced when you create a rule.
	//
	// example:
	//
	// system::user_defined
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) GetCheckingConfig() *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig {
	return s.CheckingConfig
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) GetDescription() *string {
	return s.Description
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) GetErrorHandlers() []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers {
	return s.ErrorHandlers
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) GetId() *int64 {
	return s.Id
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) GetName() *string {
	return s.Name
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) GetSamplingConfig() *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig {
	return s.SamplingConfig
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) GetSeverity() *string {
	return s.Severity
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) GetTarget() *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget {
	return s.Target
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetCheckingConfig(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.CheckingConfig = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetDescription(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.Description = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetEnabled(v bool) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.Enabled = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetErrorHandlers(v []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.ErrorHandlers = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetId(v int64) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.Id = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetName(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.Name = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetProjectId(v int64) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.ProjectId = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetSamplingConfig(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.SamplingConfig = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetSeverity(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.Severity = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetTarget(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.Target = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) SetTemplateCode(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule {
	s.TemplateCode = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig struct {
	// The method that is used to query the referenced samples. To obtain some types of thresholds, you need to query reference samples and perform aggregate operations on the reference values. In this example, an expression is used to indicate the query method of referenced samples.
	//
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// The threshold settings.
	Thresholds *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
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
	// FIXED
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) GetReferencedSamplesFilter() *string {
	return s.ReferencedSamplesFilter
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) GetThresholds() *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds {
	return s.Thresholds
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) GetType() *string {
	return s.Type
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) SetReferencedSamplesFilter(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig {
	s.ReferencedSamplesFilter = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) SetThresholds(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig {
	s.Thresholds = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) SetType(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig {
	s.Type = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds struct {
	// The threshold settings for critical alerts.
	Critical *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	// The expected threshold setting.
	Expected *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	// The threshold settings for normal alerts.
	Warned *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned `json:"Warned,omitempty" xml:"Warned,omitempty" type:"Struct"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) GetCritical() *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical {
	return s.Critical
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) GetExpected() *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected {
	return s.Expected
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) GetWarned() *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned {
	return s.Warned
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) SetCritical(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds {
	s.Critical = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) SetExpected(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds {
	s.Expected = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) SetWarned(v *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds {
	s.Warned = v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical struct {
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

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical) GetExpression() *string {
	return s.Expression
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical) GetOperator() *string {
	return s.Operator
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical) GetValue() *string {
	return s.Value
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical) SetExpression(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical {
	s.Expression = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical) SetOperator(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical {
	s.Operator = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical) SetValue(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical {
	s.Value = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected struct {
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

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected) GetExpression() *string {
	return s.Expression
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected) GetOperator() *string {
	return s.Operator
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected) GetValue() *string {
	return s.Value
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected) SetExpression(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected {
	s.Expression = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected) SetOperator(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected {
	s.Operator = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected) SetValue(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected {
	s.Value = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned struct {
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

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned) GetExpression() *string {
	return s.Expression
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned) GetOperator() *string {
	return s.Operator
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned) GetValue() *string {
	return s.Value
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned) SetExpression(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned {
	s.Expression = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned) SetOperator(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned {
	s.Operator = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned) SetValue(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned {
	s.Value = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsWarned) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers struct {
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
	// SAVE_ERROR_DATA
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers) GetErrorDataFilter() *string {
	return s.ErrorDataFilter
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers) GetType() *string {
	return s.Type
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers) SetErrorDataFilter(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers {
	s.ErrorDataFilter = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers) SetType(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers {
	s.Type = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig struct {
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
	// COUNT
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// The parameters required for sampling.
	//
	// example:
	//
	// { "columns": [ "id", "name" ] }
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

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) GetMetric() *string {
	return s.Metric
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) GetMetricParameters() *string {
	return s.MetricParameters
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) GetSamplingFilter() *string {
	return s.SamplingFilter
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) GetSettingConfig() *string {
	return s.SettingConfig
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) SetMetric(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig {
	s.Metric = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) SetMetricParameters(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig {
	s.MetricParameters = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) SetSamplingFilter(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig {
	s.SamplingFilter = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) SetSettingConfig(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig {
	s.SettingConfig = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig) Validate() error {
	return dara.Validate(s)
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget struct {
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
	// MAX_COMPUTE
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
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
	// TABLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) String() string {
	return dara.Prettify(s)
}

func (s ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) GoString() string {
	return s.String()
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) GetTableGuid() *string {
	return s.TableGuid
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) GetType() *string {
	return s.Type
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) SetDatabaseType(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget {
	s.DatabaseType = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) SetTableGuid(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget {
	s.TableGuid = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) SetType(v string) *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget {
	s.Type = &v
	return s
}

func (s *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget) Validate() error {
	return dara.Validate(s)
}
