// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInterceptionRuleSwitchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyInterceptionRuleSwitchRequest
	GetClusterId() *string
	SetRuleIds(v string) *ModifyInterceptionRuleSwitchRequest
	GetRuleIds() *string
	SetRuleSwitch(v int32) *ModifyInterceptionRuleSwitchRequest
	GetRuleSwitch() *int32
}

type ModifyInterceptionRuleSwitchRequest struct {
	// The ID of the cluster that you want to modify.
	//
	// > You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to obtain this parameter.
	//
	// example:
	//
	// ce5c29aba99694ade9ba85dc620b4****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The list of rule IDs to operate on. Separate multiple IDs with commas (,).
	//
	// > You can call the [ListInterceptionRulePage](~~ListInterceptionRulePage~~) operation to obtain this parameter.
	//
	// example:
	//
	// 403287
	RuleIds *string `json:"RuleIds,omitempty" xml:"RuleIds,omitempty"`
	// The switch status of the rule. Valid values:
	//
	// - **1**: Enabled.
	//
	// - **0**: Disabled.
	//
	// example:
	//
	// 1
	RuleSwitch *int32 `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
}

func (s ModifyInterceptionRuleSwitchRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInterceptionRuleSwitchRequest) GoString() string {
	return s.String()
}

func (s *ModifyInterceptionRuleSwitchRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyInterceptionRuleSwitchRequest) GetRuleIds() *string {
	return s.RuleIds
}

func (s *ModifyInterceptionRuleSwitchRequest) GetRuleSwitch() *int32 {
	return s.RuleSwitch
}

func (s *ModifyInterceptionRuleSwitchRequest) SetClusterId(v string) *ModifyInterceptionRuleSwitchRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyInterceptionRuleSwitchRequest) SetRuleIds(v string) *ModifyInterceptionRuleSwitchRequest {
	s.RuleIds = &v
	return s
}

func (s *ModifyInterceptionRuleSwitchRequest) SetRuleSwitch(v int32) *ModifyInterceptionRuleSwitchRequest {
	s.RuleSwitch = &v
	return s
}

func (s *ModifyInterceptionRuleSwitchRequest) Validate() error {
	return dara.Validate(s)
}
