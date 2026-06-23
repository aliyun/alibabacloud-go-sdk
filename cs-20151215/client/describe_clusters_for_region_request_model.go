// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClustersForRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClustersForRegionRequest
	GetClusterId() *string
	SetClusterSpec(v string) *DescribeClustersForRegionRequest
	GetClusterSpec() *string
	SetClusterType(v string) *DescribeClustersForRegionRequest
	GetClusterType() *string
	SetName(v string) *DescribeClustersForRegionRequest
	GetName() *string
	SetPageNumber(v int64) *DescribeClustersForRegionRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeClustersForRegionRequest
	GetPageSize() *int64
	SetProfile(v string) *DescribeClustersForRegionRequest
	GetProfile() *string
}

type DescribeClustersForRegionRequest struct {
	// The cluster ID.
	//
	// example:
	//
	// c8155823d057948c69a****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// Queries clusters of a specified specification. Valid values:
	//
	// - `ack.standard`: Basic
	//
	// - `ack.pro.small`: Pro
	//
	// - `ack.pro.xlarge`: Pro XL
	//
	// - `ack.pro.2xlarge`: Pro 2XL
	//
	// - `ack.pro.4xlarge`: Pro 4XL (contact customer service to add your account to the whitelist)
	//
	// Pro XL, Pro 2XL, and Pro 4XL are three tiers provided by <props="china">[ACK Pro Provisioned Control Plane](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane)<props="intl">[ACK Pro Provisioned Control Plane](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane). By pre-allocating and dedicating control plane resources, these tiers ensure that API concurrency and Pod scheduling capabilities remain at a consistently high level. They are suitable for AI training and inference, ultra-large-scale clusters, and mission-critical workloads.
	//
	// For information about cluster management fees for Pro and Provisioned Control Plane editions, see <props="china">[Cluster management fees](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee)<props="intl">[Cluster management fees](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee).
	//
	// example:
	//
	// ack.standard
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// Queries clusters of a specified type. Valid values:
	//
	// - Kubernetes: ACK dedicated cluster.
	//
	// - ManagedKubernetes: ACK managed cluster types, including ACK managed clusters (ACK Pro and ACK Basic), ACK Serverless clusters (Pro and Basic), ACK Edge clusters (Pro and Basic), and ACK Lingjun clusters (Pro).
	//
	// - ExternalKubernetes: registered cluster.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// Fuzzy search by cluster name.
	//
	// example:
	//
	// test-cluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number.
	//
	// example:
	//
	// 10
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 3
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// Queries clusters of a specified subtype. Valid values:
	//
	// - Default: ACK managed cluster, including ACK Pro and ACK Basic.
	//
	// - Edge: ACK Edge cluster, including ACK Edge Pro and ACK Edge Basic.
	//
	// - Serverless: ACK Serverless cluster, including ACK Serverless Pro and ACK Serverless Basic.
	//
	// - LingJun: ACK Lingjun cluster, available in Pro.
	//
	// example:
	//
	// Serverless
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
}

func (s DescribeClustersForRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersForRegionRequest) GoString() string {
	return s.String()
}

func (s *DescribeClustersForRegionRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClustersForRegionRequest) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *DescribeClustersForRegionRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeClustersForRegionRequest) GetName() *string {
	return s.Name
}

func (s *DescribeClustersForRegionRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeClustersForRegionRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeClustersForRegionRequest) GetProfile() *string {
	return s.Profile
}

func (s *DescribeClustersForRegionRequest) SetClusterId(v string) *DescribeClustersForRegionRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClustersForRegionRequest) SetClusterSpec(v string) *DescribeClustersForRegionRequest {
	s.ClusterSpec = &v
	return s
}

func (s *DescribeClustersForRegionRequest) SetClusterType(v string) *DescribeClustersForRegionRequest {
	s.ClusterType = &v
	return s
}

func (s *DescribeClustersForRegionRequest) SetName(v string) *DescribeClustersForRegionRequest {
	s.Name = &v
	return s
}

func (s *DescribeClustersForRegionRequest) SetPageNumber(v int64) *DescribeClustersForRegionRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClustersForRegionRequest) SetPageSize(v int64) *DescribeClustersForRegionRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClustersForRegionRequest) SetProfile(v string) *DescribeClustersForRegionRequest {
	s.Profile = &v
	return s
}

func (s *DescribeClustersForRegionRequest) Validate() error {
	return dara.Validate(s)
}
