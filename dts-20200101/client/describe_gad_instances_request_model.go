// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGadInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbEngineTypes(v string) *DescribeGadInstancesRequest
	GetDbEngineTypes() *string
	SetInstanceName(v string) *DescribeGadInstancesRequest
	GetInstanceName() *string
	SetMasterDbInstanceId(v string) *DescribeGadInstancesRequest
	GetMasterDbInstanceId() *string
	SetOwnerId(v string) *DescribeGadInstancesRequest
	GetOwnerId() *string
	SetPageNumber(v int32) *DescribeGadInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGadInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeGadInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeGadInstancesRequest
	GetResourceGroupId() *string
	SetSlaveDbInstanceId(v string) *DescribeGadInstancesRequest
	GetSlaveDbInstanceId() *string
}

type DescribeGadInstancesRequest struct {
	DbEngineTypes *string `json:"DbEngineTypes,omitempty" xml:"DbEngineTypes,omitempty"`
	// example:
	//
	// test
	InstanceName       *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MasterDbInstanceId *string `json:"MasterDbInstanceId,omitempty" xml:"MasterDbInstanceId,omitempty"`
	OwnerId            *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId   *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SlaveDbInstanceId *string `json:"SlaveDbInstanceId,omitempty" xml:"SlaveDbInstanceId,omitempty"`
}

func (s DescribeGadInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGadInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeGadInstancesRequest) GetDbEngineTypes() *string {
	return s.DbEngineTypes
}

func (s *DescribeGadInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeGadInstancesRequest) GetMasterDbInstanceId() *string {
	return s.MasterDbInstanceId
}

func (s *DescribeGadInstancesRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeGadInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGadInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGadInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGadInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeGadInstancesRequest) GetSlaveDbInstanceId() *string {
	return s.SlaveDbInstanceId
}

func (s *DescribeGadInstancesRequest) SetDbEngineTypes(v string) *DescribeGadInstancesRequest {
	s.DbEngineTypes = &v
	return s
}

func (s *DescribeGadInstancesRequest) SetInstanceName(v string) *DescribeGadInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeGadInstancesRequest) SetMasterDbInstanceId(v string) *DescribeGadInstancesRequest {
	s.MasterDbInstanceId = &v
	return s
}

func (s *DescribeGadInstancesRequest) SetOwnerId(v string) *DescribeGadInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeGadInstancesRequest) SetPageNumber(v int32) *DescribeGadInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGadInstancesRequest) SetPageSize(v int32) *DescribeGadInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGadInstancesRequest) SetRegionId(v string) *DescribeGadInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeGadInstancesRequest) SetResourceGroupId(v string) *DescribeGadInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeGadInstancesRequest) SetSlaveDbInstanceId(v string) *DescribeGadInstancesRequest {
	s.SlaveDbInstanceId = &v
	return s
}

func (s *DescribeGadInstancesRequest) Validate() error {
	return dara.Validate(s)
}
