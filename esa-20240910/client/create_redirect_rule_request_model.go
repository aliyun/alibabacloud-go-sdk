// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRedirectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReserveQueryString(v string) *CreateRedirectRuleRequest
	GetReserveQueryString() *string
	SetRule(v string) *CreateRedirectRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateRedirectRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateRedirectRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateRedirectRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateRedirectRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateRedirectRuleRequest
	GetSiteVersion() *int32
	SetStatusCode(v string) *CreateRedirectRuleRequest
	GetStatusCode() *string
	SetTargetUrl(v string) *CreateRedirectRuleRequest
	GetTargetUrl() *string
	SetType(v string) *CreateRedirectRuleRequest
	GetType() *string
}

type CreateRedirectRuleRequest struct {
	// This parameter is required.
	ReserveQueryString *string `json:"ReserveQueryString,omitempty" xml:"ReserveQueryString,omitempty"`
	Rule               *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable         *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName           *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence           *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// This parameter is required.
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// This parameter is required.
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRedirectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRedirectRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRedirectRuleRequest) GetReserveQueryString() *string {
	return s.ReserveQueryString
}

func (s *CreateRedirectRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateRedirectRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateRedirectRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateRedirectRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateRedirectRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateRedirectRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateRedirectRuleRequest) GetStatusCode() *string {
	return s.StatusCode
}

func (s *CreateRedirectRuleRequest) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *CreateRedirectRuleRequest) GetType() *string {
	return s.Type
}

func (s *CreateRedirectRuleRequest) SetReserveQueryString(v string) *CreateRedirectRuleRequest {
	s.ReserveQueryString = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetRule(v string) *CreateRedirectRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetRuleEnable(v string) *CreateRedirectRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetRuleName(v string) *CreateRedirectRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetSequence(v int32) *CreateRedirectRuleRequest {
	s.Sequence = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetSiteId(v int64) *CreateRedirectRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetSiteVersion(v int32) *CreateRedirectRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetStatusCode(v string) *CreateRedirectRuleRequest {
	s.StatusCode = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetTargetUrl(v string) *CreateRedirectRuleRequest {
	s.TargetUrl = &v
	return s
}

func (s *CreateRedirectRuleRequest) SetType(v string) *CreateRedirectRuleRequest {
	s.Type = &v
	return s
}

func (s *CreateRedirectRuleRequest) Validate() error {
	return dara.Validate(s)
}
