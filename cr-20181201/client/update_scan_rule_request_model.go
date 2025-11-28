// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScanRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateScanRuleRequest
	GetInstanceId() *string
	SetNamespaces(v []*string) *UpdateScanRuleRequest
	GetNamespaces() []*string
	SetRepoNames(v []*string) *UpdateScanRuleRequest
	GetRepoNames() []*string
	SetRepoTagFilterPattern(v string) *UpdateScanRuleRequest
	GetRepoTagFilterPattern() *string
	SetRuleName(v string) *UpdateScanRuleRequest
	GetRuleName() *string
	SetScanRuleId(v string) *UpdateScanRuleRequest
	GetScanRuleId() *string
	SetScanScope(v string) *UpdateScanRuleRequest
	GetScanScope() *string
	SetTriggerType(v string) *UpdateScanRuleRequest
	GetTriggerType() *string
}

type UpdateScanRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cri-4abntrj42twd****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// if can be null:
	// true
	Namespaces []*string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty" type:"Repeated"`
	// if can be null:
	// true
	RepoNames []*string `json:"RepoNames,omitempty" xml:"RepoNames,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// prod-.*
	RepoTagFilterPattern *string `json:"RepoTagFilterPattern,omitempty" xml:"RepoTagFilterPattern,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// scan-test
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// crscnr-3qmkeiuggfpjkfrq
	ScanRuleId *string `json:"ScanRuleId,omitempty" xml:"ScanRuleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// REPO
	ScanScope *string `json:"ScanScope,omitempty" xml:"ScanScope,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AUTO
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s UpdateScanRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateScanRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateScanRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateScanRuleRequest) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *UpdateScanRuleRequest) GetRepoNames() []*string {
	return s.RepoNames
}

func (s *UpdateScanRuleRequest) GetRepoTagFilterPattern() *string {
	return s.RepoTagFilterPattern
}

func (s *UpdateScanRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateScanRuleRequest) GetScanRuleId() *string {
	return s.ScanRuleId
}

func (s *UpdateScanRuleRequest) GetScanScope() *string {
	return s.ScanScope
}

func (s *UpdateScanRuleRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *UpdateScanRuleRequest) SetInstanceId(v string) *UpdateScanRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateScanRuleRequest) SetNamespaces(v []*string) *UpdateScanRuleRequest {
	s.Namespaces = v
	return s
}

func (s *UpdateScanRuleRequest) SetRepoNames(v []*string) *UpdateScanRuleRequest {
	s.RepoNames = v
	return s
}

func (s *UpdateScanRuleRequest) SetRepoTagFilterPattern(v string) *UpdateScanRuleRequest {
	s.RepoTagFilterPattern = &v
	return s
}

func (s *UpdateScanRuleRequest) SetRuleName(v string) *UpdateScanRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateScanRuleRequest) SetScanRuleId(v string) *UpdateScanRuleRequest {
	s.ScanRuleId = &v
	return s
}

func (s *UpdateScanRuleRequest) SetScanScope(v string) *UpdateScanRuleRequest {
	s.ScanScope = &v
	return s
}

func (s *UpdateScanRuleRequest) SetTriggerType(v string) *UpdateScanRuleRequest {
	s.TriggerType = &v
	return s
}

func (s *UpdateScanRuleRequest) Validate() error {
	return dara.Validate(s)
}
