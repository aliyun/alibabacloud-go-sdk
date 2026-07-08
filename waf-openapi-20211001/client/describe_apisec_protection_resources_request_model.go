// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecProtectionResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApisecStatus(v int32) *DescribeApisecProtectionResourcesRequest
	GetApisecStatus() *int32
	SetInstanceId(v string) *DescribeApisecProtectionResourcesRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *DescribeApisecProtectionResourcesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeApisecProtectionResourcesRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeApisecProtectionResourcesRequest
	GetRegionId() *string
	SetResource(v string) *DescribeApisecProtectionResourcesRequest
	GetResource() *string
	SetResourceManagerResourceGroupId(v string) *DescribeApisecProtectionResourcesRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeApisecProtectionResourcesRequest struct {
	// The global switch for API security.
	//
	// example:
	//
	// 1
	ApisecStatus *int32 `json:"ApisecStatus,omitempty" xml:"ApisecStatus,omitempty"`
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-0xldbqt****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number to return when you perform a paged query. Default value: **1**, which indicates the first page. For more information about paging, see the API reference.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page when you perform a paged query. Default value: **10**, which indicates 10 entries per page. For more information about paging, see the API reference.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region in which the WAF instance is deployed. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the protected object.
	//
	// example:
	//
	// cwaf-***-waf
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeApisecProtectionResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecProtectionResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeApisecProtectionResourcesRequest) GetApisecStatus() *int32 {
	return s.ApisecStatus
}

func (s *DescribeApisecProtectionResourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeApisecProtectionResourcesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeApisecProtectionResourcesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeApisecProtectionResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeApisecProtectionResourcesRequest) GetResource() *string {
	return s.Resource
}

func (s *DescribeApisecProtectionResourcesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeApisecProtectionResourcesRequest) SetApisecStatus(v int32) *DescribeApisecProtectionResourcesRequest {
	s.ApisecStatus = &v
	return s
}

func (s *DescribeApisecProtectionResourcesRequest) SetInstanceId(v string) *DescribeApisecProtectionResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeApisecProtectionResourcesRequest) SetPageNumber(v int64) *DescribeApisecProtectionResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeApisecProtectionResourcesRequest) SetPageSize(v int64) *DescribeApisecProtectionResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeApisecProtectionResourcesRequest) SetRegionId(v string) *DescribeApisecProtectionResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeApisecProtectionResourcesRequest) SetResource(v string) *DescribeApisecProtectionResourcesRequest {
	s.Resource = &v
	return s
}

func (s *DescribeApisecProtectionResourcesRequest) SetResourceManagerResourceGroupId(v string) *DescribeApisecProtectionResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeApisecProtectionResourcesRequest) Validate() error {
	return dara.Validate(s)
}
