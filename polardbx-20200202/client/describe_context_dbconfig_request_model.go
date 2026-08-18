// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContextDBConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeContextDBConfigRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeContextDBConfigRequest
	GetRegionId() *string
}

type DescribeContextDBConfigRequest struct {
	// The instance ID.
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

func (s DescribeContextDBConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeContextDBConfigRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeContextDBConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeContextDBConfigRequest) SetDBInstanceName(v string) *DescribeContextDBConfigRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeContextDBConfigRequest) SetRegionId(v string) *DescribeContextDBConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContextDBConfigRequest) Validate() error {
	return dara.Validate(s)
}
