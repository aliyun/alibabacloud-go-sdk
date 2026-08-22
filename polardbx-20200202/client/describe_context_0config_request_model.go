// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContext0ConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeContext0ConfigRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeContext0ConfigRequest
	GetRegionId() *string
}

type DescribeContext0ConfigRequest struct {
	// 关联的 PolarDB-X 实例名（pxc- 前缀）
	//
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

func (s DescribeContext0ConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContext0ConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeContext0ConfigRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeContext0ConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeContext0ConfigRequest) SetDBInstanceName(v string) *DescribeContext0ConfigRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeContext0ConfigRequest) SetRegionId(v string) *DescribeContext0ConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContext0ConfigRequest) Validate() error {
	return dara.Validate(s)
}
