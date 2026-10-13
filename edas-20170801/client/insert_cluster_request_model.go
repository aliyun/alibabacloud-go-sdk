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
	// - 2: ECS cluster.
	//
	// - 3: EDAS built-in Kubernetes cluster.
	//
	// - 5: Kubernetes cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	ClusterType *int32 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The IaaS resource provider.
	//
	// <note>When you select Alibaba Cloud, you must use uppercase Pinyin, such as `ALIYUN`.</note>
	//
	// example:
	//
	// ALIYUN
	IaasProvider *string `json:"IaasProvider,omitempty" xml:"IaasProvider,omitempty"`
	// The region ID of the custom namespace. Format: `physical region:custom namespace identifier`, such as `cn-hangzhou:test`.
	//
	// example:
	//
	// cn-beijing:td****
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	// The network type. Valid values:
	//
	// - 1: Classic network.
	//
	// - 2: VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	NetworkMode *int32 `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// **[Deprecated]*	- The CPU overselling ratio for Docker clusters. Valid values:
	//
	// - 2: 1:2 ratio.
	//
	// - 4: 1:4 ratio.
	//
	// - 8: 1:8 ratio.
	//
	// example:
	//
	// 2
	OversoldFactor *int32 `json:"OversoldFactor,omitempty" xml:"OversoldFactor,omitempty"`
	// The ID of the VPC. This parameter is required if the network type is set to VPC.
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
