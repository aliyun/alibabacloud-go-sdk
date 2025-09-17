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
	// The mode in which you want to add the domain name to WAF. Set the value to share.
	//
	// 	- **share:*	- adds the domain name to WAF in CNAME record mode. This is the default value.
	//
	// example:
	//
	// share
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The domain name whose access configurations you want to modify.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain   *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
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
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
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
	return dara.Validate(s)
}

type ModifyDomainRequestListen struct {
	// The ID of the certificate that you want to add.
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
	// The custom cipher suites that you want to add. This parameter is available only if you set **CipherSuite*	- to **99**.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Specifies whether to support TLS 1.3. This parameter is available only if you specify **HttpsPorts**. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// Specifies whether to enable the exclusive IP address feature. This parameter is available only if you set **IPv6Enabled*	- to false and **ProtectionResource*	- to **share**. Valid values:
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Specifies whether to enable force redirect from HTTP to HTTPS for received requests. This parameter is available only if you specify **HttpsPorts*	- and leave **HttpPorts*	- empty. Valid values:
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
	// The HTTP listener ports. Specify the value in the [**port1,port2,...**] format.
	HttpPorts []*int32 `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	// The HTTPS listener ports. Specify the value in the [**port1,port2,...**] format.
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
	// example:
	//
	// true
	SM2AccessOnly *bool `json:"SM2AccessOnly,omitempty" xml:"SM2AccessOnly,omitempty"`
	// The ID of the SM certificate that you want to add. This parameter is available only if you set SM2Enabled to true.
	//
	// example:
	//
	// 123-cn-hangzhou
	SM2CertId *string `json:"SM2CertId,omitempty" xml:"SM2CertId,omitempty"`
	// Specifies whether to add an SM certificate.
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
	// 2
	XffHeaderMode *int32 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// The custom header fields that are used to obtain the originating IP address of a client. Specify the value in the **["header1","header2",...]*	- format.
	//
	// >  This parameter is required only if you set **XffHeaderMode*	- to 2.
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
	BackendPorts []*ModifyDomainRequestRedirectBackendPorts `json:"BackendPorts,omitempty" xml:"BackendPorts,omitempty" type:"Repeated"`
	// The IP addresses or domain names of the origin server. You cannot specify both IP addresses and domain names. If you specify domain names, the domain names can be resolved only to IPv4 addresses.
	//
	// 	- If you specify IP addresses, specify the value in the **["ip1","ip2",...]*	- format. You can enter up to 20 IP addresses.
	//
	// 	- If you specify domain names, specify the value in the **["domain"]*	- format. You can enter up to 20 domain names.
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
	// iphash
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
	RequestHeaders []*ModifyDomainRequestRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
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
	// >  This parameter is required only if you set **SniEnabled*	- to true.
	//
	// example:
	//
	// www.aliyundoc.com
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// if can be null:
	// true
	WLProxyClientIp *bool `json:"WLProxyClientIp,omitempty" xml:"WLProxyClientIp,omitempty"`
	// if can be null:
	// true
	WebServerType *bool `json:"WebServerType,omitempty" xml:"WebServerType,omitempty"`
	// The timeout period of write connections. Unit: seconds. Valid values: 1 to 3600.
	//
	// example:
	//
	// 200
	WriteTimeout *int32 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
	// if can be null:
	// true
	XClientIp *bool `json:"XClientIp,omitempty" xml:"XClientIp,omitempty"`
	// if can be null:
	// true
	XTrueIp *bool `json:"XTrueIp,omitempty" xml:"XTrueIp,omitempty"`
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
	return dara.Validate(s)
}

type ModifyDomainRequestRedirectBackendPorts struct {
	BackendPort *int32  `json:"BackendPort,omitempty" xml:"BackendPort,omitempty"`
	ListenPort  *int32  `json:"ListenPort,omitempty" xml:"ListenPort,omitempty"`
	Protocol    *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
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
