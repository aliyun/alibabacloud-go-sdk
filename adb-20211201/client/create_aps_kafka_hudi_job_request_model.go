// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsKafkaHudiJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrossRole(v string) *CreateApsKafkaHudiJobRequest
	GetAcrossRole() *string
	SetAcrossUid(v string) *CreateApsKafkaHudiJobRequest
	GetAcrossUid() *string
	SetAdvancedConfig(v string) *CreateApsKafkaHudiJobRequest
	GetAdvancedConfig() *string
	SetColumns(v []*CreateApsKafkaHudiJobRequestColumns) *CreateApsKafkaHudiJobRequest
	GetColumns() []*CreateApsKafkaHudiJobRequestColumns
	SetDBClusterId(v string) *CreateApsKafkaHudiJobRequest
	GetDBClusterId() *string
	SetDataOutputFormat(v string) *CreateApsKafkaHudiJobRequest
	GetDataOutputFormat() *string
	SetDatasourceId(v int64) *CreateApsKafkaHudiJobRequest
	GetDatasourceId() *int64
	SetDbName(v string) *CreateApsKafkaHudiJobRequest
	GetDbName() *string
	SetFullComputeUnit(v string) *CreateApsKafkaHudiJobRequest
	GetFullComputeUnit() *string
	SetHudiAdvancedConfig(v string) *CreateApsKafkaHudiJobRequest
	GetHudiAdvancedConfig() *string
	SetIncrementalComputeUnit(v string) *CreateApsKafkaHudiJobRequest
	GetIncrementalComputeUnit() *string
	SetJsonParseLevel(v int32) *CreateApsKafkaHudiJobRequest
	GetJsonParseLevel() *int32
	SetKafkaClusterId(v string) *CreateApsKafkaHudiJobRequest
	GetKafkaClusterId() *string
	SetKafkaTopic(v string) *CreateApsKafkaHudiJobRequest
	GetKafkaTopic() *string
	SetLakehouseId(v int64) *CreateApsKafkaHudiJobRequest
	GetLakehouseId() *int64
	SetMaxOffsetsPerTrigger(v int64) *CreateApsKafkaHudiJobRequest
	GetMaxOffsetsPerTrigger() *int64
	SetOssLocation(v string) *CreateApsKafkaHudiJobRequest
	GetOssLocation() *string
	SetOutputFormat(v string) *CreateApsKafkaHudiJobRequest
	GetOutputFormat() *string
	SetPartitionSpecs(v []map[string]interface{}) *CreateApsKafkaHudiJobRequest
	GetPartitionSpecs() []map[string]interface{}
	SetPrimaryKeyDefinition(v string) *CreateApsKafkaHudiJobRequest
	GetPrimaryKeyDefinition() *string
	SetRegionId(v string) *CreateApsKafkaHudiJobRequest
	GetRegionId() *string
	SetResourceGroup(v string) *CreateApsKafkaHudiJobRequest
	GetResourceGroup() *string
	SetSourceRegionId(v string) *CreateApsKafkaHudiJobRequest
	GetSourceRegionId() *string
	SetStartingOffsets(v string) *CreateApsKafkaHudiJobRequest
	GetStartingOffsets() *string
	SetTableName(v string) *CreateApsKafkaHudiJobRequest
	GetTableName() *string
	SetTargetGenerateRule(v string) *CreateApsKafkaHudiJobRequest
	GetTargetGenerateRule() *string
	SetTargetType(v string) *CreateApsKafkaHudiJobRequest
	GetTargetType() *string
	SetWorkloadName(v string) *CreateApsKafkaHudiJobRequest
	GetWorkloadName() *string
}

type CreateApsKafkaHudiJobRequest struct {
	// example:
	//
	// aps
	AcrossRole *string `json:"AcrossRole,omitempty" xml:"AcrossRole,omitempty"`
	// example:
	//
	// 123************
	AcrossUid *string `json:"AcrossUid,omitempty" xml:"AcrossUid,omitempty"`
	// example:
	//
	// -
	AdvancedConfig *string `json:"AdvancedConfig,omitempty" xml:"AdvancedConfig,omitempty"`
	// This parameter is required.
	Columns []*CreateApsKafkaHudiJobRequestColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// Single
	DataOutputFormat *string `json:"DataOutputFormat,omitempty" xml:"DataOutputFormat,omitempty"`
	// example:
	//
	// 1
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testDB
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// 2ACU
	FullComputeUnit *string `json:"FullComputeUnit,omitempty" xml:"FullComputeUnit,omitempty"`
	// example:
	//
	// hoodie.keep.min.commits=20
	HudiAdvancedConfig *string `json:"HudiAdvancedConfig,omitempty" xml:"HudiAdvancedConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2ACU
	IncrementalComputeUnit *string `json:"IncrementalComputeUnit,omitempty" xml:"IncrementalComputeUnit,omitempty"`
	// example:
	//
	// 0
	JsonParseLevel *int32 `json:"JsonParseLevel,omitempty" xml:"JsonParseLevel,omitempty"`
	// example:
	//
	// xxx
	KafkaClusterId *string `json:"KafkaClusterId,omitempty" xml:"KafkaClusterId,omitempty"`
	// example:
	//
	// test
	KafkaTopic *string `json:"KafkaTopic,omitempty" xml:"KafkaTopic,omitempty"`
	// example:
	//
	// 123
	LakehouseId *int64 `json:"LakehouseId,omitempty" xml:"LakehouseId,omitempty"`
	// example:
	//
	// 50000
	MaxOffsetsPerTrigger *int64 `json:"MaxOffsetsPerTrigger,omitempty" xml:"MaxOffsetsPerTrigger,omitempty"`
	// example:
	//
	// oss://test-xx-zzz/yyy/
	OssLocation *string `json:"OssLocation,omitempty" xml:"OssLocation,omitempty"`
	// example:
	//
	// HUDI
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// if can be null:
	// true
	PartitionSpecs []map[string]interface{} `json:"PartitionSpecs,omitempty" xml:"PartitionSpecs,omitempty" type:"Repeated"`
	// example:
	//
	// "Strategy": "mapping"
	PrimaryKeyDefinition *string `json:"PrimaryKeyDefinition,omitempty" xml:"PrimaryKeyDefinition,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aps
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// example:
	//
	// cn-hangzhou
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// begincursor
	StartingOffsets *string `json:"StartingOffsets,omitempty" xml:"StartingOffsets,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testTB
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// xxx
	TargetGenerateRule *string `json:"TargetGenerateRule,omitempty" xml:"TargetGenerateRule,omitempty"`
	// example:
	//
	// OSS
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	WorkloadName *string `json:"WorkloadName,omitempty" xml:"WorkloadName,omitempty"`
}

func (s CreateApsKafkaHudiJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApsKafkaHudiJobRequest) GoString() string {
	return s.String()
}

func (s *CreateApsKafkaHudiJobRequest) GetAcrossRole() *string {
	return s.AcrossRole
}

func (s *CreateApsKafkaHudiJobRequest) GetAcrossUid() *string {
	return s.AcrossUid
}

func (s *CreateApsKafkaHudiJobRequest) GetAdvancedConfig() *string {
	return s.AdvancedConfig
}

func (s *CreateApsKafkaHudiJobRequest) GetColumns() []*CreateApsKafkaHudiJobRequestColumns {
	return s.Columns
}

func (s *CreateApsKafkaHudiJobRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApsKafkaHudiJobRequest) GetDataOutputFormat() *string {
	return s.DataOutputFormat
}

func (s *CreateApsKafkaHudiJobRequest) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *CreateApsKafkaHudiJobRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateApsKafkaHudiJobRequest) GetFullComputeUnit() *string {
	return s.FullComputeUnit
}

func (s *CreateApsKafkaHudiJobRequest) GetHudiAdvancedConfig() *string {
	return s.HudiAdvancedConfig
}

func (s *CreateApsKafkaHudiJobRequest) GetIncrementalComputeUnit() *string {
	return s.IncrementalComputeUnit
}

func (s *CreateApsKafkaHudiJobRequest) GetJsonParseLevel() *int32 {
	return s.JsonParseLevel
}

func (s *CreateApsKafkaHudiJobRequest) GetKafkaClusterId() *string {
	return s.KafkaClusterId
}

func (s *CreateApsKafkaHudiJobRequest) GetKafkaTopic() *string {
	return s.KafkaTopic
}

func (s *CreateApsKafkaHudiJobRequest) GetLakehouseId() *int64 {
	return s.LakehouseId
}

func (s *CreateApsKafkaHudiJobRequest) GetMaxOffsetsPerTrigger() *int64 {
	return s.MaxOffsetsPerTrigger
}

func (s *CreateApsKafkaHudiJobRequest) GetOssLocation() *string {
	return s.OssLocation
}

func (s *CreateApsKafkaHudiJobRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *CreateApsKafkaHudiJobRequest) GetPartitionSpecs() []map[string]interface{} {
	return s.PartitionSpecs
}

func (s *CreateApsKafkaHudiJobRequest) GetPrimaryKeyDefinition() *string {
	return s.PrimaryKeyDefinition
}

func (s *CreateApsKafkaHudiJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApsKafkaHudiJobRequest) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *CreateApsKafkaHudiJobRequest) GetSourceRegionId() *string {
	return s.SourceRegionId
}

func (s *CreateApsKafkaHudiJobRequest) GetStartingOffsets() *string {
	return s.StartingOffsets
}

func (s *CreateApsKafkaHudiJobRequest) GetTableName() *string {
	return s.TableName
}

func (s *CreateApsKafkaHudiJobRequest) GetTargetGenerateRule() *string {
	return s.TargetGenerateRule
}

func (s *CreateApsKafkaHudiJobRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateApsKafkaHudiJobRequest) GetWorkloadName() *string {
	return s.WorkloadName
}

func (s *CreateApsKafkaHudiJobRequest) SetAcrossRole(v string) *CreateApsKafkaHudiJobRequest {
	s.AcrossRole = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetAcrossUid(v string) *CreateApsKafkaHudiJobRequest {
	s.AcrossUid = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetAdvancedConfig(v string) *CreateApsKafkaHudiJobRequest {
	s.AdvancedConfig = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetColumns(v []*CreateApsKafkaHudiJobRequestColumns) *CreateApsKafkaHudiJobRequest {
	s.Columns = v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetDBClusterId(v string) *CreateApsKafkaHudiJobRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetDataOutputFormat(v string) *CreateApsKafkaHudiJobRequest {
	s.DataOutputFormat = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetDatasourceId(v int64) *CreateApsKafkaHudiJobRequest {
	s.DatasourceId = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetDbName(v string) *CreateApsKafkaHudiJobRequest {
	s.DbName = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetFullComputeUnit(v string) *CreateApsKafkaHudiJobRequest {
	s.FullComputeUnit = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetHudiAdvancedConfig(v string) *CreateApsKafkaHudiJobRequest {
	s.HudiAdvancedConfig = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetIncrementalComputeUnit(v string) *CreateApsKafkaHudiJobRequest {
	s.IncrementalComputeUnit = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetJsonParseLevel(v int32) *CreateApsKafkaHudiJobRequest {
	s.JsonParseLevel = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetKafkaClusterId(v string) *CreateApsKafkaHudiJobRequest {
	s.KafkaClusterId = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetKafkaTopic(v string) *CreateApsKafkaHudiJobRequest {
	s.KafkaTopic = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetLakehouseId(v int64) *CreateApsKafkaHudiJobRequest {
	s.LakehouseId = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetMaxOffsetsPerTrigger(v int64) *CreateApsKafkaHudiJobRequest {
	s.MaxOffsetsPerTrigger = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetOssLocation(v string) *CreateApsKafkaHudiJobRequest {
	s.OssLocation = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetOutputFormat(v string) *CreateApsKafkaHudiJobRequest {
	s.OutputFormat = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetPartitionSpecs(v []map[string]interface{}) *CreateApsKafkaHudiJobRequest {
	s.PartitionSpecs = v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetPrimaryKeyDefinition(v string) *CreateApsKafkaHudiJobRequest {
	s.PrimaryKeyDefinition = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetRegionId(v string) *CreateApsKafkaHudiJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetResourceGroup(v string) *CreateApsKafkaHudiJobRequest {
	s.ResourceGroup = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetSourceRegionId(v string) *CreateApsKafkaHudiJobRequest {
	s.SourceRegionId = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetStartingOffsets(v string) *CreateApsKafkaHudiJobRequest {
	s.StartingOffsets = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetTableName(v string) *CreateApsKafkaHudiJobRequest {
	s.TableName = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetTargetGenerateRule(v string) *CreateApsKafkaHudiJobRequest {
	s.TargetGenerateRule = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetTargetType(v string) *CreateApsKafkaHudiJobRequest {
	s.TargetType = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) SetWorkloadName(v string) *CreateApsKafkaHudiJobRequest {
	s.WorkloadName = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequest) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateApsKafkaHudiJobRequestColumns struct {
	// example:
	//
	// b
	MapName *string `json:"MapName,omitempty" xml:"MapName,omitempty"`
	// example:
	//
	// string
	MapType *string `json:"MapType,omitempty" xml:"MapType,omitempty"`
	// example:
	//
	// a
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// string
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateApsKafkaHudiJobRequestColumns) String() string {
	return dara.Prettify(s)
}

func (s CreateApsKafkaHudiJobRequestColumns) GoString() string {
	return s.String()
}

func (s *CreateApsKafkaHudiJobRequestColumns) GetMapName() *string {
	return s.MapName
}

func (s *CreateApsKafkaHudiJobRequestColumns) GetMapType() *string {
	return s.MapType
}

func (s *CreateApsKafkaHudiJobRequestColumns) GetName() *string {
	return s.Name
}

func (s *CreateApsKafkaHudiJobRequestColumns) GetType() *string {
	return s.Type
}

func (s *CreateApsKafkaHudiJobRequestColumns) SetMapName(v string) *CreateApsKafkaHudiJobRequestColumns {
	s.MapName = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequestColumns) SetMapType(v string) *CreateApsKafkaHudiJobRequestColumns {
	s.MapType = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequestColumns) SetName(v string) *CreateApsKafkaHudiJobRequestColumns {
	s.Name = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequestColumns) SetType(v string) *CreateApsKafkaHudiJobRequestColumns {
	s.Type = &v
	return s
}

func (s *CreateApsKafkaHudiJobRequestColumns) Validate() error {
	return dara.Validate(s)
}
