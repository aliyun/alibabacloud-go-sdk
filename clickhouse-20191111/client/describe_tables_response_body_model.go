// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeTablesResponseBodyItems) *DescribeTablesResponseBody
	GetItems() *DescribeTablesResponseBodyItems
	SetRequestId(v string) *DescribeTablesResponseBody
	GetRequestId() *string
}

type DescribeTablesResponseBody struct {
	// The information about the tables.
	Items *DescribeTablesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 05321590-BB65-4720-8CB6-8218E041CDD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBody) GetItems() *DescribeTablesResponseBodyItems {
	return s.Items
}

func (s *DescribeTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTablesResponseBody) SetItems(v *DescribeTablesResponseBodyItems) *DescribeTablesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeTablesResponseBody) SetRequestId(v string) *DescribeTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTablesResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTablesResponseBodyItems struct {
	Table []*DescribeTablesResponseBodyItemsTable `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeTablesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyItems) GetTable() []*DescribeTablesResponseBodyItemsTable {
	return s.Table
}

func (s *DescribeTablesResponseBodyItems) SetTable(v []*DescribeTablesResponseBodyItemsTable) *DescribeTablesResponseBodyItems {
	s.Table = v
	return s
}

func (s *DescribeTablesResponseBodyItems) Validate() error {
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

type DescribeTablesResponseBodyItemsTable struct {
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
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTablesResponseBodyItemsTable) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablesResponseBodyItemsTable) GoString() string {
	return s.String()
}

func (s *DescribeTablesResponseBodyItemsTable) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeTablesResponseBodyItemsTable) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeTablesResponseBodyItemsTable) GetTableName() *string {
	return s.TableName
}

func (s *DescribeTablesResponseBodyItemsTable) SetDBClusterId(v string) *DescribeTablesResponseBodyItemsTable {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsTable) SetSchemaName(v string) *DescribeTablesResponseBodyItemsTable {
	s.SchemaName = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsTable) SetTableName(v string) *DescribeTablesResponseBodyItemsTable {
	s.TableName = &v
	return s
}

func (s *DescribeTablesResponseBodyItemsTable) Validate() error {
	return dara.Validate(s)
}
