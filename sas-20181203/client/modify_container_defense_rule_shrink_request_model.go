// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContainerDefenseRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyContainerDefenseRuleShrinkRequest
	GetDescription() *string
	SetRuleAction(v int32) *ModifyContainerDefenseRuleShrinkRequest
	GetRuleAction() *int32
	SetRuleId(v int64) *ModifyContainerDefenseRuleShrinkRequest
	GetRuleId() *int64
	SetRuleName(v string) *ModifyContainerDefenseRuleShrinkRequest
	GetRuleName() *string
	SetRuleSwitch(v int32) *ModifyContainerDefenseRuleShrinkRequest
	GetRuleSwitch() *int32
	SetRuleType(v int32) *ModifyContainerDefenseRuleShrinkRequest
	GetRuleType() *int32
	SetScope(v []*ModifyContainerDefenseRuleShrinkRequestScope) *ModifyContainerDefenseRuleShrinkRequest
	GetScope() []*ModifyContainerDefenseRuleShrinkRequestScope
	SetWhitelistShrink(v string) *ModifyContainerDefenseRuleShrinkRequest
	GetWhitelistShrink() *string
}

type ModifyContainerDefenseRuleShrinkRequest struct {
	// The description of the rule.
	//
	// example:
	//
	// Prevent non-mirror programs from starting in containers
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The action specified in the rule. Valid values:
	//
	// 	- **1**: alert
	//
	// 	- **2**: block
	//
	// example:
	//
	// 1
	RuleAction *int32 `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The ID of the rule.
	//
	// >  You can call the [ListContainerDefenseRule](https://help.aliyun.com/document_detail/2590599.html) operation to query the IDs of rules.
	//
	// example:
	//
	// 123
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// text-001
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The status of the rule. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	RuleSwitch *int32 `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The type of the rule. Valid values:
	//
	// 	- 1: system rule
	//
	// 	- 2: custom rule
	//
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The effective scope of the rule.
	Scope []*ModifyContainerDefenseRuleShrinkRequestScope `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
	// The whitelist.
	WhitelistShrink *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s ModifyContainerDefenseRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerDefenseRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyContainerDefenseRuleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyContainerDefenseRuleShrinkRequest) GetRuleAction() *int32 {
	return s.RuleAction
}

func (s *ModifyContainerDefenseRuleShrinkRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ModifyContainerDefenseRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyContainerDefenseRuleShrinkRequest) GetRuleSwitch() *int32 {
	return s.RuleSwitch
}

func (s *ModifyContainerDefenseRuleShrinkRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *ModifyContainerDefenseRuleShrinkRequest) GetScope() []*ModifyContainerDefenseRuleShrinkRequestScope {
	return s.Scope
}

func (s *ModifyContainerDefenseRuleShrinkRequest) GetWhitelistShrink() *string {
	return s.WhitelistShrink
}

func (s *ModifyContainerDefenseRuleShrinkRequest) SetDescription(v string) *ModifyContainerDefenseRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *ModifyContainerDefenseRuleShrinkRequest) SetRuleAction(v int32) *ModifyContainerDefenseRuleShrinkRequest {
	s.RuleAction = &v
	return s
}

func (s *ModifyContainerDefenseRuleShrinkRequest) SetRuleId(v int64) *ModifyContainerDefenseRuleShrinkRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyContainerDefenseRuleShrinkRequest) SetRuleName(v string) *ModifyContainerDefenseRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyContainerDefenseRuleShrinkRequest) SetRuleSwitch(v int32) *ModifyContainerDefenseRuleShrinkRequest {
	s.RuleSwitch = &v
	return s
}

func (s *ModifyContainerDefenseRuleShrinkRequest) SetRuleType(v int32) *ModifyContainerDefenseRuleShrinkRequest {
	s.RuleType = &v
	return s
}

func (s *ModifyContainerDefenseRuleShrinkRequest) SetScope(v []*ModifyContainerDefenseRuleShrinkRequestScope) *ModifyContainerDefenseRuleShrinkRequest {
	s.Scope = v
	return s
}

func (s *ModifyContainerDefenseRuleShrinkRequest) SetWhitelistShrink(v string) *ModifyContainerDefenseRuleShrinkRequest {
	s.WhitelistShrink = &v
	return s
}

func (s *ModifyContainerDefenseRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyContainerDefenseRuleShrinkRequestScope struct {
	// Specifies whether to include all namespaces. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	AllNamespace *int32 `json:"AllNamespace,omitempty" xml:"AllNamespace,omitempty"`
	// The ID of the cluster on which the rule takes effect.
	//
	// >  You can call the [DescribeGroupedContainerInstances](~~DescribeGroupedContainerInstances~~) operation to query the IDs of clusters.
	//
	// example:
	//
	// c54b***1501
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The namespaces.
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
}

func (s ModifyContainerDefenseRuleShrinkRequestScope) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerDefenseRuleShrinkRequestScope) GoString() string {
	return s.String()
}

func (s *ModifyContainerDefenseRuleShrinkRequestScope) GetAllNamespace() *int32 {
	return s.AllNamespace
}

func (s *ModifyContainerDefenseRuleShrinkRequestScope) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyContainerDefenseRuleShrinkRequestScope) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *ModifyContainerDefenseRuleShrinkRequestScope) SetAllNamespace(v int32) *ModifyContainerDefenseRuleShrinkRequestScope {
	s.AllNamespace = &v
	return s
}

func (s *ModifyContainerDefenseRuleShrinkRequestScope) SetClusterId(v string) *ModifyContainerDefenseRuleShrinkRequestScope {
	s.ClusterId = &v
	return s
}

func (s *ModifyContainerDefenseRuleShrinkRequestScope) SetNamespaces(v []*string) *ModifyContainerDefenseRuleShrinkRequestScope {
	s.Namespaces = v
	return s
}

func (s *ModifyContainerDefenseRuleShrinkRequestScope) Validate() error {
	return dara.Validate(s)
}
