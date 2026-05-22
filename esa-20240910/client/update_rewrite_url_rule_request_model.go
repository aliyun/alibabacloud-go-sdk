// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRewriteUrlRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateRewriteUrlRuleRequest
	GetConfigId() *int64
	SetQueryString(v string) *UpdateRewriteUrlRuleRequest
	GetQueryString() *string
	SetRewriteQueryStringType(v string) *UpdateRewriteUrlRuleRequest
	GetRewriteQueryStringType() *string
	SetRewriteUriType(v string) *UpdateRewriteUrlRuleRequest
	GetRewriteUriType() *string
	SetRule(v string) *UpdateRewriteUrlRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateRewriteUrlRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateRewriteUrlRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateRewriteUrlRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateRewriteUrlRuleRequest
	GetSiteId() *int64
	SetUri(v string) *UpdateRewriteUrlRuleRequest
	GetUri() *string
}

type UpdateRewriteUrlRuleRequest struct {
	// This parameter is required.
	ConfigId               *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	QueryString            *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	RewriteQueryStringType *string `json:"RewriteQueryStringType,omitempty" xml:"RewriteQueryStringType,omitempty"`
	// if can be null:
	// false
	RewriteUriType *string `json:"RewriteUriType,omitempty" xml:"RewriteUriType,omitempty"`
	Rule           *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable     *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName       *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence       *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	SiteId *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	Uri    *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s UpdateRewriteUrlRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRewriteUrlRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRewriteUrlRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateRewriteUrlRuleRequest) GetQueryString() *string {
	return s.QueryString
}

func (s *UpdateRewriteUrlRuleRequest) GetRewriteQueryStringType() *string {
	return s.RewriteQueryStringType
}

func (s *UpdateRewriteUrlRuleRequest) GetRewriteUriType() *string {
	return s.RewriteUriType
}

func (s *UpdateRewriteUrlRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateRewriteUrlRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateRewriteUrlRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateRewriteUrlRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateRewriteUrlRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateRewriteUrlRuleRequest) GetUri() *string {
	return s.Uri
}

func (s *UpdateRewriteUrlRuleRequest) SetConfigId(v int64) *UpdateRewriteUrlRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateRewriteUrlRuleRequest) SetQueryString(v string) *UpdateRewriteUrlRuleRequest {
	s.QueryString = &v
	return s
}

func (s *UpdateRewriteUrlRuleRequest) SetRewriteQueryStringType(v string) *UpdateRewriteUrlRuleRequest {
	s.RewriteQueryStringType = &v
	return s
}

func (s *UpdateRewriteUrlRuleRequest) SetRewriteUriType(v string) *UpdateRewriteUrlRuleRequest {
	s.RewriteUriType = &v
	return s
}

func (s *UpdateRewriteUrlRuleRequest) SetRule(v string) *UpdateRewriteUrlRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateRewriteUrlRuleRequest) SetRuleEnable(v string) *UpdateRewriteUrlRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateRewriteUrlRuleRequest) SetRuleName(v string) *UpdateRewriteUrlRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateRewriteUrlRuleRequest) SetSequence(v int32) *UpdateRewriteUrlRuleRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateRewriteUrlRuleRequest) SetSiteId(v int64) *UpdateRewriteUrlRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateRewriteUrlRuleRequest) SetUri(v string) *UpdateRewriteUrlRuleRequest {
	s.Uri = &v
	return s
}

func (s *UpdateRewriteUrlRuleRequest) Validate() error {
	return dara.Validate(s)
}
