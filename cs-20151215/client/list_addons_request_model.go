// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListAddonsRequest
	GetClusterId() *string
	SetClusterSpec(v string) *ListAddonsRequest
	GetClusterSpec() *string
	SetClusterType(v string) *ListAddonsRequest
	GetClusterType() *string
	SetClusterVersion(v string) *ListAddonsRequest
	GetClusterVersion() *string
	SetProfile(v string) *ListAddonsRequest
	GetProfile() *string
	SetRegionId(v string) *ListAddonsRequest
	GetRegionId() *string
}

type ListAddonsRequest struct {
	// The cluster ID.
	//
	// If you specify a cluster ID, the system queries the list of available components for the specified cluster, and all other parameters are ignored.
	//
	// example:
	//
	// c02b3e03be10643e8a644a843ffcb****
	ClusterId *string `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	// The cluster specification.
	//
	// If cluster_id is specified, this parameter is ignored.
	//
	// The region_id, cluster_type, profile, cluster_spec, and cluster_version parameters must be specified together.
	//
	// example:
	//
	// ack.pro.small
	ClusterSpec *string `json:"cluster_spec,omitempty" xml:"cluster_spec,omitempty"`
	// The cluster type.
	//
	// If cluster_id is specified, this parameter is ignored.
	//
	// The region_id, cluster_type, profile, cluster_spec, and cluster_version parameters must be specified together.
	//
	// example:
	//
	// ManagedKubernetes
	ClusterType *string `json:"cluster_type,omitempty" xml:"cluster_type,omitempty"`
	// The cluster version.
	//
	// If cluster_id is specified, this parameter is ignored.
	//
	// The region_id, cluster_type, profile, cluster_spec, and cluster_version parameters must be specified together.
	//
	// example:
	//
	// 1.32.1-aliyun.1
	ClusterVersion *string `json:"cluster_version,omitempty" xml:"cluster_version,omitempty"`
	// The cluster subtype.
	//
	// If cluster_id is specified, this parameter is ignored.
	//
	// The region_id, cluster_type, profile, cluster_spec, and cluster_version parameters must be specified together.
	//
	// example:
	//
	// Default
	Profile *string `json:"profile,omitempty" xml:"profile,omitempty"`
	// The region.
	//
	// If cluster_id is specified, this parameter is ignored.
	//
	// The region_id, cluster_type, profile, cluster_spec, and cluster_version parameters must be specified together.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"region_id,omitempty" xml:"region_id,omitempty"`
}

func (s ListAddonsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAddonsRequest) GoString() string {
	return s.String()
}

func (s *ListAddonsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListAddonsRequest) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *ListAddonsRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListAddonsRequest) GetClusterVersion() *string {
	return s.ClusterVersion
}

func (s *ListAddonsRequest) GetProfile() *string {
	return s.Profile
}

func (s *ListAddonsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAddonsRequest) SetClusterId(v string) *ListAddonsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAddonsRequest) SetClusterSpec(v string) *ListAddonsRequest {
	s.ClusterSpec = &v
	return s
}

func (s *ListAddonsRequest) SetClusterType(v string) *ListAddonsRequest {
	s.ClusterType = &v
	return s
}

func (s *ListAddonsRequest) SetClusterVersion(v string) *ListAddonsRequest {
	s.ClusterVersion = &v
	return s
}

func (s *ListAddonsRequest) SetProfile(v string) *ListAddonsRequest {
	s.Profile = &v
	return s
}

func (s *ListAddonsRequest) SetRegionId(v string) *ListAddonsRequest {
	s.RegionId = &v
	return s
}

func (s *ListAddonsRequest) Validate() error {
	return dara.Validate(s)
}
