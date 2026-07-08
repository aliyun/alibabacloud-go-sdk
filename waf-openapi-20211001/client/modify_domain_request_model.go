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
	// The access mode of the WAF instance. Valid values:
	//
	// - **share*	- (default): onboarding by using a CNAME record.
	//
	// - **hybrid_cloud_cname**: onboarding by using a hybrid cloud CNAME record.
	//
	// example:
	//
	// share
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The domain name whose configurations you want to modify.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the domain name.
	//
	// example:
	//
	// www.aliyundoc.com-waf
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The listening settings.
	//
	// This parameter is required.
	Listen *ModifyDomainRequestListen `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	// The forwarding settings.
	//
	// This parameter is required.
	Redirect *ModifyDomainRequestRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
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
	// The ID of the certificate.
	//
	// example:
	//
	// 123
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The type of the cipher suite. This parameter is available only when **HttpsPorts*	- is not empty, which indicates that the domain uses HTTPS. Valid values:
	//
	// - **1**: all cipher suites.
	//
	// - **2**: strong cipher suites. You can select this value only when you set **TLSVersion*	- to **tlsv1.2**.
	//
	// - **99**: custom cipher suites.
	//
	// example:
	//
	// 2
	CipherSuite *int32 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suites. This parameter is available only when you set **CipherSuite*	- to **99**.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Indicates whether TLS 1.3 is supported. This parameter is available only when **HttpsPorts*	- is not empty, which indicates that the domain uses HTTPS. Valid values:
	//
	// - **true**: TLS 1.3 is supported.
	//
	// - **false**: TLS 1.3 is not supported.
	//
	// example:
	//
	// true
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// Indicates whether an exclusive IP address is enabled. This parameter is available only when you set **IPv6Enabled*	- to false and **ProtectionResource*	- to **share**. Valid values:
	//
	// - **true**: An exclusive IP address is enabled.
	//
	// - **false*	- (default): An exclusive IP address is disabled.
	//
	// example:
	//
	// true
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Indicates whether forced HTTPS redirection is enabled. This parameter is available only when the domain uses HTTPS but not HTTP. Valid values:
	//
	// - **true**: Forced HTTPS redirection is enabled.
	//
	// - **false**: Forced HTTPS redirection is disabled.
	//
	// example:
	//
	// true
	FocusHttps *bool `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	// Indicates whether subdomains are included in the HTTP Strict Transport Security (HSTS) policy. Valid values:
	//
	// - **true**: Subdomains are included.
	//
	// - **false**: Subdomains are not included.
	//
	// example:
	//
	// false
	HstsIncludeSubDomain *bool `json:"HstsIncludeSubDomain,omitempty" xml:"HstsIncludeSubDomain,omitempty"`
	// The time-to-live (TTL) of the HSTS policy. Unit: seconds.
	//
	// example:
	//
	// 365000
	HstsMaxAge *int64 `json:"HstsMaxAge,omitempty" xml:"HstsMaxAge,omitempty"`
	// Indicates whether HSTS preload is enabled. Default value: false. Valid values:
	//
	// - true: HSTS preload is enabled.
	//
	// - false: HSTS preload is disabled.
	//
	// example:
	//
	// false
	HstsPreload *bool `json:"HstsPreload,omitempty" xml:"HstsPreload,omitempty"`
	// Indicates whether HTTP/2 is enabled. This parameter is available only when **HttpsPorts*	- is not empty, which indicates that the domain uses HTTPS. Valid values:
	//
	// - **true**: HTTP/2 is enabled.
	//
	// - **false*	- (default): HTTP/2 is disabled.
	//
	// example:
	//
	// true
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The HTTP listening ports. The format is [**port1,port2,...**].
	HttpPorts []*int32 `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	// The HTTPS listening ports. The format is [**port1,port2,...**].
	HttpsPorts []*int32 `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	// Indicates whether IPv6 is enabled. Valid values:
	//
	// - **true**: IPv6 is enabled.
	//
	// - **false*	- (default): IPv6 is disabled.
	//
	// example:
	//
	// true
	IPv6Enabled *bool `json:"IPv6Enabled,omitempty" xml:"IPv6Enabled,omitempty"`
	// The type of the protection resource. Valid values:
	//
	// - **share*	- (default): a shared cluster.
	//
	// - **gslb**: a shared cluster with global server load balancing.
	//
	// example:
	//
	// share
	ProtectionResource *string `json:"ProtectionResource,omitempty" xml:"ProtectionResource,omitempty"`
	// Indicates whether access is allowed only from SM clients. This parameter is available only when SM2Enabled is set to true. Valid values:
	//
	// - true: Only SM clients can access the website.
	//
	// - false: Both SM and non-SM clients can access the website.
	//
	// example:
	//
	// true
	SM2AccessOnly *bool `json:"SM2AccessOnly,omitempty" xml:"SM2AccessOnly,omitempty"`
	// The ID of the SM certificate. This parameter is required only when SM2Enabled is set to true.
	//
	// example:
	//
	// 123-cn-hangzhou
	SM2CertId *string `json:"SM2CertId,omitempty" xml:"SM2CertId,omitempty"`
	// Indicates whether SM certificates are enabled.
	//
	// example:
	//
	// true
	SM2Enabled *bool `json:"SM2Enabled,omitempty" xml:"SM2Enabled,omitempty"`
	// The TLS version. This parameter is available only when **HttpsPorts*	- is not empty, which indicates that the domain uses HTTPS. Valid values:
	//
	// - **tlsv1**
	//
	// - **tlsv1.1**
	//
	// - **tlsv1.2**
	//
	// example:
	//
	// tlsv1
	TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
	// The method that WAF uses to obtain the real IP address of a client. Valid values:
	//
	// - **0*	- (default): WAF obtains the client IP address from the TCP connection. This option is suitable if no Layer 7 proxies are deployed in front of WAF.
	//
	// - **1**: WAF obtains the client IP address from the first value of the X-Forwarded-For (XFF) header.
	//
	// - **2**: WAF obtains the client IP address from a custom header field.
	//
	// example:
	//
	// 2
	XffHeaderMode *int32 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// The custom header fields that are used to obtain the client IP address. The format is [**"header1","header2",...**].
	//
	// > This parameter is required only when you set **XffHeaderMode*	- to 2.
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
	// The custom port mappings for back-to-origin.
	BackendPorts []*ModifyDomainRequestRedirectBackendPorts `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
	// The IP addresses or domain names of the origin server. You can specify only one type of address. If you specify a domain name, only IPv4 is supported. IPv6 is not supported.
	//
	// - IP addresses: The format is [**"ip1","ip2",...**]. You can specify up to 20 IP addresses.
	//
	// - Domain names: The format is [**"domain"**]. You can specify up to 20 domain names.
	Backends []*string `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	// The IP addresses or domain names of the backup origin server.
	BackupBackends []*string `json:"BackupBackends,omitempty" xml:"BackupBackends,omitempty" type:"Repeated"`
	// Indicates whether public cloud disaster recovery is enabled. Valid values:
	//
	// - **true**: Public cloud disaster recovery is enabled.
	//
	// - **false*	- (default): Public cloud disaster recovery is disabled.
	//
	// example:
	//
	// true
	CnameEnabled *bool `json:"CnameEnabled,omitempty" xml:"CnameEnabled,omitempty"`
	// The connection timeout period. Unit: seconds. Valid values: 1 to 3600. Default value: 5.
	//
	// example:
	//
	// 120
	ConnectTimeout *int32 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// Indicates whether forced HTTP back-to-origin is enabled. This parameter is available only when you specify **HttpsPorts**. Valid values:
	//
	// - **true**: Forced HTTP back-to-origin is enabled.
	//
	// - **false**: Forced HTTP back-to-origin is disabled.
	//
	// example:
	//
	// true
	FocusHttpBackend *bool `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	// Indicates whether HTTP/2 is enabled for back-to-origin requests. Valid values:
	//
	// - **true**: HTTP/2 is enabled for back-to-origin requests.
	//
	// - **false**: HTTP/2 is disabled for back-to-origin requests.
	//
	// example:
	//
	// true
	Http2Origin *bool `json:"Http2Origin,omitempty" xml:"Http2Origin,omitempty"`
	// The maximum number of concurrent HTTP/2 back-to-origin requests. Valid values: 1 to 512. Default value: 2.
	//
	// example:
	//
	// 128
	Http2OriginMaxConcurrency *int32 `json:"Http2OriginMaxConcurrency,omitempty" xml:"Http2OriginMaxConcurrency,omitempty"`
	// Indicates whether persistent connections are enabled. Valid values:
	//
	// - **true*	- (default): Persistent connections are enabled.
	//
	// - **false**: Persistent connections are disabled.
	//
	// example:
	//
	// true
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// The number of requests that can be reused in a persistent connection. Valid values: 60 to 1000. Default value: 1000.
	//
	// > This parameter is available only when you enable persistent connections.
	//
	// example:
	//
	// 1000
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The idle timeout for a persistent connection. Unit: seconds. Valid values: 1 to 60. Default value: 15.
	//
	// > This parameter specifies the amount of time that an idle persistent connection can remain open.
	//
	// example:
	//
	// 15
	KeepaliveTimeout *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The load balancing algorithm for back-to-origin requests. Valid values:
	//
	// - **iphash**: the IP hash algorithm.
	//
	// - **roundRobin**: the round-robin algorithm.
	//
	// - **leastTime**: the least time algorithm. This value is available only when you set **ProtectionResource*	- to **gslb**.
	//
	// This parameter is required.
	//
	// example:
	//
	// iphash
	Loadbalance *string `json:"Loadbalance,omitempty" xml:"Loadbalance,omitempty"`
	// The maximum size of a request body. Unit: GB. Valid values: 2 to 10. Default value: 2.
	//
	// > This parameter is supported only by the WAF Ultimate edition.
	//
	// example:
	//
	// 2
	MaxBodySize *int32 `json:"MaxBodySize,omitempty" xml:"MaxBodySize,omitempty"`
	// Indicates whether the Proxy Protocol is enabled to preserve client IP addresses.
	//
	// - **true**: The Proxy Protocol is enabled. If you select this option, you can view the client IP address on the origin server.
	//
	// - **false**: The Proxy Protocol is disabled.
	//
	// example:
	//
	// false
	ProxyProtocol *bool `json:"ProxyProtocol,omitempty" xml:"ProxyProtocol,omitempty"`
	// The read timeout period. Unit: seconds. Valid values: 1 to 3600. Default value: 120.
	//
	// example:
	//
	// 200
	ReadTimeout *int32 `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	// The custom header fields and their values for traffic marking.
	//
	// WAF adds these fields and values to the request headers when traffic passes through WAF. This helps backend services identify and collect statistics on WAF-processed traffic.
	RequestHeaders []*ModifyDomainRequestRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// Indicates whether forwarding requests to the origin server are retried when the requests fail. Valid values:
	//
	// - **true*	- (default): Requests are retried.
	//
	// - **false**: Requests are not retried.
	//
	// example:
	//
	// true
	Retry *bool `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// The forwarding rules for a hybrid cloud deployment. This parameter is a string that contains a JSON array. Each element in the array is a struct that contains the following fields:
	//
	// - **rs**: an array of strings. The back-to-origin IP addresses or CNAMEs.
	//
	// - **backupRs**: an array of strings. The backup back-to-origin IP addresses or CNAMEs. This field is required. If you do not want to specify backup addresses, set it to [].
	//
	// - **location**: a string. The name of the protection node.
	//
	// - **locationId**: a long integer. The ID of the protection node.
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
	// Indicates whether back-to-origin Server Name Indication (SNI) is enabled. This parameter is available only when you specify **HttpsPorts**. Valid values:
	//
	// - **true**: Back-to-origin SNI is enabled.
	//
	// - **false*	- (default): Back-to-origin SNI is disabled.
	//
	// example:
	//
	// true
	SniEnabled *bool `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	// The value of the custom SNI field. If you do not set this parameter, the value of the **Host*	- field from the request header is used by default. You typically do not need to set this parameter unless your business requires a custom SNI value for back-to-origin requests.
	//
	// > This parameter is available only when you set **SniEnabled*	- to true.
	//
	// example:
	//
	// www.aliyundoc.com
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// Indicates whether WAF is allowed to overwrite the WL-Proxy-Client-IP header. Valid values:
	//
	// - **true*	- (default): WAF is allowed to overwrite the header.
	//
	// - **false**: WAF is not allowed to overwrite the header.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	WLProxyClientIp *bool `json:"WLProxyClientIp,omitempty" xml:"WLProxyClientIp,omitempty"`
	// Indicates whether WAF is allowed to overwrite the Web-Server-Type header. Valid values:
	//
	// - **true*	- (default): WAF is allowed to overwrite the header.
	//
	// - **false**: WAF is not allowed to overwrite the header.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	WebServerType *bool `json:"WebServerType,omitempty" xml:"WebServerType,omitempty"`
	// The write timeout period. Unit: seconds. Valid values: 1 to 3600. Default value: 120.
	//
	// example:
	//
	// 200
	WriteTimeout *int32 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
	// Indicates whether WAF is allowed to overwrite the X-Client-IP header. Valid values:
	//
	// - **true*	- (default): WAF is allowed to overwrite the header.
	//
	// - **false**: WAF is not allowed to overwrite the header.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	XClientIp *bool `json:"XClientIp,omitempty" xml:"XClientIp,omitempty"`
	// Indicates whether WAF is allowed to overwrite the X-True-IP header. Valid values:
	//
	// - **true*	- (default): WAF is allowed to overwrite the header.
	//
	// - **false**: WAF is not allowed to overwrite the header.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	XTrueIp *bool `json:"XTrueIp,omitempty" xml:"XTrueIp,omitempty"`
	// Indicates whether the X-Forwarded-Proto header is used to pass the protocol used by WAF to the origin server. Valid values:
	//
	// - **true*	- (default): The WAF protocol is passed.
	//
	// - **false**: The WAF protocol is not passed.
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
	// The port of the origin server.
	//
	// example:
	//
	// 80
	BackendPort *int32 `json:"BackendPort,omitempty" xml:"BackendPort,omitempty"`
	// The WAF listening port.
	//
	// example:
	//
	// 80
	ListenPort *int32 `json:"ListenPort,omitempty" xml:"ListenPort,omitempty"`
	// The protocol of the listening port. Valid values:
	//
	// - **http**: HTTP.
	//
	// - **https**: HTTPS.
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
	// The value of the custom header field.
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
