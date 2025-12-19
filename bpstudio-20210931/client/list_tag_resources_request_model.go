// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ListTagResourcesRequest
	GetClientToken() *string
	SetNextToken(v string) *ListTagResourcesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListTagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *ListTagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *ListTagResourcesRequest
	GetResourceType() *string
	SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest
	GetTag() []*ListTagResourcesRequestTag
}

type ListTagResourcesRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 1600765710019
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource IDs. You can specify a maximum number of 50 IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// Application
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags. A maximum of 20 tags are supported.
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ListTagResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *ListTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesRequest) GetTag() []*ListTagResourcesRequestTag {
	return s.Tag
}

func (s *ListTagResourcesRequest) SetClientToken(v string) *ListTagResourcesRequest {
	s.ClientToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListTagResourcesRequest) Validate() error {
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

type ListTagResourcesRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// Label 1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// Label value 1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *ListTagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}
