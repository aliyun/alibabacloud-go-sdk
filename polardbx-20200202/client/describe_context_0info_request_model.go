// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContext0InfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeContext0InfoRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeContext0InfoRequest
	GetRegionId() *string
}

type DescribeContext0InfoRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeContext0InfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0InfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeContext0InfoRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeContext0InfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeContext0InfoRequest) SetDBInstanceName(v string) *DescribeContext0InfoRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeContext0InfoRequest) SetRegionId(v string) *DescribeContext0InfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContext0InfoRequest) Validate() error {
	return dara.Validate(s)
}
