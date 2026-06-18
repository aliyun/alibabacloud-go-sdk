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
	// A list of configurations.
	Configs []*ListOriginRulesResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
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

type ListOriginRulesResponseBodyConfigs struct {
	// The configuration ID.
	//
	// example:
	//
	// 395386449776640
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. You can use this parameter to query for global or rule-specific configurations. Valid values:
	//
	// - `global`: The global configuration.
	//
	// - `rule`: A rule-specific configuration.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Overrides the DNS record for the origin request.
	//
	// example:
	//
	// test.example.com
	DnsRecord *string `json:"DnsRecord,omitempty" xml:"DnsRecord,omitempty"`
	// Specifies whether to follow 302 redirects from the origin. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
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
	// Specifies whether to retain the original request parameters when following a redirect. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	Follow302RetainArgs *string `json:"Follow302RetainArgs,omitempty" xml:"Follow302RetainArgs,omitempty"`
	// Specifies whether to retain the original request header when following a redirect. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	Follow302RetainHeader *string `json:"Follow302RetainHeader,omitempty" xml:"Follow302RetainHeader,omitempty"`
	// The host to use for the origin request after following a 302 redirect.
	//
	// example:
	//
	// test.com
	Follow302TargetHost *string `json:"Follow302TargetHost,omitempty" xml:"Follow302TargetHost,omitempty"`
	// The `Host` header carried in the origin request.
	//
	// example:
	//
	// origin.example.com
	OriginHost *string `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
	// The origin server port used for origin requests over HTTP.
	//
	// example:
	//
	// 8080
	OriginHttpPort *string `json:"OriginHttpPort,omitempty" xml:"OriginHttpPort,omitempty"`
	// The origin server port used for origin requests over HTTPS.
	//
	// example:
	//
	// 4433
	OriginHttpsPort *string `json:"OriginHttpsPort,omitempty" xml:"OriginHttpsPort,omitempty"`
	// Specifies whether mTLS is enabled. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	OriginMtls *string `json:"OriginMtls,omitempty" xml:"OriginMtls,omitempty"`
	// The read timeout, in seconds, for the origin server.
	//
	// example:
	//
	// 10
	OriginReadTimeout *string `json:"OriginReadTimeout,omitempty" xml:"OriginReadTimeout,omitempty"`
	// The protocol used for origin requests. Valid values:
	//
	// - `http`: Use the HTTP protocol for origin requests.
	//
	// - `https`: Use the HTTPS protocol for origin requests.
	//
	// - `follow`: Use the same protocol as the client request.
	//
	// example:
	//
	// http
	OriginScheme *string `json:"OriginScheme,omitempty" xml:"OriginScheme,omitempty"`
	// The SNI carried in the origin request.
	//
	// example:
	//
	// origin.example.com
	OriginSni *string `json:"OriginSni,omitempty" xml:"OriginSni,omitempty"`
	// Specifies whether to verify the origin server certificate. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	OriginVerify *string `json:"OriginVerify,omitempty" xml:"OriginVerify,omitempty"`
	// Specifies whether to use range-based requests to retrieve files from the origin. Valid values:
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
	// The size of each chunk for range requests.
	//
	// example:
	//
	// 1MB
	RangeChunkSize *string `json:"RangeChunkSize,omitempty" xml:"RangeChunkSize,omitempty"`
	// The rule content, which uses a conditional expression to match user requests. This parameter is not required for global configurations. It supports two use cases:
	//
	// - To match all incoming requests, set the value to `true`.
	//
	// - To match specific requests, set the value to a custom expression, such as `(http.host eq "video.example.com")`.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the rule is enabled. This parameter is not required for global configurations. Valid values:
	//
	// - `on`: Enabled.
	//
	// - `off`: Disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. This parameter is not required for global configurations.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule execution order. Lower values indicate higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The site configuration version. If versioning is enabled for the site, this parameter specifies which version to use. The default is 0.
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
