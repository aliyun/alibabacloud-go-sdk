// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpsBasicConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCiphersuite(v string) *GetHttpsBasicConfigurationResponseBody
	GetCiphersuite() *string
	SetCiphersuiteGroup(v string) *GetHttpsBasicConfigurationResponseBody
	GetCiphersuiteGroup() *string
	SetConfigId(v int64) *GetHttpsBasicConfigurationResponseBody
	GetConfigId() *int64
	SetConfigType(v string) *GetHttpsBasicConfigurationResponseBody
	GetConfigType() *string
	SetHttp2(v string) *GetHttpsBasicConfigurationResponseBody
	GetHttp2() *string
	SetHttp3(v string) *GetHttpsBasicConfigurationResponseBody
	GetHttp3() *string
	SetHttps(v string) *GetHttpsBasicConfigurationResponseBody
	GetHttps() *string
	SetOcspStapling(v string) *GetHttpsBasicConfigurationResponseBody
	GetOcspStapling() *string
	SetRequestId(v string) *GetHttpsBasicConfigurationResponseBody
	GetRequestId() *string
	SetRule(v string) *GetHttpsBasicConfigurationResponseBody
	GetRule() *string
	SetRuleEnable(v string) *GetHttpsBasicConfigurationResponseBody
	GetRuleEnable() *string
	SetRuleName(v string) *GetHttpsBasicConfigurationResponseBody
	GetRuleName() *string
	SetSequence(v int32) *GetHttpsBasicConfigurationResponseBody
	GetSequence() *int32
	SetTls10(v string) *GetHttpsBasicConfigurationResponseBody
	GetTls10() *string
	SetTls11(v string) *GetHttpsBasicConfigurationResponseBody
	GetTls11() *string
	SetTls12(v string) *GetHttpsBasicConfigurationResponseBody
	GetTls12() *string
	SetTls13(v string) *GetHttpsBasicConfigurationResponseBody
	GetTls13() *string
}

type GetHttpsBasicConfigurationResponseBody struct {
	Ciphersuite      *string `json:"Ciphersuite,omitempty" xml:"Ciphersuite,omitempty"`
	CiphersuiteGroup *string `json:"CiphersuiteGroup,omitempty" xml:"CiphersuiteGroup,omitempty"`
	ConfigId         *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	ConfigType       *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	Http2            *string `json:"Http2,omitempty" xml:"Http2,omitempty"`
	Http3            *string `json:"Http3,omitempty" xml:"Http3,omitempty"`
	Https            *string `json:"Https,omitempty" xml:"Https,omitempty"`
	OcspStapling     *string `json:"OcspStapling,omitempty" xml:"OcspStapling,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Rule             *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable       *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName         *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence         *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	Tls10            *string `json:"Tls10,omitempty" xml:"Tls10,omitempty"`
	Tls11            *string `json:"Tls11,omitempty" xml:"Tls11,omitempty"`
	Tls12            *string `json:"Tls12,omitempty" xml:"Tls12,omitempty"`
	Tls13            *string `json:"Tls13,omitempty" xml:"Tls13,omitempty"`
}

func (s GetHttpsBasicConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHttpsBasicConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetHttpsBasicConfigurationResponseBody) GetCiphersuite() *string {
	return s.Ciphersuite
}

func (s *GetHttpsBasicConfigurationResponseBody) GetCiphersuiteGroup() *string {
	return s.CiphersuiteGroup
}

func (s *GetHttpsBasicConfigurationResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetHttpsBasicConfigurationResponseBody) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetHttpsBasicConfigurationResponseBody) GetHttp2() *string {
	return s.Http2
}

func (s *GetHttpsBasicConfigurationResponseBody) GetHttp3() *string {
	return s.Http3
}

func (s *GetHttpsBasicConfigurationResponseBody) GetHttps() *string {
	return s.Https
}

func (s *GetHttpsBasicConfigurationResponseBody) GetOcspStapling() *string {
	return s.OcspStapling
}

func (s *GetHttpsBasicConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHttpsBasicConfigurationResponseBody) GetRule() *string {
	return s.Rule
}

func (s *GetHttpsBasicConfigurationResponseBody) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *GetHttpsBasicConfigurationResponseBody) GetRuleName() *string {
	return s.RuleName
}

func (s *GetHttpsBasicConfigurationResponseBody) GetSequence() *int32 {
	return s.Sequence
}

func (s *GetHttpsBasicConfigurationResponseBody) GetTls10() *string {
	return s.Tls10
}

func (s *GetHttpsBasicConfigurationResponseBody) GetTls11() *string {
	return s.Tls11
}

func (s *GetHttpsBasicConfigurationResponseBody) GetTls12() *string {
	return s.Tls12
}

func (s *GetHttpsBasicConfigurationResponseBody) GetTls13() *string {
	return s.Tls13
}

func (s *GetHttpsBasicConfigurationResponseBody) SetCiphersuite(v string) *GetHttpsBasicConfigurationResponseBody {
	s.Ciphersuite = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetCiphersuiteGroup(v string) *GetHttpsBasicConfigurationResponseBody {
	s.CiphersuiteGroup = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetConfigId(v int64) *GetHttpsBasicConfigurationResponseBody {
	s.ConfigId = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetConfigType(v string) *GetHttpsBasicConfigurationResponseBody {
	s.ConfigType = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetHttp2(v string) *GetHttpsBasicConfigurationResponseBody {
	s.Http2 = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetHttp3(v string) *GetHttpsBasicConfigurationResponseBody {
	s.Http3 = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetHttps(v string) *GetHttpsBasicConfigurationResponseBody {
	s.Https = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetOcspStapling(v string) *GetHttpsBasicConfigurationResponseBody {
	s.OcspStapling = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetRequestId(v string) *GetHttpsBasicConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetRule(v string) *GetHttpsBasicConfigurationResponseBody {
	s.Rule = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetRuleEnable(v string) *GetHttpsBasicConfigurationResponseBody {
	s.RuleEnable = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetRuleName(v string) *GetHttpsBasicConfigurationResponseBody {
	s.RuleName = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetSequence(v int32) *GetHttpsBasicConfigurationResponseBody {
	s.Sequence = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetTls10(v string) *GetHttpsBasicConfigurationResponseBody {
	s.Tls10 = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetTls11(v string) *GetHttpsBasicConfigurationResponseBody {
	s.Tls11 = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetTls12(v string) *GetHttpsBasicConfigurationResponseBody {
	s.Tls12 = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) SetTls13(v string) *GetHttpsBasicConfigurationResponseBody {
	s.Tls13 = &v
	return s
}

func (s *GetHttpsBasicConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
