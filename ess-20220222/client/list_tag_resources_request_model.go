// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagResourcesRequest
	GetNextToken() *string
	SetOwnerId(v int64) *ListTagResourcesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListTagResourcesRequest
	GetRegionId() *string
	SetResourceIds(v []*string) *ListTagResourcesRequest
	GetResourceIds() []*string
	SetResourceOwnerAccount(v string) *ListTagResourcesRequest
	GetResourceOwnerAccount() *string
	SetResourceType(v string) *ListTagResourcesRequest
	GetResourceType() *string
	SetTags(v []*ListTagResourcesRequestTags) *ListTagResourcesRequest
	GetTags() []*ListTagResourcesRequestTags
}

type ListTagResourcesRequest struct {
	// The token that determines the start point of the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2679950.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the Auto Scaling resources. You can specify 1 to 50 resource IDs.
	ResourceIds          []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The resource type. Set the value to scalinggroup.
	//
	// This parameter is required.
	//
	// example:
	//
	// scalinggroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Details of the tags.
	Tags []*ListTagResourcesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagResourcesRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *ListTagResourcesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesRequest) GetTags() []*ListTagResourcesRequestTags {
	return s.Tags
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceIds(v []*string) *ListTagResourcesRequest {
	s.ResourceIds = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v []*ListTagResourcesRequestTags) *ListTagResourcesRequest {
	s.Tags = v
	return s
}

func (s *ListTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}

type ListTagResourcesRequestTags struct {
	// The key of the tag. The key is used to perform an exact match of Auto Scaling resources. The key must be 1 to 128 characters in length.
	//
	// `Tags` is used to perform an exact match of Auto Scaling resources to which the specified tags are added. Specify a tag in the key-value pair format.
	//
	// 	- If you specify only `Tags.Key`, all resources whose tags contain the specified tag key are returned.
	//
	// 	- If you specify only `Tags.Value`, the `MissingParameter.TagKey` error is reported.
	//
	// 	- If you specify multiple key-value pairs at the same time, only Auto Scaling resources that match all the tag keys and tag values are returned.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag. The value is used to perform an exact match of Auto Scaling resources. The value can be up to 128 characters in length.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequestTags) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListTagResourcesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListTagResourcesRequestTags) SetKey(v string) *ListTagResourcesRequestTags {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTags) SetValue(v string) *ListTagResourcesRequestTags {
	s.Value = &v
	return s
}

func (s *ListTagResourcesRequestTags) Validate() error {
	return dara.Validate(s)
}
