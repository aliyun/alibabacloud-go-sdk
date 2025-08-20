// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdbMySqlIndexesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAdbMySqlIndexesRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeAdbMySqlIndexesRequest
	GetRegionId() *string
	SetSchema(v string) *DescribeAdbMySqlIndexesRequest
	GetSchema() *string
	SetTableName(v string) *DescribeAdbMySqlIndexesRequest
	GetTableName() *string
}

type DescribeAdbMySqlIndexesRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Warehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-uf6wjk5xxxxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/612393.html) operation to query the most recent region list.
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
	// tpch_oss
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The name of the table.
	//
	// >  If you leave this parameter empty, the information about all the current tables in the cluster is returned.
	//
	// example:
	//
	// orders
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAdbMySqlIndexesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlIndexesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlIndexesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAdbMySqlIndexesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAdbMySqlIndexesRequest) GetSchema() *string {
	return s.Schema
}

func (s *DescribeAdbMySqlIndexesRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAdbMySqlIndexesRequest) SetDBClusterId(v string) *DescribeAdbMySqlIndexesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAdbMySqlIndexesRequest) SetRegionId(v string) *DescribeAdbMySqlIndexesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAdbMySqlIndexesRequest) SetSchema(v string) *DescribeAdbMySqlIndexesRequest {
	s.Schema = &v
	return s
}

func (s *DescribeAdbMySqlIndexesRequest) SetTableName(v string) *DescribeAdbMySqlIndexesRequest {
	s.TableName = &v
	return s
}

func (s *DescribeAdbMySqlIndexesRequest) Validate() error {
	return dara.Validate(s)
}
