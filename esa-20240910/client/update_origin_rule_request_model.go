// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOriginRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *UpdateOriginRuleRequest
	GetConfigId() *int64
	SetDnsRecord(v string) *UpdateOriginRuleRequest
	GetDnsRecord() *string
	SetFollow302Enable(v string) *UpdateOriginRuleRequest
	GetFollow302Enable() *string
	SetFollow302MaxTries(v string) *UpdateOriginRuleRequest
	GetFollow302MaxTries() *string
	SetFollow302RetainArgs(v string) *UpdateOriginRuleRequest
	GetFollow302RetainArgs() *string
	SetFollow302RetainHeader(v string) *UpdateOriginRuleRequest
	GetFollow302RetainHeader() *string
	SetFollow302TargetHost(v string) *UpdateOriginRuleRequest
	GetFollow302TargetHost() *string
	SetOriginHost(v string) *UpdateOriginRuleRequest
	GetOriginHost() *string
	SetOriginHttpPort(v string) *UpdateOriginRuleRequest
	GetOriginHttpPort() *string
	SetOriginHttpsPort(v string) *UpdateOriginRuleRequest
	GetOriginHttpsPort() *string
	SetOriginMtls(v string) *UpdateOriginRuleRequest
	GetOriginMtls() *string
	SetOriginReadTimeout(v string) *UpdateOriginRuleRequest
	GetOriginReadTimeout() *string
	SetOriginScheme(v string) *UpdateOriginRuleRequest
	GetOriginScheme() *string
	SetOriginSni(v string) *UpdateOriginRuleRequest
	GetOriginSni() *string
	SetOriginVerify(v string) *UpdateOriginRuleRequest
	GetOriginVerify() *string
	SetRange(v string) *UpdateOriginRuleRequest
	GetRange() *string
	SetRangeChunkSize(v string) *UpdateOriginRuleRequest
	GetRangeChunkSize() *string
	SetRule(v string) *UpdateOriginRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateOriginRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateOriginRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateOriginRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateOriginRuleRequest
	GetSiteId() *int64
}

type UpdateOriginRuleRequest struct {
	// Configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Rewrite the DNS resolution record of the origin request.
	//
	// example:
	//
	// test.example.com
	DnsRecord             *string `json:"DnsRecord,omitempty" xml:"DnsRecord,omitempty"`
	Follow302Enable       *string `json:"Follow302Enable,omitempty" xml:"Follow302Enable,omitempty"`
	Follow302MaxTries     *string `json:"Follow302MaxTries,omitempty" xml:"Follow302MaxTries,omitempty"`
	Follow302RetainArgs   *string `json:"Follow302RetainArgs,omitempty" xml:"Follow302RetainArgs,omitempty"`
	Follow302RetainHeader *string `json:"Follow302RetainHeader,omitempty" xml:"Follow302RetainHeader,omitempty"`
	Follow302TargetHost   *string `json:"Follow302TargetHost,omitempty" xml:"Follow302TargetHost,omitempty"`
	// The HOST carried in the origin request.
	//
	// example:
	//
	// origin.example.com
	OriginHost *string `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
	// Port of the origin server when using HTTP protocol for origin pull.
	//
	// example:
	//
	// 8080
	OriginHttpPort *string `json:"OriginHttpPort,omitempty" xml:"OriginHttpPort,omitempty"`
	// Port of the origin server when using HTTPS protocol for origin pull.
	//
	// example:
	//
	// 4433
	OriginHttpsPort *string `json:"OriginHttpsPort,omitempty" xml:"OriginHttpsPort,omitempty"`
	// mTLS switch. Valid values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	OriginMtls        *string `json:"OriginMtls,omitempty" xml:"OriginMtls,omitempty"`
	OriginReadTimeout *string `json:"OriginReadTimeout,omitempty" xml:"OriginReadTimeout,omitempty"`
	// Protocol used for the origin request. Valid values:
	//
	// - http: Use HTTP protocol for origin pull.
	//
	// - https: Use HTTPS protocol for origin pull.
	//
	// - follow: Follow the client\\"s protocol for origin pull.
	//
	// example:
	//
	// http
	OriginScheme *string `json:"OriginScheme,omitempty" xml:"OriginScheme,omitempty"`
	// SNI carried in the origin request.
	//
	// example:
	//
	// origin.example.com
	OriginSni *string `json:"OriginSni,omitempty" xml:"OriginSni,omitempty"`
	// Origin certificate verification switch. Valid values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	OriginVerify *string `json:"OriginVerify,omitempty" xml:"OriginVerify,omitempty"`
	// Use range chunking for origin pull file download. Valid values:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// - force: Force.
	//
	// example:
	//
	// on
	Range          *string `json:"Range,omitempty" xml:"Range,omitempty"`
	RangeChunkSize *string `json:"RangeChunkSize,omitempty" xml:"RangeChunkSize,omitempty"`
	// Rule content, used to match user requests with conditional expressions. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, for example: (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. Valid values:
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
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateOriginRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginRuleRequest) GoString() string {
	return s.String()
}

func (s *UpdateOriginRuleRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateOriginRuleRequest) GetDnsRecord() *string {
	return s.DnsRecord
}

func (s *UpdateOriginRuleRequest) GetFollow302Enable() *string {
	return s.Follow302Enable
}

func (s *UpdateOriginRuleRequest) GetFollow302MaxTries() *string {
	return s.Follow302MaxTries
}

func (s *UpdateOriginRuleRequest) GetFollow302RetainArgs() *string {
	return s.Follow302RetainArgs
}

func (s *UpdateOriginRuleRequest) GetFollow302RetainHeader() *string {
	return s.Follow302RetainHeader
}

func (s *UpdateOriginRuleRequest) GetFollow302TargetHost() *string {
	return s.Follow302TargetHost
}

func (s *UpdateOriginRuleRequest) GetOriginHost() *string {
	return s.OriginHost
}

func (s *UpdateOriginRuleRequest) GetOriginHttpPort() *string {
	return s.OriginHttpPort
}

func (s *UpdateOriginRuleRequest) GetOriginHttpsPort() *string {
	return s.OriginHttpsPort
}

func (s *UpdateOriginRuleRequest) GetOriginMtls() *string {
	return s.OriginMtls
}

func (s *UpdateOriginRuleRequest) GetOriginReadTimeout() *string {
	return s.OriginReadTimeout
}

func (s *UpdateOriginRuleRequest) GetOriginScheme() *string {
	return s.OriginScheme
}

func (s *UpdateOriginRuleRequest) GetOriginSni() *string {
	return s.OriginSni
}

func (s *UpdateOriginRuleRequest) GetOriginVerify() *string {
	return s.OriginVerify
}

func (s *UpdateOriginRuleRequest) GetRange() *string {
	return s.Range
}

func (s *UpdateOriginRuleRequest) GetRangeChunkSize() *string {
	return s.RangeChunkSize
}

func (s *UpdateOriginRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateOriginRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateOriginRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateOriginRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateOriginRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateOriginRuleRequest) SetConfigId(v int64) *UpdateOriginRuleRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetDnsRecord(v string) *UpdateOriginRuleRequest {
	s.DnsRecord = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetFollow302Enable(v string) *UpdateOriginRuleRequest {
	s.Follow302Enable = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetFollow302MaxTries(v string) *UpdateOriginRuleRequest {
	s.Follow302MaxTries = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetFollow302RetainArgs(v string) *UpdateOriginRuleRequest {
	s.Follow302RetainArgs = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetFollow302RetainHeader(v string) *UpdateOriginRuleRequest {
	s.Follow302RetainHeader = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetFollow302TargetHost(v string) *UpdateOriginRuleRequest {
	s.Follow302TargetHost = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetOriginHost(v string) *UpdateOriginRuleRequest {
	s.OriginHost = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetOriginHttpPort(v string) *UpdateOriginRuleRequest {
	s.OriginHttpPort = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetOriginHttpsPort(v string) *UpdateOriginRuleRequest {
	s.OriginHttpsPort = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetOriginMtls(v string) *UpdateOriginRuleRequest {
	s.OriginMtls = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetOriginReadTimeout(v string) *UpdateOriginRuleRequest {
	s.OriginReadTimeout = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetOriginScheme(v string) *UpdateOriginRuleRequest {
	s.OriginScheme = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetOriginSni(v string) *UpdateOriginRuleRequest {
	s.OriginSni = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetOriginVerify(v string) *UpdateOriginRuleRequest {
	s.OriginVerify = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetRange(v string) *UpdateOriginRuleRequest {
	s.Range = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetRangeChunkSize(v string) *UpdateOriginRuleRequest {
	s.RangeChunkSize = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetRule(v string) *UpdateOriginRuleRequest {
	s.Rule = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetRuleEnable(v string) *UpdateOriginRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetRuleName(v string) *UpdateOriginRuleRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetSequence(v int32) *UpdateOriginRuleRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateOriginRuleRequest) SetSiteId(v int64) *UpdateOriginRuleRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateOriginRuleRequest) Validate() error {
	return dara.Validate(s)
}
