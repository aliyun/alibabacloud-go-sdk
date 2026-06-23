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
	// - `Default`: managed cluster.
	//
	// - `Serverless`: serverless cluster.
	//
	// - `Edge`: edge cluster.
	//
	// example:
	//
	// Default
	ClusterProfile *string `json:"cluster_profile,omitempty" xml:"cluster_profile,omitempty"`
	// After you set `cluster_type` to `ManagedKubernetes` and configure `profile`, you can further specify the cluster specification.
	//
	// - `ack.pro.small`: Pro cluster.
	//
	// - `ack.standard`: Basic cluster (selected by default if this parameter is left empty).
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// - `Kubernetes`: ACK dedicated cluster.
	//
	// - `ManagedKubernetes`: ACK managed cluster types, including ACK managed clusters (ACK Pro and ACK Basic), ACK Serverless clusters (Pro and Basic), ACK Edge clusters (Pro and Basic), and ACK Lingjun clusters (Pro).
	//
	// - `ExternalKubernetes`: registered cluster.
	//
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// The cluster version.
	//
	// example:
	//
	// 1.24.6-aliyun.1
	ClusterVersion *string `json:"cluster_version,omitempty" xml:"cluster_version,omitempty"`
	// The ID of the region where the cluster resides.
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
