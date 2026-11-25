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
	// - **share**: CNAME access. This is the default value.
	//
	// - **hybrid_cloud_cname**: Hybrid Cloud CNAME access.
	//
	// example:
	//
	// share
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The domain name.
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
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The listener configurations.
	//
	// This parameter is required.
	Listen *ModifyDomainRequestListen `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	// The forwarding configurations.
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
	// The ID of the certificate to be added.
	//
	// example:
	//
	// 123
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The type of cipher suite to be added. This parameter takes effect only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
	//
	// - **1**: adds all cipher suites.
	//
	// - **2**: adds strong cipher suites. You can select this value only when **TLSVersion*	- is set to **tlsv1.2**.
	//
	// - **99**: adds custom cipher suites. You can select this value only when **TLSVersion*	- is not set to **tlsv1.3**.
	//
	// example:
	//
	// 2
	CipherSuite *int32 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suites to be added. This parameter takes effect only when **CipherSuite*	- is set to **99**.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Specifies whether to support TLS 1.3. Valid values:
	//
	// - **true**: supports TLS 1.3.
	//
	// - **false**: does not support TLS 1.3.
	//
	// > This parameter takes effect only when HttpsPorts is not empty, which indicates that the domain name uses HTTPS. If TLSVersion is set to tlsv1.3, this parameter must be set to true.
	//
	// example:
	//
	// true
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// Specifies whether to enable an exclusive IP address. This parameter takes effect only when **IPv6Enabled*	- is set to false, which indicates that IPv6 is disabled, and **ProtectionResource*	- is set to **share**, which indicates that a shared cluster is used. Valid values:
	//
	// - **true**: enables an exclusive IP address.
	//
	// - **false**: disables an exclusive IP address. This is the default value.
	//
	// example:
	//
	// true
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Specifies whether to enable force redirect to HTTPS. This parameter takes effect only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS, and **HttpPorts*	- is empty, which indicates that the domain name does not use HTTP. Valid values:
	//
	// - **true**: enables force redirect to HTTPS.
	//
	// - **false**: disables force redirect to HTTPS.
	//
	// example:
	//
	// true
	FocusHttps *bool `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	// Specifies whether HTTP Strict Transport Security (HSTS) includes subdomains. Valid values:
	//
	// - **true**: includes subdomains.
	//
	// - **false**: does not include subdomains.
	//
	// example:
	//
	// false
	HstsIncludeSubDomain *bool `json:"HstsIncludeSubDomain,omitempty" xml:"HstsIncludeSubDomain,omitempty"`
	// The expiration time of HSTS. Unit: seconds.
	//
	// example:
	//
	// 365000
	HstsMaxAge *int64 `json:"HstsMaxAge,omitempty" xml:"HstsMaxAge,omitempty"`
	// Specifies whether to enable HSTS preload. By default, this feature is disabled. Valid values:
	//
	// - true: enables HSTS preload.
	//
	// - false: disables HSTS preload.
	//
	// example:
	//
	// false
	HstsPreload *bool `json:"HstsPreload,omitempty" xml:"HstsPreload,omitempty"`
	// Specifies whether to enable HTTP/2. This parameter takes effect only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
	//
	// - **true**: enables HTTP/2.
	//
	// - **false**: disables HTTP/2. This is the default value.
	//
	// example:
	//
	// true
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The listening ports for HTTP. Specify the ports in the [**port1,port2,...**] format.
	HttpPorts []*int32 `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	// The listening ports for HTTPS. Specify the ports in the [**port1,port2,...**] format.
	HttpsPorts []*int32 `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	// Specifies whether to enable IPv6. Valid values:
	//
	// - **true**: enables IPv6.
	//
	// - **false**: disables IPv6. This is the default value.
	//
	// example:
	//
	// true
	IPv6Enabled *bool `json:"IPv6Enabled,omitempty" xml:"IPv6Enabled,omitempty"`
	// The type of protection resource to be used. Valid values:
	//
	// - **share**: uses a shared cluster. This is the default value.
	//
	// - **gslb**: uses intelligent load balancing of a shared cluster.
	//
	// example:
	//
	// share
	ProtectionResource *string `json:"ProtectionResource,omitempty" xml:"ProtectionResource,omitempty"`
	// Specifies whether to allow access only from clients that support SM2. This parameter takes effect only when SM2Enabled is set to true. Valid values:
	//
	// - true: allows access only from clients that support SM2.
	//
	// - false: allows access from clients that support SM2 and clients that do not support SM2.
	//
	// example:
	//
	// true
	SM2AccessOnly *bool `json:"SM2AccessOnly,omitempty" xml:"SM2AccessOnly,omitempty"`
	// The ID of the SM2 certificate to be added. This parameter takes effect only when SM2Enabled is set to true.
	//
	// example:
	//
	// 123-cn-hangzhou
	SM2CertId *string `json:"SM2CertId,omitempty" xml:"SM2CertId,omitempty"`
	// Specifies whether to enable the SM2 certificate.
	//
	// example:
	//
	// true
	SM2Enabled *bool `json:"SM2Enabled,omitempty" xml:"SM2Enabled,omitempty"`
	// The TLS version. This parameter takes effect only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
	//
	// - **tlsv1**: supports TLS 1.0 and later. This version provides the highest compatibility but lower security.
	//
	// - **tlsv1.1**: supports TLS 1.1 and later. This version provides good compatibility and security.
	//
	// - **tlsv1.2**: supports TLS 1.2 and later. This version provides good compatibility and the highest security.
	//
	// - **tlsv1.3**: supports only TLS 1.3. This version provides the highest security but lower compatibility.
	//
	// example:
	//
	// tlsv1
	TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
	// The method for WAF to obtain the actual IP address of the client.
	//
	// example:
	//
	// 2
	XffHeaderMode *int32 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// The list of custom fields used to obtain the client IP address. Specify the fields in the [**"header1","header2",...**] format.
	//
	// > You must specify this parameter only when **XffHeaderMode*	- is set to 2, which indicates that WAF reads the value of a custom field that you specify in the request header as the client IP address.
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
	// The custom back-to-origin configurations.
	BackendConfig []*ModifyDomainRequestRedirectBackendConfig `json:"BackendConfig,omitempty" xml:"BackendConfig,omitempty" type:"Repeated"`
	// The custom port configurations.
	BackendPorts []*ModifyDomainRequestRedirectBackendPorts `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
	// The IP addresses of the origin servers or the domain names to which requests are forwarded for the domain name. You can specify either IP addresses or domain names. If you specify domain names, only IPv4 addresses are supported. IPv6 addresses are not supported.
	//
	// - To specify IP addresses of origin servers, use the [**"ip1","ip2",...**] format. You can add up to 20 IP addresses.
	//
	// - To specify domain names, use the [**"domain"**] format. You can add up to 20 domain names.
	Backends []*string `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	// The IP addresses of the secondary origin servers or the secondary domain names to which requests are forwarded for the domain name.
	BackupBackends []*string `json:"BackupBackends,omitempty" xml:"BackupBackends,omitempty" type:"Repeated"`
	// Specifies whether to enable disaster recovery on Alibaba Cloud. Valid values:
	//
	// - **true**: enables disaster recovery on Alibaba Cloud.
	//
	// - **false**: disables disaster recovery on Alibaba Cloud. This is the default value.
	//
	// example:
	//
	// true
	CnameEnabled *bool `json:"CnameEnabled,omitempty" xml:"CnameEnabled,omitempty"`
	// The connection timeout period. Unit: seconds.
	//
	// Valid values: 1 to 3600.
	//
	// Default value: 5.
	//
	// example:
	//
	// 120
	ConnectTimeout *int32 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// Specifies whether to enable force redirect to HTTP for back-to-origin requests. This parameter takes effect only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
	//
	// - **true**: enables force redirect to HTTP for back-to-origin requests.
	//
	// - **false**: disables force redirect to HTTP for back-to-origin requests.
	//
	// example:
	//
	// true
	FocusHttpBackend *bool `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	// Specifies whether to enable HTTP/2 for back-to-origin requests. Valid values:
	//
	// - **true**: enables HTTP/2 for back-to-origin requests.
	//
	// - **false**: disables HTTP/2 for back-to-origin requests.
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
	// Specifies whether to maintain persistent connections. Valid values:
	//
	// - **true**: maintains persistent connections. This is the default value.
	//
	// - **false**: does not maintain persistent connections.
	//
	// example:
	//
	// true
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// The number of requests that can be reused over a persistent connection. Valid values: 60 to 1000. Default value: 1000.
	//
	// > This parameter specifies the number of requests that can be multiplexed over a persistent connection after persistent connections are enabled.
	//
	// example:
	//
	// 1000
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The timeout period for an idle persistent connection. Valid values: 1 to 60. Default value: 15. Unit: seconds.
	//
	// > This parameter specifies the duration after which an idle persistent connection is released.
	//
	// example:
	//
	// 15
	KeepaliveTimeout *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The load balancing algorithm used for back-to-origin requests. Valid values:
	//
	// - **iphash**: the IP hash algorithm.
	//
	// - **roundRobin**: the round-robin algorithm.
	//
	// - **leastTime**: the least time algorithm. You can select this value only when **ProtectionResource*	- is set to **gslb**, which indicates that intelligent load balancing of a shared cluster is used.
	//
	// This parameter is required.
	//
	// example:
	//
	// iphash
	Loadbalance *string `json:"Loadbalance,omitempty" xml:"Loadbalance,omitempty"`
	// The maximum size of the request body. Valid values: 2 to 10. Default value: 2. Unit: GB.
	//
	// > This parameter is supported only by WAF Ultimate Edition.
	//
	// example:
	//
	// 2
	MaxBodySize *int32 `json:"MaxBodySize,omitempty" xml:"MaxBodySize,omitempty"`
	// Specifies whether to enable the feature of preserving the originating IP address of the client. Valid values:
	//
	// - **true**: enables the feature. After this feature is enabled, backend services can view the originating IP address of the client.
	//
	// - **false**: disables the feature.
	//
	// example:
	//
	// false
	ProxyProtocol *bool `json:"ProxyProtocol,omitempty" xml:"ProxyProtocol,omitempty"`
	// The read timeout period. Unit: seconds.
	//
	// Valid values: 1 to 3600.
	//
	// Default value: 120.
	//
	// example:
	//
	// 200
	ReadTimeout *int32 `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	// The traffic marking fields and values for the domain name, which are used to mark the traffic processed by WAF.
	//
	// By specifying custom request header fields and their values, WAF automatically adds the specified custom field values to the request headers as traffic marks when the access traffic of the domain name passes through WAF. This facilitates the collection of related information by backend services.
	RequestHeaders []*ModifyDomainRequestRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// Specifies whether to retry when WAF fails to forward requests to the origin server. Valid values:
	//
	// - **true**: retries. This is the default value.
	//
	// - **false**: does not retry.
	//
	// example:
	//
	// true
	Retry *bool `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// The forwarding rules for Hybrid Cloud WAF. The value is a string transformed from a JSON array. Each element in the JSON array is a struct that contains the following fields:
	//
	// - **rs**: Array | The list of back-to-origin IP addresses or CNAME records.
	//
	// - **backupRs**: Array | The list of secondary back-to-origin IP addresses or CNAME records. This field is required. An empty array [] indicates that no secondary back-to-origin addresses are configured.
	//
	// - **location**: String | The name of the protection node.
	//
	// - **locationId**: Long | The ID of the protection node.
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
	// Specifies whether to enable Server Name Indication (SNI) for back-to-origin requests. This parameter takes effect only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
	//
	// - **true**: enables SNI for back-to-origin requests.
	//
	// - **false**: disables SNI for back-to-origin requests. This is the default value.
	//
	// example:
	//
	// true
	SniEnabled *bool `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	// The value of the custom SNI extension field. If you do not specify this parameter, the value of the **Host*	- field in the request header is used as the value of the SNI extension field by default.
	//
	// In most cases, you do not need to customize the SNI. You can customize the SNI only when your service has special configuration requirements and you want WAF to use an SNI that is different from the actual request Host in back-to-origin requests.
	//
	// > You must specify this parameter only when **SniEnabled*	- is set to true, which indicates that a custom SNI is used for back-to-origin requests.
	//
	// example:
	//
	// www.aliyundoc.com
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// Specifies whether to allow WAF to overwrite the WL-Proxy-Client-IP header. Valid values:
	//
	// - **true**: allows WAF to overwrite the header. This is the default value.
	//
	// - **false**: does not allow WAF to overwrite the header.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	WLProxyClientIp *bool `json:"WLProxyClientIp,omitempty" xml:"WLProxyClientIp,omitempty"`
	// Specifies whether to allow WAF to overwrite the Web-Server-Type header. Valid values:
	//
	// - **true**: allows WAF to overwrite the header. This is the default value.
	//
	// - **false**: does not allow WAF to overwrite the header.
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
	// Valid values: 1 to 3600.
	//
	// Default value: 120.
	//
	// example:
	//
	// 200
	WriteTimeout *int32 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
	// Specifies whether to allow WAF to overwrite the X-Client-IP header. Valid values:
	//
	// - **true**: allows WAF to overwrite the header. This is the default value.
	//
	// - **false**: does not allow WAF to overwrite the header.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	XClientIp *bool `json:"XClientIp,omitempty" xml:"XClientIp,omitempty"`
	// Specifies whether to allow WAF to overwrite the X-True-IP header. Valid values:
	//
	// - **true**: allows WAF to overwrite the header. This is the default value.
	//
	// - **false**: does not allow WAF to overwrite the header.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	XTrueIp *bool `json:"XTrueIp,omitempty" xml:"XTrueIp,omitempty"`
	// Specifies whether to pass the protocol used by WAF in the X-Forwarded-Proto header. Valid values:
	//
	// - **true**: passes the protocol used by WAF. This is the default value.
	//
	// - **false**: does not pass the protocol used by WAF.
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

func (s *ModifyDomainRequestRedirect) GetBackendConfig() []*ModifyDomainRequestRedirectBackendConfig {
	return s.BackendConfig
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

func (s *ModifyDomainRequestRedirect) SetBackendConfig(v []*ModifyDomainRequestRedirectBackendConfig) *ModifyDomainRequestRedirect {
	s.BackendConfig = v
	return s
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
	if s.BackendConfig != nil {
		for _, item := range s.BackendConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type ModifyDomainRequestRedirectBackendConfig struct {
	// The origin server.
	//
	// example:
	//
	// 1.1.XX.XX
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
	// The weight.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ModifyDomainRequestRedirectBackendConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainRequestRedirectBackendConfig) GoString() string {
	return s.String()
}

func (s *ModifyDomainRequestRedirectBackendConfig) GetBackend() *string {
	return s.Backend
}

func (s *ModifyDomainRequestRedirectBackendConfig) GetWeight() *int32 {
	return s.Weight
}

func (s *ModifyDomainRequestRedirectBackendConfig) SetBackend(v string) *ModifyDomainRequestRedirectBackendConfig {
	s.Backend = &v
	return s
}

func (s *ModifyDomainRequestRedirectBackendConfig) SetWeight(v int32) *ModifyDomainRequestRedirectBackendConfig {
	s.Weight = &v
	return s
}

func (s *ModifyDomainRequestRedirectBackendConfig) Validate() error {
	return dara.Validate(s)
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
	// The value of the custom request header field.
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
