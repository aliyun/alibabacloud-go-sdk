// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddonRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeAddonRequest
	GetClusterId() *string
	SetClusterSpec(v string) *DescribeAddonRequest
	GetClusterSpec() *string
	SetClusterType(v string) *DescribeAddonRequest
	GetClusterType() *string
	SetClusterVersion(v string) *DescribeAddonRequest
	GetClusterVersion() *string
	SetProfile(v string) *DescribeAddonRequest
	GetProfile() *string
	SetRegionId(v string) *DescribeAddonRequest
	GetRegionId() *string
	SetVersion(v string) *DescribeAddonRequest
	GetVersion() *string
}

type DescribeAddonRequest struct {
	// Cluster ID.
	//
	// When a cluster ID is specified, the list of available components for the cluster is queried, and other parameters are ignored.
	//
	// example:
	//
	// c02b3e03be10643e8a644a843ffcb****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// Cluster specification.
	//
	// If cluster_id is specified, this parameter is ignored.
	//
	// The five parameters region_id, cluster_type, profile, cluster_spec, and cluster_version must be specified together.
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// Cluster type.
	//
	// If cluster_id is specified, this parameter is ignored.
	//
	// The five parameters region_id, cluster_type, profile, cluster_spec, and cluster_version must be specified together.
	//
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// Cluster version.
	//
	// If cluster_id is specified, this parameter is ignored.
	//
	// The five parameters region_id, cluster_type, profile, cluster_spec, and cluster_version must be specified together.
	//
	// example:
	//
	// 1.26.3-aliyun.1
	ClusterVersion *string `json:"cluster_version,omitempty" xml:"cluster_version,omitempty"`
	// Cluster subtype.
	//
	// If cluster_id is specified, this parameter is ignored.
	//
	// The five parameters region_id, cluster_type, profile, cluster_spec, and cluster_version must be specified together.
	//
	// example:
	//
	// Default
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// Region.
	//
	// If cluster_id is specified, this parameter is ignored.
	//
	// The five parameters region_id, cluster_type, profile, cluster_spec, and cluster_version must be specified together.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// Component version. If not specified, the latest available version of the component is queried.
	//
	// example:
	//
	// v1.9.3.10-7dfca203-aliyun
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeAddonRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddonRequest) GoString() string {
	return s.String()
}

func (s *DescribeAddonRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeAddonRequest) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *DescribeAddonRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeAddonRequest) GetClusterVersion() *string {
	return s.ClusterVersion
}

func (s *DescribeAddonRequest) GetProfile() *string {
	return s.Profile
}

func (s *DescribeAddonRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAddonRequest) GetVersion() *string {
	return s.Version
}

func (s *DescribeAddonRequest) SetClusterId(v string) *DescribeAddonRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeAddonRequest) SetClusterSpec(v string) *DescribeAddonRequest {
	s.ClusterSpec = &v
	return s
}

func (s *DescribeAddonRequest) SetClusterType(v string) *DescribeAddonRequest {
	s.ClusterType = &v
	return s
}

func (s *DescribeAddonRequest) SetClusterVersion(v string) *DescribeAddonRequest {
	s.ClusterVersion = &v
	return s
}

func (s *DescribeAddonRequest) SetProfile(v string) *DescribeAddonRequest {
	s.Profile = &v
	return s
}

func (s *DescribeAddonRequest) SetRegionId(v string) *DescribeAddonRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAddonRequest) SetVersion(v string) *DescribeAddonRequest {
	s.Version = &v
	return s
}

func (s *DescribeAddonRequest) Validate() error {
	return dara.Validate(s)
}
