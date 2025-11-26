// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoScalingPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAutoScalingPolicyResponseBody
	GetRequestId() *string
	SetScalingPolicy(v *GetAutoScalingPolicyResponseBodyScalingPolicy) *GetAutoScalingPolicyResponseBody
	GetScalingPolicy() *GetAutoScalingPolicyResponseBodyScalingPolicy
}

type GetAutoScalingPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The auto scaling policy.
	ScalingPolicy *GetAutoScalingPolicyResponseBodyScalingPolicy `json:"ScalingPolicy,omitempty" xml:"ScalingPolicy,omitempty" type:"Struct"`
}

func (s GetAutoScalingPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoScalingPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutoScalingPolicyResponseBody) GetScalingPolicy() *GetAutoScalingPolicyResponseBodyScalingPolicy {
	return s.ScalingPolicy
}

func (s *GetAutoScalingPolicyResponseBody) SetRequestId(v string) *GetAutoScalingPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoScalingPolicyResponseBody) SetScalingPolicy(v *GetAutoScalingPolicyResponseBodyScalingPolicy) *GetAutoScalingPolicyResponseBody {
	s.ScalingPolicy = v
	return s
}

func (s *GetAutoScalingPolicyResponseBody) Validate() error {
	if s.ScalingPolicy != nil {
		if err := s.ScalingPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutoScalingPolicyResponseBodyScalingPolicy struct {
	// The cluster ID.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The maximum and minimum number of nodes in the node group.
	Constraints *GetAutoScalingPolicyResponseBodyScalingPolicyConstraints `json:"Constraints,omitempty" xml:"Constraints,omitempty" type:"Struct"`
	// The ID of the node group.
	//
	// example:
	//
	// ng-869471354ecd****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The ID of the scaling policy.
	//
	// example:
	//
	// asp-asduwe23znl***
	ScalingPolicyId *string `json:"ScalingPolicyId,omitempty" xml:"ScalingPolicyId,omitempty"`
	// The auto scaling rules.
	ScalingRules []*GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules `json:"ScalingRules,omitempty" xml:"ScalingRules,omitempty" type:"Repeated"`
}

func (s GetAutoScalingPolicyResponseBodyScalingPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingPolicyResponseBodyScalingPolicy) GoString() string {
	return s.String()
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicy) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicy) GetConstraints() *GetAutoScalingPolicyResponseBodyScalingPolicyConstraints {
	return s.Constraints
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicy) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicy) GetScalingPolicyId() *string {
	return s.ScalingPolicyId
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicy) GetScalingRules() []*GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules {
	return s.ScalingRules
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicy) SetClusterId(v string) *GetAutoScalingPolicyResponseBodyScalingPolicy {
	s.ClusterId = &v
	return s
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicy) SetConstraints(v *GetAutoScalingPolicyResponseBodyScalingPolicyConstraints) *GetAutoScalingPolicyResponseBodyScalingPolicy {
	s.Constraints = v
	return s
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicy) SetNodeGroupId(v string) *GetAutoScalingPolicyResponseBodyScalingPolicy {
	s.NodeGroupId = &v
	return s
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicy) SetScalingPolicyId(v string) *GetAutoScalingPolicyResponseBodyScalingPolicy {
	s.ScalingPolicyId = &v
	return s
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicy) SetScalingRules(v []*GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) *GetAutoScalingPolicyResponseBodyScalingPolicy {
	s.ScalingRules = v
	return s
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicy) Validate() error {
	if s.Constraints != nil {
		if err := s.Constraints.Validate(); err != nil {
			return err
		}
	}
	if s.ScalingRules != nil {
		for _, item := range s.ScalingRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAutoScalingPolicyResponseBodyScalingPolicyConstraints struct {
	// The maximum number of nodes in the node group. Default value: 2000.
	//
	// example:
	//
	// 2000
	MaxCapacity *int32 `json:"MaxCapacity,omitempty" xml:"MaxCapacity,omitempty"`
	// The minimum number of nodes in the node group. Default value: 0.
	//
	// example:
	//
	// 0
	MinCapacity *int32 `json:"MinCapacity,omitempty" xml:"MinCapacity,omitempty"`
}

func (s GetAutoScalingPolicyResponseBodyScalingPolicyConstraints) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingPolicyResponseBodyScalingPolicyConstraints) GoString() string {
	return s.String()
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyConstraints) GetMaxCapacity() *int32 {
	return s.MaxCapacity
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyConstraints) GetMinCapacity() *int32 {
	return s.MinCapacity
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyConstraints) SetMaxCapacity(v int32) *GetAutoScalingPolicyResponseBodyScalingPolicyConstraints {
	s.MaxCapacity = &v
	return s
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyConstraints) SetMinCapacity(v int32) *GetAutoScalingPolicyResponseBodyScalingPolicyConstraints {
	s.MinCapacity = &v
	return s
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyConstraints) Validate() error {
	return dara.Validate(s)
}

type GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules struct {
	// The scaling type. Valid values:
	//
	// 	- SCALE_OUT
	//
	// 	- SCALE_IN
	//
	// example:
	//
	// SCALE_OUT
	ActivityType *string `json:"ActivityType,omitempty" xml:"ActivityType,omitempty"`
	// The adjustment type.
	//
	// example:
	//
	// CHANGE_IN_CAPACITY
	AdjustmentType *string `json:"AdjustmentType,omitempty" xml:"AdjustmentType,omitempty"`
	// The adjustment value. The parameter value must be a positive integer, which indicates the number of instances that you want to add or remove.
	//
	// example:
	//
	// 100
	AdjustmentValue   *int32             `json:"AdjustmentValue,omitempty" xml:"AdjustmentValue,omitempty"`
	CollationTimeZone *CollationTimeZone `json:"CollationTimeZone,omitempty" xml:"CollationTimeZone,omitempty"`
	// The description of load-based scaling.
	MetricsTrigger *MetricsTrigger `json:"MetricsTrigger,omitempty" xml:"MetricsTrigger,omitempty"`
	// The name of the auto scaling rule.
	//
	// example:
	//
	// scaling-out-memory
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The description of time-based scaling.
	TimeTrigger *TimeTrigger `json:"TimeTrigger,omitempty" xml:"TimeTrigger,omitempty"`
	// The type of the auto scaling rule. Valid values:
	//
	// 	- TIME_TRIGGER: time-based scaling
	//
	// 	- METRICS_TRIGGER: load-based scaling
	//
	// example:
	//
	// TIME_TRIGGER
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) String() string {
	return dara.Prettify(s)
}

func (s GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) GoString() string {
	return s.String()
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) GetActivityType() *string {
	return s.ActivityType
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) GetAdjustmentType() *string {
	return s.AdjustmentType
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) GetAdjustmentValue() *int32 {
	return s.AdjustmentValue
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) GetCollationTimeZone() *CollationTimeZone {
	return s.CollationTimeZone
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) GetMetricsTrigger() *MetricsTrigger {
	return s.MetricsTrigger
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) GetRuleName() *string {
	return s.RuleName
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) GetTimeTrigger() *TimeTrigger {
	return s.TimeTrigger
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) GetTriggerType() *string {
	return s.TriggerType
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) SetActivityType(v string) *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules {
	s.ActivityType = &v
	return s
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) SetAdjustmentType(v string) *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules {
	s.AdjustmentType = &v
	return s
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) SetAdjustmentValue(v int32) *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules {
	s.AdjustmentValue = &v
	return s
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) SetCollationTimeZone(v *CollationTimeZone) *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules {
	s.CollationTimeZone = v
	return s
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) SetMetricsTrigger(v *MetricsTrigger) *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules {
	s.MetricsTrigger = v
	return s
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) SetRuleName(v string) *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules {
	s.RuleName = &v
	return s
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) SetTimeTrigger(v *TimeTrigger) *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules {
	s.TimeTrigger = v
	return s
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) SetTriggerType(v string) *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules {
	s.TriggerType = &v
	return s
}

func (s *GetAutoScalingPolicyResponseBodyScalingPolicyScalingRules) Validate() error {
	if s.CollationTimeZone != nil {
		if err := s.CollationTimeZone.Validate(); err != nil {
			return err
		}
	}
	if s.MetricsTrigger != nil {
		if err := s.MetricsTrigger.Validate(); err != nil {
			return err
		}
	}
	if s.TimeTrigger != nil {
		if err := s.TimeTrigger.Validate(); err != nil {
			return err
		}
	}
	return nil
}
