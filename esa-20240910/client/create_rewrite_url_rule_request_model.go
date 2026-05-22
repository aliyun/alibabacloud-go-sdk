// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRewriteUrlRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQueryString(v string) *CreateRewriteUrlRuleRequest
	GetQueryString() *string
	SetRewriteQueryStringType(v string) *CreateRewriteUrlRuleRequest
	GetRewriteQueryStringType() *string
	SetRewriteUriType(v string) *CreateRewriteUrlRuleRequest
	GetRewriteUriType() *string
	SetRule(v string) *CreateRewriteUrlRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateRewriteUrlRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateRewriteUrlRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateRewriteUrlRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateRewriteUrlRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateRewriteUrlRuleRequest
	GetSiteVersion() *int32
	SetUri(v string) *CreateRewriteUrlRuleRequest
	GetUri() *string
}

type CreateRewriteUrlRuleRequest struct {
	QueryString *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	// if can be null:
	// false
	RewriteQueryStringType *string `json:"RewriteQueryStringType,omitempty" xml:"RewriteQueryStringType,omitempty"`
	// if can be null:
	// false
	RewriteUriType *string `json:"RewriteUriType,omitempty" xml:"RewriteUriType,omitempty"`
	Rule           *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable     *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName       *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence       *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	SiteId      *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteVersion *int32  `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	Uri         *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s CreateRewriteUrlRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRewriteUrlRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateRewriteUrlRuleRequest) GetQueryString() *string {
	return s.QueryString
}

func (s *CreateRewriteUrlRuleRequest) GetRewriteQueryStringType() *string {
	return s.RewriteQueryStringType
}

func (s *CreateRewriteUrlRuleRequest) GetRewriteUriType() *string {
	return s.RewriteUriType
}

func (s *CreateRewriteUrlRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateRewriteUrlRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateRewriteUrlRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateRewriteUrlRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateRewriteUrlRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateRewriteUrlRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateRewriteUrlRuleRequest) GetUri() *string {
	return s.Uri
}

func (s *CreateRewriteUrlRuleRequest) SetQueryString(v string) *CreateRewriteUrlRuleRequest {
	s.QueryString = &v
	return s
}

func (s *CreateRewriteUrlRuleRequest) SetRewriteQueryStringType(v string) *CreateRewriteUrlRuleRequest {
	s.RewriteQueryStringType = &v
	return s
}

func (s *CreateRewriteUrlRuleRequest) SetRewriteUriType(v string) *CreateRewriteUrlRuleRequest {
	s.RewriteUriType = &v
	return s
}

func (s *CreateRewriteUrlRuleRequest) SetRule(v string) *CreateRewriteUrlRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateRewriteUrlRuleRequest) SetRuleEnable(v string) *CreateRewriteUrlRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateRewriteUrlRuleRequest) SetRuleName(v string) *CreateRewriteUrlRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateRewriteUrlRuleRequest) SetSequence(v int32) *CreateRewriteUrlRuleRequest {
	s.Sequence = &v
	return s
}

func (s *CreateRewriteUrlRuleRequest) SetSiteId(v int64) *CreateRewriteUrlRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateRewriteUrlRuleRequest) SetSiteVersion(v int32) *CreateRewriteUrlRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateRewriteUrlRuleRequest) SetUri(v string) *CreateRewriteUrlRuleRequest {
	s.Uri = &v
	return s
}

func (s *CreateRewriteUrlRuleRequest) Validate() error {
	return dara.Validate(s)
}
