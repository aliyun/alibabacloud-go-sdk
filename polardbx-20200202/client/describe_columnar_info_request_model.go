// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnarInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeColumnarInfoRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeColumnarInfoRequest
	GetRegionId() *string
}

type DescribeColumnarInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeColumnarInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnarInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeColumnarInfoRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeColumnarInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeColumnarInfoRequest) SetDBInstanceName(v string) *DescribeColumnarInfoRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeColumnarInfoRequest) SetRegionId(v string) *DescribeColumnarInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeColumnarInfoRequest) Validate() error {
	return dara.Validate(s)
}
