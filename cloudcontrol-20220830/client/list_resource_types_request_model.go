// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListResourceTypesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListResourceTypesRequest
	GetNextToken() *string
	SetResourceTypes(v []*string) *ListResourceTypesRequest
	GetResourceTypes() []*string
}

type ListResourceTypesRequest struct {
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
	ResourceTypes []*string `json:"resourceTypes,omitempty" xml:"resourceTypes,omitempty" type:"Repeated"`
}

func (s ListResourceTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceTypesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListResourceTypesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListResourceTypesRequest) GetResourceTypes() []*string {
	return s.ResourceTypes
}

func (s *ListResourceTypesRequest) SetMaxResults(v int32) *ListResourceTypesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListResourceTypesRequest) SetNextToken(v string) *ListResourceTypesRequest {
	s.NextToken = &v
	return s
}

func (s *ListResourceTypesRequest) SetResourceTypes(v []*string) *ListResourceTypesRequest {
	s.ResourceTypes = v
	return s
}

func (s *ListResourceTypesRequest) Validate() error {
	return dara.Validate(s)
}
