// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceNetInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceNetInfoRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeDBInstanceNetInfoRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DescribeDBInstanceNetInfoRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstanceNetInfoRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// selectdb-cn-7213cjv****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstanceNetInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceNetInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceNetInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceNetInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstanceNetInfoRequest) SetDBInstanceId(v string) *DescribeDBInstanceNetInfoRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetRegionId(v string) *DescribeDBInstanceNetInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) SetResourceOwnerId(v int64) *DescribeDBInstanceNetInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoRequest) Validate() error {
	return dara.Validate(s)
}
