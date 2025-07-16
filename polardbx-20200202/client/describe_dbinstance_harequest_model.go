// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceHARequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeDBInstanceHARequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeDBInstanceHARequest
	GetRegionId() *string
}

type DescribeDBInstanceHARequest struct {
	// This parameter is required.
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstanceHARequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceHARequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceHARequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceHARequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceHARequest) SetDBInstanceName(v string) *DescribeDBInstanceHARequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceHARequest) SetRegionId(v string) *DescribeDBInstanceHARequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceHARequest) Validate() error {
	return dara.Validate(s)
}
