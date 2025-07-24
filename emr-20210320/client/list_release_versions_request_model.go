// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReleaseVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterType(v string) *ListReleaseVersionsRequest
	GetClusterType() *string
	SetIaasType(v string) *ListReleaseVersionsRequest
	GetIaasType() *string
	SetRegionId(v string) *ListReleaseVersionsRequest
	GetRegionId() *string
}

type ListReleaseVersionsRequest struct {
	// The type of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// DATALAKE
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The type of the IaaS resource.
	//
	// example:
	//
	// ECS
	IaasType *string `json:"IaasType,omitempty" xml:"IaasType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListReleaseVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListReleaseVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListReleaseVersionsRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListReleaseVersionsRequest) GetIaasType() *string {
	return s.IaasType
}

func (s *ListReleaseVersionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListReleaseVersionsRequest) SetClusterType(v string) *ListReleaseVersionsRequest {
	s.ClusterType = &v
	return s
}

func (s *ListReleaseVersionsRequest) SetIaasType(v string) *ListReleaseVersionsRequest {
	s.IaasType = &v
	return s
}

func (s *ListReleaseVersionsRequest) SetRegionId(v string) *ListReleaseVersionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListReleaseVersionsRequest) Validate() error {
	return dara.Validate(s)
}
