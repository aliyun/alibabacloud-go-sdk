// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedClusterId(v string) *ModifyDedicatedClusterRequest
	GetDedicatedClusterId() *string
	SetDedicatedClusterName(v string) *ModifyDedicatedClusterRequest
	GetDedicatedClusterName() *string
	SetInstanceId(v string) *ModifyDedicatedClusterRequest
	GetInstanceId() *string
	SetOversoldRatio(v int32) *ModifyDedicatedClusterRequest
	GetOversoldRatio() *int32
	SetOwnerId(v string) *ModifyDedicatedClusterRequest
	GetOwnerId() *string
	SetRegionId(v string) *ModifyDedicatedClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyDedicatedClusterRequest
	GetResourceGroupId() *string
}

type ModifyDedicatedClusterRequest struct {
	// The ID of the cluster.
	//
	// >  You must specify one of the **InstanceId*	- and **DedicatedClusterId*	- parameters.
	//
	// example:
	//
	// dtscluster_h3fl1cs217sx952
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// dtscluster_test_001
	DedicatedClusterName *string `json:"DedicatedClusterName,omitempty" xml:"DedicatedClusterName,omitempty"`
	// The ID of the instance.
	//
	// >  You must specify one of the **InstanceId*	- and **DedicatedClusterId*	- parameters.
	//
	// example:
	//
	// rm-bp1162kryivb8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The overcommit ratio. Unit: %.
	//
	// example:
	//
	// 150
	OversoldRatio *int32  `json:"OversoldRatio,omitempty" xml:"OversoldRatio,omitempty"`
	OwnerId       *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the Data Transmission Service (DTS) instance resides.
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

func (s ModifyDedicatedClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedClusterRequest) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *ModifyDedicatedClusterRequest) GetDedicatedClusterName() *string {
	return s.DedicatedClusterName
}

func (s *ModifyDedicatedClusterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDedicatedClusterRequest) GetOversoldRatio() *int32 {
	return s.OversoldRatio
}

func (s *ModifyDedicatedClusterRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyDedicatedClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDedicatedClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDedicatedClusterRequest) SetDedicatedClusterId(v string) *ModifyDedicatedClusterRequest {
	s.DedicatedClusterId = &v
	return s
}

func (s *ModifyDedicatedClusterRequest) SetDedicatedClusterName(v string) *ModifyDedicatedClusterRequest {
	s.DedicatedClusterName = &v
	return s
}

func (s *ModifyDedicatedClusterRequest) SetInstanceId(v string) *ModifyDedicatedClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDedicatedClusterRequest) SetOversoldRatio(v int32) *ModifyDedicatedClusterRequest {
	s.OversoldRatio = &v
	return s
}

func (s *ModifyDedicatedClusterRequest) SetOwnerId(v string) *ModifyDedicatedClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedClusterRequest) SetRegionId(v string) *ModifyDedicatedClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDedicatedClusterRequest) SetResourceGroupId(v string) *ModifyDedicatedClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDedicatedClusterRequest) Validate() error {
	return dara.Validate(s)
}
