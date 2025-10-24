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
	// The mode in which you want to add the domain name to WAF. Valid values:
	//
	// 	- **share:*	- adds the domain name to WAF in CNAME record mode. This is the default value.
	//
	// 	- **hybrid_cloud_cname:*	- adds the domain name to WAF in hybrid cloud reverse proxy mode.
	//
	// example:
	//
	// share
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The domain name that you want to add to WAF.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
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
	Listen *CreateDomainRequestListen `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	// The forwarding configurations.
	//
	// This parameter is required.
	Redirect *CreateDomainRequestRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// The region where the WAF instance resides. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
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
	// The ID of the certificate that you want to add. This parameter is available only if you specify **HttpsPorts**.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 123
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The type of the cipher suites that you want to add. This parameter is available only if you specify **HttpsPorts**. Valid values:
	//
	// 	- **1**: all cipher suites.
	//
	// 	- **2**: strong cipher suites. This value is available only if you set **TLSVersion*	- to **tlsv1.2**.
	//
	// 	- **99**: custom cipher suites.
	//
	// example:
	//
	// 2
	CipherSuite *int32 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suites that you want to add.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Specifies whether to support TLS 1.3. This parameter is available only if you specify **HttpsPorts**. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// Specifies whether to enable the exclusive IP address feature. This parameter is available only if you set **IPv6Enabled*	- to **false*	- and **ProtectionResource*	- to **share**. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Specifies whether to enable force redirect from HTTP to HTTPS for received requests. This parameter is available only if you specify HttpsPorts and leave HttpPorts empty. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	FocusHttps *bool `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	// Specifies whether to enable HTTP/2. This parameter is available only if you specify **HttpsPorts**. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The HTTP listener ports.
	HttpPorts []*int32 `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	// The HTTPS listener ports.
	HttpsPorts []*int32 `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	// Specifies whether to enable IPv6 protection. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	IPv6Enabled *bool `json:"IPv6Enabled,omitempty" xml:"IPv6Enabled,omitempty"`
	// The type of the protection resource. Valid values:
	//
	// 	- **share*	- (default): a shared cluster.
	//
	// 	- **gslb**: shared cluster-based intelligent load balancing.
	//
	// example:
	//
	// share
	ProtectionResource *string `json:"ProtectionResource,omitempty" xml:"ProtectionResource,omitempty"`
	// Specifies whether to allow access only from SM certificate-based clients. This parameter is available only if you set SM2Enabled to true.
	//
	// 	- true
	//
	// 	- false
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	SM2AccessOnly *bool `json:"SM2AccessOnly,omitempty" xml:"SM2AccessOnly,omitempty"`
	// The ID of the SM certificate that you want to add. This parameter is available only if you set SM2Enabled to true.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 123-cn-hangzhou
	SM2CertId *string `json:"SM2CertId,omitempty" xml:"SM2CertId,omitempty"`
	// Specifies whether to add an SM certificate.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	SM2Enabled *bool `json:"SM2Enabled,omitempty" xml:"SM2Enabled,omitempty"`
	// The Transport Layer Security (TLS) version that you want to add. This parameter is available only if you specify **HttpsPorts**. Valid values:
	//
	// 	- **tlsv1**
	//
	// 	- **tlsv1.1**
	//
	// 	- **tlsv1.2**
	//
	// example:
	//
	// tlsv1
	TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
	// The method that is used to obtain the originating IP address of a client. Valid values:
	//
	// 	- **0*	- (default): Client traffic is not filtered by a Layer 7 proxy before the traffic reaches WAF.
	//
	// 	- **1**: WAF reads the first value of the X-Forwarded-For (XFF) header field as the originating IP address of the client.
	//
	// 	- **2**: WAF reads the value of a custom header field as the originating IP address of the client.
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
	BackendPorts []*CreateDomainRequestRedirectBackendPorts `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
	// The IP addresses or domain names of the origin server.
	Backends []*string `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	// The secondary IP addresses or domain names of the origin server.
	BackupBackends []*string `json:"BackupBackends,omitempty" xml:"BackupBackends,omitempty" type:"Repeated"`
	// Specifies whether to enable the public cloud disaster recovery feature. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	CnameEnabled *bool `json:"CnameEnabled,omitempty" xml:"CnameEnabled,omitempty"`
	// The timeout period of connections. Unit: seconds. Valid values: 1 to 3600.
	//
	// example:
	//
	// 120
	ConnectTimeout *int32 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// Specifies whether to enable force redirect from HTTPS to HTTP for back-to-origin requests. This parameter is available only if you specify **HttpsPorts**. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	FocusHttpBackend *bool `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	// Specifies whether to enable the persistent connection feature. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// The number of reused persistent connections. Valid values: 60 to 1000.
	//
	// >  This parameter specifies the number of persistent connections that can be reused after you enable the persistent connection feature.
	//
	// example:
	//
	// 1000
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The timeout period of idle persistent connections. Valid values: 1 to 60. Default value: 15. Unit: seconds.
	//
	// >  This parameter specifies the period of time after which an idle persistent connection is closed.
	//
	// example:
	//
	// 15
	KeepaliveTimeout *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The load balancing algorithm that you want to use to forward requests to the origin server. Valid values:
	//
	// 	- **iphash**
	//
	// 	- **roundRobin**
	//
	// 	- **leastTime**: This value is available only if you set **ProtectionResource*	- to **gslb**.
	//
	// This parameter is required.
	//
	// example:
	//
	// roundRobin
	Loadbalance *string `json:"Loadbalance,omitempty" xml:"Loadbalance,omitempty"`
	// The timeout period of read connections. Unit: seconds. Valid values: 1 to 3600.
	//
	// example:
	//
	// 200
	ReadTimeout *int32 `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	// The custom header fields, which are key-value pairs. The fields are used to mark requests that pass through WAF.
	//
	// When a request passes through WAF, WAF automatically adds the custom header fields to the request to mark the request. This way, the backend service can identify requests that are processed by WAF.
	RequestHeaders []*CreateDomainRequestRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// Specifies whether WAF retries if WAF fails to forward requests to the origin server. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Retry *bool `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// The forwarding rules for the hybrid cloud mode. The value is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that contains the following fields:
	//
	// 	- **rs**: the back-to-origin IP addresses or CNAMEs. Data type: array.
	//
	// 	- **location**: the name of the protection node. Data type: string.
	//
	// 	- **locationId**: the ID of the protection node. Data type: long.
	//
	// example:
	//
	// [
	//
	//       {
	//
	//             "rs": [
	//
	//                   "1.1.XX.XX"
	//
	//             ],
	//
	//             "locationId": 535,
	//
	//             "location": "test1111"
	//
	//       }
	//
	// ]
	RoutingRules *string `json:"RoutingRules,omitempty" xml:"RoutingRules,omitempty"`
	// Specifies whether to enable the Server Name Indication (SNI) feature for back-to-origin requests. This parameter is available only if you specify **HttpsPorts**. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	SniEnabled *bool `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	// The custom value of the SNI field. If you do not specify this parameter, the value of the **Host*	- header field is automatically used. In most cases, you do not need to specify a custom value for the SNI field. However, if you want WAF to use an SNI field whose value is different from the value of the Host header field in back-to-origin requests, you can specify a custom value for the SNI field.
	//
	// >  This parameter is required only if you set **SniEnabled*	- to **true**.
	//
	// example:
	//
	// www.aliyundoc.com
	SniHost         *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	WLProxyClientIp *bool   `json:"WLProxyClientIp,omitempty" xml:"WLProxyClientIp,omitempty"`
	WebServerType   *bool   `json:"WebServerType,omitempty" xml:"WebServerType,omitempty"`
	// The timeout period of write connections. Unit: seconds. Valid values: 1 to 3600.
	//
	// example:
	//
	// 200
	WriteTimeout *int32 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
	XClientIp    *bool  `json:"XClientIp,omitempty" xml:"XClientIp,omitempty"`
	XTrueIp      *bool  `json:"XTrueIp,omitempty" xml:"XTrueIp,omitempty"`
	// Specifies whether to use the X-Forward-For-Proto header field to pass the protocol used by WAF to forward requests to the origin server. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false**
	//
	// example:
	//
	// true
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
	BackendPort *int32  `json:"BackendPort,omitempty" xml:"BackendPort,omitempty"`
	ListenPort  *int32  `json:"ListenPort,omitempty" xml:"ListenPort,omitempty"`
	Protocol    *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
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
	// The key of the tag.
	//
	// example:
	//
	// Tagkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
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
