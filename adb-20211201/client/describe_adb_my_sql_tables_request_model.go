// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdbMySqlTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAdbMySqlTablesRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeAdbMySqlTablesRequest
	GetRegionId() *string
	SetSchema(v string) *DescribeAdbMySqlTablesRequest
	GetSchema() *string
}

type DescribeAdbMySqlTablesRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
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
	// if can be null:
	// false
	//
	// example:
	//
	// adb_demo
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
}

func (s DescribeAdbMySqlTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlTablesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAdbMySqlTablesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAdbMySqlTablesRequest) GetSchema() *string {
	return s.Schema
}

func (s *DescribeAdbMySqlTablesRequest) SetDBClusterId(v string) *DescribeAdbMySqlTablesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAdbMySqlTablesRequest) SetRegionId(v string) *DescribeAdbMySqlTablesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAdbMySqlTablesRequest) SetSchema(v string) *DescribeAdbMySqlTablesRequest {
	s.Schema = &v
	return s
}

func (s *DescribeAdbMySqlTablesRequest) Validate() error {
	return dara.Validate(s)
}
