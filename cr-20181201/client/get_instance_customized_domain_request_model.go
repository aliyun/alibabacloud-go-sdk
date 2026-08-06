// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceCustomizedDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *GetInstanceCustomizedDomainRequest
	GetDomain() *string
	SetInstanceId(v string) *GetInstanceCustomizedDomainRequest
	GetInstanceId() *string
	SetModuleName(v string) *GetInstanceCustomizedDomainRequest
	GetModuleName() *string
}

type GetInstanceCustomizedDomainRequest struct {
	// The custom domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// registry.ugnas.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
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

func (s GetInstanceCustomizedDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceCustomizedDomainRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceCustomizedDomainRequest) GetDomain() *string {
	return s.Domain
}

func (s *GetInstanceCustomizedDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceCustomizedDomainRequest) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetInstanceCustomizedDomainRequest) SetDomain(v string) *GetInstanceCustomizedDomainRequest {
	s.Domain = &v
	return s
}

func (s *GetInstanceCustomizedDomainRequest) SetInstanceId(v string) *GetInstanceCustomizedDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceCustomizedDomainRequest) SetModuleName(v string) *GetInstanceCustomizedDomainRequest {
	s.ModuleName = &v
	return s
}

func (s *GetInstanceCustomizedDomainRequest) Validate() error {
	return dara.Validate(s)
}
