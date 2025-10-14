// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTransformStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeTransformStatusRequest
	GetDBInstanceName() *string
	SetQueryReport(v bool) *DescribeTransformStatusRequest
	GetQueryReport() *bool
	SetRegionId(v string) *DescribeTransformStatusRequest
	GetRegionId() *string
}

type DescribeTransformStatusRequest struct {
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// true
	QueryReport *bool `json:"QueryReport,omitempty" xml:"QueryReport,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeTransformStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTransformStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeTransformStatusRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeTransformStatusRequest) GetQueryReport() *bool {
	return s.QueryReport
}

func (s *DescribeTransformStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTransformStatusRequest) SetDBInstanceName(v string) *DescribeTransformStatusRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeTransformStatusRequest) SetQueryReport(v bool) *DescribeTransformStatusRequest {
	s.QueryReport = &v
	return s
}

func (s *DescribeTransformStatusRequest) SetRegionId(v string) *DescribeTransformStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTransformStatusRequest) Validate() error {
	return dara.Validate(s)
}
