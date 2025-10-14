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
	SetTotalCount(v int32) *ListTagResourcesResponseBody
	GetTotalCount() *int32
}

type ListTagResourcesResponseBody struct {
	// The pagination token that is used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// AAAAAYwsxxxxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-280B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The tags of the resource.
	TagResources []*ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListTagResourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
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

func (s *ListTagResourcesResponseBody) SetTotalCount(v int32) *ListTagResourcesResponseBody {
	s.TotalCount = &v
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
	// The resource ID.
	//
	// example:
	//
	// example.com
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// site
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key.
	//
	// example:
	//
	// env
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// value
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
