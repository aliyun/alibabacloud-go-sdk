// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest
	GetConfigId() *int64
	SetRequestHeaderModificationShrink(v string) *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest
	GetRequestHeaderModificationShrink() *string
	SetRule(v string) *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest
	GetSiteId() *int64
}

type UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 419717024278528
	ConfigId                        *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
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
	// 568181195163328
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GetRequestHeaderModificationShrink() *string {
	return s.RequestHeaderModificationShrink
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) SetConfigId(v int64) *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) SetRequestHeaderModificationShrink(v string) *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest {
	s.RequestHeaderModificationShrink = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) SetRule(v string) *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest {
	s.Rule = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) SetRuleEnable(v string) *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) SetRuleName(v string) *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) SetSequence(v int32) *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) SetSiteId(v int64) *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateHttpIncomingRequestHeaderModificationRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
