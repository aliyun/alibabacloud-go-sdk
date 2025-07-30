// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBClusterBindingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteDBClusterBindingRequest
	GetDBClusterId() *string
	SetDBClusterIdBak(v string) *DeleteDBClusterBindingRequest
	GetDBClusterIdBak() *string
	SetDBInstanceId(v string) *DeleteDBClusterBindingRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DeleteDBClusterBindingRequest
	GetRegionId() *string
}

type DeleteDBClusterBindingRequest struct {
	// The ID of Cluster 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv2ez-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The ID of Cluster 2.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-xxxb9f2w-be
	DBClusterIdBak *string `json:"DBClusterIdBak,omitempty" xml:"DBClusterIdBak,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv2ez
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDBClusterBindingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBClusterBindingRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBClusterBindingRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDBClusterBindingRequest) GetDBClusterIdBak() *string {
	return s.DBClusterIdBak
}

func (s *DeleteDBClusterBindingRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBClusterBindingRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDBClusterBindingRequest) SetDBClusterId(v string) *DeleteDBClusterBindingRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBClusterBindingRequest) SetDBClusterIdBak(v string) *DeleteDBClusterBindingRequest {
	s.DBClusterIdBak = &v
	return s
}

func (s *DeleteDBClusterBindingRequest) SetDBInstanceId(v string) *DeleteDBClusterBindingRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBClusterBindingRequest) SetRegionId(v string) *DeleteDBClusterBindingRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDBClusterBindingRequest) Validate() error {
	return dara.Validate(s)
}
