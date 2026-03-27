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
	// The information about the business metadata that is related to DataWorks, including the usage notes, tags, categories, ancestor tasks, and extended information.
	BusinessMetadata *TableBusinessMetadata `json:"BusinessMetadata,omitempty" xml:"BusinessMetadata,omitempty" type:"Struct"`
	// The comments.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The creation time. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1736852168000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The table ID. For more information, see [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// The common format of this parameter is `${Entity type}:${Instance ID or escaped URL}:${Catalog identifier}:${Database name}:${Schema name}:${Table name}`. If a level does not exist, specify an empty string as a placeholder.
	//
	// >  For MaxCompute and DLF tables, specify an empty string at the Instance ID level as a placeholder. For MaxCompute tables, specify a MaxCompute project name at the Database name level. If the three-layer model is enabled for your MaxCompute project, you must specify a schema name at the Schema name level. Otherwise, you can specify an empty string at the Schema name level as a placeholder.
	//
	// >  For StarRocks tables, specify a catalog name at the Catalog identifier level. For DLF tables, specify a catalog ID at the Catalog identifier level. Other types of tables do not support the Catalog identifier level, and you can specify an empty string as a placeholder.
	//
	// You can configure this parameter in one of the following formats based on your table type:
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
	// > \\
	//
	// `instance_id`: the ID of an instance. If the related data source is added to DataWorks in Alibaba Cloud instance mode, you must configure this parameter.\\
	//
	// `encoded_jdbc_url`: the JDBC connection string that is URL-encoded. If the related data source is added to DataWorks in connection string mode, you must configure this parameter.\\
	//
	// `catalog_id`: the ID of a DLF catalog.\\
	//
	// `project_name`: the name of a MaxCompute project.\\
	//
	// `database_name`: the name of a database.\\
	//
	// `schema_name`: the name of a schema. For a MaxCompute table, this parameter is required only if the three-layer model is enabled for the MaxCompute project to which the table belongs. If the schema feature is not enabled for the MaxCompute project, specify an empty string for this parameter as a placeholder.\\
	//
	// `table_name`: the name of a table.
	//
	// example:
	//
	// maxcompute-table:123456XXX::test_project::test_tbl
	//
	// dlf-table:123456XXX:test_catalog:test_db::test_tbl
	//
	// hms-table:c-abc123xxx::test_db::test_tbl
	//
	// holo-table:h-abc123xxx::test_db:test_schema:test_tbl
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The modification time. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1736852168000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The table name.
	//
	// example:
	//
	// test_tbl
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of a parent metadata entity. For more information, see [Concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// 	- For data source types that support schemas, such as `MaxCompute, Hologres, PostgreSQL, SQL Server, HybridDB for PostgreSQL, and Oracle`, the `ParentMetaEntityId` parameter specifies the schema of the database to which the table belongs. In this case, the common format of this parameter is `${Entity type}:${Instance ID or escaped URL}:${Catalog identifier}:${Database name}:${Schema name}`. If a level does not exist, leave the level empty. For a MaxCompute data table, you must make sure that the three-layer model is enabled for the MaxCompute project to which the table belongs.
	//
	// 	- For other data source types that do not support schemas, the `ParentMetaEntityId` parameter specifies the database to which the table belongs. In this case, the common format of this parameter is `${Entity type}:${Instance ID or escaped URL}:${Catalog identifier}:${Database name}`. If a level does not exist, leave the level empty.
	//
	// >  For MaxCompute and DLF tables, specify an empty string at the Instance ID level as a placeholder. For MaxCompute tables, specify a MaxCompute project name at the Database name level.
	//
	// >  For StarRocks tables, specify a catalog name at the Catalog identifier level. For DLF tables, specify a catalog ID at the Catalog identifier level. Other types of tables do not support the Catalog identifier level, and you can specify an empty string as a placeholder.
	//
	// You can configure this parameter in one of the following formats based on your table type:
	//
	// `maxcompute-project:::project_name`
	//
	// `maxcompute-schema:::project_name:schema_name` (Three-layer model enabled for the MaxCompute project)
	//
	// `dlf-database::catalog_id:database_name`
	//
	// `hms-database:instance_id::database_name`
	//
	// `holo-schema:instance_id::database_name:schema_name`
	//
	// `mysql-database:(instance_id|encoded_jdbc_url)::database_name`
	//
	// > \\
	//
	// `instance_id`: the ID of an instance. If the related data source is added to DataWorks in Alibaba Cloud instance mode, you must configure this parameter.\\
	//
	// `encoded_jdbc_url`: the JDBC connection string that is URL-encoded. If the related data source is added to DataWorks in connection string mode, you must configure this parameter.\\
	//
	// `catalog_id`: the ID of a DLF catalog.\\
	//
	// `project_name`: the name of a MaxCompute project.\\
	//
	// `database_name`: the name of a database.\\
	//
	// `schema_name`: the name of a schema.
	//
	// example:
	//
	// maxcompute-schema:123456XXX::test_project_with_schema:default
	//
	// maxcompute-project:123456XXX::test_project_without_schema
	//
	// dlf-database:123456XXX:test_catalog:test_db
	//
	// hms-database:c-abc123xxx::test_db
	//
	// holo-schema:h-abc123xxx::test_db:test_schema
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
	// The partition keys. If the table is a non-partitioned table, leave this parameter empty.
	PartitionKeys []*string `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty" type:"Repeated"`
	// The table type. The value of this parameter is related to the type of metadata crawler.
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
	// The categories.
	Categories [][]*TableBusinessMetadataCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The extended information. Only MaxCompute tables supports this parameter.
	Extension *TableBusinessMetadataExtension `json:"Extension,omitempty" xml:"Extension,omitempty" type:"Struct"`
	// The usage notes.
	Readme *string `json:"Readme,omitempty" xml:"Readme,omitempty"`
	// The tags.
	Tags []*TableBusinessMetadataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ancestor tasks.
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
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The parent category ID. You can leave this parameter empty.
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
	// The type of the environment. Valid values:
	//
	// 	- Prod
	//
	// 	- Dev
	//
	// example:
	//
	// Dev
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The number of times the table is added to favorites.
	//
	// example:
	//
	// 0
	FavorCount *int64 `json:"FavorCount,omitempty" xml:"FavorCount,omitempty"`
	// The DataWorks workspace ID.
	//
	// example:
	//
	// 234
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The number of times the table is read.
	//
	// example:
	//
	// 0
	ReadCount *int64 `json:"ReadCount,omitempty" xml:"ReadCount,omitempty"`
	// The number of times the table is viewed.
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
	// The tag key. You cannot leave this parameter empty.
	//
	// example:
	//
	// tag_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. You can leave this parameter empty.
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
	// The ancestor task ID.
	//
	// example:
	//
	// 123456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ancestor task name.
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
	// Specifies whether the table is a compressed table. Valid values: true and false.
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
	// The storage location of the table.
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
	// The information about parameters.
	Parameters map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The implementation class of SerDe.
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
