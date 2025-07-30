// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllDBInstanceClassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeAllDBInstanceClassRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DescribeAllDBInstanceClassRequest
	GetResourceOwnerId() *int64
}

type DescribeAllDBInstanceClassRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAllDBInstanceClassRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDBInstanceClassRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllDBInstanceClassRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAllDBInstanceClassRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAllDBInstanceClassRequest) SetRegionId(v string) *DescribeAllDBInstanceClassRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAllDBInstanceClassRequest) SetResourceOwnerId(v int64) *DescribeAllDBInstanceClassRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAllDBInstanceClassRequest) Validate() error {
	return dara.Validate(s)
}
