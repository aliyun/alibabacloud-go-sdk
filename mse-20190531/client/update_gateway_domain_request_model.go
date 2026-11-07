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
	SetCaCertIdentifier(v string) *UpdateGatewayDomainRequest
	GetCaCertIdentifier() *string
	SetCertIdentifier(v string) *UpdateGatewayDomainRequest
	GetCertIdentifier() *string
	SetEnableWaf(v bool) *UpdateGatewayDomainRequest
	GetEnableWaf() *bool
	SetGatewayUniqueId(v string) *UpdateGatewayDomainRequest
	GetGatewayUniqueId() *string
	SetHttp2(v string) *UpdateGatewayDomainRequest
	GetHttp2() *string
	SetId(v int64) *UpdateGatewayDomainRequest
	GetId() *int64
	SetMtlsEnabled(v bool) *UpdateGatewayDomainRequest
	GetMtlsEnabled() *bool
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
	// The language of the returned information. Valid values:
	//
	// - zh: Chinese (default).
	//
	// - en: English.
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The CA certificate content.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIFazCCA1OgAwIBAgIRAIIQz7DSQONZRGPgu2OCiwAwDQYJKoZIhvcNAQELBQAw
	//
	// TzELMAkGA1UEBhMCVVMxKTAnBgNVBAoTIEludGVybmV0IFNlY3VyaXR5IFJlc2Vh
	//
	// cmNoIEdyb3VwMRUwEwYDVQQDEwxJU1JHIFJvb3QgWDEwHhcNMTUwNjA0MTEwNDM4
	//
	// WhcNMzUwNjA0MTEwNDM4WjBPMQswCQYDVQQGEwJVUzEpMCcGA1UEChMgSW50ZXJu
	//
	// ZXQgU2VjdXJpdHkgUmVzZWFyY2ggR3JvdXAxFTATBgNVBAMTDElTUkcgUm9vdCBY
	//
	// MTCCAiIwDQYJKoZIhvcNAQEBBQADggIPADCCAgoCggIBAK3oJHP0FDfzm54rVygc
	//
	// h77ct984kIxuPOZXoHj3dcKi/vVqbvYATyjb3miGbESTtrFj/RQSa78f0uoxmyF+
	//
	// 0TM8ukj13Xnfs7j/EvEhmkvBioZxaUpmZmyPfjxwv60pIgbz5MDmgK7iS4+3mX6U
	//
	// A5/TR5d8mUgjU+g4rk8Kb4Mu0UlXjIB0ttov0DiNewNwIRt18jA8+o+u3dpjq+sW
	//
	// T8KOEUt+zwvo/7V3LvSye0rgTBIlDHCNAymg4VMk7BPZ7hm/ELNKjD+Jo2FR3qyH
	//
	// B5T0Y3HsLuJvW5iB4YlcNHlsdu87kGJ55tukmi8mxdAQ4Q7e2RCOFvu396j3x+UC
	//
	// B5iPNgiV5+I3lg02dZ77DnKxHZu8A/lJBdiB3QW0KtZB6awBdpUKD9jf1b0SHzUv
	//
	// KBds0pjBqAlkd25HN7rOrFleaJ1/ctaJxQZBKT5ZPt0m9STJEadao0xAH0ahmbWn
	//
	// OlFuhjuefXKnEgV4We0+UXgVCwOPjdAvBbI+e0ocS3MFEvzG6uBQE3xDk3SzynTn
	//
	// jh8BCNAw1FtxNrQHusEwMFxIt4I7mKZ9YIqioymCzLq9gwQbooMDQaHWBfEbwrbw
	//
	// qHyGO0aoSCqI3Haadr8faqU9GY/rOPNk3sgrDQoo//fb4hVC1CLQJ13hef4Y53CI
	//
	// rU7m2Ys6xt0nUW7/vGT1M0NPAgMBAAGjQjBAMA4GA1UdDwEB/wQEAwIBBjAPBgNV
	//
	// HRMBAf8EBTADAQH/MB0GA1UdDgQWBBR5tFnme7bl5AFzgAiIyBpY9umbbjANBgkq
	//
	// hkiG9w0BAQsFAAOCAgEAVR9YqbyyqFDQDLHYGmkgJykIrGF1XIpu+ILlaS/V9lZL
	//
	// ubhzEFnTIZd+50xx+7LSYK05qAvqFyFWhfFQDlnrzuBZ6brJFe+GnY+EgPbk6ZGQ
	//
	// 3BebYhtF8GaV0nxvwuo77x/Py9auJ/GpsMiu/X1+mvoiBOv/2X/qkSsisRcOj/KK
	//
	// NFtY2PwByVS5uCbMiogziUwthDyC3+6WVwW6LLv3xLfHTjuCvjHIInNzktHCgKQ5
	//
	// ORAzI4JMPJ+GslWYHb4phowim57iaztXOoJwTdwJx4nLCgdNbOhdjsnvzqvHu7Ur
	//
	// TkXWStAmzOVyyghqpZXjFaH3pO3JLF+l+/+sKAIuvtd7u+Nxe5AW0wdeRlN8NwdC
	//
	// jNPElpzVmbUq4JUagEiuTDkHzsxHpFKVK7q4+63SM1N95R1NbdWhscdCb+ZAJzVc
	//
	// oyi3B43njTOQ5yOf+1CceWxG1bQVs5ZufpsMljq4Ui0/1lvh+wjChP4kqKOJ2qxq
	//
	// 4RgqsahDYVvTH9w7jXbyLeiNdd8XM2w9U/t7y0Ff/9yi0GE44Za4rF2LN9d11TPA
	//
	// mRGunUHBcnWEvgJBQl9nJEiU0Zsnvgc/ubhPgXRR4Xq37Z0j4r7g1SgEEzwxA57d
	//
	// emyPxgcYxn/eR44/KJ4EBs+lVDR3veyJm+kXQ99b21/+jh5Xos1AnX5iItreGCc=
	//
	// -----END CERTIFICATE-----
	CaCertIdentifier *string `json:"CaCertIdentifier,omitempty" xml:"CaCertIdentifier,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// 6209108-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// Specifies whether to enable Web Application Firewall (WAF) for the gateway domain name.
	EnableWaf *bool `json:"EnableWaf,omitempty" xml:"EnableWaf,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-86575c0bc9f04ecfbacb92b8e392****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// Specifies whether to enable HTTP/2. Valid values:
	//
	// - `open`: enables HTTP/2.
	//
	// - `close`: disables HTTP/2.
	//
	// - `globalConfig`: uses the global configuration.
	//
	// example:
	//
	// close
	Http2 *string `json:"Http2,omitempty" xml:"Http2,omitempty"`
	// The ID of the domain name to be updated.
	//
	// example:
	//
	// 94
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Specifies whether to enable mTLS.
	MtlsEnabled *bool `json:"MtlsEnabled,omitempty" xml:"MtlsEnabled,omitempty"`
	// Specifies whether to force HTTPS.
	//
	// example:
	//
	// false
	MustHttps *bool `json:"MustHttps,omitempty" xml:"MustHttps,omitempty"`
	// The protocol type. Valid values:
	//
	// - HTTPS
	//
	// - HTTP
	//
	// example:
	//
	// HTTPS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The TLS cipher suite configuration.
	TlsCipherSuitesConfigJSON *UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON `json:"TlsCipherSuitesConfigJSON,omitempty" xml:"TlsCipherSuitesConfigJSON,omitempty" type:"Struct"`
	// The maximum TLS version.
	//
	// example:
	//
	// TLS 1.3
	TlsMax *string `json:"TlsMax,omitempty" xml:"TlsMax,omitempty"`
	// The minimum TLS version.
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

func (s *UpdateGatewayDomainRequest) GetCaCertIdentifier() *string {
	return s.CaCertIdentifier
}

func (s *UpdateGatewayDomainRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *UpdateGatewayDomainRequest) GetEnableWaf() *bool {
	return s.EnableWaf
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

func (s *UpdateGatewayDomainRequest) GetMtlsEnabled() *bool {
	return s.MtlsEnabled
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

func (s *UpdateGatewayDomainRequest) SetCaCertIdentifier(v string) *UpdateGatewayDomainRequest {
	s.CaCertIdentifier = &v
	return s
}

func (s *UpdateGatewayDomainRequest) SetCertIdentifier(v string) *UpdateGatewayDomainRequest {
	s.CertIdentifier = &v
	return s
}

func (s *UpdateGatewayDomainRequest) SetEnableWaf(v bool) *UpdateGatewayDomainRequest {
	s.EnableWaf = &v
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

func (s *UpdateGatewayDomainRequest) SetMtlsEnabled(v bool) *UpdateGatewayDomainRequest {
	s.MtlsEnabled = &v
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
	if s.TlsCipherSuitesConfigJSON != nil {
		if err := s.TlsCipherSuitesConfigJSON.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGatewayDomainRequestTlsCipherSuitesConfigJSON struct {
	// The cipher suite configuration type. Valid values:
	//
	// - DEFAULT
	//
	// - CUSTOM
	//
	// example:
	//
	// DEFAULT
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The list of cipher suite names.
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
