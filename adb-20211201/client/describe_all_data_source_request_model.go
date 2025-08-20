// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAllDataSourceRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeAllDataSourceRequest
	GetRegionId() *string
	SetSchemaName(v string) *DescribeAllDataSourceRequest
	GetSchemaName() *string
	SetTableName(v string) *DescribeAllDataSourceRequest
	GetTableName() *string
}

type DescribeAllDataSourceRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1pke2pcfavw****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
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
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAllDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAllDataSourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAllDataSourceRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeAllDataSourceRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAllDataSourceRequest) SetDBClusterId(v string) *DescribeAllDataSourceRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetRegionId(v string) *DescribeAllDataSourceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetSchemaName(v string) *DescribeAllDataSourceRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeAllDataSourceRequest) SetTableName(v string) *DescribeAllDataSourceRequest {
	s.TableName = &v
	return s
}

func (s *DescribeAllDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
