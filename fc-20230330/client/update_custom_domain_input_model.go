// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomDomainInput interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v *AuthConfig) *UpdateCustomDomainInput
	GetAuthConfig() *AuthConfig
	SetCertConfig(v *CertConfig) *UpdateCustomDomainInput
	GetCertConfig() *CertConfig
	SetCorsConfig(v *CORSConfig) *UpdateCustomDomainInput
	GetCorsConfig() *CORSConfig
	SetProtocol(v string) *UpdateCustomDomainInput
	GetProtocol() *string
	SetRouteConfig(v *RouteConfig) *UpdateCustomDomainInput
	GetRouteConfig() *RouteConfig
	SetTlsConfig(v *TLSConfig) *UpdateCustomDomainInput
	GetTlsConfig() *TLSConfig
	SetWafConfig(v *WAFConfig) *UpdateCustomDomainInput
	GetWafConfig() *WAFConfig
}

type UpdateCustomDomainInput struct {
	// The authentication configuration.
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// Information about the HTTPS certificate.
	CertConfig *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	CorsConfig *CORSConfig `json:"corsConfig,omitempty" xml:"corsConfig,omitempty"`
	// The protocol type that the domain name supports. \\`HTTP\\`: supports only the HTTP protocol. \\`HTTPS\\`: supports only the HTTPS protocol. \\`HTTP,HTTPS\\`: supports both HTTP and HTTPS protocols.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The route table that maps the access paths of the custom domain name to functions.
	RouteConfig *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
	// The TLS configuration.
	TlsConfig *TLSConfig `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
	// The Web Application Firewall (WAF) configuration.
	WafConfig *WAFConfig `json:"wafConfig,omitempty" xml:"wafConfig,omitempty"`
}

func (s UpdateCustomDomainInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomDomainInput) GoString() string {
	return s.String()
}

func (s *UpdateCustomDomainInput) GetAuthConfig() *AuthConfig {
	return s.AuthConfig
}

func (s *UpdateCustomDomainInput) GetCertConfig() *CertConfig {
	return s.CertConfig
}

func (s *UpdateCustomDomainInput) GetCorsConfig() *CORSConfig {
	return s.CorsConfig
}

func (s *UpdateCustomDomainInput) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateCustomDomainInput) GetRouteConfig() *RouteConfig {
	return s.RouteConfig
}

func (s *UpdateCustomDomainInput) GetTlsConfig() *TLSConfig {
	return s.TlsConfig
}

func (s *UpdateCustomDomainInput) GetWafConfig() *WAFConfig {
	return s.WafConfig
}

func (s *UpdateCustomDomainInput) SetAuthConfig(v *AuthConfig) *UpdateCustomDomainInput {
	s.AuthConfig = v
	return s
}

func (s *UpdateCustomDomainInput) SetCertConfig(v *CertConfig) *UpdateCustomDomainInput {
	s.CertConfig = v
	return s
}

func (s *UpdateCustomDomainInput) SetCorsConfig(v *CORSConfig) *UpdateCustomDomainInput {
	s.CorsConfig = v
	return s
}

func (s *UpdateCustomDomainInput) SetProtocol(v string) *UpdateCustomDomainInput {
	s.Protocol = &v
	return s
}

func (s *UpdateCustomDomainInput) SetRouteConfig(v *RouteConfig) *UpdateCustomDomainInput {
	s.RouteConfig = v
	return s
}

func (s *UpdateCustomDomainInput) SetTlsConfig(v *TLSConfig) *UpdateCustomDomainInput {
	s.TlsConfig = v
	return s
}

func (s *UpdateCustomDomainInput) SetWafConfig(v *WAFConfig) *UpdateCustomDomainInput {
	s.WafConfig = v
	return s
}

func (s *UpdateCustomDomainInput) Validate() error {
	if s.AuthConfig != nil {
		if err := s.AuthConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CertConfig != nil {
		if err := s.CertConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CorsConfig != nil {
		if err := s.CorsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RouteConfig != nil {
		if err := s.RouteConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TlsConfig != nil {
		if err := s.TlsConfig.Validate(); err != nil {
			return err
		}
	}
	if s.WafConfig != nil {
		if err := s.WafConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
