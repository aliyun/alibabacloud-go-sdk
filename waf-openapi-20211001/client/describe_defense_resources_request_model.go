// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeDefenseResourcesRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *DescribeDefenseResourcesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDefenseResourcesRequest
	GetPageSize() *int32
	SetQuery(v string) *DescribeDefenseResourcesRequest
	GetQuery() *string
	SetRegionId(v string) *DescribeDefenseResourcesRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourcesRequest
	GetResourceManagerResourceGroupId() *string
	SetTag(v []*DescribeDefenseResourcesRequestTag) *DescribeDefenseResourcesRequest
	GetTag() []*DescribeDefenseResourcesRequestTag
}

type DescribeDefenseResourcesRequest struct {
	// The ID of the Web Application Firewall (WAF) instance.
	//
	// >  You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to obtain the ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_cdnsdf3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number of the paginated results Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of results per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The query conditions. Specify the value of this parameter as a string in the JSON format.
	//
	// >  The results vary based on the query condition. For more information, see the "**Query parameters**" section in this topic.
	//
	// example:
	//
	// {\\"product\\":\\"waf\\"}
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The region ID of the WAF instance. Valid values:
	//
	// 	- **cn-hangzhou**: The Chinese mainland.
	//
	// 	- **ap-southeast-1**: Outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// The tag of the resource. You can specify up to 20 tags.
	Tag []*DescribeDefenseResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDefenseResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourcesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDefenseResourcesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDefenseResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDefenseResourcesRequest) GetQuery() *string {
	return s.Query
}

func (s *DescribeDefenseResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDefenseResourcesRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeDefenseResourcesRequest) GetTag() []*DescribeDefenseResourcesRequestTag {
	return s.Tag
}

func (s *DescribeDefenseResourcesRequest) SetInstanceId(v string) *DescribeDefenseResourcesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeDefenseResourcesRequest) SetPageNumber(v int32) *DescribeDefenseResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDefenseResourcesRequest) SetPageSize(v int32) *DescribeDefenseResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDefenseResourcesRequest) SetQuery(v string) *DescribeDefenseResourcesRequest {
	s.Query = &v
	return s
}

func (s *DescribeDefenseResourcesRequest) SetRegionId(v string) *DescribeDefenseResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDefenseResourcesRequest) SetResourceManagerResourceGroupId(v string) *DescribeDefenseResourcesRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeDefenseResourcesRequest) SetTag(v []*DescribeDefenseResourcesRequestTag) *DescribeDefenseResourcesRequest {
	s.Tag = v
	return s
}

func (s *DescribeDefenseResourcesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeDefenseResourcesRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// Tagkey1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TagValue1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDefenseResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDefenseResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDefenseResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDefenseResourcesRequestTag) SetKey(v string) *DescribeDefenseResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDefenseResourcesRequestTag) SetValue(v string) *DescribeDefenseResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDefenseResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
