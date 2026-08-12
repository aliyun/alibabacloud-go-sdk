// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceInstanceCertsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeResourceInstanceCertsRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *DescribeResourceInstanceCertsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeResourceInstanceCertsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeResourceInstanceCertsRequest
	GetRegionId() *string
	SetResourceInstanceId(v string) *DescribeResourceInstanceCertsRequest
	GetResourceInstanceId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeResourceInstanceCertsRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeResourceInstanceCertsRequest struct {
	// The WAF instance ID.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the current WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-*
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number to return in a paged query. Default value: **1**, which indicates the first page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in a paged query. Default value: **10**, which indicates 10 entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The cloud service instance ID.
	//
	// example:
	//
	// lb-bp1*****jqnnqk5uj2p
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The Alibaba Cloud resource group ID.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeResourceInstanceCertsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceInstanceCertsRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceInstanceCertsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeResourceInstanceCertsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeResourceInstanceCertsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeResourceInstanceCertsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeResourceInstanceCertsRequest) GetResourceInstanceId() *string {
	return s.ResourceInstanceId
}

func (s *DescribeResourceInstanceCertsRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeResourceInstanceCertsRequest) SetInstanceId(v string) *DescribeResourceInstanceCertsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeResourceInstanceCertsRequest) SetPageNumber(v int64) *DescribeResourceInstanceCertsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeResourceInstanceCertsRequest) SetPageSize(v int64) *DescribeResourceInstanceCertsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeResourceInstanceCertsRequest) SetRegionId(v string) *DescribeResourceInstanceCertsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeResourceInstanceCertsRequest) SetResourceInstanceId(v string) *DescribeResourceInstanceCertsRequest {
	s.ResourceInstanceId = &v
	return s
}

func (s *DescribeResourceInstanceCertsRequest) SetResourceManagerResourceGroupId(v string) *DescribeResourceInstanceCertsRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeResourceInstanceCertsRequest) Validate() error {
	return dara.Validate(s)
}
