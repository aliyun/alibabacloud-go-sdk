// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachCcnInstanceToCenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcnId(v string) *AttachCcnInstanceToCenRequest
	GetCcnId() *string
	SetCenId(v string) *AttachCcnInstanceToCenRequest
	GetCenId() *string
	SetCidrAllocationMode(v string) *AttachCcnInstanceToCenRequest
	GetCidrAllocationMode() *string
	SetRegionBindingList(v []*AttachCcnInstanceToCenRequestRegionBindingList) *AttachCcnInstanceToCenRequest
	GetRegionBindingList() []*AttachCcnInstanceToCenRequestRegionBindingList
	SetRegionId(v string) *AttachCcnInstanceToCenRequest
	GetRegionId() *string
	SetSubnet(v string) *AttachCcnInstanceToCenRequest
	GetSubnet() *string
}

type AttachCcnInstanceToCenRequest struct {
	// The ID of the Cloud Connect Network (CCN) to attach.
	//
	// example:
	//
	// ccn-isdjvvkexkrpk*****
	CcnId *string `json:"CcnId,omitempty" xml:"CcnId,omitempty"`
	// The ID of the CEN instance to which you want to attach the CCN instance.
	//
	// example:
	//
	// cen-joimmi1s2ob3rdxw79
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The CIDR block allocation mode. Valid values: - PER_REGION: Custom CIDR blocks are allocated for each region. - UNIFIED: A unified CIDR block is allocated.
	//
	// example:
	//
	// UNIFIED
	CidrAllocationMode *string `json:"CidrAllocationMode,omitempty" xml:"CidrAllocationMode,omitempty"`
	// The list of transit routers to be newly attached to the CCN instance.
	RegionBindingList []*AttachCcnInstanceToCenRequestRegionBindingList `json:"RegionBindingList,omitempty" xml:"RegionBindingList,omitempty" type:"Repeated"`
	// The region ID of the Cloud Connect Network (CCN) instance. You can invoke the DescribeRegions operation to query the regions supported by Smart Access Gateway and the corresponding region IDs.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Internet CIDR block used to attach the Cloud Connect Network (CCN) to the Cloud Enterprise Network (CEN).
	//
	// example:
	//
	// 172.16.55.0/24
	Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s AttachCcnInstanceToCenRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachCcnInstanceToCenRequest) GoString() string {
	return s.String()
}

func (s *AttachCcnInstanceToCenRequest) GetCcnId() *string {
	return s.CcnId
}

func (s *AttachCcnInstanceToCenRequest) GetCenId() *string {
	return s.CenId
}

func (s *AttachCcnInstanceToCenRequest) GetCidrAllocationMode() *string {
	return s.CidrAllocationMode
}

func (s *AttachCcnInstanceToCenRequest) GetRegionBindingList() []*AttachCcnInstanceToCenRequestRegionBindingList {
	return s.RegionBindingList
}

func (s *AttachCcnInstanceToCenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachCcnInstanceToCenRequest) GetSubnet() *string {
	return s.Subnet
}

func (s *AttachCcnInstanceToCenRequest) SetCcnId(v string) *AttachCcnInstanceToCenRequest {
	s.CcnId = &v
	return s
}

func (s *AttachCcnInstanceToCenRequest) SetCenId(v string) *AttachCcnInstanceToCenRequest {
	s.CenId = &v
	return s
}

func (s *AttachCcnInstanceToCenRequest) SetCidrAllocationMode(v string) *AttachCcnInstanceToCenRequest {
	s.CidrAllocationMode = &v
	return s
}

func (s *AttachCcnInstanceToCenRequest) SetRegionBindingList(v []*AttachCcnInstanceToCenRequestRegionBindingList) *AttachCcnInstanceToCenRequest {
	s.RegionBindingList = v
	return s
}

func (s *AttachCcnInstanceToCenRequest) SetRegionId(v string) *AttachCcnInstanceToCenRequest {
	s.RegionId = &v
	return s
}

func (s *AttachCcnInstanceToCenRequest) SetSubnet(v string) *AttachCcnInstanceToCenRequest {
	s.Subnet = &v
	return s
}

func (s *AttachCcnInstanceToCenRequest) Validate() error {
	if s.RegionBindingList != nil {
		for _, item := range s.RegionBindingList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachCcnInstanceToCenRequestRegionBindingList struct {
	// The custom CIDR block for each region when the CidrAllocationMode parameter is set to PER_REGION.
	//
	// example:
	//
	// 192.168.60.0/24
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	// The region ID of the transit router.
	//
	// example:
	//
	// cn-shanghai-finance-1
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the transit router used for forward routing.
	//
	// example:
	//
	// tr-wz9h1y152p7bajkehs5a3
	TransitRouterId *string `json:"TransitRouterId,omitempty" xml:"TransitRouterId,omitempty"`
}

func (s AttachCcnInstanceToCenRequestRegionBindingList) String() string {
	return dara.Prettify(s)
}

func (s AttachCcnInstanceToCenRequestRegionBindingList) GoString() string {
	return s.String()
}

func (s *AttachCcnInstanceToCenRequestRegionBindingList) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *AttachCcnInstanceToCenRequestRegionBindingList) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachCcnInstanceToCenRequestRegionBindingList) GetTransitRouterId() *string {
	return s.TransitRouterId
}

func (s *AttachCcnInstanceToCenRequestRegionBindingList) SetCidrBlock(v string) *AttachCcnInstanceToCenRequestRegionBindingList {
	s.CidrBlock = &v
	return s
}

func (s *AttachCcnInstanceToCenRequestRegionBindingList) SetRegionId(v string) *AttachCcnInstanceToCenRequestRegionBindingList {
	s.RegionId = &v
	return s
}

func (s *AttachCcnInstanceToCenRequestRegionBindingList) SetTransitRouterId(v string) *AttachCcnInstanceToCenRequestRegionBindingList {
	s.TransitRouterId = &v
	return s
}

func (s *AttachCcnInstanceToCenRequestRegionBindingList) Validate() error {
	return dara.Validate(s)
}
