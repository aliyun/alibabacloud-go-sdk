// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetOriginRuleResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetOriginRuleResponseBody
	GetConfigType() *string
	SetDnsRecord(v string) *GetOriginRuleResponseBody
	GetDnsRecord() *string
	SetFollow302Enable(v string) *GetOriginRuleResponseBody
	GetFollow302Enable() *string
	SetFollow302MaxTries(v string) *GetOriginRuleResponseBody
	GetFollow302MaxTries() *string
	SetFollow302RetainArgs(v string) *GetOriginRuleResponseBody
	GetFollow302RetainArgs() *string
	SetFollow302RetainHeader(v string) *GetOriginRuleResponseBody
	GetFollow302RetainHeader() *string
	SetFollow302TargetHost(v string) *GetOriginRuleResponseBody
	GetFollow302TargetHost() *string
	SetOriginHost(v string) *GetOriginRuleResponseBody
	GetOriginHost() *string
	SetOriginHttpPort(v string) *GetOriginRuleResponseBody
	GetOriginHttpPort() *string
	SetOriginHttpsPort(v string) *GetOriginRuleResponseBody
	GetOriginHttpsPort() *string
	SetOriginMtls(v string) *GetOriginRuleResponseBody
	GetOriginMtls() *string
	SetOriginReadTimeout(v string) *GetOriginRuleResponseBody
	GetOriginReadTimeout() *string
	SetOriginScheme(v string) *GetOriginRuleResponseBody
	GetOriginScheme() *string
	SetOriginSni(v string) *GetOriginRuleResponseBody
	GetOriginSni() *string
	SetOriginVerify(v string) *GetOriginRuleResponseBody
	GetOriginVerify() *string
	SetRange(v string) *GetOriginRuleResponseBody
	GetRange() *string
	SetRangeChunkSize(v string) *GetOriginRuleResponseBody
	GetRangeChunkSize() *string
	SetRequestId(v string) *GetOriginRuleResponseBody
	GetRequestId() *string
	SetRule(v string) *GetOriginRuleResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetOriginRuleResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetOriginRuleResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetOriginRuleResponseBody
	GetSequence() *int32
	SetSiteVersion(v int32) *GetOriginRuleResponseBody
	GetSiteVersion() *int32
}

type GetOriginRuleResponseBody struct {
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type to query. Valid values:
	//
	// 	- global: global configurations.
	//
	// 	- rule: rule configurations.
	//
	// example:
	//
	// rule
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The hostname that overrides the resolved hostname of an incoming request.
	//
	// example:
	//
	// test.example.com
	DnsRecord *string `json:"DnsRecord,omitempty" xml:"DnsRecord,omitempty"`
	// 302 follow switch. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	Follow302Enable *string `json:"Follow302Enable,omitempty" xml:"Follow302Enable,omitempty"`
	// The maximum number of 302 follow times. Valid values: 1 to 5.
	//
	// example:
	//
	// 1
	Follow302MaxTries *string `json:"Follow302MaxTries,omitempty" xml:"Follow302MaxTries,omitempty"`
	// Retains the original request parameters switch. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	Follow302RetainArgs *string `json:"Follow302RetainArgs,omitempty" xml:"Follow302RetainArgs,omitempty"`
	// Retain the original request header switch. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	Follow302RetainHeader *string `json:"Follow302RetainHeader,omitempty" xml:"Follow302RetainHeader,omitempty"`
	// Target Host
	//
	// example:
	//
	// test.com
	Follow302TargetHost *string `json:"Follow302TargetHost,omitempty" xml:"Follow302TargetHost,omitempty"`
	// The Host header in origin requests.
	//
	// example:
	//
	// origin.example.com
	OriginHost *string `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
	// The origin port that is accessed when the HTTP protocol is used to back to the origin.
	//
	// example:
	//
	// 8080
	OriginHttpPort *string `json:"OriginHttpPort,omitempty" xml:"OriginHttpPort,omitempty"`
	// The origin port that is accessed when the HTTPS protocol is used to back to the origin.
	//
	// example:
	//
	// 4433
	OriginHttpsPort *string `json:"OriginHttpsPort,omitempty" xml:"OriginHttpsPort,omitempty"`
	// The MTLS switch. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	OriginMtls *string `json:"OriginMtls,omitempty" xml:"OriginMtls,omitempty"`
	// The read timeout period (in seconds) on the origin.
	//
	// example:
	//
	// 10
	OriginReadTimeout *string `json:"OriginReadTimeout,omitempty" xml:"OriginReadTimeout,omitempty"`
	// The protocol used for origin requests. Valid values:
	//
	// 	- http: HTTP.
	//
	// 	- https: HTTPS.
	//
	// 	- follow: follows the protocol used by the client.
	//
	// example:
	//
	// http
	OriginScheme *string `json:"OriginScheme,omitempty" xml:"OriginScheme,omitempty"`
	// The SNI in origin requests.
	//
	// example:
	//
	// origin.example.com
	OriginSni *string `json:"OriginSni,omitempty" xml:"OriginSni,omitempty"`
	// The origin certificate verification switch. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	OriginVerify *string `json:"OriginVerify,omitempty" xml:"OriginVerify,omitempty"`
	// Use range requests to download an object from the source. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// 	- force: enables range origin fetch by force.
	//
	// example:
	//
	// on
	Range *string `json:"Range,omitempty" xml:"Range,omitempty"`
	// The size of the range part. Valid values:
	//
	// 	- 512KB
	//
	// 	- 1MB
	//
	// 	- 2MB
	//
	// 	- 4MB
	//
	// example:
	//
	// 1MB
	RangeChunkSize *string `json:"RangeChunkSize,omitempty" xml:"RangeChunkSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The content of the rule. A conditional expression is used to match a user request. You do not need to set this parameter when you add global configuration. Use cases:
	//
	// 	- true: Match all incoming requests.
	//
	// 	- Set the value to a custom expression, for example, (http.host eq "video.example.com"): Match the specified request.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. Valid values: You do not need to set this parameter when you add global configuration. Valid values:
	//
	// 	- on
	//
	// 	- off
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. You do not need to set this parameter when you add global configurations.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule execution order. The smaller the value, the higher the priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The version number of the configurations. You can use this parameter to specify a version of your website to apply the feature settings. By default, version 0 is used.
	//
	// example:
	//
	// 1
	SiteVersion *int32 `json:"SiteVersion,omitempty" xml:"SiteVersion,omitempty"`
}

func (s GetOriginRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOriginRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetOriginRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetOriginRuleResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetOriginRuleResponseBody) GetDnsRecord() *string {
	return s.DnsRecord
}

func (s *GetOriginRuleResponseBody) GetFollow302Enable() *string {
	return s.Follow302Enable
}

func (s *GetOriginRuleResponseBody) GetFollow302MaxTries() *string {
	return s.Follow302MaxTries
}

func (s *GetOriginRuleResponseBody) GetFollow302RetainArgs() *string {
	return s.Follow302RetainArgs
}

func (s *GetOriginRuleResponseBody) GetFollow302RetainHeader() *string {
	return s.Follow302RetainHeader
}

func (s *GetOriginRuleResponseBody) GetFollow302TargetHost() *string {
	return s.Follow302TargetHost
}

func (s *GetOriginRuleResponseBody) GetOriginHost() *string {
	return s.OriginHost
}

func (s *GetOriginRuleResponseBody) GetOriginHttpPort() *string {
	return s.OriginHttpPort
}

func (s *GetOriginRuleResponseBody) GetOriginHttpsPort() *string {
	return s.OriginHttpsPort
}

func (s *GetOriginRuleResponseBody) GetOriginMtls() *string {
	return s.OriginMtls
}

func (s *GetOriginRuleResponseBody) GetOriginReadTimeout() *string {
	return s.OriginReadTimeout
}

func (s *GetOriginRuleResponseBody) GetOriginScheme() *string {
	return s.OriginScheme
}

func (s *GetOriginRuleResponseBody) GetOriginSni() *string {
	return s.OriginSni
}

func (s *GetOriginRuleResponseBody) GetOriginVerify() *string {
	return s.OriginVerify
}

func (s *GetOriginRuleResponseBody) GetRange() *string {
	return s.Range
}

func (s *GetOriginRuleResponseBody) GetRangeChunkSize() *string {
	return s.RangeChunkSize
}

func (s *GetOriginRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOriginRuleResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetOriginRuleResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetOriginRuleResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetOriginRuleResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetOriginRuleResponseBody) GetSiteVersion() *int32 {
	return s.SiteVersion
}

func (s *GetOriginRuleResponseBody) SetConfigId(v int64) *GetOriginRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetConfigType(v string) *GetOriginRuleResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetDnsRecord(v string) *GetOriginRuleResponseBody {
	s.DnsRecord = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetFollow302Enable(v string) *GetOriginRuleResponseBody {
	s.Follow302Enable = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetFollow302MaxTries(v string) *GetOriginRuleResponseBody {
	s.Follow302MaxTries = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetFollow302RetainArgs(v string) *GetOriginRuleResponseBody {
	s.Follow302RetainArgs = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetFollow302RetainHeader(v string) *GetOriginRuleResponseBody {
	s.Follow302RetainHeader = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetFollow302TargetHost(v string) *GetOriginRuleResponseBody {
	s.Follow302TargetHost = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetOriginHost(v string) *GetOriginRuleResponseBody {
	s.OriginHost = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetOriginHttpPort(v string) *GetOriginRuleResponseBody {
	s.OriginHttpPort = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetOriginHttpsPort(v string) *GetOriginRuleResponseBody {
	s.OriginHttpsPort = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetOriginMtls(v string) *GetOriginRuleResponseBody {
	s.OriginMtls = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetOriginReadTimeout(v string) *GetOriginRuleResponseBody {
	s.OriginReadTimeout = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetOriginScheme(v string) *GetOriginRuleResponseBody {
	s.OriginScheme = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetOriginSni(v string) *GetOriginRuleResponseBody {
	s.OriginSni = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetOriginVerify(v string) *GetOriginRuleResponseBody {
	s.OriginVerify = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetRange(v string) *GetOriginRuleResponseBody {
	s.Range = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetRangeChunkSize(v string) *GetOriginRuleResponseBody {
	s.RangeChunkSize = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetRequestId(v string) *GetOriginRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetRule(v string) *GetOriginRuleResponseBody {
	s.Rule = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetRuleEnable(v string) *GetOriginRuleResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetRuleName(v string) *GetOriginRuleResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetSequence(v int32) *GetOriginRuleResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetOriginRuleResponseBody) SetSiteVersion(v int32) *GetOriginRuleResponseBody {
	s.SiteVersion = &v
	return s
}

func (s *GetOriginRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
