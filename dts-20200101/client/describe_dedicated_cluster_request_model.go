// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedClusterId(v string) *DescribeDedicatedClusterRequest
	GetDedicatedClusterId() *string
	SetOwnerId(v string) *DescribeDedicatedClusterRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeDedicatedClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDedicatedClusterRequest
	GetResourceGroupId() *string
}

type DescribeDedicatedClusterRequest struct {
	// The ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// dtsCluster****
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	OwnerId            *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDedicatedClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterRequest) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *DescribeDedicatedClusterRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeDedicatedClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDedicatedClusterRequest) SetDedicatedClusterId(v string) *DescribeDedicatedClusterRequest {
	s.DedicatedClusterId = &v
	return s
}

func (s *DescribeDedicatedClusterRequest) SetOwnerId(v string) *DescribeDedicatedClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedClusterRequest) SetRegionId(v string) *DescribeDedicatedClusterRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedClusterRequest) SetResourceGroupId(v string) *DescribeDedicatedClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDedicatedClusterRequest) Validate() error {
	return dara.Validate(s)
}
