// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVirtualClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActiveClusterId(v string) *CreateVirtualClusterRequest
	GetActiveClusterId() *string
	SetClusterName(v string) *CreateVirtualClusterRequest
	GetClusterName() *string
	SetDBInstanceId(v string) *CreateVirtualClusterRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *CreateVirtualClusterRequest
	GetRegionId() *string
	SetStandbyClusterId(v string) *CreateVirtualClusterRequest
	GetStandbyClusterId() *string
}

type CreateVirtualClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	//
	// -be
	ActiveClusterId *string `json:"ActiveClusterId,omitempty" xml:"ActiveClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vcg_demo
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-26a3cjv****
	//
	// -be
	StandbyClusterId *string `json:"StandbyClusterId,omitempty" xml:"StandbyClusterId,omitempty"`
}

func (s CreateVirtualClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVirtualClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateVirtualClusterRequest) GetActiveClusterId() *string {
	return s.ActiveClusterId
}

func (s *CreateVirtualClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateVirtualClusterRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateVirtualClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateVirtualClusterRequest) GetStandbyClusterId() *string {
	return s.StandbyClusterId
}

func (s *CreateVirtualClusterRequest) SetActiveClusterId(v string) *CreateVirtualClusterRequest {
	s.ActiveClusterId = &v
	return s
}

func (s *CreateVirtualClusterRequest) SetClusterName(v string) *CreateVirtualClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateVirtualClusterRequest) SetDBInstanceId(v string) *CreateVirtualClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateVirtualClusterRequest) SetRegionId(v string) *CreateVirtualClusterRequest {
	s.RegionId = &v
	return s
}

func (s *CreateVirtualClusterRequest) SetStandbyClusterId(v string) *CreateVirtualClusterRequest {
	s.StandbyClusterId = &v
	return s
}

func (s *CreateVirtualClusterRequest) Validate() error {
	return dara.Validate(s)
}
