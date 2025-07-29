// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterType(v string) *DescribeClustersRequest
	GetClusterType() *string
	SetName(v string) *DescribeClustersRequest
	GetName() *string
	SetResourceGroupId(v string) *DescribeClustersRequest
	GetResourceGroupId() *string
}

type DescribeClustersRequest struct {
	// The cluster type.
	//
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
	// The cluster name based on which the system performs fuzzy searches among the clusters that belong to the current Alibaba Cloud account.
	//
	// example:
	//
	// test
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	ResourceGroupId *string `json:"resource_group_id,omitempty" xml:"resource_group_id,omitempty"`
}

func (s DescribeClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeClustersRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeClustersRequest) GetName() *string {
	return s.Name
}

func (s *DescribeClustersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeClustersRequest) SetClusterType(v string) *DescribeClustersRequest {
	s.ClusterType = &v
	return s
}

func (s *DescribeClustersRequest) SetName(v string) *DescribeClustersRequest {
	s.Name = &v
	return s
}

func (s *DescribeClustersRequest) SetResourceGroupId(v string) *DescribeClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClustersRequest) Validate() error {
	return dara.Validate(s)
}
