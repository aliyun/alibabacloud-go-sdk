// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListTagResourcesRequest
	GetCategory() *string
	SetNextToken(v string) *ListTagResourcesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListTagResourcesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListTagResourcesRequest
	GetOwnerId() *int64
	SetPageSize(v int32) *ListTagResourcesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListTagResourcesRequest
	GetRegionId() *string
	SetResourceARN(v []*string) *ListTagResourcesRequest
	GetResourceARN() []*string
	SetResourceOwnerAccount(v string) *ListTagResourcesRequest
	GetResourceOwnerAccount() *string
	SetTags(v string) *ListTagResourcesRequest
	GetTags() *string
}

type ListTagResourcesRequest struct {
	// The type of the tag. Valid values:
	//
	// 	- Custom
	//
	// 	- System
	//
	// 	- All
	//
	// Default value: All.
	//
	// example:
	//
	// Custom
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The token that is used to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of entries to return on each page.
	//
	// Maximum value: 1000. Default value: 50.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID.
	//
	// 	- If the resources belong to a service that is centrally deployed, set the value to the region ID of the resources by referring to [Regions supported by tag-related operations on resources of centrally deployed Alibaba Cloud services](https://help.aliyun.com/document_detail/2579691.html).
	//
	// 	- If the resources belong to a service that is not centrally deployed, set the value to the region ID of the resources.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of a resource.
	//
	// Valid values of N: 1 to 50.
	//
	// ARN format: `arn:acs:${ProductCode}:${Region}:${Account}:${ResourceType}/${ResourceId}` Fields:
	//
	// 	- `ProductCode`: the service code. You can set this field to a value obtained from the response of the [ListSupportResourceTypes](https://help.aliyun.com/document_detail/2330915.html) operation.
	//
	// 	- `Region`: the region ID of the resource. You can set this field to an asterisk (\\*) to indicate the current region.
	//
	// 	- `Account`: the ID of the Alibaba Cloud account to which the resource belongs. You can set this field to an asterisk (\\*) to indicate the current Alibaba Cloud account.
	//
	// 	- `ResourceType`: the resource type. You can set this field to a value obtained from the response of the [ListSupportResourceTypes](https://help.aliyun.com/document_detail/2330915.html) operation.
	//
	// 	- `ResourceId`: the ID of the resource.
	//
	// >  You can set `ProductCode` and `ResourceType` in ResourceARN to values defined in Resource Group, ActionTrail, or Resource Center.
	//
	// example:
	//
	// arn:acs:ecs:cn-hangzhou:123456789****:instance/i-bp15hr53jws84akg****
	ResourceARN          []*string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The key-value pairs of tags. You can specify 1 to 10 key-value pairs.
	//
	// If you specify multiple tags, the system queries the resources to which all these tags are added.
	//
	// Limits:
	//
	// 	- A tag key must be 1 to 128 characters in length.
	//
	// 	- A tag value must be 1 to 128 characters in length.
	//
	// 	- Tag keys and tag values are case-sensitive.
	//
	// 	- Each tag key on a resource can have only one tag value. If you create a tag that has the same key as an existing tag, the value of the existing tag is overwritten.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) GetCategory() *string {
	return s.Category
}

func (s *ListTagResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListTagResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListTagResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagResourcesRequest) GetResourceARN() []*string {
	return s.ResourceARN
}

func (s *ListTagResourcesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListTagResourcesRequest) GetTags() *string {
	return s.Tags
}

func (s *ListTagResourcesRequest) SetCategory(v string) *ListTagResourcesRequest {
	s.Category = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetPageSize(v int32) *ListTagResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceARN(v []*string) *ListTagResourcesRequest {
	s.ResourceARN = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetTags(v string) *ListTagResourcesRequest {
	s.Tags = &v
	return s
}

func (s *ListTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
