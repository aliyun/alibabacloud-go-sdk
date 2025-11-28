// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScanRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateScanRuleRequest
	GetInstanceId() *string
	SetNamespaces(v []*string) *CreateScanRuleRequest
	GetNamespaces() []*string
	SetRepoNames(v []*string) *CreateScanRuleRequest
	GetRepoNames() []*string
	SetRepoTagFilterPattern(v string) *CreateScanRuleRequest
	GetRepoTagFilterPattern() *string
	SetRuleName(v string) *CreateScanRuleRequest
	GetRuleName() *string
	SetScanScope(v string) *CreateScanRuleRequest
	GetScanScope() *string
	SetScanType(v string) *CreateScanRuleRequest
	GetScanType() *string
	SetTriggerType(v string) *CreateScanRuleRequest
	GetTriggerType() *string
}

type CreateScanRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cri-dqwc**********
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
	// .*
	RepoTagFilterPattern *string `json:"RepoTagFilterPattern,omitempty" xml:"RepoTagFilterPattern,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// NAMESPACE
	ScanScope *string `json:"ScanScope,omitempty" xml:"ScanScope,omitempty"`
	// example:
	//
	// VUL
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AUTO
	TriggerType *string `json:"TriggerType,omitempty" xml:"TriggerType,omitempty"`
}

func (s CreateScanRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScanRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateScanRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateScanRuleRequest) GetNamespaces() []*string {
	return s.Namespaces
}

func (s *CreateScanRuleRequest) GetRepoNames() []*string {
	return s.RepoNames
}

func (s *CreateScanRuleRequest) GetRepoTagFilterPattern() *string {
	return s.RepoTagFilterPattern
}

func (s *CreateScanRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateScanRuleRequest) GetScanScope() *string {
	return s.ScanScope
}

func (s *CreateScanRuleRequest) GetScanType() *string {
	return s.ScanType
}

func (s *CreateScanRuleRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *CreateScanRuleRequest) SetInstanceId(v string) *CreateScanRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScanRuleRequest) SetNamespaces(v []*string) *CreateScanRuleRequest {
	s.Namespaces = v
	return s
}

func (s *CreateScanRuleRequest) SetRepoNames(v []*string) *CreateScanRuleRequest {
	s.RepoNames = v
	return s
}

func (s *CreateScanRuleRequest) SetRepoTagFilterPattern(v string) *CreateScanRuleRequest {
	s.RepoTagFilterPattern = &v
	return s
}

func (s *CreateScanRuleRequest) SetRuleName(v string) *CreateScanRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateScanRuleRequest) SetScanScope(v string) *CreateScanRuleRequest {
	s.ScanScope = &v
	return s
}

func (s *CreateScanRuleRequest) SetScanType(v string) *CreateScanRuleRequest {
	s.ScanType = &v
	return s
}

func (s *CreateScanRuleRequest) SetTriggerType(v string) *CreateScanRuleRequest {
	s.TriggerType = &v
	return s
}

func (s *CreateScanRuleRequest) Validate() error {
	return dara.Validate(s)
}
