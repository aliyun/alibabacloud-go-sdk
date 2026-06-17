// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetricRuleResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOverwrite(v string) *CreateMetricRuleResourcesRequest
	GetOverwrite() *string
	SetResources(v string) *CreateMetricRuleResourcesRequest
	GetResources() *string
	SetRuleId(v string) *CreateMetricRuleResourcesRequest
	GetRuleId() *string
}

type CreateMetricRuleResourcesRequest struct {
	// Specifies whether to overwrite. Valid values:
	//
	// - true: overwrites. The resources submitted this time overwrite the previously associated resources. That is, full modification is performed.
	//
	// - false: does not overwrite. The resources submitted this time do not overwrite the previously associated resources (the associated resources are the historical associated resources plus the resources submitted this time). That is, incremental modification is performed.
	//
	// example:
	//
	// false
	Overwrite *string `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The associated resources. The value is in the JSON array format.
	//
	// > A maximum of 100 resource instances can be added at a time, and an alert rule can be associated with a maximum of 3,000 instances.
	//
	// >
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"instanceId":"i-a2d5q7pm3f9yr29e****"}]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The ID of the alert rule.
	//
	// example:
	//
	// i-2ze3w55tr2rcpejpcfap_59c96b85-0339-4f35-ba66-ae4e34d3****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s CreateMetricRuleResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMetricRuleResourcesRequest) GoString() string {
	return s.String()
}

func (s *CreateMetricRuleResourcesRequest) GetOverwrite() *string {
	return s.Overwrite
}

func (s *CreateMetricRuleResourcesRequest) GetResources() *string {
	return s.Resources
}

func (s *CreateMetricRuleResourcesRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateMetricRuleResourcesRequest) SetOverwrite(v string) *CreateMetricRuleResourcesRequest {
	s.Overwrite = &v
	return s
}

func (s *CreateMetricRuleResourcesRequest) SetResources(v string) *CreateMetricRuleResourcesRequest {
	s.Resources = &v
	return s
}

func (s *CreateMetricRuleResourcesRequest) SetRuleId(v string) *CreateMetricRuleResourcesRequest {
	s.RuleId = &v
	return s
}

func (s *CreateMetricRuleResourcesRequest) Validate() error {
	return dara.Validate(s)
}
