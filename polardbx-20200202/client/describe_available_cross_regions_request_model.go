// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableCrossRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeAvailableCrossRegionsRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeAvailableCrossRegionsRequest
	GetRegionId() *string
}

type DescribeAvailableCrossRegionsRequest struct {
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAvailableCrossRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableCrossRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCrossRegionsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeAvailableCrossRegionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAvailableCrossRegionsRequest) SetDBInstanceName(v string) *DescribeAvailableCrossRegionsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAvailableCrossRegionsRequest) SetRegionId(v string) *DescribeAvailableCrossRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableCrossRegionsRequest) Validate() error {
	return dara.Validate(s)
}
