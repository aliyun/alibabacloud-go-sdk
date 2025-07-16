// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceSSLRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeDBInstanceSSLRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeDBInstanceSSLRequest
	GetRegionId() *string
}

type DescribeDBInstanceSSLRequest struct {
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstanceSSLRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceSSLRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceSSLRequest) SetDBInstanceName(v string) *DescribeDBInstanceSSLRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceSSLRequest) SetRegionId(v string) *DescribeDBInstanceSSLRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceSSLRequest) Validate() error {
	return dara.Validate(s)
}
