// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyFilter(v string) *ListTagKeysRequest
	GetKeyFilter() *string
	SetMaxResults(v int32) *ListTagKeysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagKeysRequest
	GetNextToken() *string
	SetResourceType(v string) *ListTagKeysRequest
	GetResourceType() *string
}

type ListTagKeysRequest struct {
	// The tag key for a fuzzy query.
	//
	// example:
	//
	// team
	KeyFilter *string `json:"KeyFilter,omitempty" xml:"KeyFilter,omitempty"`
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
}

func (s ListTagKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) GetKeyFilter() *string {
	return s.KeyFilter
}

func (s *ListTagKeysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagKeysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagKeysRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagKeysRequest) SetKeyFilter(v string) *ListTagKeysRequest {
	s.KeyFilter = &v
	return s
}

func (s *ListTagKeysRequest) SetMaxResults(v int32) *ListTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagKeysRequest) Validate() error {
	return dara.Validate(s)
}
