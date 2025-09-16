// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpIncomingRequestHeaderModificationRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRequestHeaderModificationShrink(v string) *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest
	GetRequestHeaderModificationShrink() *string
	SetRule(v string) *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest
	GetSiteVersion() *int32
}

type CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest struct {
	// This parameter is required.
	RequestHeaderModificationShrink *string `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty"`
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
	// 478016908379824
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GetRequestHeaderModificationShrink() *string {
	return s.RequestHeaderModificationShrink
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) SetRequestHeaderModificationShrink(v string) *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest {
	s.RequestHeaderModificationShrink = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) SetRule(v string) *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest {
	s.Rule = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) SetRuleEnable(v string) *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) SetRuleName(v string) *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) SetSequence(v int32) *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest {
	s.Sequence = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) SetSiteId(v int64) *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) SetSiteVersion(v int32) *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateHttpIncomingRequestHeaderModificationRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
