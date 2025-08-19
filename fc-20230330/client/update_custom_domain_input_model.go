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
	AuthConfig *AuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	CertConfig *CertConfig `json:"certConfig,omitempty" xml:"certConfig,omitempty"`
	// example:
	//
	// HTTP
	Protocol    *string      `json:"protocol,omitempty" xml:"protocol,omitempty"`
	RouteConfig *RouteConfig `json:"routeConfig,omitempty" xml:"routeConfig,omitempty"`
	TlsConfig   *TLSConfig   `json:"tlsConfig,omitempty" xml:"tlsConfig,omitempty"`
	WafConfig   *WAFConfig   `json:"wafConfig,omitempty" xml:"wafConfig,omitempty"`
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
	return dara.Validate(s)
}
