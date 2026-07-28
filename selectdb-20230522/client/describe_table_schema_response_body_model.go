// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateStatement(v string) *DescribeTableSchemaResponseBody
	GetCreateStatement() *string
	SetDatabase(v string) *DescribeTableSchemaResponseBody
	GetDatabase() *string
	SetRequestId(v string) *DescribeTableSchemaResponseBody
	GetRequestId() *string
	SetTable(v string) *DescribeTableSchemaResponseBody
	GetTable() *string
}

type DescribeTableSchemaResponseBody struct {
	// example:
	//
	// CREATE TABLE test_tb
	//
	// (
	//
	//     k1 TINYINT,
	//
	//     k2 DECIMAL(10, 2) DEFAULT "10.05",
	//
	//     k3 CHAR(10) COMMENT "string column",
	//
	//     k4 INT NOT NULL DEFAULT "1" COMMENT "int column"
	//
	// )
	//
	// COMMENT "my first table"
	//
	// DISTRIBUTED BY HASH(k1) BUCKETS 16
	CreateStatement *string `json:"CreateStatement,omitempty" xml:"CreateStatement,omitempty"`
	// example:
	//
	// test_db
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// example:
	//
	// 4773E4EC-025D-509F-AEA9-D53123FDFB0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// test_tb
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeTableSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableSchemaResponseBody) GetCreateStatement() *string {
	return s.CreateStatement
}

func (s *DescribeTableSchemaResponseBody) GetDatabase() *string {
	return s.Database
}

func (s *DescribeTableSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTableSchemaResponseBody) GetTable() *string {
	return s.Table
}

func (s *DescribeTableSchemaResponseBody) SetCreateStatement(v string) *DescribeTableSchemaResponseBody {
	s.CreateStatement = &v
	return s
}

func (s *DescribeTableSchemaResponseBody) SetDatabase(v string) *DescribeTableSchemaResponseBody {
	s.Database = &v
	return s
}

func (s *DescribeTableSchemaResponseBody) SetRequestId(v string) *DescribeTableSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableSchemaResponseBody) SetTable(v string) *DescribeTableSchemaResponseBody {
	s.Table = &v
	return s
}

func (s *DescribeTableSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}
