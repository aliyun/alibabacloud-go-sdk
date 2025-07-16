// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeSecurityIpsRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeSecurityIpsRequest
	GetRegionId() *string
}

type DescribeSecurityIpsRequest struct {
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

func (s DescribeSecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSecurityIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityIpsRequest) SetDBInstanceName(v string) *DescribeSecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSecurityIpsRequest) SetRegionId(v string) *DescribeSecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
