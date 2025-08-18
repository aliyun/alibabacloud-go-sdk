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
	// Configuration ID. It can be obtained by calling the [ListRewriteUrlRules](https://help.aliyun.com/document_detail/2867480.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3528160969****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The query string after rewriting.
	//
	// example:
	//
	// example=123
	QueryString *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	// Query string rewrite type. Value range:
	//
	// - static: Static mode.
	//
	// - dynamic: Dynamic mode.
	//
	// example:
	//
	// static
	RewriteQueryStringType *string `json:"RewriteQueryStringType,omitempty" xml:"RewriteQueryStringType,omitempty"`
	// URI rewrite type. Value range:
	//
	// - static: Static mode.
	//
	// - dynamic: Dynamic mode.
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
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, for example: (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// Rule name. This parameter is not required when adding a global configuration.
	//
	// example:
	//
	// example=123
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The target URI after rewriting.
	//
	// example:
	//
	// /image/example.jpg
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
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
