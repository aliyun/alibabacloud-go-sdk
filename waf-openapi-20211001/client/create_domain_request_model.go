// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessType(v string) *CreateDomainRequest
	GetAccessType() *string
	SetDomain(v string) *CreateDomainRequest
	GetDomain() *string
	SetInstanceId(v string) *CreateDomainRequest
	GetInstanceId() *string
	SetListen(v *CreateDomainRequestListen) *CreateDomainRequest
	GetListen() *CreateDomainRequestListen
	SetRedirect(v *CreateDomainRequestRedirect) *CreateDomainRequest
	GetRedirect() *CreateDomainRequestRedirect
	SetRegionId(v string) *CreateDomainRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreateDomainRequest
	GetResourceManagerResourceGroupId() *string
	SetTag(v []*CreateDomainRequestTag) *CreateDomainRequest
	GetTag() []*CreateDomainRequestTag
}

type CreateDomainRequest struct {
	// The access type of the WAF instance. Valid values:
	//
	// - **share*	- (default): onboarding by using a CNAME record.
	//
	// - **hybrid_cloud_cname**: onboarding by using a hybrid cloud CNAME record.
	//
	// example:
	//
	// share
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The domain name that you want to add.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
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
	// The listening configurations.
	//
	// This parameter is required.
	Listen *CreateDomainRequestListen `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	// The forwarding configurations.
	//
	// This parameter is required.
	Redirect *CreateDomainRequestRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: The Chinese mainland.
	//
	// - **ap-southeast-1**: Outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The tags. You can specify up to 20 tags.
	Tag []*CreateDomainRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) GetAccessType() *string {
	return s.AccessType
}

func (s *CreateDomainRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDomainRequest) GetListen() *CreateDomainRequestListen {
	return s.Listen
}

func (s *CreateDomainRequest) GetRedirect() *CreateDomainRequestRedirect {
	return s.Redirect
}

func (s *CreateDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDomainRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateDomainRequest) GetTag() []*CreateDomainRequestTag {
	return s.Tag
}

func (s *CreateDomainRequest) SetAccessType(v string) *CreateDomainRequest {
	s.AccessType = &v
	return s
}

func (s *CreateDomainRequest) SetDomain(v string) *CreateDomainRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainRequest) SetInstanceId(v string) *CreateDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDomainRequest) SetListen(v *CreateDomainRequestListen) *CreateDomainRequest {
	s.Listen = v
	return s
}

func (s *CreateDomainRequest) SetRedirect(v *CreateDomainRequestRedirect) *CreateDomainRequest {
	s.Redirect = v
	return s
}

func (s *CreateDomainRequest) SetRegionId(v string) *CreateDomainRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDomainRequest) SetResourceManagerResourceGroupId(v string) *CreateDomainRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateDomainRequest) SetTag(v []*CreateDomainRequestTag) *CreateDomainRequest {
	s.Tag = v
	return s
}

func (s *CreateDomainRequest) Validate() error {
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDomainRequestListen struct {
	// The ID of the certificate. This parameter is available only if you specify **HttpsPorts*	- (indicating that the domain name uses HTTPS).
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 123
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The type of the cipher suite. This parameter is available only if you specify **HttpsPorts*	- (indicating that the domain name uses HTTPS). Valid values:
	//
	// - **1**: All cipher suites.
	//
	// - **2**: Strong cipher suites. This value is available only if you set **TLSVersion*	- to **tlsv1.2**.
	//
	// - **99**: Custom cipher suites.
	//
	// example:
	//
	// 2
	CipherSuite *int32 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suites.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Indicates whether TLS 1.3 is supported. This parameter is available only if you specify **HttpsPorts*	- (indicating that the domain name uses HTTPS). Valid values:
	//
	// - **true**: TLS 1.3 is supported.
	//
	// - **false**: TLS 1.3 is not supported.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// Indicates whether the exclusive IP address feature is enabled. This parameter is available only if you set **IPv6Enabled*	- to **false*	- (indicating that IPv6 is disabled) and **ProtectionResource*	- to **share*	- (indicating that a shared cluster is used). Valid values:
	//
	// - **true**: The exclusive IP address feature is enabled.
	//
	// - **false*	- (default): The exclusive IP address feature is disabled.
	//
	// example:
	//
	// true
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Indicates whether force redirect from HTTP to HTTPS is enabled for received requests. This parameter is available only if you specify HttpsPorts and leave HttpPorts empty. Valid values:
	//
	// - **true**: Force redirect from HTTP to HTTPS is enabled.
	//
	// - **false**: Force redirect from HTTP to HTTPS is disabled.
	//
	// example:
	//
	// true
	FocusHttps *bool `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	// Indicates whether HSTS includes subdomains. Valid values:
	//
	// - **true**: HSTS includes subdomains.
	//
	// - **false**: HSTS does not include subdomains.
	HstsIncludeSubDomain *bool `json:"HstsIncludeSubDomain,omitempty" xml:"HstsIncludeSubDomain,omitempty"`
	// The time-to-live (TTL) for HSTS. Unit: seconds.
	//
	// example:
	//
	// 365000
	HstsMaxAge *int64 `json:"HstsMaxAge,omitempty" xml:"HstsMaxAge,omitempty"`
	// Indicates whether HSTS preloading is enabled. Valid values:
	//
	// - **true**: HSTS preloading is enabled.
	//
	// - **false**: HSTS preloading is disabled.
	//
	// example:
	//
	// false
	HstsPreload *bool `json:"HstsPreload,omitempty" xml:"HstsPreload,omitempty"`
	// Indicates whether HTTP/2 is enabled. This parameter is available only if you specify **HttpsPorts*	- (indicating that the domain name uses HTTPS). Valid values:
	//
	// - **true**: HTTP/2 is enabled.
	//
	// - **false*	- (default): HTTP/2 is disabled.
	//
	// example:
	//
	// true
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The listening ports for HTTP.
	HttpPorts []*int32 `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	// The listening ports for HTTPS.
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
	// The type of protection resource. Valid values:
	//
	// - **share*	- (default): Shared cluster.
	//
	// - **gslb**: Intelligent load balancing for a shared cluster.
	//
	// example:
	//
	// share
	ProtectionResource *string `json:"ProtectionResource,omitempty" xml:"ProtectionResource,omitempty"`
	// Indicates whether access is restricted to SM certificate-based clients only. This parameter is available only if you set SM2Enabled to true. Valid values:
	//
	// - **true**: Only SM certificate-based clients can access the domain.
	//
	// - **false**: Both SM certificate-based and non-SM certificate-based clients can access the domain.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	SM2AccessOnly *bool `json:"SM2AccessOnly,omitempty" xml:"SM2AccessOnly,omitempty"`
	// The ID of the SM certificate. This parameter is available only if you set SM2Enabled to true.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 123-cn-hangzhou
	SM2CertId *string `json:"SM2CertId,omitempty" xml:"SM2CertId,omitempty"`
	// Indicates whether SM certificate-based encryption is enabled.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	SM2Enabled *bool `json:"SM2Enabled,omitempty" xml:"SM2Enabled,omitempty"`
	// The minimum Transport Layer Security (TLS) version. This parameter is available only if you specify **HttpsPorts*	- (indicating that the domain name uses HTTPS). Valid values:
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
	// The method that WAF uses to obtain the originating IP address of a client. Valid values:
	//
	// - **0*	- (default): The client traffic does not pass through other Layer 7 proxies before it reaches WAF.
	//
	// - **1**: WAF uses the first value in the X-Forwarded-For (XFF) header as the client IP address.
	//
	// - **2**: WAF uses the value of a custom header field that you specify as the client IP address.
	//
	// example:
	//
	// 1
	XffHeaderMode *int32 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// The custom header fields that are used to obtain the originating IP address of a client.
	XffHeaders []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
}

func (s CreateDomainRequestListen) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainRequestListen) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestListen) GetCertId() *string {
	return s.CertId
}

func (s *CreateDomainRequestListen) GetCipherSuite() *int32 {
	return s.CipherSuite
}

func (s *CreateDomainRequestListen) GetCustomCiphers() []*string {
	return s.CustomCiphers
}

func (s *CreateDomainRequestListen) GetEnableTLSv3() *bool {
	return s.EnableTLSv3
}

func (s *CreateDomainRequestListen) GetExclusiveIp() *bool {
	return s.ExclusiveIp
}

func (s *CreateDomainRequestListen) GetFocusHttps() *bool {
	return s.FocusHttps
}

func (s *CreateDomainRequestListen) GetHstsIncludeSubDomain() *bool {
	return s.HstsIncludeSubDomain
}

func (s *CreateDomainRequestListen) GetHstsMaxAge() *int64 {
	return s.HstsMaxAge
}

func (s *CreateDomainRequestListen) GetHstsPreload() *bool {
	return s.HstsPreload
}

func (s *CreateDomainRequestListen) GetHttp2Enabled() *bool {
	return s.Http2Enabled
}

func (s *CreateDomainRequestListen) GetHttpPorts() []*int32 {
	return s.HttpPorts
}

func (s *CreateDomainRequestListen) GetHttpsPorts() []*int32 {
	return s.HttpsPorts
}

func (s *CreateDomainRequestListen) GetIPv6Enabled() *bool {
	return s.IPv6Enabled
}

func (s *CreateDomainRequestListen) GetProtectionResource() *string {
	return s.ProtectionResource
}

func (s *CreateDomainRequestListen) GetSM2AccessOnly() *bool {
	return s.SM2AccessOnly
}

func (s *CreateDomainRequestListen) GetSM2CertId() *string {
	return s.SM2CertId
}

func (s *CreateDomainRequestListen) GetSM2Enabled() *bool {
	return s.SM2Enabled
}

func (s *CreateDomainRequestListen) GetTLSVersion() *string {
	return s.TLSVersion
}

func (s *CreateDomainRequestListen) GetXffHeaderMode() *int32 {
	return s.XffHeaderMode
}

func (s *CreateDomainRequestListen) GetXffHeaders() []*string {
	return s.XffHeaders
}

func (s *CreateDomainRequestListen) SetCertId(v string) *CreateDomainRequestListen {
	s.CertId = &v
	return s
}

func (s *CreateDomainRequestListen) SetCipherSuite(v int32) *CreateDomainRequestListen {
	s.CipherSuite = &v
	return s
}

func (s *CreateDomainRequestListen) SetCustomCiphers(v []*string) *CreateDomainRequestListen {
	s.CustomCiphers = v
	return s
}

func (s *CreateDomainRequestListen) SetEnableTLSv3(v bool) *CreateDomainRequestListen {
	s.EnableTLSv3 = &v
	return s
}

func (s *CreateDomainRequestListen) SetExclusiveIp(v bool) *CreateDomainRequestListen {
	s.ExclusiveIp = &v
	return s
}

func (s *CreateDomainRequestListen) SetFocusHttps(v bool) *CreateDomainRequestListen {
	s.FocusHttps = &v
	return s
}

func (s *CreateDomainRequestListen) SetHstsIncludeSubDomain(v bool) *CreateDomainRequestListen {
	s.HstsIncludeSubDomain = &v
	return s
}

func (s *CreateDomainRequestListen) SetHstsMaxAge(v int64) *CreateDomainRequestListen {
	s.HstsMaxAge = &v
	return s
}

func (s *CreateDomainRequestListen) SetHstsPreload(v bool) *CreateDomainRequestListen {
	s.HstsPreload = &v
	return s
}

func (s *CreateDomainRequestListen) SetHttp2Enabled(v bool) *CreateDomainRequestListen {
	s.Http2Enabled = &v
	return s
}

func (s *CreateDomainRequestListen) SetHttpPorts(v []*int32) *CreateDomainRequestListen {
	s.HttpPorts = v
	return s
}

func (s *CreateDomainRequestListen) SetHttpsPorts(v []*int32) *CreateDomainRequestListen {
	s.HttpsPorts = v
	return s
}

func (s *CreateDomainRequestListen) SetIPv6Enabled(v bool) *CreateDomainRequestListen {
	s.IPv6Enabled = &v
	return s
}

func (s *CreateDomainRequestListen) SetProtectionResource(v string) *CreateDomainRequestListen {
	s.ProtectionResource = &v
	return s
}

func (s *CreateDomainRequestListen) SetSM2AccessOnly(v bool) *CreateDomainRequestListen {
	s.SM2AccessOnly = &v
	return s
}

func (s *CreateDomainRequestListen) SetSM2CertId(v string) *CreateDomainRequestListen {
	s.SM2CertId = &v
	return s
}

func (s *CreateDomainRequestListen) SetSM2Enabled(v bool) *CreateDomainRequestListen {
	s.SM2Enabled = &v
	return s
}

func (s *CreateDomainRequestListen) SetTLSVersion(v string) *CreateDomainRequestListen {
	s.TLSVersion = &v
	return s
}

func (s *CreateDomainRequestListen) SetXffHeaderMode(v int32) *CreateDomainRequestListen {
	s.XffHeaderMode = &v
	return s
}

func (s *CreateDomainRequestListen) SetXffHeaders(v []*string) *CreateDomainRequestListen {
	s.XffHeaders = v
	return s
}

func (s *CreateDomainRequestListen) Validate() error {
	return dara.Validate(s)
}

type CreateDomainRequestRedirect struct {
	// The custom port configuration.
	BackendPorts []*CreateDomainRequestRedirectBackendPorts `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
	// The IP address or domain name of the origin server.
	Backends []*string `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	// The secondary IP address or domain name of the origin server.
	BackupBackends []*string `json:"BackupBackends,omitempty" xml:"BackupBackends,omitempty" type:"Repeated"`
	// Indicates whether the public cloud disaster recovery feature is enabled for the domain name. Valid values:
	//
	// - **true**: The public cloud disaster recovery feature is enabled.
	//
	// - **false*	- (default): The public cloud disaster recovery feature is disabled.
	//
	// example:
	//
	// true
	CnameEnabled *bool `json:"CnameEnabled,omitempty" xml:"CnameEnabled,omitempty"`
	// The timeout period for connections. Unit: seconds. Valid values: 1 to 3600. Default value: 5.
	//
	// example:
	//
	// 120
	ConnectTimeout *int32 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// Indicates whether force redirect from HTTPS to HTTP is enabled for back-to-origin requests. This parameter is available only if you specify **HttpsPorts*	- (indicating that the domain name uses HTTPS). Valid values:
	//
	// - **true**: Force redirect from HTTPS to HTTP is enabled for back-to-origin requests.
	//
	// - **false**: Force redirect from HTTPS to HTTP is disabled for back-to-origin requests.
	//
	// example:
	//
	// true
	FocusHttpBackend *bool `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	// Indicates whether HTTP/2 is enabled for back-to-origin traffic. Valid values:
	//
	// - **true**: HTTP/2 is enabled for back-to-origin traffic.
	//
	// - **false**: HTTP/2 is disabled for back-to-origin traffic.
	//
	// example:
	//
	// true
	Http2Origin *bool `json:"Http2Origin,omitempty" xml:"Http2Origin,omitempty"`
	// The maximum number of concurrent HTTP/2 back-to-origin requests. Valid values: 1 to 512. Default value: 128.
	//
	// example:
	//
	// 128
	Http2OriginMaxConcurrency *int32 `json:"Http2OriginMaxConcurrency,omitempty" xml:"Http2OriginMaxConcurrency,omitempty"`
	// Indicates whether the persistent connection feature is enabled. Valid values:
	//
	// - **true*	- (default): The persistent connection feature is enabled.
	//
	// - **false**: The persistent connection feature is disabled.
	//
	// example:
	//
	// true
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// The number of reused persistent connections. Valid values: 60 to 1000. Default value: 1000.
	//
	// > The number of reused persistent connections after the persistent connection feature is enabled.
	//
	// example:
	//
	// 1000
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The timeout period of idle persistent connections. Valid values: 1 to 60. Default value: 15. Unit: seconds.
	//
	// > This parameter specifies the time for which a reused persistent connection can remain in the Idle state before the persistent connection is closed.
	//
	// example:
	//
	// 15
	KeepaliveTimeout *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The load balancing algorithm that you want to use when WAF forwards requests to the origin server. Valid values:
	//
	// - **iphash**: IP hash algorithm.
	//
	// - **roundRobin**: Round-robin algorithm.
	//
	// - **leastTime**: Least Time algorithm. This value is available only if you set **ProtectionResource*	- to **gslb*	- (indicating that intelligent load balancing for a shared cluster is used).
	//
	// This parameter is required.
	//
	// example:
	//
	// roundRobin
	Loadbalance *string `json:"Loadbalance,omitempty" xml:"Loadbalance,omitempty"`
	// The maximum size of a request body. Valid values: 2 to 10. Default value: 2. Unit: GB.
	//
	// > This parameter is supported only by the Ultimate edition.
	//
	// example:
	//
	// 2
	MaxBodySize *int32 `json:"MaxBodySize,omitempty" xml:"MaxBodySize,omitempty"`
	// Indicates whether the Proxy Protocol feature is enabled to preserve the client\\"s source IP address. Valid values:
	//
	// - **true**: The Proxy Protocol feature is enabled. After this feature is enabled, backend services can view the original IP address of the client.
	//
	// - **false**: The Proxy Protocol feature is disabled.
	//
	// example:
	//
	// false
	ProxyProtocol *bool `json:"ProxyProtocol,omitempty" xml:"ProxyProtocol,omitempty"`
	// The timeout period for read operations. Unit: seconds. Valid values: 1 to 3600. Default value: 120.
	//
	// example:
	//
	// 200
	ReadTimeout *int32 `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	// The key-value pairs that you want to use to label the requests that pass through the WAF instance.
	//
	// When a request passes through WAF, WAF automatically adds the custom header fields to the request to mark the request. This way, the backend service can identify requests that are processed by WAF.
	RequestHeaders []*CreateDomainRequestRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// Indicates whether WAF retries when WAF fails to forward requests to the origin server. Valid values:
	//
	// - **true*	- (default): WAF retries.
	//
	// - **false**: WAF does not retry.
	//
	// example:
	//
	// true
	Retry *bool `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// The forwarding rules for hybrid cloud mode. The value contains the following fields:
	//
	// - **rs**: The IP addresses or canonical names of the origin servers.
	//
	// - **backupRs**: The backup IP addresses or canonical names of the origin servers. Required. An empty array [] means no backup is configured.
	//
	// - **location**: The name of the protection node.
	//
	// - **locationId**: The ID of the protection node.
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
	// Indicates whether the Server Name Indication (SNI) feature is enabled for back-to-origin requests. This parameter is available only if you specify **HttpsPorts*	- (indicating that the domain name uses HTTPS). Valid values:
	//
	// - **true**: The SNI feature is enabled for back-to-origin requests.
	//
	// - **false*	- (default): The SNI feature is disabled for back-to-origin requests.
	//
	// example:
	//
	// true
	SniEnabled *bool `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	// The custom value of the SNI field. If you do not specify this parameter, the value of the **Host*	- header in the request is used as the value of the SNI field. This parameter is optional. If you want WAF to use an SNI field value that is different from the Host field value in back-to-origin requests, you can specify a custom value for the SNI field.
	//
	// > This parameter is required only if you set **SniEnabled*	- to **true*	- (indicating that the SNI feature is enabled for back-to-origin requests).
	//
	// example:
	//
	// www.aliyundoc.com
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// Indicates whether WAF is allowed to overwrite the WL-Proxy-Client-IP header. Valid values:
	//
	// - **true*	- (default): WAF overwrites the header.
	//
	// - **false**: WAF does not overwrite the header.
	//
	// example:
	//
	// true
	WLProxyClientIp *bool `json:"WLProxyClientIp,omitempty" xml:"WLProxyClientIp,omitempty"`
	// Indicates whether WAF is allowed to overwrite the Web-Server-Type header. Valid values:
	//
	// - **true*	- (default): WAF overwrites the header.
	//
	// - **false**: WAF does not overwrite the header.
	//
	// example:
	//
	// true
	WebServerType *bool `json:"WebServerType,omitempty" xml:"WebServerType,omitempty"`
	// The timeout period for write operations. Unit: seconds. Valid values: 1 to 3600. Default value: 120.
	//
	// example:
	//
	// 200
	WriteTimeout *int32 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
	// Indicates whether WAF is allowed to overwrite the X-Client-IP header. Valid values:
	//
	// - **true*	- (default): WAF overwrites the header.
	//
	// - **false**: WAF does not overwrite the header.
	//
	// example:
	//
	// true
	XClientIp *bool `json:"XClientIp,omitempty" xml:"XClientIp,omitempty"`
	// Indicates whether WAF is allowed to overwrite the X-True-IP header. Valid values:
	//
	// - **true*	- (default): WAF overwrites the header.
	//
	// - **false**: WAF does not overwrite the header.
	//
	// example:
	//
	// true
	XTrueIp *bool `json:"XTrueIp,omitempty" xml:"XTrueIp,omitempty"`
	// Indicates whether the X-Forward-For-Proto header is used to identify the protocol used by WAF to forward requests to the origin server. Valid values:
	//
	// - **true*	- (default): The X-Forward-For-Proto header is used to identify the protocol.
	//
	// - **false**: The X-Forward-For-Proto header is not used.
	//
	// example:
	//
	// false
	XffProto *bool `json:"XffProto,omitempty" xml:"XffProto,omitempty"`
}

func (s CreateDomainRequestRedirect) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainRequestRedirect) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestRedirect) GetBackendPorts() []*CreateDomainRequestRedirectBackendPorts {
	return s.BackendPorts
}

func (s *CreateDomainRequestRedirect) GetBackends() []*string {
	return s.Backends
}

func (s *CreateDomainRequestRedirect) GetBackupBackends() []*string {
	return s.BackupBackends
}

func (s *CreateDomainRequestRedirect) GetCnameEnabled() *bool {
	return s.CnameEnabled
}

func (s *CreateDomainRequestRedirect) GetConnectTimeout() *int32 {
	return s.ConnectTimeout
}

func (s *CreateDomainRequestRedirect) GetFocusHttpBackend() *bool {
	return s.FocusHttpBackend
}

func (s *CreateDomainRequestRedirect) GetHttp2Origin() *bool {
	return s.Http2Origin
}

func (s *CreateDomainRequestRedirect) GetHttp2OriginMaxConcurrency() *int32 {
	return s.Http2OriginMaxConcurrency
}

func (s *CreateDomainRequestRedirect) GetKeepalive() *bool {
	return s.Keepalive
}

func (s *CreateDomainRequestRedirect) GetKeepaliveRequests() *int32 {
	return s.KeepaliveRequests
}

func (s *CreateDomainRequestRedirect) GetKeepaliveTimeout() *int32 {
	return s.KeepaliveTimeout
}

func (s *CreateDomainRequestRedirect) GetLoadbalance() *string {
	return s.Loadbalance
}

func (s *CreateDomainRequestRedirect) GetMaxBodySize() *int32 {
	return s.MaxBodySize
}

func (s *CreateDomainRequestRedirect) GetProxyProtocol() *bool {
	return s.ProxyProtocol
}

func (s *CreateDomainRequestRedirect) GetReadTimeout() *int32 {
	return s.ReadTimeout
}

func (s *CreateDomainRequestRedirect) GetRequestHeaders() []*CreateDomainRequestRedirectRequestHeaders {
	return s.RequestHeaders
}

func (s *CreateDomainRequestRedirect) GetRetry() *bool {
	return s.Retry
}

func (s *CreateDomainRequestRedirect) GetRoutingRules() *string {
	return s.RoutingRules
}

func (s *CreateDomainRequestRedirect) GetSniEnabled() *bool {
	return s.SniEnabled
}

func (s *CreateDomainRequestRedirect) GetSniHost() *string {
	return s.SniHost
}

func (s *CreateDomainRequestRedirect) GetWLProxyClientIp() *bool {
	return s.WLProxyClientIp
}

func (s *CreateDomainRequestRedirect) GetWebServerType() *bool {
	return s.WebServerType
}

func (s *CreateDomainRequestRedirect) GetWriteTimeout() *int32 {
	return s.WriteTimeout
}

func (s *CreateDomainRequestRedirect) GetXClientIp() *bool {
	return s.XClientIp
}

func (s *CreateDomainRequestRedirect) GetXTrueIp() *bool {
	return s.XTrueIp
}

func (s *CreateDomainRequestRedirect) GetXffProto() *bool {
	return s.XffProto
}

func (s *CreateDomainRequestRedirect) SetBackendPorts(v []*CreateDomainRequestRedirectBackendPorts) *CreateDomainRequestRedirect {
	s.BackendPorts = v
	return s
}

func (s *CreateDomainRequestRedirect) SetBackends(v []*string) *CreateDomainRequestRedirect {
	s.Backends = v
	return s
}

func (s *CreateDomainRequestRedirect) SetBackupBackends(v []*string) *CreateDomainRequestRedirect {
	s.BackupBackends = v
	return s
}

func (s *CreateDomainRequestRedirect) SetCnameEnabled(v bool) *CreateDomainRequestRedirect {
	s.CnameEnabled = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetConnectTimeout(v int32) *CreateDomainRequestRedirect {
	s.ConnectTimeout = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetFocusHttpBackend(v bool) *CreateDomainRequestRedirect {
	s.FocusHttpBackend = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetHttp2Origin(v bool) *CreateDomainRequestRedirect {
	s.Http2Origin = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetHttp2OriginMaxConcurrency(v int32) *CreateDomainRequestRedirect {
	s.Http2OriginMaxConcurrency = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetKeepalive(v bool) *CreateDomainRequestRedirect {
	s.Keepalive = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetKeepaliveRequests(v int32) *CreateDomainRequestRedirect {
	s.KeepaliveRequests = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetKeepaliveTimeout(v int32) *CreateDomainRequestRedirect {
	s.KeepaliveTimeout = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetLoadbalance(v string) *CreateDomainRequestRedirect {
	s.Loadbalance = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetMaxBodySize(v int32) *CreateDomainRequestRedirect {
	s.MaxBodySize = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetProxyProtocol(v bool) *CreateDomainRequestRedirect {
	s.ProxyProtocol = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetReadTimeout(v int32) *CreateDomainRequestRedirect {
	s.ReadTimeout = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetRequestHeaders(v []*CreateDomainRequestRedirectRequestHeaders) *CreateDomainRequestRedirect {
	s.RequestHeaders = v
	return s
}

func (s *CreateDomainRequestRedirect) SetRetry(v bool) *CreateDomainRequestRedirect {
	s.Retry = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetRoutingRules(v string) *CreateDomainRequestRedirect {
	s.RoutingRules = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetSniEnabled(v bool) *CreateDomainRequestRedirect {
	s.SniEnabled = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetSniHost(v string) *CreateDomainRequestRedirect {
	s.SniHost = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetWLProxyClientIp(v bool) *CreateDomainRequestRedirect {
	s.WLProxyClientIp = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetWebServerType(v bool) *CreateDomainRequestRedirect {
	s.WebServerType = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetWriteTimeout(v int32) *CreateDomainRequestRedirect {
	s.WriteTimeout = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetXClientIp(v bool) *CreateDomainRequestRedirect {
	s.XClientIp = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetXTrueIp(v bool) *CreateDomainRequestRedirect {
	s.XTrueIp = &v
	return s
}

func (s *CreateDomainRequestRedirect) SetXffProto(v bool) *CreateDomainRequestRedirect {
	s.XffProto = &v
	return s
}

func (s *CreateDomainRequestRedirect) Validate() error {
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

type CreateDomainRequestRedirectBackendPorts struct {
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
	// The protocol of the listener port. Valid values:
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

func (s CreateDomainRequestRedirectBackendPorts) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainRequestRedirectBackendPorts) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestRedirectBackendPorts) GetBackendPort() *int32 {
	return s.BackendPort
}

func (s *CreateDomainRequestRedirectBackendPorts) GetListenPort() *int32 {
	return s.ListenPort
}

func (s *CreateDomainRequestRedirectBackendPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateDomainRequestRedirectBackendPorts) SetBackendPort(v int32) *CreateDomainRequestRedirectBackendPorts {
	s.BackendPort = &v
	return s
}

func (s *CreateDomainRequestRedirectBackendPorts) SetListenPort(v int32) *CreateDomainRequestRedirectBackendPorts {
	s.ListenPort = &v
	return s
}

func (s *CreateDomainRequestRedirectBackendPorts) SetProtocol(v string) *CreateDomainRequestRedirectBackendPorts {
	s.Protocol = &v
	return s
}

func (s *CreateDomainRequestRedirectBackendPorts) Validate() error {
	return dara.Validate(s)
}

type CreateDomainRequestRedirectRequestHeaders struct {
	// The key of the custom header field.
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

func (s CreateDomainRequestRedirectRequestHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainRequestRedirectRequestHeaders) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestRedirectRequestHeaders) GetKey() *string {
	return s.Key
}

func (s *CreateDomainRequestRedirectRequestHeaders) GetValue() *string {
	return s.Value
}

func (s *CreateDomainRequestRedirectRequestHeaders) SetKey(v string) *CreateDomainRequestRedirectRequestHeaders {
	s.Key = &v
	return s
}

func (s *CreateDomainRequestRedirectRequestHeaders) SetValue(v string) *CreateDomainRequestRedirectRequestHeaders {
	s.Value = &v
	return s
}

func (s *CreateDomainRequestRedirectRequestHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateDomainRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// Tagkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TagValue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDomainRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainRequestTag) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateDomainRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateDomainRequestTag) SetKey(v string) *CreateDomainRequestTag {
	s.Key = &v
	return s
}

func (s *CreateDomainRequestTag) SetValue(v string) *CreateDomainRequestTag {
	s.Value = &v
	return s
}

func (s *CreateDomainRequestTag) Validate() error {
	return dara.Validate(s)
}
