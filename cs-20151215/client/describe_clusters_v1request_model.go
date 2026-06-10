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
	// c3fb96524f9274b4495df0f12a6b5****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The cluster specification. This parameter is valid only when `cluster_type` is set to `ManagedKubernetes` and the `profile` parameter is specified. Valid values:
	//
	// - `ack.standard`: Standard
	//
	// - `ack.pro.small`: Pro
	//
	// - `ack.pro.xlarge`: Pro XL
	//
	// - `ack.pro.2xlarge`: Pro 2XL
	//
	// - `ack.pro.4xlarge`: Pro 4XL (Contact customer service to enable this option.)
	//
	// Pro XL, Pro 2XL, and Pro 4XL are three tiers provided by the <props="china">[ACK Pro provisioned control plane](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane)<props="intl">[ACK Pro provisioned control plane](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane). These tiers pre-allocate and dedicate control plane resources to ensure a consistently high, predictable level of performance for API concurrency and pod scheduling. They are suitable for AI training and inference, ultra-large-scale clusters, and mission-critical workloads.
	//
	// For information about the cluster management fees for Pro and provisioned control plane editions, see <props="china">[Cluster management fee](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee)<props="intl">[Cluster management fee](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee).
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// The cluster type.
	//
	// - `Kubernetes`: an ACK dedicated cluster.
	//
	// - `ManagedKubernetes`: an ACK managed cluster. This type includes ACK managed clusters (Pro and Standard), ACK Serverless clusters (Pro and Standard), ACK Edge clusters (Pro and Standard), and ACK Lingjun clusters (Pro).
	//
	// - `ExternalKubernetes`: a registered cluster.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// cluster-demo
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// When `cluster_type` is set to `ManagedKubernetes`, you can further specify a sub-type of the cluster.
	//
	// - `Default`: an ACK managed cluster. This includes ACK Pro and ACK Standard clusters.
	//
	// - `Edge`: an ACK Edge cluster. This includes ACK Edge Pro and ACK Edge Standard clusters.
	//
	// - `Serverless`: an ACK Serverless cluster. This includes ACK Serverless Pro and ACK Serverless Standard clusters.
	//
	// - `Lingjun`: an ACK Lingjun cluster (Pro edition).
	//
	// example:
	//
	// Default
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// The ID of the region to which the clusters belong.
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
