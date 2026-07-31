// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsKafkaHudiJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcrossRole(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetAcrossRole() *string
	SetAcrossUid(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetAcrossUid() *string
	SetAdvancedConfig(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetAdvancedConfig() *string
	SetColumnsShrink(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetColumnsShrink() *string
	SetDBClusterId(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetDBClusterId() *string
	SetDataFormatType(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetDataFormatType() *string
	SetDataOutputFormat(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetDataOutputFormat() *string
	SetDatasourceId(v int64) *CreateApsKafkaHudiJobShrinkRequest
	GetDatasourceId() *int64
	SetDbName(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetDbName() *string
	SetFullComputeUnit(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetFullComputeUnit() *string
	SetHudiAdvancedConfig(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetHudiAdvancedConfig() *string
	SetIncrementalComputeUnit(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetIncrementalComputeUnit() *string
	SetJsonParseLevel(v int32) *CreateApsKafkaHudiJobShrinkRequest
	GetJsonParseLevel() *int32
	SetKafkaClusterId(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetKafkaClusterId() *string
	SetKafkaTopic(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetKafkaTopic() *string
	SetLakehouseId(v int64) *CreateApsKafkaHudiJobShrinkRequest
	GetLakehouseId() *int64
	SetMaxOffsetsPerTrigger(v int64) *CreateApsKafkaHudiJobShrinkRequest
	GetMaxOffsetsPerTrigger() *int64
	SetOssLocation(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetOssLocation() *string
	SetOutputFormat(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetOutputFormat() *string
	SetPartitionSpecsShrink(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetPartitionSpecsShrink() *string
	SetPrimaryKeyDefinition(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetPrimaryKeyDefinition() *string
	SetRegionId(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetRegionId() *string
	SetResourceGroup(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetResourceGroup() *string
	SetSourceRegionId(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetSourceRegionId() *string
	SetStartingOffsets(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetStartingOffsets() *string
	SetTableName(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetTableName() *string
	SetTargetGenerateRule(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetTargetGenerateRule() *string
	SetTargetType(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetTargetType() *string
	SetWorkloadName(v string) *CreateApsKafkaHudiJobShrinkRequest
	GetWorkloadName() *string
}

type CreateApsKafkaHudiJobShrinkRequest struct {
	// The RAM role of a trusted entity that is an Alibaba Cloud account. For more information about how to create a RAM role, see Create a RAM role for a trusted Alibaba Cloud account.
	//
	// The Alibaba Cloud account that owns the AnalyticDB for MySQL cluster must be added as a trusted account to the RAM role.
	//
	// example:
	//
	// aps
	AcrossRole *string `json:"AcrossRole,omitempty" xml:"AcrossRole,omitempty"`
	// The ID of the Alibaba Cloud account to which the source Kafka instance belongs.
	//
	// example:
	//
	// 123************
	AcrossUid *string `json:"AcrossUid,omitempty" xml:"AcrossUid,omitempty"`
	// The advanced configuration.
	//
	// example:
	//
	// -
	AdvancedConfig *string `json:"AdvancedConfig,omitempty" xml:"AdvancedConfig,omitempty"`
	// The column information.
	//
	// This parameter is required.
	ColumnsShrink *string `json:"Columns,omitempty" xml:"Columns,omitempty"`
	// The cluster ID.
	//
	// > Call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to view the cluster IDs of all AnalyticDB for MySQL Data Lakehouse Edition (V3.0) clusters in the destination region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The Kafka message type. Valid values: json, general_canal_json, mongo_canal_json, dataworks_json, and shareplex_json.
	//
	// example:
	//
	// json
	DataFormatType *string `json:"DataFormatType,omitempty" xml:"DataFormatType,omitempty"`
	// The valid values and their descriptions are as follows:
	//
	// Single: The source is a single-line JSON record.
	//
	// Multi: The source is a JSON array. A single JSON record is returned as the output.
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
	// The user-defined name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// testDB
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The configuration for full synchronization.
	//
	// example:
	//
	// 2ACU
	FullComputeUnit *string `json:"FullComputeUnit,omitempty" xml:"FullComputeUnit,omitempty"`
	// The Hudi configuration for the destination.
	//
	// example:
	//
	// hoodie.keep.min.commits=20
	HudiAdvancedConfig *string `json:"HudiAdvancedConfig,omitempty" xml:"HudiAdvancedConfig,omitempty"`
	// The configuration for incremental synchronization.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ACU
	IncrementalComputeUnit *string `json:"IncrementalComputeUnit,omitempty" xml:"IncrementalComputeUnit,omitempty"`
	// The number of nested JSON layers to parse. Valid values:
	//
	// 0: No parsing is performed.
	//
	// 1: One layer is parsed.
	//
	// 2: Two layers are parsed.
	//
	// 3: Three layers are parsed.
	//
	// 4: Four layers are parsed.
	//
	// By default, one layer is parsed. For more information about the JSON parsing policy for nested data, see JSON parsing levels and schema field inference examples.
	//
	// example:
	//
	// 0
	JsonParseLevel *int32 `json:"JsonParseLevel,omitempty" xml:"JsonParseLevel,omitempty"`
	// The ID of the Kafka instance. Obtain the ID from the Kafka console.
	//
	// example:
	//
	// xxx
	KafkaClusterId *string `json:"KafkaClusterId,omitempty" xml:"KafkaClusterId,omitempty"`
	// The ID of the Kafka topic. Obtain the ID from the Kafka console.
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
	// The number of entries to consume in a single batch.
	//
	// example:
	//
	// 50000
	MaxOffsetsPerTrigger *int64 `json:"MaxOffsetsPerTrigger,omitempty" xml:"MaxOffsetsPerTrigger,omitempty"`
	// The destination lakehouse address. This must be a complete OSS path.
	//
	// example:
	//
	// oss://test-xx-zzz/yyy/
	OssLocation *string `json:"OssLocation,omitempty" xml:"OssLocation,omitempty"`
	// The output data format.
	//
	// example:
	//
	// HUDI
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// The partition information.
	//
	// if can be null:
	// true
	PartitionSpecsShrink *string `json:"PartitionSpecs,omitempty" xml:"PartitionSpecs,omitempty"`
	// The primary key settings. This parameter supports the UUID policy and the mapping policy. The policies are described as follows.
	//
	// UUID policy: "Strategy": "uuid".
	//
	// Mapping policy:
	//
	// "Strategy": "mapping",
	//
	// "Values":[
	//
	// "f1",
	//
	// "f2"
	//
	// ],
	//
	// "RecordVersionField","xxx"
	//
	// \\`RecordVersionField\\` specifies the Hudi record version.
	//
	// example:
	//
	// "Strategy": "mapping"
	PrimaryKeyDefinition *string `json:"PrimaryKeyDefinition,omitempty" xml:"PrimaryKeyDefinition,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// aps
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	SourceRegionId *string `json:"SourceRegionId,omitempty" xml:"SourceRegionId,omitempty"`
	// The initial consumer offset for Kafka.
	//
	// Valid values:
	//
	// begin_cursor, end_cursor, and timestamp.
	//
	// These values correspond to the earliest offset, the latest offset, and a specified time.
	//
	// This parameter is required.
	//
	// example:
	//
	// begincursor
	StartingOffsets *string `json:"StartingOffsets,omitempty" xml:"StartingOffsets,omitempty"`
	// The user-defined name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// testTB
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The generation rule for the destination.
	//
	// example:
	//
	// xxx
	TargetGenerateRule *string `json:"TargetGenerateRule,omitempty" xml:"TargetGenerateRule,omitempty"`
	// The type of the destination.
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

func (s CreateApsKafkaHudiJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApsKafkaHudiJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetAcrossRole() *string {
	return s.AcrossRole
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetAcrossUid() *string {
	return s.AcrossUid
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetAdvancedConfig() *string {
	return s.AdvancedConfig
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetColumnsShrink() *string {
	return s.ColumnsShrink
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetDataFormatType() *string {
	return s.DataFormatType
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetDataOutputFormat() *string {
	return s.DataOutputFormat
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetDbName() *string {
	return s.DbName
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetFullComputeUnit() *string {
	return s.FullComputeUnit
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetHudiAdvancedConfig() *string {
	return s.HudiAdvancedConfig
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetIncrementalComputeUnit() *string {
	return s.IncrementalComputeUnit
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetJsonParseLevel() *int32 {
	return s.JsonParseLevel
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetKafkaClusterId() *string {
	return s.KafkaClusterId
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetKafkaTopic() *string {
	return s.KafkaTopic
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetLakehouseId() *int64 {
	return s.LakehouseId
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetMaxOffsetsPerTrigger() *int64 {
	return s.MaxOffsetsPerTrigger
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetOssLocation() *string {
	return s.OssLocation
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetPartitionSpecsShrink() *string {
	return s.PartitionSpecsShrink
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetPrimaryKeyDefinition() *string {
	return s.PrimaryKeyDefinition
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetSourceRegionId() *string {
	return s.SourceRegionId
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetStartingOffsets() *string {
	return s.StartingOffsets
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetTableName() *string {
	return s.TableName
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetTargetGenerateRule() *string {
	return s.TargetGenerateRule
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateApsKafkaHudiJobShrinkRequest) GetWorkloadName() *string {
	return s.WorkloadName
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetAcrossRole(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.AcrossRole = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetAcrossUid(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.AcrossUid = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetAdvancedConfig(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.AdvancedConfig = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetColumnsShrink(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.ColumnsShrink = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetDBClusterId(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetDataFormatType(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.DataFormatType = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetDataOutputFormat(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.DataOutputFormat = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetDatasourceId(v int64) *CreateApsKafkaHudiJobShrinkRequest {
	s.DatasourceId = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetDbName(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.DbName = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetFullComputeUnit(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.FullComputeUnit = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetHudiAdvancedConfig(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.HudiAdvancedConfig = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetIncrementalComputeUnit(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.IncrementalComputeUnit = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetJsonParseLevel(v int32) *CreateApsKafkaHudiJobShrinkRequest {
	s.JsonParseLevel = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetKafkaClusterId(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.KafkaClusterId = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetKafkaTopic(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.KafkaTopic = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetLakehouseId(v int64) *CreateApsKafkaHudiJobShrinkRequest {
	s.LakehouseId = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetMaxOffsetsPerTrigger(v int64) *CreateApsKafkaHudiJobShrinkRequest {
	s.MaxOffsetsPerTrigger = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetOssLocation(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.OssLocation = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetOutputFormat(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.OutputFormat = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetPartitionSpecsShrink(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.PartitionSpecsShrink = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetPrimaryKeyDefinition(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.PrimaryKeyDefinition = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetRegionId(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetResourceGroup(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.ResourceGroup = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetSourceRegionId(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.SourceRegionId = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetStartingOffsets(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.StartingOffsets = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetTableName(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.TableName = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetTargetGenerateRule(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.TargetGenerateRule = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetTargetType(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.TargetType = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) SetWorkloadName(v string) *CreateApsKafkaHudiJobShrinkRequest {
	s.WorkloadName = &v
	return s
}

func (s *CreateApsKafkaHudiJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
