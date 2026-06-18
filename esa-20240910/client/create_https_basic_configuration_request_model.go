// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpsBasicConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCiphersuite(v string) *CreateHttpsBasicConfigurationRequest
	GetCiphersuite() *string
	SetCiphersuiteGroup(v string) *CreateHttpsBasicConfigurationRequest
	GetCiphersuiteGroup() *string
	SetHttp2(v string) *CreateHttpsBasicConfigurationRequest
	GetHttp2() *string
	SetHttp3(v string) *CreateHttpsBasicConfigurationRequest
	GetHttp3() *string
	SetHttps(v string) *CreateHttpsBasicConfigurationRequest
	GetHttps() *string
	SetOcspStapling(v string) *CreateHttpsBasicConfigurationRequest
	GetOcspStapling() *string
	SetRule(v string) *CreateHttpsBasicConfigurationRequest
	GetRule() *string
	SetRuleEnable(v string) *CreateHttpsBasicConfigurationRequest
	GetRuleEnable() *string
	SetRuleName(v string) *CreateHttpsBasicConfigurationRequest
	GetRuleName() *string
	SetSequence(v int32) *CreateHttpsBasicConfigurationRequest
	GetSequence() *int32
	SetSiteId(v int64) *CreateHttpsBasicConfigurationRequest
	GetSiteId() *int64
	SetTls10(v string) *CreateHttpsBasicConfigurationRequest
	GetTls10() *string
	SetTls11(v string) *CreateHttpsBasicConfigurationRequest
	GetTls11() *string
	SetTls12(v string) *CreateHttpsBasicConfigurationRequest
	GetTls12() *string
	SetTls13(v string) *CreateHttpsBasicConfigurationRequest
	GetTls13() *string
}

type CreateHttpsBasicConfigurationRequest struct {
	// Custom cipher suite, specifying the specific encryption algorithms selected when CiphersuiteGroup is set to custom.
	//
	// example:
	//
	// TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
	Ciphersuite *string `json:"Ciphersuite,omitempty" xml:"Ciphersuite,omitempty"`
	// Cipher suite group. All cipher suites are used by default. Valid values:
	//
	// - all: all cipher suites.
	//
	// - strict: strong cipher suites.
	//
	// - custom: custom cipher suites.
	//
	// example:
	//
	// all
	CiphersuiteGroup *string `json:"CiphersuiteGroup,omitempty" xml:"CiphersuiteGroup,omitempty"`
	// Whether to enable HTTP/2. Enabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Http2 *string `json:"Http2,omitempty" xml:"Http2,omitempty"`
	// Whether to enable HTTP/3. Enabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Http3 *string `json:"Http3,omitempty" xml:"Http3,omitempty"`
	// Whether to enable HTTPS. Enabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Https *string `json:"Https,omitempty" xml:"Https,omitempty"`
	// Whether to enable OCSP stapling. Disabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	OcspStapling *string `json:"OcspStapling,omitempty" xml:"OcspStapling,omitempty"`
	// Rule content, using conditional expressions to match user requests. This parameter does not need to be set when adding a global configuration. There are two usage scenarios:
	//
	// - Match all incoming requests: set the value to true
	//
	// - Match specified requests: set the value to a custom expression, for example: (http.host eq "video.example.com")
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Rule switch. This parameter does not need to be set when adding a global configuration. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
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
	// Rule execution order. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Whether to enable TLS 1.0. Disabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Tls10 *string `json:"Tls10,omitempty" xml:"Tls10,omitempty"`
	// Whether to enable TLS 1.1. Enabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Tls11 *string `json:"Tls11,omitempty" xml:"Tls11,omitempty"`
	// Whether to enable TLS 1.2. Enabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Tls12 *string `json:"Tls12,omitempty" xml:"Tls12,omitempty"`
	// Whether to enable TLS 1.3. Enabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Tls13 *string `json:"Tls13,omitempty" xml:"Tls13,omitempty"`
}

func (s CreateHttpsBasicConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpsBasicConfigurationRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpsBasicConfigurationRequest) GetCiphersuite() *string {
	return s.Ciphersuite
}

func (s *CreateHttpsBasicConfigurationRequest) GetCiphersuiteGroup() *string {
	return s.CiphersuiteGroup
}

func (s *CreateHttpsBasicConfigurationRequest) GetHttp2() *string {
	return s.Http2
}

func (s *CreateHttpsBasicConfigurationRequest) GetHttp3() *string {
	return s.Http3
}

func (s *CreateHttpsBasicConfigurationRequest) GetHttps() *string {
	return s.Https
}

func (s *CreateHttpsBasicConfigurationRequest) GetOcspStapling() *string {
	return s.OcspStapling
}

func (s *CreateHttpsBasicConfigurationRequest) GetRule() *string {
	return s.Rule
}

func (s *CreateHttpsBasicConfigurationRequest) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *CreateHttpsBasicConfigurationRequest) GetRuleName() *string {
	return s.RuleName
}

func (s *CreateHttpsBasicConfigurationRequest) GetSequence() *int32 {
	return s.Sequence
}

func (s *CreateHttpsBasicConfigurationRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateHttpsBasicConfigurationRequest) GetTls10() *string {
	return s.Tls10
}

func (s *CreateHttpsBasicConfigurationRequest) GetTls11() *string {
	return s.Tls11
}

func (s *CreateHttpsBasicConfigurationRequest) GetTls12() *string {
	return s.Tls12
}

func (s *CreateHttpsBasicConfigurationRequest) GetTls13() *string {
	return s.Tls13
}

func (s *CreateHttpsBasicConfigurationRequest) SetCiphersuite(v string) *CreateHttpsBasicConfigurationRequest {
	s.Ciphersuite = &v
	return s
}

func (s *CreateHttpsBasicConfigurationRequest) SetCiphersuiteGroup(v string) *CreateHttpsBasicConfigurationRequest {
	s.CiphersuiteGroup = &v
	return s
}

func (s *CreateHttpsBasicConfigurationRequest) SetHttp2(v string) *CreateHttpsBasicConfigurationRequest {
	s.Http2 = &v
	return s
}

func (s *CreateHttpsBasicConfigurationRequest) SetHttp3(v string) *CreateHttpsBasicConfigurationRequest {
	s.Http3 = &v
	return s
}

func (s *CreateHttpsBasicConfigurationRequest) SetHttps(v string) *CreateHttpsBasicConfigurationRequest {
	s.Https = &v
	return s
}

func (s *CreateHttpsBasicConfigurationRequest) SetOcspStapling(v string) *CreateHttpsBasicConfigurationRequest {
	s.OcspStapling = &v
	return s
}

func (s *CreateHttpsBasicConfigurationRequest) SetRule(v string) *CreateHttpsBasicConfigurationRequest {
	s.Rule = &v
	return s
}

func (s *CreateHttpsBasicConfigurationRequest) SetRuleEnable(v string) *CreateHttpsBasicConfigurationRequest {
	s.RuleEnable = &v
	return s
}

func (s *CreateHttpsBasicConfigurationRequest) SetRuleName(v string) *CreateHttpsBasicConfigurationRequest {
	s.RuleName = &v
	return s
}

func (s *CreateHttpsBasicConfigurationRequest) SetSequence(v int32) *CreateHttpsBasicConfigurationRequest {
	s.Sequence = &v
	return s
}

func (s *CreateHttpsBasicConfigurationRequest) SetSiteId(v int64) *CreateHttpsBasicConfigurationRequest {
	s.SiteId = &v
	return s
}

func (s *CreateHttpsBasicConfigurationRequest) SetTls10(v string) *CreateHttpsBasicConfigurationRequest {
	s.Tls10 = &v
	return s
}

func (s *CreateHttpsBasicConfigurationRequest) SetTls11(v string) *CreateHttpsBasicConfigurationRequest {
	s.Tls11 = &v
	return s
}

func (s *CreateHttpsBasicConfigurationRequest) SetTls12(v string) *CreateHttpsBasicConfigurationRequest {
	s.Tls12 = &v
	return s
}

func (s *CreateHttpsBasicConfigurationRequest) SetTls13(v string) *CreateHttpsBasicConfigurationRequest {
	s.Tls13 = &v
	return s
}

func (s *CreateHttpsBasicConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
