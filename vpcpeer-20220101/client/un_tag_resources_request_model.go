// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UnTagResourcesRequest
	GetAll() *bool
	SetClientToken(v string) *UnTagResourcesRequest
	GetClientToken() *string
	SetRegionId(v string) *UnTagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *UnTagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *UnTagResourcesRequest
	GetResourceType() *string
	SetTagKey(v []*string) *UnTagResourcesRequest
	GetTagKey() []*string
}

type UnTagResourcesRequest struct {
	// Specifies whether to remove all tags from the resource. Valid values:
	//
	// - **true**: Removes all tags from the resource.
	//
	// - **false*	- (default): Does not remove all tags from the resource.
	//
	// example:
	//
	// false
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the token value on your client and ensure that it is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is unique.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the region where the tag resides.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource. You can specify up to 20 resource IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type. Valid value: **PeerConnection**, which indicates a VPC peering connection.
	//
	// This parameter is required.
	//
	// example:
	//
	// PeerConnection
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key that you want to remove. You can specify up to 20 tag keys. If you need to pass this parameter, you can enter an empty string.
	//
	// A tag key can be up to 128 characters in length and cannot start with `aliyun` or `acs:`. It cannot contain `http://` or `https://`.
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UnTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UnTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UnTagResourcesRequest) GetAll() *bool {
	return s.All
}

func (s *UnTagResourcesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UnTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UnTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UnTagResourcesRequest) GetTagKey() []*string {
	return s.TagKey
}

func (s *UnTagResourcesRequest) SetAll(v bool) *UnTagResourcesRequest {
	s.All = &v
	return s
}

func (s *UnTagResourcesRequest) SetClientToken(v string) *UnTagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *UnTagResourcesRequest) SetRegionId(v string) *UnTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UnTagResourcesRequest) SetResourceId(v []*string) *UnTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UnTagResourcesRequest) SetResourceType(v string) *UnTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UnTagResourcesRequest) SetTagKey(v []*string) *UnTagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UnTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
