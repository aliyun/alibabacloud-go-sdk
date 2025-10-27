// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContainerDefenseRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyContainerDefenseRuleRequest
	GetDescription() *string
	SetRuleAction(v int32) *ModifyContainerDefenseRuleRequest
	GetRuleAction() *int32
	SetRuleId(v int64) *ModifyContainerDefenseRuleRequest
	GetRuleId() *int64
	SetRuleName(v string) *ModifyContainerDefenseRuleRequest
	GetRuleName() *string
	SetRuleSwitch(v int32) *ModifyContainerDefenseRuleRequest
	GetRuleSwitch() *int32
	SetRuleType(v int32) *ModifyContainerDefenseRuleRequest
	GetRuleType() *int32
	SetScope(v []*ModifyContainerDefenseRuleRequestScope) *ModifyContainerDefenseRuleRequest
	GetScope() []*ModifyContainerDefenseRuleRequestScope
	SetWhitelist(v *ModifyContainerDefenseRuleRequestWhitelist) *ModifyContainerDefenseRuleRequest
	GetWhitelist() *ModifyContainerDefenseRuleRequestWhitelist
}

type ModifyContainerDefenseRuleRequest struct {
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
	Scope []*ModifyContainerDefenseRuleRequestScope `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
	// The whitelist.
	Whitelist *ModifyContainerDefenseRuleRequestWhitelist `json:"Whitelist,omitempty" xml:"Whitelist,omitempty" type:"Struct"`
}

func (s ModifyContainerDefenseRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerDefenseRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyContainerDefenseRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyContainerDefenseRuleRequest) GetRuleAction() *int32 {
	return s.RuleAction
}

func (s *ModifyContainerDefenseRuleRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ModifyContainerDefenseRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *ModifyContainerDefenseRuleRequest) GetRuleSwitch() *int32 {
	return s.RuleSwitch
}

func (s *ModifyContainerDefenseRuleRequest) GetRuleType() *int32 {
	return s.RuleType
}

func (s *ModifyContainerDefenseRuleRequest) GetScope() []*ModifyContainerDefenseRuleRequestScope {
	return s.Scope
}

func (s *ModifyContainerDefenseRuleRequest) GetWhitelist() *ModifyContainerDefenseRuleRequestWhitelist {
	return s.Whitelist
}

func (s *ModifyContainerDefenseRuleRequest) SetDescription(v string) *ModifyContainerDefenseRuleRequest {
	s.Description = &v
	return s
}

func (s *ModifyContainerDefenseRuleRequest) SetRuleAction(v int32) *ModifyContainerDefenseRuleRequest {
	s.RuleAction = &v
	return s
}

func (s *ModifyContainerDefenseRuleRequest) SetRuleId(v int64) *ModifyContainerDefenseRuleRequest {
	s.RuleId = &v
	return s
}

func (s *ModifyContainerDefenseRuleRequest) SetRuleName(v string) *ModifyContainerDefenseRuleRequest {
	s.RuleName = &v
	return s
}

func (s *ModifyContainerDefenseRuleRequest) SetRuleSwitch(v int32) *ModifyContainerDefenseRuleRequest {
	s.RuleSwitch = &v
	return s
}

func (s *ModifyContainerDefenseRuleRequest) SetRuleType(v int32) *ModifyContainerDefenseRuleRequest {
	s.RuleType = &v
	return s
}

func (s *ModifyContainerDefenseRuleRequest) SetScope(v []*ModifyContainerDefenseRuleRequestScope) *ModifyContainerDefenseRuleRequest {
	s.Scope = v
	return s
}

func (s *ModifyContainerDefenseRuleRequest) SetWhitelist(v *ModifyContainerDefenseRuleRequestWhitelist) *ModifyContainerDefenseRuleRequest {
	s.Whitelist = v
	return s
}

func (s *ModifyContainerDefenseRuleRequest) Validate() error {
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

type ModifyContainerDefenseRuleRequestScope struct {
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

func (s ModifyContainerDefenseRuleRequestScope) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerDefenseRuleRequestScope) GoString() string {
	return s.String()
}

func (s *ModifyContainerDefenseRuleRequestScope) GetAllNamespace() *int32 {
	return s.AllNamespace
}

func (s *ModifyContainerDefenseRuleRequestScope) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyContainerDefenseRuleRequestScope) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *ModifyContainerDefenseRuleRequestScope) SetAllNamespace(v int32) *ModifyContainerDefenseRuleRequestScope {
	s.AllNamespace = &v
	return s
}

func (s *ModifyContainerDefenseRuleRequestScope) SetClusterId(v string) *ModifyContainerDefenseRuleRequestScope {
	s.ClusterId = &v
	return s
}

func (s *ModifyContainerDefenseRuleRequestScope) SetNamespaces(v []*string) *ModifyContainerDefenseRuleRequestScope {
	s.Namespaces = v
	return s
}

func (s *ModifyContainerDefenseRuleRequestScope) Validate() error {
	return dara.Validate(s)
}

type ModifyContainerDefenseRuleRequestWhitelist struct {
	// Deprecated
	//
	// The hash values of the files that need to be added to the whitelist.
	//
	// >  This parameter is not supported.
	Hash []*string `json:"Hash,omitempty" xml:"Hash,omitempty" type:"Repeated"`
	// The images that need to be added to the whitelist.
	Image []*string `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
	// The paths to the files that need to be added to the whitelist.
	Path []*string `json:"Path,omitempty" xml:"Path,omitempty" type:"Repeated"`
}

func (s ModifyContainerDefenseRuleRequestWhitelist) String() string {
	return dara.Prettify(s)
}

func (s ModifyContainerDefenseRuleRequestWhitelist) GoString() string {
	return s.String()
}

func (s *ModifyContainerDefenseRuleRequestWhitelist) GetHash() []*string {
	return s.Hash
}

func (s *ModifyContainerDefenseRuleRequestWhitelist) GetImage() []*string {
	return s.Image
}

func (s *ModifyContainerDefenseRuleRequestWhitelist) GetPath() []*string {
	return s.Path
}

func (s *ModifyContainerDefenseRuleRequestWhitelist) SetHash(v []*string) *ModifyContainerDefenseRuleRequestWhitelist {
	s.Hash = v
	return s
}

func (s *ModifyContainerDefenseRuleRequestWhitelist) SetImage(v []*string) *ModifyContainerDefenseRuleRequestWhitelist {
	s.Image = v
	return s
}

func (s *ModifyContainerDefenseRuleRequestWhitelist) SetPath(v []*string) *ModifyContainerDefenseRuleRequestWhitelist {
	s.Path = v
	return s
}

func (s *ModifyContainerDefenseRuleRequestWhitelist) Validate() error {
	return dara.Validate(s)
}
