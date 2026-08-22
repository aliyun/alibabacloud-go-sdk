// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContext0SecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeContext0SecurityIpsRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeContext0SecurityIpsRequest
	GetRegionId() *string
}

type DescribeContext0SecurityIpsRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The region in which the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeContext0SecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0SecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContext0SecurityIpsRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeContext0SecurityIpsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeContext0SecurityIpsRequest) SetDBInstanceName(v string) *DescribeContext0SecurityIpsRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeContext0SecurityIpsRequest) SetRegionId(v string) *DescribeContext0SecurityIpsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContext0SecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
