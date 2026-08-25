// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCacheSize(v string) *ModifyDBClusterRequest
	GetCacheSize() *string
	SetClusterNodeCount(v int32) *ModifyDBClusterRequest
	GetClusterNodeCount() *int32
	SetClusterNodeType(v string) *ModifyDBClusterRequest
	GetClusterNodeType() *string
	SetDBClusterClass(v string) *ModifyDBClusterRequest
	GetDBClusterClass() *string
	SetDBClusterId(v string) *ModifyDBClusterRequest
	GetDBClusterId() *string
	SetDBInstanceId(v string) *ModifyDBClusterRequest
	GetDBInstanceId() *string
	SetEngine(v string) *ModifyDBClusterRequest
	GetEngine() *string
	SetRegionId(v string) *ModifyDBClusterRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *ModifyDBClusterRequest
	GetResourceOwnerId() *int64
	SetScaleMax(v float64) *ModifyDBClusterRequest
	GetScaleMax() *float64
	SetScaleMin(v float64) *ModifyDBClusterRequest
	GetScaleMin() *float64
}

type ModifyDBClusterRequest struct {
	// The reserved cache size.
	//
	// example:
	//
	// 200
	CacheSize        *string `json:"CacheSize,omitempty" xml:"CacheSize,omitempty"`
	ClusterNodeCount *int32  `json:"ClusterNodeCount,omitempty" xml:"ClusterNodeCount,omitempty"`
	ClusterNodeType  *string `json:"ClusterNodeType,omitempty" xml:"ClusterNodeType,omitempty"`
	// The cluster specifications. Valid values:
	//
	// - **selectdb.xlarge**: 4 cores, 32 GB memory.
	//
	// - **selectdb.2xlarge**: 8 cores, 64 GB memory.
	//
	// - **selectdb.4xlarge**: 16 cores, 128 GB memory.
	//
	// - **selectdb.8xlarge**: 32 cores, 256 GB memory.
	//
	// - **selectdb.16xlarge**: 16 cores, 512 GB memory.
	//
	// - **selectdb.24xlarge**: 24 cores, 768 GB memory.
	//
	// - **selectdb.32xlarge**: 32 cores, 1024 GB memory.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb.2xlarge
	DBClusterClass *string `json:"DBClusterClass,omitempty" xml:"DBClusterClass,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-xxxb9f2w-be
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database type. Set the value to selectdb.
	//
	// example:
	//
	// selectdb
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64   `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ScaleMax        *float64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	ScaleMin        *float64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
}

func (s ModifyDBClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterRequest) GetCacheSize() *string {
	return s.CacheSize
}

func (s *ModifyDBClusterRequest) GetClusterNodeCount() *int32 {
	return s.ClusterNodeCount
}

func (s *ModifyDBClusterRequest) GetClusterNodeType() *string {
	return s.ClusterNodeType
}

func (s *ModifyDBClusterRequest) GetDBClusterClass() *string {
	return s.DBClusterClass
}

func (s *ModifyDBClusterRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBClusterRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBClusterRequest) GetEngine() *string {
	return s.Engine
}

func (s *ModifyDBClusterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBClusterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBClusterRequest) GetScaleMax() *float64 {
	return s.ScaleMax
}

func (s *ModifyDBClusterRequest) GetScaleMin() *float64 {
	return s.ScaleMin
}

func (s *ModifyDBClusterRequest) SetCacheSize(v string) *ModifyDBClusterRequest {
	s.CacheSize = &v
	return s
}

func (s *ModifyDBClusterRequest) SetClusterNodeCount(v int32) *ModifyDBClusterRequest {
	s.ClusterNodeCount = &v
	return s
}

func (s *ModifyDBClusterRequest) SetClusterNodeType(v string) *ModifyDBClusterRequest {
	s.ClusterNodeType = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterClass(v string) *ModifyDBClusterRequest {
	s.DBClusterClass = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBClusterId(v string) *ModifyDBClusterRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetDBInstanceId(v string) *ModifyDBClusterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetEngine(v string) *ModifyDBClusterRequest {
	s.Engine = &v
	return s
}

func (s *ModifyDBClusterRequest) SetRegionId(v string) *ModifyDBClusterRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetResourceOwnerId(v int64) *ModifyDBClusterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBClusterRequest) SetScaleMax(v float64) *ModifyDBClusterRequest {
	s.ScaleMax = &v
	return s
}

func (s *ModifyDBClusterRequest) SetScaleMin(v float64) *ModifyDBClusterRequest {
	s.ScaleMin = &v
	return s
}

func (s *ModifyDBClusterRequest) Validate() error {
	return dara.Validate(s)
}
