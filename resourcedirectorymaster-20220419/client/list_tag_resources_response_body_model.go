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
	SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody
	GetTagResources() []*ListTagResourcesResponseBodyTagResources
}

type ListTagResourcesResponseBody struct {
	// Indicates whether the next query is required.```` Valid values:
	//
	// 	- If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the `next query` is not required.
	//
	// 	- If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.````
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8054B059-6B36-53BF-AA45-B8C9A0ED05AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesResponseBody) GetTagResources() []*ListTagResourcesResponseBodyTagResources {
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

func (s *ListTagResourcesResponseBody) SetTagResources(v []*ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBody) Validate() error {
	if s.TagResources != nil {
		for _, item := range s.TagResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagResourcesResponseBodyTagResources struct {
	// The Alibaba Cloud account ID of the member.
	//
	// example:
	//
	// 1098***
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the object whose tags are queried. Valid values:
	//
	// 	- Account: member
	//
	// example:
	//
	// Account
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// k1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// k1
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagResourcesResponseBodyTagResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesResponseBodyTagResources) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagResourcesResponseBodyTagResources) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceId(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetResourceType(v string) *ListTagResourcesResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagKey(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagValue(v string) *ListTagResourcesResponseBodyTagResources {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) Validate() error {
	return dara.Validate(s)
}
