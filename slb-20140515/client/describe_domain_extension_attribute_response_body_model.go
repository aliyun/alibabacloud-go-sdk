// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainExtensionAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DescribeDomainExtensionAttributeResponseBody
	GetDomain() *string
	SetDomainExtensionId(v string) *DescribeDomainExtensionAttributeResponseBody
	GetDomainExtensionId() *string
	SetListenerPort(v int32) *DescribeDomainExtensionAttributeResponseBody
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *DescribeDomainExtensionAttributeResponseBody
	GetLoadBalancerId() *string
	SetRequestId(v string) *DescribeDomainExtensionAttributeResponseBody
	GetRequestId() *string
	SetServerCertificateId(v string) *DescribeDomainExtensionAttributeResponseBody
	GetServerCertificateId() *string
}

type DescribeDomainExtensionAttributeResponseBody struct {
	// The domain name.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the additional certificate.
	//
	// example:
	//
	// de-bp1rp7ta1****
	DomainExtensionId *string `json:"DomainExtensionId,omitempty" xml:"DomainExtensionId,omitempty"`
	// The frontend port of the HTTPS listener that is configured for the SLB instance. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 443
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the SLB instance.
	//
	// example:
	//
	// lb-bp1o94dp5i6*****earr9g6d1l
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 48C1B671-C6DB-4DDE-9B30-10557E36CDE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the server certificate used by the domain name.
	//
	// example:
	//
	// 231579085529123_166f82******_1714763408_709981430
	ServerCertificateId *string `json:"ServerCertificateId,omitempty" xml:"ServerCertificateId,omitempty"`
}

func (s DescribeDomainExtensionAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainExtensionAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainExtensionAttributeResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDomainExtensionAttributeResponseBody) GetDomainExtensionId() *string {
	return s.DomainExtensionId
}

func (s *DescribeDomainExtensionAttributeResponseBody) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeDomainExtensionAttributeResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeDomainExtensionAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainExtensionAttributeResponseBody) GetServerCertificateId() *string {
	return s.ServerCertificateId
}

func (s *DescribeDomainExtensionAttributeResponseBody) SetDomain(v string) *DescribeDomainExtensionAttributeResponseBody {
	s.Domain = &v
	return s
}

func (s *DescribeDomainExtensionAttributeResponseBody) SetDomainExtensionId(v string) *DescribeDomainExtensionAttributeResponseBody {
	s.DomainExtensionId = &v
	return s
}

func (s *DescribeDomainExtensionAttributeResponseBody) SetListenerPort(v int32) *DescribeDomainExtensionAttributeResponseBody {
	s.ListenerPort = &v
	return s
}

func (s *DescribeDomainExtensionAttributeResponseBody) SetLoadBalancerId(v string) *DescribeDomainExtensionAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeDomainExtensionAttributeResponseBody) SetRequestId(v string) *DescribeDomainExtensionAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainExtensionAttributeResponseBody) SetServerCertificateId(v string) *DescribeDomainExtensionAttributeResponseBody {
	s.ServerCertificateId = &v
	return s
}

func (s *DescribeDomainExtensionAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
