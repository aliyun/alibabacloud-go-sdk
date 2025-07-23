// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *CreateClusterRequest
	GetClusterName() *string
	SetMasterInstanceId(v string) *CreateClusterRequest
	GetMasterInstanceId() *string
	SetRegionId(v string) *CreateClusterRequest
	GetRegionId() *string
}

type CreateClusterRequest struct {
	// The name of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cluster_on****
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The ID of the master HSM.
	//
	// This parameter is required.
	//
	// example:
	//
	// hsm_intl-sg-uz63ixak****
	MasterInstanceId *string `json:"MasterInstanceId,omitempty" xml:"MasterInstanceId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateClusterRequest) GetMasterInstanceId() *string {
	return s.MasterInstanceId
}

func (s *CreateClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetMasterInstanceId(v string) *CreateClusterRequest {
	s.MasterInstanceId = &v
	return s
}

func (s *CreateClusterRequest) SetRegionId(v string) *CreateClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateClusterRequest) Validate() error {
	return dara.Validate(s)
}
