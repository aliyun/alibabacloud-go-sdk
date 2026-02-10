// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListLiveTagResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListLiveTagResourcesResponseBody
	GetRequestId() *string
	SetTagResources(v *ListLiveTagResourcesResponseBodyTagResources) *ListLiveTagResourcesResponseBody
	GetTagResources() *ListLiveTagResourcesResponseBodyTagResources
}

type ListLiveTagResourcesResponseBody struct {
	// example:
	//
	// 6a5e8f4fae643e70d1a2ff1827cd91bd
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId    *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListLiveTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListLiveTagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveTagResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLiveTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLiveTagResourcesResponseBody) GetTagResources() *ListLiveTagResourcesResponseBodyTagResources {
	return s.TagResources
}

func (s *ListLiveTagResourcesResponseBody) SetNextToken(v string) *ListLiveTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListLiveTagResourcesResponseBody) SetRequestId(v string) *ListLiveTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLiveTagResourcesResponseBody) SetTagResources(v *ListLiveTagResourcesResponseBodyTagResources) *ListLiveTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *ListLiveTagResourcesResponseBody) Validate() error {
	if s.TagResources != nil {
		if err := s.TagResources.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListLiveTagResourcesResponseBodyTagResources struct {
	TagResource []*ListLiveTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListLiveTagResourcesResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListLiveTagResourcesResponseBodyTagResources) GetTagResource() []*ListLiveTagResourcesResponseBodyTagResourcesTagResource {
	return s.TagResource
}

func (s *ListLiveTagResourcesResponseBodyTagResources) SetTagResource(v []*ListLiveTagResourcesResponseBodyTagResourcesTagResource) *ListLiveTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

func (s *ListLiveTagResourcesResponseBodyTagResources) Validate() error {
	if s.TagResource != nil {
		for _, item := range s.TagResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListLiveTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListLiveTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return dara.Prettify(s)
}

func (s ListLiveTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListLiveTagResourcesResponseBodyTagResourcesTagResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListLiveTagResourcesResponseBodyTagResourcesTagResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListLiveTagResourcesResponseBodyTagResourcesTagResource) GetTagKey() *string {
	return s.TagKey
}

func (s *ListLiveTagResourcesResponseBodyTagResourcesTagResource) GetTagValue() *string {
	return s.TagValue
}

func (s *ListLiveTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListLiveTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListLiveTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListLiveTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListLiveTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListLiveTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListLiveTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListLiveTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListLiveTagResourcesResponseBodyTagResourcesTagResource) Validate() error {
	return dara.Validate(s)
}
