// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContainerDefenseRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *AddContainerDefenseRuleShrinkRequest
	GetDescription() *string
	SetRuleAction(v int32) *AddContainerDefenseRuleShrinkRequest
	GetRuleAction() *int32
	SetRuleId(v int64) *AddContainerDefenseRuleShrinkRequest
	GetRuleId() *int64
	SetRuleName(v string) *AddContainerDefenseRuleShrinkRequest
	GetRuleName() *string
	SetRuleSwitch(v int32) *AddContainerDefenseRuleShrinkRequest
	GetRuleSwitch() *int32
	SetRuleType(v int32) *AddContainerDefenseRuleShrinkRequest
	GetRuleType() *int32
	SetScope(v []*AddContainerDefenseRuleShrinkRequestScope) *AddContainerDefenseRuleShrinkRequest
	GetScope() []*AddContainerDefenseRuleShrinkRequestScope
	SetWhitelistShrink(v string) *AddContainerDefenseRuleShrinkRequest
	GetWhitelistShrink() *string
}

type AddContainerDefenseRuleShrinkRequest struct {
	// The description of the rule.
	//
	// example:
	//
	// test-proc-defense
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The action that is performed when the rule is hit. Valid values:
	//
	// 	- **1**: alert
	//
	// 	- **2**: block
	//
	// example:
	//
	// 1
	RuleAction *int32 `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The ID of the rule. You do not need to manually specify the ID.
	//
	// example:
	//
	// 500018
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// auto-test-rule-lt9umq
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The switch of the rule. Valid values:
	//
	// 	- **0**: off
	//
	// 	- **1**: on
	//
	// example:
	//
	// 1
	RuleSwitch *int32 `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The rule type. Valid values:
	//
	// 	- 2: user-defined rules
	//
	// > Only the value 2 is supported.
	//
	// example:
	//
	// 2
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The scope.
	Scope []*AddContainerDefenseRuleShrinkRequestScope `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
	// The whitelist.
	WhitelistShrink *string `json:"Whitelist,omitempty" xml:"Whitelist,omitempty"`
}

func (s AddContainerDefenseRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddContainerDefenseRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddContainerDefenseRuleShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *AddContainerDefenseRuleShrinkRequest) GetRuleAction() *int32 {
	return s.RuleAction
}

func (s *AddContainerDefenseRuleShrinkRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *AddContainerDefenseRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *AddContainerDefenseRuleShrinkRequest) GetRuleSwitch() *int32 {
	return s.RuleSwitch
}

func (s *AddContainerDefenseRuleShrinkRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *AddContainerDefenseRuleShrinkRequest) GetScope() []*AddContainerDefenseRuleShrinkRequestScope {
	return s.Scope
}

func (s *AddContainerDefenseRuleShrinkRequest) GetWhitelistShrink() *string {
	return s.WhitelistShrink
}

func (s *AddContainerDefenseRuleShrinkRequest) SetDescription(v string) *AddContainerDefenseRuleShrinkRequest {
	s.Description = &v
	return s
}

func (s *AddContainerDefenseRuleShrinkRequest) SetRuleAction(v int32) *AddContainerDefenseRuleShrinkRequest {
	s.RuleAction = &v
	return s
}

func (s *AddContainerDefenseRuleShrinkRequest) SetRuleId(v int64) *AddContainerDefenseRuleShrinkRequest {
	s.RuleId = &v
	return s
}

func (s *AddContainerDefenseRuleShrinkRequest) SetRuleName(v string) *AddContainerDefenseRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *AddContainerDefenseRuleShrinkRequest) SetRuleSwitch(v int32) *AddContainerDefenseRuleShrinkRequest {
	s.RuleSwitch = &v
	return s
}

func (s *AddContainerDefenseRuleShrinkRequest) SetRuleType(v int32) *AddContainerDefenseRuleShrinkRequest {
	s.RuleType = &v
	return s
}

func (s *AddContainerDefenseRuleShrinkRequest) SetScope(v []*AddContainerDefenseRuleShrinkRequestScope) *AddContainerDefenseRuleShrinkRequest {
	s.Scope = v
	return s
}

func (s *AddContainerDefenseRuleShrinkRequest) SetWhitelistShrink(v string) *AddContainerDefenseRuleShrinkRequest {
	s.WhitelistShrink = &v
	return s
}

func (s *AddContainerDefenseRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type AddContainerDefenseRuleShrinkRequestScope struct {
	// Specifies whether to include all namespaces. Valid values:
	//
	// 	- **0**: You can use the Namespaces parameter to specify the namespaces to include.
	//
	// 	- **1**: All namespaces are included.
	//
	// example:
	//
	// 0
	AllNamespace *int32 `json:"AllNamespace,omitempty" xml:"AllNamespace,omitempty"`
	// The ID of the cluster.
	//
	// >  You can call the [DescribeGroupedContainerInstances](https://help.aliyun.com/document_detail/182997.html) operation to query the IDs of clusters.
	//
	// example:
	//
	// 8e2***75b
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The namespaces to include.
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
}

func (s AddContainerDefenseRuleShrinkRequestScope) String() string {
	return dara.Prettify(s)
}

func (s AddContainerDefenseRuleShrinkRequestScope) GoString() string {
	return s.String()
}

func (s *AddContainerDefenseRuleShrinkRequestScope) GetAllNamespace() *int32 {
	return s.AllNamespace
}

func (s *AddContainerDefenseRuleShrinkRequestScope) GetClusterId() *string {
	return s.ClusterId
}

func (s *AddContainerDefenseRuleShrinkRequestScope) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *AddContainerDefenseRuleShrinkRequestScope) SetAllNamespace(v int32) *AddContainerDefenseRuleShrinkRequestScope {
	s.AllNamespace = &v
	return s
}

func (s *AddContainerDefenseRuleShrinkRequestScope) SetClusterId(v string) *AddContainerDefenseRuleShrinkRequestScope {
	s.ClusterId = &v
	return s
}

func (s *AddContainerDefenseRuleShrinkRequestScope) SetNamespaces(v []*string) *AddContainerDefenseRuleShrinkRequestScope {
	s.Namespaces = v
	return s
}

func (s *AddContainerDefenseRuleShrinkRequestScope) Validate() error {
	return dara.Validate(s)
}
