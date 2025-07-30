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
	SetResourceOwnerId(v int64) *DescribeSecurityIPListRequest
	GetResourceOwnerId() *int64
}

type DescribeSecurityIPListRequest struct {
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

func (s *DescribeSecurityIPListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSecurityIPListRequest) SetDBInstanceId(v string) *DescribeSecurityIPListRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSecurityIPListRequest) SetRegionId(v string) *DescribeSecurityIPListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityIPListRequest) SetResourceOwnerId(v int64) *DescribeSecurityIPListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecurityIPListRequest) Validate() error {
	return dara.Validate(s)
}
