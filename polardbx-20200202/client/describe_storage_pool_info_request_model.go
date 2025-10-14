// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStoragePoolInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeStoragePoolInfoRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *DescribeStoragePoolInfoRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeStoragePoolInfoRequest
	GetResourceGroupId() *string
}

type DescribeStoragePoolInfoRequest struct {
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
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeStoragePoolInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStoragePoolInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeStoragePoolInfoRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeStoragePoolInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStoragePoolInfoRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeStoragePoolInfoRequest) SetDBInstanceName(v string) *DescribeStoragePoolInfoRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeStoragePoolInfoRequest) SetRegionId(v string) *DescribeStoragePoolInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStoragePoolInfoRequest) SetResourceGroupId(v string) *DescribeStoragePoolInfoRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeStoragePoolInfoRequest) Validate() error {
	return dara.Validate(s)
}
