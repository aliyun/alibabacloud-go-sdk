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
	DnsRecord             *string `json:"DnsRecord,omitempty" xml:"DnsRecord,omitempty"`
	Follow302Enable       *string `json:"Follow302Enable,omitempty" xml:"Follow302Enable,omitempty"`
	Follow302MaxTries     *string `json:"Follow302MaxTries,omitempty" xml:"Follow302MaxTries,omitempty"`
	Follow302RetainArgs   *string `json:"Follow302RetainArgs,omitempty" xml:"Follow302RetainArgs,omitempty"`
	Follow302RetainHeader *string `json:"Follow302RetainHeader,omitempty" xml:"Follow302RetainHeader,omitempty"`
	Follow302TargetHost   *string `json:"Follow302TargetHost,omitempty" xml:"Follow302TargetHost,omitempty"`
	OriginHost            *string `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
	OriginHttpPort        *string `json:"OriginHttpPort,omitempty" xml:"OriginHttpPort,omitempty"`
	OriginHttpsPort       *string `json:"OriginHttpsPort,omitempty" xml:"OriginHttpsPort,omitempty"`
	OriginMtls            *string `json:"OriginMtls,omitempty" xml:"OriginMtls,omitempty"`
	OriginReadTimeout     *string `json:"OriginReadTimeout,omitempty" xml:"OriginReadTimeout,omitempty"`
	OriginScheme          *string `json:"OriginScheme,omitempty" xml:"OriginScheme,omitempty"`
	OriginSni             *string `json:"OriginSni,omitempty" xml:"OriginSni,omitempty"`
	OriginVerify          *string `json:"OriginVerify,omitempty" xml:"OriginVerify,omitempty"`
	Range                 *string `json:"Range,omitempty" xml:"Range,omitempty"`
	RangeChunkSize        *string `json:"RangeChunkSize,omitempty" xml:"RangeChunkSize,omitempty"`
	Rule                  *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable            *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName              *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence              *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	SiteId      *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
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
