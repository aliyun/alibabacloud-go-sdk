// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListVodTagResourcesRequest
	GetNextToken() *string
	SetOwnerId(v int64) *ListVodTagResourcesRequest
	GetOwnerId() *int64
	SetResourceId(v []*string) *ListVodTagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *ListVodTagResourcesRequest
	GetResourceType() *string
	SetTag(v []*ListVodTagResourcesRequestTag) *ListVodTagResourcesRequest
	GetTag() []*ListVodTagResourcesRequestTag
	SetTagOwnerBid(v string) *ListVodTagResourcesRequest
	GetTagOwnerBid() *string
	SetTagOwnerUid(v string) *ListVodTagResourcesRequest
	GetTagOwnerUid() *string
}

type ListVodTagResourcesRequest struct {
	NextToken  *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId    *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// This parameter is required.
	ResourceType *string                          `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag          []*ListVodTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	TagOwnerBid  *string                          `json:"TagOwnerBid,omitempty" xml:"TagOwnerBid,omitempty"`
	TagOwnerUid  *string                          `json:"TagOwnerUid,omitempty" xml:"TagOwnerUid,omitempty"`
}

func (s ListVodTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVodTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListVodTagResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVodTagResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListVodTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *ListVodTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListVodTagResourcesRequest) GetTag() []*ListVodTagResourcesRequestTag {
	return s.Tag
}

func (s *ListVodTagResourcesRequest) GetTagOwnerBid() *string {
	return s.TagOwnerBid
}

func (s *ListVodTagResourcesRequest) GetTagOwnerUid() *string {
	return s.TagOwnerUid
}

func (s *ListVodTagResourcesRequest) SetNextToken(v string) *ListVodTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListVodTagResourcesRequest) SetOwnerId(v int64) *ListVodTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListVodTagResourcesRequest) SetResourceId(v []*string) *ListVodTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListVodTagResourcesRequest) SetResourceType(v string) *ListVodTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListVodTagResourcesRequest) SetTag(v []*ListVodTagResourcesRequestTag) *ListVodTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListVodTagResourcesRequest) SetTagOwnerBid(v string) *ListVodTagResourcesRequest {
	s.TagOwnerBid = &v
	return s
}

func (s *ListVodTagResourcesRequest) SetTagOwnerUid(v string) *ListVodTagResourcesRequest {
	s.TagOwnerUid = &v
	return s
}

func (s *ListVodTagResourcesRequest) Validate() error {
	return dara.Validate(s)
}

type ListVodTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVodTagResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListVodTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListVodTagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListVodTagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListVodTagResourcesRequestTag) SetKey(v string) *ListVodTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListVodTagResourcesRequestTag) SetValue(v string) *ListVodTagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *ListVodTagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
