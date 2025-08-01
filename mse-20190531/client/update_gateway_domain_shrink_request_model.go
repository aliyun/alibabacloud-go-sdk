// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayDomainShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayDomainShrinkRequest
	GetAcceptLanguage() *string
	SetCertIdentifier(v string) *UpdateGatewayDomainShrinkRequest
	GetCertIdentifier() *string
	SetGatewayUniqueId(v string) *UpdateGatewayDomainShrinkRequest
	GetGatewayUniqueId() *string
	SetHttp2(v string) *UpdateGatewayDomainShrinkRequest
	GetHttp2() *string
	SetId(v int64) *UpdateGatewayDomainShrinkRequest
	GetId() *int64
	SetMustHttps(v bool) *UpdateGatewayDomainShrinkRequest
	GetMustHttps() *bool
	SetProtocol(v string) *UpdateGatewayDomainShrinkRequest
	GetProtocol() *string
	SetTlsCipherSuitesConfigJSONShrink(v string) *UpdateGatewayDomainShrinkRequest
	GetTlsCipherSuitesConfigJSONShrink() *string
	SetTlsMax(v string) *UpdateGatewayDomainShrinkRequest
	GetTlsMax() *string
	SetTlsMin(v string) *UpdateGatewayDomainShrinkRequest
	GetTlsMin() *string
}

type UpdateGatewayDomainShrinkRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 6209108-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-86575c0bc9f04ecfbacb92b8e392****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// Specifies whether to enable `HTTP/2`.
	//
	// 	- `open`: `HTTP/2` is enabled.
	//
	// 	- `close`: `HTTP/2` is disabled.
	//
	// 	- `globalConfig`: Global configurations are used.
	//
	// example:
	//
	// close
	Http2 *string `json:"Http2,omitempty" xml:"Http2,omitempty"`
	// The ID of the domain name that you want to update.
	//
	// example:
	//
	// 94
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether to forcibly use HTTPS.
	//
	// example:
	//
	// false
	MustHttps *bool `json:"MustHttps,omitempty" xml:"MustHttps,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- HTTPS
	//
	// 	- HTTP
	//
	// example:
	//
	// HTTPS
	Protocol                        *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	TlsCipherSuitesConfigJSONShrink *string `json:"TlsCipherSuitesConfigJSON,omitempty" xml:"TlsCipherSuitesConfigJSON,omitempty"`
	// The maximum version of Transport Layer Security (TLS).
	//
	// example:
	//
	// TLS 1.3
	TlsMax *string `json:"TlsMax,omitempty" xml:"TlsMax,omitempty"`
	// The minimum version of TLS.
	//
	// example:
	//
	// TLS 1.0
	TlsMin *string `json:"TlsMin,omitempty" xml:"TlsMin,omitempty"`
}

func (s UpdateGatewayDomainShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayDomainShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayDomainShrinkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayDomainShrinkRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *UpdateGatewayDomainShrinkRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayDomainShrinkRequest) GetHttp2() *string {
	return s.Http2
}

func (s *UpdateGatewayDomainShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayDomainShrinkRequest) GetMustHttps() *bool {
	return s.MustHttps
}

func (s *UpdateGatewayDomainShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateGatewayDomainShrinkRequest) GetTlsCipherSuitesConfigJSONShrink() *string {
	return s.TlsCipherSuitesConfigJSONShrink
}

func (s *UpdateGatewayDomainShrinkRequest) GetTlsMax() *string {
	return s.TlsMax
}

func (s *UpdateGatewayDomainShrinkRequest) GetTlsMin() *string {
	return s.TlsMin
}

func (s *UpdateGatewayDomainShrinkRequest) SetAcceptLanguage(v string) *UpdateGatewayDomainShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayDomainShrinkRequest) SetCertIdentifier(v string) *UpdateGatewayDomainShrinkRequest {
	s.CertIdentifier = &v
	return s
}

func (s *UpdateGatewayDomainShrinkRequest) SetGatewayUniqueId(v string) *UpdateGatewayDomainShrinkRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayDomainShrinkRequest) SetHttp2(v string) *UpdateGatewayDomainShrinkRequest {
	s.Http2 = &v
	return s
}

func (s *UpdateGatewayDomainShrinkRequest) SetId(v int64) *UpdateGatewayDomainShrinkRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayDomainShrinkRequest) SetMustHttps(v bool) *UpdateGatewayDomainShrinkRequest {
	s.MustHttps = &v
	return s
}

func (s *UpdateGatewayDomainShrinkRequest) SetProtocol(v string) *UpdateGatewayDomainShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateGatewayDomainShrinkRequest) SetTlsCipherSuitesConfigJSONShrink(v string) *UpdateGatewayDomainShrinkRequest {
	s.TlsCipherSuitesConfigJSONShrink = &v
	return s
}

func (s *UpdateGatewayDomainShrinkRequest) SetTlsMax(v string) *UpdateGatewayDomainShrinkRequest {
	s.TlsMax = &v
	return s
}

func (s *UpdateGatewayDomainShrinkRequest) SetTlsMin(v string) *UpdateGatewayDomainShrinkRequest {
	s.TlsMin = &v
	return s
}

func (s *UpdateGatewayDomainShrinkRequest) Validate() error {
	return dara.Validate(s)
}
