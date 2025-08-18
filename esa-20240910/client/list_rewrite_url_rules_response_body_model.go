// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRewriteUrlRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListRewriteUrlRulesResponseBodyConfigs) *ListRewriteUrlRulesResponseBody
	GetConfigs() []*ListRewriteUrlRulesResponseBodyConfigs
	SetPageNumber(v int32) *ListRewriteUrlRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRewriteUrlRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRewriteUrlRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListRewriteUrlRulesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListRewriteUrlRulesResponseBody
	GetTotalPage() *int32
}

type ListRewriteUrlRulesResponseBody struct {
	// List of rewrite URL configurations.
	Configs []*ListRewriteUrlRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The size of the page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-280B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of items.
	//
	// example:
	//
	// 8
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListRewriteUrlRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRewriteUrlRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRewriteUrlRulesResponseBody) GetConfigs() []*ListRewriteUrlRulesResponseBodyConfigs {
	return s.Configs
}

func (s *ListRewriteUrlRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRewriteUrlRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRewriteUrlRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRewriteUrlRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRewriteUrlRulesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListRewriteUrlRulesResponseBody) SetConfigs(v []*ListRewriteUrlRulesResponseBodyConfigs) *ListRewriteUrlRulesResponseBody {
	s.Configs = v
	return s
}

func (s *ListRewriteUrlRulesResponseBody) SetPageNumber(v int32) *ListRewriteUrlRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBody) SetPageSize(v int32) *ListRewriteUrlRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBody) SetRequestId(v string) *ListRewriteUrlRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBody) SetTotalCount(v int32) *ListRewriteUrlRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBody) SetTotalPage(v int32) *ListRewriteUrlRulesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRewriteUrlRulesResponseBodyConfigs struct {
	// Configuration ID.
	//
	// example:
	//
	// 39538644977****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type. Value range:
	//
	// - global: Global configuration;
	//
	// - rule: Rule configuration;
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
	// example:
	//
	// static
	RewriteUriType *string `json:"RewriteUriType,omitempty" xml:"RewriteUriType,omitempty"`
	// Rule content, using conditional expressions to match user requests. Not required when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, e.g., (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. Not required when adding a global configuration. Value range:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// Rule name. Not required when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// Rule execution order. The smaller the value, the higher the priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the effective version of the configuration, defaulting to version 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Target URI after rewriting.
	//
	// example:
	//
	// /image.example.com/index.html
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ListRewriteUrlRulesResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListRewriteUrlRulesResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) GetQueryString() *string {
	return s.QueryString
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) GetRewriteQueryStringType() *string {
	return s.RewriteQueryStringType
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) GetRewriteUriType() *string {
	return s.RewriteUriType
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) GetUri() *string {
	return s.Uri
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) SetConfigId(v int64) *ListRewriteUrlRulesResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) SetConfigType(v string) *ListRewriteUrlRulesResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) SetQueryString(v string) *ListRewriteUrlRulesResponseBodyConfigs {
	s.QueryString = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) SetRewriteQueryStringType(v string) *ListRewriteUrlRulesResponseBodyConfigs {
	s.RewriteQueryStringType = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) SetRewriteUriType(v string) *ListRewriteUrlRulesResponseBodyConfigs {
	s.RewriteUriType = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) SetRule(v string) *ListRewriteUrlRulesResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) SetRuleEnable(v string) *ListRewriteUrlRulesResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) SetRuleName(v string) *ListRewriteUrlRulesResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) SetSequence(v int32) *ListRewriteUrlRulesResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) SetSiteVersion(v int32) *ListRewriteUrlRulesResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) SetUri(v string) *ListRewriteUrlRulesResponseBodyConfigs {
	s.Uri = &v
	return s
}

func (s *ListRewriteUrlRulesResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
