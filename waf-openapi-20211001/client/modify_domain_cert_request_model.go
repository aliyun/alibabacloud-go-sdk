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
	// This parameter is required.
	//
	// example:
	//
	// 123-cn-hangzhou
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// example:
	//
	// 1
	CipherSuite   *string   `json:"CipherSuite,omitempty" xml:"CipherSuite,omitempty"`
	CustomCiphers []*string `json:"CustomCiphers,omitempty" xml:"CustomCiphers,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// false
	EnableTLSv3 *bool `json:"EnableTLSv3,omitempty" xml:"EnableTLSv3,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-ww**b06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
