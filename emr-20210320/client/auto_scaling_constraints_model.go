// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutoScalingConstraints interface {
	dara.Model
	String() string
	GoString() string
	SetAutoScalingMetricUnits(v []*MetricUnitValue) *AutoScalingConstraints
	GetAutoScalingMetricUnits() []*MetricUnitValue
	SetDefaultMetricTriggeredRules(v []*ScalingRule) *AutoScalingConstraints
	GetDefaultMetricTriggeredRules() []*ScalingRule
	SetMaxAdjustmentValue(v int32) *AutoScalingConstraints
	GetMaxAdjustmentValue() *int32
	SetMaxByLoadRuleCount(v int32) *AutoScalingConstraints
	GetMaxByLoadRuleCount() *int32
	SetMaxByTimeRuleCount(v int32) *AutoScalingConstraints
	GetMaxByTimeRuleCount() *int32
	SetSupportMetricTags(v []*AutoScalingConstraintsSupportMetricTags) *AutoScalingConstraints
	GetSupportMetricTags() []*AutoScalingConstraintsSupportMetricTags
	SetSupportMetrics(v []*string) *AutoScalingConstraints
	GetSupportMetrics() []*string
	SetSupportRuleTypes(v []*string) *AutoScalingConstraints
	GetSupportRuleTypes() []*string
}

type AutoScalingConstraints struct {
	// 按负载伸缩指标单位描述。
	AutoScalingMetricUnits []*MetricUnitValue `json:"AutoScalingMetricUnits,omitempty" xml:"AutoScalingMetricUnits,omitempty" type:"Repeated"`
	// 默认按负载弹性伸缩规则列表
	DefaultMetricTriggeredRules []*ScalingRule `json:"DefaultMetricTriggeredRules,omitempty" xml:"DefaultMetricTriggeredRules,omitempty" type:"Repeated"`
	// 单次伸缩活动最大扩缩容节点数量。
	//
	// example:
	//
	// 1000
	MaxAdjustmentValue *int32 `json:"MaxAdjustmentValue,omitempty" xml:"MaxAdjustmentValue,omitempty"`
	// 按负载规则数量最大值。
	//
	// example:
	//
	// 10
	MaxByLoadRuleCount *int32 `json:"MaxByLoadRuleCount,omitempty" xml:"MaxByLoadRuleCount,omitempty"`
	// 按时间规则数量最大值。
	//
	// example:
	//
	// 10
	MaxByTimeRuleCount *int32 `json:"MaxByTimeRuleCount,omitempty" xml:"MaxByTimeRuleCount,omitempty"`
	// 支持的按负载弹性伸缩指标Tag列表。
	SupportMetricTags []*AutoScalingConstraintsSupportMetricTags `json:"SupportMetricTags,omitempty" xml:"SupportMetricTags,omitempty" type:"Repeated"`
	// 支持的按负载弹性伸缩指标列表。
	//
	// example:
	//
	// ["YarnRootAvailableMemoryUsage","YarnRootAvailableVcores"]
	SupportMetrics []*string `json:"SupportMetrics,omitempty" xml:"SupportMetrics,omitempty" type:"Repeated"`
	// 支持的弹性伸缩规则类型。
	//
	// example:
	//
	// ["TIME_TRIGGER","METRICS_TRIGGER"]
	SupportRuleTypes []*string `json:"SupportRuleTypes,omitempty" xml:"SupportRuleTypes,omitempty" type:"Repeated"`
}

func (s AutoScalingConstraints) String() string {
	return dara.Prettify(s)
}

func (s AutoScalingConstraints) GoString() string {
	return s.String()
}

func (s *AutoScalingConstraints) GetAutoScalingMetricUnits() []*MetricUnitValue {
	return s.AutoScalingMetricUnits
}

func (s *AutoScalingConstraints) GetDefaultMetricTriggeredRules() []*ScalingRule {
	return s.DefaultMetricTriggeredRules
}

func (s *AutoScalingConstraints) GetMaxAdjustmentValue() *int32 {
	return s.MaxAdjustmentValue
}

func (s *AutoScalingConstraints) GetMaxByLoadRuleCount() *int32 {
	return s.MaxByLoadRuleCount
}

func (s *AutoScalingConstraints) GetMaxByTimeRuleCount() *int32 {
	return s.MaxByTimeRuleCount
}

func (s *AutoScalingConstraints) GetSupportMetricTags() []*AutoScalingConstraintsSupportMetricTags {
	return s.SupportMetricTags
}

func (s *AutoScalingConstraints) GetSupportMetrics() []*string {
	return s.SupportMetrics
}

func (s *AutoScalingConstraints) GetSupportRuleTypes() []*string {
	return s.SupportRuleTypes
}

func (s *AutoScalingConstraints) SetAutoScalingMetricUnits(v []*MetricUnitValue) *AutoScalingConstraints {
	s.AutoScalingMetricUnits = v
	return s
}

func (s *AutoScalingConstraints) SetDefaultMetricTriggeredRules(v []*ScalingRule) *AutoScalingConstraints {
	s.DefaultMetricTriggeredRules = v
	return s
}

func (s *AutoScalingConstraints) SetMaxAdjustmentValue(v int32) *AutoScalingConstraints {
	s.MaxAdjustmentValue = &v
	return s
}

func (s *AutoScalingConstraints) SetMaxByLoadRuleCount(v int32) *AutoScalingConstraints {
	s.MaxByLoadRuleCount = &v
	return s
}

func (s *AutoScalingConstraints) SetMaxByTimeRuleCount(v int32) *AutoScalingConstraints {
	s.MaxByTimeRuleCount = &v
	return s
}

func (s *AutoScalingConstraints) SetSupportMetricTags(v []*AutoScalingConstraintsSupportMetricTags) *AutoScalingConstraints {
	s.SupportMetricTags = v
	return s
}

func (s *AutoScalingConstraints) SetSupportMetrics(v []*string) *AutoScalingConstraints {
	s.SupportMetrics = v
	return s
}

func (s *AutoScalingConstraints) SetSupportRuleTypes(v []*string) *AutoScalingConstraints {
	s.SupportRuleTypes = v
	return s
}

func (s *AutoScalingConstraints) Validate() error {
	return dara.Validate(s)
}

type AutoScalingConstraintsSupportMetricTags struct {
	// 指标名称。
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// 指标Tag。
	Tags []*Tag `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s AutoScalingConstraintsSupportMetricTags) String() string {
	return dara.Prettify(s)
}

func (s AutoScalingConstraintsSupportMetricTags) GoString() string {
	return s.String()
}

func (s *AutoScalingConstraintsSupportMetricTags) GetMetricName() *string {
	return s.MetricName
}

func (s *AutoScalingConstraintsSupportMetricTags) GetTags() []*Tag {
	return s.Tags
}

func (s *AutoScalingConstraintsSupportMetricTags) SetMetricName(v string) *AutoScalingConstraintsSupportMetricTags {
	s.MetricName = &v
	return s
}

func (s *AutoScalingConstraintsSupportMetricTags) SetTags(v []*Tag) *AutoScalingConstraintsSupportMetricTags {
	s.Tags = v
	return s
}

func (s *AutoScalingConstraintsSupportMetricTags) Validate() error {
	return dara.Validate(s)
}
