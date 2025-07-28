// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v []*DescribeHybridCloudResourcesResponseBodyDomains) *DescribeHybridCloudResourcesResponseBody
	GetDomains() []*DescribeHybridCloudResourcesResponseBodyDomains
	SetRequestId(v string) *DescribeHybridCloudResourcesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeHybridCloudResourcesResponseBody
	GetTotalCount() *int64
}

type DescribeHybridCloudResourcesResponseBody struct {
	// The domain names.
	Domains []*DescribeHybridCloudResourcesResponseBodyDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 98D2AA9A-5959-5CCD-83E3-B6606232A2BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 24
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeHybridCloudResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourcesResponseBody) GetDomains() []*DescribeHybridCloudResourcesResponseBodyDomains {
	return s.Domains
}

func (s *DescribeHybridCloudResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudResourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeHybridCloudResourcesResponseBody) SetDomains(v []*DescribeHybridCloudResourcesResponseBodyDomains) *DescribeHybridCloudResourcesResponseBody {
	s.Domains = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBody) SetRequestId(v string) *DescribeHybridCloudResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBody) SetTotalCount(v int64) *DescribeHybridCloudResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridCloudResourcesResponseBodyDomains struct {
	// The CNAME assigned by WAF.
	//
	// >  This parameter is returned only if the value of **CnameEnabled*	- is true.
	//
	// example:
	//
	// 50fqmu1ci7g0xtiyxnrhgx6qdhmn****.yundunwaf5.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The domain name.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The access ID.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The listeners.
	Listen *DescribeHybridCloudResourcesResponseBodyDomainsListen `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	// The configurations of the forwarding rule.
	Redirect *DescribeHybridCloudResourcesResponseBodyDomainsRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfmvtc5z52****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The status of the domain name. Valid values:
	//
	// 	- **1:*	- The domain name is in a normal state.
	//
	// 	- **2:*	- The domain name is being created.
	//
	// 	- **3:*	- The domain name is being modified.
	//
	// 	- **4:*	- The domain name is being released.
	//
	// 	- **5:*	- WAF no longer forwards the traffic of the domain name.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 130715431409****
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeHybridCloudResourcesResponseBodyDomains) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudResourcesResponseBodyDomains) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) GetCname() *string {
	return s.Cname
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) GetDomain() *string {
	return s.Domain
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) GetId() *int64 {
	return s.Id
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) GetListen() *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	return s.Listen
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) GetRedirect() *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	return s.Redirect
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) GetUid() *string {
	return s.Uid
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetCname(v string) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.Cname = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetDomain(v string) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.Domain = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetId(v int64) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.Id = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetListen(v *DescribeHybridCloudResourcesResponseBodyDomainsListen) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.Listen = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetRedirect(v *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.Redirect = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetStatus(v int32) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.Status = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) SetUid(v string) *DescribeHybridCloudResourcesResponseBodyDomains {
	s.Uid = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomains) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridCloudResourcesResponseBodyDomainsListen struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 72***76-cn-hangzhou
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The types of cipher suites that are added. Valid values:
	//
	// 	- **1:*	- all cipher suites.
	//
	// 	- **2:*	- strong cipher suites.
	//
	// 	- **99:*	- custom cipher suites.
	//
	// example:
	//
	// 1
	CipherSuite *int32 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suites.
	//
	// >  This parameter is returned only if the value of **CipherSuite*	- is **99**.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Indicates whether TLS 1.3 is supported. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// Indicates whether exclusive IP addresses are supported. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// Indicates whether the HTTP to HTTPS redirection feature is enabled for the domain name. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	FocusHttps *bool `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	// Indicates whether HTTP/2 is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The HTTP listener ports.
	HttpPorts []*int64 `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	// The HTTPS listener ports.
	HttpsPorts []*int64 `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	// Specifies whether to enable IPv6. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	IPv6Enabled *bool `json:"IPv6Enabled,omitempty" xml:"IPv6Enabled,omitempty"`
	// The type of the protection resource. Valid values:
	//
	// 	- **share:*	- shared cluster.
	//
	// 	- **gslb:*	- shared cluster-based intelligent load balancing.
	//
	// example:
	//
	// share
	ProtectionResource *string `json:"ProtectionResource,omitempty" xml:"ProtectionResource,omitempty"`
	// The version of the Transport Layer Security (TLS) protocol. Valid values:
	//
	// 	- **tlsv1**
	//
	// 	- **tlsv1.1**
	//
	// 	- **tlsv1.2**
	//
	// example:
	//
	// tlsv1.2
	TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
	// The method that is used to obtain the actual IP address of a client. Valid values:
	//
	// 	- **0**: No Layer 7 proxies are deployed in front of WAF.
	//
	// 	- **1**: WAF reads the first value of the X-Forwarded-For (XFF) header field as the actual IP address of the client.
	//
	// 	- **2**: WAF reads the value of a custom header field as the actual IP address of the client.
	//
	// example:
	//
	// 0
	XffHeaderMode *int32 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// The custom header fields that are used to obtain the actual IP addresses of clients. The value is in the ["header1","header2",...] format.
	//
	// >  This parameter is returned only if the value of **XffHeaderMode*	- is 2.
	XffHeaders []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
}

func (s DescribeHybridCloudResourcesResponseBodyDomainsListen) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudResourcesResponseBodyDomainsListen) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) GetCertId() *string {
	return s.CertId
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) GetCipherSuite() *int32 {
	return s.CipherSuite
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) GetCustomCiphers() []*string {
	return s.CustomCiphers
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) GetEnableTLSv3() *bool {
	return s.EnableTLSv3
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) GetExclusiveIp() *bool {
	return s.ExclusiveIp
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) GetFocusHttps() *bool {
	return s.FocusHttps
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) GetHttp2Enabled() *bool {
	return s.Http2Enabled
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) GetHttpPorts() []*int64 {
	return s.HttpPorts
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) GetHttpsPorts() []*int64 {
	return s.HttpsPorts
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) GetIPv6Enabled() *bool {
	return s.IPv6Enabled
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) GetProtectionResource() *string {
	return s.ProtectionResource
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) GetTLSVersion() *string {
	return s.TLSVersion
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) GetXffHeaderMode() *int32 {
	return s.XffHeaderMode
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) GetXffHeaders() []*string {
	return s.XffHeaders
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetCertId(v string) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.CertId = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetCipherSuite(v int32) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.CipherSuite = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetCustomCiphers(v []*string) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.CustomCiphers = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetEnableTLSv3(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.EnableTLSv3 = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetExclusiveIp(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.ExclusiveIp = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetFocusHttps(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.FocusHttps = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetHttp2Enabled(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.Http2Enabled = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetHttpPorts(v []*int64) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.HttpPorts = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetHttpsPorts(v []*int64) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.HttpsPorts = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetIPv6Enabled(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.IPv6Enabled = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetProtectionResource(v string) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.ProtectionResource = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetTLSVersion(v string) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.TLSVersion = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetXffHeaderMode(v int32) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.XffHeaderMode = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) SetXffHeaders(v []*string) *DescribeHybridCloudResourcesResponseBodyDomainsListen {
	s.XffHeaders = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsListen) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridCloudResourcesResponseBodyDomainsRedirect struct {
	// The IP addresses or domain names of the origin server.
	Backends []*string `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	// Indicates whether the public cloud disaster recovery feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CnameEnabled *bool `json:"CnameEnabled,omitempty" xml:"CnameEnabled,omitempty"`
	// The timeout period for connections. Unit: seconds. Valid values: 5 to 120.
	//
	// example:
	//
	// 120
	ConnectTimeout *int64 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// Indicates whether the HTTPS to HTTP redirection feature is enabled for back-to-origin requests. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	FocusHttpBackend *bool `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	// Indicates whether the persistent connection feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// The number of reused persistent connections. Valid values: 60 to 1000.
	//
	// >  This parameter indicates the number of reused persistent connections after the persistent connection feature is enabled.
	//
	// example:
	//
	// 1000
	KeepaliveRequests *int64 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The timeout period for persistent connections that are in the Idle state. Unit: seconds. Valid values: 1 to 60. Default value: 15.
	//
	// >  This parameter indicates the period of time during which a reused persistent connection can remain in the Idle state before the persistent connection is released.
	//
	// example:
	//
	// 15
	KeepaliveTimeout *int64 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The load balancing algorithm that is used to forward requests to the origin server. Valid values:
	//
	// 	- **iphash**
	//
	// 	- **roundRobin**
	//
	// 	- **leastTime**
	//
	// example:
	//
	// iphash
	Loadbalance *string `json:"Loadbalance,omitempty" xml:"Loadbalance,omitempty"`
	// The timeout period for read connections. Unit: seconds. Valid values: 5 to 1800.
	//
	// example:
	//
	// 200
	ReadTimeout *int64 `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	// The key-value pair that is used to label requests that pass through WAF.
	RequestHeaders []*DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// Indicates whether WAF retries forwarding requests if requests fail to be forwarded to the origin server. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Retry *bool `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// The forwarding rules that are configured for the domain name. This parameter is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that contains the following fields:
	//
	// 	- **rs**: the back-to-origin IP addresses or CNAMEs. The value is of the ARRAY type.
	//
	// 	- **location**: the name of the protection node. The value is of the STRING type.
	//
	// 	- **locationId**: the ID of the protection node. The value is of the LONG type.
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
	// Indicates whether the origin Server Name Indication (SNI) feature is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	SniEnabled *bool `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	// The value of the custom SNI field. If the parameter is left empty, the value of the **Host*	- field in the request header is automatically used as the value of the SNI field.
	//
	// >  This parameter is returned only if the value of **SniEnabled*	- is **true**.
	//
	// example:
	//
	// www.aliyundoc.com
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// The timeout period for write connections. Unit: seconds. Valid values: 5 to 1800.
	//
	// example:
	//
	// 200
	WriteTimeout *int64 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
}

func (s DescribeHybridCloudResourcesResponseBodyDomainsRedirect) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GetBackends() []*string {
	return s.Backends
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GetCnameEnabled() *bool {
	return s.CnameEnabled
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GetConnectTimeout() *int64 {
	return s.ConnectTimeout
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GetFocusHttpBackend() *bool {
	return s.FocusHttpBackend
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GetKeepalive() *bool {
	return s.Keepalive
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GetKeepaliveRequests() *int64 {
	return s.KeepaliveRequests
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GetKeepaliveTimeout() *int64 {
	return s.KeepaliveTimeout
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GetLoadbalance() *string {
	return s.Loadbalance
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GetReadTimeout() *int64 {
	return s.ReadTimeout
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GetRequestHeaders() []*DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders {
	return s.RequestHeaders
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GetRetry() *bool {
	return s.Retry
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GetRoutingRules() *string {
	return s.RoutingRules
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GetSniEnabled() *bool {
	return s.SniEnabled
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GetSniHost() *string {
	return s.SniHost
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) GetWriteTimeout() *int64 {
	return s.WriteTimeout
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetBackends(v []*string) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.Backends = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetCnameEnabled(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.CnameEnabled = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetConnectTimeout(v int64) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.ConnectTimeout = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetFocusHttpBackend(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.FocusHttpBackend = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetKeepalive(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.Keepalive = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetKeepaliveRequests(v int64) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.KeepaliveRequests = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetKeepaliveTimeout(v int64) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.KeepaliveTimeout = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetLoadbalance(v string) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.Loadbalance = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetReadTimeout(v int64) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.ReadTimeout = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetRequestHeaders(v []*DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.RequestHeaders = v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetRetry(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.Retry = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetRoutingRules(v string) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.RoutingRules = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetSniEnabled(v bool) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.SniEnabled = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetSniHost(v string) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.SniHost = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) SetWriteTimeout(v int64) *DescribeHybridCloudResourcesResponseBodyDomainsRedirect {
	s.WriteTimeout = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirect) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders struct {
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

func (s DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders) GetKey() *string {
	return s.Key
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders) GetValue() *string {
	return s.Value
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders) SetKey(v string) *DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders {
	s.Key = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders) SetValue(v string) *DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders {
	s.Value = &v
	return s
}

func (s *DescribeHybridCloudResourcesResponseBodyDomainsRedirectRequestHeaders) Validate() error {
	return dara.Validate(s)
}
