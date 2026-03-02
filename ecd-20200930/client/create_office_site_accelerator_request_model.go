// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOfficeSiteAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccelerateRegion(v []*CreateOfficeSiteAcceleratorRequestAccelerateRegion) *CreateOfficeSiteAcceleratorRequest
	GetAccelerateRegion() []*CreateOfficeSiteAcceleratorRequestAccelerateRegion
	SetName(v string) *CreateOfficeSiteAcceleratorRequest
	GetName() *string
	SetOfficeSiteId(v string) *CreateOfficeSiteAcceleratorRequest
	GetOfficeSiteId() *string
	SetRegionId(v string) *CreateOfficeSiteAcceleratorRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateOfficeSiteAcceleratorRequest
	GetResourceGroupId() *string
}

type CreateOfficeSiteAcceleratorRequest struct {
	// The regions to include in global acceleration.
	//
	// This parameter is required.
	AccelerateRegion []*CreateOfficeSiteAcceleratorRequestAccelerateRegion `json:"AccelerateRegion,omitempty" xml:"AccelerateRegion,omitempty" type:"Repeated"`
	// The name of the GA instance.
	//
	// example:
	//
	// testGA
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The office network ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai+dir-259382****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-3mtuc28rx95lx****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateOfficeSiteAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOfficeSiteAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *CreateOfficeSiteAcceleratorRequest) GetAccelerateRegion() []*CreateOfficeSiteAcceleratorRequestAccelerateRegion {
	return s.AccelerateRegion
}

func (s *CreateOfficeSiteAcceleratorRequest) GetName() *string {
	return s.Name
}

func (s *CreateOfficeSiteAcceleratorRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateOfficeSiteAcceleratorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateOfficeSiteAcceleratorRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateOfficeSiteAcceleratorRequest) SetAccelerateRegion(v []*CreateOfficeSiteAcceleratorRequestAccelerateRegion) *CreateOfficeSiteAcceleratorRequest {
	s.AccelerateRegion = v
	return s
}

func (s *CreateOfficeSiteAcceleratorRequest) SetName(v string) *CreateOfficeSiteAcceleratorRequest {
	s.Name = &v
	return s
}

func (s *CreateOfficeSiteAcceleratorRequest) SetOfficeSiteId(v string) *CreateOfficeSiteAcceleratorRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateOfficeSiteAcceleratorRequest) SetRegionId(v string) *CreateOfficeSiteAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *CreateOfficeSiteAcceleratorRequest) SetResourceGroupId(v string) *CreateOfficeSiteAcceleratorRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateOfficeSiteAcceleratorRequest) Validate() error {
	if s.AccelerateRegion != nil {
		for _, item := range s.AccelerateRegion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateOfficeSiteAcceleratorRequestAccelerateRegion struct {
	// The ID of the region to include in global acceleration.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-southeast-1
	AccelerateRegionId *string `json:"AccelerateRegionId,omitempty" xml:"AccelerateRegionId,omitempty"`
	// The bandwidth that you want to allocate to the acceleration region. Unit: Mbit/s.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The IP version used to connect to the GA instance.
	//
	// >  Only pay-as-you-go standard GA instances support `DUAL_STACK`.
	//
	// Valid values:
	//
	// 	- DUAL_STACK: IPv4 and IPv6.
	//
	// 	- IPv6: IPv6.
	//
	// 	- IPv4 (default): IPv4.
	//
	// example:
	//
	// IPv4
	IpVersion *string `json:"IpVersion,omitempty" xml:"IpVersion,omitempty"`
	// The line type of the elastic IP address (EIP) in the acceleration region.
	//
	// >
	//
	// 	- This parameter is required only if the bandwidth metering method of the GA instance is **pay-by-data transfer**.
	//
	// 	- Different acceleration regions support different line types of EIPs.
	//
	// Valid values:
	//
	// 	- BGP: BGP (Multi-ISP) lines.
	//
	// 	- BGP_PRO: BGP (Multi-ISP) Pro.
	//
	// This parameter is required.
	//
	// example:
	//
	// BGP
	IspType *string `json:"IspType,omitempty" xml:"IspType,omitempty"`
}

func (s CreateOfficeSiteAcceleratorRequestAccelerateRegion) String() string {
	return dara.Prettify(s)
}

func (s CreateOfficeSiteAcceleratorRequestAccelerateRegion) GoString() string {
	return s.String()
}

func (s *CreateOfficeSiteAcceleratorRequestAccelerateRegion) GetAccelerateRegionId() *string {
	return s.AccelerateRegionId
}

func (s *CreateOfficeSiteAcceleratorRequestAccelerateRegion) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *CreateOfficeSiteAcceleratorRequestAccelerateRegion) GetIpVersion() *string {
	return s.IpVersion
}

func (s *CreateOfficeSiteAcceleratorRequestAccelerateRegion) GetIspType() *string {
	return s.IspType
}

func (s *CreateOfficeSiteAcceleratorRequestAccelerateRegion) SetAccelerateRegionId(v string) *CreateOfficeSiteAcceleratorRequestAccelerateRegion {
	s.AccelerateRegionId = &v
	return s
}

func (s *CreateOfficeSiteAcceleratorRequestAccelerateRegion) SetBandwidth(v int32) *CreateOfficeSiteAcceleratorRequestAccelerateRegion {
	s.Bandwidth = &v
	return s
}

func (s *CreateOfficeSiteAcceleratorRequestAccelerateRegion) SetIpVersion(v string) *CreateOfficeSiteAcceleratorRequestAccelerateRegion {
	s.IpVersion = &v
	return s
}

func (s *CreateOfficeSiteAcceleratorRequestAccelerateRegion) SetIspType(v string) *CreateOfficeSiteAcceleratorRequestAccelerateRegion {
	s.IspType = &v
	return s
}

func (s *CreateOfficeSiteAcceleratorRequestAccelerateRegion) Validate() error {
	return dara.Validate(s)
}
