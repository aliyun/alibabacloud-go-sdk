// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVodTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListVodTagResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVodTagResourcesResponseBody
	GetRequestId() *string
	SetTagResources(v *ListVodTagResourcesResponseBodyTagResources) *ListVodTagResourcesResponseBody
	GetTagResources() *ListVodTagResourcesResponseBodyTagResources
}

type ListVodTagResourcesResponseBody struct {
	NextToken    *string                                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListVodTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListVodTagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVodTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListVodTagResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVodTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVodTagResourcesResponseBody) GetTagResources() *ListVodTagResourcesResponseBodyTagResources {
	return s.TagResources
}

func (s *ListVodTagResourcesResponseBody) SetNextToken(v string) *ListVodTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVodTagResourcesResponseBody) SetRequestId(v string) *ListVodTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVodTagResourcesResponseBody) SetTagResources(v *ListVodTagResourcesResponseBodyTagResources) *ListVodTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *ListVodTagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVodTagResourcesResponseBodyTagResources struct {
	TagResource []*ListVodTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListVodTagResourcesResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s ListVodTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListVodTagResourcesResponseBodyTagResources) GetTagResource() []*ListVodTagResourcesResponseBodyTagResourcesTagResource {
	return s.TagResource
}

func (s *ListVodTagResourcesResponseBodyTagResources) SetTagResource(v []*ListVodTagResourcesResponseBodyTagResourcesTagResource) *ListVodTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

func (s *ListVodTagResourcesResponseBodyTagResources) Validate() error {
	return dara.Validate(s)
}

type ListVodTagResourcesResponseBodyTagResourcesTagResource struct {
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListVodTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return dara.Prettify(s)
}

func (s ListVodTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListVodTagResourcesResponseBodyTagResourcesTagResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListVodTagResourcesResponseBodyTagResourcesTagResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListVodTagResourcesResponseBodyTagResourcesTagResource) GetTagKey() *string {
	return s.TagKey
}

func (s *ListVodTagResourcesResponseBodyTagResourcesTagResource) GetTagValue() *string {
	return s.TagValue
}

func (s *ListVodTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListVodTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListVodTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListVodTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListVodTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListVodTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListVodTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListVodTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListVodTagResourcesResponseBodyTagResourcesTagResource) Validate() error {
	return dara.Validate(s)
}
