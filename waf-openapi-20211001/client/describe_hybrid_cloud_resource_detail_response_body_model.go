// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudResourceDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v *DescribeHybridCloudResourceDetailResponseBodyDomain) *DescribeHybridCloudResourceDetailResponseBody
	GetDomain() *DescribeHybridCloudResourceDetailResponseBodyDomain
	SetRequestId(v string) *DescribeHybridCloudResourceDetailResponseBody
	GetRequestId() *string
}

type DescribeHybridCloudResourceDetailResponseBody struct {
	Domain *DescribeHybridCloudResourceDetailResponseBodyDomain `json:"Domain,omitempty" xml:"Domain,omitempty" type:"Struct"`
	// example:
	//
	// 66A98669-CC6E-4F3E-80A6-3014697B11AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHybridCloudResourceDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudResourceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourceDetailResponseBody) GetDomain() *DescribeHybridCloudResourceDetailResponseBodyDomain {
	return s.Domain
}

func (s *DescribeHybridCloudResourceDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHybridCloudResourceDetailResponseBody) SetDomain(v *DescribeHybridCloudResourceDetailResponseBodyDomain) *DescribeHybridCloudResourceDetailResponseBody {
	s.Domain = v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBody) SetRequestId(v string) *DescribeHybridCloudResourceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridCloudResourceDetailResponseBodyDomain struct {
	// CNAME
	//
	// example:
	//
	// kdmqyi3ck7xogegxpiyfpb0fj21mgkxn.****.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// example:
	//
	// www.*****.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// id
	//
	// example:
	//
	// 31323
	Id       *int64                                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	Listen   *DescribeHybridCloudResourceDetailResponseBodyDomainListen   `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	Redirect *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// example:
	//
	// rg-***aby
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1046011128270720
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s DescribeHybridCloudResourceDetailResponseBodyDomain) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudResourceDetailResponseBodyDomain) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) GetCname() *string {
	return s.Cname
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) GetDomain() *string {
	return s.Domain
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) GetId() *int64 {
	return s.Id
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) GetListen() *DescribeHybridCloudResourceDetailResponseBodyDomainListen {
	return s.Listen
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) GetRedirect() *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	return s.Redirect
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) GetUid() *string {
	return s.Uid
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) SetCname(v string) *DescribeHybridCloudResourceDetailResponseBodyDomain {
	s.Cname = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) SetDomain(v string) *DescribeHybridCloudResourceDetailResponseBodyDomain {
	s.Domain = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) SetId(v int64) *DescribeHybridCloudResourceDetailResponseBodyDomain {
	s.Id = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) SetListen(v *DescribeHybridCloudResourceDetailResponseBodyDomainListen) *DescribeHybridCloudResourceDetailResponseBodyDomain {
	s.Listen = v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) SetRedirect(v *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) *DescribeHybridCloudResourceDetailResponseBodyDomain {
	s.Redirect = v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudResourceDetailResponseBodyDomain {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) SetStatus(v int32) *DescribeHybridCloudResourceDetailResponseBodyDomain {
	s.Status = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) SetUid(v string) *DescribeHybridCloudResourceDetailResponseBodyDomain {
	s.Uid = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomain) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridCloudResourceDetailResponseBodyDomainListen struct {
	// example:
	//
	// 19312542-cn-hangzhou
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// example:
	//
	// 0
	CipherSuite   *int32    `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// example:
	//
	// true
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// example:
	//
	// true
	ExclusiveIp *bool `json:"ExclusiveIp,omitempty" xml:"ExclusiveIp,omitempty"`
	// example:
	//
	// true
	FocusHttps *bool `json:"FocusHttps,omitempty" xml:"FocusHttps,omitempty"`
	// example:
	//
	// true
	Http2Enabled *bool    `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	HttpPorts    []*int64 `json:"HttpPorts,omitempty" xml:"HttpPorts,omitempty" type:"Repeated"`
	HttpsPorts   []*int64 `json:"HttpsPorts,omitempty" xml:"HttpsPorts,omitempty" type:"Repeated"`
	// example:
	//
	// true
	IPv6Enabled *bool `json:"IPv6Enabled,omitempty" xml:"IPv6Enabled,omitempty"`
	// example:
	//
	// share
	ProtectionResource *string `json:"ProtectionResource,omitempty" xml:"ProtectionResource,omitempty"`
	// example:
	//
	// tlsv1
	TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
	// example:
	//
	// 1
	XffHeaderMode *int32    `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	XffHeaders    []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
}

func (s DescribeHybridCloudResourceDetailResponseBodyDomainListen) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudResourceDetailResponseBodyDomainListen) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) GetCertId() *string {
	return s.CertId
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) GetCipherSuite() *int32 {
	return s.CipherSuite
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) GetCustomCiphers() []*string {
	return s.CustomCiphers
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) GetEnableTLSv3() *bool {
	return s.EnableTLSv3
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) GetExclusiveIp() *bool {
	return s.ExclusiveIp
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) GetFocusHttps() *bool {
	return s.FocusHttps
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) GetHttp2Enabled() *bool {
	return s.Http2Enabled
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) GetHttpPorts() []*int64 {
	return s.HttpPorts
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) GetHttpsPorts() []*int64 {
	return s.HttpsPorts
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) GetIPv6Enabled() *bool {
	return s.IPv6Enabled
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) GetProtectionResource() *string {
	return s.ProtectionResource
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) GetTLSVersion() *string {
	return s.TLSVersion
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) GetXffHeaderMode() *int32 {
	return s.XffHeaderMode
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) GetXffHeaders() []*string {
	return s.XffHeaders
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) SetCertId(v string) *DescribeHybridCloudResourceDetailResponseBodyDomainListen {
	s.CertId = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) SetCipherSuite(v int32) *DescribeHybridCloudResourceDetailResponseBodyDomainListen {
	s.CipherSuite = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) SetCustomCiphers(v []*string) *DescribeHybridCloudResourceDetailResponseBodyDomainListen {
	s.CustomCiphers = v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) SetEnableTLSv3(v bool) *DescribeHybridCloudResourceDetailResponseBodyDomainListen {
	s.EnableTLSv3 = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) SetExclusiveIp(v bool) *DescribeHybridCloudResourceDetailResponseBodyDomainListen {
	s.ExclusiveIp = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) SetFocusHttps(v bool) *DescribeHybridCloudResourceDetailResponseBodyDomainListen {
	s.FocusHttps = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) SetHttp2Enabled(v bool) *DescribeHybridCloudResourceDetailResponseBodyDomainListen {
	s.Http2Enabled = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) SetHttpPorts(v []*int64) *DescribeHybridCloudResourceDetailResponseBodyDomainListen {
	s.HttpPorts = v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) SetHttpsPorts(v []*int64) *DescribeHybridCloudResourceDetailResponseBodyDomainListen {
	s.HttpsPorts = v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) SetIPv6Enabled(v bool) *DescribeHybridCloudResourceDetailResponseBodyDomainListen {
	s.IPv6Enabled = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) SetProtectionResource(v string) *DescribeHybridCloudResourceDetailResponseBodyDomainListen {
	s.ProtectionResource = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) SetTLSVersion(v string) *DescribeHybridCloudResourceDetailResponseBodyDomainListen {
	s.TLSVersion = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) SetXffHeaderMode(v int32) *DescribeHybridCloudResourceDetailResponseBodyDomainListen {
	s.XffHeaderMode = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) SetXffHeaders(v []*string) *DescribeHybridCloudResourceDetailResponseBodyDomainListen {
	s.XffHeaders = v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainListen) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridCloudResourceDetailResponseBodyDomainRedirect struct {
	Backends []*string `json:"Backends,omitempty" xml:"Backends,omitempty" type:"Repeated"`
	// example:
	//
	// true
	CnameEnabled *bool `json:"CnameEnabled,omitempty" xml:"CnameEnabled,omitempty"`
	// example:
	//
	// 1
	ConnectTimeout *int64 `json:"ConnectTimeout,omitempty" xml:"ConnectTimeout,omitempty"`
	// example:
	//
	// true
	FocusHttpBackend *bool `json:"FocusHttpBackend,omitempty" xml:"FocusHttpBackend,omitempty"`
	// example:
	//
	// true
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// example:
	//
	// 1000
	KeepaliveRequests *int64 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// example:
	//
	// 1
	KeepaliveTimeout *int64 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// example:
	//
	// iphash
	Loadbalance *string `json:"Loadbalance,omitempty" xml:"Loadbalance,omitempty"`
	// example:
	//
	// 1
	ReadTimeout    *int64                                                                       `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	RequestHeaders []*DescribeHybridCloudResourceDetailResponseBodyDomainRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Retry *bool `json:"Retry,omitempty" xml:"Retry,omitempty"`
	// example:
	//
	// [{\\"backupRs\\":[],\\"location\\":\\"v3-test\\",\\"locationId\\":1148,\\"rs\\":[\\"39.98.217.197\\",\\"2.2.2.2\\"]}]
	RoutingRules *string `json:"RoutingRules,omitempty" xml:"RoutingRules,omitempty"`
	// example:
	//
	// true
	SniEnabled *bool `json:"SniEnabled,omitempty" xml:"SniEnabled,omitempty"`
	// example:
	//
	// eew111
	SniHost *string `json:"SniHost,omitempty" xml:"SniHost,omitempty"`
	// example:
	//
	// 1
	WriteTimeout *int64 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
}

func (s DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GetBackends() []*string {
	return s.Backends
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GetCnameEnabled() *bool {
	return s.CnameEnabled
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GetConnectTimeout() *int64 {
	return s.ConnectTimeout
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GetFocusHttpBackend() *bool {
	return s.FocusHttpBackend
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GetKeepalive() *bool {
	return s.Keepalive
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GetKeepaliveRequests() *int64 {
	return s.KeepaliveRequests
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GetKeepaliveTimeout() *int64 {
	return s.KeepaliveTimeout
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GetLoadbalance() *string {
	return s.Loadbalance
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GetReadTimeout() *int64 {
	return s.ReadTimeout
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GetRequestHeaders() []*DescribeHybridCloudResourceDetailResponseBodyDomainRedirectRequestHeaders {
	return s.RequestHeaders
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GetRetry() *bool {
	return s.Retry
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GetRoutingRules() *string {
	return s.RoutingRules
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GetSniEnabled() *bool {
	return s.SniEnabled
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GetSniHost() *string {
	return s.SniHost
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) GetWriteTimeout() *int64 {
	return s.WriteTimeout
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) SetBackends(v []*string) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	s.Backends = v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) SetCnameEnabled(v bool) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	s.CnameEnabled = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) SetConnectTimeout(v int64) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	s.ConnectTimeout = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) SetFocusHttpBackend(v bool) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	s.FocusHttpBackend = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) SetKeepalive(v bool) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	s.Keepalive = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) SetKeepaliveRequests(v int64) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	s.KeepaliveRequests = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) SetKeepaliveTimeout(v int64) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	s.KeepaliveTimeout = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) SetLoadbalance(v string) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	s.Loadbalance = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) SetReadTimeout(v int64) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	s.ReadTimeout = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) SetRequestHeaders(v []*DescribeHybridCloudResourceDetailResponseBodyDomainRedirectRequestHeaders) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	s.RequestHeaders = v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) SetRetry(v bool) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	s.Retry = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) SetRoutingRules(v string) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	s.RoutingRules = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) SetSniEnabled(v bool) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	s.SniEnabled = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) SetSniHost(v string) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	s.SniHost = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) SetWriteTimeout(v int64) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect {
	s.WriteTimeout = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirect) Validate() error {
	return dara.Validate(s)
}

type DescribeHybridCloudResourceDetailResponseBodyDomainRedirectRequestHeaders struct {
	// example:
	//
	// L2x1ZmZ5L2NvcmUvYXBwcy9tLnl1bmR1bi53YWYuMS9wbHVnaW5z
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 9506360478730
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeHybridCloudResourceDetailResponseBodyDomainRedirectRequestHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudResourceDetailResponseBodyDomainRedirectRequestHeaders) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirectRequestHeaders) GetKey() *string {
	return s.Key
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirectRequestHeaders) GetValue() *string {
	return s.Value
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirectRequestHeaders) SetKey(v string) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirectRequestHeaders {
	s.Key = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirectRequestHeaders) SetValue(v string) *DescribeHybridCloudResourceDetailResponseBodyDomainRedirectRequestHeaders {
	s.Value = &v
	return s
}

func (s *DescribeHybridCloudResourceDetailResponseBodyDomainRedirectRequestHeaders) Validate() error {
	return dara.Validate(s)
}
