// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWebCustomDomainInput interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultForwardingAppName(v string) *UpdateWebCustomDomainInput
	GetDefaultForwardingAppName() *string
	SetProtocol(v string) *UpdateWebCustomDomainInput
	GetProtocol() *string
	SetRouteConfig(v *RouteConfig) *UpdateWebCustomDomainInput
	GetRouteConfig() *RouteConfig
	SetWebCertConfig(v *WebCertConfig) *UpdateWebCustomDomainInput
	GetWebCertConfig() *WebCertConfig
	SetWebTLSConfig(v *WebTLSConfig) *UpdateWebCustomDomainInput
	GetWebTLSConfig() *WebTLSConfig
	SetWebWAFConfig(v *WebWAFConfig) *UpdateWebCustomDomainInput
	GetWebWAFConfig() *WebWAFConfig
}

type UpdateWebCustomDomainInput struct {
	DefaultForwardingAppName *string `json:"DefaultForwardingAppName,omitempty" xml:"DefaultForwardingAppName,omitempty"`
	// example:
	//
	// HTTP
	Protocol      *string        `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RouteConfig   *RouteConfig   `json:"RouteConfig,omitempty" xml:"RouteConfig,omitempty"`
	WebCertConfig *WebCertConfig `json:"WebCertConfig,omitempty" xml:"WebCertConfig,omitempty"`
	WebTLSConfig  *WebTLSConfig  `json:"WebTLSConfig,omitempty" xml:"WebTLSConfig,omitempty"`
	WebWAFConfig  *WebWAFConfig  `json:"WebWAFConfig,omitempty" xml:"WebWAFConfig,omitempty"`
}

func (s UpdateWebCustomDomainInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateWebCustomDomainInput) GoString() string {
	return s.String()
}

func (s *UpdateWebCustomDomainInput) GetDefaultForwardingAppName() *string {
	return s.DefaultForwardingAppName
}

func (s *UpdateWebCustomDomainInput) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateWebCustomDomainInput) GetRouteConfig() *RouteConfig {
	return s.RouteConfig
}

func (s *UpdateWebCustomDomainInput) GetWebCertConfig() *WebCertConfig {
	return s.WebCertConfig
}

func (s *UpdateWebCustomDomainInput) GetWebTLSConfig() *WebTLSConfig {
	return s.WebTLSConfig
}

func (s *UpdateWebCustomDomainInput) GetWebWAFConfig() *WebWAFConfig {
	return s.WebWAFConfig
}

func (s *UpdateWebCustomDomainInput) SetDefaultForwardingAppName(v string) *UpdateWebCustomDomainInput {
	s.DefaultForwardingAppName = &v
	return s
}

func (s *UpdateWebCustomDomainInput) SetProtocol(v string) *UpdateWebCustomDomainInput {
	s.Protocol = &v
	return s
}

func (s *UpdateWebCustomDomainInput) SetRouteConfig(v *RouteConfig) *UpdateWebCustomDomainInput {
	s.RouteConfig = v
	return s
}

func (s *UpdateWebCustomDomainInput) SetWebCertConfig(v *WebCertConfig) *UpdateWebCustomDomainInput {
	s.WebCertConfig = v
	return s
}

func (s *UpdateWebCustomDomainInput) SetWebTLSConfig(v *WebTLSConfig) *UpdateWebCustomDomainInput {
	s.WebTLSConfig = v
	return s
}

func (s *UpdateWebCustomDomainInput) SetWebWAFConfig(v *WebWAFConfig) *UpdateWebCustomDomainInput {
	s.WebWAFConfig = v
	return s
}

func (s *UpdateWebCustomDomainInput) Validate() error {
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
