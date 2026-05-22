// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpRequestHeaderModificationRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateHttpRequestHeaderModificationRuleShrinkRequest
	GetConfigId() *int64
	SetRequestHeaderModificationShrink(v string) *UpdateHttpRequestHeaderModificationRuleShrinkRequest
	GetRequestHeaderModificationShrink() *string
	SetRule(v string) *UpdateHttpRequestHeaderModificationRuleShrinkRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateHttpRequestHeaderModificationRuleShrinkRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateHttpRequestHeaderModificationRuleShrinkRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateHttpRequestHeaderModificationRuleShrinkRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateHttpRequestHeaderModificationRuleShrinkRequest
	GetSiteId() *int64
}

type UpdateHttpRequestHeaderModificationRuleShrinkRequest struct {
	// This parameter is required.
	ConfigId                        *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	RequestHeaderModificationShrink *string `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty"`
	Rule                            *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable                      *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName                        *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence                        *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateHttpRequestHeaderModificationRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpRequestHeaderModificationRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpRequestHeaderModificationRuleShrinkRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateHttpRequestHeaderModificationRuleShrinkRequest) GetRequestHeaderModificationShrink() *string {
	return s.RequestHeaderModificationShrink
}

func (s *UpdateHttpRequestHeaderModificationRuleShrinkRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateHttpRequestHeaderModificationRuleShrinkRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateHttpRequestHeaderModificationRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateHttpRequestHeaderModificationRuleShrinkRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateHttpRequestHeaderModificationRuleShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateHttpRequestHeaderModificationRuleShrinkRequest) SetConfigId(v int64) *UpdateHttpRequestHeaderModificationRuleShrinkRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleShrinkRequest) SetRequestHeaderModificationShrink(v string) *UpdateHttpRequestHeaderModificationRuleShrinkRequest {
	s.RequestHeaderModificationShrink = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleShrinkRequest) SetRule(v string) *UpdateHttpRequestHeaderModificationRuleShrinkRequest {
	s.Rule = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleShrinkRequest) SetRuleEnable(v string) *UpdateHttpRequestHeaderModificationRuleShrinkRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleShrinkRequest) SetRuleName(v string) *UpdateHttpRequestHeaderModificationRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleShrinkRequest) SetSequence(v int32) *UpdateHttpRequestHeaderModificationRuleShrinkRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleShrinkRequest) SetSiteId(v int64) *UpdateHttpRequestHeaderModificationRuleShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateHttpRequestHeaderModificationRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
