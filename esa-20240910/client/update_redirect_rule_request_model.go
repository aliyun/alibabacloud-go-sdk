// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRedirectRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateRedirectRuleRequest
	GetConfigId() *int64
	SetReserveQueryString(v string) *UpdateRedirectRuleRequest
	GetReserveQueryString() *string
	SetRule(v string) *UpdateRedirectRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateRedirectRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateRedirectRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateRedirectRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateRedirectRuleRequest
	GetSiteId() *int64
	SetStatusCode(v string) *UpdateRedirectRuleRequest
	GetStatusCode() *string
	SetTargetUrl(v string) *UpdateRedirectRuleRequest
	GetTargetUrl() *string
	SetType(v string) *UpdateRedirectRuleRequest
	GetType() *string
}

type UpdateRedirectRuleRequest struct {
	// This parameter is required.
	ConfigId           *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ReserveQueryString *string `json:"ReserveQueryString,omitempty" xml:"ReserveQueryString,omitempty"`
	Rule               *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable         *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName           *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence           *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	SiteId     *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	TargetUrl  *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRedirectRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRedirectRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRedirectRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateRedirectRuleRequest) GetReserveQueryString() *string {
	return s.ReserveQueryString
}

func (s *UpdateRedirectRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateRedirectRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateRedirectRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateRedirectRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateRedirectRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateRedirectRuleRequest) GetStatusCode() *string {
	return s.StatusCode
}

func (s *UpdateRedirectRuleRequest) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *UpdateRedirectRuleRequest) GetType() *string {
	return s.Type
}

func (s *UpdateRedirectRuleRequest) SetConfigId(v int64) *UpdateRedirectRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetReserveQueryString(v string) *UpdateRedirectRuleRequest {
	s.ReserveQueryString = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetRule(v string) *UpdateRedirectRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetRuleEnable(v string) *UpdateRedirectRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetRuleName(v string) *UpdateRedirectRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetSequence(v int32) *UpdateRedirectRuleRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetSiteId(v int64) *UpdateRedirectRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetStatusCode(v string) *UpdateRedirectRuleRequest {
	s.StatusCode = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetTargetUrl(v string) *UpdateRedirectRuleRequest {
	s.TargetUrl = &v
	return s
}

func (s *UpdateRedirectRuleRequest) SetType(v string) *UpdateRedirectRuleRequest {
	s.Type = &v
	return s
}

func (s *UpdateRedirectRuleRequest) Validate() error {
	return dara.Validate(s)
}
