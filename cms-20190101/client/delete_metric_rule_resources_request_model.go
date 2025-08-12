// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRuleResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResources(v string) *DeleteMetricRuleResourcesRequest
	GetResources() *string
	SetRuleId(v string) *DeleteMetricRuleResourcesRequest
	GetRuleId() *string
}

type DeleteMetricRuleResourcesRequest struct {
	// The resources that are associated with the alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"instanceId":"i-uf6hm9lnlzsarrc7****"}]
	Resources *string `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// The ID of the alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// rr-bp18017n6iolv****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteMetricRuleResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRuleResourcesRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleResourcesRequest) GetResources() *string {
	return s.Resources
}

func (s *DeleteMetricRuleResourcesRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteMetricRuleResourcesRequest) SetResources(v string) *DeleteMetricRuleResourcesRequest {
	s.Resources = &v
	return s
}

func (s *DeleteMetricRuleResourcesRequest) SetRuleId(v string) *DeleteMetricRuleResourcesRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteMetricRuleResourcesRequest) Validate() error {
	return dara.Validate(s)
}
