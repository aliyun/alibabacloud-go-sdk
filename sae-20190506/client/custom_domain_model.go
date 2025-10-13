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
	SetCertConfig(v *CertConfig) *CustomDomain
	GetCertConfig() *CertConfig
	SetCreatedTime(v string) *CustomDomain
	GetCreatedTime() *string
	SetDomainName(v string) *CustomDomain
	GetDomainName() *string
	SetKeepFullPath(v bool) *CustomDomain
	GetKeepFullPath() *bool
	SetLastModifiedTime(v string) *CustomDomain
	GetLastModifiedTime() *string
	SetNamespaceID(v string) *CustomDomain
	GetNamespaceID() *string
	SetProtocol(v string) *CustomDomain
	GetProtocol() *string
	SetRequestId(v string) *CustomDomain
	GetRequestId() *string
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
	AccountId        *string      `json:"accountId,omitempty" xml:"accountId,omitempty"`
	ApiVersion       *string      `json:"apiVersion,omitempty" xml:"apiVersion,omitempty"`
	CertConfig       *CertConfig  `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	CreatedTime      *string      `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	DomainName       *string      `json:"domainName,omitempty" xml:"domainName,omitempty"`
	KeepFullPath     *bool        `json:"keepFullPath,omitempty" xml:"keepFullPath,omitempty"`
	LastModifiedTime *string      `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	NamespaceID      *string      `json:"namespaceID,omitempty" xml:"namespaceID,omitempty"`
	Protocol         *string      `json:"protocol,omitempty" xml:"protocol,omitempty"`
	RequestId        *string      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	RouteConfig      *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
	SubdomainCount   *string      `json:"subdomainCount,omitempty" xml:"subdomainCount,omitempty"`
	TlsConfig        *TLSConfig   `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
	WafConfig        *WAFConfig   `json:"wafConfig,omitempty" xml:"wafConfig,omitempty"`
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

func (s *CustomDomain) GetCertConfig() *CertConfig {
	return s.CertConfig
}

func (s *CustomDomain) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *CustomDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *CustomDomain) GetKeepFullPath() *bool {
	return s.KeepFullPath
}

func (s *CustomDomain) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *CustomDomain) GetNamespaceID() *string {
	return s.NamespaceID
}

func (s *CustomDomain) GetProtocol() *string {
	return s.Protocol
}

func (s *CustomDomain) GetRequestId() *string {
	return s.RequestId
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

func (s *CustomDomain) SetKeepFullPath(v bool) *CustomDomain {
	s.KeepFullPath = &v
	return s
}

func (s *CustomDomain) SetLastModifiedTime(v string) *CustomDomain {
	s.LastModifiedTime = &v
	return s
}

func (s *CustomDomain) SetNamespaceID(v string) *CustomDomain {
	s.NamespaceID = &v
	return s
}

func (s *CustomDomain) SetProtocol(v string) *CustomDomain {
	s.Protocol = &v
	return s
}

func (s *CustomDomain) SetRequestId(v string) *CustomDomain {
	s.RequestId = &v
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
	if s.CertConfig != nil {
		if err := s.CertConfig.Validate(); err != nil {
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
