// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnarVersionListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeColumnarVersionListRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeColumnarVersionListRequest
	GetRegionId() *string
}

type DescribeColumnarVersionListRequest struct {
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeColumnarVersionListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnarVersionListRequest) GoString() string {
	return s.String()
}

func (s *DescribeColumnarVersionListRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeColumnarVersionListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeColumnarVersionListRequest) SetDBInstanceName(v string) *DescribeColumnarVersionListRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeColumnarVersionListRequest) SetRegionId(v string) *DescribeColumnarVersionListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeColumnarVersionListRequest) Validate() error {
	return dara.Validate(s)
}
