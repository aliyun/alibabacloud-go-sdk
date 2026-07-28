// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScalingRule interface {
	dara.Model
	String() string
	GoString() string
	SetAdjustInfos(v []*ScalingRuleAdjustInfos) *ScalingRule
	GetAdjustInfos() []*ScalingRuleAdjustInfos
	SetCronStr(v string) *ScalingRule
	GetCronStr() *string
	SetDisabled(v bool) *ScalingRule
	GetDisabled() *bool
	SetEndTime(v int64) *ScalingRule
	GetEndTime() *int64
	SetRuleId(v string) *ScalingRule
	GetRuleId() *string
	SetScalingRuleName(v string) *ScalingRule
	GetScalingRuleName() *string
	SetStartTime(v int64) *ScalingRule
	GetStartTime() *int64
	SetTimeZone(v string) *ScalingRule
	GetTimeZone() *string
}

type ScalingRule struct {
	// The adjustment information of the scaling rule.
	AdjustInfos []*ScalingRuleAdjustInfos `json:"adjustInfos,omitempty" xml:"adjustInfos,omitempty" type:"Repeated"`
	// The cron expression for the scaling schedule.
	CronStr *string `json:"cronStr,omitempty" xml:"cronStr,omitempty"`
	// Indicates whether the scaling rule is disabled.
	Disabled *bool `json:"disabled,omitempty" xml:"disabled,omitempty"`
	// The end time of the scaling rule. Unit: milliseconds.
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The ID of the scaling rule.
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The name of the scaling rule.
	ScalingRuleName *string `json:"scalingRuleName,omitempty" xml:"scalingRuleName,omitempty"`
	// The start time of the scaling rule. Unit: milliseconds.
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The time zone of the scaling rule.
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s ScalingRule) String() string {
	return dara.Prettify(s)
}

func (s ScalingRule) GoString() string {
	return s.String()
}

func (s *ScalingRule) GetAdjustInfos() []*ScalingRuleAdjustInfos {
	return s.AdjustInfos
}

func (s *ScalingRule) GetCronStr() *string {
	return s.CronStr
}

func (s *ScalingRule) GetDisabled() *bool {
	return s.Disabled
}

func (s *ScalingRule) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ScalingRule) GetRuleId() *string {
	return s.RuleId
}

func (s *ScalingRule) GetScalingRuleName() *string {
	return s.ScalingRuleName
}

func (s *ScalingRule) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ScalingRule) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ScalingRule) SetAdjustInfos(v []*ScalingRuleAdjustInfos) *ScalingRule {
	s.AdjustInfos = v
	return s
}

func (s *ScalingRule) SetCronStr(v string) *ScalingRule {
	s.CronStr = &v
	return s
}

func (s *ScalingRule) SetDisabled(v bool) *ScalingRule {
	s.Disabled = &v
	return s
}

func (s *ScalingRule) SetEndTime(v int64) *ScalingRule {
	s.EndTime = &v
	return s
}

func (s *ScalingRule) SetRuleId(v string) *ScalingRule {
	s.RuleId = &v
	return s
}

func (s *ScalingRule) SetScalingRuleName(v string) *ScalingRule {
	s.ScalingRuleName = &v
	return s
}

func (s *ScalingRule) SetStartTime(v int64) *ScalingRule {
	s.StartTime = &v
	return s
}

func (s *ScalingRule) SetTimeZone(v string) *ScalingRule {
	s.TimeZone = &v
	return s
}

func (s *ScalingRule) Validate() error {
	if s.AdjustInfos != nil {
		for _, item := range s.AdjustInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ScalingRuleAdjustInfos struct {
	// The type of the component to be adjusted.
	ComponentType *string `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// The target value for the adjustment.
	TargetValue *string `json:"targetValue,omitempty" xml:"targetValue,omitempty"`
}

func (s ScalingRuleAdjustInfos) String() string {
	return dara.Prettify(s)
}

func (s ScalingRuleAdjustInfos) GoString() string {
	return s.String()
}

func (s *ScalingRuleAdjustInfos) GetComponentType() *string {
	return s.ComponentType
}

func (s *ScalingRuleAdjustInfos) GetTargetValue() *string {
	return s.TargetValue
}

func (s *ScalingRuleAdjustInfos) SetComponentType(v string) *ScalingRuleAdjustInfos {
	s.ComponentType = &v
	return s
}

func (s *ScalingRuleAdjustInfos) SetTargetValue(v string) *ScalingRuleAdjustInfos {
	s.TargetValue = &v
	return s
}

func (s *ScalingRuleAdjustInfos) Validate() error {
	return dara.Validate(s)
}
