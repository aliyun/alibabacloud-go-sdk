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
	// The ID of the cluster. If you specify a cluster ID, only components used in the cluster are queried. Other parameters are ignored.
	//
	// example:
	//
	// c02b3e03be10643e8a644a843ffcb****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The specifications of the cluster. If cluster_id is specified, this parameter is ignored. You must specify the region_id, cluster_type, profile, cluster_spec, and cluster_version parameters at the same time.
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// The type of the cluster. If cluster_id is specified, this parameter is ignored. You must specify the region_id, cluster_type, profile, cluster_spec, and cluster_version parameters at the same time.
	//
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// The version of the cluster. If cluster_id is specified, this parameter is ignored. You must specify the region_id, cluster_type, profile, cluster_spec, and cluster_version parameters at the same time.
	//
	// example:
	//
	// 1.26.3-aliyun.1
	ClusterVersion *string `json:"cluster_version,omitempty" xml:"cluster_version,omitempty"`
	// The subtype of the cluster. If cluster_id is specified, this parameter is ignored. You must specify the region_id, cluster_type, profile, cluster_spec, and cluster_version parameters at the same time.
	//
	// example:
	//
	// Default
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// The region ID. If cluster_id is specified, this parameter is ignored. You must specify the region_id, cluster_type, profile, cluster_spec, and cluster_version parameters at the same time.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
	// The version of the component. If you do not specify this parameter, the latest version of the component is queried.
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
