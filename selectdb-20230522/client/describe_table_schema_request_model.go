// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableSchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeTableSchemaRequest
	GetDBInstanceId() *string
	SetDatabase(v string) *DescribeTableSchemaRequest
	GetDatabase() *string
	SetRegionId(v string) *DescribeTableSchemaRequest
	GetRegionId() *string
	SetTable(v string) *DescribeTableSchemaRequest
	GetTable() *string
}

type DescribeTableSchemaRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_db
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_tb
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeTableSchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableSchemaRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableSchemaRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeTableSchemaRequest) GetDatabase() *string {
	return s.Database
}

func (s *DescribeTableSchemaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTableSchemaRequest) GetTable() *string {
	return s.Table
}

func (s *DescribeTableSchemaRequest) SetDBInstanceId(v string) *DescribeTableSchemaRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeTableSchemaRequest) SetDatabase(v string) *DescribeTableSchemaRequest {
	s.Database = &v
	return s
}

func (s *DescribeTableSchemaRequest) SetRegionId(v string) *DescribeTableSchemaRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableSchemaRequest) SetTable(v string) *DescribeTableSchemaRequest {
	s.Table = &v
	return s
}

func (s *DescribeTableSchemaRequest) Validate() error {
	return dara.Validate(s)
}
