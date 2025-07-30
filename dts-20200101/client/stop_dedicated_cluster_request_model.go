// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDedicatedClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedClusterId(v string) *StopDedicatedClusterRequest
	GetDedicatedClusterId() *string
	SetDedicatedClusterName(v string) *StopDedicatedClusterRequest
	GetDedicatedClusterName() *string
	SetInstanceId(v string) *StopDedicatedClusterRequest
	GetInstanceId() *string
	SetOwnerId(v string) *StopDedicatedClusterRequest
	GetOwnerId() *string
	SetRegionId(v string) *StopDedicatedClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *StopDedicatedClusterRequest
	GetResourceGroupId() *string
}

type StopDedicatedClusterRequest struct {
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
	OwnerId    *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. This parameter is a global parameter and not required.
	//
	// example:
	//
	// The resource group ID. This parameter is a global parameter and not required.
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s StopDedicatedClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDedicatedClusterRequest) GoString() string {
	return s.String()
}

func (s *StopDedicatedClusterRequest) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *StopDedicatedClusterRequest) GetDedicatedClusterName() *string {
	return s.DedicatedClusterName
}

func (s *StopDedicatedClusterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopDedicatedClusterRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *StopDedicatedClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopDedicatedClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *StopDedicatedClusterRequest) SetDedicatedClusterId(v string) *StopDedicatedClusterRequest {
	s.DedicatedClusterId = &v
	return s
}

func (s *StopDedicatedClusterRequest) SetDedicatedClusterName(v string) *StopDedicatedClusterRequest {
	s.DedicatedClusterName = &v
	return s
}

func (s *StopDedicatedClusterRequest) SetInstanceId(v string) *StopDedicatedClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *StopDedicatedClusterRequest) SetOwnerId(v string) *StopDedicatedClusterRequest {
	s.OwnerId = &v
	return s
}

func (s *StopDedicatedClusterRequest) SetRegionId(v string) *StopDedicatedClusterRequest {
	s.RegionId = &v
	return s
}

func (s *StopDedicatedClusterRequest) SetResourceGroupId(v string) *StopDedicatedClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *StopDedicatedClusterRequest) Validate() error {
	return dara.Validate(s)
}
