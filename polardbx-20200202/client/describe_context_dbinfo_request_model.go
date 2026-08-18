// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContextDBInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeContextDBInfoRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeContextDBInfoRequest
	GetRegionId() *string
}

type DescribeContextDBInfoRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeContextDBInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeContextDBInfoRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeContextDBInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeContextDBInfoRequest) SetDBInstanceName(v string) *DescribeContextDBInfoRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeContextDBInfoRequest) SetRegionId(v string) *DescribeContextDBInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContextDBInfoRequest) Validate() error {
	return dara.Validate(s)
}
