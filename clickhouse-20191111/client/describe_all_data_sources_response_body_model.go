// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColumns(v *DescribeAllDataSourcesResponseBodyColumns) *DescribeAllDataSourcesResponseBody
	GetColumns() *DescribeAllDataSourcesResponseBodyColumns
	SetRequestId(v string) *DescribeAllDataSourcesResponseBody
	GetRequestId() *string
	SetSchemas(v *DescribeAllDataSourcesResponseBodySchemas) *DescribeAllDataSourcesResponseBody
	GetSchemas() *DescribeAllDataSourcesResponseBodySchemas
	SetTables(v *DescribeAllDataSourcesResponseBodyTables) *DescribeAllDataSourcesResponseBody
	GetTables() *DescribeAllDataSourcesResponseBodyTables
}

type DescribeAllDataSourcesResponseBody struct {
	// Details of the columns.
	Columns *DescribeAllDataSourcesResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 75EA41D7-025A-50A6-9287-359A91399F1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the databases.
	Schemas *DescribeAllDataSourcesResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Struct"`
	// The information about the tables.
	Tables *DescribeAllDataSourcesResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
}

func (s DescribeAllDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponseBody) GetColumns() *DescribeAllDataSourcesResponseBodyColumns {
	return s.Columns
}

func (s *DescribeAllDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAllDataSourcesResponseBody) GetSchemas() *DescribeAllDataSourcesResponseBodySchemas {
	return s.Schemas
}

func (s *DescribeAllDataSourcesResponseBody) GetTables() *DescribeAllDataSourcesResponseBodyTables {
	return s.Tables
}

func (s *DescribeAllDataSourcesResponseBody) SetColumns(v *DescribeAllDataSourcesResponseBodyColumns) *DescribeAllDataSourcesResponseBody {
	s.Columns = v
	return s
}

func (s *DescribeAllDataSourcesResponseBody) SetRequestId(v string) *DescribeAllDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBody) SetSchemas(v *DescribeAllDataSourcesResponseBodySchemas) *DescribeAllDataSourcesResponseBody {
	s.Schemas = v
	return s
}

func (s *DescribeAllDataSourcesResponseBody) SetTables(v *DescribeAllDataSourcesResponseBodyTables) *DescribeAllDataSourcesResponseBody {
	s.Tables = v
	return s
}

func (s *DescribeAllDataSourcesResponseBody) Validate() error {
	if s.Columns != nil {
		if err := s.Columns.Validate(); err != nil {
			return err
		}
	}
	if s.Schemas != nil {
		if err := s.Schemas.Validate(); err != nil {
			return err
		}
	}
	if s.Tables != nil {
		if err := s.Tables.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAllDataSourcesResponseBodyColumns struct {
	Column []*DescribeAllDataSourcesResponseBodyColumnsColumn `json:"Column,omitempty" xml:"Column,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourcesResponseBodyColumns) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourcesResponseBodyColumns) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponseBodyColumns) GetColumn() []*DescribeAllDataSourcesResponseBodyColumnsColumn {
	return s.Column
}

func (s *DescribeAllDataSourcesResponseBodyColumns) SetColumn(v []*DescribeAllDataSourcesResponseBodyColumnsColumn) *DescribeAllDataSourcesResponseBodyColumns {
	s.Column = v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyColumns) Validate() error {
	if s.Column != nil {
		for _, item := range s.Column {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAllDataSourcesResponseBodyColumnsColumn struct {
	// Indicates whether the column is an auto-increment column. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	AutoIncrementColumn *bool `json:"AutoIncrementColumn,omitempty" xml:"AutoIncrementColumn,omitempty"`
	// The column name.
	//
	// example:
	//
	// name
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Indicates whether the column is the primary key of the table. Valid values:
	//
	// 	- **true**: The column is the primary key of the table.
	//
	// 	- **false**: The column is not the primary key of the table.
	//
	// example:
	//
	// true
	PrimaryKey *bool `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name.
	//
	// example:
	//
	// table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The column type.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAllDataSourcesResponseBodyColumnsColumn) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourcesResponseBodyColumnsColumn) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) GetAutoIncrementColumn() *bool {
	return s.AutoIncrementColumn
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) GetColumnName() *string {
	return s.ColumnName
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) GetPrimaryKey() *bool {
	return s.PrimaryKey
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) GetType() *string {
	return s.Type
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) SetAutoIncrementColumn(v bool) *DescribeAllDataSourcesResponseBodyColumnsColumn {
	s.AutoIncrementColumn = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) SetColumnName(v string) *DescribeAllDataSourcesResponseBodyColumnsColumn {
	s.ColumnName = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) SetDBClusterId(v string) *DescribeAllDataSourcesResponseBodyColumnsColumn {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) SetPrimaryKey(v bool) *DescribeAllDataSourcesResponseBodyColumnsColumn {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) SetSchemaName(v string) *DescribeAllDataSourcesResponseBodyColumnsColumn {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) SetTableName(v string) *DescribeAllDataSourcesResponseBodyColumnsColumn {
	s.TableName = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) SetType(v string) *DescribeAllDataSourcesResponseBodyColumnsColumn {
	s.Type = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyColumnsColumn) Validate() error {
	return dara.Validate(s)
}

type DescribeAllDataSourcesResponseBodySchemas struct {
	Schema []*DescribeAllDataSourcesResponseBodySchemasSchema `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourcesResponseBodySchemas) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourcesResponseBodySchemas) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponseBodySchemas) GetSchema() []*DescribeAllDataSourcesResponseBodySchemasSchema {
	return s.Schema
}

func (s *DescribeAllDataSourcesResponseBodySchemas) SetSchema(v []*DescribeAllDataSourcesResponseBodySchemasSchema) *DescribeAllDataSourcesResponseBodySchemas {
	s.Schema = v
	return s
}

func (s *DescribeAllDataSourcesResponseBodySchemas) Validate() error {
	if s.Schema != nil {
		for _, item := range s.Schema {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAllDataSourcesResponseBodySchemasSchema struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
}

func (s DescribeAllDataSourcesResponseBodySchemasSchema) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourcesResponseBodySchemasSchema) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponseBodySchemasSchema) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAllDataSourcesResponseBodySchemasSchema) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeAllDataSourcesResponseBodySchemasSchema) SetDBClusterId(v string) *DescribeAllDataSourcesResponseBodySchemasSchema {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodySchemasSchema) SetSchemaName(v string) *DescribeAllDataSourcesResponseBodySchemasSchema {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodySchemasSchema) Validate() error {
	return dara.Validate(s)
}

type DescribeAllDataSourcesResponseBodyTables struct {
	Table []*DescribeAllDataSourcesResponseBodyTablesTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourcesResponseBodyTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourcesResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponseBodyTables) GetTable() []*DescribeAllDataSourcesResponseBodyTablesTable {
	return s.Table
}

func (s *DescribeAllDataSourcesResponseBodyTables) SetTable(v []*DescribeAllDataSourcesResponseBodyTablesTable) *DescribeAllDataSourcesResponseBodyTables {
	s.Table = v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyTables) Validate() error {
	if s.Table != nil {
		for _, item := range s.Table {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAllDataSourcesResponseBodyTablesTable struct {
	// The cluster ID.
	//
	// example:
	//
	// cc-bp108z124a8o7****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The database name.
	//
	// example:
	//
	// database
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The table name.
	//
	// example:
	//
	// table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAllDataSourcesResponseBodyTablesTable) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourcesResponseBodyTablesTable) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourcesResponseBodyTablesTable) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAllDataSourcesResponseBodyTablesTable) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeAllDataSourcesResponseBodyTablesTable) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAllDataSourcesResponseBodyTablesTable) SetDBClusterId(v string) *DescribeAllDataSourcesResponseBodyTablesTable {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyTablesTable) SetSchemaName(v string) *DescribeAllDataSourcesResponseBodyTablesTable {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyTablesTable) SetTableName(v string) *DescribeAllDataSourcesResponseBodyTablesTable {
	s.TableName = &v
	return s
}

func (s *DescribeAllDataSourcesResponseBodyTablesTable) Validate() error {
	return dara.Validate(s)
}
