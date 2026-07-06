// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTable interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessMetadata(v *TableBusinessMetadata) *Table
	GetBusinessMetadata() *TableBusinessMetadata
	SetComment(v string) *Table
	GetComment() *string
	SetCreateTime(v int64) *Table
	GetCreateTime() *int64
	SetId(v string) *Table
	GetId() *string
	SetModifyTime(v int64) *Table
	GetModifyTime() *int64
	SetName(v string) *Table
	GetName() *string
	SetParentMetaEntityId(v string) *Table
	GetParentMetaEntityId() *string
	SetPartitionKeys(v []*string) *Table
	GetPartitionKeys() []*string
	SetStatisticsInfos(v map[string]*string) *Table
	GetStatisticsInfos() map[string]*string
	SetTableType(v string) *Table
	GetTableType() *string
	SetTechnicalMetadata(v *TableTechnicalMetadata) *Table
	GetTechnicalMetadata() *TableTechnicalMetadata
}

type Table struct {
	// The business metadata related to DataWorks, including usage instructions, tags, categories, upstream production nodes, and extended information.
	BusinessMetadata *TableBusinessMetadata `json:"BusinessMetadata,omitempty" xml:"BusinessMetadata,omitempty" type:"Struct"`
	// The comment.
	//
	// example:
	//
	// 测试表
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The creation time, in millisecond-level timestamp.
	//
	// example:
	//
	// 1736852168000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID. For more information, see [Metadata entity concepts](https://help.aliyun.com/document_detail/2880092.html).
	//
	// The format is `${EntityType}:${instance ID or encoded URL}:${DataFolderIdentity}:${DatabaseName}:${PatternName}:${TableName}`. Use an empty character as a placeholder for levels that do not exist.
	//
	// > For maxcompute and dlf types, use an empty string as a placeholder for the instance ID. For the maxcompute type, the database name is the MaxCompute project name. Projects with the three-layer model enabled require a schema name. For projects without the three-layer model enabled, use an empty string as a placeholder for the schema name.
	//
	// > For the starrocks type, the data catalog identifier is the catalog name. For the dlf type, the data catalog identifier is the catalog ID. Other types do not support the catalog level. Use an empty string as a placeholder.
	//
	// The following examples show the ID formats for common types:
	//
	// `maxcompute-table:::project_name:[schema_name]:table_name`
	//
	// `dlf-table::catalog_id:database_name::table_name`
	//
	// `hms-table:instance_id::database_name::table_name`
	//
	// `holo-table:instance_id::database_name:schema_name:table_name`
	//
	// `mysql-table:(instance_id|encoded_jdbc_url)::database_name::table_name`
	//
	// > Where
	//
	// `instance_id`: The instance ID. This is required when the data source is registered in instance mode.
	//
	// `encoded_jdbc_url`: The URL-encoded JDBC connection string. This is required when the data source is registered by using a connection string.
	//
	// `catalog_id`: The DLF catalog ID.
	//
	// `project_name`: The MaxCompute project name.
	//
	// `database_name`: The database name.
	//
	// `schema_name`: The schema name. For the maxcompute type, this is required only when the three-layer model is enabled for the project. If the three-layer model is not enabled, use an empty string as a placeholder.
	//
	// `table_name`: The table name.
	//
	// example:
	//
	// dlf-table::catalog_id:database_name::table_name
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The modification time, in millisecond-level timestamp.
	//
	// example:
	//
	// 1736852168000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name.
	//
	// example:
	//
	// table_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parent-level metadata entity ID. For more information, see [Metadata entity concepts](https://help.aliyun.com/document_detail/2880092.html).
	//
	// - For types that support schemas (`maxcompute/holo/postgresql/sqlserver/hybriddb_for_postgresql/oracle, where the maxcompute type requires the Layer 3 model to be enabled for the project`), ParentMetaEntityId is the database pattern to which the table belongs. The format is `${EntityType}:${instance ID or encoded URL}:${DataFolderIdentity}:${DatabaseName}:${PatternName}`. Use an empty character as a placeholder for levels that do not exist.
	//
	// - For other types, ParentMetaEntityId is the database to which the table belongs. The format is `${EntityType}:${instance ID or encoded URL}:${DataFolderIdentity}:${DatabaseName}`. Use an empty character as a placeholder for levels that do not exist.
	//
	// > For maxcompute and dlf types, use an empty string as a placeholder for the instance ID. For the maxcompute type, the database name is the MaxCompute project name.
	//
	// > For the starrocks type, the data catalog identifier is the catalog name. For the dlf type, the data catalog identifier is the catalog ID. Other types do not support the catalog level. Use an empty string as a placeholder.
	//
	// The following examples show the ParentMetaEntityId formats for common types:
	//
	//
	//
	// `maxcompute-project:::project_name`
	//
	// `maxcompute-schema:::project_name:schema_name` (only when the three-layer model is enabled for the project)
	//
	// `dlf-database::catalog_id:database_name`
	//
	// `hms-database:instance_id::database_name`
	//
	// `holo-schema:instance_id::database_name:schema_name`
	//
	// `mysql-database:(instance_id|encoded_jdbc_url)::database_name`
	//
	// > Where
	//
	// `instance_id`: The instance ID. This is required when the data source is registered in instance mode.
	//
	// `encoded_jdbc_url`: The URL-encoded JDBC connection string. This is required when the data source is registered by using a connection string.
	//
	// `catalog_id`: The DLF catalog ID.
	//
	// `project_name`: The MaxCompute project name.
	//
	// `database_name`: The database name.
	//
	// `schema_name`: The schema name.
	//
	// example:
	//
	// dlf-database::catalog_id:database_name
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// The list of partition keys. This is empty for non-partitioned tables.
	PartitionKeys   []*string          `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty" type:"Repeated"`
	StatisticsInfos map[string]*string `json:"StatisticsInfos,omitempty" xml:"StatisticsInfos,omitempty"`
	// The table type. The valid values depend on the metadata collector type.
	//
	// example:
	//
	// TABLE
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
	// The technical metadata.
	TechnicalMetadata *TableTechnicalMetadata `json:"TechnicalMetadata,omitempty" xml:"TechnicalMetadata,omitempty" type:"Struct"`
}

func (s Table) String() string {
	return dara.Prettify(s)
}

func (s Table) GoString() string {
	return s.String()
}

func (s *Table) GetBusinessMetadata() *TableBusinessMetadata {
	return s.BusinessMetadata
}

func (s *Table) GetComment() *string {
	return s.Comment
}

func (s *Table) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *Table) GetId() *string {
	return s.Id
}

func (s *Table) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *Table) GetName() *string {
	return s.Name
}

func (s *Table) GetParentMetaEntityId() *string {
	return s.ParentMetaEntityId
}

func (s *Table) GetPartitionKeys() []*string {
	return s.PartitionKeys
}

func (s *Table) GetStatisticsInfos() map[string]*string {
	return s.StatisticsInfos
}

func (s *Table) GetTableType() *string {
	return s.TableType
}

func (s *Table) GetTechnicalMetadata() *TableTechnicalMetadata {
	return s.TechnicalMetadata
}

func (s *Table) SetBusinessMetadata(v *TableBusinessMetadata) *Table {
	s.BusinessMetadata = v
	return s
}

func (s *Table) SetComment(v string) *Table {
	s.Comment = &v
	return s
}

func (s *Table) SetCreateTime(v int64) *Table {
	s.CreateTime = &v
	return s
}

func (s *Table) SetId(v string) *Table {
	s.Id = &v
	return s
}

func (s *Table) SetModifyTime(v int64) *Table {
	s.ModifyTime = &v
	return s
}

func (s *Table) SetName(v string) *Table {
	s.Name = &v
	return s
}

func (s *Table) SetParentMetaEntityId(v string) *Table {
	s.ParentMetaEntityId = &v
	return s
}

func (s *Table) SetPartitionKeys(v []*string) *Table {
	s.PartitionKeys = v
	return s
}

func (s *Table) SetStatisticsInfos(v map[string]*string) *Table {
	s.StatisticsInfos = v
	return s
}

func (s *Table) SetTableType(v string) *Table {
	s.TableType = &v
	return s
}

func (s *Table) SetTechnicalMetadata(v *TableTechnicalMetadata) *Table {
	s.TechnicalMetadata = v
	return s
}

func (s *Table) Validate() error {
	if s.BusinessMetadata != nil {
		if err := s.BusinessMetadata.Validate(); err != nil {
			return err
		}
	}
	if s.TechnicalMetadata != nil {
		if err := s.TechnicalMetadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TableBusinessMetadata struct {
	// The list of categories.
	Categories [][]*TableBusinessMetadataCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The custom attribute values, where key is the custom attribute identifier and value is the list of attribute values.
	CustomAttributes map[string][]*string `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty"`
	// The extension information. Currently only supported for MaxCompute type.
	Extension *TableBusinessMetadataExtension `json:"Extension,omitempty" xml:"Extension,omitempty" type:"Struct"`
	// The usage instructions.
	//
	// example:
	//
	// ## 使用说明
	Readme *string `json:"Readme,omitempty" xml:"Readme,omitempty"`
	// The list of tags.
	Tags []*TableBusinessMetadataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The list of upstream nodes.
	UpstreamTasks []*TableBusinessMetadataUpstreamTasks `json:"UpstreamTasks,omitempty" xml:"UpstreamTasks,omitempty" type:"Repeated"`
}

func (s TableBusinessMetadata) String() string {
	return dara.Prettify(s)
}

func (s TableBusinessMetadata) GoString() string {
	return s.String()
}

func (s *TableBusinessMetadata) GetCategories() [][]*TableBusinessMetadataCategories {
	return s.Categories
}

func (s *TableBusinessMetadata) GetCustomAttributes() map[string][]*string {
	return s.CustomAttributes
}

func (s *TableBusinessMetadata) GetExtension() *TableBusinessMetadataExtension {
	return s.Extension
}

func (s *TableBusinessMetadata) GetReadme() *string {
	return s.Readme
}

func (s *TableBusinessMetadata) GetTags() []*TableBusinessMetadataTags {
	return s.Tags
}

func (s *TableBusinessMetadata) GetUpstreamTasks() []*TableBusinessMetadataUpstreamTasks {
	return s.UpstreamTasks
}

func (s *TableBusinessMetadata) SetCategories(v [][]*TableBusinessMetadataCategories) *TableBusinessMetadata {
	s.Categories = v
	return s
}

func (s *TableBusinessMetadata) SetCustomAttributes(v map[string][]*string) *TableBusinessMetadata {
	s.CustomAttributes = v
	return s
}

func (s *TableBusinessMetadata) SetExtension(v *TableBusinessMetadataExtension) *TableBusinessMetadata {
	s.Extension = v
	return s
}

func (s *TableBusinessMetadata) SetReadme(v string) *TableBusinessMetadata {
	s.Readme = &v
	return s
}

func (s *TableBusinessMetadata) SetTags(v []*TableBusinessMetadataTags) *TableBusinessMetadata {
	s.Tags = v
	return s
}

func (s *TableBusinessMetadata) SetUpstreamTasks(v []*TableBusinessMetadataUpstreamTasks) *TableBusinessMetadata {
	s.UpstreamTasks = v
	return s
}

func (s *TableBusinessMetadata) Validate() error {
	if s.Extension != nil {
		if err := s.Extension.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UpstreamTasks != nil {
		for _, item := range s.UpstreamTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TableBusinessMetadataCategories struct {
	// The category ID.
	//
	// example:
	//
	// CATEGORY.456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name.
	//
	// example:
	//
	// 测试类目
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parent category ID. This parameter can be empty.
	//
	// example:
	//
	// CATEGORY.123
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s TableBusinessMetadataCategories) String() string {
	return dara.Prettify(s)
}

func (s TableBusinessMetadataCategories) GoString() string {
	return s.String()
}

func (s *TableBusinessMetadataCategories) GetId() *string {
	return s.Id
}

func (s *TableBusinessMetadataCategories) GetName() *string {
	return s.Name
}

func (s *TableBusinessMetadataCategories) GetParentId() *string {
	return s.ParentId
}

func (s *TableBusinessMetadataCategories) SetId(v string) *TableBusinessMetadataCategories {
	s.Id = &v
	return s
}

func (s *TableBusinessMetadataCategories) SetName(v string) *TableBusinessMetadataCategories {
	s.Name = &v
	return s
}

func (s *TableBusinessMetadataCategories) SetParentId(v string) *TableBusinessMetadataCategories {
	s.ParentId = &v
	return s
}

func (s *TableBusinessMetadataCategories) Validate() error {
	return dara.Validate(s)
}

type TableBusinessMetadataExtension struct {
	// The environment type. Valid values:
	//
	// - Prod: Production environment.
	//
	// - Dev: Development environment.
	//
	// example:
	//
	// Dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The favorite count.
	//
	// example:
	//
	// 0
	FavorCount *int64 `json:"FavorCount,omitempty" xml:"FavorCount,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The read count.
	//
	// example:
	//
	// 0
	ReadCount *int64 `json:"ReadCount,omitempty" xml:"ReadCount,omitempty"`
	// The view count.
	//
	// example:
	//
	// 0
	ViewCount *int64 `json:"ViewCount,omitempty" xml:"ViewCount,omitempty"`
}

func (s TableBusinessMetadataExtension) String() string {
	return dara.Prettify(s)
}

func (s TableBusinessMetadataExtension) GoString() string {
	return s.String()
}

func (s *TableBusinessMetadataExtension) GetEnvType() *string {
	return s.EnvType
}

func (s *TableBusinessMetadataExtension) GetFavorCount() *int64 {
	return s.FavorCount
}

func (s *TableBusinessMetadataExtension) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *TableBusinessMetadataExtension) GetReadCount() *int64 {
	return s.ReadCount
}

func (s *TableBusinessMetadataExtension) GetViewCount() *int64 {
	return s.ViewCount
}

func (s *TableBusinessMetadataExtension) SetEnvType(v string) *TableBusinessMetadataExtension {
	s.EnvType = &v
	return s
}

func (s *TableBusinessMetadataExtension) SetFavorCount(v int64) *TableBusinessMetadataExtension {
	s.FavorCount = &v
	return s
}

func (s *TableBusinessMetadataExtension) SetProjectId(v int64) *TableBusinessMetadataExtension {
	s.ProjectId = &v
	return s
}

func (s *TableBusinessMetadataExtension) SetReadCount(v int64) *TableBusinessMetadataExtension {
	s.ReadCount = &v
	return s
}

func (s *TableBusinessMetadataExtension) SetViewCount(v int64) *TableBusinessMetadataExtension {
	s.ViewCount = &v
	return s
}

func (s *TableBusinessMetadataExtension) Validate() error {
	return dara.Validate(s)
}

type TableBusinessMetadataTags struct {
	// The tag key. This parameter cannot be empty.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. This parameter can be empty.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// tag_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TableBusinessMetadataTags) String() string {
	return dara.Prettify(s)
}

func (s TableBusinessMetadataTags) GoString() string {
	return s.String()
}

func (s *TableBusinessMetadataTags) GetKey() *string {
	return s.Key
}

func (s *TableBusinessMetadataTags) GetValue() *string {
	return s.Value
}

func (s *TableBusinessMetadataTags) SetKey(v string) *TableBusinessMetadataTags {
	s.Key = &v
	return s
}

func (s *TableBusinessMetadataTags) SetValue(v string) *TableBusinessMetadataTags {
	s.Value = &v
	return s
}

func (s *TableBusinessMetadataTags) Validate() error {
	return dara.Validate(s)
}

type TableBusinessMetadataUpstreamTasks struct {
	// The node ID.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The node name.
	//
	// example:
	//
	// test_task
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s TableBusinessMetadataUpstreamTasks) String() string {
	return dara.Prettify(s)
}

func (s TableBusinessMetadataUpstreamTasks) GoString() string {
	return s.String()
}

func (s *TableBusinessMetadataUpstreamTasks) GetId() *int64 {
	return s.Id
}

func (s *TableBusinessMetadataUpstreamTasks) GetName() *string {
	return s.Name
}

func (s *TableBusinessMetadataUpstreamTasks) SetId(v int64) *TableBusinessMetadataUpstreamTasks {
	s.Id = &v
	return s
}

func (s *TableBusinessMetadataUpstreamTasks) SetName(v string) *TableBusinessMetadataUpstreamTasks {
	s.Name = &v
	return s
}

func (s *TableBusinessMetadataUpstreamTasks) Validate() error {
	return dara.Validate(s)
}

type TableTechnicalMetadata struct {
	// Indicates whether the table is compressed.
	//
	// example:
	//
	// false
	Compressed *bool `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	// The input format.
	//
	// example:
	//
	// org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat
	InputFormat *string `json:"InputFormat,omitempty" xml:"InputFormat,omitempty"`
	// The storage location.
	//
	// example:
	//
	// oss://test-bucket/test_tbl
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The output format.
	//
	// example:
	//
	// org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// The owner.
	//
	// example:
	//
	// test_user
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The parameter information.
	Parameters map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The class used by SerDe.
	//
	// example:
	//
	// org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe
	SerializationLibrary *string `json:"SerializationLibrary,omitempty" xml:"SerializationLibrary,omitempty"`
}

func (s TableTechnicalMetadata) String() string {
	return dara.Prettify(s)
}

func (s TableTechnicalMetadata) GoString() string {
	return s.String()
}

func (s *TableTechnicalMetadata) GetCompressed() *bool {
	return s.Compressed
}

func (s *TableTechnicalMetadata) GetInputFormat() *string {
	return s.InputFormat
}

func (s *TableTechnicalMetadata) GetLocation() *string {
	return s.Location
}

func (s *TableTechnicalMetadata) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *TableTechnicalMetadata) GetOwner() *string {
	return s.Owner
}

func (s *TableTechnicalMetadata) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *TableTechnicalMetadata) GetSerializationLibrary() *string {
	return s.SerializationLibrary
}

func (s *TableTechnicalMetadata) SetCompressed(v bool) *TableTechnicalMetadata {
	s.Compressed = &v
	return s
}

func (s *TableTechnicalMetadata) SetInputFormat(v string) *TableTechnicalMetadata {
	s.InputFormat = &v
	return s
}

func (s *TableTechnicalMetadata) SetLocation(v string) *TableTechnicalMetadata {
	s.Location = &v
	return s
}

func (s *TableTechnicalMetadata) SetOutputFormat(v string) *TableTechnicalMetadata {
	s.OutputFormat = &v
	return s
}

func (s *TableTechnicalMetadata) SetOwner(v string) *TableTechnicalMetadata {
	s.Owner = &v
	return s
}

func (s *TableTechnicalMetadata) SetParameters(v map[string]*string) *TableTechnicalMetadata {
	s.Parameters = v
	return s
}

func (s *TableTechnicalMetadata) SetSerializationLibrary(v string) *TableTechnicalMetadata {
	s.SerializationLibrary = &v
	return s
}

func (s *TableTechnicalMetadata) Validate() error {
	return dara.Validate(s)
}
