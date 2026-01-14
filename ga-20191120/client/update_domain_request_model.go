// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *UpdateDomainRequest
	GetDomain() *string
	SetRegionId(v string) *UpdateDomainRequest
	GetRegionId() *string
	SetTargetDomain(v string) *UpdateDomainRequest
	GetTargetDomain() *string
}

type UpdateDomainRequest struct {
	// The new accelerated domain name.
	//
	// Only primary domain names are supported, such as `example.net`.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.net
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The accelerated domain name to be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	TargetDomain *string `json:"TargetDomain,omitempty" xml:"TargetDomain,omitempty"`
}

func (s UpdateDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateDomainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDomainRequest) GetTargetDomain() *string {
	return s.TargetDomain
}

func (s *UpdateDomainRequest) SetDomain(v string) *UpdateDomainRequest {
	s.Domain = &v
	return s
}

func (s *UpdateDomainRequest) SetRegionId(v string) *UpdateDomainRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDomainRequest) SetTargetDomain(v string) *UpdateDomainRequest {
	s.TargetDomain = &v
	return s
}

func (s *UpdateDomainRequest) Validate() error {
	return dara.Validate(s)
}
