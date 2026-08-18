// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContextDBSecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeContextDBSecurityIpsRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeContextDBSecurityIpsRequest
	GetRegionId() *string
}

type DescribeContextDBSecurityIpsRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
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

func (s DescribeContextDBSecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContextDBSecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContextDBSecurityIpsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeContextDBSecurityIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeContextDBSecurityIpsRequest) SetDBInstanceName(v string) *DescribeContextDBSecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeContextDBSecurityIpsRequest) SetRegionId(v string) *DescribeContextDBSecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContextDBSecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
