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
	// Specifies the cluster specification. Valid values:
	//
	// - `ack.standard`: Basic edition.
	//
	// - `ack.pro.small`: Pro edition.
	//
	// - `ack.pro.xlarge`: Pro XL.
	//
	// - `ack.pro.2xlarge`: Pro 2XL.
	//
	// - `ack.pro.4xlarge`: Pro 4XL (requires whitelist approval from customer service).
	//
	// Pro XL, Pro 2XL, and Pro 4XL are the three tiers provided by <props="china">[ACK Pro provisioned control plane](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane)<props="intl">[ACK Pro provisioned control plane](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane). By pre-allocating and dedicating control plane resources, these tiers ensure that API concurrency and pod scheduling capabilities remain consistently at a high level. They are suitable for AI training and inference, ultra-large-scale clusters, and mission-critical workloads.
	//
	// For cluster management fees of the Pro edition and provisioned control plane editions, refer to <props="china">[Cluster management fees](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee)<props="intl">[Cluster management fees](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee).
	//
	// example:
	//
	// ack.standard
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// Specifies the cluster type. Valid values:
	//
	// - Kubernetes: ACK dedicated clusters.
	//
	// - ManagedKubernetes: ACK managed clusters, including ACK managed clusters (ACK Pro and ACK Basic), ACK Serverless clusters (Pro and Basic), ACK Edge clusters (Pro and Basic), and ACK Lingjun clusters (Pro).
	//
	// - ExternalKubernetes: registered clusters.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// Specifies the cluster name for fuzzy match.
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
	// The number of entries per page.
	//
	// example:
	//
	// 3
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// Specifies the cluster subtype. Valid values:
	//
	// - Default: ACK managed clusters, including ACK Pro and ACK Basic.
	//
	// - Edge: ACK Edge clusters, including ACK Edge Pro and ACK Edge Basic.
	//
	// - Serverless: ACK Serverless clusters, including ACK Serverless Pro and ACK Serverless Basic.
	//
	// - LingJun: ACK Lingjun clusters, available only in Pro edition.
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
