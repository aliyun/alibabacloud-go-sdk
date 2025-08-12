// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventRuleTargetListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeEventRuleTargetListRequest
	GetRegionId() *string
	SetRuleName(v string) *DescribeEventRuleTargetListRequest
	GetRuleName() *string
}

type DescribeEventRuleTargetListRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the event-triggered alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// testRule
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DescribeEventRuleTargetListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventRuleTargetListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventRuleTargetListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEventRuleTargetListRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DescribeEventRuleTargetListRequest) SetRegionId(v string) *DescribeEventRuleTargetListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEventRuleTargetListRequest) SetRuleName(v string) *DescribeEventRuleTargetListRequest {
	s.RuleName = &v
	return s
}

func (s *DescribeEventRuleTargetListRequest) Validate() error {
	return dara.Validate(s)
}
