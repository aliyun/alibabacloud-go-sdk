// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScheduledScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v bool) *ModifyScheduledScalingRuleRequest
	GetEnable() *bool
	SetInstanceId(v string) *ModifyScheduledScalingRuleRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyScheduledScalingRuleRequest
	GetRegionId() *string
	SetRuleName(v string) *ModifyScheduledScalingRuleRequest
	GetRuleName() *string
}

type ModifyScheduledScalingRuleRequest struct {
	// Specifies whether to enable the scheduled scaling rule. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// >  If the scaling task is scheduled to execute only once and you want to enable the scheduled scaling rule, make sure that the value of this parameter is at least 30 minutes later than the current point in time.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// contact-id
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s ModifyScheduledScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScheduledScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyScheduledScalingRuleRequest) GetEnable() *bool {
	return s.Enable
}

func (s *ModifyScheduledScalingRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyScheduledScalingRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyScheduledScalingRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyScheduledScalingRuleRequest) SetEnable(v bool) *ModifyScheduledScalingRuleRequest {
	s.Enable = &v
	return s
}

func (s *ModifyScheduledScalingRuleRequest) SetInstanceId(v string) *ModifyScheduledScalingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyScheduledScalingRuleRequest) SetRegionId(v string) *ModifyScheduledScalingRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyScheduledScalingRuleRequest) SetRuleName(v string) *ModifyScheduledScalingRuleRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyScheduledScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
