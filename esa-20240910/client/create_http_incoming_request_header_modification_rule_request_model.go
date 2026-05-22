// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpIncomingRequestHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestHeaderModification(v []*CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) *CreateHttpIncomingRequestHeaderModificationRuleRequest
	GetRequestHeaderModification() []*CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification
	SetRule(v string) *CreateHttpIncomingRequestHeaderModificationRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateHttpIncomingRequestHeaderModificationRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateHttpIncomingRequestHeaderModificationRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateHttpIncomingRequestHeaderModificationRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateHttpIncomingRequestHeaderModificationRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateHttpIncomingRequestHeaderModificationRuleRequest
	GetSiteVersion() *int32
}

type CreateHttpIncomingRequestHeaderModificationRuleRequest struct {
	// This parameter is required.
	RequestHeaderModification []*CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty" type:"Repeated"`
	Rule                      *string                                                                            `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable                *string                                                                            `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName                  *string                                                                            `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence                  *int32                                                                             `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateHttpIncomingRequestHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpIncomingRequestHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequest) GetRequestHeaderModification() []*CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification {
	return s.RequestHeaderModification
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequest) SetRequestHeaderModification(v []*CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) *CreateHttpIncomingRequestHeaderModificationRuleRequest {
	s.RequestHeaderModification = v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequest) SetRule(v string) *CreateHttpIncomingRequestHeaderModificationRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequest) SetRuleEnable(v string) *CreateHttpIncomingRequestHeaderModificationRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequest) SetRuleName(v string) *CreateHttpIncomingRequestHeaderModificationRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequest) SetSequence(v int32) *CreateHttpIncomingRequestHeaderModificationRuleRequest {
	s.Sequence = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequest) SetSiteId(v int64) *CreateHttpIncomingRequestHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequest) SetSiteVersion(v int32) *CreateHttpIncomingRequestHeaderModificationRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequest) Validate() error {
	if s.RequestHeaderModification != nil {
		for _, item := range s.RequestHeaderModification {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification struct {
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) GoString() string {
	return s.String()
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) GetName() *string {
	return s.Name
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) GetType() *string {
	return s.Type
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) GetValue() *string {
	return s.Value
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) SetName(v string) *CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Name = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) SetOperation(v string) *CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Operation = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) SetType(v string) *CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Type = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) SetValue(v string) *CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Value = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleRequestRequestHeaderModification) Validate() error {
	return dara.Validate(s)
}
