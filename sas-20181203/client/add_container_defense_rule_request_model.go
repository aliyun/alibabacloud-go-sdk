// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddContainerDefenseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *AddContainerDefenseRuleRequest
	GetDescription() *string
	SetRuleAction(v int32) *AddContainerDefenseRuleRequest
	GetRuleAction() *int32
	SetRuleId(v int64) *AddContainerDefenseRuleRequest
	GetRuleId() *int64
	SetRuleName(v string) *AddContainerDefenseRuleRequest
	GetRuleName() *string
	SetRuleSwitch(v int32) *AddContainerDefenseRuleRequest
	GetRuleSwitch() *int32
	SetRuleType(v int32) *AddContainerDefenseRuleRequest
	GetRuleType() *int32
	SetScope(v []*AddContainerDefenseRuleRequestScope) *AddContainerDefenseRuleRequest
	GetScope() []*AddContainerDefenseRuleRequestScope
	SetWhitelist(v *AddContainerDefenseRuleRequestWhitelist) *AddContainerDefenseRuleRequest
	GetWhitelist() *AddContainerDefenseRuleRequestWhitelist
}

type AddContainerDefenseRuleRequest struct {
	// The description.
	//
	// example:
	//
	// test-proc-defense
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The action to take when the rule is matched. Valid values:
	//
	// - **1**: Alert.
	//
	// - **2**: Block.
	//
	// example:
	//
	// 1
	RuleAction *int32 `json:"RuleAction,omitempty" xml:"RuleAction,omitempty"`
	// The rule ID. You do not need to specify this parameter when creating a rule.
	//
	// example:
	//
	// 500018
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// auto-test-rule-lt9umq
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule switch. Valid values:
	//
	// - **0**: Disabled.
	//
	// - **1**: Enabled.
	//
	// example:
	//
	// 1
	RuleSwitch *int32 `json:"RuleSwitch,omitempty" xml:"RuleSwitch,omitempty"`
	// The rule type. Valid values:
	//
	// - 2: user rule
	//
	// 	Notice: Only the value 2 is supported.
	//
	// example:
	//
	// 2
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// The scope. This parameter is required. Specify at least one Scope entry, such as Scope.1.AllNamespace=1, which indicates that the rule applies to all namespaces. If this parameter is not specified, the API returns a 400 error.
	Scope []*AddContainerDefenseRuleRequestScope `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
	// The whitelist.
	Whitelist *AddContainerDefenseRuleRequestWhitelist `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Struct"`
}

func (s AddContainerDefenseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddContainerDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *AddContainerDefenseRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *AddContainerDefenseRuleRequest) GetRuleAction() *int32 {
	return s.RuleAction
}

func (s *AddContainerDefenseRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *AddContainerDefenseRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *AddContainerDefenseRuleRequest) GetRuleSwitch() *int32 {
	return s.RuleSwitch
}

func (s *AddContainerDefenseRuleRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *AddContainerDefenseRuleRequest) GetScope() []*AddContainerDefenseRuleRequestScope {
	return s.Scope
}

func (s *AddContainerDefenseRuleRequest) GetWhitelist() *AddContainerDefenseRuleRequestWhitelist {
	return s.Whitelist
}

func (s *AddContainerDefenseRuleRequest) SetDescription(v string) *AddContainerDefenseRuleRequest {
	s.Description = &v
	return s
}

func (s *AddContainerDefenseRuleRequest) SetRuleAction(v int32) *AddContainerDefenseRuleRequest {
	s.RuleAction = &v
	return s
}

func (s *AddContainerDefenseRuleRequest) SetRuleId(v int64) *AddContainerDefenseRuleRequest {
	s.RuleId = &v
	return s
}

func (s *AddContainerDefenseRuleRequest) SetRuleName(v string) *AddContainerDefenseRuleRequest {
	s.RuleName = &v
	return s
}

func (s *AddContainerDefenseRuleRequest) SetRuleSwitch(v int32) *AddContainerDefenseRuleRequest {
	s.RuleSwitch = &v
	return s
}

func (s *AddContainerDefenseRuleRequest) SetRuleType(v int32) *AddContainerDefenseRuleRequest {
	s.RuleType = &v
	return s
}

func (s *AddContainerDefenseRuleRequest) SetScope(v []*AddContainerDefenseRuleRequestScope) *AddContainerDefenseRuleRequest {
	s.Scope = v
	return s
}

func (s *AddContainerDefenseRuleRequest) SetWhitelist(v *AddContainerDefenseRuleRequestWhitelist) *AddContainerDefenseRuleRequest {
	s.Whitelist = v
	return s
}

func (s *AddContainerDefenseRuleRequest) Validate() error {
	if s.Scope != nil {
		for _, item := range s.Scope {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Whitelist != nil {
		if err := s.Whitelist.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddContainerDefenseRuleRequestScope struct {
	// Specifies whether to include all namespaces. Valid values:
	//
	// - **0**: Specifies the namespaces to include by using the Namespaces parameter.
	//
	// - **1**: Includes all namespaces.
	//
	// example:
	//
	// 0
	AllNamespace *int32 `json:"AllNamespace,omitempty" xml:"AllNamespace,omitempty"`
	// The cluster ID.
	//
	// > You can call the [DescribeGroupedContainerInstances](https://help.aliyun.com/document_detail/182997.html) operation to obtain this parameter.
	//
	// example:
	//
	// 8e2***75b
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The list of included namespaces.
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
}

func (s AddContainerDefenseRuleRequestScope) String() string {
	return dara.Prettify(s)
}

func (s AddContainerDefenseRuleRequestScope) GoString() string {
	return s.String()
}

func (s *AddContainerDefenseRuleRequestScope) GetAllNamespace() *int32 {
	return s.AllNamespace
}

func (s *AddContainerDefenseRuleRequestScope) GetClusterId() *string {
	return s.ClusterId
}

func (s *AddContainerDefenseRuleRequestScope) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *AddContainerDefenseRuleRequestScope) SetAllNamespace(v int32) *AddContainerDefenseRuleRequestScope {
	s.AllNamespace = &v
	return s
}

func (s *AddContainerDefenseRuleRequestScope) SetClusterId(v string) *AddContainerDefenseRuleRequestScope {
	s.ClusterId = &v
	return s
}

func (s *AddContainerDefenseRuleRequestScope) SetNamespaces(v []*string) *AddContainerDefenseRuleRequestScope {
	s.Namespaces = v
	return s
}

func (s *AddContainerDefenseRuleRequestScope) Validate() error {
	return dara.Validate(s)
}

type AddContainerDefenseRuleRequestWhitelist struct {
	// Deprecated
	//
	// The file hash.	Notice: This parameter is not supported.
	Hash []*string `json:"Hash,omitempty" xml:"Hash,omitempty" type:"Repeated"`
	// The list of images to whitelist.
	Image []*string `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
	// The list of file paths to whitelist.
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
}

func (s AddContainerDefenseRuleRequestWhitelist) String() string {
	return dara.Prettify(s)
}

func (s AddContainerDefenseRuleRequestWhitelist) GoString() string {
	return s.String()
}

func (s *AddContainerDefenseRuleRequestWhitelist) GetHash() []*string {
	return s.Hash
}

func (s *AddContainerDefenseRuleRequestWhitelist) GetImage() []*string {
	return s.Image
}

func (s *AddContainerDefenseRuleRequestWhitelist) GetPath() []*string {
	return s.Path
}

func (s *AddContainerDefenseRuleRequestWhitelist) SetHash(v []*string) *AddContainerDefenseRuleRequestWhitelist {
	s.Hash = v
	return s
}

func (s *AddContainerDefenseRuleRequestWhitelist) SetImage(v []*string) *AddContainerDefenseRuleRequestWhitelist {
	s.Image = v
	return s
}

func (s *AddContainerDefenseRuleRequestWhitelist) SetPath(v []*string) *AddContainerDefenseRuleRequestWhitelist {
	s.Path = v
	return s
}

func (s *AddContainerDefenseRuleRequestWhitelist) Validate() error {
	return dara.Validate(s)
}
