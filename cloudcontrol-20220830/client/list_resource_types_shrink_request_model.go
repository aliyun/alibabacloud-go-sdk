// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceTypesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceTypesShrinkRequest
	GetNextToken() *string
	SetResourceTypesShrink(v string) *ListResourceTypesShrinkRequest
	GetResourceTypesShrink() *string
}

type ListResourceTypesShrinkRequest struct {
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// ECS::Disk
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The information about the resource types.
	ResourceTypesShrink *string `json:"resourceTypes,omitempty" xml:"resourceTypes,omitempty"`
}

func (s ListResourceTypesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceTypesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceTypesShrinkRequest) GetResourceTypesShrink() *string {
	return s.ResourceTypesShrink
}

func (s *ListResourceTypesShrinkRequest) SetMaxResults(v int32) *ListResourceTypesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetNextToken(v string) *ListResourceTypesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) SetResourceTypesShrink(v string) *ListResourceTypesShrinkRequest {
	s.ResourceTypesShrink = &v
	return s
}

func (s *ListResourceTypesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
