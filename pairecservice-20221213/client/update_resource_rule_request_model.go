// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateResourceRuleRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateResourceRuleRequest
	GetInstanceId() *string
	SetMetricOperationType(v string) *UpdateResourceRuleRequest
	GetMetricOperationType() *string
	SetMetricPullInfo(v string) *UpdateResourceRuleRequest
	GetMetricPullInfo() *string
	SetMetricPullPeriod(v string) *UpdateResourceRuleRequest
	GetMetricPullPeriod() *string
	SetName(v string) *UpdateResourceRuleRequest
	GetName() *string
	SetRuleComputingDefinition(v string) *UpdateResourceRuleRequest
	GetRuleComputingDefinition() *string
}

type UpdateResourceRuleRequest struct {
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
}

func (s UpdateResourceRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateResourceRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateResourceRuleRequest) GetMetricOperationType() *string {
	return s.MetricOperationType
}

func (s *UpdateResourceRuleRequest) GetMetricPullInfo() *string {
	return s.MetricPullInfo
}

func (s *UpdateResourceRuleRequest) GetMetricPullPeriod() *string {
	return s.MetricPullPeriod
}

func (s *UpdateResourceRuleRequest) GetName() *string {
	return s.Name
}

func (s *UpdateResourceRuleRequest) GetRuleComputingDefinition() *string {
	return s.RuleComputingDefinition
}

func (s *UpdateResourceRuleRequest) SetDescription(v string) *UpdateResourceRuleRequest {
	s.Description = &v
	return s
}

func (s *UpdateResourceRuleRequest) SetInstanceId(v string) *UpdateResourceRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateResourceRuleRequest) SetMetricOperationType(v string) *UpdateResourceRuleRequest {
	s.MetricOperationType = &v
	return s
}

func (s *UpdateResourceRuleRequest) SetMetricPullInfo(v string) *UpdateResourceRuleRequest {
	s.MetricPullInfo = &v
	return s
}

func (s *UpdateResourceRuleRequest) SetMetricPullPeriod(v string) *UpdateResourceRuleRequest {
	s.MetricPullPeriod = &v
	return s
}

func (s *UpdateResourceRuleRequest) SetName(v string) *UpdateResourceRuleRequest {
	s.Name = &v
	return s
}

func (s *UpdateResourceRuleRequest) SetRuleComputingDefinition(v string) *UpdateResourceRuleRequest {
	s.RuleComputingDefinition = &v
	return s
}

func (s *UpdateResourceRuleRequest) Validate() error {
	return dara.Validate(s)
}
