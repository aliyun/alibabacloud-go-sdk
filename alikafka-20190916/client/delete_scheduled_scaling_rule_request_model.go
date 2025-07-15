// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteScheduledScalingRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteScheduledScalingRuleRequest
	GetRegionId() *string
	SetRuleName(v string) *DeleteScheduledScalingRuleRequest
	GetRuleName() *string
}

type DeleteScheduledScalingRuleRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_serverless-cn-vxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the scheduled scaling rule.
	//
	// >  You can delete only rules that are disabled and rules that are scheduled only once and have been executed.
	//
	// This parameter is required.
	//
	// example:
	//
	// rule-name-test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DeleteScheduledScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteScheduledScalingRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteScheduledScalingRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteScheduledScalingRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *DeleteScheduledScalingRuleRequest) SetInstanceId(v string) *DeleteScheduledScalingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteScheduledScalingRuleRequest) SetRegionId(v string) *DeleteScheduledScalingRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteScheduledScalingRuleRequest) SetRuleName(v string) *DeleteScheduledScalingRuleRequest {
	s.RuleName = &v
	return s
}

func (s *DeleteScheduledScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
