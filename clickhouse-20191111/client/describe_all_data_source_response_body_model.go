// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColumns(v *DescribeAllDataSourceResponseBodyColumns) *DescribeAllDataSourceResponseBody
	GetColumns() *DescribeAllDataSourceResponseBodyColumns
	SetRequestId(v string) *DescribeAllDataSourceResponseBody
	GetRequestId() *string
	SetSchemas(v *DescribeAllDataSourceResponseBodySchemas) *DescribeAllDataSourceResponseBody
	GetSchemas() *DescribeAllDataSourceResponseBodySchemas
	SetTables(v *DescribeAllDataSourceResponseBodyTables) *DescribeAllDataSourceResponseBody
	GetTables() *DescribeAllDataSourceResponseBodyTables
}

type DescribeAllDataSourceResponseBody struct {
	// The information about the columns.
	Columns *DescribeAllDataSourceResponseBodyColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the databases.
	Schemas *DescribeAllDataSourceResponseBodySchemas `json:"Schemas,omitempty" xml:"Schemas,omitempty" type:"Struct"`
	// The information about the tables.
	Tables *DescribeAllDataSourceResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Struct"`
}

func (s DescribeAllDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBody) GetColumns() *DescribeAllDataSourceResponseBodyColumns {
	return s.Columns
}

func (s *DescribeAllDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAllDataSourceResponseBody) GetSchemas() *DescribeAllDataSourceResponseBodySchemas {
	return s.Schemas
}

func (s *DescribeAllDataSourceResponseBody) GetTables() *DescribeAllDataSourceResponseBodyTables {
	return s.Tables
}

func (s *DescribeAllDataSourceResponseBody) SetColumns(v *DescribeAllDataSourceResponseBodyColumns) *DescribeAllDataSourceResponseBody {
	s.Columns = v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetRequestId(v string) *DescribeAllDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetSchemas(v *DescribeAllDataSourceResponseBodySchemas) *DescribeAllDataSourceResponseBody {
	s.Schemas = v
	return s
}

func (s *DescribeAllDataSourceResponseBody) SetTables(v *DescribeAllDataSourceResponseBodyTables) *DescribeAllDataSourceResponseBody {
	s.Tables = v
	return s
}

func (s *DescribeAllDataSourceResponseBody) Validate() error {
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

type DescribeAllDataSourceResponseBodyColumns struct {
	Column []*DescribeAllDataSourceResponseBodyColumnsColumn `json:"Column,omitempty" xml:"Column,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourceResponseBodyColumns) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyColumns) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyColumns) GetColumn() []*DescribeAllDataSourceResponseBodyColumnsColumn {
	return s.Column
}

func (s *DescribeAllDataSourceResponseBodyColumns) SetColumn(v []*DescribeAllDataSourceResponseBodyColumnsColumn) *DescribeAllDataSourceResponseBodyColumns {
	s.Column = v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumns) Validate() error {
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

type DescribeAllDataSourceResponseBodyColumnsColumn struct {
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
	// 	- **true**
	//
	// 	- **false**
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
	// The type of the column.
	//
	// example:
	//
	// String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAllDataSourceResponseBodyColumnsColumn) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyColumnsColumn) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) GetAutoIncrementColumn() *bool {
	return s.AutoIncrementColumn
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) GetColumnName() *string {
	return s.ColumnName
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) GetPrimaryKey() *bool {
	return s.PrimaryKey
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) GetType() *string {
	return s.Type
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetAutoIncrementColumn(v bool) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.AutoIncrementColumn = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetColumnName(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.ColumnName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetDBClusterId(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetPrimaryKey(v bool) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetSchemaName(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetTableName(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.TableName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) SetType(v string) *DescribeAllDataSourceResponseBodyColumnsColumn {
	s.Type = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyColumnsColumn) Validate() error {
	return dara.Validate(s)
}

type DescribeAllDataSourceResponseBodySchemas struct {
	Schema []*DescribeAllDataSourceResponseBodySchemasSchema `json:"Schema,omitempty" xml:"Schema,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourceResponseBodySchemas) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodySchemas) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodySchemas) GetSchema() []*DescribeAllDataSourceResponseBodySchemasSchema {
	return s.Schema
}

func (s *DescribeAllDataSourceResponseBodySchemas) SetSchema(v []*DescribeAllDataSourceResponseBodySchemasSchema) *DescribeAllDataSourceResponseBodySchemas {
	s.Schema = v
	return s
}

func (s *DescribeAllDataSourceResponseBodySchemas) Validate() error {
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

type DescribeAllDataSourceResponseBodySchemasSchema struct {
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

func (s DescribeAllDataSourceResponseBodySchemasSchema) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodySchemasSchema) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodySchemasSchema) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAllDataSourceResponseBodySchemasSchema) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeAllDataSourceResponseBodySchemasSchema) SetDBClusterId(v string) *DescribeAllDataSourceResponseBodySchemasSchema {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodySchemasSchema) SetSchemaName(v string) *DescribeAllDataSourceResponseBodySchemasSchema {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodySchemasSchema) Validate() error {
	return dara.Validate(s)
}

type DescribeAllDataSourceResponseBodyTables struct {
	Table []*DescribeAllDataSourceResponseBodyTablesTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeAllDataSourceResponseBodyTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyTables) GetTable() []*DescribeAllDataSourceResponseBodyTablesTable {
	return s.Table
}

func (s *DescribeAllDataSourceResponseBodyTables) SetTable(v []*DescribeAllDataSourceResponseBodyTablesTable) *DescribeAllDataSourceResponseBodyTables {
	s.Table = v
	return s
}

func (s *DescribeAllDataSourceResponseBodyTables) Validate() error {
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

type DescribeAllDataSourceResponseBodyTablesTable struct {
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

func (s DescribeAllDataSourceResponseBodyTablesTable) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourceResponseBodyTablesTable) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) SetDBClusterId(v string) *DescribeAllDataSourceResponseBodyTablesTable {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) SetSchemaName(v string) *DescribeAllDataSourceResponseBodyTablesTable {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) SetTableName(v string) *DescribeAllDataSourceResponseBodyTablesTable {
	s.TableName = &v
	return s
}

func (s *DescribeAllDataSourceResponseBodyTablesTable) Validate() error {
	return dara.Validate(s)
}
