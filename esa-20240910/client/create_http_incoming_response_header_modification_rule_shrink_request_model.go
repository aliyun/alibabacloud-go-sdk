// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpIncomingResponseHeaderModificationRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResponseHeaderModificationShrink(v string) *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest
	GetResponseHeaderModificationShrink() *string
	SetRule(v string) *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest
	GetSiteVersion() *int32
}

type CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest struct {
	// This parameter is required.
	ResponseHeaderModificationShrink *string `json:"ResponseHeaderModification,omitempty" xml:"ResponseHeaderModification,omitempty"`
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
	// 608665779308176
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GetResponseHeaderModificationShrink() *string {
	return s.ResponseHeaderModificationShrink
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) SetResponseHeaderModificationShrink(v string) *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest {
	s.ResponseHeaderModificationShrink = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) SetRule(v string) *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest {
	s.Rule = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) SetRuleEnable(v string) *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) SetRuleName(v string) *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) SetSequence(v int32) *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest {
	s.Sequence = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) SetSiteId(v int64) *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) SetSiteVersion(v int32) *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateHttpIncomingResponseHeaderModificationRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
