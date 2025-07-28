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
	// The ID of the WAF instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_v2_public_cn-*
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
	// The region in which the WAF instance is deployed. Valid values:
	//
	// 	- **cn-hangzhou**: Chinese mainland.
	//
	// 	- **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// lb-bp1*****jqnnqk5uj2p
	ResourceInstanceId *string `json:"ResourceInstanceId,omitempty" xml:"ResourceInstanceId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
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
