// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpRequestHeaderModificationRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestHeaderModificationShrink(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest
	GetRequestHeaderModificationShrink() *string
	SetRule(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateHttpRequestHeaderModificationRuleShrinkRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateHttpRequestHeaderModificationRuleShrinkRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateHttpRequestHeaderModificationRuleShrinkRequest
	GetSiteVersion() *int32
}

type CreateHttpRequestHeaderModificationRuleShrinkRequest struct {
	// This parameter is required.
	RequestHeaderModificationShrink *string `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty"`
	Rule                            *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable                      *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName                        *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence                        *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateHttpRequestHeaderModificationRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpRequestHeaderModificationRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) GetRequestHeaderModificationShrink() *string {
	return s.RequestHeaderModificationShrink
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) SetRequestHeaderModificationShrink(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest {
	s.RequestHeaderModificationShrink = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) SetRule(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest {
	s.Rule = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) SetRuleEnable(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) SetRuleName(v string) *CreateHttpRequestHeaderModificationRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) SetSequence(v int32) *CreateHttpRequestHeaderModificationRuleShrinkRequest {
	s.Sequence = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) SetSiteId(v int64) *CreateHttpRequestHeaderModificationRuleShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) SetSiteVersion(v int32) *CreateHttpRequestHeaderModificationRuleShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateHttpRequestHeaderModificationRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
