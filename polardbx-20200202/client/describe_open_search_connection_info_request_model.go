// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchConnectionInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeOpenSearchConnectionInfoRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeOpenSearchConnectionInfoRequest
	GetRegionId() *string
}

type DescribeOpenSearchConnectionInfoRequest struct {
	// The instance name.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeOpenSearchConnectionInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchConnectionInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchConnectionInfoRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeOpenSearchConnectionInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOpenSearchConnectionInfoRequest) SetDBInstanceName(v string) *DescribeOpenSearchConnectionInfoRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoRequest) SetRegionId(v string) *DescribeOpenSearchConnectionInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOpenSearchConnectionInfoRequest) Validate() error {
	return dara.Validate(s)
}
