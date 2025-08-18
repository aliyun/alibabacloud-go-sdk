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
	// The query string after rewriting.
	//
	// example:
	//
	// example=123
	QueryString *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	// Query string rewrite type. Value range:
	//
	// - static: static mode.
	//
	// - dynamic: dynamic mode.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// static
	RewriteQueryStringType *string `json:"RewriteQueryStringType,omitempty" xml:"RewriteQueryStringType,omitempty"`
	// URI rewrite type. Value range:
	//
	// - static: static mode.
	//
	// - dynamic: dynamic mode.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// static
	RewriteUriType *string `json:"RewriteUriType,omitempty" xml:"RewriteUriType,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: set the value to true
	//
	// - Match specific requests: set the value to a custom expression, for example: (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. Value range:
	//
	// - on: enable.
	//
	// - off: disable.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// Rule name. This parameter is not required when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the effective version of the site configuration, defaulting to version 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// The target URI after rewriting.
	//
	// example:
	//
	// /image/example.jpg
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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
