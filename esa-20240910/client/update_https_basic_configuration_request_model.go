// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpsBasicConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCiphersuite(v string) *UpdateHttpsBasicConfigurationRequest
	GetCiphersuite() *string
	SetCiphersuiteGroup(v string) *UpdateHttpsBasicConfigurationRequest
	GetCiphersuiteGroup() *string
	SetConfigId(v int64) *UpdateHttpsBasicConfigurationRequest
	GetConfigId() *int64
	SetHttp2(v string) *UpdateHttpsBasicConfigurationRequest
	GetHttp2() *string
	SetHttp3(v string) *UpdateHttpsBasicConfigurationRequest
	GetHttp3() *string
	SetHttps(v string) *UpdateHttpsBasicConfigurationRequest
	GetHttps() *string
	SetOcspStapling(v string) *UpdateHttpsBasicConfigurationRequest
	GetOcspStapling() *string
	SetRule(v string) *UpdateHttpsBasicConfigurationRequest
	GetRule() *string
	SetRuleEnable(v string) *UpdateHttpsBasicConfigurationRequest
	GetRuleEnable() *string
	SetRuleName(v string) *UpdateHttpsBasicConfigurationRequest
	GetRuleName() *string
	SetSequence(v int32) *UpdateHttpsBasicConfigurationRequest
	GetSequence() *int32
	SetSiteId(v int64) *UpdateHttpsBasicConfigurationRequest
	GetSiteId() *int64
	SetTls10(v string) *UpdateHttpsBasicConfigurationRequest
	GetTls10() *string
	SetTls11(v string) *UpdateHttpsBasicConfigurationRequest
	GetTls11() *string
	SetTls12(v string) *UpdateHttpsBasicConfigurationRequest
	GetTls12() *string
	SetTls13(v string) *UpdateHttpsBasicConfigurationRequest
	GetTls13() *string
}

type UpdateHttpsBasicConfigurationRequest struct {
	Ciphersuite      *string `json:"Ciphersuite,omitempty" xml:"Ciphersuite,omitempty"`
	CiphersuiteGroup *string `json:"CiphersuiteGroup,omitempty" xml:"CiphersuiteGroup,omitempty"`
	// This parameter is required.
	ConfigId     *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Http2        *string `json:"Http2,omitempty" xml:"Http2,omitempty"`
	Http3        *string `json:"Http3,omitempty" xml:"Http3,omitempty"`
	Https        *string `json:"Https,omitempty" xml:"Https,omitempty"`
	OcspStapling *string `json:"OcspStapling,omitempty" xml:"OcspStapling,omitempty"`
	Rule         *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable   *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName     *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence     *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// This parameter is required.
	SiteId *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	Tls10  *string `json:"Tls10,omitempty" xml:"Tls10,omitempty"`
	Tls11  *string `json:"Tls11,omitempty" xml:"Tls11,omitempty"`
	Tls12  *string `json:"Tls12,omitempty" xml:"Tls12,omitempty"`
	Tls13  *string `json:"Tls13,omitempty" xml:"Tls13,omitempty"`
}

func (s UpdateHttpsBasicConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpsBasicConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpsBasicConfigurationRequest) GetCiphersuite() *string {
	return s.Ciphersuite
}

func (s *UpdateHttpsBasicConfigurationRequest) GetCiphersuiteGroup() *string {
	return s.CiphersuiteGroup
}

func (s *UpdateHttpsBasicConfigurationRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *UpdateHttpsBasicConfigurationRequest) GetHttp2() *string {
	return s.Http2
}

func (s *UpdateHttpsBasicConfigurationRequest) GetHttp3() *string {
	return s.Http3
}

func (s *UpdateHttpsBasicConfigurationRequest) GetHttps() *string {
	return s.Https
}

func (s *UpdateHttpsBasicConfigurationRequest) GetOcspStapling() *string {
	return s.OcspStapling
}

func (s *UpdateHttpsBasicConfigurationRequest) GetRule() *string {
	return s.Rule
}

func (s *UpdateHttpsBasicConfigurationRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *UpdateHttpsBasicConfigurationRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *UpdateHttpsBasicConfigurationRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *UpdateHttpsBasicConfigurationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateHttpsBasicConfigurationRequest) GetTls10() *string {
	return s.Tls10
}

func (s *UpdateHttpsBasicConfigurationRequest) GetTls11() *string {
	return s.Tls11
}

func (s *UpdateHttpsBasicConfigurationRequest) GetTls12() *string {
	return s.Tls12
}

func (s *UpdateHttpsBasicConfigurationRequest) GetTls13() *string {
	return s.Tls13
}

func (s *UpdateHttpsBasicConfigurationRequest) SetCiphersuite(v string) *UpdateHttpsBasicConfigurationRequest {
	s.Ciphersuite = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) SetCiphersuiteGroup(v string) *UpdateHttpsBasicConfigurationRequest {
	s.CiphersuiteGroup = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) SetConfigId(v int64) *UpdateHttpsBasicConfigurationRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) SetHttp2(v string) *UpdateHttpsBasicConfigurationRequest {
	s.Http2 = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) SetHttp3(v string) *UpdateHttpsBasicConfigurationRequest {
	s.Http3 = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) SetHttps(v string) *UpdateHttpsBasicConfigurationRequest {
	s.Https = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) SetOcspStapling(v string) *UpdateHttpsBasicConfigurationRequest {
	s.OcspStapling = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) SetRule(v string) *UpdateHttpsBasicConfigurationRequest {
	s.Rule = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) SetRuleEnable(v string) *UpdateHttpsBasicConfigurationRequest {
	s.RuleEnable = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) SetRuleName(v string) *UpdateHttpsBasicConfigurationRequest {
	s.RuleName = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) SetSequence(v int32) *UpdateHttpsBasicConfigurationRequest {
	s.Sequence = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) SetSiteId(v int64) *UpdateHttpsBasicConfigurationRequest {
	s.SiteId = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) SetTls10(v string) *UpdateHttpsBasicConfigurationRequest {
	s.Tls10 = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) SetTls11(v string) *UpdateHttpsBasicConfigurationRequest {
	s.Tls11 = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) SetTls12(v string) *UpdateHttpsBasicConfigurationRequest {
	s.Tls12 = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) SetTls13(v string) *UpdateHttpsBasicConfigurationRequest {
	s.Tls13 = &v
	return s
}

func (s *UpdateHttpsBasicConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
