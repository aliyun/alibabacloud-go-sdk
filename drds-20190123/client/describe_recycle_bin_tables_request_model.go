// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecycleBinTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DescribeRecycleBinTablesRequest
	GetDbName() *string
	SetDrdsInstanceId(v string) *DescribeRecycleBinTablesRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *DescribeRecycleBinTablesRequest
	GetRegionId() *string
}

type DescribeRecycleBinTablesRequest struct {
	// The name of the database.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds***********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRecycleBinTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecycleBinTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinTablesRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeRecycleBinTablesRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeRecycleBinTablesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRecycleBinTablesRequest) SetDbName(v string) *DescribeRecycleBinTablesRequest {
	s.DbName = &v
	return s
}

func (s *DescribeRecycleBinTablesRequest) SetDrdsInstanceId(v string) *DescribeRecycleBinTablesRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRecycleBinTablesRequest) SetRegionId(v string) *DescribeRecycleBinTablesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRecycleBinTablesRequest) Validate() error {
	return dara.Validate(s)
}
