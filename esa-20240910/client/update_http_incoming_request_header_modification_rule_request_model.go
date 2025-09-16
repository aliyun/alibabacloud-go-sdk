// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpIncomingRequestHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateHttpIncomingRequestHeaderModificationRuleRequest
	GetConfigId() *int64
	SetRequestHeaderModification(v []*UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) *UpdateHttpIncomingRequestHeaderModificationRuleRequest
	GetRequestHeaderModification() []*UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification
	SetRule(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateHttpIncomingRequestHeaderModificationRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateHttpIncomingRequestHeaderModificationRuleRequest
	GetSiteId() *int64
}

type UpdateHttpIncomingRequestHeaderModificationRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 419717024278528
	ConfigId                  *int64                                                                             `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	RequestHeaderModification []*UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty" type:"Repeated"`
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 568181195163328
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateHttpIncomingRequestHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpIncomingRequestHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) GetRequestHeaderModification() []*UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification {
	return s.RequestHeaderModification
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) SetConfigId(v int64) *UpdateHttpIncomingRequestHeaderModificationRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) SetRequestHeaderModification(v []*UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) *UpdateHttpIncomingRequestHeaderModificationRuleRequest {
	s.RequestHeaderModification = v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) SetRule(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) SetRuleEnable(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) SetRuleName(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) SetSequence(v int32) *UpdateHttpIncomingRequestHeaderModificationRuleRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) SetSiteId(v int64) *UpdateHttpIncomingRequestHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification struct {
	// This parameter is required.
	//
	// example:
	//
	// headerName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// headerValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) GoString() string {
	return s.String()
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) GetName() *string {
	return s.Name
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) GetType() *string {
	return s.Type
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) GetValue() *string {
	return s.Value
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) SetName(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Name = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) SetOperation(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Operation = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) SetType(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Type = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) SetValue(v string) *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Value = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) Validate() error {
	return dara.Validate(s)
}
