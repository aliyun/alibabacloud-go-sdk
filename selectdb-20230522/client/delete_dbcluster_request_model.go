// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteDBClusterRequest
	GetDBClusterId() *string
	SetDBInstanceId(v string) *DeleteDBClusterRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DeleteDBClusterRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteDBClusterRequest
	GetResourceGroupId() *string
	SetResourceOwnerId(v int64) *DeleteDBClusterRequest
	GetResourceOwnerId() *int64
}

type DeleteDBClusterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-xxxb9f2w-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// 代表资源一级ID的资源属性字段
	//
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
	// 代表资源组的资源属性字段
	//
	// example:
	//
	// rg-4690g37929****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDBClusterRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDBClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteDBClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBClusterRequest) SetDBClusterId(v string) *DeleteDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetDBInstanceId(v string) *DeleteDBClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetRegionId(v string) *DeleteDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetResourceGroupId(v string) *DeleteDBClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteDBClusterRequest) SetResourceOwnerId(v int64) *DeleteDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
