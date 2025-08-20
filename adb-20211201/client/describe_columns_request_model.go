// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeColumnsRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeColumnsRequest
	GetRegionId() *string
	SetSchemaName(v string) *DescribeColumnsRequest
	GetSchemaName() *string
	SetTableName(v string) *DescribeColumnsRequest
	GetTableName() *string
}

type DescribeColumnsRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1xxxxxxxx47
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

func (s DescribeColumnsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnsRequest) GoString() string {
	return s.String()
}

func (s *DescribeColumnsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeColumnsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeColumnsRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeColumnsRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeColumnsRequest) SetDBClusterId(v string) *DescribeColumnsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeColumnsRequest) SetRegionId(v string) *DescribeColumnsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeColumnsRequest) SetSchemaName(v string) *DescribeColumnsRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeColumnsRequest) SetTableName(v string) *DescribeColumnsRequest {
	s.TableName = &v
	return s
}

func (s *DescribeColumnsRequest) Validate() error {
	return dara.Validate(s)
}
