// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpIncomingResponseHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResponseHeaderModification(v []*CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) *CreateHttpIncomingResponseHeaderModificationRuleRequest
	GetResponseHeaderModification() []*CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification
	SetRule(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateHttpIncomingResponseHeaderModificationRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateHttpIncomingResponseHeaderModificationRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateHttpIncomingResponseHeaderModificationRuleRequest
	GetSiteVersion() *int32
}

type CreateHttpIncomingResponseHeaderModificationRuleRequest struct {
	// This parameter is required.
	ResponseHeaderModification []*CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification `json:"ResponseHeaderModification,omitempty" xml:"ResponseHeaderModification,omitempty" type:"Repeated"`
	Rule                       *string                                                                              `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable                 *string                                                                              `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName                   *string                                                                              `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence                   *int32                                                                               `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateHttpIncomingResponseHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpIncomingResponseHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) GetResponseHeaderModification() []*CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification {
	return s.ResponseHeaderModification
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) SetResponseHeaderModification(v []*CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) *CreateHttpIncomingResponseHeaderModificationRuleRequest {
	s.ResponseHeaderModification = v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) SetRule(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) SetRuleEnable(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) SetRuleName(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) SetSequence(v int32) *CreateHttpIncomingResponseHeaderModificationRuleRequest {
	s.Sequence = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) SetSiteId(v int64) *CreateHttpIncomingResponseHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) SetSiteVersion(v int32) *CreateHttpIncomingResponseHeaderModificationRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequest) Validate() error {
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

type CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification struct {
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) GoString() string {
	return s.String()
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) GetName() *string {
	return s.Name
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) GetType() *string {
	return s.Type
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) GetValue() *string {
	return s.Value
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) SetName(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Name = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) SetOperation(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Operation = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) SetType(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Type = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) SetValue(v string) *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification {
	s.Value = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleRequestResponseHeaderModification) Validate() error {
	return dara.Validate(s)
}
