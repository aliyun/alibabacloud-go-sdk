// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsHiveJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedConfig(v string) *CreateApsHiveJobRequest
	GetAdvancedConfig() *string
	SetConflictStrategy(v string) *CreateApsHiveJobRequest
	GetConflictStrategy() *string
	SetDBClusterId(v string) *CreateApsHiveJobRequest
	GetDBClusterId() *string
	SetDatasourceId(v int64) *CreateApsHiveJobRequest
	GetDatasourceId() *int64
	SetFullComputeUnit(v string) *CreateApsHiveJobRequest
	GetFullComputeUnit() *string
	SetOssLocation(v string) *CreateApsHiveJobRequest
	GetOssLocation() *string
	SetParallelism(v int32) *CreateApsHiveJobRequest
	GetParallelism() *int32
	SetRegionId(v string) *CreateApsHiveJobRequest
	GetRegionId() *string
	SetResourceGroup(v string) *CreateApsHiveJobRequest
	GetResourceGroup() *string
	SetSyncAllowExpression(v string) *CreateApsHiveJobRequest
	GetSyncAllowExpression() *string
	SetSyncDenyExpression(v string) *CreateApsHiveJobRequest
	GetSyncDenyExpression() *string
	SetTargetType(v string) *CreateApsHiveJobRequest
	GetTargetType() *string
	SetWorkloadName(v string) *CreateApsHiveJobRequest
	GetWorkloadName() *string
}

type CreateApsHiveJobRequest struct {
	// The advanced configurations.
	//
	// example:
	//
	// -
	AdvancedConfig *string `json:"AdvancedConfig,omitempty" xml:"AdvancedConfig,omitempty"`
	// The policy to handle tables with the same name in the destination cluster.
	//
	// example:
	//
	// Intercept: reports error and aborts.
	//
	// Ignore: ignores and continues migrating the relevant tables.
	//
	// Skip: skips related tables and only migrates other tables.
	ConflictStrategy *string `json:"ConflictStrategy,omitempty" xml:"ConflictStrategy,omitempty"`
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 40
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The number of AnalyticDB compute units (ACUs) required for data migration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 16
	FullComputeUnit *string `json:"FullComputeUnit,omitempty" xml:"FullComputeUnit,omitempty"`
	// The path of the destination data lakehouse in an Object Storage Service (OSS) bucket.
	//
	// This parameter is required.
	OssLocation *string `json:"OssLocation,omitempty" xml:"OssLocation,omitempty"`
	// The number of tasks that are allowed in parallel.
	//
	// example:
	//
	// 8
	Parallelism *int32 `json:"Parallelism,omitempty" xml:"Parallelism,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/143074.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_default
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The expression that allows objects to be synchronized.
	//
	// example:
	//
	// *
	SyncAllowExpression *string `json:"SyncAllowExpression,omitempty" xml:"SyncAllowExpression,omitempty"`
	// The expression that denies objects to be synchronized.
	//
	// example:
	//
	// abc
	SyncDenyExpression *string `json:"SyncDenyExpression,omitempty" xml:"SyncDenyExpression,omitempty"`
	// The destination type.
	//
	// example:
	//
	// OSS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The name of the workload.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx-20240224100253
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s CreateApsHiveJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApsHiveJobRequest) GoString() string {
	return s.String()
}

func (s *CreateApsHiveJobRequest) GetAdvancedConfig() *string {
	return s.AdvancedConfig
}

func (s *CreateApsHiveJobRequest) GetConflictStrategy() *string {
	return s.ConflictStrategy
}

func (s *CreateApsHiveJobRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApsHiveJobRequest) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *CreateApsHiveJobRequest) GetFullComputeUnit() *string {
	return s.FullComputeUnit
}

func (s *CreateApsHiveJobRequest) GetOssLocation() *string {
	return s.OssLocation
}

func (s *CreateApsHiveJobRequest) GetParallelism() *int32 {
	return s.Parallelism
}

func (s *CreateApsHiveJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApsHiveJobRequest) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *CreateApsHiveJobRequest) GetSyncAllowExpression() *string {
	return s.SyncAllowExpression
}

func (s *CreateApsHiveJobRequest) GetSyncDenyExpression() *string {
	return s.SyncDenyExpression
}

func (s *CreateApsHiveJobRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateApsHiveJobRequest) GetWorkloadName() *string {
	return s.WorkloadName
}

func (s *CreateApsHiveJobRequest) SetAdvancedConfig(v string) *CreateApsHiveJobRequest {
	s.AdvancedConfig = &v
	return s
}

func (s *CreateApsHiveJobRequest) SetConflictStrategy(v string) *CreateApsHiveJobRequest {
	s.ConflictStrategy = &v
	return s
}

func (s *CreateApsHiveJobRequest) SetDBClusterId(v string) *CreateApsHiveJobRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateApsHiveJobRequest) SetDatasourceId(v int64) *CreateApsHiveJobRequest {
	s.DatasourceId = &v
	return s
}

func (s *CreateApsHiveJobRequest) SetFullComputeUnit(v string) *CreateApsHiveJobRequest {
	s.FullComputeUnit = &v
	return s
}

func (s *CreateApsHiveJobRequest) SetOssLocation(v string) *CreateApsHiveJobRequest {
	s.OssLocation = &v
	return s
}

func (s *CreateApsHiveJobRequest) SetParallelism(v int32) *CreateApsHiveJobRequest {
	s.Parallelism = &v
	return s
}

func (s *CreateApsHiveJobRequest) SetRegionId(v string) *CreateApsHiveJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApsHiveJobRequest) SetResourceGroup(v string) *CreateApsHiveJobRequest {
	s.ResourceGroup = &v
	return s
}

func (s *CreateApsHiveJobRequest) SetSyncAllowExpression(v string) *CreateApsHiveJobRequest {
	s.SyncAllowExpression = &v
	return s
}

func (s *CreateApsHiveJobRequest) SetSyncDenyExpression(v string) *CreateApsHiveJobRequest {
	s.SyncDenyExpression = &v
	return s
}

func (s *CreateApsHiveJobRequest) SetTargetType(v string) *CreateApsHiveJobRequest {
	s.TargetType = &v
	return s
}

func (s *CreateApsHiveJobRequest) SetWorkloadName(v string) *CreateApsHiveJobRequest {
	s.WorkloadName = &v
	return s
}

func (s *CreateApsHiveJobRequest) Validate() error {
	return dara.Validate(s)
}
