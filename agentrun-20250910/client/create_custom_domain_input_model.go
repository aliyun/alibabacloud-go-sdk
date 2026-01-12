// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomDomainInput interface {
	dara.Model
	String() string
	GoString() string
	SetCertConfig(v *CertConfig) *CreateCustomDomainInput
	GetCertConfig() *CertConfig
	SetDescription(v string) *CreateCustomDomainInput
	GetDescription() *string
	SetDomainName(v string) *CreateCustomDomainInput
	GetDomainName() *string
	SetDomainType(v string) *CreateCustomDomainInput
	GetDomainType() *string
	SetProtocol(v string) *CreateCustomDomainInput
	GetProtocol() *string
	SetRouteConfig(v *RouteConfig) *CreateCustomDomainInput
	GetRouteConfig() *RouteConfig
	SetTlsConfig(v *TLSConfig) *CreateCustomDomainInput
	GetTlsConfig() *TLSConfig
}

type CreateCustomDomainInput struct {
	// HTTPS 证书的信息。
	//
	// example:
	//
	// {}
	CertConfig *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 域名。填写已在阿里云备案或接入备案的自定义域名名称。
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"domainName,omitempty" xml:"domainName,omitempty"`
	DomainType *string `json:"domainType,omitempty" xml:"domainType,omitempty"`
	// 域名支持的协议类型：● HTTP：仅支持 HTTP 协议。● HTTPS：仅支持 HTTPS 协议。● HTTP,HTTPS：支持 HTTP 及 HTTPS 协议。
	//
	// This parameter is required.
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
}

func (s CreateCustomDomainInput) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomDomainInput) GoString() string {
	return s.String()
}

func (s *CreateCustomDomainInput) GetCertConfig() *CertConfig {
	return s.CertConfig
}

func (s *CreateCustomDomainInput) GetDescription() *string {
	return s.Description
}

func (s *CreateCustomDomainInput) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateCustomDomainInput) GetDomainType() *string {
	return s.DomainType
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

func (s *CreateCustomDomainInput) SetCertConfig(v *CertConfig) *CreateCustomDomainInput {
	s.CertConfig = v
	return s
}

func (s *CreateCustomDomainInput) SetDescription(v string) *CreateCustomDomainInput {
	s.Description = &v
	return s
}

func (s *CreateCustomDomainInput) SetDomainName(v string) *CreateCustomDomainInput {
	s.DomainName = &v
	return s
}

func (s *CreateCustomDomainInput) SetDomainType(v string) *CreateCustomDomainInput {
	s.DomainType = &v
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

func (s *CreateCustomDomainInput) Validate() error {
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
