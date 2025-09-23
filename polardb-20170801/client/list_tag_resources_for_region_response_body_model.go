// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesForRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagResourcesForRegionResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagResourcesForRegionResponseBody
	GetRequestId() *string
	SetTagResources(v *ListTagResourcesForRegionResponseBodyTagResources) *ListTagResourcesForRegionResponseBody
	GetTagResources() *ListTagResourcesForRegionResponseBodyTagResources
}

type ListTagResourcesForRegionResponseBody struct {
	// example:
	//
	// 212db86sca4384811e0b5e8707e******
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 688C04E4-23F8-409F-8A38-B954D5******
	RequestId    *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagResources *ListTagResourcesForRegionResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesForRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesForRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesForRegionResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesForRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagResourcesForRegionResponseBody) GetTagResources() *ListTagResourcesForRegionResponseBodyTagResources {
	return s.TagResources
}

func (s *ListTagResourcesForRegionResponseBody) SetNextToken(v string) *ListTagResourcesForRegionResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesForRegionResponseBody) SetRequestId(v string) *ListTagResourcesForRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesForRegionResponseBody) SetTagResources(v *ListTagResourcesForRegionResponseBodyTagResources) *ListTagResourcesForRegionResponseBody {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesForRegionResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTagResourcesForRegionResponseBodyTagResources struct {
	TagResource []*ListTagResourcesForRegionResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesForRegionResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesForRegionResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesForRegionResponseBodyTagResources) GetTagResource() []*ListTagResourcesForRegionResponseBodyTagResourcesTagResource {
	return s.TagResource
}

func (s *ListTagResourcesForRegionResponseBodyTagResources) SetTagResource(v []*ListTagResourcesForRegionResponseBodyTagResourcesTagResource) *ListTagResourcesForRegionResponseBodyTagResources {
	s.TagResource = v
	return s
}

func (s *ListTagResourcesForRegionResponseBodyTagResources) Validate() error {
	return dara.Validate(s)
}

type ListTagResourcesForRegionResponseBodyTagResourcesTagResource struct {
	// example:
	//
	// pc-****************
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// type
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesForRegionResponseBodyTagResourcesTagResource) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesForRegionResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesForRegionResponseBodyTagResourcesTagResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagResourcesForRegionResponseBodyTagResourcesTagResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesForRegionResponseBodyTagResourcesTagResource) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagResourcesForRegionResponseBodyTagResourcesTagResource) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTagResourcesForRegionResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesForRegionResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesForRegionResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesForRegionResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesForRegionResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesForRegionResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesForRegionResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesForRegionResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesForRegionResponseBodyTagResourcesTagResource) Validate() error {
	return dara.Validate(s)
}
