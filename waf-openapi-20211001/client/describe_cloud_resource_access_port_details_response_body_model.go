// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudResourceAccessPortDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPortDetails(v []*DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) *DescribeCloudResourceAccessPortDetailsResponseBody
	GetAccessPortDetails() []*DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails
	SetRequestId(v string) *DescribeCloudResourceAccessPortDetailsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCloudResourceAccessPortDetailsResponseBody
	GetTotalCount() *int32
}

type DescribeCloudResourceAccessPortDetailsResponseBody struct {
	// The details of the ports of cloud services that are added to WAF.
	AccessPortDetails []*DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails `json:"AccessPortDetails,omitempty" xml:"AccessPortDetails,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 2EFCFE18-78F8-5079-B312-07***48B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudResourceAccessPortDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceAccessPortDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBody) GetAccessPortDetails() []*DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	return s.AccessPortDetails
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBody) SetAccessPortDetails(v []*DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) *DescribeCloudResourceAccessPortDetailsResponseBody {
	s.AccessPortDetails = v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBody) SetRequestId(v string) *DescribeCloudResourceAccessPortDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBody) SetTotalCount(v int32) *DescribeCloudResourceAccessPortDetailsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBody) Validate() error {
	if s.AccessPortDetails != nil {
		for _, item := range s.AccessPortDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails struct {
	// The certificates that are associated with the ports of cloud services.
	Certificates []*DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates `json:"Certificates,omitempty" xml:"Certificates,omitempty" type:"Repeated"`
	// The type of the cipher suites. Valid values:
	//
	// 	- **1**: all cipher suites.
	//
	// 	- **2**: strong cipher suites.
	//
	// 	- **99**: custom cipher suites.
	//
	// example:
	//
	// 1
	CipherSuite *int32 `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suites that you want to add. This parameter is available only if you set **CipherSuite*	- to **99**.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// Indicates whether to support TLS 1.3. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// Indicates whether to enable HTTP/2. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Http2Enabled *bool `json:"Http2Enabled,omitempty" xml:"Http2Enabled,omitempty"`
	// Indicates whether to enable the persistent connection feature. Valid values:
	//
	// 	- **true*	- (default)
	//
	// 	- **false:**
	//
	// example:
	//
	// true
	Keepalive *bool `json:"Keepalive,omitempty" xml:"Keepalive,omitempty"`
	// The number of reused persistent connections. Valid values: 60 to 1000.
	//
	// >  This parameter specifies the number of requests that reuse persistent connections after you enable the persistent connection feature.
	//
	// example:
	//
	// 1000
	KeepaliveRequests *int32 `json:"KeepaliveRequests,omitempty" xml:"KeepaliveRequests,omitempty"`
	// The timeout period for idle persistent connections. Valid values: 10 to 3600. Default value: 15. Unit: seconds.
	//
	// >  If no new requests are initiated over the idle persistent connection within the specified timeout period, the connection is closed.
	//
	// example:
	//
	// 10
	KeepaliveTimeout *int32 `json:"KeepaliveTimeout,omitempty" xml:"KeepaliveTimeout,omitempty"`
	// The custom header field that you want to use to label requests that are processed by WAF.
	//
	// >  This parameter is returned only when the traffic marking feature is enabled for the domain name.
	LogHeaders []*DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsLogHeaders `json:"LogHeaders,omitempty" xml:"LogHeaders,omitempty" type:"Repeated"`
	// The ID of the Alibaba Cloud account to which the resource belongs.
	//
	// example:
	//
	// 123
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The port of the cloud service that is added to WAF.
	//
	// example:
	//
	// 443
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The type of the protocol. Valid values:
	//
	// 	- **http**
	//
	// 	- **https**
	//
	// example:
	//
	// https
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The timeout period for read connections. Unit: seconds. Valid values: 1 to 3600.
	//
	// example:
	//
	// 5
	ReadTimeout *int32 `json:"ReadTimeout,omitempty" xml:"ReadTimeout,omitempty"`
	// The status of the domain name. Valid values:
	//
	// 	- **1**: indicates that the port is available.
	//
	// 	- **2**: indicates that the port is being created.
	//
	// 	- **3**: indicates that the port is being modified.
	//
	// 	- **4**: indicates that the port is being released.
	//
	// example:
	//
	// 1
	Status           *int32                                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	SubStatus        *string                                                                                `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	SubStatusDetails []*DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails `json:"SubStatusDetails,omitempty" xml:"SubStatusDetails,omitempty" type:"Repeated"`
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
	// tlsv1
	TLSVersion *string `json:"TLSVersion,omitempty" xml:"TLSVersion,omitempty"`
	// The timeout period for write connections. Unit: seconds. Valid values: 1 to 3600.
	//
	// example:
	//
	// 1
	WriteTimeout *int32 `json:"WriteTimeout,omitempty" xml:"WriteTimeout,omitempty"`
	// The method that WAF uses to obtain the originating IP address of a client. Valid values:
	//
	// 	- **0**: No Layer 7 proxies are deployed in front of WAF.
	//
	// 	- **1**: WAF reads the first value of the X-Forwarded-For (XFF) header field as the originating IP address of the client.
	//
	// 	- **2**: WAF reads the value of a custom header field as the originating IP address of the client.
	//
	// example:
	//
	// 0
	XffHeaderMode *int32 `json:"XffHeaderMode,omitempty" xml:"XffHeaderMode,omitempty"`
	// The custom header field that is used to obtain the originating IP address of a client. Specify the value in the ["header1","header2",...] format.
	//
	// >  This parameter is required only if you set **XffHeaderMode*	- to 2.
	XffHeaders []*string `json:"XffHeaders,omitempty" xml:"XffHeaders,omitempty" type:"Repeated"`
	// Indicates whether to use the X-Forward-For-Proto header to identify the protocol used by WAF to forward requests to the origin server. Valid values:
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

func (s DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetCertificates() []*DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates {
	return s.Certificates
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetCipherSuite() *int32 {
	return s.CipherSuite
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetCustomCiphers() []*string {
	return s.CustomCiphers
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetEnableTLSv3() *bool {
	return s.EnableTLSv3
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetHttp2Enabled() *bool {
	return s.Http2Enabled
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetKeepalive() *bool {
	return s.Keepalive
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetKeepaliveRequests() *int32 {
	return s.KeepaliveRequests
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetKeepaliveTimeout() *int32 {
	return s.KeepaliveTimeout
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetLogHeaders() []*DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsLogHeaders {
	return s.LogHeaders
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetPort() *int32 {
	return s.Port
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetReadTimeout() *int32 {
	return s.ReadTimeout
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetSubStatus() *string {
	return s.SubStatus
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetSubStatusDetails() []*DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails {
	return s.SubStatusDetails
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetTLSVersion() *string {
	return s.TLSVersion
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetWriteTimeout() *int32 {
	return s.WriteTimeout
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetXffHeaderMode() *int32 {
	return s.XffHeaderMode
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetXffHeaders() []*string {
	return s.XffHeaders
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) GetXffProto() *bool {
	return s.XffProto
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetCertificates(v []*DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.Certificates = v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetCipherSuite(v int32) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.CipherSuite = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetCustomCiphers(v []*string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.CustomCiphers = v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetEnableTLSv3(v bool) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.EnableTLSv3 = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetHttp2Enabled(v bool) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.Http2Enabled = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetKeepalive(v bool) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.Keepalive = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetKeepaliveRequests(v int32) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.KeepaliveRequests = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetKeepaliveTimeout(v int32) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.KeepaliveTimeout = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetLogHeaders(v []*DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsLogHeaders) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.LogHeaders = v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetOwnerUserId(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetPort(v int32) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.Port = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetProtocol(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.Protocol = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetReadTimeout(v int32) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.ReadTimeout = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetStatus(v int32) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.Status = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetSubStatus(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.SubStatus = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetSubStatusDetails(v []*DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.SubStatusDetails = v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetTLSVersion(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.TLSVersion = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetWriteTimeout(v int32) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.WriteTimeout = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetXffHeaderMode(v int32) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.XffHeaderMode = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetXffHeaders(v []*string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.XffHeaders = v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) SetXffProto(v bool) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails {
	s.XffProto = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetails) Validate() error {
	if s.Certificates != nil {
		for _, item := range s.Certificates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.LogHeaders != nil {
		for _, item := range s.LogHeaders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SubStatusDetails != nil {
		for _, item := range s.SubStatusDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates struct {
	// The type of the HTTPS certificate. Valid values:
	//
	// 	- **default**: default certificate.
	//
	// 	- **extension**: additional certificate.
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
	// The name of the certificate.
	//
	// example:
	//
	// cert-name1
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
}

func (s DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates) GetAppliedType() *string {
	return s.AppliedType
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates) GetCertificateId() *string {
	return s.CertificateId
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates) GetCertificateName() *string {
	return s.CertificateName
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates) SetAppliedType(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates {
	s.AppliedType = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates) SetCertificateId(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates {
	s.CertificateId = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates) SetCertificateName(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates {
	s.CertificateName = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsCertificates) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsLogHeaders struct {
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

func (s DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsLogHeaders) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsLogHeaders) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsLogHeaders) GetKey() *string {
	return s.Key
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsLogHeaders) GetValue() *string {
	return s.Value
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsLogHeaders) SetKey(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsLogHeaders {
	s.Key = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsLogHeaders) SetValue(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsLogHeaders {
	s.Value = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsLogHeaders) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails struct {
	AppliedType     *string `json:"AppliedType,omitempty" xml:"AppliedType,omitempty"`
	CertId          *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertName        *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CommonName      *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	Domain          *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	ExpireTime      *int64  `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ProductCertId   *string `json:"ProductCertId,omitempty" xml:"ProductCertId,omitempty"`
	ProductCertName *string `json:"ProductCertName,omitempty" xml:"ProductCertName,omitempty"`
	ReasonCode      *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
}

func (s DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) GoString() string {
	return s.String()
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) GetAppliedType() *string {
	return s.AppliedType
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) GetCertId() *string {
	return s.CertId
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) GetCertName() *string {
	return s.CertName
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) GetDomain() *string {
	return s.Domain
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) GetProductCertId() *string {
	return s.ProductCertId
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) GetProductCertName() *string {
	return s.ProductCertName
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) SetAppliedType(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails {
	s.AppliedType = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) SetCertId(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails {
	s.CertId = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) SetCertName(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails {
	s.CertName = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) SetCommonName(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails {
	s.CommonName = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) SetDomain(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails {
	s.Domain = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) SetExpireTime(v int64) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails {
	s.ExpireTime = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) SetProductCertId(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails {
	s.ProductCertId = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) SetProductCertName(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails {
	s.ProductCertName = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) SetReasonCode(v string) *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails {
	s.ReasonCode = &v
	return s
}

func (s *DescribeCloudResourceAccessPortDetailsResponseBodyAccessPortDetailsSubStatusDetails) Validate() error {
	return dara.Validate(s)
}
