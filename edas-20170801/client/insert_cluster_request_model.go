// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *InsertClusterRequest
	GetClusterName() *string
	SetClusterType(v int32) *InsertClusterRequest
	GetClusterType() *int32
	SetIaasProvider(v string) *InsertClusterRequest
	GetIaasProvider() *string
	SetLogicalRegionId(v string) *InsertClusterRequest
	GetLogicalRegionId() *string
	SetNetworkMode(v int32) *InsertClusterRequest
	GetNetworkMode() *int32
	SetOversoldFactor(v int32) *InsertClusterRequest
	GetOversoldFactor() *int32
	SetVpcId(v string) *InsertClusterRequest
	GetVpcId() *string
}

type InsertClusterRequest struct {
	// The name of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****_product_test2
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The type of the cluster. Valid values:
	//
	// 	- 2: Elastic Compute Service (ECS) cluster
	//
	// 	- 3: self-managed Kubernetes cluster in Enterprise Distributed Application Service (EDAS)
	//
	// 	- 5: Kubernetes cluster
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ClusterType *int32 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The provider of Infrastructure as a Service (IaaS) resources that are used in the cluster.
	//
	// When you use Alibaba Cloud, set the value to `ALIYUN`. The value is case-sensitive.
	//
	// example:
	//
	// ALIYUN
	IaasProvider *string `json:"IaasProvider,omitempty" xml:"IaasProvider,omitempty"`
	// The ID of the custom namespace. The ID is in the `physical region ID:custom namespace identifier` format. Example: `cn-hangzhou:test`.
	//
	// example:
	//
	// cn-beijing:td****
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	// The network type of the cluster. Valid values:
	//
	// 	- 1: classic network
	//
	// 	- 2: virtual private cloud (VPC)
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	NetworkMode *int32 `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// **This parameter is deprecated.*	- The CPU overcommit ratio supported by a Docker cluster. Valid values:
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
	// The ID of the VPC. This parameter is required if you set the NetworkMode parameter to 2.
	//
	// example:
	//
	// vpc-2zef6ob8mrlzv8x3q****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s InsertClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertClusterRequest) GoString() string {
	return s.String()
}

func (s *InsertClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *InsertClusterRequest) GetClusterType() *int32 {
	return s.ClusterType
}

func (s *InsertClusterRequest) GetIaasProvider() *string {
	return s.IaasProvider
}

func (s *InsertClusterRequest) GetLogicalRegionId() *string {
	return s.LogicalRegionId
}

func (s *InsertClusterRequest) GetNetworkMode() *int32 {
	return s.NetworkMode
}

func (s *InsertClusterRequest) GetOversoldFactor() *int32 {
	return s.OversoldFactor
}

func (s *InsertClusterRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *InsertClusterRequest) SetClusterName(v string) *InsertClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *InsertClusterRequest) SetClusterType(v int32) *InsertClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *InsertClusterRequest) SetIaasProvider(v string) *InsertClusterRequest {
	s.IaasProvider = &v
	return s
}

func (s *InsertClusterRequest) SetLogicalRegionId(v string) *InsertClusterRequest {
	s.LogicalRegionId = &v
	return s
}

func (s *InsertClusterRequest) SetNetworkMode(v int32) *InsertClusterRequest {
	s.NetworkMode = &v
	return s
}

func (s *InsertClusterRequest) SetOversoldFactor(v int32) *InsertClusterRequest {
	s.OversoldFactor = &v
	return s
}

func (s *InsertClusterRequest) SetVpcId(v string) *InsertClusterRequest {
	s.VpcId = &v
	return s
}

func (s *InsertClusterRequest) Validate() error {
	return dara.Validate(s)
}
