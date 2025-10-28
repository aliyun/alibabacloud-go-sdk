// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListStackGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListStackGroupsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListStackGroupsRequest
	GetPageSize() *int64
	SetRegionId(v string) *ListStackGroupsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ListStackGroupsRequest
	GetResourceGroupId() *string
	SetStatus(v string) *ListStackGroupsRequest
	GetStatus() *string
	SetTags(v []*ListStackGroupsRequestTags) *ListStackGroupsRequest
	GetTags() []*ListStackGroupsRequestTags
}

type ListStackGroupsRequest struct {
	// The number of the page to return.
	//
	// 	- Pages start from page 1.
	//
	// 	- Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// 	- Valid values: 1 to 50.
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the stack group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. If you do not specify this parameter, the stack groups in all the resource groups are queried.
	//
	// > To obtain the resource group ID, go to the **Resource Group*	- page in the **Resource Management*	- console. For more information, see [View the basic information about a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-acfmzawhxxcj****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The state of the stack group. If you do not specify this parameter, the stack groups in all states in the specified region are queried.
	//
	// Valid values:
	//
	// 	- ACTIVE
	//
	// 	- DELETED
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags that are added to the stack group.
	Tags []*ListStackGroupsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListStackGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListStackGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListStackGroupsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListStackGroupsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListStackGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListStackGroupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListStackGroupsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListStackGroupsRequest) GetTags() []*ListStackGroupsRequestTags {
	return s.Tags
}

func (s *ListStackGroupsRequest) SetPageNumber(v int64) *ListStackGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStackGroupsRequest) SetPageSize(v int64) *ListStackGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStackGroupsRequest) SetRegionId(v string) *ListStackGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListStackGroupsRequest) SetResourceGroupId(v string) *ListStackGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListStackGroupsRequest) SetStatus(v string) *ListStackGroupsRequest {
	s.Status = &v
	return s
}

func (s *ListStackGroupsRequest) SetTags(v []*ListStackGroupsRequestTags) *ListStackGroupsRequest {
	s.Tags = v
	return s
}

func (s *ListStackGroupsRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListStackGroupsRequestTags struct {
	// The key of the tag that is added to the stack group.
	//
	// > Tags is optional. If you specify Tags, you must specify Tags.N.Key.
	//
	// This parameter is required.
	//
	// example:
	//
	// usage
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that is added to the stack group.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListStackGroupsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListStackGroupsRequestTags) GoString() string {
	return s.String()
}

func (s *ListStackGroupsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListStackGroupsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListStackGroupsRequestTags) SetKey(v string) *ListStackGroupsRequestTags {
	s.Key = &v
	return s
}

func (s *ListStackGroupsRequestTags) SetValue(v string) *ListStackGroupsRequestTags {
	s.Value = &v
	return s
}

func (s *ListStackGroupsRequestTags) Validate() error {
	return dara.Validate(s)
}
