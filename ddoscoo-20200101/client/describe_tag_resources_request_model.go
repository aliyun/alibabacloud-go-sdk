// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeTagResourcesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeTagResourcesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeTagResourcesRequest
	GetResourceGroupId() *string
	SetResourceIds(v []*string) *DescribeTagResourcesRequest
	GetResourceIds() []*string
	SetResourceType(v string) *DescribeTagResourcesRequest
	GetResourceType() *string
	SetTags(v []*DescribeTagResourcesRequestTags) *DescribeTagResourcesRequest
	GetTags() []*DescribeTagResourcesRequestTags
}

type DescribeTagResourcesRequest struct {
	// The query token. Set the value to the value of **NextToken*	- that is returned in the last call.
	//
	// > You do not need to configure this parameter if you call this operation for the first time.
	//
	// example:
	//
	// RGuYpqDdKhzXb8C3.D1BwQgc1tMBsoxdGiEKHHUUCf****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the instance. Set the value to **cn-hangzhou**, which indicates an Anti-DDoS Proxy (Chinese Mainland) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the instance belongs in Resource Management.
	//
	// If you do not configure this parameter, the instance belongs to the default resource group.
	//
	// example:
	//
	// rg-acfm2pz25js****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IDs of the Anti-DDoS Proxy (Chinese Mainland) instances that you want to query.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	// The type of the resource to which the tag belongs. Set the value to **INSTANCE**, which indicates an Anti-DDoS Pro instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// An array consisting of tags that you want to query. Each tag consists of a tag **key*	- and a tag **value**.
	Tags []*DescribeTagResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribeTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTagResourcesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeTagResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *DescribeTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTagResourcesRequest) GetTags() []*DescribeTagResourcesRequestTags {
	return s.Tags
}

func (s *DescribeTagResourcesRequest) SetNextToken(v string) *DescribeTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTagResourcesRequest) SetRegionId(v string) *DescribeTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagResourcesRequest) SetResourceGroupId(v string) *DescribeTagResourcesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeTagResourcesRequest) SetResourceIds(v []*string) *DescribeTagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *DescribeTagResourcesRequest) SetResourceType(v string) *DescribeTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagResourcesRequest) SetTags(v []*DescribeTagResourcesRequestTags) *DescribeTagResourcesRequest {
	s.Tags = v
	return s
}

func (s *DescribeTagResourcesRequest) Validate() error {
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

type DescribeTagResourcesRequestTags struct {
	// The key of the tag that you want to query.
	//
	// >
	//
	// 	- You must specify at least one of the **ResourceIds.N*	- and **Tags.N.Key*	- parameters.
	//
	// 	- You can call the [DescribeTagKeys](https://help.aliyun.com/document_detail/159785.html) operation to query all tag keys.
	//
	// example:
	//
	// testkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag that you want to query.
	//
	// example:
	//
	// testvalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeTagResourcesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *DescribeTagResourcesRequestTags) GetKey() *string {
	return s.Key
}

func (s *DescribeTagResourcesRequestTags) GetValue() *string {
	return s.Value
}

func (s *DescribeTagResourcesRequestTags) SetKey(v string) *DescribeTagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *DescribeTagResourcesRequestTags) SetValue(v string) *DescribeTagResourcesRequestTags {
	s.Value = &v
	return s
}

func (s *DescribeTagResourcesRequestTags) Validate() error {
	return dara.Validate(s)
}
