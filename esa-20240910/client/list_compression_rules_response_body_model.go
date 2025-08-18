// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCompressionRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListCompressionRulesResponseBodyConfigs) *ListCompressionRulesResponseBody
	GetConfigs() []*ListCompressionRulesResponseBodyConfigs
	SetPageNumber(v int32) *ListCompressionRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCompressionRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCompressionRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCompressionRulesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListCompressionRulesResponseBody
	GetTotalPage() *int32
}

type ListCompressionRulesResponseBody struct {
	// List of compression rule configurations.
	Configs []*ListCompressionRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// Current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListCompressionRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCompressionRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCompressionRulesResponseBody) GetConfigs() []*ListCompressionRulesResponseBodyConfigs {
	return s.Configs
}

func (s *ListCompressionRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCompressionRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCompressionRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCompressionRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCompressionRulesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListCompressionRulesResponseBody) SetConfigs(v []*ListCompressionRulesResponseBodyConfigs) *ListCompressionRulesResponseBody {
	s.Configs = v
	return s
}

func (s *ListCompressionRulesResponseBody) SetPageNumber(v int32) *ListCompressionRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCompressionRulesResponseBody) SetPageSize(v int32) *ListCompressionRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCompressionRulesResponseBody) SetRequestId(v string) *ListCompressionRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCompressionRulesResponseBody) SetTotalCount(v int32) *ListCompressionRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCompressionRulesResponseBody) SetTotalPage(v int32) *ListCompressionRulesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListCompressionRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCompressionRulesResponseBodyConfigs struct {
	// Brotli compression. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	Brotli *string `json:"Brotli,omitempty" xml:"Brotli,omitempty"`
	// Configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type. Possible values:
	//
	// - global: Global configuration.
	//
	// - rule: Rule-based configuration.
	//
	// example:
	//
	// rule
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Gzip compression. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	Gzip *string `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
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
	// Rule switch. This parameter is not required when adding a global configuration. Possible values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
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
	// Rule execution order. The smaller the value, the higher the priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the effective version of the site configuration, defaulting to version 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Zstd compression. Value range: - on: Enable. - off: Disable.
	//
	// example:
	//
	// on
	Zstd *string `json:"Zstd,omitempty" xml:"Zstd,omitempty"`
}

func (s ListCompressionRulesResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListCompressionRulesResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListCompressionRulesResponseBodyConfigs) GetBrotli() *string {
	return s.Brotli
}

func (s *ListCompressionRulesResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListCompressionRulesResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListCompressionRulesResponseBodyConfigs) GetGzip() *string {
	return s.Gzip
}

func (s *ListCompressionRulesResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListCompressionRulesResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListCompressionRulesResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListCompressionRulesResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListCompressionRulesResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListCompressionRulesResponseBodyConfigs) GetZstd() *string {
	return s.Zstd
}

func (s *ListCompressionRulesResponseBodyConfigs) SetBrotli(v string) *ListCompressionRulesResponseBodyConfigs {
	s.Brotli = &v
	return s
}

func (s *ListCompressionRulesResponseBodyConfigs) SetConfigId(v int64) *ListCompressionRulesResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListCompressionRulesResponseBodyConfigs) SetConfigType(v string) *ListCompressionRulesResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListCompressionRulesResponseBodyConfigs) SetGzip(v string) *ListCompressionRulesResponseBodyConfigs {
	s.Gzip = &v
	return s
}

func (s *ListCompressionRulesResponseBodyConfigs) SetRule(v string) *ListCompressionRulesResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListCompressionRulesResponseBodyConfigs) SetRuleEnable(v string) *ListCompressionRulesResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListCompressionRulesResponseBodyConfigs) SetRuleName(v string) *ListCompressionRulesResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListCompressionRulesResponseBodyConfigs) SetSequence(v int32) *ListCompressionRulesResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListCompressionRulesResponseBodyConfigs) SetSiteVersion(v int32) *ListCompressionRulesResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListCompressionRulesResponseBodyConfigs) SetZstd(v string) *ListCompressionRulesResponseBodyConfigs {
	s.Zstd = &v
	return s
}

func (s *ListCompressionRulesResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
