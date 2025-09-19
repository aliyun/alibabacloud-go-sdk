// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInterceptionRuleDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetInterceptionRuleDetailRequest
	GetClusterId() *string
	SetRuleId(v string) *GetInterceptionRuleDetailRequest
	GetRuleId() *string
}

type GetInterceptionRuleDetailRequest struct {
	// The ID of the container cluster.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of container clusters.
	//
	// This parameter is required.
	//
	// example:
	//
	// c7f60fdabc84xxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The ID of the rule.
	//
	// > You can call the [ListInterceptionRulePage](~~ListInterceptionRulePage~~) operation to query the IDs of rules.
	//
	// This parameter is required.
	//
	// example:
	//
	// 500002
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s GetInterceptionRuleDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInterceptionRuleDetailRequest) GoString() string {
	return s.String()
}

func (s *GetInterceptionRuleDetailRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetInterceptionRuleDetailRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *GetInterceptionRuleDetailRequest) SetClusterId(v string) *GetInterceptionRuleDetailRequest {
	s.ClusterId = &v
	return s
}

func (s *GetInterceptionRuleDetailRequest) SetRuleId(v string) *GetInterceptionRuleDetailRequest {
	s.RuleId = &v
	return s
}

func (s *GetInterceptionRuleDetailRequest) Validate() error {
	return dara.Validate(s)
}
