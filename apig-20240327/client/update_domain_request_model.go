// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaCertIdentifier(v string) *UpdateDomainRequest
	GetCaCertIdentifier() *string
	SetCertIdentifier(v string) *UpdateDomainRequest
	GetCertIdentifier() *string
	SetClientCACert(v string) *UpdateDomainRequest
	GetClientCACert() *string
	SetDomainScope(v string) *UpdateDomainRequest
	GetDomainScope() *string
	SetForceHttps(v bool) *UpdateDomainRequest
	GetForceHttps() *bool
	SetHttp2Option(v string) *UpdateDomainRequest
	GetHttp2Option() *string
	SetMTLSEnabled(v bool) *UpdateDomainRequest
	GetMTLSEnabled() *bool
	SetProtocol(v string) *UpdateDomainRequest
	GetProtocol() *string
	SetTlsCipherSuitesConfig(v *TlsCipherSuitesConfig) *UpdateDomainRequest
	GetTlsCipherSuitesConfig() *TlsCipherSuitesConfig
	SetTlsMax(v string) *UpdateDomainRequest
	GetTlsMax() *string
	SetTlsMin(v string) *UpdateDomainRequest
	GetTlsMin() *string
}

type UpdateDomainRequest struct {
	// The CA certificate identifier.
	//
	// example:
	//
	// 1ef1da5f-38ed-69b3-****-037781890265
	CaCertIdentifier *string `json:"caCertIdentifier,omitempty" xml:"caCertIdentifier,omitempty"`
	// The certificate identifier.
	//
	// example:
	//
	// 1ef1da5f-38ed-69b3-****-037781890265
	CertIdentifier *string `json:"certIdentifier,omitempty" xml:"certIdentifier,omitempty"`
	// The client CA certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIFBTCCAu2gAwIBAgIUORLpYPGSFD1YOP6PMbE7Wd/mpTQwDQYJKoZIhvcNAQEL
	//
	// BQAwE************************************************2VwVOJ2gqX3
	//
	// YuGaxvIbDy0iQJ1GMerPRyzJTeVEtdIKT29u0PdFRr4KZWom35qX7G4=
	//
	// -----END CERTIFICATE-----
	ClientCACert *string `json:"clientCACert,omitempty" xml:"clientCACert,omitempty"`
	// The domain name scope.
	//
	// example:
	//
	// Dedicated
	DomainScope *string `json:"domainScope,omitempty" xml:"domainScope,omitempty"`
	// Specifies whether to enable forced HTTPS redirect when the HTTPS protocol type is set. When the protocol is HTTPS, forceHttps is required.
	//
	// example:
	//
	// false
	ForceHttps *bool `json:"forceHttps,omitempty" xml:"forceHttps,omitempty"`
	// The HTTP/2 settings.
	//
	// example:
	//
	// Open
	Http2Option *string `json:"http2Option,omitempty" xml:"http2Option,omitempty"`
	// Specifies whether to enable mTLS mutual authentication.
	MTLSEnabled *bool `json:"mTLSEnabled,omitempty" xml:"mTLSEnabled,omitempty"`
	// The protocol type supported by the domain name.
	//
	// - HTTP: Only the HTTP protocol is supported.
	//
	// - HTTPS: Only the HTTPS protocol is supported.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The cipher suite configuration.
	TlsCipherSuitesConfig *TlsCipherSuitesConfig `json:"tlsCipherSuitesConfig,omitempty" xml:"tlsCipherSuitesConfig,omitempty"`
	// The maximum TLS protocol version. TLS 1.3 is the maximum supported version.
	//
	// example:
	//
	// TLS 1.3
	TlsMax *string `json:"tlsMax,omitempty" xml:"tlsMax,omitempty"`
	// The minimum TLS protocol version. TLS 1.0 is the minimum supported version.
	//
	// example:
	//
	// TLS 1.0
	TlsMin *string `json:"tlsMin,omitempty" xml:"tlsMin,omitempty"`
}

func (s UpdateDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainRequest) GetCaCertIdentifier() *string {
	return s.CaCertIdentifier
}

func (s *UpdateDomainRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *UpdateDomainRequest) GetClientCACert() *string {
	return s.ClientCACert
}

func (s *UpdateDomainRequest) GetDomainScope() *string {
	return s.DomainScope
}

func (s *UpdateDomainRequest) GetForceHttps() *bool {
	return s.ForceHttps
}

func (s *UpdateDomainRequest) GetHttp2Option() *string {
	return s.Http2Option
}

func (s *UpdateDomainRequest) GetMTLSEnabled() *bool {
	return s.MTLSEnabled
}

func (s *UpdateDomainRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateDomainRequest) GetTlsCipherSuitesConfig() *TlsCipherSuitesConfig {
	return s.TlsCipherSuitesConfig
}

func (s *UpdateDomainRequest) GetTlsMax() *string {
	return s.TlsMax
}

func (s *UpdateDomainRequest) GetTlsMin() *string {
	return s.TlsMin
}

func (s *UpdateDomainRequest) SetCaCertIdentifier(v string) *UpdateDomainRequest {
	s.CaCertIdentifier = &v
	return s
}

func (s *UpdateDomainRequest) SetCertIdentifier(v string) *UpdateDomainRequest {
	s.CertIdentifier = &v
	return s
}

func (s *UpdateDomainRequest) SetClientCACert(v string) *UpdateDomainRequest {
	s.ClientCACert = &v
	return s
}

func (s *UpdateDomainRequest) SetDomainScope(v string) *UpdateDomainRequest {
	s.DomainScope = &v
	return s
}

func (s *UpdateDomainRequest) SetForceHttps(v bool) *UpdateDomainRequest {
	s.ForceHttps = &v
	return s
}

func (s *UpdateDomainRequest) SetHttp2Option(v string) *UpdateDomainRequest {
	s.Http2Option = &v
	return s
}

func (s *UpdateDomainRequest) SetMTLSEnabled(v bool) *UpdateDomainRequest {
	s.MTLSEnabled = &v
	return s
}

func (s *UpdateDomainRequest) SetProtocol(v string) *UpdateDomainRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateDomainRequest) SetTlsCipherSuitesConfig(v *TlsCipherSuitesConfig) *UpdateDomainRequest {
	s.TlsCipherSuitesConfig = v
	return s
}

func (s *UpdateDomainRequest) SetTlsMax(v string) *UpdateDomainRequest {
	s.TlsMax = &v
	return s
}

func (s *UpdateDomainRequest) SetTlsMin(v string) *UpdateDomainRequest {
	s.TlsMin = &v
	return s
}

func (s *UpdateDomainRequest) Validate() error {
	if s.TlsCipherSuitesConfig != nil {
		if err := s.TlsCipherSuitesConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
