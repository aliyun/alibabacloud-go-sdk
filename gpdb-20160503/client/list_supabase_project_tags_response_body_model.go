// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupabaseProjectTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListSupabaseProjectTagsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSupabaseProjectTagsResponseBody
	GetRequestId() *string
	SetTagResources(v []*ListSupabaseProjectTagsResponseBodyTagResources) *ListSupabaseProjectTagsResponseBody
	GetTagResources() []*ListSupabaseProjectTagsResponseBodyTagResources
}

type ListSupabaseProjectTagsResponseBody struct {
	// The pagination token for the next page of results. This parameter is not returned if no more results are available.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34b32a0a-08ef-4a87-b6be-cdd9f56fc3ad
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of tagged resources.
	TagResources []*ListSupabaseProjectTagsResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s ListSupabaseProjectTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseProjectTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSupabaseProjectTagsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupabaseProjectTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSupabaseProjectTagsResponseBody) GetTagResources() []*ListSupabaseProjectTagsResponseBodyTagResources {
	return s.TagResources
}

func (s *ListSupabaseProjectTagsResponseBody) SetNextToken(v string) *ListSupabaseProjectTagsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSupabaseProjectTagsResponseBody) SetRequestId(v string) *ListSupabaseProjectTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSupabaseProjectTagsResponseBody) SetTagResources(v []*ListSupabaseProjectTagsResponseBodyTagResources) *ListSupabaseProjectTagsResponseBody {
	s.TagResources = v
	return s
}

func (s *ListSupabaseProjectTagsResponseBody) Validate() error {
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

type ListSupabaseProjectTagsResponseBodyTagResources struct {
	// The instance ID.
	//
	// example:
	//
	// spb-xxxxxxxxxx
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
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

func (s ListSupabaseProjectTagsResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s ListSupabaseProjectTagsResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListSupabaseProjectTagsResponseBodyTagResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListSupabaseProjectTagsResponseBodyTagResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListSupabaseProjectTagsResponseBodyTagResources) GetTagKey() *string {
	return s.TagKey
}

func (s *ListSupabaseProjectTagsResponseBodyTagResources) GetTagValue() *string {
	return s.TagValue
}

func (s *ListSupabaseProjectTagsResponseBodyTagResources) SetResourceId(v string) *ListSupabaseProjectTagsResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *ListSupabaseProjectTagsResponseBodyTagResources) SetResourceType(v string) *ListSupabaseProjectTagsResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *ListSupabaseProjectTagsResponseBodyTagResources) SetTagKey(v string) *ListSupabaseProjectTagsResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *ListSupabaseProjectTagsResponseBodyTagResources) SetTagValue(v string) *ListSupabaseProjectTagsResponseBodyTagResources {
	s.TagValue = &v
	return s
}

func (s *ListSupabaseProjectTagsResponseBodyTagResources) Validate() error {
	return dara.Validate(s)
}
