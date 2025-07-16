// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeDBInstanceAttributeRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeDBInstanceAttributeRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDBInstanceAttributeRequest
	GetResourceGroupId() *string
}

type DescribeDBInstanceAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-aekzaxhezhs5***
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDBInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeDBInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDBInstanceAttributeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstanceAttributeRequest) SetDBInstanceName(v string) *DescribeDBInstanceAttributeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetRegionId(v string) *DescribeDBInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetResourceGroupId(v string) *DescribeDBInstanceAttributeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}
