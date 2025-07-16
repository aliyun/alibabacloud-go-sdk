// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDistributeTableListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeDistributeTableListRequest
	GetDBInstanceName() *string
	SetDbName(v string) *DescribeDistributeTableListRequest
	GetDbName() *string
	SetRegionId(v string) *DescribeDistributeTableListRequest
	GetRegionId() *string
}

type DescribeDistributeTableListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sbtest1
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDistributeTableListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistributeTableListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDistributeTableListRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDistributeTableListRequest) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDistributeTableListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDistributeTableListRequest) SetDBInstanceName(v string) *DescribeDistributeTableListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDistributeTableListRequest) SetDbName(v string) *DescribeDistributeTableListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeDistributeTableListRequest) SetRegionId(v string) *DescribeDistributeTableListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDistributeTableListRequest) Validate() error {
	return dara.Validate(s)
}
