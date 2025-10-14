// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeShowStorageInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeShowStorageInfoRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeShowStorageInfoRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeShowStorageInfoRequest
	GetResourceGroupId() *string
}

type DescribeShowStorageInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-**************
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeShowStorageInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeShowStorageInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeShowStorageInfoRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeShowStorageInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeShowStorageInfoRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeShowStorageInfoRequest) SetDBInstanceName(v string) *DescribeShowStorageInfoRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeShowStorageInfoRequest) SetRegionId(v string) *DescribeShowStorageInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeShowStorageInfoRequest) SetResourceGroupId(v string) *DescribeShowStorageInfoRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeShowStorageInfoRequest) Validate() error {
	return dara.Validate(s)
}
