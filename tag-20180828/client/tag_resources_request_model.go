// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *TagResourcesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *TagResourcesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *TagResourcesRequest
	GetRegionId() *string
	SetResourceARN(v []*string) *TagResourcesRequest
	GetResourceARN() []*string
	SetResourceOwnerAccount(v string) *TagResourcesRequest
	GetResourceOwnerAccount() *string
	SetTags(v string) *TagResourcesRequest
	GetTags() *string
}

type TagResourcesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// 	- If the resources belong to a service that is centrally deployed, set the value to `cn-hangzhou` or to the region ID of the resources by referring to [Regions supported by tag-related operations on resources of centrally deployed Alibaba Cloud services](https://help.aliyun.com/document_detail/2579691.html).
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
	// This parameter is required.
	//
	// example:
	//
	// arn:acs:vpc:cn-hangzhou:123456789****:vpc/vpc-bp19dd90tkt6tz7wu****
	ResourceARN          []*string `json:"ResourceARN,omitempty" xml:"ResourceARN,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The key-value pairs of tags. You can specify 1 to 10 key-value pairs.
	//
	// If you specify multiple tags, the system adds all the tags to the specified resources.
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
	// This parameter is required.
	//
	// example:
	//
	// {"k1":"v1","k2":"v2"}
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s TagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *TagResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *TagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TagResourcesRequest) GetResourceARN() []*string {
	return s.ResourceARN
}

func (s *TagResourcesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *TagResourcesRequest) GetTags() *string {
	return s.Tags
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceARN(v []*string) *TagResourcesRequest {
	s.ResourceARN = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetTags(v string) *TagResourcesRequest {
	s.Tags = &v
	return s
}

func (s *TagResourcesRequest) Validate() error {
	return dara.Validate(s)
}
