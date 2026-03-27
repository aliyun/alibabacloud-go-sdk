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
	SetTableId(v string) *Column
	GetTableId() *string
	SetType(v string) *Column
	GetType() *string
}

type Column struct {
	// Business metadata.
	BusinessMetadata *ColumnBusinessMetadata `json:"BusinessMetadata,omitempty" xml:"BusinessMetadata,omitempty" type:"Struct"`
	// The comment.
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// Specifies whether the column is a foreign key (only supported by MaxCompute).
	//
	// example:
	//
	// false
	ForeignKey *bool `json:"ForeignKey,omitempty" xml:"ForeignKey,omitempty"`
	// The ID. For more information, see [Description of concepts related to metadata entities](https://help.aliyun.com/document_detail/2880092.html).
	//
	// The format is: `${EntityType}:${Instance ID or encoded URL}:${Catalog Identifier}:${Database name}:${Schema name}:${Table Name}:${Column name}`. Use empty strings as placeholders for non-existent hierarchy levels.
	//
	// >  For the MaxCompute and DLF types, use an empty string as the placeholder for the instance ID. For MaxCompute, the database name refers to the MaxCompute project name. If the project has schema enabled, you must specify the schema name. Otherwise, use an empty string as the placeholder for the schema name.
	//
	// >  For StarRocks, the catalog identifier is the catalog name. For DLF, it is the catalog ID. Other types do not support the catalog level and you can use an empty string as a placeholder.
	//
	// Examples of ID formats for common types are as follows:
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
	// > \\
	//
	// `instance_id`: The instance ID, required when the data source is registered in instance mode.\\
	//
	// `encoded_jdbc_url`: The URL-encoded JDBC connection string, which is required when the data source is registered via a connection string.\\
	//
	// `catalog_id`: The DLF catalog ID.\\
	//
	// `project_name`: The MaxCompute project name.\\
	//
	// `database_name`: The database name.\\
	//
	// `schema_name`: The schema name. For the MaxCompute type, this is required only if the project has enabled schema; otherwise, use an empty string as a placeholder.\\
	//
	// `table_name`: The table name.\\
	//
	// `column_name`: The field name.
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
	// Specifies whether the column is a partition key.
	//
	// example:
	//
	// false
	PartitionKey *bool `json:"PartitionKey,omitempty" xml:"PartitionKey,omitempty"`
	// The position of the field.
	//
	// example:
	//
	// 1
	Position *int32 `json:"Position,omitempty" xml:"Position,omitempty"`
	// Specifies whether the column is a primary key (only supported by MaxCompute).
	//
	// example:
	//
	// false
	PrimaryKey *bool `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The table ID. You can refer to the `Table` object.
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
	// A business-level description of the field (supported only by MaxCompute, HMS (EMR clusters) and DLF.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ColumnBusinessMetadata) String() string {
	return dara.Prettify(s)
}

func (s ColumnBusinessMetadata) GoString() string {
	return s.String()
}

func (s *ColumnBusinessMetadata) GetDescription() *string {
	return s.Description
}

func (s *ColumnBusinessMetadata) SetDescription(v string) *ColumnBusinessMetadata {
	s.Description = &v
	return s
}

func (s *ColumnBusinessMetadata) Validate() error {
	return dara.Validate(s)
}
