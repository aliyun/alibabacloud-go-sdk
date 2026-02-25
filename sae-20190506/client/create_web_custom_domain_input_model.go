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
	// The name of the application to which data is forwarded by the domain name by default.
	//
	// example:
	//
	// demo-app
	DefaultForwardingAppName *string `json:"DefaultForwardingAppName,omitempty" xml:"DefaultForwardingAppName,omitempty"`
	// The domain name. Enter a custom domain name that has obtained an Internet content provider (ICP) filing in the Alibaba Cloud ICP Filing system, or a custom domain name whose ICP filing information includes Alibaba Cloud as a service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The protocol type that is supported by the custom domain name. Valid values:
	//
	// 	- **HTTP**: Only HTTP is supported.
	//
	// 	- **HTTPS**: Only HTTPS is supported.
	//
	// 	- **HTTP,HTTPS**: Both HTTP and HTTPS are supported.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The route configurations.
	RouteConfig *RouteConfig `json:"RouteConfig,omitempty" xml:"RouteConfig,omitempty"`
	// The information about the HTTPS certificate.
	WebCertConfig *WebCertConfig `json:"WebCertConfig,omitempty" xml:"WebCertConfig,omitempty"`
	// The Transport Layer Security (TLS) configurations.
	WebTLSConfig *WebTLSConfig `json:"WebTLSConfig,omitempty" xml:"WebTLSConfig,omitempty"`
	// The Web Application Firewall (WAF) configurations.
	WebWAFConfig *WebWAFConfig `json:"WebWAFConfig,omitempty" xml:"WebWAFConfig,omitempty"`
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
	if s.RouteConfig != nil {
		if err := s.RouteConfig.Validate(); err != nil {
			return err
		}
	}
	if s.WebCertConfig != nil {
		if err := s.WebCertConfig.Validate(); err != nil {
			return err
		}
	}
	if s.WebTLSConfig != nil {
		if err := s.WebTLSConfig.Validate(); err != nil {
			return err
		}
	}
	if s.WebWAFConfig != nil {
		if err := s.WebWAFConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
