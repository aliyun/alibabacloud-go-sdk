// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *ModifyDomainRequest
	GetAccessType() *string
	SetDomain(v string) *ModifyDomainRequest
	GetDomain() *string
	SetDomainId(v string) *ModifyDomainRequest
	GetDomainId() *string
	SetInstanceId(v string) *ModifyDomainRequest
	GetInstanceId() *string
	SetListen(v *ModifyDomainRequestListen) *ModifyDomainRequest
	GetListen() *ModifyDomainRequestListen
	SetRedirect(v *ModifyDomainRequestRedirect) *ModifyDomainRequest
	GetRedirect() *ModifyDomainRequestRedirect
	SetRegionId(v string) *ModifyDomainRequest
	GetRegionId() *string
}

type ModifyDomainRequest struct {
	// The access type of the WAF instance. Valid values:
	//
	// example:
	//
	// share
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The domain name to operate on.
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
	// The ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The listening configuration.
	//
	// This parameter is required.
	Listen *ModifyDomainRequestListen `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	// The forwarding configuration.
	//
	// This parameter is required.
	Redirect *ModifyDomainRequestRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// The region where the WAF instance resides. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainRequest) GetAccessType() *string {
	return s.AccessType
}

func (s *ModifyDomainRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyDomainRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *ModifyDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDomainRequest) GetListen() *ModifyDomainRequestListen {
	return s.Listen
}

func (s *ModifyDomainRequest) GetRedirect() *ModifyDomainRequestRedirect {
	return s.Redirect
}

func (s *ModifyDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDomainRequest) SetAccessType(v string) *ModifyDomainRequest {
	s.AccessType = &v
	return s
}

func (s *ModifyDomainRequest) SetDomain(v string) *ModifyDomainRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainRequest) SetDomainId(v string) *ModifyDomainRequest {
	s.DomainId = &v
	return s
}

func (s *ModifyDomainRequest) SetInstanceId(v string) *ModifyDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDomainRequest) SetListen(v *ModifyDomainRequestListen) *ModifyDomainRequest {
	s.Listen = v
	return s
}

func (s *ModifyDomainRequest) SetRedirect(v *ModifyDomainRequestRedirect) *ModifyDomainRequest {
	s.Redirect = v
	return s
}

func (s *ModifyDomainRequest) SetRegionId(v string) *ModifyDomainRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDomainRequest) Validate() error {
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
	return nil
}

type ModifyDomainRequestListen struct {
	// The ID of the certificate to add.
	//
	// example:
	//
	// 123
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The type of cipher suite to add. This parameter is available only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
	//
	// example:
	//
	// 2
	CipherSuite *int32 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The specific custom cipher suites to add. This parameter is available only when **CipherSuite*	- is set to **99**.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Specifies whether to support TLS 1.3. Valid values:
	//
	// example:
	//
	// true
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// Specifies whether to enable an exclusive IP address. This parameter is available only when **IPv6Enabled*	- is set to false and **ProtectionResource*	- is set to **share**, which indicates that a shared cluster is used. Valid values:
	//
	// example:
	//
	// true
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Specifies whether to enable forced HTTPS redirect. This parameter is available only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS, and **HttpPorts*	- is empty, which indicates that the domain name does not use HTTP. Valid values:
	//
	// example:
	//
	// true
	FocusHttps *bool `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	// Specifies whether HSTS includes subdomains. Valid values:
	//
	// example:
	//
	// false
	HstsIncludeSubDomain *bool `json:"HstsIncludeSubDomain,omitempty" xml:"HstsIncludeSubDomain,omitempty"`
	// The HSTS expiration time. Unit: seconds.
	//
	// example:
	//
	// 365000
	HstsMaxAge *int64 `json:"HstsMaxAge,omitempty" xml:"HstsMaxAge,omitempty"`
	// Specifies whether to enable HSTS preloading. This feature is disabled by default. Valid values:
	//
	// example:
	//
	// false
	HstsPreload *bool `json:"HstsPreload,omitempty" xml:"HstsPreload,omitempty"`
	// Specifies whether to enable HTTP/2. This parameter is available only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
	//
	// example:
	//
	// true
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The listening ports for HTTP. Use the [**port1,port2,...**] format.
	HttpPorts []*int32 `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	// The listening ports for HTTPS. Use the [**port1,port2,...**] format.
	HttpsPorts []*int32 `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	// Specifies whether to enable IPv6. Valid values:
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
	// Specifies whether to allow only SM2 client access. This parameter is available only when SM2Enable is set to true.
	//
	// example:
	//
	// true
	SM2AccessOnly *bool `json:"SM2AccessOnly,omitempty" xml:"SM2AccessOnly,omitempty"`
	// The ID of the SM2 certificate to add. This parameter is available only when SM2Enable is set to true.
	//
	// example:
	//
	// 123-cn-hangzhou
	SM2CertId *string `json:"SM2CertId,omitempty" xml:"SM2CertId,omitempty"`
	// Specifies whether to enable SM2 certificates.
	//
	// example:
	//
	// true
	SM2Enabled *bool `json:"SM2Enabled,omitempty" xml:"SM2Enabled,omitempty"`
	// The TLS version to add. This parameter is available only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
	//
	// example:
	//
	// tlsv1
	TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
	// The method that WAF uses to obtain the originating IP address of the client. Valid values:
	//
	// example:
	//
	// 2
	XffHeaderMode *int32 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// The list of custom header fields used to obtain the client IP address. Use the [**"header1","header2",...**] format.
	XffHeaders []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
}

func (s ModifyDomainRequestListen) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainRequestListen) GoString() string {
	return s.String()
}

func (s *ModifyDomainRequestListen) GetCertId() *string {
	return s.CertId
}

func (s *ModifyDomainRequestListen) GetCipherSuite() *int32 {
	return s.CipherSuite
}

func (s *ModifyDomainRequestListen) GetCustomCiphers() []*string {
	return s.CustomCiphers
}

func (s *ModifyDomainRequestListen) GetEnableTLSv3() *bool {
	return s.EnableTLSv3
}

func (s *ModifyDomainRequestListen) GetExclusiveIp() *bool {
	return s.ExclusiveIp
}

func (s *ModifyDomainRequestListen) GetFocusHttps() *bool {
	return s.FocusHttps
}

func (s *ModifyDomainRequestListen) GetHstsIncludeSubDomain() *bool {
	return s.HstsIncludeSubDomain
}

func (s *ModifyDomainRequestListen) GetHstsMaxAge() *int64 {
	return s.HstsMaxAge
}

func (s *ModifyDomainRequestListen) GetHstsPreload() *bool {
	return s.HstsPreload
}

func (s *ModifyDomainRequestListen) GetHttp2Enabled() *bool {
	return s.Http2Enabled
}

func (s *ModifyDomainRequestListen) GetHttpPorts() []*int32 {
	return s.HttpPorts
}

func (s *ModifyDomainRequestListen) GetHttpsPorts() []*int32 {
	return s.HttpsPorts
}

func (s *ModifyDomainRequestListen) GetIPv6Enabled() *bool {
	return s.IPv6Enabled
}

func (s *ModifyDomainRequestListen) GetProtectionResource() *string {
	return s.ProtectionResource
}

func (s *ModifyDomainRequestListen) GetSM2AccessOnly() *bool {
	return s.SM2AccessOnly
}

func (s *ModifyDomainRequestListen) GetSM2CertId() *string {
	return s.SM2CertId
}

func (s *ModifyDomainRequestListen) GetSM2Enabled() *bool {
	return s.SM2Enabled
}

func (s *ModifyDomainRequestListen) GetTLSVersion() *string {
	return s.TLSVersion
}

func (s *ModifyDomainRequestListen) GetXffHeaderMode() *int32 {
	return s.XffHeaderMode
}

func (s *ModifyDomainRequestListen) GetXffHeaders() []*string {
	return s.XffHeaders
}

func (s *ModifyDomainRequestListen) SetCertId(v string) *ModifyDomainRequestListen {
	s.CertId = &v
	return s
}

func (s *ModifyDomainRequestListen) SetCipherSuite(v int32) *ModifyDomainRequestListen {
	s.CipherSuite = &v
	return s
}

func (s *ModifyDomainRequestListen) SetCustomCiphers(v []*string) *ModifyDomainRequestListen {
	s.CustomCiphers = v
	return s
}

func (s *ModifyDomainRequestListen) SetEnableTLSv3(v bool) *ModifyDomainRequestListen {
	s.EnableTLSv3 = &v
	return s
}

func (s *ModifyDomainRequestListen) SetExclusiveIp(v bool) *ModifyDomainRequestListen {
	s.ExclusiveIp = &v
	return s
}

func (s *ModifyDomainRequestListen) SetFocusHttps(v bool) *ModifyDomainRequestListen {
	s.FocusHttps = &v
	return s
}

func (s *ModifyDomainRequestListen) SetHstsIncludeSubDomain(v bool) *ModifyDomainRequestListen {
	s.HstsIncludeSubDomain = &v
	return s
}

func (s *ModifyDomainRequestListen) SetHstsMaxAge(v int64) *ModifyDomainRequestListen {
	s.HstsMaxAge = &v
	return s
}

func (s *ModifyDomainRequestListen) SetHstsPreload(v bool) *ModifyDomainRequestListen {
	s.HstsPreload = &v
	return s
}

func (s *ModifyDomainRequestListen) SetHttp2Enabled(v bool) *ModifyDomainRequestListen {
	s.Http2Enabled = &v
	return s
}

func (s *ModifyDomainRequestListen) SetHttpPorts(v []*int32) *ModifyDomainRequestListen {
	s.HttpPorts = v
	return s
}

func (s *ModifyDomainRequestListen) SetHttpsPorts(v []*int32) *ModifyDomainRequestListen {
	s.HttpsPorts = v
	return s
}

func (s *ModifyDomainRequestListen) SetIPv6Enabled(v bool) *ModifyDomainRequestListen {
	s.IPv6Enabled = &v
	return s
}

func (s *ModifyDomainRequestListen) SetProtectionResource(v string) *ModifyDomainRequestListen {
	s.ProtectionResource = &v
	return s
}

func (s *ModifyDomainRequestListen) SetSM2AccessOnly(v bool) *ModifyDomainRequestListen {
	s.SM2AccessOnly = &v
	return s
}

func (s *ModifyDomainRequestListen) SetSM2CertId(v string) *ModifyDomainRequestListen {
	s.SM2CertId = &v
	return s
}

func (s *ModifyDomainRequestListen) SetSM2Enabled(v bool) *ModifyDomainRequestListen {
	s.SM2Enabled = &v
	return s
}

func (s *ModifyDomainRequestListen) SetTLSVersion(v string) *ModifyDomainRequestListen {
	s.TLSVersion = &v
	return s
}

func (s *ModifyDomainRequestListen) SetXffHeaderMode(v int32) *ModifyDomainRequestListen {
	s.XffHeaderMode = &v
	return s
}

func (s *ModifyDomainRequestListen) SetXffHeaders(v []*string) *ModifyDomainRequestListen {
	s.XffHeaders = v
	return s
}

func (s *ModifyDomainRequestListen) Validate() error {
	return dara.Validate(s)
}

type ModifyDomainRequestRedirect struct {
	// The custom port configuration.
	BackendPorts []*ModifyDomainRequestRedirectBackendPorts `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
	// The IP addresses or domain names of the origin servers that correspond to the domain name. You can specify only IP addresses or domain names, not both. When the back-to-origin address is a domain name, only IPv4 is supported. IPv6 is not supported.
	Backends []*string `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	// The IP addresses or domain names of the secondary origin servers that correspond to the domain name.
	BackupBackends []*string `json:"BackupBackends,omitempty" xml:"BackupBackends,omitempty" type:"Repeated"`
	// Specifies whether to enable public cloud disaster recovery. Valid values:
	//
	// example:
	//
	// true
	CnameEnabled *bool `json:"CnameEnabled,omitempty" xml:"CnameEnabled,omitempty"`
	// The connection timeout period. Unit: seconds.
	//
	// example:
	//
	// 120
	ConnectTimeout *int32 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// Specifies whether to enable forced HTTP back-to-origin. This parameter is available only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
	//
	// example:
	//
	// true
	FocusHttpBackend *bool `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	// Specifies whether to enable HTTP/2 back-to-origin. Valid values:
	//
	// example:
	//
	// true
	Http2Origin *bool `json:"Http2Origin,omitempty" xml:"Http2Origin,omitempty"`
	// The maximum number of concurrent HTTP/2 back-to-origin connections. Valid values: 1 to 512. Default value: 2.
	//
	// example:
	//
	// 128
	Http2OriginMaxConcurrency *int32 `json:"Http2OriginMaxConcurrency,omitempty" xml:"Http2OriginMaxConcurrency,omitempty"`
	// Specifies whether to enable persistent connections. Valid values:
	//
	// example:
	//
	// true
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// The number of requests that reuse a persistent connection. Valid values: 60 to 1000. Default value: 1000.
	//
	// example:
	//
	// 1000
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The idle persistent connection timeout period. Valid values: 1 to 60. Default value: 15. Unit: seconds.
	//
	// example:
	//
	// 15
	KeepaliveTimeout *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The load balancing algorithm used for back-to-origin. Valid values:
	//
	// This parameter is required.
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
	// Indicates whether the client source IP preservation feature is enabled.
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
	// The traffic tag fields and values of the domain name, used to tag traffic processed by WAF.
	RequestHeaders []*ModifyDomainRequestRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// Specifies whether to retry when WAF fails to forward requests to the origin server. Valid values:
	//
	// example:
	//
	// true
	Retry *bool `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// The hybrid cloud forwarding rules. The value is a string converted from a JSON array. Each element in the JSON array is a struct that contains the following fields:
	//
	// example:
	//
	// [
	//
	//     {
	//
	//         "rs":
	//
	//         [
	//
	//             "1.1.XX.XX"
	//
	//         ],
	//
	//         "backupRs":
	//
	//         [
	//
	//             "2.2.XX.XX"
	//
	//         ],
	//
	//         "locationId": 535,
	//
	//         "location": "test1111"
	//
	//     }
	//
	// ]
	RoutingRules *string `json:"RoutingRules,omitempty" xml:"RoutingRules,omitempty"`
	// Specifies whether to enable back-to-origin SNI. This parameter is available only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
	//
	// example:
	//
	// true
	SniEnabled *bool `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	// The value of the custom SNI extension field. If this parameter is not specified, the value of the **Host*	- field in the request header is used as the SNI extension field value by default.
	//
	// example:
	//
	// www.aliyundoc.com
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// Specifies whether to allow WAF to overwrite WL-Proxy-Client-IP. Valid values:
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	WLProxyClientIp *bool `json:"WLProxyClientIp,omitempty" xml:"WLProxyClientIp,omitempty"`
	// Specifies whether to allow WAF to overwrite Web-Server-Type. Valid values:
	//
	// if can be null:
	// true
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
	// Specifies whether to allow WAF to overwrite X-Client-IP. Valid values:
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	XClientIp *bool `json:"XClientIp,omitempty" xml:"XClientIp,omitempty"`
	// Specifies whether to allow WAF to overwrite X-True-IP. Valid values:
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	XTrueIp *bool `json:"XTrueIp,omitempty" xml:"XTrueIp,omitempty"`
	// Specifies whether X-Forward-For-Proto passes the WAF protocol. Valid values:
	//
	// example:
	//
	// true
	XffProto *bool `json:"XffProto,omitempty" xml:"XffProto,omitempty"`
}

func (s ModifyDomainRequestRedirect) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainRequestRedirect) GoString() string {
	return s.String()
}

func (s *ModifyDomainRequestRedirect) GetBackendPorts() []*ModifyDomainRequestRedirectBackendPorts {
	return s.BackendPorts
}

func (s *ModifyDomainRequestRedirect) GetBackends() []*string {
	return s.Backends
}

func (s *ModifyDomainRequestRedirect) GetBackupBackends() []*string {
	return s.BackupBackends
}

func (s *ModifyDomainRequestRedirect) GetCnameEnabled() *bool {
	return s.CnameEnabled
}

func (s *ModifyDomainRequestRedirect) GetConnectTimeout() *int32 {
	return s.ConnectTimeout
}

func (s *ModifyDomainRequestRedirect) GetFocusHttpBackend() *bool {
	return s.FocusHttpBackend
}

func (s *ModifyDomainRequestRedirect) GetHttp2Origin() *bool {
	return s.Http2Origin
}

func (s *ModifyDomainRequestRedirect) GetHttp2OriginMaxConcurrency() *int32 {
	return s.Http2OriginMaxConcurrency
}

func (s *ModifyDomainRequestRedirect) GetKeepalive() *bool {
	return s.Keepalive
}

func (s *ModifyDomainRequestRedirect) GetKeepaliveRequests() *int32 {
	return s.KeepaliveRequests
}

func (s *ModifyDomainRequestRedirect) GetKeepaliveTimeout() *int32 {
	return s.KeepaliveTimeout
}

func (s *ModifyDomainRequestRedirect) GetLoadbalance() *string {
	return s.Loadbalance
}

func (s *ModifyDomainRequestRedirect) GetMaxBodySize() *int32 {
	return s.MaxBodySize
}

func (s *ModifyDomainRequestRedirect) GetProxyProtocol() *bool {
	return s.ProxyProtocol
}

func (s *ModifyDomainRequestRedirect) GetReadTimeout() *int32 {
	return s.ReadTimeout
}

func (s *ModifyDomainRequestRedirect) GetRequestHeaders() []*ModifyDomainRequestRedirectRequestHeaders {
	return s.RequestHeaders
}

func (s *ModifyDomainRequestRedirect) GetRetry() *bool {
	return s.Retry
}

func (s *ModifyDomainRequestRedirect) GetRoutingRules() *string {
	return s.RoutingRules
}

func (s *ModifyDomainRequestRedirect) GetSniEnabled() *bool {
	return s.SniEnabled
}

func (s *ModifyDomainRequestRedirect) GetSniHost() *string {
	return s.SniHost
}

func (s *ModifyDomainRequestRedirect) GetWLProxyClientIp() *bool {
	return s.WLProxyClientIp
}

func (s *ModifyDomainRequestRedirect) GetWebServerType() *bool {
	return s.WebServerType
}

func (s *ModifyDomainRequestRedirect) GetWriteTimeout() *int32 {
	return s.WriteTimeout
}

func (s *ModifyDomainRequestRedirect) GetXClientIp() *bool {
	return s.XClientIp
}

func (s *ModifyDomainRequestRedirect) GetXTrueIp() *bool {
	return s.XTrueIp
}

func (s *ModifyDomainRequestRedirect) GetXffProto() *bool {
	return s.XffProto
}

func (s *ModifyDomainRequestRedirect) SetBackendPorts(v []*ModifyDomainRequestRedirectBackendPorts) *ModifyDomainRequestRedirect {
	s.BackendPorts = v
	return s
}

func (s *ModifyDomainRequestRedirect) SetBackends(v []*string) *ModifyDomainRequestRedirect {
	s.Backends = v
	return s
}

func (s *ModifyDomainRequestRedirect) SetBackupBackends(v []*string) *ModifyDomainRequestRedirect {
	s.BackupBackends = v
	return s
}

func (s *ModifyDomainRequestRedirect) SetCnameEnabled(v bool) *ModifyDomainRequestRedirect {
	s.CnameEnabled = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetConnectTimeout(v int32) *ModifyDomainRequestRedirect {
	s.ConnectTimeout = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetFocusHttpBackend(v bool) *ModifyDomainRequestRedirect {
	s.FocusHttpBackend = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetHttp2Origin(v bool) *ModifyDomainRequestRedirect {
	s.Http2Origin = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetHttp2OriginMaxConcurrency(v int32) *ModifyDomainRequestRedirect {
	s.Http2OriginMaxConcurrency = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetKeepalive(v bool) *ModifyDomainRequestRedirect {
	s.Keepalive = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetKeepaliveRequests(v int32) *ModifyDomainRequestRedirect {
	s.KeepaliveRequests = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetKeepaliveTimeout(v int32) *ModifyDomainRequestRedirect {
	s.KeepaliveTimeout = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetLoadbalance(v string) *ModifyDomainRequestRedirect {
	s.Loadbalance = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetMaxBodySize(v int32) *ModifyDomainRequestRedirect {
	s.MaxBodySize = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetProxyProtocol(v bool) *ModifyDomainRequestRedirect {
	s.ProxyProtocol = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetReadTimeout(v int32) *ModifyDomainRequestRedirect {
	s.ReadTimeout = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetRequestHeaders(v []*ModifyDomainRequestRedirectRequestHeaders) *ModifyDomainRequestRedirect {
	s.RequestHeaders = v
	return s
}

func (s *ModifyDomainRequestRedirect) SetRetry(v bool) *ModifyDomainRequestRedirect {
	s.Retry = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetRoutingRules(v string) *ModifyDomainRequestRedirect {
	s.RoutingRules = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetSniEnabled(v bool) *ModifyDomainRequestRedirect {
	s.SniEnabled = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetSniHost(v string) *ModifyDomainRequestRedirect {
	s.SniHost = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetWLProxyClientIp(v bool) *ModifyDomainRequestRedirect {
	s.WLProxyClientIp = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetWebServerType(v bool) *ModifyDomainRequestRedirect {
	s.WebServerType = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetWriteTimeout(v int32) *ModifyDomainRequestRedirect {
	s.WriteTimeout = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetXClientIp(v bool) *ModifyDomainRequestRedirect {
	s.XClientIp = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetXTrueIp(v bool) *ModifyDomainRequestRedirect {
	s.XTrueIp = &v
	return s
}

func (s *ModifyDomainRequestRedirect) SetXffProto(v bool) *ModifyDomainRequestRedirect {
	s.XffProto = &v
	return s
}

func (s *ModifyDomainRequestRedirect) Validate() error {
	if s.BackendPorts != nil {
		for _, item := range s.BackendPorts {
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

type ModifyDomainRequestRedirectBackendPorts struct {
	// The back-to-origin port.
	//
	// example:
	//
	// 80
	BackendPort *int32 `json:"BackendPort,omitempty" xml:"BackendPort,omitempty"`
	// The listening port.
	//
	// example:
	//
	// 80
	ListenPort *int32 `json:"ListenPort,omitempty" xml:"ListenPort,omitempty"`
	// The protocol of the listening port. Valid values:
	//
	// example:
	//
	// http
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
}

func (s ModifyDomainRequestRedirectBackendPorts) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainRequestRedirectBackendPorts) GoString() string {
	return s.String()
}

func (s *ModifyDomainRequestRedirectBackendPorts) GetBackendPort() *int32 {
	return s.BackendPort
}

func (s *ModifyDomainRequestRedirectBackendPorts) GetListenPort() *int32 {
	return s.ListenPort
}

func (s *ModifyDomainRequestRedirectBackendPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *ModifyDomainRequestRedirectBackendPorts) SetBackendPort(v int32) *ModifyDomainRequestRedirectBackendPorts {
	s.BackendPort = &v
	return s
}

func (s *ModifyDomainRequestRedirectBackendPorts) SetListenPort(v int32) *ModifyDomainRequestRedirectBackendPorts {
	s.ListenPort = &v
	return s
}

func (s *ModifyDomainRequestRedirectBackendPorts) SetProtocol(v string) *ModifyDomainRequestRedirectBackendPorts {
	s.Protocol = &v
	return s
}

func (s *ModifyDomainRequestRedirectBackendPorts) Validate() error {
	return dara.Validate(s)
}

type ModifyDomainRequestRedirectRequestHeaders struct {
	// The custom request header field.
	//
	// example:
	//
	// aaa
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value set for the custom request header field.
	//
	// example:
	//
	// bbb
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyDomainRequestRedirectRequestHeaders) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainRequestRedirectRequestHeaders) GoString() string {
	return s.String()
}

func (s *ModifyDomainRequestRedirectRequestHeaders) GetKey() *string {
	return s.Key
}

func (s *ModifyDomainRequestRedirectRequestHeaders) GetValue() *string {
	return s.Value
}

func (s *ModifyDomainRequestRedirectRequestHeaders) SetKey(v string) *ModifyDomainRequestRedirectRequestHeaders {
	s.Key = &v
	return s
}

func (s *ModifyDomainRequestRedirectRequestHeaders) SetValue(v string) *ModifyDomainRequestRedirectRequestHeaders {
	s.Value = &v
	return s
}

func (s *ModifyDomainRequestRedirectRequestHeaders) Validate() error {
	return dara.Validate(s)
}
