// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClustersV1Request interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClustersV1Request
	GetClusterId() *string
	SetClusterSpec(v string) *DescribeClustersV1Request
	GetClusterSpec() *string
	SetClusterType(v string) *DescribeClustersV1Request
	GetClusterType() *string
	SetName(v string) *DescribeClustersV1Request
	GetName() *string
	SetPageNumber(v int64) *DescribeClustersV1Request
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeClustersV1Request
	GetPageSize() *int64
	SetProfile(v string) *DescribeClustersV1Request
	GetProfile() *string
	SetRegionId(v string) *DescribeClustersV1Request
	GetRegionId() *string
}

type DescribeClustersV1Request struct {
	// The cluster ID.
	//
	// example:
	//
	// ca418e5e6fa2849d78301341700axxxxx
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// After you set `cluster_type` to `ManagedKubernetes` and configure the `profile` parameter, you can further specify the edition of the cluster. Valid values:
	//
	// 	- `ack.pro.small`: ACK Pro cluster.
	//
	// 	- `ack.standard`: ACK Basic cluster. If you leave the parameter empty, ACK Basic cluster is selected.
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// The type of the instance.
	//
	// 	- `Kubernetes`: ACK dedicated cluster.
	//
	// 	- `ManagedKubernetes`: ACK managed cluster. ACK managed clusters include ACK Basic clusters, ACK Pro clusters, ACK Serverless Basic clusters, ACK Serverless Pro clusters, ACK Edge Basic clusters, ACK Edge Pro clusters, and ACK Lingjun Pro clusters.
	//
	// 	- `ExternalKubernetes`: registered cluster
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// cluster-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 3
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// If you set `cluster_type` to `ManagedKubernetes`, an ACK managed cluster is created. In this case, you can further specify the cluster edition. Valid values:
	//
	// 	- `Default`: ACK managed cluster. ACK managed clusters include ACK Basic clusters and ACK Pro clusters.
	//
	// 	- `Edge`: ACK Edge cluster. ACK Edge clusters include ACK Edge Basic clusters and ACK Edge Pro clusters.
	//
	// 	- `Serverless`: ACK Serverless cluster. ACK Serverless clusters include ACK Serverless Basic clusters and ACK Serverless Pro clusters.
	//
	// 	- `Lingjun`: ACK Lingjun Pro cluster.
	//
	// example:
	//
	// Default
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// The region ID of the clusters. You can use this parameter to query all clusters in the specified region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
}

func (s DescribeClustersV1Request) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersV1Request) GoString() string {
	return s.String()
}

func (s *DescribeClustersV1Request) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClustersV1Request) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *DescribeClustersV1Request) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeClustersV1Request) GetName() *string {
	return s.Name
}

func (s *DescribeClustersV1Request) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeClustersV1Request) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeClustersV1Request) GetProfile() *string {
	return s.Profile
}

func (s *DescribeClustersV1Request) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClustersV1Request) SetClusterId(v string) *DescribeClustersV1Request {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersV1Request) SetClusterSpec(v string) *DescribeClustersV1Request {
	s.ClusterSpec = &v
	return s
}

func (s *DescribeClustersV1Request) SetClusterType(v string) *DescribeClustersV1Request {
	s.ClusterType = &v
	return s
}

func (s *DescribeClustersV1Request) SetName(v string) *DescribeClustersV1Request {
	s.Name = &v
	return s
}

func (s *DescribeClustersV1Request) SetPageNumber(v int64) *DescribeClustersV1Request {
	s.PageNumber = &v
	return s
}

func (s *DescribeClustersV1Request) SetPageSize(v int64) *DescribeClustersV1Request {
	s.PageSize = &v
	return s
}

func (s *DescribeClustersV1Request) SetProfile(v string) *DescribeClustersV1Request {
	s.Profile = &v
	return s
}

func (s *DescribeClustersV1Request) SetRegionId(v string) *DescribeClustersV1Request {
	s.RegionId = &v
	return s
}

func (s *DescribeClustersV1Request) Validate() error {
	return dara.Validate(s)
}
