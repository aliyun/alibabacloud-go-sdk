// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSilencePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateOrUpdateSilencePolicyResponseBody
	GetRequestId() *string
	SetSilencePolicy(v *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) *CreateOrUpdateSilencePolicyResponseBody
	GetSilencePolicy() *CreateOrUpdateSilencePolicyResponseBodySilencePolicy
}

type CreateOrUpdateSilencePolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A5EC8221-08F2-4C95-9AF1-49FD998C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The silence policy.
	SilencePolicy *CreateOrUpdateSilencePolicyResponseBodySilencePolicy `json:"SilencePolicy,omitempty" xml:"SilencePolicy,omitempty" type:"Struct"`
}

func (s CreateOrUpdateSilencePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSilencePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSilencePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateSilencePolicyResponseBody) GetSilencePolicy() *CreateOrUpdateSilencePolicyResponseBodySilencePolicy {
	return s.SilencePolicy
}

func (s *CreateOrUpdateSilencePolicyResponseBody) SetRequestId(v string) *CreateOrUpdateSilencePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBody) SetSilencePolicy(v *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) *CreateOrUpdateSilencePolicyResponseBody {
	s.SilencePolicy = v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBody) Validate() error {
	if s.SilencePolicy != nil {
		if err := s.SilencePolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrUpdateSilencePolicyResponseBodySilencePolicy struct {
	// The effective type. Valid values: PERMANENT: The policy is effective permanently. CYCLE_EFFECT: The policy is effective cyclically. CUSTOM_TIME: The policy is effective during a custom time period.
	//
	// example:
	//
	// PERMANENT
	EffectiveTimeType *string `json:"EffectiveTimeType,omitempty" xml:"EffectiveTimeType,omitempty"`
	// The ID of the silence policy.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// A list of matching rules.
	MatchingRules []*CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules `json:"MatchingRules,omitempty" xml:"MatchingRules,omitempty" type:"Repeated"`
	// The name of the silence policy.
	//
	// example:
	//
	// silencepolicy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to enable the silence policy. Valid values: enable and disable.
	//
	// example:
	//
	// enable
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Effective period. Valid values: DAY: daily WEEK: weekly
	//
	// example:
	//
	// WEEK
	TimePeriod *string `json:"TimePeriod,omitempty" xml:"TimePeriod,omitempty"`
	// The time period during which the silence policy is effective.
	//
	// example:
	//
	// [{startWeek: "1", startTime: "00:00", endTime: "11:59", endWeek:"7"}]
	TimeSlots *string `json:"TimeSlots,omitempty" xml:"TimeSlots,omitempty"`
}

func (s CreateOrUpdateSilencePolicyResponseBodySilencePolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSilencePolicyResponseBodySilencePolicy) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) GetEffectiveTimeType() *string {
	return s.EffectiveTimeType
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) GetId() *int64 {
	return s.Id
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) GetMatchingRules() []*CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules {
	return s.MatchingRules
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) GetName() *string {
	return s.Name
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) GetState() *string {
	return s.State
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) GetTimePeriod() *string {
	return s.TimePeriod
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) GetTimeSlots() *string {
	return s.TimeSlots
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) SetEffectiveTimeType(v string) *CreateOrUpdateSilencePolicyResponseBodySilencePolicy {
	s.EffectiveTimeType = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) SetId(v int64) *CreateOrUpdateSilencePolicyResponseBodySilencePolicy {
	s.Id = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) SetMatchingRules(v []*CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules) *CreateOrUpdateSilencePolicyResponseBodySilencePolicy {
	s.MatchingRules = v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) SetName(v string) *CreateOrUpdateSilencePolicyResponseBodySilencePolicy {
	s.Name = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) SetState(v string) *CreateOrUpdateSilencePolicyResponseBodySilencePolicy {
	s.State = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) SetTimePeriod(v string) *CreateOrUpdateSilencePolicyResponseBodySilencePolicy {
	s.TimePeriod = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) SetTimeSlots(v string) *CreateOrUpdateSilencePolicyResponseBodySilencePolicy {
	s.TimeSlots = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicy) Validate() error {
	if s.MatchingRules != nil {
		for _, item := range s.MatchingRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules struct {
	// A list of matching conditions.
	MatchingConditions []*CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions `json:"MatchingConditions,omitempty" xml:"MatchingConditions,omitempty" type:"Repeated"`
}

func (s CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules) GetMatchingConditions() []*CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions {
	return s.MatchingConditions
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules) SetMatchingConditions(v []*CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules {
	s.MatchingConditions = v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRules) Validate() error {
	if s.MatchingConditions != nil {
		for _, item := range s.MatchingConditions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions struct {
	// The key of the matching condition.
	//
	// example:
	//
	// altertname
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The logical operator of the matching condition. Valid values:
	//
	// 	- `eq`: equal to
	//
	// 	- `neq`: not equal to
	//
	// 	- `in`: contains
	//
	// 	- `nin`: does not contain
	//
	// 	- `re`: regular expression match
	//
	// 	- `nre`: regular expression mismatch
	//
	// example:
	//
	// eq
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
	// The value of the matching condition.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) GetKey() *string {
	return s.Key
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) GetOperator() *string {
	return s.Operator
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) GetValue() *string {
	return s.Value
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) SetKey(v string) *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions {
	s.Key = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) SetOperator(v string) *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions {
	s.Operator = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) SetValue(v string) *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions {
	s.Value = &v
	return s
}

func (s *CreateOrUpdateSilencePolicyResponseBodySilencePolicyMatchingRulesMatchingConditions) Validate() error {
	return dara.Validate(s)
}
