// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScalingRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyScalingRuleRequest
	GetInstanceId() *string
	SetNewTriggerType(v string) *ModifyScalingRuleRequest
	GetNewTriggerType() *string
	SetNodeGroupId(v string) *ModifyScalingRuleRequest
	GetNodeGroupId() *string
	SetOldTriggerType(v string) *ModifyScalingRuleRequest
	GetOldTriggerType() *string
	SetRule(v string) *ModifyScalingRuleRequest
	GetRule() *string
	SetScalingRuleId(v string) *ModifyScalingRuleRequest
	GetScalingRuleId() *string
}

type ModifyScalingRuleRequest struct {
	// The instance ID.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new trigger type.
	//
	// This parameter is required.
	//
	// example:
	//
	// TIME_TRIGGER
	NewTriggerType *string `json:"NewTriggerType,omitempty" xml:"NewTriggerType,omitempty"`
	// The ID of the compute group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-3d5ce6454354****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The original trigger type.
	//
	// This parameter is required.
	//
	// example:
	//
	// TIME_TRIGGER
	OldTriggerType *string `json:"OldTriggerType,omitempty" xml:"OldTriggerType,omitempty"`
	// The content of the rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "nodeNumber": 1,
	//
	//   "scalingRuleName": "test1",
	//
	//   "scalingOutRule": {
	//
	//     "year": 2026,
	//
	//     "month": 3,
	//
	//     "day": 2,
	//
	//     "hour": 5,
	//
	//     "minute": 50,
	//
	//     "second": 0,
	//
	//     "recurrenceInterval": 1,
	//
	//     "recurrenceType": "DAILY",
	//
	//     "recurrenceValues": null
	//
	//   },
	//
	//   "scalingInRule": {
	//
	//     "year": 2026,
	//
	//     "month": 3,
	//
	//     "day": 2,
	//
	//     "hour": 6,
	//
	//     "minute": 50,
	//
	//     "second": 0,
	//
	//     "recurrenceInterval": 1,
	//
	//     "recurrenceType": "DAILY",
	//
	//     "recurrenceValues": null
	//
	//   }
	//
	// }
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The ID of the scaling rule to be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// asr-bp10s4t9l21u9qtyrtco
	ScalingRuleId *string `json:"ScalingRuleId,omitempty" xml:"ScalingRuleId,omitempty"`
}

func (s ModifyScalingRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyScalingRuleRequest) GetNewTriggerType() *string {
	return s.NewTriggerType
}

func (s *ModifyScalingRuleRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ModifyScalingRuleRequest) GetOldTriggerType() *string {
	return s.OldTriggerType
}

func (s *ModifyScalingRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *ModifyScalingRuleRequest) GetScalingRuleId() *string {
	return s.ScalingRuleId
}

func (s *ModifyScalingRuleRequest) SetInstanceId(v string) *ModifyScalingRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetNewTriggerType(v string) *ModifyScalingRuleRequest {
	s.NewTriggerType = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetNodeGroupId(v string) *ModifyScalingRuleRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetOldTriggerType(v string) *ModifyScalingRuleRequest {
	s.OldTriggerType = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetRule(v string) *ModifyScalingRuleRequest {
	s.Rule = &v
	return s
}

func (s *ModifyScalingRuleRequest) SetScalingRuleId(v string) *ModifyScalingRuleRequest {
	s.ScalingRuleId = &v
	return s
}

func (s *ModifyScalingRuleRequest) Validate() error {
	return dara.Validate(s)
}
