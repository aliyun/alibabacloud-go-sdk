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
	SetResult(v []*ListTagResourcesResponseBodyResult) *ListTagResourcesResponseBody
	GetResult() []*ListTagResourcesResponseBodyResult
}

type ListTagResourcesResponseBody struct {
	// The token that is used to retrieve the next page.
	//
	// example:
	//
	// 20
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The resources.
	Result []*ListTagResourcesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
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

func (s *ListTagResourcesResponseBody) GetResult() []*ListTagResourcesResponseBodyResult {
	return s.Result
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetResult(v []*ListTagResourcesResponseBodyResult) *ListTagResourcesResponseBody {
	s.Result = v
	return s
}

func (s *ListTagResourcesResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagResourcesResponseBodyResult struct {
	// The ID of the resource.
	//
	// example:
	//
	// 54041
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// hostGroup
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// GENIE_FUNCTION
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// ALLOW
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyResult) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagResourcesResponseBodyResult) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesResponseBodyResult) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagResourcesResponseBodyResult) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTagResourcesResponseBodyResult) SetResourceId(v string) *ListTagResourcesResponseBodyResult {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyResult) SetResourceType(v string) *ListTagResourcesResponseBodyResult {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyResult) SetTagKey(v string) *ListTagResourcesResponseBodyResult {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyResult) SetTagValue(v string) *ListTagResourcesResponseBodyResult {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
