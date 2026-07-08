// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDomainCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *ModifyDomainCertRequest
	GetCertId() *string
	SetCipherSuite(v string) *ModifyDomainCertRequest
	GetCipherSuite() *string
	SetCustomCiphers(v []*string) *ModifyDomainCertRequest
	GetCustomCiphers() []*string
	SetDomain(v string) *ModifyDomainCertRequest
	GetDomain() *string
	SetEnableTLSv3(v bool) *ModifyDomainCertRequest
	GetEnableTLSv3() *bool
	SetInstanceId(v string) *ModifyDomainCertRequest
	GetInstanceId() *string
	SetRegionId(v string) *ModifyDomainCertRequest
	GetRegionId() *string
	SetTLSVersion(v string) *ModifyDomainCertRequest
	GetTLSVersion() *string
}

type ModifyDomainCertRequest struct {
	// The ID of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123-cn-hangzhou
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The type of the cipher suite. Valid values:
	//
	// - **1**: all cipher suites.
	//
	// - **2**: strong cipher suites.
	//
	// - **99**: custom cipher suites.
	//
	// example:
	//
	// 1
	CipherSuite *string `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	// The custom cipher suites. This parameter is available only when you set **CipherSuite*	- to **99**.
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// The domain name that is added to WAF in CNAME record mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether to enable TLS 1.3. Valid values:
	//
	// - **true**: TLS 1.3 is enabled.
	//
	// - **false**: TLS 1.3 is disabled.
	//
	// example:
	//
	// false
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// The ID of the WAF instance.
	//
	// > Call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-ww**b06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	// The Transport Layer Security (TLS) version. Valid values:
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
}

func (s ModifyDomainCertRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDomainCertRequest) GoString() string {
	return s.String()
}

func (s *ModifyDomainCertRequest) GetCertId() *string {
	return s.CertId
}

func (s *ModifyDomainCertRequest) GetCipherSuite() *string {
	return s.CipherSuite
}

func (s *ModifyDomainCertRequest) GetCustomCiphers() []*string {
	return s.CustomCiphers
}

func (s *ModifyDomainCertRequest) GetDomain() *string {
	return s.Domain
}

func (s *ModifyDomainCertRequest) GetEnableTLSv3() *bool {
	return s.EnableTLSv3
}

func (s *ModifyDomainCertRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDomainCertRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDomainCertRequest) GetTLSVersion() *string {
	return s.TLSVersion
}

func (s *ModifyDomainCertRequest) SetCertId(v string) *ModifyDomainCertRequest {
	s.CertId = &v
	return s
}

func (s *ModifyDomainCertRequest) SetCipherSuite(v string) *ModifyDomainCertRequest {
	s.CipherSuite = &v
	return s
}

func (s *ModifyDomainCertRequest) SetCustomCiphers(v []*string) *ModifyDomainCertRequest {
	s.CustomCiphers = v
	return s
}

func (s *ModifyDomainCertRequest) SetDomain(v string) *ModifyDomainCertRequest {
	s.Domain = &v
	return s
}

func (s *ModifyDomainCertRequest) SetEnableTLSv3(v bool) *ModifyDomainCertRequest {
	s.EnableTLSv3 = &v
	return s
}

func (s *ModifyDomainCertRequest) SetInstanceId(v string) *ModifyDomainCertRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDomainCertRequest) SetRegionId(v string) *ModifyDomainCertRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDomainCertRequest) SetTLSVersion(v string) *ModifyDomainCertRequest {
	s.TLSVersion = &v
	return s
}

func (s *ModifyDomainCertRequest) Validate() error {
	return dara.Validate(s)
}
