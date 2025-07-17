// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaCertIdentifier(v string) *CreateDomainRequest
	GetCaCertIdentifier() *string
	SetCertIdentifier(v string) *CreateDomainRequest
	GetCertIdentifier() *string
	SetClientCACert(v string) *CreateDomainRequest
	GetClientCACert() *string
	SetForceHttps(v bool) *CreateDomainRequest
	GetForceHttps() *bool
	SetGatewayType(v string) *CreateDomainRequest
	GetGatewayType() *string
	SetHttp2Option(v string) *CreateDomainRequest
	GetHttp2Option() *string
	SetMTLSEnabled(v bool) *CreateDomainRequest
	GetMTLSEnabled() *bool
	SetName(v string) *CreateDomainRequest
	GetName() *string
	SetProtocol(v string) *CreateDomainRequest
	GetProtocol() *string
	SetResourceGroupId(v string) *CreateDomainRequest
	GetResourceGroupId() *string
	SetTlsCipherSuitesConfig(v *TlsCipherSuitesConfig) *CreateDomainRequest
	GetTlsCipherSuitesConfig() *TlsCipherSuitesConfig
	SetTlsMax(v string) *CreateDomainRequest
	GetTlsMax() *string
	SetTlsMin(v string) *CreateDomainRequest
	GetTlsMin() *string
}

type CreateDomainRequest struct {
	// The CA certificate ID.
	//
	// example:
	//
	// 1ef1da5f-38ed-69b3-****-037781890265
	CaCertIdentifier *string `json:"caCertIdentifier,omitempty" xml:"caCertIdentifier,omitempty"`
	// The certificate ID.
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
	// Specifies whether to enable forcible HTTPS redirection.
	//
	// example:
	//
	// false
	ForceHttps  *bool   `json:"forceHttps,omitempty" xml:"forceHttps,omitempty"`
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// The HTTP/2 configuration.
	//
	// Valid values:
	//
	// 	- GlobalConfig
	//
	// 	- Close
	//
	// 	- Open
	//
	// example:
	//
	// Open
	Http2Option *string `json:"http2Option,omitempty" xml:"http2Option,omitempty"`
	// Specifies whether to enable mutual authentication.
	MTLSEnabled *bool `json:"mTLSEnabled,omitempty" xml:"mTLSEnabled,omitempty"`
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc.com
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The protocol type supported by the domain name.
	//
	// 	- HTTP: Only HTTP is supported.
	//
	// 	- HTTPS: Only HTTPS is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The [resource group ID](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-aekzoiafjtr7zyq
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The cipher suite configuration.
	TlsCipherSuitesConfig *TlsCipherSuitesConfig `json:"tlsCipherSuitesConfig,omitempty" xml:"tlsCipherSuitesConfig,omitempty"`
	// The maximum version of the TLS protocol. Up to TLS 1.3 is supported.
	//
	// example:
	//
	// TLS1.3
	TlsMax *string `json:"tlsMax,omitempty" xml:"tlsMax,omitempty"`
	// The minimum version of the TLS protocol. Down to TLS 1.0 is supported.
	//
	// example:
	//
	// TLS1.0
	TlsMin *string `json:"tlsMin,omitempty" xml:"tlsMin,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) GetCaCertIdentifier() *string {
	return s.CaCertIdentifier
}

func (s *CreateDomainRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *CreateDomainRequest) GetClientCACert() *string {
	return s.ClientCACert
}

func (s *CreateDomainRequest) GetForceHttps() *bool {
	return s.ForceHttps
}

func (s *CreateDomainRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *CreateDomainRequest) GetHttp2Option() *string {
	return s.Http2Option
}

func (s *CreateDomainRequest) GetMTLSEnabled() *bool {
	return s.MTLSEnabled
}

func (s *CreateDomainRequest) GetName() *string {
	return s.Name
}

func (s *CreateDomainRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateDomainRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateDomainRequest) GetTlsCipherSuitesConfig() *TlsCipherSuitesConfig {
	return s.TlsCipherSuitesConfig
}

func (s *CreateDomainRequest) GetTlsMax() *string {
	return s.TlsMax
}

func (s *CreateDomainRequest) GetTlsMin() *string {
	return s.TlsMin
}

func (s *CreateDomainRequest) SetCaCertIdentifier(v string) *CreateDomainRequest {
	s.CaCertIdentifier = &v
	return s
}

func (s *CreateDomainRequest) SetCertIdentifier(v string) *CreateDomainRequest {
	s.CertIdentifier = &v
	return s
}

func (s *CreateDomainRequest) SetClientCACert(v string) *CreateDomainRequest {
	s.ClientCACert = &v
	return s
}

func (s *CreateDomainRequest) SetForceHttps(v bool) *CreateDomainRequest {
	s.ForceHttps = &v
	return s
}

func (s *CreateDomainRequest) SetGatewayType(v string) *CreateDomainRequest {
	s.GatewayType = &v
	return s
}

func (s *CreateDomainRequest) SetHttp2Option(v string) *CreateDomainRequest {
	s.Http2Option = &v
	return s
}

func (s *CreateDomainRequest) SetMTLSEnabled(v bool) *CreateDomainRequest {
	s.MTLSEnabled = &v
	return s
}

func (s *CreateDomainRequest) SetName(v string) *CreateDomainRequest {
	s.Name = &v
	return s
}

func (s *CreateDomainRequest) SetProtocol(v string) *CreateDomainRequest {
	s.Protocol = &v
	return s
}

func (s *CreateDomainRequest) SetResourceGroupId(v string) *CreateDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateDomainRequest) SetTlsCipherSuitesConfig(v *TlsCipherSuitesConfig) *CreateDomainRequest {
	s.TlsCipherSuitesConfig = v
	return s
}

func (s *CreateDomainRequest) SetTlsMax(v string) *CreateDomainRequest {
	s.TlsMax = &v
	return s
}

func (s *CreateDomainRequest) SetTlsMin(v string) *CreateDomainRequest {
	s.TlsMin = &v
	return s
}

func (s *CreateDomainRequest) Validate() error {
	return dara.Validate(s)
}
