// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceCustomizedDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *UpdateInstanceCustomizedDomainRequest
	GetCertId() *string
	SetCertRegionId(v string) *UpdateInstanceCustomizedDomainRequest
	GetCertRegionId() *string
	SetDomain(v string) *UpdateInstanceCustomizedDomainRequest
	GetDomain() *string
	SetInstanceId(v string) *UpdateInstanceCustomizedDomainRequest
	GetInstanceId() *string
	SetModuleName(v string) *UpdateInstanceCustomizedDomainRequest
	GetModuleName() *string
}

type UpdateInstanceCustomizedDomainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 21118442
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// example:
	//
	// cn-hangzhou
	CertRegionId *string `json:"CertRegionId,omitempty" xml:"CertRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// registry-cn-guangzhou.ack.aliyuncs.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cri-h0tvspmnglam5jw5
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Registry
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s UpdateInstanceCustomizedDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceCustomizedDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceCustomizedDomainRequest) GetCertId() *string {
	return s.CertId
}

func (s *UpdateInstanceCustomizedDomainRequest) GetCertRegionId() *string {
	return s.CertRegionId
}

func (s *UpdateInstanceCustomizedDomainRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateInstanceCustomizedDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceCustomizedDomainRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *UpdateInstanceCustomizedDomainRequest) SetCertId(v string) *UpdateInstanceCustomizedDomainRequest {
	s.CertId = &v
	return s
}

func (s *UpdateInstanceCustomizedDomainRequest) SetCertRegionId(v string) *UpdateInstanceCustomizedDomainRequest {
	s.CertRegionId = &v
	return s
}

func (s *UpdateInstanceCustomizedDomainRequest) SetDomain(v string) *UpdateInstanceCustomizedDomainRequest {
	s.Domain = &v
	return s
}

func (s *UpdateInstanceCustomizedDomainRequest) SetInstanceId(v string) *UpdateInstanceCustomizedDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceCustomizedDomainRequest) SetModuleName(v string) *UpdateInstanceCustomizedDomainRequest {
	s.ModuleName = &v
	return s
}

func (s *UpdateInstanceCustomizedDomainRequest) Validate() error {
	return dara.Validate(s)
}
