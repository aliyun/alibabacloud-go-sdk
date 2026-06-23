// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCompressionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrotli(v string) *CreateCompressionRuleRequest
	GetBrotli() *string
	SetGzip(v string) *CreateCompressionRuleRequest
	GetGzip() *string
	SetRule(v string) *CreateCompressionRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateCompressionRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateCompressionRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateCompressionRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateCompressionRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateCompressionRuleRequest
	GetSiteVersion() *int32
	SetZstd(v string) *CreateCompressionRuleRequest
	GetZstd() *string
}

type CreateCompressionRuleRequest struct {
	// Brotli compression. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	Brotli *string `json:"Brotli,omitempty" xml:"Brotli,omitempty"`
	// Gzip compression. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
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
	// - on: Enabled.
	//
	// - off: Disabled.
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
	// The site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1231231221***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with configuration version management enabled, you can use this parameter to specify the site version on which the configuration takes effect. The default value is version 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Zstd compression. Valid values:
	//
	// - on: Enabled.
	//
	// - off: Disabled.
	//
	// example:
	//
	// on
	Zstd *string `json:"Zstd,omitempty" xml:"Zstd,omitempty"`
}

func (s CreateCompressionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCompressionRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateCompressionRuleRequest) GetBrotli() *string {
	return s.Brotli
}

func (s *CreateCompressionRuleRequest) GetGzip() *string {
	return s.Gzip
}

func (s *CreateCompressionRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateCompressionRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateCompressionRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateCompressionRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateCompressionRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateCompressionRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateCompressionRuleRequest) GetZstd() *string {
	return s.Zstd
}

func (s *CreateCompressionRuleRequest) SetBrotli(v string) *CreateCompressionRuleRequest {
	s.Brotli = &v
	return s
}

func (s *CreateCompressionRuleRequest) SetGzip(v string) *CreateCompressionRuleRequest {
	s.Gzip = &v
	return s
}

func (s *CreateCompressionRuleRequest) SetRule(v string) *CreateCompressionRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateCompressionRuleRequest) SetRuleEnable(v string) *CreateCompressionRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateCompressionRuleRequest) SetRuleName(v string) *CreateCompressionRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateCompressionRuleRequest) SetSequence(v int32) *CreateCompressionRuleRequest {
	s.Sequence = &v
	return s
}

func (s *CreateCompressionRuleRequest) SetSiteId(v int64) *CreateCompressionRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateCompressionRuleRequest) SetSiteVersion(v int32) *CreateCompressionRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateCompressionRuleRequest) SetZstd(v string) *CreateCompressionRuleRequest {
	s.Zstd = &v
	return s
}

func (s *CreateCompressionRuleRequest) Validate() error {
	return dara.Validate(s)
}
