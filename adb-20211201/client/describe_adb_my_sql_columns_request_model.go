// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdbMySqlColumnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAdbMySqlColumnsRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeAdbMySqlColumnsRequest
	GetRegionId() *string
	SetSchema(v string) *DescribeAdbMySqlColumnsRequest
	GetSchema() *string
	SetTableName(v string) *DescribeAdbMySqlColumnsRequest
	GetTableName() *string
}

type DescribeAdbMySqlColumnsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the IDs of all AnalyticDB for MySQL clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1r053byu48p****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/454314.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adb_demo
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAdbMySqlColumnsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlColumnsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlColumnsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAdbMySqlColumnsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAdbMySqlColumnsRequest) GetSchema() *string {
	return s.Schema
}

func (s *DescribeAdbMySqlColumnsRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAdbMySqlColumnsRequest) SetDBClusterId(v string) *DescribeAdbMySqlColumnsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAdbMySqlColumnsRequest) SetRegionId(v string) *DescribeAdbMySqlColumnsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAdbMySqlColumnsRequest) SetSchema(v string) *DescribeAdbMySqlColumnsRequest {
	s.Schema = &v
	return s
}

func (s *DescribeAdbMySqlColumnsRequest) SetTableName(v string) *DescribeAdbMySqlColumnsRequest {
	s.TableName = &v
	return s
}

func (s *DescribeAdbMySqlColumnsRequest) Validate() error {
	return dara.Validate(s)
}
