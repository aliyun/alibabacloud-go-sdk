// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdbMySqlTableMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAdbMySqlTableMetaRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeAdbMySqlTableMetaRequest
	GetRegionId() *string
	SetSchema(v string) *DescribeAdbMySqlTableMetaRequest
	GetSchema() *string
	SetTableName(v string) *DescribeAdbMySqlTableMetaRequest
	GetTableName() *string
}

type DescribeAdbMySqlTableMetaRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/129857.html) operation to query the IDs of all AnalyticDB for MySQL Data Lakehouse Edition clusters within a region.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp11q28kvl688****
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
	// This parameter is required.
	//
	// example:
	//
	// adb_demo
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The name of the table.
	//
	// >  If you leave this parameter empty, the information about all the current tables in the cluster is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAdbMySqlTableMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlTableMetaRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlTableMetaRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAdbMySqlTableMetaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAdbMySqlTableMetaRequest) GetSchema() *string {
	return s.Schema
}

func (s *DescribeAdbMySqlTableMetaRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAdbMySqlTableMetaRequest) SetDBClusterId(v string) *DescribeAdbMySqlTableMetaRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaRequest) SetRegionId(v string) *DescribeAdbMySqlTableMetaRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaRequest) SetSchema(v string) *DescribeAdbMySqlTableMetaRequest {
	s.Schema = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaRequest) SetTableName(v string) *DescribeAdbMySqlTableMetaRequest {
	s.TableName = &v
	return s
}

func (s *DescribeAdbMySqlTableMetaRequest) Validate() error {
	return dara.Validate(s)
}
