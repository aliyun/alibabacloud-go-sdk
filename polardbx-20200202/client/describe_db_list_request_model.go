// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDbListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeDbListRequest
	GetDBInstanceName() *string
	SetDBName(v string) *DescribeDbListRequest
	GetDBName() *string
	SetRegionId(v string) *DescribeDbListRequest
	GetRegionId() *string
}

type DescribeDbListRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// db_name
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDbListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDbListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDbListRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDbListRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeDbListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDbListRequest) SetDBInstanceName(v string) *DescribeDbListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDbListRequest) SetDBName(v string) *DescribeDbListRequest {
	s.DBName = &v
	return s
}

func (s *DescribeDbListRequest) SetRegionId(v string) *DescribeDbListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDbListRequest) Validate() error {
	return dara.Validate(s)
}
