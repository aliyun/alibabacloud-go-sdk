// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityIPListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSecurityIPListRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeSecurityIPListRequest
	GetRegionId() *string
}

type DescribeSecurityIPListRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-bp100p4q1g9z3****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSecurityIPListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIPListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIPListRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSecurityIPListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityIPListRequest) SetDBInstanceId(v string) *DescribeSecurityIPListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSecurityIPListRequest) SetRegionId(v string) *DescribeSecurityIPListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityIPListRequest) Validate() error {
	return dara.Validate(s)
}
