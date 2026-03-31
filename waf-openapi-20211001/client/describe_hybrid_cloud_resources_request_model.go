// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHybridCloudResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackend(v string) *DescribeHybridCloudResourcesRequest
	GetBackend() *string
	SetCnameEnabled(v bool) *DescribeHybridCloudResourcesRequest
	GetCnameEnabled() *bool
	SetDomain(v string) *DescribeHybridCloudResourcesRequest
	GetDomain() *string
	SetInstanceId(v string) *DescribeHybridCloudResourcesRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *DescribeHybridCloudResourcesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeHybridCloudResourcesRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeHybridCloudResourcesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudResourcesRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeHybridCloudResourcesRequest struct {
	// The back-to-origin IP address or domain name.
	//
	// example:
	//
	// 1.1.XX.XX
	Backend *string `json:"Backend,omitempty" xml:"Backend,omitempty"`
	// Specifies whether the public cloud disaster recovery feature is enabled for the domain name. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CnameEnabled *bool `json:"CnameEnabled,omitempty" xml:"CnameEnabled,omitempty"`
	// The domain name that you want to query.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v3prepaid_public_cn-***********
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
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou:*	- the Chinese mainland.
	//
	// 	- **ap-southeast-1:*	- outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmvtc5z52****
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeHybridCloudResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHybridCloudResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeHybridCloudResourcesRequest) GetBackend() *string {
	return s.Backend
}

func (s *DescribeHybridCloudResourcesRequest) GetCnameEnabled() *bool {
	return s.CnameEnabled
}

func (s *DescribeHybridCloudResourcesRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeHybridCloudResourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeHybridCloudResourcesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeHybridCloudResourcesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeHybridCloudResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHybridCloudResourcesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeHybridCloudResourcesRequest) SetBackend(v string) *DescribeHybridCloudResourcesRequest {
	s.Backend = &v
	return s
}

func (s *DescribeHybridCloudResourcesRequest) SetCnameEnabled(v bool) *DescribeHybridCloudResourcesRequest {
	s.CnameEnabled = &v
	return s
}

func (s *DescribeHybridCloudResourcesRequest) SetDomain(v string) *DescribeHybridCloudResourcesRequest {
	s.Domain = &v
	return s
}

func (s *DescribeHybridCloudResourcesRequest) SetInstanceId(v string) *DescribeHybridCloudResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeHybridCloudResourcesRequest) SetPageNumber(v int64) *DescribeHybridCloudResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHybridCloudResourcesRequest) SetPageSize(v int64) *DescribeHybridCloudResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHybridCloudResourcesRequest) SetRegionId(v string) *DescribeHybridCloudResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHybridCloudResourcesRequest) SetResourceManagerResourceGroupId(v string) *DescribeHybridCloudResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeHybridCloudResourcesRequest) Validate() error {
	return dara.Validate(s)
}
