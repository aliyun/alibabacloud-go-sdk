// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayDomainDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetGatewayDomainDetailResponseBody
	GetCode() *int32
	SetData(v *GetGatewayDomainDetailResponseBodyData) *GetGatewayDomainDetailResponseBody
	GetData() *GetGatewayDomainDetailResponseBodyData
	SetHttpStatusCode(v int32) *GetGatewayDomainDetailResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetGatewayDomainDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayDomainDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetGatewayDomainDetailResponseBody
	GetSuccess() *bool
}

type GetGatewayDomainDetailResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 403
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data.
	Data *GetGatewayDomainDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 88B83302-CD88-54D3-8DF2-208BFDC73F0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - `true`: The request is successful.
	//
	// - `false`: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetGatewayDomainDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayDomainDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetGatewayDomainDetailResponseBody) GetData() *GetGatewayDomainDetailResponseBodyData {
	return s.Data
}

func (s *GetGatewayDomainDetailResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetGatewayDomainDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayDomainDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayDomainDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetGatewayDomainDetailResponseBody) SetCode(v int32) *GetGatewayDomainDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBody) SetData(v *GetGatewayDomainDetailResponseBodyData) *GetGatewayDomainDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayDomainDetailResponseBody) SetHttpStatusCode(v int32) *GetGatewayDomainDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBody) SetMessage(v string) *GetGatewayDomainDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBody) SetRequestId(v string) *GetGatewayDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBody) SetSuccess(v bool) *GetGatewayDomainDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayDomainDetailResponseBodyData struct {
	// The certificate expiration time.
	//
	// example:
	//
	// 2021-04-01 02:35:12
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The algorithm.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The certificate issuance time.
	//
	// example:
	//
	// 2021-04-01 02:35:12
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
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
	// The certificate identifier.
	//
	// example:
	//
	// 234-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The certificate name.
	//
	// example:
	//
	// test
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The public domain name.
	//
	// example:
	//
	// name
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// Indicates whether WAF protection is enabled for the domain name.
	EnableWaf *bool `json:"EnableWaf,omitempty" xml:"EnableWaf,omitempty"`
	// The gateway ID.
	//
	// example:
	//
	// 1
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-7ea3da97b96543e19f6c597cd4a9****
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The certificate expiration time in GMT.
	//
	// example:
	//
	// 2021-04-01 02:35:12
	GmtAfter *string `json:"GmtAfter,omitempty" xml:"GmtAfter,omitempty"`
	// The certificate issuance time in GMT.
	//
	// example:
	//
	// 2021-04-01 02:35:12
	GmtBefore *string `json:"GmtBefore,omitempty" xml:"GmtBefore,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2022-01-07T10:07:57.000+0000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether HTTP/2 is enabled. Valid values:
	//
	// - `open`: HTTP/2 is enabled.
	//
	// - `close`: HTTP/2 is disabled.
	//
	// - `globalConfig`: The global configuration is used.
	//
	// example:
	//
	// close
	Http2 *string `json:"Http2,omitempty" xml:"Http2,omitempty"`
	// The ID.
	//
	// example:
	//
	// 12
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the domain name is managed.
	IsManaged *bool `json:"IsManaged,omitempty" xml:"IsManaged,omitempty"`
	// The issuer.
	//
	// example:
	//
	// Istio
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// Indicates whether mTLS is enabled.
	MtlsEnabled *bool `json:"MtlsEnabled,omitempty" xml:"MtlsEnabled,omitempty"`
	// Indicates whether HTTPS is forced.
	//
	// example:
	//
	// true
	MustHttps *bool `json:"MustHttps,omitempty" xml:"MustHttps,omitempty"`
	// The name.
	//
	// example:
	//
	// test.com
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The gateway protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The subject alternative names.
	//
	// example:
	//
	// test.com
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// Indicates whether the gateway supports WAF protection.
	SupportWaf *bool `json:"SupportWaf,omitempty" xml:"SupportWaf,omitempty"`
	// The TLS cipher suite configuration, including the configuration type and the list of cipher suite names.
	TlsCipherSuitesConfig *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig `json:"TlsCipherSuitesConfig,omitempty" xml:"TlsCipherSuitesConfig,omitempty" type:"Struct"`
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

func (s GetGatewayDomainDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayDomainDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayDomainDetailResponseBodyData) GetAfterDate() *int64 {
	return s.AfterDate
}

func (s *GetGatewayDomainDetailResponseBodyData) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *GetGatewayDomainDetailResponseBodyData) GetBeforeDate() *int64 {
	return s.BeforeDate
}

func (s *GetGatewayDomainDetailResponseBodyData) GetCaCertIdentifier() *string {
	return s.CaCertIdentifier
}

func (s *GetGatewayDomainDetailResponseBodyData) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *GetGatewayDomainDetailResponseBodyData) GetCertName() *string {
	return s.CertName
}

func (s *GetGatewayDomainDetailResponseBodyData) GetCommonName() *string {
	return s.CommonName
}

func (s *GetGatewayDomainDetailResponseBodyData) GetEnableWaf() *bool {
	return s.EnableWaf
}

func (s *GetGatewayDomainDetailResponseBodyData) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *GetGatewayDomainDetailResponseBodyData) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *GetGatewayDomainDetailResponseBodyData) GetGmtAfter() *string {
	return s.GmtAfter
}

func (s *GetGatewayDomainDetailResponseBodyData) GetGmtBefore() *string {
	return s.GmtBefore
}

func (s *GetGatewayDomainDetailResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetGatewayDomainDetailResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetGatewayDomainDetailResponseBodyData) GetHttp2() *string {
	return s.Http2
}

func (s *GetGatewayDomainDetailResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetGatewayDomainDetailResponseBodyData) GetIsManaged() *bool {
	return s.IsManaged
}

func (s *GetGatewayDomainDetailResponseBodyData) GetIssuer() *string {
	return s.Issuer
}

func (s *GetGatewayDomainDetailResponseBodyData) GetMtlsEnabled() *bool {
	return s.MtlsEnabled
}

func (s *GetGatewayDomainDetailResponseBodyData) GetMustHttps() *bool {
	return s.MustHttps
}

func (s *GetGatewayDomainDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetGatewayDomainDetailResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetGatewayDomainDetailResponseBodyData) GetSans() *string {
	return s.Sans
}

func (s *GetGatewayDomainDetailResponseBodyData) GetSupportWaf() *bool {
	return s.SupportWaf
}

func (s *GetGatewayDomainDetailResponseBodyData) GetTlsCipherSuitesConfig() *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig {
	return s.TlsCipherSuitesConfig
}

func (s *GetGatewayDomainDetailResponseBodyData) GetTlsMax() *string {
	return s.TlsMax
}

func (s *GetGatewayDomainDetailResponseBodyData) GetTlsMin() *string {
	return s.TlsMin
}

func (s *GetGatewayDomainDetailResponseBodyData) SetAfterDate(v int64) *GetGatewayDomainDetailResponseBodyData {
	s.AfterDate = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetAlgorithm(v string) *GetGatewayDomainDetailResponseBodyData {
	s.Algorithm = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetBeforeDate(v int64) *GetGatewayDomainDetailResponseBodyData {
	s.BeforeDate = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetCaCertIdentifier(v string) *GetGatewayDomainDetailResponseBodyData {
	s.CaCertIdentifier = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetCertIdentifier(v string) *GetGatewayDomainDetailResponseBodyData {
	s.CertIdentifier = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetCertName(v string) *GetGatewayDomainDetailResponseBodyData {
	s.CertName = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetCommonName(v string) *GetGatewayDomainDetailResponseBodyData {
	s.CommonName = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetEnableWaf(v bool) *GetGatewayDomainDetailResponseBodyData {
	s.EnableWaf = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetGatewayId(v int64) *GetGatewayDomainDetailResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetGatewayUniqueId(v string) *GetGatewayDomainDetailResponseBodyData {
	s.GatewayUniqueId = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetGmtAfter(v string) *GetGatewayDomainDetailResponseBodyData {
	s.GmtAfter = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetGmtBefore(v string) *GetGatewayDomainDetailResponseBodyData {
	s.GmtBefore = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetGmtCreate(v string) *GetGatewayDomainDetailResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetGmtModified(v string) *GetGatewayDomainDetailResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetHttp2(v string) *GetGatewayDomainDetailResponseBodyData {
	s.Http2 = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetId(v int64) *GetGatewayDomainDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetIsManaged(v bool) *GetGatewayDomainDetailResponseBodyData {
	s.IsManaged = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetIssuer(v string) *GetGatewayDomainDetailResponseBodyData {
	s.Issuer = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetMtlsEnabled(v bool) *GetGatewayDomainDetailResponseBodyData {
	s.MtlsEnabled = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetMustHttps(v bool) *GetGatewayDomainDetailResponseBodyData {
	s.MustHttps = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetName(v string) *GetGatewayDomainDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetProtocol(v string) *GetGatewayDomainDetailResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetSans(v string) *GetGatewayDomainDetailResponseBodyData {
	s.Sans = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetSupportWaf(v bool) *GetGatewayDomainDetailResponseBodyData {
	s.SupportWaf = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetTlsCipherSuitesConfig(v *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) *GetGatewayDomainDetailResponseBodyData {
	s.TlsCipherSuitesConfig = v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetTlsMax(v string) *GetGatewayDomainDetailResponseBodyData {
	s.TlsMax = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) SetTlsMin(v string) *GetGatewayDomainDetailResponseBodyData {
	s.TlsMin = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyData) Validate() error {
	if s.TlsCipherSuitesConfig != nil {
		if err := s.TlsCipherSuitesConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig struct {
	// The TLS cipher suite configuration type. DEFAULT indicates the default configuration, and CUSTOM indicates the custom configuration.
	//
	// example:
	//
	// DEFAULT
	ConfigType *string `json:"ConfigType,omitempty" xml:"ConfigType,omitempty"`
	// The list of TLS cipher suite names.
	TlsCipherSuites []*string `json:"TlsCipherSuites,omitempty" xml:"TlsCipherSuites,omitempty" type:"Repeated"`
}

func (s GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) GoString() string {
	return s.String()
}

func (s *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) GetConfigType() *string {
	return s.ConfigType
}

func (s *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) GetTlsCipherSuites() []*string {
	return s.TlsCipherSuites
}

func (s *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) SetConfigType(v string) *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig {
	s.ConfigType = &v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) SetTlsCipherSuites(v []*string) *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig {
	s.TlsCipherSuites = v
	return s
}

func (s *GetGatewayDomainDetailResponseBodyDataTlsCipherSuitesConfig) Validate() error {
	return dara.Validate(s)
}
