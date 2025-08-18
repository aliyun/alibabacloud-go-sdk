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
	// Configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Configuration type, which can be used to query global or rule configurations. Value range:
	//
	// - global: Query global configuration.
	//
	// - rule: Query rule configuration.
	//
	// example:
	//
	// rule
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
	// Port of the origin server accessed when using the HTTP protocol for origin.
	//
	// example:
	//
	// 8080
	OriginHttpPort *string `json:"OriginHttpPort,omitempty" xml:"OriginHttpPort,omitempty"`
	// Port of the origin server accessed when using the HTTPS protocol for origin.
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
	// SNI carried in the origin request.
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
	// Use the range chunk method for origin file download. Value range:
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
	// Request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter does not need to be set when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, e.g., (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter does not need to be set when adding a global configuration. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// Rule name. This parameter does not need to be set when adding a global configuration.
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
	// Version number of the site configuration. For sites with version management enabled, you can use this parameter to specify the effective version of the site configuration, defaulting to version 0.
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
