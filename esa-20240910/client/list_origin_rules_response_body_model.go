// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOriginRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListOriginRulesResponseBodyConfigs) *ListOriginRulesResponseBody
	GetConfigs() []*ListOriginRulesResponseBodyConfigs
	SetPageNumber(v int32) *ListOriginRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOriginRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListOriginRulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListOriginRulesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListOriginRulesResponseBody
	GetTotalPage() *int32
}

type ListOriginRulesResponseBody struct {
	// Response body configuration.
	Configs []*ListOriginRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListOriginRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOriginRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOriginRulesResponseBody) GetConfigs() []*ListOriginRulesResponseBodyConfigs {
	return s.Configs
}

func (s *ListOriginRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOriginRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOriginRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOriginRulesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOriginRulesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListOriginRulesResponseBody) SetConfigs(v []*ListOriginRulesResponseBodyConfigs) *ListOriginRulesResponseBody {
	s.Configs = v
	return s
}

func (s *ListOriginRulesResponseBody) SetPageNumber(v int32) *ListOriginRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListOriginRulesResponseBody) SetPageSize(v int32) *ListOriginRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListOriginRulesResponseBody) SetRequestId(v string) *ListOriginRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOriginRulesResponseBody) SetTotalCount(v int32) *ListOriginRulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOriginRulesResponseBody) SetTotalPage(v int32) *ListOriginRulesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListOriginRulesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOriginRulesResponseBodyConfigs struct {
	// Configuration ID.
	//
	// example:
	//
	// 395386449776640
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, which can be used to query global or rule configurations. Value range:
	//
	// - global: Query global configuration.
	//
	// - rule: Query rule configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Rewrite the DNS resolution record for the origin request.
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
	// HOST carried in the origin request.
	//
	// example:
	//
	// origin.example.com
	OriginHost *string `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
	// The port of the origin server to access when using the HTTP protocol for origin requests.
	//
	// example:
	//
	// 8080
	OriginHttpPort *string `json:"OriginHttpPort,omitempty" xml:"OriginHttpPort,omitempty"`
	// The port of the origin server to access when using the HTTPS protocol for origin requests.
	//
	// example:
	//
	// 4433
	OriginHttpsPort *string `json:"OriginHttpsPort,omitempty" xml:"OriginHttpsPort,omitempty"`
	// mTLS switch. Value range:
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
	// Protocol used for the origin request. Value range:
	//
	// - http: Use HTTP protocol for origin.
	//
	// - https: Use HTTPS protocol for origin.
	//
	// - follow: Follow the client\\"s protocol for origin.
	//
	// example:
	//
	// http
	OriginScheme *string `json:"OriginScheme,omitempty" xml:"OriginScheme,omitempty"`
	// SNI carried in the back-to-origin request.
	//
	// example:
	//
	// origin.example.com
	OriginSni *string `json:"OriginSni,omitempty" xml:"OriginSni,omitempty"`
	// Origin certificate verification switch. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	OriginVerify *string `json:"OriginVerify,omitempty" xml:"OriginVerify,omitempty"`
	// Use range slicing to download files from the origin. Value range:
	//
	// - on: Enable
	//
	// - off: Disable
	//
	// - force: Force
	//
	// example:
	//
	// on
	Range          *string `json:"Range,omitempty" xml:"Range,omitempty"`
	RangeChunkSize *string `json:"RangeChunkSize,omitempty" xml:"RangeChunkSize,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, e.g., (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. Value range:
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
	// Version number of the site configuration. For sites with version management enabled, this parameter can specify the version of the site for which the configuration is effective, defaulting to version 0.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s ListOriginRulesResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListOriginRulesResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListOriginRulesResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListOriginRulesResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListOriginRulesResponseBodyConfigs) GetDnsRecord() *string {
	return s.DnsRecord
}

func (s *ListOriginRulesResponseBodyConfigs) GetFollow302Enable() *string {
	return s.Follow302Enable
}

func (s *ListOriginRulesResponseBodyConfigs) GetFollow302MaxTries() *string {
	return s.Follow302MaxTries
}

func (s *ListOriginRulesResponseBodyConfigs) GetFollow302RetainArgs() *string {
	return s.Follow302RetainArgs
}

func (s *ListOriginRulesResponseBodyConfigs) GetFollow302RetainHeader() *string {
	return s.Follow302RetainHeader
}

func (s *ListOriginRulesResponseBodyConfigs) GetFollow302TargetHost() *string {
	return s.Follow302TargetHost
}

func (s *ListOriginRulesResponseBodyConfigs) GetOriginHost() *string {
	return s.OriginHost
}

func (s *ListOriginRulesResponseBodyConfigs) GetOriginHttpPort() *string {
	return s.OriginHttpPort
}

func (s *ListOriginRulesResponseBodyConfigs) GetOriginHttpsPort() *string {
	return s.OriginHttpsPort
}

func (s *ListOriginRulesResponseBodyConfigs) GetOriginMtls() *string {
	return s.OriginMtls
}

func (s *ListOriginRulesResponseBodyConfigs) GetOriginReadTimeout() *string {
	return s.OriginReadTimeout
}

func (s *ListOriginRulesResponseBodyConfigs) GetOriginScheme() *string {
	return s.OriginScheme
}

func (s *ListOriginRulesResponseBodyConfigs) GetOriginSni() *string {
	return s.OriginSni
}

func (s *ListOriginRulesResponseBodyConfigs) GetOriginVerify() *string {
	return s.OriginVerify
}

func (s *ListOriginRulesResponseBodyConfigs) GetRange() *string {
	return s.Range
}

func (s *ListOriginRulesResponseBodyConfigs) GetRangeChunkSize() *string {
	return s.RangeChunkSize
}

func (s *ListOriginRulesResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListOriginRulesResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListOriginRulesResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListOriginRulesResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListOriginRulesResponseBodyConfigs) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *ListOriginRulesResponseBodyConfigs) SetConfigId(v int64) *ListOriginRulesResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetConfigType(v string) *ListOriginRulesResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetDnsRecord(v string) *ListOriginRulesResponseBodyConfigs {
	s.DnsRecord = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetFollow302Enable(v string) *ListOriginRulesResponseBodyConfigs {
	s.Follow302Enable = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetFollow302MaxTries(v string) *ListOriginRulesResponseBodyConfigs {
	s.Follow302MaxTries = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetFollow302RetainArgs(v string) *ListOriginRulesResponseBodyConfigs {
	s.Follow302RetainArgs = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetFollow302RetainHeader(v string) *ListOriginRulesResponseBodyConfigs {
	s.Follow302RetainHeader = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetFollow302TargetHost(v string) *ListOriginRulesResponseBodyConfigs {
	s.Follow302TargetHost = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetOriginHost(v string) *ListOriginRulesResponseBodyConfigs {
	s.OriginHost = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetOriginHttpPort(v string) *ListOriginRulesResponseBodyConfigs {
	s.OriginHttpPort = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetOriginHttpsPort(v string) *ListOriginRulesResponseBodyConfigs {
	s.OriginHttpsPort = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetOriginMtls(v string) *ListOriginRulesResponseBodyConfigs {
	s.OriginMtls = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetOriginReadTimeout(v string) *ListOriginRulesResponseBodyConfigs {
	s.OriginReadTimeout = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetOriginScheme(v string) *ListOriginRulesResponseBodyConfigs {
	s.OriginScheme = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetOriginSni(v string) *ListOriginRulesResponseBodyConfigs {
	s.OriginSni = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetOriginVerify(v string) *ListOriginRulesResponseBodyConfigs {
	s.OriginVerify = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetRange(v string) *ListOriginRulesResponseBodyConfigs {
	s.Range = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetRangeChunkSize(v string) *ListOriginRulesResponseBodyConfigs {
	s.RangeChunkSize = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetRule(v string) *ListOriginRulesResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetRuleEnable(v string) *ListOriginRulesResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetRuleName(v string) *ListOriginRulesResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetSequence(v int32) *ListOriginRulesResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) SetSiteVersion(v int32) *ListOriginRulesResponseBodyConfigs {
	s.SiteVersion = &v
	return s
}

func (s *ListOriginRulesResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
