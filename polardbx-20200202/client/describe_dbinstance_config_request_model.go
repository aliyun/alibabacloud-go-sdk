// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigName(v string) *DescribeDBInstanceConfigRequest
	GetConfigName() *string
	SetDBInstanceName(v string) *DescribeDBInstanceConfigRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeDBInstanceConfigRequest
	GetRegionId() *string
}

type DescribeDBInstanceConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// htap
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-*********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
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

func (s *DescribeDBInstanceConfigRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *DescribeDBInstanceConfigRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceConfigRequest) SetConfigName(v string) *DescribeDBInstanceConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *DescribeDBInstanceConfigRequest) SetDBInstanceName(v string) *DescribeDBInstanceConfigRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceConfigRequest) SetRegionId(v string) *DescribeDBInstanceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
