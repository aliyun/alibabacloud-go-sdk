// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDomainResponseBody
	GetCode() *string
	SetData(v *GetDomainResponseBodyData) *GetDomainResponseBody
	GetData() *GetDomainResponseBodyData
	SetMessage(v string) *GetDomainResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDomainResponseBody
	GetRequestId() *string
}

type GetDomainResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data returned.
	Data *GetDomainResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message returned.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which is used to trace the API call link.
	//
	// example:
	//
	// 3ACFC7A7-45A9-58CF-B2D5-765B60254695
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDomainResponseBody) GetData() *GetDomainResponseBodyData {
	return s.Data
}

func (s *GetDomainResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDomainResponseBody) SetCode(v string) *GetDomainResponseBody {
	s.Code = &v
	return s
}

func (s *GetDomainResponseBody) SetData(v *GetDomainResponseBodyData) *GetDomainResponseBody {
	s.Data = v
	return s
}

func (s *GetDomainResponseBody) SetMessage(v string) *GetDomainResponseBody {
	s.Message = &v
	return s
}

func (s *GetDomainResponseBody) SetRequestId(v string) *GetDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDomainResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDomainResponseBodyData struct {
	// The encryption algorithm.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// The CA certificate ID.
	//
	// example:
	//
	// 876****-cn-hangzhou
	CaCertIdentifier *string `json:"caCertIdentifier,omitempty" xml:"caCertIdentifier,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// 645****-cn-hangzhou
	CertIdentifier *string `json:"certIdentifier,omitempty" xml:"certIdentifier,omitempty"`
	// The certificate name.
	//
	// example:
	//
	// test-cert
	CertName *string `json:"certName,omitempty" xml:"certName,omitempty"`
	// The client CA certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/mpTQwDQYJKoZIhvcNAQEL
	//
	// BxSbrGeJ8i0576Gn7Qezyho9abZOUhGaPeoB
	//
	// AIHWWl428uUSG/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	//
	// yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy+ZMJ8r4swA4swHwYDVR0jBBgwFoAU
	//
	// qroVyYKk7ylhcSn+ZMJ8r4swA4swDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0B
	//
	// xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx=
	//
	// -----END CERTIFICATE-----
	ClientCACert *string `json:"clientCACert,omitempty" xml:"clientCACert,omitempty"`
	// The creation source.
	//
	// Valid values:
	//
	// 	- Console
	//
	// 	- Ingress
	//
	// example:
	//
	// Console
	CreateFrom *string `json:"createFrom,omitempty" xml:"createFrom,omitempty"`
	// The creation timestamp.
	//
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// Indicates whether the domain name is the default domain name.
	//
	// example:
	//
	// false
	Default *bool `json:"default,omitempty" xml:"default,omitempty"`
	// The ID of the domain name.
	//
	// example:
	//
	// d-cq1m3utlhtgvgkv7sitg
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// Indicates whether forcible HTTPS redirection is enabled.
	//
	// example:
	//
	// false
	ForceHttps *bool `json:"forceHttps,omitempty" xml:"forceHttps,omitempty"`
	// The HTTP/2 configuration.
	//
	// Valid values:
	//
	// 	- GlobalConfig
	//
	// 	- Close
	//
	// 	- Open
	//
	// example:
	//
	// Open
	Http2Option *string `json:"http2Option,omitempty" xml:"http2Option,omitempty"`
	// The certificate issuer.
	//
	// example:
	//
	// Alibaba
	Issuer *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	// Indicates whether mutual authentication is enabled.
	//
	// Valid values:
	//
	// 	- false
	//
	// 	- true
	//
	// example:
	//
	// false
	MTLSEnabled *bool `json:"mTLSEnabled,omitempty" xml:"mTLSEnabled,omitempty"`
	// The domain name.
	//
	// example:
	//
	// abc.com
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The expiration time of the certificate.
	//
	// example:
	//
	// 1719386834548
	NotAfterTimstamp *int64 `json:"notAfterTimstamp,omitempty" xml:"notAfterTimstamp,omitempty"`
	// The time when the certificate started to take effect.
	//
	// example:
	//
	// 1719386834548
	NotBeforeTimestamp *int64 `json:"notBeforeTimestamp,omitempty" xml:"notBeforeTimestamp,omitempty"`
	// The supported protocol. Valid values:
	//
	// 	- HTTP: Only HTTP is supported.
	//
	// 	- HTTPS: Only HTTPS is supported.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aekzvlxzgo5b4si
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// All domain names that are bound to the certificate.
	//
	// example:
	//
	// aliyun.com
	Sans *string `json:"sans,omitempty" xml:"sans,omitempty"`
	// The information about online resources.
	StatisticsInfo *GetDomainResponseBodyDataStatisticsInfo `json:"statisticsInfo,omitempty" xml:"statisticsInfo,omitempty" type:"Struct"`
	// The cipher suite configuration.
	TlsCipherSuitesConfig *TlsCipherSuitesConfig `json:"tlsCipherSuitesConfig,omitempty" xml:"tlsCipherSuitesConfig,omitempty"`
	// The maximum version of the TLS protocol. Up to TLS 1.3 is supported.
	//
	// example:
	//
	// TLS 1.3
	TlsMax *string `json:"tlsMax,omitempty" xml:"tlsMax,omitempty"`
	// The minimum version of the TLS protocol. Down to TLS 1.0 is supported.
	//
	// example:
	//
	// TLS 1.0
	TlsMin *string `json:"tlsMin,omitempty" xml:"tlsMin,omitempty"`
	// The update timestamp.
	//
	// example:
	//
	// 1719386834548
	Updatetimestamp *int64 `json:"updatetimestamp,omitempty" xml:"updatetimestamp,omitempty"`
}

func (s GetDomainResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDomainResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBodyData) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *GetDomainResponseBodyData) GetCaCertIdentifier() *string {
	return s.CaCertIdentifier
}

func (s *GetDomainResponseBodyData) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *GetDomainResponseBodyData) GetCertName() *string {
	return s.CertName
}

func (s *GetDomainResponseBodyData) GetClientCACert() *string {
	return s.ClientCACert
}

func (s *GetDomainResponseBodyData) GetCreateFrom() *string {
	return s.CreateFrom
}

func (s *GetDomainResponseBodyData) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetDomainResponseBodyData) GetDefault() *bool {
	return s.Default
}

func (s *GetDomainResponseBodyData) GetDomainId() *string {
	return s.DomainId
}

func (s *GetDomainResponseBodyData) GetForceHttps() *bool {
	return s.ForceHttps
}

func (s *GetDomainResponseBodyData) GetHttp2Option() *string {
	return s.Http2Option
}

func (s *GetDomainResponseBodyData) GetIssuer() *string {
	return s.Issuer
}

func (s *GetDomainResponseBodyData) GetMTLSEnabled() *bool {
	return s.MTLSEnabled
}

func (s *GetDomainResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetDomainResponseBodyData) GetNotAfterTimstamp() *int64 {
	return s.NotAfterTimstamp
}

func (s *GetDomainResponseBodyData) GetNotBeforeTimestamp() *int64 {
	return s.NotBeforeTimestamp
}

func (s *GetDomainResponseBodyData) GetProtocol() *string {
	return s.Protocol
}

func (s *GetDomainResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetDomainResponseBodyData) GetSans() *string {
	return s.Sans
}

func (s *GetDomainResponseBodyData) GetStatisticsInfo() *GetDomainResponseBodyDataStatisticsInfo {
	return s.StatisticsInfo
}

func (s *GetDomainResponseBodyData) GetTlsCipherSuitesConfig() *TlsCipherSuitesConfig {
	return s.TlsCipherSuitesConfig
}

func (s *GetDomainResponseBodyData) GetTlsMax() *string {
	return s.TlsMax
}

func (s *GetDomainResponseBodyData) GetTlsMin() *string {
	return s.TlsMin
}

func (s *GetDomainResponseBodyData) GetUpdatetimestamp() *int64 {
	return s.Updatetimestamp
}

func (s *GetDomainResponseBodyData) SetAlgorithm(v string) *GetDomainResponseBodyData {
	s.Algorithm = &v
	return s
}

func (s *GetDomainResponseBodyData) SetCaCertIdentifier(v string) *GetDomainResponseBodyData {
	s.CaCertIdentifier = &v
	return s
}

func (s *GetDomainResponseBodyData) SetCertIdentifier(v string) *GetDomainResponseBodyData {
	s.CertIdentifier = &v
	return s
}

func (s *GetDomainResponseBodyData) SetCertName(v string) *GetDomainResponseBodyData {
	s.CertName = &v
	return s
}

func (s *GetDomainResponseBodyData) SetClientCACert(v string) *GetDomainResponseBodyData {
	s.ClientCACert = &v
	return s
}

func (s *GetDomainResponseBodyData) SetCreateFrom(v string) *GetDomainResponseBodyData {
	s.CreateFrom = &v
	return s
}

func (s *GetDomainResponseBodyData) SetCreateTimestamp(v int64) *GetDomainResponseBodyData {
	s.CreateTimestamp = &v
	return s
}

func (s *GetDomainResponseBodyData) SetDefault(v bool) *GetDomainResponseBodyData {
	s.Default = &v
	return s
}

func (s *GetDomainResponseBodyData) SetDomainId(v string) *GetDomainResponseBodyData {
	s.DomainId = &v
	return s
}

func (s *GetDomainResponseBodyData) SetForceHttps(v bool) *GetDomainResponseBodyData {
	s.ForceHttps = &v
	return s
}

func (s *GetDomainResponseBodyData) SetHttp2Option(v string) *GetDomainResponseBodyData {
	s.Http2Option = &v
	return s
}

func (s *GetDomainResponseBodyData) SetIssuer(v string) *GetDomainResponseBodyData {
	s.Issuer = &v
	return s
}

func (s *GetDomainResponseBodyData) SetMTLSEnabled(v bool) *GetDomainResponseBodyData {
	s.MTLSEnabled = &v
	return s
}

func (s *GetDomainResponseBodyData) SetName(v string) *GetDomainResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetDomainResponseBodyData) SetNotAfterTimstamp(v int64) *GetDomainResponseBodyData {
	s.NotAfterTimstamp = &v
	return s
}

func (s *GetDomainResponseBodyData) SetNotBeforeTimestamp(v int64) *GetDomainResponseBodyData {
	s.NotBeforeTimestamp = &v
	return s
}

func (s *GetDomainResponseBodyData) SetProtocol(v string) *GetDomainResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *GetDomainResponseBodyData) SetResourceGroupId(v string) *GetDomainResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *GetDomainResponseBodyData) SetSans(v string) *GetDomainResponseBodyData {
	s.Sans = &v
	return s
}

func (s *GetDomainResponseBodyData) SetStatisticsInfo(v *GetDomainResponseBodyDataStatisticsInfo) *GetDomainResponseBodyData {
	s.StatisticsInfo = v
	return s
}

func (s *GetDomainResponseBodyData) SetTlsCipherSuitesConfig(v *TlsCipherSuitesConfig) *GetDomainResponseBodyData {
	s.TlsCipherSuitesConfig = v
	return s
}

func (s *GetDomainResponseBodyData) SetTlsMax(v string) *GetDomainResponseBodyData {
	s.TlsMax = &v
	return s
}

func (s *GetDomainResponseBodyData) SetTlsMin(v string) *GetDomainResponseBodyData {
	s.TlsMin = &v
	return s
}

func (s *GetDomainResponseBodyData) SetUpdatetimestamp(v int64) *GetDomainResponseBodyData {
	s.Updatetimestamp = &v
	return s
}

func (s *GetDomainResponseBodyData) Validate() error {
	if s.StatisticsInfo != nil {
		if err := s.StatisticsInfo.Validate(); err != nil {
			return err
		}
	}
	if s.TlsCipherSuitesConfig != nil {
		if err := s.TlsCipherSuitesConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDomainResponseBodyDataStatisticsInfo struct {
	// The resource statistics.
	ResourceStatistics []*ResourceStatistic `json:"resourceStatistics,omitempty" xml:"resourceStatistics,omitempty" type:"Repeated"`
	// The total number of resources.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s GetDomainResponseBodyDataStatisticsInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDomainResponseBodyDataStatisticsInfo) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBodyDataStatisticsInfo) GetResourceStatistics() []*ResourceStatistic {
	return s.ResourceStatistics
}

func (s *GetDomainResponseBodyDataStatisticsInfo) GetTotalCount() *string {
	return s.TotalCount
}

func (s *GetDomainResponseBodyDataStatisticsInfo) SetResourceStatistics(v []*ResourceStatistic) *GetDomainResponseBodyDataStatisticsInfo {
	s.ResourceStatistics = v
	return s
}

func (s *GetDomainResponseBodyDataStatisticsInfo) SetTotalCount(v string) *GetDomainResponseBodyDataStatisticsInfo {
	s.TotalCount = &v
	return s
}

func (s *GetDomainResponseBodyDataStatisticsInfo) Validate() error {
	if s.ResourceStatistics != nil {
		for _, item := range s.ResourceStatistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
