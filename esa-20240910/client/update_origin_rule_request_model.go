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
	// This parameter is required.
	ConfigId              *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
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
