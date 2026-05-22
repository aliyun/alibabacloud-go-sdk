// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCompressionRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrotli(v string) *UpdateCompressionRuleRequest
	GetBrotli() *string
	SetConfigId(v int64) *UpdateCompressionRuleRequest
	GetConfigId() *int64
	SetGzip(v string) *UpdateCompressionRuleRequest
	GetGzip() *string
	SetRule(v string) *UpdateCompressionRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateCompressionRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateCompressionRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateCompressionRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateCompressionRuleRequest
	GetSiteId() *int64
	SetZstd(v string) *UpdateCompressionRuleRequest
	GetZstd() *string
}

type UpdateCompressionRuleRequest struct {
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
	// Configuration ID. It can be obtained by calling the [ListCompressionRules](~~ListCompressionRules~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
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
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
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

func (s UpdateCompressionRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCompressionRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateCompressionRuleRequest) GetBrotli() *string {
	return s.Brotli
}

func (s *UpdateCompressionRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateCompressionRuleRequest) GetGzip() *string {
	return s.Gzip
}

func (s *UpdateCompressionRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateCompressionRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateCompressionRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateCompressionRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateCompressionRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateCompressionRuleRequest) GetZstd() *string {
	return s.Zstd
}

func (s *UpdateCompressionRuleRequest) SetBrotli(v string) *UpdateCompressionRuleRequest {
	s.Brotli = &v
	return s
}

func (s *UpdateCompressionRuleRequest) SetConfigId(v int64) *UpdateCompressionRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateCompressionRuleRequest) SetGzip(v string) *UpdateCompressionRuleRequest {
	s.Gzip = &v
	return s
}

func (s *UpdateCompressionRuleRequest) SetRule(v string) *UpdateCompressionRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateCompressionRuleRequest) SetRuleEnable(v string) *UpdateCompressionRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateCompressionRuleRequest) SetRuleName(v string) *UpdateCompressionRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateCompressionRuleRequest) SetSequence(v int32) *UpdateCompressionRuleRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateCompressionRuleRequest) SetSiteId(v int64) *UpdateCompressionRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateCompressionRuleRequest) SetZstd(v string) *UpdateCompressionRuleRequest {
	s.Zstd = &v
	return s
}

func (s *UpdateCompressionRuleRequest) Validate() error {
	return dara.Validate(s)
}
