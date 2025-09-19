// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContainerDefenseRuleSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRuleIds(v []*int64) *ModifyContainerDefenseRuleSwitchRequest
	GetRuleIds() []*int64
	SetRuleSwitch(v int32) *ModifyContainerDefenseRuleSwitchRequest
	GetRuleSwitch() *int32
}

type ModifyContainerDefenseRuleSwitchRequest struct {
	// The IDs of the rules.
	RuleIds []*int64 `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
	// The status of the rule. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 0
	RuleSwitch *int32 `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
}

func (s ModifyContainerDefenseRuleSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerDefenseRuleSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyContainerDefenseRuleSwitchRequest) GetRuleIds() []*int64 {
	return s.RuleIds
}

func (s *ModifyContainerDefenseRuleSwitchRequest) GetRuleSwitch() *int32 {
	return s.RuleSwitch
}

func (s *ModifyContainerDefenseRuleSwitchRequest) SetRuleIds(v []*int64) *ModifyContainerDefenseRuleSwitchRequest {
	s.RuleIds = v
	return s
}

func (s *ModifyContainerDefenseRuleSwitchRequest) SetRuleSwitch(v int32) *ModifyContainerDefenseRuleSwitchRequest {
	s.RuleSwitch = &v
	return s
}

func (s *ModifyContainerDefenseRuleSwitchRequest) Validate() error {
	return dara.Validate(s)
}
