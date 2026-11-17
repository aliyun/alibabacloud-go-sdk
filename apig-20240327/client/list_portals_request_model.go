// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPortalsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPortalsRequest
	GetMaxResults() *int32
	SetName(v string) *ListPortalsRequest
	GetName() *string
	SetNextToken(v string) *ListPortalsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListPortalsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPortalsRequest
	GetPageSize() *int32
}

type ListPortalsRequest struct {
	// The maximum number of entries to return per request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The portal name, which is the Name property of the portal resource. This corresponds to the name field in the response list items.
	//
	// example:
	//
	// developer-portal
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// next-token-xxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListPortalsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPortalsRequest) GoString() string {
	return s.String()
}

func (s *ListPortalsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPortalsRequest) GetName() *string {
	return s.Name
}

func (s *ListPortalsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPortalsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPortalsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPortalsRequest) SetMaxResults(v int32) *ListPortalsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPortalsRequest) SetName(v string) *ListPortalsRequest {
	s.Name = &v
	return s
}

func (s *ListPortalsRequest) SetNextToken(v string) *ListPortalsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPortalsRequest) SetPageNumber(v int32) *ListPortalsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPortalsRequest) SetPageSize(v int32) *ListPortalsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPortalsRequest) Validate() error {
	return dara.Validate(s)
}
