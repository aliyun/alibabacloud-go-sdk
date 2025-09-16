// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpIncomingResponseHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateHttpIncomingResponseHeaderModificationRuleRequest
	GetConfigId() *int64
	SetResponseHeaderModification(v []*UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) *UpdateHttpIncomingResponseHeaderModificationRuleRequest
	GetResponseHeaderModification() []*UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification
	SetRule(v string) *UpdateHttpIncomingResponseHeaderModificationRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateHttpIncomingResponseHeaderModificationRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateHttpIncomingResponseHeaderModificationRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateHttpIncomingResponseHeaderModificationRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateHttpIncomingResponseHeaderModificationRuleRequest
	GetSiteId() *int64
}

type UpdateHttpIncomingResponseHeaderModificationRuleRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 352816096987136
	ConfigId                   *int64                                                                               `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ResponseHeaderModification []*UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification `json:"ResponseHeaderModification,omitempty" xml:"ResponseHeaderModification,omitempty" type:"Repeated"`
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
	// 498607398028944
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateHttpIncomingResponseHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpIncomingResponseHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequest) GetResponseHeaderModification() []*UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification {
	return s.ResponseHeaderModification
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequest) SetConfigId(v int64) *UpdateHttpIncomingResponseHeaderModificationRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequest) SetResponseHeaderModification(v []*UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) *UpdateHttpIncomingResponseHeaderModificationRuleRequest {
	s.ResponseHeaderModification = v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequest) SetRule(v string) *UpdateHttpIncomingResponseHeaderModificationRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequest) SetRuleEnable(v string) *UpdateHttpIncomingResponseHeaderModificationRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequest) SetRuleName(v string) *UpdateHttpIncomingResponseHeaderModificationRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequest) SetSequence(v int32) *UpdateHttpIncomingResponseHeaderModificationRuleRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequest) SetSiteId(v int64) *UpdateHttpIncomingResponseHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification struct {
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

func (s UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) GoString() string {
	return s.String()
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) GetName() *string {
	return s.Name
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) GetType() *string {
	return s.Type
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) GetValue() *string {
	return s.Value
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) SetName(v string) *UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Name = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) SetOperation(v string) *UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Operation = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) SetType(v string) *UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Type = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) SetValue(v string) *UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Value = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) Validate() error {
	return dara.Validate(s)
}
