// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRewriteUrlRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetRewriteUrlRuleResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetRewriteUrlRuleResponseBody
	GetConfigType() *string
	SetQueryString(v string) *GetRewriteUrlRuleResponseBody
	GetQueryString() *string
	SetRequestId(v string) *GetRewriteUrlRuleResponseBody
	GetRequestId() *string
	SetRewriteQueryStringType(v string) *GetRewriteUrlRuleResponseBody
	GetRewriteQueryStringType() *string
	SetRewriteUriType(v string) *GetRewriteUrlRuleResponseBody
	GetRewriteUriType() *string
	SetRule(v string) *GetRewriteUrlRuleResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetRewriteUrlRuleResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetRewriteUrlRuleResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetRewriteUrlRuleResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetRewriteUrlRuleResponseBody
	GetSiteVersion() *int32
	SetUri(v string) *GetRewriteUrlRuleResponseBody
	GetUri() *string
}

type GetRewriteUrlRuleResponseBody struct {
	// Configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type. Possible values:
	//
	// - global: Global configuration;
	//
	// - rule: Rule-based configuration;
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The rewritten query string.
	//
	// example:
	//
	// example=123
	QueryString *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Query string rewrite type. Possible values:
	//
	// - static: Static mode.
	//
	// - dynamic: Dynamic mode.
	//
	// example:
	//
	// static
	RewriteQueryStringType *string `json:"RewriteQueryStringType,omitempty" xml:"RewriteQueryStringType,omitempty"`
	// URI rewrite type. Possible values:
	//
	// - static: Static mode.
	//
	// - dynamic: Dynamic mode.
	//
	// example:
	//
	// static
	RewriteUriType *string `json:"RewriteUriType,omitempty" xml:"RewriteUriType,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter does not need to be set when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, e.g., (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter does not need to be set when adding a global configuration. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// Rule name. This parameter does not need to be set when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Rule execution order. The smaller the value, the higher the priority for execution.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the effective version of the site, defaulting to version 0.
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

func (s GetRewriteUrlRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRewriteUrlRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRewriteUrlRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetRewriteUrlRuleResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetRewriteUrlRuleResponseBody) GetQueryString() *string {
	return s.QueryString
}

func (s *GetRewriteUrlRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRewriteUrlRuleResponseBody) GetRewriteQueryStringType() *string {
	return s.RewriteQueryStringType
}

func (s *GetRewriteUrlRuleResponseBody) GetRewriteUriType() *string {
	return s.RewriteUriType
}

func (s *GetRewriteUrlRuleResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetRewriteUrlRuleResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetRewriteUrlRuleResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetRewriteUrlRuleResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetRewriteUrlRuleResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetRewriteUrlRuleResponseBody) GetUri() *string {
	return s.Uri
}

func (s *GetRewriteUrlRuleResponseBody) SetConfigId(v int64) *GetRewriteUrlRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetRewriteUrlRuleResponseBody) SetConfigType(v string) *GetRewriteUrlRuleResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetRewriteUrlRuleResponseBody) SetQueryString(v string) *GetRewriteUrlRuleResponseBody {
	s.QueryString = &v
	return s
}

func (s *GetRewriteUrlRuleResponseBody) SetRequestId(v string) *GetRewriteUrlRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRewriteUrlRuleResponseBody) SetRewriteQueryStringType(v string) *GetRewriteUrlRuleResponseBody {
	s.RewriteQueryStringType = &v
	return s
}

func (s *GetRewriteUrlRuleResponseBody) SetRewriteUriType(v string) *GetRewriteUrlRuleResponseBody {
	s.RewriteUriType = &v
	return s
}

func (s *GetRewriteUrlRuleResponseBody) SetRule(v string) *GetRewriteUrlRuleResponseBody {
	s.Rule = &v
	return s
}

func (s *GetRewriteUrlRuleResponseBody) SetRuleEnable(v string) *GetRewriteUrlRuleResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetRewriteUrlRuleResponseBody) SetRuleName(v string) *GetRewriteUrlRuleResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetRewriteUrlRuleResponseBody) SetSequence(v int32) *GetRewriteUrlRuleResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetRewriteUrlRuleResponseBody) SetSiteVersion(v int32) *GetRewriteUrlRuleResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetRewriteUrlRuleResponseBody) SetUri(v string) *GetRewriteUrlRuleResponseBody {
	s.Uri = &v
	return s
}

func (s *GetRewriteUrlRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
