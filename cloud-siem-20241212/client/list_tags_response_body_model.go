// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTagsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagsResponseBody
	GetRequestId() *string
	SetTags(v []*ListTagsResponseBodyTags) *ListTagsResponseBody
	GetTags() []*ListTagsResponseBodyTags
}

type ListTagsResponseBody struct {
	// The maximum number of results to return when using NextToken-based pagination. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query. Leave this parameter empty for the first query or if no more results exist. If a next query is available, set this parameter to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The entity tags. The value is a JSON array string in the following format:
	//
	// `"[{"tagKey1":"tagValue1"},{"tagKey2":"tagValue2"}]"`
	Tags []*ListTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagsResponseBody) GetTags() []*ListTagsResponseBodyTags {
	return s.Tags
}

func (s *ListTagsResponseBody) SetMaxResults(v int32) *ListTagsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagsResponseBody) SetNextToken(v string) *ListTagsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagsResponseBody) SetRequestId(v string) *ListTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagsResponseBody) SetTags(v []*ListTagsResponseBodyTags) *ListTagsResponseBody {
	s.Tags = v
	return s
}

func (s *ListTagsResponseBody) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagsResponseBodyTags struct {
	// The number of quick queries returned on the current page.
	//
	// example:
	//
	// 8
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The key of the tag to query. You can specify multiple keys. N is a positive integer.
	//
	// example:
	//
	// False
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s ListTagsResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s ListTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTags) GetCount() *int64 {
	return s.Count
}

func (s *ListTagsResponseBodyTags) GetTag() *string {
	return s.Tag
}

func (s *ListTagsResponseBodyTags) SetCount(v int64) *ListTagsResponseBodyTags {
	s.Count = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetTag(v string) *ListTagsResponseBodyTags {
	s.Tag = &v
	return s
}

func (s *ListTagsResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
