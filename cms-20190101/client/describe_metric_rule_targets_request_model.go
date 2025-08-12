// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetricRuleTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeMetricRuleTargetsRequest
	GetRegionId() *string
	SetRuleId(v string) *DescribeMetricRuleTargetsRequest
	GetRuleId() *string
}

type DescribeMetricRuleTargetsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the alert rule.
	//
	// For information about how to obtain the ID of an alert rule, see [DescribeMetricRuleList](https://help.aliyun.com/document_detail/114941.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ae06917_75a8c43178ab66****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DescribeMetricRuleTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetricRuleTargetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMetricRuleTargetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMetricRuleTargetsRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DescribeMetricRuleTargetsRequest) SetRegionId(v string) *DescribeMetricRuleTargetsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMetricRuleTargetsRequest) SetRuleId(v string) *DescribeMetricRuleTargetsRequest {
	s.RuleId = &v
	return s
}

func (s *DescribeMetricRuleTargetsRequest) Validate() error {
	return dara.Validate(s)
}
