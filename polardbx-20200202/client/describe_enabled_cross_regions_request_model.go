// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnabledCrossRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeEnabledCrossRegionsRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeEnabledCrossRegionsRequest
	GetRegionId() *string
}

type DescribeEnabledCrossRegionsRequest struct {
	// example:
	//
	// pxc-bjxxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEnabledCrossRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnabledCrossRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnabledCrossRegionsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeEnabledCrossRegionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEnabledCrossRegionsRequest) SetDBInstanceName(v string) *DescribeEnabledCrossRegionsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeEnabledCrossRegionsRequest) SetRegionId(v string) *DescribeEnabledCrossRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEnabledCrossRegionsRequest) Validate() error {
	return dara.Validate(s)
}
