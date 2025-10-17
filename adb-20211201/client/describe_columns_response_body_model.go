// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeColumnsResponseBodyItems) *DescribeColumnsResponseBody
	GetItems() *DescribeColumnsResponseBodyItems
	SetRequestId(v string) *DescribeColumnsResponseBody
	GetRequestId() *string
}

type DescribeColumnsResponseBody struct {
	// The queried columns.
	Items *DescribeColumnsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-XXX442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeColumnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBody) GetItems() *DescribeColumnsResponseBodyItems {
	return s.Items
}

func (s *DescribeColumnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeColumnsResponseBody) SetItems(v *DescribeColumnsResponseBodyItems) *DescribeColumnsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeColumnsResponseBody) SetRequestId(v string) *DescribeColumnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeColumnsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeColumnsResponseBodyItems struct {
	Column []*DescribeColumnsResponseBodyItemsColumn `json:"Column,omitempty" xml:"Column,omitempty" type:"Repeated"`
}

func (s DescribeColumnsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBodyItems) GetColumn() []*DescribeColumnsResponseBodyItemsColumn {
	return s.Column
}

func (s *DescribeColumnsResponseBodyItems) SetColumn(v []*DescribeColumnsResponseBodyItemsColumn) *DescribeColumnsResponseBodyItems {
	s.Column = v
	return s
}

func (s *DescribeColumnsResponseBodyItems) Validate() error {
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

type DescribeColumnsResponseBodyItemsColumn struct {
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
	// The name of the column.
	//
	// example:
	//
	// id
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// example:
	//
	// amv-bp111m2cfrdl1****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Indicates whether the column is the primary key of the table. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	PrimaryKey *bool `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adb_demo
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The data type of the column.
	//
	// example:
	//
	// bigint
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeColumnsResponseBodyItemsColumn) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnsResponseBodyItemsColumn) GoString() string {
	return s.String()
}

func (s *DescribeColumnsResponseBodyItemsColumn) GetAutoIncrementColumn() *bool {
	return s.AutoIncrementColumn
}

func (s *DescribeColumnsResponseBodyItemsColumn) GetColumnName() *string {
	return s.ColumnName
}

func (s *DescribeColumnsResponseBodyItemsColumn) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeColumnsResponseBodyItemsColumn) GetPrimaryKey() *bool {
	return s.PrimaryKey
}

func (s *DescribeColumnsResponseBodyItemsColumn) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeColumnsResponseBodyItemsColumn) GetTableName() *string {
	return s.TableName
}

func (s *DescribeColumnsResponseBodyItemsColumn) GetType() *string {
	return s.Type
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetAutoIncrementColumn(v bool) *DescribeColumnsResponseBodyItemsColumn {
	s.AutoIncrementColumn = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetColumnName(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.ColumnName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetDBClusterId(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.DBClusterId = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetPrimaryKey(v bool) *DescribeColumnsResponseBodyItemsColumn {
	s.PrimaryKey = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetSchemaName(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.SchemaName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetTableName(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.TableName = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) SetType(v string) *DescribeColumnsResponseBodyItemsColumn {
	s.Type = &v
	return s
}

func (s *DescribeColumnsResponseBodyItemsColumn) Validate() error {
	return dara.Validate(s)
}
