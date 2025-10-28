// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterList(v *ListClusterResponseBodyClusterList) *ListClusterResponseBody
	GetClusterList() *ListClusterResponseBodyClusterList
	SetCode(v int32) *ListClusterResponseBody
	GetCode() *int32
	SetMessage(v string) *ListClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListClusterResponseBody
	GetRequestId() *string
}

type ListClusterResponseBody struct {
	// The clusters.
	ClusterList *ListClusterResponseBodyClusterList `json:"ClusterList,omitempty" xml:"ClusterList,omitempty" type:"Struct"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1053-08e4-47a5-b2ab-5c0323de****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterResponseBody) GetClusterList() *ListClusterResponseBodyClusterList {
	return s.ClusterList
}

func (s *ListClusterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClusterResponseBody) SetClusterList(v *ListClusterResponseBodyClusterList) *ListClusterResponseBody {
	s.ClusterList = v
	return s
}

func (s *ListClusterResponseBody) SetCode(v int32) *ListClusterResponseBody {
	s.Code = &v
	return s
}

func (s *ListClusterResponseBody) SetMessage(v string) *ListClusterResponseBody {
	s.Message = &v
	return s
}

func (s *ListClusterResponseBody) SetRequestId(v string) *ListClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClusterResponseBody) Validate() error {
	if s.ClusterList != nil {
		if err := s.ClusterList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListClusterResponseBodyClusterList struct {
	Cluster []*ListClusterResponseBodyClusterListCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Repeated"`
}

func (s ListClusterResponseBodyClusterList) String() string {
	return dara.Prettify(s)
}

func (s ListClusterResponseBodyClusterList) GoString() string {
	return s.String()
}

func (s *ListClusterResponseBodyClusterList) GetCluster() []*ListClusterResponseBodyClusterListCluster {
	return s.Cluster
}

func (s *ListClusterResponseBodyClusterList) SetCluster(v []*ListClusterResponseBodyClusterListCluster) *ListClusterResponseBodyClusterList {
	s.Cluster = v
	return s
}

func (s *ListClusterResponseBodyClusterList) Validate() error {
	if s.Cluster != nil {
		for _, item := range s.Cluster {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterResponseBodyClusterListCluster struct {
	// The ID of the cluster in EDAS.
	//
	// example:
	//
	// b98b5919-c111-4dad-9f74-7233********
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// cluster-test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The type of the cluster. Valid values:
	//
	// 	- 0: regular Docker cluster
	//
	// 	- 1: Swarm cluster
	//
	// 	- 2: Elastic Compute Service (ECS) cluster
	//
	// 	- 3: self-managed Kubernetes cluster in Enterprise Distributed Application Service (EDAS)
	//
	// 	- 4: cluster in which Pandora automatically registers applications
	//
	// 	- 5: ACK cluster
	//
	// example:
	//
	// 2
	ClusterType *int32 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The total number of CPU cores.
	//
	// example:
	//
	// 2
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of used CPU cores.
	//
	// example:
	//
	// 1
	CpuUsed *int32 `json:"CpuUsed,omitempty" xml:"CpuUsed,omitempty"`
	// The timestamp when the cluster was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1502888064561
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the cluster in Container Service for Kubernetes (ACK).
	//
	// example:
	//
	// c2ce************b9203a9
	CsClusterId *string `json:"CsClusterId,omitempty" xml:"CsClusterId,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The provider of the cluster.
	//
	// example:
	//
	// ALIYUN
	IaasProvider *string `json:"IaasProvider,omitempty" xml:"IaasProvider,omitempty"`
	// The total size of memory. Unit: MB.
	//
	// example:
	//
	// 3072
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The size of used memory. Unit: MB.
	//
	// example:
	//
	// 200
	MemUsed *int32 `json:"MemUsed,omitempty" xml:"MemUsed,omitempty"`
	// The network type of the cluster. Valid values:
	//
	// 	- 1: classic network
	//
	// 	- 2: virtual private cloud (VPC)
	//
	// example:
	//
	// 1
	NetworkMode *int32 `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 2
	NodeNum *int32 `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// The CPU overcommit ratio that is supported by a Docker cluster. Valid values:
	//
	// 	- 1: 1:1, which means that CPU resources are not overcommitted.
	//
	// 	- 2: 1:2, which means that CPU resources are overcommitted by 1:2.
	//
	// 	- 4: 1:4, which means that CPU resources are overcommitted by 1:4.
	//
	// 	- 8: 1:8, which means that CPU resources are overcommitted by 1:8.
	//
	// example:
	//
	// 1
	OversoldFactor *int32 `json:"OversoldFactor,omitempty" xml:"OversoldFactor,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// 461
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The timestamp when the cluster was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1533820823203
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-23727****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListClusterResponseBodyClusterListCluster) String() string {
	return dara.Prettify(s)
}

func (s ListClusterResponseBodyClusterListCluster) GoString() string {
	return s.String()
}

func (s *ListClusterResponseBodyClusterListCluster) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterResponseBodyClusterListCluster) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListClusterResponseBodyClusterListCluster) GetClusterType() *int32 {
	return s.ClusterType
}

func (s *ListClusterResponseBodyClusterListCluster) GetCpu() *int32 {
	return s.Cpu
}

func (s *ListClusterResponseBodyClusterListCluster) GetCpuUsed() *int32 {
	return s.CpuUsed
}

func (s *ListClusterResponseBodyClusterListCluster) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListClusterResponseBodyClusterListCluster) GetCsClusterId() *string {
	return s.CsClusterId
}

func (s *ListClusterResponseBodyClusterListCluster) GetDescription() *string {
	return s.Description
}

func (s *ListClusterResponseBodyClusterListCluster) GetIaasProvider() *string {
	return s.IaasProvider
}

func (s *ListClusterResponseBodyClusterListCluster) GetMem() *int32 {
	return s.Mem
}

func (s *ListClusterResponseBodyClusterListCluster) GetMemUsed() *int32 {
	return s.MemUsed
}

func (s *ListClusterResponseBodyClusterListCluster) GetNetworkMode() *int32 {
	return s.NetworkMode
}

func (s *ListClusterResponseBodyClusterListCluster) GetNodeNum() *int32 {
	return s.NodeNum
}

func (s *ListClusterResponseBodyClusterListCluster) GetOversoldFactor() *int32 {
	return s.OversoldFactor
}

func (s *ListClusterResponseBodyClusterListCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *ListClusterResponseBodyClusterListCluster) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListClusterResponseBodyClusterListCluster) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListClusterResponseBodyClusterListCluster) GetVpcId() *string {
	return s.VpcId
}

func (s *ListClusterResponseBodyClusterListCluster) SetClusterId(v string) *ListClusterResponseBodyClusterListCluster {
	s.ClusterId = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetClusterName(v string) *ListClusterResponseBodyClusterListCluster {
	s.ClusterName = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetClusterType(v int32) *ListClusterResponseBodyClusterListCluster {
	s.ClusterType = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetCpu(v int32) *ListClusterResponseBodyClusterListCluster {
	s.Cpu = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetCpuUsed(v int32) *ListClusterResponseBodyClusterListCluster {
	s.CpuUsed = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetCreateTime(v int64) *ListClusterResponseBodyClusterListCluster {
	s.CreateTime = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetCsClusterId(v string) *ListClusterResponseBodyClusterListCluster {
	s.CsClusterId = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetDescription(v string) *ListClusterResponseBodyClusterListCluster {
	s.Description = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetIaasProvider(v string) *ListClusterResponseBodyClusterListCluster {
	s.IaasProvider = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetMem(v int32) *ListClusterResponseBodyClusterListCluster {
	s.Mem = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetMemUsed(v int32) *ListClusterResponseBodyClusterListCluster {
	s.MemUsed = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetNetworkMode(v int32) *ListClusterResponseBodyClusterListCluster {
	s.NetworkMode = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetNodeNum(v int32) *ListClusterResponseBodyClusterListCluster {
	s.NodeNum = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetOversoldFactor(v int32) *ListClusterResponseBodyClusterListCluster {
	s.OversoldFactor = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetRegionId(v string) *ListClusterResponseBodyClusterListCluster {
	s.RegionId = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetResourceGroupId(v string) *ListClusterResponseBodyClusterListCluster {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetUpdateTime(v int64) *ListClusterResponseBodyClusterListCluster {
	s.UpdateTime = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) SetVpcId(v string) *ListClusterResponseBodyClusterListCluster {
	s.VpcId = &v
	return s
}

func (s *ListClusterResponseBodyClusterListCluster) Validate() error {
	return dara.Validate(s)
}
