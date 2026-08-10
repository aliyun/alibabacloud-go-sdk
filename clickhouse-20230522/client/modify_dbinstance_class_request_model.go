// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceClassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoScaleConfig(v *ModifyDBInstanceClassRequestAutoScaleConfig) *ModifyDBInstanceClassRequest
	GetAutoScaleConfig() *ModifyDBInstanceClassRequestAutoScaleConfig
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
	// The automatic horizontal scaling configuration.
	//
	// if can be null:
	// true
	AutoScaleConfig *ModifyDBInstanceClassRequestAutoScaleConfig `json:"AutoScaleConfig,omitempty" xml:"AutoScaleConfig,omitempty" type:"Struct"`
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
	// The maximum value for serverless node elastic scaling. Valid values: 4 to 32. The value must be greater than the minimum value.
	//
	// example:
	//
	// 4
	NodeScaleMax *int32 `json:"NodeScaleMax,omitempty" xml:"NodeScaleMax,omitempty"`
	// The minimum value for serverless node elastic scaling. Valid values: 4 to 32.
	//
	// example:
	//
	// 32
	NodeScaleMin *int32 `json:"NodeScaleMin,omitempty" xml:"NodeScaleMin,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The maximum value for serverless elastic scaling.
	//
	// example:
	//
	// 32
	ScaleMax *int64 `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// The minimum value for serverless elastic scaling.
	//
	// example:
	//
	// 8
	ScaleMin *int64 `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
	// The pre-purchased storage quota, in GB.
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

func (s ModifyDBInstanceClassRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassRequest) GetAutoScaleConfig() *ModifyDBInstanceClassRequestAutoScaleConfig {
	return s.AutoScaleConfig
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

func (s *ModifyDBInstanceClassRequest) SetAutoScaleConfig(v *ModifyDBInstanceClassRequestAutoScaleConfig) *ModifyDBInstanceClassRequest {
	s.AutoScaleConfig = v
	return s
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
	if s.AutoScaleConfig != nil {
		if err := s.AutoScaleConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDBInstanceClassRequestAutoScaleConfig struct {
	// The number of nodes available for burstable horizontal scaling.
	//
	// example:
	//
	// 2
	BurstNum *int32 `json:"BurstNum,omitempty" xml:"BurstNum,omitempty"`
	// The configuration status. Valid values:
	//
	// - disable: disabled.
	//
	// - enable: enabled.
	//
	// example:
	//
	// enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch information.
	VSwitchInfos []*ModifyDBInstanceClassRequestAutoScaleConfigVSwitchInfos `json:"VSwitchInfos,omitempty" xml:"VSwitchInfos,omitempty" type:"Repeated"`
}

func (s ModifyDBInstanceClassRequestAutoScaleConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceClassRequestAutoScaleConfig) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassRequestAutoScaleConfig) GetBurstNum() *int32 {
	return s.BurstNum
}

func (s *ModifyDBInstanceClassRequestAutoScaleConfig) GetStatus() *string {
	return s.Status
}

func (s *ModifyDBInstanceClassRequestAutoScaleConfig) GetVSwitchInfos() []*ModifyDBInstanceClassRequestAutoScaleConfigVSwitchInfos {
	return s.VSwitchInfos
}

func (s *ModifyDBInstanceClassRequestAutoScaleConfig) SetBurstNum(v int32) *ModifyDBInstanceClassRequestAutoScaleConfig {
	s.BurstNum = &v
	return s
}

func (s *ModifyDBInstanceClassRequestAutoScaleConfig) SetStatus(v string) *ModifyDBInstanceClassRequestAutoScaleConfig {
	s.Status = &v
	return s
}

func (s *ModifyDBInstanceClassRequestAutoScaleConfig) SetVSwitchInfos(v []*ModifyDBInstanceClassRequestAutoScaleConfigVSwitchInfos) *ModifyDBInstanceClassRequestAutoScaleConfig {
	s.VSwitchInfos = v
	return s
}

func (s *ModifyDBInstanceClassRequestAutoScaleConfig) Validate() error {
	if s.VSwitchInfos != nil {
		for _, item := range s.VSwitchInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDBInstanceClassRequestAutoScaleConfigVSwitchInfos struct {
	// The vSwitch IDs in the zone.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The zone of the hot pool.
	//
	// example:
	//
	// cn-beijing-XXX
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyDBInstanceClassRequestAutoScaleConfigVSwitchInfos) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceClassRequestAutoScaleConfigVSwitchInfos) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceClassRequestAutoScaleConfigVSwitchInfos) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *ModifyDBInstanceClassRequestAutoScaleConfigVSwitchInfos) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyDBInstanceClassRequestAutoScaleConfigVSwitchInfos) SetVSwitchIds(v []*string) *ModifyDBInstanceClassRequestAutoScaleConfigVSwitchInfos {
	s.VSwitchIds = v
	return s
}

func (s *ModifyDBInstanceClassRequestAutoScaleConfigVSwitchInfos) SetZoneId(v string) *ModifyDBInstanceClassRequestAutoScaleConfigVSwitchInfos {
	s.ZoneId = &v
	return s
}

func (s *ModifyDBInstanceClassRequestAutoScaleConfigVSwitchInfos) Validate() error {
	return dara.Validate(s)
}
