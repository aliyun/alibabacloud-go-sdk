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
	// The cluster ID. If you specify a cluster ID, only components used in the specified cluster are queried. Other parameters are ignored.
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
	// The region of the cluster. If cluster_id is specified, this parameter is ignored. You must specify the region_id, cluster_type, profile, cluster_spec, and cluster_version parameters at the same time.
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
