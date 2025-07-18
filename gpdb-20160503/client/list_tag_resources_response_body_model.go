// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagResourcesResponseBody
	GetRequestId() *string
	SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody
	GetTagResources() *ListTagResourcesResponseBodyTagResources
}

type ListTagResourcesResponseBody struct {
	// The token used to perform the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5414A4E5-4C36-4461-95FC-23757A20B5F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about the instances and tags, including the instance IDs, instance modes, and tag key-value pairs.
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagResourcesResponseBody) GetTagResources() *ListTagResourcesResponseBodyTagResources {
	return s.TagResources
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) GetTagResource() []*ListTagResourcesResponseBodyTagResourcesTagResource {
	return s.TagResource
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) Validate() error {
	return dara.Validate(s)
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	// The ID of the instance.
	//
	// example:
	//
	// gp-xxxxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The storage mode of the instance.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) Validate() error {
	return dara.Validate(s)
}
