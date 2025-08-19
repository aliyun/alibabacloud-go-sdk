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
	SetCreatedTime(v string) *CustomDomain
	GetCreatedTime() *string
	SetDomainName(v string) *CustomDomain
	GetDomainName() *string
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
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	// example:
	//
	// 2023-03-30
	ApiVersion *string     `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	CertConfig *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	// example:
	//
	// 2023-03-30T08:02:19Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// example:
	//
	// example.com
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	// example:
	//
	// 2023-03-30T08:02:19Z
	LastModifiedTime *string `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// example:
	//
	// HTTP
	Protocol    *string      `json:"protocol,omitempty" xml:"protocol,omitempty"`
	RouteConfig *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
	// example:
	//
	// 1
	SubdomainCount *string    `json:"subdomainCount,omitempty" xml:"subdomainCount,omitempty"`
	TlsConfig      *TLSConfig `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
	WafConfig      *WAFConfig `json:"wafConfig,omitempty" xml:"wafConfig,omitempty"`
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

func (s *CustomDomain) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *CustomDomain) GetDomainName() *string {
	return s.DomainName
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

func (s *CustomDomain) SetCreatedTime(v string) *CustomDomain {
	s.CreatedTime = &v
	return s
}

func (s *CustomDomain) SetDomainName(v string) *CustomDomain {
	s.DomainName = &v
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
	return dara.Validate(s)
}
