// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecProtectionGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApisecStatus(v int32) *DescribeApisecProtectionGroupsRequest
	GetApisecStatus() *int32
	SetInstanceId(v string) *DescribeApisecProtectionGroupsRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *DescribeApisecProtectionGroupsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeApisecProtectionGroupsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeApisecProtectionGroupsRequest
	GetRegionId() *string
	SetResourceGroup(v string) *DescribeApisecProtectionGroupsRequest
	GetResourceGroup() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecProtectionGroupsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeApisecProtectionGroupsRequest struct {
	// The switch of the API security module.
	//
	// example:
	//
	// 1
	ApisecStatus *int32 `json:"ApisecStatus,omitempty" xml:"ApisecStatus,omitempty"`
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the Web Application Firewall (WAF) instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the protected object group to which the protected object belongs.
	//
	// example:
	//
	// group1
	ResourceGroup *string `json:"ResourceGroup,omitempty" xml:"ResourceGroup,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeApisecProtectionGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecProtectionGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecProtectionGroupsRequest) GetApisecStatus() *int32 {
	return s.ApisecStatus
}

func (s *DescribeApisecProtectionGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecProtectionGroupsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeApisecProtectionGroupsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeApisecProtectionGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecProtectionGroupsRequest) GetResourceGroup() *string {
	return s.ResourceGroup
}

func (s *DescribeApisecProtectionGroupsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecProtectionGroupsRequest) SetApisecStatus(v int32) *DescribeApisecProtectionGroupsRequest {
	s.ApisecStatus = &v
	return s
}

func (s *DescribeApisecProtectionGroupsRequest) SetInstanceId(v string) *DescribeApisecProtectionGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecProtectionGroupsRequest) SetPageNumber(v int64) *DescribeApisecProtectionGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisecProtectionGroupsRequest) SetPageSize(v int64) *DescribeApisecProtectionGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisecProtectionGroupsRequest) SetRegionId(v string) *DescribeApisecProtectionGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecProtectionGroupsRequest) SetResourceGroup(v string) *DescribeApisecProtectionGroupsRequest {
	s.ResourceGroup = &v
	return s
}

func (s *DescribeApisecProtectionGroupsRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecProtectionGroupsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecProtectionGroupsRequest) Validate() error {
	return dara.Validate(s)
}
