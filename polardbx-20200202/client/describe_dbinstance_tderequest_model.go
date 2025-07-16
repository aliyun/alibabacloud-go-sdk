// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceTDERequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeDBInstanceTDERequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeDBInstanceTDERequest
	GetRegionId() *string
}

type DescribeDBInstanceTDERequest struct {
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstanceTDERequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceTDERequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceTDERequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceTDERequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceTDERequest) SetDBInstanceName(v string) *DescribeDBInstanceTDERequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceTDERequest) SetRegionId(v string) *DescribeDBInstanceTDERequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceTDERequest) Validate() error {
	return dara.Validate(s)
}
