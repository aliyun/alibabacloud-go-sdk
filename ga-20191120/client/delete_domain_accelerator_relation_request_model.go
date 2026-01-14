// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainAcceleratorRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorIds(v []*string) *DeleteDomainAcceleratorRelationRequest
	GetAcceleratorIds() []*string
	SetDomain(v string) *DeleteDomainAcceleratorRelationRequest
	GetDomain() *string
	SetRegionId(v string) *DeleteDomainAcceleratorRelationRequest
	GetRegionId() *string
}

type DeleteDomainAcceleratorRelationRequest struct {
	// The ID of the GA instance to be disassociated. You can specify up to 50 IDs.
	//
	// If you leave this parameter empty, all GA instances associated with the specified domain name are disassociated.
	AcceleratorIds []*string `json:"AcceleratorIds,omitempty" xml:"AcceleratorIds,omitempty" type:"Repeated"`
	// The accelerated domain name to be disassociated.
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

func (s DeleteDomainAcceleratorRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainAcceleratorRelationRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainAcceleratorRelationRequest) GetAcceleratorIds() []*string {
	return s.AcceleratorIds
}

func (s *DeleteDomainAcceleratorRelationRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteDomainAcceleratorRelationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDomainAcceleratorRelationRequest) SetAcceleratorIds(v []*string) *DeleteDomainAcceleratorRelationRequest {
	s.AcceleratorIds = v
	return s
}

func (s *DeleteDomainAcceleratorRelationRequest) SetDomain(v string) *DeleteDomainAcceleratorRelationRequest {
	s.Domain = &v
	return s
}

func (s *DeleteDomainAcceleratorRelationRequest) SetRegionId(v string) *DeleteDomainAcceleratorRelationRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDomainAcceleratorRelationRequest) Validate() error {
	return dara.Validate(s)
}
