// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertDetail(v *DescribeDomainDetailResponseBodyCertDetail) *DescribeDomainDetailResponseBody
	GetCertDetail() *DescribeDomainDetailResponseBodyCertDetail
	SetCname(v string) *DescribeDomainDetailResponseBody
	GetCname() *string
	SetDomain(v string) *DescribeDomainDetailResponseBody
	GetDomain() *string
	SetDomainId(v string) *DescribeDomainDetailResponseBody
	GetDomainId() *string
	SetListen(v *DescribeDomainDetailResponseBodyListen) *DescribeDomainDetailResponseBody
	GetListen() *DescribeDomainDetailResponseBodyListen
	SetRedirect(v *DescribeDomainDetailResponseBodyRedirect) *DescribeDomainDetailResponseBody
	GetRedirect() *DescribeDomainDetailResponseBodyRedirect
	SetRequestId(v string) *DescribeDomainDetailResponseBody
	GetRequestId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDomainDetailResponseBody
	GetResourceManagerResourceGroupId() *string
	SetSM2CertDetail(v *DescribeDomainDetailResponseBodySM2CertDetail) *DescribeDomainDetailResponseBody
	GetSM2CertDetail() *DescribeDomainDetailResponseBodySM2CertDetail
	SetStatus(v int32) *DescribeDomainDetailResponseBody
	GetStatus() *int32
}

type DescribeDomainDetailResponseBody struct {
	// The SSL certificate details.
	CertDetail *DescribeDomainDetailResponseBodyCertDetail `json:"CertDetail,omitempty" xml:"CertDetail,omitempty" type:"Struct"`
	// The CNAME assigned by WAF to the domain name.
	//
	// example:
	//
	// xxxxxcvdaf.****.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The domain name.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The domain name ID.
	//
	// example:
	//
	// www.aliyundoc.com-waf
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The listener configuration.
	Listen *DescribeDomainDetailResponseBodyListen `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	// The forwarding configuration.
	Redirect *DescribeDomainDetailResponseBodyRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BAEF9CA9-66A0-533E-BD09-5D5D7AA8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud resource group ID.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The SM2 certificate information.
	SM2CertDetail *DescribeDomainDetailResponseBodySM2CertDetail `json:"SM2CertDetail,omitempty" xml:"SM2CertDetail,omitempty" type:"Struct"`
	// The status of the domain name. Valid values:
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDomainDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBody) GetCertDetail() *DescribeDomainDetailResponseBodyCertDetail {
	return s.CertDetail
}

func (s *DescribeDomainDetailResponseBody) GetCname() *string {
	return s.Cname
}

func (s *DescribeDomainDetailResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainDetailResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *DescribeDomainDetailResponseBody) GetListen() *DescribeDomainDetailResponseBodyListen {
	return s.Listen
}

func (s *DescribeDomainDetailResponseBody) GetRedirect() *DescribeDomainDetailResponseBodyRedirect {
	return s.Redirect
}

func (s *DescribeDomainDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainDetailResponseBody) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDomainDetailResponseBody) GetSM2CertDetail() *DescribeDomainDetailResponseBodySM2CertDetail {
	return s.SM2CertDetail
}

func (s *DescribeDomainDetailResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDomainDetailResponseBody) SetCertDetail(v *DescribeDomainDetailResponseBodyCertDetail) *DescribeDomainDetailResponseBody {
	s.CertDetail = v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetCname(v string) *DescribeDomainDetailResponseBody {
	s.Cname = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetDomain(v string) *DescribeDomainDetailResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetDomainId(v string) *DescribeDomainDetailResponseBody {
	s.DomainId = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetListen(v *DescribeDomainDetailResponseBodyListen) *DescribeDomainDetailResponseBody {
	s.Listen = v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetRedirect(v *DescribeDomainDetailResponseBodyRedirect) *DescribeDomainDetailResponseBody {
	s.Redirect = v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetRequestId(v string) *DescribeDomainDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetResourceManagerResourceGroupId(v string) *DescribeDomainDetailResponseBody {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetSM2CertDetail(v *DescribeDomainDetailResponseBodySM2CertDetail) *DescribeDomainDetailResponseBody {
	s.SM2CertDetail = v
	return s
}

func (s *DescribeDomainDetailResponseBody) SetStatus(v int32) *DescribeDomainDetailResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDomainDetailResponseBody) Validate() error {
	if s.CertDetail != nil {
		if err := s.CertDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Listen != nil {
		if err := s.Listen.Validate(); err != nil {
			return err
		}
	}
	if s.Redirect != nil {
		if err := s.Redirect.Validate(); err != nil {
			return err
		}
	}
	if s.SM2CertDetail != nil {
		if err := s.SM2CertDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDomainDetailResponseBodyCertDetail struct {
	// The common name (CN).
	//
	// example:
	//
	// test.aliyundoc.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The time when the certificate expires. The value is a UNIX timestamp (UTC). Unit: milliseconds.
	//
	// example:
	//
	// 1685590400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The SSL certificate ID.
	//
	// example:
	//
	// 123-cn-hangzhou
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The certificate name.
	//
	// example:
	//
	// test-cert-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// All domain names that are bound to the certificate.
	Sans []*string `json:"Sans,omitempty" xml:"Sans,omitempty" type:"Repeated"`
	// The effective period of the certificate. The value is in the format of a UNIX timestamp (UTC). Unit: milliseconds.
	//
	// example:
	//
	// 1677772800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainDetailResponseBodyCertDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyCertDetail) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyCertDetail) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeDomainDetailResponseBodyCertDetail) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainDetailResponseBodyCertDetail) GetId() *string {
	return s.Id
}

func (s *DescribeDomainDetailResponseBodyCertDetail) GetName() *string {
	return s.Name
}

func (s *DescribeDomainDetailResponseBodyCertDetail) GetSans() []*string {
	return s.Sans
}

func (s *DescribeDomainDetailResponseBodyCertDetail) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainDetailResponseBodyCertDetail) SetCommonName(v string) *DescribeDomainDetailResponseBodyCertDetail {
	s.CommonName = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyCertDetail) SetEndTime(v int64) *DescribeDomainDetailResponseBodyCertDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyCertDetail) SetId(v string) *DescribeDomainDetailResponseBodyCertDetail {
	s.Id = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyCertDetail) SetName(v string) *DescribeDomainDetailResponseBodyCertDetail {
	s.Name = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyCertDetail) SetSans(v []*string) *DescribeDomainDetailResponseBodyCertDetail {
	s.Sans = v
	return s
}

func (s *DescribeDomainDetailResponseBodyCertDetail) SetStartTime(v int64) *DescribeDomainDetailResponseBodyCertDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyCertDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainDetailResponseBodyListen struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 123
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The type of the cipher suite. Valid values:
	//
	// example:
	//
	// 2
	CipherSuite *int64 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suites.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Indicates whether TLS 1.3 is supported. Valid values:
	//
	// example:
	//
	// true
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// Indicates whether the exclusive IP address feature is enabled. Valid values:
	//
	// example:
	//
	// true
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Indicates whether HTTPS forced redirect is enabled. Valid values:
	//
	// example:
	//
	// true
	FocusHttps *bool `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	// Indicates whether HSTS includes subdomains. Valid values:
	HstsIncludeSubDomain *bool `json:"HstsIncludeSubDomain,omitempty" xml:"HstsIncludeSubDomain,omitempty"`
	// The HSTS expiration time. Unit: seconds.
	//
	// example:
	//
	// 365000
	HstsMaxAge *int64 `json:"HstsMaxAge,omitempty" xml:"HstsMaxAge,omitempty"`
	// Indicates whether HSTS preloading is enabled. This feature is disabled by default. Valid values:
	//
	// example:
	//
	// false
	HstsPreload *bool `json:"HstsPreload,omitempty" xml:"HstsPreload,omitempty"`
	// Indicates whether HTTP/2 is enabled. Valid values:
	//
	// example:
	//
	// true
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The listening port for the HTTP protocol.
	HttpPorts []*int64 `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	// The listening port for the HTTPS protocol.
	HttpsPorts []*int64 `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	// Indicates whether IPv6 is enabled. Valid values:
	//
	// example:
	//
	// true
	IPv6Enabled *bool `json:"IPv6Enabled,omitempty" xml:"IPv6Enabled,omitempty"`
	// The type of protection resource to use. Valid values:
	//
	// example:
	//
	// share
	ProtectionResource *string `json:"ProtectionResource,omitempty" xml:"ProtectionResource,omitempty"`
	// Indicates whether only SM client access is allowed. This parameter is used only when SM2Enable is set to true.
	//
	// example:
	//
	// true
	SM2AccessOnly *bool `json:"SM2AccessOnly,omitempty" xml:"SM2AccessOnly,omitempty"`
	// The ID of the China Encryption Standard (SM) certificate to add. This parameter is used only when SM2Enable is set to true.
	//
	// example:
	//
	// 123-cn-hangzhou
	SM2CertId *string `json:"SM2CertId,omitempty" xml:"SM2CertId,omitempty"`
	// Indicates whether the China Encryption Standard (SM) certificate is enabled. Valid values:
	//
	// example:
	//
	// true
	SM2Enabled *bool `json:"SM2Enabled,omitempty" xml:"SM2Enabled,omitempty"`
	// The TLS version. Valid values:
	//
	// example:
	//
	// tlsv1.2
	TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
	// The method that WAF uses to obtain the originating IP address of the client. Valid values:
	//
	// example:
	//
	// 2
	XffHeaderMode *int64 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// The list of custom header fields used to obtain the client IP address.
	XffHeaders []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
}

func (s DescribeDomainDetailResponseBodyListen) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyListen) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyListen) GetCertId() *string {
	return s.CertId
}

func (s *DescribeDomainDetailResponseBodyListen) GetCipherSuite() *int64 {
	return s.CipherSuite
}

func (s *DescribeDomainDetailResponseBodyListen) GetCustomCiphers() []*string {
	return s.CustomCiphers
}

func (s *DescribeDomainDetailResponseBodyListen) GetEnableTLSv3() *bool {
	return s.EnableTLSv3
}

func (s *DescribeDomainDetailResponseBodyListen) GetExclusiveIp() *bool {
	return s.ExclusiveIp
}

func (s *DescribeDomainDetailResponseBodyListen) GetFocusHttps() *bool {
	return s.FocusHttps
}

func (s *DescribeDomainDetailResponseBodyListen) GetHstsIncludeSubDomain() *bool {
	return s.HstsIncludeSubDomain
}

func (s *DescribeDomainDetailResponseBodyListen) GetHstsMaxAge() *int64 {
	return s.HstsMaxAge
}

func (s *DescribeDomainDetailResponseBodyListen) GetHstsPreload() *bool {
	return s.HstsPreload
}

func (s *DescribeDomainDetailResponseBodyListen) GetHttp2Enabled() *bool {
	return s.Http2Enabled
}

func (s *DescribeDomainDetailResponseBodyListen) GetHttpPorts() []*int64 {
	return s.HttpPorts
}

func (s *DescribeDomainDetailResponseBodyListen) GetHttpsPorts() []*int64 {
	return s.HttpsPorts
}

func (s *DescribeDomainDetailResponseBodyListen) GetIPv6Enabled() *bool {
	return s.IPv6Enabled
}

func (s *DescribeDomainDetailResponseBodyListen) GetProtectionResource() *string {
	return s.ProtectionResource
}

func (s *DescribeDomainDetailResponseBodyListen) GetSM2AccessOnly() *bool {
	return s.SM2AccessOnly
}

func (s *DescribeDomainDetailResponseBodyListen) GetSM2CertId() *string {
	return s.SM2CertId
}

func (s *DescribeDomainDetailResponseBodyListen) GetSM2Enabled() *bool {
	return s.SM2Enabled
}

func (s *DescribeDomainDetailResponseBodyListen) GetTLSVersion() *string {
	return s.TLSVersion
}

func (s *DescribeDomainDetailResponseBodyListen) GetXffHeaderMode() *int64 {
	return s.XffHeaderMode
}

func (s *DescribeDomainDetailResponseBodyListen) GetXffHeaders() []*string {
	return s.XffHeaders
}

func (s *DescribeDomainDetailResponseBodyListen) SetCertId(v string) *DescribeDomainDetailResponseBodyListen {
	s.CertId = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetCipherSuite(v int64) *DescribeDomainDetailResponseBodyListen {
	s.CipherSuite = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetCustomCiphers(v []*string) *DescribeDomainDetailResponseBodyListen {
	s.CustomCiphers = v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetEnableTLSv3(v bool) *DescribeDomainDetailResponseBodyListen {
	s.EnableTLSv3 = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetExclusiveIp(v bool) *DescribeDomainDetailResponseBodyListen {
	s.ExclusiveIp = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetFocusHttps(v bool) *DescribeDomainDetailResponseBodyListen {
	s.FocusHttps = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetHstsIncludeSubDomain(v bool) *DescribeDomainDetailResponseBodyListen {
	s.HstsIncludeSubDomain = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetHstsMaxAge(v int64) *DescribeDomainDetailResponseBodyListen {
	s.HstsMaxAge = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetHstsPreload(v bool) *DescribeDomainDetailResponseBodyListen {
	s.HstsPreload = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetHttp2Enabled(v bool) *DescribeDomainDetailResponseBodyListen {
	s.Http2Enabled = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetHttpPorts(v []*int64) *DescribeDomainDetailResponseBodyListen {
	s.HttpPorts = v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetHttpsPorts(v []*int64) *DescribeDomainDetailResponseBodyListen {
	s.HttpsPorts = v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetIPv6Enabled(v bool) *DescribeDomainDetailResponseBodyListen {
	s.IPv6Enabled = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetProtectionResource(v string) *DescribeDomainDetailResponseBodyListen {
	s.ProtectionResource = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetSM2AccessOnly(v bool) *DescribeDomainDetailResponseBodyListen {
	s.SM2AccessOnly = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetSM2CertId(v string) *DescribeDomainDetailResponseBodyListen {
	s.SM2CertId = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetSM2Enabled(v bool) *DescribeDomainDetailResponseBodyListen {
	s.SM2Enabled = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetTLSVersion(v string) *DescribeDomainDetailResponseBodyListen {
	s.TLSVersion = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetXffHeaderMode(v int64) *DescribeDomainDetailResponseBodyListen {
	s.XffHeaderMode = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) SetXffHeaders(v []*string) *DescribeDomainDetailResponseBodyListen {
	s.XffHeaders = v
	return s
}

func (s *DescribeDomainDetailResponseBodyListen) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainDetailResponseBodyRedirect struct {
	// The list of backup origin server IP addresses or back-to-origin domain names for the domain name.
	BackUpBackendList []*string `json:"BackUpBackendList,omitempty" xml:"BackUpBackendList,omitempty" type:"Repeated"`
	// The list of origin server IP addresses or back-to-origin domain names for the domain name.
	BackendList []*string `json:"BackendList,omitempty" xml:"BackendList,omitempty" type:"Repeated"`
	// The custom port configuration. By default, the port is the same as the listener port.
	BackendPorts []*DescribeDomainDetailResponseBodyRedirectBackendPorts `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The back-to-origin addresses of the domain name.
	Backends []*DescribeDomainDetailResponseBodyRedirectBackends `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The secondary back-to-origin addresses of the domain name.
	BackupBackends []*DescribeDomainDetailResponseBodyRedirectBackupBackends `json:"BackupBackends,omitempty" xml:"BackupBackends,omitempty" type:"Repeated"`
	// The connection timeout period. Unit: seconds.
	//
	// example:
	//
	// 120
	ConnectTimeout *int32 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// Indicates whether forced HTTP back-to-origin is enabled. Valid values:
	//
	// example:
	//
	// true
	FocusHttpBackend *bool `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	// The HTTP/2 back-to-origin setting.
	//
	// example:
	//
	// true
	Http2Origin *bool `json:"Http2Origin,omitempty" xml:"Http2Origin,omitempty"`
	// The number of concurrent connections for HTTP/2 back-to-origin.
	//
	// example:
	//
	// 128
	Http2OriginMaxConcurrency *int32 `json:"Http2OriginMaxConcurrency,omitempty" xml:"Http2OriginMaxConcurrency,omitempty"`
	// Indicates whether persistent connections are enabled. Valid values:
	//
	// example:
	//
	// true
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// The number of requests that reuse a persistent connection. Valid values: 60 to 1000.
	//
	// example:
	//
	// 1000
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The idle timeout period for persistent connections. Valid values: 1 to 60. Default value: 15. Unit: seconds.
	//
	// example:
	//
	// 15
	KeepaliveTimeout *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The load balancing algorithm used for back-to-origin. Valid values:
	//
	// example:
	//
	// iphash
	Loadbalance *string `json:"Loadbalance,omitempty" xml:"Loadbalance,omitempty"`
	// The maximum request body size. Valid values: 2 to 10. Default value: 2. Unit: GB.
	//
	// example:
	//
	// 2
	MaxBodySize *int32 `json:"MaxBodySize,omitempty" xml:"MaxBodySize,omitempty"`
	// Indicates whether the feature for preserving the originating IP address of the client is enabled.
	//
	// example:
	//
	// false
	ProxyProtocol *bool `json:"ProxyProtocol,omitempty" xml:"ProxyProtocol,omitempty"`
	// The read timeout period. Unit: seconds.
	//
	// example:
	//
	// 200
	ReadTimeout *int32 `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	// The traffic mark fields and values of the domain name, which are used to mark traffic processed by WAF.
	RequestHeaders []*DescribeDomainDetailResponseBodyRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// Indicates whether WAF retries when back-to-origin fails. Valid values:
	//
	// example:
	//
	// true
	Retry *bool `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// Indicates whether back-to-origin SNI is enabled. Valid values:
	//
	// example:
	//
	// true
	SniEnabled *bool `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	// The value of the custom SNI extension field.
	//
	// example:
	//
	// www.aliyundoc.com
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// Indicates whether WAF is allowed to overwrite WL-Proxy-Client-IP. Valid values:
	//
	// example:
	//
	// true
	WLProxyClientIp *bool `json:"WLProxyClientIp,omitempty" xml:"WLProxyClientIp,omitempty"`
	// Indicates whether WAF is allowed to overwrite Web-Server-Type. Valid values:
	//
	// example:
	//
	// true
	WebServerType *bool `json:"WebServerType,omitempty" xml:"WebServerType,omitempty"`
	// The write timeout period. Unit: seconds.
	//
	// example:
	//
	// 200
	WriteTimeout *int32 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
	// Indicates whether WAF is allowed to overwrite X-Client-IP. Valid values:
	//
	// example:
	//
	// true
	XClientIp *bool `json:"XClientIp,omitempty" xml:"XClientIp,omitempty"`
	// Indicates whether WAF is allowed to overwrite X-True-IP. Valid values:
	//
	// example:
	//
	// true
	XTrueIp *bool `json:"XTrueIp,omitempty" xml:"XTrueIp,omitempty"`
	// Indicates whether X-Forward-For-Proto is used to pass the protocol used by WAF. Valid values:
	//
	// example:
	//
	// true
	XffProto *bool `json:"XffProto,omitempty" xml:"XffProto,omitempty"`
}

func (s DescribeDomainDetailResponseBodyRedirect) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyRedirect) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetBackUpBackendList() []*string {
	return s.BackUpBackendList
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetBackendList() []*string {
	return s.BackendList
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetBackendPorts() []*DescribeDomainDetailResponseBodyRedirectBackendPorts {
	return s.BackendPorts
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetBackends() []*DescribeDomainDetailResponseBodyRedirectBackends {
	return s.Backends
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetBackupBackends() []*DescribeDomainDetailResponseBodyRedirectBackupBackends {
	return s.BackupBackends
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetConnectTimeout() *int32 {
	return s.ConnectTimeout
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetFocusHttpBackend() *bool {
	return s.FocusHttpBackend
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetHttp2Origin() *bool {
	return s.Http2Origin
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetHttp2OriginMaxConcurrency() *int32 {
	return s.Http2OriginMaxConcurrency
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetKeepalive() *bool {
	return s.Keepalive
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetKeepaliveRequests() *int32 {
	return s.KeepaliveRequests
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetKeepaliveTimeout() *int32 {
	return s.KeepaliveTimeout
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetLoadbalance() *string {
	return s.Loadbalance
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetMaxBodySize() *int32 {
	return s.MaxBodySize
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetProxyProtocol() *bool {
	return s.ProxyProtocol
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetReadTimeout() *int32 {
	return s.ReadTimeout
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetRequestHeaders() []*DescribeDomainDetailResponseBodyRedirectRequestHeaders {
	return s.RequestHeaders
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetRetry() *bool {
	return s.Retry
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetSniEnabled() *bool {
	return s.SniEnabled
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetSniHost() *string {
	return s.SniHost
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetWLProxyClientIp() *bool {
	return s.WLProxyClientIp
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetWebServerType() *bool {
	return s.WebServerType
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetWriteTimeout() *int32 {
	return s.WriteTimeout
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetXClientIp() *bool {
	return s.XClientIp
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetXTrueIp() *bool {
	return s.XTrueIp
}

func (s *DescribeDomainDetailResponseBodyRedirect) GetXffProto() *bool {
	return s.XffProto
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetBackUpBackendList(v []*string) *DescribeDomainDetailResponseBodyRedirect {
	s.BackUpBackendList = v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetBackendList(v []*string) *DescribeDomainDetailResponseBodyRedirect {
	s.BackendList = v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetBackendPorts(v []*DescribeDomainDetailResponseBodyRedirectBackendPorts) *DescribeDomainDetailResponseBodyRedirect {
	s.BackendPorts = v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetBackends(v []*DescribeDomainDetailResponseBodyRedirectBackends) *DescribeDomainDetailResponseBodyRedirect {
	s.Backends = v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetBackupBackends(v []*DescribeDomainDetailResponseBodyRedirectBackupBackends) *DescribeDomainDetailResponseBodyRedirect {
	s.BackupBackends = v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetConnectTimeout(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.ConnectTimeout = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetFocusHttpBackend(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.FocusHttpBackend = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetHttp2Origin(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.Http2Origin = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetHttp2OriginMaxConcurrency(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.Http2OriginMaxConcurrency = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetKeepalive(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.Keepalive = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetKeepaliveRequests(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.KeepaliveRequests = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetKeepaliveTimeout(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.KeepaliveTimeout = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetLoadbalance(v string) *DescribeDomainDetailResponseBodyRedirect {
	s.Loadbalance = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetMaxBodySize(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.MaxBodySize = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetProxyProtocol(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.ProxyProtocol = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetReadTimeout(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.ReadTimeout = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetRequestHeaders(v []*DescribeDomainDetailResponseBodyRedirectRequestHeaders) *DescribeDomainDetailResponseBodyRedirect {
	s.RequestHeaders = v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetRetry(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.Retry = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetSniEnabled(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.SniEnabled = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetSniHost(v string) *DescribeDomainDetailResponseBodyRedirect {
	s.SniHost = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetWLProxyClientIp(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.WLProxyClientIp = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetWebServerType(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.WebServerType = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetWriteTimeout(v int32) *DescribeDomainDetailResponseBodyRedirect {
	s.WriteTimeout = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetXClientIp(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.XClientIp = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetXTrueIp(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.XTrueIp = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) SetXffProto(v bool) *DescribeDomainDetailResponseBodyRedirect {
	s.XffProto = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirect) Validate() error {
	if s.BackendPorts != nil {
		for _, item := range s.BackendPorts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Backends != nil {
		for _, item := range s.Backends {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.BackupBackends != nil {
		for _, item := range s.BackupBackends {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RequestHeaders != nil {
		for _, item := range s.RequestHeaders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainDetailResponseBodyRedirectBackendPorts struct {
	// The back-to-origin port.
	//
	// example:
	//
	// 80
	BackendPort *int32 `json:"BackendPort,omitempty" xml:"BackendPort,omitempty"`
	// The listener port.
	//
	// example:
	//
	// 80
	ListenPort *int32 `json:"ListenPort,omitempty" xml:"ListenPort,omitempty"`
	// The protocol type of the listener port. Valid values:
	//
	// example:
	//
	// http
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s DescribeDomainDetailResponseBodyRedirectBackendPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyRedirectBackendPorts) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyRedirectBackendPorts) GetBackendPort() *int32 {
	return s.BackendPort
}

func (s *DescribeDomainDetailResponseBodyRedirectBackendPorts) GetListenPort() *int32 {
	return s.ListenPort
}

func (s *DescribeDomainDetailResponseBodyRedirectBackendPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeDomainDetailResponseBodyRedirectBackendPorts) SetBackendPort(v int32) *DescribeDomainDetailResponseBodyRedirectBackendPorts {
	s.BackendPort = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirectBackendPorts) SetListenPort(v int32) *DescribeDomainDetailResponseBodyRedirectBackendPorts {
	s.ListenPort = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirectBackendPorts) SetProtocol(v string) *DescribeDomainDetailResponseBodyRedirectBackendPorts {
	s.Protocol = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirectBackendPorts) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainDetailResponseBodyRedirectBackends struct {
	// The IP address or domain name of the origin server for the domain name.
	//
	// example:
	//
	// 1.1.XX.XX
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
}

func (s DescribeDomainDetailResponseBodyRedirectBackends) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyRedirectBackends) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyRedirectBackends) GetBackend() *string {
	return s.Backend
}

func (s *DescribeDomainDetailResponseBodyRedirectBackends) SetBackend(v string) *DescribeDomainDetailResponseBodyRedirectBackends {
	s.Backend = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirectBackends) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainDetailResponseBodyRedirectBackupBackends struct {
	// The IP address or domain name of the secondary origin server for the domain name.
	//
	// example:
	//
	// [
	//
	//     "1.1.XX.XX",
	//
	//     "2.2.XX.XX"
	//
	// ]
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
}

func (s DescribeDomainDetailResponseBodyRedirectBackupBackends) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyRedirectBackupBackends) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyRedirectBackupBackends) GetBackend() *string {
	return s.Backend
}

func (s *DescribeDomainDetailResponseBodyRedirectBackupBackends) SetBackend(v string) *DescribeDomainDetailResponseBodyRedirectBackupBackends {
	s.Backend = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirectBackupBackends) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainDetailResponseBodyRedirectRequestHeaders struct {
	// The custom request header field.
	//
	// example:
	//
	// aaa
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom request header field.
	//
	// example:
	//
	// bbb
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDomainDetailResponseBodyRedirectRequestHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailResponseBodyRedirectRequestHeaders) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodyRedirectRequestHeaders) GetKey() *string {
	return s.Key
}

func (s *DescribeDomainDetailResponseBodyRedirectRequestHeaders) GetValue() *string {
	return s.Value
}

func (s *DescribeDomainDetailResponseBodyRedirectRequestHeaders) SetKey(v string) *DescribeDomainDetailResponseBodyRedirectRequestHeaders {
	s.Key = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirectRequestHeaders) SetValue(v string) *DescribeDomainDetailResponseBodyRedirectRequestHeaders {
	s.Value = &v
	return s
}

func (s *DescribeDomainDetailResponseBodyRedirectRequestHeaders) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainDetailResponseBodySM2CertDetail struct {
	// The common name (CN).
	//
	// example:
	//
	// test.aliyundoc.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The time when the certificate expires. The value is a UNIX timestamp (UTC). Unit: milliseconds.
	//
	// example:
	//
	// 1665590400000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The SSL certificate ID.
	//
	// example:
	//
	// 123-cn-hangzhou
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The certificate name.
	//
	// example:
	//
	// test-sm2-cert-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// All domain names that are bound to the certificate.
	Sans []*string `json:"Sans,omitempty" xml:"Sans,omitempty" type:"Repeated"`
	// The effective period of the certificate. The value is in the format of a UNIX timestamp (UTC). Unit: milliseconds.
	//
	// example:
	//
	// 1657551525000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainDetailResponseBodySM2CertDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainDetailResponseBodySM2CertDetail) GoString() string {
	return s.String()
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) GetId() *string {
	return s.Id
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) GetName() *string {
	return s.Name
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) GetSans() []*string {
	return s.Sans
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) SetCommonName(v string) *DescribeDomainDetailResponseBodySM2CertDetail {
	s.CommonName = &v
	return s
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) SetEndTime(v int64) *DescribeDomainDetailResponseBodySM2CertDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) SetId(v string) *DescribeDomainDetailResponseBodySM2CertDetail {
	s.Id = &v
	return s
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) SetName(v string) *DescribeDomainDetailResponseBodySM2CertDetail {
	s.Name = &v
	return s
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) SetSans(v []*string) *DescribeDomainDetailResponseBodySM2CertDetail {
	s.Sans = v
	return s
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) SetStartTime(v int64) *DescribeDomainDetailResponseBodySM2CertDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainDetailResponseBodySM2CertDetail) Validate() error {
	return dara.Validate(s)
}
