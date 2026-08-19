// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iColumn interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessMetadata(v *ColumnBusinessMetadata) *Column
	GetBusinessMetadata() *ColumnBusinessMetadata
	SetComment(v string) *Column
	GetComment() *string
	SetForeignKey(v bool) *Column
	GetForeignKey() *bool
	SetId(v string) *Column
	GetId() *string
	SetName(v string) *Column
	GetName() *string
	SetPartitionKey(v bool) *Column
	GetPartitionKey() *bool
	SetPosition(v int32) *Column
	GetPosition() *int32
	SetPrimaryKey(v bool) *Column
	GetPrimaryKey() *bool
	SetStatisticsInfos(v map[string]*string) *Column
	GetStatisticsInfos() map[string]*string
	SetTableId(v string) *Column
	GetTableId() *string
	SetType(v string) *Column
	GetType() *string
}

type Column struct {
	// The business metadata.
	BusinessMetadata *ColumnBusinessMetadata `json:"BusinessMetadata,omitempty" xml:"BusinessMetadata,omitempty" type:"Struct"`
	// The comment.
	//
	// example:
	//
	// 字段1
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Indicates whether the column is a foreign key. Currently, only MaxCompute is supported.
	//
	// example:
	//
	// false
	ForeignKey *bool `json:"ForeignKey,omitempty" xml:"ForeignKey,omitempty"`
	// The ID. For more information, see [Metadata entity concepts](https://help.aliyun.com/document_detail/2880092.html).
	//
	// The format is `${EntityType}:${instance ID or encoded URL}:${DataCatalogIdentity}:${DatabaseName}:${PatternName}:${TableName}:${ColumnName}`. Use an empty character as a placeholder for levels that do not exist.
	//
	// > For MaxCompute and DLF types, use an empty string as a placeholder for the instance ID. For MaxCompute, the database name is the MaxCompute project name. Projects with the three-layer model enabled must include the schema name. For projects without the three-layer model enabled, use an empty string as a placeholder for the schema name.
	//
	// > For StarRocks, the data catalog identifier is the catalog name. For DLF, the data catalog identifier is the catalog ID. Other types do not support the catalog level, and you can use an empty string as a placeholder.
	//
	// The following examples show the ID formats for several common types:
	//
	// `maxcompute-column:::project_name:[schema_name]:table_name:column_name`
	//
	// `dlf-column::catalog_id:database_name::table_name:column_name`
	//
	// `hms-column:instance_id::database_name::table_name:column_name`
	//
	// `holo-column:instance_id::database_name:schema_name:table_name:column_name`
	//
	// `mysql-column:(instance_id|encoded_jdbc_url)::database_name::table_name:column_name`
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
	// `schema_name`: The schema name. For MaxCompute, this is required only when the three-layer model is enabled for the project. If the three-layer model is not enabled, use an empty string as a placeholder.
	//
	// `table_name`: The table name.
	//
	// `column_name`: The column name.
	//
	// example:
	//
	// maxcompute-column:123456::test_project:default:test_tbl:col1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name.
	//
	// example:
	//
	// col1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the column is a partition key.
	//
	// example:
	//
	// false
	PartitionKey *bool `json:"PartitionKey,omitempty" xml:"PartitionKey,omitempty"`
	// The position.
	//
	// example:
	//
	// 1
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
	// Indicates whether the column is a primary key. Currently, only MaxCompute is supported.
	//
	// example:
	//
	// false
	PrimaryKey      *bool              `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	StatisticsInfos map[string]*string `json:"StatisticsInfos,omitempty" xml:"StatisticsInfos,omitempty"`
	// The table ID. For more information, see the `Table` object.
	//
	// example:
	//
	// maxcompute-table:123456::test_project:default:test_tbl
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
	// The type.
	//
	// example:
	//
	// bigint
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s Column) String() string {
	return dara.Prettify(s)
}

func (s Column) GoString() string {
	return s.String()
}

func (s *Column) GetBusinessMetadata() *ColumnBusinessMetadata {
	return s.BusinessMetadata
}

func (s *Column) GetComment() *string {
	return s.Comment
}

func (s *Column) GetForeignKey() *bool {
	return s.ForeignKey
}

func (s *Column) GetId() *string {
	return s.Id
}

func (s *Column) GetName() *string {
	return s.Name
}

func (s *Column) GetPartitionKey() *bool {
	return s.PartitionKey
}

func (s *Column) GetPosition() *int32 {
	return s.Position
}

func (s *Column) GetPrimaryKey() *bool {
	return s.PrimaryKey
}

func (s *Column) GetStatisticsInfos() map[string]*string {
	return s.StatisticsInfos
}

func (s *Column) GetTableId() *string {
	return s.TableId
}

func (s *Column) GetType() *string {
	return s.Type
}

func (s *Column) SetBusinessMetadata(v *ColumnBusinessMetadata) *Column {
	s.BusinessMetadata = v
	return s
}

func (s *Column) SetComment(v string) *Column {
	s.Comment = &v
	return s
}

func (s *Column) SetForeignKey(v bool) *Column {
	s.ForeignKey = &v
	return s
}

func (s *Column) SetId(v string) *Column {
	s.Id = &v
	return s
}

func (s *Column) SetName(v string) *Column {
	s.Name = &v
	return s
}

func (s *Column) SetPartitionKey(v bool) *Column {
	s.PartitionKey = &v
	return s
}

func (s *Column) SetPosition(v int32) *Column {
	s.Position = &v
	return s
}

func (s *Column) SetPrimaryKey(v bool) *Column {
	s.PrimaryKey = &v
	return s
}

func (s *Column) SetStatisticsInfos(v map[string]*string) *Column {
	s.StatisticsInfos = v
	return s
}

func (s *Column) SetTableId(v string) *Column {
	s.TableId = &v
	return s
}

func (s *Column) SetType(v string) *Column {
	s.Type = &v
	return s
}

func (s *Column) Validate() error {
	if s.BusinessMetadata != nil {
		if err := s.BusinessMetadata.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ColumnBusinessMetadata struct {
	// The custom attribute values, where key is the custom attribute identifier and value is the attribute value list.
	CustomAttributes map[string][]*string `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty"`
	// The business description of the field. Currently, only MaxCompute, HMS (EMR cluster), and DLF types are supported.
	//
	// example:
	//
	// 字段1的业务描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ColumnBusinessMetadata) String() string {
	return dara.Prettify(s)
}

func (s ColumnBusinessMetadata) GoString() string {
	return s.String()
}

func (s *ColumnBusinessMetadata) GetCustomAttributes() map[string][]*string {
	return s.CustomAttributes
}

func (s *ColumnBusinessMetadata) GetDescription() *string {
	return s.Description
}

func (s *ColumnBusinessMetadata) SetCustomAttributes(v map[string][]*string) *ColumnBusinessMetadata {
	s.CustomAttributes = v
	return s
}

func (s *ColumnBusinessMetadata) SetDescription(v string) *ColumnBusinessMetadata {
	s.Description = &v
	return s
}

func (s *ColumnBusinessMetadata) Validate() error {
	return dara.Validate(s)
}
