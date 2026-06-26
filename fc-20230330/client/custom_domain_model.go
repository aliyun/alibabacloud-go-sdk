// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomDomain interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CustomDomain
	GetAccountId() *string
	SetApiVersion(v string) *CustomDomain
	GetApiVersion() *string
	SetAuthConfig(v *AuthConfig) *CustomDomain
	GetAuthConfig() *AuthConfig
	SetCertConfig(v *CertConfig) *CustomDomain
	GetCertConfig() *CertConfig
	SetCorsConfig(v *CORSConfig) *CustomDomain
	GetCorsConfig() *CORSConfig
	SetCreatedTime(v string) *CustomDomain
	GetCreatedTime() *string
	SetDomainName(v string) *CustomDomain
	GetDomainName() *string
	SetIsE2B(v bool) *CustomDomain
	GetIsE2B() *bool
	SetLastModifiedTime(v string) *CustomDomain
	GetLastModifiedTime() *string
	SetProtocol(v string) *CustomDomain
	GetProtocol() *string
	SetRouteConfig(v *RouteConfig) *CustomDomain
	GetRouteConfig() *RouteConfig
	SetSubdomainCount(v string) *CustomDomain
	GetSubdomainCount() *string
	SetTlsConfig(v *TLSConfig) *CustomDomain
	GetTlsConfig() *TLSConfig
	SetWafConfig(v *WAFConfig) *CustomDomain
	GetWafConfig() *WAFConfig
}

type CustomDomain struct {
	// The ID of the Alibaba Cloud account (primary account).
	//
	// example:
	//
	// 186851234023****
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// The API version of Function Compute.
	//
	// example:
	//
	// 2023-03-30
	ApiVersion *string `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	// The authentication configuration.
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	// The HTTPS certificate configuration.
	CertConfig *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	// The cross-origin resource sharing (CORS) configuration.
	CorsConfig *CORSConfig `json:"corsConfig,omitempty" xml:"corsConfig,omitempty"`
	// The time when the custom domain was created.
	//
	// example:
	//
	// 2023-03-30T08:02:19Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	IsE2B      *bool   `json:"isE2B,omitempty" xml:"isE2B,omitempty"`
	// The time when the custom domain was last modified.
	//
	// example:
	//
	// 2023-03-30T08:02:19Z
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// The protocols that are supported by the domain name. Valid values: HTTP (HTTP only), HTTPS (HTTPS only), and HTTP,HTTPS (both HTTP and HTTPS).
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The route table that maps paths to functions.
	RouteConfig *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
	// The number of subdomains.
	//
	// example:
	//
	// 1
	SubdomainCount *string `json:"subdomainCount,omitempty" xml:"subdomainCount,omitempty"`
	// The TLS configuration.
	TlsConfig *TLSConfig `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
	// The Web Application Firewall (WAF) aconfiguration.
	WafConfig *WAFConfig `json:"wafConfig,omitempty" xml:"wafConfig,omitempty"`
}

func (s CustomDomain) String() string {
	return dara.Prettify(s)
}

func (s CustomDomain) GoString() string {
	return s.String()
}

func (s *CustomDomain) GetAccountId() *string {
	return s.AccountId
}

func (s *CustomDomain) GetApiVersion() *string {
	return s.ApiVersion
}

func (s *CustomDomain) GetAuthConfig() *AuthConfig {
	return s.AuthConfig
}

func (s *CustomDomain) GetCertConfig() *CertConfig {
	return s.CertConfig
}

func (s *CustomDomain) GetCorsConfig() *CORSConfig {
	return s.CorsConfig
}

func (s *CustomDomain) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *CustomDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *CustomDomain) GetIsE2B() *bool {
	return s.IsE2B
}

func (s *CustomDomain) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *CustomDomain) GetProtocol() *string {
	return s.Protocol
}

func (s *CustomDomain) GetRouteConfig() *RouteConfig {
	return s.RouteConfig
}

func (s *CustomDomain) GetSubdomainCount() *string {
	return s.SubdomainCount
}

func (s *CustomDomain) GetTlsConfig() *TLSConfig {
	return s.TlsConfig
}

func (s *CustomDomain) GetWafConfig() *WAFConfig {
	return s.WafConfig
}

func (s *CustomDomain) SetAccountId(v string) *CustomDomain {
	s.AccountId = &v
	return s
}

func (s *CustomDomain) SetApiVersion(v string) *CustomDomain {
	s.ApiVersion = &v
	return s
}

func (s *CustomDomain) SetAuthConfig(v *AuthConfig) *CustomDomain {
	s.AuthConfig = v
	return s
}

func (s *CustomDomain) SetCertConfig(v *CertConfig) *CustomDomain {
	s.CertConfig = v
	return s
}

func (s *CustomDomain) SetCorsConfig(v *CORSConfig) *CustomDomain {
	s.CorsConfig = v
	return s
}

func (s *CustomDomain) SetCreatedTime(v string) *CustomDomain {
	s.CreatedTime = &v
	return s
}

func (s *CustomDomain) SetDomainName(v string) *CustomDomain {
	s.DomainName = &v
	return s
}

func (s *CustomDomain) SetIsE2B(v bool) *CustomDomain {
	s.IsE2B = &v
	return s
}

func (s *CustomDomain) SetLastModifiedTime(v string) *CustomDomain {
	s.LastModifiedTime = &v
	return s
}

func (s *CustomDomain) SetProtocol(v string) *CustomDomain {
	s.Protocol = &v
	return s
}

func (s *CustomDomain) SetRouteConfig(v *RouteConfig) *CustomDomain {
	s.RouteConfig = v
	return s
}

func (s *CustomDomain) SetSubdomainCount(v string) *CustomDomain {
	s.SubdomainCount = &v
	return s
}

func (s *CustomDomain) SetTlsConfig(v *TLSConfig) *CustomDomain {
	s.TlsConfig = v
	return s
}

func (s *CustomDomain) SetWafConfig(v *WAFConfig) *CustomDomain {
	s.WafConfig = v
	return s
}

func (s *CustomDomain) Validate() error {
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
