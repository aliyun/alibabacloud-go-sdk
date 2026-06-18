// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOriginRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDnsRecord(v string) *CreateOriginRuleRequest
	GetDnsRecord() *string
	SetFollow302Enable(v string) *CreateOriginRuleRequest
	GetFollow302Enable() *string
	SetFollow302MaxTries(v string) *CreateOriginRuleRequest
	GetFollow302MaxTries() *string
	SetFollow302RetainArgs(v string) *CreateOriginRuleRequest
	GetFollow302RetainArgs() *string
	SetFollow302RetainHeader(v string) *CreateOriginRuleRequest
	GetFollow302RetainHeader() *string
	SetFollow302TargetHost(v string) *CreateOriginRuleRequest
	GetFollow302TargetHost() *string
	SetOriginHost(v string) *CreateOriginRuleRequest
	GetOriginHost() *string
	SetOriginHttpPort(v string) *CreateOriginRuleRequest
	GetOriginHttpPort() *string
	SetOriginHttpsPort(v string) *CreateOriginRuleRequest
	GetOriginHttpsPort() *string
	SetOriginMtls(v string) *CreateOriginRuleRequest
	GetOriginMtls() *string
	SetOriginReadTimeout(v string) *CreateOriginRuleRequest
	GetOriginReadTimeout() *string
	SetOriginScheme(v string) *CreateOriginRuleRequest
	GetOriginScheme() *string
	SetOriginSni(v string) *CreateOriginRuleRequest
	GetOriginSni() *string
	SetOriginVerify(v string) *CreateOriginRuleRequest
	GetOriginVerify() *string
	SetRange(v string) *CreateOriginRuleRequest
	GetRange() *string
	SetRangeChunkSize(v string) *CreateOriginRuleRequest
	GetRangeChunkSize() *string
	SetRule(v string) *CreateOriginRuleRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateOriginRuleRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateOriginRuleRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateOriginRuleRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateOriginRuleRequest
	GetSiteId() *int64
	SetSiteVersion(v int32) *CreateOriginRuleRequest
	GetSiteVersion() *int32
}

type CreateOriginRuleRequest struct {
	// Overrides the DNS Record for Origin requests.
	//
	// example:
	//
	// test.example.com
	DnsRecord *string `json:"DnsRecord,omitempty" xml:"DnsRecord,omitempty"`
	// Specifies whether to follow 302 redirects from the Origin. Valid values:
	//
	// - `on`: Enables following redirects.
	//
	// - `off`: Disables following redirects.
	//
	// example:
	//
	// on
	Follow302Enable *string `json:"Follow302Enable,omitempty" xml:"Follow302Enable,omitempty"`
	// The maximum number of 302 redirects to follow. Valid range: 1 to 5.
	//
	// example:
	//
	// 1
	Follow302MaxTries *string `json:"Follow302MaxTries,omitempty" xml:"Follow302MaxTries,omitempty"`
	// Specifies whether to retain the original request parameters when following a 302 redirect. Valid values:
	//
	// - `on`: Retains the request parameters.
	//
	// - `off`: Does not retain the request parameters.
	//
	// example:
	//
	// on
	Follow302RetainArgs *string `json:"Follow302RetainArgs,omitempty" xml:"Follow302RetainArgs,omitempty"`
	// Specifies whether to retain the original request header when following a 302 redirect. Valid values:
	//
	// - `on`: Retains the request header.
	//
	// - `off`: Does not retain the request header.
	//
	// example:
	//
	// on
	Follow302RetainHeader *string `json:"Follow302RetainHeader,omitempty" xml:"Follow302RetainHeader,omitempty"`
	// The Origin `Host` header to use after a 302 redirect.
	//
	// example:
	//
	// test.com
	Follow302TargetHost *string `json:"Follow302TargetHost,omitempty" xml:"Follow302TargetHost,omitempty"`
	// The `Host` header to use in Origin requests.
	//
	// example:
	//
	// origin.example.com
	OriginHost *string `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
	// The Origin Port to use for HTTP Origin requests.
	//
	// example:
	//
	// 8080
	OriginHttpPort *string `json:"OriginHttpPort,omitempty" xml:"OriginHttpPort,omitempty"`
	// The Origin Port to use for HTTPS Origin requests.
	//
	// example:
	//
	// 4433
	OriginHttpsPort *string `json:"OriginHttpsPort,omitempty" xml:"OriginHttpsPort,omitempty"`
	// Specifies whether to enable mutual Transport Layer Security (mTLS) for Origin connections. Valid values:
	//
	// - `on`: Enables mTLS.
	//
	// - `off`: Disables mTLS.
	//
	// example:
	//
	// on
	OriginMtls *string `json:"OriginMtls,omitempty" xml:"OriginMtls,omitempty"`
	// The Origin read timeout in seconds.
	//
	// example:
	//
	// 10
	OriginReadTimeout *string `json:"OriginReadTimeout,omitempty" xml:"OriginReadTimeout,omitempty"`
	// The protocol for Origin requests. Valid values:
	//
	// - `http`: Uses the HTTP protocol.
	//
	// - `https`: Uses the HTTPS protocol.
	//
	// - `follow`: Uses the same protocol as the client request.
	//
	// example:
	//
	// http
	OriginScheme *string `json:"OriginScheme,omitempty" xml:"OriginScheme,omitempty"`
	// The Server Name Indication (SNI) to use in Origin requests.
	//
	// example:
	//
	// origin.example.com
	OriginSni *string `json:"OriginSni,omitempty" xml:"OriginSni,omitempty"`
	// Specifies whether to enable Origin Certificate Verification. Valid values:
	//
	// - `on`: Enables verification.
	//
	// - `off`: Disables verification.
	//
	// example:
	//
	// on
	OriginVerify *string `json:"OriginVerify,omitempty" xml:"OriginVerify,omitempty"`
	// Specifies whether to use range requests to download files from the Origin. Valid values:
	//
	// - `on`: Enables range requests.
	//
	// - `off`: Disables range requests.
	//
	// - `force`: Forces range requests.
	//
	// example:
	//
	// on
	Range *string `json:"Range,omitempty" xml:"Range,omitempty"`
	// The size of each chunk for range requests. Valid values:
	//
	// - `512KB`
	//
	// - `1MB`
	//
	// - `2MB`
	//
	// - `4MB`
	//
	// example:
	//
	// 1MB
	RangeChunkSize *string `json:"RangeChunkSize,omitempty" xml:"RangeChunkSize,omitempty"`
	// The content of the rule, a Conditional Expression that matches user requests. Not required when creating a Global Configuration. There are two scenarios:
	//
	// - To match all incoming requests, set the value to `true`.
	//
	// - To match specific requests, set the value to a custom expression. Example: `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. Not required when creating a Global Configuration. Valid values:
	//
	// - `on`: Enables the rule.
	//
	// - `off`: Disables the rule.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. Not required when creating a Global Configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The priority of the rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The ID of the site. You can obtain this ID by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 340035003106221
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// For sites with version management enabled, this specifies the version to which the configuration applies. The default value is 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s CreateOriginRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOriginRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateOriginRuleRequest) GetDnsRecord() *string {
	return s.DnsRecord
}

func (s *CreateOriginRuleRequest) GetFollow302Enable() *string {
	return s.Follow302Enable
}

func (s *CreateOriginRuleRequest) GetFollow302MaxTries() *string {
	return s.Follow302MaxTries
}

func (s *CreateOriginRuleRequest) GetFollow302RetainArgs() *string {
	return s.Follow302RetainArgs
}

func (s *CreateOriginRuleRequest) GetFollow302RetainHeader() *string {
	return s.Follow302RetainHeader
}

func (s *CreateOriginRuleRequest) GetFollow302TargetHost() *string {
	return s.Follow302TargetHost
}

func (s *CreateOriginRuleRequest) GetOriginHost() *string {
	return s.OriginHost
}

func (s *CreateOriginRuleRequest) GetOriginHttpPort() *string {
	return s.OriginHttpPort
}

func (s *CreateOriginRuleRequest) GetOriginHttpsPort() *string {
	return s.OriginHttpsPort
}

func (s *CreateOriginRuleRequest) GetOriginMtls() *string {
	return s.OriginMtls
}

func (s *CreateOriginRuleRequest) GetOriginReadTimeout() *string {
	return s.OriginReadTimeout
}

func (s *CreateOriginRuleRequest) GetOriginScheme() *string {
	return s.OriginScheme
}

func (s *CreateOriginRuleRequest) GetOriginSni() *string {
	return s.OriginSni
}

func (s *CreateOriginRuleRequest) GetOriginVerify() *string {
	return s.OriginVerify
}

func (s *CreateOriginRuleRequest) GetRange() *string {
	return s.Range
}

func (s *CreateOriginRuleRequest) GetRangeChunkSize() *string {
	return s.RangeChunkSize
}

func (s *CreateOriginRuleRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateOriginRuleRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateOriginRuleRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateOriginRuleRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateOriginRuleRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateOriginRuleRequest) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *CreateOriginRuleRequest) SetDnsRecord(v string) *CreateOriginRuleRequest {
	s.DnsRecord = &v
	return s
}

func (s *CreateOriginRuleRequest) SetFollow302Enable(v string) *CreateOriginRuleRequest {
	s.Follow302Enable = &v
	return s
}

func (s *CreateOriginRuleRequest) SetFollow302MaxTries(v string) *CreateOriginRuleRequest {
	s.Follow302MaxTries = &v
	return s
}

func (s *CreateOriginRuleRequest) SetFollow302RetainArgs(v string) *CreateOriginRuleRequest {
	s.Follow302RetainArgs = &v
	return s
}

func (s *CreateOriginRuleRequest) SetFollow302RetainHeader(v string) *CreateOriginRuleRequest {
	s.Follow302RetainHeader = &v
	return s
}

func (s *CreateOriginRuleRequest) SetFollow302TargetHost(v string) *CreateOriginRuleRequest {
	s.Follow302TargetHost = &v
	return s
}

func (s *CreateOriginRuleRequest) SetOriginHost(v string) *CreateOriginRuleRequest {
	s.OriginHost = &v
	return s
}

func (s *CreateOriginRuleRequest) SetOriginHttpPort(v string) *CreateOriginRuleRequest {
	s.OriginHttpPort = &v
	return s
}

func (s *CreateOriginRuleRequest) SetOriginHttpsPort(v string) *CreateOriginRuleRequest {
	s.OriginHttpsPort = &v
	return s
}

func (s *CreateOriginRuleRequest) SetOriginMtls(v string) *CreateOriginRuleRequest {
	s.OriginMtls = &v
	return s
}

func (s *CreateOriginRuleRequest) SetOriginReadTimeout(v string) *CreateOriginRuleRequest {
	s.OriginReadTimeout = &v
	return s
}

func (s *CreateOriginRuleRequest) SetOriginScheme(v string) *CreateOriginRuleRequest {
	s.OriginScheme = &v
	return s
}

func (s *CreateOriginRuleRequest) SetOriginSni(v string) *CreateOriginRuleRequest {
	s.OriginSni = &v
	return s
}

func (s *CreateOriginRuleRequest) SetOriginVerify(v string) *CreateOriginRuleRequest {
	s.OriginVerify = &v
	return s
}

func (s *CreateOriginRuleRequest) SetRange(v string) *CreateOriginRuleRequest {
	s.Range = &v
	return s
}

func (s *CreateOriginRuleRequest) SetRangeChunkSize(v string) *CreateOriginRuleRequest {
	s.RangeChunkSize = &v
	return s
}

func (s *CreateOriginRuleRequest) SetRule(v string) *CreateOriginRuleRequest {
	s.Rule = &v
	return s
}

func (s *CreateOriginRuleRequest) SetRuleEnable(v string) *CreateOriginRuleRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateOriginRuleRequest) SetRuleName(v string) *CreateOriginRuleRequest {
	s.RuleName = &v
	return s
}

func (s *CreateOriginRuleRequest) SetSequence(v int32) *CreateOriginRuleRequest {
	s.Sequence = &v
	return s
}

func (s *CreateOriginRuleRequest) SetSiteId(v int64) *CreateOriginRuleRequest {
	s.SiteId = &v
	return s
}

func (s *CreateOriginRuleRequest) SetSiteVersion(v int32) *CreateOriginRuleRequest {
	s.SiteVersion = &v
	return s
}

func (s *CreateOriginRuleRequest) Validate() error {
	return dara.Validate(s)
}
