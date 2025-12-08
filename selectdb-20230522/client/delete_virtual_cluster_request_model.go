// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteVirtualClusterRequest
	GetDBClusterId() *string
	SetDBInstanceId(v string) *DeleteVirtualClusterRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DeleteVirtualClusterRequest
	GetRegionId() *string
}

type DeleteVirtualClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-vcg-72vz***-be
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
}

func (s DeleteVirtualClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteVirtualClusterRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteVirtualClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteVirtualClusterRequest) SetDBClusterId(v string) *DeleteVirtualClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteVirtualClusterRequest) SetDBInstanceId(v string) *DeleteVirtualClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteVirtualClusterRequest) SetRegionId(v string) *DeleteVirtualClusterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteVirtualClusterRequest) Validate() error {
	return dara.Validate(s)
}
