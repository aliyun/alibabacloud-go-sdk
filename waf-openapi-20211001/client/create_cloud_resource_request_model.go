// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateCloudResourceRequest
	GetInstanceId() *string
	SetListen(v *CreateCloudResourceRequestListen) *CreateCloudResourceRequest
	GetListen() *CreateCloudResourceRequestListen
	SetOwnerUserId(v string) *CreateCloudResourceRequest
	GetOwnerUserId() *string
	SetRedirect(v *CreateCloudResourceRequestRedirect) *CreateCloudResourceRequest
	GetRedirect() *CreateCloudResourceRequestRedirect
	SetRegionId(v string) *CreateCloudResourceRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *CreateCloudResourceRequest
	GetResourceManagerResourceGroupId() *string
	SetTag(v []*CreateCloudResourceRequestTag) *CreateCloudResourceRequest
	GetTag() []*CreateCloudResourceRequestTag
}

type CreateCloudResourceRequest struct {
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The listener configurations.
	//
	// This parameter is required.
	Listen *CreateCloudResourceRequestListen `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 123
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The forwarding configurations.
	Redirect *CreateCloudResourceRequestRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: the Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
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
	Tag []*CreateCloudResourceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateCloudResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateCloudResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateCloudResourceRequest) GetListen() *CreateCloudResourceRequestListen {
	return s.Listen
}

func (s *CreateCloudResourceRequest) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *CreateCloudResourceRequest) GetRedirect() *CreateCloudResourceRequestRedirect {
	return s.Redirect
}

func (s *CreateCloudResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCloudResourceRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *CreateCloudResourceRequest) GetTag() []*CreateCloudResourceRequestTag {
	return s.Tag
}

func (s *CreateCloudResourceRequest) SetInstanceId(v string) *CreateCloudResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateCloudResourceRequest) SetListen(v *CreateCloudResourceRequestListen) *CreateCloudResourceRequest {
	s.Listen = v
	return s
}

func (s *CreateCloudResourceRequest) SetOwnerUserId(v string) *CreateCloudResourceRequest {
	s.OwnerUserId = &v
	return s
}

func (s *CreateCloudResourceRequest) SetRedirect(v *CreateCloudResourceRequestRedirect) *CreateCloudResourceRequest {
	s.Redirect = v
	return s
}

func (s *CreateCloudResourceRequest) SetRegionId(v string) *CreateCloudResourceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCloudResourceRequest) SetResourceManagerResourceGroupId(v string) *CreateCloudResourceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *CreateCloudResourceRequest) SetTag(v []*CreateCloudResourceRequestTag) *CreateCloudResourceRequest {
	s.Tag = v
	return s
}

func (s *CreateCloudResourceRequest) Validate() error {
	return dara.Validate(s)
}

type CreateCloudResourceRequestListen struct {
	// The certificates.
	Certificates []*CreateCloudResourceRequestListenCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
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
	// 1
	CipherSuite *int32 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suites that you want to add. This parameter is available only if you set **CipherSuite*	- to **99**.
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
	// The port of the cloud service.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type. Valid values:
	//
	// 	- **http**
	//
	// 	- **https**
	//
	// This parameter is required.
	//
	// example:
	//
	// http
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The instance ID of the cloud service.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1*****
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The type of the cloud service that you want to add. Valid values:
	//
	// 	- **clb4**: Layer 4 CLB.
	//
	// 	- **clb7**: Layer 7 CLB.
	//
	// 	- **ecs**: ECS.
	//
	// 	- **nlb**: Network Load Balancer (NLB).
	//
	// This parameter is required.
	//
	// example:
	//
	// clb4
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
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
}

func (s CreateCloudResourceRequestListen) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudResourceRequestListen) GoString() string {
	return s.String()
}

func (s *CreateCloudResourceRequestListen) GetCertificates() []*CreateCloudResourceRequestListenCertificates {
	return s.Certificates
}

func (s *CreateCloudResourceRequestListen) GetCipherSuite() *int32 {
	return s.CipherSuite
}

func (s *CreateCloudResourceRequestListen) GetCustomCiphers() []*string {
	return s.CustomCiphers
}

func (s *CreateCloudResourceRequestListen) GetEnableTLSv3() *bool {
	return s.EnableTLSv3
}

func (s *CreateCloudResourceRequestListen) GetHttp2Enabled() *bool {
	return s.Http2Enabled
}

func (s *CreateCloudResourceRequestListen) GetPort() *int32 {
	return s.Port
}

func (s *CreateCloudResourceRequestListen) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateCloudResourceRequestListen) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *CreateCloudResourceRequestListen) GetResourceProduct() *string {
	return s.ResourceProduct
}

func (s *CreateCloudResourceRequestListen) GetTLSVersion() *string {
	return s.TLSVersion
}

func (s *CreateCloudResourceRequestListen) SetCertificates(v []*CreateCloudResourceRequestListenCertificates) *CreateCloudResourceRequestListen {
	s.Certificates = v
	return s
}

func (s *CreateCloudResourceRequestListen) SetCipherSuite(v int32) *CreateCloudResourceRequestListen {
	s.CipherSuite = &v
	return s
}

func (s *CreateCloudResourceRequestListen) SetCustomCiphers(v []*string) *CreateCloudResourceRequestListen {
	s.CustomCiphers = v
	return s
}

func (s *CreateCloudResourceRequestListen) SetEnableTLSv3(v bool) *CreateCloudResourceRequestListen {
	s.EnableTLSv3 = &v
	return s
}

func (s *CreateCloudResourceRequestListen) SetHttp2Enabled(v bool) *CreateCloudResourceRequestListen {
	s.Http2Enabled = &v
	return s
}

func (s *CreateCloudResourceRequestListen) SetPort(v int32) *CreateCloudResourceRequestListen {
	s.Port = &v
	return s
}

func (s *CreateCloudResourceRequestListen) SetProtocol(v string) *CreateCloudResourceRequestListen {
	s.Protocol = &v
	return s
}

func (s *CreateCloudResourceRequestListen) SetResourceInstanceId(v string) *CreateCloudResourceRequestListen {
	s.ResourceInstanceId = &v
	return s
}

func (s *CreateCloudResourceRequestListen) SetResourceProduct(v string) *CreateCloudResourceRequestListen {
	s.ResourceProduct = &v
	return s
}

func (s *CreateCloudResourceRequestListen) SetTLSVersion(v string) *CreateCloudResourceRequestListen {
	s.TLSVersion = &v
	return s
}

func (s *CreateCloudResourceRequestListen) Validate() error {
	return dara.Validate(s)
}

type CreateCloudResourceRequestListenCertificates struct {
	// The type of the certificate. Valid values:
	//
	// 	- **default**: default certificate.
	//
	// 	- **extension**: additional certificate.
	//
	// example:
	//
	// default
	AppliedType *string `json:"AppliedType,omitempty" xml:"AppliedType,omitempty"`
	// The ID of the certificate that you want to add.
	//
	// >  You can call the [DescribeCertificates](https://help.aliyun.com/document_detail/160783.html) operation to query the IDs of all SSL certificates that are associated with a domain name.
	//
	// example:
	//
	// 123-cn-hangzhou
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s CreateCloudResourceRequestListenCertificates) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudResourceRequestListenCertificates) GoString() string {
	return s.String()
}

func (s *CreateCloudResourceRequestListenCertificates) GetAppliedType() *string {
	return s.AppliedType
}

func (s *CreateCloudResourceRequestListenCertificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *CreateCloudResourceRequestListenCertificates) SetAppliedType(v string) *CreateCloudResourceRequestListenCertificates {
	s.AppliedType = &v
	return s
}

func (s *CreateCloudResourceRequestListenCertificates) SetCertificateId(v string) *CreateCloudResourceRequestListenCertificates {
	s.CertificateId = &v
	return s
}

func (s *CreateCloudResourceRequestListenCertificates) Validate() error {
	return dara.Validate(s)
}

type CreateCloudResourceRequestRedirect struct {
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
	// The timeout period of idle persistent connections. Valid values: 10 to 3600. Default value: 3600. Unit: seconds.
	//
	// >  This parameter specifies the period of time after which an idle persistent connection is closed.
	//
	// example:
	//
	// 15
	KeepaliveTimeout *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The timeout period of read connections. Unit: seconds. Valid values: 1 to 3600.
	//
	// example:
	//
	// 1
	ReadTimeout *int32 `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	// The custom header fields. Specify the value in the [**{"k":"*key*","v":"*value*"}**] format. ***key**	- specifies the key of a custom header field. ***value**	- specifies the value of a custom header field.
	//
	// >  If a request contains a custom header field, WAF overwrites the original value of the field with the specified value.
	RequestHeaders []*CreateCloudResourceRequestRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// The timeout period of write connections. Unit: seconds. Valid values: 1 to 3600.
	//
	// example:
	//
	// 1
	WriteTimeout *int32 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
	// The method that is used to obtain the originating IP address of a client. Valid values:
	//
	// 	- **0**: No Layer 7 proxies are deployed in front of WAF.
	//
	// 	- **1**: WAF reads the first value of the X-Forwarded-For (XFF) header field as the originating IP address of the client.
	//
	// 	- **2**: WAF reads the value of a custom header field as the originating IP address of the client.
	//
	// example:
	//
	// 1
	XffHeaderMode *int32 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// The custom header fields that are used to obtain the originating IP address of a client. Specify the value in the **["header1","header2",...]*	- format.
	//
	// >  This parameter is required only if you set **XffHeaderMode*	- to 2.
	XffHeaders []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
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

func (s CreateCloudResourceRequestRedirect) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudResourceRequestRedirect) GoString() string {
	return s.String()
}

func (s *CreateCloudResourceRequestRedirect) GetKeepalive() *bool {
	return s.Keepalive
}

func (s *CreateCloudResourceRequestRedirect) GetKeepaliveRequests() *int32 {
	return s.KeepaliveRequests
}

func (s *CreateCloudResourceRequestRedirect) GetKeepaliveTimeout() *int32 {
	return s.KeepaliveTimeout
}

func (s *CreateCloudResourceRequestRedirect) GetReadTimeout() *int32 {
	return s.ReadTimeout
}

func (s *CreateCloudResourceRequestRedirect) GetRequestHeaders() []*CreateCloudResourceRequestRedirectRequestHeaders {
	return s.RequestHeaders
}

func (s *CreateCloudResourceRequestRedirect) GetWriteTimeout() *int32 {
	return s.WriteTimeout
}

func (s *CreateCloudResourceRequestRedirect) GetXffHeaderMode() *int32 {
	return s.XffHeaderMode
}

func (s *CreateCloudResourceRequestRedirect) GetXffHeaders() []*string {
	return s.XffHeaders
}

func (s *CreateCloudResourceRequestRedirect) GetXffProto() *bool {
	return s.XffProto
}

func (s *CreateCloudResourceRequestRedirect) SetKeepalive(v bool) *CreateCloudResourceRequestRedirect {
	s.Keepalive = &v
	return s
}

func (s *CreateCloudResourceRequestRedirect) SetKeepaliveRequests(v int32) *CreateCloudResourceRequestRedirect {
	s.KeepaliveRequests = &v
	return s
}

func (s *CreateCloudResourceRequestRedirect) SetKeepaliveTimeout(v int32) *CreateCloudResourceRequestRedirect {
	s.KeepaliveTimeout = &v
	return s
}

func (s *CreateCloudResourceRequestRedirect) SetReadTimeout(v int32) *CreateCloudResourceRequestRedirect {
	s.ReadTimeout = &v
	return s
}

func (s *CreateCloudResourceRequestRedirect) SetRequestHeaders(v []*CreateCloudResourceRequestRedirectRequestHeaders) *CreateCloudResourceRequestRedirect {
	s.RequestHeaders = v
	return s
}

func (s *CreateCloudResourceRequestRedirect) SetWriteTimeout(v int32) *CreateCloudResourceRequestRedirect {
	s.WriteTimeout = &v
	return s
}

func (s *CreateCloudResourceRequestRedirect) SetXffHeaderMode(v int32) *CreateCloudResourceRequestRedirect {
	s.XffHeaderMode = &v
	return s
}

func (s *CreateCloudResourceRequestRedirect) SetXffHeaders(v []*string) *CreateCloudResourceRequestRedirect {
	s.XffHeaders = v
	return s
}

func (s *CreateCloudResourceRequestRedirect) SetXffProto(v bool) *CreateCloudResourceRequestRedirect {
	s.XffProto = &v
	return s
}

func (s *CreateCloudResourceRequestRedirect) Validate() error {
	return dara.Validate(s)
}

type CreateCloudResourceRequestRedirectRequestHeaders struct {
	// The key of the custom header field.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom header field.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCloudResourceRequestRedirectRequestHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudResourceRequestRedirectRequestHeaders) GoString() string {
	return s.String()
}

func (s *CreateCloudResourceRequestRedirectRequestHeaders) GetKey() *string {
	return s.Key
}

func (s *CreateCloudResourceRequestRedirectRequestHeaders) GetValue() *string {
	return s.Value
}

func (s *CreateCloudResourceRequestRedirectRequestHeaders) SetKey(v string) *CreateCloudResourceRequestRedirectRequestHeaders {
	s.Key = &v
	return s
}

func (s *CreateCloudResourceRequestRedirectRequestHeaders) SetValue(v string) *CreateCloudResourceRequestRedirectRequestHeaders {
	s.Value = &v
	return s
}

func (s *CreateCloudResourceRequestRedirectRequestHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateCloudResourceRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// TagKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// TagValue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCloudResourceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudResourceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCloudResourceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCloudResourceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCloudResourceRequestTag) SetKey(v string) *CreateCloudResourceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCloudResourceRequestTag) SetValue(v string) *CreateCloudResourceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCloudResourceRequestTag) Validate() error {
	return dara.Validate(s)
}
