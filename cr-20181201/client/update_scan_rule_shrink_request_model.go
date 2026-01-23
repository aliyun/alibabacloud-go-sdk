// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScanRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateScanRuleShrinkRequest
	GetInstanceId() *string
	SetNamespacesShrink(v string) *UpdateScanRuleShrinkRequest
	GetNamespacesShrink() *string
	SetRepoNamesShrink(v string) *UpdateScanRuleShrinkRequest
	GetRepoNamesShrink() *string
	SetRepoTagFilterPattern(v string) *UpdateScanRuleShrinkRequest
	GetRepoTagFilterPattern() *string
	SetRuleName(v string) *UpdateScanRuleShrinkRequest
	GetRuleName() *string
	SetScanRuleId(v string) *UpdateScanRuleShrinkRequest
	GetScanRuleId() *string
	SetScanScope(v string) *UpdateScanRuleShrinkRequest
	GetScanScope() *string
	SetTriggerType(v string) *UpdateScanRuleShrinkRequest
	GetTriggerType() *string
}

type UpdateScanRuleShrinkRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-4abntrj42twd****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of namespaces.
	//
	// 	- When the scan scope is NAMESPACE, this parameter cannot be empty.
	//
	// 	- If the scan scope is REPO, you must specify a unique Namespace for this parameter.
	//
	// if can be null:
	// true
	NamespacesShrink *string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty"`
	// The list of repositories.
	//
	// 	- When the scan scope is NAMESPACE, this parameter must be empty.
	//
	// 	- When the scan scope is REPO, this parameter cannot be empty.
	//
	// if can be null:
	// true
	RepoNamesShrink *string `json:"RepoNames,omitempty" xml:"RepoNames,omitempty"`
	// The tag filtering rules.
	//
	// This parameter is required.
	//
	// example:
	//
	// prod-.*
	RepoTagFilterPattern *string `json:"RepoTagFilterPattern,omitempty" xml:"RepoTagFilterPattern,omitempty"`
	// The rule name.
	//
	// This parameter is required.
	//
	// example:
	//
	// scan-test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crscnr-3qmkeiuggfpjkfrq
	ScanRuleId *string `json:"ScanRuleId,omitempty" xml:"ScanRuleId,omitempty"`
	// The scan scope.
	//
	// Valid values:
	//
	// 	- NAMESPACE: namespace.
	//
	// 	- REPO: repository.
	//
	// This parameter is required.
	//
	// example:
	//
	// REPO
	ScanScope *string `json:"ScanScope,omitempty" xml:"ScanScope,omitempty"`
	// The trigger type.
	//
	// Valid values:
	//
	// 	- AUTO: automatically triggers.
	//
	// 	- MANUAL: manually triggers.
	//
	// This parameter is required.
	//
	// example:
	//
	// AUTO
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s UpdateScanRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScanRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateScanRuleShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateScanRuleShrinkRequest) GetNamespacesShrink() *string {
	return s.NamespacesShrink
}

func (s *UpdateScanRuleShrinkRequest) GetRepoNamesShrink() *string {
	return s.RepoNamesShrink
}

func (s *UpdateScanRuleShrinkRequest) GetRepoTagFilterPattern() *string {
	return s.RepoTagFilterPattern
}

func (s *UpdateScanRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateScanRuleShrinkRequest) GetScanRuleId() *string {
	return s.ScanRuleId
}

func (s *UpdateScanRuleShrinkRequest) GetScanScope() *string {
	return s.ScanScope
}

func (s *UpdateScanRuleShrinkRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *UpdateScanRuleShrinkRequest) SetInstanceId(v string) *UpdateScanRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateScanRuleShrinkRequest) SetNamespacesShrink(v string) *UpdateScanRuleShrinkRequest {
	s.NamespacesShrink = &v
	return s
}

func (s *UpdateScanRuleShrinkRequest) SetRepoNamesShrink(v string) *UpdateScanRuleShrinkRequest {
	s.RepoNamesShrink = &v
	return s
}

func (s *UpdateScanRuleShrinkRequest) SetRepoTagFilterPattern(v string) *UpdateScanRuleShrinkRequest {
	s.RepoTagFilterPattern = &v
	return s
}

func (s *UpdateScanRuleShrinkRequest) SetRuleName(v string) *UpdateScanRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateScanRuleShrinkRequest) SetScanRuleId(v string) *UpdateScanRuleShrinkRequest {
	s.ScanRuleId = &v
	return s
}

func (s *UpdateScanRuleShrinkRequest) SetScanScope(v string) *UpdateScanRuleShrinkRequest {
	s.ScanScope = &v
	return s
}

func (s *UpdateScanRuleShrinkRequest) SetTriggerType(v string) *UpdateScanRuleShrinkRequest {
	s.TriggerType = &v
	return s
}

func (s *UpdateScanRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
