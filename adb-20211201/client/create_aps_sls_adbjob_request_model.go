// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsSlsADBJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrossRole(v string) *CreateApsSlsADBJobRequest
	GetAcrossRole() *string
	SetAcrossUid(v string) *CreateApsSlsADBJobRequest
	GetAcrossUid() *string
	SetAdvancedConfig(v string) *CreateApsSlsADBJobRequest
	GetAdvancedConfig() *string
	SetColumns(v []*CreateApsSlsADBJobRequestColumns) *CreateApsSlsADBJobRequest
	GetColumns() []*CreateApsSlsADBJobRequestColumns
	SetDBClusterId(v string) *CreateApsSlsADBJobRequest
	GetDBClusterId() *string
	SetDatasourceId(v int64) *CreateApsSlsADBJobRequest
	GetDatasourceId() *int64
	SetDbName(v string) *CreateApsSlsADBJobRequest
	GetDbName() *string
	SetDirtyDataHandleMode(v string) *CreateApsSlsADBJobRequest
	GetDirtyDataHandleMode() *string
	SetDirtyDataProcessPattern(v string) *CreateApsSlsADBJobRequest
	GetDirtyDataProcessPattern() *string
	SetExactlyOnce(v string) *CreateApsSlsADBJobRequest
	GetExactlyOnce() *string
	SetFullComputeUnit(v string) *CreateApsSlsADBJobRequest
	GetFullComputeUnit() *string
	SetHudiAdvancedConfig(v string) *CreateApsSlsADBJobRequest
	GetHudiAdvancedConfig() *string
	SetIncrementalComputeUnit(v string) *CreateApsSlsADBJobRequest
	GetIncrementalComputeUnit() *string
	SetLakehouseId(v int64) *CreateApsSlsADBJobRequest
	GetLakehouseId() *int64
	SetMaxOffsetsPerTrigger(v int64) *CreateApsSlsADBJobRequest
	GetMaxOffsetsPerTrigger() *int64
	SetOssLocation(v string) *CreateApsSlsADBJobRequest
	GetOssLocation() *string
	SetOutputFormat(v string) *CreateApsSlsADBJobRequest
	GetOutputFormat() *string
	SetPartitionSpecs(v []map[string]interface{}) *CreateApsSlsADBJobRequest
	GetPartitionSpecs() []map[string]interface{}
	SetPassword(v string) *CreateApsSlsADBJobRequest
	GetPassword() *string
	SetPrimaryKeyDefinition(v string) *CreateApsSlsADBJobRequest
	GetPrimaryKeyDefinition() *string
	SetProject(v string) *CreateApsSlsADBJobRequest
	GetProject() *string
	SetRegionId(v string) *CreateApsSlsADBJobRequest
	GetRegionId() *string
	SetResourceGroup(v string) *CreateApsSlsADBJobRequest
	GetResourceGroup() *string
	SetSourceRegionId(v string) *CreateApsSlsADBJobRequest
	GetSourceRegionId() *string
	SetStartingOffsets(v string) *CreateApsSlsADBJobRequest
	GetStartingOffsets() *string
	SetStore(v string) *CreateApsSlsADBJobRequest
	GetStore() *string
	SetTableName(v string) *CreateApsSlsADBJobRequest
	GetTableName() *string
	SetTargetGenerateRule(v string) *CreateApsSlsADBJobRequest
	GetTargetGenerateRule() *string
	SetTargetType(v string) *CreateApsSlsADBJobRequest
	GetTargetType() *string
	SetUnixTimestampConvert(v *CreateApsSlsADBJobRequestUnixTimestampConvert) *CreateApsSlsADBJobRequest
	GetUnixTimestampConvert() *CreateApsSlsADBJobRequestUnixTimestampConvert
	SetUserName(v string) *CreateApsSlsADBJobRequest
	GetUserName() *string
	SetWorkloadName(v string) *CreateApsSlsADBJobRequest
	GetWorkloadName() *string
}

type CreateApsSlsADBJobRequest struct {
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
	Columns []*CreateApsSlsADBJobRequestColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
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
	PartitionSpecs []map[string]interface{} `json:"PartitionSpecs,omitempty" xml:"PartitionSpecs,omitempty" type:"Repeated"`
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
	UnixTimestampConvert *CreateApsSlsADBJobRequestUnixTimestampConvert `json:"UnixTimestampConvert,omitempty" xml:"UnixTimestampConvert,omitempty" type:"Struct"`
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

func (s CreateApsSlsADBJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApsSlsADBJobRequest) GoString() string {
	return s.String()
}

func (s *CreateApsSlsADBJobRequest) GetAcrossRole() *string {
	return s.AcrossRole
}

func (s *CreateApsSlsADBJobRequest) GetAcrossUid() *string {
	return s.AcrossUid
}

func (s *CreateApsSlsADBJobRequest) GetAdvancedConfig() *string {
	return s.AdvancedConfig
}

func (s *CreateApsSlsADBJobRequest) GetColumns() []*CreateApsSlsADBJobRequestColumns {
	return s.Columns
}

func (s *CreateApsSlsADBJobRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApsSlsADBJobRequest) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *CreateApsSlsADBJobRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateApsSlsADBJobRequest) GetDirtyDataHandleMode() *string {
	return s.DirtyDataHandleMode
}

func (s *CreateApsSlsADBJobRequest) GetDirtyDataProcessPattern() *string {
	return s.DirtyDataProcessPattern
}

func (s *CreateApsSlsADBJobRequest) GetExactlyOnce() *string {
	return s.ExactlyOnce
}

func (s *CreateApsSlsADBJobRequest) GetFullComputeUnit() *string {
	return s.FullComputeUnit
}

func (s *CreateApsSlsADBJobRequest) GetHudiAdvancedConfig() *string {
	return s.HudiAdvancedConfig
}

func (s *CreateApsSlsADBJobRequest) GetIncrementalComputeUnit() *string {
	return s.IncrementalComputeUnit
}

func (s *CreateApsSlsADBJobRequest) GetLakehouseId() *int64 {
	return s.LakehouseId
}

func (s *CreateApsSlsADBJobRequest) GetMaxOffsetsPerTrigger() *int64 {
	return s.MaxOffsetsPerTrigger
}

func (s *CreateApsSlsADBJobRequest) GetOssLocation() *string {
	return s.OssLocation
}

func (s *CreateApsSlsADBJobRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *CreateApsSlsADBJobRequest) GetPartitionSpecs() []map[string]interface{} {
	return s.PartitionSpecs
}

func (s *CreateApsSlsADBJobRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateApsSlsADBJobRequest) GetPrimaryKeyDefinition() *string {
	return s.PrimaryKeyDefinition
}

func (s *CreateApsSlsADBJobRequest) GetProject() *string {
	return s.Project
}

func (s *CreateApsSlsADBJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApsSlsADBJobRequest) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *CreateApsSlsADBJobRequest) GetSourceRegionId() *string {
	return s.SourceRegionId
}

func (s *CreateApsSlsADBJobRequest) GetStartingOffsets() *string {
	return s.StartingOffsets
}

func (s *CreateApsSlsADBJobRequest) GetStore() *string {
	return s.Store
}

func (s *CreateApsSlsADBJobRequest) GetTableName() *string {
	return s.TableName
}

func (s *CreateApsSlsADBJobRequest) GetTargetGenerateRule() *string {
	return s.TargetGenerateRule
}

func (s *CreateApsSlsADBJobRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateApsSlsADBJobRequest) GetUnixTimestampConvert() *CreateApsSlsADBJobRequestUnixTimestampConvert {
	return s.UnixTimestampConvert
}

func (s *CreateApsSlsADBJobRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateApsSlsADBJobRequest) GetWorkloadName() *string {
	return s.WorkloadName
}

func (s *CreateApsSlsADBJobRequest) SetAcrossRole(v string) *CreateApsSlsADBJobRequest {
	s.AcrossRole = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetAcrossUid(v string) *CreateApsSlsADBJobRequest {
	s.AcrossUid = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetAdvancedConfig(v string) *CreateApsSlsADBJobRequest {
	s.AdvancedConfig = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetColumns(v []*CreateApsSlsADBJobRequestColumns) *CreateApsSlsADBJobRequest {
	s.Columns = v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetDBClusterId(v string) *CreateApsSlsADBJobRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetDatasourceId(v int64) *CreateApsSlsADBJobRequest {
	s.DatasourceId = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetDbName(v string) *CreateApsSlsADBJobRequest {
	s.DbName = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetDirtyDataHandleMode(v string) *CreateApsSlsADBJobRequest {
	s.DirtyDataHandleMode = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetDirtyDataProcessPattern(v string) *CreateApsSlsADBJobRequest {
	s.DirtyDataProcessPattern = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetExactlyOnce(v string) *CreateApsSlsADBJobRequest {
	s.ExactlyOnce = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetFullComputeUnit(v string) *CreateApsSlsADBJobRequest {
	s.FullComputeUnit = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetHudiAdvancedConfig(v string) *CreateApsSlsADBJobRequest {
	s.HudiAdvancedConfig = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetIncrementalComputeUnit(v string) *CreateApsSlsADBJobRequest {
	s.IncrementalComputeUnit = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetLakehouseId(v int64) *CreateApsSlsADBJobRequest {
	s.LakehouseId = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetMaxOffsetsPerTrigger(v int64) *CreateApsSlsADBJobRequest {
	s.MaxOffsetsPerTrigger = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetOssLocation(v string) *CreateApsSlsADBJobRequest {
	s.OssLocation = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetOutputFormat(v string) *CreateApsSlsADBJobRequest {
	s.OutputFormat = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetPartitionSpecs(v []map[string]interface{}) *CreateApsSlsADBJobRequest {
	s.PartitionSpecs = v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetPassword(v string) *CreateApsSlsADBJobRequest {
	s.Password = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetPrimaryKeyDefinition(v string) *CreateApsSlsADBJobRequest {
	s.PrimaryKeyDefinition = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetProject(v string) *CreateApsSlsADBJobRequest {
	s.Project = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetRegionId(v string) *CreateApsSlsADBJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetResourceGroup(v string) *CreateApsSlsADBJobRequest {
	s.ResourceGroup = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetSourceRegionId(v string) *CreateApsSlsADBJobRequest {
	s.SourceRegionId = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetStartingOffsets(v string) *CreateApsSlsADBJobRequest {
	s.StartingOffsets = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetStore(v string) *CreateApsSlsADBJobRequest {
	s.Store = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetTableName(v string) *CreateApsSlsADBJobRequest {
	s.TableName = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetTargetGenerateRule(v string) *CreateApsSlsADBJobRequest {
	s.TargetGenerateRule = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetTargetType(v string) *CreateApsSlsADBJobRequest {
	s.TargetType = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetUnixTimestampConvert(v *CreateApsSlsADBJobRequestUnixTimestampConvert) *CreateApsSlsADBJobRequest {
	s.UnixTimestampConvert = v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetUserName(v string) *CreateApsSlsADBJobRequest {
	s.UserName = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) SetWorkloadName(v string) *CreateApsSlsADBJobRequest {
	s.WorkloadName = &v
	return s
}

func (s *CreateApsSlsADBJobRequest) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UnixTimestampConvert != nil {
		if err := s.UnixTimestampConvert.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateApsSlsADBJobRequestColumns struct {
	// The name of the mapping.
	//
	// example:
	//
	// test
	MapName *string `json:"MapName,omitempty" xml:"MapName,omitempty"`
	// The type of the mapping.
	//
	// example:
	//
	// int
	MapType *string `json:"MapType,omitempty" xml:"MapType,omitempty"`
	// The name of the column.
	//
	// example:
	//
	// id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The data type of the column.
	//
	// example:
	//
	// bigint
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateApsSlsADBJobRequestColumns) String() string {
	return dara.Prettify(s)
}

func (s CreateApsSlsADBJobRequestColumns) GoString() string {
	return s.String()
}

func (s *CreateApsSlsADBJobRequestColumns) GetMapName() *string {
	return s.MapName
}

func (s *CreateApsSlsADBJobRequestColumns) GetMapType() *string {
	return s.MapType
}

func (s *CreateApsSlsADBJobRequestColumns) GetName() *string {
	return s.Name
}

func (s *CreateApsSlsADBJobRequestColumns) GetType() *string {
	return s.Type
}

func (s *CreateApsSlsADBJobRequestColumns) SetMapName(v string) *CreateApsSlsADBJobRequestColumns {
	s.MapName = &v
	return s
}

func (s *CreateApsSlsADBJobRequestColumns) SetMapType(v string) *CreateApsSlsADBJobRequestColumns {
	s.MapType = &v
	return s
}

func (s *CreateApsSlsADBJobRequestColumns) SetName(v string) *CreateApsSlsADBJobRequestColumns {
	s.Name = &v
	return s
}

func (s *CreateApsSlsADBJobRequestColumns) SetType(v string) *CreateApsSlsADBJobRequestColumns {
	s.Type = &v
	return s
}

func (s *CreateApsSlsADBJobRequestColumns) Validate() error {
	return dara.Validate(s)
}

type CreateApsSlsADBJobRequestUnixTimestampConvert struct {
	// Specifies whether to enable the conversion of timestamps.
	//
	// example:
	//
	// false
	Convert *string `json:"Convert,omitempty" xml:"Convert,omitempty"`
	// The format of the timestamp.
	//
	// example:
	//
	// yyyyMMdd
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// Specifies whether to enable the timestamp conversion.
	//
	// example:
	//
	// false
	Transform *bool `json:"Transform,omitempty" xml:"Transform,omitempty"`
}

func (s CreateApsSlsADBJobRequestUnixTimestampConvert) String() string {
	return dara.Prettify(s)
}

func (s CreateApsSlsADBJobRequestUnixTimestampConvert) GoString() string {
	return s.String()
}

func (s *CreateApsSlsADBJobRequestUnixTimestampConvert) GetConvert() *string {
	return s.Convert
}

func (s *CreateApsSlsADBJobRequestUnixTimestampConvert) GetFormat() *string {
	return s.Format
}

func (s *CreateApsSlsADBJobRequestUnixTimestampConvert) GetTransform() *bool {
	return s.Transform
}

func (s *CreateApsSlsADBJobRequestUnixTimestampConvert) SetConvert(v string) *CreateApsSlsADBJobRequestUnixTimestampConvert {
	s.Convert = &v
	return s
}

func (s *CreateApsSlsADBJobRequestUnixTimestampConvert) SetFormat(v string) *CreateApsSlsADBJobRequestUnixTimestampConvert {
	s.Format = &v
	return s
}

func (s *CreateApsSlsADBJobRequestUnixTimestampConvert) SetTransform(v bool) *CreateApsSlsADBJobRequestUnixTimestampConvert {
	s.Transform = &v
	return s
}

func (s *CreateApsSlsADBJobRequestUnixTimestampConvert) Validate() error {
	return dara.Validate(s)
}
