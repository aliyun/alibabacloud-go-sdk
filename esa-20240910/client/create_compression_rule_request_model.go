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
	// Brotli compression. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	Brotli *string `json:"Brotli,omitempty" xml:"Brotli,omitempty"`
	// Gzip compression. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	Gzip *string `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - To match all incoming requests: Set the value to true
	//
	// - To match specific requests: Set the value to a custom expression, for example: (http.host eq \\"video.example.com\\")
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
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1231231221***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the effective version of the site configuration, defaulting to version 0.
	//
	// example:
	//
	// 0
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
	// Zstd compression. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
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
