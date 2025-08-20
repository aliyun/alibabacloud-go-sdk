// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeTablesRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeTablesRequest
	GetRegionId() *string
	SetSchemaName(v string) *DescribeTablesRequest
	GetSchemaName() *string
}

type DescribeTablesRequest struct {
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
}

func (s DescribeTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTablesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeTablesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTablesRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeTablesRequest) SetDBClusterId(v string) *DescribeTablesRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTablesRequest) SetRegionId(v string) *DescribeTablesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTablesRequest) SetSchemaName(v string) *DescribeTablesRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeTablesRequest) Validate() error {
	return dara.Validate(s)
}
