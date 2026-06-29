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
	// The list of compression rule configurations.
	Configs []*ListCompressionRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The current page number, which is the same as the PageNumber request parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
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
	if s.Configs != nil {
		for _, item := range s.Configs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCompressionRulesResponseBodyConfigs struct {
	// Brotli compression. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Brotli *string `json:"Brotli,omitempty" xml:"Brotli,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. Valid values:
	//
	// - global: global configuration.
	//
	// - rule: rule configuration.
	//
	// example:
	//
	// rule
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Gzip compression. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Gzip *string `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
	// The rule content, which uses a conditional expression to match user requests. You do not need to set this parameter when adding a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specified requests: Set the value to a custom expression, for example, (http.host eq \\"video.example.com\\").
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. You do not need to set this parameter when adding a global configuration. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. You do not need to set this parameter when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule execution order. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. Default value: 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Zstd compression. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
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
