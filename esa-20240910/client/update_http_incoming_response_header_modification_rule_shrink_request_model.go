// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest
	GetConfigId() *int64
	SetResponseHeaderModificationShrink(v string) *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest
	GetResponseHeaderModificationShrink() *string
	SetRule(v string) *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest
	GetSiteId() *int64
}

type UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 352816096987136
	ConfigId                         *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
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
	// 498607398028944
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GetResponseHeaderModificationShrink() *string {
	return s.ResponseHeaderModificationShrink
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) SetConfigId(v int64) *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) SetResponseHeaderModificationShrink(v string) *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest {
	s.ResponseHeaderModificationShrink = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) SetRule(v string) *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest {
	s.Rule = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) SetRuleEnable(v string) *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) SetRuleName(v string) *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) SetSequence(v int32) *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) SetSiteId(v int64) *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateHttpIncomingResponseHeaderModificationRuleShrinkRequest) Validate() error {
	return dara.Validate(s)
}
