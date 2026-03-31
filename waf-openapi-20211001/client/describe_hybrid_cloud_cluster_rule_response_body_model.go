// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudClusterRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterRule(v *DescribeHybridCloudClusterRuleResponseBodyClusterRule) *DescribeHybridCloudClusterRuleResponseBody
	GetClusterRule() *DescribeHybridCloudClusterRuleResponseBodyClusterRule
	SetRequestId(v string) *DescribeHybridCloudClusterRuleResponseBody
	GetRequestId() *string
}

type DescribeHybridCloudClusterRuleResponseBody struct {
	// The details of the rule.
	ClusterRule *DescribeHybridCloudClusterRuleResponseBodyClusterRule `json:"ClusterRule,omitempty" xml:"ClusterRule,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1F29A6D2-9EB6-526D-A997-36888**99CB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHybridCloudClusterRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClusterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClusterRuleResponseBody) GetClusterRule() *DescribeHybridCloudClusterRuleResponseBodyClusterRule {
	return s.ClusterRule
}

func (s *DescribeHybridCloudClusterRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudClusterRuleResponseBody) SetClusterRule(v *DescribeHybridCloudClusterRuleResponseBodyClusterRule) *DescribeHybridCloudClusterRuleResponseBody {
	s.ClusterRule = v
	return s
}

func (s *DescribeHybridCloudClusterRuleResponseBody) SetRequestId(v string) *DescribeHybridCloudClusterRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudClusterRuleResponseBody) Validate() error {
	if s.ClusterRule != nil {
		if err := s.ClusterRule.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeHybridCloudClusterRuleResponseBodyClusterRule struct {
	ClusterRuleResourceId *string `json:"ClusterRuleResourceId,omitempty" xml:"ClusterRuleResourceId,omitempty"`
	// The configuration of the rule.
	//
	// example:
	//
	// {\\"check_mode\\":\\"all\\",\\"exclude\\":{\\"exact\\":[],\\"regex\\":[]}}
	RuleConfig *string `json:"RuleConfig,omitempty" xml:"RuleConfig,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **on**: enabled.
	//
	// 	- **off**: disabled.
	//
	// example:
	//
	// on
	RuleStatus *string `json:"RuleStatus,omitempty" xml:"RuleStatus,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- **pullin**: The traffic redirection rule of the hybrid cloud cluster.
	//
	// example:
	//
	// pullin
	RuleType *string `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
}

func (s DescribeHybridCloudClusterRuleResponseBodyClusterRule) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudClusterRuleResponseBodyClusterRule) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudClusterRuleResponseBodyClusterRule) GetClusterRuleResourceId() *string {
	return s.ClusterRuleResourceId
}

func (s *DescribeHybridCloudClusterRuleResponseBodyClusterRule) GetRuleConfig() *string {
	return s.RuleConfig
}

func (s *DescribeHybridCloudClusterRuleResponseBodyClusterRule) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *DescribeHybridCloudClusterRuleResponseBodyClusterRule) GetRuleType() *string {
	return s.RuleType
}

func (s *DescribeHybridCloudClusterRuleResponseBodyClusterRule) SetClusterRuleResourceId(v string) *DescribeHybridCloudClusterRuleResponseBodyClusterRule {
	s.ClusterRuleResourceId = &v
	return s
}

func (s *DescribeHybridCloudClusterRuleResponseBodyClusterRule) SetRuleConfig(v string) *DescribeHybridCloudClusterRuleResponseBodyClusterRule {
	s.RuleConfig = &v
	return s
}

func (s *DescribeHybridCloudClusterRuleResponseBodyClusterRule) SetRuleStatus(v string) *DescribeHybridCloudClusterRuleResponseBodyClusterRule {
	s.RuleStatus = &v
	return s
}

func (s *DescribeHybridCloudClusterRuleResponseBodyClusterRule) SetRuleType(v string) *DescribeHybridCloudClusterRuleResponseBodyClusterRule {
	s.RuleType = &v
	return s
}

func (s *DescribeHybridCloudClusterRuleResponseBodyClusterRule) Validate() error {
	return dara.Validate(s)
}
