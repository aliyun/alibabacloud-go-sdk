// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCluster(v *GetClusterResponseBodyCluster) *GetClusterResponseBody
	GetCluster() *GetClusterResponseBodyCluster
	SetCode(v int32) *GetClusterResponseBody
	GetCode() *int32
	SetMessage(v string) *GetClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetClusterResponseBody
	GetRequestId() *string
}

type GetClusterResponseBody struct {
	// The information about the cluster.
	Cluster *GetClusterResponseBodyCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Struct"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// d76db491
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBody) GetCluster() *GetClusterResponseBodyCluster {
	return s.Cluster
}

func (s *GetClusterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClusterResponseBody) SetCluster(v *GetClusterResponseBodyCluster) *GetClusterResponseBody {
	s.Cluster = v
	return s
}

func (s *GetClusterResponseBody) SetCode(v int32) *GetClusterResponseBody {
	s.Code = &v
	return s
}

func (s *GetClusterResponseBody) SetMessage(v string) *GetClusterResponseBody {
	s.Message = &v
	return s
}

func (s *GetClusterResponseBody) SetRequestId(v string) *GetClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClusterResponseBody) Validate() error {
	if s.Cluster != nil {
		if err := s.Cluster.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetClusterResponseBodyCluster struct {
	// The ID of the cluster.
	//
	// example:
	//
	// 5439271a-015b-433d-befb-d76d****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The import status of the cluster. Valid values:
	//
	// 	- 1: The cluster is imported.
	//
	// 	- 2: The cluster fails to be imported.
	//
	// 	- 3: The cluster is being imported.
	//
	// 	- 4: The cluster is deleted.
	//
	// 	- 0: The cluster is not imported.
	//
	// example:
	//
	// 0
	ClusterImportStatus *int32 `json:"ClusterImportStatus,omitempty" xml:"ClusterImportStatus,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// ClusterTest
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The type of the cluster. Valid values:
	//
	// 	- 0: regular Docker cluster
	//
	// 	- 1: Swarm cluster
	//
	// 	- 2: Elastic Compute Service (ECS) cluster
	//
	// 	- 3: self-managed Kubernetes cluster in EDAS
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
	// 4
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The number of used CPU cores.
	//
	// example:
	//
	// 2
	CpuUsed *int32 `json:"CpuUsed,omitempty" xml:"CpuUsed,omitempty"`
	// The time when the cluster was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1570708232145
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the Container Service for Kubernetes (ACK) cluster.
	//
	// example:
	//
	// c2ce62869f4d4466b920312315f05****
	CsClusterId *string `json:"CsClusterId,omitempty" xml:"CsClusterId,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The provider of Infrastructure as a Service (IaaS) resources used in the cluster.
	//
	// example:
	//
	// ALIYUN
	IaasProvider *string `json:"IaasProvider,omitempty" xml:"IaasProvider,omitempty"`
	// The total size of memory. Unit: MB.
	//
	// example:
	//
	// 2048
	Mem *int32 `json:"Mem,omitempty" xml:"Mem,omitempty"`
	// The size of used memory. Unit: MB.
	//
	// example:
	//
	// 1024
	MemUsed *int32 `json:"MemUsed,omitempty" xml:"MemUsed,omitempty"`
	// The network type of the cluster. Valid values:
	//
	// 	- 1: classic network
	//
	// 	- 2: virtual private cloud (VPC)
	//
	// example:
	//
	// 2
	NetworkMode *int32 `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// The number of ECS instances.
	//
	// example:
	//
	// 4
	NodeNum *int32 `json:"NodeNum,omitempty" xml:"NodeNum,omitempty"`
	// The overcommit ratio supported by a Docker cluster. Valid values:
	//
	// 	- 1: 1:1, which means that resources are not overcommitted.
	//
	// 	- 2: 1:2, which means that resources are overcommitted by 1:2.
	//
	// 	- 4: 1:4, which means that resources are overcommitted by 1:4.
	//
	// 	- 8: 1:8, which means that resources are overcommitted by 1:8.
	//
	// example:
	//
	// 2
	OversoldFactor *int32 `json:"OversoldFactor,omitempty" xml:"OversoldFactor,omitempty"`
	// The ID of the region where the cluster resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The subtype of the Kubernetes cluster. Valid values: ManagedKubernetes, Ask, and ExternalKubernetes. ManagedKubernetes refers to the ACK cluster. Ask refers to the Serverless Kubernetes (ASK) cluster. ExternalKubernetes refers to the external cluster.
	//
	// example:
	//
	// ManagedKubernetes
	SubClusterType *string `json:"SubClusterType,omitempty" xml:"SubClusterType,omitempty"`
	// The time when the cluster was last modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1570708232145
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-xxxxz1mlwpb****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetClusterResponseBodyCluster) String() string {
	return dara.Prettify(s)
}

func (s GetClusterResponseBodyCluster) GoString() string {
	return s.String()
}

func (s *GetClusterResponseBodyCluster) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetClusterResponseBodyCluster) GetClusterImportStatus() *int32 {
	return s.ClusterImportStatus
}

func (s *GetClusterResponseBodyCluster) GetClusterName() *string {
	return s.ClusterName
}

func (s *GetClusterResponseBodyCluster) GetClusterType() *int32 {
	return s.ClusterType
}

func (s *GetClusterResponseBodyCluster) GetCpu() *int32 {
	return s.Cpu
}

func (s *GetClusterResponseBodyCluster) GetCpuUsed() *int32 {
	return s.CpuUsed
}

func (s *GetClusterResponseBodyCluster) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetClusterResponseBodyCluster) GetCsClusterId() *string {
	return s.CsClusterId
}

func (s *GetClusterResponseBodyCluster) GetDescription() *string {
	return s.Description
}

func (s *GetClusterResponseBodyCluster) GetIaasProvider() *string {
	return s.IaasProvider
}

func (s *GetClusterResponseBodyCluster) GetMem() *int32 {
	return s.Mem
}

func (s *GetClusterResponseBodyCluster) GetMemUsed() *int32 {
	return s.MemUsed
}

func (s *GetClusterResponseBodyCluster) GetNetworkMode() *int32 {
	return s.NetworkMode
}

func (s *GetClusterResponseBodyCluster) GetNodeNum() *int32 {
	return s.NodeNum
}

func (s *GetClusterResponseBodyCluster) GetOversoldFactor() *int32 {
	return s.OversoldFactor
}

func (s *GetClusterResponseBodyCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *GetClusterResponseBodyCluster) GetSubClusterType() *string {
	return s.SubClusterType
}

func (s *GetClusterResponseBodyCluster) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetClusterResponseBodyCluster) GetVpcId() *string {
	return s.VpcId
}

func (s *GetClusterResponseBodyCluster) SetClusterId(v string) *GetClusterResponseBodyCluster {
	s.ClusterId = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetClusterImportStatus(v int32) *GetClusterResponseBodyCluster {
	s.ClusterImportStatus = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetClusterName(v string) *GetClusterResponseBodyCluster {
	s.ClusterName = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetClusterType(v int32) *GetClusterResponseBodyCluster {
	s.ClusterType = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetCpu(v int32) *GetClusterResponseBodyCluster {
	s.Cpu = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetCpuUsed(v int32) *GetClusterResponseBodyCluster {
	s.CpuUsed = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetCreateTime(v int64) *GetClusterResponseBodyCluster {
	s.CreateTime = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetCsClusterId(v string) *GetClusterResponseBodyCluster {
	s.CsClusterId = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetDescription(v string) *GetClusterResponseBodyCluster {
	s.Description = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetIaasProvider(v string) *GetClusterResponseBodyCluster {
	s.IaasProvider = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetMem(v int32) *GetClusterResponseBodyCluster {
	s.Mem = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetMemUsed(v int32) *GetClusterResponseBodyCluster {
	s.MemUsed = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetNetworkMode(v int32) *GetClusterResponseBodyCluster {
	s.NetworkMode = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetNodeNum(v int32) *GetClusterResponseBodyCluster {
	s.NodeNum = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetOversoldFactor(v int32) *GetClusterResponseBodyCluster {
	s.OversoldFactor = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetRegionId(v string) *GetClusterResponseBodyCluster {
	s.RegionId = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetSubClusterType(v string) *GetClusterResponseBodyCluster {
	s.SubClusterType = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetUpdateTime(v int64) *GetClusterResponseBodyCluster {
	s.UpdateTime = &v
	return s
}

func (s *GetClusterResponseBodyCluster) SetVpcId(v string) *GetClusterResponseBodyCluster {
	s.VpcId = &v
	return s
}

func (s *GetClusterResponseBodyCluster) Validate() error {
	return dara.Validate(s)
}
