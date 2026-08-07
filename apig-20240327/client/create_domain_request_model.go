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
	SetDomainScope(v string) *CreateDomainRequest
	GetDomainScope() *string
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
	// The CA certificate identifier. This parameter is optional for Dedicated with HTTPS. This parameter is not allowed for Serverless and is not validated for Dedicated with HTTP.
	//
	// example:
	//
	// 1ef1da5f-38ed-69b3-****-037781890265
	CaCertIdentifier *string `json:"caCertIdentifier,omitempty" xml:"caCertIdentifier,omitempty"`
	// The certificate identifier. This parameter is required for Dedicated with HTTPS and must pass validation. This parameter is not allowed for Serverless and is not validated for Dedicated with HTTP.
	//
	// example:
	//
	// 1ef1da5f-38ed-69b3-****-037781890265
	CertIdentifier *string `json:"certIdentifier,omitempty" xml:"certIdentifier,omitempty"`
	// The client CA certificate. This parameter is conditionally required for Dedicated with HTTPS (when MTLSEnabled is set to true). This parameter is not allowed for Serverless and is not validated for Dedicated with HTTP.
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
	// The domain name scope. Valid values:
	//
	// - Dedicated: dedicated gateway domain name.
	//
	// - Serverless: Serverless gateway domain name.
	//
	// Default value: Dedicated.
	//
	// example:
	//
	// Dedicated
	DomainScope *string `json:"domainScope,omitempty" xml:"domainScope,omitempty"`
	// Specifies whether to enable forced HTTPS redirect for the HTTPS protocol type. This parameter is required for Serverless and for Dedicated with HTTPS. This parameter is not validated for Dedicated with HTTP.
	//
	// example:
	//
	// false
	ForceHttps *bool `json:"forceHttps,omitempty" xml:"forceHttps,omitempty"`
	// The gateway type. If not specified, the default value is API.
	//
	// example:
	//
	// API
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// The HTTP/2 setting. Valid values:
	//
	// - GlobalConfig: follows the global configuration.
	//
	// - Open: enabled.
	//
	// - Close: disabled.
	//
	// Default value: GlobalConfig. This setting is supported only for HTTPS domain names in the Dedicated scope.
	//
	// example:
	//
	// Open
	Http2Option *string `json:"http2Option,omitempty" xml:"http2Option,omitempty"`
	// Specifies whether to enable mTLS mutual authentication. This parameter is optional for Dedicated with HTTPS. When set to true, ClientCACert is required. This parameter is not allowed for Serverless.
	MTLSEnabled *bool `json:"mTLSEnabled,omitempty" xml:"mTLSEnabled,omitempty"`
	// The domain name. The name must be 1 to 128 characters in length, such as abc.com.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc.com
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The protocol type used by the domain name. Valid values: HTTP and HTTPS. This parameter is required for the Dedicated scope and is not allowed for the Serverless scope.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The [resource group ID](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-aekzhiv7derfweq
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The TLS cipher suite configuration, including the configuration type, cipher suite names, and supported TLS versions. This configuration is supported only for HTTPS domain names in the Dedicated scope.
	TlsCipherSuitesConfig *TlsCipherSuitesConfig `json:"tlsCipherSuitesConfig,omitempty" xml:"tlsCipherSuitesConfig,omitempty"`
	// The maximum TLS protocol version. This parameter is optional for Dedicated with HTTPS. If not specified, the value is derived from TlsMin. The value must be greater than or equal to TlsMin. This parameter is not allowed for Serverless.
	//
	// example:
	//
	// TLS1.3
	TlsMax *string `json:"tlsMax,omitempty" xml:"tlsMax,omitempty"`
	// The minimum TLS protocol version. This parameter is optional for Dedicated with HTTPS. If not specified, the default value is TLS 1.0. Valid values: TLS 1.0 to TLS 1.3, compatible with TLSv1.x. This parameter is not allowed for Serverless.
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

func (s *CreateDomainRequest) GetDomainScope() *string {
	return s.DomainScope
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

func (s *CreateDomainRequest) SetDomainScope(v string) *CreateDomainRequest {
	s.DomainScope = &v
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
	if s.TlsCipherSuitesConfig != nil {
		if err := s.TlsCipherSuitesConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
