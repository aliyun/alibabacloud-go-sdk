// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeServiceAccountRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeServiceAccountRequest
	GetRegionId() *string
}

type DescribeServiceAccountRequest struct {
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

func (s DescribeServiceAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceAccountRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceAccountRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeServiceAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServiceAccountRequest) SetDBInstanceName(v string) *DescribeServiceAccountRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeServiceAccountRequest) SetRegionId(v string) *DescribeServiceAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeServiceAccountRequest) Validate() error {
	return dara.Validate(s)
}
