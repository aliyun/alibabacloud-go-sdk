// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *CreateDomainRequest
	GetDomain() *string
	SetFiling(v *CreateDomainRequestFiling) *CreateDomainRequest
	GetFiling() *CreateDomainRequestFiling
	SetInstanceId(v string) *CreateDomainRequest
	GetInstanceId() *string
}

type CreateDomainRequest struct {
	// The domain name of the website.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Registration information parameters.
	Filing *CreateDomainRequestFiling `json:"Filing,omitempty" xml:"Filing,omitempty" type:"Struct"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateDomainRequest) GetFiling() *CreateDomainRequestFiling {
	return s.Filing
}

func (s *CreateDomainRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDomainRequest) SetDomain(v string) *CreateDomainRequest {
	s.Domain = &v
	return s
}

func (s *CreateDomainRequest) SetFiling(v *CreateDomainRequestFiling) *CreateDomainRequest {
	s.Filing = v
	return s
}

func (s *CreateDomainRequest) SetInstanceId(v string) *CreateDomainRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDomainRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDomainRequestFiling struct {
	// Record number associated with the domain name.
	//
	// example:
	//
	// 浙xx-xxxxxx
	IcpNumber *string `json:"IcpNumber,omitempty" xml:"IcpNumber,omitempty"`
}

func (s CreateDomainRequestFiling) String() string {
	return dara.Prettify(s)
}

func (s CreateDomainRequestFiling) GoString() string {
	return s.String()
}

func (s *CreateDomainRequestFiling) GetIcpNumber() *string {
	return s.IcpNumber
}

func (s *CreateDomainRequestFiling) SetIcpNumber(v string) *CreateDomainRequestFiling {
	s.IcpNumber = &v
	return s
}

func (s *CreateDomainRequestFiling) Validate() error {
	return dara.Validate(s)
}
