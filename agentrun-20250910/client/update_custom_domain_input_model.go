// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomDomainInput interface {
	dara.Model
	String() string
	GoString() string
	SetCertConfig(v *CertConfig) *UpdateCustomDomainInput
	GetCertConfig() *CertConfig
	SetDescription(v string) *UpdateCustomDomainInput
	GetDescription() *string
	SetProtocol(v string) *UpdateCustomDomainInput
	GetProtocol() *string
	SetRouteConfig(v *RouteConfig) *UpdateCustomDomainInput
	GetRouteConfig() *RouteConfig
	SetTlsConfig(v *TLSConfig) *UpdateCustomDomainInput
	GetTlsConfig() *TLSConfig
}

type UpdateCustomDomainInput struct {
	// HTTPS 证书的信息。
	//
	// example:
	//
	// {}
	CertConfig *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
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
}

func (s UpdateCustomDomainInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomDomainInput) GoString() string {
	return s.String()
}

func (s *UpdateCustomDomainInput) GetCertConfig() *CertConfig {
	return s.CertConfig
}

func (s *UpdateCustomDomainInput) GetDescription() *string {
	return s.Description
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

func (s *UpdateCustomDomainInput) SetCertConfig(v *CertConfig) *UpdateCustomDomainInput {
	s.CertConfig = v
	return s
}

func (s *UpdateCustomDomainInput) SetDescription(v string) *UpdateCustomDomainInput {
	s.Description = &v
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

func (s *UpdateCustomDomainInput) Validate() error {
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
