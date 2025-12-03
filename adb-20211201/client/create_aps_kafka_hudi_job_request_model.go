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
	SetDataFormatType(v string) *CreateApsKafkaHudiJobRequest
	GetDataFormatType() *string
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
	// The Resource Access Management (RAM) role that is created for the trusted Alibaba Cloud account. For more information, see Create a RAM role for a trusted Alibaba Cloud account. The ARN of the RAM role that grants AnalyticDB for MySQL permission to access resources in the source account. Required for cross-account data ingestion.
	//
	// example:
	//
	// aps
	AcrossRole *string `json:"AcrossRole,omitempty" xml:"AcrossRole,omitempty"`
	// The ID of the Alibaba Cloud account to which the source Kafka belongs.
	//
	// example:
	//
	// 123************
	AcrossUid *string `json:"AcrossUid,omitempty" xml:"AcrossUid,omitempty"`
	// The advanced configurations.
	//
	// example:
	//
	// -
	AdvancedConfig *string `json:"AdvancedConfig,omitempty" xml:"AdvancedConfig,omitempty"`
	// The column information.
	//
	// This parameter is required.
	Columns []*CreateApsKafkaHudiJobRequestColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all clusters in a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId    *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	DataFormatType *string `json:"DataFormatType,omitempty" xml:"DataFormatType,omitempty"`
	// Enumeration value and description. Single: The source is a single-row JSON record. Multi: source is a JSON array. Output a single JSON record.
	//
	// example:
	//
	// Single
	DataOutputFormat *string `json:"DataOutputFormat,omitempty" xml:"DataOutputFormat,omitempty"`
	// The data source ID.
	//
	// example:
	//
	// 1
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// The name of the user-defined database.
	//
	// This parameter is required.
	//
	// example:
	//
	// testDB
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The full synchronization configuration.
	//
	// example:
	//
	// 2ACU
	FullComputeUnit *string `json:"FullComputeUnit,omitempty" xml:"FullComputeUnit,omitempty"`
	// The HUDI configuration of the destination.
	//
	// example:
	//
	// hoodie.keep.min.commits=20
	HudiAdvancedConfig *string `json:"HudiAdvancedConfig,omitempty" xml:"HudiAdvancedConfig,omitempty"`
	// The incremental synchronization configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ACU
	IncrementalComputeUnit *string `json:"IncrementalComputeUnit,omitempty" xml:"IncrementalComputeUnit,omitempty"`
	// The number of layers that are parsed for nested JSON fields. Valid values: 0: Nested JSON fields are not parsed. 1: parses one layer. 2: Two layers are parsed. 3: Three layers are parsed. 4: Four layers are parsed. By default, one layer is parsed. For more information about how nested JSON fields are parsed, see the Examples of schema fields parsed with different numbers of layers section of this topic.
	//
	// example:
	//
	// 0
	JsonParseLevel *int32 `json:"JsonParseLevel,omitempty" xml:"JsonParseLevel,omitempty"`
	// The ID of the Apache Kafka instance. You can get it in the Kafka console.
	//
	// example:
	//
	// xxx
	KafkaClusterId *string `json:"KafkaClusterId,omitempty" xml:"KafkaClusterId,omitempty"`
	// Kafka Topic ID. You can get it in the Kafka console.
	//
	// example:
	//
	// test
	KafkaTopic *string `json:"KafkaTopic,omitempty" xml:"KafkaTopic,omitempty"`
	// The ID of the lakehouse.
	//
	// example:
	//
	// 123
	LakehouseId *int64 `json:"LakehouseId,omitempty" xml:"LakehouseId,omitempty"`
	// The maximum number of records to fetch in a single batch.
	//
	// example:
	//
	// 50000
	MaxOffsetsPerTrigger *int64 `json:"MaxOffsetsPerTrigger,omitempty" xml:"MaxOffsetsPerTrigger,omitempty"`
	// The path of the destination data lakehouse in an Object Storage Service (OSS) bucket.
	//
	// example:
	//
	// oss://test-xx-zzz/yyy/
	OssLocation *string `json:"OssLocation,omitempty" xml:"OssLocation,omitempty"`
	// The format of the output data.
	//
	// example:
	//
	// HUDI
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// The partition information.
	//
	// if can be null:
	// true
	PartitionSpecs []map[string]interface{} `json:"PartitionSpecs,omitempty" xml:"PartitionSpecs,omitempty" type:"Repeated"`
	// The primary key settings. Contains the uuid policy and mapping policy. The explanation is as follows. Uuid policy: "Strategy": "uuid". Mapping policy: "Strategy": "mapping", "Values":[ "f1", "f2" ], "RecordVersionField","xxx" The meaning of the RecordVersionField is the HUDI record version.
	//
	// example:
	//
	// "Strategy": "mapping"
	PrimaryKeyDefinition *string `json:"PrimaryKeyDefinition,omitempty" xml:"PrimaryKeyDefinition,omitempty"`
	// The region ID of the cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// aps
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// 地域ID。
	//
	// example:
	//
	// cn-hangzhou
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// Specifies the position from which to start consuming messages. Valid values: begin_cursor/end_cursor/timestamp Each corresponds to the earliest /latest /specified time respectively.
	//
	// This parameter is required.
	//
	// example:
	//
	// begincursor
	StartingOffsets *string `json:"StartingOffsets,omitempty" xml:"StartingOffsets,omitempty"`
	// The name of the user-defined table.
	//
	// This parameter is required.
	//
	// example:
	//
	// testTB
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The rules for generating the destination database.
	//
	// example:
	//
	// xxx
	TargetGenerateRule *string `json:"TargetGenerateRule,omitempty" xml:"TargetGenerateRule,omitempty"`
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

func (s *CreateApsKafkaHudiJobRequest) GetDataFormatType() *string {
	return s.DataFormatType
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

func (s *CreateApsKafkaHudiJobRequest) SetDataFormatType(v string) *CreateApsKafkaHudiJobRequest {
	s.DataFormatType = &v
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
	// The name of the partition column in the destination table.
	//
	// example:
	//
	// b
	MapName *string `json:"MapName,omitempty" xml:"MapName,omitempty"`
	// The desired format for the destination partition column.
	//
	// example:
	//
	// string
	MapType *string `json:"MapType,omitempty" xml:"MapType,omitempty"`
	// The name of the source column to use for partitioning.
	//
	// example:
	//
	// a
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The format of the source field. See the table below for valid values.
	//
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
