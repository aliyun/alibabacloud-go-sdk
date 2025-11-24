// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*DescribeServiceMeshClustersResponseBodyClusters) *DescribeServiceMeshClustersResponseBody
	GetClusters() []*DescribeServiceMeshClustersResponseBodyClusters
	SetNumberOfClusters(v int64) *DescribeServiceMeshClustersResponseBody
	GetNumberOfClusters() *int64
	SetRequestId(v string) *DescribeServiceMeshClustersResponseBody
	GetRequestId() *string
}

type DescribeServiceMeshClustersResponseBody struct {
	// The queried clusters.
	Clusters []*DescribeServiceMeshClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The total number of ACK clusters in the current Region.
	//
	// example:
	//
	// 5
	NumberOfClusters *int64 `json:"NumberOfClusters,omitempty" xml:"NumberOfClusters,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeServiceMeshClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshClustersResponseBody) GetClusters() []*DescribeServiceMeshClustersResponseBodyClusters {
	return s.Clusters
}

func (s *DescribeServiceMeshClustersResponseBody) GetNumberOfClusters() *int64 {
	return s.NumberOfClusters
}

func (s *DescribeServiceMeshClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeServiceMeshClustersResponseBody) SetClusters(v []*DescribeServiceMeshClustersResponseBodyClusters) *DescribeServiceMeshClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeServiceMeshClustersResponseBody) SetNumberOfClusters(v int64) *DescribeServiceMeshClustersResponseBody {
	s.NumberOfClusters = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBody) SetRequestId(v string) *DescribeServiceMeshClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBody) Validate() error {
	if s.Clusters != nil {
		for _, item := range s.Clusters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeServiceMeshClustersResponseBodyClusters struct {
	// The domain name of the cluster.
	//
	// example:
	//
	// c.com
	ClusterDomain *string `json:"ClusterDomain,omitempty" xml:"ClusterDomain,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// c80f45444b3da447da60a911390c2****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster type.
	//
	// example:
	//
	// Ask
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 2020-05-12T15:38:16+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The error message about the cluster.
	//
	// example:
	//
	// fail
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates that the cluster is available or the reason why the cluster cannot be added to the ASM instance. Valid values:
	//
	// 	- `0`: The cluster can be added to the ASM instance.
	//
	// 	- `1`: The cluster cannot be added to the ASM instance because you do not have administrator permissions on the cluster.
	//
	// 	- `2`: The cluster cannot be added to the ASM instance because the cluster and the ASM instance reside in different VPCs between which no private connections are built.
	//
	// 	- `3`: The CIDR block of the cluster conflicts with that of the ASM instance.
	//
	// 	- `4`: The cluster has a namespace that is named istio-system.
	//
	// example:
	//
	// 0
	ForbiddenFlag *int64 `json:"ForbiddenFlag,omitempty" xml:"ForbiddenFlag,omitempty"`
	// The reason why the cluster on the data plane cannot be added to the ASM instance. The value is a JSON string in the following format:
	//
	//     [
	//
	//       {
	//
	//         "cluster": "cdd55bd6e054b4c6ca18ec02614******",
	//
	//         "object": "Pod",
	//
	//         "cidr": "172.16.0.0/24"
	//
	//       },
	//
	//       {
	//
	//         "cluster": "cfa37fdf7cb1641e1976f8293ac******",
	//
	//         "object": "Pod",
	//
	//         "cidr": "172.16.0.0/24"
	//
	//       }
	//
	//     ]
	//
	// In the preceding example, the CIDR block `172.16.0.0/24` of the pod in the `cdd55bd6e054b4c6ca18ec02614******` cluster conflicts with the CIDR block `172.16.0.0/24` of the pod in the `cfa37fdf7cb1641e1976f8293ac******` cluster.
	//
	// Valid values of the `object` parameter:
	//
	// 	- `Pod`
	//
	// 	- `Service`
	//
	// 	- `VSwitch`
	//
	// 	- `VPC`
	//
	// 	- `VPC CIDR`
	//
	// example:
	//
	// [{"cluster":"cdd55bd6e054b4c6ca18ec02614******", "object":"Pod", "cidr":"172.16.0.0/24"},{"cluster":"cfa37fdf7cb1641e1976f8293ac******", "object":"Pod", "cidr":"172.16.0.0/24"}]
	ForbiddenInfo *string `json:"ForbiddenInfo,omitempty" xml:"ForbiddenInfo,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// ask1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region in which the cluster resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ASM instance ID.
	//
	// example:
	//
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-bp197668l6iupljy****
	SgId *string `json:"SgId,omitempty" xml:"SgId,omitempty"`
	// The state of the cluster. Valid values:
	//
	// 	- `running`: The cluster is running.
	//
	// 	- `starting`: The cluster is starting.
	//
	// 	- `stopping`: The cluster is being stopped.
	//
	// 	- `stopped`: The cluster is stopped.
	//
	// 	- `failed`: The cluster fails to be run.
	//
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The time when the cluster was last modified.
	//
	// example:
	//
	// 2020-05-12T15:38:16+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The version number of the cluster.
	//
	// example:
	//
	// v1.16.6-aliyun.1
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-8vbrwmt95b4zf6wf7****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeServiceMeshClustersResponseBodyClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) GetClusterDomain() *string {
	return s.ClusterDomain
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) GetForbiddenFlag() *int64 {
	return s.ForbiddenFlag
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) GetForbiddenInfo() *string {
	return s.ForbiddenInfo
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) GetName() *string {
	return s.Name
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) GetSgId() *string {
	return s.SgId
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) GetState() *string {
	return s.State
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) GetVersion() *string {
	return s.Version
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetClusterDomain(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.ClusterDomain = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetClusterId(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetClusterType(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.ClusterType = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetCreationTime(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.CreationTime = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetErrorMessage(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetForbiddenFlag(v int64) *DescribeServiceMeshClustersResponseBodyClusters {
	s.ForbiddenFlag = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetForbiddenInfo(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.ForbiddenInfo = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetName(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.Name = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetRegionId(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.RegionId = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetServiceMeshId(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.ServiceMeshId = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetSgId(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.SgId = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetState(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.State = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetUpdateTime(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.UpdateTime = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetVersion(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.Version = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) SetVpcId(v string) *DescribeServiceMeshClustersResponseBodyClusters {
	s.VpcId = &v
	return s
}

func (s *DescribeServiceMeshClustersResponseBodyClusters) Validate() error {
	return dara.Validate(s)
}
