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
	// > You can call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The listening configuration.
	//
	// This parameter is required.
	Listen *CreateCloudResourceRequestListen `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	// The UID of the current resource ownership.
	//
	// example:
	//
	// 123
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The forwarding configuration.
	Redirect *CreateCloudResourceRequestRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
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
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The list of tags. A maximum of 20 tags can be specified.
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

type CreateCloudResourceRequestListen struct {
	// The list of certificate IDs.
	Certificates []*CreateCloudResourceRequestListenCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The type of cipher suite to add. This parameter is used only when **HttpsPorts*	- is not empty, which indicates that the domain name uses the HTTPS protocol. Valid values:
	//
	// - **1**: all cipher suites.
	//
	// - **2**: strong cipher suites. This value is available only when **TLSVersion*	- is set to **tlsv1.2**.
	//
	// - **99**: custom cipher suites. This value is available only when TLSVersion is not set to tlsv1.3.
	//
	// example:
	//
	// 1
	CipherSuite *int32 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The specific custom cipher suites to add. This parameter is used only when **CipherSuite*	- is set to **99**.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// The domain name connected to WAF.
	//
	// > This parameter is required only when the cloud service type is ddos. For other service types, leave this field empty.
	//
	// example:
	//
	// www.c**sw.net
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether TLS 1.3 is supported. Valid values:
	//
	// - **true**: TLS 1.3 is supported.
	//
	// - **false**: TLS 1.3 is not supported.
	//
	// > This parameter is used only when HttpsPorts is not empty, which indicates that the domain name uses the HTTPS protocol. When TLSVersion is set to tlsv1.3, this value must be true.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// Specifies whether to enable HTTP/2. This parameter is used only when **HttpsPorts*	- is not empty, which indicates that the domain name uses the HTTPS protocol. Valid values:
	//
	// - **true**: HTTP/2 is enabled.
	//
	// - **false*	- (default): HTTP/2 is not enabled.
	//
	// example:
	//
	// true
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// The port of the cloud service connected to WAF.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type.
	//
	// example:
	//
	// http
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The ID of the cloud service instance.
	//
	// example:
	//
	// lb-bp1*****
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The cloud service type.
	//
	// This parameter is required.
	//
	// example:
	//
	// clb4
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
	// The region ID of the cloud service.
	//
	//
	//
	// > This parameter is required when the instance ID to be connected has not been synchronized to WAF.
	//
	// example:
	//
	// cn-hangzhou
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// The TLS version to add. This parameter is used only when **HttpsPorts*	- is not empty, which indicates that the domain name uses the HTTPS protocol. Valid values:
	//
	// - **tlsv1**: Supports TLS 1.0 and later. Highest compatibility and lowest security.
	//
	// - **tlsv1.1**: Supports TLS 1.1 and later. Good compatibility and good security.
	//
	// - **tlsv1.2**: Supports TLS 1.2 and later. Good compatibility and highest security.
	//
	// - **tlsv1.3**: Supports only TLS 1.3. Highest security and lowest compatibility.
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

func (s *CreateCloudResourceRequestListen) GetDomain() *string {
	return s.Domain
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

func (s *CreateCloudResourceRequestListen) GetResourceRegionId() *string {
	return s.ResourceRegionId
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

func (s *CreateCloudResourceRequestListen) SetDomain(v string) *CreateCloudResourceRequestListen {
	s.Domain = &v
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

func (s *CreateCloudResourceRequestListen) SetResourceRegionId(v string) *CreateCloudResourceRequestListen {
	s.ResourceRegionId = &v
	return s
}

func (s *CreateCloudResourceRequestListen) SetTLSVersion(v string) *CreateCloudResourceRequestListen {
	s.TLSVersion = &v
	return s
}

func (s *CreateCloudResourceRequestListen) Validate() error {
	if s.Certificates != nil {
		for _, item := range s.Certificates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCloudResourceRequestListenCertificates struct {
	// The certificate type for the HTTPS protocol. Valid values:
	//
	// - **default**: default certificate.
	//
	// - **extension**: extension certificate.
	//
	// example:
	//
	// default
	AppliedType *string `json:"AppliedType,omitempty" xml:"AppliedType,omitempty"`
	// The ID of the certificate to add.
	//
	// > You can call [DescribeResourceInstanceCerts](https://help.aliyun.com/document_detail/2718120.html) to query the IDs of all SSL certificates associated with the cloud service instance.
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
	// The number of requests that reuse persistent connections. Valid values: 60 to 1000.
	//
	// > After persistent connections are enabled, this parameter specifies how many requests can reuse persistent connections.
	//
	// example:
	//
	// 1000
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The timeout period of idle persistent connections. Valid values: 10 to 3600. Default value: 3600. Unit: seconds.
	//
	// > Specifies how long an idle persistent connection can remain open before it is released.
	//
	// example:
	//
	// 3600
	KeepaliveTimeout *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The maximum request body size. Valid values: 2 to 10. Default value: 2. Unit: GB.
	//
	// > Only Ultimate Edition supports this parameter.
	//
	// example:
	//
	// 2
	MaxBodySize *int32 `json:"MaxBodySize,omitempty" xml:"MaxBodySize,omitempty"`
	// The read timeout period. Unit: seconds.
	//
	// Valid values: 1 to 3600.
	//
	// example:
	//
	// 1
	ReadTimeout *int32 `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	// The value of this parameter is in the format of [**{"k":"_key_","v":"_value_"}**], where **_key_*	- specifies the custom request header field and **_value_*	- specifies the value set for the field.
	//
	// > If the custom header field already exists in the request, the system overwrites the value of the custom field in the request with the specified traffic tag value.
	RequestHeaders []*CreateCloudResourceRequestRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// The write timeout period. Unit: seconds.
	//
	// Valid values: 1 to 3600.
	//
	// example:
	//
	// 1
	WriteTimeout *int32 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
	// The method that WAF uses to obtain the originating IP address of the client. Valid values:
	//
	// - **0**: No Layer 7 proxy is deployed in front of WAF.
	//
	// - **1**: WAF reads the first value of the X-Forwarded-For (XFF) header field as the client IP address.
	//
	// - **2**: WAF reads the value of a custom header field that you specify as the client IP address.
	//
	// example:
	//
	// 1
	XffHeaderMode *int32 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// The list of custom header fields used to obtain the client IP address, in the format of [**"header1","header2",……**].
	//
	// > This parameter is required only when **XffHeaderMode*	- is set to 2, which indicates that WAF reads the value of a custom header field that you specify as the client IP address.
	XffHeaders []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
	// Specifies whether to use X-Forward-For-Proto to pass the protocol of WAF. Valid values:
	//
	// - **true*	- (default): The protocol of WAF is passed.
	//
	// - **false**: The protocol of WAF is not passed.
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

func (s *CreateCloudResourceRequestRedirect) GetMaxBodySize() *int32 {
	return s.MaxBodySize
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

func (s *CreateCloudResourceRequestRedirect) SetMaxBodySize(v int32) *CreateCloudResourceRequestRedirect {
	s.MaxBodySize = &v
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

type CreateCloudResourceRequestRedirectRequestHeaders struct {
	// The specified custom request header field.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value set for the custom request header field.
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
	// The tag key.
	//
	// example:
	//
	// TagKey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
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
