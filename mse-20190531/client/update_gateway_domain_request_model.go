// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayDomainRequest
	GetAcceptLanguage() *string
	SetCertIdentifier(v string) *UpdateGatewayDomainRequest
	GetCertIdentifier() *string
	SetGatewayUniqueId(v string) *UpdateGatewayDomainRequest
	GetGatewayUniqueId() *string
	SetHttp2(v string) *UpdateGatewayDomainRequest
	GetHttp2() *string
	SetId(v int64) *UpdateGatewayDomainRequest
	GetId() *int64
	SetMustHttps(v bool) *UpdateGatewayDomainRequest
	GetMustHttps() *bool
	SetProtocol(v string) *UpdateGatewayDomainRequest
	GetProtocol() *string
	SetTlsCipherSuitesConfigJSON(v *UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON) *UpdateGatewayDomainRequest
	GetTlsCipherSuitesConfigJSON() *UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON
	SetTlsMax(v string) *UpdateGatewayDomainRequest
	GetTlsMax() *string
	SetTlsMin(v string) *UpdateGatewayDomainRequest
	GetTlsMin() *string
}

type UpdateGatewayDomainRequest struct {
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
	Protocol                  *string                                              `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	TlsCipherSuitesConfigJSON *UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON `json:"TlsCipherSuitesConfigJSON,omitempty" xml:"TlsCipherSuitesConfigJSON,omitempty" type:"Struct"`
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

func (s UpdateGatewayDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayDomainRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayDomainRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *UpdateGatewayDomainRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayDomainRequest) GetHttp2() *string {
	return s.Http2
}

func (s *UpdateGatewayDomainRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayDomainRequest) GetMustHttps() *bool {
	return s.MustHttps
}

func (s *UpdateGatewayDomainRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateGatewayDomainRequest) GetTlsCipherSuitesConfigJSON() *UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON {
	return s.TlsCipherSuitesConfigJSON
}

func (s *UpdateGatewayDomainRequest) GetTlsMax() *string {
	return s.TlsMax
}

func (s *UpdateGatewayDomainRequest) GetTlsMin() *string {
	return s.TlsMin
}

func (s *UpdateGatewayDomainRequest) SetAcceptLanguage(v string) *UpdateGatewayDomainRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayDomainRequest) SetCertIdentifier(v string) *UpdateGatewayDomainRequest {
	s.CertIdentifier = &v
	return s
}

func (s *UpdateGatewayDomainRequest) SetGatewayUniqueId(v string) *UpdateGatewayDomainRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayDomainRequest) SetHttp2(v string) *UpdateGatewayDomainRequest {
	s.Http2 = &v
	return s
}

func (s *UpdateGatewayDomainRequest) SetId(v int64) *UpdateGatewayDomainRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayDomainRequest) SetMustHttps(v bool) *UpdateGatewayDomainRequest {
	s.MustHttps = &v
	return s
}

func (s *UpdateGatewayDomainRequest) SetProtocol(v string) *UpdateGatewayDomainRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateGatewayDomainRequest) SetTlsCipherSuitesConfigJSON(v *UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON) *UpdateGatewayDomainRequest {
	s.TlsCipherSuitesConfigJSON = v
	return s
}

func (s *UpdateGatewayDomainRequest) SetTlsMax(v string) *UpdateGatewayDomainRequest {
	s.TlsMax = &v
	return s
}

func (s *UpdateGatewayDomainRequest) SetTlsMin(v string) *UpdateGatewayDomainRequest {
	s.TlsMin = &v
	return s
}

func (s *UpdateGatewayDomainRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON struct {
	ConfigType      *string   `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	TlsCipherSuites []*string `json:"TlsCipherSuites,omitempty" xml:"TlsCipherSuites,omitempty" type:"Repeated"`
}

func (s UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON) GoString() string {
	return s.String()
}

func (s *UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON) GetConfigType() *string {
	return s.ConfigType
}

func (s *UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON) GetTlsCipherSuites() []*string {
	return s.TlsCipherSuites
}

func (s *UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON) SetConfigType(v string) *UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON {
	s.ConfigType = &v
	return s
}

func (s *UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON) SetTlsCipherSuites(v []*string) *UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON {
	s.TlsCipherSuites = v
	return s
}

func (s *UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON) Validate() error {
	return dara.Validate(s)
}
