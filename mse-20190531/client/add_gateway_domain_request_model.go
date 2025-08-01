// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGatewayDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *AddGatewayDomainRequest
	GetAcceptLanguage() *string
	SetCertIdentifier(v string) *AddGatewayDomainRequest
	GetCertIdentifier() *string
	SetGatewayUniqueId(v string) *AddGatewayDomainRequest
	GetGatewayUniqueId() *string
	SetHttp2(v string) *AddGatewayDomainRequest
	GetHttp2() *string
	SetMustHttps(v bool) *AddGatewayDomainRequest
	GetMustHttps() *bool
	SetName(v string) *AddGatewayDomainRequest
	GetName() *string
	SetProtocol(v string) *AddGatewayDomainRequest
	GetProtocol() *string
	SetTlsCipherSuitesConfigJSON(v *AddGatewayDomainRequestTlsCipherSuitesConfigJSON) *AddGatewayDomainRequest
	GetTlsCipherSuitesConfigJSON() *AddGatewayDomainRequestTlsCipherSuitesConfigJSON
	SetTlsMax(v string) *AddGatewayDomainRequest
	GetTlsMax() *string
	SetTlsMin(v string) *AddGatewayDomainRequest
	GetTlsMin() *string
}

type AddGatewayDomainRequest struct {
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
	// 6828169-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-86575c0bc9f04ecfbacb92b8e392****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// Specifies whether to enable `HTTP/2`.
	//
	// 	- `open`: enables `HTTP/2`
	//
	// 	- `close`: disables `HTTP/2`
	//
	// 	- `globalConfig`: uses global configurations
	//
	// example:
	//
	// close
	Http2 *string `json:"Http2,omitempty" xml:"Http2,omitempty"`
	// Specifies whether to enable HTTPS.
	//
	// example:
	//
	// true
	MustHttps *bool `json:"MustHttps,omitempty" xml:"MustHttps,omitempty"`
	// The domain name.
	//
	// example:
	//
	// test.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- `HTTP`
	//
	// 	- `HTTPS`
	//
	// example:
	//
	// HTTP
	Protocol                  *string                                           `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	TlsCipherSuitesConfigJSON *AddGatewayDomainRequestTlsCipherSuitesConfigJSON `json:"TlsCipherSuitesConfigJSON,omitempty" xml:"TlsCipherSuitesConfigJSON,omitempty" type:"Struct"`
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

func (s AddGatewayDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayDomainRequest) GoString() string {
	return s.String()
}

func (s *AddGatewayDomainRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *AddGatewayDomainRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *AddGatewayDomainRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *AddGatewayDomainRequest) GetHttp2() *string {
	return s.Http2
}

func (s *AddGatewayDomainRequest) GetMustHttps() *bool {
	return s.MustHttps
}

func (s *AddGatewayDomainRequest) GetName() *string {
	return s.Name
}

func (s *AddGatewayDomainRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *AddGatewayDomainRequest) GetTlsCipherSuitesConfigJSON() *AddGatewayDomainRequestTlsCipherSuitesConfigJSON {
	return s.TlsCipherSuitesConfigJSON
}

func (s *AddGatewayDomainRequest) GetTlsMax() *string {
	return s.TlsMax
}

func (s *AddGatewayDomainRequest) GetTlsMin() *string {
	return s.TlsMin
}

func (s *AddGatewayDomainRequest) SetAcceptLanguage(v string) *AddGatewayDomainRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *AddGatewayDomainRequest) SetCertIdentifier(v string) *AddGatewayDomainRequest {
	s.CertIdentifier = &v
	return s
}

func (s *AddGatewayDomainRequest) SetGatewayUniqueId(v string) *AddGatewayDomainRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *AddGatewayDomainRequest) SetHttp2(v string) *AddGatewayDomainRequest {
	s.Http2 = &v
	return s
}

func (s *AddGatewayDomainRequest) SetMustHttps(v bool) *AddGatewayDomainRequest {
	s.MustHttps = &v
	return s
}

func (s *AddGatewayDomainRequest) SetName(v string) *AddGatewayDomainRequest {
	s.Name = &v
	return s
}

func (s *AddGatewayDomainRequest) SetProtocol(v string) *AddGatewayDomainRequest {
	s.Protocol = &v
	return s
}

func (s *AddGatewayDomainRequest) SetTlsCipherSuitesConfigJSON(v *AddGatewayDomainRequestTlsCipherSuitesConfigJSON) *AddGatewayDomainRequest {
	s.TlsCipherSuitesConfigJSON = v
	return s
}

func (s *AddGatewayDomainRequest) SetTlsMax(v string) *AddGatewayDomainRequest {
	s.TlsMax = &v
	return s
}

func (s *AddGatewayDomainRequest) SetTlsMin(v string) *AddGatewayDomainRequest {
	s.TlsMin = &v
	return s
}

func (s *AddGatewayDomainRequest) Validate() error {
	return dara.Validate(s)
}

type AddGatewayDomainRequestTlsCipherSuitesConfigJSON struct {
	ConfigType      *string   `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	TlsCipherSuites []*string `json:"TlsCipherSuites,omitempty" xml:"TlsCipherSuites,omitempty" type:"Repeated"`
}

func (s AddGatewayDomainRequestTlsCipherSuitesConfigJSON) String() string {
	return dara.Prettify(s)
}

func (s AddGatewayDomainRequestTlsCipherSuitesConfigJSON) GoString() string {
	return s.String()
}

func (s *AddGatewayDomainRequestTlsCipherSuitesConfigJSON) GetConfigType() *string {
	return s.ConfigType
}

func (s *AddGatewayDomainRequestTlsCipherSuitesConfigJSON) GetTlsCipherSuites() []*string {
	return s.TlsCipherSuites
}

func (s *AddGatewayDomainRequestTlsCipherSuitesConfigJSON) SetConfigType(v string) *AddGatewayDomainRequestTlsCipherSuitesConfigJSON {
	s.ConfigType = &v
	return s
}

func (s *AddGatewayDomainRequestTlsCipherSuitesConfigJSON) SetTlsCipherSuites(v []*string) *AddGatewayDomainRequestTlsCipherSuitesConfigJSON {
	s.TlsCipherSuites = v
	return s
}

func (s *AddGatewayDomainRequestTlsCipherSuitesConfigJSON) Validate() error {
	return dara.Validate(s)
}
