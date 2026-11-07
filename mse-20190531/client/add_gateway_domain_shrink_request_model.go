// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayDomainShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddGatewayDomainShrinkRequest
	GetAcceptLanguage() *string
	SetCaCertIdentifier(v string) *AddGatewayDomainShrinkRequest
	GetCaCertIdentifier() *string
	SetCertIdentifier(v string) *AddGatewayDomainShrinkRequest
	GetCertIdentifier() *string
	SetGatewayUniqueId(v string) *AddGatewayDomainShrinkRequest
	GetGatewayUniqueId() *string
	SetHttp2(v string) *AddGatewayDomainShrinkRequest
	GetHttp2() *string
	SetMtlsEnabled(v bool) *AddGatewayDomainShrinkRequest
	GetMtlsEnabled() *bool
	SetMustHttps(v bool) *AddGatewayDomainShrinkRequest
	GetMustHttps() *bool
	SetName(v string) *AddGatewayDomainShrinkRequest
	GetName() *string
	SetProtocol(v string) *AddGatewayDomainShrinkRequest
	GetProtocol() *string
	SetTlsCipherSuitesConfigJSONShrink(v string) *AddGatewayDomainShrinkRequest
	GetTlsCipherSuitesConfigJSONShrink() *string
	SetTlsMax(v string) *AddGatewayDomainShrinkRequest
	GetTlsMax() *string
	SetTlsMin(v string) *AddGatewayDomainShrinkRequest
	GetTlsMin() *string
}

type AddGatewayDomainShrinkRequest struct {
	// The language of the response. Valid values:
	//
	// - zh: Chinese. This is the default value.
	//
	// - en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The content of the CA certificate.
	CaCertIdentifier *string `json:"CaCertIdentifier,omitempty" xml:"CaCertIdentifier,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// 6828169-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-86575c0bc9f04ecfbacb92b8e392****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// Specifies whether to enable HTTP/2. Valid values:
	//
	// - `open`: Enables HTTP/2.
	//
	// - `close`: Disables HTTP/2.
	//
	// - `globalConfig`: Uses the global configuration.
	//
	// example:
	//
	// close
	Http2 *string `json:"Http2,omitempty" xml:"Http2,omitempty"`
	// Specifies whether to enable mTLS.
	MtlsEnabled *bool `json:"MtlsEnabled,omitempty" xml:"MtlsEnabled,omitempty"`
	// Specifies whether to enable HTTPS.
	//
	// example:
	//
	// true
	MustHttps *bool `json:"MustHttps,omitempty" xml:"MustHttps,omitempty"`
	// The domain name.
	//
	// example:
	//
	// test.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The protocol type. Valid values:
	//
	// - `HTTP`
	//
	// - `HTTPS`
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The TLS cipher suite configuration.
	TlsCipherSuitesConfigJSONShrink *string `json:"TlsCipherSuitesConfigJSON,omitempty" xml:"TlsCipherSuitesConfigJSON,omitempty"`
	// The maximum TLS version.
	//
	// example:
	//
	// TLS 1.3
	TlsMax *string `json:"TlsMax,omitempty" xml:"TlsMax,omitempty"`
	// The minimum TLS version.
	//
	// example:
	//
	// TLS 1.0
	TlsMin *string `json:"TlsMin,omitempty" xml:"TlsMin,omitempty"`
}

func (s AddGatewayDomainShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayDomainShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddGatewayDomainShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddGatewayDomainShrinkRequest) GetCaCertIdentifier() *string {
	return s.CaCertIdentifier
}

func (s *AddGatewayDomainShrinkRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *AddGatewayDomainShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddGatewayDomainShrinkRequest) GetHttp2() *string {
	return s.Http2
}

func (s *AddGatewayDomainShrinkRequest) GetMtlsEnabled() *bool {
	return s.MtlsEnabled
}

func (s *AddGatewayDomainShrinkRequest) GetMustHttps() *bool {
	return s.MustHttps
}

func (s *AddGatewayDomainShrinkRequest) GetName() *string {
	return s.Name
}

func (s *AddGatewayDomainShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *AddGatewayDomainShrinkRequest) GetTlsCipherSuitesConfigJSONShrink() *string {
	return s.TlsCipherSuitesConfigJSONShrink
}

func (s *AddGatewayDomainShrinkRequest) GetTlsMax() *string {
	return s.TlsMax
}

func (s *AddGatewayDomainShrinkRequest) GetTlsMin() *string {
	return s.TlsMin
}

func (s *AddGatewayDomainShrinkRequest) SetAcceptLanguage(v string) *AddGatewayDomainShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddGatewayDomainShrinkRequest) SetCaCertIdentifier(v string) *AddGatewayDomainShrinkRequest {
	s.CaCertIdentifier = &v
	return s
}

func (s *AddGatewayDomainShrinkRequest) SetCertIdentifier(v string) *AddGatewayDomainShrinkRequest {
	s.CertIdentifier = &v
	return s
}

func (s *AddGatewayDomainShrinkRequest) SetGatewayUniqueId(v string) *AddGatewayDomainShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddGatewayDomainShrinkRequest) SetHttp2(v string) *AddGatewayDomainShrinkRequest {
	s.Http2 = &v
	return s
}

func (s *AddGatewayDomainShrinkRequest) SetMtlsEnabled(v bool) *AddGatewayDomainShrinkRequest {
	s.MtlsEnabled = &v
	return s
}

func (s *AddGatewayDomainShrinkRequest) SetMustHttps(v bool) *AddGatewayDomainShrinkRequest {
	s.MustHttps = &v
	return s
}

func (s *AddGatewayDomainShrinkRequest) SetName(v string) *AddGatewayDomainShrinkRequest {
	s.Name = &v
	return s
}

func (s *AddGatewayDomainShrinkRequest) SetProtocol(v string) *AddGatewayDomainShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *AddGatewayDomainShrinkRequest) SetTlsCipherSuitesConfigJSONShrink(v string) *AddGatewayDomainShrinkRequest {
	s.TlsCipherSuitesConfigJSONShrink = &v
	return s
}

func (s *AddGatewayDomainShrinkRequest) SetTlsMax(v string) *AddGatewayDomainShrinkRequest {
	s.TlsMax = &v
	return s
}

func (s *AddGatewayDomainShrinkRequest) SetTlsMin(v string) *AddGatewayDomainShrinkRequest {
	s.TlsMin = &v
	return s
}

func (s *AddGatewayDomainShrinkRequest) Validate() error {
	return dara.Validate(s)
}
