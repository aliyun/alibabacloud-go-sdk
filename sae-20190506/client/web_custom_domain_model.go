// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebCustomDomain interface {
	dara.Model
	String() string
	GoString() string
	SetCreatedTime(v string) *WebCustomDomain
	GetCreatedTime() *string
	SetDefaultForwardingAppName(v string) *WebCustomDomain
	GetDefaultForwardingAppName() *string
	SetDomainName(v string) *WebCustomDomain
	GetDomainName() *string
	SetLastModifiedTime(v string) *WebCustomDomain
	GetLastModifiedTime() *string
	SetNamespaceId(v string) *WebCustomDomain
	GetNamespaceId() *string
	SetProtocol(v string) *WebCustomDomain
	GetProtocol() *string
	SetRouteConfig(v *RouteConfig) *WebCustomDomain
	GetRouteConfig() *RouteConfig
	SetWebCertConfig(v *WebCertConfig) *WebCustomDomain
	GetWebCertConfig() *WebCertConfig
	SetWebTLSConfig(v *WebTLSConfig) *WebCustomDomain
	GetWebTLSConfig() *WebTLSConfig
	SetWebWAFConfig(v *WebWAFConfig) *WebCustomDomain
	GetWebWAFConfig() *WebWAFConfig
	SetAccountId(v string) *WebCustomDomain
	GetAccountId() *string
}

type WebCustomDomain struct {
	// The time when the custom domain name was created.
	//
	// example:
	//
	// 2023-03-30T08:02:19Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The name of the application that is forwarded by the domain name by default.
	//
	// example:
	//
	// appxxxxx
	DefaultForwardingAppName *string `json:"DefaultForwardingAppName,omitempty" xml:"DefaultForwardingAppName,omitempty"`
	// The domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The time when the custom domain name was last updated.
	//
	// example:
	//
	// 2023-03-30T08:02:19Z
	LastModifiedTime *string `json:"LastModifiedTime,omitempty" xml:"LastModifiedTime,omitempty"`
	// The ID of the namespace to which the domain name belongs.
	//
	// example:
	//
	// cn-chengdu
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The protocol type that is supported by the custom domain name. HTTP: Only HTTP is supported. HTTPS: Only HTTPS is supported. HTTP, HTTPS: Both HTTP and HTTPS are supported.
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
	// The WAF configurations.
	WebWAFConfig *WebWAFConfig `json:"WebWAFConfig,omitempty" xml:"WebWAFConfig,omitempty"`
	// The ID of your Alibaba Cloud account
	//
	// example:
	//
	// 123xxxxxx
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s WebCustomDomain) String() string {
	return dara.Prettify(s)
}

func (s WebCustomDomain) GoString() string {
	return s.String()
}

func (s *WebCustomDomain) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *WebCustomDomain) GetDefaultForwardingAppName() *string {
	return s.DefaultForwardingAppName
}

func (s *WebCustomDomain) GetDomainName() *string {
	return s.DomainName
}

func (s *WebCustomDomain) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *WebCustomDomain) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *WebCustomDomain) GetProtocol() *string {
	return s.Protocol
}

func (s *WebCustomDomain) GetRouteConfig() *RouteConfig {
	return s.RouteConfig
}

func (s *WebCustomDomain) GetWebCertConfig() *WebCertConfig {
	return s.WebCertConfig
}

func (s *WebCustomDomain) GetWebTLSConfig() *WebTLSConfig {
	return s.WebTLSConfig
}

func (s *WebCustomDomain) GetWebWAFConfig() *WebWAFConfig {
	return s.WebWAFConfig
}

func (s *WebCustomDomain) GetAccountId() *string {
	return s.AccountId
}

func (s *WebCustomDomain) SetCreatedTime(v string) *WebCustomDomain {
	s.CreatedTime = &v
	return s
}

func (s *WebCustomDomain) SetDefaultForwardingAppName(v string) *WebCustomDomain {
	s.DefaultForwardingAppName = &v
	return s
}

func (s *WebCustomDomain) SetDomainName(v string) *WebCustomDomain {
	s.DomainName = &v
	return s
}

func (s *WebCustomDomain) SetLastModifiedTime(v string) *WebCustomDomain {
	s.LastModifiedTime = &v
	return s
}

func (s *WebCustomDomain) SetNamespaceId(v string) *WebCustomDomain {
	s.NamespaceId = &v
	return s
}

func (s *WebCustomDomain) SetProtocol(v string) *WebCustomDomain {
	s.Protocol = &v
	return s
}

func (s *WebCustomDomain) SetRouteConfig(v *RouteConfig) *WebCustomDomain {
	s.RouteConfig = v
	return s
}

func (s *WebCustomDomain) SetWebCertConfig(v *WebCertConfig) *WebCustomDomain {
	s.WebCertConfig = v
	return s
}

func (s *WebCustomDomain) SetWebTLSConfig(v *WebTLSConfig) *WebCustomDomain {
	s.WebTLSConfig = v
	return s
}

func (s *WebCustomDomain) SetWebWAFConfig(v *WebWAFConfig) *WebCustomDomain {
	s.WebWAFConfig = v
	return s
}

func (s *WebCustomDomain) SetAccountId(v string) *WebCustomDomain {
	s.AccountId = &v
	return s
}

func (s *WebCustomDomain) Validate() error {
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
