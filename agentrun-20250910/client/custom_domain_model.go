// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomDomain interface {
	dara.Model
	String() string
	GoString() string
	SetCertConfig(v *CertConfig) *CustomDomain
	GetCertConfig() *CertConfig
	SetCreatedAt(v string) *CustomDomain
	GetCreatedAt() *string
	SetDescription(v string) *CustomDomain
	GetDescription() *string
	SetDomainName(v string) *CustomDomain
	GetDomainName() *string
	SetDomainType(v string) *CustomDomain
	GetDomainType() *string
	SetProtocol(v string) *CustomDomain
	GetProtocol() *string
	SetRouteConfig(v *RouteConfig) *CustomDomain
	GetRouteConfig() *RouteConfig
	SetTlsConfig(v *TLSConfig) *CustomDomain
	GetTlsConfig() *TLSConfig
	SetUpdatedAt(v string) *CustomDomain
	GetUpdatedAt() *string
}

type CustomDomain struct {
	// HTTPS 证书的信息。
	//
	// example:
	//
	// {}
	CertConfig *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 2025-01-10T10:30:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 域名。填写已在阿里云备案或接入备案的自定义域名名称。
	//
	// example:
	//
	// example.com
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	DomainType *string `json:"domainType,omitempty" xml:"domainType,omitempty"`
	// 域名支持的协议类型：● HTTP：仅支持 HTTP 协议。● HTTPS：仅支持 HTTPS 协议。● HTTP,HTTPS：支持 HTTP 及 HTTPS 协议。
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// 路由表：自定义域名访问时的 PATH 到 资源 的映射。
	//
	// example:
	//
	// {}
	RouteConfig *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
	// TLS 配置信息。
	//
	// example:
	//
	// {}
	TlsConfig *TLSConfig `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
	// 更新时间
	//
	// example:
	//
	// 2025-01-10T11:45:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s CustomDomain) String() string {
	return dara.Prettify(s)
}

func (s CustomDomain) GoString() string {
	return s.String()
}

func (s *CustomDomain) GetCertConfig() *CertConfig {
	return s.CertConfig
}

func (s *CustomDomain) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CustomDomain) GetDescription() *string {
	return s.Description
}

func (s *CustomDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *CustomDomain) GetDomainType() *string {
	return s.DomainType
}

func (s *CustomDomain) GetProtocol() *string {
	return s.Protocol
}

func (s *CustomDomain) GetRouteConfig() *RouteConfig {
	return s.RouteConfig
}

func (s *CustomDomain) GetTlsConfig() *TLSConfig {
	return s.TlsConfig
}

func (s *CustomDomain) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CustomDomain) SetCertConfig(v *CertConfig) *CustomDomain {
	s.CertConfig = v
	return s
}

func (s *CustomDomain) SetCreatedAt(v string) *CustomDomain {
	s.CreatedAt = &v
	return s
}

func (s *CustomDomain) SetDescription(v string) *CustomDomain {
	s.Description = &v
	return s
}

func (s *CustomDomain) SetDomainName(v string) *CustomDomain {
	s.DomainName = &v
	return s
}

func (s *CustomDomain) SetDomainType(v string) *CustomDomain {
	s.DomainType = &v
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

func (s *CustomDomain) SetTlsConfig(v *TLSConfig) *CustomDomain {
	s.TlsConfig = v
	return s
}

func (s *CustomDomain) SetUpdatedAt(v string) *CustomDomain {
	s.UpdatedAt = &v
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
	return nil
}
