// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer7CertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCert(v string) *ConfigLayer7CertRequest
	GetCert() *string
	SetCertId(v int32) *ConfigLayer7CertRequest
	GetCertId() *int32
	SetCertIdentifier(v string) *ConfigLayer7CertRequest
	GetCertIdentifier() *string
	SetCertName(v string) *ConfigLayer7CertRequest
	GetCertName() *string
	SetCertRegion(v string) *ConfigLayer7CertRequest
	GetCertRegion() *string
	SetDomain(v string) *ConfigLayer7CertRequest
	GetDomain() *string
	SetKey(v string) *ConfigLayer7CertRequest
	GetKey() *string
	SetResourceGroupId(v string) *ConfigLayer7CertRequest
	GetResourceGroupId() *string
}

type ConfigLayer7CertRequest struct {
	// example:
	//
	// xx
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// example:
	//
	// 1234
	CertId         *int32  `json:"CertId,omitempty" xml:"CertId,omitempty"`
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// example:
	//
	// testCertName
	CertName   *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	CertRegion *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.aliyun.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// xx
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ConfigLayer7CertRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer7CertRequest) GoString() string {
	return s.String()
}

func (s *ConfigLayer7CertRequest) GetCert() *string {
	return s.Cert
}

func (s *ConfigLayer7CertRequest) GetCertId() *int32 {
	return s.CertId
}

func (s *ConfigLayer7CertRequest) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *ConfigLayer7CertRequest) GetCertName() *string {
	return s.CertName
}

func (s *ConfigLayer7CertRequest) GetCertRegion() *string {
	return s.CertRegion
}

func (s *ConfigLayer7CertRequest) GetDomain() *string {
	return s.Domain
}

func (s *ConfigLayer7CertRequest) GetKey() *string {
	return s.Key
}

func (s *ConfigLayer7CertRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ConfigLayer7CertRequest) SetCert(v string) *ConfigLayer7CertRequest {
	s.Cert = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetCertId(v int32) *ConfigLayer7CertRequest {
	s.CertId = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetCertIdentifier(v string) *ConfigLayer7CertRequest {
	s.CertIdentifier = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetCertName(v string) *ConfigLayer7CertRequest {
	s.CertName = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetCertRegion(v string) *ConfigLayer7CertRequest {
	s.CertRegion = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetDomain(v string) *ConfigLayer7CertRequest {
	s.Domain = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetKey(v string) *ConfigLayer7CertRequest {
	s.Key = &v
	return s
}

func (s *ConfigLayer7CertRequest) SetResourceGroupId(v string) *ConfigLayer7CertRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ConfigLayer7CertRequest) Validate() error {
	return dara.Validate(s)
}
