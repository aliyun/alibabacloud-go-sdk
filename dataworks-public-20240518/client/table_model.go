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
	SetTableType(v string) *Table
	GetTableType() *string
	SetTechnicalMetadata(v *TableTechnicalMetadata) *Table
	GetTechnicalMetadata() *TableTechnicalMetadata
}

type Table struct {
	// The business metadata. This parameter is specific to DataWorks and includes instructions, tags, categories, upstream tasks, and extended information.
	BusinessMetadata *TableBusinessMetadata `json:"BusinessMetadata,omitempty" xml:"BusinessMetadata,omitempty" type:"Struct"`
	// The comment on the table.
	//
	// example:
	//
	// 测试表
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The table creation time, provided as a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1736852168000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the entity. For more information, see [Metadata entity concepts](https://help.aliyun.com/document_detail/2880092.html).
	//
	// The format is `${EntityType}:${instance ID or escaped URL}:${data catalog identifier}:${database name}:${schema name}:${table name}`. Use an empty string as a placeholder for any non-existent level.
	//
	// > For MaxCompute and DLF data types, use an empty string as a placeholder for the instance ID. For MaxCompute, the database name is the MaxCompute project name. You must provide a schema name for projects where the three-layer model is enabled. If the model is not enabled, use an empty string as a placeholder for the schema name.
	//
	// > For StarRocks data types, the data catalog identifier is the catalog name. For DLF data types, the data catalog identifier is the catalog ID. Other data types do not support the catalog level. For these types, use an empty string as a placeholder.
	//
	// The following are the ID formats for several common data types:
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
	// > Placeholder descriptions:<br>
	//
	// > `instance_id`: The instance ID. This is required when the data source is registered in instance mode.<br>
	//
	// > `encoded_jdbc_url`: The URL-encoded JDBC connection string. This is required when the data source is registered by using a connection string.<br>
	//
	// > `catalog_id`: The DLF catalog ID.<br>
	//
	// > `project_name`: The MaxCompute project name.<br>
	//
	// > `database_name`: The database name.<br>
	//
	// > `schema_name`: The schema name. For the MaxCompute data type, this is required only if the project has the three-layer model enabled. Otherwise, use an empty string as a placeholder.<br>
	//
	// > `table_name`: The table name.<br><br><br><br><br><br><br>
	//
	// example:
	//
	// dlf-table::catalog_id:database_name::table_name
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time the table was last modified, provided as a Unix timestamp in milliseconds.
	//
	// example:
	//
	// 1736852168000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// table_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the parent metadata entity. For more information, see [Metadata entity concepts](https://help.aliyun.com/document_detail/2880092.html).
	//
	// - For data types that support schemas, such as `maxcompute/holo/postgresql/sqlserver/hybriddb_for_postgresql/oracle`, `ParentMetaEntityId` specifies the table\\"s database schema. For the MaxCompute data type, this applies only to MaxCompute projects with the three-layer model enabled. The format is `${EntityType}:${instance ID or escaped URL}:${data catalog identifier}:${database name}:${schema name}`. Use an empty string as a placeholder for any non-existent level.
	//
	// - For other data types, `ParentMetaEntityId` specifies the table\\"s database. The format is `${EntityType}:${instance ID or escaped URL}:${data catalog identifier}:${database name}`. Use an empty string as a placeholder for any non-existent level.
	//
	// > For MaxCompute and DLF data types, use an empty string as a placeholder for the instance ID. For the MaxCompute data type, the database name is the MaxCompute project name.
	//
	// > For StarRocks data types, the data catalog identifier is the catalog name. For DLF data types, the data catalog identifier is the catalog ID. Other data types do not support the catalog level. For these types, use an empty string as a placeholder.
	//
	// The following are the formats of `ParentMetaEntityId` for several common data types:
	//
	// `maxcompute-project:::project_name`
	//
	// `maxcompute-schema:::project_name:schema_name` (Only for projects with the three-layer model enabled)
	//
	// `dlf-database::catalog_id:database_name`
	//
	// `hms-database:instance_id::database_name`
	//
	// `holo-schema:instance_id::database_name:schema_name`
	//
	// `mysql-database:(instance_id|encoded_jdbc_url)::database_name`
	//
	// > Placeholder descriptions:<br>
	//
	// > `instance_id`: The instance ID. This is required when the data source is registered in instance mode.<br>
	//
	// > `encoded_jdbc_url`: The URL-encoded JDBC connection string. This is required when the data source is registered by using a connection string.<br>
	//
	// > `catalog_id`: The DLF catalog ID.<br>
	//
	// > `project_name`: The MaxCompute project name.<br>
	//
	// > `database_name`: The database name.<br>
	//
	// > `schema_name`: The schema name.<br><br><br><br><br><br>
	//
	// example:
	//
	// dlf-database::catalog_id:database_name
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// The list of partition keys. This parameter is empty for non-partitioned tables.
	PartitionKeys []*string `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty" type:"Repeated"`
	// The table type. The value depends on the type of metadata collector.
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
	// The list of categories to which the table belongs.
	Categories [][]*TableBusinessMetadataCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// A map of custom attribute identifiers to lists of their corresponding values.
	CustomAttributes map[string][]*string `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty"`
	// Extended information. This parameter is supported only for the MaxCompute data type.
	Extension *TableBusinessMetadataExtension `json:"Extension,omitempty" xml:"Extension,omitempty" type:"Struct"`
	// The instructions for use.
	//
	// example:
	//
	// ## 使用说明
	Readme *string `json:"Readme,omitempty" xml:"Readme,omitempty"`
	// The list of tags.
	Tags []*TableBusinessMetadataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The list of upstream tasks.
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
	// The category name.
	//
	// example:
	//
	// 测试类目
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parent category\\"s ID. This can be an empty string.
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
	// - Prod: The production environment.
	//
	// - Dev: The development environment.
	//
	// example:
	//
	// Dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The number of times the table was favorited.
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
	// The number of reads.
	//
	// example:
	//
	// 0
	ReadCount *int64 `json:"ReadCount,omitempty" xml:"ReadCount,omitempty"`
	// The number of views.
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
	// The tag key. This value cannot be empty.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. This can be an empty string.
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
	// The task ID.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The task name.
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
	// The table owner.
	//
	// example:
	//
	// test_user
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The parameter information.
	Parameters map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The class used by the serializer/deserializer (SerDe).
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
