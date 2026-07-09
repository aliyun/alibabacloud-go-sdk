// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceCustomizedDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *CreateInstanceCustomizedDomainRequest
	GetCertId() *string
	SetCertRegionId(v string) *CreateInstanceCustomizedDomainRequest
	GetCertRegionId() *string
	SetDomain(v string) *CreateInstanceCustomizedDomainRequest
	GetDomain() *string
	SetInstanceId(v string) *CreateInstanceCustomizedDomainRequest
	GetInstanceId() *string
	SetModuleName(v string) *CreateInstanceCustomizedDomainRequest
	GetModuleName() *string
}

type CreateInstanceCustomizedDomainRequest struct {
	// The ID of the custom domain name certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21428921
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The region to which the certificate belongs.
	//
	// example:
	//
	// cn-hangzhou
	CertRegionId *string `json:"CertRegionId,omitempty" xml:"CertRegionId,omitempty"`
	// The custom domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The custom module name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Registry
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s CreateInstanceCustomizedDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceCustomizedDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceCustomizedDomainRequest) GetCertId() *string {
	return s.CertId
}

func (s *CreateInstanceCustomizedDomainRequest) GetCertRegionId() *string {
	return s.CertRegionId
}

func (s *CreateInstanceCustomizedDomainRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateInstanceCustomizedDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceCustomizedDomainRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *CreateInstanceCustomizedDomainRequest) SetCertId(v string) *CreateInstanceCustomizedDomainRequest {
	s.CertId = &v
	return s
}

func (s *CreateInstanceCustomizedDomainRequest) SetCertRegionId(v string) *CreateInstanceCustomizedDomainRequest {
	s.CertRegionId = &v
	return s
}

func (s *CreateInstanceCustomizedDomainRequest) SetDomain(v string) *CreateInstanceCustomizedDomainRequest {
	s.Domain = &v
	return s
}

func (s *CreateInstanceCustomizedDomainRequest) SetInstanceId(v string) *CreateInstanceCustomizedDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceCustomizedDomainRequest) SetModuleName(v string) *CreateInstanceCustomizedDomainRequest {
	s.ModuleName = &v
	return s
}

func (s *CreateInstanceCustomizedDomainRequest) Validate() error {
	return dara.Validate(s)
}
