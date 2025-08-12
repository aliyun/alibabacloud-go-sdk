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
	// Specifies whether to overwrite existing resources. Valid values:
	//
	// 	- true: The resources submitted this time overwrite the previously associated resources.
	//
	// 	- false: The resources submitted this time do not overwrite the previously associated resources. The associated resources after submission include the previously associated resources and the resources submitted this time.
	//
	// example:
	//
	// false
	Overwrite *string `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The resources that are associated with the alert rule. Set the value to a JSON array.
	//
	// >  You can add up to 100 resources each time. An alert rule can be associated with up to 3,000 resources.
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
