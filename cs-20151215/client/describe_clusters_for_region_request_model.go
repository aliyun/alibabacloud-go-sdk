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
	// The ID of the cluster.
	//
	// example:
	//
	// c8155823d057948c69a****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The specification of the cluster. Valid values:
	//
	// - `ack.standard`: Standard Edition
	//
	// - `ack.pro.small`: Pro Edition
	//
	// - `ack.pro.xlarge`: Pro XL
	//
	// - `ack.pro.2xlarge`: Pro 2XL
	//
	// - `ack.pro.4xlarge`: Pro 4XL (To use this specification, you must submit a ticket.)
	//
	// Pro XL, Pro 2XL, and Pro 4XL are three specifications available for the <props="china">[ACK Pro provisioned control plane](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane)<props="intl">[ACK Pro provisioned control plane](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/user-guide/ack-pro-provisioned-control-plane). These specifications ensure a high and deterministic level of API concurrency and Pod scheduling capabilities by pre-allocating and dedicating control plane resources. They are suitable for AI training and inference, large-scale clusters, and mission-critical workloads.
	//
	// For information about the cluster management fees for Pro Edition and provisioned control plane clusters, see <props="china">[cluster management fee](https://help.aliyun.com/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee)<props="intl">[cluster management fee](https://www.alibabacloud.com/help/ack/ack-managed-and-ack-dedicated/product-overview/cluster-management-fee).
	//
	// example:
	//
	// ack.standard
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// The type of the cluster. Valid values:
	//
	// - `Kubernetes`: an ACK dedicated cluster.
	//
	// - `ManagedKubernetes`: an ACK managed cluster. This includes ACK managed clusters (Pro and Standard Editions), ACK Serverless clusters (Pro and Standard Editions), ACK Edge clusters (Pro and Standard Editions), and ACK Lingjun clusters (Pro Edition).
	//
	// - `ExternalKubernetes`: a registered cluster.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// The name of the cluster. Fuzzy search is supported.
	//
	// example:
	//
	// test-cluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number to return.
	//
	// example:
	//
	// 10
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 3
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The subtype of the cluster. Valid values:
	//
	// - `Default`: ACK managed clusters, including Pro and Standard Editions.
	//
	// - `Edge`: ACK Edge clusters, including Pro and Standard Editions.
	//
	// - `Serverless`: ACK Serverless clusters, including Pro and Standard Editions.
	//
	// - `LingJun`: ACK Lingjun clusters, available in the Pro Edition.
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
