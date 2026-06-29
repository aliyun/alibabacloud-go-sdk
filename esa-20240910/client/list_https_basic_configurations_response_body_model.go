// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpsBasicConfigurationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v []*ListHttpsBasicConfigurationsResponseBodyConfigs) *ListHttpsBasicConfigurationsResponseBody
	GetConfigs() []*ListHttpsBasicConfigurationsResponseBodyConfigs
	SetPageNumber(v int32) *ListHttpsBasicConfigurationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHttpsBasicConfigurationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListHttpsBasicConfigurationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListHttpsBasicConfigurationsResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListHttpsBasicConfigurationsResponseBody
	GetTotalPage() *int32
}

type ListHttpsBasicConfigurationsResponseBody struct {
	// The configurations returned.
	Configs []*ListHttpsBasicConfigurationsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Repeated"`
	// The current page number, which is the same as the PageNumber request parameter.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EDBD3EB3-97DA-5465-AEF5-8DCA5DC5E395
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListHttpsBasicConfigurationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHttpsBasicConfigurationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHttpsBasicConfigurationsResponseBody) GetConfigs() []*ListHttpsBasicConfigurationsResponseBodyConfigs {
	return s.Configs
}

func (s *ListHttpsBasicConfigurationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpsBasicConfigurationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpsBasicConfigurationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHttpsBasicConfigurationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHttpsBasicConfigurationsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListHttpsBasicConfigurationsResponseBody) SetConfigs(v []*ListHttpsBasicConfigurationsResponseBodyConfigs) *ListHttpsBasicConfigurationsResponseBody {
	s.Configs = v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBody) SetPageNumber(v int32) *ListHttpsBasicConfigurationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBody) SetPageSize(v int32) *ListHttpsBasicConfigurationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBody) SetRequestId(v string) *ListHttpsBasicConfigurationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBody) SetTotalCount(v int32) *ListHttpsBasicConfigurationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBody) SetTotalPage(v int32) *ListHttpsBasicConfigurationsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBody) Validate() error {
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

type ListHttpsBasicConfigurationsResponseBodyConfigs struct {
	// The custom cipher suite, which specifies the encryption algorithms selected when CiphersuiteGroup is set to custom.
	//
	// example:
	//
	// TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
	Ciphersuite *string `json:"Ciphersuite,omitempty" xml:"Ciphersuite,omitempty"`
	// The cipher suite group. All cipher suites are enabled by default. Valid values:
	//
	// - all: all cipher suites.
	//
	// - strict: strong cipher suites.
	//
	// - custom: custom cipher suites.
	//
	// example:
	//
	// strict
	CiphersuiteGroup *string `json:"CiphersuiteGroup,omitempty" xml:"CiphersuiteGroup,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 395386449776640
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration type. You can use this parameter to query global or rule configurations. Valid values:
	//
	// - global: queries global configurations.
	//
	// - rule: queries rule configurations.
	//
	// example:
	//
	// global
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// Specifies whether to enable HTTP/2. HTTP/2 is enabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Http2 *string `json:"Http2,omitempty" xml:"Http2,omitempty"`
	// Specifies whether to enable HTTP/3. HTTP/3 is enabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Http3 *string `json:"Http3,omitempty" xml:"Http3,omitempty"`
	// Specifies whether to enable HTTPS. HTTPS is enabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Https *string `json:"Https,omitempty" xml:"Https,omitempty"`
	// Specifies whether to enable OCSP stapling. OCSP stapling is disabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	OcspStapling *string `json:"OcspStapling,omitempty" xml:"OcspStapling,omitempty"`
	// The rule content, which uses a conditional expression to match user requests. You do not need to set this parameter when adding a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: set the value to true.
	//
	// - Match specified requests: set the value to a custom expression, such as (http.host eq \\"video.example.com\\").
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. You do not need to set this parameter when adding a global configuration. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. You do not need to set this parameter when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The rule execution order. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Specifies whether to enable TLS 1.0. TLS 1.0 is disabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Tls10 *string `json:"Tls10,omitempty" xml:"Tls10,omitempty"`
	// Specifies whether to enable TLS 1.1. TLS 1.1 is disabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Tls11 *string `json:"Tls11,omitempty" xml:"Tls11,omitempty"`
	// Specifies whether to enable TLS 1.2. TLS 1.2 is disabled by default. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	Tls12 *string `json:"Tls12,omitempty" xml:"Tls12,omitempty"`
	// Specifies whether to enable TLS 1.3. TLS 1.3 is disabled by default. Valid values:
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

func (s ListHttpsBasicConfigurationsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListHttpsBasicConfigurationsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetCiphersuite() *string {
	return s.Ciphersuite
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetCiphersuiteGroup() *string {
	return s.CiphersuiteGroup
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetConfigType() *string {
	return s.ConfigType
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetHttp2() *string {
	return s.Http2
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetHttp3() *string {
	return s.Http3
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetHttps() *string {
	return s.Https
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetOcspStapling() *string {
	return s.OcspStapling
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetRule() *string {
	return s.Rule
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetRuleName() *string {
	return s.RuleName
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetSequence() *int32 {
	return s.Sequence
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetTls10() *string {
	return s.Tls10
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetTls11() *string {
	return s.Tls11
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetTls12() *string {
	return s.Tls12
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) GetTls13() *string {
	return s.Tls13
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetCiphersuite(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Ciphersuite = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetCiphersuiteGroup(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.CiphersuiteGroup = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetConfigId(v int64) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.ConfigId = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetConfigType(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.ConfigType = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetHttp2(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Http2 = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetHttp3(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Http3 = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetHttps(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Https = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetOcspStapling(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.OcspStapling = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetRule(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Rule = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetRuleEnable(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.RuleEnable = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetRuleName(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.RuleName = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetSequence(v int32) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Sequence = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetTls10(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Tls10 = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetTls11(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Tls11 = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetTls12(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Tls12 = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) SetTls13(v string) *ListHttpsBasicConfigurationsResponseBodyConfigs {
	s.Tls13 = &v
	return s
}

func (s *ListHttpsBasicConfigurationsResponseBodyConfigs) Validate() error {
	return dara.Validate(s)
}
