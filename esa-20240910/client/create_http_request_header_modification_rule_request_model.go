// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpRequestHeaderModificationRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestHeaderModification(v []*CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) *CreateHttpRequestHeaderModificationRuleRequest
	GetRequestHeaderModification() []*CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification
	SetRule(v string) *CreateHttpRequestHeaderModificationRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateHttpRequestHeaderModificationRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateHttpRequestHeaderModificationRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateHttpRequestHeaderModificationRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateHttpRequestHeaderModificationRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateHttpRequestHeaderModificationRuleRequest
	GetSiteVersion() *int32
}

type CreateHttpRequestHeaderModificationRuleRequest struct {
	// This parameter is required.
	RequestHeaderModification []*CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty" type:"Repeated"`
	Rule                      *string                                                                    `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable                *string                                                                    `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName                  *string                                                                    `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence                  *int32                                                                     `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateHttpRequestHeaderModificationRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpRequestHeaderModificationRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpRequestHeaderModificationRuleRequest) GetRequestHeaderModification() []*CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification {
	return s.RequestHeaderModification
}

func (s *CreateHttpRequestHeaderModificationRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateHttpRequestHeaderModificationRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateHttpRequestHeaderModificationRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateHttpRequestHeaderModificationRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateHttpRequestHeaderModificationRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateHttpRequestHeaderModificationRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateHttpRequestHeaderModificationRuleRequest) SetRequestHeaderModification(v []*CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) *CreateHttpRequestHeaderModificationRuleRequest {
	s.RequestHeaderModification = v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleRequest) SetRule(v string) *CreateHttpRequestHeaderModificationRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleRequest) SetRuleEnable(v string) *CreateHttpRequestHeaderModificationRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleRequest) SetRuleName(v string) *CreateHttpRequestHeaderModificationRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleRequest) SetSequence(v int32) *CreateHttpRequestHeaderModificationRuleRequest {
	s.Sequence = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleRequest) SetSiteId(v int64) *CreateHttpRequestHeaderModificationRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleRequest) SetSiteVersion(v int32) *CreateHttpRequestHeaderModificationRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleRequest) Validate() error {
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

type CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification struct {
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) GoString() string {
	return s.String()
}

func (s *CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) GetName() *string {
	return s.Name
}

func (s *CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) GetType() *string {
	return s.Type
}

func (s *CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) GetValue() *string {
	return s.Value
}

func (s *CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) SetName(v string) *CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Name = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) SetOperation(v string) *CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Operation = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) SetType(v string) *CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Type = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) SetValue(v string) *CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification {
	s.Value = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleRequestRequestHeaderModification) Validate() error {
	return dara.Validate(s)
}
