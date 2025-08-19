// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagValuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTagValuesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagValuesRequest
	GetNextToken() *string
	SetResourceType(v string) *ListTagValuesRequest
	GetResourceType() *string
	SetTagKey(v string) *ListTagValuesRequest
	GetTagKey() *string
	SetValueFilter(v string) *ListTagValuesRequest
	GetValueFilter() *string
}

type ListTagValuesRequest struct {
	// The maximum number of entries to return for a single request.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resource type.
	//
	// The value Account indicates the members of the resource directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// Account
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tag key. This parameter specifies a filter condition for the query.
	//
	// This parameter is required.
	//
	// example:
	//
	// k1
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value for a fuzzy query.
	//
	// example:
	//
	// v1
	ValueFilter *string `json:"ValueFilter,omitempty" xml:"ValueFilter,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagValuesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagValuesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagValuesRequest) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagValuesRequest) GetValueFilter() *string {
	return s.ValueFilter
}

func (s *ListTagValuesRequest) SetMaxResults(v int32) *ListTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) SetTagKey(v string) *ListTagValuesRequest {
	s.TagKey = &v
	return s
}

func (s *ListTagValuesRequest) SetValueFilter(v string) *ListTagValuesRequest {
	s.ValueFilter = &v
	return s
}

func (s *ListTagValuesRequest) Validate() error {
	return dara.Validate(s)
}
