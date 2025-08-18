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
	// Custom cipher suite, indicating the specific encryption algorithm selected when CiphersuiteGroup is set to custom.
	//
	// example:
	//
	// TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
	Ciphersuite *string `json:"Ciphersuite,omitempty" xml:"Ciphersuite,omitempty"`
	// Cipher suite group. Default is all cipher suites. Value range:
	//
	// - all: All cipher suites.
	//
	// - strict: Strong cipher suites.
	//
	// - custom: Custom cipher suites.
	//
	// example:
	//
	// all
	CiphersuiteGroup *string `json:"CiphersuiteGroup,omitempty" xml:"CiphersuiteGroup,omitempty"`
	// Configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3528160969****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Whether to enable HTTP2. Default is enabled. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	Http2 *string `json:"Http2,omitempty" xml:"Http2,omitempty"`
	// Whether to enable HTTP3. Default is enabled. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	Http3 *string `json:"Http3,omitempty" xml:"Http3,omitempty"`
	// Whether to enable HTTPS. Default is enabled. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	Https *string `json:"Https,omitempty" xml:"Https,omitempty"`
	// Whether to enable OCSP. Default is disabled. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	OcspStapling *string `json:"OcspStapling,omitempty" xml:"OcspStapling,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter is not required when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: Set the value to true
	//
	// - Match specific requests: Set the value to a custom expression, for example: (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter is not required when adding a global configuration. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
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
	Sequence *int32  `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1231231221****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Whether to enable TLS1.0. Default is disabled. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	Tls10 *string `json:"Tls10,omitempty" xml:"Tls10,omitempty"`
	// Whether to enable TLS1.1. Default is disabled. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	Tls11 *string `json:"Tls11,omitempty" xml:"Tls11,omitempty"`
	// Whether to enable TLS1.2. Default is disabled. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	Tls12 *string `json:"Tls12,omitempty" xml:"Tls12,omitempty"`
	// Whether to enable TLS1.3. Default is disabled. Value range:
	//
	// - on: Enable.
	//
	// - off: Disable.
	//
	// example:
	//
	// on
	Tls13 *string `json:"Tls13,omitempty" xml:"Tls13,omitempty"`
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
