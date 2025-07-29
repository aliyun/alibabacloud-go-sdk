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
	// The specification of the clusters to query. Valid values:
	//
	// 	- ack.pro.small: ACK Pro clusters.
	//
	// 	- ack.standard: ACK Basic clusters.
	//
	// example:
	//
	// ack.standard
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// The type of the clusters to query. Valid values:
	//
	// 	- Kubernetes: ACK dedicated clusters.
	//
	// 	- ManagedKubernetes: ACK managed clusters. ACK managed clusters include ACK Basic clusters, ACK Pro clusters, ACK Serverless Basic clusters, ACK Serverless Pro clusters, ACK Edge Basic clusters, ACK Edge Pro clusters, and ACK Lingjun Pro clusters.
	//
	// 	- ExternalKubernetes: registered clusters.
	//
	// example:
	//
	// Kubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// Perform a fuzzy search by using the cluster name.
	//
	// example:
	//
	// test-cluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of pages.
	//
	// example:
	//
	// 10
	PageNumber *int64 `json:"page_number,omitempty" xml:"page_number,omitempty"`
	// The number of records on each page.
	//
	// example:
	//
	// 3
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// The subtype of the clusters to query. Valid values:
	//
	// 	- Default: ACK managed clusters. ACK managed clusters include ACK Basic clusters and ACK Pro clusters.
	//
	// 	- Edge: ACK Edge clusters. ACK Edge clusters include ACK Edge Basic clusters and ACK Edge Pro clusters.
	//
	// 	- Serverless: ACK Serverless clusters. ACK Serverless clusters include ACK Serverless Basic clusters and ACK Serverless Pro clusters.
	//
	// 	- Lingjun: ACK Lingjun Pro clusters.
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
