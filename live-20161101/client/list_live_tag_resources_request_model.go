// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListLiveTagResourcesRequest
	GetNextToken() *string
	SetOwnerId(v int64) *ListLiveTagResourcesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ListLiveTagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *ListLiveTagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *ListLiveTagResourcesRequest
	GetResourceType() *string
	SetTag(v []*ListLiveTagResourcesRequestTag) *ListLiveTagResourcesRequest
	GetTag() []*ListLiveTagResourcesRequestTag
	SetTagOwnerBid(v string) *ListLiveTagResourcesRequest
	GetTagOwnerBid() *string
	SetTagOwnerUid(v string) *ListLiveTagResourcesRequest
	GetTagOwnerUid() *string
}

type ListLiveTagResourcesRequest struct {
	// The token for the next query.
	//
	// example:
	//
	// q2j8bLtBdhONLRkgaPBa6A==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of resource IDs. ResourceId and Tag cannot both be empty.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// DOMAIN
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The list of tags. ResourceId and Tag cannot both be empty.
	Tag []*ListLiveTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The Bid of the tag owner.
	//
	// example:
	//
	// 26842
	TagOwnerBid *string `json:"TagOwnerBid,omitempty" xml:"TagOwnerBid,omitempty"`
	// The Alibaba Cloud account ID to which the tag belongs.
	//
	// example:
	//
	// xxx1234xxx
	TagOwnerUid *string `json:"TagOwnerUid,omitempty" xml:"TagOwnerUid,omitempty"`
}

func (s ListLiveTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListLiveTagResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLiveTagResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListLiveTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLiveTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *ListLiveTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListLiveTagResourcesRequest) GetTag() []*ListLiveTagResourcesRequestTag {
	return s.Tag
}

func (s *ListLiveTagResourcesRequest) GetTagOwnerBid() *string {
	return s.TagOwnerBid
}

func (s *ListLiveTagResourcesRequest) GetTagOwnerUid() *string {
	return s.TagOwnerUid
}

func (s *ListLiveTagResourcesRequest) SetNextToken(v string) *ListLiveTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListLiveTagResourcesRequest) SetOwnerId(v int64) *ListLiveTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListLiveTagResourcesRequest) SetRegionId(v string) *ListLiveTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListLiveTagResourcesRequest) SetResourceId(v []*string) *ListLiveTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListLiveTagResourcesRequest) SetResourceType(v string) *ListLiveTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListLiveTagResourcesRequest) SetTag(v []*ListLiveTagResourcesRequestTag) *ListLiveTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListLiveTagResourcesRequest) SetTagOwnerBid(v string) *ListLiveTagResourcesRequest {
	s.TagOwnerBid = &v
	return s
}

func (s *ListLiveTagResourcesRequest) SetTagOwnerUid(v string) *ListLiveTagResourcesRequest {
	s.TagOwnerUid = &v
	return s
}

func (s *ListLiveTagResourcesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLiveTagResourcesRequestTag struct {
	// The tag key of the resource.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value that corresponds to the tag key.
	//
	// example:
	//
	// dev
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListLiveTagResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListLiveTagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListLiveTagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListLiveTagResourcesRequestTag) SetKey(v string) *ListLiveTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListLiveTagResourcesRequestTag) SetValue(v string) *ListLiveTagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *ListLiveTagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
