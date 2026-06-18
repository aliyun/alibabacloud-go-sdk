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
	// The configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Overrides the DNS record for the origin request.
	//
	// example:
	//
	// test.example.com
	DnsRecord *string `json:"DnsRecord,omitempty" xml:"DnsRecord,omitempty"`
	// Indicates whether to follow 302 redirects for origin requests. Valid values:
	//
	// - `on`: Follows 302 redirects.
	//
	// - `off`: Does not follow 302 redirects.
	//
	// example:
	//
	// on
	Follow302Enable *string `json:"Follow302Enable,omitempty" xml:"Follow302Enable,omitempty"`
	// The maximum number of 302 redirects to follow. The value must be an integer from 1 to 5.
	//
	// example:
	//
	// 1
	Follow302MaxTries *string `json:"Follow302MaxTries,omitempty" xml:"Follow302MaxTries,omitempty"`
	// Indicates whether to retain the original request parameters when following a 302 redirect. Valid values:
	//
	// - `on`: Retains the parameters.
	//
	// - `off`: Does not retain the parameters.
	//
	// example:
	//
	// on
	Follow302RetainArgs *string `json:"Follow302RetainArgs,omitempty" xml:"Follow302RetainArgs,omitempty"`
	// Indicates whether to retain the original request headers when following a 302 redirect. Valid values:
	//
	// - `on`: Retains the headers.
	//
	// - `off`: Does not retain the headers.
	//
	// example:
	//
	// on
	Follow302RetainHeader *string `json:"Follow302RetainHeader,omitempty" xml:"Follow302RetainHeader,omitempty"`
	// The `Host` header for the redirected origin request.
	//
	// example:
	//
	// test.com
	Follow302TargetHost *string `json:"Follow302TargetHost,omitempty" xml:"Follow302TargetHost,omitempty"`
	// The `Host` header for the origin request.
	//
	// example:
	//
	// origin.example.com
	OriginHost *string `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
	// The origin server port for HTTP requests.
	//
	// example:
	//
	// 8080
	OriginHttpPort *string `json:"OriginHttpPort,omitempty" xml:"OriginHttpPort,omitempty"`
	// The origin server port for HTTPS requests.
	//
	// example:
	//
	// 4433
	OriginHttpsPort *string `json:"OriginHttpsPort,omitempty" xml:"OriginHttpsPort,omitempty"`
	// Indicates whether to enable mutual TLS (mTLS) for origin requests. Valid values:
	//
	// - `on`: Enables mTLS.
	//
	// - `off`: Disables mTLS.
	//
	// example:
	//
	// on
	OriginMtls *string `json:"OriginMtls,omitempty" xml:"OriginMtls,omitempty"`
	// The read timeout period for the origin server, in seconds.
	//
	// example:
	//
	// 10
	OriginReadTimeout *string `json:"OriginReadTimeout,omitempty" xml:"OriginReadTimeout,omitempty"`
	// The protocol for origin requests. Valid values:
	//
	// - `http`: The origin request uses HTTP.
	//
	// - `https`: The origin request uses HTTPS.
	//
	// - `follow`: The origin request uses the same protocol as the client request.
	//
	// example:
	//
	// http
	OriginScheme *string `json:"OriginScheme,omitempty" xml:"OriginScheme,omitempty"`
	// The Server Name Indication (SNI) for the origin request.
	//
	// example:
	//
	// origin.example.com
	OriginSni *string `json:"OriginSni,omitempty" xml:"OriginSni,omitempty"`
	// Indicates whether to enable origin certificate verification. Valid values:
	//
	// - `on`: Enables verification.
	//
	// - `off`: Disables verification.
	//
	// example:
	//
	// on
	OriginVerify *string `json:"OriginVerify,omitempty" xml:"OriginVerify,omitempty"`
	// Indicates whether to use range requests when fetching files from the origin server. Valid values:
	//
	// - `on`: Enables range requests.
	//
	// - `off`: Disables range requests.
	//
	// - `force`: Enforces range requests.
	//
	// example:
	//
	// on
	Range *string `json:"Range,omitempty" xml:"Range,omitempty"`
	// The size of each chunk for range requests. Valid values:
	//
	// - 512KB
	//
	// - 1MB
	//
	// - 2MB
	//
	// - 4MB
	//
	// example:
	//
	// 1MB
	RangeChunkSize *string `json:"RangeChunkSize,omitempty" xml:"RangeChunkSize,omitempty"`
	// The rule content, which is a conditional expression that matches user requests. This parameter is not required when you add a global configuration.
	//
	// - To match all incoming requests, set the value to `true`.
	//
	// - To match specific requests, set the value to a custom expression, such as `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Indicates whether the rule is enabled. This parameter is not required when you add a global configuration. Valid values:
	//
	// - `on`: Enables the rule.
	//
	// - `off`: Disables the rule.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. This parameter is not required when you add a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution priority of the rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The site ID. You can get this ID by calling the [ListSites](~~ListSites~~) operation.
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
