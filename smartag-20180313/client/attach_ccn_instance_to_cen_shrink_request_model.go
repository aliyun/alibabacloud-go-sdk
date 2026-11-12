// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachCcnInstanceToCenShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCcnId(v string) *AttachCcnInstanceToCenShrinkRequest
	GetCcnId() *string
	SetCenId(v string) *AttachCcnInstanceToCenShrinkRequest
	GetCenId() *string
	SetCidrAllocationMode(v string) *AttachCcnInstanceToCenShrinkRequest
	GetCidrAllocationMode() *string
	SetRegionBindingListShrink(v string) *AttachCcnInstanceToCenShrinkRequest
	GetRegionBindingListShrink() *string
	SetRegionId(v string) *AttachCcnInstanceToCenShrinkRequest
	GetRegionId() *string
	SetSubnet(v string) *AttachCcnInstanceToCenShrinkRequest
	GetSubnet() *string
}

type AttachCcnInstanceToCenShrinkRequest struct {
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
	RegionBindingListShrink *string `json:"RegionBindingList,omitempty" xml:"RegionBindingList,omitempty"`
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

func (s AttachCcnInstanceToCenShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachCcnInstanceToCenShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachCcnInstanceToCenShrinkRequest) GetCcnId() *string {
	return s.CcnId
}

func (s *AttachCcnInstanceToCenShrinkRequest) GetCenId() *string {
	return s.CenId
}

func (s *AttachCcnInstanceToCenShrinkRequest) GetCidrAllocationMode() *string {
	return s.CidrAllocationMode
}

func (s *AttachCcnInstanceToCenShrinkRequest) GetRegionBindingListShrink() *string {
	return s.RegionBindingListShrink
}

func (s *AttachCcnInstanceToCenShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachCcnInstanceToCenShrinkRequest) GetSubnet() *string {
	return s.Subnet
}

func (s *AttachCcnInstanceToCenShrinkRequest) SetCcnId(v string) *AttachCcnInstanceToCenShrinkRequest {
	s.CcnId = &v
	return s
}

func (s *AttachCcnInstanceToCenShrinkRequest) SetCenId(v string) *AttachCcnInstanceToCenShrinkRequest {
	s.CenId = &v
	return s
}

func (s *AttachCcnInstanceToCenShrinkRequest) SetCidrAllocationMode(v string) *AttachCcnInstanceToCenShrinkRequest {
	s.CidrAllocationMode = &v
	return s
}

func (s *AttachCcnInstanceToCenShrinkRequest) SetRegionBindingListShrink(v string) *AttachCcnInstanceToCenShrinkRequest {
	s.RegionBindingListShrink = &v
	return s
}

func (s *AttachCcnInstanceToCenShrinkRequest) SetRegionId(v string) *AttachCcnInstanceToCenShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *AttachCcnInstanceToCenShrinkRequest) SetSubnet(v string) *AttachCcnInstanceToCenShrinkRequest {
	s.Subnet = &v
	return s
}

func (s *AttachCcnInstanceToCenShrinkRequest) Validate() error {
	return dara.Validate(s)
}
