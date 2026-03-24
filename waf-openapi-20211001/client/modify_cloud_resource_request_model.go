// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudResourceId(v string) *ModifyCloudResourceRequest
	GetCloudResourceId() *string
	SetInstanceId(v string) *ModifyCloudResourceRequest
	GetInstanceId() *string
	SetListen(v *ModifyCloudResourceRequestListen) *ModifyCloudResourceRequest
	GetListen() *ModifyCloudResourceRequestListen
	SetRedirect(v *ModifyCloudResourceRequestRedirect) *ModifyCloudResourceRequest
	GetRedirect() *ModifyCloudResourceRequestRedirect
	SetRegionId(v string) *ModifyCloudResourceRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *ModifyCloudResourceRequest
	GetResourceManagerResourceGroupId() *string
}

type ModifyCloudResourceRequest struct {
	// The ID of the cloud resource that is added to WAF.
	//
	// > Call [CreateCloudResource](https://help.aliyun.com/document_detail/2839876.html) to add a cloud resource. The resource ID is included in the response.
	//
	// example:
	//
	// lb-***-80-clb7
	CloudResourceId *string `json:"CloudResourceId,omitempty" xml:"CloudResourceId,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The listener configuration.
	//
	// This parameter is required.
	Listen *ModifyCloudResourceRequestListen `json:"Listen,omitempty" xml:"Listen,omitempty" type:"Struct"`
	// The forwarding configuration.
	Redirect *ModifyCloudResourceRequestRedirect `json:"Redirect,omitempty" xml:"Redirect,omitempty" type:"Struct"`
	// The region of the WAF instance. Valid values:
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
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s ModifyCloudResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudResourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyCloudResourceRequest) GetCloudResourceId() *string {
	return s.CloudResourceId
}

func (s *ModifyCloudResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyCloudResourceRequest) GetListen() *ModifyCloudResourceRequestListen {
	return s.Listen
}

func (s *ModifyCloudResourceRequest) GetRedirect() *ModifyCloudResourceRequestRedirect {
	return s.Redirect
}

func (s *ModifyCloudResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCloudResourceRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ModifyCloudResourceRequest) SetCloudResourceId(v string) *ModifyCloudResourceRequest {
	s.CloudResourceId = &v
	return s
}

func (s *ModifyCloudResourceRequest) SetInstanceId(v string) *ModifyCloudResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyCloudResourceRequest) SetListen(v *ModifyCloudResourceRequestListen) *ModifyCloudResourceRequest {
	s.Listen = v
	return s
}

func (s *ModifyCloudResourceRequest) SetRedirect(v *ModifyCloudResourceRequestRedirect) *ModifyCloudResourceRequest {
	s.Redirect = v
	return s
}

func (s *ModifyCloudResourceRequest) SetRegionId(v string) *ModifyCloudResourceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCloudResourceRequest) SetResourceManagerResourceGroupId(v string) *ModifyCloudResourceRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ModifyCloudResourceRequest) Validate() error {
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

type ModifyCloudResourceRequestListen struct {
	// The certificate information.
	Certificates []*ModifyCloudResourceRequestListenCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The type of the cipher suite to add. This parameter applies only when you use the HTTPS protocol. Valid values:
	//
	// - **1**: adds all cipher suites.
	//
	// - **2**: adds strong cipher suites. This value is available only when **TLSVersion*	- is set to **tlsv1.2**.
	//
	// - **99**: adds custom cipher suites.
	//
	// example:
	//
	// 1
	CipherSuite *int32 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suites.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Indicates whether TLS 1.3 is supported. This parameter applies only when you use the HTTPS protocol. Valid values:
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
	// Indicates whether HTTP/2 is enabled. This parameter applies only when you use the HTTPS protocol. Valid values:
	//
	// - **true**: enables HTTP/2.
	//
	// - **false*	- (default): disables HTTP/2.
	//
	// example:
	//
	// true
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// Deprecated
	//
	// The listening port of the cloud service instance that is added to WAF.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The protocol type. Valid values:
	//
	// - **http**: HTTP.
	//
	// - **https**: HTTPS.
	//
	// This parameter is required.
	//
	// example:
	//
	// http
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Deprecated
	//
	// The ID of the cloud service instance.
	//
	// example:
	//
	// lb-***
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// Deprecated
	//
	// The type of the cloud service. Valid values:
	//
	// - **clb4**: Layer 4 Classic Load Balancer (CLB).
	//
	// - **clb7**: Layer 7 CLB.
	//
	// - **ecs**: Elastic Compute Service (ECS).
	//
	// - **nlb**: Network Load Balancer (NLB).
	//
	// example:
	//
	// clb7
	ResourceProduct *string `json:"ResourceProduct,omitempty" xml:"ResourceProduct,omitempty"`
	// The Transport Layer Security (TLS) version. This parameter applies only when you use the HTTPS protocol. Valid values:
	//
	// - **tlsv1**
	//
	// - **tlsv1.1**
	//
	// - **tlsv1.2**
	//
	// example:
	//
	// tlsv1.2
	TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
}

func (s ModifyCloudResourceRequestListen) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudResourceRequestListen) GoString() string {
	return s.String()
}

func (s *ModifyCloudResourceRequestListen) GetCertificates() []*ModifyCloudResourceRequestListenCertificates {
	return s.Certificates
}

func (s *ModifyCloudResourceRequestListen) GetCipherSuite() *int32 {
	return s.CipherSuite
}

func (s *ModifyCloudResourceRequestListen) GetCustomCiphers() []*string {
	return s.CustomCiphers
}

func (s *ModifyCloudResourceRequestListen) GetEnableTLSv3() *bool {
	return s.EnableTLSv3
}

func (s *ModifyCloudResourceRequestListen) GetHttp2Enabled() *bool {
	return s.Http2Enabled
}

func (s *ModifyCloudResourceRequestListen) GetPort() *int32 {
	return s.Port
}

func (s *ModifyCloudResourceRequestListen) GetProtocol() *string {
	return s.Protocol
}

func (s *ModifyCloudResourceRequestListen) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *ModifyCloudResourceRequestListen) GetResourceProduct() *string {
	return s.ResourceProduct
}

func (s *ModifyCloudResourceRequestListen) GetTLSVersion() *string {
	return s.TLSVersion
}

func (s *ModifyCloudResourceRequestListen) SetCertificates(v []*ModifyCloudResourceRequestListenCertificates) *ModifyCloudResourceRequestListen {
	s.Certificates = v
	return s
}

func (s *ModifyCloudResourceRequestListen) SetCipherSuite(v int32) *ModifyCloudResourceRequestListen {
	s.CipherSuite = &v
	return s
}

func (s *ModifyCloudResourceRequestListen) SetCustomCiphers(v []*string) *ModifyCloudResourceRequestListen {
	s.CustomCiphers = v
	return s
}

func (s *ModifyCloudResourceRequestListen) SetEnableTLSv3(v bool) *ModifyCloudResourceRequestListen {
	s.EnableTLSv3 = &v
	return s
}

func (s *ModifyCloudResourceRequestListen) SetHttp2Enabled(v bool) *ModifyCloudResourceRequestListen {
	s.Http2Enabled = &v
	return s
}

func (s *ModifyCloudResourceRequestListen) SetPort(v int32) *ModifyCloudResourceRequestListen {
	s.Port = &v
	return s
}

func (s *ModifyCloudResourceRequestListen) SetProtocol(v string) *ModifyCloudResourceRequestListen {
	s.Protocol = &v
	return s
}

func (s *ModifyCloudResourceRequestListen) SetResourceInstanceId(v string) *ModifyCloudResourceRequestListen {
	s.ResourceInstanceId = &v
	return s
}

func (s *ModifyCloudResourceRequestListen) SetResourceProduct(v string) *ModifyCloudResourceRequestListen {
	s.ResourceProduct = &v
	return s
}

func (s *ModifyCloudResourceRequestListen) SetTLSVersion(v string) *ModifyCloudResourceRequestListen {
	s.TLSVersion = &v
	return s
}

func (s *ModifyCloudResourceRequestListen) Validate() error {
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

type ModifyCloudResourceRequestListenCertificates struct {
	// The type of the certificate for the HTTPS protocol. Valid values:
	//
	// - **default**: a default certificate.
	//
	// - **extension**: an extension certificate.
	//
	// example:
	//
	// default
	AppliedType *string `json:"AppliedType,omitempty" xml:"AppliedType,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 123-cn-hangzhou
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s ModifyCloudResourceRequestListenCertificates) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudResourceRequestListenCertificates) GoString() string {
	return s.String()
}

func (s *ModifyCloudResourceRequestListenCertificates) GetAppliedType() *string {
	return s.AppliedType
}

func (s *ModifyCloudResourceRequestListenCertificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *ModifyCloudResourceRequestListenCertificates) SetAppliedType(v string) *ModifyCloudResourceRequestListenCertificates {
	s.AppliedType = &v
	return s
}

func (s *ModifyCloudResourceRequestListenCertificates) SetCertificateId(v string) *ModifyCloudResourceRequestListenCertificates {
	s.CertificateId = &v
	return s
}

func (s *ModifyCloudResourceRequestListenCertificates) Validate() error {
	return dara.Validate(s)
}

type ModifyCloudResourceRequestRedirect struct {
	// Indicates whether persistent connections are enabled. Valid values:
	//
	// - **true*	- (default): enables persistent connections.
	//
	// - **false**: disables persistent connections.
	//
	// example:
	//
	// true
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// The maximum number of requests that can be served through one persistent connection. Valid values: 60 to 1000.
	//
	// example:
	//
	// 1000
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The timeout period for an idle persistent connection. Valid values: 10 to 3600. Default value: 3600. Unit: seconds.
	//
	// example:
	//
	// 15
	KeepaliveTimeout *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The maximum size of a request body. Valid values: 2 to 10. Default value: 2. Unit: GB.
	//
	// example:
	//
	// 2
	MaxBodySize *int32 `json:"MaxBodySize,omitempty" xml:"MaxBodySize,omitempty"`
	// The read timeout period. Unit: seconds. Valid values: 1 to 3600.
	//
	// example:
	//
	// 1
	ReadTimeout *int32 `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	// The custom header fields used to mark traffic that is processed by WAF.
	RequestHeaders []*ModifyCloudResourceRequestRedirectRequestHeaders `json:"RequestHeaders,omitempty" xml:"RequestHeaders,omitempty" type:"Repeated"`
	// The write timeout period. Unit: seconds. Valid values: 1 to 3600.
	//
	// example:
	//
	// 1
	WriteTimeout *int32 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
	// The method that WAF uses to obtain the real IP address of a client. Valid values:
	//
	// - **0**: WAF obtains the real IP address of the client from the request. Use this value when no Layer 7 proxy resides before WAF.
	//
	// - **1**: WAF reads the first value of the X-Forwarded-For (XFF) header as the client IP address.
	//
	// - **2**: WAF reads the value of a custom header field as the client IP address.
	//
	// example:
	//
	// 0
	XffHeaderMode *int32 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// The custom header fields that are used to obtain the client IP address.
	//
	// > This parameter is required only when **XffHeaderMode*	- is set to **2**.
	XffHeaders []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
	// Indicates whether the X-Forwarded-Proto header is used to pass the protocol used by WAF. Valid values:
	//
	// - **true*	- (default): passes the protocol.
	//
	// - **false**: does not pass the protocol.
	//
	// example:
	//
	// true
	XffProto *bool `json:"XffProto,omitempty" xml:"XffProto,omitempty"`
}

func (s ModifyCloudResourceRequestRedirect) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudResourceRequestRedirect) GoString() string {
	return s.String()
}

func (s *ModifyCloudResourceRequestRedirect) GetKeepalive() *bool {
	return s.Keepalive
}

func (s *ModifyCloudResourceRequestRedirect) GetKeepaliveRequests() *int32 {
	return s.KeepaliveRequests
}

func (s *ModifyCloudResourceRequestRedirect) GetKeepaliveTimeout() *int32 {
	return s.KeepaliveTimeout
}

func (s *ModifyCloudResourceRequestRedirect) GetMaxBodySize() *int32 {
	return s.MaxBodySize
}

func (s *ModifyCloudResourceRequestRedirect) GetReadTimeout() *int32 {
	return s.ReadTimeout
}

func (s *ModifyCloudResourceRequestRedirect) GetRequestHeaders() []*ModifyCloudResourceRequestRedirectRequestHeaders {
	return s.RequestHeaders
}

func (s *ModifyCloudResourceRequestRedirect) GetWriteTimeout() *int32 {
	return s.WriteTimeout
}

func (s *ModifyCloudResourceRequestRedirect) GetXffHeaderMode() *int32 {
	return s.XffHeaderMode
}

func (s *ModifyCloudResourceRequestRedirect) GetXffHeaders() []*string {
	return s.XffHeaders
}

func (s *ModifyCloudResourceRequestRedirect) GetXffProto() *bool {
	return s.XffProto
}

func (s *ModifyCloudResourceRequestRedirect) SetKeepalive(v bool) *ModifyCloudResourceRequestRedirect {
	s.Keepalive = &v
	return s
}

func (s *ModifyCloudResourceRequestRedirect) SetKeepaliveRequests(v int32) *ModifyCloudResourceRequestRedirect {
	s.KeepaliveRequests = &v
	return s
}

func (s *ModifyCloudResourceRequestRedirect) SetKeepaliveTimeout(v int32) *ModifyCloudResourceRequestRedirect {
	s.KeepaliveTimeout = &v
	return s
}

func (s *ModifyCloudResourceRequestRedirect) SetMaxBodySize(v int32) *ModifyCloudResourceRequestRedirect {
	s.MaxBodySize = &v
	return s
}

func (s *ModifyCloudResourceRequestRedirect) SetReadTimeout(v int32) *ModifyCloudResourceRequestRedirect {
	s.ReadTimeout = &v
	return s
}

func (s *ModifyCloudResourceRequestRedirect) SetRequestHeaders(v []*ModifyCloudResourceRequestRedirectRequestHeaders) *ModifyCloudResourceRequestRedirect {
	s.RequestHeaders = v
	return s
}

func (s *ModifyCloudResourceRequestRedirect) SetWriteTimeout(v int32) *ModifyCloudResourceRequestRedirect {
	s.WriteTimeout = &v
	return s
}

func (s *ModifyCloudResourceRequestRedirect) SetXffHeaderMode(v int32) *ModifyCloudResourceRequestRedirect {
	s.XffHeaderMode = &v
	return s
}

func (s *ModifyCloudResourceRequestRedirect) SetXffHeaders(v []*string) *ModifyCloudResourceRequestRedirect {
	s.XffHeaders = v
	return s
}

func (s *ModifyCloudResourceRequestRedirect) SetXffProto(v bool) *ModifyCloudResourceRequestRedirect {
	s.XffProto = &v
	return s
}

func (s *ModifyCloudResourceRequestRedirect) Validate() error {
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

type ModifyCloudResourceRequestRedirectRequestHeaders struct {
	// The custom request header field.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the custom request header field.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyCloudResourceRequestRedirectRequestHeaders) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudResourceRequestRedirectRequestHeaders) GoString() string {
	return s.String()
}

func (s *ModifyCloudResourceRequestRedirectRequestHeaders) GetKey() *string {
	return s.Key
}

func (s *ModifyCloudResourceRequestRedirectRequestHeaders) GetValue() *string {
	return s.Value
}

func (s *ModifyCloudResourceRequestRedirectRequestHeaders) SetKey(v string) *ModifyCloudResourceRequestRedirectRequestHeaders {
	s.Key = &v
	return s
}

func (s *ModifyCloudResourceRequestRedirectRequestHeaders) SetValue(v string) *ModifyCloudResourceRequestRedirectRequestHeaders {
	s.Value = &v
	return s
}

func (s *ModifyCloudResourceRequestRedirectRequestHeaders) Validate() error {
	return dara.Validate(s)
}
