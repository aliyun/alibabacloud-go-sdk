// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceAttributeRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeDBInstanceAttributeRequest
	GetRegionId() *string
}

type DescribeDBInstanceAttributeRequest struct {
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
	// Valid values:
	//
	// 	- cn-beijing
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceAttributeRequest) SetDBInstanceId(v string) *DescribeDBInstanceAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetRegionId(v string) *DescribeDBInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
