// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateResourceRuleRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateResourceRuleRequest
	GetInstanceId() *string
	SetMetricOperationType(v string) *CreateResourceRuleRequest
	GetMetricOperationType() *string
	SetMetricPullInfo(v string) *CreateResourceRuleRequest
	GetMetricPullInfo() *string
	SetMetricPullPeriod(v string) *CreateResourceRuleRequest
	GetMetricPullPeriod() *string
	SetName(v string) *CreateResourceRuleRequest
	GetName() *string
	SetRuleComputingDefinition(v string) *CreateResourceRuleRequest
	GetRuleComputingDefinition() *string
	SetRuleItems(v []*CreateResourceRuleRequestRuleItems) *CreateResourceRuleRequest
	GetRuleItems() []*CreateResourceRuleRequestRuleItems
}

type CreateResourceRuleRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MetricOperationType *string `json:"MetricOperationType,omitempty" xml:"MetricOperationType,omitempty"`
	MetricPullInfo      *string `json:"MetricPullInfo,omitempty" xml:"MetricPullInfo,omitempty"`
	MetricPullPeriod    *string `json:"MetricPullPeriod,omitempty" xml:"MetricPullPeriod,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	RuleComputingDefinition *string `json:"RuleComputingDefinition,omitempty" xml:"RuleComputingDefinition,omitempty"`
	// This parameter is required.
	RuleItems []*CreateResourceRuleRequestRuleItems `json:"RuleItems,omitempty" xml:"RuleItems,omitempty" type:"Repeated"`
}

func (s CreateResourceRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateResourceRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateResourceRuleRequest) GetMetricOperationType() *string {
	return s.MetricOperationType
}

func (s *CreateResourceRuleRequest) GetMetricPullInfo() *string {
	return s.MetricPullInfo
}

func (s *CreateResourceRuleRequest) GetMetricPullPeriod() *string {
	return s.MetricPullPeriod
}

func (s *CreateResourceRuleRequest) GetName() *string {
	return s.Name
}

func (s *CreateResourceRuleRequest) GetRuleComputingDefinition() *string {
	return s.RuleComputingDefinition
}

func (s *CreateResourceRuleRequest) GetRuleItems() []*CreateResourceRuleRequestRuleItems {
	return s.RuleItems
}

func (s *CreateResourceRuleRequest) SetDescription(v string) *CreateResourceRuleRequest {
	s.Description = &v
	return s
}

func (s *CreateResourceRuleRequest) SetInstanceId(v string) *CreateResourceRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateResourceRuleRequest) SetMetricOperationType(v string) *CreateResourceRuleRequest {
	s.MetricOperationType = &v
	return s
}

func (s *CreateResourceRuleRequest) SetMetricPullInfo(v string) *CreateResourceRuleRequest {
	s.MetricPullInfo = &v
	return s
}

func (s *CreateResourceRuleRequest) SetMetricPullPeriod(v string) *CreateResourceRuleRequest {
	s.MetricPullPeriod = &v
	return s
}

func (s *CreateResourceRuleRequest) SetName(v string) *CreateResourceRuleRequest {
	s.Name = &v
	return s
}

func (s *CreateResourceRuleRequest) SetRuleComputingDefinition(v string) *CreateResourceRuleRequest {
	s.RuleComputingDefinition = &v
	return s
}

func (s *CreateResourceRuleRequest) SetRuleItems(v []*CreateResourceRuleRequestRuleItems) *CreateResourceRuleRequest {
	s.RuleItems = v
	return s
}

func (s *CreateResourceRuleRequest) Validate() error {
	return dara.Validate(s)
}

type CreateResourceRuleRequestRuleItems struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	MaxValue *float64 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// This parameter is required.
	MinValue *float64 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateResourceRuleRequestRuleItems) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceRuleRequestRuleItems) GoString() string {
	return s.String()
}

func (s *CreateResourceRuleRequestRuleItems) GetDescription() *string {
	return s.Description
}

func (s *CreateResourceRuleRequestRuleItems) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *CreateResourceRuleRequestRuleItems) GetMinValue() *float64 {
	return s.MinValue
}

func (s *CreateResourceRuleRequestRuleItems) GetName() *string {
	return s.Name
}

func (s *CreateResourceRuleRequestRuleItems) GetValue() *float64 {
	return s.Value
}

func (s *CreateResourceRuleRequestRuleItems) SetDescription(v string) *CreateResourceRuleRequestRuleItems {
	s.Description = &v
	return s
}

func (s *CreateResourceRuleRequestRuleItems) SetMaxValue(v float64) *CreateResourceRuleRequestRuleItems {
	s.MaxValue = &v
	return s
}

func (s *CreateResourceRuleRequestRuleItems) SetMinValue(v float64) *CreateResourceRuleRequestRuleItems {
	s.MinValue = &v
	return s
}

func (s *CreateResourceRuleRequestRuleItems) SetName(v string) *CreateResourceRuleRequestRuleItems {
	s.Name = &v
	return s
}

func (s *CreateResourceRuleRequestRuleItems) SetValue(v float64) *CreateResourceRuleRequestRuleItems {
	s.Value = &v
	return s
}

func (s *CreateResourceRuleRequestRuleItems) Validate() error {
	return dara.Validate(s)
}
