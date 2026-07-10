// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceClassShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoScaleConfigShrink(v string) *ModifyDBInstanceClassShrinkRequest
	GetAutoScaleConfigShrink() *string
	SetComputingGroupId(v string) *ModifyDBInstanceClassShrinkRequest
	GetComputingGroupId() *string
	SetDBInstanceId(v string) *ModifyDBInstanceClassShrinkRequest
	GetDBInstanceId() *string
	SetNodeCount(v int32) *ModifyDBInstanceClassShrinkRequest
	GetNodeCount() *int32
	SetNodeScaleMax(v int32) *ModifyDBInstanceClassShrinkRequest
	GetNodeScaleMax() *int32
	SetNodeScaleMin(v int32) *ModifyDBInstanceClassShrinkRequest
	GetNodeScaleMin() *int32
	SetRegionId(v string) *ModifyDBInstanceClassShrinkRequest
	GetRegionId() *string
	SetScaleMax(v int64) *ModifyDBInstanceClassShrinkRequest
	GetScaleMax() *int64
	SetScaleMin(v int64) *ModifyDBInstanceClassShrinkRequest
	GetScaleMin() *int64
	SetStorageQuota(v int64) *ModifyDBInstanceClassShrinkRequest
	GetStorageQuota() *int64
	SetStorageType(v string) *ModifyDBInstanceClassShrinkRequest
	GetStorageType() *string
}

type ModifyDBInstanceClassShrinkRequest struct {
	// The autoscaling configuration for the compute group.
	//
	// if can be null:
	// true
	AutoScaleConfigShrink *string `json:"AutoScaleConfig,omitempty" xml:"AutoScaleConfig,omitempty"`
	// The compute group ID.
	//
	// example:
	//
	// cc-gs5j3sua77*******-clickhouse
	ComputingGroupId *string `json:"ComputingGroupId,omitempty" xml:"ComputingGroupId,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-gs5j3sua77*******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The number of nodes. Valid values: 2 to 16.
	//
	// example:
	//
	// 2
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The maximum capacity per node for serverless autoscaling. Valid values: 4 to 32. This value must be greater than the minimum value.
	//
	// example:
	//
	// 4
	NodeScaleMax *int32 `json:"NodeScaleMax,omitempty" xml:"NodeScaleMax,omitempty"`
	// The minimum capacity per node for serverless autoscaling. Valid values: 4 to 32.
	//
	// example:
	//
	// 32
	NodeScaleMin *int32 `json:"NodeScaleMin,omitempty" xml:"NodeScaleMin,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The maximum capacity for serverless autoscaling.
	//
	// example:
	//
	// 32
	ScaleMax *int64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum capacity for serverless autoscaling.
	//
	// example:
	//
	// 8
	ScaleMin *int64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The pre-purchased storage capacity in GB.
	//
	// example:
	//
	// 100
	StorageQuota *int64 `json:"StorageQuota,omitempty" xml:"StorageQuota,omitempty"`
	// The storage type.
	//
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s ModifyDBInstanceClassShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceClassShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassShrinkRequest) GetAutoScaleConfigShrink() *string {
	return s.AutoScaleConfigShrink
}

func (s *ModifyDBInstanceClassShrinkRequest) GetComputingGroupId() *string {
	return s.ComputingGroupId
}

func (s *ModifyDBInstanceClassShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceClassShrinkRequest) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *ModifyDBInstanceClassShrinkRequest) GetNodeScaleMax() *int32 {
	return s.NodeScaleMax
}

func (s *ModifyDBInstanceClassShrinkRequest) GetNodeScaleMin() *int32 {
	return s.NodeScaleMin
}

func (s *ModifyDBInstanceClassShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBInstanceClassShrinkRequest) GetScaleMax() *int64 {
	return s.ScaleMax
}

func (s *ModifyDBInstanceClassShrinkRequest) GetScaleMin() *int64 {
	return s.ScaleMin
}

func (s *ModifyDBInstanceClassShrinkRequest) GetStorageQuota() *int64 {
	return s.StorageQuota
}

func (s *ModifyDBInstanceClassShrinkRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *ModifyDBInstanceClassShrinkRequest) SetAutoScaleConfigShrink(v string) *ModifyDBInstanceClassShrinkRequest {
	s.AutoScaleConfigShrink = &v
	return s
}

func (s *ModifyDBInstanceClassShrinkRequest) SetComputingGroupId(v string) *ModifyDBInstanceClassShrinkRequest {
	s.ComputingGroupId = &v
	return s
}

func (s *ModifyDBInstanceClassShrinkRequest) SetDBInstanceId(v string) *ModifyDBInstanceClassShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceClassShrinkRequest) SetNodeCount(v int32) *ModifyDBInstanceClassShrinkRequest {
	s.NodeCount = &v
	return s
}

func (s *ModifyDBInstanceClassShrinkRequest) SetNodeScaleMax(v int32) *ModifyDBInstanceClassShrinkRequest {
	s.NodeScaleMax = &v
	return s
}

func (s *ModifyDBInstanceClassShrinkRequest) SetNodeScaleMin(v int32) *ModifyDBInstanceClassShrinkRequest {
	s.NodeScaleMin = &v
	return s
}

func (s *ModifyDBInstanceClassShrinkRequest) SetRegionId(v string) *ModifyDBInstanceClassShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceClassShrinkRequest) SetScaleMax(v int64) *ModifyDBInstanceClassShrinkRequest {
	s.ScaleMax = &v
	return s
}

func (s *ModifyDBInstanceClassShrinkRequest) SetScaleMin(v int64) *ModifyDBInstanceClassShrinkRequest {
	s.ScaleMin = &v
	return s
}

func (s *ModifyDBInstanceClassShrinkRequest) SetStorageQuota(v int64) *ModifyDBInstanceClassShrinkRequest {
	s.StorageQuota = &v
	return s
}

func (s *ModifyDBInstanceClassShrinkRequest) SetStorageType(v string) *ModifyDBInstanceClassShrinkRequest {
	s.StorageType = &v
	return s
}

func (s *ModifyDBInstanceClassShrinkRequest) Validate() error {
	return dara.Validate(s)
}
