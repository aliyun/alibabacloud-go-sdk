// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvDropMetricsRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *DescribeEnvDropMetricsRuleRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *DescribeEnvDropMetricsRuleRequest
	GetRegionId() *string
}

type DescribeEnvDropMetricsRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// env-xxxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEnvDropMetricsRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvDropMetricsRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnvDropMetricsRuleRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DescribeEnvDropMetricsRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEnvDropMetricsRuleRequest) SetEnvironmentId(v string) *DescribeEnvDropMetricsRuleRequest {
	s.EnvironmentId = &v
	return s
}

func (s *DescribeEnvDropMetricsRuleRequest) SetRegionId(v string) *DescribeEnvDropMetricsRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEnvDropMetricsRuleRequest) Validate() error {
	return dara.Validate(s)
}
