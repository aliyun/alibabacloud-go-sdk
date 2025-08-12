// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRuleTargetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteMetricRuleTargetsRequest
	GetRegionId() *string
	SetRuleId(v string) *DeleteMetricRuleTargetsRequest
	GetRuleId() *string
	SetTargetIds(v []*string) *DeleteMetricRuleTargetsRequest
	GetTargetIds() []*string
}

type DeleteMetricRuleTargetsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the alert rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// ruleId-xxxxxx
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The resource IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	TargetIds []*string `json:"TargetIds,omitempty" xml:"TargetIds,omitempty" type:"Repeated"`
}

func (s DeleteMetricRuleTargetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRuleTargetsRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleTargetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteMetricRuleTargetsRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteMetricRuleTargetsRequest) GetTargetIds() []*string {
	return s.TargetIds
}

func (s *DeleteMetricRuleTargetsRequest) SetRegionId(v string) *DeleteMetricRuleTargetsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMetricRuleTargetsRequest) SetRuleId(v string) *DeleteMetricRuleTargetsRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteMetricRuleTargetsRequest) SetTargetIds(v []*string) *DeleteMetricRuleTargetsRequest {
	s.TargetIds = v
	return s
}

func (s *DeleteMetricRuleTargetsRequest) Validate() error {
	return dara.Validate(s)
}
