// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScanRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateScanRuleShrinkRequest
	GetInstanceId() *string
	SetNamespacesShrink(v string) *CreateScanRuleShrinkRequest
	GetNamespacesShrink() *string
	SetRepoNamesShrink(v string) *CreateScanRuleShrinkRequest
	GetRepoNamesShrink() *string
	SetRepoTagFilterPattern(v string) *CreateScanRuleShrinkRequest
	GetRepoTagFilterPattern() *string
	SetRuleName(v string) *CreateScanRuleShrinkRequest
	GetRuleName() *string
	SetScanScope(v string) *CreateScanRuleShrinkRequest
	GetScanScope() *string
	SetScanType(v string) *CreateScanRuleShrinkRequest
	GetScanType() *string
	SetTriggerType(v string) *CreateScanRuleShrinkRequest
	GetTriggerType() *string
}

type CreateScanRuleShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cri-dqwc**********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// if can be null:
	// true
	NamespacesShrink *string `json:"Namespaces,omitempty" xml:"Namespaces,omitempty"`
	// if can be null:
	// true
	RepoNamesShrink *string `json:"RepoNames,omitempty" xml:"RepoNames,omitempty"`
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

func (s CreateScanRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateScanRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateScanRuleShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateScanRuleShrinkRequest) GetNamespacesShrink() *string {
	return s.NamespacesShrink
}

func (s *CreateScanRuleShrinkRequest) GetRepoNamesShrink() *string {
	return s.RepoNamesShrink
}

func (s *CreateScanRuleShrinkRequest) GetRepoTagFilterPattern() *string {
	return s.RepoTagFilterPattern
}

func (s *CreateScanRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateScanRuleShrinkRequest) GetScanScope() *string {
	return s.ScanScope
}

func (s *CreateScanRuleShrinkRequest) GetScanType() *string {
	return s.ScanType
}

func (s *CreateScanRuleShrinkRequest) GetTriggerType() *string {
	return s.TriggerType
}

func (s *CreateScanRuleShrinkRequest) SetInstanceId(v string) *CreateScanRuleShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateScanRuleShrinkRequest) SetNamespacesShrink(v string) *CreateScanRuleShrinkRequest {
	s.NamespacesShrink = &v
	return s
}

func (s *CreateScanRuleShrinkRequest) SetRepoNamesShrink(v string) *CreateScanRuleShrinkRequest {
	s.RepoNamesShrink = &v
	return s
}

func (s *CreateScanRuleShrinkRequest) SetRepoTagFilterPattern(v string) *CreateScanRuleShrinkRequest {
	s.RepoTagFilterPattern = &v
	return s
}

func (s *CreateScanRuleShrinkRequest) SetRuleName(v string) *CreateScanRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *CreateScanRuleShrinkRequest) SetScanScope(v string) *CreateScanRuleShrinkRequest {
	s.ScanScope = &v
	return s
}

func (s *CreateScanRuleShrinkRequest) SetScanType(v string) *CreateScanRuleShrinkRequest {
	s.ScanType = &v
	return s
}

func (s *CreateScanRuleShrinkRequest) SetTriggerType(v string) *CreateScanRuleShrinkRequest {
	s.TriggerType = &v
	return s
}

func (s *CreateScanRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
