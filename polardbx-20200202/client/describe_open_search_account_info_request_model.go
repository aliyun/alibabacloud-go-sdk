// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchAccountInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeOpenSearchAccountInfoRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeOpenSearchAccountInfoRequest
	GetRegionId() *string
}

type DescribeOpenSearchAccountInfoRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxx
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

func (s DescribeOpenSearchAccountInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchAccountInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchAccountInfoRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeOpenSearchAccountInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeOpenSearchAccountInfoRequest) SetDBInstanceName(v string) *DescribeOpenSearchAccountInfoRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoRequest) SetRegionId(v string) *DescribeOpenSearchAccountInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoRequest) Validate() error {
	return dara.Validate(s)
}
