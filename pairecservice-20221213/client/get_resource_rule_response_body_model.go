// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *GetResourceRuleResponseBody
	GetDescription() *string
	SetMetricOperationType(v string) *GetResourceRuleResponseBody
	GetMetricOperationType() *string
	SetMetricPullInfo(v string) *GetResourceRuleResponseBody
	GetMetricPullInfo() *string
	SetMetricPullPeriod(v string) *GetResourceRuleResponseBody
	GetMetricPullPeriod() *string
	SetName(v string) *GetResourceRuleResponseBody
	GetName() *string
	SetRequestId(v string) *GetResourceRuleResponseBody
	GetRequestId() *string
	SetResourceRuleId(v string) *GetResourceRuleResponseBody
	GetResourceRuleId() *string
	SetRuleComputingDefinition(v string) *GetResourceRuleResponseBody
	GetRuleComputingDefinition() *string
	SetRuleItems(v []*GetResourceRuleResponseBodyRuleItems) *GetResourceRuleResponseBody
	GetRuleItems() []*GetResourceRuleResponseBodyRuleItems
}

type GetResourceRuleResponseBody struct {
	Description             *string                                 `json:"Description,omitempty" xml:"Description,omitempty"`
	MetricOperationType     *string                                 `json:"MetricOperationType,omitempty" xml:"MetricOperationType,omitempty"`
	MetricPullInfo          *string                                 `json:"MetricPullInfo,omitempty" xml:"MetricPullInfo,omitempty"`
	MetricPullPeriod        *string                                 `json:"MetricPullPeriod,omitempty" xml:"MetricPullPeriod,omitempty"`
	Name                    *string                                 `json:"Name,omitempty" xml:"Name,omitempty"`
	RequestId               *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceRuleId          *string                                 `json:"ResourceRuleId,omitempty" xml:"ResourceRuleId,omitempty"`
	RuleComputingDefinition *string                                 `json:"RuleComputingDefinition,omitempty" xml:"RuleComputingDefinition,omitempty"`
	RuleItems               []*GetResourceRuleResponseBodyRuleItems `json:"RuleItems,omitempty" xml:"RuleItems,omitempty" type:"Repeated"`
}

func (s GetResourceRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceRuleResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetResourceRuleResponseBody) GetMetricOperationType() *string {
	return s.MetricOperationType
}

func (s *GetResourceRuleResponseBody) GetMetricPullInfo() *string {
	return s.MetricPullInfo
}

func (s *GetResourceRuleResponseBody) GetMetricPullPeriod() *string {
	return s.MetricPullPeriod
}

func (s *GetResourceRuleResponseBody) GetName() *string {
	return s.Name
}

func (s *GetResourceRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceRuleResponseBody) GetResourceRuleId() *string {
	return s.ResourceRuleId
}

func (s *GetResourceRuleResponseBody) GetRuleComputingDefinition() *string {
	return s.RuleComputingDefinition
}

func (s *GetResourceRuleResponseBody) GetRuleItems() []*GetResourceRuleResponseBodyRuleItems {
	return s.RuleItems
}

func (s *GetResourceRuleResponseBody) SetDescription(v string) *GetResourceRuleResponseBody {
	s.Description = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetMetricOperationType(v string) *GetResourceRuleResponseBody {
	s.MetricOperationType = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetMetricPullInfo(v string) *GetResourceRuleResponseBody {
	s.MetricPullInfo = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetMetricPullPeriod(v string) *GetResourceRuleResponseBody {
	s.MetricPullPeriod = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetName(v string) *GetResourceRuleResponseBody {
	s.Name = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetRequestId(v string) *GetResourceRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetResourceRuleId(v string) *GetResourceRuleResponseBody {
	s.ResourceRuleId = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetRuleComputingDefinition(v string) *GetResourceRuleResponseBody {
	s.RuleComputingDefinition = &v
	return s
}

func (s *GetResourceRuleResponseBody) SetRuleItems(v []*GetResourceRuleResponseBodyRuleItems) *GetResourceRuleResponseBody {
	s.RuleItems = v
	return s
}

func (s *GetResourceRuleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetResourceRuleResponseBodyRuleItems struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxValue    *string `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinValue    *string `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Value       *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetResourceRuleResponseBodyRuleItems) String() string {
	return dara.Prettify(s)
}

func (s GetResourceRuleResponseBodyRuleItems) GoString() string {
	return s.String()
}

func (s *GetResourceRuleResponseBodyRuleItems) GetDescription() *string {
	return s.Description
}

func (s *GetResourceRuleResponseBodyRuleItems) GetMaxValue() *string {
	return s.MaxValue
}

func (s *GetResourceRuleResponseBodyRuleItems) GetMinValue() *string {
	return s.MinValue
}

func (s *GetResourceRuleResponseBodyRuleItems) GetName() *string {
	return s.Name
}

func (s *GetResourceRuleResponseBodyRuleItems) GetValue() *string {
	return s.Value
}

func (s *GetResourceRuleResponseBodyRuleItems) SetDescription(v string) *GetResourceRuleResponseBodyRuleItems {
	s.Description = &v
	return s
}

func (s *GetResourceRuleResponseBodyRuleItems) SetMaxValue(v string) *GetResourceRuleResponseBodyRuleItems {
	s.MaxValue = &v
	return s
}

func (s *GetResourceRuleResponseBodyRuleItems) SetMinValue(v string) *GetResourceRuleResponseBodyRuleItems {
	s.MinValue = &v
	return s
}

func (s *GetResourceRuleResponseBodyRuleItems) SetName(v string) *GetResourceRuleResponseBodyRuleItems {
	s.Name = &v
	return s
}

func (s *GetResourceRuleResponseBodyRuleItems) SetValue(v string) *GetResourceRuleResponseBodyRuleItems {
	s.Value = &v
	return s
}

func (s *GetResourceRuleResponseBodyRuleItems) Validate() error {
	return dara.Validate(s)
}
