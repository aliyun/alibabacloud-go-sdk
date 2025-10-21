// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeEndpointsRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeEndpointsRequest
	GetRegionId() *string
}

type DescribeEndpointsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndpointsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndpointsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeEndpointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEndpointsRequest) SetDBInstanceId(v string) *DescribeEndpointsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeEndpointsRequest) SetRegionId(v string) *DescribeEndpointsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
