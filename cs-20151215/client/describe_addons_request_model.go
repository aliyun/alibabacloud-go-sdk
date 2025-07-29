// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterProfile(v string) *DescribeAddonsRequest
	GetClusterProfile() *string
	SetClusterSpec(v string) *DescribeAddonsRequest
	GetClusterSpec() *string
	SetClusterType(v string) *DescribeAddonsRequest
	GetClusterType() *string
	SetClusterVersion(v string) *DescribeAddonsRequest
	GetClusterVersion() *string
	SetRegion(v string) *DescribeAddonsRequest
	GetRegion() *string
}

type DescribeAddonsRequest struct {
	// The cluster type. Valid values:
	//
	// 	- `Default`: ACK managed cluster
	//
	// 	- `Serverless`: ACK Serverless cluster
	//
	// 	- `Edge`: ACK Edge cluster
	//
	// example:
	//
	// Default
	ClusterProfile *string `json:"cluster_profile,omitempty" xml:"cluster_profile,omitempty"`
	// If you set `cluster_type` to `ManagedKubernetes` and specify `profile`, you can further specify the edition of the cluster. Valid values:
	//
	// 	- `ack.pro.small`: creates an ACK Pro cluster.
	//
	// 	- `ack.standard`: creates an ACK Basic cluster. If you leave the parameter empty, an ACK Basic cluster is created.
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// 	- `Kubernetes`: ACK dedicated cluster.
	//
	// 	- `ManagedKubernetes`: ACK managed cluster. ACK managed clusters include ACK Basic clusters, ACK Pro clusters, ACK Serverless Basic clusters, ACK Serverless Pro clusters, ACK Edge Basic clusters, ACK Edge Pro clusters, and ACK Lingjun Pro clusters.
	//
	// 	- `ExternalKubernetes`: registered cluster.
	//
	// example:
	//
	// kubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// The cluster version.
	//
	// example:
	//
	// 1.24.6-aliyun.1
	ClusterVersion *string `json:"cluster_version,omitempty" xml:"cluster_version,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
}

func (s DescribeAddonsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddonsRequest) GetClusterProfile() *string {
	return s.ClusterProfile
}

func (s *DescribeAddonsRequest) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *DescribeAddonsRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeAddonsRequest) GetClusterVersion() *string {
	return s.ClusterVersion
}

func (s *DescribeAddonsRequest) GetRegion() *string {
	return s.Region
}

func (s *DescribeAddonsRequest) SetClusterProfile(v string) *DescribeAddonsRequest {
	s.ClusterProfile = &v
	return s
}

func (s *DescribeAddonsRequest) SetClusterSpec(v string) *DescribeAddonsRequest {
	s.ClusterSpec = &v
	return s
}

func (s *DescribeAddonsRequest) SetClusterType(v string) *DescribeAddonsRequest {
	s.ClusterType = &v
	return s
}

func (s *DescribeAddonsRequest) SetClusterVersion(v string) *DescribeAddonsRequest {
	s.ClusterVersion = &v
	return s
}

func (s *DescribeAddonsRequest) SetRegion(v string) *DescribeAddonsRequest {
	s.Region = &v
	return s
}

func (s *DescribeAddonsRequest) Validate() error {
	return dara.Validate(s)
}
