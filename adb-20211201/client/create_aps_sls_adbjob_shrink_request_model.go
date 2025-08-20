// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsSlsADBJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrossRole(v string) *CreateApsSlsADBJobShrinkRequest
	GetAcrossRole() *string
	SetAcrossUid(v string) *CreateApsSlsADBJobShrinkRequest
	GetAcrossUid() *string
	SetAdvancedConfig(v string) *CreateApsSlsADBJobShrinkRequest
	GetAdvancedConfig() *string
	SetColumnsShrink(v string) *CreateApsSlsADBJobShrinkRequest
	GetColumnsShrink() *string
	SetDBClusterId(v string) *CreateApsSlsADBJobShrinkRequest
	GetDBClusterId() *string
	SetDatasourceId(v int64) *CreateApsSlsADBJobShrinkRequest
	GetDatasourceId() *int64
	SetDbName(v string) *CreateApsSlsADBJobShrinkRequest
	GetDbName() *string
	SetDirtyDataHandleMode(v string) *CreateApsSlsADBJobShrinkRequest
	GetDirtyDataHandleMode() *string
	SetDirtyDataProcessPattern(v string) *CreateApsSlsADBJobShrinkRequest
	GetDirtyDataProcessPattern() *string
	SetExactlyOnce(v string) *CreateApsSlsADBJobShrinkRequest
	GetExactlyOnce() *string
	SetFullComputeUnit(v string) *CreateApsSlsADBJobShrinkRequest
	GetFullComputeUnit() *string
	SetHudiAdvancedConfig(v string) *CreateApsSlsADBJobShrinkRequest
	GetHudiAdvancedConfig() *string
	SetIncrementalComputeUnit(v string) *CreateApsSlsADBJobShrinkRequest
	GetIncrementalComputeUnit() *string
	SetLakehouseId(v int64) *CreateApsSlsADBJobShrinkRequest
	GetLakehouseId() *int64
	SetMaxOffsetsPerTrigger(v int64) *CreateApsSlsADBJobShrinkRequest
	GetMaxOffsetsPerTrigger() *int64
	SetOssLocation(v string) *CreateApsSlsADBJobShrinkRequest
	GetOssLocation() *string
	SetOutputFormat(v string) *CreateApsSlsADBJobShrinkRequest
	GetOutputFormat() *string
	SetPartitionSpecsShrink(v string) *CreateApsSlsADBJobShrinkRequest
	GetPartitionSpecsShrink() *string
	SetPassword(v string) *CreateApsSlsADBJobShrinkRequest
	GetPassword() *string
	SetPrimaryKeyDefinition(v string) *CreateApsSlsADBJobShrinkRequest
	GetPrimaryKeyDefinition() *string
	SetProject(v string) *CreateApsSlsADBJobShrinkRequest
	GetProject() *string
	SetRegionId(v string) *CreateApsSlsADBJobShrinkRequest
	GetRegionId() *string
	SetResourceGroup(v string) *CreateApsSlsADBJobShrinkRequest
	GetResourceGroup() *string
	SetSourceRegionId(v string) *CreateApsSlsADBJobShrinkRequest
	GetSourceRegionId() *string
	SetStartingOffsets(v string) *CreateApsSlsADBJobShrinkRequest
	GetStartingOffsets() *string
	SetStore(v string) *CreateApsSlsADBJobShrinkRequest
	GetStore() *string
	SetTableName(v string) *CreateApsSlsADBJobShrinkRequest
	GetTableName() *string
	SetTargetGenerateRule(v string) *CreateApsSlsADBJobShrinkRequest
	GetTargetGenerateRule() *string
	SetTargetType(v string) *CreateApsSlsADBJobShrinkRequest
	GetTargetType() *string
	SetUnixTimestampConvertShrink(v string) *CreateApsSlsADBJobShrinkRequest
	GetUnixTimestampConvertShrink() *string
	SetUserName(v string) *CreateApsSlsADBJobShrinkRequest
	GetUserName() *string
	SetWorkloadName(v string) *CreateApsSlsADBJobShrinkRequest
	GetWorkloadName() *string
}

type CreateApsSlsADBJobShrinkRequest struct {
	// The name of the cross-account role.
	//
	// example:
	//
	// test-role
	AcrossRole *string `json:"AcrossRole,omitempty" xml:"AcrossRole,omitempty"`
	// The cross-account UID.
	//
	// example:
	//
	// 123456
	AcrossUid *string `json:"AcrossUid,omitempty" xml:"AcrossUid,omitempty"`
	// The advanced configurations.
	//
	// example:
	//
	// -
	AdvancedConfig *string `json:"AdvancedConfig,omitempty" xml:"AdvancedConfig,omitempty"`
	// The information about columns.
	//
	// This parameter is required.
	//
	// example:
	//
	// -
	ColumnsShrink *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-*********
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 327
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbName
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The dirty data processing mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// -
	DirtyDataHandleMode *string `json:"DirtyDataHandleMode,omitempty" xml:"DirtyDataHandleMode,omitempty"`
	// The dirty data processing mode.
	//
	// example:
	//
	// STOP
	DirtyDataProcessPattern *string `json:"DirtyDataProcessPattern,omitempty" xml:"DirtyDataProcessPattern,omitempty"`
	// Specifies whether to enable the consistency check.
	//
	// example:
	//
	// false
	ExactlyOnce *string `json:"ExactlyOnce,omitempty" xml:"ExactlyOnce,omitempty"`
	// The number of full AnalyticDB compute units (ACUs).
	//
	// example:
	//
	// 16
	FullComputeUnit *string `json:"FullComputeUnit,omitempty" xml:"FullComputeUnit,omitempty"`
	// The advanced configurations of Hudi.
	//
	// example:
	//
	// -
	HudiAdvancedConfig *string `json:"HudiAdvancedConfig,omitempty" xml:"HudiAdvancedConfig,omitempty"`
	// The number of increment ACUs.
	//
	// example:
	//
	// 168
	IncrementalComputeUnit *string `json:"IncrementalComputeUnit,omitempty" xml:"IncrementalComputeUnit,omitempty"`
	// The lakehouse ID.
	//
	// example:
	//
	// 123
	LakehouseId *int64 `json:"LakehouseId,omitempty" xml:"LakehouseId,omitempty"`
	// The latest offset.
	//
	// example:
	//
	// -
	MaxOffsetsPerTrigger *int64 `json:"MaxOffsetsPerTrigger,omitempty" xml:"MaxOffsetsPerTrigger,omitempty"`
	// The Object Storage Service (OSS) URL.
	//
	// example:
	//
	// oss://test*
	OssLocation *string `json:"OssLocation,omitempty" xml:"OssLocation,omitempty"`
	// The format of the output file.
	//
	// example:
	//
	// -
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// The information about partition.
	//
	// example:
	//
	// -
	PartitionSpecsShrink *string `json:"PartitionSpecs,omitempty" xml:"PartitionSpecs,omitempty"`
	// The password of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_user
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The definition of the primary key.
	//
	// example:
	//
	// -
	PrimaryKeyDefinition *string `json:"PrimaryKeyDefinition,omitempty" xml:"PrimaryKeyDefinition,omitempty"`
	// The name of the SLS project.
	//
	// example:
	//
	// test
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the resource group.
	//
	// example:
	//
	// test
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The region ID of the source cluster.
	//
	// example:
	//
	// cn-beijing
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// The start offset.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	StartingOffsets *string `json:"StartingOffsets,omitempty" xml:"StartingOffsets,omitempty"`
	// The SLS Logstore.
	//
	// example:
	//
	// test
	Store *string `json:"Store,omitempty" xml:"Store,omitempty"`
	// The name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The rules for generating the destination database.
	//
	// example:
	//
	// -
	TargetGenerateRule *string `json:"TargetGenerateRule,omitempty" xml:"TargetGenerateRule,omitempty"`
	// The destination type.
	//
	// example:
	//
	// ADB
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The timestamp conversion.
	//
	// example:
	//
	// -
	UnixTimestampConvertShrink *string `json:"UnixTimestampConvert,omitempty" xml:"UnixTimestampConvert,omitempty"`
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The name of the workload.
	//
	// This parameter is required.
	//
	// example:
	//
	// sls-******
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s CreateApsSlsADBJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApsSlsADBJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApsSlsADBJobShrinkRequest) GetAcrossRole() *string {
	return s.AcrossRole
}

func (s *CreateApsSlsADBJobShrinkRequest) GetAcrossUid() *string {
	return s.AcrossUid
}

func (s *CreateApsSlsADBJobShrinkRequest) GetAdvancedConfig() *string {
	return s.AdvancedConfig
}

func (s *CreateApsSlsADBJobShrinkRequest) GetColumnsShrink() *string {
	return s.ColumnsShrink
}

func (s *CreateApsSlsADBJobShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApsSlsADBJobShrinkRequest) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *CreateApsSlsADBJobShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateApsSlsADBJobShrinkRequest) GetDirtyDataHandleMode() *string {
	return s.DirtyDataHandleMode
}

func (s *CreateApsSlsADBJobShrinkRequest) GetDirtyDataProcessPattern() *string {
	return s.DirtyDataProcessPattern
}

func (s *CreateApsSlsADBJobShrinkRequest) GetExactlyOnce() *string {
	return s.ExactlyOnce
}

func (s *CreateApsSlsADBJobShrinkRequest) GetFullComputeUnit() *string {
	return s.FullComputeUnit
}

func (s *CreateApsSlsADBJobShrinkRequest) GetHudiAdvancedConfig() *string {
	return s.HudiAdvancedConfig
}

func (s *CreateApsSlsADBJobShrinkRequest) GetIncrementalComputeUnit() *string {
	return s.IncrementalComputeUnit
}

func (s *CreateApsSlsADBJobShrinkRequest) GetLakehouseId() *int64 {
	return s.LakehouseId
}

func (s *CreateApsSlsADBJobShrinkRequest) GetMaxOffsetsPerTrigger() *int64 {
	return s.MaxOffsetsPerTrigger
}

func (s *CreateApsSlsADBJobShrinkRequest) GetOssLocation() *string {
	return s.OssLocation
}

func (s *CreateApsSlsADBJobShrinkRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *CreateApsSlsADBJobShrinkRequest) GetPartitionSpecsShrink() *string {
	return s.PartitionSpecsShrink
}

func (s *CreateApsSlsADBJobShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateApsSlsADBJobShrinkRequest) GetPrimaryKeyDefinition() *string {
	return s.PrimaryKeyDefinition
}

func (s *CreateApsSlsADBJobShrinkRequest) GetProject() *string {
	return s.Project
}

func (s *CreateApsSlsADBJobShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApsSlsADBJobShrinkRequest) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *CreateApsSlsADBJobShrinkRequest) GetSourceRegionId() *string {
	return s.SourceRegionId
}

func (s *CreateApsSlsADBJobShrinkRequest) GetStartingOffsets() *string {
	return s.StartingOffsets
}

func (s *CreateApsSlsADBJobShrinkRequest) GetStore() *string {
	return s.Store
}

func (s *CreateApsSlsADBJobShrinkRequest) GetTableName() *string {
	return s.TableName
}

func (s *CreateApsSlsADBJobShrinkRequest) GetTargetGenerateRule() *string {
	return s.TargetGenerateRule
}

func (s *CreateApsSlsADBJobShrinkRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateApsSlsADBJobShrinkRequest) GetUnixTimestampConvertShrink() *string {
	return s.UnixTimestampConvertShrink
}

func (s *CreateApsSlsADBJobShrinkRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateApsSlsADBJobShrinkRequest) GetWorkloadName() *string {
	return s.WorkloadName
}

func (s *CreateApsSlsADBJobShrinkRequest) SetAcrossRole(v string) *CreateApsSlsADBJobShrinkRequest {
	s.AcrossRole = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetAcrossUid(v string) *CreateApsSlsADBJobShrinkRequest {
	s.AcrossUid = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetAdvancedConfig(v string) *CreateApsSlsADBJobShrinkRequest {
	s.AdvancedConfig = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetColumnsShrink(v string) *CreateApsSlsADBJobShrinkRequest {
	s.ColumnsShrink = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetDBClusterId(v string) *CreateApsSlsADBJobShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetDatasourceId(v int64) *CreateApsSlsADBJobShrinkRequest {
	s.DatasourceId = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetDbName(v string) *CreateApsSlsADBJobShrinkRequest {
	s.DbName = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetDirtyDataHandleMode(v string) *CreateApsSlsADBJobShrinkRequest {
	s.DirtyDataHandleMode = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetDirtyDataProcessPattern(v string) *CreateApsSlsADBJobShrinkRequest {
	s.DirtyDataProcessPattern = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetExactlyOnce(v string) *CreateApsSlsADBJobShrinkRequest {
	s.ExactlyOnce = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetFullComputeUnit(v string) *CreateApsSlsADBJobShrinkRequest {
	s.FullComputeUnit = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetHudiAdvancedConfig(v string) *CreateApsSlsADBJobShrinkRequest {
	s.HudiAdvancedConfig = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetIncrementalComputeUnit(v string) *CreateApsSlsADBJobShrinkRequest {
	s.IncrementalComputeUnit = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetLakehouseId(v int64) *CreateApsSlsADBJobShrinkRequest {
	s.LakehouseId = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetMaxOffsetsPerTrigger(v int64) *CreateApsSlsADBJobShrinkRequest {
	s.MaxOffsetsPerTrigger = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetOssLocation(v string) *CreateApsSlsADBJobShrinkRequest {
	s.OssLocation = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetOutputFormat(v string) *CreateApsSlsADBJobShrinkRequest {
	s.OutputFormat = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetPartitionSpecsShrink(v string) *CreateApsSlsADBJobShrinkRequest {
	s.PartitionSpecsShrink = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetPassword(v string) *CreateApsSlsADBJobShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetPrimaryKeyDefinition(v string) *CreateApsSlsADBJobShrinkRequest {
	s.PrimaryKeyDefinition = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetProject(v string) *CreateApsSlsADBJobShrinkRequest {
	s.Project = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetRegionId(v string) *CreateApsSlsADBJobShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetResourceGroup(v string) *CreateApsSlsADBJobShrinkRequest {
	s.ResourceGroup = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetSourceRegionId(v string) *CreateApsSlsADBJobShrinkRequest {
	s.SourceRegionId = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetStartingOffsets(v string) *CreateApsSlsADBJobShrinkRequest {
	s.StartingOffsets = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetStore(v string) *CreateApsSlsADBJobShrinkRequest {
	s.Store = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetTableName(v string) *CreateApsSlsADBJobShrinkRequest {
	s.TableName = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetTargetGenerateRule(v string) *CreateApsSlsADBJobShrinkRequest {
	s.TargetGenerateRule = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetTargetType(v string) *CreateApsSlsADBJobShrinkRequest {
	s.TargetType = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetUnixTimestampConvertShrink(v string) *CreateApsSlsADBJobShrinkRequest {
	s.UnixTimestampConvertShrink = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetUserName(v string) *CreateApsSlsADBJobShrinkRequest {
	s.UserName = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) SetWorkloadName(v string) *CreateApsSlsADBJobShrinkRequest {
	s.WorkloadName = &v
	return s
}

func (s *CreateApsSlsADBJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
