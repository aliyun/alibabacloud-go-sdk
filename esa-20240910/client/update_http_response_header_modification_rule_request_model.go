// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpResponseHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateHttpResponseHeaderModificationRuleRequest
	GetConfigId() *int64
	SetResponseHeaderModification(v []*UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) *UpdateHttpResponseHeaderModificationRuleRequest
	GetResponseHeaderModification() []*UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification
	SetRule(v string) *UpdateHttpResponseHeaderModificationRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateHttpResponseHeaderModificationRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateHttpResponseHeaderModificationRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateHttpResponseHeaderModificationRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateHttpResponseHeaderModificationRuleRequest
	GetSiteId() *int64
}

type UpdateHttpResponseHeaderModificationRuleRequest struct {
	// This parameter is required.
	ConfigId                   *int64                                                                       `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ResponseHeaderModification []*UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification `json:"ResponseHeaderModification,omitempty" xml:"ResponseHeaderModification,omitempty" type:"Repeated"`
	Rule                       *string                                                                      `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable                 *string                                                                      `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName                   *string                                                                      `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence                   *int32                                                                       `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateHttpResponseHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpResponseHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpResponseHeaderModificationRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateHttpResponseHeaderModificationRuleRequest) GetResponseHeaderModification() []*UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification {
	return s.ResponseHeaderModification
}

func (s *UpdateHttpResponseHeaderModificationRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateHttpResponseHeaderModificationRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateHttpResponseHeaderModificationRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateHttpResponseHeaderModificationRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateHttpResponseHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateHttpResponseHeaderModificationRuleRequest) SetConfigId(v int64) *UpdateHttpResponseHeaderModificationRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleRequest) SetResponseHeaderModification(v []*UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) *UpdateHttpResponseHeaderModificationRuleRequest {
	s.ResponseHeaderModification = v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleRequest) SetRule(v string) *UpdateHttpResponseHeaderModificationRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleRequest) SetRuleEnable(v string) *UpdateHttpResponseHeaderModificationRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleRequest) SetRuleName(v string) *UpdateHttpResponseHeaderModificationRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleRequest) SetSequence(v int32) *UpdateHttpResponseHeaderModificationRuleRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleRequest) SetSiteId(v int64) *UpdateHttpResponseHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleRequest) Validate() error {
	if s.ResponseHeaderModification != nil {
		for _, item := range s.ResponseHeaderModification {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification struct {
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) GoString() string {
	return s.String()
}

func (s *UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) GetName() *string {
	return s.Name
}

func (s *UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) GetType() *string {
	return s.Type
}

func (s *UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) GetValue() *string {
	return s.Value
}

func (s *UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) SetName(v string) *UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Name = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) SetOperation(v string) *UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Operation = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) SetType(v string) *UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Type = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) SetValue(v string) *UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Value = &v
	return s
}

func (s *UpdateHttpResponseHeaderModificationRuleRequestResponseHeaderModification) Validate() error {
	return dara.Validate(s)
}
