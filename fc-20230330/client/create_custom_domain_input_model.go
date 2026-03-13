// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomDomainInput interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v *AuthConfig) *CreateCustomDomainInput
	GetAuthConfig() *AuthConfig
	SetCertConfig(v *CertConfig) *CreateCustomDomainInput
	GetCertConfig() *CertConfig
	SetCorsConfig(v *CORSConfig) *CreateCustomDomainInput
	GetCorsConfig() *CORSConfig
	SetDomainName(v string) *CreateCustomDomainInput
	GetDomainName() *string
	SetProtocol(v string) *CreateCustomDomainInput
	GetProtocol() *string
	SetRouteConfig(v *RouteConfig) *CreateCustomDomainInput
	GetRouteConfig() *RouteConfig
	SetTlsConfig(v *TLSConfig) *CreateCustomDomainInput
	GetTlsConfig() *TLSConfig
	SetWafConfig(v *WAFConfig) *CreateCustomDomainInput
	GetWafConfig() *WAFConfig
}

type CreateCustomDomainInput struct {
	// The configuration of permission authentication.
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// The configuration of the HTTPS certificate.
	CertConfig *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	CorsConfig *CORSConfig `json:"corsConfig,omitempty" xml:"corsConfig,omitempty"`
	// The domain name. Enter a custom domain name that has obtained an Internet content provider (ICP) filing in the Alibaba Cloud ICP Filing system, or a custom domain name whose ICP filing information includes Alibaba Cloud as a service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// The protocol type that is supported by the custom domain name. Valid values: HTTP HTTPS HTTP,HTTPS
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The route table that maps paths to functions when the function is invoked by using the custom domain name.
	RouteConfig *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
	// The Transport Layer Security (TLS) configuration.
	TlsConfig *TLSConfig `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
	// The Web Application Firewall (WAF) configuration.
	WafConfig *WAFConfig `json:"wafConfig,omitempty" xml:"wafConfig,omitempty"`
}

func (s CreateCustomDomainInput) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomDomainInput) GoString() string {
	return s.String()
}

func (s *CreateCustomDomainInput) GetAuthConfig() *AuthConfig {
	return s.AuthConfig
}

func (s *CreateCustomDomainInput) GetCertConfig() *CertConfig {
	return s.CertConfig
}

func (s *CreateCustomDomainInput) GetCorsConfig() *CORSConfig {
	return s.CorsConfig
}

func (s *CreateCustomDomainInput) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateCustomDomainInput) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateCustomDomainInput) GetRouteConfig() *RouteConfig {
	return s.RouteConfig
}

func (s *CreateCustomDomainInput) GetTlsConfig() *TLSConfig {
	return s.TlsConfig
}

func (s *CreateCustomDomainInput) GetWafConfig() *WAFConfig {
	return s.WafConfig
}

func (s *CreateCustomDomainInput) SetAuthConfig(v *AuthConfig) *CreateCustomDomainInput {
	s.AuthConfig = v
	return s
}

func (s *CreateCustomDomainInput) SetCertConfig(v *CertConfig) *CreateCustomDomainInput {
	s.CertConfig = v
	return s
}

func (s *CreateCustomDomainInput) SetCorsConfig(v *CORSConfig) *CreateCustomDomainInput {
	s.CorsConfig = v
	return s
}

func (s *CreateCustomDomainInput) SetDomainName(v string) *CreateCustomDomainInput {
	s.DomainName = &v
	return s
}

func (s *CreateCustomDomainInput) SetProtocol(v string) *CreateCustomDomainInput {
	s.Protocol = &v
	return s
}

func (s *CreateCustomDomainInput) SetRouteConfig(v *RouteConfig) *CreateCustomDomainInput {
	s.RouteConfig = v
	return s
}

func (s *CreateCustomDomainInput) SetTlsConfig(v *TLSConfig) *CreateCustomDomainInput {
	s.TlsConfig = v
	return s
}

func (s *CreateCustomDomainInput) SetWafConfig(v *WAFConfig) *CreateCustomDomainInput {
	s.WafConfig = v
	return s
}

func (s *CreateCustomDomainInput) Validate() error {
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
