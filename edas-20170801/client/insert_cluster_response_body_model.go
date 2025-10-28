// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCluster(v *InsertClusterResponseBodyCluster) *InsertClusterResponseBody
	GetCluster() *InsertClusterResponseBodyCluster
	SetCode(v int32) *InsertClusterResponseBody
	GetCode() *int32
	SetMessage(v string) *InsertClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsertClusterResponseBody
	GetRequestId() *string
}

type InsertClusterResponseBody struct {
	// The information about the cluster that was created.
	Cluster *InsertClusterResponseBodyCluster `json:"Cluster,omitempty" xml:"Cluster,omitempty" type:"Struct"`
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
	// b197-40ab-9155-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InsertClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertClusterResponseBody) GoString() string {
	return s.String()
}

func (s *InsertClusterResponseBody) GetCluster() *InsertClusterResponseBodyCluster {
	return s.Cluster
}

func (s *InsertClusterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InsertClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsertClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertClusterResponseBody) SetCluster(v *InsertClusterResponseBodyCluster) *InsertClusterResponseBody {
	s.Cluster = v
	return s
}

func (s *InsertClusterResponseBody) SetCode(v int32) *InsertClusterResponseBody {
	s.Code = &v
	return s
}

func (s *InsertClusterResponseBody) SetMessage(v string) *InsertClusterResponseBody {
	s.Message = &v
	return s
}

func (s *InsertClusterResponseBody) SetRequestId(v string) *InsertClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertClusterResponseBody) Validate() error {
	if s.Cluster != nil {
		if err := s.Cluster.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertClusterResponseBodyCluster struct {
	// The ID of cluster.
	//
	// example:
	//
	// 8705ad13-5d86-47fc-b68f-257b59ed****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// ****_product_test2
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The type of the cluster. Valid values:
	//
	// 	- 2: ECS cluster
	//
	// 	- 3: self-managed Kubernetes cluster in EDAS
	//
	// 	- 5: Kubernetes cluster
	//
	// example:
	//
	// 2
	ClusterType *int32 `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The provider of the IaaS resources that are used in the cluster.
	//
	// example:
	//
	// ALIYUN
	IaasProvider *string `json:"IaasProvider,omitempty" xml:"IaasProvider,omitempty"`
	// The network type of the cluster. Valid values:
	//
	// 	- 1: classic network
	//
	// 	- 2\\. VPC
	//
	// example:
	//
	// 2
	NetworkMode *int32 `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// **This parameter is deprecated.*	- The CPU overcommit ratio supported by the Docker cluster. Valid values:
	//
	// 	- 2: 1:2, which means that resources are overcommitted by 1:2.
	//
	// 	- 4: 1:4, which means that resources are overcommitted by 1:4.
	//
	// 	- 8: 1:8, which means that resources are overcommitted by 1:8.
	//
	// example:
	//
	// 1
	OversoldFactor *int32 `json:"OversoldFactor,omitempty" xml:"OversoldFactor,omitempty"`
	// The ID of the region in which the cluster resides.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-2zef6ob8mrlzv8x3q****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s InsertClusterResponseBodyCluster) String() string {
	return dara.Prettify(s)
}

func (s InsertClusterResponseBodyCluster) GoString() string {
	return s.String()
}

func (s *InsertClusterResponseBodyCluster) GetClusterId() *string {
	return s.ClusterId
}

func (s *InsertClusterResponseBodyCluster) GetClusterName() *string {
	return s.ClusterName
}

func (s *InsertClusterResponseBodyCluster) GetClusterType() *int32 {
	return s.ClusterType
}

func (s *InsertClusterResponseBodyCluster) GetIaasProvider() *string {
	return s.IaasProvider
}

func (s *InsertClusterResponseBodyCluster) GetNetworkMode() *int32 {
	return s.NetworkMode
}

func (s *InsertClusterResponseBodyCluster) GetOversoldFactor() *int32 {
	return s.OversoldFactor
}

func (s *InsertClusterResponseBodyCluster) GetRegionId() *string {
	return s.RegionId
}

func (s *InsertClusterResponseBodyCluster) GetVpcId() *string {
	return s.VpcId
}

func (s *InsertClusterResponseBodyCluster) SetClusterId(v string) *InsertClusterResponseBodyCluster {
	s.ClusterId = &v
	return s
}

func (s *InsertClusterResponseBodyCluster) SetClusterName(v string) *InsertClusterResponseBodyCluster {
	s.ClusterName = &v
	return s
}

func (s *InsertClusterResponseBodyCluster) SetClusterType(v int32) *InsertClusterResponseBodyCluster {
	s.ClusterType = &v
	return s
}

func (s *InsertClusterResponseBodyCluster) SetIaasProvider(v string) *InsertClusterResponseBodyCluster {
	s.IaasProvider = &v
	return s
}

func (s *InsertClusterResponseBodyCluster) SetNetworkMode(v int32) *InsertClusterResponseBodyCluster {
	s.NetworkMode = &v
	return s
}

func (s *InsertClusterResponseBodyCluster) SetOversoldFactor(v int32) *InsertClusterResponseBodyCluster {
	s.OversoldFactor = &v
	return s
}

func (s *InsertClusterResponseBodyCluster) SetRegionId(v string) *InsertClusterResponseBodyCluster {
	s.RegionId = &v
	return s
}

func (s *InsertClusterResponseBodyCluster) SetVpcId(v string) *InsertClusterResponseBodyCluster {
	s.VpcId = &v
	return s
}

func (s *InsertClusterResponseBodyCluster) Validate() error {
	return dara.Validate(s)
}
