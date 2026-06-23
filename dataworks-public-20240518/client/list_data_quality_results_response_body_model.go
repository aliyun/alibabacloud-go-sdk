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
	// 数据质量校验结果分页查询结果。
	PagingInfo *ListDataQualityResultsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// API请求ID。
	//
	// example:
	//
	// 691CA452-D37A-****
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
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataQualityResultsResponseBodyPagingInfo struct {
	// 质量校验结果列表。
	DataQualityResults []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResults `json:"DataQualityResults,omitempty" xml:"DataQualityResults,omitempty" type:"Repeated"`
	// 页码。
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 页大小。
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 总条数。
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
	if s.DataQualityResults != nil {
		for _, item := range s.DataQualityResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResults struct {
	// 校验结果生成时间。
	//
	// example:
	//
	// 1708284916414
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 本次校验的详情。
	Details []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// 校验结果ID。
	//
	// example:
	//
	// 16033
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 校验开始时，规则配置快照。
	Rule *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Struct"`
	// 本次校验所使用的样本值。
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
	// 校验结果状态：
	//
	// - Running
	//
	// - Error
	//
	// - Passed
	//
	// - Warned
	//
	// - Critical
	//
	// example:
	//
	// Passed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 校验任务实例ID。
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

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsDetails struct {
	// 最终用来与阈值比较的值。
	//
	// example:
	//
	// 100.0
	CheckedValue *string `json:"CheckedValue,omitempty" xml:"CheckedValue,omitempty"`
	// 使用引用的样本，用来参与CheckedValue计算的基准值。
	//
	// example:
	//
	// 0.0
	ReferencedValue *string `json:"ReferencedValue,omitempty" xml:"ReferencedValue,omitempty"`
	// 最终的比较结果状态：
	//
	// - Error
	//
	// - Passed
	//
	// - Warned
	//
	// - Critical
	//
	// example:
	//
	// Passed
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
	// 样本校验设置。
	CheckingConfig *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig `json:"CheckingConfig,omitempty" xml:"CheckingConfig,omitempty" type:"Struct"`
	// 规则描述信息，最长500个字符。
	//
	// example:
	//
	// this is a odps _sql task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 规则是否启用。
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// 质量规则校验问题处理器。
	ErrorHandlers []*ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleErrorHandlers `json:"ErrorHandlers,omitempty" xml:"ErrorHandlers,omitempty" type:"Repeated"`
	// 规则ID。
	//
	// example:
	//
	// 100001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 规则名称，数字、英文字母、汉字、半角全角标点符号组合，最长255个字符。
	//
	// example:
	//
	// The table cannot be empty.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// DataWorks项目空间ID。
	//
	// example:
	//
	// 100
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 样本采集所需的设置。
	SamplingConfig *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleSamplingConfig `json:"SamplingConfig,omitempty" xml:"SamplingConfig,omitempty" type:"Struct"`
	// 规则对于业务的等级（对应页面上的强弱规则）：
	//
	// - High
	//
	// - Normal
	//
	// example:
	//
	// Normal
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// 规则所监控的对象。
	Target *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleTarget `json:"Target,omitempty" xml:"Target,omitempty" type:"Struct"`
	// 创建规则时所引用的规则模板Code。
	//
	// example:
	//
	// SYSTEM:user_defined_sql
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

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfig struct {
	// 有些类型的阈值需要查询出一些参考样本，然后对参考样本的值进行汇总得出进行比较的阈值，这里使用一个表达式来表示参考样本的查询方式。
	//
	// example:
	//
	// { "bizdate": [ "-1", "-7", "-1m" ] }
	ReferencedSamplesFilter *string `json:"ReferencedSamplesFilter,omitempty" xml:"ReferencedSamplesFilter,omitempty"`
	// 阈值设置。
	Thresholds *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds `json:"Thresholds,omitempty" xml:"Thresholds,omitempty" type:"Struct"`
	// 阈值计算方式：
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
	if s.Thresholds != nil {
		if err := s.Thresholds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholds struct {
	// 严重警告的阈值设置。
	Critical *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical `json:"Critical,omitempty" xml:"Critical,omitempty" type:"Struct"`
	// 期望的阈值设置。
	Expected *ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsExpected `json:"Expected,omitempty" xml:"Expected,omitempty" type:"Struct"`
	// 普通警告的阈值设置。
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

type ListDataQualityResultsResponseBodyPagingInfoDataQualityResultsRuleCheckingConfigThresholdsCritical struct {
	// 阈值表达式。
	//
	// 波动率类型规则必须使用表达式方式表示波动阈值。如：
	//
	// - 波动上升大于0.01： $checkValue > 0.01
	//
	// - 波动下降大于0.01：$checkValue < -0.01
	//
	// - 波动率绝对值：abs($checkValue) > 0.01
	//
	// 固定值类型规则也可以使用表达式方式配置阈值，如果同时配置，表达式优先级高于Operator和Value。
	//
	// example:
	//
	// $checkValue > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// 比较符：
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
	// 阈值数值。
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
	// 阈值表达式。
	//
	// 波动率类型规则必须使用表达式方式表示波动阈值。如：
	//
	// - 波动上升大于0.01： $checkValue > 0.01
	//
	// - 波动下降大于0.01：$checkValue < -0.01
	//
	// - 波动率绝对值：abs($checkValue) > 0.01
	//
	// 固定值类型规则也可以使用表达式方式配置阈值，如果同时配置，表达式优先级高于Operator和Value。
	//
	// example:
	//
	// $checkValue > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// 比较符：
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
	// 阈值数值。
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
	// 阈值表达式。
	//
	// 波动率类型规则必须使用表达式方式表示波动阈值。如：
	//
	// - 波动上升大于0.01： $checkValue > 0.01
	//
	// - 波动下降大于0.01：$checkValue < -0.01
	//
	// - 波动率绝对值：abs($checkValue) > 0.01
	//
	// 固定值类型规则也可以使用表达式方式配置阈值，如果同时配置，表达式优先级高于Operator和Value。
	//
	// example:
	//
	// $checkValue > 0.01
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// 比较符：
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
	// 阈值数值。
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
	// 如果是自定义SQL规则，需要用户指定SQL来过滤问题数据。
	//
	// example:
	//
	// SELECT 	- FROM tb_api_log WHERE id IS NULL
	ErrorDataFilter *string `json:"ErrorDataFilter,omitempty" xml:"ErrorDataFilter,omitempty"`
	// 处理器类型。
	//
	// - SaveErrorData
	//
	// example:
	//
	// SaveErrorData
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
	// 采样的指标名称：
	//
	// - Count：表行数
	//
	// - Min：字段最小值
	//
	// - Max：字段最大值
	//
	// - Avg：字段均值
	//
	// - DistinctCount：字段唯一值个数
	//
	// - DistinctPercent：字段唯一值个数与数据行数占比
	//
	// - DuplicatedCount：字段重复值个数
	//
	// - DuplicatedPercent：字段重复值个数与数据行数占比
	//
	// - TableSize：表大小
	//
	// - NullValueCount：字段为空的行数
	//
	// - NullValuePercent：字段为空的比例
	//
	// - GroupCount：按字段值聚合后每个值与对应的数据行数
	//
	// - CountNotIn：枚举值不匹配行数
	//
	// - CountDistinctNotIn：枚举值不匹配唯一值个数
	//
	// - UserDefinedSql：通过自定义SQL做样本采集
	//
	// example:
	//
	// Count
	Metric *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	// 样本采集时，所需的参数。
	//
	// example:
	//
	// { "columns": [ "id", "name" ] }
	MetricParameters *string `json:"MetricParameters,omitempty" xml:"MetricParameters,omitempty"`
	// 采样时，对不关注的数据进行二次过滤的条件，最多16777215个字符。
	//
	// example:
	//
	// id IS NULL
	SamplingFilter *string `json:"SamplingFilter,omitempty" xml:"SamplingFilter,omitempty"`
	// 具体执行采样语句前，插入执行的一些运行时参数设置语句，最长1000个字符。目前只支持MaxCompute。
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
	// 表类型的数据集，表所属的数据库类型：
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
	// 表在数据地图中的唯一ID。
	//
	// example:
	//
	// odps.unit_test.tb_unit_test
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// 监控对象类型：
	//
	// - Table
	//
	// example:
	//
	// Table
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
