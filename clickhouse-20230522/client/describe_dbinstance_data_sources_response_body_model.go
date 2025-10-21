// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDBInstanceDataSourcesResponseBodyData) *DescribeDBInstanceDataSourcesResponseBody
	GetData() *DescribeDBInstanceDataSourcesResponseBodyData
	SetRequestId(v string) *DescribeDBInstanceDataSourcesResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceDataSourcesResponseBody struct {
	// The returned result.
	Data *DescribeDBInstanceDataSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F543E6CC-6868-523D-8D28-0E92CF977ED2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSourcesResponseBody) GetData() *DescribeDBInstanceDataSourcesResponseBodyData {
	return s.Data
}

func (s *DescribeDBInstanceDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceDataSourcesResponseBody) SetData(v *DescribeDBInstanceDataSourcesResponseBodyData) *DescribeDBInstanceDataSourcesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBody) SetRequestId(v string) *DescribeDBInstanceDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceDataSourcesResponseBodyData struct {
	// The columns.
	Columns []*DescribeDBInstanceDataSourcesResponseBodyDataColumns `json:"Columns,omitempty" xml:"Columns,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The account.
	//
	// example:
	//
	// default
	Schemas *string `json:"Schemas,omitempty" xml:"Schemas,omitempty"`
	// The tables.
	Tables []*string `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s DescribeDBInstanceDataSourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDataSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSourcesResponseBodyData) GetColumns() []*DescribeDBInstanceDataSourcesResponseBodyDataColumns {
	return s.Columns
}

func (s *DescribeDBInstanceDataSourcesResponseBodyData) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceDataSourcesResponseBodyData) GetSchemas() *string {
	return s.Schemas
}

func (s *DescribeDBInstanceDataSourcesResponseBodyData) GetTables() []*string {
	return s.Tables
}

func (s *DescribeDBInstanceDataSourcesResponseBodyData) SetColumns(v []*DescribeDBInstanceDataSourcesResponseBodyDataColumns) *DescribeDBInstanceDataSourcesResponseBodyData {
	s.Columns = v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyData) SetDBInstanceId(v string) *DescribeDBInstanceDataSourcesResponseBodyData {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyData) SetSchemas(v string) *DescribeDBInstanceDataSourcesResponseBodyData {
	s.Schemas = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyData) SetTables(v []*string) *DescribeDBInstanceDataSourcesResponseBodyData {
	s.Tables = v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyData) Validate() error {
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

type DescribeDBInstanceDataSourcesResponseBodyDataColumns struct {
	// The column name.
	//
	// example:
	//
	// c31
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The description of the database account.
	//
	// example:
	//
	// Used for test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The database name.
	//
	// example:
	//
	// dbtest
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// Indicates whether the column is the primary key of the table. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The table name.
	//
	// example:
	//
	// tableTest
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The type of the stored data.
	//
	// example:
	//
	// UInt64
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDBInstanceDataSourcesResponseBodyDataColumns) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDataSourcesResponseBodyDataColumns) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) GetColumnName() *string {
	return s.ColumnName
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) GetComment() *string {
	return s.Comment
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) GetDBName() *string {
	return s.DBName
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) GetTableName() *string {
	return s.TableName
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) GetType() *string {
	return s.Type
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) SetColumnName(v string) *DescribeDBInstanceDataSourcesResponseBodyDataColumns {
	s.ColumnName = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) SetComment(v string) *DescribeDBInstanceDataSourcesResponseBodyDataColumns {
	s.Comment = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) SetDBName(v string) *DescribeDBInstanceDataSourcesResponseBodyDataColumns {
	s.DBName = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) SetPrimaryKey(v string) *DescribeDBInstanceDataSourcesResponseBodyDataColumns {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) SetTableName(v string) *DescribeDBInstanceDataSourcesResponseBodyDataColumns {
	s.TableName = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) SetType(v string) *DescribeDBInstanceDataSourcesResponseBodyDataColumns {
	s.Type = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesResponseBodyDataColumns) Validate() error {
	return dara.Validate(s)
}
