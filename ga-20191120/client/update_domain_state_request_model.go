// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *UpdateDomainStateRequest
	GetDomain() *string
	SetRegionId(v string) *UpdateDomainStateRequest
	GetRegionId() *string
}

type UpdateDomainStateRequest struct {
	// The accelerated domain name whose ICP filing status you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateDomainStateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainStateRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainStateRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateDomainStateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDomainStateRequest) SetDomain(v string) *UpdateDomainStateRequest {
	s.Domain = &v
	return s
}

func (s *UpdateDomainStateRequest) SetRegionId(v string) *UpdateDomainStateRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDomainStateRequest) Validate() error {
	return dara.Validate(s)
}
