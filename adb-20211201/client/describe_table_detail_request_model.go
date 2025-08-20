// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeTableDetailRequest
	GetDBClusterId() *string
	SetRegionId(v string) *DescribeTableDetailRequest
	GetRegionId() *string
	SetSchemaName(v string) *DescribeTableDetailRequest
	GetSchemaName() *string
	SetTableName(v string) *DescribeTableDetailRequest
	GetTableName() *string
}

type DescribeTableDetailRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// am-bp1xxxxxxxx47
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The region ID.
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

func (s DescribeTableDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableDetailRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeTableDetailRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTableDetailRequest) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeTableDetailRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeTableDetailRequest) SetDBClusterId(v string) *DescribeTableDetailRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeTableDetailRequest) SetRegionId(v string) *DescribeTableDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableDetailRequest) SetSchemaName(v string) *DescribeTableDetailRequest {
	s.SchemaName = &v
	return s
}

func (s *DescribeTableDetailRequest) SetTableName(v string) *DescribeTableDetailRequest {
	s.TableName = &v
	return s
}

func (s *DescribeTableDetailRequest) Validate() error {
	return dara.Validate(s)
}
