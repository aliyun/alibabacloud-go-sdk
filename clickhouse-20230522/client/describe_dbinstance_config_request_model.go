// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceConfigRequest
	GetDBInstanceId() *string
	SetRegionId(v string) *DescribeDBInstanceConfigRequest
	GetRegionId() *string
}

type DescribeDBInstanceConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc-wz9go4x*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDBInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceConfigRequest) SetDBInstanceId(v string) *DescribeDBInstanceConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceConfigRequest) SetRegionId(v string) *DescribeDBInstanceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
