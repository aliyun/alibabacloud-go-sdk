// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompressionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBrotli(v string) *GetCompressionRuleResponseBody
	GetBrotli() *string
	SetConfigId(v int64) *GetCompressionRuleResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetCompressionRuleResponseBody
	GetConfigType() *string
	SetGzip(v string) *GetCompressionRuleResponseBody
	GetGzip() *string
	SetRequestId(v string) *GetCompressionRuleResponseBody
	GetRequestId() *string
	SetRule(v string) *GetCompressionRuleResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetCompressionRuleResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetCompressionRuleResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetCompressionRuleResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetCompressionRuleResponseBody
	GetSiteVersion() *int32
	SetZstd(v string) *GetCompressionRuleResponseBody
	GetZstd() *string
}

type GetCompressionRuleResponseBody struct {
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
	// The ID of the configuration.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The type of the configuration. Valid values:
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
	// The ID of the request.
	//
	// example:
	//
	// 186C6DF2-D96A-5102-B04E-FB92C16C9867
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The rule content, which uses a conditional expression to match user requests. You do not need to set this parameter when adding a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: set the value to true.
	//
	// - Match specific requests: set the value to a custom expression, such as (http.host eq "video.example.com").
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
	// The name of the rule. You do not need to set this parameter when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution sequence of the rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The version number of the site configuration. For sites with configuration version management enabled, you can use this parameter to specify the site version for which the configuration takes effect. The default value is version 0.
	//
	// example:
	//
	// 0
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

func (s GetCompressionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCompressionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetCompressionRuleResponseBody) GetBrotli() *string {
	return s.Brotli
}

func (s *GetCompressionRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetCompressionRuleResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetCompressionRuleResponseBody) GetGzip() *string {
	return s.Gzip
}

func (s *GetCompressionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCompressionRuleResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetCompressionRuleResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetCompressionRuleResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetCompressionRuleResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetCompressionRuleResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetCompressionRuleResponseBody) GetZstd() *string {
	return s.Zstd
}

func (s *GetCompressionRuleResponseBody) SetBrotli(v string) *GetCompressionRuleResponseBody {
	s.Brotli = &v
	return s
}

func (s *GetCompressionRuleResponseBody) SetConfigId(v int64) *GetCompressionRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetCompressionRuleResponseBody) SetConfigType(v string) *GetCompressionRuleResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetCompressionRuleResponseBody) SetGzip(v string) *GetCompressionRuleResponseBody {
	s.Gzip = &v
	return s
}

func (s *GetCompressionRuleResponseBody) SetRequestId(v string) *GetCompressionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCompressionRuleResponseBody) SetRule(v string) *GetCompressionRuleResponseBody {
	s.Rule = &v
	return s
}

func (s *GetCompressionRuleResponseBody) SetRuleEnable(v string) *GetCompressionRuleResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetCompressionRuleResponseBody) SetRuleName(v string) *GetCompressionRuleResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetCompressionRuleResponseBody) SetSequence(v int32) *GetCompressionRuleResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetCompressionRuleResponseBody) SetSiteVersion(v int32) *GetCompressionRuleResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetCompressionRuleResponseBody) SetZstd(v string) *GetCompressionRuleResponseBody {
	s.Zstd = &v
	return s
}

func (s *GetCompressionRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
