// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushResourceRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *PushResourceRuleResponseBody
	GetDescription() *string
	SetMetricOperationType(v string) *PushResourceRuleResponseBody
	GetMetricOperationType() *string
	SetMetricPullInfo(v string) *PushResourceRuleResponseBody
	GetMetricPullInfo() *string
	SetMetricPullPeriod(v string) *PushResourceRuleResponseBody
	GetMetricPullPeriod() *string
	SetName(v string) *PushResourceRuleResponseBody
	GetName() *string
	SetRequestId(v string) *PushResourceRuleResponseBody
	GetRequestId() *string
	SetResourceRuleId(v string) *PushResourceRuleResponseBody
	GetResourceRuleId() *string
	SetRuleComputingDefinition(v string) *PushResourceRuleResponseBody
	GetRuleComputingDefinition() *string
	SetRuleItems(v []*PushResourceRuleResponseBodyRuleItems) *PushResourceRuleResponseBody
	GetRuleItems() []*PushResourceRuleResponseBodyRuleItems
}

type PushResourceRuleResponseBody struct {
	Description             *string                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	MetricOperationType     *string                                  `json:"MetricOperationType,omitempty" xml:"MetricOperationType,omitempty"`
	MetricPullInfo          *string                                  `json:"MetricPullInfo,omitempty" xml:"MetricPullInfo,omitempty"`
	MetricPullPeriod        *string                                  `json:"MetricPullPeriod,omitempty" xml:"MetricPullPeriod,omitempty"`
	Name                    *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId               *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceRuleId          *string                                  `json:"ResourceRuleId,omitempty" xml:"ResourceRuleId,omitempty"`
	RuleComputingDefinition *string                                  `json:"RuleComputingDefinition,omitempty" xml:"RuleComputingDefinition,omitempty"`
	RuleItems               []*PushResourceRuleResponseBodyRuleItems `json:"RuleItems,omitempty" xml:"RuleItems,omitempty" type:"Repeated"`
}

func (s PushResourceRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushResourceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PushResourceRuleResponseBody) GetDescription() *string {
	return s.Description
}

func (s *PushResourceRuleResponseBody) GetMetricOperationType() *string {
	return s.MetricOperationType
}

func (s *PushResourceRuleResponseBody) GetMetricPullInfo() *string {
	return s.MetricPullInfo
}

func (s *PushResourceRuleResponseBody) GetMetricPullPeriod() *string {
	return s.MetricPullPeriod
}

func (s *PushResourceRuleResponseBody) GetName() *string {
	return s.Name
}

func (s *PushResourceRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushResourceRuleResponseBody) GetResourceRuleId() *string {
	return s.ResourceRuleId
}

func (s *PushResourceRuleResponseBody) GetRuleComputingDefinition() *string {
	return s.RuleComputingDefinition
}

func (s *PushResourceRuleResponseBody) GetRuleItems() []*PushResourceRuleResponseBodyRuleItems {
	return s.RuleItems
}

func (s *PushResourceRuleResponseBody) SetDescription(v string) *PushResourceRuleResponseBody {
	s.Description = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetMetricOperationType(v string) *PushResourceRuleResponseBody {
	s.MetricOperationType = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetMetricPullInfo(v string) *PushResourceRuleResponseBody {
	s.MetricPullInfo = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetMetricPullPeriod(v string) *PushResourceRuleResponseBody {
	s.MetricPullPeriod = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetName(v string) *PushResourceRuleResponseBody {
	s.Name = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetRequestId(v string) *PushResourceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetResourceRuleId(v string) *PushResourceRuleResponseBody {
	s.ResourceRuleId = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetRuleComputingDefinition(v string) *PushResourceRuleResponseBody {
	s.RuleComputingDefinition = &v
	return s
}

func (s *PushResourceRuleResponseBody) SetRuleItems(v []*PushResourceRuleResponseBodyRuleItems) *PushResourceRuleResponseBody {
	s.RuleItems = v
	return s
}

func (s *PushResourceRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type PushResourceRuleResponseBodyRuleItems struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxValue    *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinValue    *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PushResourceRuleResponseBodyRuleItems) String() string {
	return dara.Prettify(s)
}

func (s PushResourceRuleResponseBodyRuleItems) GoString() string {
	return s.String()
}

func (s *PushResourceRuleResponseBodyRuleItems) GetDescription() *string {
	return s.Description
}

func (s *PushResourceRuleResponseBodyRuleItems) GetMaxValue() *string {
	return s.MaxValue
}

func (s *PushResourceRuleResponseBodyRuleItems) GetMinValue() *string {
	return s.MinValue
}

func (s *PushResourceRuleResponseBodyRuleItems) GetName() *string {
	return s.Name
}

func (s *PushResourceRuleResponseBodyRuleItems) GetValue() *string {
	return s.Value
}

func (s *PushResourceRuleResponseBodyRuleItems) SetDescription(v string) *PushResourceRuleResponseBodyRuleItems {
	s.Description = &v
	return s
}

func (s *PushResourceRuleResponseBodyRuleItems) SetMaxValue(v string) *PushResourceRuleResponseBodyRuleItems {
	s.MaxValue = &v
	return s
}

func (s *PushResourceRuleResponseBodyRuleItems) SetMinValue(v string) *PushResourceRuleResponseBodyRuleItems {
	s.MinValue = &v
	return s
}

func (s *PushResourceRuleResponseBodyRuleItems) SetName(v string) *PushResourceRuleResponseBodyRuleItems {
	s.Name = &v
	return s
}

func (s *PushResourceRuleResponseBodyRuleItems) SetValue(v string) *PushResourceRuleResponseBodyRuleItems {
	s.Value = &v
	return s
}

func (s *PushResourceRuleResponseBodyRuleItems) Validate() error {
	return dara.Validate(s)
}
