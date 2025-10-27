// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceClassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComputingGroupId(v string) *ModifyDBInstanceClassRequest
	GetComputingGroupId() *string
	SetDBInstanceId(v string) *ModifyDBInstanceClassRequest
	GetDBInstanceId() *string
	SetNodeCount(v int32) *ModifyDBInstanceClassRequest
	GetNodeCount() *int32
	SetNodeScaleMax(v int32) *ModifyDBInstanceClassRequest
	GetNodeScaleMax() *int32
	SetNodeScaleMin(v int32) *ModifyDBInstanceClassRequest
	GetNodeScaleMin() *int32
	SetRegionId(v string) *ModifyDBInstanceClassRequest
	GetRegionId() *string
	SetScaleMax(v int64) *ModifyDBInstanceClassRequest
	GetScaleMax() *int64
	SetScaleMin(v int64) *ModifyDBInstanceClassRequest
	GetScaleMin() *int64
	SetStorageQuota(v int64) *ModifyDBInstanceClassRequest
	GetStorageQuota() *int64
	SetStorageType(v string) *ModifyDBInstanceClassRequest
	GetStorageType() *string
}

type ModifyDBInstanceClassRequest struct {
	ComputingGroupId *string `json:"ComputingGroupId,omitempty" xml:"ComputingGroupId,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// example:
	//
	// 2
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// example:
	//
	// 4
	NodeScaleMax *int32 `json:"NodeScaleMax,omitempty" xml:"NodeScaleMax,omitempty"`
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
	// The maximum capacity for elastic scaling.
	//
	// example:
	//
	// 32
	ScaleMax *int64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum capacity for elastic scaling.
	//
	// example:
	//
	// 2
	ScaleMin *int64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// example:
	//
	// 100
	StorageQuota *int64 `json:"StorageQuota,omitempty" xml:"StorageQuota,omitempty"`
	// example:
	//
	// OSS
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s ModifyDBInstanceClassRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassRequest) GetComputingGroupId() *string {
	return s.ComputingGroupId
}

func (s *ModifyDBInstanceClassRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceClassRequest) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *ModifyDBInstanceClassRequest) GetNodeScaleMax() *int32 {
	return s.NodeScaleMax
}

func (s *ModifyDBInstanceClassRequest) GetNodeScaleMin() *int32 {
	return s.NodeScaleMin
}

func (s *ModifyDBInstanceClassRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBInstanceClassRequest) GetScaleMax() *int64 {
	return s.ScaleMax
}

func (s *ModifyDBInstanceClassRequest) GetScaleMin() *int64 {
	return s.ScaleMin
}

func (s *ModifyDBInstanceClassRequest) GetStorageQuota() *int64 {
	return s.StorageQuota
}

func (s *ModifyDBInstanceClassRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *ModifyDBInstanceClassRequest) SetComputingGroupId(v string) *ModifyDBInstanceClassRequest {
	s.ComputingGroupId = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetDBInstanceId(v string) *ModifyDBInstanceClassRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetNodeCount(v int32) *ModifyDBInstanceClassRequest {
	s.NodeCount = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetNodeScaleMax(v int32) *ModifyDBInstanceClassRequest {
	s.NodeScaleMax = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetNodeScaleMin(v int32) *ModifyDBInstanceClassRequest {
	s.NodeScaleMin = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetRegionId(v string) *ModifyDBInstanceClassRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetScaleMax(v int64) *ModifyDBInstanceClassRequest {
	s.ScaleMax = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetScaleMin(v int64) *ModifyDBInstanceClassRequest {
	s.ScaleMin = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetStorageQuota(v int64) *ModifyDBInstanceClassRequest {
	s.StorageQuota = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) SetStorageType(v string) *ModifyDBInstanceClassRequest {
	s.StorageType = &v
	return s
}

func (s *ModifyDBInstanceClassRequest) Validate() error {
	return dara.Validate(s)
}
