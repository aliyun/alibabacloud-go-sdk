// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRemediationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigRuleId(v string) *DescribeRemediationRequest
	GetConfigRuleId() *string
	SetRemediationId(v string) *DescribeRemediationRequest
	GetRemediationId() *string
}

type DescribeRemediationRequest struct {
	// The rule ID.
	//
	// example:
	//
	// cr-3184626622af003****
	ConfigRuleId *string `json:"ConfigRuleId,omitempty" xml:"ConfigRuleId,omitempty"`
	// The ID of the remediation configuration.
	//
	// example:
	//
	// crr-f381cf0c1c2f004e****
	RemediationId *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
}

func (s DescribeRemediationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRemediationRequest) GoString() string {
	return s.String()
}

func (s *DescribeRemediationRequest) GetConfigRuleId() *string {
	return s.ConfigRuleId
}

func (s *DescribeRemediationRequest) GetRemediationId() *string {
	return s.RemediationId
}

func (s *DescribeRemediationRequest) SetConfigRuleId(v string) *DescribeRemediationRequest {
	s.ConfigRuleId = &v
	return s
}

func (s *DescribeRemediationRequest) SetRemediationId(v string) *DescribeRemediationRequest {
	s.RemediationId = &v
	return s
}

func (s *DescribeRemediationRequest) Validate() error {
	return dara.Validate(s)
}
