// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceDataSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceDataSourcesRequest
	GetDBInstanceId() *string
	SetDBName(v string) *DescribeDBInstanceDataSourcesRequest
	GetDBName() *string
	SetRegionId(v string) *DescribeDBInstanceDataSourcesRequest
	GetRegionId() *string
	SetTableName(v string) *DescribeDBInstanceDataSourcesRequest
	GetTableName() *string
}

type DescribeDBInstanceDataSourcesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The database name.
	//
	// example:
	//
	// dbtest
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The table name.
	//
	// example:
	//
	// tableTest
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeDBInstanceDataSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSourcesRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceDataSourcesRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeDBInstanceDataSourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceDataSourcesRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeDBInstanceDataSourcesRequest) SetDBInstanceId(v string) *DescribeDBInstanceDataSourcesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesRequest) SetDBName(v string) *DescribeDBInstanceDataSourcesRequest {
	s.DBName = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesRequest) SetRegionId(v string) *DescribeDBInstanceDataSourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesRequest) SetTableName(v string) *DescribeDBInstanceDataSourcesRequest {
	s.TableName = &v
	return s
}

func (s *DescribeDBInstanceDataSourcesRequest) Validate() error {
	return dara.Validate(s)
}
