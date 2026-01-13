// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListResourceRulesResponseBody
	GetRequestId() *string
	SetResourceRules(v []*ListResourceRulesResponseBodyResourceRules) *ListResourceRulesResponseBody
	GetResourceRules() []*ListResourceRulesResponseBodyResourceRules
	SetTotalCount(v int64) *ListResourceRulesResponseBody
	GetTotalCount() *int64
}

type ListResourceRulesResponseBody struct {
	RequestId     *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceRules []*ListResourceRulesResponseBodyResourceRules `json:"ResourceRules,omitempty" xml:"ResourceRules,omitempty" type:"Repeated"`
	TotalCount    *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceRulesResponseBody) GetResourceRules() []*ListResourceRulesResponseBodyResourceRules {
	return s.ResourceRules
}

func (s *ListResourceRulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListResourceRulesResponseBody) SetRequestId(v string) *ListResourceRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceRulesResponseBody) SetResourceRules(v []*ListResourceRulesResponseBodyResourceRules) *ListResourceRulesResponseBody {
	s.ResourceRules = v
	return s
}

func (s *ListResourceRulesResponseBody) SetTotalCount(v int64) *ListResourceRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceRulesResponseBody) Validate() error {
	if s.ResourceRules != nil {
		for _, item := range s.ResourceRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceRulesResponseBodyResourceRules struct {
	Description             *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	MetricOperationType     *string                                                `json:"MetricOperationType,omitempty" xml:"MetricOperationType,omitempty"`
	MetricPullInfo          *string                                                `json:"MetricPullInfo,omitempty" xml:"MetricPullInfo,omitempty"`
	MetricPullPeriod        *string                                                `json:"MetricPullPeriod,omitempty" xml:"MetricPullPeriod,omitempty"`
	Name                    *string                                                `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceRuleId          *string                                                `json:"ResourceRuleId,omitempty" xml:"ResourceRuleId,omitempty"`
	RuleComputingDefinition *string                                                `json:"RuleComputingDefinition,omitempty" xml:"RuleComputingDefinition,omitempty"`
	RuleItems               []*ListResourceRulesResponseBodyResourceRulesRuleItems `json:"RuleItems,omitempty" xml:"RuleItems,omitempty" type:"Repeated"`
}

func (s ListResourceRulesResponseBodyResourceRules) String() string {
	return dara.Prettify(s)
}

func (s ListResourceRulesResponseBodyResourceRules) GoString() string {
	return s.String()
}

func (s *ListResourceRulesResponseBodyResourceRules) GetDescription() *string {
	return s.Description
}

func (s *ListResourceRulesResponseBodyResourceRules) GetMetricOperationType() *string {
	return s.MetricOperationType
}

func (s *ListResourceRulesResponseBodyResourceRules) GetMetricPullInfo() *string {
	return s.MetricPullInfo
}

func (s *ListResourceRulesResponseBodyResourceRules) GetMetricPullPeriod() *string {
	return s.MetricPullPeriod
}

func (s *ListResourceRulesResponseBodyResourceRules) GetName() *string {
	return s.Name
}

func (s *ListResourceRulesResponseBodyResourceRules) GetResourceRuleId() *string {
	return s.ResourceRuleId
}

func (s *ListResourceRulesResponseBodyResourceRules) GetRuleComputingDefinition() *string {
	return s.RuleComputingDefinition
}

func (s *ListResourceRulesResponseBodyResourceRules) GetRuleItems() []*ListResourceRulesResponseBodyResourceRulesRuleItems {
	return s.RuleItems
}

func (s *ListResourceRulesResponseBodyResourceRules) SetDescription(v string) *ListResourceRulesResponseBodyResourceRules {
	s.Description = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRules) SetMetricOperationType(v string) *ListResourceRulesResponseBodyResourceRules {
	s.MetricOperationType = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRules) SetMetricPullInfo(v string) *ListResourceRulesResponseBodyResourceRules {
	s.MetricPullInfo = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRules) SetMetricPullPeriod(v string) *ListResourceRulesResponseBodyResourceRules {
	s.MetricPullPeriod = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRules) SetName(v string) *ListResourceRulesResponseBodyResourceRules {
	s.Name = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRules) SetResourceRuleId(v string) *ListResourceRulesResponseBodyResourceRules {
	s.ResourceRuleId = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRules) SetRuleComputingDefinition(v string) *ListResourceRulesResponseBodyResourceRules {
	s.RuleComputingDefinition = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRules) SetRuleItems(v []*ListResourceRulesResponseBodyResourceRulesRuleItems) *ListResourceRulesResponseBodyResourceRules {
	s.RuleItems = v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRules) Validate() error {
	if s.RuleItems != nil {
		for _, item := range s.RuleItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceRulesResponseBodyResourceRulesRuleItems struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxValue    *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinValue    *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListResourceRulesResponseBodyResourceRulesRuleItems) String() string {
	return dara.Prettify(s)
}

func (s ListResourceRulesResponseBodyResourceRulesRuleItems) GoString() string {
	return s.String()
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) GetDescription() *string {
	return s.Description
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) GetMaxValue() *string {
	return s.MaxValue
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) GetMinValue() *string {
	return s.MinValue
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) GetName() *string {
	return s.Name
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) GetValue() *string {
	return s.Value
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) SetDescription(v string) *ListResourceRulesResponseBodyResourceRulesRuleItems {
	s.Description = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) SetMaxValue(v string) *ListResourceRulesResponseBodyResourceRulesRuleItems {
	s.MaxValue = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) SetMinValue(v string) *ListResourceRulesResponseBodyResourceRulesRuleItems {
	s.MinValue = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) SetName(v string) *ListResourceRulesResponseBodyResourceRulesRuleItems {
	s.Name = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) SetValue(v string) *ListResourceRulesResponseBodyResourceRulesRuleItems {
	s.Value = &v
	return s
}

func (s *ListResourceRulesResponseBodyResourceRulesRuleItems) Validate() error {
	return dara.Validate(s)
}
