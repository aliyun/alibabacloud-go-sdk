// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVirtualClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActiveClusterId(v string) *ModifyVirtualClusterRequest
	GetActiveClusterId() *string
	SetDBClusterId(v string) *ModifyVirtualClusterRequest
	GetDBClusterId() *string
	SetDBInstanceId(v string) *ModifyVirtualClusterRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *ModifyVirtualClusterRequest
	GetRegionId() *string
	SetStandbyClusterId(v string) *ModifyVirtualClusterRequest
	GetStandbyClusterId() *string
}

type ModifyVirtualClusterRequest struct {
	// example:
	//
	// selectdb-o2yg***-be
	ActiveClusterId *string `json:"ActiveClusterId,omitempty" xml:"ActiveClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// selectdb-vcg-b****-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
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
	// example:
	//
	// selectdb-pu6y****-be
	StandbyClusterId *string `json:"StandbyClusterId,omitempty" xml:"StandbyClusterId,omitempty"`
}

func (s ModifyVirtualClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVirtualClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyVirtualClusterRequest) GetActiveClusterId() *string {
	return s.ActiveClusterId
}

func (s *ModifyVirtualClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyVirtualClusterRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyVirtualClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyVirtualClusterRequest) GetStandbyClusterId() *string {
	return s.StandbyClusterId
}

func (s *ModifyVirtualClusterRequest) SetActiveClusterId(v string) *ModifyVirtualClusterRequest {
	s.ActiveClusterId = &v
	return s
}

func (s *ModifyVirtualClusterRequest) SetDBClusterId(v string) *ModifyVirtualClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyVirtualClusterRequest) SetDBInstanceId(v string) *ModifyVirtualClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyVirtualClusterRequest) SetRegionId(v string) *ModifyVirtualClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyVirtualClusterRequest) SetStandbyClusterId(v string) *ModifyVirtualClusterRequest {
	s.StandbyClusterId = &v
	return s
}

func (s *ModifyVirtualClusterRequest) Validate() error {
	return dara.Validate(s)
}
