// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWebCustomDomainInput interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultForwardingAppName(v string) *CreateWebCustomDomainInput
	GetDefaultForwardingAppName() *string
	SetDomainName(v string) *CreateWebCustomDomainInput
	GetDomainName() *string
	SetProtocol(v string) *CreateWebCustomDomainInput
	GetProtocol() *string
	SetRouteConfig(v *RouteConfig) *CreateWebCustomDomainInput
	GetRouteConfig() *RouteConfig
	SetWebCertConfig(v *WebCertConfig) *CreateWebCustomDomainInput
	GetWebCertConfig() *WebCertConfig
	SetWebTLSConfig(v *WebTLSConfig) *CreateWebCustomDomainInput
	GetWebTLSConfig() *WebTLSConfig
	SetWebWAFConfig(v *WebWAFConfig) *CreateWebCustomDomainInput
	GetWebWAFConfig() *WebWAFConfig
}

type CreateWebCustomDomainInput struct {
	DefaultForwardingAppName *string `json:"DefaultForwardingAppName,omitempty" xml:"DefaultForwardingAppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// HTTP
	Protocol      *string        `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RouteConfig   *RouteConfig   `json:"RouteConfig,omitempty" xml:"RouteConfig,omitempty"`
	WebCertConfig *WebCertConfig `json:"WebCertConfig,omitempty" xml:"WebCertConfig,omitempty"`
	WebTLSConfig  *WebTLSConfig  `json:"WebTLSConfig,omitempty" xml:"WebTLSConfig,omitempty"`
	WebWAFConfig  *WebWAFConfig  `json:"WebWAFConfig,omitempty" xml:"WebWAFConfig,omitempty"`
}

func (s CreateWebCustomDomainInput) String() string {
	return dara.Prettify(s)
}

func (s CreateWebCustomDomainInput) GoString() string {
	return s.String()
}

func (s *CreateWebCustomDomainInput) GetDefaultForwardingAppName() *string {
	return s.DefaultForwardingAppName
}

func (s *CreateWebCustomDomainInput) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateWebCustomDomainInput) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateWebCustomDomainInput) GetRouteConfig() *RouteConfig {
	return s.RouteConfig
}

func (s *CreateWebCustomDomainInput) GetWebCertConfig() *WebCertConfig {
	return s.WebCertConfig
}

func (s *CreateWebCustomDomainInput) GetWebTLSConfig() *WebTLSConfig {
	return s.WebTLSConfig
}

func (s *CreateWebCustomDomainInput) GetWebWAFConfig() *WebWAFConfig {
	return s.WebWAFConfig
}

func (s *CreateWebCustomDomainInput) SetDefaultForwardingAppName(v string) *CreateWebCustomDomainInput {
	s.DefaultForwardingAppName = &v
	return s
}

func (s *CreateWebCustomDomainInput) SetDomainName(v string) *CreateWebCustomDomainInput {
	s.DomainName = &v
	return s
}

func (s *CreateWebCustomDomainInput) SetProtocol(v string) *CreateWebCustomDomainInput {
	s.Protocol = &v
	return s
}

func (s *CreateWebCustomDomainInput) SetRouteConfig(v *RouteConfig) *CreateWebCustomDomainInput {
	s.RouteConfig = v
	return s
}

func (s *CreateWebCustomDomainInput) SetWebCertConfig(v *WebCertConfig) *CreateWebCustomDomainInput {
	s.WebCertConfig = v
	return s
}

func (s *CreateWebCustomDomainInput) SetWebTLSConfig(v *WebTLSConfig) *CreateWebCustomDomainInput {
	s.WebTLSConfig = v
	return s
}

func (s *CreateWebCustomDomainInput) SetWebWAFConfig(v *WebWAFConfig) *CreateWebCustomDomainInput {
	s.WebWAFConfig = v
	return s
}

func (s *CreateWebCustomDomainInput) Validate() error {
	return dara.Validate(s)
}
