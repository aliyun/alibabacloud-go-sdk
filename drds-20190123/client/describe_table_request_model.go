// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeTableRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeTableRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *DescribeTableRequest
	GetRegionId() *string
	SetTableName(v string) *DescribeTableRequest
	GetTableName() *string
}

type DescribeTableRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds_flashback
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds*********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region where the PolarDB-X 1.0 instance is created.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the table.
	//
	// This parameter is required.
	//
	// example:
	//
	// employee_split
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableRequest) GoString() string {
	return s.String()
}

func (s *DescribeTableRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeTableRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *DescribeTableRequest) SetDbName(v string) *DescribeTableRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTableRequest) SetDrdsInstanceId(v string) *DescribeTableRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeTableRequest) SetRegionId(v string) *DescribeTableRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTableRequest) SetTableName(v string) *DescribeTableRequest {
	s.TableName = &v
	return s
}

func (s *DescribeTableRequest) Validate() error {
	return dara.Validate(s)
}
