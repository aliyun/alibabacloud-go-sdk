// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAckClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*DescribeAckClustersResponseBodyClusters) *DescribeAckClustersResponseBody
	GetClusters() []*DescribeAckClustersResponseBodyClusters
	SetPageNo(v int32) *DescribeAckClustersResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *DescribeAckClustersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAckClustersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAckClustersResponseBody
	GetTotalCount() *int32
}

type DescribeAckClustersResponseBody struct {
	Clusters []*DescribeAckClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// C5DDD596-1191-5F36-A504-8733045A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAckClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClustersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAckClustersResponseBody) GetClusters() []*DescribeAckClustersResponseBodyClusters {
	return s.Clusters
}

func (s *DescribeAckClustersResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *DescribeAckClustersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAckClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAckClustersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAckClustersResponseBody) SetClusters(v []*DescribeAckClustersResponseBodyClusters) *DescribeAckClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *DescribeAckClustersResponseBody) SetPageNo(v int32) *DescribeAckClustersResponseBody {
	s.PageNo = &v
	return s
}

func (s *DescribeAckClustersResponseBody) SetPageSize(v int32) *DescribeAckClustersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAckClustersResponseBody) SetRequestId(v string) *DescribeAckClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAckClustersResponseBody) SetTotalCount(v int32) *DescribeAckClustersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAckClustersResponseBody) Validate() error {
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

type DescribeAckClustersResponseBodyClusters struct {
	// example:
	//
	// cb0f5640b1b2d404cad6ba21509d7847b
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// ack-cluster-name
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// example:
	//
	// 135809047715****
	MemberUid *string `json:"MemberUid,omitempty" xml:"MemberUid,omitempty"`
	// example:
	//
	// terway-eniip
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// Default
	Profile *string `json:"Profile,omitempty" xml:"Profile,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// running
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// example:
	//
	// vpc-2vcg932hsxsxuqbgl****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeAckClustersResponseBodyClusters) String() string {
	return dara.Prettify(s)
}

func (s DescribeAckClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *DescribeAckClustersResponseBodyClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeAckClustersResponseBodyClusters) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeAckClustersResponseBodyClusters) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *DescribeAckClustersResponseBodyClusters) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeAckClustersResponseBodyClusters) GetMemberUid() *string {
	return s.MemberUid
}

func (s *DescribeAckClustersResponseBodyClusters) GetNetwork() *string {
	return s.Network
}

func (s *DescribeAckClustersResponseBodyClusters) GetProfile() *string {
	return s.Profile
}

func (s *DescribeAckClustersResponseBodyClusters) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAckClustersResponseBodyClusters) GetState() *string {
	return s.State
}

func (s *DescribeAckClustersResponseBodyClusters) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeAckClustersResponseBodyClusters) SetClusterId(v string) *DescribeAckClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *DescribeAckClustersResponseBodyClusters) SetClusterName(v string) *DescribeAckClustersResponseBodyClusters {
	s.ClusterName = &v
	return s
}

func (s *DescribeAckClustersResponseBodyClusters) SetClusterSpec(v string) *DescribeAckClustersResponseBodyClusters {
	s.ClusterSpec = &v
	return s
}

func (s *DescribeAckClustersResponseBodyClusters) SetClusterType(v string) *DescribeAckClustersResponseBodyClusters {
	s.ClusterType = &v
	return s
}

func (s *DescribeAckClustersResponseBodyClusters) SetMemberUid(v string) *DescribeAckClustersResponseBodyClusters {
	s.MemberUid = &v
	return s
}

func (s *DescribeAckClustersResponseBodyClusters) SetNetwork(v string) *DescribeAckClustersResponseBodyClusters {
	s.Network = &v
	return s
}

func (s *DescribeAckClustersResponseBodyClusters) SetProfile(v string) *DescribeAckClustersResponseBodyClusters {
	s.Profile = &v
	return s
}

func (s *DescribeAckClustersResponseBodyClusters) SetRegionId(v string) *DescribeAckClustersResponseBodyClusters {
	s.RegionId = &v
	return s
}

func (s *DescribeAckClustersResponseBodyClusters) SetState(v string) *DescribeAckClustersResponseBodyClusters {
	s.State = &v
	return s
}

func (s *DescribeAckClustersResponseBodyClusters) SetVpcId(v string) *DescribeAckClustersResponseBodyClusters {
	s.VpcId = &v
	return s
}

func (s *DescribeAckClustersResponseBodyClusters) Validate() error {
	return dara.Validate(s)
}
