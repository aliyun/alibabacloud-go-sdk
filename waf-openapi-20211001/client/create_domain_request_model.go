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
	// - **share*	- (default): CNAME access.
	//
	// - **hybrid_cloud_cname**: hybrid cloud CNAME access.
	//
	// example:
	//
	// share
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The domain name to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the current WAF instance.
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
	Listen *CreateDomainRequestListen `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	// The forwarding configuration.
	//
	// This parameter is required.
	Redirect *CreateDomainRequestRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// The region where the WAF instance is deployed. Valid values:
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
	// The Alibaba Cloud resource group ID.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The tag list, which contains a maximum of 20 items.
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
	// The ID of the certificate to add. This parameter is used only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 123
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The type of cipher suite to add. This parameter is used only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
	//
	// - **1**: all cipher suites.
	//
	// - **2**: strong cipher suites. This value is available only when **TLSVersion*	- is set to **tlsv1.2**.
	//
	// - **99**: custom cipher suites.
	//
	// example:
	//
	// 2
	CipherSuite *int32 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suites to add.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Specifies whether to support TLS 1.3. This parameter is used only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
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
	// Specifies whether to enable an exclusive IP address. This parameter is used only when **IPv6Enabled*	- is set to **false*	- (which indicates that IPv6 is not enabled) and **ProtectionResource*	- is set to **share*	- (which indicates that a shared cluster is used). Valid values:
	//
	// - **true**: An exclusive IP address is enabled.
	//
	// - **false*	- (default): An exclusive IP address is not enabled.
	//
	// example:
	//
	// true
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Specifies whether to enable forced HTTPS redirect. This parameter is used only when HttpsPorts is not empty (which indicates that the domain name uses HTTPS) and HttpPorts is empty (which indicates that the domain name does not use HTTP). Valid values:
	//
	// - **true**: Forced HTTPS redirect is enabled.
	//
	// - **false**: Forced HTTPS redirect is not enabled.
	//
	// example:
	//
	// true
	FocusHttps *bool `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	// Specifies whether HSTS includes subdomains. Valid values:
	//
	// - **true**: Enabled.
	//
	// - **false**: Not enabled.
	HstsIncludeSubDomain *bool `json:"HstsIncludeSubDomain,omitempty" xml:"HstsIncludeSubDomain,omitempty"`
	// The HSTS expiration time. Unit: seconds.
	//
	// example:
	//
	// 365000
	HstsMaxAge *int64 `json:"HstsMaxAge,omitempty" xml:"HstsMaxAge,omitempty"`
	// Specifies whether to enable HSTS preloading. This feature is disabled by default. Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// example:
	//
	// false
	HstsPreload *bool `json:"HstsPreload,omitempty" xml:"HstsPreload,omitempty"`
	// Specifies whether to enable HTTP/2. This parameter is used only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
	//
	// - **true**: HTTP/2 is enabled.
	//
	// - **false*	- (default): HTTP/2 is not enabled.
	//
	// example:
	//
	// true
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The listening ports for HTTP.
	HttpPorts []*int32 `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	// The listening ports for HTTPS.
	HttpsPorts []*int32 `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	// Specifies whether to enable IPv6. Valid values:
	//
	// - **true**: IPv6 is enabled.
	//
	// - **false*	- (default): IPv6 is not enabled.
	//
	// example:
	//
	// true
	IPv6Enabled *bool `json:"IPv6Enabled,omitempty" xml:"IPv6Enabled,omitempty"`
	// The type of protection resource to use. Valid values:
	//
	// - **share*	- (default): shared cluster.
	//
	// - **gslb**: shared cluster-based intelligent load balancing.
	//
	// example:
	//
	// share
	ProtectionResource *string `json:"ProtectionResource,omitempty" xml:"ProtectionResource,omitempty"`
	// Specifies whether only China Encryption (SM) clients can access the domain name. This parameter is used only when SM2Enabled is set to true.
	//
	// - true: Only China Encryption (SM) clients can access the domain name.
	//
	// - false: Both China Encryption (SM) and non-China Encryption (SM) clients can access the domain name.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	SM2AccessOnly *bool `json:"SM2AccessOnly,omitempty" xml:"SM2AccessOnly,omitempty"`
	// The ID of the China Encryption (SM) certificate to add. This parameter is used only when SM2Enabled is set to true.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 123-cn-hangzhou
	SM2CertId *string `json:"SM2CertId,omitempty" xml:"SM2CertId,omitempty"`
	// Specifies whether to enable the China Encryption (SM) certificate.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	SM2Enabled *bool `json:"SM2Enabled,omitempty" xml:"SM2Enabled,omitempty"`
	// The TLS version to add. This parameter is used only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
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
	// The method that WAF uses to obtain the originating IP address of the client. Valid values:
	//
	// - **0*	- (default): The client traffic has not been forwarded by any Layer 7 proxy before reaching WAF.
	//
	// - **1**: WAF reads the first value in the X-Forwarded-For (XFF) header as the client IP address.
	//
	// - **2**: WAF reads the value of a custom header field that you specify as the client IP address.
	//
	// example:
	//
	// 1
	XffHeaderMode *int32 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// The list of custom header fields used to obtain the client IP address.
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
	// The IP addresses or back-to-origin domain names of the origin server corresponding to the domain name.
	Backends []*string `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	// The backup origin server IP addresses or back-to-origin domain names corresponding to the domain name.
	BackupBackends []*string `json:"BackupBackends,omitempty" xml:"BackupBackends,omitempty" type:"Repeated"`
	// Specifies whether to enable public cloud disaster recovery. Valid values:
	//
	// - **true**: Public cloud disaster recovery is enabled.
	//
	// - **false*	- (default): Public cloud disaster recovery is not enabled.
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
	// Specifies whether to enable forced HTTP back-to-origin. This parameter is used only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
	//
	// - **true**: Forced HTTP back-to-origin is enabled.
	//
	// - **false**: Forced HTTP back-to-origin is not enabled.
	//
	// example:
	//
	// true
	FocusHttpBackend *bool `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	// Specifies whether to enable HTTP/2 back-to-origin. Valid values:
	//
	// - **true**: HTTP/2 back-to-origin is enabled.
	//
	// - **false**: HTTP/2 back-to-origin is not enabled.
	//
	// example:
	//
	// true
	Http2Origin *bool `json:"Http2Origin,omitempty" xml:"Http2Origin,omitempty"`
	// The maximum number of concurrent connections for HTTP/2 back-to-origin. Valid values: 1 to 512. Default value: 128.
	//
	// example:
	//
	// 128
	Http2OriginMaxConcurrency *int32 `json:"Http2OriginMaxConcurrency,omitempty" xml:"Http2OriginMaxConcurrency,omitempty"`
	// Specifies whether to enable persistent connections. Valid values:
	//
	// - **true*	- (default): Persistent connections are enabled.
	//
	// - **false**: Persistent connections are not enabled.
	//
	// example:
	//
	// true
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// The number of requests that reuse persistent connections. Valid values: 60 to 1000. Default value: 1000.
	//
	// > The number of persistent connections to reuse after persistent connections are enabled.
	//
	// example:
	//
	// 1000
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The idle timeout period for persistent connections. Valid values: 1 to 60. Default value: 15. Unit: seconds.
	//
	// > The idle time after which a reused persistent connection is released.
	//
	// example:
	//
	// 15
	KeepaliveTimeout *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The load balancing algorithm used for back-to-origin. Valid values:
	//
	// - **iphash**: IP hash algorithm.
	//
	// - **roundRobin**: round-robin algorithm.
	//
	// - **leastTime**: Least Time algorithm. This value is available only when **ProtectionResource*	- is set to **gslb**, which indicates that the shared cluster-based intelligent load balancing is used.
	//
	// This parameter is required.
	//
	// example:
	//
	// roundRobin
	Loadbalance *string `json:"Loadbalance,omitempty" xml:"Loadbalance,omitempty"`
	// The maximum request body size. Valid values: 2 to 10. Default value: 2. Unit: GB.
	//
	// > Only the Ultimate Edition supports this feature.
	//
	// example:
	//
	// 2
	MaxBodySize *int32 `json:"MaxBodySize,omitempty" xml:"MaxBodySize,omitempty"`
	// Specifies whether the client source IP preservation feature is enabled.
	//
	// - **true**: The client source IP preservation feature is enabled. After this feature is enabled, the backend service can view the originating IP address of the client.
	//
	// - **false**: The client source IP preservation feature is not enabled.
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
	// The traffic mark field and value of the domain name, used to mark traffic processed by WAF.
	//
	// By specifying custom request header fields and corresponding values, when the access traffic of the domain name passes through WAF, WAF automatically adds the specified custom field values to the request header as traffic marks, which helps the backend service collect relevant information.
	RequestHeaders []*CreateDomainRequestRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// Specifies whether to retry when WAF fails to forward requests to the origin server. Valid values:
	//
	// - **true*	- (default): Retry is enabled.
	//
	// - **false**: Retry is not enabled.
	//
	// example:
	//
	// true
	Retry *bool `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// The hybrid cloud forwarding rules. The value is a string converted from a JSON array. Each element in the JSON array is a structure that contains the following fields:
	//
	// - **rs**: Array type | The list of back-to-origin IP addresses or back-to-origin CNAMEs.
	//
	// - **backupRs**: Array type | The list of backup back-to-origin IP addresses or back-to-origin CNAMEs. This field is required. Use [] to indicate that no backup is set.
	//
	// - **location**: String type | The name of the protection node.
	//
	// - **locationId**: Long type | The ID of the protection node.
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
	// Specifies whether to enable back-to-origin SNI. This parameter is used only when **HttpsPorts*	- is not empty, which indicates that the domain name uses HTTPS. Valid values:
	//
	// - **true**: Back-to-origin SNI is enabled.
	//
	// - **false*	- (default): Back-to-origin SNI is not enabled.
	//
	// example:
	//
	// true
	SniEnabled *bool `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	// The value of the custom SNI extension field. If this parameter is not set, the value of the **Host*	- field in the request header is used as the SNI extension field value by default.
	//
	// In most cases, you do not need to customize the SNI unless your service has special configuration requirements and you want WAF to use an SNI that is different from the actual request Host in back-to-origin requests (the custom SNI set here).
	//
	// > This parameter is required only when **SniEnabled*	- is set to **true**, which indicates that back-to-origin SNI is enabled.
	//
	// example:
	//
	// www.aliyundoc.com
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// Specifies whether to allow WAF to overwrite WL-Proxy-Client-IP. Valid values:
	//
	// - **true*	- (default): WAF is allowed to overwrite.
	//
	// - **false**: WAF is not allowed to overwrite.
	//
	// example:
	//
	// true
	WLProxyClientIp *bool `json:"WLProxyClientIp,omitempty" xml:"WLProxyClientIp,omitempty"`
	// Specifies whether to allow WAF to overwrite Web-Server-Type. Valid values:
	//
	// - **true*	- (default): WAF is allowed to overwrite.
	//
	// - **false**: WAF is not allowed to overwrite.
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
	// Specifies whether to allow WAF to overwrite X-Client-IP. Valid values:
	//
	// - **true*	- (default): WAF is allowed to overwrite.
	//
	// - **false**: WAF is not allowed to overwrite.
	//
	// example:
	//
	// true
	XClientIp *bool `json:"XClientIp,omitempty" xml:"XClientIp,omitempty"`
	// Specifies whether to allow WAF to overwrite X-True-IP. Valid values:
	//
	// - **true*	- (default): WAF is allowed to overwrite.
	//
	// - **false**: WAF is not allowed to overwrite.
	//
	// example:
	//
	// true
	XTrueIp *bool `json:"XTrueIp,omitempty" xml:"XTrueIp,omitempty"`
	// Specifies whether to use X-Forward-For-Proto to pass the protocol used by WAF. Valid values:
	//
	// - **true*	- (default): The protocol used by WAF is passed.
	//
	// - **false**: The protocol used by WAF is not passed.
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
	// The listening port.
	//
	// example:
	//
	// 80
	ListenPort *int32 `json:"ListenPort,omitempty" xml:"ListenPort,omitempty"`
	// The protocol of the listening port. Valid values:
	//
	// - **http**: The protocol of the listening port is HTTP.
	//
	// - **https**: The protocol of the listening port is HTTPS.
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
	// The specified custom request header field.
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
